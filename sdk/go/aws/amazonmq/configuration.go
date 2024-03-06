// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package amazonmq

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AmazonMQ::Configuration
//
// Deprecated: Configuration is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Configuration struct {
	pulumi.CustomResourceState

	Arn                    pulumi.StringOutput    `pulumi:"arn"`
	AuthenticationStrategy pulumi.StringPtrOutput `pulumi:"authenticationStrategy"`
	AwsId                  pulumi.StringOutput    `pulumi:"awsId"`
	Data                   pulumi.StringOutput    `pulumi:"data"`
	Description            pulumi.StringPtrOutput `pulumi:"description"`
	EngineType             pulumi.StringOutput    `pulumi:"engineType"`
	EngineVersion          pulumi.StringOutput    `pulumi:"engineVersion"`
	Name                   pulumi.StringOutput    `pulumi:"name"`
	Revision               pulumi.IntOutput       `pulumi:"revision"`
	Tags                   aws.TagArrayOutput     `pulumi:"tags"`
}

// NewConfiguration registers a new resource with the given unique name, arguments, and options.
func NewConfiguration(ctx *pulumi.Context,
	name string, args *ConfigurationArgs, opts ...pulumi.ResourceOption) (*Configuration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Data == nil {
		return nil, errors.New("invalid value for required argument 'Data'")
	}
	if args.EngineType == nil {
		return nil, errors.New("invalid value for required argument 'EngineType'")
	}
	if args.EngineVersion == nil {
		return nil, errors.New("invalid value for required argument 'EngineVersion'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"authenticationStrategy",
		"engineType",
		"engineVersion",
		"name",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Configuration
	err := ctx.RegisterResource("aws-native:amazonmq:Configuration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfiguration gets an existing Configuration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationState, opts ...pulumi.ResourceOption) (*Configuration, error) {
	var resource Configuration
	err := ctx.ReadResource("aws-native:amazonmq:Configuration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Configuration resources.
type configurationState struct {
}

type ConfigurationState struct {
}

func (ConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationState)(nil)).Elem()
}

type configurationArgs struct {
	AuthenticationStrategy *string   `pulumi:"authenticationStrategy"`
	Data                   string    `pulumi:"data"`
	Description            *string   `pulumi:"description"`
	EngineType             string    `pulumi:"engineType"`
	EngineVersion          string    `pulumi:"engineVersion"`
	Name                   *string   `pulumi:"name"`
	Tags                   []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a Configuration resource.
type ConfigurationArgs struct {
	AuthenticationStrategy pulumi.StringPtrInput
	Data                   pulumi.StringInput
	Description            pulumi.StringPtrInput
	EngineType             pulumi.StringInput
	EngineVersion          pulumi.StringInput
	Name                   pulumi.StringPtrInput
	Tags                   aws.TagArrayInput
}

func (ConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationArgs)(nil)).Elem()
}

type ConfigurationInput interface {
	pulumi.Input

	ToConfigurationOutput() ConfigurationOutput
	ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput
}

func (*Configuration) ElementType() reflect.Type {
	return reflect.TypeOf((**Configuration)(nil)).Elem()
}

func (i *Configuration) ToConfigurationOutput() ConfigurationOutput {
	return i.ToConfigurationOutputWithContext(context.Background())
}

func (i *Configuration) ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationOutput)
}

type ConfigurationOutput struct{ *pulumi.OutputState }

func (ConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Configuration)(nil)).Elem()
}

func (o ConfigurationOutput) ToConfigurationOutput() ConfigurationOutput {
	return o
}

func (o ConfigurationOutput) ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput {
	return o
}

func (o ConfigurationOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o ConfigurationOutput) AuthenticationStrategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringPtrOutput { return v.AuthenticationStrategy }).(pulumi.StringPtrOutput)
}

func (o ConfigurationOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

func (o ConfigurationOutput) Data() pulumi.StringOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringOutput { return v.Data }).(pulumi.StringOutput)
}

func (o ConfigurationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ConfigurationOutput) EngineType() pulumi.StringOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringOutput { return v.EngineType }).(pulumi.StringOutput)
}

func (o ConfigurationOutput) EngineVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringOutput { return v.EngineVersion }).(pulumi.StringOutput)
}

func (o ConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConfigurationOutput) Revision() pulumi.IntOutput {
	return o.ApplyT(func(v *Configuration) pulumi.IntOutput { return v.Revision }).(pulumi.IntOutput)
}

func (o ConfigurationOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *Configuration) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationInput)(nil)).Elem(), &Configuration{})
	pulumi.RegisterOutputType(ConfigurationOutput{})
}
