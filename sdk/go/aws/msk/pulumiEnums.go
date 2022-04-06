// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package msk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ClusterEncryptionInTransitClientBroker string

const (
	ClusterEncryptionInTransitClientBrokerTls          = ClusterEncryptionInTransitClientBroker("TLS")
	ClusterEncryptionInTransitClientBrokerTlsPlaintext = ClusterEncryptionInTransitClientBroker("TLS_PLAINTEXT")
	ClusterEncryptionInTransitClientBrokerPlaintext    = ClusterEncryptionInTransitClientBroker("PLAINTEXT")
)

func (ClusterEncryptionInTransitClientBroker) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterEncryptionInTransitClientBroker)(nil)).Elem()
}

func (e ClusterEncryptionInTransitClientBroker) ToClusterEncryptionInTransitClientBrokerOutput() ClusterEncryptionInTransitClientBrokerOutput {
	return pulumi.ToOutput(e).(ClusterEncryptionInTransitClientBrokerOutput)
}

func (e ClusterEncryptionInTransitClientBroker) ToClusterEncryptionInTransitClientBrokerOutputWithContext(ctx context.Context) ClusterEncryptionInTransitClientBrokerOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ClusterEncryptionInTransitClientBrokerOutput)
}

func (e ClusterEncryptionInTransitClientBroker) ToClusterEncryptionInTransitClientBrokerPtrOutput() ClusterEncryptionInTransitClientBrokerPtrOutput {
	return e.ToClusterEncryptionInTransitClientBrokerPtrOutputWithContext(context.Background())
}

func (e ClusterEncryptionInTransitClientBroker) ToClusterEncryptionInTransitClientBrokerPtrOutputWithContext(ctx context.Context) ClusterEncryptionInTransitClientBrokerPtrOutput {
	return ClusterEncryptionInTransitClientBroker(e).ToClusterEncryptionInTransitClientBrokerOutputWithContext(ctx).ToClusterEncryptionInTransitClientBrokerPtrOutputWithContext(ctx)
}

func (e ClusterEncryptionInTransitClientBroker) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClusterEncryptionInTransitClientBroker) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClusterEncryptionInTransitClientBroker) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ClusterEncryptionInTransitClientBroker) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ClusterEncryptionInTransitClientBrokerOutput struct{ *pulumi.OutputState }

func (ClusterEncryptionInTransitClientBrokerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterEncryptionInTransitClientBroker)(nil)).Elem()
}

func (o ClusterEncryptionInTransitClientBrokerOutput) ToClusterEncryptionInTransitClientBrokerOutput() ClusterEncryptionInTransitClientBrokerOutput {
	return o
}

func (o ClusterEncryptionInTransitClientBrokerOutput) ToClusterEncryptionInTransitClientBrokerOutputWithContext(ctx context.Context) ClusterEncryptionInTransitClientBrokerOutput {
	return o
}

func (o ClusterEncryptionInTransitClientBrokerOutput) ToClusterEncryptionInTransitClientBrokerPtrOutput() ClusterEncryptionInTransitClientBrokerPtrOutput {
	return o.ToClusterEncryptionInTransitClientBrokerPtrOutputWithContext(context.Background())
}

func (o ClusterEncryptionInTransitClientBrokerOutput) ToClusterEncryptionInTransitClientBrokerPtrOutputWithContext(ctx context.Context) ClusterEncryptionInTransitClientBrokerPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterEncryptionInTransitClientBroker) *ClusterEncryptionInTransitClientBroker {
		return &v
	}).(ClusterEncryptionInTransitClientBrokerPtrOutput)
}

func (o ClusterEncryptionInTransitClientBrokerOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ClusterEncryptionInTransitClientBrokerOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClusterEncryptionInTransitClientBroker) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ClusterEncryptionInTransitClientBrokerOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClusterEncryptionInTransitClientBrokerOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClusterEncryptionInTransitClientBroker) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ClusterEncryptionInTransitClientBrokerPtrOutput struct{ *pulumi.OutputState }

func (ClusterEncryptionInTransitClientBrokerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterEncryptionInTransitClientBroker)(nil)).Elem()
}

func (o ClusterEncryptionInTransitClientBrokerPtrOutput) ToClusterEncryptionInTransitClientBrokerPtrOutput() ClusterEncryptionInTransitClientBrokerPtrOutput {
	return o
}

func (o ClusterEncryptionInTransitClientBrokerPtrOutput) ToClusterEncryptionInTransitClientBrokerPtrOutputWithContext(ctx context.Context) ClusterEncryptionInTransitClientBrokerPtrOutput {
	return o
}

func (o ClusterEncryptionInTransitClientBrokerPtrOutput) Elem() ClusterEncryptionInTransitClientBrokerOutput {
	return o.ApplyT(func(v *ClusterEncryptionInTransitClientBroker) ClusterEncryptionInTransitClientBroker {
		if v != nil {
			return *v
		}
		var ret ClusterEncryptionInTransitClientBroker
		return ret
	}).(ClusterEncryptionInTransitClientBrokerOutput)
}

