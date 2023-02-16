// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SES::ReceiptRule
 *
 * @deprecated ReceiptRule is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class ReceiptRule extends pulumi.CustomResource {
    /**
     * Get an existing ReceiptRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ReceiptRule {
        pulumi.log.warn("ReceiptRule is deprecated: ReceiptRule is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new ReceiptRule(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ses:ReceiptRule';

    /**
     * Returns true if the given object is an instance of ReceiptRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReceiptRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReceiptRule.__pulumiType;
    }

    public readonly after!: pulumi.Output<string | undefined>;
    public readonly rule!: pulumi.Output<outputs.ses.ReceiptRuleRule>;
    public readonly ruleSetName!: pulumi.Output<string>;

    /**
     * Create a ReceiptRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated ReceiptRule is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: ReceiptRuleArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("ReceiptRule is deprecated: ReceiptRule is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.rule === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rule'");
            }
            if ((!args || args.ruleSetName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleSetName'");
            }
            resourceInputs["after"] = args ? args.after : undefined;
            resourceInputs["rule"] = args ? args.rule : undefined;
            resourceInputs["ruleSetName"] = args ? args.ruleSetName : undefined;
        } else {
            resourceInputs["after"] = undefined /*out*/;
            resourceInputs["rule"] = undefined /*out*/;
            resourceInputs["ruleSetName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ReceiptRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ReceiptRule resource.
 */
export interface ReceiptRuleArgs {
    after?: pulumi.Input<string>;
    rule: pulumi.Input<inputs.ses.ReceiptRuleRuleArgs>;
    ruleSetName: pulumi.Input<string>;
}
