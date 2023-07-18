// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mediaconvert

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type JobTemplateAccelerationSettings struct {
	Mode string `pulumi:"mode"`
}

// JobTemplateAccelerationSettingsInput is an input type that accepts JobTemplateAccelerationSettingsArgs and JobTemplateAccelerationSettingsOutput values.
// You can construct a concrete instance of `JobTemplateAccelerationSettingsInput` via:
//
//	JobTemplateAccelerationSettingsArgs{...}
type JobTemplateAccelerationSettingsInput interface {
	pulumi.Input

	ToJobTemplateAccelerationSettingsOutput() JobTemplateAccelerationSettingsOutput
	ToJobTemplateAccelerationSettingsOutputWithContext(context.Context) JobTemplateAccelerationSettingsOutput
}

type JobTemplateAccelerationSettingsArgs struct {
	Mode pulumi.StringInput `pulumi:"mode"`
}

func (JobTemplateAccelerationSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobTemplateAccelerationSettings)(nil)).Elem()
}

func (i JobTemplateAccelerationSettingsArgs) ToJobTemplateAccelerationSettingsOutput() JobTemplateAccelerationSettingsOutput {
	return i.ToJobTemplateAccelerationSettingsOutputWithContext(context.Background())
}

func (i JobTemplateAccelerationSettingsArgs) ToJobTemplateAccelerationSettingsOutputWithContext(ctx context.Context) JobTemplateAccelerationSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobTemplateAccelerationSettingsOutput)
}

func (i JobTemplateAccelerationSettingsArgs) ToJobTemplateAccelerationSettingsPtrOutput() JobTemplateAccelerationSettingsPtrOutput {
	return i.ToJobTemplateAccelerationSettingsPtrOutputWithContext(context.Background())
}

func (i JobTemplateAccelerationSettingsArgs) ToJobTemplateAccelerationSettingsPtrOutputWithContext(ctx context.Context) JobTemplateAccelerationSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobTemplateAccelerationSettingsOutput).ToJobTemplateAccelerationSettingsPtrOutputWithContext(ctx)
}

// JobTemplateAccelerationSettingsPtrInput is an input type that accepts JobTemplateAccelerationSettingsArgs, JobTemplateAccelerationSettingsPtr and JobTemplateAccelerationSettingsPtrOutput values.
// You can construct a concrete instance of `JobTemplateAccelerationSettingsPtrInput` via:
//
//	        JobTemplateAccelerationSettingsArgs{...}
//
//	or:
//
//	        nil
type JobTemplateAccelerationSettingsPtrInput interface {
	pulumi.Input

	ToJobTemplateAccelerationSettingsPtrOutput() JobTemplateAccelerationSettingsPtrOutput
	ToJobTemplateAccelerationSettingsPtrOutputWithContext(context.Context) JobTemplateAccelerationSettingsPtrOutput
}

type jobTemplateAccelerationSettingsPtrType JobTemplateAccelerationSettingsArgs

func JobTemplateAccelerationSettingsPtr(v *JobTemplateAccelerationSettingsArgs) JobTemplateAccelerationSettingsPtrInput {
	return (*jobTemplateAccelerationSettingsPtrType)(v)
}

func (*jobTemplateAccelerationSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobTemplateAccelerationSettings)(nil)).Elem()
}

func (i *jobTemplateAccelerationSettingsPtrType) ToJobTemplateAccelerationSettingsPtrOutput() JobTemplateAccelerationSettingsPtrOutput {
	return i.ToJobTemplateAccelerationSettingsPtrOutputWithContext(context.Background())
}

func (i *jobTemplateAccelerationSettingsPtrType) ToJobTemplateAccelerationSettingsPtrOutputWithContext(ctx context.Context) JobTemplateAccelerationSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobTemplateAccelerationSettingsPtrOutput)
}

type JobTemplateAccelerationSettingsOutput struct{ *pulumi.OutputState }

