// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package verifiedpermissions

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type IdentitySourceCognitoUserPoolConfiguration struct {
	ClientIds   []string `pulumi:"clientIds"`
	UserPoolArn string   `pulumi:"userPoolArn"`
}

// IdentitySourceCognitoUserPoolConfigurationInput is an input type that accepts IdentitySourceCognitoUserPoolConfigurationArgs and IdentitySourceCognitoUserPoolConfigurationOutput values.
// You can construct a concrete instance of `IdentitySourceCognitoUserPoolConfigurationInput` via:
//
//	IdentitySourceCognitoUserPoolConfigurationArgs{...}
type IdentitySourceCognitoUserPoolConfigurationInput interface {
	pulumi.Input

	ToIdentitySourceCognitoUserPoolConfigurationOutput() IdentitySourceCognitoUserPoolConfigurationOutput
	ToIdentitySourceCognitoUserPoolConfigurationOutputWithContext(context.Context) IdentitySourceCognitoUserPoolConfigurationOutput
}

type IdentitySourceCognitoUserPoolConfigurationArgs struct {
	ClientIds   pulumi.StringArrayInput `pulumi:"clientIds"`
	UserPoolArn pulumi.StringInput      `pulumi:"userPoolArn"`
}

func (IdentitySourceCognitoUserPoolConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentitySourceCognitoUserPoolConfiguration)(nil)).Elem()
}

func (i IdentitySourceCognitoUserPoolConfigurationArgs) ToIdentitySourceCognitoUserPoolConfigurationOutput() IdentitySourceCognitoUserPoolConfigurationOutput {
	return i.ToIdentitySourceCognitoUserPoolConfigurationOutputWithContext(context.Background())
}

func (i IdentitySourceCognitoUserPoolConfigurationArgs) ToIdentitySourceCognitoUserPoolConfigurationOutputWithContext(ctx context.Context) IdentitySourceCognitoUserPoolConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentitySourceCognitoUserPoolConfigurationOutput)
}

type IdentitySourceCognitoUserPoolConfigurationOutput struct{ *pulumi.OutputState }

func (IdentitySourceCognitoUserPoolConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentitySourceCognitoUserPoolConfiguration)(nil)).Elem()
}

func (o IdentitySourceCognitoUserPoolConfigurationOutput) ToIdentitySourceCognitoUserPoolConfigurationOutput() IdentitySourceCognitoUserPoolConfigurationOutput {
	return o
}

func (o IdentitySourceCognitoUserPoolConfigurationOutput) ToIdentitySourceCognitoUserPoolConfigurationOutputWithContext(ctx context.Context) IdentitySourceCognitoUserPoolConfigurationOutput {
	return o
}

func (o IdentitySourceCognitoUserPoolConfigurationOutput) ClientIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IdentitySourceCognitoUserPoolConfiguration) []string { return v.ClientIds }).(pulumi.StringArrayOutput)
}

func (o IdentitySourceCognitoUserPoolConfigurationOutput) UserPoolArn() pulumi.StringOutput {
	return o.ApplyT(func(v IdentitySourceCognitoUserPoolConfiguration) string { return v.UserPoolArn }).(pulumi.StringOutput)
}

type IdentitySourceCognitoUserPoolConfigurationPtrOutput struct{ *pulumi.OutputState }

func (IdentitySourceCognitoUserPoolConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentitySourceCognitoUserPoolConfiguration)(nil)).Elem()
}

func (o IdentitySourceCognitoUserPoolConfigurationPtrOutput) ToIdentitySourceCognitoUserPoolConfigurationPtrOutput() IdentitySourceCognitoUserPoolConfigurationPtrOutput {
	return o
}

func (o IdentitySourceCognitoUserPoolConfigurationPtrOutput) ToIdentitySourceCognitoUserPoolConfigurationPtrOutputWithContext(ctx context.Context) IdentitySourceCognitoUserPoolConfigurationPtrOutput {
	return o
}

func (o IdentitySourceCognitoUserPoolConfigurationPtrOutput) Elem() IdentitySourceCognitoUserPoolConfigurationOutput {
	return o.ApplyT(func(v *IdentitySourceCognitoUserPoolConfiguration) IdentitySourceCognitoUserPoolConfiguration {
		if v != nil {
			return *v
		}
		var ret IdentitySourceCognitoUserPoolConfiguration
		return ret
	}).(IdentitySourceCognitoUserPoolConfigurationOutput)
}

