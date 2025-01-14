// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package finspace

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

// Additional parameters to identify Federation mode
type EnvironmentFederationParameters struct {
	// SAML metadata URL to link with the Environment
	ApplicationCallBackUrl *string `pulumi:"applicationCallBackUrl"`
	// Attribute map for SAML configuration
	AttributeMap []EnvironmentFederationParametersAttributeMapItemProperties `pulumi:"attributeMap"`
	// Federation provider name to link with the Environment
	FederationProviderName *string `pulumi:"federationProviderName"`
	// SAML metadata URL to link with the Environment
	FederationUrn *string `pulumi:"federationUrn"`
	// SAML metadata document to link the federation provider to the Environment
	SamlMetadataDocument *string `pulumi:"samlMetadataDocument"`
	// SAML metadata URL to link with the Environment
	SamlMetadataUrl *string `pulumi:"samlMetadataUrl"`
}

// EnvironmentFederationParametersInput is an input type that accepts EnvironmentFederationParametersArgs and EnvironmentFederationParametersOutput values.
// You can construct a concrete instance of `EnvironmentFederationParametersInput` via:
//
//	EnvironmentFederationParametersArgs{...}
type EnvironmentFederationParametersInput interface {
	pulumi.Input

	ToEnvironmentFederationParametersOutput() EnvironmentFederationParametersOutput
	ToEnvironmentFederationParametersOutputWithContext(context.Context) EnvironmentFederationParametersOutput
}

// Additional parameters to identify Federation mode
type EnvironmentFederationParametersArgs struct {
	// SAML metadata URL to link with the Environment
	ApplicationCallBackUrl pulumi.StringPtrInput `pulumi:"applicationCallBackUrl"`
	// Attribute map for SAML configuration
	AttributeMap EnvironmentFederationParametersAttributeMapItemPropertiesArrayInput `pulumi:"attributeMap"`
	// Federation provider name to link with the Environment
	FederationProviderName pulumi.StringPtrInput `pulumi:"federationProviderName"`
	// SAML metadata URL to link with the Environment
	FederationUrn pulumi.StringPtrInput `pulumi:"federationUrn"`
	// SAML metadata document to link the federation provider to the Environment
	SamlMetadataDocument pulumi.StringPtrInput `pulumi:"samlMetadataDocument"`
	// SAML metadata URL to link with the Environment
	SamlMetadataUrl pulumi.StringPtrInput `pulumi:"samlMetadataUrl"`
}

func (EnvironmentFederationParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentFederationParameters)(nil)).Elem()
}

func (i EnvironmentFederationParametersArgs) ToEnvironmentFederationParametersOutput() EnvironmentFederationParametersOutput {
	return i.ToEnvironmentFederationParametersOutputWithContext(context.Background())
}

func (i EnvironmentFederationParametersArgs) ToEnvironmentFederationParametersOutputWithContext(ctx context.Context) EnvironmentFederationParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentFederationParametersOutput)
}

func (i EnvironmentFederationParametersArgs) ToEnvironmentFederationParametersPtrOutput() EnvironmentFederationParametersPtrOutput {
	return i.ToEnvironmentFederationParametersPtrOutputWithContext(context.Background())
}

func (i EnvironmentFederationParametersArgs) ToEnvironmentFederationParametersPtrOutputWithContext(ctx context.Context) EnvironmentFederationParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentFederationParametersOutput).ToEnvironmentFederationParametersPtrOutputWithContext(ctx)
}

// EnvironmentFederationParametersPtrInput is an input type that accepts EnvironmentFederationParametersArgs, EnvironmentFederationParametersPtr and EnvironmentFederationParametersPtrOutput values.
// You can construct a concrete instance of `EnvironmentFederationParametersPtrInput` via:
//
//	        EnvironmentFederationParametersArgs{...}
//
//	or:
//
//	        nil
type EnvironmentFederationParametersPtrInput interface {
	pulumi.Input

	ToEnvironmentFederationParametersPtrOutput() EnvironmentFederationParametersPtrOutput
	ToEnvironmentFederationParametersPtrOutputWithContext(context.Context) EnvironmentFederationParametersPtrOutput
}

type environmentFederationParametersPtrType EnvironmentFederationParametersArgs

