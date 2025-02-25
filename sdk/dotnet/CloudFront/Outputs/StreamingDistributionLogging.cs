// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Outputs
{

    [OutputType]
    public sealed class StreamingDistributionLogging
    {
        public readonly string Bucket;
        public readonly bool Enabled;
        public readonly string Prefix;

        [OutputConstructor]
        private StreamingDistributionLogging(
            string bucket,

            bool enabled,

            string prefix)
        {
            Bucket = bucket;
            Enabled = enabled;
            Prefix = prefix;
        }
    }
}
