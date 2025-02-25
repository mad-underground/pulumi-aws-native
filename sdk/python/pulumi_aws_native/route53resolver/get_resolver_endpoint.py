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
    'GetResolverEndpointResult',
    'AwaitableGetResolverEndpointResult',
    'get_resolver_endpoint',
    'get_resolver_endpoint_output',
]

@pulumi.output_type
class GetResolverEndpointResult:
    def __init__(__self__, arn=None, host_vpc_id=None, ip_address_count=None, ip_addresses=None, name=None, protocols=None, resolver_endpoint_id=None, resolver_endpoint_type=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if host_vpc_id and not isinstance(host_vpc_id, str):
            raise TypeError("Expected argument 'host_vpc_id' to be a str")
        pulumi.set(__self__, "host_vpc_id", host_vpc_id)
        if ip_address_count and not isinstance(ip_address_count, str):
            raise TypeError("Expected argument 'ip_address_count' to be a str")
        pulumi.set(__self__, "ip_address_count", ip_address_count)
        if ip_addresses and not isinstance(ip_addresses, list):
            raise TypeError("Expected argument 'ip_addresses' to be a list")
        pulumi.set(__self__, "ip_addresses", ip_addresses)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if protocols and not isinstance(protocols, list):
            raise TypeError("Expected argument 'protocols' to be a list")
        pulumi.set(__self__, "protocols", protocols)
        if resolver_endpoint_id and not isinstance(resolver_endpoint_id, str):
            raise TypeError("Expected argument 'resolver_endpoint_id' to be a str")
        pulumi.set(__self__, "resolver_endpoint_id", resolver_endpoint_id)
        if resolver_endpoint_type and not isinstance(resolver_endpoint_type, str):
            raise TypeError("Expected argument 'resolver_endpoint_type' to be a str")
        pulumi.set(__self__, "resolver_endpoint_type", resolver_endpoint_type)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="hostVpcId")
    def host_vpc_id(self) -> Optional[str]:
        return pulumi.get(self, "host_vpc_id")

    @property
    @pulumi.getter(name="ipAddressCount")
    def ip_address_count(self) -> Optional[str]:
        return pulumi.get(self, "ip_address_count")

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> Optional[Sequence['outputs.ResolverEndpointIpAddressRequest']]:
        return pulumi.get(self, "ip_addresses")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def protocols(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "protocols")

    @property
    @pulumi.getter(name="resolverEndpointId")
    def resolver_endpoint_id(self) -> Optional[str]:
        return pulumi.get(self, "resolver_endpoint_id")

    @property
    @pulumi.getter(name="resolverEndpointType")
    def resolver_endpoint_type(self) -> Optional[str]:
        return pulumi.get(self, "resolver_endpoint_type")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        return pulumi.get(self, "tags")


class AwaitableGetResolverEndpointResult(GetResolverEndpointResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResolverEndpointResult(
            arn=self.arn,
            host_vpc_id=self.host_vpc_id,
            ip_address_count=self.ip_address_count,
            ip_addresses=self.ip_addresses,
            name=self.name,
            protocols=self.protocols,
            resolver_endpoint_id=self.resolver_endpoint_id,
            resolver_endpoint_type=self.resolver_endpoint_type,
            tags=self.tags)


def get_resolver_endpoint(resolver_endpoint_id: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetResolverEndpointResult:
    """
    Resource Type definition for AWS::Route53Resolver::ResolverEndpoint
    """
    __args__ = dict()
    __args__['resolverEndpointId'] = resolver_endpoint_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:route53resolver:getResolverEndpoint', __args__, opts=opts, typ=GetResolverEndpointResult).value

    return AwaitableGetResolverEndpointResult(
        arn=pulumi.get(__ret__, 'arn'),
        host_vpc_id=pulumi.get(__ret__, 'host_vpc_id'),
        ip_address_count=pulumi.get(__ret__, 'ip_address_count'),
        ip_addresses=pulumi.get(__ret__, 'ip_addresses'),
        name=pulumi.get(__ret__, 'name'),
        protocols=pulumi.get(__ret__, 'protocols'),
        resolver_endpoint_id=pulumi.get(__ret__, 'resolver_endpoint_id'),
        resolver_endpoint_type=pulumi.get(__ret__, 'resolver_endpoint_type'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_resolver_endpoint)
def get_resolver_endpoint_output(resolver_endpoint_id: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetResolverEndpointResult]:
    """
    Resource Type definition for AWS::Route53Resolver::ResolverEndpoint
    """
    ...
