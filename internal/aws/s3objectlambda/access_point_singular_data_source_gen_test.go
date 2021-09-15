// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package s3objectlambda_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSS3ObjectLambdaAccessPointDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::S3ObjectLambda::AccessPoint", "awscc_s3objectlambda_access_point", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSS3ObjectLambdaAccessPointDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::S3ObjectLambda::AccessPoint", "awscc_s3objectlambda_access_point", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}