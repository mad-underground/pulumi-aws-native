// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GroundStation.Outputs
{

    [OutputType]
    public sealed class DataflowEndpointGroupEndpointDetails
    {
        public readonly Outputs.DataflowEndpointGroupAwsGroundStationAgentEndpoint? AwsGroundStationAgentEndpoint;
        public readonly Outputs.DataflowEndpointGroupDataflowEndpoint? Endpoint;
        public readonly Outputs.DataflowEndpointGroupSecurityDetails? SecurityDetails;

        [OutputConstructor]
        private DataflowEndpointGroupEndpointDetails(
            Outputs.DataflowEndpointGroupAwsGroundStationAgentEndpoint? awsGroundStationAgentEndpoint,

            Outputs.DataflowEndpointGroupDataflowEndpoint? endpoint,

            Outputs.DataflowEndpointGroupSecurityDetails? securityDetails)
        {
            AwsGroundStationAgentEndpoint = awsGroundStationAgentEndpoint;
            Endpoint = endpoint;
            SecurityDetails = securityDetails;
        }
    }
}
