# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetResourceGroupResult',
    'AwaitableGetResourceGroupResult',
    'get_resource_group',
    'get_resource_group_output',
]

@pulumi.output_type
class GetResourceGroupResult:
    def __init__(__self__, arn=None, id=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")


class AwaitableGetResourceGroupResult(GetResourceGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResourceGroupResult(
            arn=self.arn,
            id=self.id)


def get_resource_group(id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetResourceGroupResult:
    """
    Resource Type definition for AWS::Inspector::ResourceGroup
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:inspector:getResourceGroup', __args__, opts=opts, typ=GetResourceGroupResult).value

    return AwaitableGetResourceGroupResult(
        arn=__ret__.arn,
        id=__ret__.id)


@_utilities.lift_output_func(get_resource_group)
def get_resource_group_output(id: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetResourceGroupResult]:
    """
    Resource Type definition for AWS::Inspector::ResourceGroup
    """
    ...
