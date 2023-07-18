// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApiGateway::GatewayResponse
//
// Deprecated: GatewayResponse is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type GatewayResponse struct {
	pulumi.CustomResourceState

	ResponseParameters pulumi.AnyOutput       `pulumi:"responseParameters"`
	ResponseTemplates  pulumi.AnyOutput       `pulumi:"responseTemplates"`
	ResponseType       pulumi.StringOutput    `pulumi:"responseType"`
	RestApiId          pulumi.StringOutput    `pulumi:"restApiId"`
	StatusCode         pulumi.StringPtrOutput `pulumi:"statusCode"`
}

// NewGatewayResponse registers a new resource with the given unique name, arguments, and options.
func NewGatewayResponse(ctx *pulumi.Context,
	name string, args *GatewayResponseArgs, opts ...pulumi.ResourceOption) (*GatewayResponse, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResponseType == nil {
		return nil, errors.New("invalid value for required argument 'ResponseType'")
	}
	if args.RestApiId == nil {
		return nil, errors.New("invalid value for required argument 'RestApiId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GatewayResponse
	err := ctx.RegisterResource("aws-native:apigateway:GatewayResponse", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayResponse gets an existing GatewayResponse resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayResponse(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayResponseState, opts ...pulumi.ResourceOption) (*GatewayResponse, error) {
	var resource GatewayResponse
	err := ctx.ReadResource("aws-native:apigateway:GatewayResponse", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayResponse resources.
type gatewayResponseState struct {
}

type GatewayResponseState struct {
}

func (GatewayResponseState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayResponseState)(nil)).Elem()
}

type gatewayResponseArgs struct {
	ResponseParameters interface{} `pulumi:"responseParameters"`
	ResponseTemplates  interface{} `pulumi:"responseTemplates"`
	ResponseType       string      `pulumi:"responseType"`
	RestApiId          string      `pulumi:"restApiId"`
	StatusCode         *string     `pulumi:"statusCode"`
}

// The set of arguments for constructing a GatewayResponse resource.
type GatewayResponseArgs struct {
	ResponseParameters pulumi.Input
	ResponseTemplates  pulumi.Input
	ResponseType       pulumi.StringInput
	RestApiId          pulumi.StringInput
	StatusCode         pulumi.StringPtrInput
}

func (GatewayResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayResponseArgs)(nil)).Elem()
}

type GatewayResponseInput interface {
	pulumi.Input

	ToGatewayResponseOutput() GatewayResponseOutput
	ToGatewayResponseOutputWithContext(ctx context.Context) GatewayResponseOutput
}

func (*GatewayResponse) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayResponse)(nil)).Elem()
}

func (i *GatewayResponse) ToGatewayResponseOutput() GatewayResponseOutput {
	return i.ToGatewayResponseOutputWithContext(context.Background())
}

func (i *GatewayResponse) ToGatewayResponseOutputWithContext(ctx context.Context) GatewayResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayResponseOutput)
}

type GatewayResponseOutput struct{ *pulumi.OutputState }

func (GatewayResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayResponse)(nil)).Elem()
}

func (o GatewayResponseOutput) ToGatewayResponseOutput() GatewayResponseOutput {
	return o
}

func (o GatewayResponseOutput) ToGatewayResponseOutputWithContext(ctx context.Context) GatewayResponseOutput {
	return o
}

func (o GatewayResponseOutput) ResponseParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *GatewayResponse) pulumi.AnyOutput { return v.ResponseParameters }).(pulumi.AnyOutput)
}

func (o GatewayResponseOutput) ResponseTemplates() pulumi.AnyOutput {
	return o.ApplyT(func(v *GatewayResponse) pulumi.AnyOutput { return v.ResponseTemplates }).(pulumi.AnyOutput)
}

func (o GatewayResponseOutput) ResponseType() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayResponse) pulumi.StringOutput { return v.ResponseType }).(pulumi.StringOutput)
}

func (o GatewayResponseOutput) RestApiId() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayResponse) pulumi.StringOutput { return v.RestApiId }).(pulumi.StringOutput)
}

func (o GatewayResponseOutput) StatusCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayResponse) pulumi.StringPtrOutput { return v.StatusCode }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayResponseInput)(nil)).Elem(), &GatewayResponse{})
	pulumi.RegisterOutputType(GatewayResponseOutput{})
}
