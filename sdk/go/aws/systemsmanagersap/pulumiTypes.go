// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package systemsmanagersap

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type ApplicationCredential struct {
	CredentialType *ApplicationCredentialCredentialType `pulumi:"credentialType"`
	DatabaseName   *string                              `pulumi:"databaseName"`
	SecretId       *string                              `pulumi:"secretId"`
}

// ApplicationCredentialInput is an input type that accepts ApplicationCredentialArgs and ApplicationCredentialOutput values.
// You can construct a concrete instance of `ApplicationCredentialInput` via:
//
//	ApplicationCredentialArgs{...}
type ApplicationCredentialInput interface {
	pulumi.Input

	ToApplicationCredentialOutput() ApplicationCredentialOutput
	ToApplicationCredentialOutputWithContext(context.Context) ApplicationCredentialOutput
}

type ApplicationCredentialArgs struct {
	CredentialType ApplicationCredentialCredentialTypePtrInput `pulumi:"credentialType"`
	DatabaseName   pulumi.StringPtrInput                       `pulumi:"databaseName"`
	SecretId       pulumi.StringPtrInput                       `pulumi:"secretId"`
}

func (ApplicationCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationCredential)(nil)).Elem()
}

func (i ApplicationCredentialArgs) ToApplicationCredentialOutput() ApplicationCredentialOutput {
	return i.ToApplicationCredentialOutputWithContext(context.Background())
}

func (i ApplicationCredentialArgs) ToApplicationCredentialOutputWithContext(ctx context.Context) ApplicationCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationCredentialOutput)
}

// ApplicationCredentialArrayInput is an input type that accepts ApplicationCredentialArray and ApplicationCredentialArrayOutput values.
// You can construct a concrete instance of `ApplicationCredentialArrayInput` via:
//
//	ApplicationCredentialArray{ ApplicationCredentialArgs{...} }
type ApplicationCredentialArrayInput interface {
	pulumi.Input

	ToApplicationCredentialArrayOutput() ApplicationCredentialArrayOutput
	ToApplicationCredentialArrayOutputWithContext(context.Context) ApplicationCredentialArrayOutput
}

type ApplicationCredentialArray []ApplicationCredentialInput

func (ApplicationCredentialArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationCredential)(nil)).Elem()
}

func (i ApplicationCredentialArray) ToApplicationCredentialArrayOutput() ApplicationCredentialArrayOutput {
	return i.ToApplicationCredentialArrayOutputWithContext(context.Background())
}

func (i ApplicationCredentialArray) ToApplicationCredentialArrayOutputWithContext(ctx context.Context) ApplicationCredentialArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationCredentialArrayOutput)
}

type ApplicationCredentialOutput struct{ *pulumi.OutputState }

func (ApplicationCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationCredential)(nil)).Elem()
}

func (o ApplicationCredentialOutput) ToApplicationCredentialOutput() ApplicationCredentialOutput {
	return o
}

func (o ApplicationCredentialOutput) ToApplicationCredentialOutputWithContext(ctx context.Context) ApplicationCredentialOutput {
	return o
}

func (o ApplicationCredentialOutput) CredentialType() ApplicationCredentialCredentialTypePtrOutput {
	return o.ApplyT(func(v ApplicationCredential) *ApplicationCredentialCredentialType { return v.CredentialType }).(ApplicationCredentialCredentialTypePtrOutput)
}

func (o ApplicationCredentialOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationCredential) *string { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

func (o ApplicationCredentialOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationCredential) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

type ApplicationCredentialArrayOutput struct{ *pulumi.OutputState }

func (ApplicationCredentialArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationCredential)(nil)).Elem()
}

func (o ApplicationCredentialArrayOutput) ToApplicationCredentialArrayOutput() ApplicationCredentialArrayOutput {
	return o
}

func (o ApplicationCredentialArrayOutput) ToApplicationCredentialArrayOutputWithContext(ctx context.Context) ApplicationCredentialArrayOutput {
	return o
}

func (o ApplicationCredentialArrayOutput) Index(i pulumi.IntInput) ApplicationCredentialOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationCredential {
		return vs[0].([]ApplicationCredential)[vs[1].(int)]
	}).(ApplicationCredentialOutput)
}

// A key-value pair to associate with a resource.
type ApplicationTag struct {
	// The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key string `pulumi:"key"`
	// The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value string `pulumi:"value"`
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationCredentialInput)(nil)).Elem(), ApplicationCredentialArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationCredentialArrayInput)(nil)).Elem(), ApplicationCredentialArray{})
	pulumi.RegisterOutputType(ApplicationCredentialOutput{})
	pulumi.RegisterOutputType(ApplicationCredentialArrayOutput{})
}
