// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devicefarm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The rule's stringified attribute.
type DevicePoolRuleAttribute string

const (
	DevicePoolRuleAttributeArn                 = DevicePoolRuleAttribute("ARN")
	DevicePoolRuleAttributePlatform            = DevicePoolRuleAttribute("PLATFORM")
	DevicePoolRuleAttributeFormFactor          = DevicePoolRuleAttribute("FORM_FACTOR")
	DevicePoolRuleAttributeManufacturer        = DevicePoolRuleAttribute("MANUFACTURER")
	DevicePoolRuleAttributeRemoteAccessEnabled = DevicePoolRuleAttribute("REMOTE_ACCESS_ENABLED")
	DevicePoolRuleAttributeRemoteDebugEnabled  = DevicePoolRuleAttribute("REMOTE_DEBUG_ENABLED")
	DevicePoolRuleAttributeAppiumVersion       = DevicePoolRuleAttribute("APPIUM_VERSION")
	DevicePoolRuleAttributeInstanceArn         = DevicePoolRuleAttribute("INSTANCE_ARN")
	DevicePoolRuleAttributeInstanceLabels      = DevicePoolRuleAttribute("INSTANCE_LABELS")
	DevicePoolRuleAttributeFleetType           = DevicePoolRuleAttribute("FLEET_TYPE")
	DevicePoolRuleAttributeOsVersion           = DevicePoolRuleAttribute("OS_VERSION")
	DevicePoolRuleAttributeModel               = DevicePoolRuleAttribute("MODEL")
	DevicePoolRuleAttributeAvailability        = DevicePoolRuleAttribute("AVAILABILITY")
)

func (DevicePoolRuleAttribute) ElementType() reflect.Type {
	return reflect.TypeOf((*DevicePoolRuleAttribute)(nil)).Elem()
}

func (e DevicePoolRuleAttribute) ToDevicePoolRuleAttributeOutput() DevicePoolRuleAttributeOutput {
	return pulumi.ToOutput(e).(DevicePoolRuleAttributeOutput)
}

func (e DevicePoolRuleAttribute) ToDevicePoolRuleAttributeOutputWithContext(ctx context.Context) DevicePoolRuleAttributeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DevicePoolRuleAttributeOutput)
}

func (e DevicePoolRuleAttribute) ToDevicePoolRuleAttributePtrOutput() DevicePoolRuleAttributePtrOutput {
	return e.ToDevicePoolRuleAttributePtrOutputWithContext(context.Background())
}

func (e DevicePoolRuleAttribute) ToDevicePoolRuleAttributePtrOutputWithContext(ctx context.Context) DevicePoolRuleAttributePtrOutput {
	return DevicePoolRuleAttribute(e).ToDevicePoolRuleAttributeOutputWithContext(ctx).ToDevicePoolRuleAttributePtrOutputWithContext(ctx)
}

func (e DevicePoolRuleAttribute) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DevicePoolRuleAttribute) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DevicePoolRuleAttribute) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DevicePoolRuleAttribute) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DevicePoolRuleAttributeOutput struct{ *pulumi.OutputState }

func (DevicePoolRuleAttributeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DevicePoolRuleAttribute)(nil)).Elem()
}

func (o DevicePoolRuleAttributeOutput) ToDevicePoolRuleAttributeOutput() DevicePoolRuleAttributeOutput {
	return o
}

func (o DevicePoolRuleAttributeOutput) ToDevicePoolRuleAttributeOutputWithContext(ctx context.Context) DevicePoolRuleAttributeOutput {
	return o
}

func (o DevicePoolRuleAttributeOutput) ToDevicePoolRuleAttributePtrOutput() DevicePoolRuleAttributePtrOutput {
	return o.ToDevicePoolRuleAttributePtrOutputWithContext(context.Background())
}

