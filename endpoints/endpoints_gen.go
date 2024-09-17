// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by internal/generate/endpoints/main.go; DO NOT EDIT.

package endpoints

import (
	"github.com/YakDriver/regexache"
)

// All known partition IDs.
const (
	AwsPartitionID      = "aws"        // AWS Standard
	AwsCnPartitionID    = "aws-cn"     // AWS China
	AwsIsoPartitionID   = "aws-iso"    // AWS ISO (US)
	AwsIsoBPartitionID  = "aws-iso-b"  // AWS ISOB (US)
	AwsIsoEPartitionID  = "aws-iso-e"  // AWS ISOE (Europe)
	AwsIsoFPartitionID  = "aws-iso-f"  // AWS ISOF
	AwsUsGovPartitionID = "aws-us-gov" // AWS GovCloud (US)
)

// All known Region IDs.
const (
	// AWS Standard partition's Regions.
	AfSouth1RegionID     = "af-south-1"     // Africa (Cape Town)
	ApEast1RegionID      = "ap-east-1"      // Asia Pacific (Hong Kong)
	ApNortheast1RegionID = "ap-northeast-1" // Asia Pacific (Tokyo)
	ApNortheast2RegionID = "ap-northeast-2" // Asia Pacific (Seoul)
	ApNortheast3RegionID = "ap-northeast-3" // Asia Pacific (Osaka)
	ApSouth1RegionID     = "ap-south-1"     // Asia Pacific (Mumbai)
	ApSouth2RegionID     = "ap-south-2"     // Asia Pacific (Hyderabad)
	ApSoutheast1RegionID = "ap-southeast-1" // Asia Pacific (Singapore)
	ApSoutheast2RegionID = "ap-southeast-2" // Asia Pacific (Sydney)
	ApSoutheast3RegionID = "ap-southeast-3" // Asia Pacific (Jakarta)
	ApSoutheast4RegionID = "ap-southeast-4" // Asia Pacific (Melbourne)
	ApSoutheast5RegionID = "ap-southeast-5" // Asia Pacific (Malaysia)
	CaCentral1RegionID   = "ca-central-1"   // Canada (Central)
	CaWest1RegionID      = "ca-west-1"      // Canada West (Calgary)
	EuCentral1RegionID   = "eu-central-1"   // Europe (Frankfurt)
	EuCentral2RegionID   = "eu-central-2"   // Europe (Zurich)
	EuNorth1RegionID     = "eu-north-1"     // Europe (Stockholm)
	EuSouth1RegionID     = "eu-south-1"     // Europe (Milan)
	EuSouth2RegionID     = "eu-south-2"     // Europe (Spain)
	EuWest1RegionID      = "eu-west-1"      // Europe (Ireland)
	EuWest2RegionID      = "eu-west-2"      // Europe (London)
	EuWest3RegionID      = "eu-west-3"      // Europe (Paris)
	IlCentral1RegionID   = "il-central-1"   // Israel (Tel Aviv)
	MeCentral1RegionID   = "me-central-1"   // Middle East (UAE)
	MeSouth1RegionID     = "me-south-1"     // Middle East (Bahrain)
	SaEast1RegionID      = "sa-east-1"      // South America (Sao Paulo)
	UsEast1RegionID      = "us-east-1"      // US East (N. Virginia)
	UsEast2RegionID      = "us-east-2"      // US East (Ohio)
	UsWest1RegionID      = "us-west-1"      // US West (N. California)
	UsWest2RegionID      = "us-west-2"      // US West (Oregon)
	// AWS China partition's Regions.
	CnNorth1RegionID     = "cn-north-1"     // China (Beijing)
	CnNorthwest1RegionID = "cn-northwest-1" // China (Ningxia)
	// AWS ISO (US) partition's Regions.
	UsIsoEast1RegionID = "us-iso-east-1" // US ISO East
	UsIsoWest1RegionID = "us-iso-west-1" // US ISO WEST
	// AWS ISOB (US) partition's Regions.
	UsIsobEast1RegionID = "us-isob-east-1" // US ISOB East (Ohio)
	// AWS ISOE (Europe) partition's Regions.
	EuIsoeWest1RegionID = "eu-isoe-west-1" // EU ISOE West
	// AWS ISOF partition's Regions.
	// AWS GovCloud (US) partition's Regions.
	UsGovEast1RegionID = "us-gov-east-1" // AWS GovCloud (US-East)
	UsGovWest1RegionID = "us-gov-west-1" // AWS GovCloud (US-West)
)

