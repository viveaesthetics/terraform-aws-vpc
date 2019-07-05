package tests

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestVpcExist(t *testing.T) {
    expectedVpcCidr := "172.17.48.0/20"

	terraformOptions := &terraform.Options {
		// The path to where our Terraform code is located
		TerraformDir: "fixture/",

		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	actualVpcCidr := terraform.Output(t, terraformOptions, "vpc_cidr_block")

	// Verify we're getting back the outputs we expect
	assert.Equal(t, expectedVpcCidr, actualVpcCidr)
}