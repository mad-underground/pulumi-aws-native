// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    public sealed class BucketRedirectRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("hostName")]
        public Input<string>? HostName { get; set; }

        [Input("httpRedirectCode")]
        public Input<string>? HttpRedirectCode { get; set; }

        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        [Input("replaceKeyPrefixWith")]
        public Input<string>? ReplaceKeyPrefixWith { get; set; }

        [Input("replaceKeyWith")]
        public Input<string>? ReplaceKeyWith { get; set; }

        public BucketRedirectRuleArgs()
        {
        }
        public static new BucketRedirectRuleArgs Empty => new BucketRedirectRuleArgs();
    }
}
