# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'JobTemplateAccelerationSettingsArgs',
    'JobTemplateHopDestinationArgs',
]

@pulumi.input_type
class JobTemplateAccelerationSettingsArgs:
    def __init__(__self__, *,
                 mode: pulumi.Input[str]):
        pulumi.set(__self__, "mode", mode)

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Input[str]:
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "mode", value)


@pulumi.input_type
class JobTemplateHopDestinationArgs:
    def __init__(__self__, *,
                 priority: Optional[pulumi.Input[int]] = None,
                 queue: Optional[pulumi.Input[str]] = None,
                 wait_minutes: Optional[pulumi.Input[int]] = None):
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if queue is not None:
            pulumi.set(__self__, "queue", queue)
        if wait_minutes is not None:
            pulumi.set(__self__, "wait_minutes", wait_minutes)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def queue(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "queue")

    @queue.setter
    def queue(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "queue", value)

    @property
    @pulumi.getter(name="waitMinutes")
    def wait_minutes(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "wait_minutes")

    @wait_minutes.setter
    def wait_minutes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "wait_minutes", value)

