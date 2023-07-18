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

// Resource Schema of AWS::EC2::IPAMPoolCidr Type
type IPAMPoolCidr struct {
	pulumi.CustomResourceState

	// Represents a single IPv4 or IPv6 CIDR
	Cidr pulumi.StringPtrOutput `pulumi:"cidr"`
	// Id of the IPAM Pool Cidr.
	IpamPoolCidrId pulumi.StringOutput `pulumi:"ipamPoolCidrId"`
	// Id of the IPAM Pool.
	IpamPoolId pulumi.StringOutput `pulumi:"ipamPoolId"`
	// The desired netmask length of the provision. If set, IPAM will choose a block of free space with this size and return the CIDR representing it.
	NetmaskLength pulumi.IntPtrOutput `pulumi:"netmaskLength"`
	// Provisioned state of the cidr.
	State pulumi.StringOutput `pulumi:"state"`
}

// NewIPAMPoolCidr registers a new resource with the given unique name, arguments, and options.
func NewIPAMPoolCidr(ctx *pulumi.Context,
	name string, args *IPAMPoolCidrArgs, opts ...pulumi.ResourceOption) (*IPAMPoolCidr, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpamPoolId == nil {
		return nil, errors.New("invalid value for required argument 'IpamPoolId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IPAMPoolCidr
	err := ctx.RegisterResource("aws-native:ec2:IPAMPoolCidr", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIPAMPoolCidr gets an existing IPAMPoolCidr resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIPAMPoolCidr(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IPAMPoolCidrState, opts ...pulumi.ResourceOption) (*IPAMPoolCidr, error) {
	var resource IPAMPoolCidr
	err := ctx.ReadResource("aws-native:ec2:IPAMPoolCidr", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IPAMPoolCidr resources.
type ipampoolCidrState struct {
}

type IPAMPoolCidrState struct {
}

func (IPAMPoolCidrState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipampoolCidrState)(nil)).Elem()
}

type ipampoolCidrArgs struct {
	// Represents a single IPv4 or IPv6 CIDR
	Cidr *string `pulumi:"cidr"`
	// Id of the IPAM Pool.
	IpamPoolId string `pulumi:"ipamPoolId"`
	// The desired netmask length of the provision. If set, IPAM will choose a block of free space with this size and return the CIDR representing it.
	NetmaskLength *int `pulumi:"netmaskLength"`
}

// The set of arguments for constructing a IPAMPoolCidr resource.
type IPAMPoolCidrArgs struct {
	// Represents a single IPv4 or IPv6 CIDR
	Cidr pulumi.StringPtrInput
	// Id of the IPAM Pool.
	IpamPoolId pulumi.StringInput
	// The desired netmask length of the provision. If set, IPAM will choose a block of free space with this size and return the CIDR representing it.
	NetmaskLength pulumi.IntPtrInput
}

func (IPAMPoolCidrArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipampoolCidrArgs)(nil)).Elem()
}

type IPAMPoolCidrInput interface {
	pulumi.Input

	ToIPAMPoolCidrOutput() IPAMPoolCidrOutput
	ToIPAMPoolCidrOutputWithContext(ctx context.Context) IPAMPoolCidrOutput
}

func (*IPAMPoolCidr) ElementType() reflect.Type {
	return reflect.TypeOf((**IPAMPoolCidr)(nil)).Elem()
}

func (i *IPAMPoolCidr) ToIPAMPoolCidrOutput() IPAMPoolCidrOutput {
	return i.ToIPAMPoolCidrOutputWithContext(context.Background())
}

func (i *IPAMPoolCidr) ToIPAMPoolCidrOutputWithContext(ctx context.Context) IPAMPoolCidrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPAMPoolCidrOutput)
}

type IPAMPoolCidrOutput struct{ *pulumi.OutputState }

func (IPAMPoolCidrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPAMPoolCidr)(nil)).Elem()
}

func (o IPAMPoolCidrOutput) ToIPAMPoolCidrOutput() IPAMPoolCidrOutput {
	return o
}

func (o IPAMPoolCidrOutput) ToIPAMPoolCidrOutputWithContext(ctx context.Context) IPAMPoolCidrOutput {
	return o
}

// Represents a single IPv4 or IPv6 CIDR
func (o IPAMPoolCidrOutput) Cidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPAMPoolCidr) pulumi.StringPtrOutput { return v.Cidr }).(pulumi.StringPtrOutput)
}

// Id of the IPAM Pool Cidr.
func (o IPAMPoolCidrOutput) IpamPoolCidrId() pulumi.StringOutput {
	return o.ApplyT(func(v *IPAMPoolCidr) pulumi.StringOutput { return v.IpamPoolCidrId }).(pulumi.StringOutput)
}

// Id of the IPAM Pool.
func (o IPAMPoolCidrOutput) IpamPoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *IPAMPoolCidr) pulumi.StringOutput { return v.IpamPoolId }).(pulumi.StringOutput)
}

// The desired netmask length of the provision. If set, IPAM will choose a block of free space with this size and return the CIDR representing it.
func (o IPAMPoolCidrOutput) NetmaskLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *IPAMPoolCidr) pulumi.IntPtrOutput { return v.NetmaskLength }).(pulumi.IntPtrOutput)
}

// Provisioned state of the cidr.
func (o IPAMPoolCidrOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *IPAMPoolCidr) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IPAMPoolCidrInput)(nil)).Elem(), &IPAMPoolCidr{})
	pulumi.RegisterOutputType(IPAMPoolCidrOutput{})
}
