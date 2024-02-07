// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package autoscaling

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AutoScaling::AutoScalingGroup
func LookupAutoScalingGroup(ctx *pulumi.Context, args *LookupAutoScalingGroupArgs, opts ...pulumi.InvokeOption) (*LookupAutoScalingGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAutoScalingGroupResult
	err := ctx.Invoke("aws-native:autoscaling:getAutoScalingGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAutoScalingGroupArgs struct {
	AutoScalingGroupName string `pulumi:"autoScalingGroupName"`
}

type LookupAutoScalingGroupResult struct {
	AvailabilityZones                []string                                     `pulumi:"availabilityZones"`
	CapacityRebalance                *bool                                        `pulumi:"capacityRebalance"`
	Context                          *string                                      `pulumi:"context"`
	Cooldown                         *string                                      `pulumi:"cooldown"`
	DefaultInstanceWarmup            *int                                         `pulumi:"defaultInstanceWarmup"`
	DesiredCapacity                  *string                                      `pulumi:"desiredCapacity"`
	DesiredCapacityType              *string                                      `pulumi:"desiredCapacityType"`
	HealthCheckGracePeriod           *int                                         `pulumi:"healthCheckGracePeriod"`
	HealthCheckType                  *string                                      `pulumi:"healthCheckType"`
	InstanceMaintenancePolicy        *AutoScalingGroupInstanceMaintenancePolicy   `pulumi:"instanceMaintenancePolicy"`
	LaunchConfigurationName          *string                                      `pulumi:"launchConfigurationName"`
	LaunchTemplate                   *AutoScalingGroupLaunchTemplateSpecification `pulumi:"launchTemplate"`
	LifecycleHookSpecificationList   []AutoScalingGroupLifecycleHookSpecification `pulumi:"lifecycleHookSpecificationList"`
	LoadBalancerNames                []string                                     `pulumi:"loadBalancerNames"`
	MaxInstanceLifetime              *int                                         `pulumi:"maxInstanceLifetime"`
	MaxSize                          *string                                      `pulumi:"maxSize"`
	MetricsCollection                []AutoScalingGroupMetricsCollection          `pulumi:"metricsCollection"`
	MinSize                          *string                                      `pulumi:"minSize"`
	MixedInstancesPolicy             *AutoScalingGroupMixedInstancesPolicy        `pulumi:"mixedInstancesPolicy"`
	NewInstancesProtectedFromScaleIn *bool                                        `pulumi:"newInstancesProtectedFromScaleIn"`
	NotificationConfiguration        *AutoScalingGroupNotificationConfiguration   `pulumi:"notificationConfiguration"`
	NotificationConfigurations       []AutoScalingGroupNotificationConfiguration  `pulumi:"notificationConfigurations"`
	PlacementGroup                   *string                                      `pulumi:"placementGroup"`
	ServiceLinkedRoleArn             *string                                      `pulumi:"serviceLinkedRoleArn"`
	Tags                             []AutoScalingGroupTagProperty                `pulumi:"tags"`
	TargetGroupArns                  []string                                     `pulumi:"targetGroupArns"`
	TerminationPolicies              []string                                     `pulumi:"terminationPolicies"`
	VpcZoneIdentifier                []string                                     `pulumi:"vpcZoneIdentifier"`
}

func LookupAutoScalingGroupOutput(ctx *pulumi.Context, args LookupAutoScalingGroupOutputArgs, opts ...pulumi.InvokeOption) LookupAutoScalingGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAutoScalingGroupResult, error) {
			args := v.(LookupAutoScalingGroupArgs)
			r, err := LookupAutoScalingGroup(ctx, &args, opts...)
			var s LookupAutoScalingGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAutoScalingGroupResultOutput)
}

type LookupAutoScalingGroupOutputArgs struct {
	AutoScalingGroupName pulumi.StringInput `pulumi:"autoScalingGroupName"`
}

func (LookupAutoScalingGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutoScalingGroupArgs)(nil)).Elem()
}

type LookupAutoScalingGroupResultOutput struct{ *pulumi.OutputState }

func (LookupAutoScalingGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutoScalingGroupResult)(nil)).Elem()
}

func (o LookupAutoScalingGroupResultOutput) ToLookupAutoScalingGroupResultOutput() LookupAutoScalingGroupResultOutput {
	return o
}

func (o LookupAutoScalingGroupResultOutput) ToLookupAutoScalingGroupResultOutputWithContext(ctx context.Context) LookupAutoScalingGroupResultOutput {
	return o
}

func (o LookupAutoScalingGroupResultOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAutoScalingGroupResult) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o LookupAutoScalingGroupResultOutput) CapacityRebalance() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAutoScalingGroupResult) *bool { return v.CapacityRebalance }).(pulumi.BoolPtrOutput)
}

func (o LookupAutoScalingGroupResultOutput) Context() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutoScalingGroupResult) *string { return v.Context }).(pulumi.StringPtrOutput)
}

func (o LookupAutoScalingGroupResultOutput) Cooldown() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutoScalingGroupResult) *string { return v.Cooldown }).(pulumi.StringPtrOutput)
}

func (o LookupAutoScalingGroupResultOutput) DefaultInstanceWarmup() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAutoScalingGroupResult) *int { return v.DefaultInstanceWarmup }).(pulumi.IntPtrOutput)
}

func (o LookupAutoScalingGroupResultOutput) DesiredCapacity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutoScalingGroupResult) *string { return v.DesiredCapacity }).(pulumi.StringPtrOutput)
}

func (o LookupAutoScalingGroupResultOutput) DesiredCapacityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutoScalingGroupResult) *string { return v.DesiredCapacityType }).(pulumi.StringPtrOutput)
}

