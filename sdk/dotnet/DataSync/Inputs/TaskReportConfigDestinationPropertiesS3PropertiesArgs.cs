// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataSync.Inputs
{

    /// <summary>
    /// Specifies the Amazon S3 bucket where DataSync uploads your task report.
    /// </summary>
    public sealed class TaskReportConfigDestinationPropertiesS3PropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the Amazon Resource Name (ARN) of the IAM policy that allows Datasync to upload a task report to your S3 bucket.
        /// </summary>
        [Input("bucketAccessRoleArn")]
        public Input<string>? BucketAccessRoleArn { get; set; }

        /// <summary>
        /// Specifies the ARN of the S3 bucket where Datasync uploads your report.
        /// </summary>
        [Input("s3BucketArn")]
        public Input<string>? S3BucketArn { get; set; }

        /// <summary>
        /// Specifies a bucket prefix for your report.
        /// </summary>
        [Input("subdirectory")]
        public Input<string>? Subdirectory { get; set; }

        public TaskReportConfigDestinationPropertiesS3PropertiesArgs()
        {
        }
        public static new TaskReportConfigDestinationPropertiesS3PropertiesArgs Empty => new TaskReportConfigDestinationPropertiesS3PropertiesArgs();
    }
}