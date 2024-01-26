// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra
{
    /// <summary>
    /// A Kendra FAQ resource
    /// </summary>
    [AwsNativeResourceType("aws-native:kendra:Faq")]
    public partial class Faq : global::Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// FAQ description
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// FAQ file format
        /// </summary>
        [Output("fileFormat")]
        public Output<Pulumi.AwsNative.Kendra.FaqFileFormat?> FileFormat { get; private set; } = null!;

        /// <summary>
        /// Index ID
        /// </summary>
        [Output("indexId")]
        public Output<string> IndexId { get; private set; } = null!;

        [Output("languageCode")]
        public Output<string?> LanguageCode { get; private set; } = null!;

        /// <summary>
        /// FAQ name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// FAQ role ARN
        /// </summary>
        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;

        /// <summary>
        /// FAQ S3 path
        /// </summary>
        [Output("s3Path")]
        public Output<Outputs.FaqS3Path> S3Path { get; private set; } = null!;

        /// <summary>
        /// Tags for labeling the FAQ
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.FaqTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Faq resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Faq(string name, FaqArgs args, CustomResourceOptions? options = null)
            : base("aws-native:kendra:Faq", name, args ?? new FaqArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Faq(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:kendra:Faq", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "description",
                    "fileFormat",
                    "indexId",
                    "name",
                    "roleArn",
                    "s3Path",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Faq resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Faq Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Faq(name, id, options);
        }
    }

    public sealed class FaqArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// FAQ description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// FAQ file format
        /// </summary>
        [Input("fileFormat")]
        public Input<Pulumi.AwsNative.Kendra.FaqFileFormat>? FileFormat { get; set; }

        /// <summary>
        /// Index ID
        /// </summary>
        [Input("indexId", required: true)]
        public Input<string> IndexId { get; set; } = null!;

        [Input("languageCode")]
        public Input<string>? LanguageCode { get; set; }

        /// <summary>
        /// FAQ name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// FAQ role ARN
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        /// <summary>
        /// FAQ S3 path
        /// </summary>
        [Input("s3Path", required: true)]
        public Input<Inputs.FaqS3PathArgs> S3Path { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.FaqTagArgs>? _tags;

        /// <summary>
        /// Tags for labeling the FAQ
        /// </summary>
        public InputList<Inputs.FaqTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.FaqTagArgs>());
            set => _tags = value;
        }

        public FaqArgs()
        {
        }
        public static new FaqArgs Empty => new FaqArgs();
    }
}
