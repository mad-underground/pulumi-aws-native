// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Create and manage wireless gateways, including LoRa gateways.
 */
export class WirelessDevice extends pulumi.CustomResource {
    /**
     * Get an existing WirelessDevice resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WirelessDevice {
        return new WirelessDevice(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:iotwireless:WirelessDevice';

    /**
     * Returns true if the given object is an instance of WirelessDevice.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelessDevice {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelessDevice.__pulumiType;
    }

    /**
     * Wireless device arn. Returned after successful create.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Wireless device Id. Returned after successful create.
     */
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    /**
     * Wireless device description
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Wireless device destination name
     */
    public readonly destinationName!: pulumi.Output<string>;
    /**
     * The date and time when the most recent uplink was received.
     */
    public readonly lastUplinkReceivedAt!: pulumi.Output<string | undefined>;
    /**
     * The combination of Package, Station and Model which represents the version of the LoRaWAN Wireless Device.
     */
    public readonly loRaWan!: pulumi.Output<outputs.iotwireless.WirelessDeviceLoRaWanDevice | undefined>;
    /**
     * Wireless device name
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * A list of key-value pairs that contain metadata for the device. Currently not supported, will not create if tags are passed.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * Thing arn. Passed into update to associate Thing with Wireless device.
     */
    public readonly thingArn!: pulumi.Output<string | undefined>;
    /**
     * Thing Arn. If there is a Thing created, this can be returned with a Get call.
     */
    public /*out*/ readonly thingName!: pulumi.Output<string>;
    /**
     * Wireless device type, currently only Sidewalk and LoRa
     */
    public readonly type!: pulumi.Output<enums.iotwireless.WirelessDeviceType>;

    /**
     * Create a WirelessDevice resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WirelessDeviceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.destinationName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationName'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["destinationName"] = args ? args.destinationName : undefined;
            resourceInputs["lastUplinkReceivedAt"] = args ? args.lastUplinkReceivedAt : undefined;
            resourceInputs["loRaWan"] = args ? args.loRaWan : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["thingArn"] = args ? args.thingArn : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["thingName"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["destinationName"] = undefined /*out*/;
            resourceInputs["lastUplinkReceivedAt"] = undefined /*out*/;
            resourceInputs["loRaWan"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["thingArn"] = undefined /*out*/;
            resourceInputs["thingName"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WirelessDevice.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a WirelessDevice resource.
 */
export interface WirelessDeviceArgs {
    /**
     * Wireless device description
     */
    description?: pulumi.Input<string>;
    /**
     * Wireless device destination name
     */
    destinationName: pulumi.Input<string>;
    /**
     * The date and time when the most recent uplink was received.
     */
    lastUplinkReceivedAt?: pulumi.Input<string>;
    /**
     * The combination of Package, Station and Model which represents the version of the LoRaWAN Wireless Device.
     */
    loRaWan?: pulumi.Input<inputs.iotwireless.WirelessDeviceLoRaWanDeviceArgs>;
    /**
     * Wireless device name
     */
    name?: pulumi.Input<string>;
    /**
     * A list of key-value pairs that contain metadata for the device. Currently not supported, will not create if tags are passed.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * Thing arn. Passed into update to associate Thing with Wireless device.
     */
    thingArn?: pulumi.Input<string>;
    /**
     * Wireless device type, currently only Sidewalk and LoRa
     */
    type: pulumi.Input<enums.iotwireless.WirelessDeviceType>;
}
