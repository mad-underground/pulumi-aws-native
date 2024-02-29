// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Outputs
{

    /// <summary>
    /// Specifies tags to use to identify a subset of objects for an Amazon S3 bucket.
    /// </summary>
    [OutputType]
    public sealed class BucketTagFilter
    {
        /// <summary>
        /// The tag key.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// The tag value.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private BucketTagFilter(
            string key,

            string value)
        {
            Key = key;
            Value = value;
        }
    }
}
