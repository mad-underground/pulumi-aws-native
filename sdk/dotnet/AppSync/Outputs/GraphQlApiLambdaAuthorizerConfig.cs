// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync.Outputs
{

    [OutputType]
    public sealed class GraphQlApiLambdaAuthorizerConfig
    {
        public readonly double? AuthorizerResultTtlInSeconds;
        public readonly string? AuthorizerUri;
        public readonly string? IdentityValidationExpression;

        [OutputConstructor]
        private GraphQlApiLambdaAuthorizerConfig(
            double? authorizerResultTtlInSeconds,

            string? authorizerUri,

            string? identityValidationExpression)
        {
            AuthorizerResultTtlInSeconds = authorizerResultTtlInSeconds;
            AuthorizerUri = authorizerUri;
            IdentityValidationExpression = identityValidationExpression;
        }
    }
}
