// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancingV2.Inputs
{

    /// <summary>
    /// Specifies a subnet for a load balancer.
    /// </summary>
    public sealed class LoadBalancerSubnetMappingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// [Network Load Balancers] The allocation ID of the Elastic IP address for an internet-facing load balancer.
        /// </summary>
        [Input("allocationId")]
        public Input<string>? AllocationId { get; set; }

        /// <summary>
        /// [Network Load Balancers] The IPv6 address.
        /// </summary>
        [Input("iPv6Address")]
        public Input<string>? IPv6Address { get; set; }

        /// <summary>
        /// [Network Load Balancers] The private IPv4 address for an internal load balancer.
        /// </summary>
        [Input("privateIPv4Address")]
        public Input<string>? PrivateIPv4Address { get; set; }

        /// <summary>
        /// The ID of the subnet.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        public LoadBalancerSubnetMappingArgs()
        {
        }
        public static new LoadBalancerSubnetMappingArgs Empty => new LoadBalancerSubnetMappingArgs();
    }
}
