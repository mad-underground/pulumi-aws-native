// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.EKS
{
    /// <summary>
    /// Resolve parameter value conflicts
    /// </summary>
    [EnumType]
    public readonly struct AddonResolveConflicts : IEquatable<AddonResolveConflicts>
    {
        private readonly string _value;

        private AddonResolveConflicts(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AddonResolveConflicts None { get; } = new AddonResolveConflicts("NONE");
        public static AddonResolveConflicts Overwrite { get; } = new AddonResolveConflicts("OVERWRITE");
        public static AddonResolveConflicts Preserve { get; } = new AddonResolveConflicts("PRESERVE");

        public static bool operator ==(AddonResolveConflicts left, AddonResolveConflicts right) => left.Equals(right);
        public static bool operator !=(AddonResolveConflicts left, AddonResolveConflicts right) => !left.Equals(right);

        public static explicit operator string(AddonResolveConflicts value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AddonResolveConflicts other && Equals(other);
        public bool Equals(AddonResolveConflicts other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Ipv4 or Ipv6. You can only specify ipv6 for 1.21 and later clusters that use version 1.10.1 or later of the Amazon VPC CNI add-on
    /// </summary>
    [EnumType]
    public readonly struct ClusterKubernetesNetworkConfigIpFamily : IEquatable<ClusterKubernetesNetworkConfigIpFamily>
    {
        private readonly string _value;

        private ClusterKubernetesNetworkConfigIpFamily(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ClusterKubernetesNetworkConfigIpFamily Ipv4 { get; } = new ClusterKubernetesNetworkConfigIpFamily("ipv4");
        public static ClusterKubernetesNetworkConfigIpFamily Ipv6 { get; } = new ClusterKubernetesNetworkConfigIpFamily("ipv6");

        public static bool operator ==(ClusterKubernetesNetworkConfigIpFamily left, ClusterKubernetesNetworkConfigIpFamily right) => left.Equals(right);
        public static bool operator !=(ClusterKubernetesNetworkConfigIpFamily left, ClusterKubernetesNetworkConfigIpFamily right) => !left.Equals(right);

        public static explicit operator string(ClusterKubernetesNetworkConfigIpFamily value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ClusterKubernetesNetworkConfigIpFamily other && Equals(other);
        public bool Equals(ClusterKubernetesNetworkConfigIpFamily other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of the identity provider configuration.
    /// </summary>
    [EnumType]
    public readonly struct IdentityProviderConfigType : IEquatable<IdentityProviderConfigType>
    {
        private readonly string _value;

        private IdentityProviderConfigType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static IdentityProviderConfigType Oidc { get; } = new IdentityProviderConfigType("oidc");

        public static bool operator ==(IdentityProviderConfigType left, IdentityProviderConfigType right) => left.Equals(right);
        public static bool operator !=(IdentityProviderConfigType left, IdentityProviderConfigType right) => !left.Equals(right);

        public static explicit operator string(IdentityProviderConfigType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is IdentityProviderConfigType other && Equals(other);
        public bool Equals(IdentityProviderConfigType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
