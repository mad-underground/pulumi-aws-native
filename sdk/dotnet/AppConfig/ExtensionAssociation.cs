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
    /// An example resource schema demonstrating some basic constructs and validation rules.
    /// </summary>
    [AwsNativeResourceType("aws-native:appconfig:ExtensionAssociation")]
    public partial class ExtensionAssociation : global::Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("extensionArn")]
        public Output<string> ExtensionArn { get; private set; } = null!;

        [Output("extensionIdentifier")]
        public Output<string?> ExtensionIdentifier { get; private set; } = null!;

        [Output("extensionVersionNumber")]
        public Output<int?> ExtensionVersionNumber { get; private set; } = null!;

        [Output("parameters")]
        public Output<ImmutableDictionary<string, string>?> Parameters { get; private set; } = null!;

        [Output("resourceArn")]
        public Output<string> ResourceArn { get; private set; } = null!;

        [Output("resourceIdentifier")]
        public Output<string?> ResourceIdentifier { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.ExtensionAssociationTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ExtensionAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ExtensionAssociation(string name, ExtensionAssociationArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:appconfig:ExtensionAssociation", name, args ?? new ExtensionAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ExtensionAssociation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:appconfig:ExtensionAssociation", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "extensionIdentifier",
                    "extensionVersionNumber",
                    "resourceIdentifier",
                    "tags[*]",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ExtensionAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ExtensionAssociation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ExtensionAssociation(name, id, options);
        }
    }

    public sealed class ExtensionAssociationArgs : global::Pulumi.ResourceArgs
    {
        [Input("extensionIdentifier")]
        public Input<string>? ExtensionIdentifier { get; set; }

        [Input("extensionVersionNumber")]
        public Input<int>? ExtensionVersionNumber { get; set; }

        [Input("parameters")]
        private InputMap<string>? _parameters;
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        [Input("resourceIdentifier")]
        public Input<string>? ResourceIdentifier { get; set; }

        [Input("tags")]
        private InputList<Inputs.ExtensionAssociationTagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Inputs.ExtensionAssociationTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ExtensionAssociationTagArgs>());
            set => _tags = value;
        }

        public ExtensionAssociationArgs()
        {
        }
        public static new ExtensionAssociationArgs Empty => new ExtensionAssociationArgs();
    }
}
