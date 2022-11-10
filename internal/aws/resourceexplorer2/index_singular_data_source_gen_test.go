// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package resourceexplorer2_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSResourceExplorer2IndexDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ResourceExplorer2::Index", "awscc_resourceexplorer2_index", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSResourceExplorer2IndexDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ResourceExplorer2::Index", "awscc_resourceexplorer2_index", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
