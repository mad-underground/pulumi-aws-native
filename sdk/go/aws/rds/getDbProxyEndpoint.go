// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource schema for AWS::RDS::DBProxyEndpoint.
func LookupDbProxyEndpoint(ctx *pulumi.Context, args *LookupDbProxyEndpointArgs, opts ...pulumi.InvokeOption) (*LookupDbProxyEndpointResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDbProxyEndpointResult
	err := ctx.Invoke("aws-native:rds:getDbProxyEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDbProxyEndpointArgs struct {
	// The identifier for the DB proxy endpoint. This name must be unique for all DB proxy endpoints owned by your AWS account in the specified AWS Region.
	DbProxyEndpointName string `pulumi:"dbProxyEndpointName"`
}

type LookupDbProxyEndpointResult struct {
	// The Amazon Resource Name (ARN) for the DB proxy endpoint.
	DbProxyEndpointArn *string `pulumi:"dbProxyEndpointArn"`
	// The endpoint that you can use to connect to the DB proxy. You include the endpoint value in the connection string for a database client application.
	Endpoint *string `pulumi:"endpoint"`
	// A value that indicates whether this endpoint is the default endpoint for the associated DB proxy. Default DB proxy endpoints always have read/write capability. Other endpoints that you associate with the DB proxy can be either read/write or read-only.
	IsDefault *bool `pulumi:"isDefault"`
	// An optional set of key-value pairs to associate arbitrary data of your choosing with the DB proxy endpoint.
	Tags []DbProxyEndpointTagFormat `pulumi:"tags"`
	// A value that indicates whether the DB proxy endpoint can be used for read/write or read-only operations.
	TargetRole *DbProxyEndpointTargetRole `pulumi:"targetRole"`
	// VPC ID to associate with the new DB proxy endpoint.
	VpcId *string `pulumi:"vpcId"`
	// VPC security group IDs to associate with the new DB proxy endpoint.
	VpcSecurityGroupIds []string `pulumi:"vpcSecurityGroupIds"`
}

func LookupDbProxyEndpointOutput(ctx *pulumi.Context, args LookupDbProxyEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupDbProxyEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDbProxyEndpointResult, error) {
			args := v.(LookupDbProxyEndpointArgs)
			r, err := LookupDbProxyEndpoint(ctx, &args, opts...)
			var s LookupDbProxyEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDbProxyEndpointResultOutput)
}

type LookupDbProxyEndpointOutputArgs struct {
	// The identifier for the DB proxy endpoint. This name must be unique for all DB proxy endpoints owned by your AWS account in the specified AWS Region.
	DbProxyEndpointName pulumi.StringInput `pulumi:"dbProxyEndpointName"`
}

func (LookupDbProxyEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDbProxyEndpointArgs)(nil)).Elem()
}

type LookupDbProxyEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupDbProxyEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDbProxyEndpointResult)(nil)).Elem()
}

func (o LookupDbProxyEndpointResultOutput) ToLookupDbProxyEndpointResultOutput() LookupDbProxyEndpointResultOutput {
	return o
}

func (o LookupDbProxyEndpointResultOutput) ToLookupDbProxyEndpointResultOutputWithContext(ctx context.Context) LookupDbProxyEndpointResultOutput {
	return o
}

func (o LookupDbProxyEndpointResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDbProxyEndpointResult] {
	return pulumix.Output[LookupDbProxyEndpointResult]{
		OutputState: o.OutputState,
	}
}

// The Amazon Resource Name (ARN) for the DB proxy endpoint.
func (o LookupDbProxyEndpointResultOutput) DbProxyEndpointArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDbProxyEndpointResult) *string { return v.DbProxyEndpointArn }).(pulumi.StringPtrOutput)
}

// The endpoint that you can use to connect to the DB proxy. You include the endpoint value in the connection string for a database client application.
func (o LookupDbProxyEndpointResultOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDbProxyEndpointResult) *string { return v.Endpoint }).(pulumi.StringPtrOutput)
}

// A value that indicates whether this endpoint is the default endpoint for the associated DB proxy. Default DB proxy endpoints always have read/write capability. Other endpoints that you associate with the DB proxy can be either read/write or read-only.
func (o LookupDbProxyEndpointResultOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDbProxyEndpointResult) *bool { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

// An optional set of key-value pairs to associate arbitrary data of your choosing with the DB proxy endpoint.
func (o LookupDbProxyEndpointResultOutput) Tags() DbProxyEndpointTagFormatArrayOutput {
	return o.ApplyT(func(v LookupDbProxyEndpointResult) []DbProxyEndpointTagFormat { return v.Tags }).(DbProxyEndpointTagFormatArrayOutput)
}

// A value that indicates whether the DB proxy endpoint can be used for read/write or read-only operations.
func (o LookupDbProxyEndpointResultOutput) TargetRole() DbProxyEndpointTargetRolePtrOutput {
	return o.ApplyT(func(v LookupDbProxyEndpointResult) *DbProxyEndpointTargetRole { return v.TargetRole }).(DbProxyEndpointTargetRolePtrOutput)
}

// VPC ID to associate with the new DB proxy endpoint.
func (o LookupDbProxyEndpointResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDbProxyEndpointResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

// VPC security group IDs to associate with the new DB proxy endpoint.
func (o LookupDbProxyEndpointResultOutput) VpcSecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDbProxyEndpointResult) []string { return v.VpcSecurityGroupIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDbProxyEndpointResultOutput{})
}