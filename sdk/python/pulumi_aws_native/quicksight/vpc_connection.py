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

__all__ = ['VPCConnectionArgs', 'VPCConnection']

@pulumi.input_type
class VPCConnectionArgs:
    def __init__(__self__, *,
                 availability_status: Optional[pulumi.Input['VPCConnectionAvailabilityStatus']] = None,
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 dns_resolvers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['VPCConnectionTagArgs']]]] = None,
                 v_pc_connection_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VPCConnection resource.
        """
        if availability_status is not None:
            pulumi.set(__self__, "availability_status", availability_status)
        if aws_account_id is not None:
            pulumi.set(__self__, "aws_account_id", aws_account_id)
        if dns_resolvers is not None:
            pulumi.set(__self__, "dns_resolvers", dns_resolvers)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)
        if security_group_ids is not None:
            pulumi.set(__self__, "security_group_ids", security_group_ids)
        if subnet_ids is not None:
            pulumi.set(__self__, "subnet_ids", subnet_ids)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if v_pc_connection_id is not None:
            pulumi.set(__self__, "v_pc_connection_id", v_pc_connection_id)

    @property
    @pulumi.getter(name="availabilityStatus")
    def availability_status(self) -> Optional[pulumi.Input['VPCConnectionAvailabilityStatus']]:
        return pulumi.get(self, "availability_status")

    @availability_status.setter
    def availability_status(self, value: Optional[pulumi.Input['VPCConnectionAvailabilityStatus']]):
        pulumi.set(self, "availability_status", value)

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "aws_account_id")

    @aws_account_id.setter
    def aws_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_account_id", value)

    @property
    @pulumi.getter(name="dnsResolvers")
    def dns_resolvers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "dns_resolvers")

    @dns_resolvers.setter
    def dns_resolvers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "dns_resolvers", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_group_ids", value)

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "subnet_ids")

    @subnet_ids.setter
    def subnet_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "subnet_ids", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VPCConnectionTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VPCConnectionTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vPCConnectionId")
    def v_pc_connection_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "v_pc_connection_id")

    @v_pc_connection_id.setter
    def v_pc_connection_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "v_pc_connection_id", value)


class VPCConnection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 availability_status: Optional[pulumi.Input['VPCConnectionAvailabilityStatus']] = None,
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 dns_resolvers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VPCConnectionTagArgs']]]]] = None,
                 v_pc_connection_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Definition of the AWS::QuickSight::VPCConnection Resource Type.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[VPCConnectionArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of the AWS::QuickSight::VPCConnection Resource Type.

        :param str resource_name: The name of the resource.
        :param VPCConnectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VPCConnectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 availability_status: Optional[pulumi.Input['VPCConnectionAvailabilityStatus']] = None,
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 dns_resolvers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VPCConnectionTagArgs']]]]] = None,
                 v_pc_connection_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VPCConnectionArgs.__new__(VPCConnectionArgs)

            __props__.__dict__["availability_status"] = availability_status
            __props__.__dict__["aws_account_id"] = aws_account_id
            __props__.__dict__["dns_resolvers"] = dns_resolvers
            __props__.__dict__["name"] = name
            __props__.__dict__["role_arn"] = role_arn
            __props__.__dict__["security_group_ids"] = security_group_ids
            __props__.__dict__["subnet_ids"] = subnet_ids
            __props__.__dict__["tags"] = tags
            __props__.__dict__["v_pc_connection_id"] = v_pc_connection_id
            __props__.__dict__["arn"] = None
            __props__.__dict__["created_time"] = None
            __props__.__dict__["last_updated_time"] = None
            __props__.__dict__["network_interfaces"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["v_pc_id"] = None
        super(VPCConnection, __self__).__init__(
            'aws-native:quicksight:VPCConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'VPCConnection':
        """
        Get an existing VPCConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = VPCConnectionArgs.__new__(VPCConnectionArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["availability_status"] = None
        __props__.__dict__["aws_account_id"] = None
        __props__.__dict__["created_time"] = None
        __props__.__dict__["dns_resolvers"] = None
        __props__.__dict__["last_updated_time"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["network_interfaces"] = None
        __props__.__dict__["role_arn"] = None
        __props__.__dict__["security_group_ids"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["subnet_ids"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["v_pc_connection_id"] = None
        __props__.__dict__["v_pc_id"] = None
        return VPCConnection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="availabilityStatus")
    def availability_status(self) -> pulumi.Output[Optional['VPCConnectionAvailabilityStatus']]:
        return pulumi.get(self, "availability_status")

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "aws_account_id")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter(name="dnsResolvers")
    def dns_resolvers(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "dns_resolvers")

    @property
    @pulumi.getter(name="lastUpdatedTime")
    def last_updated_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "last_updated_time")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkInterfaces")
    def network_interfaces(self) -> pulumi.Output[Sequence['outputs.VPCConnectionNetworkInterface']]:
        return pulumi.get(self, "network_interfaces")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output['VPCConnectionResourceStatus']:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.VPCConnectionTag']]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vPCConnectionId")
    def v_pc_connection_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "v_pc_connection_id")

    @property
    @pulumi.getter(name="vPCId")
    def v_pc_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "v_pc_id")

