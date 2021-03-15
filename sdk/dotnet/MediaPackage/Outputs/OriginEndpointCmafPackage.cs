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
    public sealed class OriginEndpointCmafPackage
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-cmafpackage.html#cfn-mediapackage-originendpoint-cmafpackage-encryption
        /// </summary>
        public readonly Outputs.OriginEndpointCmafEncryption? Encryption;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-cmafpackage.html#cfn-mediapackage-originendpoint-cmafpackage-hlsmanifests
        /// </summary>
        public readonly ImmutableArray<Outputs.OriginEndpointHlsManifest> HlsManifests;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-cmafpackage.html#cfn-mediapackage-originendpoint-cmafpackage-segmentdurationseconds
        /// </summary>
        public readonly int? SegmentDurationSeconds;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-cmafpackage.html#cfn-mediapackage-originendpoint-cmafpackage-segmentprefix
        /// </summary>
        public readonly string? SegmentPrefix;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-cmafpackage.html#cfn-mediapackage-originendpoint-cmafpackage-streamselection
        /// </summary>
        public readonly Outputs.OriginEndpointStreamSelection? StreamSelection;

        [OutputConstructor]
        private OriginEndpointCmafPackage(
            Outputs.OriginEndpointCmafEncryption? encryption,

            ImmutableArray<Outputs.OriginEndpointHlsManifest> hlsManifests,

            int? segmentDurationSeconds,

            string? segmentPrefix,

            Outputs.OriginEndpointStreamSelection? streamSelection)
        {
            Encryption = encryption;
            HlsManifests = hlsManifests;
            SegmentDurationSeconds = segmentDurationSeconds;
            SegmentPrefix = segmentPrefix;
            StreamSelection = streamSelection;
        }
    }
}
