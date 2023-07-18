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
    'GetWorkflowResult',
    'AwaitableGetWorkflowResult',
    'get_workflow',
    'get_workflow_output',
]

@pulumi.output_type
class GetWorkflowResult:
    def __init__(__self__, arn=None, creation_time=None, description=None, id=None, name=None, status=None, tags=None, type=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if creation_time and not isinstance(creation_time, str):
            raise TypeError("Expected argument 'creation_time' to be a str")
        pulumi.set(__self__, "creation_time", creation_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[str]:
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def status(self) -> Optional['WorkflowStatus']:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional['outputs.WorkflowTagMap']:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> Optional['WorkflowType']:
        return pulumi.get(self, "type")


class AwaitableGetWorkflowResult(GetWorkflowResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWorkflowResult(
            arn=self.arn,
            creation_time=self.creation_time,
            description=self.description,
            id=self.id,
            name=self.name,
            status=self.status,
            tags=self.tags,
            type=self.type)


def get_workflow(id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWorkflowResult:
    """
    Definition of AWS::Omics::Workflow Resource Type
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:omics:getWorkflow', __args__, opts=opts, typ=GetWorkflowResult).value

    return AwaitableGetWorkflowResult(
        arn=pulumi.get(__ret__, 'arn'),
        creation_time=pulumi.get(__ret__, 'creation_time'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'))


@_utilities.lift_output_func(get_workflow)
def get_workflow_output(id: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetWorkflowResult]:
    """
    Definition of AWS::Omics::Workflow Resource Type
    """
    ...
