// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotfleethub

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::IoTFleetHub::Application
type Application struct {
	pulumi.CustomResourceState

	// The ARN of the application.
	ApplicationArn pulumi.StringOutput `pulumi:"applicationArn"`
	// When the Application was created
	ApplicationCreationDate pulumi.IntOutput `pulumi:"applicationCreationDate"`
	// Application Description, should be between 1 and 2048 characters.
	ApplicationDescription pulumi.StringPtrOutput `pulumi:"applicationDescription"`
	// The ID of the application.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// When the Application was last updated
	ApplicationLastUpdateDate pulumi.IntOutput `pulumi:"applicationLastUpdateDate"`
	// Application Name, should be between 1 and 256 characters.
	ApplicationName pulumi.StringOutput `pulumi:"applicationName"`
	// The current state of the application.
	ApplicationState pulumi.StringOutput `pulumi:"applicationState"`
	// The URL of the application.
	ApplicationUrl pulumi.StringOutput `pulumi:"applicationUrl"`
	// A message indicating why Create or Delete Application failed.
	ErrorMessage pulumi.StringOutput `pulumi:"errorMessage"`
	// The ARN of the role that the web application assumes when it interacts with AWS IoT Core. For more info on configuring this attribute, see https://docs.aws.amazon.com/iot/latest/apireference/API_iotfleethub_CreateApplication.html#API_iotfleethub_CreateApplication_RequestSyntax
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// The AWS SSO application generated client ID (used with AWS SSO APIs).
	SsoClientId pulumi.StringOutput `pulumi:"ssoClientId"`
	// A list of key-value pairs that contain metadata for the application.
	Tags ApplicationTagArrayOutput `pulumi:"tags"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Application
	err := ctx.RegisterResource("aws-native:iotfleethub:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("aws-native:iotfleethub:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type applicationState struct {
}

type ApplicationState struct {
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	// Application Description, should be between 1 and 2048 characters.
	ApplicationDescription *string `pulumi:"applicationDescription"`
	// Application Name, should be between 1 and 256 characters.
	ApplicationName *string `pulumi:"applicationName"`
	// The ARN of the role that the web application assumes when it interacts with AWS IoT Core. For more info on configuring this attribute, see https://docs.aws.amazon.com/iot/latest/apireference/API_iotfleethub_CreateApplication.html#API_iotfleethub_CreateApplication_RequestSyntax
	RoleArn string `pulumi:"roleArn"`
	// A list of key-value pairs that contain metadata for the application.
	Tags []ApplicationTag `pulumi:"tags"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// Application Description, should be between 1 and 2048 characters.
	ApplicationDescription pulumi.StringPtrInput
	// Application Name, should be between 1 and 256 characters.
	ApplicationName pulumi.StringPtrInput
	// The ARN of the role that the web application assumes when it interacts with AWS IoT Core. For more info on configuring this attribute, see https://docs.aws.amazon.com/iot/latest/apireference/API_iotfleethub_CreateApplication.html#API_iotfleethub_CreateApplication_RequestSyntax
	RoleArn pulumi.StringInput
	// A list of key-value pairs that contain metadata for the application.
	Tags ApplicationTagArrayInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}

type ApplicationInput interface {
	pulumi.Input

	ToApplicationOutput() ApplicationOutput
	ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput
}

func (*Application) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (i *Application) ToApplicationOutput() ApplicationOutput {
	return i.ToApplicationOutputWithContext(context.Background())
}

func (i *Application) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOutput)
}

type ApplicationOutput struct{ *pulumi.OutputState }

func (ApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (o ApplicationOutput) ToApplicationOutput() ApplicationOutput {
	return o
}

func (o ApplicationOutput) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return o
}

// The ARN of the application.
func (o ApplicationOutput) ApplicationArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.ApplicationArn }).(pulumi.StringOutput)
}

// When the Application was created
func (o ApplicationOutput) ApplicationCreationDate() pulumi.IntOutput {
	return o.ApplyT(func(v *Application) pulumi.IntOutput { return v.ApplicationCreationDate }).(pulumi.IntOutput)
}

// Application Description, should be between 1 and 2048 characters.
func (o ApplicationOutput) ApplicationDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.ApplicationDescription }).(pulumi.StringPtrOutput)
}

// The ID of the application.
func (o ApplicationOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

// When the Application was last updated
func (o ApplicationOutput) ApplicationLastUpdateDate() pulumi.IntOutput {
	return o.ApplyT(func(v *Application) pulumi.IntOutput { return v.ApplicationLastUpdateDate }).(pulumi.IntOutput)
}

// Application Name, should be between 1 and 256 characters.
func (o ApplicationOutput) ApplicationName() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.ApplicationName }).(pulumi.StringOutput)
}

// The current state of the application.
func (o ApplicationOutput) ApplicationState() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.ApplicationState }).(pulumi.StringOutput)
}

// The URL of the application.
func (o ApplicationOutput) ApplicationUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.ApplicationUrl }).(pulumi.StringOutput)
}

// A message indicating why Create or Delete Application failed.
func (o ApplicationOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.ErrorMessage }).(pulumi.StringOutput)
}

// The ARN of the role that the web application assumes when it interacts with AWS IoT Core. For more info on configuring this attribute, see https://docs.aws.amazon.com/iot/latest/apireference/API_iotfleethub_CreateApplication.html#API_iotfleethub_CreateApplication_RequestSyntax
func (o ApplicationOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

// The AWS SSO application generated client ID (used with AWS SSO APIs).
func (o ApplicationOutput) SsoClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.SsoClientId }).(pulumi.StringOutput)
}

// A list of key-value pairs that contain metadata for the application.
func (o ApplicationOutput) Tags() ApplicationTagArrayOutput {
	return o.ApplyT(func(v *Application) ApplicationTagArrayOutput { return v.Tags }).(ApplicationTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationInput)(nil)).Elem(), &Application{})
	pulumi.RegisterOutputType(ApplicationOutput{})
}
