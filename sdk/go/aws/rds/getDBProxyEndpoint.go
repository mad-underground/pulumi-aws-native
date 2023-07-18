// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::RDS::DBProxyEndpoint.
func LookupDBProxyEndpoint(ctx *pulumi.Context, args *LookupDBProxyEndpointArgs, opts ...pulumi.InvokeOption) (*LookupDBProxyEndpointResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDBProxyEndpointResult
	err := ctx.Invoke("aws-native:rds:getDBProxyEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDBProxyEndpointArgs struct {
	// The identifier for the DB proxy endpoint. This name must be unique for all DB proxy endpoints owned by your AWS account in the specified AWS Region.
	DBProxyEndpointName string `pulumi:"dBProxyEndpointName"`
}

type LookupDBProxyEndpointResult struct {
	// The Amazon Resource Name (ARN) for the DB proxy endpoint.
	DBProxyEndpointArn *string `pulumi:"dBProxyEndpointArn"`
	// The endpoint that you can use to connect to the DB proxy. You include the endpoint value in the connection string for a database client application.
	Endpoint *string `pulumi:"endpoint"`
	// A value that indicates whether this endpoint is the default endpoint for the associated DB proxy. Default DB proxy endpoints always have read/write capability. Other endpoints that you associate with the DB proxy can be either read/write or read-only.
	IsDefault *bool `pulumi:"isDefault"`
	// An optional set of key-value pairs to associate arbitrary data of your choosing with the DB proxy endpoint.
	Tags []DBProxyEndpointTagFormat `pulumi:"tags"`
	// A value that indicates whether the DB proxy endpoint can be used for read/write or read-only operations.
	TargetRole *DBProxyEndpointTargetRole `pulumi:"targetRole"`
	// VPC ID to associate with the new DB proxy endpoint.
	VpcId *string `pulumi:"vpcId"`
	// VPC security group IDs to associate with the new DB proxy endpoint.
	VpcSecurityGroupIds []string `pulumi:"vpcSecurityGroupIds"`
}

func LookupDBProxyEndpointOutput(ctx *pulumi.Context, args LookupDBProxyEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupDBProxyEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDBProxyEndpointResult, error) {
			args := v.(LookupDBProxyEndpointArgs)
			r, err := LookupDBProxyEndpoint(ctx, &args, opts...)
			var s LookupDBProxyEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDBProxyEndpointResultOutput)
}

type LookupDBProxyEndpointOutputArgs struct {
	// The identifier for the DB proxy endpoint. This name must be unique for all DB proxy endpoints owned by your AWS account in the specified AWS Region.
	DBProxyEndpointName pulumi.StringInput `pulumi:"dBProxyEndpointName"`
}

func (LookupDBProxyEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDBProxyEndpointArgs)(nil)).Elem()
}

type LookupDBProxyEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupDBProxyEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDBProxyEndpointResult)(nil)).Elem()
}

func (o LookupDBProxyEndpointResultOutput) ToLookupDBProxyEndpointResultOutput() LookupDBProxyEndpointResultOutput {
	return o
}

func (o LookupDBProxyEndpointResultOutput) ToLookupDBProxyEndpointResultOutputWithContext(ctx context.Context) LookupDBProxyEndpointResultOutput {
	return o
}

// The Amazon Resource Name (ARN) for the DB proxy endpoint.
func (o LookupDBProxyEndpointResultOutput) DBProxyEndpointArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBProxyEndpointResult) *string { return v.DBProxyEndpointArn }).(pulumi.StringPtrOutput)
}

// The endpoint that you can use to connect to the DB proxy. You include the endpoint value in the connection string for a database client application.
func (o LookupDBProxyEndpointResultOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBProxyEndpointResult) *string { return v.Endpoint }).(pulumi.StringPtrOutput)
}

// A value that indicates whether this endpoint is the default endpoint for the associated DB proxy. Default DB proxy endpoints always have read/write capability. Other endpoints that you associate with the DB proxy can be either read/write or read-only.
func (o LookupDBProxyEndpointResultOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDBProxyEndpointResult) *bool { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

// An optional set of key-value pairs to associate arbitrary data of your choosing with the DB proxy endpoint.
func (o LookupDBProxyEndpointResultOutput) Tags() DBProxyEndpointTagFormatArrayOutput {
	return o.ApplyT(func(v LookupDBProxyEndpointResult) []DBProxyEndpointTagFormat { return v.Tags }).(DBProxyEndpointTagFormatArrayOutput)
}

// A value that indicates whether the DB proxy endpoint can be used for read/write or read-only operations.
func (o LookupDBProxyEndpointResultOutput) TargetRole() DBProxyEndpointTargetRolePtrOutput {
	return o.ApplyT(func(v LookupDBProxyEndpointResult) *DBProxyEndpointTargetRole { return v.TargetRole }).(DBProxyEndpointTargetRolePtrOutput)
}

// VPC ID to associate with the new DB proxy endpoint.
func (o LookupDBProxyEndpointResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBProxyEndpointResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

// VPC security group IDs to associate with the new DB proxy endpoint.
func (o LookupDBProxyEndpointResultOutput) VpcSecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDBProxyEndpointResult) []string { return v.VpcSecurityGroupIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDBProxyEndpointResultOutput{})
}
