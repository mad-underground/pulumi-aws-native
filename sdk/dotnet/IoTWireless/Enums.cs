// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.IoTWireless
{
    /// <summary>
    /// Must be RuleName
    /// </summary>
    [EnumType]
    public readonly struct DestinationExpressionType : IEquatable<DestinationExpressionType>
    {
        private readonly string _value;

        private DestinationExpressionType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DestinationExpressionType RuleName { get; } = new DestinationExpressionType("RuleName");
        public static DestinationExpressionType MqttTopic { get; } = new DestinationExpressionType("MqttTopic");

        public static bool operator ==(DestinationExpressionType left, DestinationExpressionType right) => left.Equals(right);
        public static bool operator !=(DestinationExpressionType left, DestinationExpressionType right) => !left.Equals(right);

        public static explicit operator string(DestinationExpressionType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DestinationExpressionType other && Equals(other);
        public bool Equals(DestinationExpressionType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct NetworkAnalyzerConfigurationLogLevel : IEquatable<NetworkAnalyzerConfigurationLogLevel>
    {
        private readonly string _value;

        private NetworkAnalyzerConfigurationLogLevel(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static NetworkAnalyzerConfigurationLogLevel Info { get; } = new NetworkAnalyzerConfigurationLogLevel("INFO");
        public static NetworkAnalyzerConfigurationLogLevel Error { get; } = new NetworkAnalyzerConfigurationLogLevel("ERROR");
        public static NetworkAnalyzerConfigurationLogLevel Disabled { get; } = new NetworkAnalyzerConfigurationLogLevel("DISABLED");

        public static bool operator ==(NetworkAnalyzerConfigurationLogLevel left, NetworkAnalyzerConfigurationLogLevel right) => left.Equals(right);
        public static bool operator !=(NetworkAnalyzerConfigurationLogLevel left, NetworkAnalyzerConfigurationLogLevel right) => !left.Equals(right);

        public static explicit operator string(NetworkAnalyzerConfigurationLogLevel value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is NetworkAnalyzerConfigurationLogLevel other && Equals(other);
        public bool Equals(NetworkAnalyzerConfigurationLogLevel other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct NetworkAnalyzerConfigurationWirelessDeviceFrameInfo : IEquatable<NetworkAnalyzerConfigurationWirelessDeviceFrameInfo>
    {
        private readonly string _value;

        private NetworkAnalyzerConfigurationWirelessDeviceFrameInfo(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static NetworkAnalyzerConfigurationWirelessDeviceFrameInfo Enabled { get; } = new NetworkAnalyzerConfigurationWirelessDeviceFrameInfo("ENABLED");
        public static NetworkAnalyzerConfigurationWirelessDeviceFrameInfo Disabled { get; } = new NetworkAnalyzerConfigurationWirelessDeviceFrameInfo("DISABLED");

        public static bool operator ==(NetworkAnalyzerConfigurationWirelessDeviceFrameInfo left, NetworkAnalyzerConfigurationWirelessDeviceFrameInfo right) => left.Equals(right);
        public static bool operator !=(NetworkAnalyzerConfigurationWirelessDeviceFrameInfo left, NetworkAnalyzerConfigurationWirelessDeviceFrameInfo right) => !left.Equals(right);

        public static explicit operator string(NetworkAnalyzerConfigurationWirelessDeviceFrameInfo value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is NetworkAnalyzerConfigurationWirelessDeviceFrameInfo other && Equals(other);
        public bool Equals(NetworkAnalyzerConfigurationWirelessDeviceFrameInfo other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The partner type
    /// </summary>
    [EnumType]
    public readonly struct PartnerAccountPartnerType : IEquatable<PartnerAccountPartnerType>
    {
        private readonly string _value;

        private PartnerAccountPartnerType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PartnerAccountPartnerType Sidewalk { get; } = new PartnerAccountPartnerType("Sidewalk");

        public static bool operator ==(PartnerAccountPartnerType left, PartnerAccountPartnerType right) => left.Equals(right);
        public static bool operator !=(PartnerAccountPartnerType left, PartnerAccountPartnerType right) => !left.Equals(right);

        public static explicit operator string(PartnerAccountPartnerType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PartnerAccountPartnerType other && Equals(other);
        public bool Equals(PartnerAccountPartnerType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// A filter to list only the wireless gateway task definitions that use this task definition type
    /// </summary>
    [EnumType]
    public readonly struct TaskDefinitionType : IEquatable<TaskDefinitionType>
    {
        private readonly string _value;

        private TaskDefinitionType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TaskDefinitionType Update { get; } = new TaskDefinitionType("UPDATE");

        public static bool operator ==(TaskDefinitionType left, TaskDefinitionType right) => left.Equals(right);
        public static bool operator !=(TaskDefinitionType left, TaskDefinitionType right) => !left.Equals(right);

        public static explicit operator string(TaskDefinitionType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TaskDefinitionType other && Equals(other);
        public bool Equals(TaskDefinitionType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Wireless device type, currently only Sidewalk and LoRa
    /// </summary>
    [EnumType]
    public readonly struct WirelessDeviceType : IEquatable<WirelessDeviceType>
    {
        private readonly string _value;

        private WirelessDeviceType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static WirelessDeviceType Sidewalk { get; } = new WirelessDeviceType("Sidewalk");
        public static WirelessDeviceType LoRaWAN { get; } = new WirelessDeviceType("LoRaWAN");

        public static bool operator ==(WirelessDeviceType left, WirelessDeviceType right) => left.Equals(right);
        public static bool operator !=(WirelessDeviceType left, WirelessDeviceType right) => !left.Equals(right);

        public static explicit operator string(WirelessDeviceType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is WirelessDeviceType other && Equals(other);
        public bool Equals(WirelessDeviceType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
