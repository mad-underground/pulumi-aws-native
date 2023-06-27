// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.Cassandra
{
    [EnumType]
    public readonly struct KeyspaceRegionListItem : IEquatable<KeyspaceRegionListItem>
    {
        private readonly string _value;

        private KeyspaceRegionListItem(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static KeyspaceRegionListItem ApNortheast1 { get; } = new KeyspaceRegionListItem("ap-northeast-1");
        public static KeyspaceRegionListItem ApNortheast2 { get; } = new KeyspaceRegionListItem("ap-northeast-2");
        public static KeyspaceRegionListItem ApSouth1 { get; } = new KeyspaceRegionListItem("ap-south-1");
        public static KeyspaceRegionListItem ApSoutheast1 { get; } = new KeyspaceRegionListItem("ap-southeast-1");
        public static KeyspaceRegionListItem ApSoutheast2 { get; } = new KeyspaceRegionListItem("ap-southeast-2");
        public static KeyspaceRegionListItem CaCentral1 { get; } = new KeyspaceRegionListItem("ca-central-1");
        public static KeyspaceRegionListItem EuCentral1 { get; } = new KeyspaceRegionListItem("eu-central-1");
        public static KeyspaceRegionListItem EuNorth1 { get; } = new KeyspaceRegionListItem("eu-north-1");
        public static KeyspaceRegionListItem EuWest1 { get; } = new KeyspaceRegionListItem("eu-west-1");
        public static KeyspaceRegionListItem EuWest2 { get; } = new KeyspaceRegionListItem("eu-west-2");
        public static KeyspaceRegionListItem EuWest3 { get; } = new KeyspaceRegionListItem("eu-west-3");
        public static KeyspaceRegionListItem SaEast1 { get; } = new KeyspaceRegionListItem("sa-east-1");
        public static KeyspaceRegionListItem UsEast1 { get; } = new KeyspaceRegionListItem("us-east-1");
        public static KeyspaceRegionListItem UsEast2 { get; } = new KeyspaceRegionListItem("us-east-2");
        public static KeyspaceRegionListItem UsWest1 { get; } = new KeyspaceRegionListItem("us-west-1");
        public static KeyspaceRegionListItem UsWest2 { get; } = new KeyspaceRegionListItem("us-west-2");

        public static bool operator ==(KeyspaceRegionListItem left, KeyspaceRegionListItem right) => left.Equals(right);
        public static bool operator !=(KeyspaceRegionListItem left, KeyspaceRegionListItem right) => !left.Equals(right);

        public static explicit operator string(KeyspaceRegionListItem value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is KeyspaceRegionListItem other && Equals(other);
        public bool Equals(KeyspaceRegionListItem other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct KeyspaceReplicationSpecificationReplicationStrategy : IEquatable<KeyspaceReplicationSpecificationReplicationStrategy>
    {
        private readonly string _value;

        private KeyspaceReplicationSpecificationReplicationStrategy(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static KeyspaceReplicationSpecificationReplicationStrategy SingleRegion { get; } = new KeyspaceReplicationSpecificationReplicationStrategy("SINGLE_REGION");
        public static KeyspaceReplicationSpecificationReplicationStrategy MultiRegion { get; } = new KeyspaceReplicationSpecificationReplicationStrategy("MULTI_REGION");

        public static bool operator ==(KeyspaceReplicationSpecificationReplicationStrategy left, KeyspaceReplicationSpecificationReplicationStrategy right) => left.Equals(right);
        public static bool operator !=(KeyspaceReplicationSpecificationReplicationStrategy left, KeyspaceReplicationSpecificationReplicationStrategy right) => !left.Equals(right);

        public static explicit operator string(KeyspaceReplicationSpecificationReplicationStrategy value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is KeyspaceReplicationSpecificationReplicationStrategy other && Equals(other);
        public bool Equals(KeyspaceReplicationSpecificationReplicationStrategy other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct TableClusteringKeyColumnOrderBy : IEquatable<TableClusteringKeyColumnOrderBy>
    {
        private readonly string _value;

        private TableClusteringKeyColumnOrderBy(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TableClusteringKeyColumnOrderBy Asc { get; } = new TableClusteringKeyColumnOrderBy("ASC");
        public static TableClusteringKeyColumnOrderBy Desc { get; } = new TableClusteringKeyColumnOrderBy("DESC");

        public static bool operator ==(TableClusteringKeyColumnOrderBy left, TableClusteringKeyColumnOrderBy right) => left.Equals(right);
        public static bool operator !=(TableClusteringKeyColumnOrderBy left, TableClusteringKeyColumnOrderBy right) => !left.Equals(right);

        public static explicit operator string(TableClusteringKeyColumnOrderBy value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TableClusteringKeyColumnOrderBy other && Equals(other);
        public bool Equals(TableClusteringKeyColumnOrderBy other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Server-side encryption type
    /// </summary>
    [EnumType]
    public readonly struct TableEncryptionType : IEquatable<TableEncryptionType>
    {
        private readonly string _value;

        private TableEncryptionType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TableEncryptionType AwsOwnedKmsKey { get; } = new TableEncryptionType("AWS_OWNED_KMS_KEY");
        public static TableEncryptionType CustomerManagedKmsKey { get; } = new TableEncryptionType("CUSTOMER_MANAGED_KMS_KEY");

        public static bool operator ==(TableEncryptionType left, TableEncryptionType right) => left.Equals(right);
        public static bool operator !=(TableEncryptionType left, TableEncryptionType right) => !left.Equals(right);

        public static explicit operator string(TableEncryptionType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TableEncryptionType other && Equals(other);
        public bool Equals(TableEncryptionType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Capacity mode for the specified table
    /// </summary>
    [EnumType]
    public readonly struct TableMode : IEquatable<TableMode>
    {
        private readonly string _value;

        private TableMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TableMode Provisioned { get; } = new TableMode("PROVISIONED");
        public static TableMode OnDemand { get; } = new TableMode("ON_DEMAND");

        public static bool operator ==(TableMode left, TableMode right) => left.Equals(right);
        public static bool operator !=(TableMode left, TableMode right) => !left.Equals(right);

        public static explicit operator string(TableMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TableMode other && Equals(other);
        public bool Equals(TableMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
