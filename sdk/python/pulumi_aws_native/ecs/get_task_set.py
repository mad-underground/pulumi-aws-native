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
    'GetTaskSetResult',
    'AwaitableGetTaskSetResult',
    'get_task_set',
    'get_task_set_output',
]

@pulumi.output_type
class GetTaskSetResult:
    def __init__(__self__, id=None, scale=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if scale and not isinstance(scale, dict):
            raise TypeError("Expected argument 'scale' to be a dict")
        pulumi.set(__self__, "scale", scale)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The ID of the task set.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def scale(self) -> Optional['outputs.TaskSetScale']:
        """
        A floating-point percentage of the desired number of tasks to place and keep running in the task set.
        """
        return pulumi.get(self, "scale")


class AwaitableGetTaskSetResult(GetTaskSetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTaskSetResult(
            id=self.id,
            scale=self.scale)


def get_task_set(cluster: Optional[str] = None,
                 id: Optional[str] = None,
                 service: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTaskSetResult:
    """
    Create a task set in the specified cluster and service. This is used when a service uses the EXTERNAL deployment controller type. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.htmlin the Amazon Elastic Container Service Developer Guide.


    :param str cluster: The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in.
    :param str id: The ID of the task set.
    :param str service: The short name or full Amazon Resource Name (ARN) of the service to create the task set in.
    """
    __args__ = dict()
    __args__['cluster'] = cluster
    __args__['id'] = id
    __args__['service'] = service
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:ecs:getTaskSet', __args__, opts=opts, typ=GetTaskSetResult).value

    return AwaitableGetTaskSetResult(
        id=__ret__.id,
        scale=__ret__.scale)


@_utilities.lift_output_func(get_task_set)
def get_task_set_output(cluster: Optional[pulumi.Input[str]] = None,
                        id: Optional[pulumi.Input[str]] = None,
                        service: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTaskSetResult]:
    """
    Create a task set in the specified cluster and service. This is used when a service uses the EXTERNAL deployment controller type. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.htmlin the Amazon Elastic Container Service Developer Guide.


    :param str cluster: The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in.
    :param str id: The ID of the task set.
    :param str service: The short name or full Amazon Resource Name (ARN) of the service to create the task set in.
    """
    ...
