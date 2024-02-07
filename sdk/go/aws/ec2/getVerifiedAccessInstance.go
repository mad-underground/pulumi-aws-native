// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::EC2::VerifiedAccessInstance resource creates an AWS EC2 Verified Access Instance.
func LookupVerifiedAccessInstance(ctx *pulumi.Context, args *LookupVerifiedAccessInstanceArgs, opts ...pulumi.InvokeOption) (*LookupVerifiedAccessInstanceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVerifiedAccessInstanceResult
	err := ctx.Invoke("aws-native:ec2:getVerifiedAccessInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVerifiedAccessInstanceArgs struct {
	// The ID of the AWS Verified Access instance.
	VerifiedAccessInstanceId string `pulumi:"verifiedAccessInstanceId"`
}

type LookupVerifiedAccessInstanceResult struct {
	// Time this Verified Access Instance was created.
	CreationTime *string `pulumi:"creationTime"`
	// A description for the AWS Verified Access instance.
	Description *string `pulumi:"description"`
	// Indicates whether FIPS is enabled
	FipsEnabled *bool `pulumi:"fipsEnabled"`
	// Time this Verified Access Instance was last updated.
	LastUpdatedTime *string `pulumi:"lastUpdatedTime"`
	// The configuration options for AWS Verified Access instances.
	LoggingConfigurations *VerifiedAccessInstanceVerifiedAccessLogs `pulumi:"loggingConfigurations"`
	// An array of key-value pairs to apply to this resource.
	Tags []VerifiedAccessInstanceTag `pulumi:"tags"`
	// The ID of the AWS Verified Access instance.
	VerifiedAccessInstanceId *string `pulumi:"verifiedAccessInstanceId"`
	// The IDs of the AWS Verified Access trust providers.
	VerifiedAccessTrustProviderIds []string `pulumi:"verifiedAccessTrustProviderIds"`
	// AWS Verified Access trust providers.
	VerifiedAccessTrustProviders []VerifiedAccessInstanceVerifiedAccessTrustProvider `pulumi:"verifiedAccessTrustProviders"`
}

func LookupVerifiedAccessInstanceOutput(ctx *pulumi.Context, args LookupVerifiedAccessInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupVerifiedAccessInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVerifiedAccessInstanceResult, error) {
			args := v.(LookupVerifiedAccessInstanceArgs)
			r, err := LookupVerifiedAccessInstance(ctx, &args, opts...)
			var s LookupVerifiedAccessInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVerifiedAccessInstanceResultOutput)
}

type LookupVerifiedAccessInstanceOutputArgs struct {
	// The ID of the AWS Verified Access instance.
	VerifiedAccessInstanceId pulumi.StringInput `pulumi:"verifiedAccessInstanceId"`
}

func (LookupVerifiedAccessInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVerifiedAccessInstanceArgs)(nil)).Elem()
}

type LookupVerifiedAccessInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupVerifiedAccessInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVerifiedAccessInstanceResult)(nil)).Elem()
}

func (o LookupVerifiedAccessInstanceResultOutput) ToLookupVerifiedAccessInstanceResultOutput() LookupVerifiedAccessInstanceResultOutput {
	return o
}

func (o LookupVerifiedAccessInstanceResultOutput) ToLookupVerifiedAccessInstanceResultOutputWithContext(ctx context.Context) LookupVerifiedAccessInstanceResultOutput {
	return o
}

// Time this Verified Access Instance was created.
func (o LookupVerifiedAccessInstanceResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVerifiedAccessInstanceResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

// A description for the AWS Verified Access instance.
func (o LookupVerifiedAccessInstanceResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVerifiedAccessInstanceResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Indicates whether FIPS is enabled
func (o LookupVerifiedAccessInstanceResultOutput) FipsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVerifiedAccessInstanceResult) *bool { return v.FipsEnabled }).(pulumi.BoolPtrOutput)
}

// Time this Verified Access Instance was last updated.
func (o LookupVerifiedAccessInstanceResultOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVerifiedAccessInstanceResult) *string { return v.LastUpdatedTime }).(pulumi.StringPtrOutput)
}

// The configuration options for AWS Verified Access instances.
func (o LookupVerifiedAccessInstanceResultOutput) LoggingConfigurations() VerifiedAccessInstanceVerifiedAccessLogsPtrOutput {
	return o.ApplyT(func(v LookupVerifiedAccessInstanceResult) *VerifiedAccessInstanceVerifiedAccessLogs {
		return v.LoggingConfigurations
	}).(VerifiedAccessInstanceVerifiedAccessLogsPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupVerifiedAccessInstanceResultOutput) Tags() VerifiedAccessInstanceTagArrayOutput {
	return o.ApplyT(func(v LookupVerifiedAccessInstanceResult) []VerifiedAccessInstanceTag { return v.Tags }).(VerifiedAccessInstanceTagArrayOutput)
}

// The ID of the AWS Verified Access instance.
func (o LookupVerifiedAccessInstanceResultOutput) VerifiedAccessInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVerifiedAccessInstanceResult) *string { return v.VerifiedAccessInstanceId }).(pulumi.StringPtrOutput)
}

// The IDs of the AWS Verified Access trust providers.
func (o LookupVerifiedAccessInstanceResultOutput) VerifiedAccessTrustProviderIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVerifiedAccessInstanceResult) []string { return v.VerifiedAccessTrustProviderIds }).(pulumi.StringArrayOutput)
}

// AWS Verified Access trust providers.
func (o LookupVerifiedAccessInstanceResultOutput) VerifiedAccessTrustProviders() VerifiedAccessInstanceVerifiedAccessTrustProviderArrayOutput {
	return o.ApplyT(func(v LookupVerifiedAccessInstanceResult) []VerifiedAccessInstanceVerifiedAccessTrustProvider {
		return v.VerifiedAccessTrustProviders
	}).(VerifiedAccessInstanceVerifiedAccessTrustProviderArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVerifiedAccessInstanceResultOutput{})
}
