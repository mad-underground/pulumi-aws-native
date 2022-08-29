// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::CustomerGateway
func LookupCustomerGateway(ctx *pulumi.Context, args *LookupCustomerGatewayArgs, opts ...pulumi.InvokeOption) (*LookupCustomerGatewayResult, error) {
	var rv LookupCustomerGatewayResult
	err := ctx.Invoke("aws-native:ec2:getCustomerGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomerGatewayArgs struct {
	// CustomerGateway ID generated after customer gateway is created. Each customer gateway has a unique ID.
	CustomerGatewayId string `pulumi:"customerGatewayId"`
}

type LookupCustomerGatewayResult struct {
	// CustomerGateway ID generated after customer gateway is created. Each customer gateway has a unique ID.
	CustomerGatewayId *string `pulumi:"customerGatewayId"`
	// One or more tags for the customer gateway.
	Tags []CustomerGatewayTag `pulumi:"tags"`
}

func LookupCustomerGatewayOutput(ctx *pulumi.Context, args LookupCustomerGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupCustomerGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomerGatewayResult, error) {
			args := v.(LookupCustomerGatewayArgs)
			r, err := LookupCustomerGateway(ctx, &args, opts...)
			var s LookupCustomerGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCustomerGatewayResultOutput)
}

type LookupCustomerGatewayOutputArgs struct {
	// CustomerGateway ID generated after customer gateway is created. Each customer gateway has a unique ID.
	CustomerGatewayId pulumi.StringInput `pulumi:"customerGatewayId"`
}

func (LookupCustomerGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomerGatewayArgs)(nil)).Elem()
}

type LookupCustomerGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupCustomerGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomerGatewayResult)(nil)).Elem()
}

func (o LookupCustomerGatewayResultOutput) ToLookupCustomerGatewayResultOutput() LookupCustomerGatewayResultOutput {
	return o
}

func (o LookupCustomerGatewayResultOutput) ToLookupCustomerGatewayResultOutputWithContext(ctx context.Context) LookupCustomerGatewayResultOutput {
	return o
}

// CustomerGateway ID generated after customer gateway is created. Each customer gateway has a unique ID.
func (o LookupCustomerGatewayResultOutput) CustomerGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomerGatewayResult) *string { return v.CustomerGatewayId }).(pulumi.StringPtrOutput)
}

// One or more tags for the customer gateway.
func (o LookupCustomerGatewayResultOutput) Tags() CustomerGatewayTagArrayOutput {
	return o.ApplyT(func(v LookupCustomerGatewayResult) []CustomerGatewayTag { return v.Tags }).(CustomerGatewayTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomerGatewayResultOutput{})
}
