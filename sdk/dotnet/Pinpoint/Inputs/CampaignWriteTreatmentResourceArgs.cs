// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint.Inputs
{

    public sealed class CampaignWriteTreatmentResourceArgs : global::Pulumi.ResourceArgs
    {
        [Input("customDeliveryConfiguration")]
        public Input<Inputs.CampaignCustomDeliveryConfigurationArgs>? CustomDeliveryConfiguration { get; set; }

        [Input("messageConfiguration")]
        public Input<Inputs.CampaignMessageConfigurationArgs>? MessageConfiguration { get; set; }

        [Input("schedule")]
        public Input<Inputs.CampaignScheduleArgs>? Schedule { get; set; }

        [Input("sizePercent")]
        public Input<int>? SizePercent { get; set; }

        [Input("templateConfiguration")]
        public Input<Inputs.CampaignTemplateConfigurationArgs>? TemplateConfiguration { get; set; }

        [Input("treatmentDescription")]
        public Input<string>? TreatmentDescription { get; set; }

        [Input("treatmentName")]
        public Input<string>? TreatmentName { get; set; }

        public CampaignWriteTreatmentResourceArgs()
        {
        }
        public static new CampaignWriteTreatmentResourceArgs Empty => new CampaignWriteTreatmentResourceArgs();
    }
}
