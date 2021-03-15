// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Outputs
{

    [OutputType]
    public sealed class StorageLensBucketLevel
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-bucketlevel.html#cfn-s3-storagelens-bucketlevel-activitymetrics
        /// </summary>
        public readonly Outputs.StorageLensActivityMetrics? ActivityMetrics;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-bucketlevel.html#cfn-s3-storagelens-bucketlevel-prefixlevel
        /// </summary>
        public readonly Outputs.StorageLensPrefixLevel? PrefixLevel;

        [OutputConstructor]
        private StorageLensBucketLevel(
            Outputs.StorageLensActivityMetrics? activityMetrics,

            Outputs.StorageLensPrefixLevel? prefixLevel)
        {
            ActivityMetrics = activityMetrics;
            PrefixLevel = prefixLevel;
        }
    }
}
