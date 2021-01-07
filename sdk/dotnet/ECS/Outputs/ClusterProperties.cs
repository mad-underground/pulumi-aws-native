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
    public sealed class ClusterProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-cluster.html#cfn-ecs-cluster-capacityproviders
        /// </summary>
        public readonly ImmutableArray<string> CapacityProviders;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-cluster.html#cfn-ecs-cluster-clustername
        /// </summary>
        public readonly string? ClusterName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-cluster.html#cfn-ecs-cluster-clustersettings
        /// </summary>
        public readonly ImmutableArray<Outputs.ClusterClusterSettings> ClusterSettings;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-cluster.html#cfn-ecs-cluster-defaultcapacityproviderstrategy
        /// </summary>
        public readonly ImmutableArray<Outputs.ClusterCapacityProviderStrategyItem> DefaultCapacityProviderStrategy;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-cluster.html#cfn-ecs-cluster-tags
        /// </summary>
        public readonly ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags;

        [OutputConstructor]
        private ClusterProperties(
            ImmutableArray<string> CapacityProviders,

            string? ClusterName,

            ImmutableArray<Outputs.ClusterClusterSettings> ClusterSettings,

            ImmutableArray<Outputs.ClusterCapacityProviderStrategyItem> DefaultCapacityProviderStrategy,

            ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags)
        {
            this.CapacityProviders = CapacityProviders;
            this.ClusterName = ClusterName;
            this.ClusterSettings = ClusterSettings;
            this.DefaultCapacityProviderStrategy = DefaultCapacityProviderStrategy;
            this.Tags = Tags;
        }
    }
}