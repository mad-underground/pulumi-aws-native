// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ResourceExplorer2
{
    /// <summary>
    /// Definition of AWS::ResourceExplorer2::View Resource Type
    /// </summary>
    [AwsNativeResourceType("aws-native:resourceexplorer2:View")]
    public partial class View : global::Pulumi.CustomResource
    {
        [Output("filters")]
        public Output<Outputs.ViewSearchFilter?> Filters { get; private set; } = null!;

        [Output("includedProperties")]
        public Output<ImmutableArray<Outputs.ViewIncludedProperty>> IncludedProperties { get; private set; } = null!;

        [Output("scope")]
        public Output<string?> Scope { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("viewArn")]
        public Output<string> ViewArn { get; private set; } = null!;

        [Output("viewName")]
        public Output<string> ViewName { get; private set; } = null!;


        /// <summary>
        /// Create a View resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public View(string name, ViewArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:resourceexplorer2:View", name, args ?? new ViewArgs(), MakeResourceOptions(options, ""))
        {
        }

        private View(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:resourceexplorer2:View", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "scope",
                    "viewName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing View resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static View Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new View(name, id, options);
        }
    }

    public sealed class ViewArgs : global::Pulumi.ResourceArgs
    {
        [Input("filters")]
        public Input<Inputs.ViewSearchFilterArgs>? Filters { get; set; }

        [Input("includedProperties")]
        private InputList<Inputs.ViewIncludedPropertyArgs>? _includedProperties;
        public InputList<Inputs.ViewIncludedPropertyArgs> IncludedProperties
        {
            get => _includedProperties ?? (_includedProperties = new InputList<Inputs.ViewIncludedPropertyArgs>());
            set => _includedProperties = value;
        }

        [Input("scope")]
        public Input<string>? Scope { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("viewName")]
        public Input<string>? ViewName { get; set; }

        public ViewArgs()
        {
        }
        public static new ViewArgs Empty => new ViewArgs();
    }
}
