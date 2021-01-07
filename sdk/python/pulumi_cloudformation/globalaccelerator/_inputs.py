# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from .. import _inputs as _root_inputs

__all__ = [
    'AcceleratorPropertiesArgs',
    'EndpointGroupEndpointConfigurationArgs',
    'EndpointGroupPortOverrideArgs',
    'EndpointGroupPropertiesArgs',
    'ListenerPortRangeArgs',
    'ListenerPropertiesArgs',
]

@pulumi.input_type
class AcceleratorPropertiesArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 enabled: Optional[pulumi.Input[bool]] = None,
                 ip_address_type: Optional[pulumi.Input[str]] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-accelerator.html
        :param pulumi.Input[str] name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-accelerator.html#cfn-globalaccelerator-accelerator-name
        :param pulumi.Input[bool] enabled: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-accelerator.html#cfn-globalaccelerator-accelerator-enabled
        :param pulumi.Input[str] ip_address_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-accelerator.html#cfn-globalaccelerator-accelerator-ipaddresstype
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-accelerator.html#cfn-globalaccelerator-accelerator-ipaddresses
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-accelerator.html#cfn-globalaccelerator-accelerator-tags
        """
        pulumi.set(__self__, "name", name)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if ip_address_type is not None:
            pulumi.set(__self__, "ip_address_type", ip_address_type)
        if ip_addresses is not None:
            pulumi.set(__self__, "ip_addresses", ip_addresses)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="Name")
    def name(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-accelerator.html#cfn-globalaccelerator-accelerator-name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="Enabled")
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-accelerator.html#cfn-globalaccelerator-accelerator-enabled
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="IpAddressType")
    def ip_address_type(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-accelerator.html#cfn-globalaccelerator-accelerator-ipaddresstype
        """
        return pulumi.get(self, "ip_address_type")

    @ip_address_type.setter
    def ip_address_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address_type", value)

    @property
    @pulumi.getter(name="IpAddresses")
    def ip_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-accelerator.html#cfn-globalaccelerator-accelerator-ipaddresses
        """
        return pulumi.get(self, "ip_addresses")

    @ip_addresses.setter
    def ip_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ip_addresses", value)

    @property
    @pulumi.getter(name="Tags")
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-accelerator.html#cfn-globalaccelerator-accelerator-tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class EndpointGroupEndpointConfigurationArgs:
    def __init__(__self__, *,
                 endpoint_id: pulumi.Input[str],
                 client_ip_preservation_enabled: Optional[pulumi.Input[bool]] = None,
                 weight: Optional[pulumi.Input[int]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-endpointgroup-endpointconfiguration.html
        :param pulumi.Input[str] endpoint_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-endpointgroup-endpointconfiguration.html#cfn-globalaccelerator-endpointgroup-endpointconfiguration-endpointid
        :param pulumi.Input[bool] client_ip_preservation_enabled: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-endpointgroup-endpointconfiguration.html#cfn-globalaccelerator-endpointgroup-endpointconfiguration-clientippreservationenabled
        :param pulumi.Input[int] weight: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-endpointgroup-endpointconfiguration.html#cfn-globalaccelerator-endpointgroup-endpointconfiguration-weight
        """
        pulumi.set(__self__, "endpoint_id", endpoint_id)
        if client_ip_preservation_enabled is not None:
            pulumi.set(__self__, "client_ip_preservation_enabled", client_ip_preservation_enabled)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter(name="EndpointId")
    def endpoint_id(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-endpointgroup-endpointconfiguration.html#cfn-globalaccelerator-endpointgroup-endpointconfiguration-endpointid
        """
        return pulumi.get(self, "endpoint_id")

    @endpoint_id.setter
    def endpoint_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint_id", value)

    @property
    @pulumi.getter(name="ClientIPPreservationEnabled")
    def client_ip_preservation_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-endpointgroup-endpointconfiguration.html#cfn-globalaccelerator-endpointgroup-endpointconfiguration-clientippreservationenabled
        """
        return pulumi.get(self, "client_ip_preservation_enabled")

    @client_ip_preservation_enabled.setter
    def client_ip_preservation_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "client_ip_preservation_enabled", value)

    @property
    @pulumi.getter(name="Weight")
    def weight(self) -> Optional[pulumi.Input[int]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-endpointgroup-endpointconfiguration.html#cfn-globalaccelerator-endpointgroup-endpointconfiguration-weight
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "weight", value)


