// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package supportapp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The severity level of a support case that a customer wants to get notified for.
type SlackChannelConfigurationNotifyOnCaseSeverity string

const (
	SlackChannelConfigurationNotifyOnCaseSeverityNone = SlackChannelConfigurationNotifyOnCaseSeverity("none")
	SlackChannelConfigurationNotifyOnCaseSeverityAll  = SlackChannelConfigurationNotifyOnCaseSeverity("all")
	SlackChannelConfigurationNotifyOnCaseSeverityHigh = SlackChannelConfigurationNotifyOnCaseSeverity("high")
)

func (SlackChannelConfigurationNotifyOnCaseSeverity) ElementType() reflect.Type {
	return reflect.TypeOf((*SlackChannelConfigurationNotifyOnCaseSeverity)(nil)).Elem()
}

func (e SlackChannelConfigurationNotifyOnCaseSeverity) ToSlackChannelConfigurationNotifyOnCaseSeverityOutput() SlackChannelConfigurationNotifyOnCaseSeverityOutput {
	return pulumi.ToOutput(e).(SlackChannelConfigurationNotifyOnCaseSeverityOutput)
}

func (e SlackChannelConfigurationNotifyOnCaseSeverity) ToSlackChannelConfigurationNotifyOnCaseSeverityOutputWithContext(ctx context.Context) SlackChannelConfigurationNotifyOnCaseSeverityOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SlackChannelConfigurationNotifyOnCaseSeverityOutput)
}

func (e SlackChannelConfigurationNotifyOnCaseSeverity) ToSlackChannelConfigurationNotifyOnCaseSeverityPtrOutput() SlackChannelConfigurationNotifyOnCaseSeverityPtrOutput {
	return e.ToSlackChannelConfigurationNotifyOnCaseSeverityPtrOutputWithContext(context.Background())
}

func (e SlackChannelConfigurationNotifyOnCaseSeverity) ToSlackChannelConfigurationNotifyOnCaseSeverityPtrOutputWithContext(ctx context.Context) SlackChannelConfigurationNotifyOnCaseSeverityPtrOutput {
	return SlackChannelConfigurationNotifyOnCaseSeverity(e).ToSlackChannelConfigurationNotifyOnCaseSeverityOutputWithContext(ctx).ToSlackChannelConfigurationNotifyOnCaseSeverityPtrOutputWithContext(ctx)
}

func (e SlackChannelConfigurationNotifyOnCaseSeverity) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SlackChannelConfigurationNotifyOnCaseSeverity) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SlackChannelConfigurationNotifyOnCaseSeverity) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SlackChannelConfigurationNotifyOnCaseSeverity) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SlackChannelConfigurationNotifyOnCaseSeverityOutput struct{ *pulumi.OutputState }

func (SlackChannelConfigurationNotifyOnCaseSeverityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SlackChannelConfigurationNotifyOnCaseSeverity)(nil)).Elem()
}

func (o SlackChannelConfigurationNotifyOnCaseSeverityOutput) ToSlackChannelConfigurationNotifyOnCaseSeverityOutput() SlackChannelConfigurationNotifyOnCaseSeverityOutput {
	return o
}

func (o SlackChannelConfigurationNotifyOnCaseSeverityOutput) ToSlackChannelConfigurationNotifyOnCaseSeverityOutputWithContext(ctx context.Context) SlackChannelConfigurationNotifyOnCaseSeverityOutput {
	return o
}

func (o SlackChannelConfigurationNotifyOnCaseSeverityOutput) ToSlackChannelConfigurationNotifyOnCaseSeverityPtrOutput() SlackChannelConfigurationNotifyOnCaseSeverityPtrOutput {
	return o.ToSlackChannelConfigurationNotifyOnCaseSeverityPtrOutputWithContext(context.Background())
}

func (o SlackChannelConfigurationNotifyOnCaseSeverityOutput) ToSlackChannelConfigurationNotifyOnCaseSeverityPtrOutputWithContext(ctx context.Context) SlackChannelConfigurationNotifyOnCaseSeverityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SlackChannelConfigurationNotifyOnCaseSeverity) *SlackChannelConfigurationNotifyOnCaseSeverity {
		return &v
	}).(SlackChannelConfigurationNotifyOnCaseSeverityPtrOutput)
}

func (o SlackChannelConfigurationNotifyOnCaseSeverityOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SlackChannelConfigurationNotifyOnCaseSeverityOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SlackChannelConfigurationNotifyOnCaseSeverity) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SlackChannelConfigurationNotifyOnCaseSeverityOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SlackChannelConfigurationNotifyOnCaseSeverityOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SlackChannelConfigurationNotifyOnCaseSeverity) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SlackChannelConfigurationNotifyOnCaseSeverityPtrOutput struct{ *pulumi.OutputState }

