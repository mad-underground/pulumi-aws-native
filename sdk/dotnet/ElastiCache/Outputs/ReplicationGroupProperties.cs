// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.ElastiCache.Outputs
{

    [OutputType]
    public sealed class ReplicationGroupProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-atrestencryptionenabled
        /// </summary>
        public readonly bool? AtRestEncryptionEnabled;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-authtoken
        /// </summary>
        public readonly string? AuthToken;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-autominorversionupgrade
        /// </summary>
        public readonly bool? AutoMinorVersionUpgrade;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-automaticfailoverenabled
        /// </summary>
        public readonly bool? AutomaticFailoverEnabled;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-cachenodetype
        /// </summary>
        public readonly string? CacheNodeType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-cacheparametergroupname
        /// </summary>
        public readonly string? CacheParameterGroupName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-cachesecuritygroupnames
        /// </summary>
        public readonly ImmutableArray<string> CacheSecurityGroupNames;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-cachesubnetgroupname
        /// </summary>
        public readonly string? CacheSubnetGroupName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-engine
        /// </summary>
        public readonly string? Engine;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-engineversion
        /// </summary>
        public readonly string? EngineVersion;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-globalreplicationgroupid
        /// </summary>
        public readonly string? GlobalReplicationGroupId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-kmskeyid
        /// </summary>
        public readonly string? KmsKeyId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-multiazenabled
        /// </summary>
        public readonly bool? MultiAZEnabled;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-nodegroupconfiguration
        /// </summary>
        public readonly ImmutableArray<Outputs.ReplicationGroupNodeGroupConfiguration> NodeGroupConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-notificationtopicarn
        /// </summary>
        public readonly string? NotificationTopicArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-numcacheclusters
        /// </summary>
        public readonly int? NumCacheClusters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-numnodegroups
        /// </summary>
        public readonly int? NumNodeGroups;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-port
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-preferredcacheclusterazs
        /// </summary>
        public readonly ImmutableArray<string> PreferredCacheClusterAZs;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-preferredmaintenancewindow
        /// </summary>
        public readonly string? PreferredMaintenanceWindow;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-primaryclusterid
        /// </summary>
        public readonly string? PrimaryClusterId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-replicaspernodegroup
        /// </summary>
        public readonly int? ReplicasPerNodeGroup;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-replicationgroupdescription
        /// </summary>
        public readonly string ReplicationGroupDescription;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-replicationgroupid
        /// </summary>
        public readonly string? ReplicationGroupId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-securitygroupids
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-snapshotarns
        /// </summary>
        public readonly ImmutableArray<string> SnapshotArns;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-snapshotname
        /// </summary>
        public readonly string? SnapshotName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-snapshotretentionlimit
        /// </summary>
        public readonly int? SnapshotRetentionLimit;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-snapshotwindow
        /// </summary>
        public readonly string? SnapshotWindow;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-snapshottingclusterid
        /// </summary>
        public readonly string? SnapshottingClusterId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-tags
        /// </summary>
        public readonly ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-transitencryptionenabled
        /// </summary>
        public readonly bool? TransitEncryptionEnabled;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-usergroupids
        /// </summary>
        public readonly ImmutableArray<string> UserGroupIds;

        [OutputConstructor]
        private ReplicationGroupProperties(
            bool? AtRestEncryptionEnabled,

            string? AuthToken,

            bool? AutoMinorVersionUpgrade,

            bool? AutomaticFailoverEnabled,

            string? CacheNodeType,

            string? CacheParameterGroupName,

            ImmutableArray<string> CacheSecurityGroupNames,

            string? CacheSubnetGroupName,

            string? Engine,

            string? EngineVersion,

            string? GlobalReplicationGroupId,

            string? KmsKeyId,

            bool? MultiAZEnabled,

            ImmutableArray<Outputs.ReplicationGroupNodeGroupConfiguration> NodeGroupConfiguration,

            string? NotificationTopicArn,

            int? NumCacheClusters,

            int? NumNodeGroups,

            int? Port,

            ImmutableArray<string> PreferredCacheClusterAZs,

            string? PreferredMaintenanceWindow,

            string? PrimaryClusterId,

            int? ReplicasPerNodeGroup,

            string ReplicationGroupDescription,

            string? ReplicationGroupId,

            ImmutableArray<string> SecurityGroupIds,

            ImmutableArray<string> SnapshotArns,

            string? SnapshotName,

            int? SnapshotRetentionLimit,

            string? SnapshotWindow,

            string? SnapshottingClusterId,

            ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags,

            bool? TransitEncryptionEnabled,

            ImmutableArray<string> UserGroupIds)
        {
            this.AtRestEncryptionEnabled = AtRestEncryptionEnabled;
            this.AuthToken = AuthToken;
            this.AutoMinorVersionUpgrade = AutoMinorVersionUpgrade;
            this.AutomaticFailoverEnabled = AutomaticFailoverEnabled;
            this.CacheNodeType = CacheNodeType;
            this.CacheParameterGroupName = CacheParameterGroupName;
            this.CacheSecurityGroupNames = CacheSecurityGroupNames;
            this.CacheSubnetGroupName = CacheSubnetGroupName;
            this.Engine = Engine;
            this.EngineVersion = EngineVersion;
            this.GlobalReplicationGroupId = GlobalReplicationGroupId;
            this.KmsKeyId = KmsKeyId;
            this.MultiAZEnabled = MultiAZEnabled;
            this.NodeGroupConfiguration = NodeGroupConfiguration;
            this.NotificationTopicArn = NotificationTopicArn;
            this.NumCacheClusters = NumCacheClusters;
            this.NumNodeGroups = NumNodeGroups;
            this.Port = Port;
            this.PreferredCacheClusterAZs = PreferredCacheClusterAZs;
            this.PreferredMaintenanceWindow = PreferredMaintenanceWindow;
            this.PrimaryClusterId = PrimaryClusterId;
            this.ReplicasPerNodeGroup = ReplicasPerNodeGroup;
            this.ReplicationGroupDescription = ReplicationGroupDescription;
            this.ReplicationGroupId = ReplicationGroupId;
            this.SecurityGroupIds = SecurityGroupIds;
            this.SnapshotArns = SnapshotArns;
            this.SnapshotName = SnapshotName;
            this.SnapshotRetentionLimit = SnapshotRetentionLimit;
            this.SnapshotWindow = SnapshotWindow;
            this.SnapshottingClusterId = SnapshottingClusterId;
            this.Tags = Tags;
            this.TransitEncryptionEnabled = TransitEncryptionEnabled;
            this.UserGroupIds = UserGroupIds;
        }
    }
}
