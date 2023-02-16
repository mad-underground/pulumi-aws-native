// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ECS.Outputs
{

    /// <summary>
    /// A capacity provider strategy consists of one or more capacity providers along with the `base` and `weight` to assign to them. A capacity provider must be associated with the cluster to be used in a capacity provider strategy. The PutClusterCapacityProviders API is used to associate a capacity provider with a cluster. Only capacity providers with an `ACTIVE` or `UPDATING` status can be used.
    /// </summary>
    [OutputType]
    public sealed class ClusterCapacityProviderStrategyItem
    {
        public readonly int? Base;
        public readonly string? CapacityProvider;
        public readonly int? Weight;

        [OutputConstructor]
        private ClusterCapacityProviderStrategyItem(
            int? @base,

            string? capacityProvider,

            int? weight)
        {
            Base = @base;
            CapacityProvider = capacityProvider;
            Weight = weight;
        }
    }
}
