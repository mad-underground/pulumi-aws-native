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
    'GetEnvironmentResult',
    'AwaitableGetEnvironmentResult',
    'get_environment',
    'get_environment_output',
]

@pulumi.output_type
class GetEnvironmentResult:
    def __init__(__self__, engine_version=None, environment_arn=None, environment_id=None, high_availability_config=None, instance_type=None, preferred_maintenance_window=None, tags=None):
        if engine_version and not isinstance(engine_version, str):
            raise TypeError("Expected argument 'engine_version' to be a str")
        pulumi.set(__self__, "engine_version", engine_version)
        if environment_arn and not isinstance(environment_arn, str):
            raise TypeError("Expected argument 'environment_arn' to be a str")
        pulumi.set(__self__, "environment_arn", environment_arn)
        if environment_id and not isinstance(environment_id, str):
            raise TypeError("Expected argument 'environment_id' to be a str")
        pulumi.set(__self__, "environment_id", environment_id)
        if high_availability_config and not isinstance(high_availability_config, dict):
            raise TypeError("Expected argument 'high_availability_config' to be a dict")
        pulumi.set(__self__, "high_availability_config", high_availability_config)
        if instance_type and not isinstance(instance_type, str):
            raise TypeError("Expected argument 'instance_type' to be a str")
        pulumi.set(__self__, "instance_type", instance_type)
        if preferred_maintenance_window and not isinstance(preferred_maintenance_window, str):
            raise TypeError("Expected argument 'preferred_maintenance_window' to be a str")
        pulumi.set(__self__, "preferred_maintenance_window", preferred_maintenance_window)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> Optional[str]:
        """
        The version of the runtime engine for the environment.
        """
        return pulumi.get(self, "engine_version")

    @property
    @pulumi.getter(name="environmentArn")
    def environment_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the runtime environment.
        """
        return pulumi.get(self, "environment_arn")

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> Optional[str]:
        """
        The unique identifier of the environment.
        """
        return pulumi.get(self, "environment_id")

    @property
    @pulumi.getter(name="highAvailabilityConfig")
    def high_availability_config(self) -> Optional['outputs.EnvironmentHighAvailabilityConfig']:
        return pulumi.get(self, "high_availability_config")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[str]:
        """
        The type of instance underlying the environment.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter(name="preferredMaintenanceWindow")
    def preferred_maintenance_window(self) -> Optional[str]:
        """
        Configures a desired maintenance window for the environment. If you do not provide a value, a random system-generated value will be assigned.
        """
        return pulumi.get(self, "preferred_maintenance_window")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Tags associated to this environment.
        """
        return pulumi.get(self, "tags")


class AwaitableGetEnvironmentResult(GetEnvironmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEnvironmentResult(
            engine_version=self.engine_version,
            environment_arn=self.environment_arn,
            environment_id=self.environment_id,
            high_availability_config=self.high_availability_config,
            instance_type=self.instance_type,
            preferred_maintenance_window=self.preferred_maintenance_window,
            tags=self.tags)


def get_environment(environment_arn: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEnvironmentResult:
    """
    Represents a runtime environment that can run migrated mainframe applications.


    :param str environment_arn: The Amazon Resource Name (ARN) of the runtime environment.
    """
    __args__ = dict()
    __args__['environmentArn'] = environment_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:m2:getEnvironment', __args__, opts=opts, typ=GetEnvironmentResult).value

    return AwaitableGetEnvironmentResult(
        engine_version=pulumi.get(__ret__, 'engine_version'),
        environment_arn=pulumi.get(__ret__, 'environment_arn'),
        environment_id=pulumi.get(__ret__, 'environment_id'),
        high_availability_config=pulumi.get(__ret__, 'high_availability_config'),
        instance_type=pulumi.get(__ret__, 'instance_type'),
        preferred_maintenance_window=pulumi.get(__ret__, 'preferred_maintenance_window'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_environment)
def get_environment_output(environment_arn: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEnvironmentResult]:
    """
    Represents a runtime environment that can run migrated mainframe applications.


    :param str environment_arn: The Amazon Resource Name (ARN) of the runtime environment.
    """
    ...
