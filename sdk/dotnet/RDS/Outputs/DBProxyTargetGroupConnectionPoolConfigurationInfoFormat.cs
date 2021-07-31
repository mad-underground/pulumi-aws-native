// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RDS.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbproxytargetgroup-connectionpoolconfigurationinfoformat.html
    /// </summary>
    [OutputType]
    public sealed class DBProxyTargetGroupConnectionPoolConfigurationInfoFormat
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbproxytargetgroup-connectionpoolconfigurationinfoformat.html#cfn-rds-dbproxytargetgroup-connectionpoolconfigurationinfoformat-connectionborrowtimeout
        /// </summary>
        public readonly int? ConnectionBorrowTimeout;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbproxytargetgroup-connectionpoolconfigurationinfoformat.html#cfn-rds-dbproxytargetgroup-connectionpoolconfigurationinfoformat-initquery
        /// </summary>
        public readonly string? InitQuery;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbproxytargetgroup-connectionpoolconfigurationinfoformat.html#cfn-rds-dbproxytargetgroup-connectionpoolconfigurationinfoformat-maxconnectionspercent
        /// </summary>
        public readonly int? MaxConnectionsPercent;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbproxytargetgroup-connectionpoolconfigurationinfoformat.html#cfn-rds-dbproxytargetgroup-connectionpoolconfigurationinfoformat-maxidleconnectionspercent
        /// </summary>
        public readonly int? MaxIdleConnectionsPercent;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbproxytargetgroup-connectionpoolconfigurationinfoformat.html#cfn-rds-dbproxytargetgroup-connectionpoolconfigurationinfoformat-sessionpinningfilters
        /// </summary>
        public readonly ImmutableArray<string> SessionPinningFilters;

        [OutputConstructor]
        private DBProxyTargetGroupConnectionPoolConfigurationInfoFormat(
            int? connectionBorrowTimeout,

            string? initQuery,

            int? maxConnectionsPercent,

            int? maxIdleConnectionsPercent,

            ImmutableArray<string> sessionPinningFilters)
        {
            ConnectionBorrowTimeout = connectionBorrowTimeout;
            InitQuery = initQuery;
            MaxConnectionsPercent = maxConnectionsPercent;
            MaxIdleConnectionsPercent = maxIdleConnectionsPercent;
            SessionPinningFilters = sessionPinningFilters;
        }
    }
}
