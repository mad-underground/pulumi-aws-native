// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SES
{
    public static class GetConfigurationSet
    {
        /// <summary>
        /// Resource schema for AWS::SES::ConfigurationSet.
        /// </summary>
        public static Task<GetConfigurationSetResult> InvokeAsync(GetConfigurationSetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetConfigurationSetResult>("aws-native:ses:getConfigurationSet", args ?? new GetConfigurationSetArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::SES::ConfigurationSet.
        /// </summary>
        public static Output<GetConfigurationSetResult> Invoke(GetConfigurationSetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetConfigurationSetResult>("aws-native:ses:getConfigurationSet", args ?? new GetConfigurationSetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConfigurationSetArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the configuration set.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetConfigurationSetArgs()
        {
        }
        public static new GetConfigurationSetArgs Empty => new GetConfigurationSetArgs();
    }

    public sealed class GetConfigurationSetInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the configuration set.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetConfigurationSetInvokeArgs()
        {
        }
        public static new GetConfigurationSetInvokeArgs Empty => new GetConfigurationSetInvokeArgs();
    }


    [OutputType]
    public sealed class GetConfigurationSetResult
    {
        public readonly Outputs.ConfigurationSetDeliveryOptions? DeliveryOptions;
        public readonly Outputs.ConfigurationSetReputationOptions? ReputationOptions;
        public readonly Outputs.ConfigurationSetSendingOptions? SendingOptions;
        public readonly Outputs.ConfigurationSetSuppressionOptions? SuppressionOptions;
        public readonly Outputs.ConfigurationSetTrackingOptions? TrackingOptions;

        [OutputConstructor]
        private GetConfigurationSetResult(
            Outputs.ConfigurationSetDeliveryOptions? deliveryOptions,

            Outputs.ConfigurationSetReputationOptions? reputationOptions,

            Outputs.ConfigurationSetSendingOptions? sendingOptions,

            Outputs.ConfigurationSetSuppressionOptions? suppressionOptions,

            Outputs.ConfigurationSetTrackingOptions? trackingOptions)
        {
            DeliveryOptions = deliveryOptions;
            ReputationOptions = reputationOptions;
            SendingOptions = sendingOptions;
            SuppressionOptions = suppressionOptions;
            TrackingOptions = trackingOptions;
        }
    }
}
