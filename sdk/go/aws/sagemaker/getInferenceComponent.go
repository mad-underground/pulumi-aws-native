// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::SageMaker::InferenceComponent
func LookupInferenceComponent(ctx *pulumi.Context, args *LookupInferenceComponentArgs, opts ...pulumi.InvokeOption) (*LookupInferenceComponentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupInferenceComponentResult
	err := ctx.Invoke("aws-native:sagemaker:getInferenceComponent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInferenceComponentArgs struct {
	InferenceComponentArn string `pulumi:"inferenceComponentArn"`
}

type LookupInferenceComponentResult struct {
	CreationTime             *string                          `pulumi:"creationTime"`
	EndpointArn              *string                          `pulumi:"endpointArn"`
	EndpointName             *string                          `pulumi:"endpointName"`
	FailureReason            *string                          `pulumi:"failureReason"`
	InferenceComponentArn    *string                          `pulumi:"inferenceComponentArn"`
	InferenceComponentName   *string                          `pulumi:"inferenceComponentName"`
	InferenceComponentStatus *InferenceComponentStatus        `pulumi:"inferenceComponentStatus"`
	LastModifiedTime         *string                          `pulumi:"lastModifiedTime"`
	RuntimeConfig            *InferenceComponentRuntimeConfig `pulumi:"runtimeConfig"`
	Specification            *InferenceComponentSpecification `pulumi:"specification"`
	Tags                     []InferenceComponentTag          `pulumi:"tags"`
	VariantName              *string                          `pulumi:"variantName"`
}

func LookupInferenceComponentOutput(ctx *pulumi.Context, args LookupInferenceComponentOutputArgs, opts ...pulumi.InvokeOption) LookupInferenceComponentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInferenceComponentResult, error) {
			args := v.(LookupInferenceComponentArgs)
			r, err := LookupInferenceComponent(ctx, &args, opts...)
			var s LookupInferenceComponentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInferenceComponentResultOutput)
}

type LookupInferenceComponentOutputArgs struct {
	InferenceComponentArn pulumi.StringInput `pulumi:"inferenceComponentArn"`
}

func (LookupInferenceComponentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInferenceComponentArgs)(nil)).Elem()
}

type LookupInferenceComponentResultOutput struct{ *pulumi.OutputState }

func (LookupInferenceComponentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInferenceComponentResult)(nil)).Elem()
}

func (o LookupInferenceComponentResultOutput) ToLookupInferenceComponentResultOutput() LookupInferenceComponentResultOutput {
	return o
}

func (o LookupInferenceComponentResultOutput) ToLookupInferenceComponentResultOutputWithContext(ctx context.Context) LookupInferenceComponentResultOutput {
	return o
}

func (o LookupInferenceComponentResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupInferenceComponentResult] {
	return pulumix.Output[LookupInferenceComponentResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupInferenceComponentResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInferenceComponentResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o LookupInferenceComponentResultOutput) EndpointArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInferenceComponentResult) *string { return v.EndpointArn }).(pulumi.StringPtrOutput)
}

func (o LookupInferenceComponentResultOutput) EndpointName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInferenceComponentResult) *string { return v.EndpointName }).(pulumi.StringPtrOutput)
}

func (o LookupInferenceComponentResultOutput) FailureReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInferenceComponentResult) *string { return v.FailureReason }).(pulumi.StringPtrOutput)
}

func (o LookupInferenceComponentResultOutput) InferenceComponentArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInferenceComponentResult) *string { return v.InferenceComponentArn }).(pulumi.StringPtrOutput)
}

func (o LookupInferenceComponentResultOutput) InferenceComponentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInferenceComponentResult) *string { return v.InferenceComponentName }).(pulumi.StringPtrOutput)
}

func (o LookupInferenceComponentResultOutput) InferenceComponentStatus() InferenceComponentStatusPtrOutput {
	return o.ApplyT(func(v LookupInferenceComponentResult) *InferenceComponentStatus { return v.InferenceComponentStatus }).(InferenceComponentStatusPtrOutput)
}

func (o LookupInferenceComponentResultOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInferenceComponentResult) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o LookupInferenceComponentResultOutput) RuntimeConfig() InferenceComponentRuntimeConfigPtrOutput {
	return o.ApplyT(func(v LookupInferenceComponentResult) *InferenceComponentRuntimeConfig { return v.RuntimeConfig }).(InferenceComponentRuntimeConfigPtrOutput)
}

func (o LookupInferenceComponentResultOutput) Specification() InferenceComponentSpecificationPtrOutput {
	return o.ApplyT(func(v LookupInferenceComponentResult) *InferenceComponentSpecification { return v.Specification }).(InferenceComponentSpecificationPtrOutput)
}

func (o LookupInferenceComponentResultOutput) Tags() InferenceComponentTagArrayOutput {
	return o.ApplyT(func(v LookupInferenceComponentResult) []InferenceComponentTag { return v.Tags }).(InferenceComponentTagArrayOutput)
}

func (o LookupInferenceComponentResultOutput) VariantName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInferenceComponentResult) *string { return v.VariantName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInferenceComponentResultOutput{})
}
