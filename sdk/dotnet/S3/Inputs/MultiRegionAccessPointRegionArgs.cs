// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    public sealed class MultiRegionAccessPointRegionArgs : global::Pulumi.ResourceArgs
    {
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        [Input("bucketAccountId")]
        public Input<string>? BucketAccountId { get; set; }

        public MultiRegionAccessPointRegionArgs()
        {
        }
        public static new MultiRegionAccessPointRegionArgs Empty => new MultiRegionAccessPointRegionArgs();
    }
}
