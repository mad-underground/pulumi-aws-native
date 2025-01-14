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
    /// Resource Type definition for AWS::EC2::VPNGatewayRoutePropagation
    /// </summary>
    [Obsolete(@"VpnGatewayRoutePropagation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:ec2:VpnGatewayRoutePropagation")]
    public partial class VpnGatewayRoutePropagation : global::Pulumi.CustomResource
    {
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        [Output("routeTableIds")]
        public Output<ImmutableArray<string>> RouteTableIds { get; private set; } = null!;

        [Output("vpnGatewayId")]
        public Output<string> VpnGatewayId { get; private set; } = null!;


        /// <summary>
        /// Create a VpnGatewayRoutePropagation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpnGatewayRoutePropagation(string name, VpnGatewayRoutePropagationArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ec2:VpnGatewayRoutePropagation", name, args ?? new VpnGatewayRoutePropagationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpnGatewayRoutePropagation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:VpnGatewayRoutePropagation", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing VpnGatewayRoutePropagation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpnGatewayRoutePropagation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VpnGatewayRoutePropagation(name, id, options);
        }
    }

    public sealed class VpnGatewayRoutePropagationArgs : global::Pulumi.ResourceArgs
    {
        [Input("routeTableIds", required: true)]
        private InputList<string>? _routeTableIds;
        public InputList<string> RouteTableIds
        {
            get => _routeTableIds ?? (_routeTableIds = new InputList<string>());
            set => _routeTableIds = value;
        }

        [Input("vpnGatewayId", required: true)]
        public Input<string> VpnGatewayId { get; set; } = null!;

        public VpnGatewayRoutePropagationArgs()
        {
        }
        public static new VpnGatewayRoutePropagationArgs Empty => new VpnGatewayRoutePropagationArgs();
    }
}
