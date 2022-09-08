// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::CodePipeline::Webhook
 */
export function getWebhook(args: GetWebhookArgs, opts?: pulumi.InvokeOptions): Promise<GetWebhookResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:codepipeline:getWebhook", {
        "id": args.id,
    }, opts);
}

export interface GetWebhookArgs {
    id: string;
}

export interface GetWebhookResult {
    readonly authentication?: string;
    readonly authenticationConfiguration?: outputs.codepipeline.WebhookAuthConfiguration;
    readonly filters?: outputs.codepipeline.WebhookFilterRule[];
    readonly id?: string;
    readonly registerWithThirdParty?: boolean;
    readonly targetAction?: string;
    readonly targetPipeline?: string;
    readonly targetPipelineVersion?: number;
    readonly url?: string;
}

export function getWebhookOutput(args: GetWebhookOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetWebhookResult> {
    return pulumi.output(args).apply(a => getWebhook(a, opts))
}

export interface GetWebhookOutputArgs {
    id: pulumi.Input<string>;
}
