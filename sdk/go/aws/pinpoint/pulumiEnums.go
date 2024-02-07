// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InAppTemplateAlignment string

const (
	InAppTemplateAlignmentLeft   = InAppTemplateAlignment("LEFT")
	InAppTemplateAlignmentCenter = InAppTemplateAlignment("CENTER")
	InAppTemplateAlignmentRight  = InAppTemplateAlignment("RIGHT")
)

func (InAppTemplateAlignment) ElementType() reflect.Type {
	return reflect.TypeOf((*InAppTemplateAlignment)(nil)).Elem()
}

func (e InAppTemplateAlignment) ToInAppTemplateAlignmentOutput() InAppTemplateAlignmentOutput {
	return pulumi.ToOutput(e).(InAppTemplateAlignmentOutput)
}

func (e InAppTemplateAlignment) ToInAppTemplateAlignmentOutputWithContext(ctx context.Context) InAppTemplateAlignmentOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InAppTemplateAlignmentOutput)
}

func (e InAppTemplateAlignment) ToInAppTemplateAlignmentPtrOutput() InAppTemplateAlignmentPtrOutput {
	return e.ToInAppTemplateAlignmentPtrOutputWithContext(context.Background())
}

func (e InAppTemplateAlignment) ToInAppTemplateAlignmentPtrOutputWithContext(ctx context.Context) InAppTemplateAlignmentPtrOutput {
	return InAppTemplateAlignment(e).ToInAppTemplateAlignmentOutputWithContext(ctx).ToInAppTemplateAlignmentPtrOutputWithContext(ctx)
}

func (e InAppTemplateAlignment) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InAppTemplateAlignment) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InAppTemplateAlignment) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InAppTemplateAlignment) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InAppTemplateAlignmentOutput struct{ *pulumi.OutputState }

func (InAppTemplateAlignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InAppTemplateAlignment)(nil)).Elem()
}

func (o InAppTemplateAlignmentOutput) ToInAppTemplateAlignmentOutput() InAppTemplateAlignmentOutput {
	return o
}

func (o InAppTemplateAlignmentOutput) ToInAppTemplateAlignmentOutputWithContext(ctx context.Context) InAppTemplateAlignmentOutput {
	return o
}

func (o InAppTemplateAlignmentOutput) ToInAppTemplateAlignmentPtrOutput() InAppTemplateAlignmentPtrOutput {
	return o.ToInAppTemplateAlignmentPtrOutputWithContext(context.Background())
}

func (o InAppTemplateAlignmentOutput) ToInAppTemplateAlignmentPtrOutputWithContext(ctx context.Context) InAppTemplateAlignmentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InAppTemplateAlignment) *InAppTemplateAlignment {
		return &v
	}).(InAppTemplateAlignmentPtrOutput)
}

func (o InAppTemplateAlignmentOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InAppTemplateAlignmentOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InAppTemplateAlignment) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InAppTemplateAlignmentOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InAppTemplateAlignmentOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InAppTemplateAlignment) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InAppTemplateAlignmentPtrOutput struct{ *pulumi.OutputState }

func (InAppTemplateAlignmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InAppTemplateAlignment)(nil)).Elem()
}

func (o InAppTemplateAlignmentPtrOutput) ToInAppTemplateAlignmentPtrOutput() InAppTemplateAlignmentPtrOutput {
	return o
}

func (o InAppTemplateAlignmentPtrOutput) ToInAppTemplateAlignmentPtrOutputWithContext(ctx context.Context) InAppTemplateAlignmentPtrOutput {
	return o
}

func (o InAppTemplateAlignmentPtrOutput) Elem() InAppTemplateAlignmentOutput {
	return o.ApplyT(func(v *InAppTemplateAlignment) InAppTemplateAlignment {
		if v != nil {
			return *v
		}
		var ret InAppTemplateAlignment
		return ret
	}).(InAppTemplateAlignmentOutput)
}

func (o InAppTemplateAlignmentPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InAppTemplateAlignmentPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *InAppTemplateAlignment) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// InAppTemplateAlignmentInput is an input type that accepts values of the InAppTemplateAlignment enum
// A concrete instance of `InAppTemplateAlignmentInput` can be one of the following:
//
//	InAppTemplateAlignmentLeft
//	InAppTemplateAlignmentCenter
//	InAppTemplateAlignmentRight
type InAppTemplateAlignmentInput interface {
	pulumi.Input

	ToInAppTemplateAlignmentOutput() InAppTemplateAlignmentOutput
	ToInAppTemplateAlignmentOutputWithContext(context.Context) InAppTemplateAlignmentOutput
}

var inAppTemplateAlignmentPtrType = reflect.TypeOf((**InAppTemplateAlignment)(nil)).Elem()

