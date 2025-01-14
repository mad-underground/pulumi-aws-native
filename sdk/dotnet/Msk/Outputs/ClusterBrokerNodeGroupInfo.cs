// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Msk.Outputs
{

    [OutputType]
    public sealed class ClusterBrokerNodeGroupInfo
    {
        public readonly string? BrokerAzDistribution;
        public readonly ImmutableArray<string> ClientSubnets;
        public readonly Outputs.ClusterConnectivityInfo? ConnectivityInfo;
        public readonly string InstanceType;
        public readonly ImmutableArray<string> SecurityGroups;
        public readonly Outputs.ClusterStorageInfo? StorageInfo;

        [OutputConstructor]
        private ClusterBrokerNodeGroupInfo(
            string? brokerAzDistribution,

            ImmutableArray<string> clientSubnets,

            Outputs.ClusterConnectivityInfo? connectivityInfo,

            string instanceType,

            ImmutableArray<string> securityGroups,

            Outputs.ClusterStorageInfo? storageInfo)
        {
            BrokerAzDistribution = brokerAzDistribution;
            ClientSubnets = clientSubnets;
            ConnectivityInfo = connectivityInfo;
            InstanceType = instanceType;
            SecurityGroups = securityGroups;
            StorageInfo = storageInfo;
        }
    }
}
