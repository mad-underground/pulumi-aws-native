// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::NatGateway
func LookupNatGateway(ctx *pulumi.Context, args *LookupNatGatewayArgs, opts ...pulumi.InvokeOption) (*LookupNatGatewayResult, error) {
	var rv LookupNatGatewayResult
	err := ctx.Invoke("aws-native:ec2:getNatGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNatGatewayArgs struct {
	NatGatewayId string `pulumi:"natGatewayId"`
}

type LookupNatGatewayResult struct {
	NatGatewayId *string         `pulumi:"natGatewayId"`
	Tags         []NatGatewayTag `pulumi:"tags"`
}

func LookupNatGatewayOutput(ctx *pulumi.Context, args LookupNatGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupNatGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNatGatewayResult, error) {
			args := v.(LookupNatGatewayArgs)
			r, err := LookupNatGateway(ctx, &args, opts...)
			var s LookupNatGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNatGatewayResultOutput)
}

type LookupNatGatewayOutputArgs struct {
	NatGatewayId pulumi.StringInput `pulumi:"natGatewayId"`
}

func (LookupNatGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNatGatewayArgs)(nil)).Elem()
}

type LookupNatGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupNatGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNatGatewayResult)(nil)).Elem()
}

func (o LookupNatGatewayResultOutput) ToLookupNatGatewayResultOutput() LookupNatGatewayResultOutput {
	return o
}

func (o LookupNatGatewayResultOutput) ToLookupNatGatewayResultOutputWithContext(ctx context.Context) LookupNatGatewayResultOutput {
	return o
}

func (o LookupNatGatewayResultOutput) NatGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) *string { return v.NatGatewayId }).(pulumi.StringPtrOutput)
}

func (o LookupNatGatewayResultOutput) Tags() NatGatewayTagArrayOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) []NatGatewayTag { return v.Tags }).(NatGatewayTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNatGatewayResultOutput{})
}
