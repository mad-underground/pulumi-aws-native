// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-cloudformation/sdk/go/cloudformation"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type AliasAttributes struct {
}

// AliasAttributesInput is an input type that accepts AliasAttributesArgs and AliasAttributesOutput values.
// You can construct a concrete instance of `AliasAttributesInput` via:
//
//          AliasAttributesArgs{...}
type AliasAttributesInput interface {
	pulumi.Input

	ToAliasAttributesOutput() AliasAttributesOutput
	ToAliasAttributesOutputWithContext(context.Context) AliasAttributesOutput
}

type AliasAttributesArgs struct {
}

func (AliasAttributesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AliasAttributes)(nil)).Elem()
}

func (i AliasAttributesArgs) ToAliasAttributesOutput() AliasAttributesOutput {
	return i.ToAliasAttributesOutputWithContext(context.Background())
}

func (i AliasAttributesArgs) ToAliasAttributesOutputWithContext(ctx context.Context) AliasAttributesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasAttributesOutput)
}

func (i AliasAttributesArgs) ToAliasAttributesPtrOutput() AliasAttributesPtrOutput {
	return i.ToAliasAttributesPtrOutputWithContext(context.Background())
}

func (i AliasAttributesArgs) ToAliasAttributesPtrOutputWithContext(ctx context.Context) AliasAttributesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasAttributesOutput).ToAliasAttributesPtrOutputWithContext(ctx)
}

// AliasAttributesPtrInput is an input type that accepts AliasAttributesArgs, AliasAttributesPtr and AliasAttributesPtrOutput values.
// You can construct a concrete instance of `AliasAttributesPtrInput` via:
//
//          AliasAttributesArgs{...}
//
//  or:
//
//          nil
type AliasAttributesPtrInput interface {
	pulumi.Input

	ToAliasAttributesPtrOutput() AliasAttributesPtrOutput
	ToAliasAttributesPtrOutputWithContext(context.Context) AliasAttributesPtrOutput
}

type aliasAttributesPtrType AliasAttributesArgs

func AliasAttributesPtr(v *AliasAttributesArgs) AliasAttributesPtrInput {
	return (*aliasAttributesPtrType)(v)
}

func (*aliasAttributesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AliasAttributes)(nil)).Elem()
}

func (i *aliasAttributesPtrType) ToAliasAttributesPtrOutput() AliasAttributesPtrOutput {
	return i.ToAliasAttributesPtrOutputWithContext(context.Background())
}

func (i *aliasAttributesPtrType) ToAliasAttributesPtrOutputWithContext(ctx context.Context) AliasAttributesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasAttributesPtrOutput)
}

type AliasAttributesOutput struct{ *pulumi.OutputState }

func (AliasAttributesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AliasAttributes)(nil)).Elem()
}

func (o AliasAttributesOutput) ToAliasAttributesOutput() AliasAttributesOutput {
	return o
}

func (o AliasAttributesOutput) ToAliasAttributesOutputWithContext(ctx context.Context) AliasAttributesOutput {
	return o
}

func (o AliasAttributesOutput) ToAliasAttributesPtrOutput() AliasAttributesPtrOutput {
	return o.ToAliasAttributesPtrOutputWithContext(context.Background())
}

func (o AliasAttributesOutput) ToAliasAttributesPtrOutputWithContext(ctx context.Context) AliasAttributesPtrOutput {
	return o.ApplyT(func(v AliasAttributes) *AliasAttributes {
		return &v
	}).(AliasAttributesPtrOutput)
}

type AliasAttributesPtrOutput struct{ *pulumi.OutputState }

func (AliasAttributesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AliasAttributes)(nil)).Elem()
}

func (o AliasAttributesPtrOutput) ToAliasAttributesPtrOutput() AliasAttributesPtrOutput {
	return o
}

func (o AliasAttributesPtrOutput) ToAliasAttributesPtrOutputWithContext(ctx context.Context) AliasAttributesPtrOutput {
	return o
}

