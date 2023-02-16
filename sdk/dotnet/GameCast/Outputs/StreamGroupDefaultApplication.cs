// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameCast.Outputs
{

    /// <summary>
    /// Information about default application running on the stream group.
    /// </summary>
    [OutputType]
    public sealed class StreamGroupDefaultApplication
    {
        /// <summary>
        /// GameCast resource ID, base62 encoded.
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private StreamGroupDefaultApplication(string? id)
        {
            Id = id;
        }
    }
}
