// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::WAFRegional::RateBasedRule
 */
export function getRateBasedRule(args: GetRateBasedRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetRateBasedRuleResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:wafregional:getRateBasedRule", {
        "id": args.id,
    }, opts);
}

export interface GetRateBasedRuleArgs {
    id: string;
}

export interface GetRateBasedRuleResult {
    readonly id?: string;
    readonly matchPredicates?: outputs.wafregional.RateBasedRulePredicate[];
    readonly rateLimit?: number;
}
/**
 * Resource Type definition for AWS::WAFRegional::RateBasedRule
 */
export function getRateBasedRuleOutput(args: GetRateBasedRuleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRateBasedRuleResult> {
    return pulumi.output(args).apply((a: any) => getRateBasedRule(a, opts))
}

export interface GetRateBasedRuleOutputArgs {
    id: pulumi.Input<string>;
}
