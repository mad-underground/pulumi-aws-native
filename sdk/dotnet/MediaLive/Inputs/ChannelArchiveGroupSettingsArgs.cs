// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Inputs
{

    public sealed class ChannelArchiveGroupSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("archiveCdnSettings")]
        public Input<Inputs.ChannelArchiveCdnSettingsArgs>? ArchiveCdnSettings { get; set; }

        [Input("destination")]
        public Input<Inputs.ChannelOutputLocationRefArgs>? Destination { get; set; }

        [Input("rolloverInterval")]
        public Input<int>? RolloverInterval { get; set; }

        public ChannelArchiveGroupSettingsArgs()
        {
        }
        public static new ChannelArchiveGroupSettingsArgs Empty => new ChannelArchiveGroupSettingsArgs();
    }
}
