// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.EKS.Outputs
{

    [OutputType]
    public sealed class ClusterKubernetesNetworkConfig
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-kubernetesnetworkconfig.html#cfn-eks-cluster-kubernetesnetworkconfig-serviceipv4cidr
        /// </summary>
        public readonly string? ServiceIpv4Cidr;

        [OutputConstructor]
        private ClusterKubernetesNetworkConfig(string? ServiceIpv4Cidr)
        {
            this.ServiceIpv4Cidr = ServiceIpv4Cidr;
        }
    }
}
