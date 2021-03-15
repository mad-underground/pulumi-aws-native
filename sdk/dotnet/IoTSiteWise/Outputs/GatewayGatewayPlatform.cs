// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTSiteWise.Outputs
{

    [OutputType]
    public sealed class GatewayGatewayPlatform
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-gateway-gatewayplatform.html#cfn-iotsitewise-gateway-gatewayplatform-greengrass
        /// </summary>
        public readonly Outputs.GatewayGreengrass Greengrass;

        [OutputConstructor]
        private GatewayGatewayPlatform(Outputs.GatewayGreengrass greengrass)
        {
            Greengrass = greengrass;
        }
    }
}
