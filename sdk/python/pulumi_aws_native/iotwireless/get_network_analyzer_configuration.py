# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'GetNetworkAnalyzerConfigurationResult',
    'AwaitableGetNetworkAnalyzerConfigurationResult',
    'get_network_analyzer_configuration',
    'get_network_analyzer_configuration_output',
]

@pulumi.output_type
class GetNetworkAnalyzerConfigurationResult:
    def __init__(__self__, arn=None, description=None, trace_content=None, wireless_devices=None, wireless_gateways=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if trace_content and not isinstance(trace_content, dict):
            raise TypeError("Expected argument 'trace_content' to be a dict")
        pulumi.set(__self__, "trace_content", trace_content)
        if wireless_devices and not isinstance(wireless_devices, list):
            raise TypeError("Expected argument 'wireless_devices' to be a list")
        pulumi.set(__self__, "wireless_devices", wireless_devices)
        if wireless_gateways and not isinstance(wireless_gateways, list):
            raise TypeError("Expected argument 'wireless_gateways' to be a list")
        pulumi.set(__self__, "wireless_gateways", wireless_gateways)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        Arn for network analyzer configuration, Returned upon successful create.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the new resource
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="traceContent")
    def trace_content(self) -> Optional['outputs.TraceContentProperties']:
        """
        Trace content for your wireless gateway and wireless device resources
        """
        return pulumi.get(self, "trace_content")

    @property
    @pulumi.getter(name="wirelessDevices")
    def wireless_devices(self) -> Optional[Sequence[str]]:
        """
        List of wireless gateway resources that have been added to the network analyzer configuration
        """
        return pulumi.get(self, "wireless_devices")

    @property
    @pulumi.getter(name="wirelessGateways")
    def wireless_gateways(self) -> Optional[Sequence[str]]:
        """
        List of wireless gateway resources that have been added to the network analyzer configuration
        """
        return pulumi.get(self, "wireless_gateways")


class AwaitableGetNetworkAnalyzerConfigurationResult(GetNetworkAnalyzerConfigurationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkAnalyzerConfigurationResult(
            arn=self.arn,
            description=self.description,
            trace_content=self.trace_content,
            wireless_devices=self.wireless_devices,
            wireless_gateways=self.wireless_gateways)


def get_network_analyzer_configuration(name: Optional[str] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNetworkAnalyzerConfigurationResult:
    """
    Create and manage NetworkAnalyzerConfiguration resource.


    :param str name: Name of the network analyzer configuration
    """
    __args__ = dict()
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:iotwireless:getNetworkAnalyzerConfiguration', __args__, opts=opts, typ=GetNetworkAnalyzerConfigurationResult).value

    return AwaitableGetNetworkAnalyzerConfigurationResult(
        arn=__ret__.arn,
        description=__ret__.description,
        trace_content=__ret__.trace_content,
        wireless_devices=__ret__.wireless_devices,
        wireless_gateways=__ret__.wireless_gateways)


@_utilities.lift_output_func(get_network_analyzer_configuration)
def get_network_analyzer_configuration_output(name: Optional[pulumi.Input[str]] = None,
                                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNetworkAnalyzerConfigurationResult]:
    """
    Create and manage NetworkAnalyzerConfiguration resource.


    :param str name: Name of the network analyzer configuration
    """
    ...
