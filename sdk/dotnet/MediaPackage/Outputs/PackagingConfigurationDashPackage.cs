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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-dashpackage.html
    /// </summary>
    [OutputType]
    public sealed class PackagingConfigurationDashPackage
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-dashpackage.html#cfn-mediapackage-packagingconfiguration-dashpackage-dashmanifests
        /// </summary>
        public readonly ImmutableArray<Outputs.PackagingConfigurationDashManifest> DashManifests;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-dashpackage.html#cfn-mediapackage-packagingconfiguration-dashpackage-encryption
        /// </summary>
        public readonly Outputs.PackagingConfigurationDashEncryption? Encryption;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-dashpackage.html#cfn-mediapackage-packagingconfiguration-dashpackage-periodtriggers
        /// </summary>
        public readonly ImmutableArray<string> PeriodTriggers;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-dashpackage.html#cfn-mediapackage-packagingconfiguration-dashpackage-segmentdurationseconds
        /// </summary>
        public readonly int? SegmentDurationSeconds;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-dashpackage.html#cfn-mediapackage-packagingconfiguration-dashpackage-segmenttemplateformat
        /// </summary>
        public readonly string? SegmentTemplateFormat;

        [OutputConstructor]
        private PackagingConfigurationDashPackage(
            ImmutableArray<Outputs.PackagingConfigurationDashManifest> dashManifests,

            Outputs.PackagingConfigurationDashEncryption? encryption,

            ImmutableArray<string> periodTriggers,

            int? segmentDurationSeconds,

            string? segmentTemplateFormat)
        {
            DashManifests = dashManifests;
            Encryption = encryption;
            PeriodTriggers = periodTriggers;
            SegmentDurationSeconds = segmentDurationSeconds;
            SegmentTemplateFormat = segmentTemplateFormat;
        }
    }
}
