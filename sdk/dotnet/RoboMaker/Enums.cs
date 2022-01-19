// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.RoboMaker
{
    /// <summary>
    /// The target architecture of the robot.
    /// </summary>
    [EnumType]
    public readonly struct RobotArchitecture : IEquatable<RobotArchitecture>
    {
        private readonly string _value;

        private RobotArchitecture(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static RobotArchitecture X8664 { get; } = new RobotArchitecture("X86_64");
        public static RobotArchitecture Arm64 { get; } = new RobotArchitecture("ARM64");
        public static RobotArchitecture Armhf { get; } = new RobotArchitecture("ARMHF");

        public static bool operator ==(RobotArchitecture left, RobotArchitecture right) => left.Equals(right);
        public static bool operator !=(RobotArchitecture left, RobotArchitecture right) => !left.Equals(right);

        public static explicit operator string(RobotArchitecture value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RobotArchitecture other && Equals(other);
        public bool Equals(RobotArchitecture other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The name of the rendering engine.
    /// </summary>
    [EnumType]
    public readonly struct SimulationApplicationRenderingEngineName : IEquatable<SimulationApplicationRenderingEngineName>
    {
        private readonly string _value;

        private SimulationApplicationRenderingEngineName(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SimulationApplicationRenderingEngineName Ogre { get; } = new SimulationApplicationRenderingEngineName("OGRE");

        public static bool operator ==(SimulationApplicationRenderingEngineName left, SimulationApplicationRenderingEngineName right) => left.Equals(right);
        public static bool operator !=(SimulationApplicationRenderingEngineName left, SimulationApplicationRenderingEngineName right) => !left.Equals(right);

        public static explicit operator string(SimulationApplicationRenderingEngineName value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SimulationApplicationRenderingEngineName other && Equals(other);
        public bool Equals(SimulationApplicationRenderingEngineName other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The name of the robot software suite.
    /// </summary>
    [EnumType]
    public readonly struct SimulationApplicationRobotSoftwareSuiteName : IEquatable<SimulationApplicationRobotSoftwareSuiteName>
    {
        private readonly string _value;

        private SimulationApplicationRobotSoftwareSuiteName(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SimulationApplicationRobotSoftwareSuiteName Ros { get; } = new SimulationApplicationRobotSoftwareSuiteName("ROS");
        public static SimulationApplicationRobotSoftwareSuiteName Ros2 { get; } = new SimulationApplicationRobotSoftwareSuiteName("ROS2");
        public static SimulationApplicationRobotSoftwareSuiteName General { get; } = new SimulationApplicationRobotSoftwareSuiteName("General");

        public static bool operator ==(SimulationApplicationRobotSoftwareSuiteName left, SimulationApplicationRobotSoftwareSuiteName right) => left.Equals(right);
        public static bool operator !=(SimulationApplicationRobotSoftwareSuiteName left, SimulationApplicationRobotSoftwareSuiteName right) => !left.Equals(right);

        public static explicit operator string(SimulationApplicationRobotSoftwareSuiteName value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SimulationApplicationRobotSoftwareSuiteName other && Equals(other);
        public bool Equals(SimulationApplicationRobotSoftwareSuiteName other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The version of the robot software suite.
    /// </summary>
    [EnumType]
    public readonly struct SimulationApplicationRobotSoftwareSuiteVersion : IEquatable<SimulationApplicationRobotSoftwareSuiteVersion>
    {
        private readonly string _value;

        private SimulationApplicationRobotSoftwareSuiteVersion(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SimulationApplicationRobotSoftwareSuiteVersion Kinetic { get; } = new SimulationApplicationRobotSoftwareSuiteVersion("Kinetic");
        public static SimulationApplicationRobotSoftwareSuiteVersion Melodic { get; } = new SimulationApplicationRobotSoftwareSuiteVersion("Melodic");
        public static SimulationApplicationRobotSoftwareSuiteVersion Dashing { get; } = new SimulationApplicationRobotSoftwareSuiteVersion("Dashing");
        public static SimulationApplicationRobotSoftwareSuiteVersion Foxy { get; } = new SimulationApplicationRobotSoftwareSuiteVersion("Foxy");

        public static bool operator ==(SimulationApplicationRobotSoftwareSuiteVersion left, SimulationApplicationRobotSoftwareSuiteVersion right) => left.Equals(right);
        public static bool operator !=(SimulationApplicationRobotSoftwareSuiteVersion left, SimulationApplicationRobotSoftwareSuiteVersion right) => !left.Equals(right);

        public static explicit operator string(SimulationApplicationRobotSoftwareSuiteVersion value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SimulationApplicationRobotSoftwareSuiteVersion other && Equals(other);
        public bool Equals(SimulationApplicationRobotSoftwareSuiteVersion other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The name of the simulation software suite.
    /// </summary>
    [EnumType]
    public readonly struct SimulationApplicationSimulationSoftwareSuiteName : IEquatable<SimulationApplicationSimulationSoftwareSuiteName>
    {
        private readonly string _value;

        private SimulationApplicationSimulationSoftwareSuiteName(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SimulationApplicationSimulationSoftwareSuiteName Gazebo { get; } = new SimulationApplicationSimulationSoftwareSuiteName("Gazebo");
        public static SimulationApplicationSimulationSoftwareSuiteName RosbagPlay { get; } = new SimulationApplicationSimulationSoftwareSuiteName("RosbagPlay");
        public static SimulationApplicationSimulationSoftwareSuiteName SimulationRuntime { get; } = new SimulationApplicationSimulationSoftwareSuiteName("SimulationRuntime");

        public static bool operator ==(SimulationApplicationSimulationSoftwareSuiteName left, SimulationApplicationSimulationSoftwareSuiteName right) => left.Equals(right);
        public static bool operator !=(SimulationApplicationSimulationSoftwareSuiteName left, SimulationApplicationSimulationSoftwareSuiteName right) => !left.Equals(right);

        public static explicit operator string(SimulationApplicationSimulationSoftwareSuiteName value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SimulationApplicationSimulationSoftwareSuiteName other && Equals(other);
        public bool Equals(SimulationApplicationSimulationSoftwareSuiteName other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The version of the simulation software suite.
    /// </summary>
    [EnumType]
    public readonly struct SimulationApplicationSimulationSoftwareSuiteVersion : IEquatable<SimulationApplicationSimulationSoftwareSuiteVersion>
    {
        private readonly string _value;

        private SimulationApplicationSimulationSoftwareSuiteVersion(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SimulationApplicationSimulationSoftwareSuiteVersion Seven { get; } = new SimulationApplicationSimulationSoftwareSuiteVersion("7");
        public static SimulationApplicationSimulationSoftwareSuiteVersion Nine { get; } = new SimulationApplicationSimulationSoftwareSuiteVersion("9");
        public static SimulationApplicationSimulationSoftwareSuiteVersion SimulationApplicationSimulationSoftwareSuiteVersion_11 { get; } = new SimulationApplicationSimulationSoftwareSuiteVersion("11");
        public static SimulationApplicationSimulationSoftwareSuiteVersion Kinetic { get; } = new SimulationApplicationSimulationSoftwareSuiteVersion("Kinetic");
        public static SimulationApplicationSimulationSoftwareSuiteVersion Melodic { get; } = new SimulationApplicationSimulationSoftwareSuiteVersion("Melodic");
        public static SimulationApplicationSimulationSoftwareSuiteVersion Dashing { get; } = new SimulationApplicationSimulationSoftwareSuiteVersion("Dashing");
        public static SimulationApplicationSimulationSoftwareSuiteVersion Foxy { get; } = new SimulationApplicationSimulationSoftwareSuiteVersion("Foxy");

        public static bool operator ==(SimulationApplicationSimulationSoftwareSuiteVersion left, SimulationApplicationSimulationSoftwareSuiteVersion right) => left.Equals(right);
        public static bool operator !=(SimulationApplicationSimulationSoftwareSuiteVersion left, SimulationApplicationSimulationSoftwareSuiteVersion right) => !left.Equals(right);

        public static explicit operator string(SimulationApplicationSimulationSoftwareSuiteVersion value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SimulationApplicationSimulationSoftwareSuiteVersion other && Equals(other);
        public bool Equals(SimulationApplicationSimulationSoftwareSuiteVersion other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The target processor architecture for the application.
    /// </summary>
    [EnumType]
    public readonly struct SimulationApplicationSourceConfigArchitecture : IEquatable<SimulationApplicationSourceConfigArchitecture>
    {
        private readonly string _value;

        private SimulationApplicationSourceConfigArchitecture(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SimulationApplicationSourceConfigArchitecture X8664 { get; } = new SimulationApplicationSourceConfigArchitecture("X86_64");
        public static SimulationApplicationSourceConfigArchitecture Arm64 { get; } = new SimulationApplicationSourceConfigArchitecture("ARM64");
        public static SimulationApplicationSourceConfigArchitecture Armhf { get; } = new SimulationApplicationSourceConfigArchitecture("ARMHF");

        public static bool operator ==(SimulationApplicationSourceConfigArchitecture left, SimulationApplicationSourceConfigArchitecture right) => left.Equals(right);
        public static bool operator !=(SimulationApplicationSourceConfigArchitecture left, SimulationApplicationSourceConfigArchitecture right) => !left.Equals(right);

        public static explicit operator string(SimulationApplicationSourceConfigArchitecture value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SimulationApplicationSourceConfigArchitecture other && Equals(other);
        public bool Equals(SimulationApplicationSourceConfigArchitecture other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
