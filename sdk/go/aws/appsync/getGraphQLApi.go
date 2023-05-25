// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appsync

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppSync::GraphQLApi
func LookupGraphQLApi(ctx *pulumi.Context, args *LookupGraphQLApiArgs, opts ...pulumi.InvokeOption) (*LookupGraphQLApiResult, error) {
	var rv LookupGraphQLApiResult
	err := ctx.Invoke("aws-native:appsync:getGraphQLApi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGraphQLApiArgs struct {
	Id string `pulumi:"id"`
}

type LookupGraphQLApiResult struct {
	AdditionalAuthenticationProviders []GraphQLApiAdditionalAuthenticationProvider `pulumi:"additionalAuthenticationProviders"`
	ApiId                             *string                                      `pulumi:"apiId"`
	Arn                               *string                                      `pulumi:"arn"`
	AuthenticationType                *string                                      `pulumi:"authenticationType"`
	GraphQLDns                        *string                                      `pulumi:"graphQLDns"`
	GraphQLUrl                        *string                                      `pulumi:"graphQLUrl"`
	Id                                *string                                      `pulumi:"id"`
	LambdaAuthorizerConfig            *GraphQLApiLambdaAuthorizerConfig            `pulumi:"lambdaAuthorizerConfig"`
	LogConfig                         *GraphQLApiLogConfig                         `pulumi:"logConfig"`
	MergedApiExecutionRoleArn         *string                                      `pulumi:"mergedApiExecutionRoleArn"`
	Name                              *string                                      `pulumi:"name"`
	OpenIDConnectConfig               *GraphQLApiOpenIDConnectConfig               `pulumi:"openIDConnectConfig"`
	OwnerContact                      *string                                      `pulumi:"ownerContact"`
	RealtimeDns                       *string                                      `pulumi:"realtimeDns"`
	RealtimeUrl                       *string                                      `pulumi:"realtimeUrl"`
	Tags                              []GraphQLApiTag                              `pulumi:"tags"`
	UserPoolConfig                    *GraphQLApiUserPoolConfig                    `pulumi:"userPoolConfig"`
	XrayEnabled                       *bool                                        `pulumi:"xrayEnabled"`
}

func LookupGraphQLApiOutput(ctx *pulumi.Context, args LookupGraphQLApiOutputArgs, opts ...pulumi.InvokeOption) LookupGraphQLApiResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGraphQLApiResult, error) {
			args := v.(LookupGraphQLApiArgs)
			r, err := LookupGraphQLApi(ctx, &args, opts...)
			var s LookupGraphQLApiResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGraphQLApiResultOutput)
}

type LookupGraphQLApiOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupGraphQLApiOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGraphQLApiArgs)(nil)).Elem()
}

type LookupGraphQLApiResultOutput struct{ *pulumi.OutputState }

func (LookupGraphQLApiResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGraphQLApiResult)(nil)).Elem()
}

func (o LookupGraphQLApiResultOutput) ToLookupGraphQLApiResultOutput() LookupGraphQLApiResultOutput {
	return o
}

func (o LookupGraphQLApiResultOutput) ToLookupGraphQLApiResultOutputWithContext(ctx context.Context) LookupGraphQLApiResultOutput {
	return o
}

func (o LookupGraphQLApiResultOutput) AdditionalAuthenticationProviders() GraphQLApiAdditionalAuthenticationProviderArrayOutput {
	return o.ApplyT(func(v LookupGraphQLApiResult) []GraphQLApiAdditionalAuthenticationProvider {
		return v.AdditionalAuthenticationProviders
	}).(GraphQLApiAdditionalAuthenticationProviderArrayOutput)
}

func (o LookupGraphQLApiResultOutput) ApiId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGraphQLApiResult) *string { return v.ApiId }).(pulumi.StringPtrOutput)
}

func (o LookupGraphQLApiResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGraphQLApiResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupGraphQLApiResultOutput) AuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGraphQLApiResult) *string { return v.AuthenticationType }).(pulumi.StringPtrOutput)
}

func (o LookupGraphQLApiResultOutput) GraphQLDns() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGraphQLApiResult) *string { return v.GraphQLDns }).(pulumi.StringPtrOutput)
}

func (o LookupGraphQLApiResultOutput) GraphQLUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGraphQLApiResult) *string { return v.GraphQLUrl }).(pulumi.StringPtrOutput)
}

func (o LookupGraphQLApiResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGraphQLApiResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupGraphQLApiResultOutput) LambdaAuthorizerConfig() GraphQLApiLambdaAuthorizerConfigPtrOutput {
	return o.ApplyT(func(v LookupGraphQLApiResult) *GraphQLApiLambdaAuthorizerConfig { return v.LambdaAuthorizerConfig }).(GraphQLApiLambdaAuthorizerConfigPtrOutput)
}

func (o LookupGraphQLApiResultOutput) LogConfig() GraphQLApiLogConfigPtrOutput {
	return o.ApplyT(func(v LookupGraphQLApiResult) *GraphQLApiLogConfig { return v.LogConfig }).(GraphQLApiLogConfigPtrOutput)
}

func (o LookupGraphQLApiResultOutput) MergedApiExecutionRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGraphQLApiResult) *string { return v.MergedApiExecutionRoleArn }).(pulumi.StringPtrOutput)
}

func (o LookupGraphQLApiResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGraphQLApiResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupGraphQLApiResultOutput) OpenIDConnectConfig() GraphQLApiOpenIDConnectConfigPtrOutput {
	return o.ApplyT(func(v LookupGraphQLApiResult) *GraphQLApiOpenIDConnectConfig { return v.OpenIDConnectConfig }).(GraphQLApiOpenIDConnectConfigPtrOutput)
}

func (o LookupGraphQLApiResultOutput) OwnerContact() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGraphQLApiResult) *string { return v.OwnerContact }).(pulumi.StringPtrOutput)
}

func (o LookupGraphQLApiResultOutput) RealtimeDns() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGraphQLApiResult) *string { return v.RealtimeDns }).(pulumi.StringPtrOutput)
}

func (o LookupGraphQLApiResultOutput) RealtimeUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGraphQLApiResult) *string { return v.RealtimeUrl }).(pulumi.StringPtrOutput)
}

func (o LookupGraphQLApiResultOutput) Tags() GraphQLApiTagArrayOutput {
	return o.ApplyT(func(v LookupGraphQLApiResult) []GraphQLApiTag { return v.Tags }).(GraphQLApiTagArrayOutput)
}

func (o LookupGraphQLApiResultOutput) UserPoolConfig() GraphQLApiUserPoolConfigPtrOutput {
	return o.ApplyT(func(v LookupGraphQLApiResult) *GraphQLApiUserPoolConfig { return v.UserPoolConfig }).(GraphQLApiUserPoolConfigPtrOutput)
}

func (o LookupGraphQLApiResultOutput) XrayEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGraphQLApiResult) *bool { return v.XrayEnabled }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGraphQLApiResultOutput{})
}
