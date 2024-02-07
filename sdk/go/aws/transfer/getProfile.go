// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package transfer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Transfer::Profile
func LookupProfile(ctx *pulumi.Context, args *LookupProfileArgs, opts ...pulumi.InvokeOption) (*LookupProfileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProfileResult
	err := ctx.Invoke("aws-native:transfer:getProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProfileArgs struct {
	// A unique identifier for the profile
	ProfileId string `pulumi:"profileId"`
}

type LookupProfileResult struct {
	// Specifies the unique Amazon Resource Name (ARN) for the profile.
	Arn *string `pulumi:"arn"`
	// AS2 identifier agreed with a trading partner.
	As2Id *string `pulumi:"as2Id"`
	// List of the certificate IDs associated with this profile to be used for encryption and signing of AS2 messages.
	CertificateIds []string `pulumi:"certificateIds"`
	// A unique identifier for the profile
	ProfileId *string `pulumi:"profileId"`
	// An array of key-value pairs to apply to this resource.
	Tags []ProfileTag `pulumi:"tags"`
}

func LookupProfileOutput(ctx *pulumi.Context, args LookupProfileOutputArgs, opts ...pulumi.InvokeOption) LookupProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProfileResult, error) {
			args := v.(LookupProfileArgs)
			r, err := LookupProfile(ctx, &args, opts...)
			var s LookupProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProfileResultOutput)
}

type LookupProfileOutputArgs struct {
	// A unique identifier for the profile
	ProfileId pulumi.StringInput `pulumi:"profileId"`
}

func (LookupProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProfileArgs)(nil)).Elem()
}

type LookupProfileResultOutput struct{ *pulumi.OutputState }

func (LookupProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProfileResult)(nil)).Elem()
}

func (o LookupProfileResultOutput) ToLookupProfileResultOutput() LookupProfileResultOutput {
	return o
}

func (o LookupProfileResultOutput) ToLookupProfileResultOutputWithContext(ctx context.Context) LookupProfileResultOutput {
	return o
}

// Specifies the unique Amazon Resource Name (ARN) for the profile.
func (o LookupProfileResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// AS2 identifier agreed with a trading partner.
func (o LookupProfileResultOutput) As2Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.As2Id }).(pulumi.StringPtrOutput)
}

// List of the certificate IDs associated with this profile to be used for encryption and signing of AS2 messages.
func (o LookupProfileResultOutput) CertificateIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupProfileResult) []string { return v.CertificateIds }).(pulumi.StringArrayOutput)
}

// A unique identifier for the profile
func (o LookupProfileResultOutput) ProfileId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.ProfileId }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupProfileResultOutput) Tags() ProfileTagArrayOutput {
	return o.ApplyT(func(v LookupProfileResult) []ProfileTag { return v.Tags }).(ProfileTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProfileResultOutput{})
}
