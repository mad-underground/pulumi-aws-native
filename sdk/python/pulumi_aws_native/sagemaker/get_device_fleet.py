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

__all__ = [
    'GetDeviceFleetResult',
    'AwaitableGetDeviceFleetResult',
    'get_device_fleet',
    'get_device_fleet_output',
]

@pulumi.output_type
class GetDeviceFleetResult:
    def __init__(__self__, description=None, output_config=None, role_arn=None, tags=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if output_config and not isinstance(output_config, dict):
            raise TypeError("Expected argument 'output_config' to be a dict")
        pulumi.set(__self__, "output_config", output_config)
        if role_arn and not isinstance(role_arn, str):
            raise TypeError("Expected argument 'role_arn' to be a str")
        pulumi.set(__self__, "role_arn", role_arn)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Description for the edge device fleet
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="outputConfig")
    def output_config(self) -> Optional['outputs.DeviceFleetEdgeOutputConfig']:
        """
        S3 bucket and an ecryption key id (if available) to store outputs for the fleet
        """
        return pulumi.get(self, "output_config")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[str]:
        """
        Role associated with the device fleet
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.DeviceFleetTag']]:
        """
        Associate tags with the resource
        """
        return pulumi.get(self, "tags")


class AwaitableGetDeviceFleetResult(GetDeviceFleetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDeviceFleetResult(
            description=self.description,
            output_config=self.output_config,
            role_arn=self.role_arn,
            tags=self.tags)


def get_device_fleet(device_fleet_name: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDeviceFleetResult:
    """
    Resource schema for AWS::SageMaker::DeviceFleet


    :param str device_fleet_name: The name of the edge device fleet
    """
    __args__ = dict()
    __args__['deviceFleetName'] = device_fleet_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:sagemaker:getDeviceFleet', __args__, opts=opts, typ=GetDeviceFleetResult).value

    return AwaitableGetDeviceFleetResult(
        description=pulumi.get(__ret__, 'description'),
        output_config=pulumi.get(__ret__, 'output_config'),
        role_arn=pulumi.get(__ret__, 'role_arn'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_device_fleet)
def get_device_fleet_output(device_fleet_name: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDeviceFleetResult]:
    """
    Resource schema for AWS::SageMaker::DeviceFleet


    :param str device_fleet_name: The name of the edge device fleet
    """
    ...
