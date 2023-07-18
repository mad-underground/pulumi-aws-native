// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::VPC
type VPC struct {
	pulumi.CustomResourceState

	// The primary IPv4 CIDR block for the VPC.
	CidrBlock pulumi.StringPtrOutput `pulumi:"cidrBlock"`
	// A list of IPv4 CIDR block association IDs for the VPC.
	CidrBlockAssociations pulumi.StringArrayOutput `pulumi:"cidrBlockAssociations"`
	// The default network ACL ID that is associated with the VPC.
	DefaultNetworkAcl pulumi.StringOutput `pulumi:"defaultNetworkAcl"`
	// The default security group ID that is associated with the VPC.
	DefaultSecurityGroup pulumi.StringOutput `pulumi:"defaultSecurityGroup"`
	// Indicates whether the instances launched in the VPC get DNS hostnames. If enabled, instances in the VPC get DNS hostnames; otherwise, they do not. Disabled by default for nondefault VPCs.
	EnableDnsHostnames pulumi.BoolPtrOutput `pulumi:"enableDnsHostnames"`
	// Indicates whether the DNS resolution is supported for the VPC. If enabled, queries to the Amazon provided DNS server at the 169.254.169.253 IP address, or the reserved IP address at the base of the VPC network range "plus two" succeed. If disabled, the Amazon provided DNS service in the VPC that resolves public DNS hostnames to IP addresses is not enabled. Enabled by default.
	EnableDnsSupport pulumi.BoolPtrOutput `pulumi:"enableDnsSupport"`
	// The allowed tenancy of instances launched into the VPC.
	//
	// "default": An instance launched into the VPC runs on shared hardware by default, unless you explicitly specify a different tenancy during instance launch.
	//
	// "dedicated": An instance launched into the VPC is a Dedicated Instance by default, unless you explicitly specify a tenancy of host during instance launch. You cannot specify a tenancy of default during instance launch.
	//
	// Updating InstanceTenancy requires no replacement only if you are updating its value from "dedicated" to "default". Updating InstanceTenancy from "default" to "dedicated" requires replacement.
	InstanceTenancy pulumi.StringPtrOutput `pulumi:"instanceTenancy"`
	// The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR
	Ipv4IpamPoolId pulumi.StringPtrOutput `pulumi:"ipv4IpamPoolId"`
	// The netmask length of the IPv4 CIDR you want to allocate to this VPC from an Amazon VPC IP Address Manager (IPAM) pool
	Ipv4NetmaskLength pulumi.IntPtrOutput `pulumi:"ipv4NetmaskLength"`
	// A list of IPv6 CIDR blocks that are associated with the VPC.
	Ipv6CidrBlocks pulumi.StringArrayOutput `pulumi:"ipv6CidrBlocks"`
	// The tags for the VPC.
	Tags VPCTagArrayOutput `pulumi:"tags"`
	// The Id for the model.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewVPC registers a new resource with the given unique name, arguments, and options.
func NewVPC(ctx *pulumi.Context,
	name string, args *VPCArgs, opts ...pulumi.ResourceOption) (*VPC, error) {
	if args == nil {
		args = &VPCArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VPC
	err := ctx.RegisterResource("aws-native:ec2:VPC", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVPC gets an existing VPC resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVPC(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VPCState, opts ...pulumi.ResourceOption) (*VPC, error) {
	var resource VPC
	err := ctx.ReadResource("aws-native:ec2:VPC", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VPC resources.
type vpcState struct {
}

type VPCState struct {
}

func (VPCState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcState)(nil)).Elem()
}

type vpcArgs struct {
	// The primary IPv4 CIDR block for the VPC.
	CidrBlock *string `pulumi:"cidrBlock"`
	// Indicates whether the instances launched in the VPC get DNS hostnames. If enabled, instances in the VPC get DNS hostnames; otherwise, they do not. Disabled by default for nondefault VPCs.
	EnableDnsHostnames *bool `pulumi:"enableDnsHostnames"`
	// Indicates whether the DNS resolution is supported for the VPC. If enabled, queries to the Amazon provided DNS server at the 169.254.169.253 IP address, or the reserved IP address at the base of the VPC network range "plus two" succeed. If disabled, the Amazon provided DNS service in the VPC that resolves public DNS hostnames to IP addresses is not enabled. Enabled by default.
	EnableDnsSupport *bool `pulumi:"enableDnsSupport"`
	// The allowed tenancy of instances launched into the VPC.
	//
	// "default": An instance launched into the VPC runs on shared hardware by default, unless you explicitly specify a different tenancy during instance launch.
	//
	// "dedicated": An instance launched into the VPC is a Dedicated Instance by default, unless you explicitly specify a tenancy of host during instance launch. You cannot specify a tenancy of default during instance launch.
	//
	// Updating InstanceTenancy requires no replacement only if you are updating its value from "dedicated" to "default". Updating InstanceTenancy from "default" to "dedicated" requires replacement.
	InstanceTenancy *string `pulumi:"instanceTenancy"`
	// The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR
	Ipv4IpamPoolId *string `pulumi:"ipv4IpamPoolId"`
	// The netmask length of the IPv4 CIDR you want to allocate to this VPC from an Amazon VPC IP Address Manager (IPAM) pool
	Ipv4NetmaskLength *int `pulumi:"ipv4NetmaskLength"`
	// The tags for the VPC.
	Tags []VPCTag `pulumi:"tags"`
}

// The set of arguments for constructing a VPC resource.
type VPCArgs struct {
	// The primary IPv4 CIDR block for the VPC.
	CidrBlock pulumi.StringPtrInput
	// Indicates whether the instances launched in the VPC get DNS hostnames. If enabled, instances in the VPC get DNS hostnames; otherwise, they do not. Disabled by default for nondefault VPCs.
	EnableDnsHostnames pulumi.BoolPtrInput
	// Indicates whether the DNS resolution is supported for the VPC. If enabled, queries to the Amazon provided DNS server at the 169.254.169.253 IP address, or the reserved IP address at the base of the VPC network range "plus two" succeed. If disabled, the Amazon provided DNS service in the VPC that resolves public DNS hostnames to IP addresses is not enabled. Enabled by default.
	EnableDnsSupport pulumi.BoolPtrInput
	// The allowed tenancy of instances launched into the VPC.
	//
	// "default": An instance launched into the VPC runs on shared hardware by default, unless you explicitly specify a different tenancy during instance launch.
	//
	// "dedicated": An instance launched into the VPC is a Dedicated Instance by default, unless you explicitly specify a tenancy of host during instance launch. You cannot specify a tenancy of default during instance launch.
	//
	// Updating InstanceTenancy requires no replacement only if you are updating its value from "dedicated" to "default". Updating InstanceTenancy from "default" to "dedicated" requires replacement.
	InstanceTenancy pulumi.StringPtrInput
	// The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR
	Ipv4IpamPoolId pulumi.StringPtrInput
	// The netmask length of the IPv4 CIDR you want to allocate to this VPC from an Amazon VPC IP Address Manager (IPAM) pool
	Ipv4NetmaskLength pulumi.IntPtrInput
	// The tags for the VPC.
	Tags VPCTagArrayInput
}

func (VPCArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcArgs)(nil)).Elem()
}

type VPCInput interface {
	pulumi.Input

	ToVPCOutput() VPCOutput
	ToVPCOutputWithContext(ctx context.Context) VPCOutput
}

func (*VPC) ElementType() reflect.Type {
	return reflect.TypeOf((**VPC)(nil)).Elem()
}

func (i *VPC) ToVPCOutput() VPCOutput {
	return i.ToVPCOutputWithContext(context.Background())
}

func (i *VPC) ToVPCOutputWithContext(ctx context.Context) VPCOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VPCOutput)
}

