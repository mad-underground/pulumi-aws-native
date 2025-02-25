// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync
{
    public static class GetGraphQlApi
    {
        /// <summary>
        /// Resource Type definition for AWS::AppSync::GraphQLApi
        /// </summary>
        public static Task<GetGraphQlApiResult> InvokeAsync(GetGraphQlApiArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGraphQlApiResult>("aws-native:appsync:getGraphQlApi", args ?? new GetGraphQlApiArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::AppSync::GraphQLApi
        /// </summary>
        public static Output<GetGraphQlApiResult> Invoke(GetGraphQlApiInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGraphQlApiResult>("aws-native:appsync:getGraphQlApi", args ?? new GetGraphQlApiInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGraphQlApiArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetGraphQlApiArgs()
        {
        }
        public static new GetGraphQlApiArgs Empty => new GetGraphQlApiArgs();
    }

    public sealed class GetGraphQlApiInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetGraphQlApiInvokeArgs()
        {
        }
        public static new GetGraphQlApiInvokeArgs Empty => new GetGraphQlApiInvokeArgs();
    }


    [OutputType]
    public sealed class GetGraphQlApiResult
    {
        public readonly ImmutableArray<Outputs.GraphQlApiAdditionalAuthenticationProvider> AdditionalAuthenticationProviders;
        public readonly string? ApiId;
        public readonly string? ApiType;
        public readonly string? Arn;
        public readonly string? AuthenticationType;
        public readonly Outputs.GraphQlApiEnhancedMetricsConfig? EnhancedMetricsConfig;
        /// <summary>
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::AppSync::GraphQLApi` for more information about the expected schema for this property.
        /// </summary>
        public readonly object? EnvironmentVariables;
        public readonly string? GraphQlDns;
        public readonly string? GraphQlEndpointArn;
        public readonly string? GraphQlUrl;
        public readonly string? Id;
        public readonly string? IntrospectionConfig;
        public readonly Outputs.GraphQlApiLambdaAuthorizerConfig? LambdaAuthorizerConfig;
        public readonly Outputs.GraphQlApiLogConfig? LogConfig;
        public readonly string? MergedApiExecutionRoleArn;
        public readonly string? Name;
        public readonly Outputs.GraphQlApiOpenIdConnectConfig? OpenIdConnectConfig;
        public readonly string? OwnerContact;
        public readonly int? QueryDepthLimit;
        public readonly string? RealtimeDns;
        public readonly string? RealtimeUrl;
        public readonly int? ResolverCountLimit;
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        public readonly Outputs.GraphQlApiUserPoolConfig? UserPoolConfig;
        public readonly string? Visibility;
        public readonly bool? XrayEnabled;

        [OutputConstructor]
        private GetGraphQlApiResult(
            ImmutableArray<Outputs.GraphQlApiAdditionalAuthenticationProvider> additionalAuthenticationProviders,

            string? apiId,

            string? apiType,

            string? arn,

            string? authenticationType,

            Outputs.GraphQlApiEnhancedMetricsConfig? enhancedMetricsConfig,

            object? environmentVariables,

            string? graphQlDns,

            string? graphQlEndpointArn,

            string? graphQlUrl,

            string? id,

            string? introspectionConfig,

            Outputs.GraphQlApiLambdaAuthorizerConfig? lambdaAuthorizerConfig,

            Outputs.GraphQlApiLogConfig? logConfig,

            string? mergedApiExecutionRoleArn,

            string? name,

            Outputs.GraphQlApiOpenIdConnectConfig? openIdConnectConfig,

            string? ownerContact,

            int? queryDepthLimit,

            string? realtimeDns,

            string? realtimeUrl,

            int? resolverCountLimit,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            Outputs.GraphQlApiUserPoolConfig? userPoolConfig,

            string? visibility,

            bool? xrayEnabled)
        {
            AdditionalAuthenticationProviders = additionalAuthenticationProviders;
            ApiId = apiId;
            ApiType = apiType;
            Arn = arn;
            AuthenticationType = authenticationType;
            EnhancedMetricsConfig = enhancedMetricsConfig;
            EnvironmentVariables = environmentVariables;
            GraphQlDns = graphQlDns;
            GraphQlEndpointArn = graphQlEndpointArn;
            GraphQlUrl = graphQlUrl;
            Id = id;
            IntrospectionConfig = introspectionConfig;
            LambdaAuthorizerConfig = lambdaAuthorizerConfig;
            LogConfig = logConfig;
            MergedApiExecutionRoleArn = mergedApiExecutionRoleArn;
            Name = name;
            OpenIdConnectConfig = openIdConnectConfig;
            OwnerContact = ownerContact;
            QueryDepthLimit = queryDepthLimit;
            RealtimeDns = realtimeDns;
            RealtimeUrl = realtimeUrl;
            ResolverCountLimit = resolverCountLimit;
            Tags = tags;
            UserPoolConfig = userPoolConfig;
            Visibility = visibility;
            XrayEnabled = xrayEnabled;
        }
    }
}