func (SlackChannelConfigurationNotifyOnCaseSeverityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SlackChannelConfigurationNotifyOnCaseSeverity)(nil)).Elem()
}

func (o SlackChannelConfigurationNotifyOnCaseSeverityPtrOutput) ToSlackChannelConfigurationNotifyOnCaseSeverityPtrOutput() SlackChannelConfigurationNotifyOnCaseSeverityPtrOutput {
	return o
}

func (o SlackChannelConfigurationNotifyOnCaseSeverityPtrOutput) ToSlackChannelConfigurationNotifyOnCaseSeverityPtrOutputWithContext(ctx context.Context) SlackChannelConfigurationNotifyOnCaseSeverityPtrOutput {
	return o
}

func (o SlackChannelConfigurationNotifyOnCaseSeverityPtrOutput) Elem() SlackChannelConfigurationNotifyOnCaseSeverityOutput {
	return o.ApplyT(func(v *SlackChannelConfigurationNotifyOnCaseSeverity) SlackChannelConfigurationNotifyOnCaseSeverity {
		if v != nil {
			return *v
		}
		var ret SlackChannelConfigurationNotifyOnCaseSeverity
		return ret
	}).(SlackChannelConfigurationNotifyOnCaseSeverityOutput)
}

func (o SlackChannelConfigurationNotifyOnCaseSeverityPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SlackChannelConfigurationNotifyOnCaseSeverityPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SlackChannelConfigurationNotifyOnCaseSeverity) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SlackChannelConfigurationNotifyOnCaseSeverityInput is an input type that accepts values of the SlackChannelConfigurationNotifyOnCaseSeverity enum
// A concrete instance of `SlackChannelConfigurationNotifyOnCaseSeverityInput` can be one of the following:
//
//	SlackChannelConfigurationNotifyOnCaseSeverityNone
//	SlackChannelConfigurationNotifyOnCaseSeverityAll
//	SlackChannelConfigurationNotifyOnCaseSeverityHigh
type SlackChannelConfigurationNotifyOnCaseSeverityInput interface {
	pulumi.Input

	ToSlackChannelConfigurationNotifyOnCaseSeverityOutput() SlackChannelConfigurationNotifyOnCaseSeverityOutput
	ToSlackChannelConfigurationNotifyOnCaseSeverityOutputWithContext(context.Context) SlackChannelConfigurationNotifyOnCaseSeverityOutput
}

var slackChannelConfigurationNotifyOnCaseSeverityPtrType = reflect.TypeOf((**SlackChannelConfigurationNotifyOnCaseSeverity)(nil)).Elem()

type SlackChannelConfigurationNotifyOnCaseSeverityPtrInput interface {
	pulumi.Input

	ToSlackChannelConfigurationNotifyOnCaseSeverityPtrOutput() SlackChannelConfigurationNotifyOnCaseSeverityPtrOutput
	ToSlackChannelConfigurationNotifyOnCaseSeverityPtrOutputWithContext(context.Context) SlackChannelConfigurationNotifyOnCaseSeverityPtrOutput
}

type slackChannelConfigurationNotifyOnCaseSeverityPtr string

func SlackChannelConfigurationNotifyOnCaseSeverityPtr(v string) SlackChannelConfigurationNotifyOnCaseSeverityPtrInput {
	return (*slackChannelConfigurationNotifyOnCaseSeverityPtr)(&v)
}

func (*slackChannelConfigurationNotifyOnCaseSeverityPtr) ElementType() reflect.Type {
	return slackChannelConfigurationNotifyOnCaseSeverityPtrType
}

func (in *slackChannelConfigurationNotifyOnCaseSeverityPtr) ToSlackChannelConfigurationNotifyOnCaseSeverityPtrOutput() SlackChannelConfigurationNotifyOnCaseSeverityPtrOutput {
	return pulumi.ToOutput(in).(SlackChannelConfigurationNotifyOnCaseSeverityPtrOutput)
}

func (in *slackChannelConfigurationNotifyOnCaseSeverityPtr) ToSlackChannelConfigurationNotifyOnCaseSeverityPtrOutputWithContext(ctx context.Context) SlackChannelConfigurationNotifyOnCaseSeverityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SlackChannelConfigurationNotifyOnCaseSeverityPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SlackChannelConfigurationNotifyOnCaseSeverityInput)(nil)).Elem(), SlackChannelConfigurationNotifyOnCaseSeverity("none"))
	pulumi.RegisterInputType(reflect.TypeOf((*SlackChannelConfigurationNotifyOnCaseSeverityPtrInput)(nil)).Elem(), SlackChannelConfigurationNotifyOnCaseSeverity("none"))
	pulumi.RegisterOutputType(SlackChannelConfigurationNotifyOnCaseSeverityOutput{})
	pulumi.RegisterOutputType(SlackChannelConfigurationNotifyOnCaseSeverityPtrOutput{})
}
