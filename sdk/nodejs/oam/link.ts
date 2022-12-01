// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Oam::Link Resource Type
 */
export class Link extends pulumi.CustomResource {
    /**
     * Get an existing Link resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Link {
        return new Link(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:oam:Link';

    /**
     * Returns true if the given object is an instance of Link.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Link {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Link.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public /*out*/ readonly label!: pulumi.Output<string>;
    public readonly labelTemplate!: pulumi.Output<string>;
    public readonly resourceTypes!: pulumi.Output<enums.oam.LinkResourceType[]>;
    public readonly sinkIdentifier!: pulumi.Output<string>;
    /**
     * Tags to apply to the link
     */
    public readonly tags!: pulumi.Output<any | undefined>;

    /**
     * Create a Link resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LinkArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.labelTemplate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'labelTemplate'");
            }
            if ((!args || args.resourceTypes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceTypes'");
            }
            if ((!args || args.sinkIdentifier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sinkIdentifier'");
            }
            resourceInputs["labelTemplate"] = args ? args.labelTemplate : undefined;
            resourceInputs["resourceTypes"] = args ? args.resourceTypes : undefined;
            resourceInputs["sinkIdentifier"] = args ? args.sinkIdentifier : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["label"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["label"] = undefined /*out*/;
            resourceInputs["labelTemplate"] = undefined /*out*/;
            resourceInputs["resourceTypes"] = undefined /*out*/;
            resourceInputs["sinkIdentifier"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Link.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Link resource.
 */
export interface LinkArgs {
    labelTemplate: pulumi.Input<string>;
    resourceTypes: pulumi.Input<pulumi.Input<enums.oam.LinkResourceType>[]>;
    sinkIdentifier: pulumi.Input<string>;
    /**
     * Tags to apply to the link
     */
    tags?: any;
}
