// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApiGateway::Method
func LookupMethod(ctx *pulumi.Context, args *LookupMethodArgs, opts ...pulumi.InvokeOption) (*LookupMethodResult, error) {
	var rv LookupMethodResult
	err := ctx.Invoke("aws-native:apigateway:getMethod", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMethodArgs struct {
	// The backend system that the method calls when it receives a request.
	HttpMethod string `pulumi:"httpMethod"`
	// The ID of an API Gateway resource.
	ResourceId string `pulumi:"resourceId"`
	// The ID of the RestApi resource in which API Gateway creates the method.
	RestApiId string `pulumi:"restApiId"`
}

type LookupMethodResult struct {
	// Indicates whether the method requires clients to submit a valid API key.
	ApiKeyRequired *bool `pulumi:"apiKeyRequired"`
	// A list of authorization scopes configured on the method.
	AuthorizationScopes []string `pulumi:"authorizationScopes"`
	// The method's authorization type.
	AuthorizationType *MethodAuthorizationType `pulumi:"authorizationType"`
	// The identifier of the authorizer to use on this method.
	AuthorizerId *string `pulumi:"authorizerId"`
	// The backend system that the method calls when it receives a request.
	Integration *MethodIntegration `pulumi:"integration"`
	// The responses that can be sent to the client who calls the method.
	MethodResponses []MethodResponse `pulumi:"methodResponses"`
	// A friendly operation name for the method.
	OperationName *string `pulumi:"operationName"`
	// The resources that are used for the request's content type. Specify request models as key-value pairs (string-to-string mapping), with a content type as the key and a Model resource name as the value.
	RequestModels interface{} `pulumi:"requestModels"`
	// The request parameters that API Gateway accepts. Specify request parameters as key-value pairs (string-to-Boolean mapping), with a source as the key and a Boolean as the value.
	RequestParameters interface{} `pulumi:"requestParameters"`
	// The ID of the associated request validator.
	RequestValidatorId *string `pulumi:"requestValidatorId"`
}

func LookupMethodOutput(ctx *pulumi.Context, args LookupMethodOutputArgs, opts ...pulumi.InvokeOption) LookupMethodResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMethodResult, error) {
			args := v.(LookupMethodArgs)
			r, err := LookupMethod(ctx, &args, opts...)
			return *r, err
		}).(LookupMethodResultOutput)
}

type LookupMethodOutputArgs struct {
	// The backend system that the method calls when it receives a request.
	HttpMethod pulumi.StringInput `pulumi:"httpMethod"`
	// The ID of an API Gateway resource.
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	// The ID of the RestApi resource in which API Gateway creates the method.
	RestApiId pulumi.StringInput `pulumi:"restApiId"`
}

func (LookupMethodOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMethodArgs)(nil)).Elem()
}

type LookupMethodResultOutput struct{ *pulumi.OutputState }

func (LookupMethodResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMethodResult)(nil)).Elem()
}

func (o LookupMethodResultOutput) ToLookupMethodResultOutput() LookupMethodResultOutput {
	return o
}

func (o LookupMethodResultOutput) ToLookupMethodResultOutputWithContext(ctx context.Context) LookupMethodResultOutput {
	return o
}

// Indicates whether the method requires clients to submit a valid API key.
func (o LookupMethodResultOutput) ApiKeyRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupMethodResult) *bool { return v.ApiKeyRequired }).(pulumi.BoolPtrOutput)
}

// A list of authorization scopes configured on the method.
func (o LookupMethodResultOutput) AuthorizationScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMethodResult) []string { return v.AuthorizationScopes }).(pulumi.StringArrayOutput)
}

// The method's authorization type.
func (o LookupMethodResultOutput) AuthorizationType() MethodAuthorizationTypePtrOutput {
	return o.ApplyT(func(v LookupMethodResult) *MethodAuthorizationType { return v.AuthorizationType }).(MethodAuthorizationTypePtrOutput)
}

// The identifier of the authorizer to use on this method.
func (o LookupMethodResultOutput) AuthorizerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMethodResult) *string { return v.AuthorizerId }).(pulumi.StringPtrOutput)
}

// The backend system that the method calls when it receives a request.
func (o LookupMethodResultOutput) Integration() MethodIntegrationPtrOutput {
	return o.ApplyT(func(v LookupMethodResult) *MethodIntegration { return v.Integration }).(MethodIntegrationPtrOutput)
}

// The responses that can be sent to the client who calls the method.
func (o LookupMethodResultOutput) MethodResponses() MethodResponseArrayOutput {
	return o.ApplyT(func(v LookupMethodResult) []MethodResponse { return v.MethodResponses }).(MethodResponseArrayOutput)
}

// A friendly operation name for the method.
func (o LookupMethodResultOutput) OperationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMethodResult) *string { return v.OperationName }).(pulumi.StringPtrOutput)
}

// The resources that are used for the request's content type. Specify request models as key-value pairs (string-to-string mapping), with a content type as the key and a Model resource name as the value.
func (o LookupMethodResultOutput) RequestModels() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupMethodResult) interface{} { return v.RequestModels }).(pulumi.AnyOutput)
}

// The request parameters that API Gateway accepts. Specify request parameters as key-value pairs (string-to-Boolean mapping), with a source as the key and a Boolean as the value.
func (o LookupMethodResultOutput) RequestParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupMethodResult) interface{} { return v.RequestParameters }).(pulumi.AnyOutput)
}

// The ID of the associated request validator.
func (o LookupMethodResultOutput) RequestValidatorId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMethodResult) *string { return v.RequestValidatorId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMethodResultOutput{})
}
