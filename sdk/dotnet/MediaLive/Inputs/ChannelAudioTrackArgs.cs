// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Inputs
{

    public sealed class ChannelAudioTrackArgs : global::Pulumi.ResourceArgs
    {
        [Input("track")]
        public Input<int>? Track { get; set; }

        public ChannelAudioTrackArgs()
        {
        }
        public static new ChannelAudioTrackArgs Empty => new ChannelAudioTrackArgs();
    }
}
