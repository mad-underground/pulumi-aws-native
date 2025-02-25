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

// The “AWS::ApiGateway::RequestValidator“ resource sets up basic validation rules for incoming requests to your API. For more information, see [Enable Basic Request Validation for an API in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-method-request-validation.html) in the *API Gateway Developer Guide*.
type RequestValidator struct {
	pulumi.CustomResourceState

	// The name of this RequestValidator
	Name               pulumi.StringPtrOutput `pulumi:"name"`
	RequestValidatorId pulumi.StringOutput    `pulumi:"requestValidatorId"`
	// The string identifier of the associated RestApi.
	RestApiId pulumi.StringOutput `pulumi:"restApiId"`
	// A Boolean flag to indicate whether to validate a request body according to the configured Model schema.
	ValidateRequestBody pulumi.BoolPtrOutput `pulumi:"validateRequestBody"`
	// A Boolean flag to indicate whether to validate request parameters (``true``) or not (``false``).
	ValidateRequestParameters pulumi.BoolPtrOutput `pulumi:"validateRequestParameters"`
}

// NewRequestValidator registers a new resource with the given unique name, arguments, and options.
func NewRequestValidator(ctx *pulumi.Context,
	name string, args *RequestValidatorArgs, opts ...pulumi.ResourceOption) (*RequestValidator, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RestApiId == nil {
		return nil, errors.New("invalid value for required argument 'RestApiId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"name",
		"restApiId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RequestValidator
	err := ctx.RegisterResource("aws-native:apigateway:RequestValidator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRequestValidator gets an existing RequestValidator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRequestValidator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RequestValidatorState, opts ...pulumi.ResourceOption) (*RequestValidator, error) {
	var resource RequestValidator
	err := ctx.ReadResource("aws-native:apigateway:RequestValidator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RequestValidator resources.
type requestValidatorState struct {
}

type RequestValidatorState struct {
}

func (RequestValidatorState) ElementType() reflect.Type {
	return reflect.TypeOf((*requestValidatorState)(nil)).Elem()
}

type requestValidatorArgs struct {
	// The name of this RequestValidator
	Name *string `pulumi:"name"`
	// The string identifier of the associated RestApi.
	RestApiId string `pulumi:"restApiId"`
	// A Boolean flag to indicate whether to validate a request body according to the configured Model schema.
	ValidateRequestBody *bool `pulumi:"validateRequestBody"`
	// A Boolean flag to indicate whether to validate request parameters (``true``) or not (``false``).
	ValidateRequestParameters *bool `pulumi:"validateRequestParameters"`
}

// The set of arguments for constructing a RequestValidator resource.
type RequestValidatorArgs struct {
	// The name of this RequestValidator
	Name pulumi.StringPtrInput
	// The string identifier of the associated RestApi.
	RestApiId pulumi.StringInput
	// A Boolean flag to indicate whether to validate a request body according to the configured Model schema.
	ValidateRequestBody pulumi.BoolPtrInput
	// A Boolean flag to indicate whether to validate request parameters (``true``) or not (``false``).
	ValidateRequestParameters pulumi.BoolPtrInput
}

func (RequestValidatorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*requestValidatorArgs)(nil)).Elem()
}

type RequestValidatorInput interface {
	pulumi.Input

	ToRequestValidatorOutput() RequestValidatorOutput
	ToRequestValidatorOutputWithContext(ctx context.Context) RequestValidatorOutput
}

func (*RequestValidator) ElementType() reflect.Type {
	return reflect.TypeOf((**RequestValidator)(nil)).Elem()
}

func (i *RequestValidator) ToRequestValidatorOutput() RequestValidatorOutput {
	return i.ToRequestValidatorOutputWithContext(context.Background())
}

func (i *RequestValidator) ToRequestValidatorOutputWithContext(ctx context.Context) RequestValidatorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestValidatorOutput)
}

type RequestValidatorOutput struct{ *pulumi.OutputState }

func (RequestValidatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RequestValidator)(nil)).Elem()
}

func (o RequestValidatorOutput) ToRequestValidatorOutput() RequestValidatorOutput {
	return o
}

func (o RequestValidatorOutput) ToRequestValidatorOutputWithContext(ctx context.Context) RequestValidatorOutput {
	return o
}

// The name of this RequestValidator
func (o RequestValidatorOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestValidator) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RequestValidatorOutput) RequestValidatorId() pulumi.StringOutput {
	return o.ApplyT(func(v *RequestValidator) pulumi.StringOutput { return v.RequestValidatorId }).(pulumi.StringOutput)
}

// The string identifier of the associated RestApi.
func (o RequestValidatorOutput) RestApiId() pulumi.StringOutput {
	return o.ApplyT(func(v *RequestValidator) pulumi.StringOutput { return v.RestApiId }).(pulumi.StringOutput)
}

// A Boolean flag to indicate whether to validate a request body according to the configured Model schema.
func (o RequestValidatorOutput) ValidateRequestBody() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RequestValidator) pulumi.BoolPtrOutput { return v.ValidateRequestBody }).(pulumi.BoolPtrOutput)
}

// A Boolean flag to indicate whether to validate request parameters (“true“) or not (“false“).
func (o RequestValidatorOutput) ValidateRequestParameters() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RequestValidator) pulumi.BoolPtrOutput { return v.ValidateRequestParameters }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RequestValidatorInput)(nil)).Elem(), &RequestValidator{})
	pulumi.RegisterOutputType(RequestValidatorOutput{})
}
