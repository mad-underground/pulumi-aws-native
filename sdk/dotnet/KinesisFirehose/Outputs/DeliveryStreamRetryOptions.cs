// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.KinesisFirehose.Outputs
{

    [OutputType]
    public sealed class DeliveryStreamRetryOptions
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-retryoptions.html#cfn-kinesisfirehose-deliverystream-retryoptions-durationinseconds
        /// </summary>
        public readonly int? DurationInSeconds;

        [OutputConstructor]
        private DeliveryStreamRetryOptions(int? DurationInSeconds)
        {
            this.DurationInSeconds = DurationInSeconds;
        }
    }
}