func (o DevicePoolRuleAttributeOutput) ToDevicePoolRuleAttributePtrOutputWithContext(ctx context.Context) DevicePoolRuleAttributePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DevicePoolRuleAttribute) *DevicePoolRuleAttribute {
		return &v
	}).(DevicePoolRuleAttributePtrOutput)
}

func (o DevicePoolRuleAttributeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DevicePoolRuleAttributeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DevicePoolRuleAttribute) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DevicePoolRuleAttributeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DevicePoolRuleAttributeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DevicePoolRuleAttribute) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DevicePoolRuleAttributePtrOutput struct{ *pulumi.OutputState }

func (DevicePoolRuleAttributePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DevicePoolRuleAttribute)(nil)).Elem()
}

func (o DevicePoolRuleAttributePtrOutput) ToDevicePoolRuleAttributePtrOutput() DevicePoolRuleAttributePtrOutput {
	return o
}

func (o DevicePoolRuleAttributePtrOutput) ToDevicePoolRuleAttributePtrOutputWithContext(ctx context.Context) DevicePoolRuleAttributePtrOutput {
	return o
}

func (o DevicePoolRuleAttributePtrOutput) Elem() DevicePoolRuleAttributeOutput {
	return o.ApplyT(func(v *DevicePoolRuleAttribute) DevicePoolRuleAttribute {
		if v != nil {
			return *v
		}
		var ret DevicePoolRuleAttribute
		return ret
	}).(DevicePoolRuleAttributeOutput)
}

func (o DevicePoolRuleAttributePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DevicePoolRuleAttributePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DevicePoolRuleAttribute) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// DevicePoolRuleAttributeInput is an input type that accepts values of the DevicePoolRuleAttribute enum
// A concrete instance of `DevicePoolRuleAttributeInput` can be one of the following:
//
//	DevicePoolRuleAttributeArn
//	DevicePoolRuleAttributePlatform
//	DevicePoolRuleAttributeFormFactor
//	DevicePoolRuleAttributeManufacturer
//	DevicePoolRuleAttributeRemoteAccessEnabled
//	DevicePoolRuleAttributeRemoteDebugEnabled
//	DevicePoolRuleAttributeAppiumVersion
//	DevicePoolRuleAttributeInstanceArn
//	DevicePoolRuleAttributeInstanceLabels
//	DevicePoolRuleAttributeFleetType
//	DevicePoolRuleAttributeOsVersion
//	DevicePoolRuleAttributeModel
//	DevicePoolRuleAttributeAvailability
type DevicePoolRuleAttributeInput interface {
	pulumi.Input

	ToDevicePoolRuleAttributeOutput() DevicePoolRuleAttributeOutput
	ToDevicePoolRuleAttributeOutputWithContext(context.Context) DevicePoolRuleAttributeOutput
}

var devicePoolRuleAttributePtrType = reflect.TypeOf((**DevicePoolRuleAttribute)(nil)).Elem()

type DevicePoolRuleAttributePtrInput interface {
	pulumi.Input

	ToDevicePoolRuleAttributePtrOutput() DevicePoolRuleAttributePtrOutput
	ToDevicePoolRuleAttributePtrOutputWithContext(context.Context) DevicePoolRuleAttributePtrOutput
}

type devicePoolRuleAttributePtr string

func DevicePoolRuleAttributePtr(v string) DevicePoolRuleAttributePtrInput {
	return (*devicePoolRuleAttributePtr)(&v)
}

func (*devicePoolRuleAttributePtr) ElementType() reflect.Type {
	return devicePoolRuleAttributePtrType
}

func (in *devicePoolRuleAttributePtr) ToDevicePoolRuleAttributePtrOutput() DevicePoolRuleAttributePtrOutput {
	return pulumi.ToOutput(in).(DevicePoolRuleAttributePtrOutput)
}

func (in *devicePoolRuleAttributePtr) ToDevicePoolRuleAttributePtrOutputWithContext(ctx context.Context) DevicePoolRuleAttributePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DevicePoolRuleAttributePtrOutput)
}

// Specifies how Device Farm compares the rule's attribute to the value.
type DevicePoolRuleOperator string

