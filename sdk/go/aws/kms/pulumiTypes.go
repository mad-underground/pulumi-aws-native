// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

// A key-value pair to associate with a resource.
type KeyTag struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key string `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value string `pulumi:"value"`
}

// KeyTagInput is an input type that accepts KeyTagArgs and KeyTagOutput values.
// You can construct a concrete instance of `KeyTagInput` via:
//
//	KeyTagArgs{...}
type KeyTagInput interface {
	pulumi.Input

	ToKeyTagOutput() KeyTagOutput
	ToKeyTagOutputWithContext(context.Context) KeyTagOutput
}

// A key-value pair to associate with a resource.
type KeyTagArgs struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key pulumi.StringInput `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value pulumi.StringInput `pulumi:"value"`
}

func (KeyTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyTag)(nil)).Elem()
}

func (i KeyTagArgs) ToKeyTagOutput() KeyTagOutput {
	return i.ToKeyTagOutputWithContext(context.Background())
}

func (i KeyTagArgs) ToKeyTagOutputWithContext(ctx context.Context) KeyTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyTagOutput)
}

// KeyTagArrayInput is an input type that accepts KeyTagArray and KeyTagArrayOutput values.
// You can construct a concrete instance of `KeyTagArrayInput` via:
//
//	KeyTagArray{ KeyTagArgs{...} }
type KeyTagArrayInput interface {
	pulumi.Input

	ToKeyTagArrayOutput() KeyTagArrayOutput
	ToKeyTagArrayOutputWithContext(context.Context) KeyTagArrayOutput
}

type KeyTagArray []KeyTagInput

func (KeyTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyTag)(nil)).Elem()
}

func (i KeyTagArray) ToKeyTagArrayOutput() KeyTagArrayOutput {
	return i.ToKeyTagArrayOutputWithContext(context.Background())
}

func (i KeyTagArray) ToKeyTagArrayOutputWithContext(ctx context.Context) KeyTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyTagArrayOutput)
}

// A key-value pair to associate with a resource.
type KeyTagOutput struct{ *pulumi.OutputState }

func (KeyTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyTag)(nil)).Elem()
}

func (o KeyTagOutput) ToKeyTagOutput() KeyTagOutput {
	return o
}

func (o KeyTagOutput) ToKeyTagOutputWithContext(ctx context.Context) KeyTagOutput {
	return o
}

// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o KeyTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v KeyTag) string { return v.Key }).(pulumi.StringOutput)
}

// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o KeyTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v KeyTag) string { return v.Value }).(pulumi.StringOutput)
}

type KeyTagArrayOutput struct{ *pulumi.OutputState }

func (KeyTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyTag)(nil)).Elem()
}

func (o KeyTagArrayOutput) ToKeyTagArrayOutput() KeyTagArrayOutput {
	return o
}

func (o KeyTagArrayOutput) ToKeyTagArrayOutputWithContext(ctx context.Context) KeyTagArrayOutput {
	return o
}

func (o KeyTagArrayOutput) Index(i pulumi.IntInput) KeyTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeyTag {
		return vs[0].([]KeyTag)[vs[1].(int)]
	}).(KeyTagOutput)
}

// A key-value pair to associate with a resource.
type ReplicaKeyTag struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key string `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value string `pulumi:"value"`
}

// ReplicaKeyTagInput is an input type that accepts ReplicaKeyTagArgs and ReplicaKeyTagOutput values.
// You can construct a concrete instance of `ReplicaKeyTagInput` via:
//
//	ReplicaKeyTagArgs{...}
type ReplicaKeyTagInput interface {
	pulumi.Input

	ToReplicaKeyTagOutput() ReplicaKeyTagOutput
	ToReplicaKeyTagOutputWithContext(context.Context) ReplicaKeyTagOutput
}

// A key-value pair to associate with a resource.
type ReplicaKeyTagArgs struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key pulumi.StringInput `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value pulumi.StringInput `pulumi:"value"`
}

func (ReplicaKeyTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicaKeyTag)(nil)).Elem()
}

func (i ReplicaKeyTagArgs) ToReplicaKeyTagOutput() ReplicaKeyTagOutput {
	return i.ToReplicaKeyTagOutputWithContext(context.Background())
}

func (i ReplicaKeyTagArgs) ToReplicaKeyTagOutputWithContext(ctx context.Context) ReplicaKeyTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicaKeyTagOutput)
}

// ReplicaKeyTagArrayInput is an input type that accepts ReplicaKeyTagArray and ReplicaKeyTagArrayOutput values.
// You can construct a concrete instance of `ReplicaKeyTagArrayInput` via:
//
//	ReplicaKeyTagArray{ ReplicaKeyTagArgs{...} }
type ReplicaKeyTagArrayInput interface {
	pulumi.Input

	ToReplicaKeyTagArrayOutput() ReplicaKeyTagArrayOutput
	ToReplicaKeyTagArrayOutputWithContext(context.Context) ReplicaKeyTagArrayOutput
}

type ReplicaKeyTagArray []ReplicaKeyTagInput

func (ReplicaKeyTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReplicaKeyTag)(nil)).Elem()
}

func (i ReplicaKeyTagArray) ToReplicaKeyTagArrayOutput() ReplicaKeyTagArrayOutput {
	return i.ToReplicaKeyTagArrayOutputWithContext(context.Background())
}

func (i ReplicaKeyTagArray) ToReplicaKeyTagArrayOutputWithContext(ctx context.Context) ReplicaKeyTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicaKeyTagArrayOutput)
}

// A key-value pair to associate with a resource.
type ReplicaKeyTagOutput struct{ *pulumi.OutputState }

func (ReplicaKeyTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicaKeyTag)(nil)).Elem()
}

func (o ReplicaKeyTagOutput) ToReplicaKeyTagOutput() ReplicaKeyTagOutput {
	return o
}

func (o ReplicaKeyTagOutput) ToReplicaKeyTagOutputWithContext(ctx context.Context) ReplicaKeyTagOutput {
	return o
}

// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o ReplicaKeyTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v ReplicaKeyTag) string { return v.Key }).(pulumi.StringOutput)
}

// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o ReplicaKeyTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ReplicaKeyTag) string { return v.Value }).(pulumi.StringOutput)
}

type ReplicaKeyTagArrayOutput struct{ *pulumi.OutputState }

func (ReplicaKeyTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReplicaKeyTag)(nil)).Elem()
}

func (o ReplicaKeyTagArrayOutput) ToReplicaKeyTagArrayOutput() ReplicaKeyTagArrayOutput {
	return o
}

func (o ReplicaKeyTagArrayOutput) ToReplicaKeyTagArrayOutputWithContext(ctx context.Context) ReplicaKeyTagArrayOutput {
	return o
}

func (o ReplicaKeyTagArrayOutput) Index(i pulumi.IntInput) ReplicaKeyTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReplicaKeyTag {
		return vs[0].([]ReplicaKeyTag)[vs[1].(int)]
	}).(ReplicaKeyTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeyTagInput)(nil)).Elem(), KeyTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyTagArrayInput)(nil)).Elem(), KeyTagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicaKeyTagInput)(nil)).Elem(), ReplicaKeyTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicaKeyTagArrayInput)(nil)).Elem(), ReplicaKeyTagArray{})
	pulumi.RegisterOutputType(KeyTagOutput{})
	pulumi.RegisterOutputType(KeyTagArrayOutput{})
	pulumi.RegisterOutputType(ReplicaKeyTagOutput{})
	pulumi.RegisterOutputType(ReplicaKeyTagArrayOutput{})
}
