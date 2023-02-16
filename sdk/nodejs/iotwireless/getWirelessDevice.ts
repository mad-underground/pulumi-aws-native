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
export function getWirelessDevice(args: GetWirelessDeviceArgs, opts?: pulumi.InvokeOptions): Promise<GetWirelessDeviceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:iotwireless:getWirelessDevice", {
        "id": args.id,
    }, opts);
}

export interface GetWirelessDeviceArgs {
    /**
     * Wireless device Id. Returned after successful create.
     */
    id: string;
}

export interface GetWirelessDeviceResult {
    /**
     * Wireless device arn. Returned after successful create.
     */
    readonly arn?: string;
    /**
     * Wireless device description
     */
    readonly description?: string;
    /**
     * Wireless device destination name
     */
    readonly destinationName?: string;
    /**
     * Wireless device Id. Returned after successful create.
     */
    readonly id?: string;
    /**
     * The date and time when the most recent uplink was received.
     */
    readonly lastUplinkReceivedAt?: string;
    /**
     * The combination of Package, Station and Model which represents the version of the LoRaWAN Wireless Device.
     */
    readonly loRaWAN?: outputs.iotwireless.WirelessDeviceLoRaWANDevice;
    /**
     * Wireless device name
     */
    readonly name?: string;
    /**
     * A list of key-value pairs that contain metadata for the device. Currently not supported, will not create if tags are passed.
     */
    readonly tags?: outputs.iotwireless.WirelessDeviceTag[];
    /**
     * Thing arn. Passed into update to associate Thing with Wireless device.
     */
    readonly thingArn?: string;
    /**
     * Thing Arn. If there is a Thing created, this can be returned with a Get call.
     */
    readonly thingName?: string;
    /**
     * Wireless device type, currently only Sidewalk and LoRa
     */
    readonly type?: enums.iotwireless.WirelessDeviceType;
}
/**
 * Create and manage wireless gateways, including LoRa gateways.
 */
export function getWirelessDeviceOutput(args: GetWirelessDeviceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetWirelessDeviceResult> {
    return pulumi.output(args).apply((a: any) => getWirelessDevice(a, opts))
}

export interface GetWirelessDeviceOutputArgs {
    /**
     * Wireless device Id. Returned after successful create.
     */
    id: pulumi.Input<string>;
}
