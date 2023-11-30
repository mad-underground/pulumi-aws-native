// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.OpenSearchServerless
{
    /// <summary>
    /// The possible types for the access policy
    /// </summary>
    [EnumType]
    public readonly struct AccessPolicyType : IEquatable<AccessPolicyType>
    {
        private readonly string _value;

        private AccessPolicyType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AccessPolicyType Data { get; } = new AccessPolicyType("data");

        public static bool operator ==(AccessPolicyType left, AccessPolicyType right) => left.Equals(right);
        public static bool operator !=(AccessPolicyType left, AccessPolicyType right) => !left.Equals(right);

        public static explicit operator string(AccessPolicyType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AccessPolicyType other && Equals(other);
        public bool Equals(AccessPolicyType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The possible standby replicas for the collection
    /// </summary>
    [EnumType]
    public readonly struct CollectionStandbyReplicas : IEquatable<CollectionStandbyReplicas>
    {
        private readonly string _value;

        private CollectionStandbyReplicas(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CollectionStandbyReplicas Enabled { get; } = new CollectionStandbyReplicas("ENABLED");
        public static CollectionStandbyReplicas Disabled { get; } = new CollectionStandbyReplicas("DISABLED");

        public static bool operator ==(CollectionStandbyReplicas left, CollectionStandbyReplicas right) => left.Equals(right);
        public static bool operator !=(CollectionStandbyReplicas left, CollectionStandbyReplicas right) => !left.Equals(right);

        public static explicit operator string(CollectionStandbyReplicas value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CollectionStandbyReplicas other && Equals(other);
        public bool Equals(CollectionStandbyReplicas other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The possible types for the collection
    /// </summary>
    [EnumType]
    public readonly struct CollectionType : IEquatable<CollectionType>
    {
        private readonly string _value;

        private CollectionType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CollectionType Search { get; } = new CollectionType("SEARCH");
        public static CollectionType Timeseries { get; } = new CollectionType("TIMESERIES");
        public static CollectionType Vectorsearch { get; } = new CollectionType("VECTORSEARCH");

        public static bool operator ==(CollectionType left, CollectionType right) => left.Equals(right);
        public static bool operator !=(CollectionType left, CollectionType right) => !left.Equals(right);

        public static explicit operator string(CollectionType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CollectionType other && Equals(other);
        public bool Equals(CollectionType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of lifecycle policy
    /// </summary>
    [EnumType]
    public readonly struct LifecyclePolicyType : IEquatable<LifecyclePolicyType>
    {
        private readonly string _value;

        private LifecyclePolicyType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static LifecyclePolicyType Retention { get; } = new LifecyclePolicyType("retention");

        public static bool operator ==(LifecyclePolicyType left, LifecyclePolicyType right) => left.Equals(right);
        public static bool operator !=(LifecyclePolicyType left, LifecyclePolicyType right) => !left.Equals(right);

        public static explicit operator string(LifecyclePolicyType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is LifecyclePolicyType other && Equals(other);
        public bool Equals(LifecyclePolicyType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Config type for security config
    /// </summary>
    [EnumType]
    public readonly struct SecurityConfigType : IEquatable<SecurityConfigType>
    {
        private readonly string _value;

        private SecurityConfigType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SecurityConfigType Saml { get; } = new SecurityConfigType("saml");

        public static bool operator ==(SecurityConfigType left, SecurityConfigType right) => left.Equals(right);
        public static bool operator !=(SecurityConfigType left, SecurityConfigType right) => !left.Equals(right);

        public static explicit operator string(SecurityConfigType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SecurityConfigType other && Equals(other);
        public bool Equals(SecurityConfigType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The possible types for the network policy
    /// </summary>
    [EnumType]
    public readonly struct SecurityPolicyType : IEquatable<SecurityPolicyType>
    {
        private readonly string _value;

        private SecurityPolicyType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SecurityPolicyType Encryption { get; } = new SecurityPolicyType("encryption");
        public static SecurityPolicyType Network { get; } = new SecurityPolicyType("network");

        public static bool operator ==(SecurityPolicyType left, SecurityPolicyType right) => left.Equals(right);
        public static bool operator !=(SecurityPolicyType left, SecurityPolicyType right) => !left.Equals(right);

        public static explicit operator string(SecurityPolicyType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SecurityPolicyType other && Equals(other);
        public bool Equals(SecurityPolicyType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
