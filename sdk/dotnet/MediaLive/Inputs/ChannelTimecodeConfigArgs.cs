// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Inputs
{

    public sealed class ChannelTimecodeConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("source")]
        public Input<string>? Source { get; set; }

        [Input("syncThreshold")]
        public Input<int>? SyncThreshold { get; set; }

        public ChannelTimecodeConfigArgs()
        {
        }
        public static new ChannelTimecodeConfigArgs Empty => new ChannelTimecodeConfigArgs();
    }
}
