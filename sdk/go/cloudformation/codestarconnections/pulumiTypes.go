// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codestarconnections

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-cloudformation/sdk/go/cloudformation"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ConnectionAttributes struct {
	ConnectionArn    string `pulumi:"ConnectionArn"`
	ConnectionStatus string `pulumi:"ConnectionStatus"`
	OwnerAccountId   string `pulumi:"OwnerAccountId"`
}

// ConnectionAttributesInput is an input type that accepts ConnectionAttributesArgs and ConnectionAttributesOutput values.
// You can construct a concrete instance of `ConnectionAttributesInput` via:
//
//          ConnectionAttributesArgs{...}
type ConnectionAttributesInput interface {
	pulumi.Input

	ToConnectionAttributesOutput() ConnectionAttributesOutput
	ToConnectionAttributesOutputWithContext(context.Context) ConnectionAttributesOutput
}

type ConnectionAttributesArgs struct {
	ConnectionArn    pulumi.StringInput `pulumi:"ConnectionArn"`
	ConnectionStatus pulumi.StringInput `pulumi:"ConnectionStatus"`
	OwnerAccountId   pulumi.StringInput `pulumi:"OwnerAccountId"`
}

func (ConnectionAttributesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionAttributes)(nil)).Elem()
}

func (i ConnectionAttributesArgs) ToConnectionAttributesOutput() ConnectionAttributesOutput {
	return i.ToConnectionAttributesOutputWithContext(context.Background())
}

func (i ConnectionAttributesArgs) ToConnectionAttributesOutputWithContext(ctx context.Context) ConnectionAttributesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionAttributesOutput)
}

func (i ConnectionAttributesArgs) ToConnectionAttributesPtrOutput() ConnectionAttributesPtrOutput {
	return i.ToConnectionAttributesPtrOutputWithContext(context.Background())
}

func (i ConnectionAttributesArgs) ToConnectionAttributesPtrOutputWithContext(ctx context.Context) ConnectionAttributesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionAttributesOutput).ToConnectionAttributesPtrOutputWithContext(ctx)
}

// ConnectionAttributesPtrInput is an input type that accepts ConnectionAttributesArgs, ConnectionAttributesPtr and ConnectionAttributesPtrOutput values.
// You can construct a concrete instance of `ConnectionAttributesPtrInput` via:
//
//          ConnectionAttributesArgs{...}
//
//  or:
//
//          nil
type ConnectionAttributesPtrInput interface {
	pulumi.Input

	ToConnectionAttributesPtrOutput() ConnectionAttributesPtrOutput
	ToConnectionAttributesPtrOutputWithContext(context.Context) ConnectionAttributesPtrOutput
}

type connectionAttributesPtrType ConnectionAttributesArgs

func ConnectionAttributesPtr(v *ConnectionAttributesArgs) ConnectionAttributesPtrInput {
	return (*connectionAttributesPtrType)(v)
}

func (*connectionAttributesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionAttributes)(nil)).Elem()
}

func (i *connectionAttributesPtrType) ToConnectionAttributesPtrOutput() ConnectionAttributesPtrOutput {
	return i.ToConnectionAttributesPtrOutputWithContext(context.Background())
}

func (i *connectionAttributesPtrType) ToConnectionAttributesPtrOutputWithContext(ctx context.Context) ConnectionAttributesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionAttributesPtrOutput)
}

type ConnectionAttributesOutput struct{ *pulumi.OutputState }

func (ConnectionAttributesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionAttributes)(nil)).Elem()
}

func (o ConnectionAttributesOutput) ToConnectionAttributesOutput() ConnectionAttributesOutput {
	return o
}

func (o ConnectionAttributesOutput) ToConnectionAttributesOutputWithContext(ctx context.Context) ConnectionAttributesOutput {
	return o
}

func (o ConnectionAttributesOutput) ToConnectionAttributesPtrOutput() ConnectionAttributesPtrOutput {
	return o.ToConnectionAttributesPtrOutputWithContext(context.Background())
}

func (o ConnectionAttributesOutput) ToConnectionAttributesPtrOutputWithContext(ctx context.Context) ConnectionAttributesPtrOutput {
	return o.ApplyT(func(v ConnectionAttributes) *ConnectionAttributes {
		return &v
	}).(ConnectionAttributesPtrOutput)
}
func (o ConnectionAttributesOutput) ConnectionArn() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionAttributes) string { return v.ConnectionArn }).(pulumi.StringOutput)
}

