// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IVSChat
{
    public static class GetLoggingConfiguration
    {
        /// <summary>
        /// Resource type definition for AWS::IVSChat::LoggingConfiguration.
        /// </summary>
        public static Task<GetLoggingConfigurationResult> InvokeAsync(GetLoggingConfigurationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLoggingConfigurationResult>("aws-native:ivschat:getLoggingConfiguration", args ?? new GetLoggingConfigurationArgs(), options.WithDefaults());

        /// <summary>
        /// Resource type definition for AWS::IVSChat::LoggingConfiguration.
        /// </summary>
        public static Output<GetLoggingConfigurationResult> Invoke(GetLoggingConfigurationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLoggingConfigurationResult>("aws-native:ivschat:getLoggingConfiguration", args ?? new GetLoggingConfigurationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLoggingConfigurationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// LoggingConfiguration ARN is automatically generated on creation and assigned as the unique identifier.
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetLoggingConfigurationArgs()
        {
        }
        public static new GetLoggingConfigurationArgs Empty => new GetLoggingConfigurationArgs();
    }

    public sealed class GetLoggingConfigurationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// LoggingConfiguration ARN is automatically generated on creation and assigned as the unique identifier.
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetLoggingConfigurationInvokeArgs()
        {
        }
        public static new GetLoggingConfigurationInvokeArgs Empty => new GetLoggingConfigurationInvokeArgs();
    }


    [OutputType]
    public sealed class GetLoggingConfigurationResult
    {
        /// <summary>
        /// LoggingConfiguration ARN is automatically generated on creation and assigned as the unique identifier.
        /// </summary>
        public readonly string? Arn;
        public readonly Outputs.LoggingConfigurationDestinationConfiguration? DestinationConfiguration;
        /// <summary>
        /// The system-generated ID of the logging configuration.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The name of the logging configuration. The value does not need to be unique.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The state of the logging configuration. When the state is ACTIVE, the configuration is ready to log chat content.
        /// </summary>
        public readonly Pulumi.AwsNative.IVSChat.LoggingConfigurationState? State;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.LoggingConfigurationTag> Tags;

        [OutputConstructor]
        private GetLoggingConfigurationResult(
            string? arn,

            Outputs.LoggingConfigurationDestinationConfiguration? destinationConfiguration,

            string? id,

            string? name,

            Pulumi.AwsNative.IVSChat.LoggingConfigurationState? state,

            ImmutableArray<Outputs.LoggingConfigurationTag> tags)
        {
            Arn = arn;
            DestinationConfiguration = destinationConfiguration;
            Id = id;
            Name = name;
            State = state;
            Tags = tags;
        }
    }
}
