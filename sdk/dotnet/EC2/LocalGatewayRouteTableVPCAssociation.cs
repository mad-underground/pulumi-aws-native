// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayroutetablevpcassociation.html
    /// </summary>
    [AwsNativeResourceType("aws-native:EC2:LocalGatewayRouteTableVPCAssociation")]
    public partial class LocalGatewayRouteTableVPCAssociation : Pulumi.CustomResource
    {
        [Output("LocalGatewayId")]
        public Output<string> LocalGatewayId { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayroutetablevpcassociation.html#cfn-ec2-localgatewayroutetablevpcassociation-localgatewayroutetableid
        /// </summary>
        [Output("LocalGatewayRouteTableId")]
        public Output<string> LocalGatewayRouteTableId { get; private set; } = null!;

        [Output("LocalGatewayRouteTableVpcAssociationId")]
        public Output<string> LocalGatewayRouteTableVpcAssociationId { get; private set; } = null!;

        [Output("State")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayroutetablevpcassociation.html#cfn-ec2-localgatewayroutetablevpcassociation-tags
        /// </summary>
        [Output("Tags")]
        public Output<Outputs.LocalGatewayRouteTableVPCAssociationTags?> Tags { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayroutetablevpcassociation.html#cfn-ec2-localgatewayroutetablevpcassociation-vpcid
        /// </summary>
        [Output("VpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a LocalGatewayRouteTableVPCAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LocalGatewayRouteTableVPCAssociation(string name, LocalGatewayRouteTableVPCAssociationArgs args, CustomResourceOptions? options = null)
            : base("aws-native:EC2:LocalGatewayRouteTableVPCAssociation", name, args ?? new LocalGatewayRouteTableVPCAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LocalGatewayRouteTableVPCAssociation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:EC2:LocalGatewayRouteTableVPCAssociation", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing LocalGatewayRouteTableVPCAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LocalGatewayRouteTableVPCAssociation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LocalGatewayRouteTableVPCAssociation(name, id, options);
        }
    }

    public sealed class LocalGatewayRouteTableVPCAssociationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayroutetablevpcassociation.html#cfn-ec2-localgatewayroutetablevpcassociation-localgatewayroutetableid
        /// </summary>
        [Input("LocalGatewayRouteTableId", required: true)]
        public Input<string> LocalGatewayRouteTableId { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayroutetablevpcassociation.html#cfn-ec2-localgatewayroutetablevpcassociation-tags
        /// </summary>
        [Input("Tags")]
        public Input<Inputs.LocalGatewayRouteTableVPCAssociationTagsArgs>? Tags { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-localgatewayroutetablevpcassociation.html#cfn-ec2-localgatewayroutetablevpcassociation-vpcid
        /// </summary>
        [Input("VpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public LocalGatewayRouteTableVPCAssociationArgs()
        {
        }
    }
}
