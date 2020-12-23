// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.RDS.Outputs
{

    [OutputType]
    public sealed class GlobalClusterProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-globalcluster.html#cfn-rds-globalcluster-deletionprotection
        /// </summary>
        public readonly bool? DeletionProtection;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-globalcluster.html#cfn-rds-globalcluster-engine
        /// </summary>
        public readonly string? Engine;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-globalcluster.html#cfn-rds-globalcluster-engineversion
        /// </summary>
        public readonly string? EngineVersion;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-globalcluster.html#cfn-rds-globalcluster-globalclusteridentifier
        /// </summary>
        public readonly string? GlobalClusterIdentifier;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-globalcluster.html#cfn-rds-globalcluster-sourcedbclusteridentifier
        /// </summary>
        public readonly string? SourceDBClusterIdentifier;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-globalcluster.html#cfn-rds-globalcluster-storageencrypted
        /// </summary>
        public readonly bool? StorageEncrypted;

        [OutputConstructor]
        private GlobalClusterProperties(
            bool? DeletionProtection,

            string? Engine,

            string? EngineVersion,

            string? GlobalClusterIdentifier,

            string? SourceDBClusterIdentifier,

            bool? StorageEncrypted)
        {
            this.DeletionProtection = DeletionProtection;
            this.Engine = Engine;
            this.EngineVersion = EngineVersion;
            this.GlobalClusterIdentifier = GlobalClusterIdentifier;
            this.SourceDBClusterIdentifier = SourceDBClusterIdentifier;
            this.StorageEncrypted = StorageEncrypted;
        }
    }
}
