// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApiGateway::Resource
func LookupResource(ctx *pulumi.Context, args *LookupResourceArgs, opts ...pulumi.InvokeOption) (*LookupResourceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupResourceResult
	err := ctx.Invoke("aws-native:apigateway:getResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourceArgs struct {
	// A unique primary identifier for a Resource
	ResourceId string `pulumi:"resourceId"`
	// The ID of the RestApi resource in which you want to create this resource..
	RestApiId string `pulumi:"restApiId"`
}

type LookupResourceResult struct {
	// A unique primary identifier for a Resource
	ResourceId *string `pulumi:"resourceId"`
}

func LookupResourceOutput(ctx *pulumi.Context, args LookupResourceOutputArgs, opts ...pulumi.InvokeOption) LookupResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResourceResult, error) {
			args := v.(LookupResourceArgs)
			r, err := LookupResource(ctx, &args, opts...)
			var s LookupResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResourceResultOutput)
}

type LookupResourceOutputArgs struct {
	// A unique primary identifier for a Resource
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	// The ID of the RestApi resource in which you want to create this resource..
	RestApiId pulumi.StringInput `pulumi:"restApiId"`
}

func (LookupResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceArgs)(nil)).Elem()
}

type LookupResourceResultOutput struct{ *pulumi.OutputState }

func (LookupResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceResult)(nil)).Elem()
}

func (o LookupResourceResultOutput) ToLookupResourceResultOutput() LookupResourceResultOutput {
	return o
}

func (o LookupResourceResultOutput) ToLookupResourceResultOutputWithContext(ctx context.Context) LookupResourceResultOutput {
	return o
}

// A unique primary identifier for a Resource
func (o LookupResourceResultOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceResult) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourceResultOutput{})
}
