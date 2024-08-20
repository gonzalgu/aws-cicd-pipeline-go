package lib

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type PipelineStageProps struct {
	awscdk.StageProps
}

func NewPipelineStage(scope constructs.Construct, stageName string, props *PipelineStageProps) awscdk.Stage {
	var sprops awscdk.StageProps
	if props != nil {
		sprops = props.StageProps
	}
	stage := awscdk.NewStage(scope, &stageName, &sprops)
	NewLambdaStack(stage, "LambdaStack", stageName, nil)
	return stage
}
