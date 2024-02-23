// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ServiceCatalogAppRegistry
{
    /// <summary>
    /// Resource Schema for AWS::ServiceCatalogAppRegistry::AttributeGroup.
    /// </summary>
    [AwsNativeResourceType("aws-native:servicecatalogappregistry:AttributeGroup")]
    public partial class AttributeGroup : global::Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::ServiceCatalogAppRegistry::AttributeGroup` for more information about the expected schema for this property.
        /// </summary>
        [Output("attributes")]
        public Output<object> Attributes { get; private set; } = null!;

        /// <summary>
        /// The description of the attribute group. 
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the attribute group. 
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a AttributeGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AttributeGroup(string name, AttributeGroupArgs args, CustomResourceOptions? options = null)
            : base("aws-native:servicecatalogappregistry:AttributeGroup", name, args ?? new AttributeGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AttributeGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:servicecatalogappregistry:AttributeGroup", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AttributeGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AttributeGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AttributeGroup(name, id, options);
        }
    }

    public sealed class AttributeGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::ServiceCatalogAppRegistry::AttributeGroup` for more information about the expected schema for this property.
        /// </summary>
        [Input("attributes", required: true)]
        public Input<object> Attributes { get; set; } = null!;

        /// <summary>
        /// The description of the attribute group. 
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the attribute group. 
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public AttributeGroupArgs()
        {
        }
        public static new AttributeGroupArgs Empty => new AttributeGroupArgs();
    }
}
