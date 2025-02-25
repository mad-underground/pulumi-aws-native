// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cognito
{
    public static class GetLogDeliveryConfiguration
    {
        /// <summary>
        /// Resource Type definition for AWS::Cognito::LogDeliveryConfiguration
        /// </summary>
        public static Task<GetLogDeliveryConfigurationResult> InvokeAsync(GetLogDeliveryConfigurationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLogDeliveryConfigurationResult>("aws-native:cognito:getLogDeliveryConfiguration", args ?? new GetLogDeliveryConfigurationArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Cognito::LogDeliveryConfiguration
        /// </summary>
        public static Output<GetLogDeliveryConfigurationResult> Invoke(GetLogDeliveryConfigurationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLogDeliveryConfigurationResult>("aws-native:cognito:getLogDeliveryConfiguration", args ?? new GetLogDeliveryConfigurationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLogDeliveryConfigurationArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetLogDeliveryConfigurationArgs()
        {
        }
        public static new GetLogDeliveryConfigurationArgs Empty => new GetLogDeliveryConfigurationArgs();
    }

    public sealed class GetLogDeliveryConfigurationInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetLogDeliveryConfigurationInvokeArgs()
        {
        }
        public static new GetLogDeliveryConfigurationInvokeArgs Empty => new GetLogDeliveryConfigurationInvokeArgs();
    }


    [OutputType]
    public sealed class GetLogDeliveryConfigurationResult
    {
        public readonly string? Id;
        public readonly ImmutableArray<Outputs.LogDeliveryConfigurationLogConfiguration> LogConfigurations;

        [OutputConstructor]
        private GetLogDeliveryConfigurationResult(
            string? id,

            ImmutableArray<Outputs.LogDeliveryConfigurationLogConfiguration> logConfigurations)
        {
            Id = id;
            LogConfigurations = logConfigurations;
        }
    }
}
