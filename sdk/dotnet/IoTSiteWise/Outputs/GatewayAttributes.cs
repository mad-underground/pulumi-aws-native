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
    public sealed class GatewayAttributes
    {
        public readonly string GatewayId;

        [OutputConstructor]
        private GatewayAttributes(string GatewayId)
        {
            this.GatewayId = GatewayId;
        }
    }
}
