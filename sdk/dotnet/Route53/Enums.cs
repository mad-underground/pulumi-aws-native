// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.Route53
{
    [EnumType]
    public readonly struct HealthCheckConfigPropertiesInsufficientDataHealthStatus : IEquatable<HealthCheckConfigPropertiesInsufficientDataHealthStatus>
    {
        private readonly string _value;

        private HealthCheckConfigPropertiesInsufficientDataHealthStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static HealthCheckConfigPropertiesInsufficientDataHealthStatus Healthy { get; } = new HealthCheckConfigPropertiesInsufficientDataHealthStatus("Healthy");
        public static HealthCheckConfigPropertiesInsufficientDataHealthStatus LastKnownStatus { get; } = new HealthCheckConfigPropertiesInsufficientDataHealthStatus("LastKnownStatus");
        public static HealthCheckConfigPropertiesInsufficientDataHealthStatus Unhealthy { get; } = new HealthCheckConfigPropertiesInsufficientDataHealthStatus("Unhealthy");

        public static bool operator ==(HealthCheckConfigPropertiesInsufficientDataHealthStatus left, HealthCheckConfigPropertiesInsufficientDataHealthStatus right) => left.Equals(right);
        public static bool operator !=(HealthCheckConfigPropertiesInsufficientDataHealthStatus left, HealthCheckConfigPropertiesInsufficientDataHealthStatus right) => !left.Equals(right);

        public static explicit operator string(HealthCheckConfigPropertiesInsufficientDataHealthStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is HealthCheckConfigPropertiesInsufficientDataHealthStatus other && Equals(other);
        public bool Equals(HealthCheckConfigPropertiesInsufficientDataHealthStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct HealthCheckConfigPropertiesType : IEquatable<HealthCheckConfigPropertiesType>
    {
        private readonly string _value;

        private HealthCheckConfigPropertiesType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static HealthCheckConfigPropertiesType Calculated { get; } = new HealthCheckConfigPropertiesType("CALCULATED");
        public static HealthCheckConfigPropertiesType CloudwatchMetric { get; } = new HealthCheckConfigPropertiesType("CLOUDWATCH_METRIC");
        public static HealthCheckConfigPropertiesType Http { get; } = new HealthCheckConfigPropertiesType("HTTP");
        public static HealthCheckConfigPropertiesType HttpStrMatch { get; } = new HealthCheckConfigPropertiesType("HTTP_STR_MATCH");
        public static HealthCheckConfigPropertiesType Https { get; } = new HealthCheckConfigPropertiesType("HTTPS");
        public static HealthCheckConfigPropertiesType HttpsStrMatch { get; } = new HealthCheckConfigPropertiesType("HTTPS_STR_MATCH");
        public static HealthCheckConfigPropertiesType Tcp { get; } = new HealthCheckConfigPropertiesType("TCP");
        public static HealthCheckConfigPropertiesType RecoveryControl { get; } = new HealthCheckConfigPropertiesType("RECOVERY_CONTROL");

        public static bool operator ==(HealthCheckConfigPropertiesType left, HealthCheckConfigPropertiesType right) => left.Equals(right);
        public static bool operator !=(HealthCheckConfigPropertiesType left, HealthCheckConfigPropertiesType right) => !left.Equals(right);

        public static explicit operator string(HealthCheckConfigPropertiesType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is HealthCheckConfigPropertiesType other && Equals(other);
        public bool Equals(HealthCheckConfigPropertiesType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// A string specifying the initial status of the key signing key (KSK). You can set the value to ACTIVE or INACTIVE.
    /// </summary>
    [EnumType]
    public readonly struct KeySigningKeyStatus : IEquatable<KeySigningKeyStatus>
    {
        private readonly string _value;

        private KeySigningKeyStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static KeySigningKeyStatus Active { get; } = new KeySigningKeyStatus("ACTIVE");
        public static KeySigningKeyStatus Inactive { get; } = new KeySigningKeyStatus("INACTIVE");

        public static bool operator ==(KeySigningKeyStatus left, KeySigningKeyStatus right) => left.Equals(right);
        public static bool operator !=(KeySigningKeyStatus left, KeySigningKeyStatus right) => !left.Equals(right);

        public static explicit operator string(KeySigningKeyStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is KeySigningKeyStatus other && Equals(other);
        public bool Equals(KeySigningKeyStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
