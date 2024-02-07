// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Schema of AWS::EC2::IPAMAllocation Type
func LookupIpamAllocation(ctx *pulumi.Context, args *LookupIpamAllocationArgs, opts ...pulumi.InvokeOption) (*LookupIpamAllocationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIpamAllocationResult
	err := ctx.Invoke("aws-native:ec2:getIpamAllocation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIpamAllocationArgs struct {
	Cidr string `pulumi:"cidr"`
	// Id of the allocation.
	IpamPoolAllocationId string `pulumi:"ipamPoolAllocationId"`
	// Id of the IPAM Pool.
	IpamPoolId string `pulumi:"ipamPoolId"`
}

type LookupIpamAllocationResult struct {
	// Id of the allocation.
	IpamPoolAllocationId *string `pulumi:"ipamPoolAllocationId"`
}

func LookupIpamAllocationOutput(ctx *pulumi.Context, args LookupIpamAllocationOutputArgs, opts ...pulumi.InvokeOption) LookupIpamAllocationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIpamAllocationResult, error) {
			args := v.(LookupIpamAllocationArgs)
			r, err := LookupIpamAllocation(ctx, &args, opts...)
			var s LookupIpamAllocationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIpamAllocationResultOutput)
}

type LookupIpamAllocationOutputArgs struct {
	Cidr pulumi.StringInput `pulumi:"cidr"`
	// Id of the allocation.
	IpamPoolAllocationId pulumi.StringInput `pulumi:"ipamPoolAllocationId"`
	// Id of the IPAM Pool.
	IpamPoolId pulumi.StringInput `pulumi:"ipamPoolId"`
}

func (LookupIpamAllocationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpamAllocationArgs)(nil)).Elem()
}

type LookupIpamAllocationResultOutput struct{ *pulumi.OutputState }

func (LookupIpamAllocationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpamAllocationResult)(nil)).Elem()
}

func (o LookupIpamAllocationResultOutput) ToLookupIpamAllocationResultOutput() LookupIpamAllocationResultOutput {
	return o
}

func (o LookupIpamAllocationResultOutput) ToLookupIpamAllocationResultOutputWithContext(ctx context.Context) LookupIpamAllocationResultOutput {
	return o
}

// Id of the allocation.
func (o LookupIpamAllocationResultOutput) IpamPoolAllocationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpamAllocationResult) *string { return v.IpamPoolAllocationId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIpamAllocationResultOutput{})
}