func EnvironmentFederationParametersPtr(v *EnvironmentFederationParametersArgs) EnvironmentFederationParametersPtrInput {
	return (*environmentFederationParametersPtrType)(v)
}

func (*environmentFederationParametersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentFederationParameters)(nil)).Elem()
}

func (i *environmentFederationParametersPtrType) ToEnvironmentFederationParametersPtrOutput() EnvironmentFederationParametersPtrOutput {
	return i.ToEnvironmentFederationParametersPtrOutputWithContext(context.Background())
}

func (i *environmentFederationParametersPtrType) ToEnvironmentFederationParametersPtrOutputWithContext(ctx context.Context) EnvironmentFederationParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentFederationParametersPtrOutput)
}

// Additional parameters to identify Federation mode
type EnvironmentFederationParametersOutput struct{ *pulumi.OutputState }

func (EnvironmentFederationParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentFederationParameters)(nil)).Elem()
}

func (o EnvironmentFederationParametersOutput) ToEnvironmentFederationParametersOutput() EnvironmentFederationParametersOutput {
	return o
}

func (o EnvironmentFederationParametersOutput) ToEnvironmentFederationParametersOutputWithContext(ctx context.Context) EnvironmentFederationParametersOutput {
	return o
}

func (o EnvironmentFederationParametersOutput) ToEnvironmentFederationParametersPtrOutput() EnvironmentFederationParametersPtrOutput {
	return o.ToEnvironmentFederationParametersPtrOutputWithContext(context.Background())
}

func (o EnvironmentFederationParametersOutput) ToEnvironmentFederationParametersPtrOutputWithContext(ctx context.Context) EnvironmentFederationParametersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnvironmentFederationParameters) *EnvironmentFederationParameters {
		return &v
	}).(EnvironmentFederationParametersPtrOutput)
}

// SAML metadata URL to link with the Environment
func (o EnvironmentFederationParametersOutput) ApplicationCallBackUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentFederationParameters) *string { return v.ApplicationCallBackUrl }).(pulumi.StringPtrOutput)
}

// Attribute map for SAML configuration
func (o EnvironmentFederationParametersOutput) AttributeMap() EnvironmentFederationParametersAttributeMapItemPropertiesArrayOutput {
	return o.ApplyT(func(v EnvironmentFederationParameters) []EnvironmentFederationParametersAttributeMapItemProperties {
		return v.AttributeMap
	}).(EnvironmentFederationParametersAttributeMapItemPropertiesArrayOutput)
}

// Federation provider name to link with the Environment
func (o EnvironmentFederationParametersOutput) FederationProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentFederationParameters) *string { return v.FederationProviderName }).(pulumi.StringPtrOutput)
}

// SAML metadata URL to link with the Environment
func (o EnvironmentFederationParametersOutput) FederationUrn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentFederationParameters) *string { return v.FederationUrn }).(pulumi.StringPtrOutput)
}

// SAML metadata document to link the federation provider to the Environment
func (o EnvironmentFederationParametersOutput) SamlMetadataDocument() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentFederationParameters) *string { return v.SamlMetadataDocument }).(pulumi.StringPtrOutput)
}

// SAML metadata URL to link with the Environment
func (o EnvironmentFederationParametersOutput) SamlMetadataUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentFederationParameters) *string { return v.SamlMetadataUrl }).(pulumi.StringPtrOutput)
}

type EnvironmentFederationParametersPtrOutput struct{ *pulumi.OutputState }

func (EnvironmentFederationParametersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentFederationParameters)(nil)).Elem()
}

func (o EnvironmentFederationParametersPtrOutput) ToEnvironmentFederationParametersPtrOutput() EnvironmentFederationParametersPtrOutput {
	return o
}

func (o EnvironmentFederationParametersPtrOutput) ToEnvironmentFederationParametersPtrOutputWithContext(ctx context.Context) EnvironmentFederationParametersPtrOutput {
	return o
}

func (o EnvironmentFederationParametersPtrOutput) Elem() EnvironmentFederationParametersOutput {
	return o.ApplyT(func(v *EnvironmentFederationParameters) EnvironmentFederationParameters {
		if v != nil {
			return *v
		}
		var ret EnvironmentFederationParameters
		return ret
	}).(EnvironmentFederationParametersOutput)
}

