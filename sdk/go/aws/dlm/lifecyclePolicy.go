// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dlm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::DLM::LifecyclePolicy
//
// Deprecated: LifecyclePolicy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type LifecyclePolicy struct {
	pulumi.CustomResourceState

	Arn                    pulumi.StringOutput                            `pulumi:"arn"`
	AwsId                  pulumi.StringOutput                            `pulumi:"awsId"`
	CopyTags               pulumi.BoolPtrOutput                           `pulumi:"copyTags"`
	CreateInterval         pulumi.IntPtrOutput                            `pulumi:"createInterval"`
	CrossRegionCopyTargets LifecyclePolicyCrossRegionCopyTargetsPtrOutput `pulumi:"crossRegionCopyTargets"`
	DefaultPolicy          pulumi.StringPtrOutput                         `pulumi:"defaultPolicy"`
	Description            pulumi.StringPtrOutput                         `pulumi:"description"`
	Exclusions             LifecyclePolicyExclusionsPtrOutput             `pulumi:"exclusions"`
	ExecutionRoleArn       pulumi.StringPtrOutput                         `pulumi:"executionRoleArn"`
	ExtendDeletion         pulumi.BoolPtrOutput                           `pulumi:"extendDeletion"`
	PolicyDetails          LifecyclePolicyPolicyDetailsPtrOutput          `pulumi:"policyDetails"`
	RetainInterval         pulumi.IntPtrOutput                            `pulumi:"retainInterval"`
	State                  pulumi.StringPtrOutput                         `pulumi:"state"`
	Tags                   aws.TagArrayOutput                             `pulumi:"tags"`
}

// NewLifecyclePolicy registers a new resource with the given unique name, arguments, and options.
func NewLifecyclePolicy(ctx *pulumi.Context,
	name string, args *LifecyclePolicyArgs, opts ...pulumi.ResourceOption) (*LifecyclePolicy, error) {
	if args == nil {
		args = &LifecyclePolicyArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LifecyclePolicy
	err := ctx.RegisterResource("aws-native:dlm:LifecyclePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLifecyclePolicy gets an existing LifecyclePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLifecyclePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LifecyclePolicyState, opts ...pulumi.ResourceOption) (*LifecyclePolicy, error) {
	var resource LifecyclePolicy
	err := ctx.ReadResource("aws-native:dlm:LifecyclePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LifecyclePolicy resources.
type lifecyclePolicyState struct {
}

type LifecyclePolicyState struct {
}

func (LifecyclePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*lifecyclePolicyState)(nil)).Elem()
}

type lifecyclePolicyArgs struct {
	CopyTags               *bool                                  `pulumi:"copyTags"`
	CreateInterval         *int                                   `pulumi:"createInterval"`
	CrossRegionCopyTargets *LifecyclePolicyCrossRegionCopyTargets `pulumi:"crossRegionCopyTargets"`
	DefaultPolicy          *string                                `pulumi:"defaultPolicy"`
	Description            *string                                `pulumi:"description"`
	Exclusions             *LifecyclePolicyExclusions             `pulumi:"exclusions"`
	ExecutionRoleArn       *string                                `pulumi:"executionRoleArn"`
	ExtendDeletion         *bool                                  `pulumi:"extendDeletion"`
	PolicyDetails          *LifecyclePolicyPolicyDetails          `pulumi:"policyDetails"`
	RetainInterval         *int                                   `pulumi:"retainInterval"`
	State                  *string                                `pulumi:"state"`
	Tags                   []aws.Tag                              `pulumi:"tags"`
}

// The set of arguments for constructing a LifecyclePolicy resource.
type LifecyclePolicyArgs struct {
	CopyTags               pulumi.BoolPtrInput
	CreateInterval         pulumi.IntPtrInput
	CrossRegionCopyTargets LifecyclePolicyCrossRegionCopyTargetsPtrInput
	DefaultPolicy          pulumi.StringPtrInput
	Description            pulumi.StringPtrInput
	Exclusions             LifecyclePolicyExclusionsPtrInput
	ExecutionRoleArn       pulumi.StringPtrInput
	ExtendDeletion         pulumi.BoolPtrInput
	PolicyDetails          LifecyclePolicyPolicyDetailsPtrInput
	RetainInterval         pulumi.IntPtrInput
	State                  pulumi.StringPtrInput
	Tags                   aws.TagArrayInput
}

func (LifecyclePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lifecyclePolicyArgs)(nil)).Elem()
}

