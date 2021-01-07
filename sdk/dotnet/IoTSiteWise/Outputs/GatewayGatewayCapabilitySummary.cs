// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.IoTSiteWise.Outputs
{

    [OutputType]
    public sealed class GatewayGatewayCapabilitySummary
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-gateway-gatewaycapabilitysummary.html#cfn-iotsitewise-gateway-gatewaycapabilitysummary-capabilityconfiguration
        /// </summary>
        public readonly string? CapabilityConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-gateway-gatewaycapabilitysummary.html#cfn-iotsitewise-gateway-gatewaycapabilitysummary-capabilitynamespace
        /// </summary>
        public readonly string CapabilityNamespace;

        [OutputConstructor]
        private GatewayGatewayCapabilitySummary(
            string? CapabilityConfiguration,

            string CapabilityNamespace)
        {
            this.CapabilityConfiguration = CapabilityConfiguration;
            this.CapabilityNamespace = CapabilityNamespace;
        }
    }
}
