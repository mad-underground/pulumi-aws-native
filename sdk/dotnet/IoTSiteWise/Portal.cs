// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.IoTSiteWise
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-portal.html
    /// </summary>
    [CloudformationResourceType("cloudformation:IoTSiteWise:Portal")]
    public partial class Portal : Pulumi.CustomResource
    {
        /// <summary>
        /// The attributes associated with the resource
        /// </summary>
        [Output("attributes")]
        public Output<Outputs.PortalAttributes> Attributes { get; private set; } = null!;

        /// <summary>
        /// An explicit logical ID for the resource
        /// </summary>
        [Output("logicalId")]
        public Output<string?> LogicalId { get; private set; } = null!;

        /// <summary>
        /// Arbitrary structured data associated with the resource
        /// </summary>
        [Output("metadata")]
        public Output<Union<System.Text.Json.JsonElement, string>?> Metadata { get; private set; } = null!;

        /// <summary>
        /// The input properties associated with the resource
        /// </summary>
        [Output("properties")]
        public Output<Outputs.PortalProperties> Properties { get; private set; } = null!;


        /// <summary>
        /// Create a Portal resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Portal(string name, PortalArgs args, CustomResourceOptions? options = null)
            : base("cloudformation:IoTSiteWise:Portal", name, args ?? new PortalArgs(), MakeResourceOptions(options, ""))
        {
        }
        internal Portal(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("cloudformation:IoTSiteWise:Portal", name, new DictionaryResourceArgs(dictionary), MakeResourceOptions(options, ""))
        {
        }

        private Portal(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("cloudformation:IoTSiteWise:Portal", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Portal resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Portal Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Portal(name, id, options);
        }
    }

    public sealed class PortalArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
        /// </summary>
        [Input("deletionPolicy")]
        public Input<string>? DeletionPolicy { get; set; }

        /// <summary>
        /// An explicit logical ID for the resource
        /// </summary>
        [Input("logicalId")]
        public Input<string>? LogicalId { get; set; }

        /// <summary>
        /// Arbitrary structured data associated with the resource
        /// </summary>
        [Input("metadata")]
        public InputUnion<System.Text.Json.JsonElement, string>? Metadata { get; set; }

        /// <summary>
        /// The input properties associated with the resource
        /// </summary>
        [Input("properties", required: true)]
        public Input<Inputs.PortalPropertiesArgs> Properties { get; set; } = null!;

        /// <summary>
        /// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
        /// </summary>
        [Input("updateReplacePolicy")]
        public Input<string>? UpdateReplacePolicy { get; set; }

        public PortalArgs()
        {
        }
    }
}
