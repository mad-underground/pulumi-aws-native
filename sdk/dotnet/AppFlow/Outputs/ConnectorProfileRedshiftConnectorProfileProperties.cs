// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Outputs
{

    [OutputType]
    public sealed class ConnectorProfileRedshiftConnectorProfileProperties
    {
        /// <summary>
        /// The name of the Amazon S3 bucket associated with Redshift.
        /// </summary>
        public readonly string BucketName;
        /// <summary>
        /// The object key for the destination bucket in which Amazon AppFlow will place the ﬁles.
        /// </summary>
        public readonly string? BucketPrefix;
        /// <summary>
        /// The unique identifier of the Amazon Redshift cluster.
        /// </summary>
        public readonly string? ClusterIdentifier;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM role that grants Amazon AppFlow access to the data through the Amazon Redshift Data API.
        /// </summary>
        public readonly string? DataApiRoleArn;
        /// <summary>
        /// The name of the Amazon Redshift database that will store the transferred data.
        /// </summary>
        public readonly string? DatabaseName;
        /// <summary>
        /// The JDBC URL of the Amazon Redshift cluster.
        /// </summary>
        public readonly string? DatabaseUrl;
        /// <summary>
        /// If Amazon AppFlow will connect to Amazon Redshift Serverless or Amazon Redshift cluster.
        /// </summary>
        public readonly bool? IsRedshiftServerless;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM role.
        /// </summary>
        public readonly string RoleArn;
        /// <summary>
        /// The name of the Amazon Redshift serverless workgroup
        /// </summary>
        public readonly string? WorkgroupName;

        [OutputConstructor]
        private ConnectorProfileRedshiftConnectorProfileProperties(
            string bucketName,

            string? bucketPrefix,

            string? clusterIdentifier,

            string? dataApiRoleArn,

            string? databaseName,

            string? databaseUrl,

            bool? isRedshiftServerless,

            string roleArn,

            string? workgroupName)
        {
            BucketName = bucketName;
            BucketPrefix = bucketPrefix;
            ClusterIdentifier = clusterIdentifier;
            DataApiRoleArn = dataApiRoleArn;
            DatabaseName = databaseName;
            DatabaseUrl = databaseUrl;
            IsRedshiftServerless = isRedshiftServerless;
            RoleArn = roleArn;
            WorkgroupName = workgroupName;
        }
    }
}
