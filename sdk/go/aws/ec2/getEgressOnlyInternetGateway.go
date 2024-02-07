// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::EgressOnlyInternetGateway
func LookupEgressOnlyInternetGateway(ctx *pulumi.Context, args *LookupEgressOnlyInternetGatewayArgs, opts ...pulumi.InvokeOption) (*LookupEgressOnlyInternetGatewayResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEgressOnlyInternetGatewayResult
	err := ctx.Invoke("aws-native:ec2:getEgressOnlyInternetGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEgressOnlyInternetGatewayArgs struct {
	// Service Generated ID of the EgressOnlyInternetGateway
	Id string `pulumi:"id"`
}

type LookupEgressOnlyInternetGatewayResult struct {
	// Service Generated ID of the EgressOnlyInternetGateway
	Id *string `pulumi:"id"`
}

func LookupEgressOnlyInternetGatewayOutput(ctx *pulumi.Context, args LookupEgressOnlyInternetGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupEgressOnlyInternetGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEgressOnlyInternetGatewayResult, error) {
			args := v.(LookupEgressOnlyInternetGatewayArgs)
			r, err := LookupEgressOnlyInternetGateway(ctx, &args, opts...)
			var s LookupEgressOnlyInternetGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEgressOnlyInternetGatewayResultOutput)
}

type LookupEgressOnlyInternetGatewayOutputArgs struct {
	// Service Generated ID of the EgressOnlyInternetGateway
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupEgressOnlyInternetGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEgressOnlyInternetGatewayArgs)(nil)).Elem()
}

type LookupEgressOnlyInternetGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupEgressOnlyInternetGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEgressOnlyInternetGatewayResult)(nil)).Elem()
}

func (o LookupEgressOnlyInternetGatewayResultOutput) ToLookupEgressOnlyInternetGatewayResultOutput() LookupEgressOnlyInternetGatewayResultOutput {
	return o
}

func (o LookupEgressOnlyInternetGatewayResultOutput) ToLookupEgressOnlyInternetGatewayResultOutputWithContext(ctx context.Context) LookupEgressOnlyInternetGatewayResultOutput {
	return o
}

// Service Generated ID of the EgressOnlyInternetGateway
func (o LookupEgressOnlyInternetGatewayResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEgressOnlyInternetGatewayResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEgressOnlyInternetGatewayResultOutput{})
}
