// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    /// <summary>
    /// The configuration object that specifies the monitoring schedule and defines the monitoring job.
    /// </summary>
    public sealed class MonitoringScheduleConfigArgs : Pulumi.ResourceArgs
    {
        [Input("monitoringJobDefinition")]
        public Input<Inputs.MonitoringScheduleMonitoringJobDefinitionArgs>? MonitoringJobDefinition { get; set; }

        /// <summary>
        /// Name of the job definition
        /// </summary>
        [Input("monitoringJobDefinitionName")]
        public Input<string>? MonitoringJobDefinitionName { get; set; }

        [Input("monitoringType")]
        public Input<Pulumi.AwsNative.SageMaker.MonitoringScheduleMonitoringType>? MonitoringType { get; set; }

        [Input("scheduleConfig")]
        public Input<Inputs.MonitoringScheduleScheduleConfigArgs>? ScheduleConfig { get; set; }

        public MonitoringScheduleConfigArgs()
        {
        }
    }
}