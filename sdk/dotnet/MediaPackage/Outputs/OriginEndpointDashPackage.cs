// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaPackage.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-dashpackage.html
    /// </summary>
    [OutputType]
    public sealed class OriginEndpointDashPackage
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-dashpackage.html#cfn-mediapackage-originendpoint-dashpackage-adtriggers
        /// </summary>
        public readonly ImmutableArray<string> AdTriggers;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-dashpackage.html#cfn-mediapackage-originendpoint-dashpackage-adsondeliveryrestrictions
        /// </summary>
        public readonly string? AdsOnDeliveryRestrictions;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-dashpackage.html#cfn-mediapackage-originendpoint-dashpackage-encryption
        /// </summary>
        public readonly Outputs.OriginEndpointDashEncryption? Encryption;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-dashpackage.html#cfn-mediapackage-originendpoint-dashpackage-manifestlayout
        /// </summary>
        public readonly string? ManifestLayout;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-dashpackage.html#cfn-mediapackage-originendpoint-dashpackage-manifestwindowseconds
        /// </summary>
        public readonly int? ManifestWindowSeconds;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-dashpackage.html#cfn-mediapackage-originendpoint-dashpackage-minbuffertimeseconds
        /// </summary>
        public readonly int? MinBufferTimeSeconds;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-dashpackage.html#cfn-mediapackage-originendpoint-dashpackage-minupdateperiodseconds
        /// </summary>
        public readonly int? MinUpdatePeriodSeconds;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-dashpackage.html#cfn-mediapackage-originendpoint-dashpackage-periodtriggers
        /// </summary>
        public readonly ImmutableArray<string> PeriodTriggers;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-dashpackage.html#cfn-mediapackage-originendpoint-dashpackage-profile
        /// </summary>
        public readonly string? Profile;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-dashpackage.html#cfn-mediapackage-originendpoint-dashpackage-segmentdurationseconds
        /// </summary>
        public readonly int? SegmentDurationSeconds;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-dashpackage.html#cfn-mediapackage-originendpoint-dashpackage-segmenttemplateformat
        /// </summary>
        public readonly string? SegmentTemplateFormat;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-dashpackage.html#cfn-mediapackage-originendpoint-dashpackage-streamselection
        /// </summary>
        public readonly Outputs.OriginEndpointStreamSelection? StreamSelection;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-dashpackage.html#cfn-mediapackage-originendpoint-dashpackage-suggestedpresentationdelayseconds
        /// </summary>
        public readonly int? SuggestedPresentationDelaySeconds;

        [OutputConstructor]
        private OriginEndpointDashPackage(
            ImmutableArray<string> adTriggers,

            string? adsOnDeliveryRestrictions,

            Outputs.OriginEndpointDashEncryption? encryption,

            string? manifestLayout,

            int? manifestWindowSeconds,

            int? minBufferTimeSeconds,

            int? minUpdatePeriodSeconds,

            ImmutableArray<string> periodTriggers,

            string? profile,

            int? segmentDurationSeconds,

            string? segmentTemplateFormat,

            Outputs.OriginEndpointStreamSelection? streamSelection,

            int? suggestedPresentationDelaySeconds)
        {
            AdTriggers = adTriggers;
            AdsOnDeliveryRestrictions = adsOnDeliveryRestrictions;
            Encryption = encryption;
            ManifestLayout = manifestLayout;
            ManifestWindowSeconds = manifestWindowSeconds;
            MinBufferTimeSeconds = minBufferTimeSeconds;
            MinUpdatePeriodSeconds = minUpdatePeriodSeconds;
            PeriodTriggers = periodTriggers;
            Profile = profile;
            SegmentDurationSeconds = segmentDurationSeconds;
            SegmentTemplateFormat = segmentTemplateFormat;
            StreamSelection = streamSelection;
            SuggestedPresentationDelaySeconds = suggestedPresentationDelaySeconds;
        }
    }
}
