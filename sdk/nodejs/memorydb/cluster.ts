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
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:memorydb:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    /**
     * The name of the Access Control List to associate with the cluster.
     */
    public readonly aCLName!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the cluster.
     */
    public /*out*/ readonly aRN!: pulumi.Output<string>;
    /**
     * A flag that enables automatic minor version upgrade when set to true.
     *
     * You cannot modify the value of AutoMinorVersionUpgrade after the cluster is created. To enable AutoMinorVersionUpgrade on a cluster you must set AutoMinorVersionUpgrade to true when you create a cluster.
     */
    public readonly autoMinorVersionUpgrade!: pulumi.Output<boolean | undefined>;
    /**
     * The cluster endpoint.
     */
    public readonly clusterEndpoint!: pulumi.Output<outputs.memorydb.ClusterEndpoint | undefined>;
    /**
     * The name of the cluster. This value must be unique as it also serves as the cluster identifier.
     */
    public readonly clusterName!: pulumi.Output<string>;
    /**
     * An optional description of the cluster.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The Redis engine version used by the cluster.
     */
    public readonly engineVersion!: pulumi.Output<string | undefined>;
    /**
     * The user-supplied name of a final cluster snapshot. This is the unique name that identifies the snapshot. MemoryDB creates the snapshot, and then deletes the cluster immediately afterward.
     */
    public readonly finalSnapshotName!: pulumi.Output<string | undefined>;
    /**
     * The ID of the KMS key used to encrypt the cluster.
     */
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    /**
     * Specifies the weekly time range during which maintenance on the cluster is performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period.
     */
    public readonly maintenanceWindow!: pulumi.Output<string | undefined>;
    /**
     * The compute and memory capacity of the nodes in the cluster.
     */
    public readonly nodeType!: pulumi.Output<string>;
    /**
     * The number of replicas to apply to each shard. The limit is 5.
     */
    public readonly numReplicasPerShard!: pulumi.Output<number | undefined>;
    /**
     * The number of shards the cluster will contain.
     */
    public readonly numShards!: pulumi.Output<number | undefined>;
    /**
     * The name of the parameter group associated with the cluster.
     */
    public readonly parameterGroupName!: pulumi.Output<string | undefined>;
    /**
     * The status of the parameter group used by the cluster.
     */
    public /*out*/ readonly parameterGroupStatus!: pulumi.Output<string>;
    /**
     * The port number on which each member of the cluster accepts connections.
     */
    public readonly port!: pulumi.Output<number | undefined>;
    /**
     * One or more Amazon VPC security groups associated with this cluster.
     */
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * A list of Amazon Resource Names (ARN) that uniquely identify the RDB snapshot files stored in Amazon S3. The snapshot files are used to populate the new cluster. The Amazon S3 object name in the ARN cannot contain any commas.
     */
    public readonly snapshotArns!: pulumi.Output<string[] | undefined>;
    /**
     * The name of a snapshot from which to restore data into the new cluster. The snapshot status changes to restoring while the new cluster is being created.
     */
    public readonly snapshotName!: pulumi.Output<string | undefined>;
    /**
     * The number of days for which MemoryDB retains automatic snapshots before deleting them. For example, if you set SnapshotRetentionLimit to 5, a snapshot that was taken today is retained for 5 days before being deleted.
     */
    public readonly snapshotRetentionLimit!: pulumi.Output<number | undefined>;
    /**
     * The daily time range (in UTC) during which MemoryDB begins taking a daily snapshot of your cluster.
     */
    public readonly snapshotWindow!: pulumi.Output<string | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic to which notifications are sent.
     */
    public readonly snsTopicArn!: pulumi.Output<string | undefined>;
    /**
     * The status of the Amazon SNS notification topic. Notifications are sent only if the status is enabled.
     */
    public readonly snsTopicStatus!: pulumi.Output<string | undefined>;
    /**
     * The status of the cluster. For example, Available, Updating, Creating.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The name of the subnet group to be used for the cluster.
     */
    public readonly subnetGroupName!: pulumi.Output<string | undefined>;
    /**
     * A flag that enables in-transit encryption when set to true.
     *
     * You cannot modify the value of TransitEncryptionEnabled after the cluster is created. To enable in-transit encryption on a cluster you must set TransitEncryptionEnabled to true when you create a cluster.
     */
    public readonly tLSEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * An array of key-value pairs to apply to this cluster.
     */
    public readonly tags!: pulumi.Output<outputs.memorydb.ClusterTag[] | undefined>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.aCLName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'aCLName'");
            }
            if ((!args || args.nodeType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeType'");
            }
            resourceInputs["aCLName"] = args ? args.aCLName : undefined;
            resourceInputs["autoMinorVersionUpgrade"] = args ? args.autoMinorVersionUpgrade : undefined;
            resourceInputs["clusterEndpoint"] = args ? args.clusterEndpoint : undefined;
            resourceInputs["clusterName"] = args ? args.clusterName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["engineVersion"] = args ? args.engineVersion : undefined;
            resourceInputs["finalSnapshotName"] = args ? args.finalSnapshotName : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            resourceInputs["maintenanceWindow"] = args ? args.maintenanceWindow : undefined;
            resourceInputs["nodeType"] = args ? args.nodeType : undefined;
            resourceInputs["numReplicasPerShard"] = args ? args.numReplicasPerShard : undefined;
            resourceInputs["numShards"] = args ? args.numShards : undefined;
            resourceInputs["parameterGroupName"] = args ? args.parameterGroupName : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["snapshotArns"] = args ? args.snapshotArns : undefined;
            resourceInputs["snapshotName"] = args ? args.snapshotName : undefined;
            resourceInputs["snapshotRetentionLimit"] = args ? args.snapshotRetentionLimit : undefined;
            resourceInputs["snapshotWindow"] = args ? args.snapshotWindow : undefined;
            resourceInputs["snsTopicArn"] = args ? args.snsTopicArn : undefined;
            resourceInputs["snsTopicStatus"] = args ? args.snsTopicStatus : undefined;
            resourceInputs["subnetGroupName"] = args ? args.subnetGroupName : undefined;
            resourceInputs["tLSEnabled"] = args ? args.tLSEnabled : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["aRN"] = undefined /*out*/;
            resourceInputs["parameterGroupStatus"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        } else {
            resourceInputs["aCLName"] = undefined /*out*/;
            resourceInputs["aRN"] = undefined /*out*/;
            resourceInputs["autoMinorVersionUpgrade"] = undefined /*out*/;
            resourceInputs["clusterEndpoint"] = undefined /*out*/;
            resourceInputs["clusterName"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["engineVersion"] = undefined /*out*/;
            resourceInputs["finalSnapshotName"] = undefined /*out*/;
            resourceInputs["kmsKeyId"] = undefined /*out*/;
            resourceInputs["maintenanceWindow"] = undefined /*out*/;
            resourceInputs["nodeType"] = undefined /*out*/;
            resourceInputs["numReplicasPerShard"] = undefined /*out*/;
            resourceInputs["numShards"] = undefined /*out*/;
            resourceInputs["parameterGroupName"] = undefined /*out*/;
            resourceInputs["parameterGroupStatus"] = undefined /*out*/;
            resourceInputs["port"] = undefined /*out*/;
            resourceInputs["securityGroupIds"] = undefined /*out*/;
            resourceInputs["snapshotArns"] = undefined /*out*/;
            resourceInputs["snapshotName"] = undefined /*out*/;
            resourceInputs["snapshotRetentionLimit"] = undefined /*out*/;
            resourceInputs["snapshotWindow"] = undefined /*out*/;
            resourceInputs["snsTopicArn"] = undefined /*out*/;
            resourceInputs["snsTopicStatus"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["subnetGroupName"] = undefined /*out*/;
            resourceInputs["tLSEnabled"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Cluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * The name of the Access Control List to associate with the cluster.
     */
    aCLName: pulumi.Input<string>;
    /**
     * A flag that enables automatic minor version upgrade when set to true.
     *
     * You cannot modify the value of AutoMinorVersionUpgrade after the cluster is created. To enable AutoMinorVersionUpgrade on a cluster you must set AutoMinorVersionUpgrade to true when you create a cluster.
     */
    autoMinorVersionUpgrade?: pulumi.Input<boolean>;
    /**
     * The cluster endpoint.
     */
    clusterEndpoint?: pulumi.Input<inputs.memorydb.ClusterEndpointArgs>;
    /**
     * The name of the cluster. This value must be unique as it also serves as the cluster identifier.
     */
    clusterName?: pulumi.Input<string>;
    /**
     * An optional description of the cluster.
     */
    description?: pulumi.Input<string>;
    /**
     * The Redis engine version used by the cluster.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * The user-supplied name of a final cluster snapshot. This is the unique name that identifies the snapshot. MemoryDB creates the snapshot, and then deletes the cluster immediately afterward.
     */
    finalSnapshotName?: pulumi.Input<string>;
    /**
     * The ID of the KMS key used to encrypt the cluster.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * Specifies the weekly time range during which maintenance on the cluster is performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period.
     */
    maintenanceWindow?: pulumi.Input<string>;
    /**
     * The compute and memory capacity of the nodes in the cluster.
     */
    nodeType: pulumi.Input<string>;
    /**
     * The number of replicas to apply to each shard. The limit is 5.
     */
    numReplicasPerShard?: pulumi.Input<number>;
    /**
     * The number of shards the cluster will contain.
     */
    numShards?: pulumi.Input<number>;
    /**
     * The name of the parameter group associated with the cluster.
     */
    parameterGroupName?: pulumi.Input<string>;
    /**
     * The port number on which each member of the cluster accepts connections.
     */
    port?: pulumi.Input<number>;
    /**
     * One or more Amazon VPC security groups associated with this cluster.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of Amazon Resource Names (ARN) that uniquely identify the RDB snapshot files stored in Amazon S3. The snapshot files are used to populate the new cluster. The Amazon S3 object name in the ARN cannot contain any commas.
     */
    snapshotArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of a snapshot from which to restore data into the new cluster. The snapshot status changes to restoring while the new cluster is being created.
     */
    snapshotName?: pulumi.Input<string>;
    /**
     * The number of days for which MemoryDB retains automatic snapshots before deleting them. For example, if you set SnapshotRetentionLimit to 5, a snapshot that was taken today is retained for 5 days before being deleted.
     */
    snapshotRetentionLimit?: pulumi.Input<number>;
    /**
     * The daily time range (in UTC) during which MemoryDB begins taking a daily snapshot of your cluster.
     */
    snapshotWindow?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic to which notifications are sent.
     */
    snsTopicArn?: pulumi.Input<string>;
    /**
     * The status of the Amazon SNS notification topic. Notifications are sent only if the status is enabled.
     */
    snsTopicStatus?: pulumi.Input<string>;
    /**
     * The name of the subnet group to be used for the cluster.
     */
    subnetGroupName?: pulumi.Input<string>;
    /**
     * A flag that enables in-transit encryption when set to true.
     *
     * You cannot modify the value of TransitEncryptionEnabled after the cluster is created. To enable in-transit encryption on a cluster you must set TransitEncryptionEnabled to true when you create a cluster.
     */
    tLSEnabled?: pulumi.Input<boolean>;
    /**
     * An array of key-value pairs to apply to this cluster.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.memorydb.ClusterTagArgs>[]>;
}
