// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Efs.Outputs
{

    /// <summary>
    /// Describes the replication configuration for a specific file system.
    /// </summary>
    [OutputType]
    public sealed class FileSystemReplicationConfiguration
    {
        /// <summary>
        /// An array of destination objects. Only one destination object is supported.
        /// </summary>
        public readonly ImmutableArray<Outputs.FileSystemReplicationDestination> Destinations;

        [OutputConstructor]
        private FileSystemReplicationConfiguration(ImmutableArray<Outputs.FileSystemReplicationDestination> destinations)
        {
            Destinations = destinations;
        }
    }
}
