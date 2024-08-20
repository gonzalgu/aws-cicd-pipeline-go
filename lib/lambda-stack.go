package lib

import (
	"path"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type LambdaStackProps struct {
	StackProps awscdk.StackProps
}

func NewLambdaStack(scope constructs.Construct, id string, stageName string, props *AwsCicdPipelineGoStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}

	/*
		wd, err := path.Dir()
		if err != nil {
			panic("error getting current directory")
		}
	*/

	stack := awscdk.NewStack(scope, &id, &sprops)
	awslambda.NewFunction(stack, jsii.String("LambdaFunction"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_NODEJS_16_X(),
		Handler: jsii.String("handler.handler"),
		Code:    awslambda.AssetCode_FromAsset(jsii.String(path.Join("lib", "lambda")), nil),
		Environment: &map[string]*string{
			"stageName": jsii.String(stageName),
		},
	})
	return stack
}
