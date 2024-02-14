package test

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// An example of how to test the Terraform module in examples/terraform-aws-s3-example using Terratest.
func TestTerraformAwsS3Example(t *testing.T) {
	t.Parallel()

	// Give this S3 Bucket a unique ID for a name tag so we can distinguish it from any other Buckets provisioned
	// in your AWS account
	expectedName := fmt.Sprintf("terratest-aws-s3-example-%s", strings.ToLower(random.UniqueId()))

	// Give this S3 Bucket an environment to operate as a part of for the purposes of resource tagging
	expectedEnvironment := "AutomatedTesting"

	// Pick a random AWS region to test in. This helps ensure your code works in all regions.
	awsRegion := "us-west-2"

	// Create terraform.auto.tfvars file
	fileContent := fmt.Sprintf("tag_bucket_name = \"%s\"\ntag_bucket_environment = \"%s\"\nwith_policy = true", expectedName, expectedEnvironment)
	err := os.WriteFile("terraform.auto.tfvars", []byte(fileContent), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Run terraform init
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "..", // Update this path to your Terraform configuration

		// Environment variables to set when running Terraform
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
		},

		// Variables to pass to our Terraform code using VAR=value environment variables
		Vars: map[string]interface{}{
			"tag_bucket_name":        expectedName,
			"tag_bucket_environment": expectedEnvironment,
			"with_policy":            true,
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply` using the terraform CLI
	initAndApplyCmd := exec.Command("terraform", "init", "&&", "terraform", "apply", "-auto-approve")
	initAndApplyCmd.Dir = terraformOptions.TerraformDir
	initAndApplyCmd.Env = os.Environ()
	initAndApplyCmd.Env = append(initAndApplyCmd.Env, fmt.Sprintf("AWS_DEFAULT_REGION=%s", awsRegion))

	// Run the init and apply commands
	output, err := initAndApplyCmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to run init and apply commands: %v\nOutput:\n%s", err, output)
	}

	// Run `terraform output` to get the value of an output variable
	bucketID := terraform.Output(t, terraformOptions, "bucket_id")

	// Verify that our Bucket has versioning enabled
	actualStatus := aws.GetS3BucketVersioning(t, awsRegion, bucketID)
	expectedStatus := "Enabled"
	assert.Equal(t, expectedStatus, actualStatus)

	// Verify that our Bucket has a policy attached
	aws.AssertS3BucketPolicyExists(t, awsRegion, bucketID)
}
