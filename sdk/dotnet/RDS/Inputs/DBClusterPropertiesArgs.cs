// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.RDS.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html
    /// </summary>
    public sealed class DBClusterPropertiesArgs : Pulumi.ResourceArgs
    {
        [Input("AssociatedRoles")]
        private InputList<Inputs.DBClusterDBClusterRoleArgs>? _AssociatedRoles;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-associatedroles
        /// </summary>
        public InputList<Inputs.DBClusterDBClusterRoleArgs> AssociatedRoles
        {
            get => _AssociatedRoles ?? (_AssociatedRoles = new InputList<Inputs.DBClusterDBClusterRoleArgs>());
            set => _AssociatedRoles = value;
        }

        [Input("AvailabilityZones")]
        private InputList<string>? _AvailabilityZones;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-availabilityzones
        /// </summary>
        public InputList<string> AvailabilityZones
        {
            get => _AvailabilityZones ?? (_AvailabilityZones = new InputList<string>());
            set => _AvailabilityZones = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-backtrackwindow
        /// </summary>
        [Input("BacktrackWindow")]
        public Input<int>? BacktrackWindow { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-backuprententionperiod
        /// </summary>
        [Input("BackupRetentionPeriod")]
        public Input<int>? BackupRetentionPeriod { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-dbclusteridentifier
        /// </summary>
        [Input("DBClusterIdentifier")]
        public Input<string>? DBClusterIdentifier { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-dbclusterparametergroupname
        /// </summary>
        [Input("DBClusterParameterGroupName")]
        public Input<string>? DBClusterParameterGroupName { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-dbsubnetgroupname
        /// </summary>
        [Input("DBSubnetGroupName")]
        public Input<string>? DBSubnetGroupName { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-databasename
        /// </summary>
        [Input("DatabaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-deletionprotection
        /// </summary>
        [Input("DeletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        [Input("EnableCloudwatchLogsExports")]
        private InputList<string>? _EnableCloudwatchLogsExports;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-enablecloudwatchlogsexports
        /// </summary>
        public InputList<string> EnableCloudwatchLogsExports
        {
            get => _EnableCloudwatchLogsExports ?? (_EnableCloudwatchLogsExports = new InputList<string>());
            set => _EnableCloudwatchLogsExports = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-enablehttpendpoint
        /// </summary>
        [Input("EnableHttpEndpoint")]
        public Input<bool>? EnableHttpEndpoint { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-enableiamdatabaseauthentication
        /// </summary>
        [Input("EnableIAMDatabaseAuthentication")]
        public Input<bool>? EnableIAMDatabaseAuthentication { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-engine
        /// </summary>
        [Input("Engine", required: true)]
        public Input<string> Engine { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-enginemode
        /// </summary>
        [Input("EngineMode")]
        public Input<string>? EngineMode { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-engineversion
        /// </summary>
        [Input("EngineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-globalclusteridentifier
        /// </summary>
        [Input("GlobalClusterIdentifier")]
        public Input<string>? GlobalClusterIdentifier { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-kmskeyid
        /// </summary>
        [Input("KmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-masteruserpassword
        /// </summary>
        [Input("MasterUserPassword")]
        public Input<string>? MasterUserPassword { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-masterusername
        /// </summary>
        [Input("MasterUsername")]
        public Input<string>? MasterUsername { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-port
        /// </summary>
        [Input("Port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-preferredbackupwindow
        /// </summary>
        [Input("PreferredBackupWindow")]
        public Input<string>? PreferredBackupWindow { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-preferredmaintenancewindow
        /// </summary>
        [Input("PreferredMaintenanceWindow")]
        public Input<string>? PreferredMaintenanceWindow { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-replicationsourceidentifier
        /// </summary>
        [Input("ReplicationSourceIdentifier")]
        public Input<string>? ReplicationSourceIdentifier { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-restoretype
        /// </summary>
        [Input("RestoreType")]
        public Input<string>? RestoreType { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-scalingconfiguration
        /// </summary>
        [Input("ScalingConfiguration")]
        public Input<Inputs.DBClusterScalingConfigurationArgs>? ScalingConfiguration { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-snapshotidentifier
        /// </summary>
        [Input("SnapshotIdentifier")]
        public Input<string>? SnapshotIdentifier { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-sourcedbclusteridentifier
        /// </summary>
        [Input("SourceDBClusterIdentifier")]
        public Input<string>? SourceDBClusterIdentifier { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-sourceregion
        /// </summary>
        [Input("SourceRegion")]
        public Input<string>? SourceRegion { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-storageencrypted
        /// </summary>
        [Input("StorageEncrypted")]
        public Input<bool>? StorageEncrypted { get; set; }

        [Input("Tags")]
        private InputList<Pulumi.Cloudformation.Inputs.TagArgs>? _Tags;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-tags
        /// </summary>
        public InputList<Pulumi.Cloudformation.Inputs.TagArgs> Tags
        {
            get => _Tags ?? (_Tags = new InputList<Pulumi.Cloudformation.Inputs.TagArgs>());
            set => _Tags = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-uselatestrestorabletime
        /// </summary>
        [Input("UseLatestRestorableTime")]
        public Input<bool>? UseLatestRestorableTime { get; set; }

        [Input("VpcSecurityGroupIds")]
        private InputList<string>? _VpcSecurityGroupIds;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#cfn-rds-dbcluster-vpcsecuritygroupids
        /// </summary>
        public InputList<string> VpcSecurityGroupIds
        {
            get => _VpcSecurityGroupIds ?? (_VpcSecurityGroupIds = new InputList<string>());
            set => _VpcSecurityGroupIds = value;
        }

        public DBClusterPropertiesArgs()
        {
        }
    }
}