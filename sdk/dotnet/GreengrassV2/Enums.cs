// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.GreengrassV2
{
    [EnumType]
    public readonly struct ComponentVersionLambdaEventSourceType : IEquatable<ComponentVersionLambdaEventSourceType>
    {
        private readonly string _value;

        private ComponentVersionLambdaEventSourceType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ComponentVersionLambdaEventSourceType PubSub { get; } = new ComponentVersionLambdaEventSourceType("PUB_SUB");
        public static ComponentVersionLambdaEventSourceType IotCore { get; } = new ComponentVersionLambdaEventSourceType("IOT_CORE");

        public static bool operator ==(ComponentVersionLambdaEventSourceType left, ComponentVersionLambdaEventSourceType right) => left.Equals(right);
        public static bool operator !=(ComponentVersionLambdaEventSourceType left, ComponentVersionLambdaEventSourceType right) => !left.Equals(right);

        public static explicit operator string(ComponentVersionLambdaEventSourceType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ComponentVersionLambdaEventSourceType other && Equals(other);
        public bool Equals(ComponentVersionLambdaEventSourceType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ComponentVersionLambdaExecutionParametersInputPayloadEncodingType : IEquatable<ComponentVersionLambdaExecutionParametersInputPayloadEncodingType>
    {
        private readonly string _value;

        private ComponentVersionLambdaExecutionParametersInputPayloadEncodingType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ComponentVersionLambdaExecutionParametersInputPayloadEncodingType Json { get; } = new ComponentVersionLambdaExecutionParametersInputPayloadEncodingType("json");
        public static ComponentVersionLambdaExecutionParametersInputPayloadEncodingType Binary { get; } = new ComponentVersionLambdaExecutionParametersInputPayloadEncodingType("binary");

        public static bool operator ==(ComponentVersionLambdaExecutionParametersInputPayloadEncodingType left, ComponentVersionLambdaExecutionParametersInputPayloadEncodingType right) => left.Equals(right);
        public static bool operator !=(ComponentVersionLambdaExecutionParametersInputPayloadEncodingType left, ComponentVersionLambdaExecutionParametersInputPayloadEncodingType right) => !left.Equals(right);

        public static explicit operator string(ComponentVersionLambdaExecutionParametersInputPayloadEncodingType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ComponentVersionLambdaExecutionParametersInputPayloadEncodingType other && Equals(other);
        public bool Equals(ComponentVersionLambdaExecutionParametersInputPayloadEncodingType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ComponentVersionLambdaFilesystemPermission : IEquatable<ComponentVersionLambdaFilesystemPermission>
    {
        private readonly string _value;

        private ComponentVersionLambdaFilesystemPermission(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ComponentVersionLambdaFilesystemPermission Ro { get; } = new ComponentVersionLambdaFilesystemPermission("ro");
        public static ComponentVersionLambdaFilesystemPermission Rw { get; } = new ComponentVersionLambdaFilesystemPermission("rw");

        public static bool operator ==(ComponentVersionLambdaFilesystemPermission left, ComponentVersionLambdaFilesystemPermission right) => left.Equals(right);
        public static bool operator !=(ComponentVersionLambdaFilesystemPermission left, ComponentVersionLambdaFilesystemPermission right) => !left.Equals(right);

        public static explicit operator string(ComponentVersionLambdaFilesystemPermission value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ComponentVersionLambdaFilesystemPermission other && Equals(other);
        public bool Equals(ComponentVersionLambdaFilesystemPermission other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ComponentVersionLambdaLinuxProcessParamsIsolationMode : IEquatable<ComponentVersionLambdaLinuxProcessParamsIsolationMode>
    {
        private readonly string _value;

        private ComponentVersionLambdaLinuxProcessParamsIsolationMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ComponentVersionLambdaLinuxProcessParamsIsolationMode GreengrassContainer { get; } = new ComponentVersionLambdaLinuxProcessParamsIsolationMode("GreengrassContainer");
        public static ComponentVersionLambdaLinuxProcessParamsIsolationMode NoContainer { get; } = new ComponentVersionLambdaLinuxProcessParamsIsolationMode("NoContainer");

        public static bool operator ==(ComponentVersionLambdaLinuxProcessParamsIsolationMode left, ComponentVersionLambdaLinuxProcessParamsIsolationMode right) => left.Equals(right);
        public static bool operator !=(ComponentVersionLambdaLinuxProcessParamsIsolationMode left, ComponentVersionLambdaLinuxProcessParamsIsolationMode right) => !left.Equals(right);

        public static explicit operator string(ComponentVersionLambdaLinuxProcessParamsIsolationMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ComponentVersionLambdaLinuxProcessParamsIsolationMode other && Equals(other);
        public bool Equals(ComponentVersionLambdaLinuxProcessParamsIsolationMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct DeploymentComponentUpdatePolicyAction : IEquatable<DeploymentComponentUpdatePolicyAction>
    {
        private readonly string _value;

        private DeploymentComponentUpdatePolicyAction(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DeploymentComponentUpdatePolicyAction NotifyComponents { get; } = new DeploymentComponentUpdatePolicyAction("NOTIFY_COMPONENTS");
        public static DeploymentComponentUpdatePolicyAction SkipNotifyComponents { get; } = new DeploymentComponentUpdatePolicyAction("SKIP_NOTIFY_COMPONENTS");

        public static bool operator ==(DeploymentComponentUpdatePolicyAction left, DeploymentComponentUpdatePolicyAction right) => left.Equals(right);
        public static bool operator !=(DeploymentComponentUpdatePolicyAction left, DeploymentComponentUpdatePolicyAction right) => !left.Equals(right);

        public static explicit operator string(DeploymentComponentUpdatePolicyAction value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DeploymentComponentUpdatePolicyAction other && Equals(other);
        public bool Equals(DeploymentComponentUpdatePolicyAction other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct DeploymentIoTJobAbortCriteriaAction : IEquatable<DeploymentIoTJobAbortCriteriaAction>
    {
        private readonly string _value;

        private DeploymentIoTJobAbortCriteriaAction(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DeploymentIoTJobAbortCriteriaAction Cancel { get; } = new DeploymentIoTJobAbortCriteriaAction("CANCEL");

        public static bool operator ==(DeploymentIoTJobAbortCriteriaAction left, DeploymentIoTJobAbortCriteriaAction right) => left.Equals(right);
        public static bool operator !=(DeploymentIoTJobAbortCriteriaAction left, DeploymentIoTJobAbortCriteriaAction right) => !left.Equals(right);

        public static explicit operator string(DeploymentIoTJobAbortCriteriaAction value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DeploymentIoTJobAbortCriteriaAction other && Equals(other);
        public bool Equals(DeploymentIoTJobAbortCriteriaAction other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct DeploymentIoTJobAbortCriteriaFailureType : IEquatable<DeploymentIoTJobAbortCriteriaFailureType>
    {
        private readonly string _value;

        private DeploymentIoTJobAbortCriteriaFailureType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DeploymentIoTJobAbortCriteriaFailureType Failed { get; } = new DeploymentIoTJobAbortCriteriaFailureType("FAILED");
        public static DeploymentIoTJobAbortCriteriaFailureType Rejected { get; } = new DeploymentIoTJobAbortCriteriaFailureType("REJECTED");
        public static DeploymentIoTJobAbortCriteriaFailureType TimedOut { get; } = new DeploymentIoTJobAbortCriteriaFailureType("TIMED_OUT");
        public static DeploymentIoTJobAbortCriteriaFailureType All { get; } = new DeploymentIoTJobAbortCriteriaFailureType("ALL");

        public static bool operator ==(DeploymentIoTJobAbortCriteriaFailureType left, DeploymentIoTJobAbortCriteriaFailureType right) => left.Equals(right);
        public static bool operator !=(DeploymentIoTJobAbortCriteriaFailureType left, DeploymentIoTJobAbortCriteriaFailureType right) => !left.Equals(right);

        public static explicit operator string(DeploymentIoTJobAbortCriteriaFailureType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DeploymentIoTJobAbortCriteriaFailureType other && Equals(other);
        public bool Equals(DeploymentIoTJobAbortCriteriaFailureType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct DeploymentPoliciesFailureHandlingPolicy : IEquatable<DeploymentPoliciesFailureHandlingPolicy>
    {
        private readonly string _value;

        private DeploymentPoliciesFailureHandlingPolicy(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DeploymentPoliciesFailureHandlingPolicy Rollback { get; } = new DeploymentPoliciesFailureHandlingPolicy("ROLLBACK");
        public static DeploymentPoliciesFailureHandlingPolicy DoNothing { get; } = new DeploymentPoliciesFailureHandlingPolicy("DO_NOTHING");

        public static bool operator ==(DeploymentPoliciesFailureHandlingPolicy left, DeploymentPoliciesFailureHandlingPolicy right) => left.Equals(right);
        public static bool operator !=(DeploymentPoliciesFailureHandlingPolicy left, DeploymentPoliciesFailureHandlingPolicy right) => !left.Equals(right);

        public static explicit operator string(DeploymentPoliciesFailureHandlingPolicy value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DeploymentPoliciesFailureHandlingPolicy other && Equals(other);
        public bool Equals(DeploymentPoliciesFailureHandlingPolicy other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
