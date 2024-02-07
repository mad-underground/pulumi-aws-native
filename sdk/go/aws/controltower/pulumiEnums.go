// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package controltower

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LandingZoneDriftStatus string

const (
	LandingZoneDriftStatusDrifted = LandingZoneDriftStatus("DRIFTED")
	LandingZoneDriftStatusInSync  = LandingZoneDriftStatus("IN_SYNC")
)

type LandingZoneDriftStatusOutput struct{ *pulumi.OutputState }

func (LandingZoneDriftStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LandingZoneDriftStatus)(nil)).Elem()
}

func (o LandingZoneDriftStatusOutput) ToLandingZoneDriftStatusOutput() LandingZoneDriftStatusOutput {
	return o
}

func (o LandingZoneDriftStatusOutput) ToLandingZoneDriftStatusOutputWithContext(ctx context.Context) LandingZoneDriftStatusOutput {
	return o
}

func (o LandingZoneDriftStatusOutput) ToLandingZoneDriftStatusPtrOutput() LandingZoneDriftStatusPtrOutput {
	return o.ToLandingZoneDriftStatusPtrOutputWithContext(context.Background())
}

func (o LandingZoneDriftStatusOutput) ToLandingZoneDriftStatusPtrOutputWithContext(ctx context.Context) LandingZoneDriftStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LandingZoneDriftStatus) *LandingZoneDriftStatus {
		return &v
	}).(LandingZoneDriftStatusPtrOutput)
}

func (o LandingZoneDriftStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LandingZoneDriftStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LandingZoneDriftStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LandingZoneDriftStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LandingZoneDriftStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LandingZoneDriftStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LandingZoneDriftStatusPtrOutput struct{ *pulumi.OutputState }

func (LandingZoneDriftStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LandingZoneDriftStatus)(nil)).Elem()
}

func (o LandingZoneDriftStatusPtrOutput) ToLandingZoneDriftStatusPtrOutput() LandingZoneDriftStatusPtrOutput {
	return o
}

func (o LandingZoneDriftStatusPtrOutput) ToLandingZoneDriftStatusPtrOutputWithContext(ctx context.Context) LandingZoneDriftStatusPtrOutput {
	return o
}

func (o LandingZoneDriftStatusPtrOutput) Elem() LandingZoneDriftStatusOutput {
	return o.ApplyT(func(v *LandingZoneDriftStatus) LandingZoneDriftStatus {
		if v != nil {
			return *v
		}
		var ret LandingZoneDriftStatus
		return ret
	}).(LandingZoneDriftStatusOutput)
}

func (o LandingZoneDriftStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LandingZoneDriftStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LandingZoneDriftStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LandingZoneStatus string

const (
	LandingZoneStatusActive     = LandingZoneStatus("ACTIVE")
	LandingZoneStatusProcessing = LandingZoneStatus("PROCESSING")
	LandingZoneStatusFailed     = LandingZoneStatus("FAILED")
)

type LandingZoneStatusOutput struct{ *pulumi.OutputState }

func (LandingZoneStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LandingZoneStatus)(nil)).Elem()
}

func (o LandingZoneStatusOutput) ToLandingZoneStatusOutput() LandingZoneStatusOutput {
	return o
}

func (o LandingZoneStatusOutput) ToLandingZoneStatusOutputWithContext(ctx context.Context) LandingZoneStatusOutput {
	return o
}

func (o LandingZoneStatusOutput) ToLandingZoneStatusPtrOutput() LandingZoneStatusPtrOutput {
	return o.ToLandingZoneStatusPtrOutputWithContext(context.Background())
}

func (o LandingZoneStatusOutput) ToLandingZoneStatusPtrOutputWithContext(ctx context.Context) LandingZoneStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LandingZoneStatus) *LandingZoneStatus {
		return &v
	}).(LandingZoneStatusPtrOutput)
}

func (o LandingZoneStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LandingZoneStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LandingZoneStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LandingZoneStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LandingZoneStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LandingZoneStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LandingZoneStatusPtrOutput struct{ *pulumi.OutputState }

func (LandingZoneStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LandingZoneStatus)(nil)).Elem()
}

func (o LandingZoneStatusPtrOutput) ToLandingZoneStatusPtrOutput() LandingZoneStatusPtrOutput {
	return o
}

func (o LandingZoneStatusPtrOutput) ToLandingZoneStatusPtrOutputWithContext(ctx context.Context) LandingZoneStatusPtrOutput {
	return o
}

func (o LandingZoneStatusPtrOutput) Elem() LandingZoneStatusOutput {
	return o.ApplyT(func(v *LandingZoneStatus) LandingZoneStatus {
		if v != nil {
			return *v
		}
		var ret LandingZoneStatus
		return ret
	}).(LandingZoneStatusOutput)
}

func (o LandingZoneStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LandingZoneStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LandingZoneStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LandingZoneDriftStatusOutput{})
	pulumi.RegisterOutputType(LandingZoneDriftStatusPtrOutput{})
	pulumi.RegisterOutputType(LandingZoneStatusOutput{})
	pulumi.RegisterOutputType(LandingZoneStatusPtrOutput{})
}
