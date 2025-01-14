// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Route53.Outputs
{

    /// <summary>
    /// A complex type that contains information about the health check.
    /// </summary>
    [OutputType]
    public sealed class HealthCheckConfigProperties
    {
        public readonly Outputs.HealthCheckAlarmIdentifier? AlarmIdentifier;
        public readonly ImmutableArray<string> ChildHealthChecks;
        public readonly bool? EnableSni;
        public readonly int? FailureThreshold;
        public readonly string? FullyQualifiedDomainName;
        public readonly int? HealthThreshold;
        public readonly Pulumi.AwsNative.Route53.HealthCheckConfigPropertiesInsufficientDataHealthStatus? InsufficientDataHealthStatus;
        public readonly bool? Inverted;
        public readonly string? IpAddress;
        public readonly bool? MeasureLatency;
        public readonly int? Port;
        public readonly ImmutableArray<string> Regions;
        public readonly int? RequestInterval;
        public readonly string? ResourcePath;
        public readonly string? RoutingControlArn;
        public readonly string? SearchString;
        public readonly Pulumi.AwsNative.Route53.HealthCheckConfigPropertiesType Type;

        [OutputConstructor]
        private HealthCheckConfigProperties(
            Outputs.HealthCheckAlarmIdentifier? alarmIdentifier,

            ImmutableArray<string> childHealthChecks,

            bool? enableSni,

            int? failureThreshold,

            string? fullyQualifiedDomainName,

            int? healthThreshold,

            Pulumi.AwsNative.Route53.HealthCheckConfigPropertiesInsufficientDataHealthStatus? insufficientDataHealthStatus,

            bool? inverted,

            string? ipAddress,

            bool? measureLatency,

            int? port,

            ImmutableArray<string> regions,

            int? requestInterval,

            string? resourcePath,

            string? routingControlArn,

            string? searchString,

            Pulumi.AwsNative.Route53.HealthCheckConfigPropertiesType type)
        {
            AlarmIdentifier = alarmIdentifier;
            ChildHealthChecks = childHealthChecks;
            EnableSni = enableSni;
            FailureThreshold = failureThreshold;
            FullyQualifiedDomainName = fullyQualifiedDomainName;
            HealthThreshold = healthThreshold;
            InsufficientDataHealthStatus = insufficientDataHealthStatus;
            Inverted = inverted;
            IpAddress = ipAddress;
            MeasureLatency = measureLatency;
            Port = port;
            Regions = regions;
            RequestInterval = requestInterval;
            ResourcePath = resourcePath;
            RoutingControlArn = routingControlArn;
            SearchString = searchString;
            Type = type;
        }
    }
}
