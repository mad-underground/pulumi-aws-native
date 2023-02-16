// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Inputs
{

    public sealed class DistributionCustomOriginConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("hTTPPort")]
        public Input<int>? HTTPPort { get; set; }

        [Input("hTTPSPort")]
        public Input<int>? HTTPSPort { get; set; }

        [Input("originKeepaliveTimeout")]
        public Input<int>? OriginKeepaliveTimeout { get; set; }

        [Input("originProtocolPolicy", required: true)]
        public Input<string> OriginProtocolPolicy { get; set; } = null!;

        [Input("originReadTimeout")]
        public Input<int>? OriginReadTimeout { get; set; }

        [Input("originSSLProtocols")]
        private InputList<string>? _originSSLProtocols;
        public InputList<string> OriginSSLProtocols
        {
            get => _originSSLProtocols ?? (_originSSLProtocols = new InputList<string>());
            set => _originSSLProtocols = value;
        }

        public DistributionCustomOriginConfigArgs()
        {
        }
        public static new DistributionCustomOriginConfigArgs Empty => new DistributionCustomOriginConfigArgs();
    }
}
