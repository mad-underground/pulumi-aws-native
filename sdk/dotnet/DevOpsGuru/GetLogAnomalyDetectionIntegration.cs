// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DevOpsGuru
{
    public static class GetLogAnomalyDetectionIntegration
    {
        /// <summary>
        /// This resource schema represents the LogAnomalyDetectionIntegration resource in the Amazon DevOps Guru.
        /// </summary>
        public static Task<GetLogAnomalyDetectionIntegrationResult> InvokeAsync(GetLogAnomalyDetectionIntegrationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLogAnomalyDetectionIntegrationResult>("aws-native:devopsguru:getLogAnomalyDetectionIntegration", args ?? new GetLogAnomalyDetectionIntegrationArgs(), options.WithDefaults());

        /// <summary>
        /// This resource schema represents the LogAnomalyDetectionIntegration resource in the Amazon DevOps Guru.
        /// </summary>
        public static Output<GetLogAnomalyDetectionIntegrationResult> Invoke(GetLogAnomalyDetectionIntegrationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLogAnomalyDetectionIntegrationResult>("aws-native:devopsguru:getLogAnomalyDetectionIntegration", args ?? new GetLogAnomalyDetectionIntegrationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLogAnomalyDetectionIntegrationArgs : global::Pulumi.InvokeArgs
    {
        [Input("accountId", required: true)]
        public string AccountId { get; set; } = null!;

        public GetLogAnomalyDetectionIntegrationArgs()
        {
        }
        public static new GetLogAnomalyDetectionIntegrationArgs Empty => new GetLogAnomalyDetectionIntegrationArgs();
    }

    public sealed class GetLogAnomalyDetectionIntegrationInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        public GetLogAnomalyDetectionIntegrationInvokeArgs()
        {
        }
        public static new GetLogAnomalyDetectionIntegrationInvokeArgs Empty => new GetLogAnomalyDetectionIntegrationInvokeArgs();
    }


    [OutputType]
    public sealed class GetLogAnomalyDetectionIntegrationResult
    {
        public readonly string? AccountId;

        [OutputConstructor]
        private GetLogAnomalyDetectionIntegrationResult(string? accountId)
        {
            AccountId = accountId;
        }
    }
}