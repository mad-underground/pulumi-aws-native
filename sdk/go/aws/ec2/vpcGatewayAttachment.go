// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::VPCGatewayAttachment
type VpcGatewayAttachment struct {
	pulumi.CustomResourceState

	// Used to identify if this resource is an Internet Gateway or Vpn Gateway Attachment
	AttachmentType pulumi.StringOutput `pulumi:"attachmentType"`
	// The ID of the internet gateway. You must specify either InternetGatewayId or VpnGatewayId, but not both.
	InternetGatewayId pulumi.StringPtrOutput `pulumi:"internetGatewayId"`
	// The ID of the VPC.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The ID of the virtual private gateway. You must specify either InternetGatewayId or VpnGatewayId, but not both.
	VpnGatewayId pulumi.StringPtrOutput `pulumi:"vpnGatewayId"`
}

// NewVpcGatewayAttachment registers a new resource with the given unique name, arguments, and options.
func NewVpcGatewayAttachment(ctx *pulumi.Context,
	name string, args *VpcGatewayAttachmentArgs, opts ...pulumi.ResourceOption) (*VpcGatewayAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"vpcId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcGatewayAttachment
	err := ctx.RegisterResource("aws-native:ec2:VpcGatewayAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcGatewayAttachment gets an existing VpcGatewayAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcGatewayAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcGatewayAttachmentState, opts ...pulumi.ResourceOption) (*VpcGatewayAttachment, error) {
	var resource VpcGatewayAttachment
	err := ctx.ReadResource("aws-native:ec2:VpcGatewayAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcGatewayAttachment resources.
type vpcGatewayAttachmentState struct {
}

type VpcGatewayAttachmentState struct {
}

func (VpcGatewayAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcGatewayAttachmentState)(nil)).Elem()
}

type vpcGatewayAttachmentArgs struct {
	// The ID of the internet gateway. You must specify either InternetGatewayId or VpnGatewayId, but not both.
	InternetGatewayId *string `pulumi:"internetGatewayId"`
	// The ID of the VPC.
	VpcId string `pulumi:"vpcId"`
	// The ID of the virtual private gateway. You must specify either InternetGatewayId or VpnGatewayId, but not both.
	VpnGatewayId *string `pulumi:"vpnGatewayId"`
}

// The set of arguments for constructing a VpcGatewayAttachment resource.
type VpcGatewayAttachmentArgs struct {
	// The ID of the internet gateway. You must specify either InternetGatewayId or VpnGatewayId, but not both.
	InternetGatewayId pulumi.StringPtrInput
	// The ID of the VPC.
	VpcId pulumi.StringInput
	// The ID of the virtual private gateway. You must specify either InternetGatewayId or VpnGatewayId, but not both.
	VpnGatewayId pulumi.StringPtrInput
}

func (VpcGatewayAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcGatewayAttachmentArgs)(nil)).Elem()
}

type VpcGatewayAttachmentInput interface {
	pulumi.Input

	ToVpcGatewayAttachmentOutput() VpcGatewayAttachmentOutput
	ToVpcGatewayAttachmentOutputWithContext(ctx context.Context) VpcGatewayAttachmentOutput
}

func (*VpcGatewayAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcGatewayAttachment)(nil)).Elem()
}

func (i *VpcGatewayAttachment) ToVpcGatewayAttachmentOutput() VpcGatewayAttachmentOutput {
	return i.ToVpcGatewayAttachmentOutputWithContext(context.Background())
}

func (i *VpcGatewayAttachment) ToVpcGatewayAttachmentOutputWithContext(ctx context.Context) VpcGatewayAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcGatewayAttachmentOutput)
}

type VpcGatewayAttachmentOutput struct{ *pulumi.OutputState }

func (VpcGatewayAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcGatewayAttachment)(nil)).Elem()
}

func (o VpcGatewayAttachmentOutput) ToVpcGatewayAttachmentOutput() VpcGatewayAttachmentOutput {
	return o
}

func (o VpcGatewayAttachmentOutput) ToVpcGatewayAttachmentOutputWithContext(ctx context.Context) VpcGatewayAttachmentOutput {
	return o
}

// Used to identify if this resource is an Internet Gateway or Vpn Gateway Attachment
func (o VpcGatewayAttachmentOutput) AttachmentType() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcGatewayAttachment) pulumi.StringOutput { return v.AttachmentType }).(pulumi.StringOutput)
}

// The ID of the internet gateway. You must specify either InternetGatewayId or VpnGatewayId, but not both.
func (o VpcGatewayAttachmentOutput) InternetGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcGatewayAttachment) pulumi.StringPtrOutput { return v.InternetGatewayId }).(pulumi.StringPtrOutput)
}

// The ID of the VPC.
func (o VpcGatewayAttachmentOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcGatewayAttachment) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// The ID of the virtual private gateway. You must specify either InternetGatewayId or VpnGatewayId, but not both.
func (o VpcGatewayAttachmentOutput) VpnGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcGatewayAttachment) pulumi.StringPtrOutput { return v.VpnGatewayId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcGatewayAttachmentInput)(nil)).Elem(), &VpcGatewayAttachment{})
	pulumi.RegisterOutputType(VpcGatewayAttachmentOutput{})
}
