// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package imagebuilder

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::ImageBuilder::DistributionConfiguration
type DistributionConfiguration struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the distribution configuration.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of the distribution configuration.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The distributions of the distribution configuration.
	Distributions DistributionConfigurationDistributionArrayOutput `pulumi:"distributions"`
	// The name of the distribution configuration.
	Name pulumi.StringOutput `pulumi:"name"`
	// The tags associated with the component.
	Tags pulumi.AnyOutput `pulumi:"tags"`
}

// NewDistributionConfiguration registers a new resource with the given unique name, arguments, and options.
func NewDistributionConfiguration(ctx *pulumi.Context,
	name string, args *DistributionConfigurationArgs, opts ...pulumi.ResourceOption) (*DistributionConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Distributions == nil {
		return nil, errors.New("invalid value for required argument 'Distributions'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DistributionConfiguration
	err := ctx.RegisterResource("aws-native:imagebuilder:DistributionConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDistributionConfiguration gets an existing DistributionConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDistributionConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DistributionConfigurationState, opts ...pulumi.ResourceOption) (*DistributionConfiguration, error) {
	var resource DistributionConfiguration
	err := ctx.ReadResource("aws-native:imagebuilder:DistributionConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DistributionConfiguration resources.
type distributionConfigurationState struct {
}

type DistributionConfigurationState struct {
}

func (DistributionConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*distributionConfigurationState)(nil)).Elem()
}

type distributionConfigurationArgs struct {
	// The description of the distribution configuration.
	Description *string `pulumi:"description"`
	// The distributions of the distribution configuration.
	Distributions []DistributionConfigurationDistribution `pulumi:"distributions"`
	// The name of the distribution configuration.
	Name *string `pulumi:"name"`
	// The tags associated with the component.
	Tags interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a DistributionConfiguration resource.
type DistributionConfigurationArgs struct {
	// The description of the distribution configuration.
	Description pulumi.StringPtrInput
	// The distributions of the distribution configuration.
	Distributions DistributionConfigurationDistributionArrayInput
	// The name of the distribution configuration.
	Name pulumi.StringPtrInput
	// The tags associated with the component.
	Tags pulumi.Input
}

func (DistributionConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*distributionConfigurationArgs)(nil)).Elem()
}

type DistributionConfigurationInput interface {
	pulumi.Input

	ToDistributionConfigurationOutput() DistributionConfigurationOutput
	ToDistributionConfigurationOutputWithContext(ctx context.Context) DistributionConfigurationOutput
}

func (*DistributionConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**DistributionConfiguration)(nil)).Elem()
}

func (i *DistributionConfiguration) ToDistributionConfigurationOutput() DistributionConfigurationOutput {
	return i.ToDistributionConfigurationOutputWithContext(context.Background())
}

func (i *DistributionConfiguration) ToDistributionConfigurationOutputWithContext(ctx context.Context) DistributionConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DistributionConfigurationOutput)
}

type DistributionConfigurationOutput struct{ *pulumi.OutputState }

func (DistributionConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DistributionConfiguration)(nil)).Elem()
}

func (o DistributionConfigurationOutput) ToDistributionConfigurationOutput() DistributionConfigurationOutput {
	return o
}

func (o DistributionConfigurationOutput) ToDistributionConfigurationOutputWithContext(ctx context.Context) DistributionConfigurationOutput {
	return o
}

// The Amazon Resource Name (ARN) of the distribution configuration.
func (o DistributionConfigurationOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributionConfiguration) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The description of the distribution configuration.
func (o DistributionConfigurationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DistributionConfiguration) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The distributions of the distribution configuration.
func (o DistributionConfigurationOutput) Distributions() DistributionConfigurationDistributionArrayOutput {
	return o.ApplyT(func(v *DistributionConfiguration) DistributionConfigurationDistributionArrayOutput {
		return v.Distributions
	}).(DistributionConfigurationDistributionArrayOutput)
}

// The name of the distribution configuration.
func (o DistributionConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributionConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The tags associated with the component.
func (o DistributionConfigurationOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v *DistributionConfiguration) pulumi.AnyOutput { return v.Tags }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DistributionConfigurationInput)(nil)).Elem(), &DistributionConfiguration{})
	pulumi.RegisterOutputType(DistributionConfigurationOutput{})
}
