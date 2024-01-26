# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['UserPoolUiCustomizationAttachmentArgs', 'UserPoolUiCustomizationAttachment']

@pulumi.input_type
class UserPoolUiCustomizationAttachmentArgs:
    def __init__(__self__, *,
                 client_id: pulumi.Input[str],
                 user_pool_id: pulumi.Input[str],
                 css: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a UserPoolUiCustomizationAttachment resource.
        """
        pulumi.set(__self__, "client_id", client_id)
        pulumi.set(__self__, "user_pool_id", user_pool_id)
        if css is not None:
            pulumi.set(__self__, "css", css)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="userPoolId")
    def user_pool_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "user_pool_id")

    @user_pool_id.setter
    def user_pool_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_pool_id", value)

    @property
    @pulumi.getter
    def css(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "css")

    @css.setter
    def css(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "css", value)


class UserPoolUiCustomizationAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 css: Optional[pulumi.Input[str]] = None,
                 user_pool_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::Cognito::UserPoolUICustomizationAttachment

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserPoolUiCustomizationAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::Cognito::UserPoolUICustomizationAttachment

        :param str resource_name: The name of the resource.
        :param UserPoolUiCustomizationAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserPoolUiCustomizationAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 css: Optional[pulumi.Input[str]] = None,
                 user_pool_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserPoolUiCustomizationAttachmentArgs.__new__(UserPoolUiCustomizationAttachmentArgs)

            if client_id is None and not opts.urn:
                raise TypeError("Missing required property 'client_id'")
            __props__.__dict__["client_id"] = client_id
            __props__.__dict__["css"] = css
            if user_pool_id is None and not opts.urn:
                raise TypeError("Missing required property 'user_pool_id'")
            __props__.__dict__["user_pool_id"] = user_pool_id
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["client_id", "user_pool_id"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(UserPoolUiCustomizationAttachment, __self__).__init__(
            'aws-native:cognito:UserPoolUiCustomizationAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'UserPoolUiCustomizationAttachment':
        """
        Get an existing UserPoolUiCustomizationAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = UserPoolUiCustomizationAttachmentArgs.__new__(UserPoolUiCustomizationAttachmentArgs)

        __props__.__dict__["client_id"] = None
        __props__.__dict__["css"] = None
        __props__.__dict__["user_pool_id"] = None
        return UserPoolUiCustomizationAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter
    def css(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "css")

    @property
    @pulumi.getter(name="userPoolId")
    def user_pool_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "user_pool_id")

