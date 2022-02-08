# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetSimulationApplicationVersionResult',
    'AwaitableGetSimulationApplicationVersionResult',
    'get_simulation_application_version',
    'get_simulation_application_version_output',
]

@pulumi.output_type
class GetSimulationApplicationVersionResult:
    def __init__(__self__, application_version=None, arn=None):
        if application_version and not isinstance(application_version, str):
            raise TypeError("Expected argument 'application_version' to be a str")
        pulumi.set(__self__, "application_version", application_version)
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)

    @property
    @pulumi.getter(name="applicationVersion")
    def application_version(self) -> Optional[str]:
        return pulumi.get(self, "application_version")

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")


class AwaitableGetSimulationApplicationVersionResult(GetSimulationApplicationVersionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSimulationApplicationVersionResult(
            application_version=self.application_version,
            arn=self.arn)


def get_simulation_application_version(arn: Optional[str] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSimulationApplicationVersionResult:
    """
    An example resource schema demonstrating some basic constructs and validation rules.
    """
    __args__ = dict()
    __args__['arn'] = arn
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:robomaker:getSimulationApplicationVersion', __args__, opts=opts, typ=GetSimulationApplicationVersionResult).value

    return AwaitableGetSimulationApplicationVersionResult(
        application_version=__ret__.application_version,
        arn=__ret__.arn)


@_utilities.lift_output_func(get_simulation_application_version)
def get_simulation_application_version_output(arn: Optional[pulumi.Input[str]] = None,
                                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSimulationApplicationVersionResult]:
    """
    An example resource schema demonstrating some basic constructs and validation rules.
    """
    ...
