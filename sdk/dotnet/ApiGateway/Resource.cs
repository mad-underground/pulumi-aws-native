// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway
{
    /// <summary>
    /// The ``AWS::ApiGateway::Resource`` resource creates a resource in an API.
    /// </summary>
    [AwsNativeResourceType("aws-native:apigateway:Resource")]
    public partial class Resource : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The parent resource's identifier.
        /// </summary>
        [Output("parentId")]
        public Output<string> ParentId { get; private set; } = null!;

        /// <summary>
        /// The last path segment for this resource.
        /// </summary>
        [Output("pathPart")]
        public Output<string> PathPart { get; private set; } = null!;

        /// <summary>
        /// A unique primary identifier for a Resource
        /// </summary>
        [Output("resourceId")]
        public Output<string> ResourceId { get; private set; } = null!;

        /// <summary>
        /// The string identifier of the associated RestApi.
        /// </summary>
        [Output("restApiId")]
        public Output<string> RestApiId { get; private set; } = null!;


        /// <summary>
        /// Create a Resource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Resource(string name, ResourceArgs args, CustomResourceOptions? options = null)
            : base("aws-native:apigateway:Resource", name, args ?? new ResourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Resource(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:apigateway:Resource", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "parentId",
                    "pathPart",
                    "restApiId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Resource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Resource Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Resource(name, id, options);
        }
    }

    public sealed class ResourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The parent resource's identifier.
        /// </summary>
        [Input("parentId", required: true)]
        public Input<string> ParentId { get; set; } = null!;

        /// <summary>
        /// The last path segment for this resource.
        /// </summary>
        [Input("pathPart", required: true)]
        public Input<string> PathPart { get; set; } = null!;

        /// <summary>
        /// The string identifier of the associated RestApi.
        /// </summary>
        [Input("restApiId", required: true)]
        public Input<string> RestApiId { get; set; } = null!;

        public ResourceArgs()
        {
        }
        public static new ResourceArgs Empty => new ResourceArgs();
    }
}
