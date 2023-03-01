// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IVSChat.Inputs
{

    /// <summary>
    /// S3 destination configuration for IVS Chat logging.
    /// </summary>
    public sealed class LoggingConfigurationS3DestinationConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the Amazon S3 bucket where chat activity will be logged.
        /// </summary>
        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        public LoggingConfigurationS3DestinationConfigurationArgs()
        {
        }
        public static new LoggingConfigurationS3DestinationConfigurationArgs Empty => new LoggingConfigurationS3DestinationConfigurationArgs();
    }
}
