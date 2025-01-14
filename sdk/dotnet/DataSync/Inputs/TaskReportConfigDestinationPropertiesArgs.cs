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
    /// Specifies where DataSync uploads your task report.
    /// </summary>
    public sealed class TaskReportConfigDestinationPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the Amazon S3 bucket where DataSync uploads your task report.
        /// </summary>
        [Input("s3")]
        public Input<Inputs.TaskReportConfigDestinationPropertiesS3PropertiesArgs>? S3 { get; set; }

        public TaskReportConfigDestinationPropertiesArgs()
        {
        }
        public static new TaskReportConfigDestinationPropertiesArgs Empty => new TaskReportConfigDestinationPropertiesArgs();
    }
}