type InAppTemplateAlignmentPtrInput interface {
	pulumi.Input

	ToInAppTemplateAlignmentPtrOutput() InAppTemplateAlignmentPtrOutput
	ToInAppTemplateAlignmentPtrOutputWithContext(context.Context) InAppTemplateAlignmentPtrOutput
}

type inAppTemplateAlignmentPtr string

func InAppTemplateAlignmentPtr(v string) InAppTemplateAlignmentPtrInput {
	return (*inAppTemplateAlignmentPtr)(&v)
}

func (*inAppTemplateAlignmentPtr) ElementType() reflect.Type {
	return inAppTemplateAlignmentPtrType
}

func (in *inAppTemplateAlignmentPtr) ToInAppTemplateAlignmentPtrOutput() InAppTemplateAlignmentPtrOutput {
	return pulumi.ToOutput(in).(InAppTemplateAlignmentPtrOutput)
}

func (in *inAppTemplateAlignmentPtr) ToInAppTemplateAlignmentPtrOutputWithContext(ctx context.Context) InAppTemplateAlignmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InAppTemplateAlignmentPtrOutput)
}

type InAppTemplateButtonAction string

const (
	InAppTemplateButtonActionLink     = InAppTemplateButtonAction("LINK")
	InAppTemplateButtonActionDeepLink = InAppTemplateButtonAction("DEEP_LINK")
	InAppTemplateButtonActionClose    = InAppTemplateButtonAction("CLOSE")
)

func (InAppTemplateButtonAction) ElementType() reflect.Type {
	return reflect.TypeOf((*InAppTemplateButtonAction)(nil)).Elem()
}

func (e InAppTemplateButtonAction) ToInAppTemplateButtonActionOutput() InAppTemplateButtonActionOutput {
	return pulumi.ToOutput(e).(InAppTemplateButtonActionOutput)
}

func (e InAppTemplateButtonAction) ToInAppTemplateButtonActionOutputWithContext(ctx context.Context) InAppTemplateButtonActionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InAppTemplateButtonActionOutput)
}

func (e InAppTemplateButtonAction) ToInAppTemplateButtonActionPtrOutput() InAppTemplateButtonActionPtrOutput {
	return e.ToInAppTemplateButtonActionPtrOutputWithContext(context.Background())
}

func (e InAppTemplateButtonAction) ToInAppTemplateButtonActionPtrOutputWithContext(ctx context.Context) InAppTemplateButtonActionPtrOutput {
	return InAppTemplateButtonAction(e).ToInAppTemplateButtonActionOutputWithContext(ctx).ToInAppTemplateButtonActionPtrOutputWithContext(ctx)
}

func (e InAppTemplateButtonAction) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InAppTemplateButtonAction) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InAppTemplateButtonAction) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InAppTemplateButtonAction) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InAppTemplateButtonActionOutput struct{ *pulumi.OutputState }

func (InAppTemplateButtonActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InAppTemplateButtonAction)(nil)).Elem()
}

func (o InAppTemplateButtonActionOutput) ToInAppTemplateButtonActionOutput() InAppTemplateButtonActionOutput {
	return o
}

func (o InAppTemplateButtonActionOutput) ToInAppTemplateButtonActionOutputWithContext(ctx context.Context) InAppTemplateButtonActionOutput {
	return o
}

func (o InAppTemplateButtonActionOutput) ToInAppTemplateButtonActionPtrOutput() InAppTemplateButtonActionPtrOutput {
	return o.ToInAppTemplateButtonActionPtrOutputWithContext(context.Background())
}

func (o InAppTemplateButtonActionOutput) ToInAppTemplateButtonActionPtrOutputWithContext(ctx context.Context) InAppTemplateButtonActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InAppTemplateButtonAction) *InAppTemplateButtonAction {
		return &v
	}).(InAppTemplateButtonActionPtrOutput)
}

func (o InAppTemplateButtonActionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InAppTemplateButtonActionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InAppTemplateButtonAction) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InAppTemplateButtonActionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InAppTemplateButtonActionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InAppTemplateButtonAction) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InAppTemplateButtonActionPtrOutput struct{ *pulumi.OutputState }

func (InAppTemplateButtonActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InAppTemplateButtonAction)(nil)).Elem()
}

func (o InAppTemplateButtonActionPtrOutput) ToInAppTemplateButtonActionPtrOutput() InAppTemplateButtonActionPtrOutput {
	return o
}

func (o InAppTemplateButtonActionPtrOutput) ToInAppTemplateButtonActionPtrOutputWithContext(ctx context.Context) InAppTemplateButtonActionPtrOutput {
	return o
}

func (o InAppTemplateButtonActionPtrOutput) Elem() InAppTemplateButtonActionOutput {
	return o.ApplyT(func(v *InAppTemplateButtonAction) InAppTemplateButtonAction {
		if v != nil {
			return *v
		}
		var ret InAppTemplateButtonAction
		return ret
	}).(InAppTemplateButtonActionOutput)
}