func (o IdentitySourceCognitoUserPoolConfigurationPtrOutput) ClientIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IdentitySourceCognitoUserPoolConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.ClientIds
	}).(pulumi.StringArrayOutput)
}

func (o IdentitySourceCognitoUserPoolConfigurationPtrOutput) UserPoolArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentitySourceCognitoUserPoolConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.UserPoolArn
	}).(pulumi.StringPtrOutput)
}

type IdentitySourceConfiguration struct {
	CognitoUserPoolConfiguration IdentitySourceCognitoUserPoolConfiguration `pulumi:"cognitoUserPoolConfiguration"`
}

// IdentitySourceConfigurationInput is an input type that accepts IdentitySourceConfigurationArgs and IdentitySourceConfigurationOutput values.
// You can construct a concrete instance of `IdentitySourceConfigurationInput` via:
//
//	IdentitySourceConfigurationArgs{...}
type IdentitySourceConfigurationInput interface {
	pulumi.Input

	ToIdentitySourceConfigurationOutput() IdentitySourceConfigurationOutput
	ToIdentitySourceConfigurationOutputWithContext(context.Context) IdentitySourceConfigurationOutput
}

type IdentitySourceConfigurationArgs struct {
	CognitoUserPoolConfiguration IdentitySourceCognitoUserPoolConfigurationInput `pulumi:"cognitoUserPoolConfiguration"`
}

func (IdentitySourceConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentitySourceConfiguration)(nil)).Elem()
}

func (i IdentitySourceConfigurationArgs) ToIdentitySourceConfigurationOutput() IdentitySourceConfigurationOutput {
	return i.ToIdentitySourceConfigurationOutputWithContext(context.Background())
}

func (i IdentitySourceConfigurationArgs) ToIdentitySourceConfigurationOutputWithContext(ctx context.Context) IdentitySourceConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentitySourceConfigurationOutput)
}

type IdentitySourceConfigurationOutput struct{ *pulumi.OutputState }

func (IdentitySourceConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentitySourceConfiguration)(nil)).Elem()
}

func (o IdentitySourceConfigurationOutput) ToIdentitySourceConfigurationOutput() IdentitySourceConfigurationOutput {
	return o
}

func (o IdentitySourceConfigurationOutput) ToIdentitySourceConfigurationOutputWithContext(ctx context.Context) IdentitySourceConfigurationOutput {
	return o
}

func (o IdentitySourceConfigurationOutput) CognitoUserPoolConfiguration() IdentitySourceCognitoUserPoolConfigurationOutput {
	return o.ApplyT(func(v IdentitySourceConfiguration) IdentitySourceCognitoUserPoolConfiguration {
		return v.CognitoUserPoolConfiguration
	}).(IdentitySourceCognitoUserPoolConfigurationOutput)
}

type IdentitySourceConfigurationPtrOutput struct{ *pulumi.OutputState }

func (IdentitySourceConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentitySourceConfiguration)(nil)).Elem()
}

func (o IdentitySourceConfigurationPtrOutput) ToIdentitySourceConfigurationPtrOutput() IdentitySourceConfigurationPtrOutput {
	return o
}

func (o IdentitySourceConfigurationPtrOutput) ToIdentitySourceConfigurationPtrOutputWithContext(ctx context.Context) IdentitySourceConfigurationPtrOutput {
	return o
}

func (o IdentitySourceConfigurationPtrOutput) Elem() IdentitySourceConfigurationOutput {
	return o.ApplyT(func(v *IdentitySourceConfiguration) IdentitySourceConfiguration {
		if v != nil {
			return *v
		}
		var ret IdentitySourceConfiguration
		return ret
	}).(IdentitySourceConfigurationOutput)
}

func (o IdentitySourceConfigurationPtrOutput) CognitoUserPoolConfiguration() IdentitySourceCognitoUserPoolConfigurationPtrOutput {
	return o.ApplyT(func(v *IdentitySourceConfiguration) *IdentitySourceCognitoUserPoolConfiguration {
		if v == nil {
			return nil
		}
		return &v.CognitoUserPoolConfiguration
	}).(IdentitySourceCognitoUserPoolConfigurationPtrOutput)
}

