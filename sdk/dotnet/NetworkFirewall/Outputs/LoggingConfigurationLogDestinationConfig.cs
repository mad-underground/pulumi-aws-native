// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.NetworkFirewall.Outputs
{

    [OutputType]
    public sealed class LoggingConfigurationLogDestinationConfig
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-loggingconfiguration-logdestinationconfig.html#cfn-networkfirewall-loggingconfiguration-logdestinationconfig-logdestination
        /// </summary>
        public readonly ImmutableDictionary<string, string> LogDestination;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-loggingconfiguration-logdestinationconfig.html#cfn-networkfirewall-loggingconfiguration-logdestinationconfig-logdestinationtype
        /// </summary>
        public readonly string LogDestinationType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-loggingconfiguration-logdestinationconfig.html#cfn-networkfirewall-loggingconfiguration-logdestinationconfig-logtype
        /// </summary>
        public readonly string LogType;

        [OutputConstructor]
        private LoggingConfigurationLogDestinationConfig(
            ImmutableDictionary<string, string> LogDestination,

            string LogDestinationType,

            string LogType)
        {
            this.LogDestination = LogDestination;
            this.LogDestinationType = LogDestinationType;
            this.LogType = LogType;
        }
    }
}
