// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Schema of AWS::EC2::IPAM Type
type Ipam struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the IPAM.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The Id of the default association to the default resource discovery, created with this IPAM.
	DefaultResourceDiscoveryAssociationId pulumi.StringOutput `pulumi:"defaultResourceDiscoveryAssociationId"`
	// The Id of the default resource discovery, created with this IPAM.
	DefaultResourceDiscoveryId pulumi.StringOutput    `pulumi:"defaultResourceDiscoveryId"`
	Description                pulumi.StringPtrOutput `pulumi:"description"`
	// Id of the IPAM.
	IpamId pulumi.StringOutput `pulumi:"ipamId"`
	// The regions IPAM is enabled for. Allows pools to be created in these regions, as well as enabling monitoring
	OperatingRegions IpamOperatingRegionArrayOutput `pulumi:"operatingRegions"`
	// The Id of the default scope for publicly routable IP space, created with this IPAM.
	PrivateDefaultScopeId pulumi.StringOutput `pulumi:"privateDefaultScopeId"`
	// The Id of the default scope for publicly routable IP space, created with this IPAM.
	PublicDefaultScopeId pulumi.StringOutput `pulumi:"publicDefaultScopeId"`
	// The count of resource discoveries associated with this IPAM.
	ResourceDiscoveryAssociationCount pulumi.IntOutput `pulumi:"resourceDiscoveryAssociationCount"`
	// The number of scopes that currently exist in this IPAM.
	ScopeCount pulumi.IntOutput `pulumi:"scopeCount"`
	// An array of key-value pairs to apply to this resource.
	Tags IpamTagArrayOutput `pulumi:"tags"`
	// The tier of the IPAM.
	Tier IpamTierPtrOutput `pulumi:"tier"`
}

// NewIpam registers a new resource with the given unique name, arguments, and options.
func NewIpam(ctx *pulumi.Context,
	name string, args *IpamArgs, opts ...pulumi.ResourceOption) (*Ipam, error) {
	if args == nil {
		args = &IpamArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Ipam
	err := ctx.RegisterResource("aws-native:ec2:Ipam", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpam gets an existing Ipam resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpam(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpamState, opts ...pulumi.ResourceOption) (*Ipam, error) {
	var resource Ipam
	err := ctx.ReadResource("aws-native:ec2:Ipam", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ipam resources.
type ipamState struct {
}

type IpamState struct {
}

func (IpamState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipamState)(nil)).Elem()
}

type ipamArgs struct {
	Description *string `pulumi:"description"`
	// The regions IPAM is enabled for. Allows pools to be created in these regions, as well as enabling monitoring
	OperatingRegions []IpamOperatingRegion `pulumi:"operatingRegions"`
	// An array of key-value pairs to apply to this resource.
	Tags []IpamTag `pulumi:"tags"`
	// The tier of the IPAM.
	Tier *IpamTier `pulumi:"tier"`
}

// The set of arguments for constructing a Ipam resource.
type IpamArgs struct {
	Description pulumi.StringPtrInput
	// The regions IPAM is enabled for. Allows pools to be created in these regions, as well as enabling monitoring
	OperatingRegions IpamOperatingRegionArrayInput
	// An array of key-value pairs to apply to this resource.
	Tags IpamTagArrayInput
	// The tier of the IPAM.
	Tier IpamTierPtrInput
}

func (IpamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipamArgs)(nil)).Elem()
}

type IpamInput interface {
	pulumi.Input

	ToIpamOutput() IpamOutput
	ToIpamOutputWithContext(ctx context.Context) IpamOutput
}

func (*Ipam) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipam)(nil)).Elem()
}

func (i *Ipam) ToIpamOutput() IpamOutput {
	return i.ToIpamOutputWithContext(context.Background())
}

func (i *Ipam) ToIpamOutputWithContext(ctx context.Context) IpamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpamOutput)
}

type IpamOutput struct{ *pulumi.OutputState }

func (IpamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipam)(nil)).Elem()
}

func (o IpamOutput) ToIpamOutput() IpamOutput {
	return o
}

func (o IpamOutput) ToIpamOutputWithContext(ctx context.Context) IpamOutput {
	return o
}

// The Amazon Resource Name (ARN) of the IPAM.
func (o IpamOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipam) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The Id of the default association to the default resource discovery, created with this IPAM.
func (o IpamOutput) DefaultResourceDiscoveryAssociationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipam) pulumi.StringOutput { return v.DefaultResourceDiscoveryAssociationId }).(pulumi.StringOutput)
}

// The Id of the default resource discovery, created with this IPAM.
func (o IpamOutput) DefaultResourceDiscoveryId() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipam) pulumi.StringOutput { return v.DefaultResourceDiscoveryId }).(pulumi.StringOutput)
}

func (o IpamOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ipam) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Id of the IPAM.
func (o IpamOutput) IpamId() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipam) pulumi.StringOutput { return v.IpamId }).(pulumi.StringOutput)
}

// The regions IPAM is enabled for. Allows pools to be created in these regions, as well as enabling monitoring
func (o IpamOutput) OperatingRegions() IpamOperatingRegionArrayOutput {
	return o.ApplyT(func(v *Ipam) IpamOperatingRegionArrayOutput { return v.OperatingRegions }).(IpamOperatingRegionArrayOutput)
}

// The Id of the default scope for publicly routable IP space, created with this IPAM.
func (o IpamOutput) PrivateDefaultScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipam) pulumi.StringOutput { return v.PrivateDefaultScopeId }).(pulumi.StringOutput)
}

// The Id of the default scope for publicly routable IP space, created with this IPAM.
func (o IpamOutput) PublicDefaultScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipam) pulumi.StringOutput { return v.PublicDefaultScopeId }).(pulumi.StringOutput)
}

// The count of resource discoveries associated with this IPAM.
func (o IpamOutput) ResourceDiscoveryAssociationCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Ipam) pulumi.IntOutput { return v.ResourceDiscoveryAssociationCount }).(pulumi.IntOutput)
}

// The number of scopes that currently exist in this IPAM.
func (o IpamOutput) ScopeCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Ipam) pulumi.IntOutput { return v.ScopeCount }).(pulumi.IntOutput)
}

// An array of key-value pairs to apply to this resource.
func (o IpamOutput) Tags() IpamTagArrayOutput {
	return o.ApplyT(func(v *Ipam) IpamTagArrayOutput { return v.Tags }).(IpamTagArrayOutput)
}

// The tier of the IPAM.
func (o IpamOutput) Tier() IpamTierPtrOutput {
	return o.ApplyT(func(v *Ipam) IpamTierPtrOutput { return v.Tier }).(IpamTierPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpamInput)(nil)).Elem(), &Ipam{})
	pulumi.RegisterOutputType(IpamOutput{})
}
