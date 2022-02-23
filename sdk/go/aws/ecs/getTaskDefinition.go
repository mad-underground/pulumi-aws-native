// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Schema describing various properties for ECS TaskDefinition
func LookupTaskDefinition(ctx *pulumi.Context, args *LookupTaskDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupTaskDefinitionResult, error) {
	var rv LookupTaskDefinitionResult
	err := ctx.Invoke("aws-native:ecs:getTaskDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTaskDefinitionArgs struct {
	// The Amazon Resource Name (ARN) of the Amazon ECS task definition
	TaskDefinitionArn string `pulumi:"taskDefinitionArn"`
}

type LookupTaskDefinitionResult struct {
	Tags []TaskDefinitionTag `pulumi:"tags"`
	// The Amazon Resource Name (ARN) of the Amazon ECS task definition
	TaskDefinitionArn *string `pulumi:"taskDefinitionArn"`
}

func LookupTaskDefinitionOutput(ctx *pulumi.Context, args LookupTaskDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupTaskDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTaskDefinitionResult, error) {
			args := v.(LookupTaskDefinitionArgs)
			r, err := LookupTaskDefinition(ctx, &args, opts...)
			return *r, err
		}).(LookupTaskDefinitionResultOutput)
}

type LookupTaskDefinitionOutputArgs struct {
	// The Amazon Resource Name (ARN) of the Amazon ECS task definition
	TaskDefinitionArn pulumi.StringInput `pulumi:"taskDefinitionArn"`
}

func (LookupTaskDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTaskDefinitionArgs)(nil)).Elem()
}

type LookupTaskDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupTaskDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTaskDefinitionResult)(nil)).Elem()
}

func (o LookupTaskDefinitionResultOutput) ToLookupTaskDefinitionResultOutput() LookupTaskDefinitionResultOutput {
	return o
}

func (o LookupTaskDefinitionResultOutput) ToLookupTaskDefinitionResultOutputWithContext(ctx context.Context) LookupTaskDefinitionResultOutput {
	return o
}

func (o LookupTaskDefinitionResultOutput) Tags() TaskDefinitionTagArrayOutput {
	return o.ApplyT(func(v LookupTaskDefinitionResult) []TaskDefinitionTag { return v.Tags }).(TaskDefinitionTagArrayOutput)
}

// The Amazon Resource Name (ARN) of the Amazon ECS task definition
func (o LookupTaskDefinitionResultOutput) TaskDefinitionArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTaskDefinitionResult) *string { return v.TaskDefinitionArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTaskDefinitionResultOutput{})
}