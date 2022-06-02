// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2
{
    /// <summary>
    /// Resource Type definition for AWS::EC2::NatGateway
    /// </summary>
    [AwsNativeResourceType("aws-native:ec2:NatGateway")]
    public partial class NatGateway : Pulumi.CustomResource
    {
        [Output("allocationId")]
        public Output<string?> AllocationId { get; private set; } = null!;

        [Output("connectivityType")]
        public Output<string?> ConnectivityType { get; private set; } = null!;

        [Output("natGatewayId")]
        public Output<string> NatGatewayId { get; private set; } = null!;

        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.NatGatewayTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a NatGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NatGateway(string name, NatGatewayArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ec2:NatGateway", name, args ?? new NatGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NatGateway(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:NatGateway", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing NatGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NatGateway Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NatGateway(name, id, options);
        }
    }

    public sealed class NatGatewayArgs : Pulumi.ResourceArgs
    {
        [Input("allocationId")]
        public Input<string>? AllocationId { get; set; }

        [Input("connectivityType")]
        public Input<string>? ConnectivityType { get; set; }

        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.NatGatewayTagArgs>? _tags;
        public InputList<Inputs.NatGatewayTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.NatGatewayTagArgs>());
            set => _tags = value;
        }

        public NatGatewayArgs()
        {
        }
    }
}
