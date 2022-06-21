// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.RedshiftServerless
{
    [EnumType]
    public readonly struct NamespaceLogExport : IEquatable<NamespaceLogExport>
    {
        private readonly string _value;

        private NamespaceLogExport(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static NamespaceLogExport Useractivitylog { get; } = new NamespaceLogExport("useractivitylog");
        public static NamespaceLogExport Userlog { get; } = new NamespaceLogExport("userlog");
        public static NamespaceLogExport Connectionlog { get; } = new NamespaceLogExport("connectionlog");

        public static bool operator ==(NamespaceLogExport left, NamespaceLogExport right) => left.Equals(right);
        public static bool operator !=(NamespaceLogExport left, NamespaceLogExport right) => !left.Equals(right);

        public static explicit operator string(NamespaceLogExport value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is NamespaceLogExport other && Equals(other);
        public bool Equals(NamespaceLogExport other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct NamespaceStatus : IEquatable<NamespaceStatus>
    {
        private readonly string _value;

        private NamespaceStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static NamespaceStatus Available { get; } = new NamespaceStatus("AVAILABLE");
        public static NamespaceStatus Modifying { get; } = new NamespaceStatus("MODIFYING");
        public static NamespaceStatus Deleting { get; } = new NamespaceStatus("DELETING");

        public static bool operator ==(NamespaceStatus left, NamespaceStatus right) => left.Equals(right);
        public static bool operator !=(NamespaceStatus left, NamespaceStatus right) => !left.Equals(right);

        public static explicit operator string(NamespaceStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is NamespaceStatus other && Equals(other);
        public bool Equals(NamespaceStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}