func (o ConnectionAttributesOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionAttributes) string { return v.ConnectionStatus }).(pulumi.StringOutput)
}

func (o ConnectionAttributesOutput) OwnerAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionAttributes) string { return v.OwnerAccountId }).(pulumi.StringOutput)
}

type ConnectionAttributesPtrOutput struct{ *pulumi.OutputState }

func (ConnectionAttributesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionAttributes)(nil)).Elem()
}

func (o ConnectionAttributesPtrOutput) ToConnectionAttributesPtrOutput() ConnectionAttributesPtrOutput {
	return o
}

func (o ConnectionAttributesPtrOutput) ToConnectionAttributesPtrOutputWithContext(ctx context.Context) ConnectionAttributesPtrOutput {
	return o
}

func (o ConnectionAttributesPtrOutput) Elem() ConnectionAttributesOutput {
	return o.ApplyT(func(v *ConnectionAttributes) ConnectionAttributes { return *v }).(ConnectionAttributesOutput)
}

func (o ConnectionAttributesPtrOutput) ConnectionArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionAttributes) *string {
		if v == nil {
			return nil
		}
		return &v.ConnectionArn
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionAttributesPtrOutput) ConnectionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionAttributes) *string {
		if v == nil {
			return nil
		}
		return &v.ConnectionStatus
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionAttributesPtrOutput) OwnerAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionAttributes) *string {
		if v == nil {
			return nil
		}
		return &v.OwnerAccountId
	}).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-connection.html
type ConnectionProperties struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-connection.html#cfn-codestarconnections-connection-connectionname
	ConnectionName string `pulumi:"ConnectionName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-connection.html#cfn-codestarconnections-connection-hostarn
	HostArn *string `pulumi:"HostArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-connection.html#cfn-codestarconnections-connection-providertype
	ProviderType *string `pulumi:"ProviderType"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-connection.html#cfn-codestarconnections-connection-tags
	Tags []cloudformation.Tag `pulumi:"Tags"`
}

// ConnectionPropertiesInput is an input type that accepts ConnectionPropertiesArgs and ConnectionPropertiesOutput values.
// You can construct a concrete instance of `ConnectionPropertiesInput` via:
//
//          ConnectionPropertiesArgs{...}
type ConnectionPropertiesInput interface {
	pulumi.Input

	ToConnectionPropertiesOutput() ConnectionPropertiesOutput
	ToConnectionPropertiesOutputWithContext(context.Context) ConnectionPropertiesOutput
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-connection.html
type ConnectionPropertiesArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-connection.html#cfn-codestarconnections-connection-connectionname
	ConnectionName pulumi.StringInput `pulumi:"ConnectionName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-connection.html#cfn-codestarconnections-connection-hostarn
	HostArn pulumi.StringPtrInput `pulumi:"HostArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-connection.html#cfn-codestarconnections-connection-providertype
	ProviderType pulumi.StringPtrInput `pulumi:"ProviderType"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-connection.html#cfn-codestarconnections-connection-tags
	Tags cloudformation.TagArrayInput `pulumi:"Tags"`
}

func (ConnectionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionProperties)(nil)).Elem()
}

func (i ConnectionPropertiesArgs) ToConnectionPropertiesOutput() ConnectionPropertiesOutput {
	return i.ToConnectionPropertiesOutputWithContext(context.Background())
}

func (i ConnectionPropertiesArgs) ToConnectionPropertiesOutputWithContext(ctx context.Context) ConnectionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionPropertiesOutput)
}

func (i ConnectionPropertiesArgs) ToConnectionPropertiesPtrOutput() ConnectionPropertiesPtrOutput {
	return i.ToConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (i ConnectionPropertiesArgs) ToConnectionPropertiesPtrOutputWithContext(ctx context.Context) ConnectionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionPropertiesOutput).ToConnectionPropertiesPtrOutputWithContext(ctx)
}

// ConnectionPropertiesPtrInput is an input type that accepts ConnectionPropertiesArgs, ConnectionPropertiesPtr and ConnectionPropertiesPtrOutput values.
// You can construct a concrete instance of `ConnectionPropertiesPtrInput` via:
//
//          ConnectionPropertiesArgs{...}
//
//  or:
//
//          nil
type ConnectionPropertiesPtrInput interface {
	pulumi.Input

	ToConnectionPropertiesPtrOutput() ConnectionPropertiesPtrOutput
	ToConnectionPropertiesPtrOutputWithContext(context.Context) ConnectionPropertiesPtrOutput
}

