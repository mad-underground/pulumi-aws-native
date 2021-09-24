// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Inputs
{

    public sealed class ChannelVideoSelectorArgs : Pulumi.ResourceArgs
    {
        [Input("colorSpace")]
        public Input<string>? ColorSpace { get; set; }

        [Input("colorSpaceSettings")]
        public Input<Inputs.ChannelVideoSelectorColorSpaceSettingsArgs>? ColorSpaceSettings { get; set; }

        [Input("colorSpaceUsage")]
        public Input<string>? ColorSpaceUsage { get; set; }

        [Input("selectorSettings")]
        public Input<Inputs.ChannelVideoSelectorSettingsArgs>? SelectorSettings { get; set; }

        public ChannelVideoSelectorArgs()
        {
        }
    }
}