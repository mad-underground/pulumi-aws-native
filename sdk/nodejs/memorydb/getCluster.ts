// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::MemoryDB::Cluster resource creates an Amazon MemoryDB Cluster.
 */
export function getCluster(args: GetClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:memorydb:getCluster", {
        "clusterName": args.clusterName,
    }, opts);
}

export interface GetClusterArgs {
    /**
     * The name of the cluster. This value must be unique as it also serves as the cluster identifier.
     */
    clusterName: string;
}

export interface GetClusterResult {
    /**
     * The name of the Access Control List to associate with the cluster.
     */
    readonly aCLName?: string;
    /**
     * The Amazon Resource Name (ARN) of the cluster.
     */
    readonly aRN?: string;
    /**
     * A flag that enables automatic minor version upgrade when set to true.
     *
     * You cannot modify the value of AutoMinorVersionUpgrade after the cluster is created. To enable AutoMinorVersionUpgrade on a cluster you must set AutoMinorVersionUpgrade to true when you create a cluster.
     */
    readonly autoMinorVersionUpgrade?: boolean;
    /**
     * The cluster endpoint.
     */
    readonly clusterEndpoint?: outputs.memorydb.ClusterEndpoint;
    /**
     * An optional description of the cluster.
     */
    readonly description?: string;
    /**
     * The Redis engine version used by the cluster.
     */
    readonly engineVersion?: string;
    /**
     * The user-supplied name of a final cluster snapshot. This is the unique name that identifies the snapshot. MemoryDB creates the snapshot, and then deletes the cluster immediately afterward.
     */
    readonly finalSnapshotName?: string;
    /**
     * Specifies the weekly time range during which maintenance on the cluster is performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period.
     */
    readonly maintenanceWindow?: string;
    /**
     * The compute and memory capacity of the nodes in the cluster.
     */
    readonly nodeType?: string;
    /**
     * The number of replicas to apply to each shard. The limit is 5.
     */
    readonly numReplicasPerShard?: number;
    /**
     * The number of shards the cluster will contain.
     */
    readonly numShards?: number;
    /**
     * The name of the parameter group associated with the cluster.
     */
    readonly parameterGroupName?: string;
    /**
     * The status of the parameter group used by the cluster.
     */
    readonly parameterGroupStatus?: string;
    /**
     * One or more Amazon VPC security groups associated with this cluster.
     */
    readonly securityGroupIds?: string[];
    /**
     * The number of days for which MemoryDB retains automatic snapshots before deleting them. For example, if you set SnapshotRetentionLimit to 5, a snapshot that was taken today is retained for 5 days before being deleted.
     */
    readonly snapshotRetentionLimit?: number;
    /**
     * The daily time range (in UTC) during which MemoryDB begins taking a daily snapshot of your cluster.
     */
    readonly snapshotWindow?: string;
    /**
     * The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic to which notifications are sent.
     */
    readonly snsTopicArn?: string;
    /**
     * The status of the Amazon SNS notification topic. Notifications are sent only if the status is enabled.
     */
    readonly snsTopicStatus?: string;
    /**
     * The status of the cluster. For example, Available, Updating, Creating.
     */
    readonly status?: string;
    /**
     * The name of the subnet group to be used for the cluster.
     */
    readonly subnetGroupName?: string;
    /**
     * An array of key-value pairs to apply to this cluster.
     */
    readonly tags?: outputs.memorydb.ClusterTag[];
}
/**
 * The AWS::MemoryDB::Cluster resource creates an Amazon MemoryDB Cluster.
 */
export function getClusterOutput(args: GetClusterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetClusterResult> {
    return pulumi.output(args).apply((a: any) => getCluster(a, opts))
}

export interface GetClusterOutputArgs {
    /**
     * The name of the cluster. This value must be unique as it also serves as the cluster identifier.
     */
    clusterName: pulumi.Input<string>;
}
