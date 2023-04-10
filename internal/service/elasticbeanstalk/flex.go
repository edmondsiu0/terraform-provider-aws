// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticbeanstalk

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
)

func flattenASG(list []*elasticbeanstalk.AutoScalingGroup) []string {
	strs := make([]string, 0, len(list))
	for _, r := range list {
		if r.Name != nil {
			strs = append(strs, *r.Name)
		}
	}
	return strs
}

func flattenLoadBalancers(list []*elasticbeanstalk.LoadBalancer) []string {
	strs := make([]string, 0, len(list))
	for _, r := range list {
		if r.Name != nil {
			strs = append(strs, *r.Name)
		}
	}
	return strs
}

func flattenInstances(list []*elasticbeanstalk.Instance) []string {
	strs := make([]string, 0, len(list))
	for _, r := range list {
		if r.Id != nil {
			strs = append(strs, *r.Id)
		}
	}
	return strs
}

func flattenLaunchConfigurations(list []*elasticbeanstalk.LaunchConfiguration) []string {
	strs := make([]string, 0, len(list))
	for _, r := range list {
		if r.Name != nil {
			strs = append(strs, *r.Name)
		}
	}
	return strs
}

func flattenQueues(list []*elasticbeanstalk.Queue) []string {
	strs := make([]string, 0, len(list))
	for _, r := range list {
		if r.URL != nil {
			strs = append(strs, *r.URL)
		}
	}
	return strs
}

func flattenTriggers(list []*elasticbeanstalk.Trigger) []string {
	strs := make([]string, 0, len(list))
	for _, r := range list {
		if r.Name != nil {
			strs = append(strs, *r.Name)
		}
	}
	return strs
}

func flattenResourceLifecycleConfig(rlc *elasticbeanstalk.ApplicationResourceLifecycleConfig) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, 1)

	anything_enabled := false
	appversion_lifecycle := make(map[string]interface{})

	if rlc.ServiceRole != nil {
		appversion_lifecycle["service_role"] = aws.StringValue(rlc.ServiceRole)
	}

	if vlc := rlc.VersionLifecycleConfig; vlc != nil {
		if mar := vlc.MaxAgeRule; mar != nil && aws.BoolValue(mar.Enabled) {
			anything_enabled = true
			appversion_lifecycle["max_age_in_days"] = aws.Int64Value(mar.MaxAgeInDays)
			appversion_lifecycle["delete_source_from_s3"] = aws.BoolValue(mar.DeleteSourceFromS3)
		}
		if mcr := vlc.MaxCountRule; mcr != nil && aws.BoolValue(mcr.Enabled) {
			anything_enabled = true
			appversion_lifecycle["max_count"] = aws.Int64Value(mcr.MaxCount)
			appversion_lifecycle["delete_source_from_s3"] = aws.BoolValue(mcr.DeleteSourceFromS3)
		}
	}

	if anything_enabled {
		result = append(result, appversion_lifecycle)
	}

	return result
}
