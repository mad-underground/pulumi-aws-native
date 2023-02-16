// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SSM.Inputs
{

    public sealed class ResourceDataSyncS3DestinationArgs : global::Pulumi.ResourceArgs
    {
        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        [Input("bucketPrefix")]
        public Input<string>? BucketPrefix { get; set; }

        [Input("bucketRegion", required: true)]
        public Input<string> BucketRegion { get; set; } = null!;

        [Input("kMSKeyArn")]
        public Input<string>? KMSKeyArn { get; set; }

        [Input("syncFormat", required: true)]
        public Input<string> SyncFormat { get; set; } = null!;

        public ResourceDataSyncS3DestinationArgs()
        {
        }
        public static new ResourceDataSyncS3DestinationArgs Empty => new ResourceDataSyncS3DestinationArgs();
    }
}
