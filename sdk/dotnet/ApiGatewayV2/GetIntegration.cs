// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGatewayV2
{
    public static class GetIntegration
    {
        /// <summary>
        /// Resource Type definition for AWS::ApiGatewayV2::Integration
        /// </summary>
        public static Task<GetIntegrationResult> InvokeAsync(GetIntegrationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIntegrationResult>("aws-native:apigatewayv2:getIntegration", args ?? new GetIntegrationArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::ApiGatewayV2::Integration
        /// </summary>
        public static Output<GetIntegrationResult> Invoke(GetIntegrationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIntegrationResult>("aws-native:apigatewayv2:getIntegration", args ?? new GetIntegrationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIntegrationArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetIntegrationArgs()
        {
        }
        public static new GetIntegrationArgs Empty => new GetIntegrationArgs();
    }

    public sealed class GetIntegrationInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetIntegrationInvokeArgs()
        {
        }
        public static new GetIntegrationInvokeArgs Empty => new GetIntegrationInvokeArgs();
    }


    [OutputType]
    public sealed class GetIntegrationResult
    {
        public readonly string? ConnectionId;
        public readonly string? ConnectionType;
        public readonly string? ContentHandlingStrategy;
        public readonly string? CredentialsArn;
        public readonly string? Description;
        public readonly string? Id;
        public readonly string? IntegrationMethod;
        public readonly string? IntegrationSubtype;
        public readonly string? IntegrationType;
        public readonly string? IntegrationUri;
        public readonly string? PassthroughBehavior;
        public readonly string? PayloadFormatVersion;
        public readonly object? RequestParameters;
        public readonly object? RequestTemplates;
        public readonly object? ResponseParameters;
        public readonly string? TemplateSelectionExpression;
        public readonly int? TimeoutInMillis;
        public readonly Outputs.IntegrationTlsConfig? TlsConfig;

        [OutputConstructor]
        private GetIntegrationResult(
            string? connectionId,

            string? connectionType,

            string? contentHandlingStrategy,

            string? credentialsArn,

            string? description,

            string? id,

            string? integrationMethod,

            string? integrationSubtype,

            string? integrationType,

            string? integrationUri,

            string? passthroughBehavior,

            string? payloadFormatVersion,

            object? requestParameters,

            object? requestTemplates,

            object? responseParameters,

            string? templateSelectionExpression,

            int? timeoutInMillis,

            Outputs.IntegrationTlsConfig? tlsConfig)
        {
            ConnectionId = connectionId;
            ConnectionType = connectionType;
            ContentHandlingStrategy = contentHandlingStrategy;
            CredentialsArn = credentialsArn;
            Description = description;
            Id = id;
            IntegrationMethod = integrationMethod;
            IntegrationSubtype = integrationSubtype;
            IntegrationType = integrationType;
            IntegrationUri = integrationUri;
            PassthroughBehavior = passthroughBehavior;
            PayloadFormatVersion = payloadFormatVersion;
            RequestParameters = requestParameters;
            RequestTemplates = requestTemplates;
            ResponseParameters = responseParameters;
            TemplateSelectionExpression = templateSelectionExpression;
            TimeoutInMillis = timeoutInMillis;
            TlsConfig = tlsConfig;
        }
    }
}
