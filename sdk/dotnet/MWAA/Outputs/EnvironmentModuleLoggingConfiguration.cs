// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.MWAA.Outputs
{

    [OutputType]
    public sealed class EnvironmentModuleLoggingConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaa-environment-moduleloggingconfiguration.html#cfn-mwaa-environment-moduleloggingconfiguration-cloudwatchloggrouparn
        /// </summary>
        public readonly string? CloudWatchLogGroupArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaa-environment-moduleloggingconfiguration.html#cfn-mwaa-environment-moduleloggingconfiguration-enabled
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaa-environment-moduleloggingconfiguration.html#cfn-mwaa-environment-moduleloggingconfiguration-loglevel
        /// </summary>
        public readonly string? LogLevel;

        [OutputConstructor]
        private EnvironmentModuleLoggingConfiguration(
            string? CloudWatchLogGroupArn,

            bool? Enabled,

            string? LogLevel)
        {
            this.CloudWatchLogGroupArn = CloudWatchLogGroupArn;
            this.Enabled = Enabled;
            this.LogLevel = LogLevel;
        }
    }
}
