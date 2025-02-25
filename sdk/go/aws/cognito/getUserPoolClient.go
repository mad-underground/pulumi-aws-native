// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Cognito::UserPoolClient
func LookupUserPoolClient(ctx *pulumi.Context, args *LookupUserPoolClientArgs, opts ...pulumi.InvokeOption) (*LookupUserPoolClientResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupUserPoolClientResult
	err := ctx.Invoke("aws-native:cognito:getUserPoolClient", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserPoolClientArgs struct {
	ClientId   string `pulumi:"clientId"`
	UserPoolId string `pulumi:"userPoolId"`
}

type LookupUserPoolClientResult struct {
	AccessTokenValidity                      *int                                  `pulumi:"accessTokenValidity"`
	AllowedOAuthFlows                        []string                              `pulumi:"allowedOAuthFlows"`
	AllowedOAuthFlowsUserPoolClient          *bool                                 `pulumi:"allowedOAuthFlowsUserPoolClient"`
	AllowedOAuthScopes                       []string                              `pulumi:"allowedOAuthScopes"`
	AnalyticsConfiguration                   *UserPoolClientAnalyticsConfiguration `pulumi:"analyticsConfiguration"`
	AuthSessionValidity                      *int                                  `pulumi:"authSessionValidity"`
	CallbackUrls                             []string                              `pulumi:"callbackUrls"`
	ClientId                                 *string                               `pulumi:"clientId"`
	ClientName                               *string                               `pulumi:"clientName"`
	ClientSecret                             *string                               `pulumi:"clientSecret"`
	DefaultRedirectUri                       *string                               `pulumi:"defaultRedirectUri"`
	EnablePropagateAdditionalUserContextData *bool                                 `pulumi:"enablePropagateAdditionalUserContextData"`
	EnableTokenRevocation                    *bool                                 `pulumi:"enableTokenRevocation"`
	ExplicitAuthFlows                        []string                              `pulumi:"explicitAuthFlows"`
	IdTokenValidity                          *int                                  `pulumi:"idTokenValidity"`
	LogoutUrls                               []string                              `pulumi:"logoutUrls"`
	Name                                     *string                               `pulumi:"name"`
	PreventUserExistenceErrors               *string                               `pulumi:"preventUserExistenceErrors"`
	ReadAttributes                           []string                              `pulumi:"readAttributes"`
	RefreshTokenValidity                     *int                                  `pulumi:"refreshTokenValidity"`
	SupportedIdentityProviders               []string                              `pulumi:"supportedIdentityProviders"`
	TokenValidityUnits                       *UserPoolClientTokenValidityUnits     `pulumi:"tokenValidityUnits"`
	WriteAttributes                          []string                              `pulumi:"writeAttributes"`
}

func LookupUserPoolClientOutput(ctx *pulumi.Context, args LookupUserPoolClientOutputArgs, opts ...pulumi.InvokeOption) LookupUserPoolClientResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserPoolClientResult, error) {
			args := v.(LookupUserPoolClientArgs)
			r, err := LookupUserPoolClient(ctx, &args, opts...)
			var s LookupUserPoolClientResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserPoolClientResultOutput)
}

type LookupUserPoolClientOutputArgs struct {
	ClientId   pulumi.StringInput `pulumi:"clientId"`
	UserPoolId pulumi.StringInput `pulumi:"userPoolId"`
}

func (LookupUserPoolClientOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserPoolClientArgs)(nil)).Elem()
}

type LookupUserPoolClientResultOutput struct{ *pulumi.OutputState }

func (LookupUserPoolClientResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserPoolClientResult)(nil)).Elem()
}

func (o LookupUserPoolClientResultOutput) ToLookupUserPoolClientResultOutput() LookupUserPoolClientResultOutput {
	return o
}

func (o LookupUserPoolClientResultOutput) ToLookupUserPoolClientResultOutputWithContext(ctx context.Context) LookupUserPoolClientResultOutput {
	return o
}

func (o LookupUserPoolClientResultOutput) AccessTokenValidity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupUserPoolClientResult) *int { return v.AccessTokenValidity }).(pulumi.IntPtrOutput)
}

