// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package medialive

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::MediaLive::Input
func LookupInput(ctx *pulumi.Context, args *LookupInputArgs, opts ...pulumi.InvokeOption) (*LookupInputResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupInputResult
	err := ctx.Invoke("aws-native:medialive:getInput", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInputArgs struct {
	Id string `pulumi:"id"`
}

type LookupInputResult struct {
	Arn                 *string                        `pulumi:"arn"`
	Destinations        []InputDestinationRequest      `pulumi:"destinations"`
	Id                  *string                        `pulumi:"id"`
	InputDevices        []InputDeviceSettings          `pulumi:"inputDevices"`
	InputSecurityGroups []string                       `pulumi:"inputSecurityGroups"`
	MediaConnectFlows   []InputMediaConnectFlowRequest `pulumi:"mediaConnectFlows"`
	Name                *string                        `pulumi:"name"`
	RoleArn             *string                        `pulumi:"roleArn"`
	Sources             []InputSourceRequest           `pulumi:"sources"`
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::MediaLive::Input` for more information about the expected schema for this property.
	Tags interface{} `pulumi:"tags"`
}

func LookupInputOutput(ctx *pulumi.Context, args LookupInputOutputArgs, opts ...pulumi.InvokeOption) LookupInputResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInputResult, error) {
			args := v.(LookupInputArgs)
			r, err := LookupInput(ctx, &args, opts...)
			var s LookupInputResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInputResultOutput)
}

type LookupInputOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupInputOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInputArgs)(nil)).Elem()
}

type LookupInputResultOutput struct{ *pulumi.OutputState }

func (LookupInputResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInputResult)(nil)).Elem()
}

func (o LookupInputResultOutput) ToLookupInputResultOutput() LookupInputResultOutput {
	return o
}

func (o LookupInputResultOutput) ToLookupInputResultOutputWithContext(ctx context.Context) LookupInputResultOutput {
	return o
}

func (o LookupInputResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInputResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupInputResultOutput) Destinations() InputDestinationRequestArrayOutput {
	return o.ApplyT(func(v LookupInputResult) []InputDestinationRequest { return v.Destinations }).(InputDestinationRequestArrayOutput)
}

func (o LookupInputResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInputResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupInputResultOutput) InputDevices() InputDeviceSettingsArrayOutput {
	return o.ApplyT(func(v LookupInputResult) []InputDeviceSettings { return v.InputDevices }).(InputDeviceSettingsArrayOutput)
}

func (o LookupInputResultOutput) InputSecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInputResult) []string { return v.InputSecurityGroups }).(pulumi.StringArrayOutput)
}

func (o LookupInputResultOutput) MediaConnectFlows() InputMediaConnectFlowRequestArrayOutput {
	return o.ApplyT(func(v LookupInputResult) []InputMediaConnectFlowRequest { return v.MediaConnectFlows }).(InputMediaConnectFlowRequestArrayOutput)
}

func (o LookupInputResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInputResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupInputResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInputResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

func (o LookupInputResultOutput) Sources() InputSourceRequestArrayOutput {
	return o.ApplyT(func(v LookupInputResult) []InputSourceRequest { return v.Sources }).(InputSourceRequestArrayOutput)
}

// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::MediaLive::Input` for more information about the expected schema for this property.
func (o LookupInputResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupInputResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInputResultOutput{})
}
