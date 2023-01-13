// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Creates an AWS Firewall Manager policy.
 */
export function getPolicy(args: GetPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:fms:getPolicy", {
        "id": args.id,
    }, opts);
}

export interface GetPolicyArgs {
    id: string;
}

export interface GetPolicyResult {
    readonly arn?: string;
    readonly excludeMap?: outputs.fms.PolicyIEMap;
    readonly excludeResourceTags?: boolean;
    readonly id?: string;
    readonly includeMap?: outputs.fms.PolicyIEMap;
    readonly policyName?: string;
    readonly remediationEnabled?: boolean;
    readonly resourceTags?: outputs.fms.PolicyResourceTag[];
    readonly resourceType?: string;
    readonly resourceTypeList?: string[];
    readonly resourcesCleanUp?: boolean;
    readonly securityServicePolicyData?: outputs.fms.PolicySecurityServicePolicyData;
    readonly tags?: outputs.fms.PolicyTag[];
}
/**
 * Creates an AWS Firewall Manager policy.
 */
export function getPolicyOutput(args: GetPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPolicyResult> {
    return pulumi.output(args).apply((a: any) => getPolicy(a, opts))
}

export interface GetPolicyOutputArgs {
    id: pulumi.Input<string>;
}
