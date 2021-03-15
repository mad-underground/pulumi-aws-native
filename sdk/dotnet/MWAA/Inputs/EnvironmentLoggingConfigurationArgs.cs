// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MWAA.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaa-environment-loggingconfiguration.html
    /// </summary>
    public sealed class EnvironmentLoggingConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaa-environment-loggingconfiguration.html#cfn-mwaa-environment-loggingconfiguration-dagprocessinglogs
        /// </summary>
        [Input("dagProcessingLogs")]
        public Input<Inputs.EnvironmentModuleLoggingConfigurationArgs>? DagProcessingLogs { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaa-environment-loggingconfiguration.html#cfn-mwaa-environment-loggingconfiguration-schedulerlogs
        /// </summary>
        [Input("schedulerLogs")]
        public Input<Inputs.EnvironmentModuleLoggingConfigurationArgs>? SchedulerLogs { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaa-environment-loggingconfiguration.html#cfn-mwaa-environment-loggingconfiguration-tasklogs
        /// </summary>
        [Input("taskLogs")]
        public Input<Inputs.EnvironmentModuleLoggingConfigurationArgs>? TaskLogs { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaa-environment-loggingconfiguration.html#cfn-mwaa-environment-loggingconfiguration-webserverlogs
        /// </summary>
        [Input("webserverLogs")]
        public Input<Inputs.EnvironmentModuleLoggingConfigurationArgs>? WebserverLogs { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaa-environment-loggingconfiguration.html#cfn-mwaa-environment-loggingconfiguration-workerlogs
        /// </summary>
        [Input("workerLogs")]
        public Input<Inputs.EnvironmentModuleLoggingConfigurationArgs>? WorkerLogs { get; set; }

        public EnvironmentLoggingConfigurationArgs()
        {
        }
    }
}
