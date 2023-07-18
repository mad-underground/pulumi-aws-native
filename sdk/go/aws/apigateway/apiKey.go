// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApiGateway::ApiKey
type ApiKey struct {
	pulumi.CustomResourceState

	// A Unique Key ID which identifies the API Key. Generated by the Create API and returned by the Read and List APIs
	APIKeyId pulumi.StringOutput `pulumi:"aPIKeyId"`
	// An AWS Marketplace customer identifier to use when integrating with the AWS SaaS Marketplace.
	CustomerId pulumi.StringPtrOutput `pulumi:"customerId"`
	// A description of the purpose of the API key.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Indicates whether the API key can be used by clients.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Specifies whether the key identifier is distinct from the created API key value. This parameter is deprecated and should not be used.
	GenerateDistinctId pulumi.BoolPtrOutput `pulumi:"generateDistinctId"`
	// A name for the API key. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the API key name.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// A list of stages to associate with this API key.
	StageKeys ApiKeyStageKeyArrayOutput `pulumi:"stageKeys"`
	// An array of arbitrary tags (key-value pairs) to associate with the API key.
	Tags ApiKeyTagArrayOutput `pulumi:"tags"`
	// The value of the API key. Must be at least 20 characters long.
	Value pulumi.StringPtrOutput `pulumi:"value"`
}

// NewApiKey registers a new resource with the given unique name, arguments, and options.
func NewApiKey(ctx *pulumi.Context,
	name string, args *ApiKeyArgs, opts ...pulumi.ResourceOption) (*ApiKey, error) {
	if args == nil {
		args = &ApiKeyArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApiKey
	err := ctx.RegisterResource("aws-native:apigateway:ApiKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiKey gets an existing ApiKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiKeyState, opts ...pulumi.ResourceOption) (*ApiKey, error) {
	var resource ApiKey
	err := ctx.ReadResource("aws-native:apigateway:ApiKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiKey resources.
type apiKeyState struct {
}

type ApiKeyState struct {
}

func (ApiKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiKeyState)(nil)).Elem()
}

type apiKeyArgs struct {
	// An AWS Marketplace customer identifier to use when integrating with the AWS SaaS Marketplace.
	CustomerId *string `pulumi:"customerId"`
	// A description of the purpose of the API key.
	Description *string `pulumi:"description"`
	// Indicates whether the API key can be used by clients.
	Enabled *bool `pulumi:"enabled"`
	// Specifies whether the key identifier is distinct from the created API key value. This parameter is deprecated and should not be used.
	GenerateDistinctId *bool `pulumi:"generateDistinctId"`
	// A name for the API key. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the API key name.
	Name *string `pulumi:"name"`
	// A list of stages to associate with this API key.
	StageKeys []ApiKeyStageKey `pulumi:"stageKeys"`
	// An array of arbitrary tags (key-value pairs) to associate with the API key.
	Tags []ApiKeyTag `pulumi:"tags"`
	// The value of the API key. Must be at least 20 characters long.
	Value *string `pulumi:"value"`
}

// The set of arguments for constructing a ApiKey resource.
type ApiKeyArgs struct {
	// An AWS Marketplace customer identifier to use when integrating with the AWS SaaS Marketplace.
	CustomerId pulumi.StringPtrInput
	// A description of the purpose of the API key.
	Description pulumi.StringPtrInput
	// Indicates whether the API key can be used by clients.
	Enabled pulumi.BoolPtrInput
	// Specifies whether the key identifier is distinct from the created API key value. This parameter is deprecated and should not be used.
	GenerateDistinctId pulumi.BoolPtrInput
	// A name for the API key. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the API key name.
	Name pulumi.StringPtrInput
	// A list of stages to associate with this API key.
	StageKeys ApiKeyStageKeyArrayInput
	// An array of arbitrary tags (key-value pairs) to associate with the API key.
	Tags ApiKeyTagArrayInput
	// The value of the API key. Must be at least 20 characters long.
	Value pulumi.StringPtrInput
}

func (ApiKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiKeyArgs)(nil)).Elem()
}

type ApiKeyInput interface {
	pulumi.Input

	ToApiKeyOutput() ApiKeyOutput
	ToApiKeyOutputWithContext(ctx context.Context) ApiKeyOutput
}

func (*ApiKey) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiKey)(nil)).Elem()
}

func (i *ApiKey) ToApiKeyOutput() ApiKeyOutput {
	return i.ToApiKeyOutputWithContext(context.Background())
}

func (i *ApiKey) ToApiKeyOutputWithContext(ctx context.Context) ApiKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiKeyOutput)
}

type ApiKeyOutput struct{ *pulumi.OutputState }

func (ApiKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiKey)(nil)).Elem()
}

func (o ApiKeyOutput) ToApiKeyOutput() ApiKeyOutput {
	return o
}

func (o ApiKeyOutput) ToApiKeyOutputWithContext(ctx context.Context) ApiKeyOutput {
	return o
}

// A Unique Key ID which identifies the API Key. Generated by the Create API and returned by the Read and List APIs
func (o ApiKeyOutput) APIKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiKey) pulumi.StringOutput { return v.APIKeyId }).(pulumi.StringOutput)
}

// An AWS Marketplace customer identifier to use when integrating with the AWS SaaS Marketplace.
func (o ApiKeyOutput) CustomerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiKey) pulumi.StringPtrOutput { return v.CustomerId }).(pulumi.StringPtrOutput)
}

// A description of the purpose of the API key.
func (o ApiKeyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiKey) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Indicates whether the API key can be used by clients.
func (o ApiKeyOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApiKey) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Specifies whether the key identifier is distinct from the created API key value. This parameter is deprecated and should not be used.
func (o ApiKeyOutput) GenerateDistinctId() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApiKey) pulumi.BoolPtrOutput { return v.GenerateDistinctId }).(pulumi.BoolPtrOutput)
}

// A name for the API key. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the API key name.
func (o ApiKeyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiKey) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// A list of stages to associate with this API key.
func (o ApiKeyOutput) StageKeys() ApiKeyStageKeyArrayOutput {
	return o.ApplyT(func(v *ApiKey) ApiKeyStageKeyArrayOutput { return v.StageKeys }).(ApiKeyStageKeyArrayOutput)
}

// An array of arbitrary tags (key-value pairs) to associate with the API key.
func (o ApiKeyOutput) Tags() ApiKeyTagArrayOutput {
	return o.ApplyT(func(v *ApiKey) ApiKeyTagArrayOutput { return v.Tags }).(ApiKeyTagArrayOutput)
}

// The value of the API key. Must be at least 20 characters long.
func (o ApiKeyOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiKey) pulumi.StringPtrOutput { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiKeyInput)(nil)).Elem(), &ApiKey{})
	pulumi.RegisterOutputType(ApiKeyOutput{})
}
