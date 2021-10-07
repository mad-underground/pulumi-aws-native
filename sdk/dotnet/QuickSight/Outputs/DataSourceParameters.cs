// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    /// <summary>
    /// &lt;p&gt;The parameters that Amazon QuickSight uses to connect to your underlying data source.
    ///             This is a variant type structure. For this structure to be valid, only one of the
    ///             attributes can be non-null.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class DataSourceParameters
    {
        public readonly Outputs.DataSourceAmazonElasticsearchParameters? AmazonElasticsearchParameters;
        public readonly Outputs.DataSourceAthenaParameters? AthenaParameters;
        public readonly Outputs.DataSourceAuroraParameters? AuroraParameters;
        public readonly Outputs.DataSourceAuroraPostgreSqlParameters? AuroraPostgreSqlParameters;
        public readonly Outputs.DataSourceMariaDbParameters? MariaDbParameters;
        public readonly Outputs.DataSourceMySqlParameters? MySqlParameters;
        public readonly Outputs.DataSourceOracleParameters? OracleParameters;
        public readonly Outputs.DataSourcePostgreSqlParameters? PostgreSqlParameters;
        public readonly Outputs.DataSourcePrestoParameters? PrestoParameters;
        public readonly Outputs.DataSourceRdsParameters? RdsParameters;
        public readonly Outputs.DataSourceRedshiftParameters? RedshiftParameters;
        public readonly Outputs.DataSourceS3Parameters? S3Parameters;
        public readonly Outputs.DataSourceSnowflakeParameters? SnowflakeParameters;
        public readonly Outputs.DataSourceSparkParameters? SparkParameters;
        public readonly Outputs.DataSourceSqlServerParameters? SqlServerParameters;
        public readonly Outputs.DataSourceTeradataParameters? TeradataParameters;

        [OutputConstructor]
        private DataSourceParameters(
            Outputs.DataSourceAmazonElasticsearchParameters? amazonElasticsearchParameters,

            Outputs.DataSourceAthenaParameters? athenaParameters,

            Outputs.DataSourceAuroraParameters? auroraParameters,

            Outputs.DataSourceAuroraPostgreSqlParameters? auroraPostgreSqlParameters,

            Outputs.DataSourceMariaDbParameters? mariaDbParameters,

            Outputs.DataSourceMySqlParameters? mySqlParameters,

            Outputs.DataSourceOracleParameters? oracleParameters,

            Outputs.DataSourcePostgreSqlParameters? postgreSqlParameters,

            Outputs.DataSourcePrestoParameters? prestoParameters,

            Outputs.DataSourceRdsParameters? rdsParameters,

            Outputs.DataSourceRedshiftParameters? redshiftParameters,

            Outputs.DataSourceS3Parameters? s3Parameters,

            Outputs.DataSourceSnowflakeParameters? snowflakeParameters,

            Outputs.DataSourceSparkParameters? sparkParameters,

            Outputs.DataSourceSqlServerParameters? sqlServerParameters,

            Outputs.DataSourceTeradataParameters? teradataParameters)
        {
            AmazonElasticsearchParameters = amazonElasticsearchParameters;
            AthenaParameters = athenaParameters;
            AuroraParameters = auroraParameters;
            AuroraPostgreSqlParameters = auroraPostgreSqlParameters;
            MariaDbParameters = mariaDbParameters;
            MySqlParameters = mySqlParameters;
            OracleParameters = oracleParameters;
            PostgreSqlParameters = postgreSqlParameters;
            PrestoParameters = prestoParameters;
            RdsParameters = rdsParameters;
            RedshiftParameters = redshiftParameters;
            S3Parameters = s3Parameters;
            SnowflakeParameters = snowflakeParameters;
            SparkParameters = sparkParameters;
            SqlServerParameters = sqlServerParameters;
            TeradataParameters = teradataParameters;
        }
    }
}