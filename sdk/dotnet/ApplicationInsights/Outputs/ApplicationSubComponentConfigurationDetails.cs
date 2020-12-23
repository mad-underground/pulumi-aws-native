// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.ApplicationInsights.Outputs
{

    [OutputType]
    public sealed class ApplicationSubComponentConfigurationDetails
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-subcomponentconfigurationdetails.html#cfn-applicationinsights-application-subcomponentconfigurationdetails-alarmmetrics
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationAlarmMetric> AlarmMetrics;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-subcomponentconfigurationdetails.html#cfn-applicationinsights-application-subcomponentconfigurationdetails-logs
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationLog> Logs;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-subcomponentconfigurationdetails.html#cfn-applicationinsights-application-subcomponentconfigurationdetails-windowsevents
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationWindowsEvent> WindowsEvents;

        [OutputConstructor]
        private ApplicationSubComponentConfigurationDetails(
            ImmutableArray<Outputs.ApplicationAlarmMetric> AlarmMetrics,

            ImmutableArray<Outputs.ApplicationLog> Logs,

            ImmutableArray<Outputs.ApplicationWindowsEvent> WindowsEvents)
        {
            this.AlarmMetrics = AlarmMetrics;
            this.Logs = Logs;
            this.WindowsEvents = WindowsEvents;
        }
    }
}
