// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppMesh.Outputs
{

    [OutputType]
    public sealed class GatewayRouteGrpcGatewayRoute
    {
        public readonly Outputs.GatewayRouteGrpcGatewayRouteAction Action;
        public readonly Outputs.GatewayRouteGrpcGatewayRouteMatch Match;

        [OutputConstructor]
        private GatewayRouteGrpcGatewayRoute(
            Outputs.GatewayRouteGrpcGatewayRouteAction action,

            Outputs.GatewayRouteGrpcGatewayRouteMatch match)
        {
            Action = action;
            Match = match;
        }
    }
}
