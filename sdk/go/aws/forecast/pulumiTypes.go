// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package forecast

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AttributesItemProperties struct {
	// Name of the dataset field
	AttributeName *string `pulumi:"attributeName"`
	// Data type of the field
	AttributeType *DatasetAttributesItemPropertiesAttributeType `pulumi:"attributeType"`
}

// AttributesItemPropertiesInput is an input type that accepts AttributesItemPropertiesArgs and AttributesItemPropertiesOutput values.
// You can construct a concrete instance of `AttributesItemPropertiesInput` via:
//
//          AttributesItemPropertiesArgs{...}
type AttributesItemPropertiesInput interface {
	pulumi.Input

	ToAttributesItemPropertiesOutput() AttributesItemPropertiesOutput
	ToAttributesItemPropertiesOutputWithContext(context.Context) AttributesItemPropertiesOutput
}

type AttributesItemPropertiesArgs struct {
	// Name of the dataset field
	AttributeName pulumi.StringPtrInput `pulumi:"attributeName"`
	// Data type of the field
	AttributeType DatasetAttributesItemPropertiesAttributeTypePtrInput `pulumi:"attributeType"`
}

func (AttributesItemPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AttributesItemProperties)(nil)).Elem()
}

func (i AttributesItemPropertiesArgs) ToAttributesItemPropertiesOutput() AttributesItemPropertiesOutput {
	return i.ToAttributesItemPropertiesOutputWithContext(context.Background())
}

func (i AttributesItemPropertiesArgs) ToAttributesItemPropertiesOutputWithContext(ctx context.Context) AttributesItemPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttributesItemPropertiesOutput)
}

// AttributesItemPropertiesArrayInput is an input type that accepts AttributesItemPropertiesArray and AttributesItemPropertiesArrayOutput values.
// You can construct a concrete instance of `AttributesItemPropertiesArrayInput` via:
//
//          AttributesItemPropertiesArray{ AttributesItemPropertiesArgs{...} }
type AttributesItemPropertiesArrayInput interface {
	pulumi.Input

	ToAttributesItemPropertiesArrayOutput() AttributesItemPropertiesArrayOutput
	ToAttributesItemPropertiesArrayOutputWithContext(context.Context) AttributesItemPropertiesArrayOutput
}

type AttributesItemPropertiesArray []AttributesItemPropertiesInput

func (AttributesItemPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AttributesItemProperties)(nil)).Elem()
}

func (i AttributesItemPropertiesArray) ToAttributesItemPropertiesArrayOutput() AttributesItemPropertiesArrayOutput {
	return i.ToAttributesItemPropertiesArrayOutputWithContext(context.Background())
}

func (i AttributesItemPropertiesArray) ToAttributesItemPropertiesArrayOutputWithContext(ctx context.Context) AttributesItemPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttributesItemPropertiesArrayOutput)
}

type AttributesItemPropertiesOutput struct{ *pulumi.OutputState }

func (AttributesItemPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AttributesItemProperties)(nil)).Elem()
}

func (o AttributesItemPropertiesOutput) ToAttributesItemPropertiesOutput() AttributesItemPropertiesOutput {
	return o
}

func (o AttributesItemPropertiesOutput) ToAttributesItemPropertiesOutputWithContext(ctx context.Context) AttributesItemPropertiesOutput {
	return o
}

// Name of the dataset field
func (o AttributesItemPropertiesOutput) AttributeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AttributesItemProperties) *string { return v.AttributeName }).(pulumi.StringPtrOutput)
}

// Data type of the field
func (o AttributesItemPropertiesOutput) AttributeType() DatasetAttributesItemPropertiesAttributeTypePtrOutput {
	return o.ApplyT(func(v AttributesItemProperties) *DatasetAttributesItemPropertiesAttributeType { return v.AttributeType }).(DatasetAttributesItemPropertiesAttributeTypePtrOutput)
}

type AttributesItemPropertiesArrayOutput struct{ *pulumi.OutputState }

func (AttributesItemPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AttributesItemProperties)(nil)).Elem()
}

func (o AttributesItemPropertiesArrayOutput) ToAttributesItemPropertiesArrayOutput() AttributesItemPropertiesArrayOutput {
	return o
}

func (o AttributesItemPropertiesArrayOutput) ToAttributesItemPropertiesArrayOutputWithContext(ctx context.Context) AttributesItemPropertiesArrayOutput {
	return o
}

func (o AttributesItemPropertiesArrayOutput) Index(i pulumi.IntInput) AttributesItemPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AttributesItemProperties {
		return vs[0].([]AttributesItemProperties)[vs[1].(int)]
	}).(AttributesItemPropertiesOutput)
}

