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
    /// Resource Type definition for AWS::EC2::TransitGatewayRoute
    /// </summary>
    [Obsolete(@"TransitGatewayRoute is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:ec2:TransitGatewayRoute")]
    public partial class TransitGatewayRoute : global::Pulumi.CustomResource
    {
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        [Output("blackhole")]
        public Output<bool?> Blackhole { get; private set; } = null!;

        [Output("destinationCidrBlock")]
        public Output<string?> DestinationCidrBlock { get; private set; } = null!;

        [Output("transitGatewayAttachmentId")]
        public Output<string?> TransitGatewayAttachmentId { get; private set; } = null!;

        [Output("transitGatewayRouteTableId")]
        public Output<string> TransitGatewayRouteTableId { get; private set; } = null!;


        /// <summary>
        /// Create a TransitGatewayRoute resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TransitGatewayRoute(string name, TransitGatewayRouteArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ec2:TransitGatewayRoute", name, args ?? new TransitGatewayRouteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TransitGatewayRoute(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:TransitGatewayRoute", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "blackhole",
                    "destinationCidrBlock",
                    "transitGatewayAttachmentId",
                    "transitGatewayRouteTableId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TransitGatewayRoute resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TransitGatewayRoute Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new TransitGatewayRoute(name, id, options);
        }
    }

    public sealed class TransitGatewayRouteArgs : global::Pulumi.ResourceArgs
    {
        [Input("blackhole")]
        public Input<bool>? Blackhole { get; set; }

        [Input("destinationCidrBlock")]
        public Input<string>? DestinationCidrBlock { get; set; }

        [Input("transitGatewayAttachmentId")]
        public Input<string>? TransitGatewayAttachmentId { get; set; }

        [Input("transitGatewayRouteTableId", required: true)]
        public Input<string> TransitGatewayRouteTableId { get; set; } = null!;

        public TransitGatewayRouteArgs()
        {
        }
        public static new TransitGatewayRouteArgs Empty => new TransitGatewayRouteArgs();
    }
}
