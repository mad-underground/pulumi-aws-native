// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::WAFRegional::WebACLAssociation
 */
export function getWebAclAssociation(args: GetWebAclAssociationArgs, opts?: pulumi.InvokeOptions): Promise<GetWebAclAssociationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:wafregional:getWebAclAssociation", {
        "id": args.id,
    }, opts);
}

export interface GetWebAclAssociationArgs {
    id: string;
}

export interface GetWebAclAssociationResult {
    readonly id?: string;
}
/**
 * Resource Type definition for AWS::WAFRegional::WebACLAssociation
 */
export function getWebAclAssociationOutput(args: GetWebAclAssociationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetWebAclAssociationResult> {
    return pulumi.output(args).apply((a: any) => getWebAclAssociation(a, opts))
}

export interface GetWebAclAssociationOutputArgs {
    id: pulumi.Input<string>;
}