// A key-value pair to associate with a resource.
type DatasetGroupTag struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key string `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value string `pulumi:"value"`
}

// DatasetGroupTagInput is an input type that accepts DatasetGroupTagArgs and DatasetGroupTagOutput values.
// You can construct a concrete instance of `DatasetGroupTagInput` via:
//
//          DatasetGroupTagArgs{...}
type DatasetGroupTagInput interface {
	pulumi.Input

	ToDatasetGroupTagOutput() DatasetGroupTagOutput
	ToDatasetGroupTagOutputWithContext(context.Context) DatasetGroupTagOutput
}

// A key-value pair to associate with a resource.
type DatasetGroupTagArgs struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key pulumi.StringInput `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value pulumi.StringInput `pulumi:"value"`
}

func (DatasetGroupTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetGroupTag)(nil)).Elem()
}

func (i DatasetGroupTagArgs) ToDatasetGroupTagOutput() DatasetGroupTagOutput {
	return i.ToDatasetGroupTagOutputWithContext(context.Background())
}

func (i DatasetGroupTagArgs) ToDatasetGroupTagOutputWithContext(ctx context.Context) DatasetGroupTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetGroupTagOutput)
}

// DatasetGroupTagArrayInput is an input type that accepts DatasetGroupTagArray and DatasetGroupTagArrayOutput values.
// You can construct a concrete instance of `DatasetGroupTagArrayInput` via:
//
//          DatasetGroupTagArray{ DatasetGroupTagArgs{...} }
type DatasetGroupTagArrayInput interface {
	pulumi.Input

	ToDatasetGroupTagArrayOutput() DatasetGroupTagArrayOutput
	ToDatasetGroupTagArrayOutputWithContext(context.Context) DatasetGroupTagArrayOutput
}

type DatasetGroupTagArray []DatasetGroupTagInput

func (DatasetGroupTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatasetGroupTag)(nil)).Elem()
}

func (i DatasetGroupTagArray) ToDatasetGroupTagArrayOutput() DatasetGroupTagArrayOutput {
	return i.ToDatasetGroupTagArrayOutputWithContext(context.Background())
}

func (i DatasetGroupTagArray) ToDatasetGroupTagArrayOutputWithContext(ctx context.Context) DatasetGroupTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetGroupTagArrayOutput)
}

// A key-value pair to associate with a resource.
type DatasetGroupTagOutput struct{ *pulumi.OutputState }

func (DatasetGroupTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetGroupTag)(nil)).Elem()
}

func (o DatasetGroupTagOutput) ToDatasetGroupTagOutput() DatasetGroupTagOutput {
	return o
}

func (o DatasetGroupTagOutput) ToDatasetGroupTagOutputWithContext(ctx context.Context) DatasetGroupTagOutput {
	return o
}

// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o DatasetGroupTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetGroupTag) string { return v.Key }).(pulumi.StringOutput)
}

// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o DatasetGroupTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetGroupTag) string { return v.Value }).(pulumi.StringOutput)
}

type DatasetGroupTagArrayOutput struct{ *pulumi.OutputState }

func (DatasetGroupTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatasetGroupTag)(nil)).Elem()
}

func (o DatasetGroupTagArrayOutput) ToDatasetGroupTagArrayOutput() DatasetGroupTagArrayOutput {
	return o
}

func (o DatasetGroupTagArrayOutput) ToDatasetGroupTagArrayOutputWithContext(ctx context.Context) DatasetGroupTagArrayOutput {
	return o
}

func (o DatasetGroupTagArrayOutput) Index(i pulumi.IntInput) DatasetGroupTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatasetGroupTag {
		return vs[0].([]DatasetGroupTag)[vs[1].(int)]
	}).(DatasetGroupTagOutput)
}

type EncryptionConfigProperties struct {
	KmsKeyArn *string `pulumi:"kmsKeyArn"`
	RoleArn   *string `pulumi:"roleArn"`
}

// EncryptionConfigPropertiesInput is an input type that accepts EncryptionConfigPropertiesArgs and EncryptionConfigPropertiesOutput values.
// You can construct a concrete instance of `EncryptionConfigPropertiesInput` via:
//
//          EncryptionConfigPropertiesArgs{...}
type EncryptionConfigPropertiesInput interface {
	pulumi.Input

	ToEncryptionConfigPropertiesOutput() EncryptionConfigPropertiesOutput
	ToEncryptionConfigPropertiesOutputWithContext(context.Context) EncryptionConfigPropertiesOutput
}

