# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'GetBridgeOutputResourceResult',
    'AwaitableGetBridgeOutputResourceResult',
    'get_bridge_output_resource',
    'get_bridge_output_resource_output',
]

@pulumi.output_type
class GetBridgeOutputResourceResult:
    def __init__(__self__, network_output=None):
        if network_output and not isinstance(network_output, dict):
            raise TypeError("Expected argument 'network_output' to be a dict")
        pulumi.set(__self__, "network_output", network_output)

    @property
    @pulumi.getter(name="networkOutput")
    def network_output(self) -> Optional['outputs.BridgeOutputResourceBridgeNetworkOutput']:
        """
        The output of the bridge.
        """
        return pulumi.get(self, "network_output")


class AwaitableGetBridgeOutputResourceResult(GetBridgeOutputResourceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBridgeOutputResourceResult(
            network_output=self.network_output)


def get_bridge_output_resource(bridge_arn: Optional[str] = None,
                               name: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBridgeOutputResourceResult:
    """
    Resource schema for AWS::MediaConnect::BridgeOutput


    :param str bridge_arn: The Amazon Resource Number (ARN) of the bridge.
    :param str name: The network output name.
    """
    __args__ = dict()
    __args__['bridgeArn'] = bridge_arn
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:mediaconnect:getBridgeOutputResource', __args__, opts=opts, typ=GetBridgeOutputResourceResult).value

    return AwaitableGetBridgeOutputResourceResult(
        network_output=__ret__.network_output)


@_utilities.lift_output_func(get_bridge_output_resource)
def get_bridge_output_resource_output(bridge_arn: Optional[pulumi.Input[str]] = None,
                                      name: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBridgeOutputResourceResult]:
    """
    Resource schema for AWS::MediaConnect::BridgeOutput


    :param str bridge_arn: The Amazon Resource Number (ARN) of the bridge.
    :param str name: The network output name.
    """
    ...