func (o ClusterEncryptionInTransitClientBrokerPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClusterEncryptionInTransitClientBrokerPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ClusterEncryptionInTransitClientBroker) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ClusterEncryptionInTransitClientBrokerInput is an input type that accepts ClusterEncryptionInTransitClientBrokerArgs and ClusterEncryptionInTransitClientBrokerOutput values.
// You can construct a concrete instance of `ClusterEncryptionInTransitClientBrokerInput` via:
//
//          ClusterEncryptionInTransitClientBrokerArgs{...}
type ClusterEncryptionInTransitClientBrokerInput interface {
	pulumi.Input

	ToClusterEncryptionInTransitClientBrokerOutput() ClusterEncryptionInTransitClientBrokerOutput
	ToClusterEncryptionInTransitClientBrokerOutputWithContext(context.Context) ClusterEncryptionInTransitClientBrokerOutput
}

var clusterEncryptionInTransitClientBrokerPtrType = reflect.TypeOf((**ClusterEncryptionInTransitClientBroker)(nil)).Elem()

type ClusterEncryptionInTransitClientBrokerPtrInput interface {
	pulumi.Input

	ToClusterEncryptionInTransitClientBrokerPtrOutput() ClusterEncryptionInTransitClientBrokerPtrOutput
	ToClusterEncryptionInTransitClientBrokerPtrOutputWithContext(context.Context) ClusterEncryptionInTransitClientBrokerPtrOutput
}

type clusterEncryptionInTransitClientBrokerPtr string

func ClusterEncryptionInTransitClientBrokerPtr(v string) ClusterEncryptionInTransitClientBrokerPtrInput {
	return (*clusterEncryptionInTransitClientBrokerPtr)(&v)
}

func (*clusterEncryptionInTransitClientBrokerPtr) ElementType() reflect.Type {
	return clusterEncryptionInTransitClientBrokerPtrType
}

func (in *clusterEncryptionInTransitClientBrokerPtr) ToClusterEncryptionInTransitClientBrokerPtrOutput() ClusterEncryptionInTransitClientBrokerPtrOutput {
	return pulumi.ToOutput(in).(ClusterEncryptionInTransitClientBrokerPtrOutput)
}

func (in *clusterEncryptionInTransitClientBrokerPtr) ToClusterEncryptionInTransitClientBrokerPtrOutputWithContext(ctx context.Context) ClusterEncryptionInTransitClientBrokerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ClusterEncryptionInTransitClientBrokerPtrOutput)
}

type ClusterEnhancedMonitoring string

const (
	ClusterEnhancedMonitoringDefault              = ClusterEnhancedMonitoring("DEFAULT")
	ClusterEnhancedMonitoringPerBroker            = ClusterEnhancedMonitoring("PER_BROKER")
	ClusterEnhancedMonitoringPerTopicPerBroker    = ClusterEnhancedMonitoring("PER_TOPIC_PER_BROKER")
	ClusterEnhancedMonitoringPerTopicPerPartition = ClusterEnhancedMonitoring("PER_TOPIC_PER_PARTITION")
)

func (ClusterEnhancedMonitoring) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterEnhancedMonitoring)(nil)).Elem()
}

func (e ClusterEnhancedMonitoring) ToClusterEnhancedMonitoringOutput() ClusterEnhancedMonitoringOutput {
	return pulumi.ToOutput(e).(ClusterEnhancedMonitoringOutput)
}

func (e ClusterEnhancedMonitoring) ToClusterEnhancedMonitoringOutputWithContext(ctx context.Context) ClusterEnhancedMonitoringOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ClusterEnhancedMonitoringOutput)
}

func (e ClusterEnhancedMonitoring) ToClusterEnhancedMonitoringPtrOutput() ClusterEnhancedMonitoringPtrOutput {
	return e.ToClusterEnhancedMonitoringPtrOutputWithContext(context.Background())
}

func (e ClusterEnhancedMonitoring) ToClusterEnhancedMonitoringPtrOutputWithContext(ctx context.Context) ClusterEnhancedMonitoringPtrOutput {
	return ClusterEnhancedMonitoring(e).ToClusterEnhancedMonitoringOutputWithContext(ctx).ToClusterEnhancedMonitoringPtrOutputWithContext(ctx)
}

func (e ClusterEnhancedMonitoring) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClusterEnhancedMonitoring) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClusterEnhancedMonitoring) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ClusterEnhancedMonitoring) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ClusterEnhancedMonitoringOutput struct{ *pulumi.OutputState }

func (ClusterEnhancedMonitoringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterEnhancedMonitoring)(nil)).Elem()
}

func (o ClusterEnhancedMonitoringOutput) ToClusterEnhancedMonitoringOutput() ClusterEnhancedMonitoringOutput {
	return o
}

