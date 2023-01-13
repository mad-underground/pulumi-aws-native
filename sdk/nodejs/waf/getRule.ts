// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::WAF::Rule
 */
export function getRule(args: GetRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetRuleResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:waf:getRule", {
        "id": args.id,
    }, opts);
}

export interface GetRuleArgs {
    id: string;
}

export interface GetRuleResult {
    readonly id?: string;
    readonly predicates?: outputs.waf.RulePredicate[];
}
/**
 * Resource Type definition for AWS::WAF::Rule
 */
export function getRuleOutput(args: GetRuleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRuleResult> {
    return pulumi.output(args).apply((a: any) => getRule(a, opts))
}

export interface GetRuleOutputArgs {
    id: pulumi.Input<string>;
}
