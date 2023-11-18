// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::MediaLive::Multiplex
 */
export class Multiplex extends pulumi.CustomResource {
    /**
     * Get an existing Multiplex resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Multiplex {
        return new Multiplex(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:medialive:Multiplex';

    /**
     * Returns true if the given object is an instance of Multiplex.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Multiplex {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Multiplex.__pulumiType;
    }

    /**
     * The unique arn of the multiplex.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A list of availability zones for the multiplex.
     */
    public readonly availabilityZones!: pulumi.Output<string[]>;
    /**
     * A list of the multiplex output destinations.
     */
    public readonly destinations!: pulumi.Output<outputs.medialive.MultiplexOutputDestination[] | undefined>;
    /**
     * Configuration for a multiplex event.
     */
    public readonly multiplexSettings!: pulumi.Output<outputs.medialive.MultiplexSettings>;
    /**
     * Name of multiplex.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The number of currently healthy pipelines.
     */
    public /*out*/ readonly pipelinesRunningCount!: pulumi.Output<number>;
    /**
     * The number of programs in the multiplex.
     */
    public /*out*/ readonly programCount!: pulumi.Output<number>;
    public /*out*/ readonly state!: pulumi.Output<enums.medialive.MultiplexState>;
    /**
     * A collection of key-value pairs.
     */
    public readonly tags!: pulumi.Output<outputs.medialive.MultiplexTags[] | undefined>;

    /**
     * Create a Multiplex resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MultiplexArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.availabilityZones === undefined) && !opts.urn) {
                throw new Error("Missing required property 'availabilityZones'");
            }
            if ((!args || args.multiplexSettings === undefined) && !opts.urn) {
                throw new Error("Missing required property 'multiplexSettings'");
            }
            resourceInputs["availabilityZones"] = args ? args.availabilityZones : undefined;
            resourceInputs["destinations"] = args ? args.destinations : undefined;
            resourceInputs["multiplexSettings"] = args ? args.multiplexSettings : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["pipelinesRunningCount"] = undefined /*out*/;
            resourceInputs["programCount"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["availabilityZones"] = undefined /*out*/;
            resourceInputs["destinations"] = undefined /*out*/;
            resourceInputs["multiplexSettings"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["pipelinesRunningCount"] = undefined /*out*/;
            resourceInputs["programCount"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["availabilityZones[*]"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Multiplex.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Multiplex resource.
 */
export interface MultiplexArgs {
    /**
     * A list of availability zones for the multiplex.
     */
    availabilityZones: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of the multiplex output destinations.
     */
    destinations?: pulumi.Input<pulumi.Input<inputs.medialive.MultiplexOutputDestinationArgs>[]>;
    /**
     * Configuration for a multiplex event.
     */
    multiplexSettings: pulumi.Input<inputs.medialive.MultiplexSettingsArgs>;
    /**
     * Name of multiplex.
     */
    name?: pulumi.Input<string>;
    /**
     * A collection of key-value pairs.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.medialive.MultiplexTagsArgs>[]>;
}