func (o AliasAttributesPtrOutput) Elem() AliasAttributesOutput {
	return o.ApplyT(func(v *AliasAttributes) AliasAttributes { return *v }).(AliasAttributesOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html
type AliasProperties struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html#cfn-kms-alias-aliasname
	AliasName string `pulumi:"AliasName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html#cfn-kms-alias-targetkeyid
	TargetKeyId string `pulumi:"TargetKeyId"`
}

// AliasPropertiesInput is an input type that accepts AliasPropertiesArgs and AliasPropertiesOutput values.
// You can construct a concrete instance of `AliasPropertiesInput` via:
//
//          AliasPropertiesArgs{...}
type AliasPropertiesInput interface {
	pulumi.Input

	ToAliasPropertiesOutput() AliasPropertiesOutput
	ToAliasPropertiesOutputWithContext(context.Context) AliasPropertiesOutput
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html
type AliasPropertiesArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html#cfn-kms-alias-aliasname
	AliasName pulumi.StringInput `pulumi:"AliasName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html#cfn-kms-alias-targetkeyid
	TargetKeyId pulumi.StringInput `pulumi:"TargetKeyId"`
}

func (AliasPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AliasProperties)(nil)).Elem()
}

func (i AliasPropertiesArgs) ToAliasPropertiesOutput() AliasPropertiesOutput {
	return i.ToAliasPropertiesOutputWithContext(context.Background())
}

func (i AliasPropertiesArgs) ToAliasPropertiesOutputWithContext(ctx context.Context) AliasPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasPropertiesOutput)
}

func (i AliasPropertiesArgs) ToAliasPropertiesPtrOutput() AliasPropertiesPtrOutput {
	return i.ToAliasPropertiesPtrOutputWithContext(context.Background())
}

func (i AliasPropertiesArgs) ToAliasPropertiesPtrOutputWithContext(ctx context.Context) AliasPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasPropertiesOutput).ToAliasPropertiesPtrOutputWithContext(ctx)
}

// AliasPropertiesPtrInput is an input type that accepts AliasPropertiesArgs, AliasPropertiesPtr and AliasPropertiesPtrOutput values.
// You can construct a concrete instance of `AliasPropertiesPtrInput` via:
//
//          AliasPropertiesArgs{...}
//
//  or:
//
//          nil
type AliasPropertiesPtrInput interface {
	pulumi.Input

	ToAliasPropertiesPtrOutput() AliasPropertiesPtrOutput
	ToAliasPropertiesPtrOutputWithContext(context.Context) AliasPropertiesPtrOutput
}

type aliasPropertiesPtrType AliasPropertiesArgs

func AliasPropertiesPtr(v *AliasPropertiesArgs) AliasPropertiesPtrInput {
	return (*aliasPropertiesPtrType)(v)
}

func (*aliasPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AliasProperties)(nil)).Elem()
}

func (i *aliasPropertiesPtrType) ToAliasPropertiesPtrOutput() AliasPropertiesPtrOutput {
	return i.ToAliasPropertiesPtrOutputWithContext(context.Background())
}

func (i *aliasPropertiesPtrType) ToAliasPropertiesPtrOutputWithContext(ctx context.Context) AliasPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasPropertiesPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html
type AliasPropertiesOutput struct{ *pulumi.OutputState }

func (AliasPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AliasProperties)(nil)).Elem()
}

func (o AliasPropertiesOutput) ToAliasPropertiesOutput() AliasPropertiesOutput {
	return o
}

func (o AliasPropertiesOutput) ToAliasPropertiesOutputWithContext(ctx context.Context) AliasPropertiesOutput {
	return o
}

func (o AliasPropertiesOutput) ToAliasPropertiesPtrOutput() AliasPropertiesPtrOutput {
	return o.ToAliasPropertiesPtrOutputWithContext(context.Background())
}

func (o AliasPropertiesOutput) ToAliasPropertiesPtrOutputWithContext(ctx context.Context) AliasPropertiesPtrOutput {
	return o.ApplyT(func(v AliasProperties) *AliasProperties {
		return &v
	}).(AliasPropertiesPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html#cfn-kms-alias-aliasname
func (o AliasPropertiesOutput) AliasName() pulumi.StringOutput {
	return o.ApplyT(func(v AliasProperties) string { return v.AliasName }).(pulumi.StringOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html#cfn-kms-alias-targetkeyid
func (o AliasPropertiesOutput) TargetKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v AliasProperties) string { return v.TargetKeyId }).(pulumi.StringOutput)
}

type AliasPropertiesPtrOutput struct{ *pulumi.OutputState }

func (AliasPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AliasProperties)(nil)).Elem()
}

func (o AliasPropertiesPtrOutput) ToAliasPropertiesPtrOutput() AliasPropertiesPtrOutput {
	return o
}

func (o AliasPropertiesPtrOutput) ToAliasPropertiesPtrOutputWithContext(ctx context.Context) AliasPropertiesPtrOutput {
	return o
}

