// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * AWS::NetworkManager::TransitGatewayPeering Resoruce Type.
 */
export class TransitGatewayPeering extends pulumi.CustomResource {
    /**
     * Get an existing TransitGatewayPeering resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TransitGatewayPeering {
        return new TransitGatewayPeering(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:networkmanager:TransitGatewayPeering';

    /**
     * Returns true if the given object is an instance of TransitGatewayPeering.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TransitGatewayPeering {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TransitGatewayPeering.__pulumiType;
    }

    /**
     * The ARN (Amazon Resource Name) of the core network that you want to peer a transit gateway to.
     */
    public /*out*/ readonly coreNetworkArn!: pulumi.Output<string>;
    /**
     * The Id of the core network that you want to peer a transit gateway to.
     */
    public readonly coreNetworkId!: pulumi.Output<string>;
    /**
     * The creation time of the transit gateway peering
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The location of the transit gateway peering
     */
    public /*out*/ readonly edgeLocation!: pulumi.Output<string>;
    /**
     * Peering owner account Id
     */
    public /*out*/ readonly ownerAccountId!: pulumi.Output<string>;
    /**
     * The Id of the transit gateway peering
     */
    public /*out*/ readonly peeringId!: pulumi.Output<string>;
    /**
     * Peering type (TransitGatewayPeering)
     */
    public /*out*/ readonly peeringType!: pulumi.Output<string>;
    /**
     * The ARN (Amazon Resource Name) of the resource that you will peer to a core network
     */
    public /*out*/ readonly resourceArn!: pulumi.Output<string>;
    /**
     * The state of the transit gateway peering
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.networkmanager.TransitGatewayPeeringTag[] | undefined>;
    /**
     * The ARN (Amazon Resource Name) of the transit gateway that you will peer to a core network
     */
    public readonly transitGatewayArn!: pulumi.Output<string>;
    /**
     * The ID of the TransitGatewayPeeringAttachment
     */
    public /*out*/ readonly transitGatewayPeeringAttachmentId!: pulumi.Output<string>;

    /**
     * Create a TransitGatewayPeering resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TransitGatewayPeeringArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.coreNetworkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'coreNetworkId'");
            }
            if ((!args || args.transitGatewayArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'transitGatewayArn'");
            }
            resourceInputs["coreNetworkId"] = args ? args.coreNetworkId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["transitGatewayArn"] = args ? args.transitGatewayArn : undefined;
            resourceInputs["coreNetworkArn"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["edgeLocation"] = undefined /*out*/;
            resourceInputs["ownerAccountId"] = undefined /*out*/;
            resourceInputs["peeringId"] = undefined /*out*/;
            resourceInputs["peeringType"] = undefined /*out*/;
            resourceInputs["resourceArn"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["transitGatewayPeeringAttachmentId"] = undefined /*out*/;
        } else {
            resourceInputs["coreNetworkArn"] = undefined /*out*/;
            resourceInputs["coreNetworkId"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["edgeLocation"] = undefined /*out*/;
            resourceInputs["ownerAccountId"] = undefined /*out*/;
            resourceInputs["peeringId"] = undefined /*out*/;
            resourceInputs["peeringType"] = undefined /*out*/;
            resourceInputs["resourceArn"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["transitGatewayArn"] = undefined /*out*/;
            resourceInputs["transitGatewayPeeringAttachmentId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TransitGatewayPeering.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a TransitGatewayPeering resource.
 */
export interface TransitGatewayPeeringArgs {
    /**
     * The Id of the core network that you want to peer a transit gateway to.
     */
    coreNetworkId: pulumi.Input<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.networkmanager.TransitGatewayPeeringTagArgs>[]>;
    /**
     * The ARN (Amazon Resource Name) of the transit gateway that you will peer to a core network
     */
    transitGatewayArn: pulumi.Input<string>;
}