func (o InAppTemplateButtonActionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InAppTemplateButtonActionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *InAppTemplateButtonAction) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// InAppTemplateButtonActionInput is an input type that accepts values of the InAppTemplateButtonAction enum
// A concrete instance of `InAppTemplateButtonActionInput` can be one of the following:
//
//	InAppTemplateButtonActionLink
//	InAppTemplateButtonActionDeepLink
//	InAppTemplateButtonActionClose
type InAppTemplateButtonActionInput interface {
	pulumi.Input

	ToInAppTemplateButtonActionOutput() InAppTemplateButtonActionOutput
	ToInAppTemplateButtonActionOutputWithContext(context.Context) InAppTemplateButtonActionOutput
}

var inAppTemplateButtonActionPtrType = reflect.TypeOf((**InAppTemplateButtonAction)(nil)).Elem()

type InAppTemplateButtonActionPtrInput interface {
	pulumi.Input

	ToInAppTemplateButtonActionPtrOutput() InAppTemplateButtonActionPtrOutput
	ToInAppTemplateButtonActionPtrOutputWithContext(context.Context) InAppTemplateButtonActionPtrOutput
}

type inAppTemplateButtonActionPtr string

func InAppTemplateButtonActionPtr(v string) InAppTemplateButtonActionPtrInput {
	return (*inAppTemplateButtonActionPtr)(&v)
}

func (*inAppTemplateButtonActionPtr) ElementType() reflect.Type {
	return inAppTemplateButtonActionPtrType
}

func (in *inAppTemplateButtonActionPtr) ToInAppTemplateButtonActionPtrOutput() InAppTemplateButtonActionPtrOutput {
	return pulumi.ToOutput(in).(InAppTemplateButtonActionPtrOutput)
}

func (in *inAppTemplateButtonActionPtr) ToInAppTemplateButtonActionPtrOutputWithContext(ctx context.Context) InAppTemplateButtonActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InAppTemplateButtonActionPtrOutput)
}

type InAppTemplateLayout string

const (
	InAppTemplateLayoutBottomBanner = InAppTemplateLayout("BOTTOM_BANNER")
	InAppTemplateLayoutTopBanner    = InAppTemplateLayout("TOP_BANNER")
	InAppTemplateLayoutOverlays     = InAppTemplateLayout("OVERLAYS")
	InAppTemplateLayoutMobileFeed   = InAppTemplateLayout("MOBILE_FEED")
	InAppTemplateLayoutMiddleBanner = InAppTemplateLayout("MIDDLE_BANNER")
	InAppTemplateLayoutCarousel     = InAppTemplateLayout("CAROUSEL")
)

func (InAppTemplateLayout) ElementType() reflect.Type {
	return reflect.TypeOf((*InAppTemplateLayout)(nil)).Elem()
}

func (e InAppTemplateLayout) ToInAppTemplateLayoutOutput() InAppTemplateLayoutOutput {
	return pulumi.ToOutput(e).(InAppTemplateLayoutOutput)
}

func (e InAppTemplateLayout) ToInAppTemplateLayoutOutputWithContext(ctx context.Context) InAppTemplateLayoutOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InAppTemplateLayoutOutput)
}

func (e InAppTemplateLayout) ToInAppTemplateLayoutPtrOutput() InAppTemplateLayoutPtrOutput {
	return e.ToInAppTemplateLayoutPtrOutputWithContext(context.Background())
}

func (e InAppTemplateLayout) ToInAppTemplateLayoutPtrOutputWithContext(ctx context.Context) InAppTemplateLayoutPtrOutput {
	return InAppTemplateLayout(e).ToInAppTemplateLayoutOutputWithContext(ctx).ToInAppTemplateLayoutPtrOutputWithContext(ctx)
}

func (e InAppTemplateLayout) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InAppTemplateLayout) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InAppTemplateLayout) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InAppTemplateLayout) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InAppTemplateLayoutOutput struct{ *pulumi.OutputState }

func (InAppTemplateLayoutOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InAppTemplateLayout)(nil)).Elem()
}

func (o InAppTemplateLayoutOutput) ToInAppTemplateLayoutOutput() InAppTemplateLayoutOutput {
	return o
}

func (o InAppTemplateLayoutOutput) ToInAppTemplateLayoutOutputWithContext(ctx context.Context) InAppTemplateLayoutOutput {
	return o
}

func (o InAppTemplateLayoutOutput) ToInAppTemplateLayoutPtrOutput() InAppTemplateLayoutPtrOutput {
	return o.ToInAppTemplateLayoutPtrOutputWithContext(context.Background())
}

