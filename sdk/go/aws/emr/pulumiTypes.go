// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package emr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An arbitrary set of tags (key-value pairs) for this EMR Studio.
type StudioTag struct {
	// The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key string `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value string `pulumi:"value"`
}

// StudioTagInput is an input type that accepts StudioTagArgs and StudioTagOutput values.
// You can construct a concrete instance of `StudioTagInput` via:
//
//          StudioTagArgs{...}
type StudioTagInput interface {
	pulumi.Input

	ToStudioTagOutput() StudioTagOutput
	ToStudioTagOutputWithContext(context.Context) StudioTagOutput
}

// An arbitrary set of tags (key-value pairs) for this EMR Studio.
type StudioTagArgs struct {
	// The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key pulumi.StringInput `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value pulumi.StringInput `pulumi:"value"`
}

func (StudioTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StudioTag)(nil)).Elem()
}

func (i StudioTagArgs) ToStudioTagOutput() StudioTagOutput {
	return i.ToStudioTagOutputWithContext(context.Background())
}

func (i StudioTagArgs) ToStudioTagOutputWithContext(ctx context.Context) StudioTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StudioTagOutput)
}

// StudioTagArrayInput is an input type that accepts StudioTagArray and StudioTagArrayOutput values.
// You can construct a concrete instance of `StudioTagArrayInput` via:
//
//          StudioTagArray{ StudioTagArgs{...} }
type StudioTagArrayInput interface {
	pulumi.Input

	ToStudioTagArrayOutput() StudioTagArrayOutput
	ToStudioTagArrayOutputWithContext(context.Context) StudioTagArrayOutput
}

type StudioTagArray []StudioTagInput

func (StudioTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StudioTag)(nil)).Elem()
}

func (i StudioTagArray) ToStudioTagArrayOutput() StudioTagArrayOutput {
	return i.ToStudioTagArrayOutputWithContext(context.Background())
}

func (i StudioTagArray) ToStudioTagArrayOutputWithContext(ctx context.Context) StudioTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StudioTagArrayOutput)
}

// An arbitrary set of tags (key-value pairs) for this EMR Studio.
type StudioTagOutput struct{ *pulumi.OutputState }

func (StudioTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StudioTag)(nil)).Elem()
}

func (o StudioTagOutput) ToStudioTagOutput() StudioTagOutput {
	return o
}

func (o StudioTagOutput) ToStudioTagOutputWithContext(ctx context.Context) StudioTagOutput {
	return o
}

// The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o StudioTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v StudioTag) string { return v.Key }).(pulumi.StringOutput)
}

// The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o StudioTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v StudioTag) string { return v.Value }).(pulumi.StringOutput)
}

type StudioTagArrayOutput struct{ *pulumi.OutputState }

func (StudioTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StudioTag)(nil)).Elem()
}

func (o StudioTagArrayOutput) ToStudioTagArrayOutput() StudioTagArrayOutput {
	return o
}

func (o StudioTagArrayOutput) ToStudioTagArrayOutputWithContext(ctx context.Context) StudioTagArrayOutput {
	return o
}

func (o StudioTagArrayOutput) Index(i pulumi.IntInput) StudioTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StudioTag {
		return vs[0].([]StudioTag)[vs[1].(int)]
	}).(StudioTagOutput)
}

func init() {
	pulumi.RegisterOutputType(StudioTagOutput{})
	pulumi.RegisterOutputType(StudioTagArrayOutput{})
}