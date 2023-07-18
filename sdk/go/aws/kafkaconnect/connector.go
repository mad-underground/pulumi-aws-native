// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kafkaconnect

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::KafkaConnect::Connector
type Connector struct {
	pulumi.CustomResourceState

	Capacity ConnectorCapacityOutput `pulumi:"capacity"`
	// Amazon Resource Name for the created Connector.
	ConnectorArn pulumi.StringOutput `pulumi:"connectorArn"`
	// The configuration for the connector.
	ConnectorConfiguration pulumi.AnyOutput `pulumi:"connectorConfiguration"`
	// A summary description of the connector.
	ConnectorDescription pulumi.StringPtrOutput `pulumi:"connectorDescription"`
	// The name of the connector.
	ConnectorName                    pulumi.StringOutput                             `pulumi:"connectorName"`
	KafkaCluster                     ConnectorKafkaClusterOutput                     `pulumi:"kafkaCluster"`
	KafkaClusterClientAuthentication ConnectorKafkaClusterClientAuthenticationOutput `pulumi:"kafkaClusterClientAuthentication"`
	KafkaClusterEncryptionInTransit  ConnectorKafkaClusterEncryptionInTransitOutput  `pulumi:"kafkaClusterEncryptionInTransit"`
	// The version of Kafka Connect. It has to be compatible with both the Kafka cluster's version and the plugins.
	KafkaConnectVersion pulumi.StringOutput           `pulumi:"kafkaConnectVersion"`
	LogDelivery         ConnectorLogDeliveryPtrOutput `pulumi:"logDelivery"`
	// List of plugins to use with the connector.
	Plugins ConnectorPluginArrayOutput `pulumi:"plugins"`
	// The Amazon Resource Name (ARN) of the IAM role used by the connector to access Amazon S3 objects and other external resources.
	ServiceExecutionRoleArn pulumi.StringOutput                   `pulumi:"serviceExecutionRoleArn"`
	WorkerConfiguration     ConnectorWorkerConfigurationPtrOutput `pulumi:"workerConfiguration"`
}

// NewConnector registers a new resource with the given unique name, arguments, and options.
func NewConnector(ctx *pulumi.Context,
	name string, args *ConnectorArgs, opts ...pulumi.ResourceOption) (*Connector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Capacity == nil {
		return nil, errors.New("invalid value for required argument 'Capacity'")
	}
	if args.ConnectorConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'ConnectorConfiguration'")
	}
	if args.KafkaCluster == nil {
		return nil, errors.New("invalid value for required argument 'KafkaCluster'")
	}
	if args.KafkaClusterClientAuthentication == nil {
		return nil, errors.New("invalid value for required argument 'KafkaClusterClientAuthentication'")
	}
	if args.KafkaClusterEncryptionInTransit == nil {
		return nil, errors.New("invalid value for required argument 'KafkaClusterEncryptionInTransit'")
	}
	if args.KafkaConnectVersion == nil {
		return nil, errors.New("invalid value for required argument 'KafkaConnectVersion'")
	}
	if args.Plugins == nil {
		return nil, errors.New("invalid value for required argument 'Plugins'")
	}
	if args.ServiceExecutionRoleArn == nil {
		return nil, errors.New("invalid value for required argument 'ServiceExecutionRoleArn'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Connector
	err := ctx.RegisterResource("aws-native:kafkaconnect:Connector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnector gets an existing Connector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectorState, opts ...pulumi.ResourceOption) (*Connector, error) {
	var resource Connector
	err := ctx.ReadResource("aws-native:kafkaconnect:Connector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connector resources.
type connectorState struct {
}

type ConnectorState struct {
}

func (ConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectorState)(nil)).Elem()
}

type connectorArgs struct {
	Capacity ConnectorCapacity `pulumi:"capacity"`
	// The configuration for the connector.
	ConnectorConfiguration interface{} `pulumi:"connectorConfiguration"`
	// A summary description of the connector.
	ConnectorDescription *string `pulumi:"connectorDescription"`
	// The name of the connector.
	ConnectorName                    *string                                   `pulumi:"connectorName"`
	KafkaCluster                     ConnectorKafkaCluster                     `pulumi:"kafkaCluster"`
	KafkaClusterClientAuthentication ConnectorKafkaClusterClientAuthentication `pulumi:"kafkaClusterClientAuthentication"`
	KafkaClusterEncryptionInTransit  ConnectorKafkaClusterEncryptionInTransit  `pulumi:"kafkaClusterEncryptionInTransit"`
	// The version of Kafka Connect. It has to be compatible with both the Kafka cluster's version and the plugins.
	KafkaConnectVersion string                `pulumi:"kafkaConnectVersion"`
	LogDelivery         *ConnectorLogDelivery `pulumi:"logDelivery"`
	// List of plugins to use with the connector.
	Plugins []ConnectorPlugin `pulumi:"plugins"`
	// The Amazon Resource Name (ARN) of the IAM role used by the connector to access Amazon S3 objects and other external resources.
	ServiceExecutionRoleArn string                        `pulumi:"serviceExecutionRoleArn"`
	WorkerConfiguration     *ConnectorWorkerConfiguration `pulumi:"workerConfiguration"`
}

// The set of arguments for constructing a Connector resource.
type ConnectorArgs struct {
	Capacity ConnectorCapacityInput
	// The configuration for the connector.
	ConnectorConfiguration pulumi.Input
	// A summary description of the connector.
	ConnectorDescription pulumi.StringPtrInput
	// The name of the connector.
	ConnectorName                    pulumi.StringPtrInput
	KafkaCluster                     ConnectorKafkaClusterInput
	KafkaClusterClientAuthentication ConnectorKafkaClusterClientAuthenticationInput
	KafkaClusterEncryptionInTransit  ConnectorKafkaClusterEncryptionInTransitInput
	// The version of Kafka Connect. It has to be compatible with both the Kafka cluster's version and the plugins.
	KafkaConnectVersion pulumi.StringInput
	LogDelivery         ConnectorLogDeliveryPtrInput
	// List of plugins to use with the connector.
	Plugins ConnectorPluginArrayInput
	// The Amazon Resource Name (ARN) of the IAM role used by the connector to access Amazon S3 objects and other external resources.
	ServiceExecutionRoleArn pulumi.StringInput
	WorkerConfiguration     ConnectorWorkerConfigurationPtrInput
}

func (ConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectorArgs)(nil)).Elem()
}

