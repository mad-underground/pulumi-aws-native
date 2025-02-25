// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisAnalytics
{
    /// <summary>
    /// Resource Type definition for AWS::KinesisAnalytics::Application
    /// </summary>
    [Obsolete(@"Application is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:kinesisanalytics:Application")]
    public partial class Application : global::Pulumi.CustomResource
    {
        [Output("applicationCode")]
        public Output<string?> ApplicationCode { get; private set; } = null!;

        [Output("applicationDescription")]
        public Output<string?> ApplicationDescription { get; private set; } = null!;

        [Output("applicationName")]
        public Output<string?> ApplicationName { get; private set; } = null!;

        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        [Output("inputs")]
        public Output<ImmutableArray<Outputs.ApplicationInput>> Inputs { get; private set; } = null!;


        /// <summary>
        /// Create a Application resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Application(string name, ApplicationArgs args, CustomResourceOptions? options = null)
            : base("aws-native:kinesisanalytics:Application", name, args ?? new ApplicationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Application(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:kinesisanalytics:Application", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "applicationName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Application resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Application Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Application(name, id, options);
        }
    }

    public sealed class ApplicationArgs : global::Pulumi.ResourceArgs
    {
        [Input("applicationCode")]
        public Input<string>? ApplicationCode { get; set; }

        [Input("applicationDescription")]
        public Input<string>? ApplicationDescription { get; set; }

        [Input("applicationName")]
        public Input<string>? ApplicationName { get; set; }

        [Input("inputs", required: true)]
        private InputList<Inputs.ApplicationInputArgs>? _inputs;
        public InputList<Inputs.ApplicationInputArgs> Inputs
        {
            get => _inputs ?? (_inputs = new InputList<Inputs.ApplicationInputArgs>());
            set => _inputs = value;
        }

        public ApplicationArgs()
        {
        }
        public static new ApplicationArgs Empty => new ApplicationArgs();
    }
}
