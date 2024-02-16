// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppConfig
{
    /// <summary>
    /// Resource Type definition for AWS::AppConfig::Extension
    /// </summary>
    [AwsNativeResourceType("aws-native:appconfig:Extension")]
    public partial class Extension : global::Pulumi.CustomResource
    {
        [Output("actions")]
        public Output<ImmutableDictionary<string, ImmutableArray<Outputs.ExtensionAction>>> Actions { get; private set; } = null!;

        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Description of the extension.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("latestVersionNumber")]
        public Output<int?> LatestVersionNumber { get; private set; } = null!;

        /// <summary>
        /// Name of the extension.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("parameters")]
        public Output<ImmutableDictionary<string, Outputs.ExtensionParameter>?> Parameters { get; private set; } = null!;

        /// <summary>
        /// An array of key-value tags to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.ExtensionTag>> Tags { get; private set; } = null!;

        [Output("versionNumber")]
        public Output<int> VersionNumber { get; private set; } = null!;


        /// <summary>
        /// Create a Extension resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Extension(string name, ExtensionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:appconfig:Extension", name, args ?? new ExtensionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Extension(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:appconfig:Extension", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "name",
                    "tags[*]",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Extension resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Extension Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Extension(name, id, options);
        }
    }

    public sealed class ExtensionArgs : global::Pulumi.ResourceArgs
    {
        [Input("actions", required: true)]
        private InputMap<ImmutableArray<Inputs.ExtensionActionArgs>>? _actions;
        public InputMap<ImmutableArray<Inputs.ExtensionActionArgs>> Actions
        {
            get => _actions ?? (_actions = new InputMap<ImmutableArray<Inputs.ExtensionActionArgs>>());
            set => _actions = value;
        }

        /// <summary>
        /// Description of the extension.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("latestVersionNumber")]
        public Input<int>? LatestVersionNumber { get; set; }

        /// <summary>
        /// Name of the extension.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameters")]
        private InputMap<Inputs.ExtensionParameterArgs>? _parameters;
        public InputMap<Inputs.ExtensionParameterArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<Inputs.ExtensionParameterArgs>());
            set => _parameters = value;
        }

        [Input("tags")]
        private InputList<Inputs.ExtensionTagArgs>? _tags;

        /// <summary>
        /// An array of key-value tags to apply to this resource.
        /// </summary>
        public InputList<Inputs.ExtensionTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ExtensionTagArgs>());
            set => _tags = value;
        }

        public ExtensionArgs()
        {
        }
        public static new ExtensionArgs Empty => new ExtensionArgs();
    }
}
