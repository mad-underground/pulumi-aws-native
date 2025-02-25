// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ivschat

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource type definition for AWS::IVSChat::LoggingConfiguration.
type LoggingConfiguration struct {
	pulumi.CustomResourceState

	// LoggingConfiguration ARN is automatically generated on creation and assigned as the unique identifier.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The system-generated ID of the logging configuration.
	AwsId                    pulumi.StringOutput                                `pulumi:"awsId"`
	DestinationConfiguration LoggingConfigurationDestinationConfigurationOutput `pulumi:"destinationConfiguration"`
	// The name of the logging configuration. The value does not need to be unique.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The state of the logging configuration. When the state is ACTIVE, the configuration is ready to log chat content.
	State LoggingConfigurationStateEnumOutput `pulumi:"state"`
	// An array of key-value pairs to apply to this resource.
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewLoggingConfiguration registers a new resource with the given unique name, arguments, and options.
func NewLoggingConfiguration(ctx *pulumi.Context,
	name string, args *LoggingConfigurationArgs, opts ...pulumi.ResourceOption) (*LoggingConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DestinationConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'DestinationConfiguration'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LoggingConfiguration
	err := ctx.RegisterResource("aws-native:ivschat:LoggingConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoggingConfiguration gets an existing LoggingConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoggingConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoggingConfigurationState, opts ...pulumi.ResourceOption) (*LoggingConfiguration, error) {
	var resource LoggingConfiguration
	err := ctx.ReadResource("aws-native:ivschat:LoggingConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoggingConfiguration resources.
type loggingConfigurationState struct {
}

type LoggingConfigurationState struct {
}

func (LoggingConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*loggingConfigurationState)(nil)).Elem()
}

type loggingConfigurationArgs struct {
	DestinationConfiguration LoggingConfigurationDestinationConfiguration `pulumi:"destinationConfiguration"`
	// The name of the logging configuration. The value does not need to be unique.
	Name *string `pulumi:"name"`
	// An array of key-value pairs to apply to this resource.
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a LoggingConfiguration resource.
type LoggingConfigurationArgs struct {
	DestinationConfiguration LoggingConfigurationDestinationConfigurationInput
	// The name of the logging configuration. The value does not need to be unique.
	Name pulumi.StringPtrInput
	// An array of key-value pairs to apply to this resource.
	Tags aws.TagArrayInput
}

func (LoggingConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loggingConfigurationArgs)(nil)).Elem()
}

type LoggingConfigurationInput interface {
	pulumi.Input

	ToLoggingConfigurationOutput() LoggingConfigurationOutput
	ToLoggingConfigurationOutputWithContext(ctx context.Context) LoggingConfigurationOutput
}

func (*LoggingConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggingConfiguration)(nil)).Elem()
}

func (i *LoggingConfiguration) ToLoggingConfigurationOutput() LoggingConfigurationOutput {
	return i.ToLoggingConfigurationOutputWithContext(context.Background())
}

func (i *LoggingConfiguration) ToLoggingConfigurationOutputWithContext(ctx context.Context) LoggingConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggingConfigurationOutput)
}

type LoggingConfigurationOutput struct{ *pulumi.OutputState }

func (LoggingConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggingConfiguration)(nil)).Elem()
}

func (o LoggingConfigurationOutput) ToLoggingConfigurationOutput() LoggingConfigurationOutput {
	return o
}

func (o LoggingConfigurationOutput) ToLoggingConfigurationOutputWithContext(ctx context.Context) LoggingConfigurationOutput {
	return o
}

// LoggingConfiguration ARN is automatically generated on creation and assigned as the unique identifier.
func (o LoggingConfigurationOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingConfiguration) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The system-generated ID of the logging configuration.
func (o LoggingConfigurationOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingConfiguration) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

func (o LoggingConfigurationOutput) DestinationConfiguration() LoggingConfigurationDestinationConfigurationOutput {
	return o.ApplyT(func(v *LoggingConfiguration) LoggingConfigurationDestinationConfigurationOutput {
		return v.DestinationConfiguration
	}).(LoggingConfigurationDestinationConfigurationOutput)
}

// The name of the logging configuration. The value does not need to be unique.
func (o LoggingConfigurationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoggingConfiguration) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The state of the logging configuration. When the state is ACTIVE, the configuration is ready to log chat content.
func (o LoggingConfigurationOutput) State() LoggingConfigurationStateEnumOutput {
	return o.ApplyT(func(v *LoggingConfiguration) LoggingConfigurationStateEnumOutput { return v.State }).(LoggingConfigurationStateEnumOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LoggingConfigurationOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *LoggingConfiguration) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoggingConfigurationInput)(nil)).Elem(), &LoggingConfiguration{})
	pulumi.RegisterOutputType(LoggingConfigurationOutput{})
}
