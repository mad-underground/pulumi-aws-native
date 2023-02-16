# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetAppBlockResult',
    'AwaitableGetAppBlockResult',
    'get_app_block',
    'get_app_block_output',
]

@pulumi.output_type
class GetAppBlockResult:
    def __init__(__self__, arn=None, created_time=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if created_time and not isinstance(created_time, str):
            raise TypeError("Expected argument 'created_time' to be a str")
        pulumi.set(__self__, "created_time", created_time)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> Optional[str]:
        return pulumi.get(self, "created_time")


class AwaitableGetAppBlockResult(GetAppBlockResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppBlockResult(
            arn=self.arn,
            created_time=self.created_time)


def get_app_block(arn: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppBlockResult:
    """
    Resource Type definition for AWS::AppStream::AppBlock
    """
    __args__ = dict()
    __args__['arn'] = arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:appstream:getAppBlock', __args__, opts=opts, typ=GetAppBlockResult).value

    return AwaitableGetAppBlockResult(
        arn=__ret__.arn,
        created_time=__ret__.created_time)


@_utilities.lift_output_func(get_app_block)
def get_app_block_output(arn: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAppBlockResult]:
    """
    Resource Type definition for AWS::AppStream::AppBlock
    """
    ...
