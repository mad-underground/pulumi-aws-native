// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package opensearchserverless

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Amazon OpenSearchServerless vpc endpoint resource
func LookupVpcEndpoint(ctx *pulumi.Context, args *LookupVpcEndpointArgs, opts ...pulumi.InvokeOption) (*LookupVpcEndpointResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpcEndpointResult
	err := ctx.Invoke("aws-native:opensearchserverless:getVpcEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVpcEndpointArgs struct {
	// The identifier of the VPC Endpoint
	Id string `pulumi:"id"`
}

type LookupVpcEndpointResult struct {
	// The identifier of the VPC Endpoint
	Id *string `pulumi:"id"`
	// The ID of one or more security groups to associate with the endpoint network interface
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The ID of one or more subnets in which to create an endpoint network interface
	SubnetIds []string `pulumi:"subnetIds"`
}

func LookupVpcEndpointOutput(ctx *pulumi.Context, args LookupVpcEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupVpcEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpcEndpointResult, error) {
			args := v.(LookupVpcEndpointArgs)
			r, err := LookupVpcEndpoint(ctx, &args, opts...)
			var s LookupVpcEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVpcEndpointResultOutput)
}

type LookupVpcEndpointOutputArgs struct {
	// The identifier of the VPC Endpoint
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupVpcEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcEndpointArgs)(nil)).Elem()
}

type LookupVpcEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupVpcEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcEndpointResult)(nil)).Elem()
}

func (o LookupVpcEndpointResultOutput) ToLookupVpcEndpointResultOutput() LookupVpcEndpointResultOutput {
	return o
}

func (o LookupVpcEndpointResultOutput) ToLookupVpcEndpointResultOutputWithContext(ctx context.Context) LookupVpcEndpointResultOutput {
	return o
}

// The identifier of the VPC Endpoint
func (o LookupVpcEndpointResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The ID of one or more security groups to associate with the endpoint network interface
func (o LookupVpcEndpointResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// The ID of one or more subnets in which to create an endpoint network interface
func (o LookupVpcEndpointResultOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) []string { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcEndpointResultOutput{})
}
