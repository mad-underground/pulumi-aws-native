// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2
{
    /// <summary>
    /// Resource Type definition for AWS::EC2::TrafficMirrorSession
    /// </summary>
    [Obsolete(@"TrafficMirrorSession is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:ec2:TrafficMirrorSession")]
    public partial class TrafficMirrorSession : global::Pulumi.CustomResource
    {
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("networkInterfaceId")]
        public Output<string> NetworkInterfaceId { get; private set; } = null!;

        [Output("packetLength")]
        public Output<int?> PacketLength { get; private set; } = null!;

        [Output("sessionNumber")]
        public Output<int> SessionNumber { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        [Output("trafficMirrorFilterId")]
        public Output<string> TrafficMirrorFilterId { get; private set; } = null!;

        [Output("trafficMirrorTargetId")]
        public Output<string> TrafficMirrorTargetId { get; private set; } = null!;

        [Output("virtualNetworkId")]
        public Output<int?> VirtualNetworkId { get; private set; } = null!;


        /// <summary>
        /// Create a TrafficMirrorSession resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TrafficMirrorSession(string name, TrafficMirrorSessionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ec2:TrafficMirrorSession", name, args ?? new TrafficMirrorSessionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TrafficMirrorSession(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:TrafficMirrorSession", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "networkInterfaceId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TrafficMirrorSession resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TrafficMirrorSession Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new TrafficMirrorSession(name, id, options);
        }
    }

    public sealed class TrafficMirrorSessionArgs : global::Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("networkInterfaceId", required: true)]
        public Input<string> NetworkInterfaceId { get; set; } = null!;

        [Input("packetLength")]
        public Input<int>? PacketLength { get; set; }

        [Input("sessionNumber", required: true)]
        public Input<int> SessionNumber { get; set; } = null!;

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        [Input("trafficMirrorFilterId", required: true)]
        public Input<string> TrafficMirrorFilterId { get; set; } = null!;

        [Input("trafficMirrorTargetId", required: true)]
        public Input<string> TrafficMirrorTargetId { get; set; } = null!;

        [Input("virtualNetworkId")]
        public Input<int>? VirtualNetworkId { get; set; }

        public TrafficMirrorSessionArgs()
        {
        }
        public static new TrafficMirrorSessionArgs Empty => new TrafficMirrorSessionArgs();
    }
}
