// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::Logs::AccountPolicy resource specifies a CloudWatch Logs AccountPolicy.
func LookupAccountPolicy(ctx *pulumi.Context, args *LookupAccountPolicyArgs, opts ...pulumi.InvokeOption) (*LookupAccountPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAccountPolicyResult
	err := ctx.Invoke("aws-native:logs:getAccountPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccountPolicyArgs struct {
	// User account id
	AccountId string `pulumi:"accountId"`
	// The name of the account policy
	PolicyName string `pulumi:"policyName"`
	// Type of the policy.
	PolicyType AccountPolicyPolicyType `pulumi:"policyType"`
}

type LookupAccountPolicyResult struct {
	// User account id
	AccountId *string `pulumi:"accountId"`
	// The body of the policy document you want to use for this topic.
	//
	// You can only add one policy per PolicyType.
	//
	// The policy must be in JSON string format.
	//
	// Length Constraints: Maximum length of 30720
	PolicyDocument *string `pulumi:"policyDocument"`
	// Scope for policy application
	Scope *AccountPolicyScope `pulumi:"scope"`
	// Log group  selection criteria to apply policy only to a subset of log groups. SelectionCriteria string can be up to 25KB and cloudwatchlogs determines the length of selectionCriteria by using its UTF-8 bytes
	SelectionCriteria *string `pulumi:"selectionCriteria"`
}

func LookupAccountPolicyOutput(ctx *pulumi.Context, args LookupAccountPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupAccountPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccountPolicyResult, error) {
			args := v.(LookupAccountPolicyArgs)
			r, err := LookupAccountPolicy(ctx, &args, opts...)
			var s LookupAccountPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccountPolicyResultOutput)
}

type LookupAccountPolicyOutputArgs struct {
	// User account id
	AccountId pulumi.StringInput `pulumi:"accountId"`
	// The name of the account policy
	PolicyName pulumi.StringInput `pulumi:"policyName"`
	// Type of the policy.
	PolicyType AccountPolicyPolicyTypeInput `pulumi:"policyType"`
}

func (LookupAccountPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountPolicyArgs)(nil)).Elem()
}

type LookupAccountPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupAccountPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountPolicyResult)(nil)).Elem()
}

func (o LookupAccountPolicyResultOutput) ToLookupAccountPolicyResultOutput() LookupAccountPolicyResultOutput {
	return o
}

func (o LookupAccountPolicyResultOutput) ToLookupAccountPolicyResultOutputWithContext(ctx context.Context) LookupAccountPolicyResultOutput {
	return o
}

// User account id
func (o LookupAccountPolicyResultOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountPolicyResult) *string { return v.AccountId }).(pulumi.StringPtrOutput)
}

// The body of the policy document you want to use for this topic.
//
// You can only add one policy per PolicyType.
//
// The policy must be in JSON string format.
//
// Length Constraints: Maximum length of 30720
func (o LookupAccountPolicyResultOutput) PolicyDocument() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountPolicyResult) *string { return v.PolicyDocument }).(pulumi.StringPtrOutput)
}

// Scope for policy application
func (o LookupAccountPolicyResultOutput) Scope() AccountPolicyScopePtrOutput {
	return o.ApplyT(func(v LookupAccountPolicyResult) *AccountPolicyScope { return v.Scope }).(AccountPolicyScopePtrOutput)
}

// Log group  selection criteria to apply policy only to a subset of log groups. SelectionCriteria string can be up to 25KB and cloudwatchlogs determines the length of selectionCriteria by using its UTF-8 bytes
func (o LookupAccountPolicyResultOutput) SelectionCriteria() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountPolicyResult) *string { return v.SelectionCriteria }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccountPolicyResultOutput{})
}
