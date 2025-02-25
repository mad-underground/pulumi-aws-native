// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::CapacityReservationFleet
func LookupCapacityReservationFleet(ctx *pulumi.Context, args *LookupCapacityReservationFleetArgs, opts ...pulumi.InvokeOption) (*LookupCapacityReservationFleetResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCapacityReservationFleetResult
	err := ctx.Invoke("aws-native:ec2:getCapacityReservationFleet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCapacityReservationFleetArgs struct {
	CapacityReservationFleetId string `pulumi:"capacityReservationFleetId"`
}

type LookupCapacityReservationFleetResult struct {
	CapacityReservationFleetId *string `pulumi:"capacityReservationFleetId"`
	NoRemoveEndDate            *bool   `pulumi:"noRemoveEndDate"`
	RemoveEndDate              *bool   `pulumi:"removeEndDate"`
	TotalTargetCapacity        *int    `pulumi:"totalTargetCapacity"`
}

func LookupCapacityReservationFleetOutput(ctx *pulumi.Context, args LookupCapacityReservationFleetOutputArgs, opts ...pulumi.InvokeOption) LookupCapacityReservationFleetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCapacityReservationFleetResult, error) {
			args := v.(LookupCapacityReservationFleetArgs)
			r, err := LookupCapacityReservationFleet(ctx, &args, opts...)
			var s LookupCapacityReservationFleetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCapacityReservationFleetResultOutput)
}

type LookupCapacityReservationFleetOutputArgs struct {
	CapacityReservationFleetId pulumi.StringInput `pulumi:"capacityReservationFleetId"`
}

func (LookupCapacityReservationFleetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCapacityReservationFleetArgs)(nil)).Elem()
}

type LookupCapacityReservationFleetResultOutput struct{ *pulumi.OutputState }

func (LookupCapacityReservationFleetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCapacityReservationFleetResult)(nil)).Elem()
}

func (o LookupCapacityReservationFleetResultOutput) ToLookupCapacityReservationFleetResultOutput() LookupCapacityReservationFleetResultOutput {
	return o
}

func (o LookupCapacityReservationFleetResultOutput) ToLookupCapacityReservationFleetResultOutputWithContext(ctx context.Context) LookupCapacityReservationFleetResultOutput {
	return o
}

func (o LookupCapacityReservationFleetResultOutput) CapacityReservationFleetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCapacityReservationFleetResult) *string { return v.CapacityReservationFleetId }).(pulumi.StringPtrOutput)
}

func (o LookupCapacityReservationFleetResultOutput) NoRemoveEndDate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupCapacityReservationFleetResult) *bool { return v.NoRemoveEndDate }).(pulumi.BoolPtrOutput)
}

func (o LookupCapacityReservationFleetResultOutput) RemoveEndDate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupCapacityReservationFleetResult) *bool { return v.RemoveEndDate }).(pulumi.BoolPtrOutput)
}

func (o LookupCapacityReservationFleetResultOutput) TotalTargetCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupCapacityReservationFleetResult) *int { return v.TotalTargetCapacity }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCapacityReservationFleetResultOutput{})
}
