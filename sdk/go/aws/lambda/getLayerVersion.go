// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Lambda::LayerVersion
func LookupLayerVersion(ctx *pulumi.Context, args *LookupLayerVersionArgs, opts ...pulumi.InvokeOption) (*LookupLayerVersionResult, error) {
	var rv LookupLayerVersionResult
	err := ctx.Invoke("aws-native:lambda:getLayerVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLayerVersionArgs struct {
	// The ARN of the version.
	Id string `pulumi:"id"`
}

type LookupLayerVersionResult struct {
	// The ARN of the version.
	Id *string `pulumi:"id"`
}

func LookupLayerVersionOutput(ctx *pulumi.Context, args LookupLayerVersionOutputArgs, opts ...pulumi.InvokeOption) LookupLayerVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLayerVersionResult, error) {
			args := v.(LookupLayerVersionArgs)
			r, err := LookupLayerVersion(ctx, &args, opts...)
			var s LookupLayerVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLayerVersionResultOutput)
}

type LookupLayerVersionOutputArgs struct {
	// The ARN of the version.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupLayerVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLayerVersionArgs)(nil)).Elem()
}

type LookupLayerVersionResultOutput struct{ *pulumi.OutputState }

func (LookupLayerVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLayerVersionResult)(nil)).Elem()
}

func (o LookupLayerVersionResultOutput) ToLookupLayerVersionResultOutput() LookupLayerVersionResultOutput {
	return o
}

func (o LookupLayerVersionResultOutput) ToLookupLayerVersionResultOutputWithContext(ctx context.Context) LookupLayerVersionResultOutput {
	return o
}

// The ARN of the version.
func (o LookupLayerVersionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLayerVersionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLayerVersionResultOutput{})
}
