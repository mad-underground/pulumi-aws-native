// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.Scheduler
{
    /// <summary>
    /// Specifies whether the task's elastic network interface receives a public IP address. You can specify ENABLED only when LaunchType in EcsParameters is set to FARGATE.
    /// </summary>
    [EnumType]
    public readonly struct ScheduleAssignPublicIp : IEquatable<ScheduleAssignPublicIp>
    {
        private readonly string _value;

        private ScheduleAssignPublicIp(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScheduleAssignPublicIp Enabled { get; } = new ScheduleAssignPublicIp("ENABLED");
        public static ScheduleAssignPublicIp Disabled { get; } = new ScheduleAssignPublicIp("DISABLED");

        public static bool operator ==(ScheduleAssignPublicIp left, ScheduleAssignPublicIp right) => left.Equals(right);
        public static bool operator !=(ScheduleAssignPublicIp left, ScheduleAssignPublicIp right) => !left.Equals(right);

        public static explicit operator string(ScheduleAssignPublicIp value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScheduleAssignPublicIp other && Equals(other);
        public bool Equals(ScheduleAssignPublicIp other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Determines whether the schedule is executed within a flexible time window.
    /// </summary>
    [EnumType]
    public readonly struct ScheduleFlexibleTimeWindowMode : IEquatable<ScheduleFlexibleTimeWindowMode>
    {
        private readonly string _value;

        private ScheduleFlexibleTimeWindowMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScheduleFlexibleTimeWindowMode Off { get; } = new ScheduleFlexibleTimeWindowMode("OFF");
        public static ScheduleFlexibleTimeWindowMode Flexible { get; } = new ScheduleFlexibleTimeWindowMode("FLEXIBLE");

        public static bool operator ==(ScheduleFlexibleTimeWindowMode left, ScheduleFlexibleTimeWindowMode right) => left.Equals(right);
        public static bool operator !=(ScheduleFlexibleTimeWindowMode left, ScheduleFlexibleTimeWindowMode right) => !left.Equals(right);

        public static explicit operator string(ScheduleFlexibleTimeWindowMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScheduleFlexibleTimeWindowMode other && Equals(other);
        public bool Equals(ScheduleFlexibleTimeWindowMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Specifies the state of the schedule group.
    /// </summary>
    [EnumType]
    public readonly struct ScheduleGroupState : IEquatable<ScheduleGroupState>
    {
        private readonly string _value;

        private ScheduleGroupState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScheduleGroupState Active { get; } = new ScheduleGroupState("ACTIVE");
        public static ScheduleGroupState Deleting { get; } = new ScheduleGroupState("DELETING");

        public static bool operator ==(ScheduleGroupState left, ScheduleGroupState right) => left.Equals(right);
        public static bool operator !=(ScheduleGroupState left, ScheduleGroupState right) => !left.Equals(right);

        public static explicit operator string(ScheduleGroupState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScheduleGroupState other && Equals(other);
        public bool Equals(ScheduleGroupState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Specifies the launch type on which your task is running. The launch type that you specify here must match one of the launch type (compatibilities) of the target task. The FARGATE value is supported only in the Regions where AWS Fargate with Amazon ECS is supported. For more information, see AWS Fargate on Amazon ECS in the Amazon Elastic Container Service Developer Guide.
    /// </summary>
    [EnumType]
    public readonly struct ScheduleLaunchType : IEquatable<ScheduleLaunchType>
    {
        private readonly string _value;

        private ScheduleLaunchType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScheduleLaunchType Ec2 { get; } = new ScheduleLaunchType("EC2");
        public static ScheduleLaunchType Fargate { get; } = new ScheduleLaunchType("FARGATE");
        public static ScheduleLaunchType External { get; } = new ScheduleLaunchType("EXTERNAL");

        public static bool operator ==(ScheduleLaunchType left, ScheduleLaunchType right) => left.Equals(right);
        public static bool operator !=(ScheduleLaunchType left, ScheduleLaunchType right) => !left.Equals(right);

        public static explicit operator string(ScheduleLaunchType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScheduleLaunchType other && Equals(other);
        public bool Equals(ScheduleLaunchType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of constraint. Use distinctInstance to ensure that each task in a particular group is running on a different container instance. Use memberOf to restrict the selection to a group of valid candidates.
    /// </summary>
    [EnumType]
    public readonly struct SchedulePlacementConstraintType : IEquatable<SchedulePlacementConstraintType>
    {
        private readonly string _value;

        private SchedulePlacementConstraintType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SchedulePlacementConstraintType DistinctInstance { get; } = new SchedulePlacementConstraintType("distinctInstance");
        public static SchedulePlacementConstraintType MemberOf { get; } = new SchedulePlacementConstraintType("memberOf");

        public static bool operator ==(SchedulePlacementConstraintType left, SchedulePlacementConstraintType right) => left.Equals(right);
        public static bool operator !=(SchedulePlacementConstraintType left, SchedulePlacementConstraintType right) => !left.Equals(right);

        public static explicit operator string(SchedulePlacementConstraintType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SchedulePlacementConstraintType other && Equals(other);
        public bool Equals(SchedulePlacementConstraintType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of placement strategy. The random placement strategy randomly places tasks on available candidates. The spread placement strategy spreads placement across available candidates evenly based on the field parameter. The binpack strategy places tasks on available candidates that have the least available amount of the resource that is specified with the field parameter. For example, if you binpack on memory, a task is placed on the instance with the least amount of remaining memory (but still enough to run the task).
    /// </summary>
    [EnumType]
    public readonly struct SchedulePlacementStrategyType : IEquatable<SchedulePlacementStrategyType>
    {
        private readonly string _value;

        private SchedulePlacementStrategyType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SchedulePlacementStrategyType Random { get; } = new SchedulePlacementStrategyType("random");
        public static SchedulePlacementStrategyType Spread { get; } = new SchedulePlacementStrategyType("spread");
        public static SchedulePlacementStrategyType Binpack { get; } = new SchedulePlacementStrategyType("binpack");

        public static bool operator ==(SchedulePlacementStrategyType left, SchedulePlacementStrategyType right) => left.Equals(right);
        public static bool operator !=(SchedulePlacementStrategyType left, SchedulePlacementStrategyType right) => !left.Equals(right);

        public static explicit operator string(SchedulePlacementStrategyType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SchedulePlacementStrategyType other && Equals(other);
        public bool Equals(SchedulePlacementStrategyType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Specifies whether to propagate the tags from the task definition to the task. If no value is specified, the tags are not propagated. Tags can only be propagated to the task during task creation. To add tags to a task after task creation, use the TagResource API action.
    /// </summary>
    [EnumType]
    public readonly struct SchedulePropagateTags : IEquatable<SchedulePropagateTags>
    {
        private readonly string _value;

        private SchedulePropagateTags(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SchedulePropagateTags TaskDefinition { get; } = new SchedulePropagateTags("TASK_DEFINITION");

        public static bool operator ==(SchedulePropagateTags left, SchedulePropagateTags right) => left.Equals(right);
        public static bool operator !=(SchedulePropagateTags left, SchedulePropagateTags right) => !left.Equals(right);

        public static explicit operator string(SchedulePropagateTags value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SchedulePropagateTags other && Equals(other);
        public bool Equals(SchedulePropagateTags other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Specifies whether the schedule is enabled or disabled.
    /// </summary>
    [EnumType]
    public readonly struct ScheduleState : IEquatable<ScheduleState>
    {
        private readonly string _value;

        private ScheduleState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScheduleState Enabled { get; } = new ScheduleState("ENABLED");
        public static ScheduleState Disabled { get; } = new ScheduleState("DISABLED");

        public static bool operator ==(ScheduleState left, ScheduleState right) => left.Equals(right);
        public static bool operator !=(ScheduleState left, ScheduleState right) => !left.Equals(right);

        public static explicit operator string(ScheduleState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScheduleState other && Equals(other);
        public bool Equals(ScheduleState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}