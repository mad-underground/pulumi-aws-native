// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ImageBuilder
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-imagerecipe.html
    /// </summary>
    [AwsNativeResourceType("aws-native:ImageBuilder:ImageRecipe")]
    public partial class ImageRecipe : Pulumi.CustomResource
    {
        [Output("Arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-imagerecipe.html#cfn-imagebuilder-imagerecipe-blockdevicemappings
        /// </summary>
        [Output("BlockDeviceMappings")]
        public Output<ImmutableArray<Outputs.ImageRecipeInstanceBlockDeviceMapping>> BlockDeviceMappings { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-imagerecipe.html#cfn-imagebuilder-imagerecipe-components
        /// </summary>
        [Output("Components")]
        public Output<ImmutableArray<Outputs.ImageRecipeComponentConfiguration>> Components { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-imagerecipe.html#cfn-imagebuilder-imagerecipe-description
        /// </summary>
        [Output("Description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-imagerecipe.html#cfn-imagebuilder-imagerecipe-name
        /// </summary>
        [Output("Name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-imagerecipe.html#cfn-imagebuilder-imagerecipe-parentimage
        /// </summary>
        [Output("ParentImage")]
        public Output<string> ParentImage { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-imagerecipe.html#cfn-imagebuilder-imagerecipe-tags
        /// </summary>
        [Output("Tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-imagerecipe.html#cfn-imagebuilder-imagerecipe-version
        /// </summary>
        [Output("Version")]
        public Output<string> Version { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-imagerecipe.html#cfn-imagebuilder-imagerecipe-workingdirectory
        /// </summary>
        [Output("WorkingDirectory")]
        public Output<string?> WorkingDirectory { get; private set; } = null!;


        /// <summary>
        /// Create a ImageRecipe resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ImageRecipe(string name, ImageRecipeArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ImageBuilder:ImageRecipe", name, args ?? new ImageRecipeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ImageRecipe(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ImageBuilder:ImageRecipe", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ImageRecipe resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ImageRecipe Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ImageRecipe(name, id, options);
        }
    }

    public sealed class ImageRecipeArgs : Pulumi.ResourceArgs
    {
        [Input("BlockDeviceMappings")]
        private InputList<Inputs.ImageRecipeInstanceBlockDeviceMappingArgs>? _BlockDeviceMappings;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-imagerecipe.html#cfn-imagebuilder-imagerecipe-blockdevicemappings
        /// </summary>
        public InputList<Inputs.ImageRecipeInstanceBlockDeviceMappingArgs> BlockDeviceMappings
        {
            get => _BlockDeviceMappings ?? (_BlockDeviceMappings = new InputList<Inputs.ImageRecipeInstanceBlockDeviceMappingArgs>());
            set => _BlockDeviceMappings = value;
        }

        [Input("Components", required: true)]
        private InputList<Inputs.ImageRecipeComponentConfigurationArgs>? _Components;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-imagerecipe.html#cfn-imagebuilder-imagerecipe-components
        /// </summary>
        public InputList<Inputs.ImageRecipeComponentConfigurationArgs> Components
        {
            get => _Components ?? (_Components = new InputList<Inputs.ImageRecipeComponentConfigurationArgs>());
            set => _Components = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-imagerecipe.html#cfn-imagebuilder-imagerecipe-description
        /// </summary>
        [Input("Description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-imagerecipe.html#cfn-imagebuilder-imagerecipe-name
        /// </summary>
        [Input("Name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-imagerecipe.html#cfn-imagebuilder-imagerecipe-parentimage
        /// </summary>
        [Input("ParentImage", required: true)]
        public Input<string> ParentImage { get; set; } = null!;

        [Input("Tags")]
        private InputMap<string>? _Tags;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-imagerecipe.html#cfn-imagebuilder-imagerecipe-tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _Tags ?? (_Tags = new InputMap<string>());
            set => _Tags = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-imagerecipe.html#cfn-imagebuilder-imagerecipe-version
        /// </summary>
        [Input("Version", required: true)]
        public Input<string> Version { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-imagerecipe.html#cfn-imagebuilder-imagerecipe-workingdirectory
        /// </summary>
        [Input("WorkingDirectory")]
        public Input<string>? WorkingDirectory { get; set; }

        public ImageRecipeArgs()
        {
        }
    }
}
