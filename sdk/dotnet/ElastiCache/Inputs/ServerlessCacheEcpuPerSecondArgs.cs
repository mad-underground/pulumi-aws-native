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
    /// The ECPU per second of the Serverless Cache.
    /// </summary>
    public sealed class ServerlessCacheEcpuPerSecondArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum ECPU per second of the Serverless Cache.
        /// </summary>
        [Input("maximum", required: true)]
        public Input<int> Maximum { get; set; } = null!;

        public ServerlessCacheEcpuPerSecondArgs()
        {
        }
        public static new ServerlessCacheEcpuPerSecondArgs Empty => new ServerlessCacheEcpuPerSecondArgs();
    }
}
