// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Wireless Device Import Tasks
 *
 * @deprecated WirelessDeviceImportTask is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class WirelessDeviceImportTask extends pulumi.CustomResource {
    /**
     * Get an existing WirelessDeviceImportTask resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WirelessDeviceImportTask {
        pulumi.log.warn("WirelessDeviceImportTask is deprecated: WirelessDeviceImportTask is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new WirelessDeviceImportTask(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:iotwireless:WirelessDeviceImportTask';

    /**
     * Returns true if the given object is an instance of WirelessDeviceImportTask.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelessDeviceImportTask {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelessDeviceImportTask.__pulumiType;
    }

    /**
     * Arn for Wireless Device Import Task, Returned upon successful start.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Id for Wireless Device Import Task, Returned upon successful start.
     */
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    /**
     * CreationDate for import task
     */
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    /**
     * Destination Name for import task
     */
    public readonly destinationName!: pulumi.Output<string>;
    /**
     * Failed Imported Devices Count
     */
    public /*out*/ readonly failedImportedDevicesCount!: pulumi.Output<number>;
    /**
     * Initialized Imported Devices Count
     */
    public /*out*/ readonly initializedImportedDevicesCount!: pulumi.Output<number>;
    /**
     * Onboarded Imported Devices Count
     */
    public /*out*/ readonly onboardedImportedDevicesCount!: pulumi.Output<number>;
    /**
     * Pending Imported Devices Count
     */
    public /*out*/ readonly pendingImportedDevicesCount!: pulumi.Output<number>;
    /**
     * sidewalk contain file for created device and role
     */
    public readonly sidewalk!: pulumi.Output<outputs.iotwireless.SidewalkProperties>;
    /**
     * Status for import task
     */
    public /*out*/ readonly status!: pulumi.Output<enums.iotwireless.WirelessDeviceImportTaskStatus>;
    /**
     * StatusReason for import task
     */
    public /*out*/ readonly statusReason!: pulumi.Output<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a WirelessDeviceImportTask resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated WirelessDeviceImportTask is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: WirelessDeviceImportTaskArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("WirelessDeviceImportTask is deprecated: WirelessDeviceImportTask is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.destinationName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationName'");
            }
            if ((!args || args.sidewalk === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sidewalk'");
            }
            resourceInputs["destinationName"] = args ? args.destinationName : undefined;
            resourceInputs["sidewalk"] = args ? args.sidewalk : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["creationDate"] = undefined /*out*/;
            resourceInputs["failedImportedDevicesCount"] = undefined /*out*/;
            resourceInputs["initializedImportedDevicesCount"] = undefined /*out*/;
            resourceInputs["onboardedImportedDevicesCount"] = undefined /*out*/;
            resourceInputs["pendingImportedDevicesCount"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["statusReason"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["creationDate"] = undefined /*out*/;
            resourceInputs["destinationName"] = undefined /*out*/;
            resourceInputs["failedImportedDevicesCount"] = undefined /*out*/;
            resourceInputs["initializedImportedDevicesCount"] = undefined /*out*/;
            resourceInputs["onboardedImportedDevicesCount"] = undefined /*out*/;
            resourceInputs["pendingImportedDevicesCount"] = undefined /*out*/;
            resourceInputs["sidewalk"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["statusReason"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WirelessDeviceImportTask.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a WirelessDeviceImportTask resource.
 */
export interface WirelessDeviceImportTaskArgs {
    /**
     * Destination Name for import task
     */
    destinationName: pulumi.Input<string>;
    /**
     * sidewalk contain file for created device and role
     */
    sidewalk: pulumi.Input<inputs.iotwireless.SidewalkPropertiesArgs>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
