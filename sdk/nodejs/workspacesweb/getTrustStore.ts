// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::WorkSpacesWeb::TrustStore Resource Type
 */
export function getTrustStore(args: GetTrustStoreArgs, opts?: pulumi.InvokeOptions): Promise<GetTrustStoreResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:workspacesweb:getTrustStore", {
        "trustStoreArn": args.trustStoreArn,
    }, opts);
}

export interface GetTrustStoreArgs {
    trustStoreArn: string;
}

export interface GetTrustStoreResult {
    readonly associatedPortalArns?: string[];
    readonly certificateList?: string[];
    readonly tags?: outputs.Tag[];
    readonly trustStoreArn?: string;
}
/**
 * Definition of AWS::WorkSpacesWeb::TrustStore Resource Type
 */
export function getTrustStoreOutput(args: GetTrustStoreOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTrustStoreResult> {
    return pulumi.output(args).apply((a: any) => getTrustStore(a, opts))
}

export interface GetTrustStoreOutputArgs {
    trustStoreArn: pulumi.Input<string>;
}
