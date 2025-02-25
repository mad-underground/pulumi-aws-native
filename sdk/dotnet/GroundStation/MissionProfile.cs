// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GroundStation
{
    /// <summary>
    /// AWS Ground Station Mission Profile resource type for CloudFormation.
    /// </summary>
    [AwsNativeResourceType("aws-native:groundstation:MissionProfile")]
    public partial class MissionProfile : global::Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        /// <summary>
        /// Post-pass time needed after the contact.
        /// </summary>
        [Output("contactPostPassDurationSeconds")]
        public Output<int?> ContactPostPassDurationSeconds { get; private set; } = null!;

        /// <summary>
        /// Pre-pass time needed before the contact.
        /// </summary>
        [Output("contactPrePassDurationSeconds")]
        public Output<int?> ContactPrePassDurationSeconds { get; private set; } = null!;

        [Output("dataflowEdges")]
        public Output<ImmutableArray<Outputs.MissionProfileDataflowEdge>> DataflowEdges { get; private set; } = null!;

        /// <summary>
        /// Visibilities with shorter duration than the specified minimum viable contact duration will be ignored when searching for available contacts.
        /// </summary>
        [Output("minimumViableContactDurationSeconds")]
        public Output<int> MinimumViableContactDurationSeconds { get; private set; } = null!;

        /// <summary>
        /// A name used to identify a mission profile.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The ARN of a KMS Key used for encrypting data during transmission from the source to destination locations.
        /// </summary>
        [Output("streamsKmsKey")]
        public Output<Outputs.MissionProfileStreamsKmsKey?> StreamsKmsKey { get; private set; } = null!;

        /// <summary>
        /// The ARN of the KMS Key or Alias Key role used to define permissions on KMS Key usage.
        /// </summary>
        [Output("streamsKmsRole")]
        public Output<string?> StreamsKmsRole { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        [Output("trackingConfigArn")]
        public Output<string> TrackingConfigArn { get; private set; } = null!;


        /// <summary>
        /// Create a MissionProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MissionProfile(string name, MissionProfileArgs args, CustomResourceOptions? options = null)
            : base("aws-native:groundstation:MissionProfile", name, args ?? new MissionProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MissionProfile(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:groundstation:MissionProfile", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing MissionProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MissionProfile Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new MissionProfile(name, id, options);
        }
    }

    public sealed class MissionProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Post-pass time needed after the contact.
        /// </summary>
        [Input("contactPostPassDurationSeconds")]
        public Input<int>? ContactPostPassDurationSeconds { get; set; }

        /// <summary>
        /// Pre-pass time needed before the contact.
        /// </summary>
        [Input("contactPrePassDurationSeconds")]
        public Input<int>? ContactPrePassDurationSeconds { get; set; }

        [Input("dataflowEdges", required: true)]
        private InputList<Inputs.MissionProfileDataflowEdgeArgs>? _dataflowEdges;
        public InputList<Inputs.MissionProfileDataflowEdgeArgs> DataflowEdges
        {
            get => _dataflowEdges ?? (_dataflowEdges = new InputList<Inputs.MissionProfileDataflowEdgeArgs>());
            set => _dataflowEdges = value;
        }

        /// <summary>
        /// Visibilities with shorter duration than the specified minimum viable contact duration will be ignored when searching for available contacts.
        /// </summary>
        [Input("minimumViableContactDurationSeconds", required: true)]
        public Input<int> MinimumViableContactDurationSeconds { get; set; } = null!;

        /// <summary>
        /// A name used to identify a mission profile.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ARN of a KMS Key used for encrypting data during transmission from the source to destination locations.
        /// </summary>
        [Input("streamsKmsKey")]
        public Input<Inputs.MissionProfileStreamsKmsKeyArgs>? StreamsKmsKey { get; set; }

        /// <summary>
        /// The ARN of the KMS Key or Alias Key role used to define permissions on KMS Key usage.
        /// </summary>
        [Input("streamsKmsRole")]
        public Input<string>? StreamsKmsRole { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        [Input("trackingConfigArn", required: true)]
        public Input<string> TrackingConfigArn { get; set; } = null!;

        public MissionProfileArgs()
        {
        }
        public static new MissionProfileArgs Empty => new MissionProfileArgs();
    }
}
