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
    public sealed class PackagingConfigurationStreamSelection
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-streamselection.html#cfn-mediapackage-packagingconfiguration-streamselection-maxvideobitspersecond
        /// </summary>
        public readonly int? MaxVideoBitsPerSecond;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-streamselection.html#cfn-mediapackage-packagingconfiguration-streamselection-minvideobitspersecond
        /// </summary>
        public readonly int? MinVideoBitsPerSecond;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-streamselection.html#cfn-mediapackage-packagingconfiguration-streamselection-streamorder
        /// </summary>
        public readonly string? StreamOrder;

        [OutputConstructor]
        private PackagingConfigurationStreamSelection(
            int? maxVideoBitsPerSecond,

            int? minVideoBitsPerSecond,

            string? streamOrder)
        {
            MaxVideoBitsPerSecond = maxVideoBitsPerSecond;
            MinVideoBitsPerSecond = minVideoBitsPerSecond;
            StreamOrder = streamOrder;
        }
    }
}
