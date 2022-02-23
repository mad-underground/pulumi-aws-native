// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGatewayV2
{
    public static class GetRouteResponse
    {
        /// <summary>
        /// Resource Type definition for AWS::ApiGatewayV2::RouteResponse
        /// </summary>
        public static Task<GetRouteResponseResult> InvokeAsync(GetRouteResponseArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRouteResponseResult>("aws-native:apigatewayv2:getRouteResponse", args ?? new GetRouteResponseArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::ApiGatewayV2::RouteResponse
        /// </summary>
        public static Output<GetRouteResponseResult> Invoke(GetRouteResponseInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRouteResponseResult>("aws-native:apigatewayv2:getRouteResponse", args ?? new GetRouteResponseInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRouteResponseArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetRouteResponseArgs()
        {
        }
    }

    public sealed class GetRouteResponseInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetRouteResponseInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRouteResponseResult
    {
        public readonly string? Id;
        public readonly string? ModelSelectionExpression;
        public readonly object? ResponseModels;
        public readonly object? ResponseParameters;
        public readonly string? RouteResponseKey;

        [OutputConstructor]
        private GetRouteResponseResult(
            string? id,

            string? modelSelectionExpression,

            object? responseModels,

            object? responseParameters,

            string? routeResponseKey)
        {
            Id = id;
            ModelSelectionExpression = modelSelectionExpression;
            ResponseModels = responseModels;
            ResponseParameters = responseParameters;
            RouteResponseKey = routeResponseKey;
        }
    }
}