// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Synthetics.Outputs
{

    [OutputType]
    public sealed class CanaryCode
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-code.html#cfn-synthetics-canary-code-handler
        /// </summary>
        public readonly string? Handler;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-code.html#cfn-synthetics-canary-code-s3bucket
        /// </summary>
        public readonly string? S3Bucket;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-code.html#cfn-synthetics-canary-code-s3key
        /// </summary>
        public readonly string? S3Key;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-code.html#cfn-synthetics-canary-code-s3objectversion
        /// </summary>
        public readonly string? S3ObjectVersion;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-code.html#cfn-synthetics-canary-code-script
        /// </summary>
        public readonly string? Script;

        [OutputConstructor]
        private CanaryCode(
            string? handler,

            string? s3Bucket,

            string? s3Key,

            string? s3ObjectVersion,

            string? script)
        {
            Handler = handler;
            S3Bucket = s3Bucket;
            S3Key = s3Key;
            S3ObjectVersion = s3ObjectVersion;
            Script = script;
        }
    }
}
