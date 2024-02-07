// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package personalize

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::Personalize::Schema.
type Schema struct {
	pulumi.CustomResourceState

	// The domain of a Domain dataset group.
	Domain SchemaDomainPtrOutput `pulumi:"domain"`
	// Name for the schema.
	Name pulumi.StringOutput `pulumi:"name"`
	// A schema in Avro JSON format.
	Schema pulumi.StringOutput `pulumi:"schema"`
	// Arn for the schema.
	SchemaArn pulumi.StringOutput `pulumi:"schemaArn"`
}

// NewSchema registers a new resource with the given unique name, arguments, and options.
func NewSchema(ctx *pulumi.Context,
	name string, args *SchemaArgs, opts ...pulumi.ResourceOption) (*Schema, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Schema == nil {
		return nil, errors.New("invalid value for required argument 'Schema'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"domain",
		"name",
		"schema",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Schema
	err := ctx.RegisterResource("aws-native:personalize:Schema", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSchema gets an existing Schema resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSchema(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SchemaState, opts ...pulumi.ResourceOption) (*Schema, error) {
	var resource Schema
	err := ctx.ReadResource("aws-native:personalize:Schema", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Schema resources.
type schemaState struct {
}

type SchemaState struct {
}

func (SchemaState) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaState)(nil)).Elem()
}

type schemaArgs struct {
	// The domain of a Domain dataset group.
	Domain *SchemaDomain `pulumi:"domain"`
	// Name for the schema.
	Name *string `pulumi:"name"`
	// A schema in Avro JSON format.
	Schema string `pulumi:"schema"`
}

// The set of arguments for constructing a Schema resource.
type SchemaArgs struct {
	// The domain of a Domain dataset group.
	Domain SchemaDomainPtrInput
	// Name for the schema.
	Name pulumi.StringPtrInput
	// A schema in Avro JSON format.
	Schema pulumi.StringInput
}

func (SchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaArgs)(nil)).Elem()
}

type SchemaInput interface {
	pulumi.Input

	ToSchemaOutput() SchemaOutput
	ToSchemaOutputWithContext(ctx context.Context) SchemaOutput
}

func (*Schema) ElementType() reflect.Type {
	return reflect.TypeOf((**Schema)(nil)).Elem()
}

func (i *Schema) ToSchemaOutput() SchemaOutput {
	return i.ToSchemaOutputWithContext(context.Background())
}

func (i *Schema) ToSchemaOutputWithContext(ctx context.Context) SchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaOutput)
}

type SchemaOutput struct{ *pulumi.OutputState }

func (SchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Schema)(nil)).Elem()
}

func (o SchemaOutput) ToSchemaOutput() SchemaOutput {
	return o
}

func (o SchemaOutput) ToSchemaOutputWithContext(ctx context.Context) SchemaOutput {
	return o
}

// The domain of a Domain dataset group.
func (o SchemaOutput) Domain() SchemaDomainPtrOutput {
	return o.ApplyT(func(v *Schema) SchemaDomainPtrOutput { return v.Domain }).(SchemaDomainPtrOutput)
}

// Name for the schema.
func (o SchemaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A schema in Avro JSON format.
func (o SchemaOutput) Schema() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.Schema }).(pulumi.StringOutput)
}

// Arn for the schema.
func (o SchemaOutput) SchemaArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.SchemaArn }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SchemaInput)(nil)).Elem(), &Schema{})
	pulumi.RegisterOutputType(SchemaOutput{})
}
