// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SageMaker::Workteam
func LookupWorkteam(ctx *pulumi.Context, args *LookupWorkteamArgs, opts ...pulumi.InvokeOption) (*LookupWorkteamResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkteamResult
	err := ctx.Invoke("aws-native:sagemaker:getWorkteam", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkteamArgs struct {
	Id string `pulumi:"id"`
}

type LookupWorkteamResult struct {
	Description               *string                            `pulumi:"description"`
	Id                        *string                            `pulumi:"id"`
	MemberDefinitions         []WorkteamMemberDefinition         `pulumi:"memberDefinitions"`
	NotificationConfiguration *WorkteamNotificationConfiguration `pulumi:"notificationConfiguration"`
	Tags                      []aws.Tag                          `pulumi:"tags"`
}

func LookupWorkteamOutput(ctx *pulumi.Context, args LookupWorkteamOutputArgs, opts ...pulumi.InvokeOption) LookupWorkteamResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkteamResult, error) {
			args := v.(LookupWorkteamArgs)
			r, err := LookupWorkteam(ctx, &args, opts...)
			var s LookupWorkteamResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkteamResultOutput)
}

type LookupWorkteamOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupWorkteamOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkteamArgs)(nil)).Elem()
}

type LookupWorkteamResultOutput struct{ *pulumi.OutputState }

func (LookupWorkteamResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkteamResult)(nil)).Elem()
}

func (o LookupWorkteamResultOutput) ToLookupWorkteamResultOutput() LookupWorkteamResultOutput {
	return o
}

func (o LookupWorkteamResultOutput) ToLookupWorkteamResultOutputWithContext(ctx context.Context) LookupWorkteamResultOutput {
	return o
}

func (o LookupWorkteamResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkteamResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupWorkteamResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkteamResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupWorkteamResultOutput) MemberDefinitions() WorkteamMemberDefinitionArrayOutput {
	return o.ApplyT(func(v LookupWorkteamResult) []WorkteamMemberDefinition { return v.MemberDefinitions }).(WorkteamMemberDefinitionArrayOutput)
}

func (o LookupWorkteamResultOutput) NotificationConfiguration() WorkteamNotificationConfigurationPtrOutput {
	return o.ApplyT(func(v LookupWorkteamResult) *WorkteamNotificationConfiguration { return v.NotificationConfiguration }).(WorkteamNotificationConfigurationPtrOutput)
}

func (o LookupWorkteamResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupWorkteamResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkteamResultOutput{})
}
