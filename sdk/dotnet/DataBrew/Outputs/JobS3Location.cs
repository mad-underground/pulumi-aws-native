// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.DataBrew.Outputs
{

    [OutputType]
    public sealed class JobS3Location
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-s3location.html#cfn-databrew-job-s3location-bucket
        /// </summary>
        public readonly string Bucket;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-s3location.html#cfn-databrew-job-s3location-key
        /// </summary>
        public readonly string? Key;

        [OutputConstructor]
        private JobS3Location(
            string Bucket,

            string? Key)
        {
            this.Bucket = Bucket;
            this.Key = Key;
        }
    }
}
