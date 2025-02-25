// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Synthetics.Inputs
{

    public sealed class CanaryCodeArgs : global::Pulumi.ResourceArgs
    {
        [Input("handler", required: true)]
        public Input<string> Handler { get; set; } = null!;

        [Input("s3Bucket")]
        public Input<string>? S3Bucket { get; set; }

        [Input("s3Key")]
        public Input<string>? S3Key { get; set; }

        [Input("s3ObjectVersion")]
        public Input<string>? S3ObjectVersion { get; set; }

        [Input("script")]
        public Input<string>? Script { get; set; }

        [Input("sourceLocationArn")]
        public Input<string>? SourceLocationArn { get; set; }

        public CanaryCodeArgs()
        {
        }
        public static new CanaryCodeArgs Empty => new CanaryCodeArgs();
    }
}
