// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    public sealed class BucketReplicationTimeArgs : global::Pulumi.ResourceArgs
    {
        [Input("status", required: true)]
        public Input<string> Status { get; set; } = null!;

        [Input("time", required: true)]
        public Input<Inputs.BucketReplicationTimeValueArgs> Time { get; set; } = null!;

        public BucketReplicationTimeArgs()
        {
        }
        public static new BucketReplicationTimeArgs Empty => new BucketReplicationTimeArgs();
    }
}
