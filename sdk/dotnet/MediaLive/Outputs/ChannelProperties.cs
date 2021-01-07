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
    public sealed class ChannelProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-channelclass
        /// </summary>
        public readonly string? ChannelClass;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-destinations
        /// </summary>
        public readonly ImmutableArray<Outputs.ChannelOutputDestination> Destinations;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-encodersettings
        /// </summary>
        public readonly Outputs.ChannelEncoderSettings? EncoderSettings;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-inputattachments
        /// </summary>
        public readonly ImmutableArray<Outputs.ChannelInputAttachment> InputAttachments;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-inputspecification
        /// </summary>
        public readonly Outputs.ChannelInputSpecification? InputSpecification;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-loglevel
        /// </summary>
        public readonly string? LogLevel;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-name
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-rolearn
        /// </summary>
        public readonly string? RoleArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-tags
        /// </summary>
        public readonly Union<System.Text.Json.JsonElement, string>? Tags;

        [OutputConstructor]
        private ChannelProperties(
            string? ChannelClass,

            ImmutableArray<Outputs.ChannelOutputDestination> Destinations,

            Outputs.ChannelEncoderSettings? EncoderSettings,

            ImmutableArray<Outputs.ChannelInputAttachment> InputAttachments,

            Outputs.ChannelInputSpecification? InputSpecification,

            string? LogLevel,

            string? Name,

            string? RoleArn,

            Union<System.Text.Json.JsonElement, string>? Tags)
        {
            this.ChannelClass = ChannelClass;
            this.Destinations = Destinations;
            this.EncoderSettings = EncoderSettings;
            this.InputAttachments = InputAttachments;
            this.InputSpecification = InputSpecification;
            this.LogLevel = LogLevel;
            this.Name = Name;
            this.RoleArn = RoleArn;
            this.Tags = Tags;
        }
    }
}
