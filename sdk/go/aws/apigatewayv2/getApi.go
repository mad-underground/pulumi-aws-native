// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The “AWS::ApiGatewayV2::Api“ resource creates an API. WebSocket APIs and HTTP APIs are supported. For more information about WebSocket APIs, see [About WebSocket APIs in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-overview.html) in the *API Gateway Developer Guide*. For more information about HTTP APIs, see [HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api.html) in the *API Gateway Developer Guide.*
func LookupApi(ctx *pulumi.Context, args *LookupApiArgs, opts ...pulumi.InvokeOption) (*LookupApiResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupApiResult
	err := ctx.Invoke("aws-native:apigatewayv2:getApi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiArgs struct {
	ApiId string `pulumi:"apiId"`
}

type LookupApiResult struct {
	ApiEndpoint *string `pulumi:"apiEndpoint"`
	ApiId       *string `pulumi:"apiId"`
	// An API key selection expression. Supported only for WebSocket APIs. See [API Key Selection Expressions](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
	ApiKeySelectionExpression *string `pulumi:"apiKeySelectionExpression"`
	// A CORS configuration. Supported only for HTTP APIs. See [Configuring CORS](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html) for more information.
	CorsConfiguration *ApiCors `pulumi:"corsConfiguration"`
	// The description of the API.
	Description *string `pulumi:"description"`
	// Specifies whether clients can invoke your API by using the default ``execute-api`` endpoint. By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint.
	DisableExecuteApiEndpoint *bool `pulumi:"disableExecuteApiEndpoint"`
	// The name of the API. Required unless you specify an OpenAPI definition for ``Body`` or ``S3BodyLocation``.
	Name *string `pulumi:"name"`
	// The route selection expression for the API. For HTTP APIs, the ``routeSelectionExpression`` must be ``${request.method} ${request.path}``. If not provided, this will be the default for HTTP APIs. This property is required for WebSocket APIs.
	RouteSelectionExpression *string `pulumi:"routeSelectionExpression"`
	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]string `pulumi:"tags"`
	// A version identifier for the API.
	Version *string `pulumi:"version"`
}

func LookupApiOutput(ctx *pulumi.Context, args LookupApiOutputArgs, opts ...pulumi.InvokeOption) LookupApiResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiResult, error) {
			args := v.(LookupApiArgs)
			r, err := LookupApi(ctx, &args, opts...)
			var s LookupApiResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiResultOutput)
}

type LookupApiOutputArgs struct {
	ApiId pulumi.StringInput `pulumi:"apiId"`
}

func (LookupApiOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiArgs)(nil)).Elem()
}

type LookupApiResultOutput struct{ *pulumi.OutputState }

func (LookupApiResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiResult)(nil)).Elem()
}

func (o LookupApiResultOutput) ToLookupApiResultOutput() LookupApiResultOutput {
	return o
}

func (o LookupApiResultOutput) ToLookupApiResultOutputWithContext(ctx context.Context) LookupApiResultOutput {
	return o
}

func (o LookupApiResultOutput) ApiEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.ApiEndpoint }).(pulumi.StringPtrOutput)
}

func (o LookupApiResultOutput) ApiId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.ApiId }).(pulumi.StringPtrOutput)
}

// An API key selection expression. Supported only for WebSocket APIs. See [API Key Selection Expressions](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
func (o LookupApiResultOutput) ApiKeySelectionExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.ApiKeySelectionExpression }).(pulumi.StringPtrOutput)
}

// A CORS configuration. Supported only for HTTP APIs. See [Configuring CORS](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html) for more information.
func (o LookupApiResultOutput) CorsConfiguration() ApiCorsPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *ApiCors { return v.CorsConfiguration }).(ApiCorsPtrOutput)
}

// The description of the API.
func (o LookupApiResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies whether clients can invoke your API by using the default “execute-api“ endpoint. By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint.
func (o LookupApiResultOutput) DisableExecuteApiEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *bool { return v.DisableExecuteApiEndpoint }).(pulumi.BoolPtrOutput)
}

// The name of the API. Required unless you specify an OpenAPI definition for “Body“ or “S3BodyLocation“.
func (o LookupApiResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The route selection expression for the API. For HTTP APIs, the “routeSelectionExpression“ must be “${request.method} ${request.path}“. If not provided, this will be the default for HTTP APIs. This property is required for WebSocket APIs.
func (o LookupApiResultOutput) RouteSelectionExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.RouteSelectionExpression }).(pulumi.StringPtrOutput)
}

// The collection of tags. Each tag element is associated with a given resource.
func (o LookupApiResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApiResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// A version identifier for the API.
func (o LookupApiResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiResultOutput{})
}
