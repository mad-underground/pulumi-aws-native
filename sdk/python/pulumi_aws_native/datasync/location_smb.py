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
from ._enums import *
from ._inputs import *

__all__ = ['LocationSMBArgs', 'LocationSMB']

@pulumi.input_type
class LocationSMBArgs:
    def __init__(__self__, *,
                 agent_arns: pulumi.Input[Sequence[pulumi.Input[str]]],
                 user: pulumi.Input[str],
                 domain: Optional[pulumi.Input[str]] = None,
                 mount_options: Optional[pulumi.Input['LocationSMBMountOptionsArgs']] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 server_hostname: Optional[pulumi.Input[str]] = None,
                 subdirectory: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['LocationSMBTagArgs']]]] = None):
        """
        The set of arguments for constructing a LocationSMB resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] agent_arns: The Amazon Resource Names (ARNs) of agents to use for a Simple Message Block (SMB) location.
        :param pulumi.Input[str] user: The user who can mount the share, has the permissions to access files and folders in the SMB share.
        :param pulumi.Input[str] domain: The name of the Windows domain that the SMB server belongs to.
        :param pulumi.Input[str] password: The password of the user who can mount the share and has the permissions to access files and folders in the SMB share.
        :param pulumi.Input[str] server_hostname: The name of the SMB server. This value is the IP address or Domain Name Service (DNS) name of the SMB server.
        :param pulumi.Input[str] subdirectory: The subdirectory in the SMB file system that is used to read data from the SMB source location or write data to the SMB destination
        :param pulumi.Input[Sequence[pulumi.Input['LocationSMBTagArgs']]] tags: An array of key-value pairs to apply to this resource.
        """
        pulumi.set(__self__, "agent_arns", agent_arns)
        pulumi.set(__self__, "user", user)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if mount_options is not None:
            pulumi.set(__self__, "mount_options", mount_options)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if server_hostname is not None:
            pulumi.set(__self__, "server_hostname", server_hostname)
        if subdirectory is not None:
            pulumi.set(__self__, "subdirectory", subdirectory)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="agentArns")
    def agent_arns(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The Amazon Resource Names (ARNs) of agents to use for a Simple Message Block (SMB) location.
        """
        return pulumi.get(self, "agent_arns")

    @agent_arns.setter
    def agent_arns(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "agent_arns", value)

    @property
    @pulumi.getter
    def user(self) -> pulumi.Input[str]:
        """
        The user who can mount the share, has the permissions to access files and folders in the SMB share.
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: pulumi.Input[str]):
        pulumi.set(self, "user", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Windows domain that the SMB server belongs to.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="mountOptions")
    def mount_options(self) -> Optional[pulumi.Input['LocationSMBMountOptionsArgs']]:
        return pulumi.get(self, "mount_options")

    @mount_options.setter
    def mount_options(self, value: Optional[pulumi.Input['LocationSMBMountOptionsArgs']]):
        pulumi.set(self, "mount_options", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password of the user who can mount the share and has the permissions to access files and folders in the SMB share.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="serverHostname")
    def server_hostname(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the SMB server. This value is the IP address or Domain Name Service (DNS) name of the SMB server.
        """
        return pulumi.get(self, "server_hostname")

    @server_hostname.setter
    def server_hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_hostname", value)

    @property
    @pulumi.getter
    def subdirectory(self) -> Optional[pulumi.Input[str]]:
        """
        The subdirectory in the SMB file system that is used to read data from the SMB source location or write data to the SMB destination
        """
        return pulumi.get(self, "subdirectory")

    @subdirectory.setter
    def subdirectory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subdirectory", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LocationSMBTagArgs']]]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LocationSMBTagArgs']]]]):
        pulumi.set(self, "tags", value)


class LocationSMB(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agent_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 mount_options: Optional[pulumi.Input[pulumi.InputType['LocationSMBMountOptionsArgs']]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 server_hostname: Optional[pulumi.Input[str]] = None,
                 subdirectory: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LocationSMBTagArgs']]]]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource schema for AWS::DataSync::LocationSMB.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] agent_arns: The Amazon Resource Names (ARNs) of agents to use for a Simple Message Block (SMB) location.
        :param pulumi.Input[str] domain: The name of the Windows domain that the SMB server belongs to.
        :param pulumi.Input[str] password: The password of the user who can mount the share and has the permissions to access files and folders in the SMB share.
        :param pulumi.Input[str] server_hostname: The name of the SMB server. This value is the IP address or Domain Name Service (DNS) name of the SMB server.
        :param pulumi.Input[str] subdirectory: The subdirectory in the SMB file system that is used to read data from the SMB source location or write data to the SMB destination
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LocationSMBTagArgs']]]] tags: An array of key-value pairs to apply to this resource.
        :param pulumi.Input[str] user: The user who can mount the share, has the permissions to access files and folders in the SMB share.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LocationSMBArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource schema for AWS::DataSync::LocationSMB.

        :param str resource_name: The name of the resource.
        :param LocationSMBArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LocationSMBArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agent_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 mount_options: Optional[pulumi.Input[pulumi.InputType['LocationSMBMountOptionsArgs']]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 server_hostname: Optional[pulumi.Input[str]] = None,
                 subdirectory: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LocationSMBTagArgs']]]]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LocationSMBArgs.__new__(LocationSMBArgs)

            if agent_arns is None and not opts.urn:
                raise TypeError("Missing required property 'agent_arns'")
            __props__.__dict__["agent_arns"] = agent_arns
            __props__.__dict__["domain"] = domain
            __props__.__dict__["mount_options"] = mount_options
            __props__.__dict__["password"] = password
            __props__.__dict__["server_hostname"] = server_hostname
            __props__.__dict__["subdirectory"] = subdirectory
            __props__.__dict__["tags"] = tags
            if user is None and not opts.urn:
                raise TypeError("Missing required property 'user'")
            __props__.__dict__["user"] = user
            __props__.__dict__["location_arn"] = None
            __props__.__dict__["location_uri"] = None
        super(LocationSMB, __self__).__init__(
            'aws-native:datasync:LocationSMB',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'LocationSMB':
        """
        Get an existing LocationSMB resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = LocationSMBArgs.__new__(LocationSMBArgs)

        __props__.__dict__["agent_arns"] = None
        __props__.__dict__["domain"] = None
        __props__.__dict__["location_arn"] = None
        __props__.__dict__["location_uri"] = None
        __props__.__dict__["mount_options"] = None
        __props__.__dict__["password"] = None
        __props__.__dict__["server_hostname"] = None
        __props__.__dict__["subdirectory"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["user"] = None
        return LocationSMB(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="agentArns")
    def agent_arns(self) -> pulumi.Output[Sequence[str]]:
        """
        The Amazon Resource Names (ARNs) of agents to use for a Simple Message Block (SMB) location.
        """
        return pulumi.get(self, "agent_arns")

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the Windows domain that the SMB server belongs to.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="locationArn")
    def location_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the SMB location that is created.
        """
        return pulumi.get(self, "location_arn")

    @property
    @pulumi.getter(name="locationUri")
    def location_uri(self) -> pulumi.Output[str]:
        """
        The URL of the SMB location that was described.
        """
        return pulumi.get(self, "location_uri")

    @property
    @pulumi.getter(name="mountOptions")
    def mount_options(self) -> pulumi.Output[Optional['outputs.LocationSMBMountOptions']]:
        return pulumi.get(self, "mount_options")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        The password of the user who can mount the share and has the permissions to access files and folders in the SMB share.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="serverHostname")
    def server_hostname(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the SMB server. This value is the IP address or Domain Name Service (DNS) name of the SMB server.
        """
        return pulumi.get(self, "server_hostname")

    @property
    @pulumi.getter
    def subdirectory(self) -> pulumi.Output[Optional[str]]:
        """
        The subdirectory in the SMB file system that is used to read data from the SMB source location or write data to the SMB destination
        """
        return pulumi.get(self, "subdirectory")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.LocationSMBTag']]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def user(self) -> pulumi.Output[str]:
        """
        The user who can mount the share, has the permissions to access files and folders in the SMB share.
        """
        return pulumi.get(self, "user")

