// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package globalaccelerator

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::GlobalAccelerator::Accelerator
func LookupAccelerator(ctx *pulumi.Context, args *LookupAcceleratorArgs, opts ...pulumi.InvokeOption) (*LookupAcceleratorResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAcceleratorResult
	err := ctx.Invoke("aws-native:globalaccelerator:getAccelerator", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAcceleratorArgs struct {
	// The Amazon Resource Name (ARN) of the accelerator.
	AcceleratorArn string `pulumi:"acceleratorArn"`
}

type LookupAcceleratorResult struct {
	// The Amazon Resource Name (ARN) of the accelerator.
	AcceleratorArn *string `pulumi:"acceleratorArn"`
	// The Domain Name System (DNS) name that Global Accelerator creates that points to your accelerator's static IPv4 addresses.
	DnsName *string `pulumi:"dnsName"`
	// The Domain Name System (DNS) name that Global Accelerator creates that points to your accelerator's static IPv4 and IPv6 addresses.
	DualStackDnsName *string `pulumi:"dualStackDnsName"`
	// Indicates whether an accelerator is enabled. The value is true or false.
	Enabled *bool `pulumi:"enabled"`
	// IP Address type.
	IpAddressType *AcceleratorIpAddressType `pulumi:"ipAddressType"`
	// The IP addresses from BYOIP Prefix pool.
	IpAddresses []string `pulumi:"ipAddresses"`
	// The IPv4 addresses assigned to the accelerator.
	Ipv4Addresses []string `pulumi:"ipv4Addresses"`
	// The IPv6 addresses assigned if the accelerator is dualstack
	Ipv6Addresses []string `pulumi:"ipv6Addresses"`
	// Name of accelerator.
	Name *string          `pulumi:"name"`
	Tags []AcceleratorTag `pulumi:"tags"`
}

func LookupAcceleratorOutput(ctx *pulumi.Context, args LookupAcceleratorOutputArgs, opts ...pulumi.InvokeOption) LookupAcceleratorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAcceleratorResult, error) {
			args := v.(LookupAcceleratorArgs)
			r, err := LookupAccelerator(ctx, &args, opts...)
			var s LookupAcceleratorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAcceleratorResultOutput)
}

type LookupAcceleratorOutputArgs struct {
	// The Amazon Resource Name (ARN) of the accelerator.
	AcceleratorArn pulumi.StringInput `pulumi:"acceleratorArn"`
}

func (LookupAcceleratorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAcceleratorArgs)(nil)).Elem()
}

type LookupAcceleratorResultOutput struct{ *pulumi.OutputState }

func (LookupAcceleratorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAcceleratorResult)(nil)).Elem()
}

func (o LookupAcceleratorResultOutput) ToLookupAcceleratorResultOutput() LookupAcceleratorResultOutput {
	return o
}

func (o LookupAcceleratorResultOutput) ToLookupAcceleratorResultOutputWithContext(ctx context.Context) LookupAcceleratorResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the accelerator.
func (o LookupAcceleratorResultOutput) AcceleratorArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAcceleratorResult) *string { return v.AcceleratorArn }).(pulumi.StringPtrOutput)
}

// The Domain Name System (DNS) name that Global Accelerator creates that points to your accelerator's static IPv4 addresses.
func (o LookupAcceleratorResultOutput) DnsName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAcceleratorResult) *string { return v.DnsName }).(pulumi.StringPtrOutput)
}

// The Domain Name System (DNS) name that Global Accelerator creates that points to your accelerator's static IPv4 and IPv6 addresses.
func (o LookupAcceleratorResultOutput) DualStackDnsName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAcceleratorResult) *string { return v.DualStackDnsName }).(pulumi.StringPtrOutput)
}

// Indicates whether an accelerator is enabled. The value is true or false.
func (o LookupAcceleratorResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAcceleratorResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// IP Address type.
func (o LookupAcceleratorResultOutput) IpAddressType() AcceleratorIpAddressTypePtrOutput {
	return o.ApplyT(func(v LookupAcceleratorResult) *AcceleratorIpAddressType { return v.IpAddressType }).(AcceleratorIpAddressTypePtrOutput)
}

// The IP addresses from BYOIP Prefix pool.
func (o LookupAcceleratorResultOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAcceleratorResult) []string { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

// The IPv4 addresses assigned to the accelerator.
func (o LookupAcceleratorResultOutput) Ipv4Addresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAcceleratorResult) []string { return v.Ipv4Addresses }).(pulumi.StringArrayOutput)
}

// The IPv6 addresses assigned if the accelerator is dualstack
func (o LookupAcceleratorResultOutput) Ipv6Addresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAcceleratorResult) []string { return v.Ipv6Addresses }).(pulumi.StringArrayOutput)
}

// Name of accelerator.
func (o LookupAcceleratorResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAcceleratorResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupAcceleratorResultOutput) Tags() AcceleratorTagArrayOutput {
	return o.ApplyT(func(v LookupAcceleratorResult) []AcceleratorTag { return v.Tags }).(AcceleratorTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAcceleratorResultOutput{})
}