func (o AliasPropertiesPtrOutput) Elem() AliasPropertiesOutput {
	return o.ApplyT(func(v *AliasProperties) AliasProperties { return *v }).(AliasPropertiesOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html#cfn-kms-alias-aliasname
func (o AliasPropertiesPtrOutput) AliasName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AliasProperties) *string {
		if v == nil {
			return nil
		}
		return &v.AliasName
	}).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html#cfn-kms-alias-targetkeyid
func (o AliasPropertiesPtrOutput) TargetKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AliasProperties) *string {
		if v == nil {
			return nil
		}
		return &v.TargetKeyId
	}).(pulumi.StringPtrOutput)
}

type KeyAttributes struct {
	Arn   string `pulumi:"Arn"`
	KeyId string `pulumi:"KeyId"`
}

// KeyAttributesInput is an input type that accepts KeyAttributesArgs and KeyAttributesOutput values.
// You can construct a concrete instance of `KeyAttributesInput` via:
//
//          KeyAttributesArgs{...}
type KeyAttributesInput interface {
	pulumi.Input

	ToKeyAttributesOutput() KeyAttributesOutput
	ToKeyAttributesOutputWithContext(context.Context) KeyAttributesOutput
}

type KeyAttributesArgs struct {
	Arn   pulumi.StringInput `pulumi:"Arn"`
	KeyId pulumi.StringInput `pulumi:"KeyId"`
}

func (KeyAttributesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyAttributes)(nil)).Elem()
}

func (i KeyAttributesArgs) ToKeyAttributesOutput() KeyAttributesOutput {
	return i.ToKeyAttributesOutputWithContext(context.Background())
}

func (i KeyAttributesArgs) ToKeyAttributesOutputWithContext(ctx context.Context) KeyAttributesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyAttributesOutput)
}

func (i KeyAttributesArgs) ToKeyAttributesPtrOutput() KeyAttributesPtrOutput {
	return i.ToKeyAttributesPtrOutputWithContext(context.Background())
}

func (i KeyAttributesArgs) ToKeyAttributesPtrOutputWithContext(ctx context.Context) KeyAttributesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyAttributesOutput).ToKeyAttributesPtrOutputWithContext(ctx)
}

// KeyAttributesPtrInput is an input type that accepts KeyAttributesArgs, KeyAttributesPtr and KeyAttributesPtrOutput values.
// You can construct a concrete instance of `KeyAttributesPtrInput` via:
//
//          KeyAttributesArgs{...}
//
//  or:
//
//          nil
type KeyAttributesPtrInput interface {
	pulumi.Input

	ToKeyAttributesPtrOutput() KeyAttributesPtrOutput
	ToKeyAttributesPtrOutputWithContext(context.Context) KeyAttributesPtrOutput
}

type keyAttributesPtrType KeyAttributesArgs

func KeyAttributesPtr(v *KeyAttributesArgs) KeyAttributesPtrInput {
	return (*keyAttributesPtrType)(v)
}

func (*keyAttributesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyAttributes)(nil)).Elem()
}

func (i *keyAttributesPtrType) ToKeyAttributesPtrOutput() KeyAttributesPtrOutput {
	return i.ToKeyAttributesPtrOutputWithContext(context.Background())
}

func (i *keyAttributesPtrType) ToKeyAttributesPtrOutputWithContext(ctx context.Context) KeyAttributesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyAttributesPtrOutput)
}

type KeyAttributesOutput struct{ *pulumi.OutputState }

func (KeyAttributesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyAttributes)(nil)).Elem()
}

func (o KeyAttributesOutput) ToKeyAttributesOutput() KeyAttributesOutput {
	return o
}

func (o KeyAttributesOutput) ToKeyAttributesOutputWithContext(ctx context.Context) KeyAttributesOutput {
	return o
}

func (o KeyAttributesOutput) ToKeyAttributesPtrOutput() KeyAttributesPtrOutput {
	return o.ToKeyAttributesPtrOutputWithContext(context.Background())
}

func (o KeyAttributesOutput) ToKeyAttributesPtrOutputWithContext(ctx context.Context) KeyAttributesPtrOutput {
	return o.ApplyT(func(v KeyAttributes) *KeyAttributes {
		return &v
	}).(KeyAttributesPtrOutput)
}
func (o KeyAttributesOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v KeyAttributes) string { return v.Arn }).(pulumi.StringOutput)
}

func (o KeyAttributesOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v KeyAttributes) string { return v.KeyId }).(pulumi.StringOutput)
}

type KeyAttributesPtrOutput struct{ *pulumi.OutputState }

func (KeyAttributesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyAttributes)(nil)).Elem()
}

