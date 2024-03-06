// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3Outposts
{
    /// <summary>
    /// Resource Type Definition for AWS::S3Outposts::Endpoint
    /// </summary>
    [AwsNativeResourceType("aws-native:s3outposts:Endpoint")]
    public partial class Endpoint : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The type of access for the on-premise network connectivity for the Outpost endpoint. To access endpoint from an on-premises network, you must specify the access type and provide the customer owned Ipv4 pool.
        /// </summary>
        [Output("accessType")]
        public Output<Pulumi.AwsNative.S3Outposts.EndpointAccessType?> AccessType { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the endpoint.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The ID of the endpoint.
        /// </summary>
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        /// <summary>
        /// The VPC CIDR committed by this endpoint.
        /// </summary>
        [Output("cidrBlock")]
        public Output<string> CidrBlock { get; private set; } = null!;

        /// <summary>
        /// The time the endpoint was created.
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// The ID of the customer-owned IPv4 pool for the Endpoint. IP addresses will be allocated from this pool for the endpoint.
        /// </summary>
        [Output("customerOwnedIpv4Pool")]
        public Output<string?> CustomerOwnedIpv4Pool { get; private set; } = null!;

        /// <summary>
        /// The failure reason, if any, for a create or delete endpoint operation.
        /// </summary>
        [Output("failedReason")]
        public Output<Outputs.EndpointFailedReason?> FailedReason { get; private set; } = null!;

        /// <summary>
        /// The network interfaces of the endpoint.
        /// </summary>
        [Output("networkInterfaces")]
        public Output<ImmutableArray<Outputs.EndpointNetworkInterface>> NetworkInterfaces { get; private set; } = null!;

        /// <summary>
        /// The id of the customer outpost on which the bucket resides.
        /// </summary>
        [Output("outpostId")]
        public Output<string> OutpostId { get; private set; } = null!;

        /// <summary>
        /// The ID of the security group to use with the endpoint.
        /// </summary>
        [Output("securityGroupId")]
        public Output<string> SecurityGroupId { get; private set; } = null!;

        [Output("status")]
        public Output<Pulumi.AwsNative.S3Outposts.EndpointStatus> Status { get; private set; } = null!;

        /// <summary>
        /// The ID of the subnet in the selected VPC. The subnet must belong to the Outpost.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;


        /// <summary>
        /// Create a Endpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Endpoint(string name, EndpointArgs args, CustomResourceOptions? options = null)
            : base("aws-native:s3outposts:Endpoint", name, args ?? new EndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Endpoint(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:s3outposts:Endpoint", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "accessType",
                    "customerOwnedIpv4Pool",
                    "outpostId",
                    "securityGroupId",
                    "subnetId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Endpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Endpoint Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Endpoint(name, id, options);
        }
    }

    public sealed class EndpointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of access for the on-premise network connectivity for the Outpost endpoint. To access endpoint from an on-premises network, you must specify the access type and provide the customer owned Ipv4 pool.
        /// </summary>
        [Input("accessType")]
        public Input<Pulumi.AwsNative.S3Outposts.EndpointAccessType>? AccessType { get; set; }

        /// <summary>
        /// The ID of the customer-owned IPv4 pool for the Endpoint. IP addresses will be allocated from this pool for the endpoint.
        /// </summary>
        [Input("customerOwnedIpv4Pool")]
        public Input<string>? CustomerOwnedIpv4Pool { get; set; }

        /// <summary>
        /// The failure reason, if any, for a create or delete endpoint operation.
        /// </summary>
        [Input("failedReason")]
        public Input<Inputs.EndpointFailedReasonArgs>? FailedReason { get; set; }

        /// <summary>
        /// The id of the customer outpost on which the bucket resides.
        /// </summary>
        [Input("outpostId", required: true)]
        public Input<string> OutpostId { get; set; } = null!;

        /// <summary>
        /// The ID of the security group to use with the endpoint.
        /// </summary>
        [Input("securityGroupId", required: true)]
        public Input<string> SecurityGroupId { get; set; } = null!;

        /// <summary>
        /// The ID of the subnet in the selected VPC. The subnet must belong to the Outpost.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        public EndpointArgs()
        {
        }
        public static new EndpointArgs Empty => new EndpointArgs();
    }
}
