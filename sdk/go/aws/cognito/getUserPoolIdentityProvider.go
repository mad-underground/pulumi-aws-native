// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::Cognito::UserPoolIdentityProvider
func LookupUserPoolIdentityProvider(ctx *pulumi.Context, args *LookupUserPoolIdentityProviderArgs, opts ...pulumi.InvokeOption) (*LookupUserPoolIdentityProviderResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupUserPoolIdentityProviderResult
	err := ctx.Invoke("aws-native:cognito:getUserPoolIdentityProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserPoolIdentityProviderArgs struct {
	ProviderName string `pulumi:"providerName"`
	UserPoolId   string `pulumi:"userPoolId"`
}

type LookupUserPoolIdentityProviderResult struct {
	AttributeMapping interface{} `pulumi:"attributeMapping"`
	IdpIdentifiers   []string    `pulumi:"idpIdentifiers"`
	ProviderDetails  interface{} `pulumi:"providerDetails"`
}

func LookupUserPoolIdentityProviderOutput(ctx *pulumi.Context, args LookupUserPoolIdentityProviderOutputArgs, opts ...pulumi.InvokeOption) LookupUserPoolIdentityProviderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserPoolIdentityProviderResult, error) {
			args := v.(LookupUserPoolIdentityProviderArgs)
			r, err := LookupUserPoolIdentityProvider(ctx, &args, opts...)
			var s LookupUserPoolIdentityProviderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserPoolIdentityProviderResultOutput)
}

type LookupUserPoolIdentityProviderOutputArgs struct {
	ProviderName pulumi.StringInput `pulumi:"providerName"`
	UserPoolId   pulumi.StringInput `pulumi:"userPoolId"`
}

func (LookupUserPoolIdentityProviderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserPoolIdentityProviderArgs)(nil)).Elem()
}

type LookupUserPoolIdentityProviderResultOutput struct{ *pulumi.OutputState }

func (LookupUserPoolIdentityProviderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserPoolIdentityProviderResult)(nil)).Elem()
}

func (o LookupUserPoolIdentityProviderResultOutput) ToLookupUserPoolIdentityProviderResultOutput() LookupUserPoolIdentityProviderResultOutput {
	return o
}

func (o LookupUserPoolIdentityProviderResultOutput) ToLookupUserPoolIdentityProviderResultOutputWithContext(ctx context.Context) LookupUserPoolIdentityProviderResultOutput {
	return o
}

func (o LookupUserPoolIdentityProviderResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupUserPoolIdentityProviderResult] {
	return pulumix.Output[LookupUserPoolIdentityProviderResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupUserPoolIdentityProviderResultOutput) AttributeMapping() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupUserPoolIdentityProviderResult) interface{} { return v.AttributeMapping }).(pulumi.AnyOutput)
}

func (o LookupUserPoolIdentityProviderResultOutput) IdpIdentifiers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUserPoolIdentityProviderResult) []string { return v.IdpIdentifiers }).(pulumi.StringArrayOutput)
}

func (o LookupUserPoolIdentityProviderResultOutput) ProviderDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupUserPoolIdentityProviderResult) interface{} { return v.ProviderDetails }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserPoolIdentityProviderResultOutput{})
}
