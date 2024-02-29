// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Specifies a route in a route table. For more information, see [Routes](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html#route-table-routes) in the *Amazon VPC User Guide*.
 *  You must specify either a destination CIDR block or prefix list ID. You must also specify exactly one of the resources as the target.
 *  If you create a route that references a transit gateway in the same template where you create the transit gateway, you must declare a dependency on the transit gateway attachment. The route table cannot use the transit gateway until it has successfully attached to the VPC. Add a [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) in the ``AWS::EC2::Route`` resource to explicitly declare a dependency on the ``AWS::EC2::TransitGatewayAttachment`` resource.
 */
export class Route extends pulumi.CustomResource {
    /**
     * Get an existing Route resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Route {
        return new Route(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ec2:Route';

    /**
     * Returns true if the given object is an instance of Route.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Route {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Route.__pulumiType;
    }

    /**
     * The ID of the carrier gateway.
     *  You can only use this option when the VPC contains a subnet which is associated with a Wavelength Zone.
     */
    public readonly carrierGatewayId!: pulumi.Output<string | undefined>;
    public /*out*/ readonly cidrBlock!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the core network.
     */
    public readonly coreNetworkArn!: pulumi.Output<string | undefined>;
    /**
     * The IPv4 CIDR address block used for the destination match. Routing decisions are based on the most specific match. We modify the specified CIDR block to its canonical form; for example, if you specify ``100.68.0.18/18``, we modify it to ``100.68.0.0/18``.
     */
    public readonly destinationCidrBlock!: pulumi.Output<string | undefined>;
    /**
     * The IPv6 CIDR block used for the destination match. Routing decisions are based on the most specific match.
     */
    public readonly destinationIpv6CidrBlock!: pulumi.Output<string | undefined>;
    /**
     * The ID of a prefix list used for the destination match.
     */
    public readonly destinationPrefixListId!: pulumi.Output<string | undefined>;
    /**
     * [IPv6 traffic only] The ID of an egress-only internet gateway.
     */
    public readonly egressOnlyInternetGatewayId!: pulumi.Output<string | undefined>;
    /**
     * The ID of an internet gateway or virtual private gateway attached to your VPC.
     */
    public readonly gatewayId!: pulumi.Output<string | undefined>;
    /**
     * The ID of a NAT instance in your VPC. The operation fails if you specify an instance ID unless exactly one network interface is attached.
     */
    public readonly instanceId!: pulumi.Output<string | undefined>;
    /**
     * The ID of the local gateway.
     */
    public readonly localGatewayId!: pulumi.Output<string | undefined>;
    /**
     * [IPv4 traffic only] The ID of a NAT gateway.
     */
    public readonly natGatewayId!: pulumi.Output<string | undefined>;
    /**
     * The ID of a network interface.
     */
    public readonly networkInterfaceId!: pulumi.Output<string | undefined>;
    /**
     * The ID of the route table for the route.
     */
    public readonly routeTableId!: pulumi.Output<string>;
    /**
     * The ID of a transit gateway.
     */
    public readonly transitGatewayId!: pulumi.Output<string | undefined>;
    /**
     * The ID of a VPC endpoint. Supported for Gateway Load Balancer endpoints only.
     */
    public readonly vpcEndpointId!: pulumi.Output<string | undefined>;
    /**
     * The ID of a VPC peering connection.
     */
    public readonly vpcPeeringConnectionId!: pulumi.Output<string | undefined>;