type VPCOutput struct{ *pulumi.OutputState }

func (VPCOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VPC)(nil)).Elem()
}

func (o VPCOutput) ToVPCOutput() VPCOutput {
	return o
}

func (o VPCOutput) ToVPCOutputWithContext(ctx context.Context) VPCOutput {
	return o
}

// The primary IPv4 CIDR block for the VPC.
func (o VPCOutput) CidrBlock() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VPC) pulumi.StringPtrOutput { return v.CidrBlock }).(pulumi.StringPtrOutput)
}

// A list of IPv4 CIDR block association IDs for the VPC.
func (o VPCOutput) CidrBlockAssociations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VPC) pulumi.StringArrayOutput { return v.CidrBlockAssociations }).(pulumi.StringArrayOutput)
}

// The default network ACL ID that is associated with the VPC.
func (o VPCOutput) DefaultNetworkAcl() pulumi.StringOutput {
	return o.ApplyT(func(v *VPC) pulumi.StringOutput { return v.DefaultNetworkAcl }).(pulumi.StringOutput)
}

// The default security group ID that is associated with the VPC.
func (o VPCOutput) DefaultSecurityGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *VPC) pulumi.StringOutput { return v.DefaultSecurityGroup }).(pulumi.StringOutput)
}

// Indicates whether the instances launched in the VPC get DNS hostnames. If enabled, instances in the VPC get DNS hostnames; otherwise, they do not. Disabled by default for nondefault VPCs.
func (o VPCOutput) EnableDnsHostnames() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VPC) pulumi.BoolPtrOutput { return v.EnableDnsHostnames }).(pulumi.BoolPtrOutput)
}

