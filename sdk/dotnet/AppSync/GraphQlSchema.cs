// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync
{
    /// <summary>
    /// Resource Type definition for AWS::AppSync::GraphQLSchema
    /// </summary>
    [Obsolete(@"GraphQlSchema is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:appsync:GraphQlSchema")]
    public partial class GraphQlSchema : global::Pulumi.CustomResource
    {
        [Output("apiId")]
        public Output<string> ApiId { get; private set; } = null!;

        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        [Output("definition")]
        public Output<string?> Definition { get; private set; } = null!;

        [Output("definitionS3Location")]
        public Output<string?> DefinitionS3Location { get; private set; } = null!;


        /// <summary>
        /// Create a GraphQlSchema resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GraphQlSchema(string name, GraphQlSchemaArgs args, CustomResourceOptions? options = null)
            : base("aws-native:appsync:GraphQlSchema", name, args ?? new GraphQlSchemaArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GraphQlSchema(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:appsync:GraphQlSchema", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "apiId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing GraphQlSchema resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GraphQlSchema Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new GraphQlSchema(name, id, options);
        }
    }

    public sealed class GraphQlSchemaArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        [Input("definition")]
        public Input<string>? Definition { get; set; }

        [Input("definitionS3Location")]
        public Input<string>? DefinitionS3Location { get; set; }

        public GraphQlSchemaArgs()
        {
        }
        public static new GraphQlSchemaArgs Empty => new GraphQlSchemaArgs();
    }
}
