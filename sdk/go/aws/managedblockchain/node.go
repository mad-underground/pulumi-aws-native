// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package managedblockchain

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ManagedBlockchain::Node
//
// Deprecated: Node is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Node struct {
	pulumi.CustomResourceState

	Arn               pulumi.StringOutput     `pulumi:"arn"`
	MemberId          pulumi.StringPtrOutput  `pulumi:"memberId"`
	NetworkId         pulumi.StringOutput     `pulumi:"networkId"`
	NodeConfiguration NodeConfigurationOutput `pulumi:"nodeConfiguration"`
	NodeId            pulumi.StringOutput     `pulumi:"nodeId"`
}

// NewNode registers a new resource with the given unique name, arguments, and options.
func NewNode(ctx *pulumi.Context,
	name string, args *NodeArgs, opts ...pulumi.ResourceOption) (*Node, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkId'")
	}
	if args.NodeConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'NodeConfiguration'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Node
	err := ctx.RegisterResource("aws-native:managedblockchain:Node", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNode gets an existing Node resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNode(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodeState, opts ...pulumi.ResourceOption) (*Node, error) {
	var resource Node
	err := ctx.ReadResource("aws-native:managedblockchain:Node", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Node resources.
type nodeState struct {
}

type NodeState struct {
}

func (NodeState) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeState)(nil)).Elem()
}

type nodeArgs struct {
	MemberId          *string           `pulumi:"memberId"`
	NetworkId         string            `pulumi:"networkId"`
	NodeConfiguration NodeConfiguration `pulumi:"nodeConfiguration"`
}

// The set of arguments for constructing a Node resource.
type NodeArgs struct {
	MemberId          pulumi.StringPtrInput
	NetworkId         pulumi.StringInput
	NodeConfiguration NodeConfigurationInput
}

func (NodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeArgs)(nil)).Elem()
}

type NodeInput interface {
	pulumi.Input

	ToNodeOutput() NodeOutput
	ToNodeOutputWithContext(ctx context.Context) NodeOutput
}

func (*Node) ElementType() reflect.Type {
	return reflect.TypeOf((**Node)(nil)).Elem()
}

func (i *Node) ToNodeOutput() NodeOutput {
	return i.ToNodeOutputWithContext(context.Background())
}

func (i *Node) ToNodeOutputWithContext(ctx context.Context) NodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeOutput)
}

type NodeOutput struct{ *pulumi.OutputState }

func (NodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Node)(nil)).Elem()
}

func (o NodeOutput) ToNodeOutput() NodeOutput {
	return o
}

func (o NodeOutput) ToNodeOutputWithContext(ctx context.Context) NodeOutput {
	return o
}

func (o NodeOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Node) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o NodeOutput) MemberId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Node) pulumi.StringPtrOutput { return v.MemberId }).(pulumi.StringPtrOutput)
}

func (o NodeOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *Node) pulumi.StringOutput { return v.NetworkId }).(pulumi.StringOutput)
}

func (o NodeOutput) NodeConfiguration() NodeConfigurationOutput {
	return o.ApplyT(func(v *Node) NodeConfigurationOutput { return v.NodeConfiguration }).(NodeConfigurationOutput)
}

func (o NodeOutput) NodeId() pulumi.StringOutput {
	return o.ApplyT(func(v *Node) pulumi.StringOutput { return v.NodeId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NodeInput)(nil)).Elem(), &Node{})
	pulumi.RegisterOutputType(NodeOutput{})
}
