// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.AppMesh.Outputs
{

    [OutputType]
    public sealed class GatewayRouteGrpcGatewayRouteMatch
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-grpcgatewayroutematch.html#cfn-appmesh-gatewayroute-grpcgatewayroutematch-servicename
        /// </summary>
        public readonly string? ServiceName;

        [OutputConstructor]
        private GatewayRouteGrpcGatewayRouteMatch(string? ServiceName)
        {
            this.ServiceName = ServiceName;
        }
    }
}
