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

// Resource Schema of AWS::EC2::IPAMResourceDiscoveryAssociation Type
type IpamResourceDiscoveryAssociation struct {
	pulumi.CustomResourceState

	// Arn of the IPAM.
	IpamArn pulumi.StringOutput `pulumi:"ipamArn"`
	// The Id of the IPAM this Resource Discovery is associated to.
	IpamId pulumi.StringOutput `pulumi:"ipamId"`
	// The home region of the IPAM.
	IpamRegion pulumi.StringOutput `pulumi:"ipamRegion"`
	// The Amazon Resource Name (ARN) of the resource discovery association is a part of.
	IpamResourceDiscoveryAssociationArn pulumi.StringOutput `pulumi:"ipamResourceDiscoveryAssociationArn"`
	// Id of the IPAM Resource Discovery Association.
	IpamResourceDiscoveryAssociationId pulumi.StringOutput `pulumi:"ipamResourceDiscoveryAssociationId"`
	// The Amazon Resource Name (ARN) of the IPAM Resource Discovery Association.
	IpamResourceDiscoveryId pulumi.StringOutput `pulumi:"ipamResourceDiscoveryId"`
	// If the Resource Discovery Association exists due as part of CreateIpam.
	IsDefault pulumi.BoolOutput `pulumi:"isDefault"`
	// The AWS Account ID for the account where the shared IPAM exists.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// The status of the resource discovery.
	ResourceDiscoveryStatus pulumi.StringOutput `pulumi:"resourceDiscoveryStatus"`
	// The operational state of the Resource Discovery Association. Related to Create/Delete activities.
	State pulumi.StringOutput `pulumi:"state"`
	// An array of key-value pairs to apply to this resource.
	Tags IpamResourceDiscoveryAssociationTagArrayOutput `pulumi:"tags"`
}

// NewIpamResourceDiscoveryAssociation registers a new resource with the given unique name, arguments, and options.
func NewIpamResourceDiscoveryAssociation(ctx *pulumi.Context,
	name string, args *IpamResourceDiscoveryAssociationArgs, opts ...pulumi.ResourceOption) (*IpamResourceDiscoveryAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpamId == nil {
		return nil, errors.New("invalid value for required argument 'IpamId'")
	}
	if args.IpamResourceDiscoveryId == nil {
		return nil, errors.New("invalid value for required argument 'IpamResourceDiscoveryId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"ipamId",
		"ipamResourceDiscoveryId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IpamResourceDiscoveryAssociation
	err := ctx.RegisterResource("aws-native:ec2:IpamResourceDiscoveryAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpamResourceDiscoveryAssociation gets an existing IpamResourceDiscoveryAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpamResourceDiscoveryAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpamResourceDiscoveryAssociationState, opts ...pulumi.ResourceOption) (*IpamResourceDiscoveryAssociation, error) {
	var resource IpamResourceDiscoveryAssociation
	err := ctx.ReadResource("aws-native:ec2:IpamResourceDiscoveryAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpamResourceDiscoveryAssociation resources.
type ipamResourceDiscoveryAssociationState struct {
}

type IpamResourceDiscoveryAssociationState struct {
}

func (IpamResourceDiscoveryAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipamResourceDiscoveryAssociationState)(nil)).Elem()
}

type ipamResourceDiscoveryAssociationArgs struct {
	// The Id of the IPAM this Resource Discovery is associated to.
	IpamId string `pulumi:"ipamId"`
	// The Amazon Resource Name (ARN) of the IPAM Resource Discovery Association.
	IpamResourceDiscoveryId string `pulumi:"ipamResourceDiscoveryId"`
	// An array of key-value pairs to apply to this resource.
	Tags []IpamResourceDiscoveryAssociationTag `pulumi:"tags"`
}

// The set of arguments for constructing a IpamResourceDiscoveryAssociation resource.
type IpamResourceDiscoveryAssociationArgs struct {
	// The Id of the IPAM this Resource Discovery is associated to.
	IpamId pulumi.StringInput
	// The Amazon Resource Name (ARN) of the IPAM Resource Discovery Association.
	IpamResourceDiscoveryId pulumi.StringInput
	// An array of key-value pairs to apply to this resource.
	Tags IpamResourceDiscoveryAssociationTagArrayInput
}

func (IpamResourceDiscoveryAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipamResourceDiscoveryAssociationArgs)(nil)).Elem()
}

type IpamResourceDiscoveryAssociationInput interface {
	pulumi.Input

	ToIpamResourceDiscoveryAssociationOutput() IpamResourceDiscoveryAssociationOutput
	ToIpamResourceDiscoveryAssociationOutputWithContext(ctx context.Context) IpamResourceDiscoveryAssociationOutput
}

func (*IpamResourceDiscoveryAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**IpamResourceDiscoveryAssociation)(nil)).Elem()
}

