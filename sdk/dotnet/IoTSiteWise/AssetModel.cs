// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTSiteWise
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html
    /// </summary>
    [AwsNativeResourceType("aws-native:IoTSiteWise:AssetModel")]
    public partial class AssetModel : Pulumi.CustomResource
    {
        [Output("assetModelArn")]
        public Output<string> AssetModelArn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodeldescription
        /// </summary>
        [Output("assetModelDescription")]
        public Output<string?> AssetModelDescription { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelhierarchies
        /// </summary>
        [Output("assetModelHierarchies")]
        public Output<ImmutableArray<Outputs.AssetModelAssetModelHierarchy>> AssetModelHierarchies { get; private set; } = null!;

        [Output("assetModelId")]
        public Output<string> AssetModelId { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelname
        /// </summary>
        [Output("assetModelName")]
        public Output<string> AssetModelName { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelproperties
        /// </summary>
        [Output("assetModelProperties")]
        public Output<ImmutableArray<Outputs.AssetModelAssetModelProperty>> AssetModelProperties { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a AssetModel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AssetModel(string name, AssetModelArgs args, CustomResourceOptions? options = null)
            : base("aws-native:IoTSiteWise:AssetModel", name, args ?? new AssetModelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AssetModel(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:IoTSiteWise:AssetModel", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AssetModel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AssetModel Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AssetModel(name, id, options);
        }
    }

    public sealed class AssetModelArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodeldescription
        /// </summary>
        [Input("assetModelDescription")]
        public Input<string>? AssetModelDescription { get; set; }

        [Input("assetModelHierarchies")]
        private InputList<Inputs.AssetModelAssetModelHierarchyArgs>? _assetModelHierarchies;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelhierarchies
        /// </summary>
        public InputList<Inputs.AssetModelAssetModelHierarchyArgs> AssetModelHierarchies
        {
            get => _assetModelHierarchies ?? (_assetModelHierarchies = new InputList<Inputs.AssetModelAssetModelHierarchyArgs>());
            set => _assetModelHierarchies = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelname
        /// </summary>
        [Input("assetModelName", required: true)]
        public Input<string> AssetModelName { get; set; } = null!;

        [Input("assetModelProperties")]
        private InputList<Inputs.AssetModelAssetModelPropertyArgs>? _assetModelProperties;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelproperties
        /// </summary>
        public InputList<Inputs.AssetModelAssetModelPropertyArgs> AssetModelProperties
        {
            get => _assetModelProperties ?? (_assetModelProperties = new InputList<Inputs.AssetModelAssetModelPropertyArgs>());
            set => _assetModelProperties = value;
        }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-tags
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public AssetModelArgs()
        {
        }
    }
}
