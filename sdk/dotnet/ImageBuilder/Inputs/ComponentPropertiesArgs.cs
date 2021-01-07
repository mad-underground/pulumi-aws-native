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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html
    /// </summary>
    public sealed class ComponentPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-changedescription
        /// </summary>
        [Input("ChangeDescription")]
        public Input<string>? ChangeDescription { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-data
        /// </summary>
        [Input("Data")]
        public Input<string>? Data { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-description
        /// </summary>
        [Input("Description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-kmskeyid
        /// </summary>
        [Input("KmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-name
        /// </summary>
        [Input("Name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-platform
        /// </summary>
        [Input("Platform", required: true)]
        public Input<string> Platform { get; set; } = null!;

        [Input("SupportedOsVersions")]
        private InputList<string>? _SupportedOsVersions;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-supportedosversions
        /// </summary>
        public InputList<string> SupportedOsVersions
        {
            get => _SupportedOsVersions ?? (_SupportedOsVersions = new InputList<string>());
            set => _SupportedOsVersions = value;
        }

        [Input("Tags")]
        private InputMap<string>? _Tags;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _Tags ?? (_Tags = new InputMap<string>());
            set => _Tags = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-uri
        /// </summary>
        [Input("Uri")]
        public Input<string>? Uri { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-version
        /// </summary>
        [Input("Version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public ComponentPropertiesArgs()
        {
        }
    }
}
