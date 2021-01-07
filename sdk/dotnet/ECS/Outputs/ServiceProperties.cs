// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.ECS.Outputs
{

    [OutputType]
    public sealed class ServiceProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-capacityproviderstrategy
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceCapacityProviderStrategyItem> CapacityProviderStrategy;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-cluster
        /// </summary>
        public readonly string? Cluster;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-deploymentconfiguration
        /// </summary>
        public readonly Outputs.ServiceDeploymentConfiguration? DeploymentConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-deploymentcontroller
        /// </summary>
        public readonly Outputs.ServiceDeploymentController? DeploymentController;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-desiredcount
        /// </summary>
        public readonly int? DesiredCount;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-enableecsmanagedtags
        /// </summary>
        public readonly bool? EnableECSManagedTags;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-healthcheckgraceperiodseconds
        /// </summary>
        public readonly int? HealthCheckGracePeriodSeconds;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-launchtype
        /// </summary>
        public readonly string? LaunchType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-loadbalancers
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceLoadBalancer> LoadBalancers;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-networkconfiguration
        /// </summary>
        public readonly Outputs.ServiceNetworkConfiguration? NetworkConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-placementconstraints
        /// </summary>
        public readonly ImmutableArray<Outputs.ServicePlacementConstraint> PlacementConstraints;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-placementstrategies
        /// </summary>
        public readonly ImmutableArray<Outputs.ServicePlacementStrategy> PlacementStrategies;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-platformversion
        /// </summary>
        public readonly string? PlatformVersion;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-propagatetags
        /// </summary>
        public readonly string? PropagateTags;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-role
        /// </summary>
        public readonly string? Role;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-schedulingstrategy
        /// </summary>
        public readonly string? SchedulingStrategy;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-servicearn
        /// </summary>
        public readonly string? ServiceArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-servicename
        /// </summary>
        public readonly string? ServiceName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-serviceregistries
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceServiceRegistry> ServiceRegistries;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-tags
        /// </summary>
        public readonly ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-taskdefinition
        /// </summary>
        public readonly string? TaskDefinition;

        [OutputConstructor]
        private ServiceProperties(
            ImmutableArray<Outputs.ServiceCapacityProviderStrategyItem> CapacityProviderStrategy,

            string? Cluster,

            Outputs.ServiceDeploymentConfiguration? DeploymentConfiguration,

            Outputs.ServiceDeploymentController? DeploymentController,

            int? DesiredCount,

            bool? EnableECSManagedTags,

            int? HealthCheckGracePeriodSeconds,

            string? LaunchType,

            ImmutableArray<Outputs.ServiceLoadBalancer> LoadBalancers,

            Outputs.ServiceNetworkConfiguration? NetworkConfiguration,

            ImmutableArray<Outputs.ServicePlacementConstraint> PlacementConstraints,

            ImmutableArray<Outputs.ServicePlacementStrategy> PlacementStrategies,

            string? PlatformVersion,

            string? PropagateTags,

            string? Role,

            string? SchedulingStrategy,

            string? ServiceArn,

            string? ServiceName,

            ImmutableArray<Outputs.ServiceServiceRegistry> ServiceRegistries,

            ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags,

            string? TaskDefinition)
        {
            this.CapacityProviderStrategy = CapacityProviderStrategy;
            this.Cluster = Cluster;
            this.DeploymentConfiguration = DeploymentConfiguration;
            this.DeploymentController = DeploymentController;
            this.DesiredCount = DesiredCount;
            this.EnableECSManagedTags = EnableECSManagedTags;
            this.HealthCheckGracePeriodSeconds = HealthCheckGracePeriodSeconds;
            this.LaunchType = LaunchType;
            this.LoadBalancers = LoadBalancers;
            this.NetworkConfiguration = NetworkConfiguration;
            this.PlacementConstraints = PlacementConstraints;
            this.PlacementStrategies = PlacementStrategies;
            this.PlatformVersion = PlatformVersion;
            this.PropagateTags = PropagateTags;
            this.Role = Role;
            this.SchedulingStrategy = SchedulingStrategy;
            this.ServiceArn = ServiceArn;
            this.ServiceName = ServiceName;
            this.ServiceRegistries = ServiceRegistries;
            this.Tags = Tags;
            this.TaskDefinition = TaskDefinition;
        }
    }
}