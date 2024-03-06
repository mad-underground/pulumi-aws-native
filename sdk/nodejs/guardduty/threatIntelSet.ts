// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::GuardDuty::ThreatIntelSet
 */
export class ThreatIntelSet extends pulumi.CustomResource {
    /**
     * Get an existing ThreatIntelSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ThreatIntelSet {
        return new ThreatIntelSet(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:guardduty:ThreatIntelSet';

    /**
     * Returns true if the given object is an instance of ThreatIntelSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ThreatIntelSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ThreatIntelSet.__pulumiType;
    }

    public readonly activate!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    public readonly detectorId!: pulumi.Output<string | undefined>;
    public readonly format!: pulumi.Output<string>;
    public readonly location!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a ThreatIntelSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ThreatIntelSetArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.format === undefined) && !opts.urn) {
                throw new Error("Missing required property 'format'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            resourceInputs["activate"] = args ? args.activate : undefined;
            resourceInputs["detectorId"] = args ? args.detectorId : undefined;
            resourceInputs["format"] = args ? args.format : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["activate"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["detectorId"] = undefined /*out*/;
            resourceInputs["format"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["detectorId", "format"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(ThreatIntelSet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ThreatIntelSet resource.
 */
export interface ThreatIntelSetArgs {
    activate?: pulumi.Input<boolean>;
    detectorId?: pulumi.Input<string>;
    format: pulumi.Input<string>;
    location: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
