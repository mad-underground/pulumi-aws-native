// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::CloudFormation::Stack resource nests a stack as a resource in a top-level template.
func LookupStack(ctx *pulumi.Context, args *LookupStackArgs, opts ...pulumi.InvokeOption) (*LookupStackResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupStackResult
	err := ctx.Invoke("aws-native:cloudformation:getStack", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStackArgs struct {
	StackId string `pulumi:"stackId"`
}

type LookupStackResult struct {
	Capabilities                []StackCapabilitiesItem `pulumi:"capabilities"`
	ChangeSetId                 *string                 `pulumi:"changeSetId"`
	CreationTime                *string                 `pulumi:"creationTime"`
	Description                 *string                 `pulumi:"description"`
	DisableRollback             *bool                   `pulumi:"disableRollback"`
	EnableTerminationProtection *bool                   `pulumi:"enableTerminationProtection"`
	LastUpdateTime              *string                 `pulumi:"lastUpdateTime"`
	NotificationArns            []string                `pulumi:"notificationArns"`
	Outputs                     []StackOutputType       `pulumi:"outputs"`
	Parameters                  map[string]string       `pulumi:"parameters"`
	ParentId                    *string                 `pulumi:"parentId"`
	RoleArn                     *string                 `pulumi:"roleArn"`
	RootId                      *string                 `pulumi:"rootId"`
	StackId                     *string                 `pulumi:"stackId"`
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::CloudFormation::Stack` for more information about the expected schema for this property.
	StackPolicyBody   interface{}  `pulumi:"stackPolicyBody"`
	StackStatus       *StackStatus `pulumi:"stackStatus"`
	StackStatusReason *string      `pulumi:"stackStatusReason"`
	Tags              []aws.Tag    `pulumi:"tags"`
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::CloudFormation::Stack` for more information about the expected schema for this property.
	TemplateBody     interface{} `pulumi:"templateBody"`
	TimeoutInMinutes *int        `pulumi:"timeoutInMinutes"`
}

func LookupStackOutput(ctx *pulumi.Context, args LookupStackOutputArgs, opts ...pulumi.InvokeOption) LookupStackResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStackResult, error) {
			args := v.(LookupStackArgs)
			r, err := LookupStack(ctx, &args, opts...)
			var s LookupStackResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStackResultOutput)
}

type LookupStackOutputArgs struct {
	StackId pulumi.StringInput `pulumi:"stackId"`
}

func (LookupStackOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStackArgs)(nil)).Elem()
}

type LookupStackResultOutput struct{ *pulumi.OutputState }

func (LookupStackResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStackResult)(nil)).Elem()
}

func (o LookupStackResultOutput) ToLookupStackResultOutput() LookupStackResultOutput {
	return o
}

func (o LookupStackResultOutput) ToLookupStackResultOutputWithContext(ctx context.Context) LookupStackResultOutput {
	return o
}

func (o LookupStackResultOutput) Capabilities() StackCapabilitiesItemArrayOutput {
	return o.ApplyT(func(v LookupStackResult) []StackCapabilitiesItem { return v.Capabilities }).(StackCapabilitiesItemArrayOutput)
}

func (o LookupStackResultOutput) ChangeSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.ChangeSetId }).(pulumi.StringPtrOutput)
}

func (o LookupStackResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o LookupStackResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupStackResultOutput) DisableRollback() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *bool { return v.DisableRollback }).(pulumi.BoolPtrOutput)
}

func (o LookupStackResultOutput) EnableTerminationProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *bool { return v.EnableTerminationProtection }).(pulumi.BoolPtrOutput)
}

func (o LookupStackResultOutput) LastUpdateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.LastUpdateTime }).(pulumi.StringPtrOutput)
}

func (o LookupStackResultOutput) NotificationArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupStackResult) []string { return v.NotificationArns }).(pulumi.StringArrayOutput)
}

func (o LookupStackResultOutput) Outputs() StackOutputTypeArrayOutput {
	return o.ApplyT(func(v LookupStackResult) []StackOutputType { return v.Outputs }).(StackOutputTypeArrayOutput)
}

func (o LookupStackResultOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStackResult) map[string]string { return v.Parameters }).(pulumi.StringMapOutput)
}

func (o LookupStackResultOutput) ParentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.ParentId }).(pulumi.StringPtrOutput)
}

func (o LookupStackResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

func (o LookupStackResultOutput) RootId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.RootId }).(pulumi.StringPtrOutput)
}

func (o LookupStackResultOutput) StackId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.StackId }).(pulumi.StringPtrOutput)
}

// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::CloudFormation::Stack` for more information about the expected schema for this property.
func (o LookupStackResultOutput) StackPolicyBody() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupStackResult) interface{} { return v.StackPolicyBody }).(pulumi.AnyOutput)
}

func (o LookupStackResultOutput) StackStatus() StackStatusPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *StackStatus { return v.StackStatus }).(StackStatusPtrOutput)
}

func (o LookupStackResultOutput) StackStatusReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.StackStatusReason }).(pulumi.StringPtrOutput)
}

func (o LookupStackResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupStackResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::CloudFormation::Stack` for more information about the expected schema for this property.
func (o LookupStackResultOutput) TemplateBody() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupStackResult) interface{} { return v.TemplateBody }).(pulumi.AnyOutput)
}

func (o LookupStackResultOutput) TimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *int { return v.TimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStackResultOutput{})
}
