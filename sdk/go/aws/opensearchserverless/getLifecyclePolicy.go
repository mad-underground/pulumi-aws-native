// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package opensearchserverless

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Amazon OpenSearchServerless lifecycle policy resource
func LookupLifecyclePolicy(ctx *pulumi.Context, args *LookupLifecyclePolicyArgs, opts ...pulumi.InvokeOption) (*LookupLifecyclePolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLifecyclePolicyResult
	err := ctx.Invoke("aws-native:opensearchserverless:getLifecyclePolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLifecyclePolicyArgs struct {
	// The name of the policy
	Name string              `pulumi:"name"`
	Type LifecyclePolicyType `pulumi:"type"`
}

type LookupLifecyclePolicyResult struct {
	// The description of the policy
	Description *string `pulumi:"description"`
	// The JSON policy document that is the content for the policy
	Policy *string `pulumi:"policy"`
}

func LookupLifecyclePolicyOutput(ctx *pulumi.Context, args LookupLifecyclePolicyOutputArgs, opts ...pulumi.InvokeOption) LookupLifecyclePolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLifecyclePolicyResult, error) {
			args := v.(LookupLifecyclePolicyArgs)
			r, err := LookupLifecyclePolicy(ctx, &args, opts...)
			var s LookupLifecyclePolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLifecyclePolicyResultOutput)
}

type LookupLifecyclePolicyOutputArgs struct {
	// The name of the policy
	Name pulumi.StringInput       `pulumi:"name"`
	Type LifecyclePolicyTypeInput `pulumi:"type"`
}

func (LookupLifecyclePolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLifecyclePolicyArgs)(nil)).Elem()
}

type LookupLifecyclePolicyResultOutput struct{ *pulumi.OutputState }

func (LookupLifecyclePolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLifecyclePolicyResult)(nil)).Elem()
}

func (o LookupLifecyclePolicyResultOutput) ToLookupLifecyclePolicyResultOutput() LookupLifecyclePolicyResultOutput {
	return o
}

func (o LookupLifecyclePolicyResultOutput) ToLookupLifecyclePolicyResultOutputWithContext(ctx context.Context) LookupLifecyclePolicyResultOutput {
	return o
}

// The description of the policy
func (o LookupLifecyclePolicyResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLifecyclePolicyResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The JSON policy document that is the content for the policy
func (o LookupLifecyclePolicyResultOutput) Policy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLifecyclePolicyResult) *string { return v.Policy }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLifecyclePolicyResultOutput{})
}
