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
    'GetNetworkSettingsResult',
    'AwaitableGetNetworkSettingsResult',
    'get_network_settings',
    'get_network_settings_output',
]

@pulumi.output_type
class GetNetworkSettingsResult:
    def __init__(__self__, associated_portal_arns=None, network_settings_arn=None, security_group_ids=None, subnet_ids=None, tags=None, vpc_id=None):
        if associated_portal_arns and not isinstance(associated_portal_arns, list):
            raise TypeError("Expected argument 'associated_portal_arns' to be a list")
        pulumi.set(__self__, "associated_portal_arns", associated_portal_arns)
        if network_settings_arn and not isinstance(network_settings_arn, str):
            raise TypeError("Expected argument 'network_settings_arn' to be a str")
        pulumi.set(__self__, "network_settings_arn", network_settings_arn)
        if security_group_ids and not isinstance(security_group_ids, list):
            raise TypeError("Expected argument 'security_group_ids' to be a list")
        pulumi.set(__self__, "security_group_ids", security_group_ids)
        if subnet_ids and not isinstance(subnet_ids, list):
            raise TypeError("Expected argument 'subnet_ids' to be a list")
        pulumi.set(__self__, "subnet_ids", subnet_ids)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="associatedPortalArns")
    def associated_portal_arns(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "associated_portal_arns")

    @property
    @pulumi.getter(name="networkSettingsArn")
    def network_settings_arn(self) -> Optional[str]:
        return pulumi.get(self, "network_settings_arn")

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.NetworkSettingsTag']]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        return pulumi.get(self, "vpc_id")


class AwaitableGetNetworkSettingsResult(GetNetworkSettingsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkSettingsResult(
            associated_portal_arns=self.associated_portal_arns,
            network_settings_arn=self.network_settings_arn,
            security_group_ids=self.security_group_ids,
            subnet_ids=self.subnet_ids,
            tags=self.tags,
            vpc_id=self.vpc_id)


def get_network_settings(network_settings_arn: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNetworkSettingsResult:
    """
    Definition of AWS::WorkSpacesWeb::NetworkSettings Resource Type
    """
    __args__ = dict()
    __args__['networkSettingsArn'] = network_settings_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:workspacesweb:getNetworkSettings', __args__, opts=opts, typ=GetNetworkSettingsResult).value

    return AwaitableGetNetworkSettingsResult(
        associated_portal_arns=pulumi.get(__ret__, 'associated_portal_arns'),
        network_settings_arn=pulumi.get(__ret__, 'network_settings_arn'),
        security_group_ids=pulumi.get(__ret__, 'security_group_ids'),
        subnet_ids=pulumi.get(__ret__, 'subnet_ids'),
        tags=pulumi.get(__ret__, 'tags'),
        vpc_id=pulumi.get(__ret__, 'vpc_id'))


@_utilities.lift_output_func(get_network_settings)
def get_network_settings_output(network_settings_arn: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNetworkSettingsResult]:
    """
    Definition of AWS::WorkSpacesWeb::NetworkSettings Resource Type
    """
    ...