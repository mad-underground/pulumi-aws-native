// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    public sealed class BucketRedirectAllRequestsToArgs : global::Pulumi.ResourceArgs
    {
        [Input("hostName", required: true)]
        public Input<string> HostName { get; set; } = null!;

        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        public BucketRedirectAllRequestsToArgs()
        {
        }
        public static new BucketRedirectAllRequestsToArgs Empty => new BucketRedirectAllRequestsToArgs();
    }
}
