// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IAM::ServiceLinkedRole
func LookupServiceLinkedRole(ctx *pulumi.Context, args *LookupServiceLinkedRoleArgs, opts ...pulumi.InvokeOption) (*LookupServiceLinkedRoleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupServiceLinkedRoleResult
	err := ctx.Invoke("aws-native:iam:getServiceLinkedRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceLinkedRoleArgs struct {
	// The name of the role.
	RoleName string `pulumi:"roleName"`
}

type LookupServiceLinkedRoleResult struct {
	// The description of the role.
	Description *string `pulumi:"description"`
	// The name of the role.
	RoleName *string `pulumi:"roleName"`
}

func LookupServiceLinkedRoleOutput(ctx *pulumi.Context, args LookupServiceLinkedRoleOutputArgs, opts ...pulumi.InvokeOption) LookupServiceLinkedRoleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceLinkedRoleResult, error) {
			args := v.(LookupServiceLinkedRoleArgs)
			r, err := LookupServiceLinkedRole(ctx, &args, opts...)
			var s LookupServiceLinkedRoleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceLinkedRoleResultOutput)
}

type LookupServiceLinkedRoleOutputArgs struct {
	// The name of the role.
	RoleName pulumi.StringInput `pulumi:"roleName"`
}

func (LookupServiceLinkedRoleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceLinkedRoleArgs)(nil)).Elem()
}

type LookupServiceLinkedRoleResultOutput struct{ *pulumi.OutputState }

func (LookupServiceLinkedRoleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceLinkedRoleResult)(nil)).Elem()
}

func (o LookupServiceLinkedRoleResultOutput) ToLookupServiceLinkedRoleResultOutput() LookupServiceLinkedRoleResultOutput {
	return o
}

func (o LookupServiceLinkedRoleResultOutput) ToLookupServiceLinkedRoleResultOutputWithContext(ctx context.Context) LookupServiceLinkedRoleResultOutput {
	return o
}

// The description of the role.
func (o LookupServiceLinkedRoleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceLinkedRoleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the role.
func (o LookupServiceLinkedRoleResultOutput) RoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceLinkedRoleResult) *string { return v.RoleName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceLinkedRoleResultOutput{})
}
