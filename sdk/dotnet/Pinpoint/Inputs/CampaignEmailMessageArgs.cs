// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint.Inputs
{

    public sealed class CampaignEmailMessageArgs : global::Pulumi.ResourceArgs
    {
        [Input("body")]
        public Input<string>? Body { get; set; }

        [Input("fromAddress")]
        public Input<string>? FromAddress { get; set; }

        [Input("htmlBody")]
        public Input<string>? HtmlBody { get; set; }

        [Input("title")]
        public Input<string>? Title { get; set; }

        public CampaignEmailMessageArgs()
        {
        }
        public static new CampaignEmailMessageArgs Empty => new CampaignEmailMessageArgs();
    }
}
