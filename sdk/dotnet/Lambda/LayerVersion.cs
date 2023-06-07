// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lambda
{
    /// <summary>
    /// Resource Type definition for AWS::Lambda::LayerVersion
    /// </summary>
    [AwsNativeResourceType("aws-native:lambda:LayerVersion")]
    public partial class LayerVersion : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of compatible function runtimes. Used for filtering with ListLayers and ListLayerVersions.
        /// </summary>
        [Output("compatibleRuntimes")]
        public Output<ImmutableArray<string>> CompatibleRuntimes { get; private set; } = null!;

        /// <summary>
        /// The function layer archive.
        /// </summary>
        [Output("content")]
        public Output<Outputs.LayerVersionContent> Content { get; private set; } = null!;

        /// <summary>
        /// The description of the version.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name or Amazon Resource Name (ARN) of the layer.
        /// </summary>
        [Output("layerName")]
        public Output<string?> LayerName { get; private set; } = null!;

        /// <summary>
        /// The layer's software license.
        /// </summary>
        [Output("licenseInfo")]
        public Output<string?> LicenseInfo { get; private set; } = null!;


        /// <summary>
        /// Create a LayerVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LayerVersion(string name, LayerVersionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:lambda:LayerVersion", name, args ?? new LayerVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LayerVersion(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:lambda:LayerVersion", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing LayerVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LayerVersion Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LayerVersion(name, id, options);
        }
    }

    public sealed class LayerVersionArgs : global::Pulumi.ResourceArgs
    {
        [Input("compatibleRuntimes")]
        private InputList<string>? _compatibleRuntimes;

        /// <summary>
        /// A list of compatible function runtimes. Used for filtering with ListLayers and ListLayerVersions.
        /// </summary>
        public InputList<string> CompatibleRuntimes
        {
            get => _compatibleRuntimes ?? (_compatibleRuntimes = new InputList<string>());
            set => _compatibleRuntimes = value;
        }

        /// <summary>
        /// The function layer archive.
        /// </summary>
        [Input("content", required: true)]
        public Input<Inputs.LayerVersionContentArgs> Content { get; set; } = null!;

        /// <summary>
        /// The description of the version.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name or Amazon Resource Name (ARN) of the layer.
        /// </summary>
        [Input("layerName")]
        public Input<string>? LayerName { get; set; }

        /// <summary>
        /// The layer's software license.
        /// </summary>
        [Input("licenseInfo")]
        public Input<string>? LicenseInfo { get; set; }

        public LayerVersionArgs()
        {
        }
        public static new LayerVersionArgs Empty => new LayerVersionArgs();
    }
}
