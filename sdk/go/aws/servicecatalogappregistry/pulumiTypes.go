// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalogappregistry

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type ApplicationTags struct {
}

// ApplicationTagsInput is an input type that accepts ApplicationTagsArgs and ApplicationTagsOutput values.
// You can construct a concrete instance of `ApplicationTagsInput` via:
//
//	ApplicationTagsArgs{...}
type ApplicationTagsInput interface {
	pulumi.Input

	ToApplicationTagsOutput() ApplicationTagsOutput
	ToApplicationTagsOutputWithContext(context.Context) ApplicationTagsOutput
}

type ApplicationTagsArgs struct {
}

func (ApplicationTagsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationTags)(nil)).Elem()
}

func (i ApplicationTagsArgs) ToApplicationTagsOutput() ApplicationTagsOutput {
	return i.ToApplicationTagsOutputWithContext(context.Background())
}

func (i ApplicationTagsArgs) ToApplicationTagsOutputWithContext(ctx context.Context) ApplicationTagsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationTagsOutput)
}

func (i ApplicationTagsArgs) ToApplicationTagsPtrOutput() ApplicationTagsPtrOutput {
	return i.ToApplicationTagsPtrOutputWithContext(context.Background())
}

func (i ApplicationTagsArgs) ToApplicationTagsPtrOutputWithContext(ctx context.Context) ApplicationTagsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationTagsOutput).ToApplicationTagsPtrOutputWithContext(ctx)
}

// ApplicationTagsPtrInput is an input type that accepts ApplicationTagsArgs, ApplicationTagsPtr and ApplicationTagsPtrOutput values.
// You can construct a concrete instance of `ApplicationTagsPtrInput` via:
//
//	        ApplicationTagsArgs{...}
//
//	or:
//
//	        nil
type ApplicationTagsPtrInput interface {
	pulumi.Input

	ToApplicationTagsPtrOutput() ApplicationTagsPtrOutput
	ToApplicationTagsPtrOutputWithContext(context.Context) ApplicationTagsPtrOutput
}

type applicationTagsPtrType ApplicationTagsArgs

func ApplicationTagsPtr(v *ApplicationTagsArgs) ApplicationTagsPtrInput {
	return (*applicationTagsPtrType)(v)
}

func (*applicationTagsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationTags)(nil)).Elem()
}

func (i *applicationTagsPtrType) ToApplicationTagsPtrOutput() ApplicationTagsPtrOutput {
	return i.ToApplicationTagsPtrOutputWithContext(context.Background())
}

func (i *applicationTagsPtrType) ToApplicationTagsPtrOutputWithContext(ctx context.Context) ApplicationTagsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationTagsPtrOutput)
}

type ApplicationTagsOutput struct{ *pulumi.OutputState }

func (ApplicationTagsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationTags)(nil)).Elem()
}

func (o ApplicationTagsOutput) ToApplicationTagsOutput() ApplicationTagsOutput {
	return o
}

func (o ApplicationTagsOutput) ToApplicationTagsOutputWithContext(ctx context.Context) ApplicationTagsOutput {
	return o
}

func (o ApplicationTagsOutput) ToApplicationTagsPtrOutput() ApplicationTagsPtrOutput {
	return o.ToApplicationTagsPtrOutputWithContext(context.Background())
}

func (o ApplicationTagsOutput) ToApplicationTagsPtrOutputWithContext(ctx context.Context) ApplicationTagsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationTags) *ApplicationTags {
		return &v
	}).(ApplicationTagsPtrOutput)
}

type ApplicationTagsPtrOutput struct{ *pulumi.OutputState }

func (ApplicationTagsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationTags)(nil)).Elem()
}

func (o ApplicationTagsPtrOutput) ToApplicationTagsPtrOutput() ApplicationTagsPtrOutput {
	return o
}

func (o ApplicationTagsPtrOutput) ToApplicationTagsPtrOutputWithContext(ctx context.Context) ApplicationTagsPtrOutput {
	return o
}

func (o ApplicationTagsPtrOutput) Elem() ApplicationTagsOutput {
	return o.ApplyT(func(v *ApplicationTags) ApplicationTags {
		if v != nil {
			return *v
		}
		var ret ApplicationTags
		return ret
	}).(ApplicationTagsOutput)
}

type AttributeGroupTags struct {
}

// AttributeGroupTagsInput is an input type that accepts AttributeGroupTagsArgs and AttributeGroupTagsOutput values.
// You can construct a concrete instance of `AttributeGroupTagsInput` via:
//
//	AttributeGroupTagsArgs{...}
type AttributeGroupTagsInput interface {
	pulumi.Input

	ToAttributeGroupTagsOutput() AttributeGroupTagsOutput
	ToAttributeGroupTagsOutputWithContext(context.Context) AttributeGroupTagsOutput
}

