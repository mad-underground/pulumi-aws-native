// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package imagebuilder

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::ImageBuilder::LifecyclePolicy
type LifecyclePolicy struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the lifecycle policy.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of the lifecycle policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The execution role of the lifecycle policy.
	ExecutionRole pulumi.StringOutput `pulumi:"executionRole"`
	// The name of the lifecycle policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// The policy details of the lifecycle policy.
	PolicyDetails LifecyclePolicyPolicyDetailArrayOutput `pulumi:"policyDetails"`
	// The resource selection of the lifecycle policy.
	ResourceSelection LifecyclePolicyResourceSelectionOutput `pulumi:"resourceSelection"`
	// The resource type of the lifecycle policy.
	ResourceType LifecyclePolicyResourceTypeOutput `pulumi:"resourceType"`
	// The status of the lifecycle policy.
	Status LifecyclePolicyStatusPtrOutput `pulumi:"status"`
	// The tags associated with the lifecycle policy.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewLifecyclePolicy registers a new resource with the given unique name, arguments, and options.
func NewLifecyclePolicy(ctx *pulumi.Context,
	name string, args *LifecyclePolicyArgs, opts ...pulumi.ResourceOption) (*LifecyclePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExecutionRole == nil {
		return nil, errors.New("invalid value for required argument 'ExecutionRole'")
	}
	if args.PolicyDetails == nil {
		return nil, errors.New("invalid value for required argument 'PolicyDetails'")
	}
	if args.ResourceSelection == nil {
		return nil, errors.New("invalid value for required argument 'ResourceSelection'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"name",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LifecyclePolicy
	err := ctx.RegisterResource("aws-native:imagebuilder:LifecyclePolicy", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws-native:imagebuilder:LifecyclePolicy", name, id, state, &resource, opts...)
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
	// The description of the lifecycle policy.
	Description *string `pulumi:"description"`
	// The execution role of the lifecycle policy.
	ExecutionRole string `pulumi:"executionRole"`
	// The name of the lifecycle policy.
	Name *string `pulumi:"name"`
	// The policy details of the lifecycle policy.
	PolicyDetails []LifecyclePolicyPolicyDetail `pulumi:"policyDetails"`
	// The resource selection of the lifecycle policy.
	ResourceSelection LifecyclePolicyResourceSelection `pulumi:"resourceSelection"`
	// The resource type of the lifecycle policy.
	ResourceType LifecyclePolicyResourceType `pulumi:"resourceType"`
	// The status of the lifecycle policy.
	Status *LifecyclePolicyStatus `pulumi:"status"`
	// The tags associated with the lifecycle policy.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a LifecyclePolicy resource.
type LifecyclePolicyArgs struct {
	// The description of the lifecycle policy.
	Description pulumi.StringPtrInput
	// The execution role of the lifecycle policy.
	ExecutionRole pulumi.StringInput
	// The name of the lifecycle policy.
	Name pulumi.StringPtrInput
	// The policy details of the lifecycle policy.
	PolicyDetails LifecyclePolicyPolicyDetailArrayInput
	// The resource selection of the lifecycle policy.
	ResourceSelection LifecyclePolicyResourceSelectionInput
	// The resource type of the lifecycle policy.
	ResourceType LifecyclePolicyResourceTypeInput
	// The status of the lifecycle policy.
	Status LifecyclePolicyStatusPtrInput
	// The tags associated with the lifecycle policy.
	Tags pulumi.StringMapInput
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

// The Amazon Resource Name (ARN) of the lifecycle policy.
func (o LifecyclePolicyOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *LifecyclePolicy) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The description of the lifecycle policy.
func (o LifecyclePolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LifecyclePolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The execution role of the lifecycle policy.
func (o LifecyclePolicyOutput) ExecutionRole() pulumi.StringOutput {
	return o.ApplyT(func(v *LifecyclePolicy) pulumi.StringOutput { return v.ExecutionRole }).(pulumi.StringOutput)
}

// The name of the lifecycle policy.
func (o LifecyclePolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LifecyclePolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The policy details of the lifecycle policy.
func (o LifecyclePolicyOutput) PolicyDetails() LifecyclePolicyPolicyDetailArrayOutput {
	return o.ApplyT(func(v *LifecyclePolicy) LifecyclePolicyPolicyDetailArrayOutput { return v.PolicyDetails }).(LifecyclePolicyPolicyDetailArrayOutput)
}

// The resource selection of the lifecycle policy.
func (o LifecyclePolicyOutput) ResourceSelection() LifecyclePolicyResourceSelectionOutput {
	return o.ApplyT(func(v *LifecyclePolicy) LifecyclePolicyResourceSelectionOutput { return v.ResourceSelection }).(LifecyclePolicyResourceSelectionOutput)
}

// The resource type of the lifecycle policy.
func (o LifecyclePolicyOutput) ResourceType() LifecyclePolicyResourceTypeOutput {
	return o.ApplyT(func(v *LifecyclePolicy) LifecyclePolicyResourceTypeOutput { return v.ResourceType }).(LifecyclePolicyResourceTypeOutput)
}

// The status of the lifecycle policy.
func (o LifecyclePolicyOutput) Status() LifecyclePolicyStatusPtrOutput {
	return o.ApplyT(func(v *LifecyclePolicy) LifecyclePolicyStatusPtrOutput { return v.Status }).(LifecyclePolicyStatusPtrOutput)
}

// The tags associated with the lifecycle policy.
func (o LifecyclePolicyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LifecyclePolicy) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LifecyclePolicyInput)(nil)).Elem(), &LifecyclePolicy{})
	pulumi.RegisterOutputType(LifecyclePolicyOutput{})
}
