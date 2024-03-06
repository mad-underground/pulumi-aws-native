// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::CloudFront::KeyValueStore
type KeyValueStore struct {
	pulumi.CustomResourceState

	Arn          pulumi.StringOutput                `pulumi:"arn"`
	AwsId        pulumi.StringOutput                `pulumi:"awsId"`
	Comment      pulumi.StringPtrOutput             `pulumi:"comment"`
	ImportSource KeyValueStoreImportSourcePtrOutput `pulumi:"importSource"`
	Name         pulumi.StringOutput                `pulumi:"name"`
	Status       pulumi.StringOutput                `pulumi:"status"`
}

// NewKeyValueStore registers a new resource with the given unique name, arguments, and options.
func NewKeyValueStore(ctx *pulumi.Context,
	name string, args *KeyValueStoreArgs, opts ...pulumi.ResourceOption) (*KeyValueStore, error) {
	if args == nil {
		args = &KeyValueStoreArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"name",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource KeyValueStore
	err := ctx.RegisterResource("aws-native:cloudfront:KeyValueStore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeyValueStore gets an existing KeyValueStore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeyValueStore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyValueStoreState, opts ...pulumi.ResourceOption) (*KeyValueStore, error) {
	var resource KeyValueStore
	err := ctx.ReadResource("aws-native:cloudfront:KeyValueStore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KeyValueStore resources.
type keyValueStoreState struct {
}

type KeyValueStoreState struct {
}

func (KeyValueStoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyValueStoreState)(nil)).Elem()
}

type keyValueStoreArgs struct {
	Comment      *string                    `pulumi:"comment"`
	ImportSource *KeyValueStoreImportSource `pulumi:"importSource"`
	Name         *string                    `pulumi:"name"`
}

// The set of arguments for constructing a KeyValueStore resource.
type KeyValueStoreArgs struct {
	Comment      pulumi.StringPtrInput
	ImportSource KeyValueStoreImportSourcePtrInput
	Name         pulumi.StringPtrInput
}

func (KeyValueStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyValueStoreArgs)(nil)).Elem()
}

type KeyValueStoreInput interface {
	pulumi.Input

	ToKeyValueStoreOutput() KeyValueStoreOutput
	ToKeyValueStoreOutputWithContext(ctx context.Context) KeyValueStoreOutput
}

func (*KeyValueStore) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyValueStore)(nil)).Elem()
}

func (i *KeyValueStore) ToKeyValueStoreOutput() KeyValueStoreOutput {
	return i.ToKeyValueStoreOutputWithContext(context.Background())
}

func (i *KeyValueStore) ToKeyValueStoreOutputWithContext(ctx context.Context) KeyValueStoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyValueStoreOutput)
}

type KeyValueStoreOutput struct{ *pulumi.OutputState }

func (KeyValueStoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyValueStore)(nil)).Elem()
}

func (o KeyValueStoreOutput) ToKeyValueStoreOutput() KeyValueStoreOutput {
	return o
}

func (o KeyValueStoreOutput) ToKeyValueStoreOutputWithContext(ctx context.Context) KeyValueStoreOutput {
	return o
}

func (o KeyValueStoreOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyValueStore) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o KeyValueStoreOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyValueStore) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

func (o KeyValueStoreOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyValueStore) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o KeyValueStoreOutput) ImportSource() KeyValueStoreImportSourcePtrOutput {
	return o.ApplyT(func(v *KeyValueStore) KeyValueStoreImportSourcePtrOutput { return v.ImportSource }).(KeyValueStoreImportSourcePtrOutput)
}

func (o KeyValueStoreOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyValueStore) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o KeyValueStoreOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyValueStore) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeyValueStoreInput)(nil)).Elem(), &KeyValueStore{})
	pulumi.RegisterOutputType(KeyValueStoreOutput{})
}
