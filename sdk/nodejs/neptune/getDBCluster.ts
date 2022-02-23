// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Neptune::DBCluster
 */
export function getDBCluster(args: GetDBClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetDBClusterResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:neptune:getDBCluster", {
        "id": args.id,
    }, opts);
}

export interface GetDBClusterArgs {
    id: string;
}

export interface GetDBClusterResult {
    readonly associatedRoles?: outputs.neptune.DBClusterRole[];
    readonly backupRetentionPeriod?: number;
    readonly clusterResourceId?: string;
    readonly dBClusterParameterGroupName?: string;
    readonly deletionProtection?: boolean;
    readonly enableCloudwatchLogsExports?: string[];
    readonly endpoint?: string;
    readonly iamAuthEnabled?: boolean;
    readonly id?: string;
    readonly port?: number;
    readonly preferredBackupWindow?: string;
    readonly preferredMaintenanceWindow?: string;
    readonly readEndpoint?: string;
    readonly tags?: outputs.neptune.DBClusterTag[];
    readonly vpcSecurityGroupIds?: string[];
}

export function getDBClusterOutput(args: GetDBClusterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDBClusterResult> {
    return pulumi.output(args).apply(a => getDBCluster(a, opts))
}

export interface GetDBClusterOutputArgs {
    id: pulumi.Input<string>;
}