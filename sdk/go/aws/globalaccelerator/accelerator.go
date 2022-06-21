// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package globalaccelerator

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::GlobalAccelerator::Accelerator
type Accelerator struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the accelerator.
	AcceleratorArn pulumi.StringOutput `pulumi:"acceleratorArn"`
	// The Domain Name System (DNS) name that Global Accelerator creates that points to your accelerator's static IP addresses.
	DnsName pulumi.StringOutput `pulumi:"dnsName"`
	// Indicates whether an accelerator is enabled. The value is true or false.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// IP Address type.
	IpAddressType AcceleratorIpAddressTypePtrOutput `pulumi:"ipAddressType"`
	// The IP addresses from BYOIP Prefix pool.
	IpAddresses pulumi.StringArrayOutput `pulumi:"ipAddresses"`
	// The IPv4 addresses assigned to the accelerator.
	Ipv4Addresses pulumi.StringArrayOutput `pulumi:"ipv4Addresses"`
	// Name of accelerator.
	Name pulumi.StringOutput       `pulumi:"name"`
	Tags AcceleratorTagArrayOutput `pulumi:"tags"`
}

// NewAccelerator registers a new resource with the given unique name, arguments, and options.
func NewAccelerator(ctx *pulumi.Context,
	name string, args *AcceleratorArgs, opts ...pulumi.ResourceOption) (*Accelerator, error) {
	if args == nil {
		args = &AcceleratorArgs{}
	}

	var resource Accelerator
	err := ctx.RegisterResource("aws-native:globalaccelerator:Accelerator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccelerator gets an existing Accelerator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccelerator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AcceleratorState, opts ...pulumi.ResourceOption) (*Accelerator, error) {
	var resource Accelerator
	err := ctx.ReadResource("aws-native:globalaccelerator:Accelerator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Accelerator resources.
type acceleratorState struct {
}

type AcceleratorState struct {
}

func (AcceleratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*acceleratorState)(nil)).Elem()
}

type acceleratorArgs struct {
	// Indicates whether an accelerator is enabled. The value is true or false.
	Enabled *bool `pulumi:"enabled"`
	// IP Address type.
	IpAddressType *AcceleratorIpAddressType `pulumi:"ipAddressType"`
	// The IP addresses from BYOIP Prefix pool.
	IpAddresses []string `pulumi:"ipAddresses"`
	// Name of accelerator.
	Name *string          `pulumi:"name"`
	Tags []AcceleratorTag `pulumi:"tags"`
}

// The set of arguments for constructing a Accelerator resource.
type AcceleratorArgs struct {
	// Indicates whether an accelerator is enabled. The value is true or false.
	Enabled pulumi.BoolPtrInput
	// IP Address type.
	IpAddressType AcceleratorIpAddressTypePtrInput
	// The IP addresses from BYOIP Prefix pool.
	IpAddresses pulumi.StringArrayInput
	// Name of accelerator.
	Name pulumi.StringPtrInput
	Tags AcceleratorTagArrayInput
}

func (AcceleratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*acceleratorArgs)(nil)).Elem()
}

type AcceleratorInput interface {
	pulumi.Input

	ToAcceleratorOutput() AcceleratorOutput
	ToAcceleratorOutputWithContext(ctx context.Context) AcceleratorOutput
}

func (*Accelerator) ElementType() reflect.Type {
	return reflect.TypeOf((**Accelerator)(nil)).Elem()
}

func (i *Accelerator) ToAcceleratorOutput() AcceleratorOutput {
	return i.ToAcceleratorOutputWithContext(context.Background())
}

func (i *Accelerator) ToAcceleratorOutputWithContext(ctx context.Context) AcceleratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AcceleratorOutput)
}

type AcceleratorOutput struct{ *pulumi.OutputState }

func (AcceleratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Accelerator)(nil)).Elem()
}

func (o AcceleratorOutput) ToAcceleratorOutput() AcceleratorOutput {
	return o
}

func (o AcceleratorOutput) ToAcceleratorOutputWithContext(ctx context.Context) AcceleratorOutput {
	return o
}

// The Amazon Resource Name (ARN) of the accelerator.
func (o AcceleratorOutput) AcceleratorArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Accelerator) pulumi.StringOutput { return v.AcceleratorArn }).(pulumi.StringOutput)
}

// The Domain Name System (DNS) name that Global Accelerator creates that points to your accelerator's static IP addresses.
func (o AcceleratorOutput) DnsName() pulumi.StringOutput {
	return o.ApplyT(func(v *Accelerator) pulumi.StringOutput { return v.DnsName }).(pulumi.StringOutput)
}

// Indicates whether an accelerator is enabled. The value is true or false.
func (o AcceleratorOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Accelerator) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// IP Address type.
func (o AcceleratorOutput) IpAddressType() AcceleratorIpAddressTypePtrOutput {
	return o.ApplyT(func(v *Accelerator) AcceleratorIpAddressTypePtrOutput { return v.IpAddressType }).(AcceleratorIpAddressTypePtrOutput)
}

// The IP addresses from BYOIP Prefix pool.
func (o AcceleratorOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Accelerator) pulumi.StringArrayOutput { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

// The IPv4 addresses assigned to the accelerator.
func (o AcceleratorOutput) Ipv4Addresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Accelerator) pulumi.StringArrayOutput { return v.Ipv4Addresses }).(pulumi.StringArrayOutput)
}

// Name of accelerator.
func (o AcceleratorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Accelerator) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AcceleratorOutput) Tags() AcceleratorTagArrayOutput {
	return o.ApplyT(func(v *Accelerator) AcceleratorTagArrayOutput { return v.Tags }).(AcceleratorTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AcceleratorInput)(nil)).Elem(), &Accelerator{})
	pulumi.RegisterOutputType(AcceleratorOutput{})
}
