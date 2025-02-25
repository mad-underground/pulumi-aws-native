// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::RefactorSpaces::Application Resource Type
 */
export function getApplication(args: GetApplicationArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:refactorspaces:getApplication", {
        "applicationIdentifier": args.applicationIdentifier,
        "environmentIdentifier": args.environmentIdentifier,
    }, opts);
}

export interface GetApplicationArgs {
    applicationIdentifier: string;
    environmentIdentifier: string;
}

export interface GetApplicationResult {
    readonly apiGatewayId?: string;
    readonly applicationIdentifier?: string;
    readonly arn?: string;
    readonly nlbArn?: string;
    readonly nlbName?: string;
    readonly proxyUrl?: string;
    readonly stageName?: string;
    /**
     * Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.
     */
    readonly tags?: outputs.Tag[];
    readonly vpcLinkId?: string;
}
/**
 * Definition of AWS::RefactorSpaces::Application Resource Type
 */
export function getApplicationOutput(args: GetApplicationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApplicationResult> {
    return pulumi.output(args).apply((a: any) => getApplication(a, opts))
}

export interface GetApplicationOutputArgs {
    applicationIdentifier: pulumi.Input<string>;
    environmentIdentifier: pulumi.Input<string>;
}
