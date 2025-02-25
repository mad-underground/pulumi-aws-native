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
    'GetViewVersionResult',
    'AwaitableGetViewVersionResult',
    'get_view_version',
    'get_view_version_output',
]

@pulumi.output_type
class GetViewVersionResult:
    def __init__(__self__, version=None, view_version_arn=None):
        if version and not isinstance(version, int):
            raise TypeError("Expected argument 'version' to be a int")
        pulumi.set(__self__, "version", version)
        if view_version_arn and not isinstance(view_version_arn, str):
            raise TypeError("Expected argument 'view_version_arn' to be a str")
        pulumi.set(__self__, "view_version_arn", view_version_arn)

    @property
    @pulumi.getter
    def version(self) -> Optional[int]:
        """
        The version of the view.
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter(name="viewVersionArn")
    def view_version_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the created view version.
        """
        return pulumi.get(self, "view_version_arn")


class AwaitableGetViewVersionResult(GetViewVersionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetViewVersionResult(
            version=self.version,
            view_version_arn=self.view_version_arn)


def get_view_version(view_version_arn: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetViewVersionResult:
    """
    Resource Type definition for AWS::Connect::ViewVersion


    :param str view_version_arn: The Amazon Resource Name (ARN) of the created view version.
    """
    __args__ = dict()
    __args__['viewVersionArn'] = view_version_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:connect:getViewVersion', __args__, opts=opts, typ=GetViewVersionResult).value

    return AwaitableGetViewVersionResult(
        version=pulumi.get(__ret__, 'version'),
        view_version_arn=pulumi.get(__ret__, 'view_version_arn'))


@_utilities.lift_output_func(get_view_version)
def get_view_version_output(view_version_arn: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetViewVersionResult]:
    """
    Resource Type definition for AWS::Connect::ViewVersion


    :param str view_version_arn: The Amazon Resource Name (ARN) of the created view version.
    """
    ...
