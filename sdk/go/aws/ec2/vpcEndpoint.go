// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::EC2::VPCEndpoint
type VpcEndpoint struct {
	pulumi.CustomResourceState

	CreationTimestamp   pulumi.StringOutput      `pulumi:"creationTimestamp"`
	DnsEntries          pulumi.StringArrayOutput `pulumi:"dnsEntries"`
	NetworkInterfaceIds pulumi.StringArrayOutput `pulumi:"networkInterfaceIds"`
	// A policy to attach to the endpoint that controls access to the service.
	PolicyDocument pulumi.StringPtrOutput `pulumi:"policyDocument"`
	// Indicate whether to associate a private hosted zone with the specified VPC.
	PrivateDnsEnabled pulumi.BoolPtrOutput `pulumi:"privateDnsEnabled"`
	// One or more route table IDs.
	RouteTableIds pulumi.StringArrayOutput `pulumi:"routeTableIds"`
	// The ID of one or more security groups to associate with the endpoint network interface.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// The service name.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// The ID of one or more subnets in which to create an endpoint network interface.
	SubnetIds       pulumi.StringArrayOutput `pulumi:"subnetIds"`
	VpcEndpointType VpcEndpointTypePtrOutput `pulumi:"vpcEndpointType"`
	// The ID of the VPC in which the endpoint will be used.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewVpcEndpoint registers a new resource with the given unique name, arguments, and options.
func NewVpcEndpoint(ctx *pulumi.Context,
	name string, args *VpcEndpointArgs, opts ...pulumi.ResourceOption) (*VpcEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"serviceName",
		"vpcEndpointType",
		"vpcId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcEndpoint
	err := ctx.RegisterResource("aws-native:ec2:VpcEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcEndpoint gets an existing VpcEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcEndpointState, opts ...pulumi.ResourceOption) (*VpcEndpoint, error) {
	var resource VpcEndpoint
	err := ctx.ReadResource("aws-native:ec2:VpcEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcEndpoint resources.
type vpcEndpointState struct {
}

type VpcEndpointState struct {
}

func (VpcEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointState)(nil)).Elem()
}

type vpcEndpointArgs struct {
	// A policy to attach to the endpoint that controls access to the service.
	PolicyDocument *string `pulumi:"policyDocument"`
	// Indicate whether to associate a private hosted zone with the specified VPC.
	PrivateDnsEnabled *bool `pulumi:"privateDnsEnabled"`
	// One or more route table IDs.
	RouteTableIds []string `pulumi:"routeTableIds"`
	// The ID of one or more security groups to associate with the endpoint network interface.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The service name.
	ServiceName string `pulumi:"serviceName"`
	// The ID of one or more subnets in which to create an endpoint network interface.
	SubnetIds       []string         `pulumi:"subnetIds"`
	VpcEndpointType *VpcEndpointType `pulumi:"vpcEndpointType"`
	// The ID of the VPC in which the endpoint will be used.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a VpcEndpoint resource.
type VpcEndpointArgs struct {
	// A policy to attach to the endpoint that controls access to the service.
	PolicyDocument pulumi.StringPtrInput
	// Indicate whether to associate a private hosted zone with the specified VPC.
	PrivateDnsEnabled pulumi.BoolPtrInput
	// One or more route table IDs.
	RouteTableIds pulumi.StringArrayInput
	// The ID of one or more security groups to associate with the endpoint network interface.
	SecurityGroupIds pulumi.StringArrayInput
	// The service name.
	ServiceName pulumi.StringInput
	// The ID of one or more subnets in which to create an endpoint network interface.
	SubnetIds       pulumi.StringArrayInput
	VpcEndpointType VpcEndpointTypePtrInput
	// The ID of the VPC in which the endpoint will be used.
	VpcId pulumi.StringInput
}

func (VpcEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointArgs)(nil)).Elem()
}

type VpcEndpointInput interface {
	pulumi.Input

	ToVpcEndpointOutput() VpcEndpointOutput
	ToVpcEndpointOutputWithContext(ctx context.Context) VpcEndpointOutput
}

func (*VpcEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpoint)(nil)).Elem()
}

func (i *VpcEndpoint) ToVpcEndpointOutput() VpcEndpointOutput {
	return i.ToVpcEndpointOutputWithContext(context.Background())
}

func (i *VpcEndpoint) ToVpcEndpointOutputWithContext(ctx context.Context) VpcEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointOutput)
}

func (i *VpcEndpoint) ToOutput(ctx context.Context) pulumix.Output[*VpcEndpoint] {
	return pulumix.Output[*VpcEndpoint]{
		OutputState: i.ToVpcEndpointOutputWithContext(ctx).OutputState,
	}
}

type VpcEndpointOutput struct{ *pulumi.OutputState }

func (VpcEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpoint)(nil)).Elem()
}

func (o VpcEndpointOutput) ToVpcEndpointOutput() VpcEndpointOutput {
	return o
}

func (o VpcEndpointOutput) ToVpcEndpointOutputWithContext(ctx context.Context) VpcEndpointOutput {
	return o
}

func (o VpcEndpointOutput) ToOutput(ctx context.Context) pulumix.Output[*VpcEndpoint] {
	return pulumix.Output[*VpcEndpoint]{
		OutputState: o.OutputState,
	}
}

func (o VpcEndpointOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

func (o VpcEndpointOutput) DnsEntries() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringArrayOutput { return v.DnsEntries }).(pulumi.StringArrayOutput)
}

func (o VpcEndpointOutput) NetworkInterfaceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringArrayOutput { return v.NetworkInterfaceIds }).(pulumi.StringArrayOutput)
}

// A policy to attach to the endpoint that controls access to the service.
func (o VpcEndpointOutput) PolicyDocument() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringPtrOutput { return v.PolicyDocument }).(pulumi.StringPtrOutput)
}

// Indicate whether to associate a private hosted zone with the specified VPC.
func (o VpcEndpointOutput) PrivateDnsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.BoolPtrOutput { return v.PrivateDnsEnabled }).(pulumi.BoolPtrOutput)
}

// One or more route table IDs.
func (o VpcEndpointOutput) RouteTableIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringArrayOutput { return v.RouteTableIds }).(pulumi.StringArrayOutput)
}

// The ID of one or more security groups to associate with the endpoint network interface.
func (o VpcEndpointOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// The service name.
func (o VpcEndpointOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// The ID of one or more subnets in which to create an endpoint network interface.
func (o VpcEndpointOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringArrayOutput { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

func (o VpcEndpointOutput) VpcEndpointType() VpcEndpointTypePtrOutput {
	return o.ApplyT(func(v *VpcEndpoint) VpcEndpointTypePtrOutput { return v.VpcEndpointType }).(VpcEndpointTypePtrOutput)
}

// The ID of the VPC in which the endpoint will be used.
func (o VpcEndpointOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointInput)(nil)).Elem(), &VpcEndpoint{})
	pulumi.RegisterOutputType(VpcEndpointOutput{})
}