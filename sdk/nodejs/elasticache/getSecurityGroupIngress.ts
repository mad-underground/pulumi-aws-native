// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ElastiCache::SecurityGroupIngress
 */
export function getSecurityGroupIngress(args: GetSecurityGroupIngressArgs, opts?: pulumi.InvokeOptions): Promise<GetSecurityGroupIngressResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:elasticache:getSecurityGroupIngress", {
        "id": args.id,
    }, opts);
}

export interface GetSecurityGroupIngressArgs {
    id: string;
}

export interface GetSecurityGroupIngressResult {
    readonly cacheSecurityGroupName?: string;
    readonly ec2SecurityGroupName?: string;
    readonly ec2SecurityGroupOwnerId?: string;
    readonly id?: string;
}
/**
 * Resource Type definition for AWS::ElastiCache::SecurityGroupIngress
 */
export function getSecurityGroupIngressOutput(args: GetSecurityGroupIngressOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSecurityGroupIngressResult> {
    return pulumi.output(args).apply((a: any) => getSecurityGroupIngress(a, opts))
}

export interface GetSecurityGroupIngressOutputArgs {
    id: pulumi.Input<string>;
}
