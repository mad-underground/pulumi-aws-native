// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::SecurityHub::Standard resource represents the implementation of an individual AWS Security Hub Standard in your account. It requires you have SecurityHub enabled before you can enable the Standard.
 */
export class Standard extends pulumi.CustomResource {
    /**
     * Get an existing Standard resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Standard {
        return new Standard(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:securityhub:Standard';

    /**
     * Returns true if the given object is an instance of Standard.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Standard {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Standard.__pulumiType;
    }

    /**
     * StandardsControls to disable from this Standard.
     */
    public readonly disabledStandardsControls!: pulumi.Output<outputs.securityhub.StandardsControl[] | undefined>;
    /**
     * The ARN of the Standard being enabled
     */
    public readonly standardsArn!: pulumi.Output<string>;
    /**
     * The ARN of the StandardsSubscription for the account ID, region, and Standard.
     */
    public /*out*/ readonly standardsSubscriptionArn!: pulumi.Output<string>;

    /**
     * Create a Standard resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StandardArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.standardsArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'standardsArn'");
            }
            resourceInputs["disabledStandardsControls"] = args ? args.disabledStandardsControls : undefined;
            resourceInputs["standardsArn"] = args ? args.standardsArn : undefined;
            resourceInputs["standardsSubscriptionArn"] = undefined /*out*/;
        } else {
            resourceInputs["disabledStandardsControls"] = undefined /*out*/;
            resourceInputs["standardsArn"] = undefined /*out*/;
            resourceInputs["standardsSubscriptionArn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Standard.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Standard resource.
 */
export interface StandardArgs {
    /**
     * StandardsControls to disable from this Standard.
     */
    disabledStandardsControls?: pulumi.Input<pulumi.Input<inputs.securityhub.StandardsControlArgs>[]>;
    /**
     * The ARN of the Standard being enabled
     */
    standardsArn: pulumi.Input<string>;
}