// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fis

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::FIS::ExperimentTemplate
func LookupExperimentTemplate(ctx *pulumi.Context, args *LookupExperimentTemplateArgs, opts ...pulumi.InvokeOption) (*LookupExperimentTemplateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupExperimentTemplateResult
	err := ctx.Invoke("aws-native:fis:getExperimentTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExperimentTemplateArgs struct {
	Id string `pulumi:"id"`
}

type LookupExperimentTemplateResult struct {
	Actions           *ExperimentTemplateActionMap         `pulumi:"actions"`
	Description       *string                              `pulumi:"description"`
	ExperimentOptions *ExperimentTemplateExperimentOptions `pulumi:"experimentOptions"`
	Id                *string                              `pulumi:"id"`
	LogConfiguration  *ExperimentTemplateLogConfiguration  `pulumi:"logConfiguration"`
	RoleArn           *string                              `pulumi:"roleArn"`
	StopConditions    []ExperimentTemplateStopCondition    `pulumi:"stopConditions"`
	Targets           *ExperimentTemplateTargetMap         `pulumi:"targets"`
}

func LookupExperimentTemplateOutput(ctx *pulumi.Context, args LookupExperimentTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupExperimentTemplateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExperimentTemplateResult, error) {
			args := v.(LookupExperimentTemplateArgs)
			r, err := LookupExperimentTemplate(ctx, &args, opts...)
			var s LookupExperimentTemplateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExperimentTemplateResultOutput)
}

type LookupExperimentTemplateOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupExperimentTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExperimentTemplateArgs)(nil)).Elem()
}

type LookupExperimentTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupExperimentTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExperimentTemplateResult)(nil)).Elem()
}

func (o LookupExperimentTemplateResultOutput) ToLookupExperimentTemplateResultOutput() LookupExperimentTemplateResultOutput {
	return o
}

func (o LookupExperimentTemplateResultOutput) ToLookupExperimentTemplateResultOutputWithContext(ctx context.Context) LookupExperimentTemplateResultOutput {
	return o
}

func (o LookupExperimentTemplateResultOutput) Actions() ExperimentTemplateActionMapPtrOutput {
	return o.ApplyT(func(v LookupExperimentTemplateResult) *ExperimentTemplateActionMap { return v.Actions }).(ExperimentTemplateActionMapPtrOutput)
}

func (o LookupExperimentTemplateResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExperimentTemplateResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupExperimentTemplateResultOutput) ExperimentOptions() ExperimentTemplateExperimentOptionsPtrOutput {
	return o.ApplyT(func(v LookupExperimentTemplateResult) *ExperimentTemplateExperimentOptions {
		return v.ExperimentOptions
	}).(ExperimentTemplateExperimentOptionsPtrOutput)
}

func (o LookupExperimentTemplateResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExperimentTemplateResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupExperimentTemplateResultOutput) LogConfiguration() ExperimentTemplateLogConfigurationPtrOutput {
	return o.ApplyT(func(v LookupExperimentTemplateResult) *ExperimentTemplateLogConfiguration { return v.LogConfiguration }).(ExperimentTemplateLogConfigurationPtrOutput)
}

func (o LookupExperimentTemplateResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExperimentTemplateResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

func (o LookupExperimentTemplateResultOutput) StopConditions() ExperimentTemplateStopConditionArrayOutput {
	return o.ApplyT(func(v LookupExperimentTemplateResult) []ExperimentTemplateStopCondition { return v.StopConditions }).(ExperimentTemplateStopConditionArrayOutput)
}

func (o LookupExperimentTemplateResultOutput) Targets() ExperimentTemplateTargetMapPtrOutput {
	return o.ApplyT(func(v LookupExperimentTemplateResult) *ExperimentTemplateTargetMap { return v.Targets }).(ExperimentTemplateTargetMapPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExperimentTemplateResultOutput{})
}
