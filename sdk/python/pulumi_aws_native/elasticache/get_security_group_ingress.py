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
    'GetSecurityGroupIngressResult',
    'AwaitableGetSecurityGroupIngressResult',
    'get_security_group_ingress',
    'get_security_group_ingress_output',
]

@pulumi.output_type
class GetSecurityGroupIngressResult:
    def __init__(__self__, cache_security_group_name=None, e_c2_security_group_name=None, e_c2_security_group_owner_id=None, id=None):
        if cache_security_group_name and not isinstance(cache_security_group_name, str):
            raise TypeError("Expected argument 'cache_security_group_name' to be a str")
        pulumi.set(__self__, "cache_security_group_name", cache_security_group_name)
        if e_c2_security_group_name and not isinstance(e_c2_security_group_name, str):
            raise TypeError("Expected argument 'e_c2_security_group_name' to be a str")
        pulumi.set(__self__, "e_c2_security_group_name", e_c2_security_group_name)
        if e_c2_security_group_owner_id and not isinstance(e_c2_security_group_owner_id, str):
            raise TypeError("Expected argument 'e_c2_security_group_owner_id' to be a str")
        pulumi.set(__self__, "e_c2_security_group_owner_id", e_c2_security_group_owner_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="cacheSecurityGroupName")
    def cache_security_group_name(self) -> Optional[str]:
        return pulumi.get(self, "cache_security_group_name")

    @property
    @pulumi.getter(name="eC2SecurityGroupName")
    def e_c2_security_group_name(self) -> Optional[str]:
        return pulumi.get(self, "e_c2_security_group_name")

    @property
    @pulumi.getter(name="eC2SecurityGroupOwnerId")
    def e_c2_security_group_owner_id(self) -> Optional[str]:
        return pulumi.get(self, "e_c2_security_group_owner_id")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")


class AwaitableGetSecurityGroupIngressResult(GetSecurityGroupIngressResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecurityGroupIngressResult(
            cache_security_group_name=self.cache_security_group_name,
            e_c2_security_group_name=self.e_c2_security_group_name,
            e_c2_security_group_owner_id=self.e_c2_security_group_owner_id,
            id=self.id)


def get_security_group_ingress(id: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSecurityGroupIngressResult:
    """
    Resource Type definition for AWS::ElastiCache::SecurityGroupIngress
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:elasticache:getSecurityGroupIngress', __args__, opts=opts, typ=GetSecurityGroupIngressResult).value

    return AwaitableGetSecurityGroupIngressResult(
        cache_security_group_name=pulumi.get(__ret__, 'cache_security_group_name'),
        e_c2_security_group_name=pulumi.get(__ret__, 'e_c2_security_group_name'),
        e_c2_security_group_owner_id=pulumi.get(__ret__, 'e_c2_security_group_owner_id'),
        id=pulumi.get(__ret__, 'id'))


@_utilities.lift_output_func(get_security_group_ingress)
def get_security_group_ingress_output(id: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSecurityGroupIngressResult]:
    """
    Resource Type definition for AWS::ElastiCache::SecurityGroupIngress
    """
    ...
