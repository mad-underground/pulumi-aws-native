// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    public sealed class BucketDefaultRetentionArgs : global::Pulumi.ResourceArgs
    {
        [Input("days")]
        public Input<int>? Days { get; set; }

        [Input("mode")]
        public Input<string>? Mode { get; set; }

        [Input("years")]
        public Input<int>? Years { get; set; }

        public BucketDefaultRetentionArgs()
        {
        }
        public static new BucketDefaultRetentionArgs Empty => new BucketDefaultRetentionArgs();
    }
}