type AttributeGroupTagsArgs struct {
}

func (AttributeGroupTagsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AttributeGroupTags)(nil)).Elem()
}

func (i AttributeGroupTagsArgs) ToAttributeGroupTagsOutput() AttributeGroupTagsOutput {
	return i.ToAttributeGroupTagsOutputWithContext(context.Background())
}

func (i AttributeGroupTagsArgs) ToAttributeGroupTagsOutputWithContext(ctx context.Context) AttributeGroupTagsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttributeGroupTagsOutput)
}

func (i AttributeGroupTagsArgs) ToAttributeGroupTagsPtrOutput() AttributeGroupTagsPtrOutput {
	return i.ToAttributeGroupTagsPtrOutputWithContext(context.Background())
}

func (i AttributeGroupTagsArgs) ToAttributeGroupTagsPtrOutputWithContext(ctx context.Context) AttributeGroupTagsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttributeGroupTagsOutput).ToAttributeGroupTagsPtrOutputWithContext(ctx)
}

// AttributeGroupTagsPtrInput is an input type that accepts AttributeGroupTagsArgs, AttributeGroupTagsPtr and AttributeGroupTagsPtrOutput values.
// You can construct a concrete instance of `AttributeGroupTagsPtrInput` via:
//
//	        AttributeGroupTagsArgs{...}
//
//	or:
//
//	        nil
type AttributeGroupTagsPtrInput interface {
	pulumi.Input

	ToAttributeGroupTagsPtrOutput() AttributeGroupTagsPtrOutput
	ToAttributeGroupTagsPtrOutputWithContext(context.Context) AttributeGroupTagsPtrOutput
}

type attributeGroupTagsPtrType AttributeGroupTagsArgs

func AttributeGroupTagsPtr(v *AttributeGroupTagsArgs) AttributeGroupTagsPtrInput {
	return (*attributeGroupTagsPtrType)(v)
}

func (*attributeGroupTagsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AttributeGroupTags)(nil)).Elem()
}

func (i *attributeGroupTagsPtrType) ToAttributeGroupTagsPtrOutput() AttributeGroupTagsPtrOutput {
	return i.ToAttributeGroupTagsPtrOutputWithContext(context.Background())
}

func (i *attributeGroupTagsPtrType) ToAttributeGroupTagsPtrOutputWithContext(ctx context.Context) AttributeGroupTagsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttributeGroupTagsPtrOutput)
}

type AttributeGroupTagsOutput struct{ *pulumi.OutputState }

func (AttributeGroupTagsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AttributeGroupTags)(nil)).Elem()
}

func (o AttributeGroupTagsOutput) ToAttributeGroupTagsOutput() AttributeGroupTagsOutput {
	return o
}

func (o AttributeGroupTagsOutput) ToAttributeGroupTagsOutputWithContext(ctx context.Context) AttributeGroupTagsOutput {
	return o
}

func (o AttributeGroupTagsOutput) ToAttributeGroupTagsPtrOutput() AttributeGroupTagsPtrOutput {
	return o.ToAttributeGroupTagsPtrOutputWithContext(context.Background())
}

func (o AttributeGroupTagsOutput) ToAttributeGroupTagsPtrOutputWithContext(ctx context.Context) AttributeGroupTagsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AttributeGroupTags) *AttributeGroupTags {
		return &v
	}).(AttributeGroupTagsPtrOutput)
}

type AttributeGroupTagsPtrOutput struct{ *pulumi.OutputState }

func (AttributeGroupTagsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttributeGroupTags)(nil)).Elem()
}

func (o AttributeGroupTagsPtrOutput) ToAttributeGroupTagsPtrOutput() AttributeGroupTagsPtrOutput {
	return o
}

func (o AttributeGroupTagsPtrOutput) ToAttributeGroupTagsPtrOutputWithContext(ctx context.Context) AttributeGroupTagsPtrOutput {
	return o
}

func (o AttributeGroupTagsPtrOutput) Elem() AttributeGroupTagsOutput {
	return o.ApplyT(func(v *AttributeGroupTags) AttributeGroupTags {
		if v != nil {
			return *v
		}
		var ret AttributeGroupTags
		return ret
	}).(AttributeGroupTagsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationTagsInput)(nil)).Elem(), ApplicationTagsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationTagsPtrInput)(nil)).Elem(), ApplicationTagsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AttributeGroupTagsInput)(nil)).Elem(), AttributeGroupTagsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AttributeGroupTagsPtrInput)(nil)).Elem(), AttributeGroupTagsArgs{})
	pulumi.RegisterOutputType(ApplicationTagsOutput{})
	pulumi.RegisterOutputType(ApplicationTagsPtrOutput{})
	pulumi.RegisterOutputType(AttributeGroupTagsOutput{})
	pulumi.RegisterOutputType(AttributeGroupTagsPtrOutput{})
}
