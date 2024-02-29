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
        /// Specifies a route in a route table. For more information, see [Routes](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html#route-table-routes) in the *Amazon VPC User Guide*.
        ///  You must specify either a destination CIDR block or prefix list ID. You must also specify exactly one of the resources as the target.
        ///  If you create a route that references a transit gateway in the same template where you create the transit gateway, you must declare a dependency on the transit gateway attachment. The route table cannot use the transit gateway until it has successfully attached to the VPC. Add a [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) in the ``AWS::EC2::Route`` resource to explicitly declare a dependency on the ``AWS::EC2::TransitGatewayAttachment`` resource.
        /// </summary>
        public static Task<GetRouteResult> InvokeAsync(GetRouteArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRouteResult>("aws-native:ec2:getRoute", args ?? new GetRouteArgs(), options.WithDefaults());

        /// <summary>
        /// Specifies a route in a route table. For more information, see [Routes](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html#route-table-routes) in the *Amazon VPC User Guide*.
        ///  You must specify either a destination CIDR block or prefix list ID. You must also specify exactly one of the resources as the target.
        ///  If you create a route that references a transit gateway in the same template where you create the transit gateway, you must declare a dependency on the transit gateway attachment. The route table cannot use the transit gateway until it has successfully attached to the VPC. Add a [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) in the ``AWS::EC2::Route`` resource to explicitly declare a dependency on the ``AWS::EC2::TransitGatewayAttachment`` resource.
        /// </summary>
        public static Output<GetRouteResult> Invoke(GetRouteInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRouteResult>("aws-native:ec2:getRoute", args ?? new GetRouteInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRouteArgs : global::Pulumi.InvokeArgs
    {
        [Input("cidrBlock", required: true)]
        public string CidrBlock { get; set; } = null!;

        /// <summary>
        /// The ID of the route table for the route.
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
        [Input("cidrBlock", required: true)]
        public Input<string> CidrBlock { get; set; } = null!;

        /// <summary>
        /// The ID of the route table for the route.
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
        ///  You can only use this option when the VPC contains a subnet which is associated with a Wavelength Zone.
        /// </summary>
        public readonly string? CarrierGatewayId;
        public readonly string? CidrBlock;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the core network.
        /// </summary>
        public readonly string? CoreNetworkArn;
        /// <summary>
        /// [IPv6 traffic only] The ID of an egress-only internet gateway.
        /// </summary>
        public readonly string? EgressOnlyInternetGatewayId;
        /// <summary>
        /// The ID of an internet gateway or virtual private gateway attached to your VPC.
        /// </summary>
        public readonly string? GatewayId;
        /// <summary>
        /// The ID of a NAT instance in your VPC. The operation fails if you specify an instance ID unless exactly one network interface is attached.
        /// </summary>
        public readonly string? InstanceId;
        /// <summary>
        /// The ID of the local gateway.
        /// </summary>
        public readonly string? LocalGatewayId;
        /// <summary>
        /// [IPv4 traffic only] The ID of a NAT gateway.
        /// </summary>
        public readonly string? NatGatewayId;
        /// <summary>
        /// The ID of a network interface.
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

            string? coreNetworkArn,

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
            CoreNetworkArn = coreNetworkArn;
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
