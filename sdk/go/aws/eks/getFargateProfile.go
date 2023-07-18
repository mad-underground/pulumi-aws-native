// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eks

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Schema for AWS::EKS::FargateProfile
func LookupFargateProfile(ctx *pulumi.Context, args *LookupFargateProfileArgs, opts ...pulumi.InvokeOption) (*LookupFargateProfileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFargateProfileResult
	err := ctx.Invoke("aws-native:eks:getFargateProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFargateProfileArgs struct {
	// Name of the Cluster
	ClusterName string `pulumi:"clusterName"`
	// Name of FargateProfile
	FargateProfileName string `pulumi:"fargateProfileName"`
}

type LookupFargateProfileResult struct {
	Arn *string `pulumi:"arn"`
	// An array of key-value pairs to apply to this resource.
	Tags []FargateProfileTag `pulumi:"tags"`
}

func LookupFargateProfileOutput(ctx *pulumi.Context, args LookupFargateProfileOutputArgs, opts ...pulumi.InvokeOption) LookupFargateProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFargateProfileResult, error) {
			args := v.(LookupFargateProfileArgs)
			r, err := LookupFargateProfile(ctx, &args, opts...)
			var s LookupFargateProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFargateProfileResultOutput)
}

type LookupFargateProfileOutputArgs struct {
	// Name of the Cluster
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// Name of FargateProfile
	FargateProfileName pulumi.StringInput `pulumi:"fargateProfileName"`
}

func (LookupFargateProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFargateProfileArgs)(nil)).Elem()
}

type LookupFargateProfileResultOutput struct{ *pulumi.OutputState }

func (LookupFargateProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFargateProfileResult)(nil)).Elem()
}

func (o LookupFargateProfileResultOutput) ToLookupFargateProfileResultOutput() LookupFargateProfileResultOutput {
	return o
}

func (o LookupFargateProfileResultOutput) ToLookupFargateProfileResultOutputWithContext(ctx context.Context) LookupFargateProfileResultOutput {
	return o
}

func (o LookupFargateProfileResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFargateProfileResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupFargateProfileResultOutput) Tags() FargateProfileTagArrayOutput {
	return o.ApplyT(func(v LookupFargateProfileResult) []FargateProfileTag { return v.Tags }).(FargateProfileTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFargateProfileResultOutput{})
}
