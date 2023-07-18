// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::Route
func LookupRoute(ctx *pulumi.Context, args *LookupRouteArgs, opts ...pulumi.InvokeOption) (*LookupRouteResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRouteResult
	err := ctx.Invoke("aws-native:ec2:getRoute", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRouteArgs struct {
	Id string `pulumi:"id"`
}

type LookupRouteResult struct {
	CarrierGatewayId            *string `pulumi:"carrierGatewayId"`
	DestinationIpv6CidrBlock    *string `pulumi:"destinationIpv6CidrBlock"`
	EgressOnlyInternetGatewayId *string `pulumi:"egressOnlyInternetGatewayId"`
	GatewayId                   *string `pulumi:"gatewayId"`
	Id                          *string `pulumi:"id"`
	InstanceId                  *string `pulumi:"instanceId"`
	LocalGatewayId              *string `pulumi:"localGatewayId"`
	NatGatewayId                *string `pulumi:"natGatewayId"`
	NetworkInterfaceId          *string `pulumi:"networkInterfaceId"`
	TransitGatewayId            *string `pulumi:"transitGatewayId"`
	VpcEndpointId               *string `pulumi:"vpcEndpointId"`
	VpcPeeringConnectionId      *string `pulumi:"vpcPeeringConnectionId"`
}

func LookupRouteOutput(ctx *pulumi.Context, args LookupRouteOutputArgs, opts ...pulumi.InvokeOption) LookupRouteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouteResult, error) {
			args := v.(LookupRouteArgs)
			r, err := LookupRoute(ctx, &args, opts...)
			var s LookupRouteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouteResultOutput)
}

type LookupRouteOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupRouteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteArgs)(nil)).Elem()
}

type LookupRouteResultOutput struct{ *pulumi.OutputState }

func (LookupRouteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteResult)(nil)).Elem()
}

func (o LookupRouteResultOutput) ToLookupRouteResultOutput() LookupRouteResultOutput {
	return o
}

func (o LookupRouteResultOutput) ToLookupRouteResultOutputWithContext(ctx context.Context) LookupRouteResultOutput {
	return o
}

func (o LookupRouteResultOutput) CarrierGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.CarrierGatewayId }).(pulumi.StringPtrOutput)
}

func (o LookupRouteResultOutput) DestinationIpv6CidrBlock() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.DestinationIpv6CidrBlock }).(pulumi.StringPtrOutput)
}

func (o LookupRouteResultOutput) EgressOnlyInternetGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.EgressOnlyInternetGatewayId }).(pulumi.StringPtrOutput)
}

func (o LookupRouteResultOutput) GatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.GatewayId }).(pulumi.StringPtrOutput)
}

func (o LookupRouteResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupRouteResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o LookupRouteResultOutput) LocalGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.LocalGatewayId }).(pulumi.StringPtrOutput)
}

func (o LookupRouteResultOutput) NatGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.NatGatewayId }).(pulumi.StringPtrOutput)
}

func (o LookupRouteResultOutput) NetworkInterfaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.NetworkInterfaceId }).(pulumi.StringPtrOutput)
}

func (o LookupRouteResultOutput) TransitGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.TransitGatewayId }).(pulumi.StringPtrOutput)
}

func (o LookupRouteResultOutput) VpcEndpointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.VpcEndpointId }).(pulumi.StringPtrOutput)
}

func (o LookupRouteResultOutput) VpcPeeringConnectionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.VpcPeeringConnectionId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouteResultOutput{})
}
