// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Cognito::UserPool
func LookupUserPool(ctx *pulumi.Context, args *LookupUserPoolArgs, opts ...pulumi.InvokeOption) (*LookupUserPoolResult, error) {
	var rv LookupUserPoolResult
	err := ctx.Invoke("aws-native:cognito:getUserPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserPoolArgs struct {
	Id string `pulumi:"id"`
}

type LookupUserPoolResult struct {
	AccountRecoverySetting      *UserPoolAccountRecoverySetting      `pulumi:"accountRecoverySetting"`
	AdminCreateUserConfig       *UserPoolAdminCreateUserConfig       `pulumi:"adminCreateUserConfig"`
	AliasAttributes             []string                             `pulumi:"aliasAttributes"`
	Arn                         *string                              `pulumi:"arn"`
	AutoVerifiedAttributes      []string                             `pulumi:"autoVerifiedAttributes"`
	DeviceConfiguration         *UserPoolDeviceConfiguration         `pulumi:"deviceConfiguration"`
	EmailConfiguration          *UserPoolEmailConfiguration          `pulumi:"emailConfiguration"`
	EmailVerificationMessage    *string                              `pulumi:"emailVerificationMessage"`
	EmailVerificationSubject    *string                              `pulumi:"emailVerificationSubject"`
	EnabledMfas                 []string                             `pulumi:"enabledMfas"`
	Id                          *string                              `pulumi:"id"`
	LambdaConfig                *UserPoolLambdaConfig                `pulumi:"lambdaConfig"`
	MfaConfiguration            *string                              `pulumi:"mfaConfiguration"`
	Policies                    *UserPoolPolicies                    `pulumi:"policies"`
	ProviderName                *string                              `pulumi:"providerName"`
	ProviderURL                 *string                              `pulumi:"providerURL"`
	Schema                      []UserPoolSchemaAttribute            `pulumi:"schema"`
	SmsAuthenticationMessage    *string                              `pulumi:"smsAuthenticationMessage"`
	SmsConfiguration            *UserPoolSmsConfiguration            `pulumi:"smsConfiguration"`
	SmsVerificationMessage      *string                              `pulumi:"smsVerificationMessage"`
	UserPoolAddOns              *UserPoolAddOns                      `pulumi:"userPoolAddOns"`
	UserPoolName                *string                              `pulumi:"userPoolName"`
	UserPoolTags                interface{}                          `pulumi:"userPoolTags"`
	UsernameAttributes          []string                             `pulumi:"usernameAttributes"`
	UsernameConfiguration       *UserPoolUsernameConfiguration       `pulumi:"usernameConfiguration"`
	VerificationMessageTemplate *UserPoolVerificationMessageTemplate `pulumi:"verificationMessageTemplate"`
}

func LookupUserPoolOutput(ctx *pulumi.Context, args LookupUserPoolOutputArgs, opts ...pulumi.InvokeOption) LookupUserPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserPoolResult, error) {
			args := v.(LookupUserPoolArgs)
			r, err := LookupUserPool(ctx, &args, opts...)
			return *r, err
		}).(LookupUserPoolResultOutput)
}

type LookupUserPoolOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupUserPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserPoolArgs)(nil)).Elem()
}

type LookupUserPoolResultOutput struct{ *pulumi.OutputState }

func (LookupUserPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserPoolResult)(nil)).Elem()
}

func (o LookupUserPoolResultOutput) ToLookupUserPoolResultOutput() LookupUserPoolResultOutput {
	return o
}

func (o LookupUserPoolResultOutput) ToLookupUserPoolResultOutputWithContext(ctx context.Context) LookupUserPoolResultOutput {
	return o
}

func (o LookupUserPoolResultOutput) AccountRecoverySetting() UserPoolAccountRecoverySettingPtrOutput {
	return o.ApplyT(func(v LookupUserPoolResult) *UserPoolAccountRecoverySetting { return v.AccountRecoverySetting }).(UserPoolAccountRecoverySettingPtrOutput)
}

func (o LookupUserPoolResultOutput) AdminCreateUserConfig() UserPoolAdminCreateUserConfigPtrOutput {
	return o.ApplyT(func(v LookupUserPoolResult) *UserPoolAdminCreateUserConfig { return v.AdminCreateUserConfig }).(UserPoolAdminCreateUserConfigPtrOutput)
}

func (o LookupUserPoolResultOutput) AliasAttributes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUserPoolResult) []string { return v.AliasAttributes }).(pulumi.StringArrayOutput)
}

func (o LookupUserPoolResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserPoolResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupUserPoolResultOutput) AutoVerifiedAttributes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUserPoolResult) []string { return v.AutoVerifiedAttributes }).(pulumi.StringArrayOutput)
}