type partitionAndRegions struct {
	partition Partition
	regions   map[string]Region
}

var (
	partitionsAndRegions = map[string]partitionAndRegions{
		AwsPartitionID: {
			partition: Partition{
				id:          AwsPartitionID,
				name:        "AWS Standard",
				dnsSuffix:   "amazonaws.com",
				regionRegex: regexache.MustCompile(`^(us|eu|ap|sa|ca|me|af|il|mx)\-\w+\-\d+$`),
			},
			regions: map[string]Region{
				AfSouth1RegionID: {
					id:          AfSouth1RegionID,
					description: "Africa (Cape Town)",
				},
				ApEast1RegionID: {
					id:          ApEast1RegionID,
					description: "Asia Pacific (Hong Kong)",
				},
				ApNortheast1RegionID: {
					id:          ApNortheast1RegionID,
					description: "Asia Pacific (Tokyo)",
				},
				ApNortheast2RegionID: {
					id:          ApNortheast2RegionID,
					description: "Asia Pacific (Seoul)",
				},
				ApNortheast3RegionID: {
					id:          ApNortheast3RegionID,
					description: "Asia Pacific (Osaka)",
				},
				ApSouth1RegionID: {
					id:          ApSouth1RegionID,
					description: "Asia Pacific (Mumbai)",
				},
				ApSouth2RegionID: {
					id:          ApSouth2RegionID,
					description: "Asia Pacific (Hyderabad)",
				},
				ApSoutheast1RegionID: {
					id:          ApSoutheast1RegionID,
					description: "Asia Pacific (Singapore)",
				},
				ApSoutheast2RegionID: {
					id:          ApSoutheast2RegionID,
					description: "Asia Pacific (Sydney)",
				},
				ApSoutheast3RegionID: {
					id:          ApSoutheast3RegionID,
					description: "Asia Pacific (Jakarta)",
				},
				ApSoutheast4RegionID: {
					id:          ApSoutheast4RegionID,
					description: "Asia Pacific (Melbourne)",
				},
				ApSoutheast5RegionID: {
					id:          ApSoutheast5RegionID,
					description: "Asia Pacific (Malaysia)",
				},
				CaCentral1RegionID: {
					id:          CaCentral1RegionID,
					description: "Canada (Central)",
				},
				CaWest1RegionID: {
					id:          CaWest1RegionID,
					description: "Canada West (Calgary)",
				},
				EuCentral1RegionID: {
					id:          EuCentral1RegionID,
					description: "Europe (Frankfurt)",
				},
				EuCentral2RegionID: {
					id:          EuCentral2RegionID,
					description: "Europe (Zurich)",
				},
				EuNorth1RegionID: {
					id:          EuNorth1RegionID,
					description: "Europe (Stockholm)",
				},
				EuSouth1RegionID: {
					id:          EuSouth1RegionID,
					description: "Europe (Milan)",
				},
				EuSouth2RegionID: {
					id:          EuSouth2RegionID,
					description: "Europe (Spain)",
				},
				EuWest1RegionID: {
					id:          EuWest1RegionID,
					description: "Europe (Ireland)",
				},
				EuWest2RegionID: {
					id:          EuWest2RegionID,
					description: "Europe (London)",
				},
				EuWest3RegionID: {
					id:          EuWest3RegionID,
					description: "Europe (Paris)",
				},
				IlCentral1RegionID: {
					id:          IlCentral1RegionID,
					description: "Israel (Tel Aviv)",
				},
				MeCentral1RegionID: {
					id:          MeCentral1RegionID,
					description: "Middle East (UAE)",
				},
				MeSouth1RegionID: {
					id:          MeSouth1RegionID,
					description: "Middle East (Bahrain)",
				},
				SaEast1RegionID: {
					id:          SaEast1RegionID,
					description: "South America (Sao Paulo)",
				},
				UsEast1RegionID: {
					id:          UsEast1RegionID,
					description: "US East (N. Virginia)",
				},
				UsEast2RegionID: {
					id:          UsEast2RegionID,
					description: "US East (Ohio)",
				},
				UsWest1RegionID: {
					id:          UsWest1RegionID,
					description: "US West (N. California)",
				},
				UsWest2RegionID: {
					id:          UsWest2RegionID,
					description: "US West (Oregon)",
				},
			},
		},
		AwsCnPartitionID: {
			partition: Partition{
				id:          AwsCnPartitionID,
				name:        "AWS China",
				dnsSuffix:   "amazonaws.com.cn",
				regionRegex: regexache.MustCompile(`^cn\-\w+\-\d+$`),
			},
			regions: map[string]Region{
				CnNorth1RegionID: {
					id:          CnNorth1RegionID,
					description: "China (Beijing)",
				},
				CnNorthwest1RegionID: {
					id:          CnNorthwest1RegionID,
					description: "China (Ningxia)",
				},
			},
		},
		AwsIsoPartitionID: {
			partition: Partition{
				id:          AwsIsoPartitionID,
				name:        "AWS ISO (US)",
				dnsSuffix:   "c2s.ic.gov",
				regionRegex: regexache.MustCompile(`^us\-iso\-\w+\-\d+$`),
			},
			regions: map[string]Region{
				UsIsoEast1RegionID: {
					id:          UsIsoEast1RegionID,
					description: "US ISO East",
				},
				UsIsoWest1RegionID: {
					id:          UsIsoWest1RegionID,
					description: "US ISO WEST",
				},
			},
		},
		AwsIsoBPartitionID: {
			partition: Partition{
				id:          AwsIsoBPartitionID,
				name:        "AWS ISOB (US)",
				dnsSuffix:   "sc2s.sgov.gov",
				regionRegex: regexache.MustCompile(`^us\-isob\-\w+\-\d+$`),
			},
			regions: map[string]Region{
				UsIsobEast1RegionID: {
					id:          UsIsobEast1RegionID,
					description: "US ISOB East (Ohio)",
				},
			},
		},
		AwsIsoEPartitionID: {
			partition: Partition{
				id:          AwsIsoEPartitionID,
				name:        "AWS ISOE (Europe)",
				dnsSuffix:   "cloud.adc-e.uk",
				regionRegex: regexache.MustCompile(`^eu\-isoe\-\w+\-\d+$`),
			},
			regions: map[string]Region{
				EuIsoeWest1RegionID: {
					id:          EuIsoeWest1RegionID,
					description: "EU ISOE West",
				},
			},
		},
		AwsIsoFPartitionID: {
			partition: Partition{
				id:          AwsIsoFPartitionID,
				name:        "AWS ISOF",
				dnsSuffix:   "csp.hci.ic.gov",
				regionRegex: regexache.MustCompile(`^us\-isof\-\w+\-\d+$`),
			},
			regions: map[string]Region{},
		},
		AwsUsGovPartitionID: {
			partition: Partition{
				id:          AwsUsGovPartitionID,
				name:        "AWS GovCloud (US)",
				dnsSuffix:   "amazonaws.com",
				regionRegex: regexache.MustCompile(`^us\-gov\-\w+\-\d+$`),
			},
			regions: map[string]Region{
				UsGovEast1RegionID: {
					id:          UsGovEast1RegionID,
					description: "AWS GovCloud (US-East)",
				},
				UsGovWest1RegionID: {
					id:          UsGovWest1RegionID,
					description: "AWS GovCloud (US-West)",
				},
			},
		},
	}
)
