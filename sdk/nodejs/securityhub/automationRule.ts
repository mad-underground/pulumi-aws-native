// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The ``AWS::SecurityHub::AutomationRule`` resource specifies an automation rule based on input parameters. For more information, see [Automation rules](https://docs.aws.amazon.com/securityhub/latest/userguide/automation-rules.html) in the *User Guide*.
 */
export class AutomationRule extends pulumi.CustomResource {
    /**
     * Get an existing AutomationRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AutomationRule {
        return new AutomationRule(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:securityhub:AutomationRule';

    /**
     * Returns true if the given object is an instance of AutomationRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AutomationRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AutomationRule.__pulumiType;
    }

    public readonly actions!: pulumi.Output<outputs.securityhub.AutomationRulesAction[] | undefined>;
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    public /*out*/ readonly createdBy!: pulumi.Output<string>;
    /**
     * A set of [Security Finding Format (ASFF)](https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-findings-format.html) finding field attributes and corresponding expected values that ASH uses to filter findings. If a rule is enabled and a finding matches the criteria specified in this parameter, ASH applies the rule action to the finding.
     */
    public readonly criteria!: pulumi.Output<outputs.securityhub.AutomationRulesFindingFilters | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly isTerminal!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly ruleArn!: pulumi.Output<string>;
    public readonly ruleName!: pulumi.Output<string | undefined>;
    public readonly ruleOrder!: pulumi.Output<number | undefined>;
    /**
     * Whether the rule is active after it is created. If this parameter is equal to ``ENABLED``, ASH applies the rule to findings and finding updates after the rule is created.
     */
    public readonly ruleStatus!: pulumi.Output<enums.securityhub.AutomationRuleRuleStatus | undefined>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a AutomationRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AutomationRuleArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["actions"] = args ? args.actions : undefined;
            resourceInputs["criteria"] = args ? args.criteria : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["isTerminal"] = args ? args.isTerminal : undefined;
            resourceInputs["ruleName"] = args ? args.ruleName : undefined;
            resourceInputs["ruleOrder"] = args ? args.ruleOrder : undefined;
            resourceInputs["ruleStatus"] = args ? args.ruleStatus : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["createdBy"] = undefined /*out*/;
            resourceInputs["ruleArn"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        } else {
            resourceInputs["actions"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["createdBy"] = undefined /*out*/;
            resourceInputs["criteria"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["isTerminal"] = undefined /*out*/;
            resourceInputs["ruleArn"] = undefined /*out*/;
            resourceInputs["ruleName"] = undefined /*out*/;
            resourceInputs["ruleOrder"] = undefined /*out*/;
            resourceInputs["ruleStatus"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AutomationRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a AutomationRule resource.
 */
export interface AutomationRuleArgs {
    actions?: pulumi.Input<pulumi.Input<inputs.securityhub.AutomationRulesActionArgs>[]>;
    /**
     * A set of [Security Finding Format (ASFF)](https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-findings-format.html) finding field attributes and corresponding expected values that ASH uses to filter findings. If a rule is enabled and a finding matches the criteria specified in this parameter, ASH applies the rule action to the finding.
     */
    criteria?: pulumi.Input<inputs.securityhub.AutomationRulesFindingFiltersArgs>;
    description?: pulumi.Input<string>;
    isTerminal?: pulumi.Input<boolean>;
    ruleName?: pulumi.Input<string>;
    ruleOrder?: pulumi.Input<number>;
    /**
     * Whether the rule is active after it is created. If this parameter is equal to ``ENABLED``, ASH applies the rule to findings and finding updates after the rule is created.
     */
    ruleStatus?: pulumi.Input<enums.securityhub.AutomationRuleRuleStatus>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