type IdentitySourceDetails struct {
	ClientIds    []string                    `pulumi:"clientIds"`
	DiscoveryUrl *string                     `pulumi:"discoveryUrl"`
	OpenIdIssuer *IdentitySourceOpenIdIssuer `pulumi:"openIdIssuer"`
	UserPoolArn  *string                     `pulumi:"userPoolArn"`
}

type IdentitySourceDetailsOutput struct{ *pulumi.OutputState }

func (IdentitySourceDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentitySourceDetails)(nil)).Elem()
}

func (o IdentitySourceDetailsOutput) ToIdentitySourceDetailsOutput() IdentitySourceDetailsOutput {
	return o
}

func (o IdentitySourceDetailsOutput) ToIdentitySourceDetailsOutputWithContext(ctx context.Context) IdentitySourceDetailsOutput {
	return o
}

func (o IdentitySourceDetailsOutput) ClientIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IdentitySourceDetails) []string { return v.ClientIds }).(pulumi.StringArrayOutput)
}

func (o IdentitySourceDetailsOutput) DiscoveryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceDetails) *string { return v.DiscoveryUrl }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceDetailsOutput) OpenIdIssuer() IdentitySourceOpenIdIssuerPtrOutput {
	return o.ApplyT(func(v IdentitySourceDetails) *IdentitySourceOpenIdIssuer { return v.OpenIdIssuer }).(IdentitySourceOpenIdIssuerPtrOutput)
}

func (o IdentitySourceDetailsOutput) UserPoolArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceDetails) *string { return v.UserPoolArn }).(pulumi.StringPtrOutput)
}

type IdentitySourceDetailsPtrOutput struct{ *pulumi.OutputState }

func (IdentitySourceDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentitySourceDetails)(nil)).Elem()
}

func (o IdentitySourceDetailsPtrOutput) ToIdentitySourceDetailsPtrOutput() IdentitySourceDetailsPtrOutput {
	return o
}

func (o IdentitySourceDetailsPtrOutput) ToIdentitySourceDetailsPtrOutputWithContext(ctx context.Context) IdentitySourceDetailsPtrOutput {
	return o
}

func (o IdentitySourceDetailsPtrOutput) Elem() IdentitySourceDetailsOutput {
	return o.ApplyT(func(v *IdentitySourceDetails) IdentitySourceDetails {
		if v != nil {
			return *v
		}
		var ret IdentitySourceDetails
		return ret
	}).(IdentitySourceDetailsOutput)
}

func (o IdentitySourceDetailsPtrOutput) ClientIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IdentitySourceDetails) []string {
		if v == nil {
			return nil
		}
		return v.ClientIds
	}).(pulumi.StringArrayOutput)
}

func (o IdentitySourceDetailsPtrOutput) DiscoveryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentitySourceDetails) *string {
		if v == nil {
			return nil
		}
		return v.DiscoveryUrl
	}).(pulumi.StringPtrOutput)
}

func (o IdentitySourceDetailsPtrOutput) OpenIdIssuer() IdentitySourceOpenIdIssuerPtrOutput {
	return o.ApplyT(func(v *IdentitySourceDetails) *IdentitySourceOpenIdIssuer {
		if v == nil {
			return nil
		}
		return v.OpenIdIssuer
	}).(IdentitySourceOpenIdIssuerPtrOutput)
}

func (o IdentitySourceDetailsPtrOutput) UserPoolArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentitySourceDetails) *string {
		if v == nil {
			return nil
		}
		return v.UserPoolArn
	}).(pulumi.StringPtrOutput)
}

type PolicyDefinition struct {
}

// PolicyDefinitionInput is an input type that accepts PolicyDefinitionArgs and PolicyDefinitionOutput values.
// You can construct a concrete instance of `PolicyDefinitionInput` via:
//
//	PolicyDefinitionArgs{...}
type PolicyDefinitionInput interface {
	pulumi.Input

	ToPolicyDefinitionOutput() PolicyDefinitionOutput
	ToPolicyDefinitionOutputWithContext(context.Context) PolicyDefinitionOutput
}

type PolicyDefinitionArgs struct {
}

func (PolicyDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyDefinition)(nil)).Elem()
}

func (i PolicyDefinitionArgs) ToPolicyDefinitionOutput() PolicyDefinitionOutput {
	return i.ToPolicyDefinitionOutputWithContext(context.Background())
}

func (i PolicyDefinitionArgs) ToPolicyDefinitionOutputWithContext(ctx context.Context) PolicyDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyDefinitionOutput)
}

