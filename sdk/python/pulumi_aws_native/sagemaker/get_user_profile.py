# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'GetUserProfileResult',
    'AwaitableGetUserProfileResult',
    'get_user_profile',
    'get_user_profile_output',
]

@pulumi.output_type
class GetUserProfileResult:
    def __init__(__self__, user_profile_arn=None, user_settings=None):
        if user_profile_arn and not isinstance(user_profile_arn, str):
            raise TypeError("Expected argument 'user_profile_arn' to be a str")
        pulumi.set(__self__, "user_profile_arn", user_profile_arn)
        if user_settings and not isinstance(user_settings, dict):
            raise TypeError("Expected argument 'user_settings' to be a dict")
        pulumi.set(__self__, "user_settings", user_settings)

    @property
    @pulumi.getter(name="userProfileArn")
    def user_profile_arn(self) -> Optional[str]:
        """
        The user profile Amazon Resource Name (ARN).
        """
        return pulumi.get(self, "user_profile_arn")

    @property
    @pulumi.getter(name="userSettings")
    def user_settings(self) -> Optional['outputs.UserProfileUserSettings']:
        """
        A collection of settings.
        """
        return pulumi.get(self, "user_settings")


class AwaitableGetUserProfileResult(GetUserProfileResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserProfileResult(
            user_profile_arn=self.user_profile_arn,
            user_settings=self.user_settings)


def get_user_profile(domain_id: Optional[str] = None,
                     user_profile_name: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserProfileResult:
    """
    Resource Type definition for AWS::SageMaker::UserProfile


    :param str domain_id: The ID of the associated Domain.
    :param str user_profile_name: A name for the UserProfile.
    """
    __args__ = dict()
    __args__['domainId'] = domain_id
    __args__['userProfileName'] = user_profile_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:sagemaker:getUserProfile', __args__, opts=opts, typ=GetUserProfileResult).value

    return AwaitableGetUserProfileResult(
        user_profile_arn=__ret__.user_profile_arn,
        user_settings=__ret__.user_settings)


@_utilities.lift_output_func(get_user_profile)
def get_user_profile_output(domain_id: Optional[pulumi.Input[str]] = None,
                            user_profile_name: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUserProfileResult]:
    """
    Resource Type definition for AWS::SageMaker::UserProfile


    :param str domain_id: The ID of the associated Domain.
    :param str user_profile_name: A name for the UserProfile.
    """
    ...