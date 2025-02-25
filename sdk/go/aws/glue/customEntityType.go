// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Glue::CustomEntityType
//
// Deprecated: CustomEntityType is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type CustomEntityType struct {
	pulumi.CustomResourceState

	AwsId        pulumi.StringOutput      `pulumi:"awsId"`
	ContextWords pulumi.StringArrayOutput `pulumi:"contextWords"`
	Name         pulumi.StringPtrOutput   `pulumi:"name"`
	RegexString  pulumi.StringPtrOutput   `pulumi:"regexString"`
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Glue::CustomEntityType` for more information about the expected schema for this property.
	Tags pulumi.AnyOutput `pulumi:"tags"`
}

// NewCustomEntityType registers a new resource with the given unique name, arguments, and options.
func NewCustomEntityType(ctx *pulumi.Context,
	name string, args *CustomEntityTypeArgs, opts ...pulumi.ResourceOption) (*CustomEntityType, error) {
	if args == nil {
		args = &CustomEntityTypeArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CustomEntityType
	err := ctx.RegisterResource("aws-native:glue:CustomEntityType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomEntityType gets an existing CustomEntityType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomEntityType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomEntityTypeState, opts ...pulumi.ResourceOption) (*CustomEntityType, error) {
	var resource CustomEntityType
	err := ctx.ReadResource("aws-native:glue:CustomEntityType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomEntityType resources.
type customEntityTypeState struct {
}

type CustomEntityTypeState struct {
}

func (CustomEntityTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*customEntityTypeState)(nil)).Elem()
}

type customEntityTypeArgs struct {
	ContextWords []string `pulumi:"contextWords"`
	Name         *string  `pulumi:"name"`
	RegexString  *string  `pulumi:"regexString"`
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Glue::CustomEntityType` for more information about the expected schema for this property.
	Tags interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a CustomEntityType resource.
type CustomEntityTypeArgs struct {
	ContextWords pulumi.StringArrayInput
	Name         pulumi.StringPtrInput
	RegexString  pulumi.StringPtrInput
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Glue::CustomEntityType` for more information about the expected schema for this property.
	Tags pulumi.Input
}

func (CustomEntityTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customEntityTypeArgs)(nil)).Elem()
}

type CustomEntityTypeInput interface {
	pulumi.Input

	ToCustomEntityTypeOutput() CustomEntityTypeOutput
	ToCustomEntityTypeOutputWithContext(ctx context.Context) CustomEntityTypeOutput
}

func (*CustomEntityType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomEntityType)(nil)).Elem()
}

func (i *CustomEntityType) ToCustomEntityTypeOutput() CustomEntityTypeOutput {
	return i.ToCustomEntityTypeOutputWithContext(context.Background())
}

func (i *CustomEntityType) ToCustomEntityTypeOutputWithContext(ctx context.Context) CustomEntityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomEntityTypeOutput)
}

type CustomEntityTypeOutput struct{ *pulumi.OutputState }

func (CustomEntityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomEntityType)(nil)).Elem()
}

func (o CustomEntityTypeOutput) ToCustomEntityTypeOutput() CustomEntityTypeOutput {
	return o
}

func (o CustomEntityTypeOutput) ToCustomEntityTypeOutputWithContext(ctx context.Context) CustomEntityTypeOutput {
	return o
}

func (o CustomEntityTypeOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomEntityType) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

func (o CustomEntityTypeOutput) ContextWords() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CustomEntityType) pulumi.StringArrayOutput { return v.ContextWords }).(pulumi.StringArrayOutput)
}

func (o CustomEntityTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomEntityType) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o CustomEntityTypeOutput) RegexString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomEntityType) pulumi.StringPtrOutput { return v.RegexString }).(pulumi.StringPtrOutput)
}

// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Glue::CustomEntityType` for more information about the expected schema for this property.
func (o CustomEntityTypeOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v *CustomEntityType) pulumi.AnyOutput { return v.Tags }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomEntityTypeInput)(nil)).Elem(), &CustomEntityType{})
	pulumi.RegisterOutputType(CustomEntityTypeOutput{})
}
