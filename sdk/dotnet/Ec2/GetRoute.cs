// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2
{
    public static class GetRoute
    {
        /// <summary>
        /// Resource Type definition for AWS::EC2::Route
        /// </summary>
        public static Task<GetRouteResult> InvokeAsync(GetRouteArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRouteResult>("aws-native:ec2:getRoute", args ?? new GetRouteArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::EC2::Route
        /// </summary>
        public static Output<GetRouteResult> Invoke(GetRouteInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRouteResult>("aws-native:ec2:getRoute", args ?? new GetRouteInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRouteArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The primary identifier of the resource generated by the service.
        /// </summary>
        [Input("cidrBlock", required: true)]
        public string CidrBlock { get; set; } = null!;

        /// <summary>
        /// The ID of the route table. The routing table must be associated with the same VPC that the virtual private gateway is attached to.
        /// </summary>
        [Input("routeTableId", required: true)]
        public string RouteTableId { get; set; } = null!;

        public GetRouteArgs()
        {
        }
        public static new GetRouteArgs Empty => new GetRouteArgs();
    }

    public sealed class GetRouteInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The primary identifier of the resource generated by the service.
        /// </summary>
        [Input("cidrBlock", required: true)]
        public Input<string> CidrBlock { get; set; } = null!;

        /// <summary>
        /// The ID of the route table. The routing table must be associated with the same VPC that the virtual private gateway is attached to.
        /// </summary>
        [Input("routeTableId", required: true)]
        public Input<string> RouteTableId { get; set; } = null!;

        public GetRouteInvokeArgs()
        {
        }
        public static new GetRouteInvokeArgs Empty => new GetRouteInvokeArgs();
    }


    [OutputType]
    public sealed class GetRouteResult
    {
        /// <summary>
        /// The ID of the carrier gateway.
        /// </summary>
        public readonly string? CarrierGatewayId;
        /// <summary>
        /// The primary identifier of the resource generated by the service.
        /// </summary>
        public readonly string? CidrBlock;
        /// <summary>
        /// The ID of the egress-only internet gateway.
        /// </summary>
        public readonly string? EgressOnlyInternetGatewayId;
        /// <summary>
        /// The ID of an internet gateway or virtual private gateway attached to your VPC.
        /// </summary>
        public readonly string? GatewayId;
        /// <summary>
        /// The ID of a NAT instance in your VPC.
        /// </summary>
        public readonly string? InstanceId;
        /// <summary>
        /// The ID of the local gateway.
        /// </summary>
        public readonly string? LocalGatewayId;
        /// <summary>
        /// The ID of a NAT gateway.
        /// </summary>
        public readonly string? NatGatewayId;
        /// <summary>
        /// The ID of the network interface.
        /// </summary>
        public readonly string? NetworkInterfaceId;
        /// <summary>
        /// The ID of a transit gateway.
        /// </summary>
        public readonly string? TransitGatewayId;
        /// <summary>
        /// The ID of a VPC endpoint. Supported for Gateway Load Balancer endpoints only.
        /// </summary>
        public readonly string? VpcEndpointId;
        /// <summary>
        /// The ID of a VPC peering connection.
        /// </summary>
        public readonly string? VpcPeeringConnectionId;

        [OutputConstructor]
        private GetRouteResult(
            string? carrierGatewayId,

            string? cidrBlock,

            string? egressOnlyInternetGatewayId,

            string? gatewayId,

            string? instanceId,

            string? localGatewayId,

            string? natGatewayId,

            string? networkInterfaceId,

            string? transitGatewayId,

            string? vpcEndpointId,

            string? vpcPeeringConnectionId)
        {
            CarrierGatewayId = carrierGatewayId;
            CidrBlock = cidrBlock;
            EgressOnlyInternetGatewayId = egressOnlyInternetGatewayId;
            GatewayId = gatewayId;
            InstanceId = instanceId;
            LocalGatewayId = localGatewayId;
            NatGatewayId = natGatewayId;
            NetworkInterfaceId = networkInterfaceId;
            TransitGatewayId = transitGatewayId;
            VpcEndpointId = vpcEndpointId;
            VpcPeeringConnectionId = vpcPeeringConnectionId;
        }
    }
}
