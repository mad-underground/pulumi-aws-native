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

// Resource Type definition for AWS::EC2::TransitGatewayAttachment
type TransitGatewayAttachment struct {
	pulumi.CustomResourceState

	// The options for the transit gateway vpc attachment.
	Options          OptionsPropertiesPtrOutput             `pulumi:"options"`
	SubnetIds        pulumi.StringArrayOutput               `pulumi:"subnetIds"`
	Tags             TransitGatewayAttachmentTagArrayOutput `pulumi:"tags"`
	TransitGatewayId pulumi.StringOutput                    `pulumi:"transitGatewayId"`
	VpcId            pulumi.StringOutput                    `pulumi:"vpcId"`
}

// NewTransitGatewayAttachment registers a new resource with the given unique name, arguments, and options.
func NewTransitGatewayAttachment(ctx *pulumi.Context,
	name string, args *TransitGatewayAttachmentArgs, opts ...pulumi.ResourceOption) (*TransitGatewayAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SubnetIds == nil {
		return nil, errors.New("invalid value for required argument 'SubnetIds'")
	}
	if args.TransitGatewayId == nil {
		return nil, errors.New("invalid value for required argument 'TransitGatewayId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TransitGatewayAttachment
	err := ctx.RegisterResource("aws-native:ec2:TransitGatewayAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransitGatewayAttachment gets an existing TransitGatewayAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransitGatewayAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransitGatewayAttachmentState, opts ...pulumi.ResourceOption) (*TransitGatewayAttachment, error) {
	var resource TransitGatewayAttachment
	err := ctx.ReadResource("aws-native:ec2:TransitGatewayAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TransitGatewayAttachment resources.
type transitGatewayAttachmentState struct {
}

type TransitGatewayAttachmentState struct {
}

func (TransitGatewayAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*transitGatewayAttachmentState)(nil)).Elem()
}

type transitGatewayAttachmentArgs struct {
	// The options for the transit gateway vpc attachment.
	Options          *OptionsProperties            `pulumi:"options"`
	SubnetIds        []string                      `pulumi:"subnetIds"`
	Tags             []TransitGatewayAttachmentTag `pulumi:"tags"`
	TransitGatewayId string                        `pulumi:"transitGatewayId"`
	VpcId            string                        `pulumi:"vpcId"`
}

// The set of arguments for constructing a TransitGatewayAttachment resource.
type TransitGatewayAttachmentArgs struct {
	// The options for the transit gateway vpc attachment.
	Options          OptionsPropertiesPtrInput
	SubnetIds        pulumi.StringArrayInput
	Tags             TransitGatewayAttachmentTagArrayInput
	TransitGatewayId pulumi.StringInput
	VpcId            pulumi.StringInput
}

func (TransitGatewayAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transitGatewayAttachmentArgs)(nil)).Elem()
}

type TransitGatewayAttachmentInput interface {
	pulumi.Input

	ToTransitGatewayAttachmentOutput() TransitGatewayAttachmentOutput
	ToTransitGatewayAttachmentOutputWithContext(ctx context.Context) TransitGatewayAttachmentOutput
}

func (*TransitGatewayAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**TransitGatewayAttachment)(nil)).Elem()
}

func (i *TransitGatewayAttachment) ToTransitGatewayAttachmentOutput() TransitGatewayAttachmentOutput {
	return i.ToTransitGatewayAttachmentOutputWithContext(context.Background())
}

func (i *TransitGatewayAttachment) ToTransitGatewayAttachmentOutputWithContext(ctx context.Context) TransitGatewayAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitGatewayAttachmentOutput)
}

type TransitGatewayAttachmentOutput struct{ *pulumi.OutputState }

func (TransitGatewayAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransitGatewayAttachment)(nil)).Elem()
}

func (o TransitGatewayAttachmentOutput) ToTransitGatewayAttachmentOutput() TransitGatewayAttachmentOutput {
	return o
}

func (o TransitGatewayAttachmentOutput) ToTransitGatewayAttachmentOutputWithContext(ctx context.Context) TransitGatewayAttachmentOutput {
	return o
}

// The options for the transit gateway vpc attachment.
func (o TransitGatewayAttachmentOutput) Options() OptionsPropertiesPtrOutput {
	return o.ApplyT(func(v *TransitGatewayAttachment) OptionsPropertiesPtrOutput { return v.Options }).(OptionsPropertiesPtrOutput)
}

func (o TransitGatewayAttachmentOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TransitGatewayAttachment) pulumi.StringArrayOutput { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

func (o TransitGatewayAttachmentOutput) Tags() TransitGatewayAttachmentTagArrayOutput {
	return o.ApplyT(func(v *TransitGatewayAttachment) TransitGatewayAttachmentTagArrayOutput { return v.Tags }).(TransitGatewayAttachmentTagArrayOutput)
}

func (o TransitGatewayAttachmentOutput) TransitGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitGatewayAttachment) pulumi.StringOutput { return v.TransitGatewayId }).(pulumi.StringOutput)
}

func (o TransitGatewayAttachmentOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitGatewayAttachment) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TransitGatewayAttachmentInput)(nil)).Elem(), &TransitGatewayAttachment{})
	pulumi.RegisterOutputType(TransitGatewayAttachmentOutput{})
}
