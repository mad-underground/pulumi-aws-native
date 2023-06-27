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
    'GetServiceLinkedRoleResult',
    'AwaitableGetServiceLinkedRoleResult',
    'get_service_linked_role',
    'get_service_linked_role_output',
]

@pulumi.output_type
class GetServiceLinkedRoleResult:
    def __init__(__self__, description=None, role_name=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if role_name and not isinstance(role_name, str):
            raise TypeError("Expected argument 'role_name' to be a str")
        pulumi.set(__self__, "role_name", role_name)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the role.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> Optional[str]:
        """
        The name of the role.
        """
        return pulumi.get(self, "role_name")


class AwaitableGetServiceLinkedRoleResult(GetServiceLinkedRoleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceLinkedRoleResult(
            description=self.description,
            role_name=self.role_name)


def get_service_linked_role(role_name: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceLinkedRoleResult:
    """
    Resource Type definition for AWS::IAM::ServiceLinkedRole


    :param str role_name: The name of the role.
    """
    __args__ = dict()
    __args__['roleName'] = role_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:iam:getServiceLinkedRole', __args__, opts=opts, typ=GetServiceLinkedRoleResult).value

    return AwaitableGetServiceLinkedRoleResult(
        description=__ret__.description,
        role_name=__ret__.role_name)


@_utilities.lift_output_func(get_service_linked_role)
def get_service_linked_role_output(role_name: Optional[pulumi.Input[str]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServiceLinkedRoleResult]:
    """
    Resource Type definition for AWS::IAM::ServiceLinkedRole


    :param str role_name: The name of the role.
    """
    ...