func (o InAppTemplateLayoutOutput) ToInAppTemplateLayoutPtrOutputWithContext(ctx context.Context) InAppTemplateLayoutPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InAppTemplateLayout) *InAppTemplateLayout {
		return &v
	}).(InAppTemplateLayoutPtrOutput)
}

func (o InAppTemplateLayoutOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InAppTemplateLayoutOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InAppTemplateLayout) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InAppTemplateLayoutOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InAppTemplateLayoutOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InAppTemplateLayout) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InAppTemplateLayoutPtrOutput struct{ *pulumi.OutputState }

func (InAppTemplateLayoutPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InAppTemplateLayout)(nil)).Elem()
}

func (o InAppTemplateLayoutPtrOutput) ToInAppTemplateLayoutPtrOutput() InAppTemplateLayoutPtrOutput {
	return o
}

func (o InAppTemplateLayoutPtrOutput) ToInAppTemplateLayoutPtrOutputWithContext(ctx context.Context) InAppTemplateLayoutPtrOutput {
	return o
}

func (o InAppTemplateLayoutPtrOutput) Elem() InAppTemplateLayoutOutput {
	return o.ApplyT(func(v *InAppTemplateLayout) InAppTemplateLayout {
		if v != nil {
			return *v
		}
		var ret InAppTemplateLayout
		return ret
	}).(InAppTemplateLayoutOutput)
}

func (o InAppTemplateLayoutPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InAppTemplateLayoutPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *InAppTemplateLayout) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// InAppTemplateLayoutInput is an input type that accepts values of the InAppTemplateLayout enum
// A concrete instance of `InAppTemplateLayoutInput` can be one of the following:
//
//	InAppTemplateLayoutBottomBanner
//	InAppTemplateLayoutTopBanner
//	InAppTemplateLayoutOverlays
//	InAppTemplateLayoutMobileFeed
//	InAppTemplateLayoutMiddleBanner
//	InAppTemplateLayoutCarousel
type InAppTemplateLayoutInput interface {
	pulumi.Input

	ToInAppTemplateLayoutOutput() InAppTemplateLayoutOutput
	ToInAppTemplateLayoutOutputWithContext(context.Context) InAppTemplateLayoutOutput
}

var inAppTemplateLayoutPtrType = reflect.TypeOf((**InAppTemplateLayout)(nil)).Elem()

type InAppTemplateLayoutPtrInput interface {
	pulumi.Input

	ToInAppTemplateLayoutPtrOutput() InAppTemplateLayoutPtrOutput
	ToInAppTemplateLayoutPtrOutputWithContext(context.Context) InAppTemplateLayoutPtrOutput
}

type inAppTemplateLayoutPtr string

func InAppTemplateLayoutPtr(v string) InAppTemplateLayoutPtrInput {
	return (*inAppTemplateLayoutPtr)(&v)
}

func (*inAppTemplateLayoutPtr) ElementType() reflect.Type {
	return inAppTemplateLayoutPtrType
}

func (in *inAppTemplateLayoutPtr) ToInAppTemplateLayoutPtrOutput() InAppTemplateLayoutPtrOutput {
	return pulumi.ToOutput(in).(InAppTemplateLayoutPtrOutput)
}

func (in *inAppTemplateLayoutPtr) ToInAppTemplateLayoutPtrOutputWithContext(ctx context.Context) InAppTemplateLayoutPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InAppTemplateLayoutPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InAppTemplateAlignmentInput)(nil)).Elem(), InAppTemplateAlignment("LEFT"))
	pulumi.RegisterInputType(reflect.TypeOf((*InAppTemplateAlignmentPtrInput)(nil)).Elem(), InAppTemplateAlignment("LEFT"))
	pulumi.RegisterInputType(reflect.TypeOf((*InAppTemplateButtonActionInput)(nil)).Elem(), InAppTemplateButtonAction("LINK"))
	pulumi.RegisterInputType(reflect.TypeOf((*InAppTemplateButtonActionPtrInput)(nil)).Elem(), InAppTemplateButtonAction("LINK"))
	pulumi.RegisterInputType(reflect.TypeOf((*InAppTemplateLayoutInput)(nil)).Elem(), InAppTemplateLayout("BOTTOM_BANNER"))
	pulumi.RegisterInputType(reflect.TypeOf((*InAppTemplateLayoutPtrInput)(nil)).Elem(), InAppTemplateLayout("BOTTOM_BANNER"))
	pulumi.RegisterOutputType(InAppTemplateAlignmentOutput{})
	pulumi.RegisterOutputType(InAppTemplateAlignmentPtrOutput{})
	pulumi.RegisterOutputType(InAppTemplateButtonActionOutput{})
	pulumi.RegisterOutputType(InAppTemplateButtonActionPtrOutput{})
	pulumi.RegisterOutputType(InAppTemplateLayoutOutput{})
	pulumi.RegisterOutputType(InAppTemplateLayoutPtrOutput{})
}
