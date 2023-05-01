// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.MSK
{
    [EnumType]
    public readonly struct ClusterEncryptionInTransitClientBroker : IEquatable<ClusterEncryptionInTransitClientBroker>
    {
        private readonly string _value;

        private ClusterEncryptionInTransitClientBroker(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ClusterEncryptionInTransitClientBroker Tls { get; } = new ClusterEncryptionInTransitClientBroker("TLS");
        public static ClusterEncryptionInTransitClientBroker TlsPlaintext { get; } = new ClusterEncryptionInTransitClientBroker("TLS_PLAINTEXT");
        public static ClusterEncryptionInTransitClientBroker Plaintext { get; } = new ClusterEncryptionInTransitClientBroker("PLAINTEXT");

        public static bool operator ==(ClusterEncryptionInTransitClientBroker left, ClusterEncryptionInTransitClientBroker right) => left.Equals(right);
        public static bool operator !=(ClusterEncryptionInTransitClientBroker left, ClusterEncryptionInTransitClientBroker right) => !left.Equals(right);

        public static explicit operator string(ClusterEncryptionInTransitClientBroker value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ClusterEncryptionInTransitClientBroker other && Equals(other);
        public bool Equals(ClusterEncryptionInTransitClientBroker other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ClusterEnhancedMonitoring : IEquatable<ClusterEnhancedMonitoring>
    {
        private readonly string _value;

        private ClusterEnhancedMonitoring(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ClusterEnhancedMonitoring Default { get; } = new ClusterEnhancedMonitoring("DEFAULT");
        public static ClusterEnhancedMonitoring PerBroker { get; } = new ClusterEnhancedMonitoring("PER_BROKER");
        public static ClusterEnhancedMonitoring PerTopicPerBroker { get; } = new ClusterEnhancedMonitoring("PER_TOPIC_PER_BROKER");
        public static ClusterEnhancedMonitoring PerTopicPerPartition { get; } = new ClusterEnhancedMonitoring("PER_TOPIC_PER_PARTITION");

        public static bool operator ==(ClusterEnhancedMonitoring left, ClusterEnhancedMonitoring right) => left.Equals(right);
        public static bool operator !=(ClusterEnhancedMonitoring left, ClusterEnhancedMonitoring right) => !left.Equals(right);

        public static explicit operator string(ClusterEnhancedMonitoring value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ClusterEnhancedMonitoring other && Equals(other);
        public bool Equals(ClusterEnhancedMonitoring other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ClusterStorageMode : IEquatable<ClusterStorageMode>
    {
        private readonly string _value;

        private ClusterStorageMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ClusterStorageMode Local { get; } = new ClusterStorageMode("LOCAL");
        public static ClusterStorageMode Tiered { get; } = new ClusterStorageMode("TIERED");

        public static bool operator ==(ClusterStorageMode left, ClusterStorageMode right) => left.Equals(right);
        public static bool operator !=(ClusterStorageMode left, ClusterStorageMode right) => !left.Equals(right);

        public static explicit operator string(ClusterStorageMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ClusterStorageMode other && Equals(other);
        public bool Equals(ClusterStorageMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of private link authentication
    /// </summary>
    [EnumType]
    public readonly struct VpcConnectionAuthentication : IEquatable<VpcConnectionAuthentication>
    {
        private readonly string _value;

        private VpcConnectionAuthentication(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static VpcConnectionAuthentication SaslIam { get; } = new VpcConnectionAuthentication("SASL_IAM");
        public static VpcConnectionAuthentication SaslScram { get; } = new VpcConnectionAuthentication("SASL_SCRAM");
        public static VpcConnectionAuthentication Tls { get; } = new VpcConnectionAuthentication("TLS");

        public static bool operator ==(VpcConnectionAuthentication left, VpcConnectionAuthentication right) => left.Equals(right);
        public static bool operator !=(VpcConnectionAuthentication left, VpcConnectionAuthentication right) => !left.Equals(right);

        public static explicit operator string(VpcConnectionAuthentication value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is VpcConnectionAuthentication other && Equals(other);
        public bool Equals(VpcConnectionAuthentication other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
