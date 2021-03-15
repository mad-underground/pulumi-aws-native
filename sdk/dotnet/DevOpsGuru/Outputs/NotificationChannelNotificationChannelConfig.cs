// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DevOpsGuru.Outputs
{

    [OutputType]
    public sealed class NotificationChannelNotificationChannelConfig
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsguru-notificationchannel-notificationchannelconfig.html#cfn-devopsguru-notificationchannel-notificationchannelconfig-sns
        /// </summary>
        public readonly Outputs.NotificationChannelSnsChannelConfig? Sns;

        [OutputConstructor]
        private NotificationChannelNotificationChannelConfig(Outputs.NotificationChannelSnsChannelConfig? sns)
        {
            Sns = sns;
        }
    }
}
