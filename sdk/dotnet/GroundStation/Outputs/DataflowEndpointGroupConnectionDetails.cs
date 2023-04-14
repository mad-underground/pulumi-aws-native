// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GroundStation.Outputs
{

    /// <summary>
    /// Egress address of AgentEndpoint with an optional mtu.
    /// </summary>
    [OutputType]
    public sealed class DataflowEndpointGroupConnectionDetails
    {
        /// <summary>
        /// Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.
        /// </summary>
        public readonly int? Mtu;
        public readonly Outputs.DataflowEndpointGroupSocketAddress? SocketAddress;

        [OutputConstructor]
        private DataflowEndpointGroupConnectionDetails(
            int? mtu,

            Outputs.DataflowEndpointGroupSocketAddress? socketAddress)
        {
            Mtu = mtu;
            SocketAddress = socketAddress;
        }
    }
}