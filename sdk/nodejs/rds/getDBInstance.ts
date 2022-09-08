// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::RDS::DBInstance resource creates an Amazon RDS DB instance.
 */
export function getDBInstance(args: GetDBInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetDBInstanceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:rds:getDBInstance", {
        "dBInstanceIdentifier": args.dBInstanceIdentifier,
    }, opts);
}

export interface GetDBInstanceArgs {
    /**
     * A name for the DB instance. If you specify a name, AWS CloudFormation converts it to lowercase. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the DB instance.
     */
    dBInstanceIdentifier: string;
}

export interface GetDBInstanceResult {
    /**
     * The amount of storage (in gigabytes) to be initially allocated for the database instance.
     */
    readonly allocatedStorage?: string;
    /**
     * A value that indicates whether major version upgrades are allowed. Changing this parameter doesn't result in an outage and the change is asynchronously applied as soon as possible.
     */
    readonly allowMajorVersionUpgrade?: boolean;
    /**
     * The AWS Identity and Access Management (IAM) roles associated with the DB instance.
     */
    readonly associatedRoles?: outputs.rds.DBInstanceRole[];
    /**
     * A value that indicates whether minor engine upgrades are applied automatically to the DB instance during the maintenance window. By default, minor engine upgrades are applied automatically.
     */
    readonly autoMinorVersionUpgrade?: boolean;
    /**
     * The Availability Zone (AZ) where the database will be created. For information on AWS Regions and Availability Zones.
     */
    readonly availabilityZone?: string;
    /**
     * The number of days for which automated backups are retained. Setting this parameter to a positive number enables backups. Setting this parameter to 0 disables automated backups.
     */
    readonly backupRetentionPeriod?: number;
    /**
     * The identifier of the CA certificate for this DB instance.
     */
    readonly cACertificateIdentifier?: string;
    /**
     * A value that indicates whether to copy tags from the DB instance to snapshots of the DB instance. By default, tags are not copied.
     */
    readonly copyTagsToSnapshot?: boolean;
    /**
     * The compute and memory capacity of the DB instance, for example, db.m4.large. Not all DB instance classes are available in all AWS Regions, or for all database engines.
     */
    readonly dBInstanceClass?: string;
    /**
     * The name of an existing DB parameter group or a reference to an AWS::RDS::DBParameterGroup resource created in the template.
     */
    readonly dBParameterGroupName?: string;
    /**
     * A list of the DB security groups to assign to the DB instance. The list can include both the name of existing DB security groups or references to AWS::RDS::DBSecurityGroup resources created in the template.
     */
    readonly dBSecurityGroups?: string[];
    /**
     * A value that indicates whether to remove automated backups immediately after the DB instance is deleted. This parameter isn't case-sensitive. The default is to remove automated backups immediately after the DB instance is deleted.
     */
    readonly deleteAutomatedBackups?: boolean;
    /**
     * A value that indicates whether the DB instance has deletion protection enabled. The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled.
     */
    readonly deletionProtection?: boolean;
    /**
     * The Active Directory directory ID to create the DB instance in. Currently, only MySQL, Microsoft SQL Server, Oracle, and PostgreSQL DB instances can be created in an Active Directory Domain.
     */
    readonly domain?: string;
    /**
     * Specify the name of the IAM role to be used when making API calls to the Directory Service.
     */
    readonly domainIAMRoleName?: string;
    /**
     * The list of log types that need to be enabled for exporting to CloudWatch Logs. The values in the list depend on the DB engine being used.
     */
    readonly enableCloudwatchLogsExports?: string[];
    /**
     * A value that indicates whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts. By default, mapping is disabled.
     */
    readonly enableIAMDatabaseAuthentication?: boolean;
    /**
     * A value that indicates whether to enable Performance Insights for the DB instance.
     */
    readonly enablePerformanceInsights?: boolean;
    /**
     * Specifies the connection endpoint.
     */
    readonly endpoint?: outputs.rds.DBInstanceEndpoint;
    /**
     * The name of the database engine that you want to use for this DB instance.
     */
    readonly engine?: string;
    /**
     * The version number of the database engine to use.
     */
    readonly engineVersion?: string;
    /**
     * The number of I/O operations per second (IOPS) that the database provisions.
     */
    readonly iops?: number;
    /**
     * License model information for this DB instance.
     */
    readonly licenseModel?: string;
    /**
     * The upper limit to which Amazon RDS can automatically scale the storage of the DB instance.
     */
    readonly maxAllocatedStorage?: number;
    /**
     * The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance. To disable collecting Enhanced Monitoring metrics, specify 0. The default is 0.
     */
    readonly monitoringInterval?: number;
    /**
     * The ARN for the IAM role that permits RDS to send enhanced monitoring metrics to Amazon CloudWatch Logs.
     */
    readonly monitoringRoleArn?: string;
    /**
     * Specifies whether the database instance is a multiple Availability Zone deployment.
     */
    readonly multiAZ?: boolean;
    /**
     * Indicates that the DB instance should be associated with the specified option group.
     */
    readonly optionGroupName?: string;
    /**
     * The AWS KMS key identifier for encryption of Performance Insights data. The KMS key ID is the Amazon Resource Name (ARN), KMS key identifier, or the KMS key alias for the KMS encryption key.
     */
    readonly performanceInsightsKMSKeyId?: string;
    /**
     * The amount of time, in days, to retain Performance Insights data. Valid values are 7 or 731 (2 years).
     */
    readonly performanceInsightsRetentionPeriod?: number;
    /**
     * The daily time range during which automated backups are created if automated backups are enabled, using the BackupRetentionPeriod parameter.
     */
    readonly preferredBackupWindow?: string;
    /**
     * he weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
     */
    readonly preferredMaintenanceWindow?: string;
    /**
     * The number of CPU cores and the number of threads per core for the DB instance class of the DB instance.
     */
    readonly processorFeatures?: outputs.rds.DBInstanceProcessorFeature[];
    /**
     * A value that specifies the order in which an Aurora Replica is promoted to the primary instance after a failure of the existing primary instance.
     */
    readonly promotionTier?: number;
    /**
     * Specifies the storage type to be associated with the DB instance.
     */
    readonly storageType?: string;
    /**
     * Tags to assign to the DB instance.
     */
    readonly tags?: outputs.rds.DBInstanceTag[];
    /**
     * The ARN from the key store with which to associate the instance for TDE encryption.
     */
    readonly tdeCredentialArn?: string;
    /**
     * A value that indicates whether the DB instance class of the DB instance uses its default processor features.
     */
    readonly useDefaultProcessorFeatures?: boolean;
    /**
     * A list of the VPC security group IDs to assign to the DB instance. The list can include both the physical IDs of existing VPC security groups and references to AWS::EC2::SecurityGroup resources created in the template.
     */
    readonly vPCSecurityGroups?: string[];
}

export function getDBInstanceOutput(args: GetDBInstanceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDBInstanceResult> {
    return pulumi.output(args).apply(a => getDBInstance(a, opts))
}

export interface GetDBInstanceOutputArgs {
    /**
     * A name for the DB instance. If you specify a name, AWS CloudFormation converts it to lowercase. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the DB instance.
     */
    dBInstanceIdentifier: pulumi.Input<string>;
}
