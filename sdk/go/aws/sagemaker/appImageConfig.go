// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SageMaker::AppImageConfig
type AppImageConfig struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the AppImageConfig.
	AppImageConfigArn pulumi.StringOutput `pulumi:"appImageConfigArn"`
	// The Name of the AppImageConfig.
	AppImageConfigName pulumi.StringOutput `pulumi:"appImageConfigName"`
	// The KernelGatewayImageConfig.
	KernelGatewayImageConfig AppImageConfigKernelGatewayImageConfigPtrOutput `pulumi:"kernelGatewayImageConfig"`
	// A list of tags to apply to the AppImageConfig.
	Tags AppImageConfigTagArrayOutput `pulumi:"tags"`
}

// NewAppImageConfig registers a new resource with the given unique name, arguments, and options.
func NewAppImageConfig(ctx *pulumi.Context,
	name string, args *AppImageConfigArgs, opts ...pulumi.ResourceOption) (*AppImageConfig, error) {
	if args == nil {
		args = &AppImageConfigArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AppImageConfig
	err := ctx.RegisterResource("aws-native:sagemaker:AppImageConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppImageConfig gets an existing AppImageConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppImageConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppImageConfigState, opts ...pulumi.ResourceOption) (*AppImageConfig, error) {
	var resource AppImageConfig
	err := ctx.ReadResource("aws-native:sagemaker:AppImageConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppImageConfig resources.
type appImageConfigState struct {
}

type AppImageConfigState struct {
}

func (AppImageConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*appImageConfigState)(nil)).Elem()
}

type appImageConfigArgs struct {
	// The Name of the AppImageConfig.
	AppImageConfigName *string `pulumi:"appImageConfigName"`
	// The KernelGatewayImageConfig.
	KernelGatewayImageConfig *AppImageConfigKernelGatewayImageConfig `pulumi:"kernelGatewayImageConfig"`
	// A list of tags to apply to the AppImageConfig.
	Tags []AppImageConfigTag `pulumi:"tags"`
}

// The set of arguments for constructing a AppImageConfig resource.
type AppImageConfigArgs struct {
	// The Name of the AppImageConfig.
	AppImageConfigName pulumi.StringPtrInput
	// The KernelGatewayImageConfig.
	KernelGatewayImageConfig AppImageConfigKernelGatewayImageConfigPtrInput
	// A list of tags to apply to the AppImageConfig.
	Tags AppImageConfigTagArrayInput
}

func (AppImageConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appImageConfigArgs)(nil)).Elem()
}

type AppImageConfigInput interface {
	pulumi.Input

	ToAppImageConfigOutput() AppImageConfigOutput
	ToAppImageConfigOutputWithContext(ctx context.Context) AppImageConfigOutput
}

func (*AppImageConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**AppImageConfig)(nil)).Elem()
}

func (i *AppImageConfig) ToAppImageConfigOutput() AppImageConfigOutput {
	return i.ToAppImageConfigOutputWithContext(context.Background())
}

func (i *AppImageConfig) ToAppImageConfigOutputWithContext(ctx context.Context) AppImageConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppImageConfigOutput)
}

type AppImageConfigOutput struct{ *pulumi.OutputState }

func (AppImageConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppImageConfig)(nil)).Elem()
}

func (o AppImageConfigOutput) ToAppImageConfigOutput() AppImageConfigOutput {
	return o
}

func (o AppImageConfigOutput) ToAppImageConfigOutputWithContext(ctx context.Context) AppImageConfigOutput {
	return o
}

// The Amazon Resource Name (ARN) of the AppImageConfig.
func (o AppImageConfigOutput) AppImageConfigArn() pulumi.StringOutput {
	return o.ApplyT(func(v *AppImageConfig) pulumi.StringOutput { return v.AppImageConfigArn }).(pulumi.StringOutput)
}

// The Name of the AppImageConfig.
func (o AppImageConfigOutput) AppImageConfigName() pulumi.StringOutput {
	return o.ApplyT(func(v *AppImageConfig) pulumi.StringOutput { return v.AppImageConfigName }).(pulumi.StringOutput)
}

// The KernelGatewayImageConfig.
func (o AppImageConfigOutput) KernelGatewayImageConfig() AppImageConfigKernelGatewayImageConfigPtrOutput {
	return o.ApplyT(func(v *AppImageConfig) AppImageConfigKernelGatewayImageConfigPtrOutput {
		return v.KernelGatewayImageConfig
	}).(AppImageConfigKernelGatewayImageConfigPtrOutput)
}

// A list of tags to apply to the AppImageConfig.
func (o AppImageConfigOutput) Tags() AppImageConfigTagArrayOutput {
	return o.ApplyT(func(v *AppImageConfig) AppImageConfigTagArrayOutput { return v.Tags }).(AppImageConfigTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppImageConfigInput)(nil)).Elem(), &AppImageConfig{})
	pulumi.RegisterOutputType(AppImageConfigOutput{})
}
