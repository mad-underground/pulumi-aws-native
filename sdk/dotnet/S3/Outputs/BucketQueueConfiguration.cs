// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.S3.Outputs
{

    [OutputType]
    public sealed class BucketQueueConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-queueconfig.html#cfn-s3-bucket-notificationconfig-queueconfig-event
        /// </summary>
        public readonly string Event;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-queueconfig.html#cfn-s3-bucket-notificationconfig-queueconfig-filter
        /// </summary>
        public readonly Outputs.BucketNotificationFilter? Filter;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-queueconfig.html#cfn-s3-bucket-notificationconfig-queueconfig-queue
        /// </summary>
        public readonly string Queue;

        [OutputConstructor]
        private BucketQueueConfiguration(
            string Event,

            Outputs.BucketNotificationFilter? Filter,

            string Queue)
        {
            this.Event = Event;
            this.Filter = Filter;
            this.Queue = Queue;
        }
    }
}