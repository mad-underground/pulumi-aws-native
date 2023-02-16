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
    'GetMacroResult',
    'AwaitableGetMacroResult',
    'get_macro',
    'get_macro_output',
]

@pulumi.output_type
class GetMacroResult:
    def __init__(__self__, description=None, function_name=None, id=None, log_group_name=None, log_role_arn=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if function_name and not isinstance(function_name, str):
            raise TypeError("Expected argument 'function_name' to be a str")
        pulumi.set(__self__, "function_name", function_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if log_group_name and not isinstance(log_group_name, str):
            raise TypeError("Expected argument 'log_group_name' to be a str")
        pulumi.set(__self__, "log_group_name", log_group_name)
        if log_role_arn and not isinstance(log_role_arn, str):
            raise TypeError("Expected argument 'log_role_arn' to be a str")
        pulumi.set(__self__, "log_role_arn", log_role_arn)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="functionName")
    def function_name(self) -> Optional[str]:
        return pulumi.get(self, "function_name")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="logGroupName")
    def log_group_name(self) -> Optional[str]:
        return pulumi.get(self, "log_group_name")

    @property
    @pulumi.getter(name="logRoleARN")
    def log_role_arn(self) -> Optional[str]:
        return pulumi.get(self, "log_role_arn")


class AwaitableGetMacroResult(GetMacroResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMacroResult(
            description=self.description,
            function_name=self.function_name,
            id=self.id,
            log_group_name=self.log_group_name,
            log_role_arn=self.log_role_arn)


def get_macro(id: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMacroResult:
    """
    Resource Type definition for AWS::CloudFormation::Macro
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:cloudformation:getMacro', __args__, opts=opts, typ=GetMacroResult).value

    return AwaitableGetMacroResult(
        description=__ret__.description,
        function_name=__ret__.function_name,
        id=__ret__.id,
        log_group_name=__ret__.log_group_name,
        log_role_arn=__ret__.log_role_arn)


@_utilities.lift_output_func(get_macro)
def get_macro_output(id: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMacroResult]:
    """
    Resource Type definition for AWS::CloudFormation::Macro
    """
    ...
