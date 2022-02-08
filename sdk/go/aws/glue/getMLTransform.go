// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Glue::MLTransform
func LookupMLTransform(ctx *pulumi.Context, args *LookupMLTransformArgs, opts ...pulumi.InvokeOption) (*LookupMLTransformResult, error) {
	var rv LookupMLTransformResult
	err := ctx.Invoke("aws-native:glue:getMLTransform", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMLTransformArgs struct {
	Id string `pulumi:"id"`
}

type LookupMLTransformResult struct {
	Description         *string                         `pulumi:"description"`
	GlueVersion         *string                         `pulumi:"glueVersion"`
	Id                  *string                         `pulumi:"id"`
	MaxCapacity         *float64                        `pulumi:"maxCapacity"`
	MaxRetries          *int                            `pulumi:"maxRetries"`
	Name                *string                         `pulumi:"name"`
	NumberOfWorkers     *int                            `pulumi:"numberOfWorkers"`
	Role                *string                         `pulumi:"role"`
	Tags                interface{}                     `pulumi:"tags"`
	Timeout             *int                            `pulumi:"timeout"`
	TransformEncryption *MLTransformTransformEncryption `pulumi:"transformEncryption"`
	TransformParameters *MLTransformTransformParameters `pulumi:"transformParameters"`
	WorkerType          *string                         `pulumi:"workerType"`
}

func LookupMLTransformOutput(ctx *pulumi.Context, args LookupMLTransformOutputArgs, opts ...pulumi.InvokeOption) LookupMLTransformResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMLTransformResult, error) {
			args := v.(LookupMLTransformArgs)
			r, err := LookupMLTransform(ctx, &args, opts...)
			return *r, err
		}).(LookupMLTransformResultOutput)
}

type LookupMLTransformOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupMLTransformOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMLTransformArgs)(nil)).Elem()
}

type LookupMLTransformResultOutput struct{ *pulumi.OutputState }

func (LookupMLTransformResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMLTransformResult)(nil)).Elem()
}

func (o LookupMLTransformResultOutput) ToLookupMLTransformResultOutput() LookupMLTransformResultOutput {
	return o
}

func (o LookupMLTransformResultOutput) ToLookupMLTransformResultOutputWithContext(ctx context.Context) LookupMLTransformResultOutput {
	return o
}

func (o LookupMLTransformResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMLTransformResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupMLTransformResultOutput) GlueVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMLTransformResult) *string { return v.GlueVersion }).(pulumi.StringPtrOutput)
}

func (o LookupMLTransformResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMLTransformResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupMLTransformResultOutput) MaxCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupMLTransformResult) *float64 { return v.MaxCapacity }).(pulumi.Float64PtrOutput)
}

func (o LookupMLTransformResultOutput) MaxRetries() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupMLTransformResult) *int { return v.MaxRetries }).(pulumi.IntPtrOutput)
}

func (o LookupMLTransformResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMLTransformResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupMLTransformResultOutput) NumberOfWorkers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupMLTransformResult) *int { return v.NumberOfWorkers }).(pulumi.IntPtrOutput)
}

func (o LookupMLTransformResultOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMLTransformResult) *string { return v.Role }).(pulumi.StringPtrOutput)
}

func (o LookupMLTransformResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupMLTransformResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

func (o LookupMLTransformResultOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupMLTransformResult) *int { return v.Timeout }).(pulumi.IntPtrOutput)
}

func (o LookupMLTransformResultOutput) TransformEncryption() MLTransformTransformEncryptionPtrOutput {
	return o.ApplyT(func(v LookupMLTransformResult) *MLTransformTransformEncryption { return v.TransformEncryption }).(MLTransformTransformEncryptionPtrOutput)
}

func (o LookupMLTransformResultOutput) TransformParameters() MLTransformTransformParametersPtrOutput {
	return o.ApplyT(func(v LookupMLTransformResult) *MLTransformTransformParameters { return v.TransformParameters }).(MLTransformTransformParametersPtrOutput)
}

func (o LookupMLTransformResultOutput) WorkerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMLTransformResult) *string { return v.WorkerType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMLTransformResultOutput{})
}
