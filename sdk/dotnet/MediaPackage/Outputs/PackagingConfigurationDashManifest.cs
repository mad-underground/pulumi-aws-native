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
    public sealed class PackagingConfigurationDashManifest
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-dashmanifest.html#cfn-mediapackage-packagingconfiguration-dashmanifest-manifestlayout
        /// </summary>
        public readonly string? ManifestLayout;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-dashmanifest.html#cfn-mediapackage-packagingconfiguration-dashmanifest-manifestname
        /// </summary>
        public readonly string? ManifestName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-dashmanifest.html#cfn-mediapackage-packagingconfiguration-dashmanifest-minbuffertimeseconds
        /// </summary>
        public readonly int? MinBufferTimeSeconds;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-dashmanifest.html#cfn-mediapackage-packagingconfiguration-dashmanifest-profile
        /// </summary>
        public readonly string? Profile;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-dashmanifest.html#cfn-mediapackage-packagingconfiguration-dashmanifest-streamselection
        /// </summary>
        public readonly Outputs.PackagingConfigurationStreamSelection? StreamSelection;

        [OutputConstructor]
        private PackagingConfigurationDashManifest(
            string? manifestLayout,

            string? manifestName,

            int? minBufferTimeSeconds,

            string? profile,

            Outputs.PackagingConfigurationStreamSelection? streamSelection)
        {
            ManifestLayout = manifestLayout;
            ManifestName = manifestName;
            MinBufferTimeSeconds = minBufferTimeSeconds;
            Profile = profile;
            StreamSelection = streamSelection;
        }
    }
}
