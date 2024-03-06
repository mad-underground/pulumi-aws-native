# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SubnetCidrBlockArgs', 'SubnetCidrBlock']

@pulumi.input_type
class SubnetCidrBlockArgs:
    def __init__(__self__, *,
                 subnet_id: pulumi.Input[str],
                 ipv6_cidr_block: Optional[pulumi.Input[str]] = None,
                 ipv6_ipam_pool_id: Optional[pulumi.Input[str]] = None,
                 ipv6_netmask_length: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a SubnetCidrBlock resource.
        :param pulumi.Input[str] subnet_id: The ID of the subnet
        :param pulumi.Input[str] ipv6_cidr_block: The IPv6 network range for the subnet, in CIDR notation. The subnet size must use a /64 prefix length
        :param pulumi.Input[str] ipv6_ipam_pool_id: The ID of an IPv6 Amazon VPC IP Address Manager (IPAM) pool from which to allocate, to get the subnet's CIDR
        :param pulumi.Input[int] ipv6_netmask_length: The netmask length of the IPv6 CIDR to allocate to the subnet from an IPAM pool
        """
        pulumi.set(__self__, "subnet_id", subnet_id)
        if ipv6_cidr_block is not None:
            pulumi.set(__self__, "ipv6_cidr_block", ipv6_cidr_block)
        if ipv6_ipam_pool_id is not None:
            pulumi.set(__self__, "ipv6_ipam_pool_id", ipv6_ipam_pool_id)
        if ipv6_netmask_length is not None:
            pulumi.set(__self__, "ipv6_netmask_length", ipv6_netmask_length)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        """
        The ID of the subnet
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="ipv6CidrBlock")
    def ipv6_cidr_block(self) -> Optional[pulumi.Input[str]]:
        """
        The IPv6 network range for the subnet, in CIDR notation. The subnet size must use a /64 prefix length
        """
        return pulumi.get(self, "ipv6_cidr_block")

    @ipv6_cidr_block.setter
    def ipv6_cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv6_cidr_block", value)

    @property
    @pulumi.getter(name="ipv6IpamPoolId")
    def ipv6_ipam_pool_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of an IPv6 Amazon VPC IP Address Manager (IPAM) pool from which to allocate, to get the subnet's CIDR
        """
        return pulumi.get(self, "ipv6_ipam_pool_id")

    @ipv6_ipam_pool_id.setter
    def ipv6_ipam_pool_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv6_ipam_pool_id", value)

    @property
    @pulumi.getter(name="ipv6NetmaskLength")
    def ipv6_netmask_length(self) -> Optional[pulumi.Input[int]]:
        """
        The netmask length of the IPv6 CIDR to allocate to the subnet from an IPAM pool
        """
        return pulumi.get(self, "ipv6_netmask_length")

    @ipv6_netmask_length.setter
    def ipv6_netmask_length(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ipv6_netmask_length", value)


class SubnetCidrBlock(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ipv6_cidr_block: Optional[pulumi.Input[str]] = None,
                 ipv6_ipam_pool_id: Optional[pulumi.Input[str]] = None,
                 ipv6_netmask_length: Optional[pulumi.Input[int]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The AWS::EC2::SubnetCidrBlock resource creates association between subnet and IPv6 CIDR

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ipv6_cidr_block: The IPv6 network range for the subnet, in CIDR notation. The subnet size must use a /64 prefix length
        :param pulumi.Input[str] ipv6_ipam_pool_id: The ID of an IPv6 Amazon VPC IP Address Manager (IPAM) pool from which to allocate, to get the subnet's CIDR
        :param pulumi.Input[int] ipv6_netmask_length: The netmask length of the IPv6 CIDR to allocate to the subnet from an IPAM pool
        :param pulumi.Input[str] subnet_id: The ID of the subnet
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SubnetCidrBlockArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The AWS::EC2::SubnetCidrBlock resource creates association between subnet and IPv6 CIDR

        :param str resource_name: The name of the resource.
        :param SubnetCidrBlockArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SubnetCidrBlockArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ipv6_cidr_block: Optional[pulumi.Input[str]] = None,
                 ipv6_ipam_pool_id: Optional[pulumi.Input[str]] = None,
                 ipv6_netmask_length: Optional[pulumi.Input[int]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SubnetCidrBlockArgs.__new__(SubnetCidrBlockArgs)

            __props__.__dict__["ipv6_cidr_block"] = ipv6_cidr_block
            __props__.__dict__["ipv6_ipam_pool_id"] = ipv6_ipam_pool_id
            __props__.__dict__["ipv6_netmask_length"] = ipv6_netmask_length
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__.__dict__["subnet_id"] = subnet_id
            __props__.__dict__["aws_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["ipv6_cidr_block", "ipv6_ipam_pool_id", "ipv6_netmask_length", "subnet_id"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(SubnetCidrBlock, __self__).__init__(
            'aws-native:ec2:SubnetCidrBlock',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'SubnetCidrBlock':
        """
        Get an existing SubnetCidrBlock resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = SubnetCidrBlockArgs.__new__(SubnetCidrBlockArgs)

        __props__.__dict__["aws_id"] = None
        __props__.__dict__["ipv6_cidr_block"] = None
        __props__.__dict__["ipv6_ipam_pool_id"] = None
        __props__.__dict__["ipv6_netmask_length"] = None
        __props__.__dict__["subnet_id"] = None
        return SubnetCidrBlock(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="awsId")
    def aws_id(self) -> pulumi.Output[str]:
        """
        Information about the IPv6 association.
        """
        return pulumi.get(self, "aws_id")

    @property
    @pulumi.getter(name="ipv6CidrBlock")
    def ipv6_cidr_block(self) -> pulumi.Output[Optional[str]]:
        """
        The IPv6 network range for the subnet, in CIDR notation. The subnet size must use a /64 prefix length
        """
        return pulumi.get(self, "ipv6_cidr_block")

    @property
    @pulumi.getter(name="ipv6IpamPoolId")
    def ipv6_ipam_pool_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of an IPv6 Amazon VPC IP Address Manager (IPAM) pool from which to allocate, to get the subnet's CIDR
        """
        return pulumi.get(self, "ipv6_ipam_pool_id")

    @property
    @pulumi.getter(name="ipv6NetmaskLength")
    def ipv6_netmask_length(self) -> pulumi.Output[Optional[int]]:
        """
        The netmask length of the IPv6 CIDR to allocate to the subnet from an IPAM pool
        """
        return pulumi.get(self, "ipv6_netmask_length")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        """
        The ID of the subnet
        """
        return pulumi.get(self, "subnet_id")