func (o ClusterEnhancedMonitoringOutput) ToClusterEnhancedMonitoringOutputWithContext(ctx context.Context) ClusterEnhancedMonitoringOutput {
	return o
}

func (o ClusterEnhancedMonitoringOutput) ToClusterEnhancedMonitoringPtrOutput() ClusterEnhancedMonitoringPtrOutput {
	return o.ToClusterEnhancedMonitoringPtrOutputWithContext(context.Background())
}

func (o ClusterEnhancedMonitoringOutput) ToClusterEnhancedMonitoringPtrOutputWithContext(ctx context.Context) ClusterEnhancedMonitoringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterEnhancedMonitoring) *ClusterEnhancedMonitoring {
		return &v
	}).(ClusterEnhancedMonitoringPtrOutput)
}

func (o ClusterEnhancedMonitoringOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ClusterEnhancedMonitoringOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClusterEnhancedMonitoring) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ClusterEnhancedMonitoringOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClusterEnhancedMonitoringOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClusterEnhancedMonitoring) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ClusterEnhancedMonitoringPtrOutput struct{ *pulumi.OutputState }

func (ClusterEnhancedMonitoringPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterEnhancedMonitoring)(nil)).Elem()
}

func (o ClusterEnhancedMonitoringPtrOutput) ToClusterEnhancedMonitoringPtrOutput() ClusterEnhancedMonitoringPtrOutput {
	return o
}

func (o ClusterEnhancedMonitoringPtrOutput) ToClusterEnhancedMonitoringPtrOutputWithContext(ctx context.Context) ClusterEnhancedMonitoringPtrOutput {
	return o
}

func (o ClusterEnhancedMonitoringPtrOutput) Elem() ClusterEnhancedMonitoringOutput {
	return o.ApplyT(func(v *ClusterEnhancedMonitoring) ClusterEnhancedMonitoring {
		if v != nil {
			return *v
		}
		var ret ClusterEnhancedMonitoring
		return ret
	}).(ClusterEnhancedMonitoringOutput)
}

func (o ClusterEnhancedMonitoringPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClusterEnhancedMonitoringPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ClusterEnhancedMonitoring) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ClusterEnhancedMonitoringInput is an input type that accepts ClusterEnhancedMonitoringArgs and ClusterEnhancedMonitoringOutput values.
// You can construct a concrete instance of `ClusterEnhancedMonitoringInput` via:
//
//          ClusterEnhancedMonitoringArgs{...}
type ClusterEnhancedMonitoringInput interface {
	pulumi.Input

	ToClusterEnhancedMonitoringOutput() ClusterEnhancedMonitoringOutput
	ToClusterEnhancedMonitoringOutputWithContext(context.Context) ClusterEnhancedMonitoringOutput
}

var clusterEnhancedMonitoringPtrType = reflect.TypeOf((**ClusterEnhancedMonitoring)(nil)).Elem()

type ClusterEnhancedMonitoringPtrInput interface {
	pulumi.Input

	ToClusterEnhancedMonitoringPtrOutput() ClusterEnhancedMonitoringPtrOutput
	ToClusterEnhancedMonitoringPtrOutputWithContext(context.Context) ClusterEnhancedMonitoringPtrOutput
}

type clusterEnhancedMonitoringPtr string

func ClusterEnhancedMonitoringPtr(v string) ClusterEnhancedMonitoringPtrInput {
	return (*clusterEnhancedMonitoringPtr)(&v)
}

func (*clusterEnhancedMonitoringPtr) ElementType() reflect.Type {
	return clusterEnhancedMonitoringPtrType
}

func (in *clusterEnhancedMonitoringPtr) ToClusterEnhancedMonitoringPtrOutput() ClusterEnhancedMonitoringPtrOutput {
	return pulumi.ToOutput(in).(ClusterEnhancedMonitoringPtrOutput)
}

func (in *clusterEnhancedMonitoringPtr) ToClusterEnhancedMonitoringPtrOutputWithContext(ctx context.Context) ClusterEnhancedMonitoringPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ClusterEnhancedMonitoringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterEncryptionInTransitClientBrokerInput)(nil)).Elem(), ClusterEncryptionInTransitClientBroker("TLS"))
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterEncryptionInTransitClientBrokerPtrInput)(nil)).Elem(), ClusterEncryptionInTransitClientBroker("TLS"))
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterEnhancedMonitoringInput)(nil)).Elem(), ClusterEnhancedMonitoring("DEFAULT"))
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterEnhancedMonitoringPtrInput)(nil)).Elem(), ClusterEnhancedMonitoring("DEFAULT"))
	pulumi.RegisterOutputType(ClusterEncryptionInTransitClientBrokerOutput{})
	pulumi.RegisterOutputType(ClusterEncryptionInTransitClientBrokerPtrOutput{})
	pulumi.RegisterOutputType(ClusterEnhancedMonitoringOutput{})
	pulumi.RegisterOutputType(ClusterEnhancedMonitoringPtrOutput{})
}