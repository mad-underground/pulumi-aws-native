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

// The “AWS::ApiGatewayV2::Model“ resource updates data model for a WebSocket API. For more information, see [Model Selection Expressions](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-model-selection-expressions) in the *API Gateway Developer Guide*.
type Model struct {
	pulumi.CustomResourceState

	// The API identifier.
	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// The content-type for the model, for example, "application/json".
	ContentType pulumi.StringPtrOutput `pulumi:"contentType"`
	// The description of the model.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	ModelId     pulumi.StringOutput    `pulumi:"modelId"`
	// The name of the model.
	Name pulumi.StringOutput `pulumi:"name"`
	// The schema for the model. For application/json models, this should be JSON schema draft 4 model.
	Schema pulumi.AnyOutput `pulumi:"schema"`
}

// NewModel registers a new resource with the given unique name, arguments, and options.
func NewModel(ctx *pulumi.Context,
	name string, args *ModelArgs, opts ...pulumi.ResourceOption) (*Model, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.Schema == nil {
		return nil, errors.New("invalid value for required argument 'Schema'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"apiId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Model
	err := ctx.RegisterResource("aws-native:apigatewayv2:Model", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetModel gets an existing Model resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetModel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModelState, opts ...pulumi.ResourceOption) (*Model, error) {
	var resource Model
	err := ctx.ReadResource("aws-native:apigatewayv2:Model", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Model resources.
type modelState struct {
}

type ModelState struct {
}

func (ModelState) ElementType() reflect.Type {
	return reflect.TypeOf((*modelState)(nil)).Elem()
}

type modelArgs struct {
	// The API identifier.
	ApiId string `pulumi:"apiId"`
	// The content-type for the model, for example, "application/json".
	ContentType *string `pulumi:"contentType"`
	// The description of the model.
	Description *string `pulumi:"description"`
	// The name of the model.
	Name *string `pulumi:"name"`
	// The schema for the model. For application/json models, this should be JSON schema draft 4 model.
	Schema interface{} `pulumi:"schema"`
}

// The set of arguments for constructing a Model resource.
type ModelArgs struct {
	// The API identifier.
	ApiId pulumi.StringInput
	// The content-type for the model, for example, "application/json".
	ContentType pulumi.StringPtrInput
	// The description of the model.
	Description pulumi.StringPtrInput
	// The name of the model.
	Name pulumi.StringPtrInput
	// The schema for the model. For application/json models, this should be JSON schema draft 4 model.
	Schema pulumi.Input
}

func (ModelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*modelArgs)(nil)).Elem()
}

type ModelInput interface {
	pulumi.Input

	ToModelOutput() ModelOutput
	ToModelOutputWithContext(ctx context.Context) ModelOutput
}

func (*Model) ElementType() reflect.Type {
	return reflect.TypeOf((**Model)(nil)).Elem()
}

func (i *Model) ToModelOutput() ModelOutput {
	return i.ToModelOutputWithContext(context.Background())
}

func (i *Model) ToModelOutputWithContext(ctx context.Context) ModelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelOutput)
}

type ModelOutput struct{ *pulumi.OutputState }

func (ModelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Model)(nil)).Elem()
}

func (o ModelOutput) ToModelOutput() ModelOutput {
	return o
}

func (o ModelOutput) ToModelOutputWithContext(ctx context.Context) ModelOutput {
	return o
}

// The API identifier.
func (o ModelOutput) ApiId() pulumi.StringOutput {
	return o.ApplyT(func(v *Model) pulumi.StringOutput { return v.ApiId }).(pulumi.StringOutput)
}

// The content-type for the model, for example, "application/json".
func (o ModelOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Model) pulumi.StringPtrOutput { return v.ContentType }).(pulumi.StringPtrOutput)
}

// The description of the model.
func (o ModelOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Model) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ModelOutput) ModelId() pulumi.StringOutput {
	return o.ApplyT(func(v *Model) pulumi.StringOutput { return v.ModelId }).(pulumi.StringOutput)
}

// The name of the model.
func (o ModelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Model) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The schema for the model. For application/json models, this should be JSON schema draft 4 model.
func (o ModelOutput) Schema() pulumi.AnyOutput {
	return o.ApplyT(func(v *Model) pulumi.AnyOutput { return v.Schema }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ModelInput)(nil)).Elem(), &Model{})
	pulumi.RegisterOutputType(ModelOutput{})
}
