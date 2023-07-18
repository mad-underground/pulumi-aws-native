// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package batch

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type schema for AWS::Batch::SchedulingPolicy
type SchedulingPolicy struct {
	pulumi.CustomResourceState

	Arn             pulumi.StringOutput                      `pulumi:"arn"`
	FairsharePolicy SchedulingPolicyFairsharePolicyPtrOutput `pulumi:"fairsharePolicy"`
	// Name of Scheduling Policy.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// A key-value pair to associate with a resource.
	Tags pulumi.AnyOutput `pulumi:"tags"`
}

// NewSchedulingPolicy registers a new resource with the given unique name, arguments, and options.
func NewSchedulingPolicy(ctx *pulumi.Context,
	name string, args *SchedulingPolicyArgs, opts ...pulumi.ResourceOption) (*SchedulingPolicy, error) {
	if args == nil {
		args = &SchedulingPolicyArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SchedulingPolicy
	err := ctx.RegisterResource("aws-native:batch:SchedulingPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSchedulingPolicy gets an existing SchedulingPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSchedulingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SchedulingPolicyState, opts ...pulumi.ResourceOption) (*SchedulingPolicy, error) {
	var resource SchedulingPolicy
	err := ctx.ReadResource("aws-native:batch:SchedulingPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SchedulingPolicy resources.
type schedulingPolicyState struct {
}

type SchedulingPolicyState struct {
}

func (SchedulingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*schedulingPolicyState)(nil)).Elem()
}

type schedulingPolicyArgs struct {
	FairsharePolicy *SchedulingPolicyFairsharePolicy `pulumi:"fairsharePolicy"`
	// Name of Scheduling Policy.
	Name *string `pulumi:"name"`
	// A key-value pair to associate with a resource.
	Tags interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a SchedulingPolicy resource.
type SchedulingPolicyArgs struct {
	FairsharePolicy SchedulingPolicyFairsharePolicyPtrInput
	// Name of Scheduling Policy.
	Name pulumi.StringPtrInput
	// A key-value pair to associate with a resource.
	Tags pulumi.Input
}

func (SchedulingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*schedulingPolicyArgs)(nil)).Elem()
}

type SchedulingPolicyInput interface {
	pulumi.Input

	ToSchedulingPolicyOutput() SchedulingPolicyOutput
	ToSchedulingPolicyOutputWithContext(ctx context.Context) SchedulingPolicyOutput
}

func (*SchedulingPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**SchedulingPolicy)(nil)).Elem()
}

func (i *SchedulingPolicy) ToSchedulingPolicyOutput() SchedulingPolicyOutput {
	return i.ToSchedulingPolicyOutputWithContext(context.Background())
}

func (i *SchedulingPolicy) ToSchedulingPolicyOutputWithContext(ctx context.Context) SchedulingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchedulingPolicyOutput)
}

type SchedulingPolicyOutput struct{ *pulumi.OutputState }

func (SchedulingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SchedulingPolicy)(nil)).Elem()
}

func (o SchedulingPolicyOutput) ToSchedulingPolicyOutput() SchedulingPolicyOutput {
	return o
}

func (o SchedulingPolicyOutput) ToSchedulingPolicyOutputWithContext(ctx context.Context) SchedulingPolicyOutput {
	return o
}

func (o SchedulingPolicyOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *SchedulingPolicy) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o SchedulingPolicyOutput) FairsharePolicy() SchedulingPolicyFairsharePolicyPtrOutput {
	return o.ApplyT(func(v *SchedulingPolicy) SchedulingPolicyFairsharePolicyPtrOutput { return v.FairsharePolicy }).(SchedulingPolicyFairsharePolicyPtrOutput)
}

// Name of Scheduling Policy.
func (o SchedulingPolicyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SchedulingPolicy) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// A key-value pair to associate with a resource.
func (o SchedulingPolicyOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v *SchedulingPolicy) pulumi.AnyOutput { return v.Tags }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SchedulingPolicyInput)(nil)).Elem(), &SchedulingPolicy{})
	pulumi.RegisterOutputType(SchedulingPolicyOutput{})
}
