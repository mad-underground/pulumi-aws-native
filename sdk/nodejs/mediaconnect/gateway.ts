// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::MediaConnect::Gateway
 */
export class Gateway extends pulumi.CustomResource {
    /**
     * Get an existing Gateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Gateway {
        return new Gateway(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:mediaconnect:Gateway';

    /**
     * Returns true if the given object is an instance of Gateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Gateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Gateway.__pulumiType;
    }

    /**
     * The range of IP addresses that contribute content or initiate output requests for flows communicating with this gateway. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
     */
    public readonly egressCidrBlocks!: pulumi.Output<string[]>;
    /**
     * The Amazon Resource Name (ARN) of the gateway.
     */
    public /*out*/ readonly gatewayArn!: pulumi.Output<string>;
    /**
     * The current status of the gateway.
     */
    public /*out*/ readonly gatewayState!: pulumi.Output<enums.mediaconnect.GatewayState>;
    /**
     * The name of the gateway. This name can not be modified after the gateway is created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The list of networks in the gateway.
     */
    public readonly networks!: pulumi.Output<outputs.mediaconnect.GatewayNetwork[]>;

    /**
     * Create a Gateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GatewayArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.egressCidrBlocks === undefined) && !opts.urn) {
                throw new Error("Missing required property 'egressCidrBlocks'");
            }
            if ((!args || args.networks === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networks'");
            }
            resourceInputs["egressCidrBlocks"] = args ? args.egressCidrBlocks : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networks"] = args ? args.networks : undefined;
            resourceInputs["gatewayArn"] = undefined /*out*/;
            resourceInputs["gatewayState"] = undefined /*out*/;
        } else {
            resourceInputs["egressCidrBlocks"] = undefined /*out*/;
            resourceInputs["gatewayArn"] = undefined /*out*/;
            resourceInputs["gatewayState"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["networks"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Gateway.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Gateway resource.
 */
export interface GatewayArgs {
    /**
     * The range of IP addresses that contribute content or initiate output requests for flows communicating with this gateway. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
     */
    egressCidrBlocks: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the gateway. This name can not be modified after the gateway is created.
     */
    name?: pulumi.Input<string>;
    /**
     * The list of networks in the gateway.
     */
    networks: pulumi.Input<pulumi.Input<inputs.mediaconnect.GatewayNetworkArgs>[]>;
}