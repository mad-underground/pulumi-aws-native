// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package verifiedpermissions

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::VerifiedPermissions::Policy Resource Type
type Policy struct {
	pulumi.CustomResourceState

	Definition    PolicyDefinitionOutput `pulumi:"definition"`
	PolicyId      pulumi.StringOutput    `pulumi:"policyId"`
	PolicyStoreId pulumi.StringPtrOutput `pulumi:"policyStoreId"`
	PolicyType    PolicyTypeOutput       `pulumi:"policyType"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Definition == nil {
		return nil, errors.New("invalid value for required argument 'Definition'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Policy
	err := ctx.RegisterResource("aws-native:verifiedpermissions:Policy", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws-native:verifiedpermissions:Policy", name, id, state, &resource, opts...)
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
	Definition    PolicyDefinition `pulumi:"definition"`
	PolicyStoreId *string          `pulumi:"policyStoreId"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	Definition    PolicyDefinitionInput
	PolicyStoreId pulumi.StringPtrInput
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

func (o PolicyOutput) Definition() PolicyDefinitionOutput {
	return o.ApplyT(func(v *Policy) PolicyDefinitionOutput { return v.Definition }).(PolicyDefinitionOutput)
}

func (o PolicyOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.PolicyId }).(pulumi.StringOutput)
}

func (o PolicyOutput) PolicyStoreId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringPtrOutput { return v.PolicyStoreId }).(pulumi.StringPtrOutput)
}

func (o PolicyOutput) PolicyType() PolicyTypeOutput {
	return o.ApplyT(func(v *Policy) PolicyTypeOutput { return v.PolicyType }).(PolicyTypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyInput)(nil)).Elem(), &Policy{})
	pulumi.RegisterOutputType(PolicyOutput{})
}