type PolicyDefinitionOutput struct{ *pulumi.OutputState }

func (PolicyDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyDefinition)(nil)).Elem()
}

func (o PolicyDefinitionOutput) ToPolicyDefinitionOutput() PolicyDefinitionOutput {
	return o
}

func (o PolicyDefinitionOutput) ToPolicyDefinitionOutputWithContext(ctx context.Context) PolicyDefinitionOutput {
	return o
}

type PolicyDefinitionPtrOutput struct{ *pulumi.OutputState }

func (PolicyDefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyDefinition)(nil)).Elem()
}

func (o PolicyDefinitionPtrOutput) ToPolicyDefinitionPtrOutput() PolicyDefinitionPtrOutput {
	return o
}

func (o PolicyDefinitionPtrOutput) ToPolicyDefinitionPtrOutputWithContext(ctx context.Context) PolicyDefinitionPtrOutput {
	return o
}

func (o PolicyDefinitionPtrOutput) Elem() PolicyDefinitionOutput {
	return o.ApplyT(func(v *PolicyDefinition) PolicyDefinition {
		if v != nil {
			return *v
		}
		var ret PolicyDefinition
		return ret
	}).(PolicyDefinitionOutput)
}

type PolicyStoreSchemaDefinition struct {
	CedarJson *string `pulumi:"cedarJson"`
}

// PolicyStoreSchemaDefinitionInput is an input type that accepts PolicyStoreSchemaDefinitionArgs and PolicyStoreSchemaDefinitionOutput values.
// You can construct a concrete instance of `PolicyStoreSchemaDefinitionInput` via:
//
//	PolicyStoreSchemaDefinitionArgs{...}
type PolicyStoreSchemaDefinitionInput interface {
	pulumi.Input

	ToPolicyStoreSchemaDefinitionOutput() PolicyStoreSchemaDefinitionOutput
	ToPolicyStoreSchemaDefinitionOutputWithContext(context.Context) PolicyStoreSchemaDefinitionOutput
}

type PolicyStoreSchemaDefinitionArgs struct {
	CedarJson pulumi.StringPtrInput `pulumi:"cedarJson"`
}

func (PolicyStoreSchemaDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyStoreSchemaDefinition)(nil)).Elem()
}

func (i PolicyStoreSchemaDefinitionArgs) ToPolicyStoreSchemaDefinitionOutput() PolicyStoreSchemaDefinitionOutput {
	return i.ToPolicyStoreSchemaDefinitionOutputWithContext(context.Background())
}

func (i PolicyStoreSchemaDefinitionArgs) ToPolicyStoreSchemaDefinitionOutputWithContext(ctx context.Context) PolicyStoreSchemaDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyStoreSchemaDefinitionOutput)
}

func (i PolicyStoreSchemaDefinitionArgs) ToPolicyStoreSchemaDefinitionPtrOutput() PolicyStoreSchemaDefinitionPtrOutput {
	return i.ToPolicyStoreSchemaDefinitionPtrOutputWithContext(context.Background())
}

func (i PolicyStoreSchemaDefinitionArgs) ToPolicyStoreSchemaDefinitionPtrOutputWithContext(ctx context.Context) PolicyStoreSchemaDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyStoreSchemaDefinitionOutput).ToPolicyStoreSchemaDefinitionPtrOutputWithContext(ctx)
}

// PolicyStoreSchemaDefinitionPtrInput is an input type that accepts PolicyStoreSchemaDefinitionArgs, PolicyStoreSchemaDefinitionPtr and PolicyStoreSchemaDefinitionPtrOutput values.
// You can construct a concrete instance of `PolicyStoreSchemaDefinitionPtrInput` via:
//
//	        PolicyStoreSchemaDefinitionArgs{...}
//
//	or:
//
//	        nil
type PolicyStoreSchemaDefinitionPtrInput interface {
	pulumi.Input

	ToPolicyStoreSchemaDefinitionPtrOutput() PolicyStoreSchemaDefinitionPtrOutput
	ToPolicyStoreSchemaDefinitionPtrOutputWithContext(context.Context) PolicyStoreSchemaDefinitionPtrOutput
}

type policyStoreSchemaDefinitionPtrType PolicyStoreSchemaDefinitionArgs

