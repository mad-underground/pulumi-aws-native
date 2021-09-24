// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::RDS::DBCluster
 *
 * @deprecated DBCluster is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class DBCluster extends pulumi.CustomResource {
    /**
     * Get an existing DBCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DBCluster {
        pulumi.log.warn("DBCluster is deprecated: DBCluster is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new DBCluster(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:rds:DBCluster';

    /**
     * Returns true if the given object is an instance of DBCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DBCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DBCluster.__pulumiType;
    }

    public readonly associatedRoles!: pulumi.Output<outputs.rds.DBClusterDBClusterRole[] | undefined>;
    public readonly availabilityZones!: pulumi.Output<string[] | undefined>;
    public readonly backtrackWindow!: pulumi.Output<number | undefined>;
    public readonly backupRetentionPeriod!: pulumi.Output<number | undefined>;
    public readonly copyTagsToSnapshot!: pulumi.Output<boolean | undefined>;
    public readonly dBClusterIdentifier!: pulumi.Output<string | undefined>;
    public readonly dBClusterParameterGroupName!: pulumi.Output<string | undefined>;
    public readonly dBSubnetGroupName!: pulumi.Output<string | undefined>;
    public readonly databaseName!: pulumi.Output<string | undefined>;
    public readonly deletionProtection!: pulumi.Output<boolean | undefined>;
    public readonly enableCloudwatchLogsExports!: pulumi.Output<string[] | undefined>;
    public readonly enableHttpEndpoint!: pulumi.Output<boolean | undefined>;
    public readonly enableIAMDatabaseAuthentication!: pulumi.Output<boolean | undefined>;
    public readonly endpointAddress!: pulumi.Output<string | undefined>;
    public readonly endpointPort!: pulumi.Output<string | undefined>;
    public readonly engine!: pulumi.Output<string>;
    public readonly engineMode!: pulumi.Output<string | undefined>;
    public readonly engineVersion!: pulumi.Output<string | undefined>;
    public readonly globalClusterIdentifier!: pulumi.Output<string | undefined>;
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    public readonly masterUserPassword!: pulumi.Output<string | undefined>;
    public readonly masterUsername!: pulumi.Output<string | undefined>;
    public readonly port!: pulumi.Output<number | undefined>;
    public readonly preferredBackupWindow!: pulumi.Output<string | undefined>;
    public readonly preferredMaintenanceWindow!: pulumi.Output<string | undefined>;
    public readonly readEndpointAddress!: pulumi.Output<string | undefined>;
    public readonly replicationSourceIdentifier!: pulumi.Output<string | undefined>;
    public readonly restoreType!: pulumi.Output<string | undefined>;
    public readonly scalingConfiguration!: pulumi.Output<outputs.rds.DBClusterScalingConfiguration | undefined>;
    public readonly snapshotIdentifier!: pulumi.Output<string | undefined>;
    public readonly sourceDBClusterIdentifier!: pulumi.Output<string | undefined>;
    public readonly sourceRegion!: pulumi.Output<string | undefined>;
    public readonly storageEncrypted!: pulumi.Output<boolean | undefined>;
    public readonly tags!: pulumi.Output<outputs.rds.DBClusterTag[] | undefined>;
    public readonly useLatestRestorableTime!: pulumi.Output<boolean | undefined>;
    public readonly vpcSecurityGroupIds!: pulumi.Output<string[] | undefined>;

    /**
     * Create a DBCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated DBCluster is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: DBClusterArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("DBCluster is deprecated: DBCluster is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.engine === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engine'");
            }
            inputs["associatedRoles"] = args ? args.associatedRoles : undefined;
            inputs["availabilityZones"] = args ? args.availabilityZones : undefined;
            inputs["backtrackWindow"] = args ? args.backtrackWindow : undefined;
            inputs["backupRetentionPeriod"] = args ? args.backupRetentionPeriod : undefined;
            inputs["copyTagsToSnapshot"] = args ? args.copyTagsToSnapshot : undefined;
            inputs["dBClusterIdentifier"] = args ? args.dBClusterIdentifier : undefined;
            inputs["dBClusterParameterGroupName"] = args ? args.dBClusterParameterGroupName : undefined;
            inputs["dBSubnetGroupName"] = args ? args.dBSubnetGroupName : undefined;
            inputs["databaseName"] = args ? args.databaseName : undefined;
            inputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            inputs["enableCloudwatchLogsExports"] = args ? args.enableCloudwatchLogsExports : undefined;
            inputs["enableHttpEndpoint"] = args ? args.enableHttpEndpoint : undefined;
            inputs["enableIAMDatabaseAuthentication"] = args ? args.enableIAMDatabaseAuthentication : undefined;
            inputs["endpointAddress"] = args ? args.endpointAddress : undefined;
            inputs["endpointPort"] = args ? args.endpointPort : undefined;
            inputs["engine"] = args ? args.engine : undefined;
            inputs["engineMode"] = args ? args.engineMode : undefined;
            inputs["engineVersion"] = args ? args.engineVersion : undefined;
            inputs["globalClusterIdentifier"] = args ? args.globalClusterIdentifier : undefined;
            inputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            inputs["masterUserPassword"] = args ? args.masterUserPassword : undefined;
            inputs["masterUsername"] = args ? args.masterUsername : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["preferredBackupWindow"] = args ? args.preferredBackupWindow : undefined;
            inputs["preferredMaintenanceWindow"] = args ? args.preferredMaintenanceWindow : undefined;
            inputs["readEndpointAddress"] = args ? args.readEndpointAddress : undefined;
            inputs["replicationSourceIdentifier"] = args ? args.replicationSourceIdentifier : undefined;
            inputs["restoreType"] = args ? args.restoreType : undefined;
            inputs["scalingConfiguration"] = args ? args.scalingConfiguration : undefined;
            inputs["snapshotIdentifier"] = args ? args.snapshotIdentifier : undefined;
            inputs["sourceDBClusterIdentifier"] = args ? args.sourceDBClusterIdentifier : undefined;
            inputs["sourceRegion"] = args ? args.sourceRegion : undefined;
            inputs["storageEncrypted"] = args ? args.storageEncrypted : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["useLatestRestorableTime"] = args ? args.useLatestRestorableTime : undefined;
            inputs["vpcSecurityGroupIds"] = args ? args.vpcSecurityGroupIds : undefined;
        } else {
            inputs["associatedRoles"] = undefined /*out*/;
            inputs["availabilityZones"] = undefined /*out*/;
            inputs["backtrackWindow"] = undefined /*out*/;
            inputs["backupRetentionPeriod"] = undefined /*out*/;
            inputs["copyTagsToSnapshot"] = undefined /*out*/;
            inputs["dBClusterIdentifier"] = undefined /*out*/;
            inputs["dBClusterParameterGroupName"] = undefined /*out*/;
            inputs["dBSubnetGroupName"] = undefined /*out*/;
            inputs["databaseName"] = undefined /*out*/;
            inputs["deletionProtection"] = undefined /*out*/;
            inputs["enableCloudwatchLogsExports"] = undefined /*out*/;
            inputs["enableHttpEndpoint"] = undefined /*out*/;
            inputs["enableIAMDatabaseAuthentication"] = undefined /*out*/;
            inputs["endpointAddress"] = undefined /*out*/;
            inputs["endpointPort"] = undefined /*out*/;
            inputs["engine"] = undefined /*out*/;
            inputs["engineMode"] = undefined /*out*/;
            inputs["engineVersion"] = undefined /*out*/;
            inputs["globalClusterIdentifier"] = undefined /*out*/;
            inputs["kmsKeyId"] = undefined /*out*/;
            inputs["masterUserPassword"] = undefined /*out*/;
            inputs["masterUsername"] = undefined /*out*/;
            inputs["port"] = undefined /*out*/;
            inputs["preferredBackupWindow"] = undefined /*out*/;
            inputs["preferredMaintenanceWindow"] = undefined /*out*/;
            inputs["readEndpointAddress"] = undefined /*out*/;
            inputs["replicationSourceIdentifier"] = undefined /*out*/;
            inputs["restoreType"] = undefined /*out*/;
            inputs["scalingConfiguration"] = undefined /*out*/;
            inputs["snapshotIdentifier"] = undefined /*out*/;
            inputs["sourceDBClusterIdentifier"] = undefined /*out*/;
            inputs["sourceRegion"] = undefined /*out*/;
            inputs["storageEncrypted"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["useLatestRestorableTime"] = undefined /*out*/;
            inputs["vpcSecurityGroupIds"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DBCluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DBCluster resource.
 */
export interface DBClusterArgs {
    associatedRoles?: pulumi.Input<pulumi.Input<inputs.rds.DBClusterDBClusterRoleArgs>[]>;
    availabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
    backtrackWindow?: pulumi.Input<number>;
    backupRetentionPeriod?: pulumi.Input<number>;
    copyTagsToSnapshot?: pulumi.Input<boolean>;
    dBClusterIdentifier?: pulumi.Input<string>;
    dBClusterParameterGroupName?: pulumi.Input<string>;
    dBSubnetGroupName?: pulumi.Input<string>;
    databaseName?: pulumi.Input<string>;
    deletionProtection?: pulumi.Input<boolean>;
    enableCloudwatchLogsExports?: pulumi.Input<pulumi.Input<string>[]>;
    enableHttpEndpoint?: pulumi.Input<boolean>;
    enableIAMDatabaseAuthentication?: pulumi.Input<boolean>;
    endpointAddress?: pulumi.Input<string>;
    endpointPort?: pulumi.Input<string>;
    engine: pulumi.Input<string>;
    engineMode?: pulumi.Input<string>;
    engineVersion?: pulumi.Input<string>;
    globalClusterIdentifier?: pulumi.Input<string>;
    kmsKeyId?: pulumi.Input<string>;
    masterUserPassword?: pulumi.Input<string>;
    masterUsername?: pulumi.Input<string>;
    port?: pulumi.Input<number>;
    preferredBackupWindow?: pulumi.Input<string>;
    preferredMaintenanceWindow?: pulumi.Input<string>;
    readEndpointAddress?: pulumi.Input<string>;
    replicationSourceIdentifier?: pulumi.Input<string>;
    restoreType?: pulumi.Input<string>;
    scalingConfiguration?: pulumi.Input<inputs.rds.DBClusterScalingConfigurationArgs>;
    snapshotIdentifier?: pulumi.Input<string>;
    sourceDBClusterIdentifier?: pulumi.Input<string>;
    sourceRegion?: pulumi.Input<string>;
    storageEncrypted?: pulumi.Input<boolean>;
    tags?: pulumi.Input<pulumi.Input<inputs.rds.DBClusterTagArgs>[]>;
    useLatestRestorableTime?: pulumi.Input<boolean>;
    vpcSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
}