// SAML metadata URL to link with the Environment
func (o EnvironmentFederationParametersPtrOutput) ApplicationCallBackUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentFederationParameters) *string {
		if v == nil {
			return nil
		}
		return v.ApplicationCallBackUrl
	}).(pulumi.StringPtrOutput)
}

// Attribute map for SAML configuration
func (o EnvironmentFederationParametersPtrOutput) AttributeMap() EnvironmentFederationParametersAttributeMapItemPropertiesArrayOutput {
	return o.ApplyT(func(v *EnvironmentFederationParameters) []EnvironmentFederationParametersAttributeMapItemProperties {
		if v == nil {
			return nil
		}
		return v.AttributeMap
	}).(EnvironmentFederationParametersAttributeMapItemPropertiesArrayOutput)
}

// Federation provider name to link with the Environment
func (o EnvironmentFederationParametersPtrOutput) FederationProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentFederationParameters) *string {
		if v == nil {
			return nil
		}
		return v.FederationProviderName
	}).(pulumi.StringPtrOutput)
}

// SAML metadata URL to link with the Environment
func (o EnvironmentFederationParametersPtrOutput) FederationUrn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentFederationParameters) *string {
		if v == nil {
			return nil
		}
		return v.FederationUrn
	}).(pulumi.StringPtrOutput)
}

// SAML metadata document to link the federation provider to the Environment
func (o EnvironmentFederationParametersPtrOutput) SamlMetadataDocument() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentFederationParameters) *string {
		if v == nil {
			return nil
		}
		return v.SamlMetadataDocument
	}).(pulumi.StringPtrOutput)
}

// SAML metadata URL to link with the Environment
func (o EnvironmentFederationParametersPtrOutput) SamlMetadataUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentFederationParameters) *string {
		if v == nil {
			return nil
		}
		return v.SamlMetadataUrl
	}).(pulumi.StringPtrOutput)
}

type EnvironmentFederationParametersAttributeMapItemProperties struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key *string `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value *string `pulumi:"value"`
}

// EnvironmentFederationParametersAttributeMapItemPropertiesInput is an input type that accepts EnvironmentFederationParametersAttributeMapItemPropertiesArgs and EnvironmentFederationParametersAttributeMapItemPropertiesOutput values.
// You can construct a concrete instance of `EnvironmentFederationParametersAttributeMapItemPropertiesInput` via:
//
//	EnvironmentFederationParametersAttributeMapItemPropertiesArgs{...}
type EnvironmentFederationParametersAttributeMapItemPropertiesInput interface {
	pulumi.Input

	ToEnvironmentFederationParametersAttributeMapItemPropertiesOutput() EnvironmentFederationParametersAttributeMapItemPropertiesOutput
	ToEnvironmentFederationParametersAttributeMapItemPropertiesOutputWithContext(context.Context) EnvironmentFederationParametersAttributeMapItemPropertiesOutput
}

type EnvironmentFederationParametersAttributeMapItemPropertiesArgs struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key pulumi.StringPtrInput `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (EnvironmentFederationParametersAttributeMapItemPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentFederationParametersAttributeMapItemProperties)(nil)).Elem()
}

func (i EnvironmentFederationParametersAttributeMapItemPropertiesArgs) ToEnvironmentFederationParametersAttributeMapItemPropertiesOutput() EnvironmentFederationParametersAttributeMapItemPropertiesOutput {
	return i.ToEnvironmentFederationParametersAttributeMapItemPropertiesOutputWithContext(context.Background())
}

func (i EnvironmentFederationParametersAttributeMapItemPropertiesArgs) ToEnvironmentFederationParametersAttributeMapItemPropertiesOutputWithContext(ctx context.Context) EnvironmentFederationParametersAttributeMapItemPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentFederationParametersAttributeMapItemPropertiesOutput)
}

