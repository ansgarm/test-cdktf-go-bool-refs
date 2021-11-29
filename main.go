package main

import (
	"cdk.tf/go/stack/generated/random"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	// The code that defines your stack goes here
	random.NewRandomProvider(stack, jsii.String("provider"), &random.RandomProviderConfig{})

	local := cdktf.NewTerraformLocal(stack, jsii.String("local_bool"), true)

	randomStr := random.NewString(stack, jsii.String("randomInput"), &random.StringConfig{
		Length: jsii.Number(10),
	})

	random.NewString(stack, jsii.String("random"), &random.StringConfig{
		Length: jsii.Number(10),
		Lower:  local.AsBoolean(), // works
		Upper:  randomStr.Lower(), // fails
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "test-cdktf-go-bool-refs")

	app.Synth()
}
