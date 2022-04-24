// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkfirewall

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource type definition for AWS::NetworkFirewall::FirewallPolicy
type FirewallPolicy struct {
	pulumi.CustomResourceState

	Description        pulumi.StringPtrOutput       `pulumi:"description"`
	FirewallPolicy     FirewallPolicyTypeOutput     `pulumi:"firewallPolicy"`
	FirewallPolicyArn  pulumi.StringOutput          `pulumi:"firewallPolicyArn"`
	FirewallPolicyId   pulumi.StringOutput          `pulumi:"firewallPolicyId"`
	FirewallPolicyName pulumi.StringOutput          `pulumi:"firewallPolicyName"`
	Tags               FirewallPolicyTagArrayOutput `pulumi:"tags"`
}

// NewFirewallPolicy registers a new resource with the given unique name, arguments, and options.
func NewFirewallPolicy(ctx *pulumi.Context,
	name string, args *FirewallPolicyArgs, opts ...pulumi.ResourceOption) (*FirewallPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FirewallPolicy == nil {
		return nil, errors.New("invalid value for required argument 'FirewallPolicy'")
	}
	var resource FirewallPolicy
	err := ctx.RegisterResource("aws-native:networkfirewall:FirewallPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallPolicy gets an existing FirewallPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallPolicyState, opts ...pulumi.ResourceOption) (*FirewallPolicy, error) {
	var resource FirewallPolicy
	err := ctx.ReadResource("aws-native:networkfirewall:FirewallPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallPolicy resources.
type firewallPolicyState struct {
}

type FirewallPolicyState struct {
}

func (FirewallPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyState)(nil)).Elem()
}

type firewallPolicyArgs struct {
	Description        *string             `pulumi:"description"`
	FirewallPolicy     FirewallPolicyType  `pulumi:"firewallPolicy"`
	FirewallPolicyName *string             `pulumi:"firewallPolicyName"`
	Tags               []FirewallPolicyTag `pulumi:"tags"`
}

// The set of arguments for constructing a FirewallPolicy resource.
type FirewallPolicyArgs struct {
	Description        pulumi.StringPtrInput
	FirewallPolicy     FirewallPolicyTypeInput
	FirewallPolicyName pulumi.StringPtrInput
	Tags               FirewallPolicyTagArrayInput
}

func (FirewallPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyArgs)(nil)).Elem()
}

type FirewallPolicyInput interface {
	pulumi.Input

	ToFirewallPolicyOutput() FirewallPolicyOutput
	ToFirewallPolicyOutputWithContext(ctx context.Context) FirewallPolicyOutput
}

func (*FirewallPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicy)(nil)).Elem()
}

func (i *FirewallPolicy) ToFirewallPolicyOutput() FirewallPolicyOutput {
	return i.ToFirewallPolicyOutputWithContext(context.Background())
}

func (i *FirewallPolicy) ToFirewallPolicyOutputWithContext(ctx context.Context) FirewallPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyOutput)
}

type FirewallPolicyOutput struct{ *pulumi.OutputState }

func (FirewallPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicy)(nil)).Elem()
}

func (o FirewallPolicyOutput) ToFirewallPolicyOutput() FirewallPolicyOutput {
	return o
}

func (o FirewallPolicyOutput) ToFirewallPolicyOutputWithContext(ctx context.Context) FirewallPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallPolicyInput)(nil)).Elem(), &FirewallPolicy{})
	pulumi.RegisterOutputType(FirewallPolicyOutput{})
}
