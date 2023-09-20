// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::EC2::VPC
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
	// The Id for the model.
	VpcId string `pulumi:"vpcId"`
}

type LookupVpcResult struct {
	// A list of IPv4 CIDR block association IDs for the VPC.
	CidrBlockAssociations []string `pulumi:"cidrBlockAssociations"`
	// The default network ACL ID that is associated with the VPC.
	DefaultNetworkAcl *string `pulumi:"defaultNetworkAcl"`
	// The default security group ID that is associated with the VPC.
	DefaultSecurityGroup *string `pulumi:"defaultSecurityGroup"`
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
	// A list of IPv6 CIDR blocks that are associated with the VPC.
	Ipv6CidrBlocks []string `pulumi:"ipv6CidrBlocks"`
	// The tags for the VPC.
	Tags []VpcTag `pulumi:"tags"`
	// The Id for the model.
	VpcId *string `pulumi:"vpcId"`
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
	// The Id for the model.
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

func (o LookupVpcResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVpcResult] {
	return pulumix.Output[LookupVpcResult]{
		OutputState: o.OutputState,
	}
}

// A list of IPv4 CIDR block association IDs for the VPC.
func (o LookupVpcResultOutput) CidrBlockAssociations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcResult) []string { return v.CidrBlockAssociations }).(pulumi.StringArrayOutput)
}

// The default network ACL ID that is associated with the VPC.
func (o LookupVpcResultOutput) DefaultNetworkAcl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcResult) *string { return v.DefaultNetworkAcl }).(pulumi.StringPtrOutput)
}

// The default security group ID that is associated with the VPC.
func (o LookupVpcResultOutput) DefaultSecurityGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcResult) *string { return v.DefaultSecurityGroup }).(pulumi.StringPtrOutput)
}

// Indicates whether the instances launched in the VPC get DNS hostnames. If enabled, instances in the VPC get DNS hostnames; otherwise, they do not. Disabled by default for nondefault VPCs.
func (o LookupVpcResultOutput) EnableDnsHostnames() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVpcResult) *bool { return v.EnableDnsHostnames }).(pulumi.BoolPtrOutput)
}

// Indicates whether the DNS resolution is supported for the VPC. If enabled, queries to the Amazon provided DNS server at the 169.254.169.253 IP address, or the reserved IP address at the base of the VPC network range "plus two" succeed. If disabled, the Amazon provided DNS service in the VPC that resolves public DNS hostnames to IP addresses is not enabled. Enabled by default.
func (o LookupVpcResultOutput) EnableDnsSupport() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVpcResult) *bool { return v.EnableDnsSupport }).(pulumi.BoolPtrOutput)
}

// The allowed tenancy of instances launched into the VPC.
//
// "default": An instance launched into the VPC runs on shared hardware by default, unless you explicitly specify a different tenancy during instance launch.
//
// "dedicated": An instance launched into the VPC is a Dedicated Instance by default, unless you explicitly specify a tenancy of host during instance launch. You cannot specify a tenancy of default during instance launch.
//
// Updating InstanceTenancy requires no replacement only if you are updating its value from "dedicated" to "default". Updating InstanceTenancy from "default" to "dedicated" requires replacement.
func (o LookupVpcResultOutput) InstanceTenancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcResult) *string { return v.InstanceTenancy }).(pulumi.StringPtrOutput)
}

// A list of IPv6 CIDR blocks that are associated with the VPC.
func (o LookupVpcResultOutput) Ipv6CidrBlocks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcResult) []string { return v.Ipv6CidrBlocks }).(pulumi.StringArrayOutput)
}

// The tags for the VPC.
func (o LookupVpcResultOutput) Tags() VpcTagArrayOutput {
	return o.ApplyT(func(v LookupVpcResult) []VpcTag { return v.Tags }).(VpcTagArrayOutput)
}

// The Id for the model.
func (o LookupVpcResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcResultOutput{})
}