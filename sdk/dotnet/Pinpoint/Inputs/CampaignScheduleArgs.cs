// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint.Inputs
{

    public sealed class CampaignScheduleArgs : Pulumi.ResourceArgs
    {
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        [Input("eventFilter")]
        public Input<Inputs.CampaignCampaignEventFilterArgs>? EventFilter { get; set; }

        [Input("frequency")]
        public Input<string>? Frequency { get; set; }

        [Input("isLocalTime")]
        public Input<bool>? IsLocalTime { get; set; }

        [Input("quietTime")]
        public Input<Inputs.CampaignQuietTimeArgs>? QuietTime { get; set; }

        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        public CampaignScheduleArgs()
        {
        }
    }
}