func (JobTemplateAccelerationSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobTemplateAccelerationSettings)(nil)).Elem()
}

func (o JobTemplateAccelerationSettingsOutput) ToJobTemplateAccelerationSettingsOutput() JobTemplateAccelerationSettingsOutput {
	return o
}

func (o JobTemplateAccelerationSettingsOutput) ToJobTemplateAccelerationSettingsOutputWithContext(ctx context.Context) JobTemplateAccelerationSettingsOutput {
	return o
}

func (o JobTemplateAccelerationSettingsOutput) ToJobTemplateAccelerationSettingsPtrOutput() JobTemplateAccelerationSettingsPtrOutput {
	return o.ToJobTemplateAccelerationSettingsPtrOutputWithContext(context.Background())
}

func (o JobTemplateAccelerationSettingsOutput) ToJobTemplateAccelerationSettingsPtrOutputWithContext(ctx context.Context) JobTemplateAccelerationSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobTemplateAccelerationSettings) *JobTemplateAccelerationSettings {
		return &v
	}).(JobTemplateAccelerationSettingsPtrOutput)
}

func (o JobTemplateAccelerationSettingsOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v JobTemplateAccelerationSettings) string { return v.Mode }).(pulumi.StringOutput)
}

type JobTemplateAccelerationSettingsPtrOutput struct{ *pulumi.OutputState }

func (JobTemplateAccelerationSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobTemplateAccelerationSettings)(nil)).Elem()
}

func (o JobTemplateAccelerationSettingsPtrOutput) ToJobTemplateAccelerationSettingsPtrOutput() JobTemplateAccelerationSettingsPtrOutput {
	return o
}

func (o JobTemplateAccelerationSettingsPtrOutput) ToJobTemplateAccelerationSettingsPtrOutputWithContext(ctx context.Context) JobTemplateAccelerationSettingsPtrOutput {
	return o
}

func (o JobTemplateAccelerationSettingsPtrOutput) Elem() JobTemplateAccelerationSettingsOutput {
	return o.ApplyT(func(v *JobTemplateAccelerationSettings) JobTemplateAccelerationSettings {
		if v != nil {
			return *v
		}
		var ret JobTemplateAccelerationSettings
		return ret
	}).(JobTemplateAccelerationSettingsOutput)
}

func (o JobTemplateAccelerationSettingsPtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobTemplateAccelerationSettings) *string {
		if v == nil {
			return nil
		}
		return &v.Mode
	}).(pulumi.StringPtrOutput)
}

type JobTemplateHopDestination struct {
	Priority    *int    `pulumi:"priority"`
	Queue       *string `pulumi:"queue"`
	WaitMinutes *int    `pulumi:"waitMinutes"`
}

// JobTemplateHopDestinationInput is an input type that accepts JobTemplateHopDestinationArgs and JobTemplateHopDestinationOutput values.
// You can construct a concrete instance of `JobTemplateHopDestinationInput` via:
//
//	JobTemplateHopDestinationArgs{...}
type JobTemplateHopDestinationInput interface {
	pulumi.Input

	ToJobTemplateHopDestinationOutput() JobTemplateHopDestinationOutput
	ToJobTemplateHopDestinationOutputWithContext(context.Context) JobTemplateHopDestinationOutput
}

type JobTemplateHopDestinationArgs struct {
	Priority    pulumi.IntPtrInput    `pulumi:"priority"`
	Queue       pulumi.StringPtrInput `pulumi:"queue"`
	WaitMinutes pulumi.IntPtrInput    `pulumi:"waitMinutes"`
}

func (JobTemplateHopDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobTemplateHopDestination)(nil)).Elem()
}

func (i JobTemplateHopDestinationArgs) ToJobTemplateHopDestinationOutput() JobTemplateHopDestinationOutput {
	return i.ToJobTemplateHopDestinationOutputWithContext(context.Background())
}

