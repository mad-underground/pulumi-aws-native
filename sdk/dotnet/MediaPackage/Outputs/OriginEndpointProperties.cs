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
    public sealed class OriginEndpointProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-originendpoint.html#cfn-mediapackage-originendpoint-authorization
        /// </summary>
        public readonly Outputs.OriginEndpointAuthorization? Authorization;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-originendpoint.html#cfn-mediapackage-originendpoint-channelid
        /// </summary>
        public readonly string ChannelId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-originendpoint.html#cfn-mediapackage-originendpoint-cmafpackage
        /// </summary>
        public readonly Outputs.OriginEndpointCmafPackage? CmafPackage;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-originendpoint.html#cfn-mediapackage-originendpoint-dashpackage
        /// </summary>
        public readonly Outputs.OriginEndpointDashPackage? DashPackage;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-originendpoint.html#cfn-mediapackage-originendpoint-description
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-originendpoint.html#cfn-mediapackage-originendpoint-hlspackage
        /// </summary>
        public readonly Outputs.OriginEndpointHlsPackage? HlsPackage;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-originendpoint.html#cfn-mediapackage-originendpoint-id
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-originendpoint.html#cfn-mediapackage-originendpoint-manifestname
        /// </summary>
        public readonly string? ManifestName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-originendpoint.html#cfn-mediapackage-originendpoint-msspackage
        /// </summary>
        public readonly Outputs.OriginEndpointMssPackage? MssPackage;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-originendpoint.html#cfn-mediapackage-originendpoint-origination
        /// </summary>
        public readonly string? Origination;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-originendpoint.html#cfn-mediapackage-originendpoint-startoverwindowseconds
        /// </summary>
        public readonly int? StartoverWindowSeconds;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-originendpoint.html#cfn-mediapackage-originendpoint-tags
        /// </summary>
        public readonly ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-originendpoint.html#cfn-mediapackage-originendpoint-timedelayseconds
        /// </summary>
        public readonly int? TimeDelaySeconds;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-originendpoint.html#cfn-mediapackage-originendpoint-whitelist
        /// </summary>
        public readonly ImmutableArray<string> Whitelist;

        [OutputConstructor]
        private OriginEndpointProperties(
            Outputs.OriginEndpointAuthorization? Authorization,

            string ChannelId,

            Outputs.OriginEndpointCmafPackage? CmafPackage,

            Outputs.OriginEndpointDashPackage? DashPackage,

            string? Description,

            Outputs.OriginEndpointHlsPackage? HlsPackage,

            string Id,

            string? ManifestName,

            Outputs.OriginEndpointMssPackage? MssPackage,

            string? Origination,

            int? StartoverWindowSeconds,

            ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags,

            int? TimeDelaySeconds,

            ImmutableArray<string> Whitelist)
        {
            this.Authorization = Authorization;
            this.ChannelId = ChannelId;
            this.CmafPackage = CmafPackage;
            this.DashPackage = DashPackage;
            this.Description = Description;
            this.HlsPackage = HlsPackage;
            this.Id = Id;
            this.ManifestName = ManifestName;
            this.MssPackage = MssPackage;
            this.Origination = Origination;
            this.StartoverWindowSeconds = StartoverWindowSeconds;
            this.Tags = Tags;
            this.TimeDelaySeconds = TimeDelaySeconds;
            this.Whitelist = Whitelist;
        }
    }
}
