// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package quicksight

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the AWS::QuickSight::VPCConnection Resource Type.
type VPCConnection struct {
	pulumi.CustomResourceState

	Arn                pulumi.StringOutput                      `pulumi:"arn"`
	AvailabilityStatus VPCConnectionAvailabilityStatusPtrOutput `pulumi:"availabilityStatus"`
	AwsAccountId       pulumi.StringPtrOutput                   `pulumi:"awsAccountId"`
	CreatedTime        pulumi.StringOutput                      `pulumi:"createdTime"`
	DnsResolvers       pulumi.StringArrayOutput                 `pulumi:"dnsResolvers"`
	LastUpdatedTime    pulumi.StringOutput                      `pulumi:"lastUpdatedTime"`
	Name               pulumi.StringPtrOutput                   `pulumi:"name"`
	NetworkInterfaces  VPCConnectionNetworkInterfaceArrayOutput `pulumi:"networkInterfaces"`
	RoleArn            pulumi.StringPtrOutput                   `pulumi:"roleArn"`
	SecurityGroupIds   pulumi.StringArrayOutput                 `pulumi:"securityGroupIds"`
	Status             VPCConnectionResourceStatusOutput        `pulumi:"status"`
	SubnetIds          pulumi.StringArrayOutput                 `pulumi:"subnetIds"`
	Tags               VPCConnectionTagArrayOutput              `pulumi:"tags"`
	VPCConnectionId    pulumi.StringPtrOutput                   `pulumi:"vPCConnectionId"`
	VPCId              pulumi.StringOutput                      `pulumi:"vPCId"`
}

// NewVPCConnection registers a new resource with the given unique name, arguments, and options.
func NewVPCConnection(ctx *pulumi.Context,
	name string, args *VPCConnectionArgs, opts ...pulumi.ResourceOption) (*VPCConnection, error) {
	if args == nil {
		args = &VPCConnectionArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VPCConnection
	err := ctx.RegisterResource("aws-native:quicksight:VPCConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVPCConnection gets an existing VPCConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVPCConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VPCConnectionState, opts ...pulumi.ResourceOption) (*VPCConnection, error) {
	var resource VPCConnection
	err := ctx.ReadResource("aws-native:quicksight:VPCConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VPCConnection resources.
type vpcconnectionState struct {
}

type VPCConnectionState struct {
}

func (VPCConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcconnectionState)(nil)).Elem()
}

type vpcconnectionArgs struct {
	AvailabilityStatus *VPCConnectionAvailabilityStatus `pulumi:"availabilityStatus"`
	AwsAccountId       *string                          `pulumi:"awsAccountId"`
	DnsResolvers       []string                         `pulumi:"dnsResolvers"`
	Name               *string                          `pulumi:"name"`
	RoleArn            *string                          `pulumi:"roleArn"`
	SecurityGroupIds   []string                         `pulumi:"securityGroupIds"`
	SubnetIds          []string                         `pulumi:"subnetIds"`
	Tags               []VPCConnectionTag               `pulumi:"tags"`
	VPCConnectionId    *string                          `pulumi:"vPCConnectionId"`
}

// The set of arguments for constructing a VPCConnection resource.
type VPCConnectionArgs struct {
	AvailabilityStatus VPCConnectionAvailabilityStatusPtrInput
	AwsAccountId       pulumi.StringPtrInput
	DnsResolvers       pulumi.StringArrayInput
	Name               pulumi.StringPtrInput
	RoleArn            pulumi.StringPtrInput
	SecurityGroupIds   pulumi.StringArrayInput
	SubnetIds          pulumi.StringArrayInput
	Tags               VPCConnectionTagArrayInput
	VPCConnectionId    pulumi.StringPtrInput
}

func (VPCConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcconnectionArgs)(nil)).Elem()
}

type VPCConnectionInput interface {
	pulumi.Input

	ToVPCConnectionOutput() VPCConnectionOutput
	ToVPCConnectionOutputWithContext(ctx context.Context) VPCConnectionOutput
}

func (*VPCConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**VPCConnection)(nil)).Elem()
}

func (i *VPCConnection) ToVPCConnectionOutput() VPCConnectionOutput {
	return i.ToVPCConnectionOutputWithContext(context.Background())
}

func (i *VPCConnection) ToVPCConnectionOutputWithContext(ctx context.Context) VPCConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VPCConnectionOutput)
}

type VPCConnectionOutput struct{ *pulumi.OutputState }

func (VPCConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VPCConnection)(nil)).Elem()
}

func (o VPCConnectionOutput) ToVPCConnectionOutput() VPCConnectionOutput {
	return o
}

func (o VPCConnectionOutput) ToVPCConnectionOutputWithContext(ctx context.Context) VPCConnectionOutput {
	return o
}

func (o VPCConnectionOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *VPCConnection) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o VPCConnectionOutput) AvailabilityStatus() VPCConnectionAvailabilityStatusPtrOutput {
	return o.ApplyT(func(v *VPCConnection) VPCConnectionAvailabilityStatusPtrOutput { return v.AvailabilityStatus }).(VPCConnectionAvailabilityStatusPtrOutput)
}

func (o VPCConnectionOutput) AwsAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VPCConnection) pulumi.StringPtrOutput { return v.AwsAccountId }).(pulumi.StringPtrOutput)
}

func (o VPCConnectionOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VPCConnection) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o VPCConnectionOutput) DnsResolvers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VPCConnection) pulumi.StringArrayOutput { return v.DnsResolvers }).(pulumi.StringArrayOutput)
}

func (o VPCConnectionOutput) LastUpdatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VPCConnection) pulumi.StringOutput { return v.LastUpdatedTime }).(pulumi.StringOutput)
}

func (o VPCConnectionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VPCConnection) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VPCConnectionOutput) NetworkInterfaces() VPCConnectionNetworkInterfaceArrayOutput {
	return o.ApplyT(func(v *VPCConnection) VPCConnectionNetworkInterfaceArrayOutput { return v.NetworkInterfaces }).(VPCConnectionNetworkInterfaceArrayOutput)
}

func (o VPCConnectionOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VPCConnection) pulumi.StringPtrOutput { return v.RoleArn }).(pulumi.StringPtrOutput)
}

func (o VPCConnectionOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VPCConnection) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o VPCConnectionOutput) Status() VPCConnectionResourceStatusOutput {
	return o.ApplyT(func(v *VPCConnection) VPCConnectionResourceStatusOutput { return v.Status }).(VPCConnectionResourceStatusOutput)
}

func (o VPCConnectionOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VPCConnection) pulumi.StringArrayOutput { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

func (o VPCConnectionOutput) Tags() VPCConnectionTagArrayOutput {
	return o.ApplyT(func(v *VPCConnection) VPCConnectionTagArrayOutput { return v.Tags }).(VPCConnectionTagArrayOutput)
}

func (o VPCConnectionOutput) VPCConnectionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VPCConnection) pulumi.StringPtrOutput { return v.VPCConnectionId }).(pulumi.StringPtrOutput)
}

func (o VPCConnectionOutput) VPCId() pulumi.StringOutput {
	return o.ApplyT(func(v *VPCConnection) pulumi.StringOutput { return v.VPCId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VPCConnectionInput)(nil)).Elem(), &VPCConnection{})
	pulumi.RegisterOutputType(VPCConnectionOutput{})
}
