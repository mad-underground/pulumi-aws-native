// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Inputs
{

    public sealed class ChannelInputChannelLevelArgs : global::Pulumi.ResourceArgs
    {
        [Input("gain")]
        public Input<int>? Gain { get; set; }

        [Input("inputChannel")]
        public Input<int>? InputChannel { get; set; }

        public ChannelInputChannelLevelArgs()
        {
        }
        public static new ChannelInputChannelLevelArgs Empty => new ChannelInputChannelLevelArgs();
    }
}
