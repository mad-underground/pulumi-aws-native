// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventschemas

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EventSchemas::Schema
type Schema struct {
	pulumi.CustomResourceState

	// The source of the schema definition.
	Content pulumi.StringOutput `pulumi:"content"`
	// A description of the schema.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The last modified time of the schema.
	LastModified pulumi.StringOutput `pulumi:"lastModified"`
	// The name of the schema registry.
	RegistryName pulumi.StringOutput `pulumi:"registryName"`
	// The ARN of the schema.
	SchemaArn pulumi.StringOutput `pulumi:"schemaArn"`
	// The name of the schema.
	SchemaName pulumi.StringPtrOutput `pulumi:"schemaName"`
	// The version number of the schema.
	SchemaVersion pulumi.StringOutput `pulumi:"schemaVersion"`
	// Tags associated with the resource.
	Tags SchemaTagsEntryArrayOutput `pulumi:"tags"`
	// The type of schema. Valid types include OpenApi3 and JSONSchemaDraft4.
	Type pulumi.StringOutput `pulumi:"type"`
	// The date the schema version was created.
	VersionCreatedDate pulumi.StringOutput `pulumi:"versionCreatedDate"`
}

// NewSchema registers a new resource with the given unique name, arguments, and options.
func NewSchema(ctx *pulumi.Context,
	name string, args *SchemaArgs, opts ...pulumi.ResourceOption) (*Schema, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"registryName",
		"schemaName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Schema
	err := ctx.RegisterResource("aws-native:eventschemas:Schema", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws-native:eventschemas:Schema", name, id, state, &resource, opts...)
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
	// The source of the schema definition.
	Content string `pulumi:"content"`
	// A description of the schema.
	Description *string `pulumi:"description"`
	// The name of the schema registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the schema.
	SchemaName *string `pulumi:"schemaName"`
	// Tags associated with the resource.
	Tags []SchemaTagsEntry `pulumi:"tags"`
	// The type of schema. Valid types include OpenApi3 and JSONSchemaDraft4.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Schema resource.
type SchemaArgs struct {
	// The source of the schema definition.
	Content pulumi.StringInput
	// A description of the schema.
	Description pulumi.StringPtrInput
	// The name of the schema registry.
	RegistryName pulumi.StringInput
	// The name of the schema.
	SchemaName pulumi.StringPtrInput
	// Tags associated with the resource.
	Tags SchemaTagsEntryArrayInput
	// The type of schema. Valid types include OpenApi3 and JSONSchemaDraft4.
	Type pulumi.StringInput
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

// The source of the schema definition.
func (o SchemaOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.Content }).(pulumi.StringOutput)
}

// A description of the schema.
func (o SchemaOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The last modified time of the schema.
func (o SchemaOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.LastModified }).(pulumi.StringOutput)
}

// The name of the schema registry.
func (o SchemaOutput) RegistryName() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.RegistryName }).(pulumi.StringOutput)
}

// The ARN of the schema.
func (o SchemaOutput) SchemaArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.SchemaArn }).(pulumi.StringOutput)
}

// The name of the schema.
func (o SchemaOutput) SchemaName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringPtrOutput { return v.SchemaName }).(pulumi.StringPtrOutput)
}

// The version number of the schema.
func (o SchemaOutput) SchemaVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.SchemaVersion }).(pulumi.StringOutput)
}

// Tags associated with the resource.
func (o SchemaOutput) Tags() SchemaTagsEntryArrayOutput {
	return o.ApplyT(func(v *Schema) SchemaTagsEntryArrayOutput { return v.Tags }).(SchemaTagsEntryArrayOutput)
}

// The type of schema. Valid types include OpenApi3 and JSONSchemaDraft4.
func (o SchemaOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The date the schema version was created.
func (o SchemaOutput) VersionCreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.VersionCreatedDate }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SchemaInput)(nil)).Elem(), &Schema{})
	pulumi.RegisterOutputType(SchemaOutput{})
}
