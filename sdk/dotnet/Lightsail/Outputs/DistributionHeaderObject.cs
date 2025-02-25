// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lightsail.Outputs
{

    /// <summary>
    /// Describes the request headers that a Lightsail distribution bases caching on.
    /// </summary>
    [OutputType]
    public sealed class DistributionHeaderObject
    {
        /// <summary>
        /// The specific headers to forward to your distribution's origin.
        /// </summary>
        public readonly ImmutableArray<string> HeadersAllowList;
        /// <summary>
        /// The headers that you want your distribution to forward to your origin and base caching on.
        /// </summary>
        public readonly string? Option;

        [OutputConstructor]
        private DistributionHeaderObject(
            ImmutableArray<string> headersAllowList,

            string? option)
        {
            HeadersAllowList = headersAllowList;
            Option = option;
        }
    }
}