const (
	DevicePoolRuleOperatorEquals              = DevicePoolRuleOperator("EQUALS")
	DevicePoolRuleOperatorLessThan            = DevicePoolRuleOperator("LESS_THAN")
	DevicePoolRuleOperatorLessThanOrEquals    = DevicePoolRuleOperator("LESS_THAN_OR_EQUALS")
	DevicePoolRuleOperatorGreaterThan         = DevicePoolRuleOperator("GREATER_THAN")
	DevicePoolRuleOperatorGreaterThanOrEquals = DevicePoolRuleOperator("GREATER_THAN_OR_EQUALS")
	DevicePoolRuleOperatorIn                  = DevicePoolRuleOperator("IN")
	DevicePoolRuleOperatorNotIn               = DevicePoolRuleOperator("NOT_IN")
	DevicePoolRuleOperatorContains            = DevicePoolRuleOperator("CONTAINS")
)

func (DevicePoolRuleOperator) ElementType() reflect.Type {
	return reflect.TypeOf((*DevicePoolRuleOperator)(nil)).Elem()
}

func (e DevicePoolRuleOperator) ToDevicePoolRuleOperatorOutput() DevicePoolRuleOperatorOutput {
	return pulumi.ToOutput(e).(DevicePoolRuleOperatorOutput)
}

func (e DevicePoolRuleOperator) ToDevicePoolRuleOperatorOutputWithContext(ctx context.Context) DevicePoolRuleOperatorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DevicePoolRuleOperatorOutput)
}

func (e DevicePoolRuleOperator) ToDevicePoolRuleOperatorPtrOutput() DevicePoolRuleOperatorPtrOutput {
	return e.ToDevicePoolRuleOperatorPtrOutputWithContext(context.Background())
}

func (e DevicePoolRuleOperator) ToDevicePoolRuleOperatorPtrOutputWithContext(ctx context.Context) DevicePoolRuleOperatorPtrOutput {
	return DevicePoolRuleOperator(e).ToDevicePoolRuleOperatorOutputWithContext(ctx).ToDevicePoolRuleOperatorPtrOutputWithContext(ctx)
}

func (e DevicePoolRuleOperator) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DevicePoolRuleOperator) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DevicePoolRuleOperator) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DevicePoolRuleOperator) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DevicePoolRuleOperatorOutput struct{ *pulumi.OutputState }

func (DevicePoolRuleOperatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DevicePoolRuleOperator)(nil)).Elem()
}

func (o DevicePoolRuleOperatorOutput) ToDevicePoolRuleOperatorOutput() DevicePoolRuleOperatorOutput {
	return o
}

func (o DevicePoolRuleOperatorOutput) ToDevicePoolRuleOperatorOutputWithContext(ctx context.Context) DevicePoolRuleOperatorOutput {
	return o
}

func (o DevicePoolRuleOperatorOutput) ToDevicePoolRuleOperatorPtrOutput() DevicePoolRuleOperatorPtrOutput {
	return o.ToDevicePoolRuleOperatorPtrOutputWithContext(context.Background())
}

func (o DevicePoolRuleOperatorOutput) ToDevicePoolRuleOperatorPtrOutputWithContext(ctx context.Context) DevicePoolRuleOperatorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DevicePoolRuleOperator) *DevicePoolRuleOperator {
		return &v
	}).(DevicePoolRuleOperatorPtrOutput)
}

func (o DevicePoolRuleOperatorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DevicePoolRuleOperatorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DevicePoolRuleOperator) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DevicePoolRuleOperatorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DevicePoolRuleOperatorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DevicePoolRuleOperator) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DevicePoolRuleOperatorPtrOutput struct{ *pulumi.OutputState }

func (DevicePoolRuleOperatorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DevicePoolRuleOperator)(nil)).Elem()
}

func (o DevicePoolRuleOperatorPtrOutput) ToDevicePoolRuleOperatorPtrOutput() DevicePoolRuleOperatorPtrOutput {
	return o
}

