// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::Host
func LookupHost(ctx *pulumi.Context, args *LookupHostArgs, opts ...pulumi.InvokeOption) (*LookupHostResult, error) {
	var rv LookupHostResult
	err := ctx.Invoke("aws-native:ec2:getHost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHostArgs struct {
	// Id of the host created.
	HostId string `pulumi:"hostId"`
}

type LookupHostResult struct {
	// Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID.
	AutoPlacement *string `pulumi:"autoPlacement"`
	// Id of the host created.
	HostId *string `pulumi:"hostId"`
	// Indicates whether to enable or disable host recovery for the Dedicated Host. Host recovery is disabled by default.
	HostRecovery *string `pulumi:"hostRecovery"`
}

func LookupHostOutput(ctx *pulumi.Context, args LookupHostOutputArgs, opts ...pulumi.InvokeOption) LookupHostResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHostResult, error) {
			args := v.(LookupHostArgs)
			r, err := LookupHost(ctx, &args, opts...)
			return *r, err
		}).(LookupHostResultOutput)
}

type LookupHostOutputArgs struct {
	// Id of the host created.
	HostId pulumi.StringInput `pulumi:"hostId"`
}

func (LookupHostOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHostArgs)(nil)).Elem()
}

type LookupHostResultOutput struct{ *pulumi.OutputState }

func (LookupHostResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHostResult)(nil)).Elem()
}

func (o LookupHostResultOutput) ToLookupHostResultOutput() LookupHostResultOutput {
	return o
}

func (o LookupHostResultOutput) ToLookupHostResultOutputWithContext(ctx context.Context) LookupHostResultOutput {
	return o
}

// Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID.
func (o LookupHostResultOutput) AutoPlacement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostResult) *string { return v.AutoPlacement }).(pulumi.StringPtrOutput)
}

// Id of the host created.
func (o LookupHostResultOutput) HostId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostResult) *string { return v.HostId }).(pulumi.StringPtrOutput)
}

// Indicates whether to enable or disable host recovery for the Dedicated Host. Host recovery is disabled by default.
func (o LookupHostResultOutput) HostRecovery() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostResult) *string { return v.HostRecovery }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHostResultOutput{})
}
