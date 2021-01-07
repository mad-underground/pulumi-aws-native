// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.MWAA.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaa-environment-moduleloggingconfiguration.html
    /// </summary>
    public sealed class EnvironmentModuleLoggingConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaa-environment-moduleloggingconfiguration.html#cfn-mwaa-environment-moduleloggingconfiguration-cloudwatchloggrouparn
        /// </summary>
        [Input("CloudWatchLogGroupArn")]
        public Input<string>? CloudWatchLogGroupArn { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaa-environment-moduleloggingconfiguration.html#cfn-mwaa-environment-moduleloggingconfiguration-enabled
        /// </summary>
        [Input("Enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaa-environment-moduleloggingconfiguration.html#cfn-mwaa-environment-moduleloggingconfiguration-loglevel
        /// </summary>
        [Input("LogLevel")]
        public Input<string>? LogLevel { get; set; }

        public EnvironmentModuleLoggingConfigurationArgs()
        {
        }
    }
}
