// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-quickconnectconfig.html
    /// </summary>
    [OutputType]
    public sealed class QuickConnectQuickConnectConfig
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-quickconnectconfig.html#cfn-connect-quickconnect-quickconnectconfig-phoneconfig
        /// </summary>
        public readonly Outputs.QuickConnectPhoneNumberQuickConnectConfig? PhoneConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-quickconnectconfig.html#cfn-connect-quickconnect-quickconnectconfig-queueconfig
        /// </summary>
        public readonly Outputs.QuickConnectQueueQuickConnectConfig? QueueConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-quickconnectconfig.html#cfn-connect-quickconnect-quickconnectconfig-quickconnecttype
        /// </summary>
        public readonly string QuickConnectType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-quickconnectconfig.html#cfn-connect-quickconnect-quickconnectconfig-userconfig
        /// </summary>
        public readonly Outputs.QuickConnectUserQuickConnectConfig? UserConfig;

        [OutputConstructor]
        private QuickConnectQuickConnectConfig(
            Outputs.QuickConnectPhoneNumberQuickConnectConfig? phoneConfig,

            Outputs.QuickConnectQueueQuickConnectConfig? queueConfig,

            string quickConnectType,

            Outputs.QuickConnectUserQuickConnectConfig? userConfig)
        {
            PhoneConfig = phoneConfig;
            QueueConfig = queueConfig;
            QuickConnectType = quickConnectType;
            UserConfig = userConfig;
        }
    }
}