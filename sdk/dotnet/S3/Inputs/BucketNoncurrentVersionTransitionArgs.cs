// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    public sealed class BucketNoncurrentVersionTransitionArgs : Pulumi.ResourceArgs
    {
        [Input("storageClass", required: true)]
        public Input<string> StorageClass { get; set; } = null!;

        [Input("transitionInDays", required: true)]
        public Input<int> TransitionInDays { get; set; } = null!;

        public BucketNoncurrentVersionTransitionArgs()
        {
        }
    }
}