// EnvironmentFederationParametersAttributeMapItemPropertiesArrayInput is an input type that accepts EnvironmentFederationParametersAttributeMapItemPropertiesArray and EnvironmentFederationParametersAttributeMapItemPropertiesArrayOutput values.
// You can construct a concrete instance of `EnvironmentFederationParametersAttributeMapItemPropertiesArrayInput` via:
//
//	EnvironmentFederationParametersAttributeMapItemPropertiesArray{ EnvironmentFederationParametersAttributeMapItemPropertiesArgs{...} }
type EnvironmentFederationParametersAttributeMapItemPropertiesArrayInput interface {
	pulumi.Input

	ToEnvironmentFederationParametersAttributeMapItemPropertiesArrayOutput() EnvironmentFederationParametersAttributeMapItemPropertiesArrayOutput
	ToEnvironmentFederationParametersAttributeMapItemPropertiesArrayOutputWithContext(context.Context) EnvironmentFederationParametersAttributeMapItemPropertiesArrayOutput
}

type EnvironmentFederationParametersAttributeMapItemPropertiesArray []EnvironmentFederationParametersAttributeMapItemPropertiesInput

func (EnvironmentFederationParametersAttributeMapItemPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentFederationParametersAttributeMapItemProperties)(nil)).Elem()
}

func (i EnvironmentFederationParametersAttributeMapItemPropertiesArray) ToEnvironmentFederationParametersAttributeMapItemPropertiesArrayOutput() EnvironmentFederationParametersAttributeMapItemPropertiesArrayOutput {
	return i.ToEnvironmentFederationParametersAttributeMapItemPropertiesArrayOutputWithContext(context.Background())
}

func (i EnvironmentFederationParametersAttributeMapItemPropertiesArray) ToEnvironmentFederationParametersAttributeMapItemPropertiesArrayOutputWithContext(ctx context.Context) EnvironmentFederationParametersAttributeMapItemPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentFederationParametersAttributeMapItemPropertiesArrayOutput)
}

type EnvironmentFederationParametersAttributeMapItemPropertiesOutput struct{ *pulumi.OutputState }

func (EnvironmentFederationParametersAttributeMapItemPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentFederationParametersAttributeMapItemProperties)(nil)).Elem()
}

func (o EnvironmentFederationParametersAttributeMapItemPropertiesOutput) ToEnvironmentFederationParametersAttributeMapItemPropertiesOutput() EnvironmentFederationParametersAttributeMapItemPropertiesOutput {
	return o
}

func (o EnvironmentFederationParametersAttributeMapItemPropertiesOutput) ToEnvironmentFederationParametersAttributeMapItemPropertiesOutputWithContext(ctx context.Context) EnvironmentFederationParametersAttributeMapItemPropertiesOutput {
	return o
}

// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o EnvironmentFederationParametersAttributeMapItemPropertiesOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentFederationParametersAttributeMapItemProperties) *string { return v.Key }).(pulumi.StringPtrOutput)
}

// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o EnvironmentFederationParametersAttributeMapItemPropertiesOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentFederationParametersAttributeMapItemProperties) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type EnvironmentFederationParametersAttributeMapItemPropertiesArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentFederationParametersAttributeMapItemPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentFederationParametersAttributeMapItemProperties)(nil)).Elem()
}

func (o EnvironmentFederationParametersAttributeMapItemPropertiesArrayOutput) ToEnvironmentFederationParametersAttributeMapItemPropertiesArrayOutput() EnvironmentFederationParametersAttributeMapItemPropertiesArrayOutput {
	return o
}

func (o EnvironmentFederationParametersAttributeMapItemPropertiesArrayOutput) ToEnvironmentFederationParametersAttributeMapItemPropertiesArrayOutputWithContext(ctx context.Context) EnvironmentFederationParametersAttributeMapItemPropertiesArrayOutput {
	return o
}

func (o EnvironmentFederationParametersAttributeMapItemPropertiesArrayOutput) Index(i pulumi.IntInput) EnvironmentFederationParametersAttributeMapItemPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EnvironmentFederationParametersAttributeMapItemProperties {
		return vs[0].([]EnvironmentFederationParametersAttributeMapItemProperties)[vs[1].(int)]
	}).(EnvironmentFederationParametersAttributeMapItemPropertiesOutput)
}

// Parameters of the first Superuser for the FinSpace Environment
type EnvironmentSuperuserParameters struct {
	// Email address
	EmailAddress *string `pulumi:"emailAddress"`
	// First name
	FirstName *string `pulumi:"firstName"`
	// Last name
	LastName *string `pulumi:"lastName"`
}

