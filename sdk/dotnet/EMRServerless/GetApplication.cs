// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EMRServerless
{
    public static class GetApplication
    {
        /// <summary>
        /// Resource schema for AWS::EMRServerless::Application Type
        /// </summary>
        public static Task<GetApplicationResult> InvokeAsync(GetApplicationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApplicationResult>("aws-native:emrserverless:getApplication", args ?? new GetApplicationArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::EMRServerless::Application Type
        /// </summary>
        public static Output<GetApplicationResult> Invoke(GetApplicationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApplicationResult>("aws-native:emrserverless:getApplication", args ?? new GetApplicationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApplicationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the EMR Serverless Application.
        /// </summary>
        [Input("applicationId", required: true)]
        public string ApplicationId { get; set; } = null!;

        public GetApplicationArgs()
        {
        }
        public static new GetApplicationArgs Empty => new GetApplicationArgs();
    }

    public sealed class GetApplicationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the EMR Serverless Application.
        /// </summary>
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        public GetApplicationInvokeArgs()
        {
        }
        public static new GetApplicationInvokeArgs Empty => new GetApplicationInvokeArgs();
    }


    [OutputType]
    public sealed class GetApplicationResult
    {
        /// <summary>
        /// The ID of the EMR Serverless Application.
        /// </summary>
        public readonly string? ApplicationId;
        public readonly Pulumi.AwsNative.EMRServerless.ApplicationArchitecture? Architecture;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the EMR Serverless Application.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// Configuration for Auto Start of Application.
        /// </summary>
        public readonly Outputs.ApplicationAutoStartConfiguration? AutoStartConfiguration;
        /// <summary>
        /// Configuration for Auto Stop of Application.
        /// </summary>
        public readonly Outputs.ApplicationAutoStopConfiguration? AutoStopConfiguration;
        /// <summary>
        /// Initial capacity initialized when an Application is started.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationInitialCapacityConfigKeyValuePair> InitialCapacity;
        /// <summary>
        /// Maximum allowed cumulative resources for an Application. No new resources will be created once the limit is hit.
        /// </summary>
        public readonly Outputs.ApplicationMaximumAllowedResources? MaximumCapacity;
        /// <summary>
        /// Network Configuration for customer VPC connectivity.
        /// </summary>
        public readonly Outputs.ApplicationNetworkConfiguration? NetworkConfiguration;
        /// <summary>
        /// Tag map with key and value
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationTag> Tags;

        [OutputConstructor]
        private GetApplicationResult(
            string? applicationId,

            Pulumi.AwsNative.EMRServerless.ApplicationArchitecture? architecture,

            string? arn,

            Outputs.ApplicationAutoStartConfiguration? autoStartConfiguration,

            Outputs.ApplicationAutoStopConfiguration? autoStopConfiguration,

            ImmutableArray<Outputs.ApplicationInitialCapacityConfigKeyValuePair> initialCapacity,

            Outputs.ApplicationMaximumAllowedResources? maximumCapacity,

            Outputs.ApplicationNetworkConfiguration? networkConfiguration,

            ImmutableArray<Outputs.ApplicationTag> tags)
        {
            ApplicationId = applicationId;
            Architecture = architecture;
            Arn = arn;
            AutoStartConfiguration = autoStartConfiguration;
            AutoStopConfiguration = autoStopConfiguration;
            InitialCapacity = initialCapacity;
            MaximumCapacity = maximumCapacity;
            NetworkConfiguration = networkConfiguration;
            Tags = tags;
        }
    }
}
