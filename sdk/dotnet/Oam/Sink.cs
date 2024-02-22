// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Oam
{
    /// <summary>
    /// Resource Type definition for AWS::Oam::Sink
    /// </summary>
    [AwsNativeResourceType("aws-native:oam:Sink")]
    public partial class Sink : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon resource name (ARN) of the ObservabilityAccessManager Sink
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The name of the ObservabilityAccessManager Sink.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The policy of this ObservabilityAccessManager Sink.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Oam::Sink` for more information about the expected schema for this property.
        /// </summary>
        [Output("policy")]
        public Output<object?> Policy { get; private set; } = null!;

        /// <summary>
        /// Tags to apply to the sink
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Sink resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Sink(string name, SinkArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:oam:Sink", name, args ?? new SinkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Sink(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:oam:Sink", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "name",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Sink resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Sink Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Sink(name, id, options);
        }
    }

    public sealed class SinkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the ObservabilityAccessManager Sink.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The policy of this ObservabilityAccessManager Sink.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Oam::Sink` for more information about the expected schema for this property.
        /// </summary>
        [Input("policy")]
        public Input<object>? Policy { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Tags to apply to the sink
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public SinkArgs()
        {
        }
        public static new SinkArgs Empty => new SinkArgs();
    }
}