type EncryptionConfigPropertiesArgs struct {
	KmsKeyArn pulumi.StringPtrInput `pulumi:"kmsKeyArn"`
	RoleArn   pulumi.StringPtrInput `pulumi:"roleArn"`
}

func (EncryptionConfigPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionConfigProperties)(nil)).Elem()
}

func (i EncryptionConfigPropertiesArgs) ToEncryptionConfigPropertiesOutput() EncryptionConfigPropertiesOutput {
	return i.ToEncryptionConfigPropertiesOutputWithContext(context.Background())
}

func (i EncryptionConfigPropertiesArgs) ToEncryptionConfigPropertiesOutputWithContext(ctx context.Context) EncryptionConfigPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionConfigPropertiesOutput)
}

func (i EncryptionConfigPropertiesArgs) ToEncryptionConfigPropertiesPtrOutput() EncryptionConfigPropertiesPtrOutput {
	return i.ToEncryptionConfigPropertiesPtrOutputWithContext(context.Background())
}

func (i EncryptionConfigPropertiesArgs) ToEncryptionConfigPropertiesPtrOutputWithContext(ctx context.Context) EncryptionConfigPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionConfigPropertiesOutput).ToEncryptionConfigPropertiesPtrOutputWithContext(ctx)
}

// EncryptionConfigPropertiesPtrInput is an input type that accepts EncryptionConfigPropertiesArgs, EncryptionConfigPropertiesPtr and EncryptionConfigPropertiesPtrOutput values.
// You can construct a concrete instance of `EncryptionConfigPropertiesPtrInput` via:
//
//          EncryptionConfigPropertiesArgs{...}
//
//  or:
//
//          nil
type EncryptionConfigPropertiesPtrInput interface {
	pulumi.Input

	ToEncryptionConfigPropertiesPtrOutput() EncryptionConfigPropertiesPtrOutput
	ToEncryptionConfigPropertiesPtrOutputWithContext(context.Context) EncryptionConfigPropertiesPtrOutput
}

type encryptionConfigPropertiesPtrType EncryptionConfigPropertiesArgs

func EncryptionConfigPropertiesPtr(v *EncryptionConfigPropertiesArgs) EncryptionConfigPropertiesPtrInput {
	return (*encryptionConfigPropertiesPtrType)(v)
}

func (*encryptionConfigPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionConfigProperties)(nil)).Elem()
}

func (i *encryptionConfigPropertiesPtrType) ToEncryptionConfigPropertiesPtrOutput() EncryptionConfigPropertiesPtrOutput {
	return i.ToEncryptionConfigPropertiesPtrOutputWithContext(context.Background())
}

func (i *encryptionConfigPropertiesPtrType) ToEncryptionConfigPropertiesPtrOutputWithContext(ctx context.Context) EncryptionConfigPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionConfigPropertiesPtrOutput)
}

type EncryptionConfigPropertiesOutput struct{ *pulumi.OutputState }

func (EncryptionConfigPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionConfigProperties)(nil)).Elem()
}

func (o EncryptionConfigPropertiesOutput) ToEncryptionConfigPropertiesOutput() EncryptionConfigPropertiesOutput {
	return o
}

func (o EncryptionConfigPropertiesOutput) ToEncryptionConfigPropertiesOutputWithContext(ctx context.Context) EncryptionConfigPropertiesOutput {
	return o
}

func (o EncryptionConfigPropertiesOutput) ToEncryptionConfigPropertiesPtrOutput() EncryptionConfigPropertiesPtrOutput {
	return o.ToEncryptionConfigPropertiesPtrOutputWithContext(context.Background())
}

func (o EncryptionConfigPropertiesOutput) ToEncryptionConfigPropertiesPtrOutputWithContext(ctx context.Context) EncryptionConfigPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionConfigProperties) *EncryptionConfigProperties {
		return &v
	}).(EncryptionConfigPropertiesPtrOutput)
}

func (o EncryptionConfigPropertiesOutput) KmsKeyArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionConfigProperties) *string { return v.KmsKeyArn }).(pulumi.StringPtrOutput)
}

func (o EncryptionConfigPropertiesOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionConfigProperties) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

type EncryptionConfigPropertiesPtrOutput struct{ *pulumi.OutputState }

func (EncryptionConfigPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionConfigProperties)(nil)).Elem()
}

func (o EncryptionConfigPropertiesPtrOutput) ToEncryptionConfigPropertiesPtrOutput() EncryptionConfigPropertiesPtrOutput {
	return o
}

