// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.DataZone
{
    /// <summary>
    /// Specifies whether the data source is enabled.
    /// </summary>
    [EnumType]
    public readonly struct DataSourceEnableSetting : IEquatable<DataSourceEnableSetting>
    {
        private readonly string _value;

        private DataSourceEnableSetting(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DataSourceEnableSetting Enabled { get; } = new DataSourceEnableSetting("ENABLED");
        public static DataSourceEnableSetting Disabled { get; } = new DataSourceEnableSetting("DISABLED");

        public static bool operator ==(DataSourceEnableSetting left, DataSourceEnableSetting right) => left.Equals(right);
        public static bool operator !=(DataSourceEnableSetting left, DataSourceEnableSetting right) => !left.Equals(right);

        public static explicit operator string(DataSourceEnableSetting value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DataSourceEnableSetting other && Equals(other);
        public bool Equals(DataSourceEnableSetting other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The search filter expression type.
    /// </summary>
    [EnumType]
    public readonly struct DataSourceFilterExpressionType : IEquatable<DataSourceFilterExpressionType>
    {
        private readonly string _value;

        private DataSourceFilterExpressionType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DataSourceFilterExpressionType Include { get; } = new DataSourceFilterExpressionType("INCLUDE");
        public static DataSourceFilterExpressionType Exclude { get; } = new DataSourceFilterExpressionType("EXCLUDE");

        public static bool operator ==(DataSourceFilterExpressionType left, DataSourceFilterExpressionType right) => left.Equals(right);
        public static bool operator !=(DataSourceFilterExpressionType left, DataSourceFilterExpressionType right) => !left.Equals(right);

        public static explicit operator string(DataSourceFilterExpressionType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DataSourceFilterExpressionType other && Equals(other);
        public bool Equals(DataSourceFilterExpressionType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The status of the data source.
    /// </summary>
    [EnumType]
    public readonly struct DataSourceStatus : IEquatable<DataSourceStatus>
    {
        private readonly string _value;

        private DataSourceStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DataSourceStatus Creating { get; } = new DataSourceStatus("CREATING");
        public static DataSourceStatus FailedCreation { get; } = new DataSourceStatus("FAILED_CREATION");
        public static DataSourceStatus Ready { get; } = new DataSourceStatus("READY");
        public static DataSourceStatus Updating { get; } = new DataSourceStatus("UPDATING");
        public static DataSourceStatus FailedUpdate { get; } = new DataSourceStatus("FAILED_UPDATE");
        public static DataSourceStatus Running { get; } = new DataSourceStatus("RUNNING");
        public static DataSourceStatus Deleting { get; } = new DataSourceStatus("DELETING");
        public static DataSourceStatus FailedDeletion { get; } = new DataSourceStatus("FAILED_DELETION");

        public static bool operator ==(DataSourceStatus left, DataSourceStatus right) => left.Equals(right);
        public static bool operator !=(DataSourceStatus left, DataSourceStatus right) => !left.Equals(right);

        public static explicit operator string(DataSourceStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DataSourceStatus other && Equals(other);
        public bool Equals(DataSourceStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of single sign-on in Amazon DataZone.
    /// </summary>
    [EnumType]
    public readonly struct DomainAuthType : IEquatable<DomainAuthType>
    {
        private readonly string _value;

        private DomainAuthType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DomainAuthType IamIdc { get; } = new DomainAuthType("IAM_IDC");
        public static DomainAuthType Disabled { get; } = new DomainAuthType("DISABLED");

        public static bool operator ==(DomainAuthType left, DomainAuthType right) => left.Equals(right);
        public static bool operator !=(DomainAuthType left, DomainAuthType right) => !left.Equals(right);

        public static explicit operator string(DomainAuthType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DomainAuthType other && Equals(other);
        public bool Equals(DomainAuthType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The status of the Amazon DataZone domain.
    /// </summary>
    [EnumType]
    public readonly struct DomainStatus : IEquatable<DomainStatus>
    {
        private readonly string _value;

        private DomainStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DomainStatus Creating { get; } = new DomainStatus("CREATING");
        public static DomainStatus Available { get; } = new DomainStatus("AVAILABLE");
        public static DomainStatus CreationFailed { get; } = new DomainStatus("CREATION_FAILED");
        public static DomainStatus Deleting { get; } = new DomainStatus("DELETING");
        public static DomainStatus Deleted { get; } = new DomainStatus("DELETED");
        public static DomainStatus DeletionFailed { get; } = new DomainStatus("DELETION_FAILED");

        public static bool operator ==(DomainStatus left, DomainStatus right) => left.Equals(right);
        public static bool operator !=(DomainStatus left, DomainStatus right) => !left.Equals(right);

        public static explicit operator string(DomainStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DomainStatus other && Equals(other);
        public bool Equals(DomainStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The single sign-on user assignment in Amazon DataZone.
    /// </summary>
    [EnumType]
    public readonly struct DomainUserAssignment : IEquatable<DomainUserAssignment>
    {
        private readonly string _value;

        private DomainUserAssignment(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DomainUserAssignment Automatic { get; } = new DomainUserAssignment("AUTOMATIC");
        public static DomainUserAssignment Manual { get; } = new DomainUserAssignment("MANUAL");

        public static bool operator ==(DomainUserAssignment left, DomainUserAssignment right) => left.Equals(right);
        public static bool operator !=(DomainUserAssignment left, DomainUserAssignment right) => !left.Equals(right);

        public static explicit operator string(DomainUserAssignment value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DomainUserAssignment other && Equals(other);
        public bool Equals(DomainUserAssignment other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The status of the Amazon DataZone environment.
    /// </summary>
    [EnumType]
    public readonly struct EnvironmentStatus : IEquatable<EnvironmentStatus>
    {
        private readonly string _value;

        private EnvironmentStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EnvironmentStatus Active { get; } = new EnvironmentStatus("ACTIVE");
        public static EnvironmentStatus Creating { get; } = new EnvironmentStatus("CREATING");
        public static EnvironmentStatus Updating { get; } = new EnvironmentStatus("UPDATING");
        public static EnvironmentStatus Deleting { get; } = new EnvironmentStatus("DELETING");
        public static EnvironmentStatus CreateFailed { get; } = new EnvironmentStatus("CREATE_FAILED");
        public static EnvironmentStatus UpdateFailed { get; } = new EnvironmentStatus("UPDATE_FAILED");
        public static EnvironmentStatus DeleteFailed { get; } = new EnvironmentStatus("DELETE_FAILED");
        public static EnvironmentStatus ValidationFailed { get; } = new EnvironmentStatus("VALIDATION_FAILED");
        public static EnvironmentStatus Suspended { get; } = new EnvironmentStatus("SUSPENDED");
        public static EnvironmentStatus Disabled { get; } = new EnvironmentStatus("DISABLED");
        public static EnvironmentStatus Expired { get; } = new EnvironmentStatus("EXPIRED");
        public static EnvironmentStatus Deleted { get; } = new EnvironmentStatus("DELETED");
        public static EnvironmentStatus Inaccessible { get; } = new EnvironmentStatus("INACCESSIBLE");

        public static bool operator ==(EnvironmentStatus left, EnvironmentStatus right) => left.Equals(right);
        public static bool operator !=(EnvironmentStatus left, EnvironmentStatus right) => !left.Equals(right);

        public static explicit operator string(EnvironmentStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EnvironmentStatus other && Equals(other);
        public bool Equals(EnvironmentStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
