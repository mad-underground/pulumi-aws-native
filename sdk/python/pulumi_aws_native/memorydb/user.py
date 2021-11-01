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
from ._inputs import *

__all__ = ['UserArgs', 'User']

@pulumi.input_type
class UserArgs:
    def __init__(__self__, *,
                 access_string: pulumi.Input[str],
                 authentication_mode: pulumi.Input['AuthenticationModePropertiesArgs'],
                 user_name: pulumi.Input[str],
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['UserTagArgs']]]] = None):
        """
        The set of arguments for constructing a User resource.
        :param pulumi.Input[str] access_string: Access permissions string used for this user account.
        :param pulumi.Input[str] user_name: The name of the user.
        :param pulumi.Input[Sequence[pulumi.Input['UserTagArgs']]] tags: An array of key-value pairs to apply to this user.
        """
        pulumi.set(__self__, "access_string", access_string)
        pulumi.set(__self__, "authentication_mode", authentication_mode)
        pulumi.set(__self__, "user_name", user_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="accessString")
    def access_string(self) -> pulumi.Input[str]:
        """
        Access permissions string used for this user account.
        """
        return pulumi.get(self, "access_string")

    @access_string.setter
    def access_string(self, value: pulumi.Input[str]):
        pulumi.set(self, "access_string", value)

    @property
    @pulumi.getter(name="authenticationMode")
    def authentication_mode(self) -> pulumi.Input['AuthenticationModePropertiesArgs']:
        return pulumi.get(self, "authentication_mode")

    @authentication_mode.setter
    def authentication_mode(self, value: pulumi.Input['AuthenticationModePropertiesArgs']):
        pulumi.set(self, "authentication_mode", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Input[str]:
        """
        The name of the user.
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UserTagArgs']]]]:
        """
        An array of key-value pairs to apply to this user.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UserTagArgs']]]]):
        pulumi.set(self, "tags", value)


class User(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_string: Optional[pulumi.Input[str]] = None,
                 authentication_mode: Optional[pulumi.Input[pulumi.InputType['AuthenticationModePropertiesArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UserTagArgs']]]]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::MemoryDB::User

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_string: Access permissions string used for this user account.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UserTagArgs']]]] tags: An array of key-value pairs to apply to this user.
        :param pulumi.Input[str] user_name: The name of the user.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::MemoryDB::User

        :param str resource_name: The name of the resource.
        :param UserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_string: Optional[pulumi.Input[str]] = None,
                 authentication_mode: Optional[pulumi.Input[pulumi.InputType['AuthenticationModePropertiesArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UserTagArgs']]]]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserArgs.__new__(UserArgs)

            if access_string is None and not opts.urn:
                raise TypeError("Missing required property 'access_string'")
            __props__.__dict__["access_string"] = access_string
            if authentication_mode is None and not opts.urn:
                raise TypeError("Missing required property 'authentication_mode'")
            __props__.__dict__["authentication_mode"] = authentication_mode
            __props__.__dict__["tags"] = tags
            if user_name is None and not opts.urn:
                raise TypeError("Missing required property 'user_name'")
            __props__.__dict__["user_name"] = user_name
            __props__.__dict__["arn"] = None
            __props__.__dict__["status"] = None
        super(User, __self__).__init__(
            'aws-native:memorydb:User',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'User':
        """
        Get an existing User resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = UserArgs.__new__(UserArgs)

        __props__.__dict__["access_string"] = None
        __props__.__dict__["arn"] = None
        __props__.__dict__["authentication_mode"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["user_name"] = None
        return User(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessString")
    def access_string(self) -> pulumi.Output[str]:
        """
        Access permissions string used for this user account.
        """
        return pulumi.get(self, "access_string")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the user account.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="authenticationMode")
    def authentication_mode(self) -> pulumi.Output['outputs.AuthenticationModeProperties']:
        return pulumi.get(self, "authentication_mode")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Indicates the user status. Can be "active", "modifying" or "deleting".
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.UserTag']]]:
        """
        An array of key-value pairs to apply to this user.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Output[str]:
        """
        The name of the user.
        """
        return pulumi.get(self, "user_name")

