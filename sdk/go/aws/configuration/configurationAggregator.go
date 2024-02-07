// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package configuration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Config::ConfigurationAggregator
type ConfigurationAggregator struct {
	pulumi.CustomResourceState

	AccountAggregationSources ConfigurationAggregatorAccountAggregationSourceArrayOutput `pulumi:"accountAggregationSources"`
	// The Amazon Resource Name (ARN) of the aggregator.
	ConfigurationAggregatorArn pulumi.StringOutput `pulumi:"configurationAggregatorArn"`
	// The name of the aggregator.
	ConfigurationAggregatorName   pulumi.StringPtrOutput                                        `pulumi:"configurationAggregatorName"`
	OrganizationAggregationSource ConfigurationAggregatorOrganizationAggregationSourcePtrOutput `pulumi:"organizationAggregationSource"`
	// The tags for the configuration aggregator.
	Tags ConfigurationAggregatorTagArrayOutput `pulumi:"tags"`
}

// NewConfigurationAggregator registers a new resource with the given unique name, arguments, and options.
func NewConfigurationAggregator(ctx *pulumi.Context,
	name string, args *ConfigurationAggregatorArgs, opts ...pulumi.ResourceOption) (*ConfigurationAggregator, error) {
	if args == nil {
		args = &ConfigurationAggregatorArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"configurationAggregatorName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ConfigurationAggregator
	err := ctx.RegisterResource("aws-native:configuration:ConfigurationAggregator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigurationAggregator gets an existing ConfigurationAggregator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigurationAggregator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationAggregatorState, opts ...pulumi.ResourceOption) (*ConfigurationAggregator, error) {
	var resource ConfigurationAggregator
	err := ctx.ReadResource("aws-native:configuration:ConfigurationAggregator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigurationAggregator resources.
type configurationAggregatorState struct {
}

type ConfigurationAggregatorState struct {
}

func (ConfigurationAggregatorState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationAggregatorState)(nil)).Elem()
}

type configurationAggregatorArgs struct {
	AccountAggregationSources []ConfigurationAggregatorAccountAggregationSource `pulumi:"accountAggregationSources"`
	// The name of the aggregator.
	ConfigurationAggregatorName   *string                                               `pulumi:"configurationAggregatorName"`
	OrganizationAggregationSource *ConfigurationAggregatorOrganizationAggregationSource `pulumi:"organizationAggregationSource"`
	// The tags for the configuration aggregator.
	Tags []ConfigurationAggregatorTag `pulumi:"tags"`
}

// The set of arguments for constructing a ConfigurationAggregator resource.
type ConfigurationAggregatorArgs struct {
	AccountAggregationSources ConfigurationAggregatorAccountAggregationSourceArrayInput
	// The name of the aggregator.
	ConfigurationAggregatorName   pulumi.StringPtrInput
	OrganizationAggregationSource ConfigurationAggregatorOrganizationAggregationSourcePtrInput
	// The tags for the configuration aggregator.
	Tags ConfigurationAggregatorTagArrayInput
}

func (ConfigurationAggregatorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationAggregatorArgs)(nil)).Elem()
}

type ConfigurationAggregatorInput interface {
	pulumi.Input

	ToConfigurationAggregatorOutput() ConfigurationAggregatorOutput
	ToConfigurationAggregatorOutputWithContext(ctx context.Context) ConfigurationAggregatorOutput
}

func (*ConfigurationAggregator) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationAggregator)(nil)).Elem()
}

func (i *ConfigurationAggregator) ToConfigurationAggregatorOutput() ConfigurationAggregatorOutput {
	return i.ToConfigurationAggregatorOutputWithContext(context.Background())
}

func (i *ConfigurationAggregator) ToConfigurationAggregatorOutputWithContext(ctx context.Context) ConfigurationAggregatorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationAggregatorOutput)
}

type ConfigurationAggregatorOutput struct{ *pulumi.OutputState }

func (ConfigurationAggregatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationAggregator)(nil)).Elem()
}

func (o ConfigurationAggregatorOutput) ToConfigurationAggregatorOutput() ConfigurationAggregatorOutput {
	return o
}

func (o ConfigurationAggregatorOutput) ToConfigurationAggregatorOutputWithContext(ctx context.Context) ConfigurationAggregatorOutput {
	return o
}

func (o ConfigurationAggregatorOutput) AccountAggregationSources() ConfigurationAggregatorAccountAggregationSourceArrayOutput {
	return o.ApplyT(func(v *ConfigurationAggregator) ConfigurationAggregatorAccountAggregationSourceArrayOutput {
		return v.AccountAggregationSources
	}).(ConfigurationAggregatorAccountAggregationSourceArrayOutput)
}

// The Amazon Resource Name (ARN) of the aggregator.
func (o ConfigurationAggregatorOutput) ConfigurationAggregatorArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationAggregator) pulumi.StringOutput { return v.ConfigurationAggregatorArn }).(pulumi.StringOutput)
}

// The name of the aggregator.
func (o ConfigurationAggregatorOutput) ConfigurationAggregatorName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationAggregator) pulumi.StringPtrOutput { return v.ConfigurationAggregatorName }).(pulumi.StringPtrOutput)
}

func (o ConfigurationAggregatorOutput) OrganizationAggregationSource() ConfigurationAggregatorOrganizationAggregationSourcePtrOutput {
	return o.ApplyT(func(v *ConfigurationAggregator) ConfigurationAggregatorOrganizationAggregationSourcePtrOutput {
		return v.OrganizationAggregationSource
	}).(ConfigurationAggregatorOrganizationAggregationSourcePtrOutput)
}

// The tags for the configuration aggregator.
func (o ConfigurationAggregatorOutput) Tags() ConfigurationAggregatorTagArrayOutput {
	return o.ApplyT(func(v *ConfigurationAggregator) ConfigurationAggregatorTagArrayOutput { return v.Tags }).(ConfigurationAggregatorTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationAggregatorInput)(nil)).Elem(), &ConfigurationAggregator{})
	pulumi.RegisterOutputType(ConfigurationAggregatorOutput{})
}
