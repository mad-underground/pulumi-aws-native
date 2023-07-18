// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aps

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// RuleGroupsNamespace schema for cloudformation.
type RuleGroupsNamespace struct {
	pulumi.CustomResourceState

	// The RuleGroupsNamespace ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The RuleGroupsNamespace data.
	Data pulumi.StringOutput `pulumi:"data"`
	// The RuleGroupsNamespace name.
	Name pulumi.StringOutput `pulumi:"name"`
	// An array of key-value pairs to apply to this resource.
	Tags RuleGroupsNamespaceTagArrayOutput `pulumi:"tags"`
	// Required to identify a specific APS Workspace associated with this RuleGroupsNamespace.
	Workspace pulumi.StringOutput `pulumi:"workspace"`
}

// NewRuleGroupsNamespace registers a new resource with the given unique name, arguments, and options.
func NewRuleGroupsNamespace(ctx *pulumi.Context,
	name string, args *RuleGroupsNamespaceArgs, opts ...pulumi.ResourceOption) (*RuleGroupsNamespace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Data == nil {
		return nil, errors.New("invalid value for required argument 'Data'")
	}
	if args.Workspace == nil {
		return nil, errors.New("invalid value for required argument 'Workspace'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RuleGroupsNamespace
	err := ctx.RegisterResource("aws-native:aps:RuleGroupsNamespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuleGroupsNamespace gets an existing RuleGroupsNamespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuleGroupsNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleGroupsNamespaceState, opts ...pulumi.ResourceOption) (*RuleGroupsNamespace, error) {
	var resource RuleGroupsNamespace
	err := ctx.ReadResource("aws-native:aps:RuleGroupsNamespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RuleGroupsNamespace resources.
type ruleGroupsNamespaceState struct {
}

type RuleGroupsNamespaceState struct {
}

func (RuleGroupsNamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleGroupsNamespaceState)(nil)).Elem()
}

type ruleGroupsNamespaceArgs struct {
	// The RuleGroupsNamespace data.
	Data string `pulumi:"data"`
	// The RuleGroupsNamespace name.
	Name *string `pulumi:"name"`
	// An array of key-value pairs to apply to this resource.
	Tags []RuleGroupsNamespaceTag `pulumi:"tags"`
	// Required to identify a specific APS Workspace associated with this RuleGroupsNamespace.
	Workspace string `pulumi:"workspace"`
}

// The set of arguments for constructing a RuleGroupsNamespace resource.
type RuleGroupsNamespaceArgs struct {
	// The RuleGroupsNamespace data.
	Data pulumi.StringInput
	// The RuleGroupsNamespace name.
	Name pulumi.StringPtrInput
	// An array of key-value pairs to apply to this resource.
	Tags RuleGroupsNamespaceTagArrayInput
	// Required to identify a specific APS Workspace associated with this RuleGroupsNamespace.
	Workspace pulumi.StringInput
}

func (RuleGroupsNamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleGroupsNamespaceArgs)(nil)).Elem()
}

type RuleGroupsNamespaceInput interface {
	pulumi.Input

	ToRuleGroupsNamespaceOutput() RuleGroupsNamespaceOutput
	ToRuleGroupsNamespaceOutputWithContext(ctx context.Context) RuleGroupsNamespaceOutput
}

func (*RuleGroupsNamespace) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleGroupsNamespace)(nil)).Elem()
}

func (i *RuleGroupsNamespace) ToRuleGroupsNamespaceOutput() RuleGroupsNamespaceOutput {
	return i.ToRuleGroupsNamespaceOutputWithContext(context.Background())
}

func (i *RuleGroupsNamespace) ToRuleGroupsNamespaceOutputWithContext(ctx context.Context) RuleGroupsNamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleGroupsNamespaceOutput)
}

type RuleGroupsNamespaceOutput struct{ *pulumi.OutputState }

func (RuleGroupsNamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleGroupsNamespace)(nil)).Elem()
}

func (o RuleGroupsNamespaceOutput) ToRuleGroupsNamespaceOutput() RuleGroupsNamespaceOutput {
	return o
}

func (o RuleGroupsNamespaceOutput) ToRuleGroupsNamespaceOutputWithContext(ctx context.Context) RuleGroupsNamespaceOutput {
	return o
}

// The RuleGroupsNamespace ARN.
func (o RuleGroupsNamespaceOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleGroupsNamespace) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The RuleGroupsNamespace data.
func (o RuleGroupsNamespaceOutput) Data() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleGroupsNamespace) pulumi.StringOutput { return v.Data }).(pulumi.StringOutput)
}

// The RuleGroupsNamespace name.
func (o RuleGroupsNamespaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleGroupsNamespace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// An array of key-value pairs to apply to this resource.
func (o RuleGroupsNamespaceOutput) Tags() RuleGroupsNamespaceTagArrayOutput {
	return o.ApplyT(func(v *RuleGroupsNamespace) RuleGroupsNamespaceTagArrayOutput { return v.Tags }).(RuleGroupsNamespaceTagArrayOutput)
}

// Required to identify a specific APS Workspace associated with this RuleGroupsNamespace.
func (o RuleGroupsNamespaceOutput) Workspace() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleGroupsNamespace) pulumi.StringOutput { return v.Workspace }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuleGroupsNamespaceInput)(nil)).Elem(), &RuleGroupsNamespace{})
	pulumi.RegisterOutputType(RuleGroupsNamespaceOutput{})
}
