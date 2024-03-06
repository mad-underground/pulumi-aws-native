// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.VpcLattice
{
    /// <summary>
    /// Associates a VPC with a service network.
    /// </summary>
    [AwsNativeResourceType("aws-native:vpclattice:ServiceNetworkVpcAssociation")]
    public partial class ServiceNetworkVpcAssociation : global::Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        [Output("serviceNetworkArn")]
        public Output<string> ServiceNetworkArn { get; private set; } = null!;

        [Output("serviceNetworkId")]
        public Output<string> ServiceNetworkId { get; private set; } = null!;

        [Output("serviceNetworkIdentifier")]
        public Output<string?> ServiceNetworkIdentifier { get; private set; } = null!;

        [Output("serviceNetworkName")]
        public Output<string> ServiceNetworkName { get; private set; } = null!;

        [Output("status")]
        public Output<Pulumi.AwsNative.VpcLattice.ServiceNetworkVpcAssociationStatus> Status { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        [Output("vpcIdentifier")]
        public Output<string?> VpcIdentifier { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceNetworkVpcAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceNetworkVpcAssociation(string name, ServiceNetworkVpcAssociationArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:vpclattice:ServiceNetworkVpcAssociation", name, args ?? new ServiceNetworkVpcAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceNetworkVpcAssociation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:vpclattice:ServiceNetworkVpcAssociation", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "serviceNetworkIdentifier",
                    "vpcIdentifier",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServiceNetworkVpcAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceNetworkVpcAssociation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServiceNetworkVpcAssociation(name, id, options);
        }
    }

    public sealed class ServiceNetworkVpcAssociationArgs : global::Pulumi.ResourceArgs
    {
        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("serviceNetworkIdentifier")]
        public Input<string>? ServiceNetworkIdentifier { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        [Input("vpcIdentifier")]
        public Input<string>? VpcIdentifier { get; set; }

        public ServiceNetworkVpcAssociationArgs()
        {
        }
        public static new ServiceNetworkVpcAssociationArgs Empty => new ServiceNetworkVpcAssociationArgs();
    }
}
