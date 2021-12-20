// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Evidently
{
    /// <summary>
    /// Resource Type definition for AWS::Evidently::Feature.
    /// </summary>
    [AwsNativeResourceType("aws-native:evidently:Feature")]
    public partial class Feature : Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("defaultVariation")]
        public Output<string?> DefaultVariation { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("entityOverrides")]
        public Output<ImmutableArray<Outputs.FeatureEntityOverride>> EntityOverrides { get; private set; } = null!;

        [Output("evaluationStrategy")]
        public Output<Pulumi.AwsNative.Evidently.FeatureEvaluationStrategy?> EvaluationStrategy { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.FeatureTag>> Tags { get; private set; } = null!;

        [Output("variations")]
        public Output<ImmutableArray<Outputs.FeatureVariationObject>> Variations { get; private set; } = null!;


        /// <summary>
        /// Create a Feature resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Feature(string name, FeatureArgs args, CustomResourceOptions? options = null)
            : base("aws-native:evidently:Feature", name, args ?? new FeatureArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Feature(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:evidently:Feature", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Feature resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Feature Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Feature(name, id, options);
        }
    }

    public sealed class FeatureArgs : Pulumi.ResourceArgs
    {
        [Input("defaultVariation")]
        public Input<string>? DefaultVariation { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("entityOverrides")]
        private InputList<Inputs.FeatureEntityOverrideArgs>? _entityOverrides;
        public InputList<Inputs.FeatureEntityOverrideArgs> EntityOverrides
        {
            get => _entityOverrides ?? (_entityOverrides = new InputList<Inputs.FeatureEntityOverrideArgs>());
            set => _entityOverrides = value;
        }

        [Input("evaluationStrategy")]
        public Input<Pulumi.AwsNative.Evidently.FeatureEvaluationStrategy>? EvaluationStrategy { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.FeatureTagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Inputs.FeatureTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.FeatureTagArgs>());
            set => _tags = value;
        }

        [Input("variations", required: true)]
        private InputList<Inputs.FeatureVariationObjectArgs>? _variations;
        public InputList<Inputs.FeatureVariationObjectArgs> Variations
        {
            get => _variations ?? (_variations = new InputList<Inputs.FeatureVariationObjectArgs>());
            set => _variations = value;
        }

        public FeatureArgs()
        {
        }
    }
}