// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    /// <summary>
    /// Information about where and how to store the results of a monitoring job.
    /// </summary>
    public sealed class ModelExplainabilityJobDefinitionS3OutputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The local path to the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring job. LocalPath is an absolute path for the output data.
        /// </summary>
        [Input("localPath", required: true)]
        public Input<string> LocalPath { get; set; } = null!;

        /// <summary>
        /// Whether to upload the results of the monitoring job continuously or after the job completes.
        /// </summary>
        [Input("s3UploadMode")]
        public Input<Pulumi.AwsNative.SageMaker.ModelExplainabilityJobDefinitionS3OutputS3UploadMode>? S3UploadMode { get; set; }

        /// <summary>
        /// A URI that identifies the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring job.
        /// </summary>
        [Input("s3Uri", required: true)]
        public Input<string> S3Uri { get; set; } = null!;

        public ModelExplainabilityJobDefinitionS3OutputArgs()
        {
        }
        public static new ModelExplainabilityJobDefinitionS3OutputArgs Empty => new ModelExplainabilityJobDefinitionS3OutputArgs();
    }
}
