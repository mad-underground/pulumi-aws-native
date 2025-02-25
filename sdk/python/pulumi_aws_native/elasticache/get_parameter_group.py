# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from .. import outputs as _root_outputs

__all__ = [
    'GetParameterGroupResult',
    'AwaitableGetParameterGroupResult',
    'get_parameter_group',
    'get_parameter_group_output',
]

@pulumi.output_type
class GetParameterGroupResult:
    def __init__(__self__, description=None, id=None, properties=None, tags=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        pulumi.set(__self__, "properties", properties)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

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
    def properties(self) -> Optional[Mapping[str, str]]:
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        return pulumi.get(self, "tags")


class AwaitableGetParameterGroupResult(GetParameterGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetParameterGroupResult(
            description=self.description,
            id=self.id,
            properties=self.properties,
            tags=self.tags)


def get_parameter_group(id: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetParameterGroupResult:
    """
    Resource Type definition for AWS::ElastiCache::ParameterGroup
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:elasticache:getParameterGroup', __args__, opts=opts, typ=GetParameterGroupResult).value

    return AwaitableGetParameterGroupResult(
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        properties=pulumi.get(__ret__, 'properties'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_parameter_group)
def get_parameter_group_output(id: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetParameterGroupResult]:
    """
    Resource Type definition for AWS::ElastiCache::ParameterGroup
    """
    ...
