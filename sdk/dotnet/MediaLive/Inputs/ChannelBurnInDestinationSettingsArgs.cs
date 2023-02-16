// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Inputs
{

    public sealed class ChannelBurnInDestinationSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("alignment")]
        public Input<string>? Alignment { get; set; }

        [Input("backgroundColor")]
        public Input<string>? BackgroundColor { get; set; }

        [Input("backgroundOpacity")]
        public Input<int>? BackgroundOpacity { get; set; }

        [Input("font")]
        public Input<Inputs.ChannelInputLocationArgs>? Font { get; set; }

        [Input("fontColor")]
        public Input<string>? FontColor { get; set; }

        [Input("fontOpacity")]
        public Input<int>? FontOpacity { get; set; }

        [Input("fontResolution")]
        public Input<int>? FontResolution { get; set; }

        [Input("fontSize")]
        public Input<string>? FontSize { get; set; }

        [Input("outlineColor")]
        public Input<string>? OutlineColor { get; set; }

        [Input("outlineSize")]
        public Input<int>? OutlineSize { get; set; }

        [Input("shadowColor")]
        public Input<string>? ShadowColor { get; set; }

        [Input("shadowOpacity")]
        public Input<int>? ShadowOpacity { get; set; }

        [Input("shadowXOffset")]
        public Input<int>? ShadowXOffset { get; set; }

        [Input("shadowYOffset")]
        public Input<int>? ShadowYOffset { get; set; }

        [Input("teletextGridControl")]
        public Input<string>? TeletextGridControl { get; set; }

        [Input("xPosition")]
        public Input<int>? XPosition { get; set; }

        [Input("yPosition")]
        public Input<int>? YPosition { get; set; }

        public ChannelBurnInDestinationSettingsArgs()
        {
        }
        public static new ChannelBurnInDestinationSettingsArgs Empty => new ChannelBurnInDestinationSettingsArgs();
    }
}
