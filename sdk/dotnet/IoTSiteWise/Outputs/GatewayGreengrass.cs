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
    public sealed class GatewayGreengrass
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-gateway-greengrass.html#cfn-iotsitewise-gateway-greengrass-grouparn
        /// </summary>
        public readonly string GroupArn;

        [OutputConstructor]
        private GatewayGreengrass(string GroupArn)
        {
            this.GroupArn = GroupArn;
        }
    }
}