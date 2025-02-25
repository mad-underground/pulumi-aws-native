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

// Specifies a VPC endpoint. A VPC endpoint provides a private connection between your VPC and an endpoint service. You can use an endpoint service provided by AWS, an MKT Partner, or another AWS accounts in your organization. For more information, see the [User Guide](https://docs.aws.amazon.com/vpc/latest/privatelink/).
//
//	An endpoint of type ``Interface`` establishes connections between the subnets in your VPC and an AWS-service, your own service, or a service hosted by another AWS-account. With an interface VPC endpoint, you specify the subnets in which to create the endpoint and the security groups to associate with the endpoint network interfaces.
//	An endpoint of type ``gateway`` serves as a target for a route in your route table for traffic destined for S3 or DDB. You can specify an endpoint policy for the endpoint, which controls access to the service from your VPC. You can also specify the VPC route tables that use the endpoint. For more information about connectivity to S3, see [W
type VpcEndpoint struct {
	pulumi.CustomResourceState

	AwsId               pulumi.StringOutput      `pulumi:"awsId"`
	CreationTimestamp   pulumi.StringOutput      `pulumi:"creationTimestamp"`
	DnsEntries          pulumi.StringArrayOutput `pulumi:"dnsEntries"`
	NetworkInterfaceIds pulumi.StringArrayOutput `pulumi:"networkInterfaceIds"`
	// An endpoint policy, which controls access to the service from the VPC. The default endpoint policy allows full access to the service. Endpoint policies are supported only for gateway and interface endpoints.
	//  For CloudFormation templates in YAML, you can provide the policy in JSON or YAML format. CFNlong converts YAML policies to JSON format before calling the API to create or modify the VPC endpoint.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::EC2::VPCEndpoint` for more information about the expected schema for this property.
	PolicyDocument pulumi.AnyOutput `pulumi:"policyDocument"`
	// Indicate whether to associate a private hosted zone with the specified VPC. The private hosted zone contains a record set for the default public DNS name for the service for the Region (for example, ``kinesis.us-east-1.amazonaws.com``), which resolves to the private IP addresses of the endpoint network interfaces in the VPC. This enables you to make requests to the default public DNS name for the service instead of the public DNS names that are automatically generated by the VPC endpoint service.
	//  To use a private hosted zone, you must set the following VPC attributes to ``true``: ``enableDnsHostnames`` and ``enableDnsSupport``.
	//  This property is supported only for interface endpoints.
	//  Default: ``false``
	PrivateDnsEnabled pulumi.BoolPtrOutput `pulumi:"privateDnsEnabled"`
	// The IDs of the route tables. Routing is supported only for gateway endpoints.
	RouteTableIds pulumi.StringArrayOutput `pulumi:"routeTableIds"`
	// The IDs of the security groups to associate with the endpoint network interfaces. If this parameter is not specified, we use the default security group for the VPC. Security groups are supported only for interface endpoints.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// The name of the endpoint service.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// The IDs of the subnets in which to create endpoint network interfaces. You must specify this property for an interface endpoint or a Gateway Load Balancer endpoint. You can't specify this property for a gateway endpoint. For a Gateway Load Balancer endpoint, you can specify only one subnet.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// The type of endpoint.
	//  Default: Gateway
	VpcEndpointType VpcEndpointTypePtrOutput `pulumi:"vpcEndpointType"`
	// The ID of the VPC.
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
	// An endpoint policy, which controls access to the service from the VPC. The default endpoint policy allows full access to the service. Endpoint policies are supported only for gateway and interface endpoints.
	//  For CloudFormation templates in YAML, you can provide the policy in JSON or YAML format. CFNlong converts YAML policies to JSON format before calling the API to create or modify the VPC endpoint.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::EC2::VPCEndpoint` for more information about the expected schema for this property.
	PolicyDocument interface{} `pulumi:"policyDocument"`
	// Indicate whether to associate a private hosted zone with the specified VPC. The private hosted zone contains a record set for the default public DNS name for the service for the Region (for example, ``kinesis.us-east-1.amazonaws.com``), which resolves to the private IP addresses of the endpoint network interfaces in the VPC. This enables you to make requests to the default public DNS name for the service instead of the public DNS names that are automatically generated by the VPC endpoint service.
	//  To use a private hosted zone, you must set the following VPC attributes to ``true``: ``enableDnsHostnames`` and ``enableDnsSupport``.
	//  This property is supported only for interface endpoints.
	//  Default: ``false``
	PrivateDnsEnabled *bool `pulumi:"privateDnsEnabled"`
	// The IDs of the route tables. Routing is supported only for gateway endpoints.
	RouteTableIds []string `pulumi:"routeTableIds"`
	// The IDs of the security groups to associate with the endpoint network interfaces. If this parameter is not specified, we use the default security group for the VPC. Security groups are supported only for interface endpoints.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The name of the endpoint service.
	ServiceName string `pulumi:"serviceName"`
	// The IDs of the subnets in which to create endpoint network interfaces. You must specify this property for an interface endpoint or a Gateway Load Balancer endpoint. You can't specify this property for a gateway endpoint. For a Gateway Load Balancer endpoint, you can specify only one subnet.
	SubnetIds []string `pulumi:"subnetIds"`
	// The type of endpoint.
	//  Default: Gateway
	VpcEndpointType *VpcEndpointType `pulumi:"vpcEndpointType"`
	// The ID of the VPC.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a VpcEndpoint resource.
type VpcEndpointArgs struct {
	// An endpoint policy, which controls access to the service from the VPC. The default endpoint policy allows full access to the service. Endpoint policies are supported only for gateway and interface endpoints.
	//  For CloudFormation templates in YAML, you can provide the policy in JSON or YAML format. CFNlong converts YAML policies to JSON format before calling the API to create or modify the VPC endpoint.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::EC2::VPCEndpoint` for more information about the expected schema for this property.
	PolicyDocument pulumi.Input
	// Indicate whether to associate a private hosted zone with the specified VPC. The private hosted zone contains a record set for the default public DNS name for the service for the Region (for example, ``kinesis.us-east-1.amazonaws.com``), which resolves to the private IP addresses of the endpoint network interfaces in the VPC. This enables you to make requests to the default public DNS name for the service instead of the public DNS names that are automatically generated by the VPC endpoint service.
	//  To use a private hosted zone, you must set the following VPC attributes to ``true``: ``enableDnsHostnames`` and ``enableDnsSupport``.
	//  This property is supported only for interface endpoints.
	//  Default: ``false``
	PrivateDnsEnabled pulumi.BoolPtrInput
	// The IDs of the route tables. Routing is supported only for gateway endpoints.
	RouteTableIds pulumi.StringArrayInput
	// The IDs of the security groups to associate with the endpoint network interfaces. If this parameter is not specified, we use the default security group for the VPC. Security groups are supported only for interface endpoints.
	SecurityGroupIds pulumi.StringArrayInput
	// The name of the endpoint service.
	ServiceName pulumi.StringInput
	// The IDs of the subnets in which to create endpoint network interfaces. You must specify this property for an interface endpoint or a Gateway Load Balancer endpoint. You can't specify this property for a gateway endpoint. For a Gateway Load Balancer endpoint, you can specify only one subnet.
	SubnetIds pulumi.StringArrayInput
	// The type of endpoint.
	//  Default: Gateway
	VpcEndpointType VpcEndpointTypePtrInput
	// The ID of the VPC.
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

func (o VpcEndpointOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
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

// An endpoint policy, which controls access to the service from the VPC. The default endpoint policy allows full access to the service. Endpoint policies are supported only for gateway and interface endpoints.
//
//	For CloudFormation templates in YAML, you can provide the policy in JSON or YAML format. CFNlong converts YAML policies to JSON format before calling the API to create or modify the VPC endpoint.
//
// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::EC2::VPCEndpoint` for more information about the expected schema for this property.
func (o VpcEndpointOutput) PolicyDocument() pulumi.AnyOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.AnyOutput { return v.PolicyDocument }).(pulumi.AnyOutput)
}

// Indicate whether to associate a private hosted zone with the specified VPC. The private hosted zone contains a record set for the default public DNS name for the service for the Region (for example, “kinesis.us-east-1.amazonaws.com“), which resolves to the private IP addresses of the endpoint network interfaces in the VPC. This enables you to make requests to the default public DNS name for the service instead of the public DNS names that are automatically generated by the VPC endpoint service.
//
//	To use a private hosted zone, you must set the following VPC attributes to ``true``: ``enableDnsHostnames`` and ``enableDnsSupport``.
//	This property is supported only for interface endpoints.
//	Default: ``false``
func (o VpcEndpointOutput) PrivateDnsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.BoolPtrOutput { return v.PrivateDnsEnabled }).(pulumi.BoolPtrOutput)
}

