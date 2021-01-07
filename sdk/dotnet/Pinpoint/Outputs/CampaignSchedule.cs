// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Pinpoint.Outputs
{

    [OutputType]
    public sealed class CampaignSchedule
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-schedule.html#cfn-pinpoint-campaign-schedule-endtime
        /// </summary>
        public readonly string? EndTime;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-schedule.html#cfn-pinpoint-campaign-schedule-eventfilter
        /// </summary>
        public readonly Outputs.CampaignCampaignEventFilter? EventFilter;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-schedule.html#cfn-pinpoint-campaign-schedule-frequency
        /// </summary>
        public readonly string? Frequency;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-schedule.html#cfn-pinpoint-campaign-schedule-islocaltime
        /// </summary>
        public readonly bool? IsLocalTime;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-schedule.html#cfn-pinpoint-campaign-schedule-quiettime
        /// </summary>
        public readonly Outputs.CampaignQuietTime? QuietTime;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-schedule.html#cfn-pinpoint-campaign-schedule-starttime
        /// </summary>
        public readonly string? StartTime;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-schedule.html#cfn-pinpoint-campaign-schedule-timezone
        /// </summary>
        public readonly string? TimeZone;

        [OutputConstructor]
        private CampaignSchedule(
            string? EndTime,

            Outputs.CampaignCampaignEventFilter? EventFilter,

            string? Frequency,

            bool? IsLocalTime,

            Outputs.CampaignQuietTime? QuietTime,

            string? StartTime,

            string? TimeZone)
        {
            this.EndTime = EndTime;
            this.EventFilter = EventFilter;
            this.Frequency = Frequency;
            this.IsLocalTime = IsLocalTime;
            this.QuietTime = QuietTime;
            this.StartTime = StartTime;
            this.TimeZone = TimeZone;
        }
    }
}