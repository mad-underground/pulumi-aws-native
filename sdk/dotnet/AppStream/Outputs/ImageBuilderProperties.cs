// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.AppStream.Outputs
{

    [OutputType]
    public sealed class ImageBuilderProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-accessendpoints
        /// </summary>
        public readonly ImmutableArray<Outputs.ImageBuilderAccessEndpoint> AccessEndpoints;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-appstreamagentversion
        /// </summary>
        public readonly string? AppstreamAgentVersion;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-description
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-displayname
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-domainjoininfo
        /// </summary>
        public readonly Outputs.ImageBuilderDomainJoinInfo? DomainJoinInfo;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-enabledefaultinternetaccess
        /// </summary>
        public readonly bool? EnableDefaultInternetAccess;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-iamrolearn
        /// </summary>
        public readonly string? IamRoleArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-imagearn
        /// </summary>
        public readonly string? ImageArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-imagename
        /// </summary>
        public readonly string? ImageName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-instancetype
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-tags
        /// </summary>
        public readonly ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-vpcconfig
        /// </summary>
        public readonly Outputs.ImageBuilderVpcConfig? VpcConfig;

        [OutputConstructor]
        private ImageBuilderProperties(
            ImmutableArray<Outputs.ImageBuilderAccessEndpoint> AccessEndpoints,

            string? AppstreamAgentVersion,

            string? Description,

            string? DisplayName,

            Outputs.ImageBuilderDomainJoinInfo? DomainJoinInfo,

            bool? EnableDefaultInternetAccess,

            string? IamRoleArn,

            string? ImageArn,

            string? ImageName,

            string InstanceType,

            string Name,

            ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags,

            Outputs.ImageBuilderVpcConfig? VpcConfig)
        {
            this.AccessEndpoints = AccessEndpoints;
            this.AppstreamAgentVersion = AppstreamAgentVersion;
            this.Description = Description;
            this.DisplayName = DisplayName;
            this.DomainJoinInfo = DomainJoinInfo;
            this.EnableDefaultInternetAccess = EnableDefaultInternetAccess;
            this.IamRoleArn = IamRoleArn;
            this.ImageArn = ImageArn;
            this.ImageName = ImageName;
            this.InstanceType = InstanceType;
            this.Name = Name;
            this.Tags = Tags;
            this.VpcConfig = VpcConfig;
        }
    }
}