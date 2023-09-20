// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2
{
    public static class GetTransitGatewayMulticastDomainAssociation
    {
        /// <summary>
        /// The AWS::EC2::TransitGatewayMulticastDomainAssociation type
        /// </summary>
        public static Task<GetTransitGatewayMulticastDomainAssociationResult> InvokeAsync(GetTransitGatewayMulticastDomainAssociationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTransitGatewayMulticastDomainAssociationResult>("aws-native:ec2:getTransitGatewayMulticastDomainAssociation", args ?? new GetTransitGatewayMulticastDomainAssociationArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::EC2::TransitGatewayMulticastDomainAssociation type
        /// </summary>
        public static Output<GetTransitGatewayMulticastDomainAssociationResult> Invoke(GetTransitGatewayMulticastDomainAssociationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTransitGatewayMulticastDomainAssociationResult>("aws-native:ec2:getTransitGatewayMulticastDomainAssociation", args ?? new GetTransitGatewayMulticastDomainAssociationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTransitGatewayMulticastDomainAssociationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The IDs of the subnets to associate with the transit gateway multicast domain.
        /// </summary>
        [Input("subnetId", required: true)]
        public string SubnetId { get; set; } = null!;

        /// <summary>
        /// The ID of the transit gateway attachment.
        /// </summary>
        [Input("transitGatewayAttachmentId", required: true)]
        public string TransitGatewayAttachmentId { get; set; } = null!;

        /// <summary>
        /// The ID of the transit gateway multicast domain.
        /// </summary>
        [Input("transitGatewayMulticastDomainId", required: true)]
        public string TransitGatewayMulticastDomainId { get; set; } = null!;

        public GetTransitGatewayMulticastDomainAssociationArgs()
        {
        }
        public static new GetTransitGatewayMulticastDomainAssociationArgs Empty => new GetTransitGatewayMulticastDomainAssociationArgs();
    }

    public sealed class GetTransitGatewayMulticastDomainAssociationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The IDs of the subnets to associate with the transit gateway multicast domain.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        /// <summary>
        /// The ID of the transit gateway attachment.
        /// </summary>
        [Input("transitGatewayAttachmentId", required: true)]
        public Input<string> TransitGatewayAttachmentId { get; set; } = null!;

        /// <summary>
        /// The ID of the transit gateway multicast domain.
        /// </summary>
        [Input("transitGatewayMulticastDomainId", required: true)]
        public Input<string> TransitGatewayMulticastDomainId { get; set; } = null!;

        public GetTransitGatewayMulticastDomainAssociationInvokeArgs()
        {
        }
        public static new GetTransitGatewayMulticastDomainAssociationInvokeArgs Empty => new GetTransitGatewayMulticastDomainAssociationInvokeArgs();
    }


    [OutputType]
    public sealed class GetTransitGatewayMulticastDomainAssociationResult
    {
        /// <summary>
        /// The ID of the resource.
        /// </summary>
        public readonly string? ResourceId;
        /// <summary>
        /// The type of resource, for example a VPC attachment.
        /// </summary>
        public readonly string? ResourceType;
        /// <summary>
        /// The state of the subnet association.
        /// </summary>
        public readonly string? State;

        [OutputConstructor]
        private GetTransitGatewayMulticastDomainAssociationResult(
            string? resourceId,

            string? resourceType,

            string? state)
        {
            ResourceId = resourceId;
            ResourceType = resourceType;
            State = state;
        }
    }
}