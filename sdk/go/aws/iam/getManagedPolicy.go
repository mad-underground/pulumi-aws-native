// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IAM::ManagedPolicy
func LookupManagedPolicy(ctx *pulumi.Context, args *LookupManagedPolicyArgs, opts ...pulumi.InvokeOption) (*LookupManagedPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupManagedPolicyResult
	err := ctx.Invoke("aws-native:iam:getManagedPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedPolicyArgs struct {
	// Amazon Resource Name (ARN) of the managed policy
	PolicyArn string `pulumi:"policyArn"`
}

type LookupManagedPolicyResult struct {
	// The number of entities (users, groups, and roles) that the policy is attached to.
	AttachmentCount *int `pulumi:"attachmentCount"`
	// The date and time, in ISO 8601 date-time format, when the policy was created.
	CreateDate *string `pulumi:"createDate"`
	// The identifier for the version of the policy that is set as the default version.
	DefaultVersionId *string `pulumi:"defaultVersionId"`
	// The name (friendly name, not ARN) of the group to attach the policy to.
	Groups []string `pulumi:"groups"`
	// Specifies whether the policy can be attached to an IAM user, group, or role.
	IsAttachable *bool `pulumi:"isAttachable"`
	// The number of entities (users and roles) for which the policy is used to set the permissions boundary.
	PermissionsBoundaryUsageCount *int `pulumi:"permissionsBoundaryUsageCount"`
	// Amazon Resource Name (ARN) of the managed policy
	PolicyArn *string `pulumi:"policyArn"`
	// The JSON policy document that you want to use as the content for the new policy.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::IAM::ManagedPolicy` for more information about the expected schema for this property.
	PolicyDocument interface{} `pulumi:"policyDocument"`
	// The stable and unique string identifying the policy.
	PolicyId *string `pulumi:"policyId"`
	// The name (friendly name, not ARN) of the role to attach the policy to.
	Roles []string `pulumi:"roles"`
	// The date and time, in ISO 8601 date-time format, when the policy was last updated.
	UpdateDate *string `pulumi:"updateDate"`
	// The name (friendly name, not ARN) of the IAM user to attach the policy to.
	Users []string `pulumi:"users"`
}

func LookupManagedPolicyOutput(ctx *pulumi.Context, args LookupManagedPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupManagedPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedPolicyResult, error) {
			args := v.(LookupManagedPolicyArgs)
			r, err := LookupManagedPolicy(ctx, &args, opts...)
			var s LookupManagedPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedPolicyResultOutput)
}

type LookupManagedPolicyOutputArgs struct {
	// Amazon Resource Name (ARN) of the managed policy
	PolicyArn pulumi.StringInput `pulumi:"policyArn"`
}

func (LookupManagedPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedPolicyArgs)(nil)).Elem()
}

type LookupManagedPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupManagedPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedPolicyResult)(nil)).Elem()
}

func (o LookupManagedPolicyResultOutput) ToLookupManagedPolicyResultOutput() LookupManagedPolicyResultOutput {
	return o
}

func (o LookupManagedPolicyResultOutput) ToLookupManagedPolicyResultOutputWithContext(ctx context.Context) LookupManagedPolicyResultOutput {
	return o
}

// The number of entities (users, groups, and roles) that the policy is attached to.
func (o LookupManagedPolicyResultOutput) AttachmentCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupManagedPolicyResult) *int { return v.AttachmentCount }).(pulumi.IntPtrOutput)
}

// The date and time, in ISO 8601 date-time format, when the policy was created.
func (o LookupManagedPolicyResultOutput) CreateDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedPolicyResult) *string { return v.CreateDate }).(pulumi.StringPtrOutput)
}

// The identifier for the version of the policy that is set as the default version.
func (o LookupManagedPolicyResultOutput) DefaultVersionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedPolicyResult) *string { return v.DefaultVersionId }).(pulumi.StringPtrOutput)
}

// The name (friendly name, not ARN) of the group to attach the policy to.
func (o LookupManagedPolicyResultOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupManagedPolicyResult) []string { return v.Groups }).(pulumi.StringArrayOutput)
}

// Specifies whether the policy can be attached to an IAM user, group, or role.
func (o LookupManagedPolicyResultOutput) IsAttachable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupManagedPolicyResult) *bool { return v.IsAttachable }).(pulumi.BoolPtrOutput)
}

// The number of entities (users and roles) for which the policy is used to set the permissions boundary.
func (o LookupManagedPolicyResultOutput) PermissionsBoundaryUsageCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupManagedPolicyResult) *int { return v.PermissionsBoundaryUsageCount }).(pulumi.IntPtrOutput)
}

// Amazon Resource Name (ARN) of the managed policy
func (o LookupManagedPolicyResultOutput) PolicyArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedPolicyResult) *string { return v.PolicyArn }).(pulumi.StringPtrOutput)
}

// The JSON policy document that you want to use as the content for the new policy.
//
// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::IAM::ManagedPolicy` for more information about the expected schema for this property.
func (o LookupManagedPolicyResultOutput) PolicyDocument() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupManagedPolicyResult) interface{} { return v.PolicyDocument }).(pulumi.AnyOutput)
}

// The stable and unique string identifying the policy.
func (o LookupManagedPolicyResultOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedPolicyResult) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

// The name (friendly name, not ARN) of the role to attach the policy to.
func (o LookupManagedPolicyResultOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupManagedPolicyResult) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

// The date and time, in ISO 8601 date-time format, when the policy was last updated.
func (o LookupManagedPolicyResultOutput) UpdateDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedPolicyResult) *string { return v.UpdateDate }).(pulumi.StringPtrOutput)
}

// The name (friendly name, not ARN) of the IAM user to attach the policy to.
func (o LookupManagedPolicyResultOutput) Users() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupManagedPolicyResult) []string { return v.Users }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedPolicyResultOutput{})
}
