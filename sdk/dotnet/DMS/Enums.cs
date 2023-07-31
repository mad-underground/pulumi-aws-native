// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.DMS
{
    /// <summary>
    /// The type of AWS DMS Serverless replication to provision using this replication configuration
    /// </summary>
    [EnumType]
    public readonly struct ReplicationConfigReplicationType : IEquatable<ReplicationConfigReplicationType>
    {
        private readonly string _value;

        private ReplicationConfigReplicationType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ReplicationConfigReplicationType FullLoad { get; } = new ReplicationConfigReplicationType("full-load");
        public static ReplicationConfigReplicationType FullLoadAndCdc { get; } = new ReplicationConfigReplicationType("full-load-and-cdc");
        public static ReplicationConfigReplicationType Cdc { get; } = new ReplicationConfigReplicationType("cdc");

        public static bool operator ==(ReplicationConfigReplicationType left, ReplicationConfigReplicationType right) => left.Equals(right);
        public static bool operator !=(ReplicationConfigReplicationType left, ReplicationConfigReplicationType right) => !left.Equals(right);

        public static explicit operator string(ReplicationConfigReplicationType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ReplicationConfigReplicationType other && Equals(other);
        public bool Equals(ReplicationConfigReplicationType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
