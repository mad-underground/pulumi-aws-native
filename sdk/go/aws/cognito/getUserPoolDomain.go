// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Cognito::UserPoolDomain
func LookupUserPoolDomain(ctx *pulumi.Context, args *LookupUserPoolDomainArgs, opts ...pulumi.InvokeOption) (*LookupUserPoolDomainResult, error) {
	var rv LookupUserPoolDomainResult
	err := ctx.Invoke("aws-native:cognito:getUserPoolDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserPoolDomainArgs struct {
	Id string `pulumi:"id"`
}

type LookupUserPoolDomainResult struct {
	CloudFrontDistribution *string                               `pulumi:"cloudFrontDistribution"`
	CustomDomainConfig     *UserPoolDomainCustomDomainConfigType `pulumi:"customDomainConfig"`
	Id                     *string                               `pulumi:"id"`
}

func LookupUserPoolDomainOutput(ctx *pulumi.Context, args LookupUserPoolDomainOutputArgs, opts ...pulumi.InvokeOption) LookupUserPoolDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserPoolDomainResult, error) {
			args := v.(LookupUserPoolDomainArgs)
			r, err := LookupUserPoolDomain(ctx, &args, opts...)
			var s LookupUserPoolDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserPoolDomainResultOutput)
}

type LookupUserPoolDomainOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupUserPoolDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserPoolDomainArgs)(nil)).Elem()
}

type LookupUserPoolDomainResultOutput struct{ *pulumi.OutputState }

func (LookupUserPoolDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserPoolDomainResult)(nil)).Elem()
}

func (o LookupUserPoolDomainResultOutput) ToLookupUserPoolDomainResultOutput() LookupUserPoolDomainResultOutput {
	return o
}

func (o LookupUserPoolDomainResultOutput) ToLookupUserPoolDomainResultOutputWithContext(ctx context.Context) LookupUserPoolDomainResultOutput {
	return o
}

func (o LookupUserPoolDomainResultOutput) CloudFrontDistribution() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserPoolDomainResult) *string { return v.CloudFrontDistribution }).(pulumi.StringPtrOutput)
}

func (o LookupUserPoolDomainResultOutput) CustomDomainConfig() UserPoolDomainCustomDomainConfigTypePtrOutput {
	return o.ApplyT(func(v LookupUserPoolDomainResult) *UserPoolDomainCustomDomainConfigType { return v.CustomDomainConfig }).(UserPoolDomainCustomDomainConfigTypePtrOutput)
}

func (o LookupUserPoolDomainResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserPoolDomainResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserPoolDomainResultOutput{})
}
