// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGatewayV2.Outputs
{

    [OutputType]
    public sealed class ApiGatewayManagedOverridesIntegrationOverrides
    {
        public readonly string? Description;
        public readonly string? IntegrationMethod;
        public readonly string? PayloadFormatVersion;
        public readonly int? TimeoutInMillis;

        [OutputConstructor]
        private ApiGatewayManagedOverridesIntegrationOverrides(
            string? description,

            string? integrationMethod,

            string? payloadFormatVersion,

            int? timeoutInMillis)
        {
            Description = description;
            IntegrationMethod = integrationMethod;
            PayloadFormatVersion = payloadFormatVersion;
            TimeoutInMillis = timeoutInMillis;
        }
    }
}