func (o EncryptionConfigPropertiesPtrOutput) ToEncryptionConfigPropertiesPtrOutputWithContext(ctx context.Context) EncryptionConfigPropertiesPtrOutput {
	return o
}

func (o EncryptionConfigPropertiesPtrOutput) Elem() EncryptionConfigPropertiesOutput {
	return o.ApplyT(func(v *EncryptionConfigProperties) EncryptionConfigProperties {
		if v != nil {
			return *v
		}
		var ret EncryptionConfigProperties
		return ret
	}).(EncryptionConfigPropertiesOutput)
}

func (o EncryptionConfigPropertiesPtrOutput) KmsKeyArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionConfigProperties) *string {
		if v == nil {
			return nil
		}
		return v.KmsKeyArn
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionConfigPropertiesPtrOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionConfigProperties) *string {
		if v == nil {
			return nil
		}
		return v.RoleArn
	}).(pulumi.StringPtrOutput)
}

type SchemaProperties struct {
	Attributes []AttributesItemProperties `pulumi:"attributes"`
}

// SchemaPropertiesInput is an input type that accepts SchemaPropertiesArgs and SchemaPropertiesOutput values.
// You can construct a concrete instance of `SchemaPropertiesInput` via:
//
//          SchemaPropertiesArgs{...}
type SchemaPropertiesInput interface {
	pulumi.Input

	ToSchemaPropertiesOutput() SchemaPropertiesOutput
	ToSchemaPropertiesOutputWithContext(context.Context) SchemaPropertiesOutput
}

type SchemaPropertiesArgs struct {
	Attributes AttributesItemPropertiesArrayInput `pulumi:"attributes"`
}

func (SchemaPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SchemaProperties)(nil)).Elem()
}

func (i SchemaPropertiesArgs) ToSchemaPropertiesOutput() SchemaPropertiesOutput {
	return i.ToSchemaPropertiesOutputWithContext(context.Background())
}

func (i SchemaPropertiesArgs) ToSchemaPropertiesOutputWithContext(ctx context.Context) SchemaPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaPropertiesOutput)
}

type SchemaPropertiesOutput struct{ *pulumi.OutputState }

func (SchemaPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SchemaProperties)(nil)).Elem()
}

func (o SchemaPropertiesOutput) ToSchemaPropertiesOutput() SchemaPropertiesOutput {
	return o
}

func (o SchemaPropertiesOutput) ToSchemaPropertiesOutputWithContext(ctx context.Context) SchemaPropertiesOutput {
	return o
}

func (o SchemaPropertiesOutput) Attributes() AttributesItemPropertiesArrayOutput {
	return o.ApplyT(func(v SchemaProperties) []AttributesItemProperties { return v.Attributes }).(AttributesItemPropertiesArrayOutput)
}

type SchemaPropertiesPtrOutput struct{ *pulumi.OutputState }

func (SchemaPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SchemaProperties)(nil)).Elem()
}

func (o SchemaPropertiesPtrOutput) ToSchemaPropertiesPtrOutput() SchemaPropertiesPtrOutput {
	return o
}

func (o SchemaPropertiesPtrOutput) ToSchemaPropertiesPtrOutputWithContext(ctx context.Context) SchemaPropertiesPtrOutput {
	return o
}

func (o SchemaPropertiesPtrOutput) Elem() SchemaPropertiesOutput {
	return o.ApplyT(func(v *SchemaProperties) SchemaProperties {
		if v != nil {
			return *v
		}
		var ret SchemaProperties
		return ret
	}).(SchemaPropertiesOutput)
}

func (o SchemaPropertiesPtrOutput) Attributes() AttributesItemPropertiesArrayOutput {
	return o.ApplyT(func(v *SchemaProperties) []AttributesItemProperties {
		if v == nil {
			return nil
		}
		return v.Attributes
	}).(AttributesItemPropertiesArrayOutput)
}

// A key-value pair to associate with a resource.
type TagsItemProperties struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

// TagsItemPropertiesInput is an input type that accepts TagsItemPropertiesArgs and TagsItemPropertiesOutput values.
// You can construct a concrete instance of `TagsItemPropertiesInput` via:
//
//          TagsItemPropertiesArgs{...}
type TagsItemPropertiesInput interface {
	pulumi.Input

	ToTagsItemPropertiesOutput() TagsItemPropertiesOutput
	ToTagsItemPropertiesOutputWithContext(context.Context) TagsItemPropertiesOutput
}

// A key-value pair to associate with a resource.
type TagsItemPropertiesArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (TagsItemPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TagsItemProperties)(nil)).Elem()
}

