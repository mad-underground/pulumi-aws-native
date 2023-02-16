// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Personalize
{
    /// <summary>
    /// Resource schema for AWS::Personalize::Schema.
    /// </summary>
    [AwsNativeResourceType("aws-native:personalize:Schema")]
    public partial class Schema : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The domain of a Domain dataset group.
        /// </summary>
        [Output("domain")]
        public Output<Pulumi.AwsNative.Personalize.SchemaDomain?> Domain { get; private set; } = null!;

        /// <summary>
        /// Name for the schema.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A schema in Avro JSON format.
        /// </summary>
        [Output("schema")]
        public Output<string> SchemaValue { get; private set; } = null!;

        /// <summary>
        /// Arn for the schema.
        /// </summary>
        [Output("schemaArn")]
        public Output<string> SchemaArn { get; private set; } = null!;


        /// <summary>
        /// Create a Schema resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Schema(string name, SchemaArgs args, CustomResourceOptions? options = null)
            : base("aws-native:personalize:Schema", name, args ?? new SchemaArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Schema(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:personalize:Schema", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Schema resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Schema Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Schema(name, id, options);
        }
    }

    public sealed class SchemaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain of a Domain dataset group.
        /// </summary>
        [Input("domain")]
        public Input<Pulumi.AwsNative.Personalize.SchemaDomain>? Domain { get; set; }

        /// <summary>
        /// Name for the schema.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A schema in Avro JSON format.
        /// </summary>
        [Input("schema", required: true)]
        public Input<string> SchemaValue { get; set; } = null!;

        public SchemaArgs()
        {
        }
        public static new SchemaArgs Empty => new SchemaArgs();
    }
}
