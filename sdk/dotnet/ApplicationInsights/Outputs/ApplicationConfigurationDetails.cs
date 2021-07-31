// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApplicationInsights.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-configurationdetails.html
    /// </summary>
    [OutputType]
    public sealed class ApplicationConfigurationDetails
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-configurationdetails.html#cfn-applicationinsights-application-configurationdetails-alarmmetrics
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationAlarmMetric> AlarmMetrics;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-configurationdetails.html#cfn-applicationinsights-application-configurationdetails-alarms
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationAlarm> Alarms;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-configurationdetails.html#cfn-applicationinsights-application-configurationdetails-jmxprometheusexporter
        /// </summary>
        public readonly Outputs.ApplicationJMXPrometheusExporter? JMXPrometheusExporter;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-configurationdetails.html#cfn-applicationinsights-application-configurationdetails-logs
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationLog> Logs;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-configurationdetails.html#cfn-applicationinsights-application-configurationdetails-windowsevents
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationWindowsEvent> WindowsEvents;

        [OutputConstructor]
        private ApplicationConfigurationDetails(
            ImmutableArray<Outputs.ApplicationAlarmMetric> alarmMetrics,

            ImmutableArray<Outputs.ApplicationAlarm> alarms,

            Outputs.ApplicationJMXPrometheusExporter? jMXPrometheusExporter,

            ImmutableArray<Outputs.ApplicationLog> logs,

            ImmutableArray<Outputs.ApplicationWindowsEvent> windowsEvents)
        {
            AlarmMetrics = alarmMetrics;
            Alarms = alarms;
            JMXPrometheusExporter = jMXPrometheusExporter;
            Logs = logs;
            WindowsEvents = windowsEvents;
        }
    }
}
