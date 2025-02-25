// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kendraranking

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type ExecutionPlanCapacityUnitsConfiguration struct {
	RescoreCapacityUnits int `pulumi:"rescoreCapacityUnits"`
}

// ExecutionPlanCapacityUnitsConfigurationInput is an input type that accepts ExecutionPlanCapacityUnitsConfigurationArgs and ExecutionPlanCapacityUnitsConfigurationOutput values.
// You can construct a concrete instance of `ExecutionPlanCapacityUnitsConfigurationInput` via:
//
//	ExecutionPlanCapacityUnitsConfigurationArgs{...}
type ExecutionPlanCapacityUnitsConfigurationInput interface {
	pulumi.Input

	ToExecutionPlanCapacityUnitsConfigurationOutput() ExecutionPlanCapacityUnitsConfigurationOutput
	ToExecutionPlanCapacityUnitsConfigurationOutputWithContext(context.Context) ExecutionPlanCapacityUnitsConfigurationOutput
}

type ExecutionPlanCapacityUnitsConfigurationArgs struct {
	RescoreCapacityUnits pulumi.IntInput `pulumi:"rescoreCapacityUnits"`
}

func (ExecutionPlanCapacityUnitsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExecutionPlanCapacityUnitsConfiguration)(nil)).Elem()
}

func (i ExecutionPlanCapacityUnitsConfigurationArgs) ToExecutionPlanCapacityUnitsConfigurationOutput() ExecutionPlanCapacityUnitsConfigurationOutput {
	return i.ToExecutionPlanCapacityUnitsConfigurationOutputWithContext(context.Background())
}

func (i ExecutionPlanCapacityUnitsConfigurationArgs) ToExecutionPlanCapacityUnitsConfigurationOutputWithContext(ctx context.Context) ExecutionPlanCapacityUnitsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExecutionPlanCapacityUnitsConfigurationOutput)
}

func (i ExecutionPlanCapacityUnitsConfigurationArgs) ToExecutionPlanCapacityUnitsConfigurationPtrOutput() ExecutionPlanCapacityUnitsConfigurationPtrOutput {
	return i.ToExecutionPlanCapacityUnitsConfigurationPtrOutputWithContext(context.Background())
}

func (i ExecutionPlanCapacityUnitsConfigurationArgs) ToExecutionPlanCapacityUnitsConfigurationPtrOutputWithContext(ctx context.Context) ExecutionPlanCapacityUnitsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExecutionPlanCapacityUnitsConfigurationOutput).ToExecutionPlanCapacityUnitsConfigurationPtrOutputWithContext(ctx)
}

// ExecutionPlanCapacityUnitsConfigurationPtrInput is an input type that accepts ExecutionPlanCapacityUnitsConfigurationArgs, ExecutionPlanCapacityUnitsConfigurationPtr and ExecutionPlanCapacityUnitsConfigurationPtrOutput values.
// You can construct a concrete instance of `ExecutionPlanCapacityUnitsConfigurationPtrInput` via:
//
//	        ExecutionPlanCapacityUnitsConfigurationArgs{...}
//
//	or:
//
//	        nil
type ExecutionPlanCapacityUnitsConfigurationPtrInput interface {
	pulumi.Input

	ToExecutionPlanCapacityUnitsConfigurationPtrOutput() ExecutionPlanCapacityUnitsConfigurationPtrOutput
	ToExecutionPlanCapacityUnitsConfigurationPtrOutputWithContext(context.Context) ExecutionPlanCapacityUnitsConfigurationPtrOutput
}

type executionPlanCapacityUnitsConfigurationPtrType ExecutionPlanCapacityUnitsConfigurationArgs

func ExecutionPlanCapacityUnitsConfigurationPtr(v *ExecutionPlanCapacityUnitsConfigurationArgs) ExecutionPlanCapacityUnitsConfigurationPtrInput {
	return (*executionPlanCapacityUnitsConfigurationPtrType)(v)
}

func (*executionPlanCapacityUnitsConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExecutionPlanCapacityUnitsConfiguration)(nil)).Elem()
}

func (i *executionPlanCapacityUnitsConfigurationPtrType) ToExecutionPlanCapacityUnitsConfigurationPtrOutput() ExecutionPlanCapacityUnitsConfigurationPtrOutput {
	return i.ToExecutionPlanCapacityUnitsConfigurationPtrOutputWithContext(context.Background())
}