@pulumi.input_type
class EndpointGroupPortOverrideArgs:
    def __init__(__self__, *,
                 endpoint_port: pulumi.Input[int],
                 listener_port: pulumi.Input[int]):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-endpointgroup-portoverride.html
        :param pulumi.Input[int] endpoint_port: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-endpointgroup-portoverride.html#cfn-globalaccelerator-endpointgroup-portoverride-endpointport
        :param pulumi.Input[int] listener_port: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-endpointgroup-portoverride.html#cfn-globalaccelerator-endpointgroup-portoverride-listenerport
        """
        pulumi.set(__self__, "endpoint_port", endpoint_port)
        pulumi.set(__self__, "listener_port", listener_port)

    @property
    @pulumi.getter(name="EndpointPort")
    def endpoint_port(self) -> pulumi.Input[int]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-endpointgroup-portoverride.html#cfn-globalaccelerator-endpointgroup-portoverride-endpointport
        """
        return pulumi.get(self, "endpoint_port")

    @endpoint_port.setter
    def endpoint_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "endpoint_port", value)

    @property
    @pulumi.getter(name="ListenerPort")
    def listener_port(self) -> pulumi.Input[int]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-endpointgroup-portoverride.html#cfn-globalaccelerator-endpointgroup-portoverride-listenerport
        """
        return pulumi.get(self, "listener_port")

    @listener_port.setter
    def listener_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "listener_port", value)


@pulumi.input_type
class EndpointGroupPropertiesArgs:
    def __init__(__self__, *,
                 endpoint_group_region: pulumi.Input[str],
                 listener_arn: pulumi.Input[str],
                 endpoint_configurations: Optional[pulumi.Input[Sequence[pulumi.Input['EndpointGroupEndpointConfigurationArgs']]]] = None,
                 health_check_interval_seconds: Optional[pulumi.Input[int]] = None,
                 health_check_path: Optional[pulumi.Input[str]] = None,
                 health_check_port: Optional[pulumi.Input[int]] = None,
                 health_check_protocol: Optional[pulumi.Input[str]] = None,
                 port_overrides: Optional[pulumi.Input[Sequence[pulumi.Input['EndpointGroupPortOverrideArgs']]]] = None,
                 threshold_count: Optional[pulumi.Input[int]] = None,
                 traffic_dial_percentage: Optional[pulumi.Input[float]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html
        :param pulumi.Input[str] endpoint_group_region: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-endpointgroupregion
        :param pulumi.Input[str] listener_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-listenerarn
        :param pulumi.Input[Sequence[pulumi.Input['EndpointGroupEndpointConfigurationArgs']]] endpoint_configurations: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-endpointconfigurations
        :param pulumi.Input[int] health_check_interval_seconds: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-healthcheckintervalseconds
        :param pulumi.Input[str] health_check_path: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-healthcheckpath
        :param pulumi.Input[int] health_check_port: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-healthcheckport
        :param pulumi.Input[str] health_check_protocol: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-healthcheckprotocol
        :param pulumi.Input[Sequence[pulumi.Input['EndpointGroupPortOverrideArgs']]] port_overrides: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-portoverrides
        :param pulumi.Input[int] threshold_count: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-thresholdcount
        :param pulumi.Input[float] traffic_dial_percentage: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-trafficdialpercentage
        """
        pulumi.set(__self__, "endpoint_group_region", endpoint_group_region)
        pulumi.set(__self__, "listener_arn", listener_arn)
        if endpoint_configurations is not None:
            pulumi.set(__self__, "endpoint_configurations", endpoint_configurations)
        if health_check_interval_seconds is not None:
            pulumi.set(__self__, "health_check_interval_seconds", health_check_interval_seconds)
        if health_check_path is not None:
            pulumi.set(__self__, "health_check_path", health_check_path)
        if health_check_port is not None:
            pulumi.set(__self__, "health_check_port", health_check_port)
        if health_check_protocol is not None:
            pulumi.set(__self__, "health_check_protocol", health_check_protocol)
        if port_overrides is not None:
            pulumi.set(__self__, "port_overrides", port_overrides)
        if threshold_count is not None:
            pulumi.set(__self__, "threshold_count", threshold_count)
        if traffic_dial_percentage is not None:
            pulumi.set(__self__, "traffic_dial_percentage", traffic_dial_percentage)

    @property
    @pulumi.getter(name="EndpointGroupRegion")
    def endpoint_group_region(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-endpointgroupregion
        """
        return pulumi.get(self, "endpoint_group_region")

    @endpoint_group_region.setter
    def endpoint_group_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint_group_region", value)

    @property
    @pulumi.getter(name="ListenerArn")
    def listener_arn(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-listenerarn
        """
        return pulumi.get(self, "listener_arn")

    @listener_arn.setter
    def listener_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "listener_arn", value)

    @property
    @pulumi.getter(name="EndpointConfigurations")
    def endpoint_configurations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EndpointGroupEndpointConfigurationArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-endpointconfigurations
        """
        return pulumi.get(self, "endpoint_configurations")

    @endpoint_configurations.setter
    def endpoint_configurations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EndpointGroupEndpointConfigurationArgs']]]]):
        pulumi.set(self, "endpoint_configurations", value)

    @property
    @pulumi.getter(name="HealthCheckIntervalSeconds")
    def health_check_interval_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-healthcheckintervalseconds
        """
        return pulumi.get(self, "health_check_interval_seconds")

    @health_check_interval_seconds.setter
    def health_check_interval_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "health_check_interval_seconds", value)

    @property
    @pulumi.getter(name="HealthCheckPath")
    def health_check_path(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-healthcheckpath
        """
        return pulumi.get(self, "health_check_path")

    @health_check_path.setter
    def health_check_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "health_check_path", value)

    @property
    @pulumi.getter(name="HealthCheckPort")
    def health_check_port(self) -> Optional[pulumi.Input[int]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-healthcheckport
        """
        return pulumi.get(self, "health_check_port")

    @health_check_port.setter
    def health_check_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "health_check_port", value)

    @property
    @pulumi.getter(name="HealthCheckProtocol")
    def health_check_protocol(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-healthcheckprotocol
        """
        return pulumi.get(self, "health_check_protocol")

    @health_check_protocol.setter
    def health_check_protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "health_check_protocol", value)

    @property
    @pulumi.getter(name="PortOverrides")
    def port_overrides(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EndpointGroupPortOverrideArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-portoverrides
        """
        return pulumi.get(self, "port_overrides")

    @port_overrides.setter
    def port_overrides(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EndpointGroupPortOverrideArgs']]]]):
        pulumi.set(self, "port_overrides", value)

    @property
    @pulumi.getter(name="ThresholdCount")
    def threshold_count(self) -> Optional[pulumi.Input[int]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-thresholdcount
        """
        return pulumi.get(self, "threshold_count")

    @threshold_count.setter
    def threshold_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "threshold_count", value)

    @property
    @pulumi.getter(name="TrafficDialPercentage")
    def traffic_dial_percentage(self) -> Optional[pulumi.Input[float]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-trafficdialpercentage
        """
        return pulumi.get(self, "traffic_dial_percentage")

    @traffic_dial_percentage.setter
    def traffic_dial_percentage(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "traffic_dial_percentage", value)


@pulumi.input_type
class ListenerPortRangeArgs:
    def __init__(__self__, *,
                 from_port: pulumi.Input[int],
                 to_port: pulumi.Input[int]):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-listener-portrange.html
        :param pulumi.Input[int] from_port: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-listener-portrange.html#cfn-globalaccelerator-listener-portrange-fromport
        :param pulumi.Input[int] to_port: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-listener-portrange.html#cfn-globalaccelerator-listener-portrange-toport
        """
        pulumi.set(__self__, "from_port", from_port)
        pulumi.set(__self__, "to_port", to_port)

    @property
    @pulumi.getter(name="FromPort")
    def from_port(self) -> pulumi.Input[int]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-listener-portrange.html#cfn-globalaccelerator-listener-portrange-fromport
        """
        return pulumi.get(self, "from_port")

    @from_port.setter
    def from_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "from_port", value)

    @property
    @pulumi.getter(name="ToPort")
    def to_port(self) -> pulumi.Input[int]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-listener-portrange.html#cfn-globalaccelerator-listener-portrange-toport
        """
        return pulumi.get(self, "to_port")

    @to_port.setter
    def to_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "to_port", value)


@pulumi.input_type
class ListenerPropertiesArgs:
    def __init__(__self__, *,
                 accelerator_arn: pulumi.Input[str],
                 port_ranges: pulumi.Input[Sequence[pulumi.Input['ListenerPortRangeArgs']]],
                 protocol: pulumi.Input[str],
                 client_affinity: Optional[pulumi.Input[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-listener.html
        :param pulumi.Input[str] accelerator_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-listener.html#cfn-globalaccelerator-listener-acceleratorarn
        :param pulumi.Input[Sequence[pulumi.Input['ListenerPortRangeArgs']]] port_ranges: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-listener.html#cfn-globalaccelerator-listener-portranges
        :param pulumi.Input[str] protocol: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-listener.html#cfn-globalaccelerator-listener-protocol
        :param pulumi.Input[str] client_affinity: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-listener.html#cfn-globalaccelerator-listener-clientaffinity
        """
        pulumi.set(__self__, "accelerator_arn", accelerator_arn)
        pulumi.set(__self__, "port_ranges", port_ranges)
        pulumi.set(__self__, "protocol", protocol)
        if client_affinity is not None:
            pulumi.set(__self__, "client_affinity", client_affinity)

    @property
    @pulumi.getter(name="AcceleratorArn")
    def accelerator_arn(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-listener.html#cfn-globalaccelerator-listener-acceleratorarn
        """
        return pulumi.get(self, "accelerator_arn")

    @accelerator_arn.setter
    def accelerator_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "accelerator_arn", value)

    @property
    @pulumi.getter(name="PortRanges")
    def port_ranges(self) -> pulumi.Input[Sequence[pulumi.Input['ListenerPortRangeArgs']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-listener.html#cfn-globalaccelerator-listener-portranges
        """
        return pulumi.get(self, "port_ranges")

    @port_ranges.setter
    def port_ranges(self, value: pulumi.Input[Sequence[pulumi.Input['ListenerPortRangeArgs']]]):
        pulumi.set(self, "port_ranges", value)

    @property
    @pulumi.getter(name="Protocol")
    def protocol(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-listener.html#cfn-globalaccelerator-listener-protocol
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="ClientAffinity")
    def client_affinity(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-listener.html#cfn-globalaccelerator-listener-clientaffinity
        """
        return pulumi.get(self, "client_affinity")

    @client_affinity.setter
    def client_affinity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_affinity", value)