func PolicyStoreSchemaDefinitionPtr(v *PolicyStoreSchemaDefinitionArgs) PolicyStoreSchemaDefinitionPtrInput {
	return (*policyStoreSchemaDefinitionPtrType)(v)
}

func (*policyStoreSchemaDefinitionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyStoreSchemaDefinition)(nil)).Elem()
}

func (i *policyStoreSchemaDefinitionPtrType) ToPolicyStoreSchemaDefinitionPtrOutput() PolicyStoreSchemaDefinitionPtrOutput {
	return i.ToPolicyStoreSchemaDefinitionPtrOutputWithContext(context.Background())
}

func (i *policyStoreSchemaDefinitionPtrType) ToPolicyStoreSchemaDefinitionPtrOutputWithContext(ctx context.Context) PolicyStoreSchemaDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyStoreSchemaDefinitionPtrOutput)
}

type PolicyStoreSchemaDefinitionOutput struct{ *pulumi.OutputState }

func (PolicyStoreSchemaDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyStoreSchemaDefinition)(nil)).Elem()
}

func (o PolicyStoreSchemaDefinitionOutput) ToPolicyStoreSchemaDefinitionOutput() PolicyStoreSchemaDefinitionOutput {
	return o
}

func (o PolicyStoreSchemaDefinitionOutput) ToPolicyStoreSchemaDefinitionOutputWithContext(ctx context.Context) PolicyStoreSchemaDefinitionOutput {
	return o
}

func (o PolicyStoreSchemaDefinitionOutput) ToPolicyStoreSchemaDefinitionPtrOutput() PolicyStoreSchemaDefinitionPtrOutput {
	return o.ToPolicyStoreSchemaDefinitionPtrOutputWithContext(context.Background())
}

func (o PolicyStoreSchemaDefinitionOutput) ToPolicyStoreSchemaDefinitionPtrOutputWithContext(ctx context.Context) PolicyStoreSchemaDefinitionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PolicyStoreSchemaDefinition) *PolicyStoreSchemaDefinition {
		return &v
	}).(PolicyStoreSchemaDefinitionPtrOutput)
}

func (o PolicyStoreSchemaDefinitionOutput) CedarJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyStoreSchemaDefinition) *string { return v.CedarJson }).(pulumi.StringPtrOutput)
}

type PolicyStoreSchemaDefinitionPtrOutput struct{ *pulumi.OutputState }

func (PolicyStoreSchemaDefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyStoreSchemaDefinition)(nil)).Elem()
}

func (o PolicyStoreSchemaDefinitionPtrOutput) ToPolicyStoreSchemaDefinitionPtrOutput() PolicyStoreSchemaDefinitionPtrOutput {
	return o
}

func (o PolicyStoreSchemaDefinitionPtrOutput) ToPolicyStoreSchemaDefinitionPtrOutputWithContext(ctx context.Context) PolicyStoreSchemaDefinitionPtrOutput {
	return o
}

func (o PolicyStoreSchemaDefinitionPtrOutput) Elem() PolicyStoreSchemaDefinitionOutput {
	return o.ApplyT(func(v *PolicyStoreSchemaDefinition) PolicyStoreSchemaDefinition {
		if v != nil {
			return *v
		}
		var ret PolicyStoreSchemaDefinition
		return ret
	}).(PolicyStoreSchemaDefinitionOutput)
}

func (o PolicyStoreSchemaDefinitionPtrOutput) CedarJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyStoreSchemaDefinition) *string {
		if v == nil {
			return nil
		}
		return v.CedarJson
	}).(pulumi.StringPtrOutput)
}

type PolicyStoreValidationSettings struct {
	Mode PolicyStoreValidationMode `pulumi:"mode"`
}

// PolicyStoreValidationSettingsInput is an input type that accepts PolicyStoreValidationSettingsArgs and PolicyStoreValidationSettingsOutput values.
// You can construct a concrete instance of `PolicyStoreValidationSettingsInput` via:
//
//	PolicyStoreValidationSettingsArgs{...}
type PolicyStoreValidationSettingsInput interface {
	pulumi.Input

	ToPolicyStoreValidationSettingsOutput() PolicyStoreValidationSettingsOutput
	ToPolicyStoreValidationSettingsOutputWithContext(context.Context) PolicyStoreValidationSettingsOutput
}

type PolicyStoreValidationSettingsArgs struct {
	Mode PolicyStoreValidationModeInput `pulumi:"mode"`
}

