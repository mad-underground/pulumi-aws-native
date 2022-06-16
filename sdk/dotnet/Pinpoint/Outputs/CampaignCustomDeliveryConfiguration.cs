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
    public sealed class CampaignCustomDeliveryConfiguration
    {
        public readonly string? DeliveryUri;
        public readonly ImmutableArray<string> EndpointTypes;

        [OutputConstructor]
        private CampaignCustomDeliveryConfiguration(
            string? deliveryUri,

            ImmutableArray<string> endpointTypes)
        {
            DeliveryUri = deliveryUri;
            EndpointTypes = endpointTypes;
        }
    }
}