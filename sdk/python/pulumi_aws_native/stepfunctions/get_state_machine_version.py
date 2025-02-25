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
    'GetStateMachineVersionResult',
    'AwaitableGetStateMachineVersionResult',
    'get_state_machine_version',
    'get_state_machine_version_output',
]

@pulumi.output_type
class GetStateMachineVersionResult:
    def __init__(__self__, arn=None, description=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")


class AwaitableGetStateMachineVersionResult(GetStateMachineVersionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStateMachineVersionResult(
            arn=self.arn,
            description=self.description)


def get_state_machine_version(arn: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStateMachineVersionResult:
    """
    Resource schema for StateMachineVersion
    """
    __args__ = dict()
    __args__['arn'] = arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:stepfunctions:getStateMachineVersion', __args__, opts=opts, typ=GetStateMachineVersionResult).value

    return AwaitableGetStateMachineVersionResult(
        arn=pulumi.get(__ret__, 'arn'),
        description=pulumi.get(__ret__, 'description'))


@_utilities.lift_output_func(get_state_machine_version)
def get_state_machine_version_output(arn: Optional[pulumi.Input[str]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetStateMachineVersionResult]:
    """
    Resource schema for StateMachineVersion
    """
    ...
