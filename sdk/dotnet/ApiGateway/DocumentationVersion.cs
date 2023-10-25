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
    /// The ``AWS::ApiGateway::DocumentationVersion`` resource creates a snapshot of the documentation for an API. For more information, see [Representation of API Documentation in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-documenting-api-content-representation.html) in the *API Gateway Developer Guide*.
    /// </summary>
    [AwsNativeResourceType("aws-native:apigateway:DocumentationVersion")]
    public partial class DocumentationVersion : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A description about the new documentation snapshot.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The version identifier of the to-be-updated documentation version.
        /// </summary>
        [Output("documentationVersion")]
        public Output<string> DocumentationVersionValue { get; private set; } = null!;

        /// <summary>
        /// The string identifier of the associated RestApi.
        /// </summary>
        [Output("restApiId")]
        public Output<string> RestApiId { get; private set; } = null!;


        /// <summary>
        /// Create a DocumentationVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DocumentationVersion(string name, DocumentationVersionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:apigateway:DocumentationVersion", name, args ?? new DocumentationVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DocumentationVersion(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:apigateway:DocumentationVersion", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "documentationVersion",
                    "restApiId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DocumentationVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DocumentationVersion Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DocumentationVersion(name, id, options);
        }
    }

    public sealed class DocumentationVersionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description about the new documentation snapshot.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The version identifier of the to-be-updated documentation version.
        /// </summary>
        [Input("documentationVersion", required: true)]
        public Input<string> DocumentationVersionValue { get; set; } = null!;

        /// <summary>
        /// The string identifier of the associated RestApi.
        /// </summary>
        [Input("restApiId", required: true)]
        public Input<string> RestApiId { get; set; } = null!;

        public DocumentationVersionArgs()
        {
        }
        public static new DocumentationVersionArgs Empty => new DocumentationVersionArgs();
    }
}
