// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfawserr

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sts/types"
	smithy "github.com/aws/smithy-go"
)

func TestErrCodeEquals(t *testing.T) {
	testCases := map[string]struct {
		Err      error
		Codes    []string
		Expected bool
	}{
		"nil error": {
			Err:      nil,
			Expected: false,
		},
		"other error": {
			Err:      fmt.Errorf("other error"),
			Expected: false,
		},
		"Top-level smithy.GenericAPIError matching first code": {
			Err:      &smithy.GenericAPIError{Code: "TestCode", Message: "TestMessage"},
			Codes:    []string{"TestCode"},
			Expected: true,
		},
		"Top-level smithy.GenericAPIError matching last code": {
			Err:      &smithy.GenericAPIError{Code: "TestCode", Message: "TestMessage"},
			Codes:    []string{"NotMatching", "TestCode"},
			Expected: true,
		},
		"Top-level smithy.GenericAPIError no code": {
			Err: &smithy.GenericAPIError{Code: "TestCode", Message: "TestMessage"},
		},
		"Top-level smithy.GenericAPIError non-matching codes": {
			Err:   &smithy.GenericAPIError{Code: "TestCode", Message: "TestMessage"},
			Codes: []string{"NotMatching", "AlsoNotMatching"},
		},
		"Wrapped smithy.GenericAPIError matching first code": {
			Err:      fmt.Errorf("test: %w", &smithy.GenericAPIError{Code: "TestCode", Message: "TestMessage"}),
			Codes:    []string{"TestCode"},
			Expected: true,
		},
		"Wrapped smithy.GenericAPIError matching last code": {
			Err:      fmt.Errorf("test: %w", &smithy.GenericAPIError{Code: "TestCode", Message: "TestMessage"}),
			Codes:    []string{"NotMatching", "TestCode"},
			Expected: true,
		},
		"Wrapped smithy.GenericAPIError non-matching codes": {
			Err:   fmt.Errorf("test: %w", &smithy.GenericAPIError{Code: "TestCode", Message: "TestMessage"}),
			Codes: []string{"NotMatching", "AlsoNotMatching"},
		},
		"Top-level sts ExpiredTokenException matching first code": {
			Err:      &types.ExpiredTokenException{ErrorCodeOverride: aws.String("TestCode"), Message: aws.String("TestMessage")},
			Codes:    []string{"TestCode"},
			Expected: true,
		},
		"Top-level sts ExpiredTokenException matching last code": {
			Err:      &types.ExpiredTokenException{ErrorCodeOverride: aws.String("TestCode"), Message: aws.String("TestMessage")},
			Codes:    []string{"NotMatching", "TestCode"},
			Expected: true,
		},
		"Wrapped sts ExpiredTokenException matching first code": {
			Err:      fmt.Errorf("test: %w", &types.ExpiredTokenException{ErrorCodeOverride: aws.String("TestCode"), Message: aws.String("TestMessage")}),
			Codes:    []string{"TestCode"},
			Expected: true,
		},
		"Wrapped sts ExpiredTokenException matching last code": {
			Err:      fmt.Errorf("test: %w", &types.ExpiredTokenException{ErrorCodeOverride: aws.String("TestCode"), Message: aws.String("TestMessage")}),
			Codes:    []string{"NotMatching", "TestCode"},
			Expected: true,
		},
	}

	for name, testCase := range testCases {
		testCase := testCase

		t.Run(name, func(t *testing.T) {
			got := ErrCodeEquals(testCase.Err, testCase.Codes...)

			if got != testCase.Expected {
				t.Errorf("got %t, expected %t", got, testCase.Expected)
			}
		})
	}
}

