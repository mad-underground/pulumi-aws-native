// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SageMaker::App
type App struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the app.
	AppArn pulumi.StringOutput `pulumi:"appArn"`
	// The name of the app.
	AppName pulumi.StringOutput `pulumi:"appName"`
	// The type of app.
	AppType AppTypeOutput `pulumi:"appType"`
	// The domain ID.
	DomainId pulumi.StringOutput `pulumi:"domainId"`
	// The instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance.
	ResourceSpec AppResourceSpecPtrOutput `pulumi:"resourceSpec"`
	// A list of tags to apply to the app.
	Tags AppTagArrayOutput `pulumi:"tags"`
	// The user profile name.
	UserProfileName pulumi.StringOutput `pulumi:"userProfileName"`
}

// NewApp registers a new resource with the given unique name, arguments, and options.
func NewApp(ctx *pulumi.Context,
	name string, args *AppArgs, opts ...pulumi.ResourceOption) (*App, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppType == nil {
		return nil, errors.New("invalid value for required argument 'AppType'")
	}
	if args.DomainId == nil {
		return nil, errors.New("invalid value for required argument 'DomainId'")
	}
	if args.UserProfileName == nil {
		return nil, errors.New("invalid value for required argument 'UserProfileName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource App
	err := ctx.RegisterResource("aws-native:sagemaker:App", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApp gets an existing App resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppState, opts ...pulumi.ResourceOption) (*App, error) {
	var resource App
	err := ctx.ReadResource("aws-native:sagemaker:App", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering App resources.
type appState struct {
}

type AppState struct {
}

func (AppState) ElementType() reflect.Type {
	return reflect.TypeOf((*appState)(nil)).Elem()
}

type appArgs struct {
	// The name of the app.
	AppName *string `pulumi:"appName"`
	// The type of app.
	AppType AppType `pulumi:"appType"`
	// The domain ID.
	DomainId string `pulumi:"domainId"`
	// The instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance.
	ResourceSpec *AppResourceSpec `pulumi:"resourceSpec"`
	// A list of tags to apply to the app.
	Tags []AppTag `pulumi:"tags"`
	// The user profile name.
	UserProfileName string `pulumi:"userProfileName"`
}

// The set of arguments for constructing a App resource.
type AppArgs struct {
	// The name of the app.
	AppName pulumi.StringPtrInput
	// The type of app.
	AppType AppTypeInput
	// The domain ID.
	DomainId pulumi.StringInput
	// The instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance.
	ResourceSpec AppResourceSpecPtrInput
	// A list of tags to apply to the app.
	Tags AppTagArrayInput
	// The user profile name.
	UserProfileName pulumi.StringInput
}

func (AppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appArgs)(nil)).Elem()
}

type AppInput interface {
	pulumi.Input

	ToAppOutput() AppOutput
	ToAppOutputWithContext(ctx context.Context) AppOutput
}

func (*App) ElementType() reflect.Type {
	return reflect.TypeOf((**App)(nil)).Elem()
}

func (i *App) ToAppOutput() AppOutput {
	return i.ToAppOutputWithContext(context.Background())
}

func (i *App) ToAppOutputWithContext(ctx context.Context) AppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppOutput)
}

type AppOutput struct{ *pulumi.OutputState }

func (AppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**App)(nil)).Elem()
}

func (o AppOutput) ToAppOutput() AppOutput {
	return o
}

func (o AppOutput) ToAppOutputWithContext(ctx context.Context) AppOutput {
	return o
}

// The Amazon Resource Name (ARN) of the app.
func (o AppOutput) AppArn() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.AppArn }).(pulumi.StringOutput)
}

// The name of the app.
func (o AppOutput) AppName() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.AppName }).(pulumi.StringOutput)
}

// The type of app.
func (o AppOutput) AppType() AppTypeOutput {
	return o.ApplyT(func(v *App) AppTypeOutput { return v.AppType }).(AppTypeOutput)
}

// The domain ID.
func (o AppOutput) DomainId() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.DomainId }).(pulumi.StringOutput)
}

// The instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance.
func (o AppOutput) ResourceSpec() AppResourceSpecPtrOutput {
	return o.ApplyT(func(v *App) AppResourceSpecPtrOutput { return v.ResourceSpec }).(AppResourceSpecPtrOutput)
}

// A list of tags to apply to the app.
func (o AppOutput) Tags() AppTagArrayOutput {
	return o.ApplyT(func(v *App) AppTagArrayOutput { return v.Tags }).(AppTagArrayOutput)
}

// The user profile name.
func (o AppOutput) UserProfileName() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.UserProfileName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppInput)(nil)).Elem(), &App{})
	pulumi.RegisterOutputType(AppOutput{})
}
