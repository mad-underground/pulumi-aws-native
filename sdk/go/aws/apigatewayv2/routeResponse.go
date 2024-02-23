// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The “AWS::ApiGatewayV2::RouteResponse“ resource creates a route response for a WebSocket API. For more information, see [Set up Route Responses for a WebSocket API in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-route-response.html) in the *API Gateway Developer Guide*.
type RouteResponse struct {
	pulumi.CustomResourceState

	// The API identifier.
	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// The model selection expression for the route response. Supported only for WebSocket APIs.
	ModelSelectionExpression pulumi.StringPtrOutput `pulumi:"modelSelectionExpression"`
	// The response models for the route response.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::ApiGatewayV2::RouteResponse` for more information about the expected schema for this property.
	ResponseModels pulumi.AnyOutput `pulumi:"responseModels"`
	// The route response parameters.
	ResponseParameters RouteResponseParameterConstraintsMapOutput `pulumi:"responseParameters"`
	// The route ID.
	RouteId         pulumi.StringOutput `pulumi:"routeId"`
	RouteResponseId pulumi.StringOutput `pulumi:"routeResponseId"`
	// The route response key.
	RouteResponseKey pulumi.StringOutput `pulumi:"routeResponseKey"`
}

// NewRouteResponse registers a new resource with the given unique name, arguments, and options.
func NewRouteResponse(ctx *pulumi.Context,
	name string, args *RouteResponseArgs, opts ...pulumi.ResourceOption) (*RouteResponse, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.RouteId == nil {
		return nil, errors.New("invalid value for required argument 'RouteId'")
	}
	if args.RouteResponseKey == nil {
		return nil, errors.New("invalid value for required argument 'RouteResponseKey'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"apiId",
		"routeId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RouteResponse
	err := ctx.RegisterResource("aws-native:apigatewayv2:RouteResponse", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouteResponse gets an existing RouteResponse resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteResponse(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteResponseState, opts ...pulumi.ResourceOption) (*RouteResponse, error) {
	var resource RouteResponse
	err := ctx.ReadResource("aws-native:apigatewayv2:RouteResponse", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouteResponse resources.
type routeResponseState struct {
}

type RouteResponseState struct {
}

func (RouteResponseState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeResponseState)(nil)).Elem()
}

type routeResponseArgs struct {
	// The API identifier.
	ApiId string `pulumi:"apiId"`
	// The model selection expression for the route response. Supported only for WebSocket APIs.
	ModelSelectionExpression *string `pulumi:"modelSelectionExpression"`
	// The response models for the route response.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::ApiGatewayV2::RouteResponse` for more information about the expected schema for this property.
	ResponseModels interface{} `pulumi:"responseModels"`
	// The route response parameters.
	ResponseParameters map[string]RouteResponseParameterConstraints `pulumi:"responseParameters"`
	// The route ID.
	RouteId string `pulumi:"routeId"`
	// The route response key.
	RouteResponseKey string `pulumi:"routeResponseKey"`
}

// The set of arguments for constructing a RouteResponse resource.
type RouteResponseArgs struct {
	// The API identifier.
	ApiId pulumi.StringInput
	// The model selection expression for the route response. Supported only for WebSocket APIs.
	ModelSelectionExpression pulumi.StringPtrInput
	// The response models for the route response.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::ApiGatewayV2::RouteResponse` for more information about the expected schema for this property.
	ResponseModels pulumi.Input
	// The route response parameters.
	ResponseParameters RouteResponseParameterConstraintsMapInput
	// The route ID.
	RouteId pulumi.StringInput
	// The route response key.
	RouteResponseKey pulumi.StringInput
}

func (RouteResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeResponseArgs)(nil)).Elem()
}

type RouteResponseInput interface {
	pulumi.Input

	ToRouteResponseOutput() RouteResponseOutput
	ToRouteResponseOutputWithContext(ctx context.Context) RouteResponseOutput
}

func (*RouteResponse) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteResponse)(nil)).Elem()
}

func (i *RouteResponse) ToRouteResponseOutput() RouteResponseOutput {
	return i.ToRouteResponseOutputWithContext(context.Background())
}

func (i *RouteResponse) ToRouteResponseOutputWithContext(ctx context.Context) RouteResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteResponseOutput)
}

type RouteResponseOutput struct{ *pulumi.OutputState }

func (RouteResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteResponse)(nil)).Elem()
}

func (o RouteResponseOutput) ToRouteResponseOutput() RouteResponseOutput {
	return o
}

func (o RouteResponseOutput) ToRouteResponseOutputWithContext(ctx context.Context) RouteResponseOutput {
	return o
}

// The API identifier.
func (o RouteResponseOutput) ApiId() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteResponse) pulumi.StringOutput { return v.ApiId }).(pulumi.StringOutput)
}

// The model selection expression for the route response. Supported only for WebSocket APIs.
func (o RouteResponseOutput) ModelSelectionExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteResponse) pulumi.StringPtrOutput { return v.ModelSelectionExpression }).(pulumi.StringPtrOutput)
}

// The response models for the route response.
//
// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::ApiGatewayV2::RouteResponse` for more information about the expected schema for this property.
func (o RouteResponseOutput) ResponseModels() pulumi.AnyOutput {
	return o.ApplyT(func(v *RouteResponse) pulumi.AnyOutput { return v.ResponseModels }).(pulumi.AnyOutput)
}

// The route response parameters.
func (o RouteResponseOutput) ResponseParameters() RouteResponseParameterConstraintsMapOutput {
	return o.ApplyT(func(v *RouteResponse) RouteResponseParameterConstraintsMapOutput { return v.ResponseParameters }).(RouteResponseParameterConstraintsMapOutput)
}

// The route ID.
func (o RouteResponseOutput) RouteId() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteResponse) pulumi.StringOutput { return v.RouteId }).(pulumi.StringOutput)
}

func (o RouteResponseOutput) RouteResponseId() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteResponse) pulumi.StringOutput { return v.RouteResponseId }).(pulumi.StringOutput)
}

// The route response key.
func (o RouteResponseOutput) RouteResponseKey() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteResponse) pulumi.StringOutput { return v.RouteResponseKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouteResponseInput)(nil)).Elem(), &RouteResponse{})
	pulumi.RegisterOutputType(RouteResponseOutput{})
}
