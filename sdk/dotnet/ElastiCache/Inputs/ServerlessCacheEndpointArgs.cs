// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElastiCache.Inputs
{

    /// <summary>
    /// The address and the port.
    /// </summary>
    public sealed class ServerlessCacheEndpointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Endpoint address.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// Endpoint port.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        public ServerlessCacheEndpointArgs()
        {
        }
        public static new ServerlessCacheEndpointArgs Empty => new ServerlessCacheEndpointArgs();
    }
}