// The IDs of the route tables. Routing is supported only for gateway endpoints.
func (o VpcEndpointOutput) RouteTableIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringArrayOutput { return v.RouteTableIds }).(pulumi.StringArrayOutput)
}

// The IDs of the security groups to associate with the endpoint network interfaces. If this parameter is not specified, we use the default security group for the VPC. Security groups are supported only for interface endpoints.
func (o VpcEndpointOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// The name of the endpoint service.
func (o VpcEndpointOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// The IDs of the subnets in which to create endpoint network interfaces. You must specify this property for an interface endpoint or a Gateway Load Balancer endpoint. You can't specify this property for a gateway endpoint. For a Gateway Load Balancer endpoint, you can specify only one subnet.
func (o VpcEndpointOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringArrayOutput { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

// The type of endpoint.
//
//	Default: Gateway
func (o VpcEndpointOutput) VpcEndpointType() VpcEndpointTypePtrOutput {
	return o.ApplyT(func(v *VpcEndpoint) VpcEndpointTypePtrOutput { return v.VpcEndpointType }).(VpcEndpointTypePtrOutput)
}

// The ID of the VPC.
func (o VpcEndpointOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointInput)(nil)).Elem(), &VpcEndpoint{})
	pulumi.RegisterOutputType(VpcEndpointOutput{})
}