// EnvironmentSuperuserParametersInput is an input type that accepts EnvironmentSuperuserParametersArgs and EnvironmentSuperuserParametersOutput values.
// You can construct a concrete instance of `EnvironmentSuperuserParametersInput` via:
//
//	EnvironmentSuperuserParametersArgs{...}
type EnvironmentSuperuserParametersInput interface {
	pulumi.Input

	ToEnvironmentSuperuserParametersOutput() EnvironmentSuperuserParametersOutput
	ToEnvironmentSuperuserParametersOutputWithContext(context.Context) EnvironmentSuperuserParametersOutput
}

// Parameters of the first Superuser for the FinSpace Environment
type EnvironmentSuperuserParametersArgs struct {
	// Email address
	EmailAddress pulumi.StringPtrInput `pulumi:"emailAddress"`
	// First name
	FirstName pulumi.StringPtrInput `pulumi:"firstName"`
	// Last name
	LastName pulumi.StringPtrInput `pulumi:"lastName"`
}

func (EnvironmentSuperuserParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentSuperuserParameters)(nil)).Elem()
}

func (i EnvironmentSuperuserParametersArgs) ToEnvironmentSuperuserParametersOutput() EnvironmentSuperuserParametersOutput {
	return i.ToEnvironmentSuperuserParametersOutputWithContext(context.Background())
}

func (i EnvironmentSuperuserParametersArgs) ToEnvironmentSuperuserParametersOutputWithContext(ctx context.Context) EnvironmentSuperuserParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentSuperuserParametersOutput)
}

func (i EnvironmentSuperuserParametersArgs) ToEnvironmentSuperuserParametersPtrOutput() EnvironmentSuperuserParametersPtrOutput {
	return i.ToEnvironmentSuperuserParametersPtrOutputWithContext(context.Background())
}

func (i EnvironmentSuperuserParametersArgs) ToEnvironmentSuperuserParametersPtrOutputWithContext(ctx context.Context) EnvironmentSuperuserParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentSuperuserParametersOutput).ToEnvironmentSuperuserParametersPtrOutputWithContext(ctx)
}

// EnvironmentSuperuserParametersPtrInput is an input type that accepts EnvironmentSuperuserParametersArgs, EnvironmentSuperuserParametersPtr and EnvironmentSuperuserParametersPtrOutput values.
// You can construct a concrete instance of `EnvironmentSuperuserParametersPtrInput` via:
//
//	        EnvironmentSuperuserParametersArgs{...}
//
//	or:
//
//	        nil
type EnvironmentSuperuserParametersPtrInput interface {
	pulumi.Input

	ToEnvironmentSuperuserParametersPtrOutput() EnvironmentSuperuserParametersPtrOutput
	ToEnvironmentSuperuserParametersPtrOutputWithContext(context.Context) EnvironmentSuperuserParametersPtrOutput
}

type environmentSuperuserParametersPtrType EnvironmentSuperuserParametersArgs

func EnvironmentSuperuserParametersPtr(v *EnvironmentSuperuserParametersArgs) EnvironmentSuperuserParametersPtrInput {
	return (*environmentSuperuserParametersPtrType)(v)
}

func (*environmentSuperuserParametersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentSuperuserParameters)(nil)).Elem()
}

func (i *environmentSuperuserParametersPtrType) ToEnvironmentSuperuserParametersPtrOutput() EnvironmentSuperuserParametersPtrOutput {
	return i.ToEnvironmentSuperuserParametersPtrOutputWithContext(context.Background())
}

func (i *environmentSuperuserParametersPtrType) ToEnvironmentSuperuserParametersPtrOutputWithContext(ctx context.Context) EnvironmentSuperuserParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentSuperuserParametersPtrOutput)
}

// Parameters of the first Superuser for the FinSpace Environment
type EnvironmentSuperuserParametersOutput struct{ *pulumi.OutputState }

func (EnvironmentSuperuserParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentSuperuserParameters)(nil)).Elem()
}

func (o EnvironmentSuperuserParametersOutput) ToEnvironmentSuperuserParametersOutput() EnvironmentSuperuserParametersOutput {
	return o
}

