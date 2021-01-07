// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.MediaPackage.Outputs
{

    [OutputType]
    public sealed class PackagingConfigurationCmafPackage
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-cmafpackage.html#cfn-mediapackage-packagingconfiguration-cmafpackage-encryption
        /// </summary>
        public readonly Outputs.PackagingConfigurationCmafEncryption? Encryption;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-cmafpackage.html#cfn-mediapackage-packagingconfiguration-cmafpackage-hlsmanifests
        /// </summary>
        public readonly ImmutableArray<Outputs.PackagingConfigurationHlsManifest> HlsManifests;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-cmafpackage.html#cfn-mediapackage-packagingconfiguration-cmafpackage-segmentdurationseconds
        /// </summary>
        public readonly int? SegmentDurationSeconds;

        [OutputConstructor]
        private PackagingConfigurationCmafPackage(
            Outputs.PackagingConfigurationCmafEncryption? Encryption,

            ImmutableArray<Outputs.PackagingConfigurationHlsManifest> HlsManifests,

            int? SegmentDurationSeconds)
        {
            this.Encryption = Encryption;
            this.HlsManifests = HlsManifests;
            this.SegmentDurationSeconds = SegmentDurationSeconds;
        }
    }
}
