// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancingV2.Outputs
{

    /// <summary>
    /// Specifies a subnet for a load balancer.
    /// </summary>
    [OutputType]
    public sealed class LoadBalancerSubnetMapping
    {
        /// <summary>
        /// [Network Load Balancers] The allocation ID of the Elastic IP address for an internet-facing load balancer.
        /// </summary>
        public readonly string? AllocationId;
        /// <summary>
        /// [Network Load Balancers] The IPv6 address.
        /// </summary>
        public readonly string? IPv6Address;
        /// <summary>
        /// [Network Load Balancers] The private IPv4 address for an internal load balancer.
        /// </summary>
        public readonly string? PrivateIPv4Address;
        /// <summary>
        /// The ID of the subnet.
        /// </summary>
        public readonly string SubnetId;

        [OutputConstructor]
        private LoadBalancerSubnetMapping(
            string? allocationId,

            string? iPv6Address,

            string? privateIPv4Address,

            string subnetId)
        {
            AllocationId = allocationId;
            IPv6Address = iPv6Address;
            PrivateIPv4Address = privateIPv4Address;
            SubnetId = subnetId;
        }
    }
}