type ConnectorInput interface {
	pulumi.Input

	ToConnectorOutput() ConnectorOutput
	ToConnectorOutputWithContext(ctx context.Context) ConnectorOutput
}

func (*Connector) ElementType() reflect.Type {
	return reflect.TypeOf((**Connector)(nil)).Elem()
}

func (i *Connector) ToConnectorOutput() ConnectorOutput {
	return i.ToConnectorOutputWithContext(context.Background())
}

func (i *Connector) ToConnectorOutputWithContext(ctx context.Context) ConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorOutput)
}

type ConnectorOutput struct{ *pulumi.OutputState }

func (ConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Connector)(nil)).Elem()
}

func (o ConnectorOutput) ToConnectorOutput() ConnectorOutput {
	return o
}

func (o ConnectorOutput) ToConnectorOutputWithContext(ctx context.Context) ConnectorOutput {
	return o
}

func (o ConnectorOutput) Capacity() ConnectorCapacityOutput {
	return o.ApplyT(func(v *Connector) ConnectorCapacityOutput { return v.Capacity }).(ConnectorCapacityOutput)
}

// Amazon Resource Name for the created Connector.
func (o ConnectorOutput) ConnectorArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Connector) pulumi.StringOutput { return v.ConnectorArn }).(pulumi.StringOutput)
}

// The configuration for the connector.
func (o ConnectorOutput) ConnectorConfiguration() pulumi.AnyOutput {
	return o.ApplyT(func(v *Connector) pulumi.AnyOutput { return v.ConnectorConfiguration }).(pulumi.AnyOutput)
}

// A summary description of the connector.
func (o ConnectorOutput) ConnectorDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connector) pulumi.StringPtrOutput { return v.ConnectorDescription }).(pulumi.StringPtrOutput)
}

// The name of the connector.
func (o ConnectorOutput) ConnectorName() pulumi.StringOutput {
	return o.ApplyT(func(v *Connector) pulumi.StringOutput { return v.ConnectorName }).(pulumi.StringOutput)
}

func (o ConnectorOutput) KafkaCluster() ConnectorKafkaClusterOutput {
	return o.ApplyT(func(v *Connector) ConnectorKafkaClusterOutput { return v.KafkaCluster }).(ConnectorKafkaClusterOutput)
}

func (o ConnectorOutput) KafkaClusterClientAuthentication() ConnectorKafkaClusterClientAuthenticationOutput {
	return o.ApplyT(func(v *Connector) ConnectorKafkaClusterClientAuthenticationOutput {
		return v.KafkaClusterClientAuthentication
	}).(ConnectorKafkaClusterClientAuthenticationOutput)
}

func (o ConnectorOutput) KafkaClusterEncryptionInTransit() ConnectorKafkaClusterEncryptionInTransitOutput {
	return o.ApplyT(func(v *Connector) ConnectorKafkaClusterEncryptionInTransitOutput {
		return v.KafkaClusterEncryptionInTransit
	}).(ConnectorKafkaClusterEncryptionInTransitOutput)
}

// The version of Kafka Connect. It has to be compatible with both the Kafka cluster's version and the plugins.
func (o ConnectorOutput) KafkaConnectVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Connector) pulumi.StringOutput { return v.KafkaConnectVersion }).(pulumi.StringOutput)
}

func (o ConnectorOutput) LogDelivery() ConnectorLogDeliveryPtrOutput {
	return o.ApplyT(func(v *Connector) ConnectorLogDeliveryPtrOutput { return v.LogDelivery }).(ConnectorLogDeliveryPtrOutput)
}

// List of plugins to use with the connector.
func (o ConnectorOutput) Plugins() ConnectorPluginArrayOutput {
	return o.ApplyT(func(v *Connector) ConnectorPluginArrayOutput { return v.Plugins }).(ConnectorPluginArrayOutput)
}

// The Amazon Resource Name (ARN) of the IAM role used by the connector to access Amazon S3 objects and other external resources.
func (o ConnectorOutput) ServiceExecutionRoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Connector) pulumi.StringOutput { return v.ServiceExecutionRoleArn }).(pulumi.StringOutput)
}

func (o ConnectorOutput) WorkerConfiguration() ConnectorWorkerConfigurationPtrOutput {
	return o.ApplyT(func(v *Connector) ConnectorWorkerConfigurationPtrOutput { return v.WorkerConfiguration }).(ConnectorWorkerConfigurationPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectorInput)(nil)).Elem(), &Connector{})
	pulumi.RegisterOutputType(ConnectorOutput{})
}
