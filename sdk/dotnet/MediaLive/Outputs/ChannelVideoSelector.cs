// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Outputs
{

    [OutputType]
    public sealed class ChannelVideoSelector
    {
        public readonly string? ColorSpace;
        public readonly Outputs.ChannelVideoSelectorColorSpaceSettings? ColorSpaceSettings;
        public readonly string? ColorSpaceUsage;
        public readonly Outputs.ChannelVideoSelectorSettings? SelectorSettings;

        [OutputConstructor]
        private ChannelVideoSelector(
            string? colorSpace,

            Outputs.ChannelVideoSelectorColorSpaceSettings? colorSpaceSettings,

            string? colorSpaceUsage,

            Outputs.ChannelVideoSelectorSettings? selectorSettings)
        {
            ColorSpace = colorSpace;
            ColorSpaceSettings = colorSpaceSettings;
            ColorSpaceUsage = colorSpaceUsage;
            SelectorSettings = selectorSettings;
        }
    }
}
