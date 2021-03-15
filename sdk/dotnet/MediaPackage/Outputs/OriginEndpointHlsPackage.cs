// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaPackage.Outputs
{

    [OutputType]
    public sealed class OriginEndpointHlsPackage
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlspackage.html#cfn-mediapackage-originendpoint-hlspackage-admarkers
        /// </summary>
        public readonly string? AdMarkers;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlspackage.html#cfn-mediapackage-originendpoint-hlspackage-adtriggers
        /// </summary>
        public readonly ImmutableArray<string> AdTriggers;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlspackage.html#cfn-mediapackage-originendpoint-hlspackage-adsondeliveryrestrictions
        /// </summary>
        public readonly string? AdsOnDeliveryRestrictions;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlspackage.html#cfn-mediapackage-originendpoint-hlspackage-encryption
        /// </summary>
        public readonly Outputs.OriginEndpointHlsEncryption? Encryption;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlspackage.html#cfn-mediapackage-originendpoint-hlspackage-includeiframeonlystream
        /// </summary>
        public readonly bool? IncludeIframeOnlyStream;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlspackage.html#cfn-mediapackage-originendpoint-hlspackage-playlisttype
        /// </summary>
        public readonly string? PlaylistType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlspackage.html#cfn-mediapackage-originendpoint-hlspackage-playlistwindowseconds
        /// </summary>
        public readonly int? PlaylistWindowSeconds;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlspackage.html#cfn-mediapackage-originendpoint-hlspackage-programdatetimeintervalseconds
        /// </summary>
        public readonly int? ProgramDateTimeIntervalSeconds;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlspackage.html#cfn-mediapackage-originendpoint-hlspackage-segmentdurationseconds
        /// </summary>
        public readonly int? SegmentDurationSeconds;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlspackage.html#cfn-mediapackage-originendpoint-hlspackage-streamselection
        /// </summary>
        public readonly Outputs.OriginEndpointStreamSelection? StreamSelection;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-hlspackage.html#cfn-mediapackage-originendpoint-hlspackage-useaudiorenditiongroup
        /// </summary>
        public readonly bool? UseAudioRenditionGroup;

        [OutputConstructor]
        private OriginEndpointHlsPackage(
            string? adMarkers,

            ImmutableArray<string> adTriggers,

            string? adsOnDeliveryRestrictions,

            Outputs.OriginEndpointHlsEncryption? encryption,

            bool? includeIframeOnlyStream,

            string? playlistType,

            int? playlistWindowSeconds,

            int? programDateTimeIntervalSeconds,

            int? segmentDurationSeconds,

            Outputs.OriginEndpointStreamSelection? streamSelection,

            bool? useAudioRenditionGroup)
        {
            AdMarkers = adMarkers;
            AdTriggers = adTriggers;
            AdsOnDeliveryRestrictions = adsOnDeliveryRestrictions;
            Encryption = encryption;
            IncludeIframeOnlyStream = includeIframeOnlyStream;
            PlaylistType = playlistType;
            PlaylistWindowSeconds = playlistWindowSeconds;
            ProgramDateTimeIntervalSeconds = programDateTimeIntervalSeconds;
            SegmentDurationSeconds = segmentDurationSeconds;
            StreamSelection = streamSelection;
            UseAudioRenditionGroup = useAudioRenditionGroup;
        }
    }
}
