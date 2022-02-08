// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::InternetGateway
func LookupInternetGateway(ctx *pulumi.Context, args *LookupInternetGatewayArgs, opts ...pulumi.InvokeOption) (*LookupInternetGatewayResult, error) {
	var rv LookupInternetGatewayResult
	err := ctx.Invoke("aws-native:ec2:getInternetGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInternetGatewayArgs struct {
	// ID of internet gateway.
	InternetGatewayId string `pulumi:"internetGatewayId"`
}

type LookupInternetGatewayResult struct {
	// ID of internet gateway.
	InternetGatewayId *string `pulumi:"internetGatewayId"`
	// Any tags to assign to the internet gateway.
	Tags []InternetGatewayTag `pulumi:"tags"`
}

func LookupInternetGatewayOutput(ctx *pulumi.Context, args LookupInternetGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupInternetGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInternetGatewayResult, error) {
			args := v.(LookupInternetGatewayArgs)
			r, err := LookupInternetGateway(ctx, &args, opts...)
			return *r, err
		}).(LookupInternetGatewayResultOutput)
}

type LookupInternetGatewayOutputArgs struct {
	// ID of internet gateway.
	InternetGatewayId pulumi.StringInput `pulumi:"internetGatewayId"`
}

func (LookupInternetGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInternetGatewayArgs)(nil)).Elem()
}

type LookupInternetGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupInternetGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInternetGatewayResult)(nil)).Elem()
}

func (o LookupInternetGatewayResultOutput) ToLookupInternetGatewayResultOutput() LookupInternetGatewayResultOutput {
	return o
}

func (o LookupInternetGatewayResultOutput) ToLookupInternetGatewayResultOutputWithContext(ctx context.Context) LookupInternetGatewayResultOutput {
	return o
}

// ID of internet gateway.
func (o LookupInternetGatewayResultOutput) InternetGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInternetGatewayResult) *string { return v.InternetGatewayId }).(pulumi.StringPtrOutput)
}

// Any tags to assign to the internet gateway.
func (o LookupInternetGatewayResultOutput) Tags() InternetGatewayTagArrayOutput {
	return o.ApplyT(func(v LookupInternetGatewayResult) []InternetGatewayTag { return v.Tags }).(InternetGatewayTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInternetGatewayResultOutput{})
}
