// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.IoTSiteWise.Outputs
{

    [OutputType]
    public sealed class AssetModelProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodeldescription
        /// </summary>
        public readonly string? AssetModelDescription;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelhierarchies
        /// </summary>
        public readonly ImmutableArray<Outputs.AssetModelAssetModelHierarchy> AssetModelHierarchies;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelname
        /// </summary>
        public readonly string AssetModelName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelproperties
        /// </summary>
        public readonly ImmutableArray<Outputs.AssetModelAssetModelProperty> Properties;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-tags
        /// </summary>
        public readonly ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags;

        [OutputConstructor]
        private AssetModelProperties(
            string? AssetModelDescription,

            ImmutableArray<Outputs.AssetModelAssetModelHierarchy> AssetModelHierarchies,

            string AssetModelName,

            ImmutableArray<Outputs.AssetModelAssetModelProperty> AssetModelProperties,

            ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags)
        {
            this.AssetModelDescription = AssetModelDescription;
            this.AssetModelHierarchies = AssetModelHierarchies;
            this.AssetModelName = AssetModelName;
            Properties = AssetModelProperties;
            this.Tags = Tags;
        }
    }
}