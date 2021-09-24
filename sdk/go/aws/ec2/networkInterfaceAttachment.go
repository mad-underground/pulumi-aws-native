// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::NetworkInterfaceAttachment
//
// Deprecated: NetworkInterfaceAttachment is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type NetworkInterfaceAttachment struct {
	pulumi.CustomResourceState

	DeleteOnTermination pulumi.BoolPtrOutput `pulumi:"deleteOnTermination"`
	DeviceIndex         pulumi.StringOutput  `pulumi:"deviceIndex"`
	InstanceId          pulumi.StringOutput  `pulumi:"instanceId"`
	NetworkInterfaceId  pulumi.StringOutput  `pulumi:"networkInterfaceId"`
}

// NewNetworkInterfaceAttachment registers a new resource with the given unique name, arguments, and options.
func NewNetworkInterfaceAttachment(ctx *pulumi.Context,
	name string, args *NetworkInterfaceAttachmentArgs, opts ...pulumi.ResourceOption) (*NetworkInterfaceAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceIndex == nil {
		return nil, errors.New("invalid value for required argument 'DeviceIndex'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.NetworkInterfaceId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkInterfaceId'")
	}
	var resource NetworkInterfaceAttachment
	err := ctx.RegisterResource("aws-native:ec2:NetworkInterfaceAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkInterfaceAttachment gets an existing NetworkInterfaceAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInterfaceAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInterfaceAttachmentState, opts ...pulumi.ResourceOption) (*NetworkInterfaceAttachment, error) {
	var resource NetworkInterfaceAttachment
	err := ctx.ReadResource("aws-native:ec2:NetworkInterfaceAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkInterfaceAttachment resources.
type networkInterfaceAttachmentState struct {
}

type NetworkInterfaceAttachmentState struct {
}

func (NetworkInterfaceAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceAttachmentState)(nil)).Elem()
}

type networkInterfaceAttachmentArgs struct {
	DeleteOnTermination *bool  `pulumi:"deleteOnTermination"`
	DeviceIndex         string `pulumi:"deviceIndex"`
	InstanceId          string `pulumi:"instanceId"`
	NetworkInterfaceId  string `pulumi:"networkInterfaceId"`
}

// The set of arguments for constructing a NetworkInterfaceAttachment resource.
type NetworkInterfaceAttachmentArgs struct {
	DeleteOnTermination pulumi.BoolPtrInput
	DeviceIndex         pulumi.StringInput
	InstanceId          pulumi.StringInput
	NetworkInterfaceId  pulumi.StringInput
}

func (NetworkInterfaceAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceAttachmentArgs)(nil)).Elem()
}

type NetworkInterfaceAttachmentInput interface {
	pulumi.Input

	ToNetworkInterfaceAttachmentOutput() NetworkInterfaceAttachmentOutput
	ToNetworkInterfaceAttachmentOutputWithContext(ctx context.Context) NetworkInterfaceAttachmentOutput
}

func (*NetworkInterfaceAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceAttachment)(nil))
}

func (i *NetworkInterfaceAttachment) ToNetworkInterfaceAttachmentOutput() NetworkInterfaceAttachmentOutput {
	return i.ToNetworkInterfaceAttachmentOutputWithContext(context.Background())
}

func (i *NetworkInterfaceAttachment) ToNetworkInterfaceAttachmentOutputWithContext(ctx context.Context) NetworkInterfaceAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceAttachmentOutput)
}

type NetworkInterfaceAttachmentOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceAttachment)(nil))
}

func (o NetworkInterfaceAttachmentOutput) ToNetworkInterfaceAttachmentOutput() NetworkInterfaceAttachmentOutput {
	return o
}

func (o NetworkInterfaceAttachmentOutput) ToNetworkInterfaceAttachmentOutputWithContext(ctx context.Context) NetworkInterfaceAttachmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NetworkInterfaceAttachmentOutput{})
}