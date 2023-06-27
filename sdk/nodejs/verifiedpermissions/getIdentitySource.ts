// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::VerifiedPermissions::IdentitySource Resource Type
 */
export function getIdentitySource(args: GetIdentitySourceArgs, opts?: pulumi.InvokeOptions): Promise<GetIdentitySourceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:verifiedpermissions:getIdentitySource", {
        "identitySourceId": args.identitySourceId,
        "policyStoreId": args.policyStoreId,
    }, opts);
}

export interface GetIdentitySourceArgs {
    identitySourceId: string;
    policyStoreId: string;
}

export interface GetIdentitySourceResult {
    readonly configuration?: outputs.verifiedpermissions.IdentitySourceConfiguration;
    readonly details?: outputs.verifiedpermissions.IdentitySourceDetails;
    readonly identitySourceId?: string;
    readonly principalEntityType?: string;
}
/**
 * Definition of AWS::VerifiedPermissions::IdentitySource Resource Type
 */
export function getIdentitySourceOutput(args: GetIdentitySourceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIdentitySourceResult> {
    return pulumi.output(args).apply((a: any) => getIdentitySource(a, opts))
}

export interface GetIdentitySourceOutputArgs {
    identitySourceId: pulumi.Input<string>;
    policyStoreId: pulumi.Input<string>;
}
