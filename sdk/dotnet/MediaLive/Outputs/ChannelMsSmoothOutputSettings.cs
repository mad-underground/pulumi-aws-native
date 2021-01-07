// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.MediaLive.Outputs
{

    [OutputType]
    public sealed class ChannelMsSmoothOutputSettings
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothoutputsettings.html#cfn-medialive-channel-mssmoothoutputsettings-h265packagingtype
        /// </summary>
        public readonly string? H265PackagingType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mssmoothoutputsettings.html#cfn-medialive-channel-mssmoothoutputsettings-namemodifier
        /// </summary>
        public readonly string? NameModifier;

        [OutputConstructor]
        private ChannelMsSmoothOutputSettings(
            string? H265PackagingType,

            string? NameModifier)
        {
            this.H265PackagingType = H265PackagingType;
            this.NameModifier = NameModifier;
        }
    }
}