type LifecyclePolicyInput interface {
	pulumi.Input

	ToLifecyclePolicyOutput() LifecyclePolicyOutput
	ToLifecyclePolicyOutputWithContext(ctx context.Context) LifecyclePolicyOutput
}

func (*LifecyclePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**LifecyclePolicy)(nil)).Elem()
}

func (i *LifecyclePolicy) ToLifecyclePolicyOutput() LifecyclePolicyOutput {
	return i.ToLifecyclePolicyOutputWithContext(context.Background())
}

func (i *LifecyclePolicy) ToLifecyclePolicyOutputWithContext(ctx context.Context) LifecyclePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LifecyclePolicyOutput)
}

type LifecyclePolicyOutput struct{ *pulumi.OutputState }

func (LifecyclePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LifecyclePolicy)(nil)).Elem()
}

func (o LifecyclePolicyOutput) ToLifecyclePolicyOutput() LifecyclePolicyOutput {
	return o
}

func (o LifecyclePolicyOutput) ToLifecyclePolicyOutputWithContext(ctx context.Context) LifecyclePolicyOutput {
	return o
}

func (o LifecyclePolicyOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *LifecyclePolicy) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o LifecyclePolicyOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *LifecyclePolicy) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

func (o LifecyclePolicyOutput) CopyTags() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LifecyclePolicy) pulumi.BoolPtrOutput { return v.CopyTags }).(pulumi.BoolPtrOutput)
}

func (o LifecyclePolicyOutput) CreateInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LifecyclePolicy) pulumi.IntPtrOutput { return v.CreateInterval }).(pulumi.IntPtrOutput)
}

func (o LifecyclePolicyOutput) CrossRegionCopyTargets() LifecyclePolicyCrossRegionCopyTargetsPtrOutput {
	return o.ApplyT(func(v *LifecyclePolicy) LifecyclePolicyCrossRegionCopyTargetsPtrOutput {
		return v.CrossRegionCopyTargets
	}).(LifecyclePolicyCrossRegionCopyTargetsPtrOutput)
}

func (o LifecyclePolicyOutput) DefaultPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LifecyclePolicy) pulumi.StringPtrOutput { return v.DefaultPolicy }).(pulumi.StringPtrOutput)
}

func (o LifecyclePolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LifecyclePolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LifecyclePolicyOutput) Exclusions() LifecyclePolicyExclusionsPtrOutput {
	return o.ApplyT(func(v *LifecyclePolicy) LifecyclePolicyExclusionsPtrOutput { return v.Exclusions }).(LifecyclePolicyExclusionsPtrOutput)
}

func (o LifecyclePolicyOutput) ExecutionRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LifecyclePolicy) pulumi.StringPtrOutput { return v.ExecutionRoleArn }).(pulumi.StringPtrOutput)
}

func (o LifecyclePolicyOutput) ExtendDeletion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LifecyclePolicy) pulumi.BoolPtrOutput { return v.ExtendDeletion }).(pulumi.BoolPtrOutput)
}

func (o LifecyclePolicyOutput) PolicyDetails() LifecyclePolicyPolicyDetailsPtrOutput {
	return o.ApplyT(func(v *LifecyclePolicy) LifecyclePolicyPolicyDetailsPtrOutput { return v.PolicyDetails }).(LifecyclePolicyPolicyDetailsPtrOutput)
}

func (o LifecyclePolicyOutput) RetainInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LifecyclePolicy) pulumi.IntPtrOutput { return v.RetainInterval }).(pulumi.IntPtrOutput)
}

func (o LifecyclePolicyOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LifecyclePolicy) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

func (o LifecyclePolicyOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *LifecyclePolicy) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LifecyclePolicyInput)(nil)).Elem(), &LifecyclePolicy{})
	pulumi.RegisterOutputType(LifecyclePolicyOutput{})
}
