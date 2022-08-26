// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RDS.Outputs
{

    [OutputType]
    public sealed class DBClusterEndpoint
    {
        /// <summary>
        /// The connection endpoint for the DB cluster.
        /// </summary>
        public readonly string? Address;
        /// <summary>
        /// The port number that will accept connections on this DB cluster.
        /// </summary>
        public readonly string? Port;

        [OutputConstructor]
        private DBClusterEndpoint(
            string? address,

            string? port)
        {
            Address = address;
            Port = port;
        }
    }
}