func (o EnvironmentSuperuserParametersOutput) ToEnvironmentSuperuserParametersOutputWithContext(ctx context.Context) EnvironmentSuperuserParametersOutput {
	return o
}

func (o EnvironmentSuperuserParametersOutput) ToEnvironmentSuperuserParametersPtrOutput() EnvironmentSuperuserParametersPtrOutput {
	return o.ToEnvironmentSuperuserParametersPtrOutputWithContext(context.Background())
}

func (o EnvironmentSuperuserParametersOutput) ToEnvironmentSuperuserParametersPtrOutputWithContext(ctx context.Context) EnvironmentSuperuserParametersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnvironmentSuperuserParameters) *EnvironmentSuperuserParameters {
		return &v
	}).(EnvironmentSuperuserParametersPtrOutput)
}

// Email address
func (o EnvironmentSuperuserParametersOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentSuperuserParameters) *string { return v.EmailAddress }).(pulumi.StringPtrOutput)
}

// First name
func (o EnvironmentSuperuserParametersOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentSuperuserParameters) *string { return v.FirstName }).(pulumi.StringPtrOutput)
}

// Last name
func (o EnvironmentSuperuserParametersOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentSuperuserParameters) *string { return v.LastName }).(pulumi.StringPtrOutput)
}

type EnvironmentSuperuserParametersPtrOutput struct{ *pulumi.OutputState }

func (EnvironmentSuperuserParametersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentSuperuserParameters)(nil)).Elem()
}

func (o EnvironmentSuperuserParametersPtrOutput) ToEnvironmentSuperuserParametersPtrOutput() EnvironmentSuperuserParametersPtrOutput {
	return o
}

func (o EnvironmentSuperuserParametersPtrOutput) ToEnvironmentSuperuserParametersPtrOutputWithContext(ctx context.Context) EnvironmentSuperuserParametersPtrOutput {
	return o
}

func (o EnvironmentSuperuserParametersPtrOutput) Elem() EnvironmentSuperuserParametersOutput {
	return o.ApplyT(func(v *EnvironmentSuperuserParameters) EnvironmentSuperuserParameters {
		if v != nil {
			return *v
		}
		var ret EnvironmentSuperuserParameters
		return ret
	}).(EnvironmentSuperuserParametersOutput)
}

// Email address
func (o EnvironmentSuperuserParametersPtrOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentSuperuserParameters) *string {
		if v == nil {
			return nil
		}
		return v.EmailAddress
	}).(pulumi.StringPtrOutput)
}

// First name
func (o EnvironmentSuperuserParametersPtrOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentSuperuserParameters) *string {
		if v == nil {
			return nil
		}
		return v.FirstName
	}).(pulumi.StringPtrOutput)
}

// Last name
func (o EnvironmentSuperuserParametersPtrOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentSuperuserParameters) *string {
		if v == nil {
			return nil
		}
		return v.LastName
	}).(pulumi.StringPtrOutput)
}

// A list of all tags for a resource.
type EnvironmentTag struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key string `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value string `pulumi:"value"`
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentFederationParametersInput)(nil)).Elem(), EnvironmentFederationParametersArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentFederationParametersPtrInput)(nil)).Elem(), EnvironmentFederationParametersArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentFederationParametersAttributeMapItemPropertiesInput)(nil)).Elem(), EnvironmentFederationParametersAttributeMapItemPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentFederationParametersAttributeMapItemPropertiesArrayInput)(nil)).Elem(), EnvironmentFederationParametersAttributeMapItemPropertiesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentSuperuserParametersInput)(nil)).Elem(), EnvironmentSuperuserParametersArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentSuperuserParametersPtrInput)(nil)).Elem(), EnvironmentSuperuserParametersArgs{})
	pulumi.RegisterOutputType(EnvironmentFederationParametersOutput{})
	pulumi.RegisterOutputType(EnvironmentFederationParametersPtrOutput{})
	pulumi.RegisterOutputType(EnvironmentFederationParametersAttributeMapItemPropertiesOutput{})
	pulumi.RegisterOutputType(EnvironmentFederationParametersAttributeMapItemPropertiesArrayOutput{})
	pulumi.RegisterOutputType(EnvironmentSuperuserParametersOutput{})
	pulumi.RegisterOutputType(EnvironmentSuperuserParametersPtrOutput{})
}
