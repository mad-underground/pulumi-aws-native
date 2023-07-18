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
    'GetSimpleADResult',
    'AwaitableGetSimpleADResult',
    'get_simple_ad',
    'get_simple_ad_output',
]

@pulumi.output_type
class GetSimpleADResult:
    def __init__(__self__, alias=None, directory_id=None, dns_ip_addresses=None, enable_sso=None):
        if alias and not isinstance(alias, str):
            raise TypeError("Expected argument 'alias' to be a str")
        pulumi.set(__self__, "alias", alias)
        if directory_id and not isinstance(directory_id, str):
            raise TypeError("Expected argument 'directory_id' to be a str")
        pulumi.set(__self__, "directory_id", directory_id)
        if dns_ip_addresses and not isinstance(dns_ip_addresses, list):
            raise TypeError("Expected argument 'dns_ip_addresses' to be a list")
        pulumi.set(__self__, "dns_ip_addresses", dns_ip_addresses)
        if enable_sso and not isinstance(enable_sso, bool):
            raise TypeError("Expected argument 'enable_sso' to be a bool")
        pulumi.set(__self__, "enable_sso", enable_sso)

    @property
    @pulumi.getter
    def alias(self) -> Optional[str]:
        """
        The alias for a directory.
        """
        return pulumi.get(self, "alias")

    @property
    @pulumi.getter(name="directoryId")
    def directory_id(self) -> Optional[str]:
        """
        The unique identifier for a directory.
        """
        return pulumi.get(self, "directory_id")

    @property
    @pulumi.getter(name="dnsIpAddresses")
    def dns_ip_addresses(self) -> Optional[Sequence[str]]:
        """
        The IP addresses of the DNS servers for the directory, such as [ "172.31.3.154", "172.31.63.203" ].
        """
        return pulumi.get(self, "dns_ip_addresses")

    @property
    @pulumi.getter(name="enableSso")
    def enable_sso(self) -> Optional[bool]:
        """
        Whether to enable single sign-on for a Simple Active Directory in AWS.
        """
        return pulumi.get(self, "enable_sso")


class AwaitableGetSimpleADResult(GetSimpleADResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSimpleADResult(
            alias=self.alias,
            directory_id=self.directory_id,
            dns_ip_addresses=self.dns_ip_addresses,
            enable_sso=self.enable_sso)


def get_simple_ad(directory_id: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSimpleADResult:
    """
    Resource Type definition for AWS::DirectoryService::SimpleAD


    :param str directory_id: The unique identifier for a directory.
    """
    __args__ = dict()
    __args__['directoryId'] = directory_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:directoryservice:getSimpleAD', __args__, opts=opts, typ=GetSimpleADResult).value

    return AwaitableGetSimpleADResult(
        alias=pulumi.get(__ret__, 'alias'),
        directory_id=pulumi.get(__ret__, 'directory_id'),
        dns_ip_addresses=pulumi.get(__ret__, 'dns_ip_addresses'),
        enable_sso=pulumi.get(__ret__, 'enable_sso'))


@_utilities.lift_output_func(get_simple_ad)
def get_simple_ad_output(directory_id: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSimpleADResult]:
    """
    Resource Type definition for AWS::DirectoryService::SimpleAD


    :param str directory_id: The unique identifier for a directory.
    """
    ...
