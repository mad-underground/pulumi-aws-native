// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * You must first create and start the CC configuration recorder in order to create CC managed rules with CFNlong. For more information, see [Managing the Configuration Recorder](https://docs.aws.amazon.com/config/latest/developerguide/stop-start-recorder.html).
 *  Adds or updates an CC rule to evaluate if your AWS resources comply with your desired configurations. For information on how many CC rules you can have per account, see [Service Limits](https://docs.aws.amazon.com/config/latest/developerguide/configlimits.html) in the *Developer Guide*.
 *  There are two types of rules: *Managed Rules* and *Custom Rules*. You can use the ``ConfigRule`` resource to create both CC Managed Rules and CC Custom Rules.
 *  CC Managed Rules are predefined, customizable rules created by CC. For a list of managed rules, see [List of Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html). If you are adding an CC managed rule, you must specify the rule's identifier for the ``SourceIdentifier`` key.
 *  CC Custom Rules are rules that you create from scratch. There are two ways to create CC custom rules: with Lambda functions ([Developer Guide](https://docs.aws.amazon.com/config/latest/developerguide/gettingstarted-concepts.html#gettingstarted-concepts-function)) and with CFNGUARDshort ([Guard GitHub Repository](https://docs.aws.amazon.com/https://github.com/aws-cloudformation/cloudformation-guard)), a policy-as-code language. CC custom rules created with LAMlong are called *Custom Lambda Rules* and CC custom rules created with CFNGUARDshort are called *Custom Policy Rules*.
 *  If you are adding a new CC Custom LAM rule, you first need to create an LAMlong function that the rule invokes to evaluate your resources. When you use the ``ConfigRule`` resource to add a Custom LAM rule to CC, you must specify the Amazon Resource Name (ARN) that LAMlong assigns to the function. You specify the ARN in the ``SourceIdentifier`` key. This key is part of the ``Source`` object, which is part of the ``ConfigRule`` object.
 *  For any new CC rule that you add, specify the ``ConfigRuleName`` in the ``ConfigRule`` object. Do not specify the ``ConfigRuleArn`` or the ``ConfigRuleId``. These values are generated by CC for new rules.
 *  If you are updating a rule that you added previously, you can specify the rule by ``ConfigRuleName``, ``ConfigRuleId``, or ``ConfigRuleArn`` in the ``ConfigRule`` data type that you use in this request.
 *  For more information about developing and using CC rules, see [Evaluating Resources with Rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config.html) in the *Developer Guide*.
 */
export class ConfigRule extends pulumi.CustomResource {
    /**
     * Get an existing ConfigRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ConfigRule {
        return new ConfigRule(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:configuration:ConfigRule';

    /**
     * Returns true if the given object is an instance of ConfigRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConfigRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConfigRule.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Indicates whether an AWS resource or CC rule is compliant and provides the number of contributors that affect the compliance.
     */
    public readonly compliance!: pulumi.Output<outputs.configuration.ComplianceProperties | undefined>;
    public /*out*/ readonly configRuleId!: pulumi.Output<string>;
    /**
     * A name for the CC rule. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the rule name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
     */
    public readonly configRuleName!: pulumi.Output<string | undefined>;
    /**
     * The description that you provide for the CC rule.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The modes the CC rule can be evaluated in. The valid values are distinct objects. By default, the value is Detective evaluation mode only.
     */
    public readonly evaluationModes!: pulumi.Output<outputs.configuration.ConfigRuleEvaluationModeConfiguration[] | undefined>;
    /**
     * A string, in JSON format, that is passed to the CC rule Lambda function.
     *
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Config::ConfigRule` for more information about the expected schema for this property.
     */
    public readonly inputParameters!: pulumi.Output<any | undefined>;
    /**
     * The maximum frequency with which CC runs evaluations for a rule. You can specify a value for ``MaximumExecutionFrequency`` when:
     *   +  You are using an AWS managed rule that is triggered at a periodic frequency.
     *   +  Your custom rule is triggered when CC delivers the configuration snapshot. For more information, see [ConfigSnapshotDeliveryProperties](https://docs.aws.amazon.com/config/latest/APIReference/API_ConfigSnapshotDeliveryProperties.html).
     *   
     *   By default, rules with a periodic trigger are evaluated every 24 hours. To change the frequency, specify a valid value for the ``MaximumExecutionFrequency`` parameter.
     */
    public readonly maximumExecutionFrequency!: pulumi.Output<string | undefined>;
    /**
     * Defines which resources can trigger an evaluation for the rule. The scope can include one or more resource types, a combination of one resource type and one resource ID, or a combination of a tag key and value. Specify a scope to constrain the resources that can trigger an evaluation for the rule. If you do not specify a scope, evaluations are triggered when any resource in the recording group changes.
     *   The scope can be empty.
     */
    public readonly scope!: pulumi.Output<outputs.configuration.ConfigRuleScope | undefined>;
    /**
     * Provides the rule owner (```` for managed rules, ``CUSTOM_POLICY`` for Custom Policy rules, and ``CUSTOM_LAMBDA`` for Custom Lambda rules), the rule identifier, and the notifications that cause the function to evaluate your AWS resources.
     */
    public readonly source!: pulumi.Output<outputs.configuration.ConfigRuleSource>;

    /**
     * Create a ConfigRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConfigRuleArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.source === undefined) && !opts.urn) {
                throw new Error("Missing required property 'source'");
            }
            resourceInputs["compliance"] = args ? args.compliance : undefined;
            resourceInputs["configRuleName"] = args ? args.configRuleName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["evaluationModes"] = args ? args.evaluationModes : undefined;
            resourceInputs["inputParameters"] = args ? args.inputParameters : undefined;
            resourceInputs["maximumExecutionFrequency"] = args ? args.maximumExecutionFrequency : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["configRuleId"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["compliance"] = undefined /*out*/;
            resourceInputs["configRuleId"] = undefined /*out*/;
            resourceInputs["configRuleName"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["evaluationModes"] = undefined /*out*/;
            resourceInputs["inputParameters"] = undefined /*out*/;
            resourceInputs["maximumExecutionFrequency"] = undefined /*out*/;
            resourceInputs["scope"] = undefined /*out*/;
            resourceInputs["source"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["configRuleName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(ConfigRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ConfigRule resource.
 */
export interface ConfigRuleArgs {
    /**
     * Indicates whether an AWS resource or CC rule is compliant and provides the number of contributors that affect the compliance.
     */
    compliance?: pulumi.Input<inputs.configuration.CompliancePropertiesArgs>;
    /**
     * A name for the CC rule. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the rule name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
     */
    configRuleName?: pulumi.Input<string>;
    /**
     * The description that you provide for the CC rule.
     */
    description?: pulumi.Input<string>;
    /**
     * The modes the CC rule can be evaluated in. The valid values are distinct objects. By default, the value is Detective evaluation mode only.
     */
    evaluationModes?: pulumi.Input<pulumi.Input<inputs.configuration.ConfigRuleEvaluationModeConfigurationArgs>[]>;
    /**
     * A string, in JSON format, that is passed to the CC rule Lambda function.
     *
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Config::ConfigRule` for more information about the expected schema for this property.
     */
    inputParameters?: any;
    /**
     * The maximum frequency with which CC runs evaluations for a rule. You can specify a value for ``MaximumExecutionFrequency`` when:
     *   +  You are using an AWS managed rule that is triggered at a periodic frequency.
     *   +  Your custom rule is triggered when CC delivers the configuration snapshot. For more information, see [ConfigSnapshotDeliveryProperties](https://docs.aws.amazon.com/config/latest/APIReference/API_ConfigSnapshotDeliveryProperties.html).
     *   
     *   By default, rules with a periodic trigger are evaluated every 24 hours. To change the frequency, specify a valid value for the ``MaximumExecutionFrequency`` parameter.
     */
    maximumExecutionFrequency?: pulumi.Input<string>;
    /**
     * Defines which resources can trigger an evaluation for the rule. The scope can include one or more resource types, a combination of one resource type and one resource ID, or a combination of a tag key and value. Specify a scope to constrain the resources that can trigger an evaluation for the rule. If you do not specify a scope, evaluations are triggered when any resource in the recording group changes.
     *   The scope can be empty.
     */
    scope?: pulumi.Input<inputs.configuration.ConfigRuleScopeArgs>;
    /**
     * Provides the rule owner (```` for managed rules, ``CUSTOM_POLICY`` for Custom Policy rules, and ``CUSTOM_LAMBDA`` for Custom Lambda rules), the rule identifier, and the notifications that cause the function to evaluate your AWS resources.
     */
    source: pulumi.Input<inputs.configuration.ConfigRuleSourceArgs>;
}
