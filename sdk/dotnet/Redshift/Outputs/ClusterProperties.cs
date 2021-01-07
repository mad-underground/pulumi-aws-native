// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Redshift.Outputs
{

    [OutputType]
    public sealed class ClusterProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-allowversionupgrade
        /// </summary>
        public readonly bool? AllowVersionUpgrade;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-automatedsnapshotretentionperiod
        /// </summary>
        public readonly int? AutomatedSnapshotRetentionPeriod;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-availabilityzone
        /// </summary>
        public readonly string? AvailabilityZone;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-clusteridentifier
        /// </summary>
        public readonly string? ClusterIdentifier;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-clusterparametergroupname
        /// </summary>
        public readonly string? ClusterParameterGroupName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-clustersecuritygroups
        /// </summary>
        public readonly ImmutableArray<string> ClusterSecurityGroups;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-clustersubnetgroupname
        /// </summary>
        public readonly string? ClusterSubnetGroupName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-clustertype
        /// </summary>
        public readonly string ClusterType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-clusterversion
        /// </summary>
        public readonly string? ClusterVersion;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-dbname
        /// </summary>
        public readonly string DBName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-elasticip
        /// </summary>
        public readonly string? ElasticIp;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-encrypted
        /// </summary>
        public readonly bool? Encrypted;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-hsmclientcertidentifier
        /// </summary>
        public readonly string? HsmClientCertificateIdentifier;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-HsmConfigurationIdentifier
        /// </summary>
        public readonly string? HsmConfigurationIdentifier;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-iamroles
        /// </summary>
        public readonly ImmutableArray<string> IamRoles;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-kmskeyid
        /// </summary>
        public readonly string? KmsKeyId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-loggingproperties
        /// </summary>
        public readonly Outputs.ClusterLoggingProperties? LoggingProperties;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-masteruserpassword
        /// </summary>
        public readonly string MasterUserPassword;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-masterusername
        /// </summary>
        public readonly string MasterUsername;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-nodetype
        /// </summary>
        public readonly string NodeType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-nodetype
        /// </summary>
        public readonly int? NumberOfNodes;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-owneraccount
        /// </summary>
        public readonly string? OwnerAccount;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-port
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-preferredmaintenancewindow
        /// </summary>
        public readonly string? PreferredMaintenanceWindow;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-publiclyaccessible
        /// </summary>
        public readonly bool? PubliclyAccessible;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-snapshotclusteridentifier
        /// </summary>
        public readonly string? SnapshotClusterIdentifier;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-snapshotidentifier
        /// </summary>
        public readonly string? SnapshotIdentifier;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-tags
        /// </summary>
        public readonly ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html#cfn-redshift-cluster-vpcsecuritygroupids
        /// </summary>
        public readonly ImmutableArray<string> VpcSecurityGroupIds;

        [OutputConstructor]
        private ClusterProperties(
            bool? AllowVersionUpgrade,

            int? AutomatedSnapshotRetentionPeriod,

            string? AvailabilityZone,

            string? ClusterIdentifier,

            string? ClusterParameterGroupName,

            ImmutableArray<string> ClusterSecurityGroups,

            string? ClusterSubnetGroupName,

            string ClusterType,

            string? ClusterVersion,

            string DBName,

            string? ElasticIp,

            bool? Encrypted,

            string? HsmClientCertificateIdentifier,

            string? HsmConfigurationIdentifier,

            ImmutableArray<string> IamRoles,

            string? KmsKeyId,

            Outputs.ClusterLoggingProperties? LoggingProperties,

            string MasterUserPassword,

            string MasterUsername,

            string NodeType,

            int? NumberOfNodes,

            string? OwnerAccount,

            int? Port,

            string? PreferredMaintenanceWindow,

            bool? PubliclyAccessible,

            string? SnapshotClusterIdentifier,

            string? SnapshotIdentifier,

            ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags,

            ImmutableArray<string> VpcSecurityGroupIds)
        {
            this.AllowVersionUpgrade = AllowVersionUpgrade;
            this.AutomatedSnapshotRetentionPeriod = AutomatedSnapshotRetentionPeriod;
            this.AvailabilityZone = AvailabilityZone;
            this.ClusterIdentifier = ClusterIdentifier;
            this.ClusterParameterGroupName = ClusterParameterGroupName;
            this.ClusterSecurityGroups = ClusterSecurityGroups;
            this.ClusterSubnetGroupName = ClusterSubnetGroupName;
            this.ClusterType = ClusterType;
            this.ClusterVersion = ClusterVersion;
            this.DBName = DBName;
            this.ElasticIp = ElasticIp;
            this.Encrypted = Encrypted;
            this.HsmClientCertificateIdentifier = HsmClientCertificateIdentifier;
            this.HsmConfigurationIdentifier = HsmConfigurationIdentifier;
            this.IamRoles = IamRoles;
            this.KmsKeyId = KmsKeyId;
            this.LoggingProperties = LoggingProperties;
            this.MasterUserPassword = MasterUserPassword;
            this.MasterUsername = MasterUsername;
            this.NodeType = NodeType;
            this.NumberOfNodes = NumberOfNodes;
            this.OwnerAccount = OwnerAccount;
            this.Port = Port;
            this.PreferredMaintenanceWindow = PreferredMaintenanceWindow;
            this.PubliclyAccessible = PubliclyAccessible;
            this.SnapshotClusterIdentifier = SnapshotClusterIdentifier;
            this.SnapshotIdentifier = SnapshotIdentifier;
            this.Tags = Tags;
            this.VpcSecurityGroupIds = VpcSecurityGroupIds;
        }
    }
}