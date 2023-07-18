// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appconfig

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppConfig::Deployment
func LookupDeployment(ctx *pulumi.Context, args *LookupDeploymentArgs, opts ...pulumi.InvokeOption) (*LookupDeploymentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDeploymentResult
	err := ctx.Invoke("aws-native:appconfig:getDeployment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeploymentArgs struct {
	Id string `pulumi:"id"`
}

type LookupDeploymentResult struct {
	Id *string `pulumi:"id"`
}

func LookupDeploymentOutput(ctx *pulumi.Context, args LookupDeploymentOutputArgs, opts ...pulumi.InvokeOption) LookupDeploymentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeploymentResult, error) {
			args := v.(LookupDeploymentArgs)
			r, err := LookupDeployment(ctx, &args, opts...)
			var s LookupDeploymentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDeploymentResultOutput)
}

type LookupDeploymentOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupDeploymentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentArgs)(nil)).Elem()
}

type LookupDeploymentResultOutput struct{ *pulumi.OutputState }

func (LookupDeploymentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentResult)(nil)).Elem()
}

func (o LookupDeploymentResultOutput) ToLookupDeploymentResultOutput() LookupDeploymentResultOutput {
	return o
}

func (o LookupDeploymentResultOutput) ToLookupDeploymentResultOutputWithContext(ctx context.Context) LookupDeploymentResultOutput {
	return o
}

func (o LookupDeploymentResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeploymentResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeploymentResultOutput{})
}
