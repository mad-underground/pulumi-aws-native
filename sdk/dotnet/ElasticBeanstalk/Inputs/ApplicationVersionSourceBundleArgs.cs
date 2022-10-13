// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticBeanstalk.Inputs
{

    public sealed class ApplicationVersionSourceBundleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon S3 bucket where the data is located.
        /// </summary>
        [Input("s3Bucket", required: true)]
        public Input<string> S3Bucket { get; set; } = null!;

        /// <summary>
        /// The Amazon S3 key where the data is located.
        /// </summary>
        [Input("s3Key", required: true)]
        public Input<string> S3Key { get; set; } = null!;

        public ApplicationVersionSourceBundleArgs()
        {
        }
        public static new ApplicationVersionSourceBundleArgs Empty => new ApplicationVersionSourceBundleArgs();
    }
}
