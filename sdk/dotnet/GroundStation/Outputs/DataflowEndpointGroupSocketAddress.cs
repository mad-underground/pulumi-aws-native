// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GroundStation.Outputs
{

    [OutputType]
    public sealed class DataflowEndpointGroupSocketAddress
    {
        public readonly string? Name;
        public readonly int? Port;

        [OutputConstructor]
        private DataflowEndpointGroupSocketAddress(
            string? name,

            int? port)
        {
            Name = name;
            Port = port;
        }
    }
}
