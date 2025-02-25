// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Lightsail::Disk
 */
export class Disk extends pulumi.CustomResource {
    /**
     * Get an existing Disk resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Disk {
        return new Disk(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:lightsail:Disk';

    /**
     * Returns true if the given object is an instance of Disk.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Disk {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Disk.__pulumiType;
    }

    /**
     * An array of objects representing the add-ons to enable for the new instance.
     */
    public readonly addOns!: pulumi.Output<outputs.lightsail.DiskAddOn[] | undefined>;
    /**
     * Name of the attached Lightsail Instance
     */
    public /*out*/ readonly attachedTo!: pulumi.Output<string>;
    /**
     * Attachment State of the Lightsail disk
     */
    public /*out*/ readonly attachmentState!: pulumi.Output<string>;
    /**
     * The Availability Zone in which to create your instance. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.
     */
    public readonly availabilityZone!: pulumi.Output<string | undefined>;
    public /*out*/ readonly diskArn!: pulumi.Output<string>;
    /**
     * The names to use for your new Lightsail disk.
     */
    public readonly diskName!: pulumi.Output<string>;
    /**
     * Iops of the Lightsail disk
     */
    public /*out*/ readonly iops!: pulumi.Output<number>;
    /**
     * Check is Disk is attached state
     */
    public /*out*/ readonly isAttached!: pulumi.Output<boolean>;
    public readonly location!: pulumi.Output<outputs.lightsail.DiskLocation | undefined>;
    /**
     * Path of the  attached Disk
     */
    public /*out*/ readonly path!: pulumi.Output<string>;
    /**
     * Resource type of Lightsail instance.
     */
    public /*out*/ readonly resourceType!: pulumi.Output<string>;
    /**
     * Size of the Lightsail disk
     */
    public readonly sizeInGb!: pulumi.Output<number>;
    /**
     * State of the Lightsail disk
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Support code to help identify any issues
     */
    public /*out*/ readonly supportCode!: pulumi.Output<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a Disk resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DiskArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.sizeInGb === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sizeInGb'");
            }
            resourceInputs["addOns"] = args ? args.addOns : undefined;
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["diskName"] = args ? args.diskName : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["sizeInGb"] = args ? args.sizeInGb : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["attachedTo"] = undefined /*out*/;
            resourceInputs["attachmentState"] = undefined /*out*/;
            resourceInputs["diskArn"] = undefined /*out*/;
            resourceInputs["iops"] = undefined /*out*/;
            resourceInputs["isAttached"] = undefined /*out*/;
            resourceInputs["path"] = undefined /*out*/;
            resourceInputs["resourceType"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["supportCode"] = undefined /*out*/;
        } else {
            resourceInputs["addOns"] = undefined /*out*/;
            resourceInputs["attachedTo"] = undefined /*out*/;
            resourceInputs["attachmentState"] = undefined /*out*/;
            resourceInputs["availabilityZone"] = undefined /*out*/;
            resourceInputs["diskArn"] = undefined /*out*/;
            resourceInputs["diskName"] = undefined /*out*/;
            resourceInputs["iops"] = undefined /*out*/;
            resourceInputs["isAttached"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["path"] = undefined /*out*/;
            resourceInputs["resourceType"] = undefined /*out*/;
            resourceInputs["sizeInGb"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["supportCode"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["availabilityZone", "diskName", "sizeInGb"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Disk.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Disk resource.
 */
export interface DiskArgs {
    /**
     * An array of objects representing the add-ons to enable for the new instance.
     */
    addOns?: pulumi.Input<pulumi.Input<inputs.lightsail.DiskAddOnArgs>[]>;
    /**
     * The Availability Zone in which to create your instance. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * The names to use for your new Lightsail disk.
     */
    diskName?: pulumi.Input<string>;
    location?: pulumi.Input<inputs.lightsail.DiskLocationArgs>;
    /**
     * Size of the Lightsail disk
     */
    sizeInGb: pulumi.Input<number>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
