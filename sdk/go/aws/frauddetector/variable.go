// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package frauddetector

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A resource schema for a Variable in Amazon Fraud Detector.
type Variable struct {
	pulumi.CustomResourceState

	// The ARN of the variable.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The time when the variable was created.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// The source of the data.
	DataSource VariableDataSourceOutput `pulumi:"dataSource"`
	// The data type.
	DataType VariableDataTypeOutput `pulumi:"dataType"`
	// The default value for the variable when no value is received.
	DefaultValue pulumi.StringOutput `pulumi:"defaultValue"`
	// The description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The time when the variable was last updated.
	LastUpdatedTime pulumi.StringOutput `pulumi:"lastUpdatedTime"`
	// The name of the variable.
	Name pulumi.StringOutput `pulumi:"name"`
	// Tags associated with this variable.
	Tags VariableTagArrayOutput `pulumi:"tags"`
	// The variable type. For more information see https://docs.aws.amazon.com/frauddetector/latest/ug/create-a-variable.html#variable-types
	VariableType VariableTypePtrOutput `pulumi:"variableType"`
}

// NewVariable registers a new resource with the given unique name, arguments, and options.
func NewVariable(ctx *pulumi.Context,
	name string, args *VariableArgs, opts ...pulumi.ResourceOption) (*Variable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataSource == nil {
		return nil, errors.New("invalid value for required argument 'DataSource'")
	}
	if args.DataType == nil {
		return nil, errors.New("invalid value for required argument 'DataType'")
	}
	if args.DefaultValue == nil {
		return nil, errors.New("invalid value for required argument 'DefaultValue'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Variable
	err := ctx.RegisterResource("aws-native:frauddetector:Variable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVariable gets an existing Variable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVariable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VariableState, opts ...pulumi.ResourceOption) (*Variable, error) {
	var resource Variable
	err := ctx.ReadResource("aws-native:frauddetector:Variable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Variable resources.
type variableState struct {
}

type VariableState struct {
}

func (VariableState) ElementType() reflect.Type {
	return reflect.TypeOf((*variableState)(nil)).Elem()
}

type variableArgs struct {
	// The source of the data.
	DataSource VariableDataSource `pulumi:"dataSource"`
	// The data type.
	DataType VariableDataType `pulumi:"dataType"`
	// The default value for the variable when no value is received.
	DefaultValue string `pulumi:"defaultValue"`
	// The description.
	Description *string `pulumi:"description"`
	// The name of the variable.
	Name *string `pulumi:"name"`
	// Tags associated with this variable.
	Tags []VariableTag `pulumi:"tags"`
	// The variable type. For more information see https://docs.aws.amazon.com/frauddetector/latest/ug/create-a-variable.html#variable-types
	VariableType *VariableType `pulumi:"variableType"`
}

// The set of arguments for constructing a Variable resource.
type VariableArgs struct {
	// The source of the data.
	DataSource VariableDataSourceInput
	// The data type.
	DataType VariableDataTypeInput
	// The default value for the variable when no value is received.
	DefaultValue pulumi.StringInput
	// The description.
	Description pulumi.StringPtrInput
	// The name of the variable.
	Name pulumi.StringPtrInput
	// Tags associated with this variable.
	Tags VariableTagArrayInput
	// The variable type. For more information see https://docs.aws.amazon.com/frauddetector/latest/ug/create-a-variable.html#variable-types
	VariableType VariableTypePtrInput
}

func (VariableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*variableArgs)(nil)).Elem()
}

type VariableInput interface {
	pulumi.Input

	ToVariableOutput() VariableOutput
	ToVariableOutputWithContext(ctx context.Context) VariableOutput
}

func (*Variable) ElementType() reflect.Type {
	return reflect.TypeOf((**Variable)(nil)).Elem()
}

func (i *Variable) ToVariableOutput() VariableOutput {
	return i.ToVariableOutputWithContext(context.Background())
}

func (i *Variable) ToVariableOutputWithContext(ctx context.Context) VariableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VariableOutput)
}

type VariableOutput struct{ *pulumi.OutputState }

func (VariableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Variable)(nil)).Elem()
}

func (o VariableOutput) ToVariableOutput() VariableOutput {
	return o
}

func (o VariableOutput) ToVariableOutputWithContext(ctx context.Context) VariableOutput {
	return o
}

// The ARN of the variable.
func (o VariableOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Variable) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The time when the variable was created.
func (o VariableOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Variable) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// The source of the data.
func (o VariableOutput) DataSource() VariableDataSourceOutput {
	return o.ApplyT(func(v *Variable) VariableDataSourceOutput { return v.DataSource }).(VariableDataSourceOutput)
}

// The data type.
func (o VariableOutput) DataType() VariableDataTypeOutput {
	return o.ApplyT(func(v *Variable) VariableDataTypeOutput { return v.DataType }).(VariableDataTypeOutput)
}

// The default value for the variable when no value is received.
func (o VariableOutput) DefaultValue() pulumi.StringOutput {
	return o.ApplyT(func(v *Variable) pulumi.StringOutput { return v.DefaultValue }).(pulumi.StringOutput)
}

// The description.
func (o VariableOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Variable) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The time when the variable was last updated.
func (o VariableOutput) LastUpdatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Variable) pulumi.StringOutput { return v.LastUpdatedTime }).(pulumi.StringOutput)
}

// The name of the variable.
func (o VariableOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Variable) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Tags associated with this variable.
func (o VariableOutput) Tags() VariableTagArrayOutput {
	return o.ApplyT(func(v *Variable) VariableTagArrayOutput { return v.Tags }).(VariableTagArrayOutput)
}

// The variable type. For more information see https://docs.aws.amazon.com/frauddetector/latest/ug/create-a-variable.html#variable-types
func (o VariableOutput) VariableType() VariableTypePtrOutput {
	return o.ApplyT(func(v *Variable) VariableTypePtrOutput { return v.VariableType }).(VariableTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VariableInput)(nil)).Elem(), &Variable{})
	pulumi.RegisterOutputType(VariableOutput{})
}
