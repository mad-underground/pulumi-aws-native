// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::EC2::VerifiedAccessGroup resource creates an AWS EC2 Verified Access Group.
func LookupVerifiedAccessGroup(ctx *pulumi.Context, args *LookupVerifiedAccessGroupArgs, opts ...pulumi.InvokeOption) (*LookupVerifiedAccessGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVerifiedAccessGroupResult
	err := ctx.Invoke("aws-native:ec2:getVerifiedAccessGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVerifiedAccessGroupArgs struct {
	// The ID of the AWS Verified Access group.
	VerifiedAccessGroupId string `pulumi:"verifiedAccessGroupId"`
}

type LookupVerifiedAccessGroupResult struct {
	// Time this Verified Access Group was created.
	CreationTime *string `pulumi:"creationTime"`
	// A description for the AWS Verified Access group.
	Description *string `pulumi:"description"`
	// Time this Verified Access Group was last updated.
	LastUpdatedTime *string `pulumi:"lastUpdatedTime"`
	// The AWS account number that owns the group.
	Owner *string `pulumi:"owner"`
	// The AWS Verified Access policy document.
	PolicyDocument *string `pulumi:"policyDocument"`
	// The status of the Verified Access policy.
	PolicyEnabled *bool `pulumi:"policyEnabled"`
	// An array of key-value pairs to apply to this resource.
	Tags []VerifiedAccessGroupTag `pulumi:"tags"`
	// The ARN of the Verified Access group.
	VerifiedAccessGroupArn *string `pulumi:"verifiedAccessGroupArn"`
	// The ID of the AWS Verified Access group.
	VerifiedAccessGroupId *string `pulumi:"verifiedAccessGroupId"`
	// The ID of the AWS Verified Access instance.
	VerifiedAccessInstanceId *string `pulumi:"verifiedAccessInstanceId"`
}

func LookupVerifiedAccessGroupOutput(ctx *pulumi.Context, args LookupVerifiedAccessGroupOutputArgs, opts ...pulumi.InvokeOption) LookupVerifiedAccessGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVerifiedAccessGroupResult, error) {
			args := v.(LookupVerifiedAccessGroupArgs)
			r, err := LookupVerifiedAccessGroup(ctx, &args, opts...)
			var s LookupVerifiedAccessGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVerifiedAccessGroupResultOutput)
}

type LookupVerifiedAccessGroupOutputArgs struct {
	// The ID of the AWS Verified Access group.
	VerifiedAccessGroupId pulumi.StringInput `pulumi:"verifiedAccessGroupId"`
}

func (LookupVerifiedAccessGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVerifiedAccessGroupArgs)(nil)).Elem()
}

type LookupVerifiedAccessGroupResultOutput struct{ *pulumi.OutputState }

func (LookupVerifiedAccessGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVerifiedAccessGroupResult)(nil)).Elem()
}

func (o LookupVerifiedAccessGroupResultOutput) ToLookupVerifiedAccessGroupResultOutput() LookupVerifiedAccessGroupResultOutput {
	return o
}

func (o LookupVerifiedAccessGroupResultOutput) ToLookupVerifiedAccessGroupResultOutputWithContext(ctx context.Context) LookupVerifiedAccessGroupResultOutput {
	return o
}

// Time this Verified Access Group was created.
func (o LookupVerifiedAccessGroupResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVerifiedAccessGroupResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

// A description for the AWS Verified Access group.
func (o LookupVerifiedAccessGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVerifiedAccessGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Time this Verified Access Group was last updated.
func (o LookupVerifiedAccessGroupResultOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVerifiedAccessGroupResult) *string { return v.LastUpdatedTime }).(pulumi.StringPtrOutput)
}

// The AWS account number that owns the group.
func (o LookupVerifiedAccessGroupResultOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVerifiedAccessGroupResult) *string { return v.Owner }).(pulumi.StringPtrOutput)
}

// The AWS Verified Access policy document.
func (o LookupVerifiedAccessGroupResultOutput) PolicyDocument() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVerifiedAccessGroupResult) *string { return v.PolicyDocument }).(pulumi.StringPtrOutput)
}

// The status of the Verified Access policy.
func (o LookupVerifiedAccessGroupResultOutput) PolicyEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVerifiedAccessGroupResult) *bool { return v.PolicyEnabled }).(pulumi.BoolPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupVerifiedAccessGroupResultOutput) Tags() VerifiedAccessGroupTagArrayOutput {
	return o.ApplyT(func(v LookupVerifiedAccessGroupResult) []VerifiedAccessGroupTag { return v.Tags }).(VerifiedAccessGroupTagArrayOutput)
}

// The ARN of the Verified Access group.
func (o LookupVerifiedAccessGroupResultOutput) VerifiedAccessGroupArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVerifiedAccessGroupResult) *string { return v.VerifiedAccessGroupArn }).(pulumi.StringPtrOutput)
}

// The ID of the AWS Verified Access group.
func (o LookupVerifiedAccessGroupResultOutput) VerifiedAccessGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVerifiedAccessGroupResult) *string { return v.VerifiedAccessGroupId }).(pulumi.StringPtrOutput)
}

// The ID of the AWS Verified Access instance.
func (o LookupVerifiedAccessGroupResultOutput) VerifiedAccessInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVerifiedAccessGroupResult) *string { return v.VerifiedAccessInstanceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVerifiedAccessGroupResultOutput{})
}