func (i *executionPlanCapacityUnitsConfigurationPtrType) ToExecutionPlanCapacityUnitsConfigurationPtrOutputWithContext(ctx context.Context) ExecutionPlanCapacityUnitsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExecutionPlanCapacityUnitsConfigurationPtrOutput)
}

type ExecutionPlanCapacityUnitsConfigurationOutput struct{ *pulumi.OutputState }

func (ExecutionPlanCapacityUnitsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExecutionPlanCapacityUnitsConfiguration)(nil)).Elem()
}

func (o ExecutionPlanCapacityUnitsConfigurationOutput) ToExecutionPlanCapacityUnitsConfigurationOutput() ExecutionPlanCapacityUnitsConfigurationOutput {
	return o
}

func (o ExecutionPlanCapacityUnitsConfigurationOutput) ToExecutionPlanCapacityUnitsConfigurationOutputWithContext(ctx context.Context) ExecutionPlanCapacityUnitsConfigurationOutput {
	return o
}

func (o ExecutionPlanCapacityUnitsConfigurationOutput) ToExecutionPlanCapacityUnitsConfigurationPtrOutput() ExecutionPlanCapacityUnitsConfigurationPtrOutput {
	return o.ToExecutionPlanCapacityUnitsConfigurationPtrOutputWithContext(context.Background())
}

func (o ExecutionPlanCapacityUnitsConfigurationOutput) ToExecutionPlanCapacityUnitsConfigurationPtrOutputWithContext(ctx context.Context) ExecutionPlanCapacityUnitsConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExecutionPlanCapacityUnitsConfiguration) *ExecutionPlanCapacityUnitsConfiguration {
		return &v
	}).(ExecutionPlanCapacityUnitsConfigurationPtrOutput)
}

func (o ExecutionPlanCapacityUnitsConfigurationOutput) RescoreCapacityUnits() pulumi.IntOutput {
	return o.ApplyT(func(v ExecutionPlanCapacityUnitsConfiguration) int { return v.RescoreCapacityUnits }).(pulumi.IntOutput)
}

type ExecutionPlanCapacityUnitsConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ExecutionPlanCapacityUnitsConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExecutionPlanCapacityUnitsConfiguration)(nil)).Elem()
}

func (o ExecutionPlanCapacityUnitsConfigurationPtrOutput) ToExecutionPlanCapacityUnitsConfigurationPtrOutput() ExecutionPlanCapacityUnitsConfigurationPtrOutput {
	return o
}

func (o ExecutionPlanCapacityUnitsConfigurationPtrOutput) ToExecutionPlanCapacityUnitsConfigurationPtrOutputWithContext(ctx context.Context) ExecutionPlanCapacityUnitsConfigurationPtrOutput {
	return o
}

func (o ExecutionPlanCapacityUnitsConfigurationPtrOutput) Elem() ExecutionPlanCapacityUnitsConfigurationOutput {
	return o.ApplyT(func(v *ExecutionPlanCapacityUnitsConfiguration) ExecutionPlanCapacityUnitsConfiguration {
		if v != nil {
			return *v
		}
		var ret ExecutionPlanCapacityUnitsConfiguration
		return ret
	}).(ExecutionPlanCapacityUnitsConfigurationOutput)
}

func (o ExecutionPlanCapacityUnitsConfigurationPtrOutput) RescoreCapacityUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExecutionPlanCapacityUnitsConfiguration) *int {
		if v == nil {
			return nil
		}
		return &v.RescoreCapacityUnits
	}).(pulumi.IntPtrOutput)
}

// A label for tagging KendraRanking resources
type ExecutionPlanTag struct {
	// A string used to identify this tag
	Key string `pulumi:"key"`
	// A string containing the value for the tag
	Value string `pulumi:"value"`
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExecutionPlanCapacityUnitsConfigurationInput)(nil)).Elem(), ExecutionPlanCapacityUnitsConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExecutionPlanCapacityUnitsConfigurationPtrInput)(nil)).Elem(), ExecutionPlanCapacityUnitsConfigurationArgs{})
	pulumi.RegisterOutputType(ExecutionPlanCapacityUnitsConfigurationOutput{})
	pulumi.RegisterOutputType(ExecutionPlanCapacityUnitsConfigurationPtrOutput{})
}
