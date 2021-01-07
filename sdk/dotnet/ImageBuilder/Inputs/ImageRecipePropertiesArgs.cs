// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.ImageBuilder.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-imagerecipe.html
    /// </summary>
    public sealed class ImageRecipePropertiesArgs : Pulumi.ResourceArgs
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

        public ImageRecipePropertiesArgs()
        {
        }
    }
}