func (i JobTemplateHopDestinationArgs) ToJobTemplateHopDestinationOutputWithContext(ctx context.Context) JobTemplateHopDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobTemplateHopDestinationOutput)
}

// JobTemplateHopDestinationArrayInput is an input type that accepts JobTemplateHopDestinationArray and JobTemplateHopDestinationArrayOutput values.
// You can construct a concrete instance of `JobTemplateHopDestinationArrayInput` via:
//
//	JobTemplateHopDestinationArray{ JobTemplateHopDestinationArgs{...} }
type JobTemplateHopDestinationArrayInput interface {
	pulumi.Input

	ToJobTemplateHopDestinationArrayOutput() JobTemplateHopDestinationArrayOutput
	ToJobTemplateHopDestinationArrayOutputWithContext(context.Context) JobTemplateHopDestinationArrayOutput
}

type JobTemplateHopDestinationArray []JobTemplateHopDestinationInput

func (JobTemplateHopDestinationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobTemplateHopDestination)(nil)).Elem()
}

func (i JobTemplateHopDestinationArray) ToJobTemplateHopDestinationArrayOutput() JobTemplateHopDestinationArrayOutput {
	return i.ToJobTemplateHopDestinationArrayOutputWithContext(context.Background())
}

func (i JobTemplateHopDestinationArray) ToJobTemplateHopDestinationArrayOutputWithContext(ctx context.Context) JobTemplateHopDestinationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobTemplateHopDestinationArrayOutput)
}

type JobTemplateHopDestinationOutput struct{ *pulumi.OutputState }

func (JobTemplateHopDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobTemplateHopDestination)(nil)).Elem()
}

func (o JobTemplateHopDestinationOutput) ToJobTemplateHopDestinationOutput() JobTemplateHopDestinationOutput {
	return o
}

func (o JobTemplateHopDestinationOutput) ToJobTemplateHopDestinationOutputWithContext(ctx context.Context) JobTemplateHopDestinationOutput {
	return o
}

func (o JobTemplateHopDestinationOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobTemplateHopDestination) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o JobTemplateHopDestinationOutput) Queue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobTemplateHopDestination) *string { return v.Queue }).(pulumi.StringPtrOutput)
}

func (o JobTemplateHopDestinationOutput) WaitMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobTemplateHopDestination) *int { return v.WaitMinutes }).(pulumi.IntPtrOutput)
}

type JobTemplateHopDestinationArrayOutput struct{ *pulumi.OutputState }

func (JobTemplateHopDestinationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobTemplateHopDestination)(nil)).Elem()
}

func (o JobTemplateHopDestinationArrayOutput) ToJobTemplateHopDestinationArrayOutput() JobTemplateHopDestinationArrayOutput {
	return o
}

func (o JobTemplateHopDestinationArrayOutput) ToJobTemplateHopDestinationArrayOutputWithContext(ctx context.Context) JobTemplateHopDestinationArrayOutput {
	return o
}

func (o JobTemplateHopDestinationArrayOutput) Index(i pulumi.IntInput) JobTemplateHopDestinationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JobTemplateHopDestination {
		return vs[0].([]JobTemplateHopDestination)[vs[1].(int)]
	}).(JobTemplateHopDestinationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*JobTemplateAccelerationSettingsInput)(nil)).Elem(), JobTemplateAccelerationSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobTemplateAccelerationSettingsPtrInput)(nil)).Elem(), JobTemplateAccelerationSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobTemplateHopDestinationInput)(nil)).Elem(), JobTemplateHopDestinationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobTemplateHopDestinationArrayInput)(nil)).Elem(), JobTemplateHopDestinationArray{})
	pulumi.RegisterOutputType(JobTemplateAccelerationSettingsOutput{})
	pulumi.RegisterOutputType(JobTemplateAccelerationSettingsPtrOutput{})
	pulumi.RegisterOutputType(JobTemplateHopDestinationOutput{})
	pulumi.RegisterOutputType(JobTemplateHopDestinationArrayOutput{})
}