// Indicates whether the DNS resolution is supported for the VPC. If enabled, queries to the Amazon provided DNS server at the 169.254.169.253 IP address, or the reserved IP address at the base of the VPC network range "plus two" succeed. If disabled, the Amazon provided DNS service in the VPC that resolves public DNS hostnames to IP addresses is not enabled. Enabled by default.
func (o VPCOutput) EnableDnsSupport() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VPC) pulumi.BoolPtrOutput { return v.EnableDnsSupport }).(pulumi.BoolPtrOutput)
}

// The allowed tenancy of instances launched into the VPC.
//
// "default": An instance launched into the VPC runs on shared hardware by default, unless you explicitly specify a different tenancy during instance launch.
//
// "dedicated": An instance launched into the VPC is a Dedicated Instance by default, unless you explicitly specify a tenancy of host during instance launch. You cannot specify a tenancy of default during instance launch.
//
// Updating InstanceTenancy requires no replacement only if you are updating its value from "dedicated" to "default". Updating InstanceTenancy from "default" to "dedicated" requires replacement.
func (o VPCOutput) InstanceTenancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VPC) pulumi.StringPtrOutput { return v.InstanceTenancy }).(pulumi.StringPtrOutput)
}

// The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR
func (o VPCOutput) Ipv4IpamPoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VPC) pulumi.StringPtrOutput { return v.Ipv4IpamPoolId }).(pulumi.StringPtrOutput)
}

// The netmask length of the IPv4 CIDR you want to allocate to this VPC from an Amazon VPC IP Address Manager (IPAM) pool
func (o VPCOutput) Ipv4NetmaskLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VPC) pulumi.IntPtrOutput { return v.Ipv4NetmaskLength }).(pulumi.IntPtrOutput)
}

// A list of IPv6 CIDR blocks that are associated with the VPC.
func (o VPCOutput) Ipv6CidrBlocks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VPC) pulumi.StringArrayOutput { return v.Ipv6CidrBlocks }).(pulumi.StringArrayOutput)
}

// The tags for the VPC.
func (o VPCOutput) Tags() VPCTagArrayOutput {
	return o.ApplyT(func(v *VPC) VPCTagArrayOutput { return v.Tags }).(VPCTagArrayOutput)
}

// The Id for the model.
func (o VPCOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *VPC) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VPCInput)(nil)).Elem(), &VPC{})
	pulumi.RegisterOutputType(VPCOutput{})
}
