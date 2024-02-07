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

// The “AWS::ApiGateway::Model“ resource defines the structure of a request or response payload for an API method.
type Model struct {
	pulumi.CustomResourceState

	// The content-type for the model.
	ContentType pulumi.StringPtrOutput `pulumi:"contentType"`
	// The description of the model.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A name for the model. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the model name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
	//   If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The string identifier of the associated RestApi.
	RestApiId pulumi.StringOutput `pulumi:"restApiId"`
	// The schema for the model. For ``application/json`` models, this should be JSON schema draft 4 model. Do not include "\*/" characters in the description of any properties because such "\*/" characters may be interpreted as the closing marker for comments in some languages, such as Java or JavaScript, causing the installation of your API's SDK generated by API Gateway to fail.
	Schema pulumi.AnyOutput `pulumi:"schema"`
}

// NewModel registers a new resource with the given unique name, arguments, and options.
func NewModel(ctx *pulumi.Context,
	name string, args *ModelArgs, opts ...pulumi.ResourceOption) (*Model, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RestApiId == nil {
		return nil, errors.New("invalid value for required argument 'RestApiId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"contentType",
		"name",
		"restApiId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Model
	err := ctx.RegisterResource("aws-native:apigateway:Model", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws-native:apigateway:Model", name, id, state, &resource, opts...)
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
	// The content-type for the model.
	ContentType *string `pulumi:"contentType"`
	// The description of the model.
	Description *string `pulumi:"description"`
	// A name for the model. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the model name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
	//   If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	Name *string `pulumi:"name"`
	// The string identifier of the associated RestApi.
	RestApiId string `pulumi:"restApiId"`
	// The schema for the model. For ``application/json`` models, this should be JSON schema draft 4 model. Do not include "\*/" characters in the description of any properties because such "\*/" characters may be interpreted as the closing marker for comments in some languages, such as Java or JavaScript, causing the installation of your API's SDK generated by API Gateway to fail.
	Schema interface{} `pulumi:"schema"`
}

// The set of arguments for constructing a Model resource.
type ModelArgs struct {
	// The content-type for the model.
	ContentType pulumi.StringPtrInput
	// The description of the model.
	Description pulumi.StringPtrInput
	// A name for the model. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the model name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
	//   If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	Name pulumi.StringPtrInput
	// The string identifier of the associated RestApi.
	RestApiId pulumi.StringInput
	// The schema for the model. For ``application/json`` models, this should be JSON schema draft 4 model. Do not include "\*/" characters in the description of any properties because such "\*/" characters may be interpreted as the closing marker for comments in some languages, such as Java or JavaScript, causing the installation of your API's SDK generated by API Gateway to fail.
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

// The content-type for the model.
func (o ModelOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Model) pulumi.StringPtrOutput { return v.ContentType }).(pulumi.StringPtrOutput)
}

// The description of the model.
func (o ModelOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Model) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A name for the model. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the model name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
//
//	If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
func (o ModelOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Model) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The string identifier of the associated RestApi.
func (o ModelOutput) RestApiId() pulumi.StringOutput {
	return o.ApplyT(func(v *Model) pulumi.StringOutput { return v.RestApiId }).(pulumi.StringOutput)
}

// The schema for the model. For “application/json“ models, this should be JSON schema draft 4 model. Do not include "\*/" characters in the description of any properties because such "\*/" characters may be interpreted as the closing marker for comments in some languages, such as Java or JavaScript, causing the installation of your API's SDK generated by API Gateway to fail.
func (o ModelOutput) Schema() pulumi.AnyOutput {
	return o.ApplyT(func(v *Model) pulumi.AnyOutput { return v.Schema }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ModelInput)(nil)).Elem(), &Model{})
	pulumi.RegisterOutputType(ModelOutput{})
}
