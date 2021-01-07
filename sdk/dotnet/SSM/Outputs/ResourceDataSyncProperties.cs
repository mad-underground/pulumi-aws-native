// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.SSM.Outputs
{

    [OutputType]
    public sealed class ResourceDataSyncProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-bucketname
        /// </summary>
        public readonly string? BucketName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-bucketprefix
        /// </summary>
        public readonly string? BucketPrefix;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-bucketregion
        /// </summary>
        public readonly string? BucketRegion;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-kmskeyarn
        /// </summary>
        public readonly string? KMSKeyArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-s3destination
        /// </summary>
        public readonly Outputs.ResourceDataSyncS3Destination? S3Destination;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-syncformat
        /// </summary>
        public readonly string? SyncFormat;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-syncname
        /// </summary>
        public readonly string SyncName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-syncsource
        /// </summary>
        public readonly Outputs.ResourceDataSyncSyncSource? SyncSource;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-synctype
        /// </summary>
        public readonly string? SyncType;

        [OutputConstructor]
        private ResourceDataSyncProperties(
            string? BucketName,

            string? BucketPrefix,

            string? BucketRegion,

            string? KMSKeyArn,

            Outputs.ResourceDataSyncS3Destination? S3Destination,

            string? SyncFormat,

            string SyncName,

            Outputs.ResourceDataSyncSyncSource? SyncSource,

            string? SyncType)
        {
            this.BucketName = BucketName;
            this.BucketPrefix = BucketPrefix;
            this.BucketRegion = BucketRegion;
            this.KMSKeyArn = KMSKeyArn;
            this.S3Destination = S3Destination;
            this.SyncFormat = SyncFormat;
            this.SyncName = SyncName;
            this.SyncSource = SyncSource;
            this.SyncType = SyncType;
        }
    }
}