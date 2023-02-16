// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sso

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for SSO PermissionSet
func LookupPermissionSet(ctx *pulumi.Context, args *LookupPermissionSetArgs, opts ...pulumi.InvokeOption) (*LookupPermissionSetResult, error) {
	var rv LookupPermissionSetResult
	err := ctx.Invoke("aws-native:sso:getPermissionSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPermissionSetArgs struct {
	// The sso instance arn that the permission set is owned.
	InstanceArn string `pulumi:"instanceArn"`
	// The permission set that the policy will be attached to
	PermissionSetArn string `pulumi:"permissionSetArn"`
}

type LookupPermissionSetResult struct {
	CustomerManagedPolicyReferences []PermissionSetCustomerManagedPolicyReference `pulumi:"customerManagedPolicyReferences"`
	// The permission set description.
	Description *string `pulumi:"description"`
	// The inline policy to put in permission set.
	InlinePolicy    interface{} `pulumi:"inlinePolicy"`
	ManagedPolicies []string    `pulumi:"managedPolicies"`
	// The permission set that the policy will be attached to
	PermissionSetArn    *string                           `pulumi:"permissionSetArn"`
	PermissionsBoundary *PermissionSetPermissionsBoundary `pulumi:"permissionsBoundary"`
	// The relay state URL that redirect links to any service in the AWS Management Console.
	RelayStateType *string `pulumi:"relayStateType"`
	// The length of time that a user can be signed in to an AWS account.
	SessionDuration *string            `pulumi:"sessionDuration"`
	Tags            []PermissionSetTag `pulumi:"tags"`
}

func LookupPermissionSetOutput(ctx *pulumi.Context, args LookupPermissionSetOutputArgs, opts ...pulumi.InvokeOption) LookupPermissionSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPermissionSetResult, error) {
			args := v.(LookupPermissionSetArgs)
			r, err := LookupPermissionSet(ctx, &args, opts...)
			var s LookupPermissionSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPermissionSetResultOutput)
}

type LookupPermissionSetOutputArgs struct {
	// The sso instance arn that the permission set is owned.
	InstanceArn pulumi.StringInput `pulumi:"instanceArn"`
	// The permission set that the policy will be attached to
	PermissionSetArn pulumi.StringInput `pulumi:"permissionSetArn"`
}

func (LookupPermissionSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPermissionSetArgs)(nil)).Elem()
}

type LookupPermissionSetResultOutput struct{ *pulumi.OutputState }

func (LookupPermissionSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPermissionSetResult)(nil)).Elem()
}

func (o LookupPermissionSetResultOutput) ToLookupPermissionSetResultOutput() LookupPermissionSetResultOutput {
	return o
}

func (o LookupPermissionSetResultOutput) ToLookupPermissionSetResultOutputWithContext(ctx context.Context) LookupPermissionSetResultOutput {
	return o
}

func (o LookupPermissionSetResultOutput) CustomerManagedPolicyReferences() PermissionSetCustomerManagedPolicyReferenceArrayOutput {
	return o.ApplyT(func(v LookupPermissionSetResult) []PermissionSetCustomerManagedPolicyReference {
		return v.CustomerManagedPolicyReferences
	}).(PermissionSetCustomerManagedPolicyReferenceArrayOutput)
}

// The permission set description.
func (o LookupPermissionSetResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPermissionSetResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The inline policy to put in permission set.
func (o LookupPermissionSetResultOutput) InlinePolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPermissionSetResult) interface{} { return v.InlinePolicy }).(pulumi.AnyOutput)
}

func (o LookupPermissionSetResultOutput) ManagedPolicies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPermissionSetResult) []string { return v.ManagedPolicies }).(pulumi.StringArrayOutput)
}

// The permission set that the policy will be attached to
func (o LookupPermissionSetResultOutput) PermissionSetArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPermissionSetResult) *string { return v.PermissionSetArn }).(pulumi.StringPtrOutput)
}

func (o LookupPermissionSetResultOutput) PermissionsBoundary() PermissionSetPermissionsBoundaryPtrOutput {
	return o.ApplyT(func(v LookupPermissionSetResult) *PermissionSetPermissionsBoundary { return v.PermissionsBoundary }).(PermissionSetPermissionsBoundaryPtrOutput)
}

// The relay state URL that redirect links to any service in the AWS Management Console.
func (o LookupPermissionSetResultOutput) RelayStateType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPermissionSetResult) *string { return v.RelayStateType }).(pulumi.StringPtrOutput)
}

// The length of time that a user can be signed in to an AWS account.
func (o LookupPermissionSetResultOutput) SessionDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPermissionSetResult) *string { return v.SessionDuration }).(pulumi.StringPtrOutput)
}

func (o LookupPermissionSetResultOutput) Tags() PermissionSetTagArrayOutput {
	return o.ApplyT(func(v LookupPermissionSetResult) []PermissionSetTag { return v.Tags }).(PermissionSetTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPermissionSetResultOutput{})
}
