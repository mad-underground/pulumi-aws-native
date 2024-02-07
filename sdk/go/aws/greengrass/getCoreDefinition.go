// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package greengrass

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Greengrass::CoreDefinition
func LookupCoreDefinition(ctx *pulumi.Context, args *LookupCoreDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupCoreDefinitionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCoreDefinitionResult
	err := ctx.Invoke("aws-native:greengrass:getCoreDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCoreDefinitionArgs struct {
	Id string `pulumi:"id"`
}

type LookupCoreDefinitionResult struct {
	Arn              *string     `pulumi:"arn"`
	Id               *string     `pulumi:"id"`
	LatestVersionArn *string     `pulumi:"latestVersionArn"`
	Name             *string     `pulumi:"name"`
	Tags             interface{} `pulumi:"tags"`
}

func LookupCoreDefinitionOutput(ctx *pulumi.Context, args LookupCoreDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupCoreDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCoreDefinitionResult, error) {
			args := v.(LookupCoreDefinitionArgs)
			r, err := LookupCoreDefinition(ctx, &args, opts...)
			var s LookupCoreDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCoreDefinitionResultOutput)
}

type LookupCoreDefinitionOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupCoreDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCoreDefinitionArgs)(nil)).Elem()
}

type LookupCoreDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupCoreDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCoreDefinitionResult)(nil)).Elem()
}

func (o LookupCoreDefinitionResultOutput) ToLookupCoreDefinitionResultOutput() LookupCoreDefinitionResultOutput {
	return o
}

func (o LookupCoreDefinitionResultOutput) ToLookupCoreDefinitionResultOutputWithContext(ctx context.Context) LookupCoreDefinitionResultOutput {
	return o
}

func (o LookupCoreDefinitionResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCoreDefinitionResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupCoreDefinitionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCoreDefinitionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupCoreDefinitionResultOutput) LatestVersionArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCoreDefinitionResult) *string { return v.LatestVersionArn }).(pulumi.StringPtrOutput)
}

func (o LookupCoreDefinitionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCoreDefinitionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupCoreDefinitionResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupCoreDefinitionResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCoreDefinitionResultOutput{})
}
