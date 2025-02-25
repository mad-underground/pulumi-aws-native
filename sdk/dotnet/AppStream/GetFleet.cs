// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppStream
{
    public static class GetFleet
    {
        /// <summary>
        /// Resource Type definition for AWS::AppStream::Fleet
        /// </summary>
        public static Task<GetFleetResult> InvokeAsync(GetFleetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFleetResult>("aws-native:appstream:getFleet", args ?? new GetFleetArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::AppStream::Fleet
        /// </summary>
        public static Output<GetFleetResult> Invoke(GetFleetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFleetResult>("aws-native:appstream:getFleet", args ?? new GetFleetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFleetArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetFleetArgs()
        {
        }
        public static new GetFleetArgs Empty => new GetFleetArgs();
    }

    public sealed class GetFleetInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetFleetInvokeArgs()
        {
        }
        public static new GetFleetInvokeArgs Empty => new GetFleetInvokeArgs();
    }


    [OutputType]
    public sealed class GetFleetResult
    {
        public readonly Outputs.FleetComputeCapacity? ComputeCapacity;
        public readonly string? Description;
        public readonly int? DisconnectTimeoutInSeconds;
        public readonly string? DisplayName;
        public readonly Outputs.FleetDomainJoinInfo? DomainJoinInfo;
        public readonly bool? EnableDefaultInternetAccess;
        public readonly string? IamRoleArn;
        public readonly string? Id;
        public readonly int? IdleDisconnectTimeoutInSeconds;
        public readonly string? ImageArn;
        public readonly string? ImageName;
        public readonly string? InstanceType;
        public readonly int? MaxConcurrentSessions;
        public readonly int? MaxSessionsPerInstance;
        public readonly int? MaxUserDurationInSeconds;
        public readonly string? Platform;
        public readonly Outputs.FleetS3Location? SessionScriptS3Location;
        public readonly string? StreamView;
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        public readonly ImmutableArray<string> UsbDeviceFilterStrings;
        public readonly Outputs.FleetVpcConfig? VpcConfig;

        [OutputConstructor]
        private GetFleetResult(
            Outputs.FleetComputeCapacity? computeCapacity,

            string? description,

            int? disconnectTimeoutInSeconds,

            string? displayName,

            Outputs.FleetDomainJoinInfo? domainJoinInfo,

            bool? enableDefaultInternetAccess,

            string? iamRoleArn,

            string? id,

            int? idleDisconnectTimeoutInSeconds,

            string? imageArn,

            string? imageName,

            string? instanceType,

            int? maxConcurrentSessions,

            int? maxSessionsPerInstance,

            int? maxUserDurationInSeconds,

            string? platform,

            Outputs.FleetS3Location? sessionScriptS3Location,

            string? streamView,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            ImmutableArray<string> usbDeviceFilterStrings,

            Outputs.FleetVpcConfig? vpcConfig)
        {
            ComputeCapacity = computeCapacity;
            Description = description;
            DisconnectTimeoutInSeconds = disconnectTimeoutInSeconds;
            DisplayName = displayName;
            DomainJoinInfo = domainJoinInfo;
            EnableDefaultInternetAccess = enableDefaultInternetAccess;
            IamRoleArn = iamRoleArn;
            Id = id;
            IdleDisconnectTimeoutInSeconds = idleDisconnectTimeoutInSeconds;
            ImageArn = imageArn;
            ImageName = imageName;
            InstanceType = instanceType;
            MaxConcurrentSessions = maxConcurrentSessions;
            MaxSessionsPerInstance = maxSessionsPerInstance;
            MaxUserDurationInSeconds = maxUserDurationInSeconds;
            Platform = platform;
            SessionScriptS3Location = sessionScriptS3Location;
            StreamView = streamView;
            Tags = tags;
            UsbDeviceFilterStrings = usbDeviceFilterStrings;
            VpcConfig = vpcConfig;
        }
    }
}
