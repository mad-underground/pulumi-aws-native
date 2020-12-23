// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.AppFlow.Outputs
{

    [OutputType]
    public sealed class FlowScheduledTriggerProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-scheduledtriggerproperties.html#cfn-appflow-flow-scheduledtriggerproperties-datapullmode
        /// </summary>
        public readonly string? DataPullMode;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-scheduledtriggerproperties.html#cfn-appflow-flow-scheduledtriggerproperties-scheduleendtime
        /// </summary>
        public readonly double? ScheduleEndTime;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-scheduledtriggerproperties.html#cfn-appflow-flow-scheduledtriggerproperties-scheduleexpression
        /// </summary>
        public readonly string ScheduleExpression;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-scheduledtriggerproperties.html#cfn-appflow-flow-scheduledtriggerproperties-schedulestarttime
        /// </summary>
        public readonly double? ScheduleStartTime;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-scheduledtriggerproperties.html#cfn-appflow-flow-scheduledtriggerproperties-timezone
        /// </summary>
        public readonly string? TimeZone;

        [OutputConstructor]
        private FlowScheduledTriggerProperties(
            string? DataPullMode,

            double? ScheduleEndTime,

            string ScheduleExpression,

            double? ScheduleStartTime,

            string? TimeZone)
        {
            this.DataPullMode = DataPullMode;
            this.ScheduleEndTime = ScheduleEndTime;
            this.ScheduleExpression = ScheduleExpression;
            this.ScheduleStartTime = ScheduleStartTime;
            this.TimeZone = TimeZone;
        }
    }
}
