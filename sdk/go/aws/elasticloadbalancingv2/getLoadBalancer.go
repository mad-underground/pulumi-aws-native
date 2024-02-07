// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancingv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ElasticLoadBalancingV2::LoadBalancer
func LookupLoadBalancer(ctx *pulumi.Context, args *LookupLoadBalancerArgs, opts ...pulumi.InvokeOption) (*LookupLoadBalancerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLoadBalancerResult
	err := ctx.Invoke("aws-native:elasticloadbalancingv2:getLoadBalancer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLoadBalancerArgs struct {
	// The Amazon Resource Name (ARN) of the load balancer.
	LoadBalancerArn string `pulumi:"loadBalancerArn"`
}

type LookupLoadBalancerResult struct {
	// The ID of the Amazon Route 53 hosted zone associated with the load balancer.
	CanonicalHostedZoneId *string `pulumi:"canonicalHostedZoneId"`
	// The public DNS name of the load balancer.
	DnsName *string `pulumi:"dnsName"`
	// Indicates whether to evaluate inbound security group rules for traffic sent to a Network Load Balancer through PrivateLink
	EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic *string `pulumi:"enforceSecurityGroupInboundRulesOnPrivateLinkTraffic"`
	// The type of IP addresses used by the subnets for your load balancer. The possible values are ipv4 (for IPv4 addresses) and dualstack (for IPv4 and IPv6 addresses).
	IpAddressType *string `pulumi:"ipAddressType"`
	// The Amazon Resource Name (ARN) of the load balancer.
	LoadBalancerArn *string `pulumi:"loadBalancerArn"`
	// The load balancer attributes.
	LoadBalancerAttributes []LoadBalancerAttribute `pulumi:"loadBalancerAttributes"`
	// The full name of the load balancer.
	LoadBalancerFullName *string `pulumi:"loadBalancerFullName"`
	// The name of the load balancer.
	LoadBalancerName *string `pulumi:"loadBalancerName"`
	// The IDs of the security groups for the load balancer.
	SecurityGroups []string `pulumi:"securityGroups"`
	// The IDs of the public subnets. You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both.
	SubnetMappings []LoadBalancerSubnetMapping `pulumi:"subnetMappings"`
	// The IDs of the public subnets. You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both. To specify an Elastic IP address, specify subnet mappings instead of subnets.
	Subnets []string `pulumi:"subnets"`
	// The tags to assign to the load balancer.
	Tags []LoadBalancerTag `pulumi:"tags"`
}

func LookupLoadBalancerOutput(ctx *pulumi.Context, args LookupLoadBalancerOutputArgs, opts ...pulumi.InvokeOption) LookupLoadBalancerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLoadBalancerResult, error) {
			args := v.(LookupLoadBalancerArgs)
			r, err := LookupLoadBalancer(ctx, &args, opts...)
			var s LookupLoadBalancerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLoadBalancerResultOutput)
}

type LookupLoadBalancerOutputArgs struct {
	// The Amazon Resource Name (ARN) of the load balancer.
	LoadBalancerArn pulumi.StringInput `pulumi:"loadBalancerArn"`
}

func (LookupLoadBalancerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoadBalancerArgs)(nil)).Elem()
}

type LookupLoadBalancerResultOutput struct{ *pulumi.OutputState }

func (LookupLoadBalancerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoadBalancerResult)(nil)).Elem()
}

func (o LookupLoadBalancerResultOutput) ToLookupLoadBalancerResultOutput() LookupLoadBalancerResultOutput {
	return o
}

func (o LookupLoadBalancerResultOutput) ToLookupLoadBalancerResultOutputWithContext(ctx context.Context) LookupLoadBalancerResultOutput {
	return o
}

// The ID of the Amazon Route 53 hosted zone associated with the load balancer.
func (o LookupLoadBalancerResultOutput) CanonicalHostedZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.CanonicalHostedZoneId }).(pulumi.StringPtrOutput)
}

// The public DNS name of the load balancer.
func (o LookupLoadBalancerResultOutput) DnsName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.DnsName }).(pulumi.StringPtrOutput)
}

// Indicates whether to evaluate inbound security group rules for traffic sent to a Network Load Balancer through PrivateLink
func (o LookupLoadBalancerResultOutput) EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string {
		return v.EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic
	}).(pulumi.StringPtrOutput)
}

// The type of IP addresses used by the subnets for your load balancer. The possible values are ipv4 (for IPv4 addresses) and dualstack (for IPv4 and IPv6 addresses).
func (o LookupLoadBalancerResultOutput) IpAddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.IpAddressType }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the load balancer.
func (o LookupLoadBalancerResultOutput) LoadBalancerArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.LoadBalancerArn }).(pulumi.StringPtrOutput)
}

// The load balancer attributes.
func (o LookupLoadBalancerResultOutput) LoadBalancerAttributes() LoadBalancerAttributeArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []LoadBalancerAttribute { return v.LoadBalancerAttributes }).(LoadBalancerAttributeArrayOutput)
}

// The full name of the load balancer.
func (o LookupLoadBalancerResultOutput) LoadBalancerFullName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.LoadBalancerFullName }).(pulumi.StringPtrOutput)
}

// The name of the load balancer.
func (o LookupLoadBalancerResultOutput) LoadBalancerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.LoadBalancerName }).(pulumi.StringPtrOutput)
}

// The IDs of the security groups for the load balancer.
func (o LookupLoadBalancerResultOutput) SecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []string { return v.SecurityGroups }).(pulumi.StringArrayOutput)
}

// The IDs of the public subnets. You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both.
func (o LookupLoadBalancerResultOutput) SubnetMappings() LoadBalancerSubnetMappingArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []LoadBalancerSubnetMapping { return v.SubnetMappings }).(LoadBalancerSubnetMappingArrayOutput)
}

// The IDs of the public subnets. You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both. To specify an Elastic IP address, specify subnet mappings instead of subnets.
func (o LookupLoadBalancerResultOutput) Subnets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []string { return v.Subnets }).(pulumi.StringArrayOutput)
}

// The tags to assign to the load balancer.
func (o LookupLoadBalancerResultOutput) Tags() LoadBalancerTagArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []LoadBalancerTag { return v.Tags }).(LoadBalancerTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLoadBalancerResultOutput{})
}
