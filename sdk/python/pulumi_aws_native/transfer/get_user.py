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
from .. import outputs as _root_outputs

__all__ = [
    'GetUserResult',
    'AwaitableGetUserResult',
    'get_user',
    'get_user_output',
]

@pulumi.output_type
class GetUserResult:
    def __init__(__self__, arn=None, home_directory=None, home_directory_mappings=None, home_directory_type=None, id=None, policy=None, posix_profile=None, role=None, ssh_public_keys=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if home_directory and not isinstance(home_directory, str):
            raise TypeError("Expected argument 'home_directory' to be a str")
        pulumi.set(__self__, "home_directory", home_directory)
        if home_directory_mappings and not isinstance(home_directory_mappings, list):
            raise TypeError("Expected argument 'home_directory_mappings' to be a list")
        pulumi.set(__self__, "home_directory_mappings", home_directory_mappings)
        if home_directory_type and not isinstance(home_directory_type, str):
            raise TypeError("Expected argument 'home_directory_type' to be a str")
        pulumi.set(__self__, "home_directory_type", home_directory_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if policy and not isinstance(policy, str):
            raise TypeError("Expected argument 'policy' to be a str")
        pulumi.set(__self__, "policy", policy)
        if posix_profile and not isinstance(posix_profile, dict):
            raise TypeError("Expected argument 'posix_profile' to be a dict")
        pulumi.set(__self__, "posix_profile", posix_profile)
        if role and not isinstance(role, str):
            raise TypeError("Expected argument 'role' to be a str")
        pulumi.set(__self__, "role", role)
        if ssh_public_keys and not isinstance(ssh_public_keys, list):
            raise TypeError("Expected argument 'ssh_public_keys' to be a list")
        pulumi.set(__self__, "ssh_public_keys", ssh_public_keys)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="homeDirectory")
    def home_directory(self) -> Optional[str]:
        return pulumi.get(self, "home_directory")

    @property
    @pulumi.getter(name="homeDirectoryMappings")
    def home_directory_mappings(self) -> Optional[Sequence['outputs.UserHomeDirectoryMapEntry']]:
        return pulumi.get(self, "home_directory_mappings")

    @property
    @pulumi.getter(name="homeDirectoryType")
    def home_directory_type(self) -> Optional[str]:
        return pulumi.get(self, "home_directory_type")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def policy(self) -> Optional[str]:
        return pulumi.get(self, "policy")

    @property
    @pulumi.getter(name="posixProfile")
    def posix_profile(self) -> Optional['outputs.UserPosixProfile']:
        return pulumi.get(self, "posix_profile")

    @property
    @pulumi.getter
    def role(self) -> Optional[str]:
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="sshPublicKeys")
    def ssh_public_keys(self) -> Optional[Sequence['outputs.UserSshPublicKey']]:
        return pulumi.get(self, "ssh_public_keys")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        return pulumi.get(self, "tags")


class AwaitableGetUserResult(GetUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserResult(
            arn=self.arn,
            home_directory=self.home_directory,
            home_directory_mappings=self.home_directory_mappings,
            home_directory_type=self.home_directory_type,
            id=self.id,
            policy=self.policy,
            posix_profile=self.posix_profile,
            role=self.role,
            ssh_public_keys=self.ssh_public_keys,
            tags=self.tags)


def get_user(id: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserResult:
    """
    Resource Type definition for AWS::Transfer::User
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:transfer:getUser', __args__, opts=opts, typ=GetUserResult).value

    return AwaitableGetUserResult(
        arn=pulumi.get(__ret__, 'arn'),
        home_directory=pulumi.get(__ret__, 'home_directory'),
        home_directory_mappings=pulumi.get(__ret__, 'home_directory_mappings'),
        home_directory_type=pulumi.get(__ret__, 'home_directory_type'),
        id=pulumi.get(__ret__, 'id'),
        policy=pulumi.get(__ret__, 'policy'),
        posix_profile=pulumi.get(__ret__, 'posix_profile'),
        role=pulumi.get(__ret__, 'role'),
        ssh_public_keys=pulumi.get(__ret__, 'ssh_public_keys'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_user)
def get_user_output(id: Optional[pulumi.Input[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUserResult]:
    """
    Resource Type definition for AWS::Transfer::User
    """
    ...
