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
    'GetScheduleGroupResult',
    'AwaitableGetScheduleGroupResult',
    'get_schedule_group',
    'get_schedule_group_output',
]

@pulumi.output_type
class GetScheduleGroupResult:
    def __init__(__self__, arn=None, creation_date=None, last_modification_date=None, state=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if creation_date and not isinstance(creation_date, str):
            raise TypeError("Expected argument 'creation_date' to be a str")
        pulumi.set(__self__, "creation_date", creation_date)
        if last_modification_date and not isinstance(last_modification_date, str):
            raise TypeError("Expected argument 'last_modification_date' to be a str")
        pulumi.set(__self__, "last_modification_date", last_modification_date)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the schedule group.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> Optional[str]:
        """
        The time at which the schedule group was created.
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter(name="lastModificationDate")
    def last_modification_date(self) -> Optional[str]:
        """
        The time at which the schedule group was last modified.
        """
        return pulumi.get(self, "last_modification_date")

    @property
    @pulumi.getter
    def state(self) -> Optional['ScheduleGroupState']:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.ScheduleGroupTag']]:
        """
        The list of tags to associate with the schedule group.
        """
        return pulumi.get(self, "tags")


class AwaitableGetScheduleGroupResult(GetScheduleGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetScheduleGroupResult(
            arn=self.arn,
            creation_date=self.creation_date,
            last_modification_date=self.last_modification_date,
            state=self.state,
            tags=self.tags)


def get_schedule_group(name: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetScheduleGroupResult:
    """
    Definition of AWS::Scheduler::ScheduleGroup Resource Type
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:scheduler:getScheduleGroup', __args__, opts=opts, typ=GetScheduleGroupResult).value

    return AwaitableGetScheduleGroupResult(
        arn=pulumi.get(__ret__, 'arn'),
        creation_date=pulumi.get(__ret__, 'creation_date'),
        last_modification_date=pulumi.get(__ret__, 'last_modification_date'),
        state=pulumi.get(__ret__, 'state'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_schedule_group)
def get_schedule_group_output(name: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetScheduleGroupResult]:
    """
    Definition of AWS::Scheduler::ScheduleGroup Resource Type
    """
    ...
