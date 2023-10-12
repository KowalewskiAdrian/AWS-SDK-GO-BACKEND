// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package config

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
	"github.com/hashicorp/aws-sdk-go-base/v2/internal/expand"
	"github.com/hashicorp/aws-sdk-go-base/v2/logging"
	"golang.org/x/net/http/httpproxy"
)

type Config struct {
	AccessKey                      string
	AllowedAccountIds              []string
	APNInfo                        *APNInfo
	AssumeRole                     *AssumeRole
	AssumeRoleWithWebIdentity      *AssumeRoleWithWebIdentity
	CallerDocumentationURL         string
	CallerName                     string
	CustomCABundle                 string
	EC2MetadataServiceEnableState  imds.ClientEnableState
	EC2MetadataServiceEndpoint     string
	EC2MetadataServiceEndpointMode string
	ForbiddenAccountIds            []string
	HTTPClient                     *http.Client
	HTTPProxy                      string
	IamEndpoint                    string
	Insecure                       bool
	Logger                         logging.Logger
	MaxRetries                     int
	Profile                        string
	Region                         string
	RetryMode                      aws.RetryMode
	SecretKey                      string
	SharedCredentialsFiles         []string
	SharedConfigFiles              []string
	SkipCredsValidation            bool
	SkipRequestingAccountId        bool
	StsEndpoint                    string
	StsRegion                      string
	SuppressDebugLog               bool
	Token                          string
	UseDualStackEndpoint           bool
	UseFIPSEndpoint                bool
	UseLegacyWorkflow              bool
	UserAgent                      UserAgentProducts
}

type AssumeRole struct {
	RoleARN           string
	Duration          time.Duration
	ExternalID        string
	Policy            string
	PolicyARNs        []string
	SessionName       string
	SourceIdentity    string
	Tags              map[string]string
	TransitiveTagKeys []string
}

func (c Config) CustomCABundleReader() (*bytes.Reader, error) {
	if c.CustomCABundle == "" {
		return nil, nil
	}
	bundleFile, err := expand.FilePath(c.CustomCABundle)
	if err != nil {
		return nil, fmt.Errorf("expanding custom CA bundle: %w", err)
	}
	bundle, err := os.ReadFile(bundleFile)
	if err != nil {
		return nil, fmt.Errorf("reading custom CA bundle: %w", err)
	}
	return bytes.NewReader(bundle), nil
}

// HTTPTransportOptions returns functional options that configures an http.Transport.
// The returned options function is called on both AWS SDKv1 and v2 default HTTP clients.
func (c Config) HTTPTransportOptions() (func(*http.Transport), error) {
	var err error
	var proxyUrl *url.URL
	if c.HTTPProxy != "" {
		proxyUrl, err = url.Parse(c.HTTPProxy)
		if err != nil {
			return nil, fmt.Errorf("error parsing HTTP proxy URL: %w", err)
		}
	}

	opts := func(tr *http.Transport) {
		tr.MaxIdleConnsPerHost = awshttp.DefaultHTTPTransportMaxIdleConnsPerHost

		tlsConfig := tr.TLSClientConfig
		if tlsConfig == nil {
			tlsConfig = &tls.Config{
				MinVersion: tls.VersionTLS12,
			}
			tr.TLSClientConfig = tlsConfig
		}

		if c.Insecure {
			tr.TLSClientConfig.InsecureSkipVerify = true
		}

		if proxyUrl != nil {
			proxyConfig := httpproxy.FromEnvironment()
			proxyConfig.HTTPProxy = proxyUrl.String()
			proxyConfig.HTTPSProxy = proxyUrl.String()
			tr.Proxy = func(req *http.Request) (*url.URL, error) {
				return proxyConfig.ProxyFunc()(req.URL)
			}
		} else {
			proxyConfig := httpproxy.FromEnvironment()
			tr.Proxy = func(req *http.Request) (*url.URL, error) {
				return proxyConfig.ProxyFunc()(req.URL)
			}
		}
	}

	return opts, nil
}

func (c Config) ResolveSharedConfigFiles() ([]string, error) {
	v, err := expand.FilePaths(c.SharedConfigFiles)
	if err != nil {
		return []string{}, fmt.Errorf("expanding shared config files: %w", err)
	}
	return v, nil
}

func (c Config) ResolveSharedCredentialsFiles() ([]string, error) {
	v, err := expand.FilePaths(c.SharedCredentialsFiles)
	if err != nil {
		return []string{}, fmt.Errorf("expanding shared credentials files: %w", err)
	}
	return v, nil
}

// VerifyAccountIDAllowed verifies an account ID is not explicitly forbidden
// or omitted from an allow list, if configured.
//
// If the AllowedAccountIds and ForbiddenAccountIds fields are both empty, this
// function will return nil.
func (c Config) VerifyAccountIDAllowed(accountID string) error {
	if len(c.ForbiddenAccountIds) > 0 {
		for _, forbiddenAccountID := range c.ForbiddenAccountIds {
			if accountID == forbiddenAccountID {
				return fmt.Errorf("AWS account ID not allowed: %s", accountID)
			}
		}
	}
	if len(c.AllowedAccountIds) > 0 {
		found := false
		for _, allowedAccountID := range c.AllowedAccountIds {
			if accountID == allowedAccountID {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("AWS account ID not allowed: %s", accountID)
		}
	}
	return nil
}

type AssumeRoleWithWebIdentity struct {
	RoleARN              string
	Duration             time.Duration
	Policy               string
	PolicyARNs           []string
	SessionName          string
	WebIdentityToken     string
	WebIdentityTokenFile string
}

func (c AssumeRoleWithWebIdentity) resolveWebIdentityTokenFile() (string, error) {
	v, err := expand.FilePath(c.WebIdentityTokenFile)
	if err != nil {
		return "", fmt.Errorf("expanding web identity token file: %w", err)
	}
	return v, nil
}

func (c AssumeRoleWithWebIdentity) HasValidTokenSource() bool {
	return c.WebIdentityToken != "" || c.WebIdentityTokenFile != ""
}

// Implements `stscreds.IdentityTokenRetriever`
func (c AssumeRoleWithWebIdentity) GetIdentityToken() ([]byte, error) {
	if c.WebIdentityToken != "" {
		return []byte(c.WebIdentityToken), nil
	}
	webIdentityTokenFile, err := c.resolveWebIdentityTokenFile()
	if err != nil {
		return nil, err
	}

	b, err := os.ReadFile(webIdentityTokenFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read file at %s: %w", webIdentityTokenFile, err)
	}

	return b, nil
}
