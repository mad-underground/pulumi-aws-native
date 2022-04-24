// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appmesh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppMesh::VirtualGateway
//
// Deprecated: VirtualGateway is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type VirtualGateway struct {
	pulumi.CustomResourceState

	Arn                pulumi.StringOutput          `pulumi:"arn"`
	MeshName           pulumi.StringOutput          `pulumi:"meshName"`
	MeshOwner          pulumi.StringPtrOutput       `pulumi:"meshOwner"`
	ResourceOwner      pulumi.StringOutput          `pulumi:"resourceOwner"`
	Spec               VirtualGatewaySpecOutput     `pulumi:"spec"`
	Tags               VirtualGatewayTagArrayOutput `pulumi:"tags"`
	Uid                pulumi.StringOutput          `pulumi:"uid"`
	VirtualGatewayName pulumi.StringPtrOutput       `pulumi:"virtualGatewayName"`
}

// NewVirtualGateway registers a new resource with the given unique name, arguments, and options.
func NewVirtualGateway(ctx *pulumi.Context,
	name string, args *VirtualGatewayArgs, opts ...pulumi.ResourceOption) (*VirtualGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MeshName == nil {
		return nil, errors.New("invalid value for required argument 'MeshName'")
	}
	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	var resource VirtualGateway
	err := ctx.RegisterResource("aws-native:appmesh:VirtualGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualGateway gets an existing VirtualGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualGatewayState, opts ...pulumi.ResourceOption) (*VirtualGateway, error) {
	var resource VirtualGateway
	err := ctx.ReadResource("aws-native:appmesh:VirtualGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualGateway resources.
type virtualGatewayState struct {
}

type VirtualGatewayState struct {
}

func (VirtualGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualGatewayState)(nil)).Elem()
}

type virtualGatewayArgs struct {
	MeshName           string              `pulumi:"meshName"`
	MeshOwner          *string             `pulumi:"meshOwner"`
	Spec               VirtualGatewaySpec  `pulumi:"spec"`
	Tags               []VirtualGatewayTag `pulumi:"tags"`
	VirtualGatewayName *string             `pulumi:"virtualGatewayName"`
}

// The set of arguments for constructing a VirtualGateway resource.
type VirtualGatewayArgs struct {
	MeshName           pulumi.StringInput
	MeshOwner          pulumi.StringPtrInput
	Spec               VirtualGatewaySpecInput
	Tags               VirtualGatewayTagArrayInput
	VirtualGatewayName pulumi.StringPtrInput
}

func (VirtualGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualGatewayArgs)(nil)).Elem()
}

type VirtualGatewayInput interface {
	pulumi.Input

	ToVirtualGatewayOutput() VirtualGatewayOutput
	ToVirtualGatewayOutputWithContext(ctx context.Context) VirtualGatewayOutput
}

func (*VirtualGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualGateway)(nil)).Elem()
}

func (i *VirtualGateway) ToVirtualGatewayOutput() VirtualGatewayOutput {
	return i.ToVirtualGatewayOutputWithContext(context.Background())
}

func (i *VirtualGateway) ToVirtualGatewayOutputWithContext(ctx context.Context) VirtualGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualGatewayOutput)
}

type VirtualGatewayOutput struct{ *pulumi.OutputState }

func (VirtualGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualGateway)(nil)).Elem()
}

func (o VirtualGatewayOutput) ToVirtualGatewayOutput() VirtualGatewayOutput {
	return o
}

func (o VirtualGatewayOutput) ToVirtualGatewayOutputWithContext(ctx context.Context) VirtualGatewayOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualGatewayInput)(nil)).Elem(), &VirtualGateway{})
	pulumi.RegisterOutputType(VirtualGatewayOutput{})
}
