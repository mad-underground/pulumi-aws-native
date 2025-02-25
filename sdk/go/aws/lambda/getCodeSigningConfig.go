// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Lambda::CodeSigningConfig.
func LookupCodeSigningConfig(ctx *pulumi.Context, args *LookupCodeSigningConfigArgs, opts ...pulumi.InvokeOption) (*LookupCodeSigningConfigResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCodeSigningConfigResult
	err := ctx.Invoke("aws-native:lambda:getCodeSigningConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCodeSigningConfigArgs struct {
	// A unique Arn for CodeSigningConfig resource
	CodeSigningConfigArn string `pulumi:"codeSigningConfigArn"`
}

type LookupCodeSigningConfigResult struct {
	// When the CodeSigningConfig is later on attached to a function, the function code will be expected to be signed by profiles from this list
	AllowedPublishers *CodeSigningConfigAllowedPublishers `pulumi:"allowedPublishers"`
	// A unique Arn for CodeSigningConfig resource
	CodeSigningConfigArn *string `pulumi:"codeSigningConfigArn"`
	// A unique identifier for CodeSigningConfig resource
	CodeSigningConfigId *string `pulumi:"codeSigningConfigId"`
	// Policies to control how to act if a signature is invalid
	CodeSigningPolicies *CodeSigningConfigCodeSigningPolicies `pulumi:"codeSigningPolicies"`
	// A description of the CodeSigningConfig
	Description *string `pulumi:"description"`
}

func LookupCodeSigningConfigOutput(ctx *pulumi.Context, args LookupCodeSigningConfigOutputArgs, opts ...pulumi.InvokeOption) LookupCodeSigningConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCodeSigningConfigResult, error) {
			args := v.(LookupCodeSigningConfigArgs)
			r, err := LookupCodeSigningConfig(ctx, &args, opts...)
			var s LookupCodeSigningConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCodeSigningConfigResultOutput)
}

type LookupCodeSigningConfigOutputArgs struct {
	// A unique Arn for CodeSigningConfig resource
	CodeSigningConfigArn pulumi.StringInput `pulumi:"codeSigningConfigArn"`
}

func (LookupCodeSigningConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCodeSigningConfigArgs)(nil)).Elem()
}

type LookupCodeSigningConfigResultOutput struct{ *pulumi.OutputState }

func (LookupCodeSigningConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCodeSigningConfigResult)(nil)).Elem()
}

func (o LookupCodeSigningConfigResultOutput) ToLookupCodeSigningConfigResultOutput() LookupCodeSigningConfigResultOutput {
	return o
}

func (o LookupCodeSigningConfigResultOutput) ToLookupCodeSigningConfigResultOutputWithContext(ctx context.Context) LookupCodeSigningConfigResultOutput {
	return o
}

// When the CodeSigningConfig is later on attached to a function, the function code will be expected to be signed by profiles from this list
func (o LookupCodeSigningConfigResultOutput) AllowedPublishers() CodeSigningConfigAllowedPublishersPtrOutput {
	return o.ApplyT(func(v LookupCodeSigningConfigResult) *CodeSigningConfigAllowedPublishers { return v.AllowedPublishers }).(CodeSigningConfigAllowedPublishersPtrOutput)
}

// A unique Arn for CodeSigningConfig resource
func (o LookupCodeSigningConfigResultOutput) CodeSigningConfigArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCodeSigningConfigResult) *string { return v.CodeSigningConfigArn }).(pulumi.StringPtrOutput)
}

// A unique identifier for CodeSigningConfig resource
func (o LookupCodeSigningConfigResultOutput) CodeSigningConfigId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCodeSigningConfigResult) *string { return v.CodeSigningConfigId }).(pulumi.StringPtrOutput)
}

// Policies to control how to act if a signature is invalid
func (o LookupCodeSigningConfigResultOutput) CodeSigningPolicies() CodeSigningConfigCodeSigningPoliciesPtrOutput {
	return o.ApplyT(func(v LookupCodeSigningConfigResult) *CodeSigningConfigCodeSigningPolicies {
		return v.CodeSigningPolicies
	}).(CodeSigningConfigCodeSigningPoliciesPtrOutput)
}

// A description of the CodeSigningConfig
func (o LookupCodeSigningConfigResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCodeSigningConfigResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCodeSigningConfigResultOutput{})
}
