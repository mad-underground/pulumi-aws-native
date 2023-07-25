// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Outputs
{

    /// <summary>
    /// Specifies a cookie as an aggregate key for a rate-based rule.
    /// </summary>
    [OutputType]
    public sealed class WebACLRateLimitCookie
    {
        /// <summary>
        /// The name of the cookie to use.
        /// </summary>
        public readonly string Name;
        public readonly ImmutableArray<Outputs.WebACLTextTransformation> TextTransformations;

        [OutputConstructor]
        private WebACLRateLimitCookie(
            string name,

            ImmutableArray<Outputs.WebACLTextTransformation> textTransformations)
        {
            Name = name;
            TextTransformations = textTransformations;
        }
    }
}