func (i *IpamResourceDiscoveryAssociation) ToIpamResourceDiscoveryAssociationOutput() IpamResourceDiscoveryAssociationOutput {
	return i.ToIpamResourceDiscoveryAssociationOutputWithContext(context.Background())
}

func (i *IpamResourceDiscoveryAssociation) ToIpamResourceDiscoveryAssociationOutputWithContext(ctx context.Context) IpamResourceDiscoveryAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpamResourceDiscoveryAssociationOutput)
}

type IpamResourceDiscoveryAssociationOutput struct{ *pulumi.OutputState }

func (IpamResourceDiscoveryAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpamResourceDiscoveryAssociation)(nil)).Elem()
}

func (o IpamResourceDiscoveryAssociationOutput) ToIpamResourceDiscoveryAssociationOutput() IpamResourceDiscoveryAssociationOutput {
	return o
}

func (o IpamResourceDiscoveryAssociationOutput) ToIpamResourceDiscoveryAssociationOutputWithContext(ctx context.Context) IpamResourceDiscoveryAssociationOutput {
	return o
}

// Arn of the IPAM.
func (o IpamResourceDiscoveryAssociationOutput) IpamArn() pulumi.StringOutput {
	return o.ApplyT(func(v *IpamResourceDiscoveryAssociation) pulumi.StringOutput { return v.IpamArn }).(pulumi.StringOutput)
}

// The Id of the IPAM this Resource Discovery is associated to.
func (o IpamResourceDiscoveryAssociationOutput) IpamId() pulumi.StringOutput {
	return o.ApplyT(func(v *IpamResourceDiscoveryAssociation) pulumi.StringOutput { return v.IpamId }).(pulumi.StringOutput)
}

// The home region of the IPAM.
func (o IpamResourceDiscoveryAssociationOutput) IpamRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *IpamResourceDiscoveryAssociation) pulumi.StringOutput { return v.IpamRegion }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the resource discovery association is a part of.
func (o IpamResourceDiscoveryAssociationOutput) IpamResourceDiscoveryAssociationArn() pulumi.StringOutput {
	return o.ApplyT(func(v *IpamResourceDiscoveryAssociation) pulumi.StringOutput {
		return v.IpamResourceDiscoveryAssociationArn
	}).(pulumi.StringOutput)
}

// Id of the IPAM Resource Discovery Association.
func (o IpamResourceDiscoveryAssociationOutput) IpamResourceDiscoveryAssociationId() pulumi.StringOutput {
	return o.ApplyT(func(v *IpamResourceDiscoveryAssociation) pulumi.StringOutput {
		return v.IpamResourceDiscoveryAssociationId
	}).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the IPAM Resource Discovery Association.
func (o IpamResourceDiscoveryAssociationOutput) IpamResourceDiscoveryId() pulumi.StringOutput {
	return o.ApplyT(func(v *IpamResourceDiscoveryAssociation) pulumi.StringOutput { return v.IpamResourceDiscoveryId }).(pulumi.StringOutput)
}

// If the Resource Discovery Association exists due as part of CreateIpam.
func (o IpamResourceDiscoveryAssociationOutput) IsDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v *IpamResourceDiscoveryAssociation) pulumi.BoolOutput { return v.IsDefault }).(pulumi.BoolOutput)
}

// The AWS Account ID for the account where the shared IPAM exists.
func (o IpamResourceDiscoveryAssociationOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *IpamResourceDiscoveryAssociation) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// The status of the resource discovery.
func (o IpamResourceDiscoveryAssociationOutput) ResourceDiscoveryStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *IpamResourceDiscoveryAssociation) pulumi.StringOutput { return v.ResourceDiscoveryStatus }).(pulumi.StringOutput)
}

// The operational state of the Resource Discovery Association. Related to Create/Delete activities.
func (o IpamResourceDiscoveryAssociationOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *IpamResourceDiscoveryAssociation) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// An array of key-value pairs to apply to this resource.
func (o IpamResourceDiscoveryAssociationOutput) Tags() IpamResourceDiscoveryAssociationTagArrayOutput {
	return o.ApplyT(func(v *IpamResourceDiscoveryAssociation) IpamResourceDiscoveryAssociationTagArrayOutput {
		return v.Tags
	}).(IpamResourceDiscoveryAssociationTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpamResourceDiscoveryAssociationInput)(nil)).Elem(), &IpamResourceDiscoveryAssociation{})
	pulumi.RegisterOutputType(IpamResourceDiscoveryAssociationOutput{})
}
