// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.VpcLattice
{
    [EnumType]
    public readonly struct AuthPolicyState : IEquatable<AuthPolicyState>
    {
        private readonly string _value;

        private AuthPolicyState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AuthPolicyState Active { get; } = new AuthPolicyState("ACTIVE");
        public static AuthPolicyState Inactive { get; } = new AuthPolicyState("INACTIVE");

        public static bool operator ==(AuthPolicyState left, AuthPolicyState right) => left.Equals(right);
        public static bool operator !=(AuthPolicyState left, AuthPolicyState right) => !left.Equals(right);

        public static explicit operator string(AuthPolicyState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AuthPolicyState other && Equals(other);
        public bool Equals(AuthPolicyState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ListenerProtocol : IEquatable<ListenerProtocol>
    {
        private readonly string _value;

        private ListenerProtocol(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ListenerProtocol Http { get; } = new ListenerProtocol("HTTP");
        public static ListenerProtocol Https { get; } = new ListenerProtocol("HTTPS");

        public static bool operator ==(ListenerProtocol left, ListenerProtocol right) => left.Equals(right);
        public static bool operator !=(ListenerProtocol left, ListenerProtocol right) => !left.Equals(right);

        public static explicit operator string(ListenerProtocol value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ListenerProtocol other && Equals(other);
        public bool Equals(ListenerProtocol other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct RuleHttpMatchMethod : IEquatable<RuleHttpMatchMethod>
    {
        private readonly string _value;

        private RuleHttpMatchMethod(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static RuleHttpMatchMethod Connect { get; } = new RuleHttpMatchMethod("CONNECT");
        public static RuleHttpMatchMethod Delete { get; } = new RuleHttpMatchMethod("DELETE");
        public static RuleHttpMatchMethod Get { get; } = new RuleHttpMatchMethod("GET");
        public static RuleHttpMatchMethod Head { get; } = new RuleHttpMatchMethod("HEAD");
        public static RuleHttpMatchMethod Options { get; } = new RuleHttpMatchMethod("OPTIONS");
        public static RuleHttpMatchMethod Post { get; } = new RuleHttpMatchMethod("POST");
        public static RuleHttpMatchMethod Put { get; } = new RuleHttpMatchMethod("PUT");
        public static RuleHttpMatchMethod Trace { get; } = new RuleHttpMatchMethod("TRACE");

        public static bool operator ==(RuleHttpMatchMethod left, RuleHttpMatchMethod right) => left.Equals(right);
        public static bool operator !=(RuleHttpMatchMethod left, RuleHttpMatchMethod right) => !left.Equals(right);

        public static explicit operator string(RuleHttpMatchMethod value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RuleHttpMatchMethod other && Equals(other);
        public bool Equals(RuleHttpMatchMethod other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ServiceAuthType : IEquatable<ServiceAuthType>
    {
        private readonly string _value;

        private ServiceAuthType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ServiceAuthType None { get; } = new ServiceAuthType("NONE");
        public static ServiceAuthType AwsIam { get; } = new ServiceAuthType("AWS_IAM");

        public static bool operator ==(ServiceAuthType left, ServiceAuthType right) => left.Equals(right);
        public static bool operator !=(ServiceAuthType left, ServiceAuthType right) => !left.Equals(right);

        public static explicit operator string(ServiceAuthType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServiceAuthType other && Equals(other);
        public bool Equals(ServiceAuthType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ServiceNetworkAuthType : IEquatable<ServiceNetworkAuthType>
    {
        private readonly string _value;

        private ServiceNetworkAuthType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ServiceNetworkAuthType None { get; } = new ServiceNetworkAuthType("NONE");
        public static ServiceNetworkAuthType AwsIam { get; } = new ServiceNetworkAuthType("AWS_IAM");

        public static bool operator ==(ServiceNetworkAuthType left, ServiceNetworkAuthType right) => left.Equals(right);
        public static bool operator !=(ServiceNetworkAuthType left, ServiceNetworkAuthType right) => !left.Equals(right);

        public static explicit operator string(ServiceNetworkAuthType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServiceNetworkAuthType other && Equals(other);
        public bool Equals(ServiceNetworkAuthType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ServiceNetworkServiceAssociationStatus : IEquatable<ServiceNetworkServiceAssociationStatus>
    {
        private readonly string _value;

        private ServiceNetworkServiceAssociationStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ServiceNetworkServiceAssociationStatus CreateInProgress { get; } = new ServiceNetworkServiceAssociationStatus("CREATE_IN_PROGRESS");
        public static ServiceNetworkServiceAssociationStatus Active { get; } = new ServiceNetworkServiceAssociationStatus("ACTIVE");
        public static ServiceNetworkServiceAssociationStatus DeleteInProgress { get; } = new ServiceNetworkServiceAssociationStatus("DELETE_IN_PROGRESS");
        public static ServiceNetworkServiceAssociationStatus CreateFailed { get; } = new ServiceNetworkServiceAssociationStatus("CREATE_FAILED");
        public static ServiceNetworkServiceAssociationStatus DeleteFailed { get; } = new ServiceNetworkServiceAssociationStatus("DELETE_FAILED");

        public static bool operator ==(ServiceNetworkServiceAssociationStatus left, ServiceNetworkServiceAssociationStatus right) => left.Equals(right);
        public static bool operator !=(ServiceNetworkServiceAssociationStatus left, ServiceNetworkServiceAssociationStatus right) => !left.Equals(right);

        public static explicit operator string(ServiceNetworkServiceAssociationStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServiceNetworkServiceAssociationStatus other && Equals(other);
        public bool Equals(ServiceNetworkServiceAssociationStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ServiceNetworkVpcAssociationStatus : IEquatable<ServiceNetworkVpcAssociationStatus>
    {
        private readonly string _value;

        private ServiceNetworkVpcAssociationStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ServiceNetworkVpcAssociationStatus CreateInProgress { get; } = new ServiceNetworkVpcAssociationStatus("CREATE_IN_PROGRESS");
        public static ServiceNetworkVpcAssociationStatus Active { get; } = new ServiceNetworkVpcAssociationStatus("ACTIVE");
        public static ServiceNetworkVpcAssociationStatus UpdateInProgress { get; } = new ServiceNetworkVpcAssociationStatus("UPDATE_IN_PROGRESS");
        public static ServiceNetworkVpcAssociationStatus DeleteInProgress { get; } = new ServiceNetworkVpcAssociationStatus("DELETE_IN_PROGRESS");
        public static ServiceNetworkVpcAssociationStatus CreateFailed { get; } = new ServiceNetworkVpcAssociationStatus("CREATE_FAILED");
        public static ServiceNetworkVpcAssociationStatus DeleteFailed { get; } = new ServiceNetworkVpcAssociationStatus("DELETE_FAILED");

        public static bool operator ==(ServiceNetworkVpcAssociationStatus left, ServiceNetworkVpcAssociationStatus right) => left.Equals(right);
        public static bool operator !=(ServiceNetworkVpcAssociationStatus left, ServiceNetworkVpcAssociationStatus right) => !left.Equals(right);

        public static explicit operator string(ServiceNetworkVpcAssociationStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServiceNetworkVpcAssociationStatus other && Equals(other);
        public bool Equals(ServiceNetworkVpcAssociationStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ServiceStatus : IEquatable<ServiceStatus>
    {
        private readonly string _value;

        private ServiceStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ServiceStatus Active { get; } = new ServiceStatus("ACTIVE");
        public static ServiceStatus CreateInProgress { get; } = new ServiceStatus("CREATE_IN_PROGRESS");
        public static ServiceStatus DeleteInProgress { get; } = new ServiceStatus("DELETE_IN_PROGRESS");
        public static ServiceStatus CreateFailed { get; } = new ServiceStatus("CREATE_FAILED");
        public static ServiceStatus DeleteFailed { get; } = new ServiceStatus("DELETE_FAILED");

        public static bool operator ==(ServiceStatus left, ServiceStatus right) => left.Equals(right);
        public static bool operator !=(ServiceStatus left, ServiceStatus right) => !left.Equals(right);

        public static explicit operator string(ServiceStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServiceStatus other && Equals(other);
        public bool Equals(ServiceStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct TargetGroupConfigProtocol : IEquatable<TargetGroupConfigProtocol>
    {
        private readonly string _value;

        private TargetGroupConfigProtocol(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TargetGroupConfigProtocol Http { get; } = new TargetGroupConfigProtocol("HTTP");
        public static TargetGroupConfigProtocol Https { get; } = new TargetGroupConfigProtocol("HTTPS");

        public static bool operator ==(TargetGroupConfigProtocol left, TargetGroupConfigProtocol right) => left.Equals(right);
        public static bool operator !=(TargetGroupConfigProtocol left, TargetGroupConfigProtocol right) => !left.Equals(right);

        public static explicit operator string(TargetGroupConfigProtocol value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TargetGroupConfigProtocol other && Equals(other);
        public bool Equals(TargetGroupConfigProtocol other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct TargetGroupConfigProtocolVersion : IEquatable<TargetGroupConfigProtocolVersion>
    {
        private readonly string _value;

        private TargetGroupConfigProtocolVersion(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TargetGroupConfigProtocolVersion Http1 { get; } = new TargetGroupConfigProtocolVersion("HTTP1");
        public static TargetGroupConfigProtocolVersion Http2 { get; } = new TargetGroupConfigProtocolVersion("HTTP2");
        public static TargetGroupConfigProtocolVersion Grpc { get; } = new TargetGroupConfigProtocolVersion("GRPC");

        public static bool operator ==(TargetGroupConfigProtocolVersion left, TargetGroupConfigProtocolVersion right) => left.Equals(right);
        public static bool operator !=(TargetGroupConfigProtocolVersion left, TargetGroupConfigProtocolVersion right) => !left.Equals(right);

        public static explicit operator string(TargetGroupConfigProtocolVersion value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TargetGroupConfigProtocolVersion other && Equals(other);
        public bool Equals(TargetGroupConfigProtocolVersion other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct TargetGroupHealthCheckConfigProtocol : IEquatable<TargetGroupHealthCheckConfigProtocol>
    {
        private readonly string _value;

        private TargetGroupHealthCheckConfigProtocol(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TargetGroupHealthCheckConfigProtocol Http { get; } = new TargetGroupHealthCheckConfigProtocol("HTTP");
        public static TargetGroupHealthCheckConfigProtocol Https { get; } = new TargetGroupHealthCheckConfigProtocol("HTTPS");

        public static bool operator ==(TargetGroupHealthCheckConfigProtocol left, TargetGroupHealthCheckConfigProtocol right) => left.Equals(right);
        public static bool operator !=(TargetGroupHealthCheckConfigProtocol left, TargetGroupHealthCheckConfigProtocol right) => !left.Equals(right);

        public static explicit operator string(TargetGroupHealthCheckConfigProtocol value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TargetGroupHealthCheckConfigProtocol other && Equals(other);
        public bool Equals(TargetGroupHealthCheckConfigProtocol other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct TargetGroupStatus : IEquatable<TargetGroupStatus>
    {
        private readonly string _value;

        private TargetGroupStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TargetGroupStatus CreateInProgress { get; } = new TargetGroupStatus("CREATE_IN_PROGRESS");
        public static TargetGroupStatus Active { get; } = new TargetGroupStatus("ACTIVE");
        public static TargetGroupStatus DeleteInProgress { get; } = new TargetGroupStatus("DELETE_IN_PROGRESS");
        public static TargetGroupStatus CreateFailed { get; } = new TargetGroupStatus("CREATE_FAILED");
        public static TargetGroupStatus DeleteFailed { get; } = new TargetGroupStatus("DELETE_FAILED");

        public static bool operator ==(TargetGroupStatus left, TargetGroupStatus right) => left.Equals(right);
        public static bool operator !=(TargetGroupStatus left, TargetGroupStatus right) => !left.Equals(right);

        public static explicit operator string(TargetGroupStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TargetGroupStatus other && Equals(other);
        public bool Equals(TargetGroupStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct TargetGroupType : IEquatable<TargetGroupType>
    {
        private readonly string _value;

        private TargetGroupType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TargetGroupType Ip { get; } = new TargetGroupType("IP");
        public static TargetGroupType Lambda { get; } = new TargetGroupType("LAMBDA");
        public static TargetGroupType Instance { get; } = new TargetGroupType("INSTANCE");

        public static bool operator ==(TargetGroupType left, TargetGroupType right) => left.Equals(right);
        public static bool operator !=(TargetGroupType left, TargetGroupType right) => !left.Equals(right);

        public static explicit operator string(TargetGroupType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TargetGroupType other && Equals(other);
        public bool Equals(TargetGroupType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