func (o LookupUserPoolClientResultOutput) AllowedOAuthFlows() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUserPoolClientResult) []string { return v.AllowedOAuthFlows }).(pulumi.StringArrayOutput)
}

func (o LookupUserPoolClientResultOutput) AllowedOAuthFlowsUserPoolClient() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupUserPoolClientResult) *bool { return v.AllowedOAuthFlowsUserPoolClient }).(pulumi.BoolPtrOutput)
}

func (o LookupUserPoolClientResultOutput) AllowedOAuthScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUserPoolClientResult) []string { return v.AllowedOAuthScopes }).(pulumi.StringArrayOutput)
}

func (o LookupUserPoolClientResultOutput) AnalyticsConfiguration() UserPoolClientAnalyticsConfigurationPtrOutput {
	return o.ApplyT(func(v LookupUserPoolClientResult) *UserPoolClientAnalyticsConfiguration {
		return v.AnalyticsConfiguration
	}).(UserPoolClientAnalyticsConfigurationPtrOutput)
}

func (o LookupUserPoolClientResultOutput) AuthSessionValidity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupUserPoolClientResult) *int { return v.AuthSessionValidity }).(pulumi.IntPtrOutput)
}

func (o LookupUserPoolClientResultOutput) CallbackUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUserPoolClientResult) []string { return v.CallbackUrls }).(pulumi.StringArrayOutput)
}

func (o LookupUserPoolClientResultOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserPoolClientResult) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o LookupUserPoolClientResultOutput) ClientName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserPoolClientResult) *string { return v.ClientName }).(pulumi.StringPtrOutput)
}

func (o LookupUserPoolClientResultOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserPoolClientResult) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o LookupUserPoolClientResultOutput) DefaultRedirectUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserPoolClientResult) *string { return v.DefaultRedirectUri }).(pulumi.StringPtrOutput)
}

func (o LookupUserPoolClientResultOutput) EnablePropagateAdditionalUserContextData() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupUserPoolClientResult) *bool { return v.EnablePropagateAdditionalUserContextData }).(pulumi.BoolPtrOutput)
}

func (o LookupUserPoolClientResultOutput) EnableTokenRevocation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupUserPoolClientResult) *bool { return v.EnableTokenRevocation }).(pulumi.BoolPtrOutput)
}

func (o LookupUserPoolClientResultOutput) ExplicitAuthFlows() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUserPoolClientResult) []string { return v.ExplicitAuthFlows }).(pulumi.StringArrayOutput)
}

func (o LookupUserPoolClientResultOutput) IdTokenValidity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupUserPoolClientResult) *int { return v.IdTokenValidity }).(pulumi.IntPtrOutput)
}

func (o LookupUserPoolClientResultOutput) LogoutUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUserPoolClientResult) []string { return v.LogoutUrls }).(pulumi.StringArrayOutput)
}

func (o LookupUserPoolClientResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserPoolClientResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupUserPoolClientResultOutput) PreventUserExistenceErrors() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserPoolClientResult) *string { return v.PreventUserExistenceErrors }).(pulumi.StringPtrOutput)
}

func (o LookupUserPoolClientResultOutput) ReadAttributes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUserPoolClientResult) []string { return v.ReadAttributes }).(pulumi.StringArrayOutput)
}

func (o LookupUserPoolClientResultOutput) RefreshTokenValidity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupUserPoolClientResult) *int { return v.RefreshTokenValidity }).(pulumi.IntPtrOutput)
}

func (o LookupUserPoolClientResultOutput) SupportedIdentityProviders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUserPoolClientResult) []string { return v.SupportedIdentityProviders }).(pulumi.StringArrayOutput)
}

func (o LookupUserPoolClientResultOutput) TokenValidityUnits() UserPoolClientTokenValidityUnitsPtrOutput {
	return o.ApplyT(func(v LookupUserPoolClientResult) *UserPoolClientTokenValidityUnits { return v.TokenValidityUnits }).(UserPoolClientTokenValidityUnitsPtrOutput)
}

func (o LookupUserPoolClientResultOutput) WriteAttributes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUserPoolClientResult) []string { return v.WriteAttributes }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserPoolClientResultOutput{})
}
