// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.NetworkFirewall.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-loggingconfiguration-logdestinationconfig.html
    /// </summary>
    public sealed class LoggingConfigurationLogDestinationConfigArgs : Pulumi.ResourceArgs
    {
        [Input("LogDestination", required: true)]
        private InputMap<string>? _LogDestination;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-loggingconfiguration-logdestinationconfig.html#cfn-networkfirewall-loggingconfiguration-logdestinationconfig-logdestination
        /// </summary>
        public InputMap<string> LogDestination
        {
            get => _LogDestination ?? (_LogDestination = new InputMap<string>());
            set => _LogDestination = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-loggingconfiguration-logdestinationconfig.html#cfn-networkfirewall-loggingconfiguration-logdestinationconfig-logdestinationtype
        /// </summary>
        [Input("LogDestinationType", required: true)]
        public Input<string> LogDestinationType { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-loggingconfiguration-logdestinationconfig.html#cfn-networkfirewall-loggingconfiguration-logdestinationconfig-logtype
        /// </summary>
        [Input("LogType", required: true)]
        public Input<string> LogType { get; set; } = null!;

        public LoggingConfigurationLogDestinationConfigArgs()
        {
        }
    }
}