func (o LookupUserPoolResultOutput) DeviceConfiguration() UserPoolDeviceConfigurationPtrOutput {
	return o.ApplyT(func(v LookupUserPoolResult) *UserPoolDeviceConfiguration { return v.DeviceConfiguration }).(UserPoolDeviceConfigurationPtrOutput)
}

func (o LookupUserPoolResultOutput) EmailConfiguration() UserPoolEmailConfigurationPtrOutput {
	return o.ApplyT(func(v LookupUserPoolResult) *UserPoolEmailConfiguration { return v.EmailConfiguration }).(UserPoolEmailConfigurationPtrOutput)
}

func (o LookupUserPoolResultOutput) EmailVerificationMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserPoolResult) *string { return v.EmailVerificationMessage }).(pulumi.StringPtrOutput)
}

func (o LookupUserPoolResultOutput) EmailVerificationSubject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserPoolResult) *string { return v.EmailVerificationSubject }).(pulumi.StringPtrOutput)
}

func (o LookupUserPoolResultOutput) EnabledMfas() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUserPoolResult) []string { return v.EnabledMfas }).(pulumi.StringArrayOutput)
}

func (o LookupUserPoolResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserPoolResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupUserPoolResultOutput) LambdaConfig() UserPoolLambdaConfigPtrOutput {
	return o.ApplyT(func(v LookupUserPoolResult) *UserPoolLambdaConfig { return v.LambdaConfig }).(UserPoolLambdaConfigPtrOutput)
}

func (o LookupUserPoolResultOutput) MfaConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserPoolResult) *string { return v.MfaConfiguration }).(pulumi.StringPtrOutput)
}

func (o LookupUserPoolResultOutput) Policies() UserPoolPoliciesPtrOutput {
	return o.ApplyT(func(v LookupUserPoolResult) *UserPoolPolicies { return v.Policies }).(UserPoolPoliciesPtrOutput)
}

func (o LookupUserPoolResultOutput) ProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserPoolResult) *string { return v.ProviderName }).(pulumi.StringPtrOutput)
}

func (o LookupUserPoolResultOutput) ProviderURL() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserPoolResult) *string { return v.ProviderURL }).(pulumi.StringPtrOutput)
}

func (o LookupUserPoolResultOutput) Schema() UserPoolSchemaAttributeArrayOutput {
	return o.ApplyT(func(v LookupUserPoolResult) []UserPoolSchemaAttribute { return v.Schema }).(UserPoolSchemaAttributeArrayOutput)
}

func (o LookupUserPoolResultOutput) SmsAuthenticationMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserPoolResult) *string { return v.SmsAuthenticationMessage }).(pulumi.StringPtrOutput)
}

func (o LookupUserPoolResultOutput) SmsConfiguration() UserPoolSmsConfigurationPtrOutput {
	return o.ApplyT(func(v LookupUserPoolResult) *UserPoolSmsConfiguration { return v.SmsConfiguration }).(UserPoolSmsConfigurationPtrOutput)
}

func (o LookupUserPoolResultOutput) SmsVerificationMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserPoolResult) *string { return v.SmsVerificationMessage }).(pulumi.StringPtrOutput)
}

func (o LookupUserPoolResultOutput) UserPoolAddOns() UserPoolAddOnsPtrOutput {
	return o.ApplyT(func(v LookupUserPoolResult) *UserPoolAddOns { return v.UserPoolAddOns }).(UserPoolAddOnsPtrOutput)
}

func (o LookupUserPoolResultOutput) UserPoolName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserPoolResult) *string { return v.UserPoolName }).(pulumi.StringPtrOutput)
}

func (o LookupUserPoolResultOutput) UserPoolTags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupUserPoolResult) interface{} { return v.UserPoolTags }).(pulumi.AnyOutput)
}

func (o LookupUserPoolResultOutput) UsernameAttributes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUserPoolResult) []string { return v.UsernameAttributes }).(pulumi.StringArrayOutput)
}

func (o LookupUserPoolResultOutput) UsernameConfiguration() UserPoolUsernameConfigurationPtrOutput {
	return o.ApplyT(func(v LookupUserPoolResult) *UserPoolUsernameConfiguration { return v.UsernameConfiguration }).(UserPoolUsernameConfigurationPtrOutput)
}

func (o LookupUserPoolResultOutput) VerificationMessageTemplate() UserPoolVerificationMessageTemplatePtrOutput {
	return o.ApplyT(func(v LookupUserPoolResult) *UserPoolVerificationMessageTemplate {
		return v.VerificationMessageTemplate
	}).(UserPoolVerificationMessageTemplatePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserPoolResultOutput{})
}
