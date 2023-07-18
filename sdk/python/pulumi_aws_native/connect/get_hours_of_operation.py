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
    'GetHoursOfOperationResult',
    'AwaitableGetHoursOfOperationResult',
    'get_hours_of_operation',
    'get_hours_of_operation_output',
]

@pulumi.output_type
class GetHoursOfOperationResult:
    def __init__(__self__, config=None, description=None, hours_of_operation_arn=None, instance_arn=None, name=None, tags=None, time_zone=None):
        if config and not isinstance(config, list):
            raise TypeError("Expected argument 'config' to be a list")
        pulumi.set(__self__, "config", config)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if hours_of_operation_arn and not isinstance(hours_of_operation_arn, str):
            raise TypeError("Expected argument 'hours_of_operation_arn' to be a str")
        pulumi.set(__self__, "hours_of_operation_arn", hours_of_operation_arn)
        if instance_arn and not isinstance(instance_arn, str):
            raise TypeError("Expected argument 'instance_arn' to be a str")
        pulumi.set(__self__, "instance_arn", instance_arn)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if time_zone and not isinstance(time_zone, str):
            raise TypeError("Expected argument 'time_zone' to be a str")
        pulumi.set(__self__, "time_zone", time_zone)

    @property
    @pulumi.getter
    def config(self) -> Optional[Sequence['outputs.HoursOfOperationConfig']]:
        """
        Configuration information for the hours of operation: day, start time, and end time.
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the hours of operation.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="hoursOfOperationArn")
    def hours_of_operation_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) for the hours of operation.
        """
        return pulumi.get(self, "hours_of_operation_arn")

    @property
    @pulumi.getter(name="instanceArn")
    def instance_arn(self) -> Optional[str]:
        """
        The identifier of the Amazon Connect instance.
        """
        return pulumi.get(self, "instance_arn")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the hours of operation.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.HoursOfOperationTag']]:
        """
        One or more tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="timeZone")
    def time_zone(self) -> Optional[str]:
        """
        The time zone of the hours of operation.
        """
        return pulumi.get(self, "time_zone")


class AwaitableGetHoursOfOperationResult(GetHoursOfOperationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHoursOfOperationResult(
            config=self.config,
            description=self.description,
            hours_of_operation_arn=self.hours_of_operation_arn,
            instance_arn=self.instance_arn,
            name=self.name,
            tags=self.tags,
            time_zone=self.time_zone)


def get_hours_of_operation(hours_of_operation_arn: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHoursOfOperationResult:
    """
    Resource Type definition for AWS::Connect::HoursOfOperation


    :param str hours_of_operation_arn: The Amazon Resource Name (ARN) for the hours of operation.
    """
    __args__ = dict()
    __args__['hoursOfOperationArn'] = hours_of_operation_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:connect:getHoursOfOperation', __args__, opts=opts, typ=GetHoursOfOperationResult).value

    return AwaitableGetHoursOfOperationResult(
        config=pulumi.get(__ret__, 'config'),
        description=pulumi.get(__ret__, 'description'),
        hours_of_operation_arn=pulumi.get(__ret__, 'hours_of_operation_arn'),
        instance_arn=pulumi.get(__ret__, 'instance_arn'),
        name=pulumi.get(__ret__, 'name'),
        tags=pulumi.get(__ret__, 'tags'),
        time_zone=pulumi.get(__ret__, 'time_zone'))


@_utilities.lift_output_func(get_hours_of_operation)
def get_hours_of_operation_output(hours_of_operation_arn: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetHoursOfOperationResult]:
    """
    Resource Type definition for AWS::Connect::HoursOfOperation


    :param str hours_of_operation_arn: The Amazon Resource Name (ARN) for the hours of operation.
    """
    ...
