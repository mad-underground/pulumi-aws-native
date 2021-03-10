// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GroundStation
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-config.html
    /// </summary>
    [AwsNativeResourceType("aws-native:GroundStation:Config")]
    public partial class Config : Pulumi.CustomResource
    {
        [Output("Arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-config.html#cfn-groundstation-config-configdata
        /// </summary>
        [Output("ConfigData")]
        public Output<Union<System.Text.Json.JsonElement, string>> ConfigData { get; private set; } = null!;

        [Output("Id")]
        public Output<string> Id { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-config.html#cfn-groundstation-config-name
        /// </summary>
        [Output("Name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-config.html#cfn-groundstation-config-tags
        /// </summary>
        [Output("Tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        [Output("Type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Config resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Config(string name, ConfigArgs args, CustomResourceOptions? options = null)
            : base("aws-native:GroundStation:Config", name, args ?? new ConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Config(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:GroundStation:Config", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Config resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Config Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Config(name, id, options);
        }
    }

    public sealed class ConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-config.html#cfn-groundstation-config-configdata
        /// </summary>
        [Input("ConfigData", required: true)]
        public InputUnion<System.Text.Json.JsonElement, string> ConfigData { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-config.html#cfn-groundstation-config-name
        /// </summary>
        [Input("Name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("Tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _Tags;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-config.html#cfn-groundstation-config-tags
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _Tags ?? (_Tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _Tags = value;
        }

        public ConfigArgs()
        {
        }
    }
}
