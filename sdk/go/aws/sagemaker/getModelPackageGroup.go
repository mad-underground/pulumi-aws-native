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

// Resource Type definition for AWS::SageMaker::ModelPackageGroup
func LookupModelPackageGroup(ctx *pulumi.Context, args *LookupModelPackageGroupArgs, opts ...pulumi.InvokeOption) (*LookupModelPackageGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupModelPackageGroupResult
	err := ctx.Invoke("aws-native:sagemaker:getModelPackageGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupModelPackageGroupArgs struct {
	ModelPackageGroupArn string `pulumi:"modelPackageGroupArn"`
}

type LookupModelPackageGroupResult struct {
	// The time at which the model package group was created.
	CreationTime         *string `pulumi:"creationTime"`
	ModelPackageGroupArn *string `pulumi:"modelPackageGroupArn"`
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::SageMaker::ModelPackageGroup` for more information about the expected schema for this property.
	ModelPackageGroupPolicy interface{} `pulumi:"modelPackageGroupPolicy"`
	// The status of a modelpackage group job.
	ModelPackageGroupStatus *ModelPackageGroupStatus `pulumi:"modelPackageGroupStatus"`
	// An array of key-value pairs to apply to this resource.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupModelPackageGroupOutput(ctx *pulumi.Context, args LookupModelPackageGroupOutputArgs, opts ...pulumi.InvokeOption) LookupModelPackageGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupModelPackageGroupResult, error) {
			args := v.(LookupModelPackageGroupArgs)
			r, err := LookupModelPackageGroup(ctx, &args, opts...)
			var s LookupModelPackageGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupModelPackageGroupResultOutput)
}

type LookupModelPackageGroupOutputArgs struct {
	ModelPackageGroupArn pulumi.StringInput `pulumi:"modelPackageGroupArn"`
}

func (LookupModelPackageGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModelPackageGroupArgs)(nil)).Elem()
}

type LookupModelPackageGroupResultOutput struct{ *pulumi.OutputState }

func (LookupModelPackageGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModelPackageGroupResult)(nil)).Elem()
}

func (o LookupModelPackageGroupResultOutput) ToLookupModelPackageGroupResultOutput() LookupModelPackageGroupResultOutput {
	return o
}

func (o LookupModelPackageGroupResultOutput) ToLookupModelPackageGroupResultOutputWithContext(ctx context.Context) LookupModelPackageGroupResultOutput {
	return o
}

// The time at which the model package group was created.
func (o LookupModelPackageGroupResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModelPackageGroupResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o LookupModelPackageGroupResultOutput) ModelPackageGroupArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModelPackageGroupResult) *string { return v.ModelPackageGroupArn }).(pulumi.StringPtrOutput)
}

// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::SageMaker::ModelPackageGroup` for more information about the expected schema for this property.
func (o LookupModelPackageGroupResultOutput) ModelPackageGroupPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupModelPackageGroupResult) interface{} { return v.ModelPackageGroupPolicy }).(pulumi.AnyOutput)
}

// The status of a modelpackage group job.
func (o LookupModelPackageGroupResultOutput) ModelPackageGroupStatus() ModelPackageGroupStatusPtrOutput {
	return o.ApplyT(func(v LookupModelPackageGroupResult) *ModelPackageGroupStatus { return v.ModelPackageGroupStatus }).(ModelPackageGroupStatusPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupModelPackageGroupResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupModelPackageGroupResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupModelPackageGroupResultOutput{})
}