func (PolicyStoreValidationSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyStoreValidationSettings)(nil)).Elem()
}

func (i PolicyStoreValidationSettingsArgs) ToPolicyStoreValidationSettingsOutput() PolicyStoreValidationSettingsOutput {
	return i.ToPolicyStoreValidationSettingsOutputWithContext(context.Background())
}

func (i PolicyStoreValidationSettingsArgs) ToPolicyStoreValidationSettingsOutputWithContext(ctx context.Context) PolicyStoreValidationSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyStoreValidationSettingsOutput)
}

type PolicyStoreValidationSettingsOutput struct{ *pulumi.OutputState }

func (PolicyStoreValidationSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyStoreValidationSettings)(nil)).Elem()
}

func (o PolicyStoreValidationSettingsOutput) ToPolicyStoreValidationSettingsOutput() PolicyStoreValidationSettingsOutput {
	return o
}

func (o PolicyStoreValidationSettingsOutput) ToPolicyStoreValidationSettingsOutputWithContext(ctx context.Context) PolicyStoreValidationSettingsOutput {
	return o
}

func (o PolicyStoreValidationSettingsOutput) Mode() PolicyStoreValidationModeOutput {
	return o.ApplyT(func(v PolicyStoreValidationSettings) PolicyStoreValidationMode { return v.Mode }).(PolicyStoreValidationModeOutput)
}

type PolicyStoreValidationSettingsPtrOutput struct{ *pulumi.OutputState }

func (PolicyStoreValidationSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyStoreValidationSettings)(nil)).Elem()
}

func (o PolicyStoreValidationSettingsPtrOutput) ToPolicyStoreValidationSettingsPtrOutput() PolicyStoreValidationSettingsPtrOutput {
	return o
}

func (o PolicyStoreValidationSettingsPtrOutput) ToPolicyStoreValidationSettingsPtrOutputWithContext(ctx context.Context) PolicyStoreValidationSettingsPtrOutput {
	return o
}

func (o PolicyStoreValidationSettingsPtrOutput) Elem() PolicyStoreValidationSettingsOutput {
	return o.ApplyT(func(v *PolicyStoreValidationSettings) PolicyStoreValidationSettings {
		if v != nil {
			return *v
		}
		var ret PolicyStoreValidationSettings
		return ret
	}).(PolicyStoreValidationSettingsOutput)
}

func (o PolicyStoreValidationSettingsPtrOutput) Mode() PolicyStoreValidationModePtrOutput {
	return o.ApplyT(func(v *PolicyStoreValidationSettings) *PolicyStoreValidationMode {
		if v == nil {
			return nil
		}
		return &v.Mode
	}).(PolicyStoreValidationModePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IdentitySourceCognitoUserPoolConfigurationInput)(nil)).Elem(), IdentitySourceCognitoUserPoolConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentitySourceConfigurationInput)(nil)).Elem(), IdentitySourceConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyDefinitionInput)(nil)).Elem(), PolicyDefinitionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyStoreSchemaDefinitionInput)(nil)).Elem(), PolicyStoreSchemaDefinitionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyStoreSchemaDefinitionPtrInput)(nil)).Elem(), PolicyStoreSchemaDefinitionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyStoreValidationSettingsInput)(nil)).Elem(), PolicyStoreValidationSettingsArgs{})
	pulumi.RegisterOutputType(IdentitySourceCognitoUserPoolConfigurationOutput{})
	pulumi.RegisterOutputType(IdentitySourceCognitoUserPoolConfigurationPtrOutput{})
	pulumi.RegisterOutputType(IdentitySourceConfigurationOutput{})
	pulumi.RegisterOutputType(IdentitySourceConfigurationPtrOutput{})
	pulumi.RegisterOutputType(IdentitySourceDetailsOutput{})
	pulumi.RegisterOutputType(IdentitySourceDetailsPtrOutput{})
	pulumi.RegisterOutputType(PolicyDefinitionOutput{})
	pulumi.RegisterOutputType(PolicyDefinitionPtrOutput{})
	pulumi.RegisterOutputType(PolicyStoreSchemaDefinitionOutput{})
	pulumi.RegisterOutputType(PolicyStoreSchemaDefinitionPtrOutput{})
	pulumi.RegisterOutputType(PolicyStoreValidationSettingsOutput{})
	pulumi.RegisterOutputType(PolicyStoreValidationSettingsPtrOutput{})
}
