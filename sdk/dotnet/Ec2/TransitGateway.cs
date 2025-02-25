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
    /// Resource Type definition for AWS::EC2::TransitGateway
    /// </summary>
    [AwsNativeResourceType("aws-native:ec2:TransitGateway")]
    public partial class TransitGateway : global::Pulumi.CustomResource
    {
        [Output("amazonSideAsn")]
        public Output<int?> AmazonSideAsn { get; private set; } = null!;

        [Output("associationDefaultRouteTableId")]
        public Output<string?> AssociationDefaultRouteTableId { get; private set; } = null!;

        [Output("autoAcceptSharedAttachments")]
        public Output<string?> AutoAcceptSharedAttachments { get; private set; } = null!;

        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        [Output("defaultRouteTableAssociation")]
        public Output<string?> DefaultRouteTableAssociation { get; private set; } = null!;

        [Output("defaultRouteTablePropagation")]
        public Output<string?> DefaultRouteTablePropagation { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("dnsSupport")]
        public Output<string?> DnsSupport { get; private set; } = null!;

        [Output("multicastSupport")]
        public Output<string?> MulticastSupport { get; private set; } = null!;

        [Output("propagationDefaultRouteTableId")]
        public Output<string?> PropagationDefaultRouteTableId { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        [Output("transitGatewayArn")]
        public Output<string> TransitGatewayArn { get; private set; } = null!;

        [Output("transitGatewayCidrBlocks")]
        public Output<ImmutableArray<string>> TransitGatewayCidrBlocks { get; private set; } = null!;

        [Output("vpnEcmpSupport")]
        public Output<string?> VpnEcmpSupport { get; private set; } = null!;


        /// <summary>
        /// Create a TransitGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TransitGateway(string name, TransitGatewayArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:ec2:TransitGateway", name, args ?? new TransitGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TransitGateway(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:TransitGateway", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "amazonSideAsn",
                    "multicastSupport",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TransitGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TransitGateway Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new TransitGateway(name, id, options);
        }
    }

    public sealed class TransitGatewayArgs : global::Pulumi.ResourceArgs
    {
        [Input("amazonSideAsn")]
        public Input<int>? AmazonSideAsn { get; set; }

        [Input("associationDefaultRouteTableId")]
        public Input<string>? AssociationDefaultRouteTableId { get; set; }

        [Input("autoAcceptSharedAttachments")]
        public Input<string>? AutoAcceptSharedAttachments { get; set; }

        [Input("defaultRouteTableAssociation")]
        public Input<string>? DefaultRouteTableAssociation { get; set; }

        [Input("defaultRouteTablePropagation")]
        public Input<string>? DefaultRouteTablePropagation { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("dnsSupport")]
        public Input<string>? DnsSupport { get; set; }

        [Input("multicastSupport")]
        public Input<string>? MulticastSupport { get; set; }

        [Input("propagationDefaultRouteTableId")]
        public Input<string>? PropagationDefaultRouteTableId { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        [Input("transitGatewayCidrBlocks")]
        private InputList<string>? _transitGatewayCidrBlocks;
        public InputList<string> TransitGatewayCidrBlocks
        {
            get => _transitGatewayCidrBlocks ?? (_transitGatewayCidrBlocks = new InputList<string>());
            set => _transitGatewayCidrBlocks = value;
        }

        [Input("vpnEcmpSupport")]
        public Input<string>? VpnEcmpSupport { get; set; }

        public TransitGatewayArgs()
        {
        }
        public static new TransitGatewayArgs Empty => new TransitGatewayArgs();
    }
}
