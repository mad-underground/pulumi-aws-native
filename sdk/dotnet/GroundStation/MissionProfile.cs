// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GroundStation
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-missionprofile.html
    /// </summary>
    [AwsNativeResourceType("aws-native:GroundStation:MissionProfile")]
    public partial class MissionProfile : Pulumi.CustomResource
    {
        [Output("Arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-missionprofile.html#cfn-groundstation-missionprofile-contactpostpassdurationseconds
        /// </summary>
        [Output("ContactPostPassDurationSeconds")]
        public Output<int?> ContactPostPassDurationSeconds { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-missionprofile.html#cfn-groundstation-missionprofile-contactprepassdurationseconds
        /// </summary>
        [Output("ContactPrePassDurationSeconds")]
        public Output<int?> ContactPrePassDurationSeconds { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-missionprofile.html#cfn-groundstation-missionprofile-dataflowedges
        /// </summary>
        [Output("DataflowEdges")]
        public Output<ImmutableArray<Outputs.MissionProfileDataflowEdge>> DataflowEdges { get; private set; } = null!;

        [Output("Id")]
        public Output<string> Id { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-missionprofile.html#cfn-groundstation-missionprofile-minimumviablecontactdurationseconds
        /// </summary>
        [Output("MinimumViableContactDurationSeconds")]
        public Output<int> MinimumViableContactDurationSeconds { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-missionprofile.html#cfn-groundstation-missionprofile-name
        /// </summary>
        [Output("Name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("Region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-missionprofile.html#cfn-groundstation-missionprofile-tags
        /// </summary>
        [Output("Tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-missionprofile.html#cfn-groundstation-missionprofile-trackingconfigarn
        /// </summary>
        [Output("TrackingConfigArn")]
        public Output<string> TrackingConfigArn { get; private set; } = null!;


        /// <summary>
        /// Create a MissionProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MissionProfile(string name, MissionProfileArgs args, CustomResourceOptions? options = null)
            : base("aws-native:GroundStation:MissionProfile", name, args ?? new MissionProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MissionProfile(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:GroundStation:MissionProfile", name, null, MakeResourceOptions(options, id))
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

    public sealed class MissionProfileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-missionprofile.html#cfn-groundstation-missionprofile-contactpostpassdurationseconds
        /// </summary>
        [Input("ContactPostPassDurationSeconds")]
        public Input<int>? ContactPostPassDurationSeconds { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-missionprofile.html#cfn-groundstation-missionprofile-contactprepassdurationseconds
        /// </summary>
        [Input("ContactPrePassDurationSeconds")]
        public Input<int>? ContactPrePassDurationSeconds { get; set; }

        [Input("DataflowEdges", required: true)]
        private InputList<Inputs.MissionProfileDataflowEdgeArgs>? _DataflowEdges;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-missionprofile.html#cfn-groundstation-missionprofile-dataflowedges
        /// </summary>
        public InputList<Inputs.MissionProfileDataflowEdgeArgs> DataflowEdges
        {
            get => _DataflowEdges ?? (_DataflowEdges = new InputList<Inputs.MissionProfileDataflowEdgeArgs>());
            set => _DataflowEdges = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-missionprofile.html#cfn-groundstation-missionprofile-minimumviablecontactdurationseconds
        /// </summary>
        [Input("MinimumViableContactDurationSeconds", required: true)]
        public Input<int> MinimumViableContactDurationSeconds { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-missionprofile.html#cfn-groundstation-missionprofile-name
        /// </summary>
        [Input("Name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("Tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _Tags;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-missionprofile.html#cfn-groundstation-missionprofile-tags
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _Tags ?? (_Tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _Tags = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-missionprofile.html#cfn-groundstation-missionprofile-trackingconfigarn
        /// </summary>
        [Input("TrackingConfigArn", required: true)]
        public Input<string> TrackingConfigArn { get; set; } = null!;

        public MissionProfileArgs()
        {
        }
    }
}
