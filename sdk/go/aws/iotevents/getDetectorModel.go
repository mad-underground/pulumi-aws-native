// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iotevents

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::IoTEvents::DetectorModel resource creates a detector model. You create a *detector model* (a model of your equipment or process) using *states*. For each state, you define conditional (Boolean) logic that evaluates the incoming inputs to detect significant events. When an event is detected, it can change the state or trigger custom-built or predefined actions using other AWS services. You can define additional events that trigger actions when entering or exiting a state and, optionally, when a condition is met. For more information, see [How to Use AWS IoT Events](https://docs.aws.amazon.com/iotevents/latest/developerguide/how-to-use-iotevents.html) in the *AWS IoT Events Developer Guide*.
func LookupDetectorModel(ctx *pulumi.Context, args *LookupDetectorModelArgs, opts ...pulumi.InvokeOption) (*LookupDetectorModelResult, error) {
	var rv LookupDetectorModelResult
	err := ctx.Invoke("aws-native:iotevents:getDetectorModel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDetectorModelArgs struct {
	// The name of the detector model.
	DetectorModelName string `pulumi:"detectorModelName"`
}

type LookupDetectorModelResult struct {
	DetectorModelDefinition *DetectorModelDefinition `pulumi:"detectorModelDefinition"`
	// A brief description of the detector model.
	DetectorModelDescription *string `pulumi:"detectorModelDescription"`
	// Information about the order in which events are evaluated and how actions are executed.
	EvaluationMethod *DetectorModelEvaluationMethod `pulumi:"evaluationMethod"`
	// The ARN of the role that grants permission to AWS IoT Events to perform its operations.
	RoleArn *string `pulumi:"roleArn"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).
	Tags []DetectorModelTag `pulumi:"tags"`
}

func LookupDetectorModelOutput(ctx *pulumi.Context, args LookupDetectorModelOutputArgs, opts ...pulumi.InvokeOption) LookupDetectorModelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDetectorModelResult, error) {
			args := v.(LookupDetectorModelArgs)
			r, err := LookupDetectorModel(ctx, &args, opts...)
			return *r, err
		}).(LookupDetectorModelResultOutput)
}

type LookupDetectorModelOutputArgs struct {
	// The name of the detector model.
	DetectorModelName pulumi.StringInput `pulumi:"detectorModelName"`
}

func (LookupDetectorModelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDetectorModelArgs)(nil)).Elem()
}

type LookupDetectorModelResultOutput struct{ *pulumi.OutputState }

func (LookupDetectorModelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDetectorModelResult)(nil)).Elem()
}

func (o LookupDetectorModelResultOutput) ToLookupDetectorModelResultOutput() LookupDetectorModelResultOutput {
	return o
}

func (o LookupDetectorModelResultOutput) ToLookupDetectorModelResultOutputWithContext(ctx context.Context) LookupDetectorModelResultOutput {
	return o
}

func (o LookupDetectorModelResultOutput) DetectorModelDefinition() DetectorModelDefinitionPtrOutput {
	return o.ApplyT(func(v LookupDetectorModelResult) *DetectorModelDefinition { return v.DetectorModelDefinition }).(DetectorModelDefinitionPtrOutput)
}

// A brief description of the detector model.
func (o LookupDetectorModelResultOutput) DetectorModelDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDetectorModelResult) *string { return v.DetectorModelDescription }).(pulumi.StringPtrOutput)
}

// Information about the order in which events are evaluated and how actions are executed.
func (o LookupDetectorModelResultOutput) EvaluationMethod() DetectorModelEvaluationMethodPtrOutput {
	return o.ApplyT(func(v LookupDetectorModelResult) *DetectorModelEvaluationMethod { return v.EvaluationMethod }).(DetectorModelEvaluationMethodPtrOutput)
}

// The ARN of the role that grants permission to AWS IoT Events to perform its operations.
func (o LookupDetectorModelResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDetectorModelResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
//
// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).
func (o LookupDetectorModelResultOutput) Tags() DetectorModelTagArrayOutput {
	return o.ApplyT(func(v LookupDetectorModelResult) []DetectorModelTag { return v.Tags }).(DetectorModelTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDetectorModelResultOutput{})
}
