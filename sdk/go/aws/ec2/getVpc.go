// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Specifies a virtual private cloud (VPC).
//
//	You can optionally request an IPv6 CIDR block for the VPC. You can request an Amazon-provided IPv6 CIDR block from Amazon's pool of IPv6 addresses, or an IPv6 CIDR block from an IPv6 address pool that you provisioned through bring your own IP addresses (BYOIP).
//	For more information, see [Virtual private clouds (VPC)](https://docs.aws.amazon.com/vpc/latest/userguide/configure-your-vpc.html) in the *Amazon VPC User Guide*.
func LookupVpc(ctx *pulumi.Context, args *LookupVpcArgs, opts ...pulumi.InvokeOption) (*LookupVpcResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpcResult
	err := ctx.Invoke("aws-native:ec2:getVpc", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVpcArgs struct {
	VpcId string `pulumi:"vpcId"`
}

type LookupVpcResult struct {
	CidrBlockAssociations []string `pulumi:"cidrBlockAssociations"`
	DefaultNetworkAcl     *string  `pulumi:"defaultNetworkAcl"`
	DefaultSecurityGroup  *string  `pulumi:"defaultSecurityGroup"`
	// Indicates whether the instances launched in the VPC get DNS hostnames. If enabled, instances in the VPC get DNS hostnames; otherwise, they do not. Disabled by default for nondefault VPCs. For more information, see [DNS attributes in your VPC](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-dns.html#vpc-dns-support).
	//  You can only enable DNS hostnames if you've enabled DNS support.
	EnableDnsHostnames *bool `pulumi:"enableDnsHostnames"`
	// Indicates whether the DNS resolution is supported for the VPC. If enabled, queries to the Amazon provided DNS server at the 169.254.169.253 IP address, or the reserved IP address at the base of the VPC network range "plus two" succeed. If disabled, the Amazon provided DNS service in the VPC that resolves public DNS hostnames to IP addresses is not enabled. Enabled by default. For more information, see [DNS attributes in your VPC](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-dns.html#vpc-dns-support).
	EnableDnsSupport *bool `pulumi:"enableDnsSupport"`
	// The allowed tenancy of instances launched into the VPC.
	//   +  ``default``: An instance launched into the VPC runs on shared hardware by default, unless you explicitly specify a different tenancy during instance launch.
	//   +  ``dedicated``: An instance launched into the VPC runs on dedicated hardware by default, unless you explicitly specify a tenancy of ``host`` during instance launch. You cannot specify a tenancy of ``default`` during instance launch.
	//
	//  Updating ``InstanceTenancy`` requires no replacement only if you are updating its value from ``dedicated`` to ``default``. Updating ``InstanceTenancy`` from ``default`` to ``dedicated`` requires replacement.
	InstanceTenancy *string  `pulumi:"instanceTenancy"`
	Ipv6CidrBlocks  []string `pulumi:"ipv6CidrBlocks"`
	// The tags for the VPC.
	Tags  []aws.Tag `pulumi:"tags"`
	VpcId *string   `pulumi:"vpcId"`
}

func LookupVpcOutput(ctx *pulumi.Context, args LookupVpcOutputArgs, opts ...pulumi.InvokeOption) LookupVpcResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpcResult, error) {
			args := v.(LookupVpcArgs)
			r, err := LookupVpc(ctx, &args, opts...)
			var s LookupVpcResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVpcResultOutput)
}

type LookupVpcOutputArgs struct {
	VpcId pulumi.StringInput `pulumi:"vpcId"`
}

func (LookupVpcOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcArgs)(nil)).Elem()
}

type LookupVpcResultOutput struct{ *pulumi.OutputState }

func (LookupVpcResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcResult)(nil)).Elem()
}

func (o LookupVpcResultOutput) ToLookupVpcResultOutput() LookupVpcResultOutput {
	return o
}

func (o LookupVpcResultOutput) ToLookupVpcResultOutputWithContext(ctx context.Context) LookupVpcResultOutput {
	return o
}

func (o LookupVpcResultOutput) CidrBlockAssociations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcResult) []string { return v.CidrBlockAssociations }).(pulumi.StringArrayOutput)
}

func (o LookupVpcResultOutput) DefaultNetworkAcl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcResult) *string { return v.DefaultNetworkAcl }).(pulumi.StringPtrOutput)
}

func (o LookupVpcResultOutput) DefaultSecurityGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcResult) *string { return v.DefaultSecurityGroup }).(pulumi.StringPtrOutput)
}

// Indicates whether the instances launched in the VPC get DNS hostnames. If enabled, instances in the VPC get DNS hostnames; otherwise, they do not. Disabled by default for nondefault VPCs. For more information, see [DNS attributes in your VPC](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-dns.html#vpc-dns-support).
//
//	You can only enable DNS hostnames if you've enabled DNS support.
func (o LookupVpcResultOutput) EnableDnsHostnames() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVpcResult) *bool { return v.EnableDnsHostnames }).(pulumi.BoolPtrOutput)
}

// Indicates whether the DNS resolution is supported for the VPC. If enabled, queries to the Amazon provided DNS server at the 169.254.169.253 IP address, or the reserved IP address at the base of the VPC network range "plus two" succeed. If disabled, the Amazon provided DNS service in the VPC that resolves public DNS hostnames to IP addresses is not enabled. Enabled by default. For more information, see [DNS attributes in your VPC](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-dns.html#vpc-dns-support).
func (o LookupVpcResultOutput) EnableDnsSupport() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVpcResult) *bool { return v.EnableDnsSupport }).(pulumi.BoolPtrOutput)
}

// The allowed tenancy of instances launched into the VPC.
//
//   - “default“: An instance launched into the VPC runs on shared hardware by default, unless you explicitly specify a different tenancy during instance launch.
//
//   - “dedicated“: An instance launched into the VPC runs on dedicated hardware by default, unless you explicitly specify a tenancy of “host“ during instance launch. You cannot specify a tenancy of “default“ during instance launch.
//
//     Updating “InstanceTenancy“ requires no replacement only if you are updating its value from “dedicated“ to “default“. Updating “InstanceTenancy“ from “default“ to “dedicated“ requires replacement.
func (o LookupVpcResultOutput) InstanceTenancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcResult) *string { return v.InstanceTenancy }).(pulumi.StringPtrOutput)
}

func (o LookupVpcResultOutput) Ipv6CidrBlocks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcResult) []string { return v.Ipv6CidrBlocks }).(pulumi.StringArrayOutput)
}

// The tags for the VPC.
func (o LookupVpcResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupVpcResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func (o LookupVpcResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcResultOutput{})
}
