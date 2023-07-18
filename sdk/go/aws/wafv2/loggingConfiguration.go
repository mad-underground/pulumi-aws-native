// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A WAFv2 Logging Configuration Resource Provider
type LoggingConfiguration struct {
	pulumi.CustomResourceState

	// The Amazon Resource Names (ARNs) of the logging destinations that you want to associate with the web ACL.
	LogDestinationConfigs pulumi.StringArrayOutput `pulumi:"logDestinationConfigs"`
	// Filtering that specifies which web requests are kept in the logs and which are dropped. You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation.
	LoggingFilter LoggingFilterPropertiesPtrOutput `pulumi:"loggingFilter"`
	// Indicates whether the logging configuration was created by AWS Firewall Manager, as part of an AWS WAF policy configuration. If true, only Firewall Manager can modify or delete the configuration.
	ManagedByFirewallManager pulumi.BoolOutput `pulumi:"managedByFirewallManager"`
	// The parts of the request that you want to keep out of the logs. For example, if you redact the HEADER field, the HEADER field in the firehose will be xxx.
	RedactedFields LoggingConfigurationFieldToMatchArrayOutput `pulumi:"redactedFields"`
	// The Amazon Resource Name (ARN) of the web ACL that you want to associate with LogDestinationConfigs.
	ResourceArn pulumi.StringOutput `pulumi:"resourceArn"`
}

// NewLoggingConfiguration registers a new resource with the given unique name, arguments, and options.
func NewLoggingConfiguration(ctx *pulumi.Context,
	name string, args *LoggingConfigurationArgs, opts ...pulumi.ResourceOption) (*LoggingConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LogDestinationConfigs == nil {
		return nil, errors.New("invalid value for required argument 'LogDestinationConfigs'")
	}
	if args.ResourceArn == nil {
		return nil, errors.New("invalid value for required argument 'ResourceArn'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LoggingConfiguration
	err := ctx.RegisterResource("aws-native:wafv2:LoggingConfiguration", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws-native:wafv2:LoggingConfiguration", name, id, state, &resource, opts...)
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
	// The Amazon Resource Names (ARNs) of the logging destinations that you want to associate with the web ACL.
	LogDestinationConfigs []string `pulumi:"logDestinationConfigs"`
	// Filtering that specifies which web requests are kept in the logs and which are dropped. You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation.
	LoggingFilter *LoggingFilterProperties `pulumi:"loggingFilter"`
	// The parts of the request that you want to keep out of the logs. For example, if you redact the HEADER field, the HEADER field in the firehose will be xxx.
	RedactedFields []LoggingConfigurationFieldToMatch `pulumi:"redactedFields"`
	// The Amazon Resource Name (ARN) of the web ACL that you want to associate with LogDestinationConfigs.
	ResourceArn string `pulumi:"resourceArn"`
}

// The set of arguments for constructing a LoggingConfiguration resource.
type LoggingConfigurationArgs struct {
	// The Amazon Resource Names (ARNs) of the logging destinations that you want to associate with the web ACL.
	LogDestinationConfigs pulumi.StringArrayInput
	// Filtering that specifies which web requests are kept in the logs and which are dropped. You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation.
	LoggingFilter LoggingFilterPropertiesPtrInput
	// The parts of the request that you want to keep out of the logs. For example, if you redact the HEADER field, the HEADER field in the firehose will be xxx.
	RedactedFields LoggingConfigurationFieldToMatchArrayInput
	// The Amazon Resource Name (ARN) of the web ACL that you want to associate with LogDestinationConfigs.
	ResourceArn pulumi.StringInput
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

// The Amazon Resource Names (ARNs) of the logging destinations that you want to associate with the web ACL.
func (o LoggingConfigurationOutput) LogDestinationConfigs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LoggingConfiguration) pulumi.StringArrayOutput { return v.LogDestinationConfigs }).(pulumi.StringArrayOutput)
}

// Filtering that specifies which web requests are kept in the logs and which are dropped. You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation.
func (o LoggingConfigurationOutput) LoggingFilter() LoggingFilterPropertiesPtrOutput {
	return o.ApplyT(func(v *LoggingConfiguration) LoggingFilterPropertiesPtrOutput { return v.LoggingFilter }).(LoggingFilterPropertiesPtrOutput)
}

// Indicates whether the logging configuration was created by AWS Firewall Manager, as part of an AWS WAF policy configuration. If true, only Firewall Manager can modify or delete the configuration.
func (o LoggingConfigurationOutput) ManagedByFirewallManager() pulumi.BoolOutput {
	return o.ApplyT(func(v *LoggingConfiguration) pulumi.BoolOutput { return v.ManagedByFirewallManager }).(pulumi.BoolOutput)
}

// The parts of the request that you want to keep out of the logs. For example, if you redact the HEADER field, the HEADER field in the firehose will be xxx.
func (o LoggingConfigurationOutput) RedactedFields() LoggingConfigurationFieldToMatchArrayOutput {
	return o.ApplyT(func(v *LoggingConfiguration) LoggingConfigurationFieldToMatchArrayOutput { return v.RedactedFields }).(LoggingConfigurationFieldToMatchArrayOutput)
}

// The Amazon Resource Name (ARN) of the web ACL that you want to associate with LogDestinationConfigs.
func (o LoggingConfigurationOutput) ResourceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingConfiguration) pulumi.StringOutput { return v.ResourceArn }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoggingConfigurationInput)(nil)).Elem(), &LoggingConfiguration{})
	pulumi.RegisterOutputType(LoggingConfigurationOutput{})
}