func (o KeyAttributesPtrOutput) ToKeyAttributesPtrOutput() KeyAttributesPtrOutput {
	return o
}

func (o KeyAttributesPtrOutput) ToKeyAttributesPtrOutputWithContext(ctx context.Context) KeyAttributesPtrOutput {
	return o
}

func (o KeyAttributesPtrOutput) Elem() KeyAttributesOutput {
	return o.ApplyT(func(v *KeyAttributes) KeyAttributes { return *v }).(KeyAttributesOutput)
}

func (o KeyAttributesPtrOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyAttributes) *string {
		if v == nil {
			return nil
		}
		return &v.Arn
	}).(pulumi.StringPtrOutput)
}

func (o KeyAttributesPtrOutput) KeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyAttributes) *string {
		if v == nil {
			return nil
		}
		return &v.KeyId
	}).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html
type KeyProperties struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-description
	Description *string `pulumi:"Description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-enablekeyrotation
	EnableKeyRotation *bool `pulumi:"EnableKeyRotation"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-enabled
	Enabled *bool `pulumi:"Enabled"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-keypolicy
	KeyPolicy interface{} `pulumi:"KeyPolicy"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-keyspec
	KeySpec *string `pulumi:"KeySpec"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-keyusage
	KeyUsage *string `pulumi:"KeyUsage"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-pendingwindowindays
	PendingWindowInDays *int `pulumi:"PendingWindowInDays"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-tags
	Tags []cloudformation.Tag `pulumi:"Tags"`
}

// KeyPropertiesInput is an input type that accepts KeyPropertiesArgs and KeyPropertiesOutput values.
// You can construct a concrete instance of `KeyPropertiesInput` via:
//
//          KeyPropertiesArgs{...}
type KeyPropertiesInput interface {
	pulumi.Input

	ToKeyPropertiesOutput() KeyPropertiesOutput
	ToKeyPropertiesOutputWithContext(context.Context) KeyPropertiesOutput
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html
type KeyPropertiesArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-description
	Description pulumi.StringPtrInput `pulumi:"Description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-enablekeyrotation
	EnableKeyRotation pulumi.BoolPtrInput `pulumi:"EnableKeyRotation"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-enabled
	Enabled pulumi.BoolPtrInput `pulumi:"Enabled"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-keypolicy
	KeyPolicy pulumi.Input `pulumi:"KeyPolicy"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-keyspec
	KeySpec pulumi.StringPtrInput `pulumi:"KeySpec"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-keyusage
	KeyUsage pulumi.StringPtrInput `pulumi:"KeyUsage"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-pendingwindowindays
	PendingWindowInDays pulumi.IntPtrInput `pulumi:"PendingWindowInDays"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-tags
	Tags cloudformation.TagArrayInput `pulumi:"Tags"`
}

func (KeyPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyProperties)(nil)).Elem()
}

func (i KeyPropertiesArgs) ToKeyPropertiesOutput() KeyPropertiesOutput {
	return i.ToKeyPropertiesOutputWithContext(context.Background())
}

func (i KeyPropertiesArgs) ToKeyPropertiesOutputWithContext(ctx context.Context) KeyPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPropertiesOutput)
}

func (i KeyPropertiesArgs) ToKeyPropertiesPtrOutput() KeyPropertiesPtrOutput {
	return i.ToKeyPropertiesPtrOutputWithContext(context.Background())
}

func (i KeyPropertiesArgs) ToKeyPropertiesPtrOutputWithContext(ctx context.Context) KeyPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPropertiesOutput).ToKeyPropertiesPtrOutputWithContext(ctx)
}

// KeyPropertiesPtrInput is an input type that accepts KeyPropertiesArgs, KeyPropertiesPtr and KeyPropertiesPtrOutput values.
// You can construct a concrete instance of `KeyPropertiesPtrInput` via:
//
//          KeyPropertiesArgs{...}
//
//  or:
//
//          nil
type KeyPropertiesPtrInput interface {
	pulumi.Input

	ToKeyPropertiesPtrOutput() KeyPropertiesPtrOutput
	ToKeyPropertiesPtrOutputWithContext(context.Context) KeyPropertiesPtrOutput
}

type keyPropertiesPtrType KeyPropertiesArgs

func KeyPropertiesPtr(v *KeyPropertiesArgs) KeyPropertiesPtrInput {
	return (*keyPropertiesPtrType)(v)
}

func (*keyPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyProperties)(nil)).Elem()
}

