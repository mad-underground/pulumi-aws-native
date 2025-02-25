// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Outputs
{

    /// <summary>
    /// Includes the JA3 fingerprint of a web request.
    /// </summary>
    [OutputType]
    public sealed class WebAclJa3Fingerprint
    {
        public readonly Pulumi.AwsNative.WaFv2.WebAclJa3FingerprintFallbackBehavior FallbackBehavior;

        [OutputConstructor]
        private WebAclJa3Fingerprint(Pulumi.AwsNative.WaFv2.WebAclJa3FingerprintFallbackBehavior fallbackBehavior)
        {
            FallbackBehavior = fallbackBehavior;
        }
    }
}
