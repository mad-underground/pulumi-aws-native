// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.MSK.Outputs
{

    [OutputType]
    public sealed class ClusterBrokerLogs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-brokerlogs.html#cfn-msk-cluster-brokerlogs-cloudwatchlogs
        /// </summary>
        public readonly Outputs.ClusterCloudWatchLogs? CloudWatchLogs;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-brokerlogs.html#cfn-msk-cluster-brokerlogs-firehose
        /// </summary>
        public readonly Outputs.ClusterFirehose? Firehose;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-brokerlogs.html#cfn-msk-cluster-brokerlogs-s3
        /// </summary>
        public readonly Outputs.ClusterS3? S3;

        [OutputConstructor]
        private ClusterBrokerLogs(
            Outputs.ClusterCloudWatchLogs? CloudWatchLogs,

            Outputs.ClusterFirehose? Firehose,

            Outputs.ClusterS3? S3)
        {
            this.CloudWatchLogs = CloudWatchLogs;
            this.Firehose = Firehose;
            this.S3 = S3;
        }
    }
}