func (i *keyPropertiesPtrType) ToKeyPropertiesPtrOutput() KeyPropertiesPtrOutput {
	return i.ToKeyPropertiesPtrOutputWithContext(context.Background())
}

func (i *keyPropertiesPtrType) ToKeyPropertiesPtrOutputWithContext(ctx context.Context) KeyPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPropertiesPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html
type KeyPropertiesOutput struct{ *pulumi.OutputState }

func (KeyPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyProperties)(nil)).Elem()
}

func (o KeyPropertiesOutput) ToKeyPropertiesOutput() KeyPropertiesOutput {
	return o
}

func (o KeyPropertiesOutput) ToKeyPropertiesOutputWithContext(ctx context.Context) KeyPropertiesOutput {
	return o
}

func (o KeyPropertiesOutput) ToKeyPropertiesPtrOutput() KeyPropertiesPtrOutput {
	return o.ToKeyPropertiesPtrOutputWithContext(context.Background())
}

func (o KeyPropertiesOutput) ToKeyPropertiesPtrOutputWithContext(ctx context.Context) KeyPropertiesPtrOutput {
	return o.ApplyT(func(v KeyProperties) *KeyProperties {
		return &v
	}).(KeyPropertiesPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-description
func (o KeyPropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-enablekeyrotation
func (o KeyPropertiesOutput) EnableKeyRotation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeyProperties) *bool { return v.EnableKeyRotation }).(pulumi.BoolPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-enabled
func (o KeyPropertiesOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeyProperties) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-keypolicy
func (o KeyPropertiesOutput) KeyPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v KeyProperties) interface{} { return v.KeyPolicy }).(pulumi.AnyOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-keyspec
func (o KeyPropertiesOutput) KeySpec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyProperties) *string { return v.KeySpec }).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-keyusage
func (o KeyPropertiesOutput) KeyUsage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyProperties) *string { return v.KeyUsage }).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-pendingwindowindays
func (o KeyPropertiesOutput) PendingWindowInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v KeyProperties) *int { return v.PendingWindowInDays }).(pulumi.IntPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-tags
func (o KeyPropertiesOutput) Tags() cloudformation.TagArrayOutput {
	return o.ApplyT(func(v KeyProperties) []cloudformation.Tag { return v.Tags }).(cloudformation.TagArrayOutput)
}

type KeyPropertiesPtrOutput struct{ *pulumi.OutputState }

func (KeyPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyProperties)(nil)).Elem()
}

func (o KeyPropertiesPtrOutput) ToKeyPropertiesPtrOutput() KeyPropertiesPtrOutput {
	return o
}

func (o KeyPropertiesPtrOutput) ToKeyPropertiesPtrOutputWithContext(ctx context.Context) KeyPropertiesPtrOutput {
	return o
}

func (o KeyPropertiesPtrOutput) Elem() KeyPropertiesOutput {
	return o.ApplyT(func(v *KeyProperties) KeyProperties { return *v }).(KeyPropertiesOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-description
func (o KeyPropertiesPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyProperties) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-enablekeyrotation
func (o KeyPropertiesPtrOutput) EnableKeyRotation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KeyProperties) *bool {
		if v == nil {
			return nil
		}
		return v.EnableKeyRotation
	}).(pulumi.BoolPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-enabled
func (o KeyPropertiesPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KeyProperties) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-keypolicy
func (o KeyPropertiesPtrOutput) KeyPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v *KeyProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.KeyPolicy
	}).(pulumi.AnyOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-keyspec
func (o KeyPropertiesPtrOutput) KeySpec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeySpec
	}).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-keyusage
func (o KeyPropertiesPtrOutput) KeyUsage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyUsage
	}).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-pendingwindowindays
func (o KeyPropertiesPtrOutput) PendingWindowInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *KeyProperties) *int {
		if v == nil {
			return nil
		}
		return v.PendingWindowInDays
	}).(pulumi.IntPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-tags
func (o KeyPropertiesPtrOutput) Tags() cloudformation.TagArrayOutput {
	return o.ApplyT(func(v *KeyProperties) []cloudformation.Tag {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(cloudformation.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(AliasAttributesOutput{})
	pulumi.RegisterOutputType(AliasAttributesPtrOutput{})
	pulumi.RegisterOutputType(AliasPropertiesOutput{})
	pulumi.RegisterOutputType(AliasPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KeyAttributesOutput{})
	pulumi.RegisterOutputType(KeyAttributesPtrOutput{})
	pulumi.RegisterOutputType(KeyPropertiesOutput{})
	pulumi.RegisterOutputType(KeyPropertiesPtrOutput{})
}
