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
    'GetMemberResult',
    'AwaitableGetMemberResult',
    'get_member',
    'get_member_output',
]

@pulumi.output_type
class GetMemberResult:
    def __init__(__self__, email=None, status=None):
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def email(self) -> Optional[str]:
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")


class AwaitableGetMemberResult(GetMemberResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMemberResult(
            email=self.email,
            status=self.status)


def get_member(detector_id: Optional[str] = None,
               member_id: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMemberResult:
    """
    Resource Type definition for AWS::GuardDuty::Member
    """
    __args__ = dict()
    __args__['detectorId'] = detector_id
    __args__['memberId'] = member_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:guardduty:getMember', __args__, opts=opts, typ=GetMemberResult).value

    return AwaitableGetMemberResult(
        email=pulumi.get(__ret__, 'email'),
        status=pulumi.get(__ret__, 'status'))


@_utilities.lift_output_func(get_member)
def get_member_output(detector_id: Optional[pulumi.Input[str]] = None,
                      member_id: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMemberResult]:
    """
    Resource Type definition for AWS::GuardDuty::Member
    """
    ...
