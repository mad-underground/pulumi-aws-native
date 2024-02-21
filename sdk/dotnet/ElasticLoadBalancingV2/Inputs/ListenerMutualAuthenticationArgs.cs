// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancingV2.Inputs
{

    public sealed class ListenerMutualAuthenticationArgs : global::Pulumi.ResourceArgs
    {
        [Input("ignoreClientCertificateExpiry")]
        public Input<bool>? IgnoreClientCertificateExpiry { get; set; }

        [Input("mode")]
        public Input<string>? Mode { get; set; }

        [Input("trustStoreArn")]
        public Input<string>? TrustStoreArn { get; set; }

        public ListenerMutualAuthenticationArgs()
        {
        }
        public static new ListenerMutualAuthenticationArgs Empty => new ListenerMutualAuthenticationArgs();
    }
}