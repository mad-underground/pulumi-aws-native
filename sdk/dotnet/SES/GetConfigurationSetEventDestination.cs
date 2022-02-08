// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SES
{
    public static class GetConfigurationSetEventDestination
    {
        /// <summary>
        /// Resource Type definition for AWS::SES::ConfigurationSetEventDestination
        /// </summary>
        public static Task<GetConfigurationSetEventDestinationResult> InvokeAsync(GetConfigurationSetEventDestinationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConfigurationSetEventDestinationResult>("aws-native:ses:getConfigurationSetEventDestination", args ?? new GetConfigurationSetEventDestinationArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SES::ConfigurationSetEventDestination
        /// </summary>
        public static Output<GetConfigurationSetEventDestinationResult> Invoke(GetConfigurationSetEventDestinationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetConfigurationSetEventDestinationResult>("aws-native:ses:getConfigurationSetEventDestination", args ?? new GetConfigurationSetEventDestinationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConfigurationSetEventDestinationArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetConfigurationSetEventDestinationArgs()
        {
        }
    }

    public sealed class GetConfigurationSetEventDestinationInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetConfigurationSetEventDestinationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetConfigurationSetEventDestinationResult
    {
        public readonly Outputs.ConfigurationSetEventDestinationEventDestination? EventDestination;
        public readonly string? Id;

        [OutputConstructor]
        private GetConfigurationSetEventDestinationResult(
            Outputs.ConfigurationSetEventDestinationEventDestination? eventDestination,

            string? id)
        {
            EventDestination = eventDestination;
            Id = id;
        }
    }
}
