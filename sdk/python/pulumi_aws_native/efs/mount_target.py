# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['MountTargetArgs', 'MountTarget']

@pulumi.input_type
class MountTargetArgs:
    def __init__(__self__, *,
                 file_system_id: pulumi.Input[str],
                 security_groups: pulumi.Input[Sequence[pulumi.Input[str]]],
                 subnet_id: pulumi.Input[str],
                 ip_address: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MountTarget resource.
        """
        pulumi.set(__self__, "file_system_id", file_system_id)
        pulumi.set(__self__, "security_groups", security_groups)
        pulumi.set(__self__, "subnet_id", subnet_id)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "file_system_id")

    @file_system_id.setter
    def file_system_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "file_system_id", value)

    @property
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "security_groups")

    @security_groups.setter
    def security_groups(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "security_groups", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)


class MountTarget(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 file_system_id: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 security_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::EFS::MountTarget

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MountTargetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::EFS::MountTarget

        :param str resource_name: The name of the resource.
        :param MountTargetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MountTargetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 file_system_id: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 security_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = MountTargetArgs.__new__(MountTargetArgs)

            if file_system_id is None and not opts.urn:
                raise TypeError("Missing required property 'file_system_id'")
            __props__.__dict__["file_system_id"] = file_system_id
            __props__.__dict__["ip_address"] = ip_address
            if security_groups is None and not opts.urn:
                raise TypeError("Missing required property 'security_groups'")
            __props__.__dict__["security_groups"] = security_groups
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__.__dict__["subnet_id"] = subnet_id
        super(MountTarget, __self__).__init__(
            'aws-native:efs:MountTarget',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'MountTarget':
        """
        Get an existing MountTarget resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = MountTargetArgs.__new__(MountTargetArgs)

        __props__.__dict__["file_system_id"] = None
        __props__.__dict__["ip_address"] = None
        __props__.__dict__["security_groups"] = None
        __props__.__dict__["subnet_id"] = None
        return MountTarget(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "file_system_id")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "security_groups")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "subnet_id")
