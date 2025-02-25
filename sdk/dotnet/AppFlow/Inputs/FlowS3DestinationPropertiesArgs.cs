// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Inputs
{

    public sealed class FlowS3DestinationPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        [Input("bucketPrefix")]
        public Input<string>? BucketPrefix { get; set; }

        [Input("s3OutputFormatConfig")]
        public Input<Inputs.FlowS3OutputFormatConfigArgs>? S3OutputFormatConfig { get; set; }

        public FlowS3DestinationPropertiesArgs()
        {
        }
        public static new FlowS3DestinationPropertiesArgs Empty => new FlowS3DestinationPropertiesArgs();
    }
}