func (o DevicePoolRuleOperatorPtrOutput) ToDevicePoolRuleOperatorPtrOutputWithContext(ctx context.Context) DevicePoolRuleOperatorPtrOutput {
	return o
}

func (o DevicePoolRuleOperatorPtrOutput) Elem() DevicePoolRuleOperatorOutput {
	return o.ApplyT(func(v *DevicePoolRuleOperator) DevicePoolRuleOperator {
		if v != nil {
			return *v
		}
		var ret DevicePoolRuleOperator
		return ret
	}).(DevicePoolRuleOperatorOutput)
}

func (o DevicePoolRuleOperatorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DevicePoolRuleOperatorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DevicePoolRuleOperator) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// DevicePoolRuleOperatorInput is an input type that accepts values of the DevicePoolRuleOperator enum
// A concrete instance of `DevicePoolRuleOperatorInput` can be one of the following:
//
//	DevicePoolRuleOperatorEquals
//	DevicePoolRuleOperatorLessThan
//	DevicePoolRuleOperatorLessThanOrEquals
//	DevicePoolRuleOperatorGreaterThan
//	DevicePoolRuleOperatorGreaterThanOrEquals
//	DevicePoolRuleOperatorIn
//	DevicePoolRuleOperatorNotIn
//	DevicePoolRuleOperatorContains
type DevicePoolRuleOperatorInput interface {
	pulumi.Input

	ToDevicePoolRuleOperatorOutput() DevicePoolRuleOperatorOutput
	ToDevicePoolRuleOperatorOutputWithContext(context.Context) DevicePoolRuleOperatorOutput
}

var devicePoolRuleOperatorPtrType = reflect.TypeOf((**DevicePoolRuleOperator)(nil)).Elem()

type DevicePoolRuleOperatorPtrInput interface {
	pulumi.Input

	ToDevicePoolRuleOperatorPtrOutput() DevicePoolRuleOperatorPtrOutput
	ToDevicePoolRuleOperatorPtrOutputWithContext(context.Context) DevicePoolRuleOperatorPtrOutput
}

type devicePoolRuleOperatorPtr string

func DevicePoolRuleOperatorPtr(v string) DevicePoolRuleOperatorPtrInput {
	return (*devicePoolRuleOperatorPtr)(&v)
}

func (*devicePoolRuleOperatorPtr) ElementType() reflect.Type {
	return devicePoolRuleOperatorPtrType
}

func (in *devicePoolRuleOperatorPtr) ToDevicePoolRuleOperatorPtrOutput() DevicePoolRuleOperatorPtrOutput {
	return pulumi.ToOutput(in).(DevicePoolRuleOperatorPtrOutput)
}

func (in *devicePoolRuleOperatorPtr) ToDevicePoolRuleOperatorPtrOutputWithContext(ctx context.Context) DevicePoolRuleOperatorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DevicePoolRuleOperatorPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DevicePoolRuleAttributeInput)(nil)).Elem(), DevicePoolRuleAttribute("ARN"))
	pulumi.RegisterInputType(reflect.TypeOf((*DevicePoolRuleAttributePtrInput)(nil)).Elem(), DevicePoolRuleAttribute("ARN"))
	pulumi.RegisterInputType(reflect.TypeOf((*DevicePoolRuleOperatorInput)(nil)).Elem(), DevicePoolRuleOperator("EQUALS"))
	pulumi.RegisterInputType(reflect.TypeOf((*DevicePoolRuleOperatorPtrInput)(nil)).Elem(), DevicePoolRuleOperator("EQUALS"))
	pulumi.RegisterOutputType(DevicePoolRuleAttributeOutput{})
	pulumi.RegisterOutputType(DevicePoolRuleAttributePtrOutput{})
	pulumi.RegisterOutputType(DevicePoolRuleOperatorOutput{})
	pulumi.RegisterOutputType(DevicePoolRuleOperatorPtrOutput{})
}
