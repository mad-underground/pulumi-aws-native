// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Wisdom::Assistant Resource Type
 */
export function getAssistant(args: GetAssistantArgs, opts?: pulumi.InvokeOptions): Promise<GetAssistantResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:wisdom:getAssistant", {
        "assistantId": args.assistantId,
    }, opts);
}

export interface GetAssistantArgs {
    assistantId: string;
}

export interface GetAssistantResult {
    readonly assistantArn?: string;
    readonly assistantId?: string;
}
/**
 * Definition of AWS::Wisdom::Assistant Resource Type
 */
export function getAssistantOutput(args: GetAssistantOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAssistantResult> {
    return pulumi.output(args).apply((a: any) => getAssistant(a, opts))
}

export interface GetAssistantOutputArgs {
    assistantId: pulumi.Input<string>;
}
