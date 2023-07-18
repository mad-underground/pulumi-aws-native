// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticbeanstalk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::ElasticBeanstalk::Application resource specifies an Elastic Beanstalk application.
type Application struct {
	pulumi.CustomResourceState

	// A name for the Elastic Beanstalk application. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name.
	ApplicationName pulumi.StringPtrOutput `pulumi:"applicationName"`
	// Your description of the application.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies an application resource lifecycle configuration to prevent your application from accumulating too many versions.
	ResourceLifecycleConfig ApplicationResourceLifecycleConfigPtrOutput `pulumi:"resourceLifecycleConfig"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		args = &ApplicationArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Application
	err := ctx.RegisterResource("aws-native:elasticbeanstalk:Application", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws-native:elasticbeanstalk:Application", name, id, state, &resource, opts...)
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
	// A name for the Elastic Beanstalk application. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name.
	ApplicationName *string `pulumi:"applicationName"`
	// Your description of the application.
	Description *string `pulumi:"description"`
	// Specifies an application resource lifecycle configuration to prevent your application from accumulating too many versions.
	ResourceLifecycleConfig *ApplicationResourceLifecycleConfig `pulumi:"resourceLifecycleConfig"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// A name for the Elastic Beanstalk application. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name.
	ApplicationName pulumi.StringPtrInput
	// Your description of the application.
	Description pulumi.StringPtrInput
	// Specifies an application resource lifecycle configuration to prevent your application from accumulating too many versions.
	ResourceLifecycleConfig ApplicationResourceLifecycleConfigPtrInput
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

// A name for the Elastic Beanstalk application. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name.
func (o ApplicationOutput) ApplicationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.ApplicationName }).(pulumi.StringPtrOutput)
}

// Your description of the application.
func (o ApplicationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies an application resource lifecycle configuration to prevent your application from accumulating too many versions.
func (o ApplicationOutput) ResourceLifecycleConfig() ApplicationResourceLifecycleConfigPtrOutput {
	return o.ApplyT(func(v *Application) ApplicationResourceLifecycleConfigPtrOutput { return v.ResourceLifecycleConfig }).(ApplicationResourceLifecycleConfigPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationInput)(nil)).Elem(), &Application{})
	pulumi.RegisterOutputType(ApplicationOutput{})
}
