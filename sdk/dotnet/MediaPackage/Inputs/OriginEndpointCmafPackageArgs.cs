// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaPackage.Inputs
{

    /// <summary>
    /// A Common Media Application Format (CMAF) packaging configuration.
    /// </summary>
    public sealed class OriginEndpointCmafPackageArgs : global::Pulumi.ResourceArgs
    {
        [Input("encryption")]
        public Input<Inputs.OriginEndpointCmafEncryptionArgs>? Encryption { get; set; }

        [Input("hlsManifests")]
        private InputList<Inputs.OriginEndpointHlsManifestArgs>? _hlsManifests;

        /// <summary>
        /// A list of HLS manifest configurations
        /// </summary>
        public InputList<Inputs.OriginEndpointHlsManifestArgs> HlsManifests
        {
            get => _hlsManifests ?? (_hlsManifests = new InputList<Inputs.OriginEndpointHlsManifestArgs>());
            set => _hlsManifests = value;
        }

        /// <summary>
        /// Duration (in seconds) of each segment. Actual segments will be rounded to the nearest multiple of the source segment duration.
        /// </summary>
        [Input("segmentDurationSeconds")]
        public Input<int>? SegmentDurationSeconds { get; set; }

        /// <summary>
        /// An optional custom string that is prepended to the name of each segment. If not specified, it defaults to the ChannelId.
        /// </summary>
        [Input("segmentPrefix")]
        public Input<string>? SegmentPrefix { get; set; }

        [Input("streamSelection")]
        public Input<Inputs.OriginEndpointStreamSelectionArgs>? StreamSelection { get; set; }

        public OriginEndpointCmafPackageArgs()
        {
        }
        public static new OriginEndpointCmafPackageArgs Empty => new OriginEndpointCmafPackageArgs();
    }
}
