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

__all__ = [
    'GetEIPResult',
    'AwaitableGetEIPResult',
    'get_eip',
    'get_eip_output',
]

@pulumi.output_type
class GetEIPResult:
    def __init__(__self__, allocation_id=None, instance_id=None, public_ip=None, public_ipv4_pool=None, tags=None):
        if allocation_id and not isinstance(allocation_id, str):
            raise TypeError("Expected argument 'allocation_id' to be a str")
        pulumi.set(__self__, "allocation_id", allocation_id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if public_ip and not isinstance(public_ip, str):
            raise TypeError("Expected argument 'public_ip' to be a str")
        pulumi.set(__self__, "public_ip", public_ip)
        if public_ipv4_pool and not isinstance(public_ipv4_pool, str):
            raise TypeError("Expected argument 'public_ipv4_pool' to be a str")
        pulumi.set(__self__, "public_ipv4_pool", public_ipv4_pool)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="allocationId")
    def allocation_id(self) -> Optional[str]:
        """
        The Allocation ID of the EIP generated by resource.
        """
        return pulumi.get(self, "allocation_id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[str]:
        """
        The ID of the instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="publicIp")
    def public_ip(self) -> Optional[str]:
        """
        The PublicIP of the EIP generated by resource.
        """
        return pulumi.get(self, "public_ip")

    @property
    @pulumi.getter(name="publicIpv4Pool")
    def public_ipv4_pool(self) -> Optional[str]:
        """
        The ID of an address pool that you own. Use this parameter to let Amazon EC2 select an address from the address pool.
        """
        return pulumi.get(self, "public_ipv4_pool")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.EIPTag']]:
        """
        Any tags assigned to the EIP.
        """
        return pulumi.get(self, "tags")


class AwaitableGetEIPResult(GetEIPResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEIPResult(
            allocation_id=self.allocation_id,
            instance_id=self.instance_id,
            public_ip=self.public_ip,
            public_ipv4_pool=self.public_ipv4_pool,
            tags=self.tags)


def get_eip(allocation_id: Optional[str] = None,
            public_ip: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEIPResult:
    """
    Resource Type definition for AWS::EC2::EIP


    :param str allocation_id: The Allocation ID of the EIP generated by resource.
    :param str public_ip: The PublicIP of the EIP generated by resource.
    """
    __args__ = dict()
    __args__['allocationId'] = allocation_id
    __args__['publicIp'] = public_ip
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:ec2:getEIP', __args__, opts=opts, typ=GetEIPResult).value

    return AwaitableGetEIPResult(
        allocation_id=pulumi.get(__ret__, 'allocation_id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        public_ip=pulumi.get(__ret__, 'public_ip'),
        public_ipv4_pool=pulumi.get(__ret__, 'public_ipv4_pool'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_eip)
def get_eip_output(allocation_id: Optional[pulumi.Input[str]] = None,
                   public_ip: Optional[pulumi.Input[str]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEIPResult]:
    """
    Resource Type definition for AWS::EC2::EIP


    :param str allocation_id: The Allocation ID of the EIP generated by resource.
    :param str public_ip: The PublicIP of the EIP generated by resource.
    """
    ...
