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
    'GetKeyGroupResult',
    'AwaitableGetKeyGroupResult',
    'get_key_group',
    'get_key_group_output',
]

@pulumi.output_type
class GetKeyGroupResult:
    def __init__(__self__, id=None, key_group_config=None, last_modified_time=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if key_group_config and not isinstance(key_group_config, dict):
            raise TypeError("Expected argument 'key_group_config' to be a dict")
        pulumi.set(__self__, "key_group_config", key_group_config)
        if last_modified_time and not isinstance(last_modified_time, str):
            raise TypeError("Expected argument 'last_modified_time' to be a str")
        pulumi.set(__self__, "last_modified_time", last_modified_time)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="keyGroupConfig")
    def key_group_config(self) -> Optional['outputs.KeyGroupConfig']:
        return pulumi.get(self, "key_group_config")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> Optional[str]:
        return pulumi.get(self, "last_modified_time")


class AwaitableGetKeyGroupResult(GetKeyGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKeyGroupResult(
            id=self.id,
            key_group_config=self.key_group_config,
            last_modified_time=self.last_modified_time)


def get_key_group(id: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetKeyGroupResult:
    """
    Resource Type definition for AWS::CloudFront::KeyGroup
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:cloudfront:getKeyGroup', __args__, opts=opts, typ=GetKeyGroupResult).value

    return AwaitableGetKeyGroupResult(
        id=__ret__.id,
        key_group_config=__ret__.key_group_config,
        last_modified_time=__ret__.last_modified_time)


@_utilities.lift_output_func(get_key_group)
def get_key_group_output(id: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetKeyGroupResult]:
    """
    Resource Type definition for AWS::CloudFront::KeyGroup
    """
    ...
