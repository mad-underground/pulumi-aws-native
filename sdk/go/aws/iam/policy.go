// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IAM::Policy
//
// Deprecated: Policy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Policy struct {
	pulumi.CustomResourceState

	Groups         pulumi.StringArrayOutput `pulumi:"groups"`
	PolicyDocument pulumi.AnyOutput         `pulumi:"policyDocument"`
	PolicyName     pulumi.StringOutput      `pulumi:"policyName"`
	Roles          pulumi.StringArrayOutput `pulumi:"roles"`
	Users          pulumi.StringArrayOutput `pulumi:"users"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyDocument == nil {
		return nil, errors.New("invalid value for required argument 'PolicyDocument'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Policy
	err := ctx.RegisterResource("aws-native:iam:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("aws-native:iam:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy resources.
type policyState struct {
}

type PolicyState struct {
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	Groups         []string    `pulumi:"groups"`
	PolicyDocument interface{} `pulumi:"policyDocument"`
	PolicyName     *string     `pulumi:"policyName"`
	Roles          []string    `pulumi:"roles"`
	Users          []string    `pulumi:"users"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	Groups         pulumi.StringArrayInput
	PolicyDocument pulumi.Input
	PolicyName     pulumi.StringPtrInput
	Roles          pulumi.StringArrayInput
	Users          pulumi.StringArrayInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}

type PolicyInput interface {
	pulumi.Input

	ToPolicyOutput() PolicyOutput
	ToPolicyOutputWithContext(ctx context.Context) PolicyOutput
}

func (*Policy) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (i *Policy) ToPolicyOutput() PolicyOutput {
	return i.ToPolicyOutputWithContext(context.Background())
}

func (i *Policy) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyOutput)
}

type PolicyOutput struct{ *pulumi.OutputState }

func (PolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (o PolicyOutput) ToPolicyOutput() PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return o
}

func (o PolicyOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringArrayOutput { return v.Groups }).(pulumi.StringArrayOutput)
}

func (o PolicyOutput) PolicyDocument() pulumi.AnyOutput {
	return o.ApplyT(func(v *Policy) pulumi.AnyOutput { return v.PolicyDocument }).(pulumi.AnyOutput)
}

func (o PolicyOutput) PolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.PolicyName }).(pulumi.StringOutput)
}

func (o PolicyOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

func (o PolicyOutput) Users() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringArrayOutput { return v.Users }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyInput)(nil)).Elem(), &Policy{})
	pulumi.RegisterOutputType(PolicyOutput{})
}
