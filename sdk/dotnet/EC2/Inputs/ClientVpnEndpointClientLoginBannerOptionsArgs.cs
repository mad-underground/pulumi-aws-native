// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Inputs
{

    public sealed class ClientVpnEndpointClientLoginBannerOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("bannerText")]
        public Input<string>? BannerText { get; set; }

        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        public ClientVpnEndpointClientLoginBannerOptionsArgs()
        {
        }
        public static new ClientVpnEndpointClientLoginBannerOptionsArgs Empty => new ClientVpnEndpointClientLoginBannerOptionsArgs();
    }
}
