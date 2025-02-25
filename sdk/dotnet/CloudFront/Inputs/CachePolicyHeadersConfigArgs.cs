// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Inputs
{

    public sealed class CachePolicyHeadersConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("headerBehavior", required: true)]
        public Input<string> HeaderBehavior { get; set; } = null!;

        [Input("headers")]
        private InputList<string>? _headers;
        public InputList<string> Headers
        {
            get => _headers ?? (_headers = new InputList<string>());
            set => _headers = value;
        }

        public CachePolicyHeadersConfigArgs()
        {
        }
        public static new CachePolicyHeadersConfigArgs Empty => new CachePolicyHeadersConfigArgs();
    }
}
