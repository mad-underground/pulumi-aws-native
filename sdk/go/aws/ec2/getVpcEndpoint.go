// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Specifies a VPC endpoint. A VPC endpoint provides a private connection between your VPC and an endpoint service. You can use an endpoint service provided by AWS, an MKT Partner, or another AWS accounts in your organization. For more information, see the [User Guide](https://docs.aws.amazon.com/vpc/latest/privatelink/).
//
//	An endpoint of type ``Interface`` establishes connections between the subnets in your VPC and an AWS-service, your own service, or a service hosted by another AWS-account. With an interface VPC endpoint, you specify the subnets in which to create the endpoint and the security groups to associate with the endpoint network interfaces.
//	An endpoint of type ``gateway`` serves as a target for a route in your route table for traffic destined for S3 or DDB. You can specify an endpoint policy for the endpoint, which controls access to the service from your VPC. You can also specify the VPC route tables that use the endpoint. For more information about connectivity to S3, see [W
func LookupVpcEndpoint(ctx *pulumi.Context, args *LookupVpcEndpointArgs, opts ...pulumi.InvokeOption) (*LookupVpcEndpointResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpcEndpointResult
	err := ctx.Invoke("aws-native:ec2:getVpcEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVpcEndpointArgs struct {
	Id string `pulumi:"id"`
}

type LookupVpcEndpointResult struct {
	CreationTimestamp   *string  `pulumi:"creationTimestamp"`
	DnsEntries          []string `pulumi:"dnsEntries"`
	Id                  *string  `pulumi:"id"`
	NetworkInterfaceIds []string `pulumi:"networkInterfaceIds"`
	// An endpoint policy, which controls access to the service from the VPC. The default endpoint policy allows full access to the service. Endpoint policies are supported only for gateway and interface endpoints.
	//  For CloudFormation templates in YAML, you can provide the policy in JSON or YAML format. CFNlong converts YAML policies to JSON format before calling the API to create or modify the VPC endpoint.
	PolicyDocument *string `pulumi:"policyDocument"`
	// Indicate whether to associate a private hosted zone with the specified VPC. The private hosted zone contains a record set for the default public DNS name for the service for the Region (for example, ``kinesis.us-east-1.amazonaws.com``), which resolves to the private IP addresses of the endpoint network interfaces in the VPC. This enables you to make requests to the default public DNS name for the service instead of the public DNS names that are automatically generated by the VPC endpoint service.
	//  To use a private hosted zone, you must set the following VPC attributes to ``true``: ``enableDnsHostnames`` and ``enableDnsSupport``.
	//  This property is supported only for interface endpoints.
	//  Default: ``false``
	PrivateDnsEnabled *bool `pulumi:"privateDnsEnabled"`
	// The IDs of the route tables. Routing is supported only for gateway endpoints.
	RouteTableIds []string `pulumi:"routeTableIds"`
	// The IDs of the security groups to associate with the endpoint network interfaces. If this parameter is not specified, we use the default security group for the VPC. Security groups are supported only for interface endpoints.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The IDs of the subnets in which to create endpoint network interfaces. You must specify this property for an interface endpoint or a Gateway Load Balancer endpoint. You can't specify this property for a gateway endpoint. For a Gateway Load Balancer endpoint, you can specify only one subnet.
	SubnetIds []string `pulumi:"subnetIds"`
}

func LookupVpcEndpointOutput(ctx *pulumi.Context, args LookupVpcEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupVpcEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpcEndpointResult, error) {
			args := v.(LookupVpcEndpointArgs)
			r, err := LookupVpcEndpoint(ctx, &args, opts...)
			var s LookupVpcEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVpcEndpointResultOutput)
}

type LookupVpcEndpointOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupVpcEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcEndpointArgs)(nil)).Elem()
}

type LookupVpcEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupVpcEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcEndpointResult)(nil)).Elem()
}

func (o LookupVpcEndpointResultOutput) ToLookupVpcEndpointResultOutput() LookupVpcEndpointResultOutput {
	return o
}

func (o LookupVpcEndpointResultOutput) ToLookupVpcEndpointResultOutputWithContext(ctx context.Context) LookupVpcEndpointResultOutput {
	return o
}

func (o LookupVpcEndpointResultOutput) CreationTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) *string { return v.CreationTimestamp }).(pulumi.StringPtrOutput)
}

func (o LookupVpcEndpointResultOutput) DnsEntries() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) []string { return v.DnsEntries }).(pulumi.StringArrayOutput)
}

func (o LookupVpcEndpointResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVpcEndpointResultOutput) NetworkInterfaceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) []string { return v.NetworkInterfaceIds }).(pulumi.StringArrayOutput)
}

// An endpoint policy, which controls access to the service from the VPC. The default endpoint policy allows full access to the service. Endpoint policies are supported only for gateway and interface endpoints.
//
//	For CloudFormation templates in YAML, you can provide the policy in JSON or YAML format. CFNlong converts YAML policies to JSON format before calling the API to create or modify the VPC endpoint.
func (o LookupVpcEndpointResultOutput) PolicyDocument() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) *string { return v.PolicyDocument }).(pulumi.StringPtrOutput)
}

// Indicate whether to associate a private hosted zone with the specified VPC. The private hosted zone contains a record set for the default public DNS name for the service for the Region (for example, “kinesis.us-east-1.amazonaws.com“), which resolves to the private IP addresses of the endpoint network interfaces in the VPC. This enables you to make requests to the default public DNS name for the service instead of the public DNS names that are automatically generated by the VPC endpoint service.
//
//	To use a private hosted zone, you must set the following VPC attributes to ``true``: ``enableDnsHostnames`` and ``enableDnsSupport``.
//	This property is supported only for interface endpoints.
//	Default: ``false``
func (o LookupVpcEndpointResultOutput) PrivateDnsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) *bool { return v.PrivateDnsEnabled }).(pulumi.BoolPtrOutput)
}

// The IDs of the route tables. Routing is supported only for gateway endpoints.
func (o LookupVpcEndpointResultOutput) RouteTableIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) []string { return v.RouteTableIds }).(pulumi.StringArrayOutput)
}

// The IDs of the security groups to associate with the endpoint network interfaces. If this parameter is not specified, we use the default security group for the VPC. Security groups are supported only for interface endpoints.
func (o LookupVpcEndpointResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// The IDs of the subnets in which to create endpoint network interfaces. You must specify this property for an interface endpoint or a Gateway Load Balancer endpoint. You can't specify this property for a gateway endpoint. For a Gateway Load Balancer endpoint, you can specify only one subnet.
func (o LookupVpcEndpointResultOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) []string { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcEndpointResultOutput{})
}
