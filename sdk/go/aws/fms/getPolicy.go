// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an AWS Firewall Manager policy.
func LookupPolicy(ctx *pulumi.Context, args *LookupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPolicyResult
	err := ctx.Invoke("aws-native:fms:getPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicyArgs struct {
	Id string `pulumi:"id"`
}

type LookupPolicyResult struct {
	Arn                       *string                          `pulumi:"arn"`
	ExcludeMap                *PolicyIeMap                     `pulumi:"excludeMap"`
	ExcludeResourceTags       *bool                            `pulumi:"excludeResourceTags"`
	Id                        *string                          `pulumi:"id"`
	IncludeMap                *PolicyIeMap                     `pulumi:"includeMap"`
	PolicyDescription         *string                          `pulumi:"policyDescription"`
	PolicyName                *string                          `pulumi:"policyName"`
	RemediationEnabled        *bool                            `pulumi:"remediationEnabled"`
	ResourceSetIds            []string                         `pulumi:"resourceSetIds"`
	ResourceTags              []PolicyResourceTag              `pulumi:"resourceTags"`
	ResourceType              *string                          `pulumi:"resourceType"`
	ResourceTypeList          []string                         `pulumi:"resourceTypeList"`
	ResourcesCleanUp          *bool                            `pulumi:"resourcesCleanUp"`
	SecurityServicePolicyData *PolicySecurityServicePolicyData `pulumi:"securityServicePolicyData"`
	Tags                      []PolicyTag                      `pulumi:"tags"`
}

func LookupPolicyOutput(ctx *pulumi.Context, args LookupPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicyResult, error) {
			args := v.(LookupPolicyArgs)
			r, err := LookupPolicy(ctx, &args, opts...)
			var s LookupPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicyResultOutput)
}

type LookupPolicyOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyArgs)(nil)).Elem()
}

type LookupPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyResult)(nil)).Elem()
}

func (o LookupPolicyResultOutput) ToLookupPolicyResultOutput() LookupPolicyResultOutput {
	return o
}

func (o LookupPolicyResultOutput) ToLookupPolicyResultOutputWithContext(ctx context.Context) LookupPolicyResultOutput {
	return o
}

func (o LookupPolicyResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyResultOutput) ExcludeMap() PolicyIeMapPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *PolicyIeMap { return v.ExcludeMap }).(PolicyIeMapPtrOutput)
}

func (o LookupPolicyResultOutput) ExcludeResourceTags() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *bool { return v.ExcludeResourceTags }).(pulumi.BoolPtrOutput)
}

func (o LookupPolicyResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyResultOutput) IncludeMap() PolicyIeMapPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *PolicyIeMap { return v.IncludeMap }).(PolicyIeMapPtrOutput)
}

func (o LookupPolicyResultOutput) PolicyDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *string { return v.PolicyDescription }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyResultOutput) PolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *string { return v.PolicyName }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyResultOutput) RemediationEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *bool { return v.RemediationEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupPolicyResultOutput) ResourceSetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPolicyResult) []string { return v.ResourceSetIds }).(pulumi.StringArrayOutput)
}

func (o LookupPolicyResultOutput) ResourceTags() PolicyResourceTagArrayOutput {
	return o.ApplyT(func(v LookupPolicyResult) []PolicyResourceTag { return v.ResourceTags }).(PolicyResourceTagArrayOutput)
}

func (o LookupPolicyResultOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyResultOutput) ResourceTypeList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPolicyResult) []string { return v.ResourceTypeList }).(pulumi.StringArrayOutput)
}

func (o LookupPolicyResultOutput) ResourcesCleanUp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *bool { return v.ResourcesCleanUp }).(pulumi.BoolPtrOutput)
}

func (o LookupPolicyResultOutput) SecurityServicePolicyData() PolicySecurityServicePolicyDataPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *PolicySecurityServicePolicyData { return v.SecurityServicePolicyData }).(PolicySecurityServicePolicyDataPtrOutput)
}

func (o LookupPolicyResultOutput) Tags() PolicyTagArrayOutput {
	return o.ApplyT(func(v LookupPolicyResult) []PolicyTag { return v.Tags }).(PolicyTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyResultOutput{})
}
