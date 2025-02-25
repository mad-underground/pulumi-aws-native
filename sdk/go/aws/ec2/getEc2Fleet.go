// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::EC2Fleet
func LookupEc2Fleet(ctx *pulumi.Context, args *LookupEc2FleetArgs, opts ...pulumi.InvokeOption) (*LookupEc2FleetResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEc2FleetResult
	err := ctx.Invoke("aws-native:ec2:getEc2Fleet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEc2FleetArgs struct {
	FleetId string `pulumi:"fleetId"`
}

type LookupEc2FleetResult struct {
	Context                         *string                                     `pulumi:"context"`
	ExcessCapacityTerminationPolicy *Ec2FleetExcessCapacityTerminationPolicy    `pulumi:"excessCapacityTerminationPolicy"`
	FleetId                         *string                                     `pulumi:"fleetId"`
	TargetCapacitySpecification     *Ec2FleetTargetCapacitySpecificationRequest `pulumi:"targetCapacitySpecification"`
}

func LookupEc2FleetOutput(ctx *pulumi.Context, args LookupEc2FleetOutputArgs, opts ...pulumi.InvokeOption) LookupEc2FleetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEc2FleetResult, error) {
			args := v.(LookupEc2FleetArgs)
			r, err := LookupEc2Fleet(ctx, &args, opts...)
			var s LookupEc2FleetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEc2FleetResultOutput)
}

type LookupEc2FleetOutputArgs struct {
	FleetId pulumi.StringInput `pulumi:"fleetId"`
}

func (LookupEc2FleetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEc2FleetArgs)(nil)).Elem()
}

type LookupEc2FleetResultOutput struct{ *pulumi.OutputState }

func (LookupEc2FleetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEc2FleetResult)(nil)).Elem()
}

func (o LookupEc2FleetResultOutput) ToLookupEc2FleetResultOutput() LookupEc2FleetResultOutput {
	return o
}

func (o LookupEc2FleetResultOutput) ToLookupEc2FleetResultOutputWithContext(ctx context.Context) LookupEc2FleetResultOutput {
	return o
}

func (o LookupEc2FleetResultOutput) Context() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEc2FleetResult) *string { return v.Context }).(pulumi.StringPtrOutput)
}

func (o LookupEc2FleetResultOutput) ExcessCapacityTerminationPolicy() Ec2FleetExcessCapacityTerminationPolicyPtrOutput {
	return o.ApplyT(func(v LookupEc2FleetResult) *Ec2FleetExcessCapacityTerminationPolicy {
		return v.ExcessCapacityTerminationPolicy
	}).(Ec2FleetExcessCapacityTerminationPolicyPtrOutput)
}

func (o LookupEc2FleetResultOutput) FleetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEc2FleetResult) *string { return v.FleetId }).(pulumi.StringPtrOutput)
}

func (o LookupEc2FleetResultOutput) TargetCapacitySpecification() Ec2FleetTargetCapacitySpecificationRequestPtrOutput {
	return o.ApplyT(func(v LookupEc2FleetResult) *Ec2FleetTargetCapacitySpecificationRequest {
		return v.TargetCapacitySpecification
	}).(Ec2FleetTargetCapacitySpecificationRequestPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEc2FleetResultOutput{})
}
