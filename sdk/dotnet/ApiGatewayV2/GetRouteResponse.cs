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
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRouteResponseResult>("aws-native:apigatewayv2:getRouteResponse", args ?? new GetRouteResponseArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::ApiGatewayV2::RouteResponse
        /// </summary>
        public static Output<GetRouteResponseResult> Invoke(GetRouteResponseInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRouteResponseResult>("aws-native:apigatewayv2:getRouteResponse", args ?? new GetRouteResponseInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRouteResponseArgs : global::Pulumi.InvokeArgs
    {
        [Input("apiId", required: true)]
        public string ApiId { get; set; } = null!;

        [Input("routeId", required: true)]
        public string RouteId { get; set; } = null!;

        [Input("routeResponseId", required: true)]
        public string RouteResponseId { get; set; } = null!;

        public GetRouteResponseArgs()
        {
        }
        public static new GetRouteResponseArgs Empty => new GetRouteResponseArgs();
    }

    public sealed class GetRouteResponseInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        [Input("routeId", required: true)]
        public Input<string> RouteId { get; set; } = null!;

        [Input("routeResponseId", required: true)]
        public Input<string> RouteResponseId { get; set; } = null!;

        public GetRouteResponseInvokeArgs()
        {
        }
        public static new GetRouteResponseInvokeArgs Empty => new GetRouteResponseInvokeArgs();
    }


    [OutputType]
    public sealed class GetRouteResponseResult
    {
        public readonly string? ModelSelectionExpression;
        public readonly object? ResponseModels;
        public readonly object? ResponseParameters;
        public readonly string? RouteResponseId;
        public readonly string? RouteResponseKey;

        [OutputConstructor]
        private GetRouteResponseResult(
            string? modelSelectionExpression,

            object? responseModels,

            object? responseParameters,

            string? routeResponseId,

            string? routeResponseKey)
        {
            ModelSelectionExpression = modelSelectionExpression;
            ResponseModels = responseModels;
            ResponseParameters = responseParameters;
            RouteResponseId = routeResponseId;
            RouteResponseKey = routeResponseKey;
        }
    }
}
