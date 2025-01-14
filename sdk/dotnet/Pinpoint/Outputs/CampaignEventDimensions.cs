// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint.Outputs
{

    [OutputType]
    public sealed class CampaignEventDimensions
    {
        public readonly object? Attributes;
        public readonly Outputs.CampaignSetDimension? EventType;
        public readonly object? Metrics;

        [OutputConstructor]
        private CampaignEventDimensions(
            object? attributes,

            Outputs.CampaignSetDimension? eventType,

            object? metrics)
        {
            Attributes = attributes;
            EventType = eventType;
            Metrics = metrics;
        }
    }
}
