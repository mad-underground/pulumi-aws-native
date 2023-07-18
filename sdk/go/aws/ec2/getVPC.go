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
func LookupVPC(ctx *pulumi.Context, args *LookupVPCArgs, opts ...pulumi.InvokeOption) (*LookupVPCResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVPCResult
	err := ctx.Invoke("aws-native:ec2:getVPC", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVPCArgs struct {
	// The Id for the model.
	VpcId string `pulumi:"vpcId"`
}

type LookupVPCResult struct {
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
	Tags []VPCTag `pulumi:"tags"`
	// The Id for the model.
	VpcId *string `pulumi:"vpcId"`
}

func LookupVPCOutput(ctx *pulumi.Context, args LookupVPCOutputArgs, opts ...pulumi.InvokeOption) LookupVPCResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVPCResult, error) {
			args := v.(LookupVPCArgs)
			r, err := LookupVPC(ctx, &args, opts...)
			var s LookupVPCResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVPCResultOutput)
}

type LookupVPCOutputArgs struct {
	// The Id for the model.
	VpcId pulumi.StringInput `pulumi:"vpcId"`
}

func (LookupVPCOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVPCArgs)(nil)).Elem()
}

type LookupVPCResultOutput struct{ *pulumi.OutputState }

func (LookupVPCResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVPCResult)(nil)).Elem()
}

func (o LookupVPCResultOutput) ToLookupVPCResultOutput() LookupVPCResultOutput {
	return o
}

func (o LookupVPCResultOutput) ToLookupVPCResultOutputWithContext(ctx context.Context) LookupVPCResultOutput {
	return o
}

// A list of IPv4 CIDR block association IDs for the VPC.
func (o LookupVPCResultOutput) CidrBlockAssociations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVPCResult) []string { return v.CidrBlockAssociations }).(pulumi.StringArrayOutput)
}

// The default network ACL ID that is associated with the VPC.
func (o LookupVPCResultOutput) DefaultNetworkAcl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVPCResult) *string { return v.DefaultNetworkAcl }).(pulumi.StringPtrOutput)
}

// The default security group ID that is associated with the VPC.
func (o LookupVPCResultOutput) DefaultSecurityGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVPCResult) *string { return v.DefaultSecurityGroup }).(pulumi.StringPtrOutput)
}

// Indicates whether the instances launched in the VPC get DNS hostnames. If enabled, instances in the VPC get DNS hostnames; otherwise, they do not. Disabled by default for nondefault VPCs.
func (o LookupVPCResultOutput) EnableDnsHostnames() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVPCResult) *bool { return v.EnableDnsHostnames }).(pulumi.BoolPtrOutput)
}

// Indicates whether the DNS resolution is supported for the VPC. If enabled, queries to the Amazon provided DNS server at the 169.254.169.253 IP address, or the reserved IP address at the base of the VPC network range "plus two" succeed. If disabled, the Amazon provided DNS service in the VPC that resolves public DNS hostnames to IP addresses is not enabled. Enabled by default.
func (o LookupVPCResultOutput) EnableDnsSupport() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVPCResult) *bool { return v.EnableDnsSupport }).(pulumi.BoolPtrOutput)
}

// The allowed tenancy of instances launched into the VPC.
//
// "default": An instance launched into the VPC runs on shared hardware by default, unless you explicitly specify a different tenancy during instance launch.
//
// "dedicated": An instance launched into the VPC is a Dedicated Instance by default, unless you explicitly specify a tenancy of host during instance launch. You cannot specify a tenancy of default during instance launch.
//
// Updating InstanceTenancy requires no replacement only if you are updating its value from "dedicated" to "default". Updating InstanceTenancy from "default" to "dedicated" requires replacement.
func (o LookupVPCResultOutput) InstanceTenancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVPCResult) *string { return v.InstanceTenancy }).(pulumi.StringPtrOutput)
}

// A list of IPv6 CIDR blocks that are associated with the VPC.
func (o LookupVPCResultOutput) Ipv6CidrBlocks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVPCResult) []string { return v.Ipv6CidrBlocks }).(pulumi.StringArrayOutput)
}

// The tags for the VPC.
func (o LookupVPCResultOutput) Tags() VPCTagArrayOutput {
	return o.ApplyT(func(v LookupVPCResult) []VPCTag { return v.Tags }).(VPCTagArrayOutput)
}

// The Id for the model.
func (o LookupVPCResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVPCResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVPCResultOutput{})
}
