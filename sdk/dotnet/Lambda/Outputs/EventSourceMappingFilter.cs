// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lambda.Outputs
{

    /// <summary>
    /// The filter object that defines parameters for ESM filtering.
    /// </summary>
    [OutputType]
    public sealed class EventSourceMappingFilter
    {
        /// <summary>
        /// The filter pattern that defines which events should be passed for invocations.
        /// </summary>
        public readonly string? Pattern;

        [OutputConstructor]
        private EventSourceMappingFilter(string? pattern)
        {
            Pattern = pattern;
        }
    }
}
