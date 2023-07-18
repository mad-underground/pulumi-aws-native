// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appstream

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppStream::ImageBuilder
func LookupImageBuilder(ctx *pulumi.Context, args *LookupImageBuilderArgs, opts ...pulumi.InvokeOption) (*LookupImageBuilderResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupImageBuilderResult
	err := ctx.Invoke("aws-native:appstream:getImageBuilder", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupImageBuilderArgs struct {
	Name string `pulumi:"name"`
}

type LookupImageBuilderResult struct {
	AccessEndpoints             []ImageBuilderAccessEndpoint `pulumi:"accessEndpoints"`
	AppstreamAgentVersion       *string                      `pulumi:"appstreamAgentVersion"`
	Description                 *string                      `pulumi:"description"`
	DisplayName                 *string                      `pulumi:"displayName"`
	DomainJoinInfo              *ImageBuilderDomainJoinInfo  `pulumi:"domainJoinInfo"`
	EnableDefaultInternetAccess *bool                        `pulumi:"enableDefaultInternetAccess"`
	IamRoleArn                  *string                      `pulumi:"iamRoleArn"`
	ImageArn                    *string                      `pulumi:"imageArn"`
	ImageName                   *string                      `pulumi:"imageName"`
	InstanceType                *string                      `pulumi:"instanceType"`
	Name                        *string                      `pulumi:"name"`
	StreamingUrl                *string                      `pulumi:"streamingUrl"`
	Tags                        []ImageBuilderTag            `pulumi:"tags"`
	VpcConfig                   *ImageBuilderVpcConfig       `pulumi:"vpcConfig"`
}

func LookupImageBuilderOutput(ctx *pulumi.Context, args LookupImageBuilderOutputArgs, opts ...pulumi.InvokeOption) LookupImageBuilderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupImageBuilderResult, error) {
			args := v.(LookupImageBuilderArgs)
			r, err := LookupImageBuilder(ctx, &args, opts...)
			var s LookupImageBuilderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupImageBuilderResultOutput)
}

type LookupImageBuilderOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupImageBuilderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImageBuilderArgs)(nil)).Elem()
}

type LookupImageBuilderResultOutput struct{ *pulumi.OutputState }

func (LookupImageBuilderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImageBuilderResult)(nil)).Elem()
}

func (o LookupImageBuilderResultOutput) ToLookupImageBuilderResultOutput() LookupImageBuilderResultOutput {
	return o
}

func (o LookupImageBuilderResultOutput) ToLookupImageBuilderResultOutputWithContext(ctx context.Context) LookupImageBuilderResultOutput {
	return o
}

func (o LookupImageBuilderResultOutput) AccessEndpoints() ImageBuilderAccessEndpointArrayOutput {
	return o.ApplyT(func(v LookupImageBuilderResult) []ImageBuilderAccessEndpoint { return v.AccessEndpoints }).(ImageBuilderAccessEndpointArrayOutput)
}

func (o LookupImageBuilderResultOutput) AppstreamAgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageBuilderResult) *string { return v.AppstreamAgentVersion }).(pulumi.StringPtrOutput)
}

func (o LookupImageBuilderResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageBuilderResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupImageBuilderResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageBuilderResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupImageBuilderResultOutput) DomainJoinInfo() ImageBuilderDomainJoinInfoPtrOutput {
	return o.ApplyT(func(v LookupImageBuilderResult) *ImageBuilderDomainJoinInfo { return v.DomainJoinInfo }).(ImageBuilderDomainJoinInfoPtrOutput)
}

func (o LookupImageBuilderResultOutput) EnableDefaultInternetAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupImageBuilderResult) *bool { return v.EnableDefaultInternetAccess }).(pulumi.BoolPtrOutput)
}

func (o LookupImageBuilderResultOutput) IamRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageBuilderResult) *string { return v.IamRoleArn }).(pulumi.StringPtrOutput)
}

func (o LookupImageBuilderResultOutput) ImageArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageBuilderResult) *string { return v.ImageArn }).(pulumi.StringPtrOutput)
}

func (o LookupImageBuilderResultOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageBuilderResult) *string { return v.ImageName }).(pulumi.StringPtrOutput)
}

func (o LookupImageBuilderResultOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageBuilderResult) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o LookupImageBuilderResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageBuilderResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupImageBuilderResultOutput) StreamingUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageBuilderResult) *string { return v.StreamingUrl }).(pulumi.StringPtrOutput)
}

func (o LookupImageBuilderResultOutput) Tags() ImageBuilderTagArrayOutput {
	return o.ApplyT(func(v LookupImageBuilderResult) []ImageBuilderTag { return v.Tags }).(ImageBuilderTagArrayOutput)
}

func (o LookupImageBuilderResultOutput) VpcConfig() ImageBuilderVpcConfigPtrOutput {
	return o.ApplyT(func(v LookupImageBuilderResult) *ImageBuilderVpcConfig { return v.VpcConfig }).(ImageBuilderVpcConfigPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupImageBuilderResultOutput{})
}
