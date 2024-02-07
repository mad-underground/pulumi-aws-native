// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package configuration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A conformance pack is a collection of AWS Config rules and remediation actions that can be easily deployed as a single entity in an account and a region or across an entire AWS Organization.
func LookupConformancePack(ctx *pulumi.Context, args *LookupConformancePackArgs, opts ...pulumi.InvokeOption) (*LookupConformancePackResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupConformancePackResult
	err := ctx.Invoke("aws-native:configuration:getConformancePack", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConformancePackArgs struct {
	// Name of the conformance pack which will be assigned as the unique identifier.
	ConformancePackName string `pulumi:"conformancePackName"`
}

type LookupConformancePackResult struct {
	// A list of ConformancePackInputParameter objects.
	ConformancePackInputParameters []ConformancePackInputParameter `pulumi:"conformancePackInputParameters"`
	// AWS Config stores intermediate files while processing conformance pack template.
	DeliveryS3Bucket *string `pulumi:"deliveryS3Bucket"`
	// The prefix for delivery S3 bucket.
	DeliveryS3KeyPrefix *string `pulumi:"deliveryS3KeyPrefix"`
}

func LookupConformancePackOutput(ctx *pulumi.Context, args LookupConformancePackOutputArgs, opts ...pulumi.InvokeOption) LookupConformancePackResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConformancePackResult, error) {
			args := v.(LookupConformancePackArgs)
			r, err := LookupConformancePack(ctx, &args, opts...)
			var s LookupConformancePackResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConformancePackResultOutput)
}

type LookupConformancePackOutputArgs struct {
	// Name of the conformance pack which will be assigned as the unique identifier.
	ConformancePackName pulumi.StringInput `pulumi:"conformancePackName"`
}

func (LookupConformancePackOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConformancePackArgs)(nil)).Elem()
}

type LookupConformancePackResultOutput struct{ *pulumi.OutputState }

func (LookupConformancePackResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConformancePackResult)(nil)).Elem()
}

func (o LookupConformancePackResultOutput) ToLookupConformancePackResultOutput() LookupConformancePackResultOutput {
	return o
}

func (o LookupConformancePackResultOutput) ToLookupConformancePackResultOutputWithContext(ctx context.Context) LookupConformancePackResultOutput {
	return o
}

// A list of ConformancePackInputParameter objects.
func (o LookupConformancePackResultOutput) ConformancePackInputParameters() ConformancePackInputParameterArrayOutput {
	return o.ApplyT(func(v LookupConformancePackResult) []ConformancePackInputParameter {
		return v.ConformancePackInputParameters
	}).(ConformancePackInputParameterArrayOutput)
}

// AWS Config stores intermediate files while processing conformance pack template.
func (o LookupConformancePackResultOutput) DeliveryS3Bucket() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConformancePackResult) *string { return v.DeliveryS3Bucket }).(pulumi.StringPtrOutput)
}

// The prefix for delivery S3 bucket.
func (o LookupConformancePackResultOutput) DeliveryS3KeyPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConformancePackResult) *string { return v.DeliveryS3KeyPrefix }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConformancePackResultOutput{})
}