func (o LookupAutoScalingGroupResultOutput) HealthCheckGracePeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAutoScalingGroupResult) *int { return v.HealthCheckGracePeriod }).(pulumi.IntPtrOutput)
}

func (o LookupAutoScalingGroupResultOutput) HealthCheckType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutoScalingGroupResult) *string { return v.HealthCheckType }).(pulumi.StringPtrOutput)
}

func (o LookupAutoScalingGroupResultOutput) InstanceMaintenancePolicy() AutoScalingGroupInstanceMaintenancePolicyPtrOutput {
	return o.ApplyT(func(v LookupAutoScalingGroupResult) *AutoScalingGroupInstanceMaintenancePolicy {
		return v.InstanceMaintenancePolicy
	}).(AutoScalingGroupInstanceMaintenancePolicyPtrOutput)
}

func (o LookupAutoScalingGroupResultOutput) LaunchConfigurationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutoScalingGroupResult) *string { return v.LaunchConfigurationName }).(pulumi.StringPtrOutput)
}

func (o LookupAutoScalingGroupResultOutput) LaunchTemplate() AutoScalingGroupLaunchTemplateSpecificationPtrOutput {
	return o.ApplyT(func(v LookupAutoScalingGroupResult) *AutoScalingGroupLaunchTemplateSpecification {
		return v.LaunchTemplate
	}).(AutoScalingGroupLaunchTemplateSpecificationPtrOutput)
}

func (o LookupAutoScalingGroupResultOutput) LifecycleHookSpecificationList() AutoScalingGroupLifecycleHookSpecificationArrayOutput {
	return o.ApplyT(func(v LookupAutoScalingGroupResult) []AutoScalingGroupLifecycleHookSpecification {
		return v.LifecycleHookSpecificationList
	}).(AutoScalingGroupLifecycleHookSpecificationArrayOutput)
}

func (o LookupAutoScalingGroupResultOutput) LoadBalancerNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAutoScalingGroupResult) []string { return v.LoadBalancerNames }).(pulumi.StringArrayOutput)
}

func (o LookupAutoScalingGroupResultOutput) MaxInstanceLifetime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAutoScalingGroupResult) *int { return v.MaxInstanceLifetime }).(pulumi.IntPtrOutput)
}

func (o LookupAutoScalingGroupResultOutput) MaxSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutoScalingGroupResult) *string { return v.MaxSize }).(pulumi.StringPtrOutput)
}

func (o LookupAutoScalingGroupResultOutput) MetricsCollection() AutoScalingGroupMetricsCollectionArrayOutput {
	return o.ApplyT(func(v LookupAutoScalingGroupResult) []AutoScalingGroupMetricsCollection { return v.MetricsCollection }).(AutoScalingGroupMetricsCollectionArrayOutput)
}

func (o LookupAutoScalingGroupResultOutput) MinSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutoScalingGroupResult) *string { return v.MinSize }).(pulumi.StringPtrOutput)
}

func (o LookupAutoScalingGroupResultOutput) MixedInstancesPolicy() AutoScalingGroupMixedInstancesPolicyPtrOutput {
	return o.ApplyT(func(v LookupAutoScalingGroupResult) *AutoScalingGroupMixedInstancesPolicy {
		return v.MixedInstancesPolicy
	}).(AutoScalingGroupMixedInstancesPolicyPtrOutput)
}

func (o LookupAutoScalingGroupResultOutput) NewInstancesProtectedFromScaleIn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAutoScalingGroupResult) *bool { return v.NewInstancesProtectedFromScaleIn }).(pulumi.BoolPtrOutput)
}

func (o LookupAutoScalingGroupResultOutput) NotificationConfiguration() AutoScalingGroupNotificationConfigurationPtrOutput {
	return o.ApplyT(func(v LookupAutoScalingGroupResult) *AutoScalingGroupNotificationConfiguration {
		return v.NotificationConfiguration
	}).(AutoScalingGroupNotificationConfigurationPtrOutput)
}

func (o LookupAutoScalingGroupResultOutput) NotificationConfigurations() AutoScalingGroupNotificationConfigurationArrayOutput {
	return o.ApplyT(func(v LookupAutoScalingGroupResult) []AutoScalingGroupNotificationConfiguration {
		return v.NotificationConfigurations
	}).(AutoScalingGroupNotificationConfigurationArrayOutput)
}

func (o LookupAutoScalingGroupResultOutput) PlacementGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutoScalingGroupResult) *string { return v.PlacementGroup }).(pulumi.StringPtrOutput)
}

func (o LookupAutoScalingGroupResultOutput) ServiceLinkedRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutoScalingGroupResult) *string { return v.ServiceLinkedRoleArn }).(pulumi.StringPtrOutput)
}

func (o LookupAutoScalingGroupResultOutput) Tags() AutoScalingGroupTagPropertyArrayOutput {
	return o.ApplyT(func(v LookupAutoScalingGroupResult) []AutoScalingGroupTagProperty { return v.Tags }).(AutoScalingGroupTagPropertyArrayOutput)
}

func (o LookupAutoScalingGroupResultOutput) TargetGroupArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAutoScalingGroupResult) []string { return v.TargetGroupArns }).(pulumi.StringArrayOutput)
}

func (o LookupAutoScalingGroupResultOutput) TerminationPolicies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAutoScalingGroupResult) []string { return v.TerminationPolicies }).(pulumi.StringArrayOutput)
}

func (o LookupAutoScalingGroupResultOutput) VpcZoneIdentifier() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAutoScalingGroupResult) []string { return v.VpcZoneIdentifier }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAutoScalingGroupResultOutput{})
}