func (i TagsItemPropertiesArgs) ToTagsItemPropertiesOutput() TagsItemPropertiesOutput {
	return i.ToTagsItemPropertiesOutputWithContext(context.Background())
}

func (i TagsItemPropertiesArgs) ToTagsItemPropertiesOutputWithContext(ctx context.Context) TagsItemPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagsItemPropertiesOutput)
}

// TagsItemPropertiesArrayInput is an input type that accepts TagsItemPropertiesArray and TagsItemPropertiesArrayOutput values.
// You can construct a concrete instance of `TagsItemPropertiesArrayInput` via:
//
//          TagsItemPropertiesArray{ TagsItemPropertiesArgs{...} }
type TagsItemPropertiesArrayInput interface {
	pulumi.Input

	ToTagsItemPropertiesArrayOutput() TagsItemPropertiesArrayOutput
	ToTagsItemPropertiesArrayOutputWithContext(context.Context) TagsItemPropertiesArrayOutput
}

type TagsItemPropertiesArray []TagsItemPropertiesInput

func (TagsItemPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TagsItemProperties)(nil)).Elem()
}

func (i TagsItemPropertiesArray) ToTagsItemPropertiesArrayOutput() TagsItemPropertiesArrayOutput {
	return i.ToTagsItemPropertiesArrayOutputWithContext(context.Background())
}

func (i TagsItemPropertiesArray) ToTagsItemPropertiesArrayOutputWithContext(ctx context.Context) TagsItemPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagsItemPropertiesArrayOutput)
}

// A key-value pair to associate with a resource.
type TagsItemPropertiesOutput struct{ *pulumi.OutputState }

func (TagsItemPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagsItemProperties)(nil)).Elem()
}

func (o TagsItemPropertiesOutput) ToTagsItemPropertiesOutput() TagsItemPropertiesOutput {
	return o
}

func (o TagsItemPropertiesOutput) ToTagsItemPropertiesOutputWithContext(ctx context.Context) TagsItemPropertiesOutput {
	return o
}

func (o TagsItemPropertiesOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v TagsItemProperties) string { return v.Key }).(pulumi.StringOutput)
}

func (o TagsItemPropertiesOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v TagsItemProperties) string { return v.Value }).(pulumi.StringOutput)
}

type TagsItemPropertiesArrayOutput struct{ *pulumi.OutputState }

func (TagsItemPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TagsItemProperties)(nil)).Elem()
}

func (o TagsItemPropertiesArrayOutput) ToTagsItemPropertiesArrayOutput() TagsItemPropertiesArrayOutput {
	return o
}

func (o TagsItemPropertiesArrayOutput) ToTagsItemPropertiesArrayOutputWithContext(ctx context.Context) TagsItemPropertiesArrayOutput {
	return o
}

func (o TagsItemPropertiesArrayOutput) Index(i pulumi.IntInput) TagsItemPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TagsItemProperties {
		return vs[0].([]TagsItemProperties)[vs[1].(int)]
	}).(TagsItemPropertiesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AttributesItemPropertiesInput)(nil)).Elem(), AttributesItemPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AttributesItemPropertiesArrayInput)(nil)).Elem(), AttributesItemPropertiesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetGroupTagInput)(nil)).Elem(), DatasetGroupTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetGroupTagArrayInput)(nil)).Elem(), DatasetGroupTagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionConfigPropertiesInput)(nil)).Elem(), EncryptionConfigPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionConfigPropertiesPtrInput)(nil)).Elem(), EncryptionConfigPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SchemaPropertiesInput)(nil)).Elem(), SchemaPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TagsItemPropertiesInput)(nil)).Elem(), TagsItemPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TagsItemPropertiesArrayInput)(nil)).Elem(), TagsItemPropertiesArray{})
	pulumi.RegisterOutputType(AttributesItemPropertiesOutput{})
	pulumi.RegisterOutputType(AttributesItemPropertiesArrayOutput{})
	pulumi.RegisterOutputType(DatasetGroupTagOutput{})
	pulumi.RegisterOutputType(DatasetGroupTagArrayOutput{})
	pulumi.RegisterOutputType(EncryptionConfigPropertiesOutput{})
	pulumi.RegisterOutputType(EncryptionConfigPropertiesPtrOutput{})
	pulumi.RegisterOutputType(SchemaPropertiesOutput{})
	pulumi.RegisterOutputType(SchemaPropertiesPtrOutput{})
	pulumi.RegisterOutputType(TagsItemPropertiesOutput{})
	pulumi.RegisterOutputType(TagsItemPropertiesArrayOutput{})
}