type connectionPropertiesPtrType ConnectionPropertiesArgs

func ConnectionPropertiesPtr(v *ConnectionPropertiesArgs) ConnectionPropertiesPtrInput {
	return (*connectionPropertiesPtrType)(v)
}

func (*connectionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionProperties)(nil)).Elem()
}

func (i *connectionPropertiesPtrType) ToConnectionPropertiesPtrOutput() ConnectionPropertiesPtrOutput {
	return i.ToConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (i *connectionPropertiesPtrType) ToConnectionPropertiesPtrOutputWithContext(ctx context.Context) ConnectionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionPropertiesPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-connection.html
type ConnectionPropertiesOutput struct{ *pulumi.OutputState }

func (ConnectionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionProperties)(nil)).Elem()
}

func (o ConnectionPropertiesOutput) ToConnectionPropertiesOutput() ConnectionPropertiesOutput {
	return o
}

func (o ConnectionPropertiesOutput) ToConnectionPropertiesOutputWithContext(ctx context.Context) ConnectionPropertiesOutput {
	return o
}

func (o ConnectionPropertiesOutput) ToConnectionPropertiesPtrOutput() ConnectionPropertiesPtrOutput {
	return o.ToConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (o ConnectionPropertiesOutput) ToConnectionPropertiesPtrOutputWithContext(ctx context.Context) ConnectionPropertiesPtrOutput {
	return o.ApplyT(func(v ConnectionProperties) *ConnectionProperties {
		return &v
	}).(ConnectionPropertiesPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-connection.html#cfn-codestarconnections-connection-connectionname
func (o ConnectionPropertiesOutput) ConnectionName() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionProperties) string { return v.ConnectionName }).(pulumi.StringOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-connection.html#cfn-codestarconnections-connection-hostarn
func (o ConnectionPropertiesOutput) HostArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionProperties) *string { return v.HostArn }).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-connection.html#cfn-codestarconnections-connection-providertype
func (o ConnectionPropertiesOutput) ProviderType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionProperties) *string { return v.ProviderType }).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-connection.html#cfn-codestarconnections-connection-tags
func (o ConnectionPropertiesOutput) Tags() cloudformation.TagArrayOutput {
	return o.ApplyT(func(v ConnectionProperties) []cloudformation.Tag { return v.Tags }).(cloudformation.TagArrayOutput)
}

type ConnectionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ConnectionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionProperties)(nil)).Elem()
}

func (o ConnectionPropertiesPtrOutput) ToConnectionPropertiesPtrOutput() ConnectionPropertiesPtrOutput {
	return o
}

func (o ConnectionPropertiesPtrOutput) ToConnectionPropertiesPtrOutputWithContext(ctx context.Context) ConnectionPropertiesPtrOutput {
	return o
}

func (o ConnectionPropertiesPtrOutput) Elem() ConnectionPropertiesOutput {
	return o.ApplyT(func(v *ConnectionProperties) ConnectionProperties { return *v }).(ConnectionPropertiesOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-connection.html#cfn-codestarconnections-connection-connectionname
func (o ConnectionPropertiesPtrOutput) ConnectionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionProperties) *string {
		if v == nil {
			return nil
		}
		return &v.ConnectionName
	}).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-connection.html#cfn-codestarconnections-connection-hostarn
func (o ConnectionPropertiesPtrOutput) HostArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionProperties) *string {
		if v == nil {
			return nil
		}
		return v.HostArn
	}).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-connection.html#cfn-codestarconnections-connection-providertype
func (o ConnectionPropertiesPtrOutput) ProviderType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionProperties) *string {
		if v == nil {
			return nil
		}
		return v.ProviderType
	}).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-connection.html#cfn-codestarconnections-connection-tags
func (o ConnectionPropertiesPtrOutput) Tags() cloudformation.TagArrayOutput {
	return o.ApplyT(func(v *ConnectionProperties) []cloudformation.Tag {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(cloudformation.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectionAttributesOutput{})
	pulumi.RegisterOutputType(ConnectionAttributesPtrOutput{})
	pulumi.RegisterOutputType(ConnectionPropertiesOutput{})
	pulumi.RegisterOutputType(ConnectionPropertiesPtrOutput{})
}
