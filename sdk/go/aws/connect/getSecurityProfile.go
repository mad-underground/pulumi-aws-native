// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package connect

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::Connect::SecurityProfile
func LookupSecurityProfile(ctx *pulumi.Context, args *LookupSecurityProfileArgs, opts ...pulumi.InvokeOption) (*LookupSecurityProfileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSecurityProfileResult
	err := ctx.Invoke("aws-native:connect:getSecurityProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecurityProfileArgs struct {
	// The Amazon Resource Name (ARN) for the security profile.
	SecurityProfileArn string `pulumi:"securityProfileArn"`
}

type LookupSecurityProfileResult struct {
	// The list of tags that a security profile uses to restrict access to resources in Amazon Connect.
	AllowedAccessControlTags []SecurityProfileTag `pulumi:"allowedAccessControlTags"`
	// The description of the security profile.
	Description *string `pulumi:"description"`
	// Permissions assigned to the security profile.
	Permissions []string `pulumi:"permissions"`
	// The Amazon Resource Name (ARN) for the security profile.
	SecurityProfileArn *string `pulumi:"securityProfileArn"`
	// The list of resources that a security profile applies tag restrictions to in Amazon Connect.
	TagRestrictedResources []string `pulumi:"tagRestrictedResources"`
	// The tags used to organize, track, or control access for this resource.
	Tags []SecurityProfileTag `pulumi:"tags"`
}

func LookupSecurityProfileOutput(ctx *pulumi.Context, args LookupSecurityProfileOutputArgs, opts ...pulumi.InvokeOption) LookupSecurityProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecurityProfileResult, error) {
			args := v.(LookupSecurityProfileArgs)
			r, err := LookupSecurityProfile(ctx, &args, opts...)
			var s LookupSecurityProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecurityProfileResultOutput)
}

type LookupSecurityProfileOutputArgs struct {
	// The Amazon Resource Name (ARN) for the security profile.
	SecurityProfileArn pulumi.StringInput `pulumi:"securityProfileArn"`
}

func (LookupSecurityProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityProfileArgs)(nil)).Elem()
}

type LookupSecurityProfileResultOutput struct{ *pulumi.OutputState }

func (LookupSecurityProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityProfileResult)(nil)).Elem()
}

func (o LookupSecurityProfileResultOutput) ToLookupSecurityProfileResultOutput() LookupSecurityProfileResultOutput {
	return o
}

func (o LookupSecurityProfileResultOutput) ToLookupSecurityProfileResultOutputWithContext(ctx context.Context) LookupSecurityProfileResultOutput {
	return o
}

func (o LookupSecurityProfileResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSecurityProfileResult] {
	return pulumix.Output[LookupSecurityProfileResult]{
		OutputState: o.OutputState,
	}
}

// The list of tags that a security profile uses to restrict access to resources in Amazon Connect.
func (o LookupSecurityProfileResultOutput) AllowedAccessControlTags() SecurityProfileTagArrayOutput {
	return o.ApplyT(func(v LookupSecurityProfileResult) []SecurityProfileTag { return v.AllowedAccessControlTags }).(SecurityProfileTagArrayOutput)
}

// The description of the security profile.
func (o LookupSecurityProfileResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityProfileResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Permissions assigned to the security profile.
func (o LookupSecurityProfileResultOutput) Permissions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSecurityProfileResult) []string { return v.Permissions }).(pulumi.StringArrayOutput)
}

// The Amazon Resource Name (ARN) for the security profile.
func (o LookupSecurityProfileResultOutput) SecurityProfileArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityProfileResult) *string { return v.SecurityProfileArn }).(pulumi.StringPtrOutput)
}

// The list of resources that a security profile applies tag restrictions to in Amazon Connect.
func (o LookupSecurityProfileResultOutput) TagRestrictedResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSecurityProfileResult) []string { return v.TagRestrictedResources }).(pulumi.StringArrayOutput)
}

// The tags used to organize, track, or control access for this resource.
func (o LookupSecurityProfileResultOutput) Tags() SecurityProfileTagArrayOutput {
	return o.ApplyT(func(v LookupSecurityProfileResult) []SecurityProfileTag { return v.Tags }).(SecurityProfileTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecurityProfileResultOutput{})
}