    /**
     * Create a Route resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouteArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.routeTableId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routeTableId'");
            }
            resourceInputs["carrierGatewayId"] = args ? args.carrierGatewayId : undefined;
            resourceInputs["coreNetworkArn"] = args ? args.coreNetworkArn : undefined;
            resourceInputs["destinationCidrBlock"] = args ? args.destinationCidrBlock : undefined;
            resourceInputs["destinationIpv6CidrBlock"] = args ? args.destinationIpv6CidrBlock : undefined;
            resourceInputs["destinationPrefixListId"] = args ? args.destinationPrefixListId : undefined;
            resourceInputs["egressOnlyInternetGatewayId"] = args ? args.egressOnlyInternetGatewayId : undefined;
            resourceInputs["gatewayId"] = args ? args.gatewayId : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["localGatewayId"] = args ? args.localGatewayId : undefined;
            resourceInputs["natGatewayId"] = args ? args.natGatewayId : undefined;
            resourceInputs["networkInterfaceId"] = args ? args.networkInterfaceId : undefined;
            resourceInputs["routeTableId"] = args ? args.routeTableId : undefined;
            resourceInputs["transitGatewayId"] = args ? args.transitGatewayId : undefined;
            resourceInputs["vpcEndpointId"] = args ? args.vpcEndpointId : undefined;
            resourceInputs["vpcPeeringConnectionId"] = args ? args.vpcPeeringConnectionId : undefined;
            resourceInputs["cidrBlock"] = undefined /*out*/;
        } else {
            resourceInputs["carrierGatewayId"] = undefined /*out*/;
            resourceInputs["cidrBlock"] = undefined /*out*/;
            resourceInputs["coreNetworkArn"] = undefined /*out*/;
            resourceInputs["destinationCidrBlock"] = undefined /*out*/;
            resourceInputs["destinationIpv6CidrBlock"] = undefined /*out*/;
            resourceInputs["destinationPrefixListId"] = undefined /*out*/;
            resourceInputs["egressOnlyInternetGatewayId"] = undefined /*out*/;
            resourceInputs["gatewayId"] = undefined /*out*/;
            resourceInputs["instanceId"] = undefined /*out*/;
            resourceInputs["localGatewayId"] = undefined /*out*/;
            resourceInputs["natGatewayId"] = undefined /*out*/;
            resourceInputs["networkInterfaceId"] = undefined /*out*/;
            resourceInputs["routeTableId"] = undefined /*out*/;
            resourceInputs["transitGatewayId"] = undefined /*out*/;
            resourceInputs["vpcEndpointId"] = undefined /*out*/;
            resourceInputs["vpcPeeringConnectionId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["destinationCidrBlock", "destinationIpv6CidrBlock", "destinationPrefixListId", "routeTableId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Route.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Route resource.
 */
export interface RouteArgs {
    /**
     * The ID of the carrier gateway.
     *  You can only use this option when the VPC contains a subnet which is associated with a Wavelength Zone.
     */
    carrierGatewayId?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the core network.
     */
    coreNetworkArn?: pulumi.Input<string>;
    /**
     * The IPv4 CIDR address block used for the destination match. Routing decisions are based on the most specific match. We modify the specified CIDR block to its canonical form; for example, if you specify ``100.68.0.18/18``, we modify it to ``100.68.0.0/18``.
     */
    destinationCidrBlock?: pulumi.Input<string>;
    /**
     * The IPv6 CIDR block used for the destination match. Routing decisions are based on the most specific match.
     */
    destinationIpv6CidrBlock?: pulumi.Input<string>;
    /**
     * The ID of a prefix list used for the destination match.
     */
    destinationPrefixListId?: pulumi.Input<string>;
    /**
     * [IPv6 traffic only] The ID of an egress-only internet gateway.
     */
    egressOnlyInternetGatewayId?: pulumi.Input<string>;
    /**
     * The ID of an internet gateway or virtual private gateway attached to your VPC.
     */
    gatewayId?: pulumi.Input<string>;
    /**
     * The ID of a NAT instance in your VPC. The operation fails if you specify an instance ID unless exactly one network interface is attached.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The ID of the local gateway.
     */
    localGatewayId?: pulumi.Input<string>;
    /**
     * [IPv4 traffic only] The ID of a NAT gateway.
     */
    natGatewayId?: pulumi.Input<string>;
    /**
     * The ID of a network interface.
     */
    networkInterfaceId?: pulumi.Input<string>;
    /**
     * The ID of the route table for the route.
     */
    routeTableId: pulumi.Input<string>;
    /**
     * The ID of a transit gateway.
     */
    transitGatewayId?: pulumi.Input<string>;
    /**
     * The ID of a VPC endpoint. Supported for Gateway Load Balancer endpoints only.
     */
    vpcEndpointId?: pulumi.Input<string>;
    /**
     * The ID of a VPC peering connection.
     */
    vpcPeeringConnectionId?: pulumi.Input<string>;
}
