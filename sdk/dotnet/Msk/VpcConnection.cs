// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Msk
{
    /// <summary>
    /// Resource Type definition for AWS::MSK::VpcConnection
    /// </summary>
    [AwsNativeResourceType("aws-native:msk:VpcConnection")]
    public partial class VpcConnection : global::Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("authentication")]
        public Output<Pulumi.AwsNative.Msk.VpcConnectionAuthentication> Authentication { get; private set; } = null!;

        [Output("clientSubnets")]
        public Output<ImmutableArray<string>> ClientSubnets { get; private set; } = null!;

        [Output("securityGroups")]
        public Output<ImmutableArray<string>> SecurityGroups { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the target cluster
        /// </summary>
        [Output("targetClusterArn")]
        public Output<string> TargetClusterArn { get; private set; } = null!;

        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a VpcConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcConnection(string name, VpcConnectionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:msk:VpcConnection", name, args ?? new VpcConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:msk:VpcConnection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "authentication",
                    "clientSubnets[*]",
                    "securityGroups[*]",
                    "targetClusterArn",
                    "vpcId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VpcConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcConnection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VpcConnection(name, id, options);
        }
    }

    public sealed class VpcConnectionArgs : global::Pulumi.ResourceArgs
    {
        [Input("authentication", required: true)]
        public Input<Pulumi.AwsNative.Msk.VpcConnectionAuthentication> Authentication { get; set; } = null!;

        [Input("clientSubnets", required: true)]
        private InputList<string>? _clientSubnets;
        public InputList<string> ClientSubnets
        {
            get => _clientSubnets ?? (_clientSubnets = new InputList<string>());
            set => _clientSubnets = value;
        }

        [Input("securityGroups", required: true)]
        private InputList<string>? _securityGroups;
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the target cluster
        /// </summary>
        [Input("targetClusterArn", required: true)]
        public Input<string> TargetClusterArn { get; set; } = null!;

        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public VpcConnectionArgs()
        {
        }
        public static new VpcConnectionArgs Empty => new VpcConnectionArgs();
    }
}