func TestErrMessageContains(t *testing.T) {
	testCases := map[string]struct {
		Err      error
		Code     string
		Message  string
		Expected bool
	}{
		"nil error": {
			Err:      nil,
			Expected: false,
		},
		"nil error code": {
			Err:      nil,
			Code:     "test",
			Expected: false,
		},
		"nil error message": {
			Err:     nil,
			Message: "test",
		},
		"nil error code and message": {
			Err:     nil,
			Code:    "test",
			Message: "test",
		},
		"other error": {
			Err:      fmt.Errorf("other error"),
			Expected: false,
		},
		"other error code and message": {
			Err:     fmt.Errorf("other error"),
			Code:    "test",
			Message: "test",
		},
		"Top-level smithy.GenericAPIError no code": {
			Err: &smithy.GenericAPIError{Code: "TestCode", Message: "TestMessage"},
		},
		"Top-level smithy.GenericAPIError matching code and no message": {
			Err:      &smithy.GenericAPIError{Code: "TestCode", Message: "TestMessage"},
			Code:     "TestCode",
			Expected: true,
		},
		"Top-level smithy.GenericAPIError matching code and matching message exact": {
			Err:      &smithy.GenericAPIError{Code: "TestCode", Message: "TestMessage"},
			Code:     "TestCode",
			Message:  "TestMessage",
			Expected: true,
		},
		"Top-level smithy.GenericAPIError non-matching code and matching message exact": {
			Err:     &smithy.GenericAPIError{Code: "TestCode", Message: "TestMessage"},
			Code:    "NotMatching",
			Message: "TestMessage",
		},
		"Top-level smithy.GenericAPIError matching code and matching message contains": {
			Err:      &smithy.GenericAPIError{Code: "TestCode", Message: "TestMessage"},
			Code:     "TestCode",
			Message:  "estMess",
			Expected: true,
		},
		"Top-level smithy.GenericAPIError matching code and non-matching message": {
			Err:     &smithy.GenericAPIError{Code: "TestCode", Message: "TestMessage"},
			Code:    "TestCode",
			Message: "NotMatching",
		},
		"Wrapped smithy.GenericAPIError matching code and no message": {
			Err:      fmt.Errorf("test: %w", &smithy.GenericAPIError{Code: "TestCode", Message: "TestMessage"}),
			Code:     "TestCode",
			Expected: true,
		},
		"Wrapped smithy.GenericAPIError matching code and matching message exact": {
			Err:      fmt.Errorf("test: %w", &smithy.GenericAPIError{Code: "TestCode", Message: "TestMessage"}),
			Code:     "TestCode",
			Message:  "TestMessage",
			Expected: true,
		},
		"Wrapped smithy.GenericAPIError non-matching code and matching message exact": {
			Err:     fmt.Errorf("test: %w", &smithy.GenericAPIError{Code: "TestCode", Message: "TestMessage"}),
			Code:    "NotMatching",
			Message: "TestMessage",
		},
		"Wrapped smithy.GenericAPIError matching code and matching message contains": {
			Err:      fmt.Errorf("test: %w", &smithy.GenericAPIError{Code: "TestCode", Message: "TestMessage"}),
			Code:     "TestCode",
			Message:  "estMess",
			Expected: true,
		},
		"Wrapped smithy.GenericAPIError matching code and non-matching message": {
			Err:     fmt.Errorf("test: %w", &smithy.GenericAPIError{Code: "TestCode", Message: "TestMessage"}),
			Code:    "TestCode",
			Message: "NotMatching",
		},
		"Top-level sts ExpiredTokenException matching code and no message": {
			Err:      &types.ExpiredTokenException{ErrorCodeOverride: aws.String("TestCode"), Message: aws.String("TestMessage")},
			Code:     "TestCode",
			Expected: true,
		},
		"Top-level sts ExpiredTokenException matching code and matching message exact": {
			Err:      &types.ExpiredTokenException{ErrorCodeOverride: aws.String("TestCode"), Message: aws.String("TestMessage")},
			Code:     "TestCode",
			Message:  "TestMessage",
			Expected: true,
		},
		"Top-level sts ExpiredTokenException non-matching code and matching message exact": {
			Err:     &types.ExpiredTokenException{ErrorCodeOverride: aws.String("TestCode"), Message: aws.String("TestMessage")},
			Code:    "NotMatching",
			Message: "TestMessage",
		},
		"Top-level sts ExpiredTokenException matching code and matching message contains": {
			Err:      &types.ExpiredTokenException{ErrorCodeOverride: aws.String("TestCode"), Message: aws.String("TestMessage")},
			Code:     "TestCode",
			Message:  "estMess",
			Expected: true,
		},
		"Top-level sts ExpiredTokenException matching code and non-matching message": {
			Err:     &types.ExpiredTokenException{ErrorCodeOverride: aws.String("TestCode"), Message: aws.String("TestMessage")},
			Code:    "TestCode",
			Message: "NotMatching",
		},
	}

	for name, testCase := range testCases {
		testCase := testCase

		t.Run(name, func(t *testing.T) {
			got := ErrMessageContains(testCase.Err, testCase.Code, testCase.Message)

			if got != testCase.Expected {
				t.Errorf("got %t, expected %t", got, testCase.Expected)
			}
		})
	}
}
