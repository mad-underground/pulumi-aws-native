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
    'GetOrganizationAdminResult',
    'AwaitableGetOrganizationAdminResult',
    'get_organization_admin',
    'get_organization_admin_output',
]

@pulumi.output_type
class GetOrganizationAdminResult:
    def __init__(__self__, graph_arn=None):
        if graph_arn and not isinstance(graph_arn, str):
            raise TypeError("Expected argument 'graph_arn' to be a str")
        pulumi.set(__self__, "graph_arn", graph_arn)

    @property
    @pulumi.getter(name="graphArn")
    def graph_arn(self) -> Optional[str]:
        """
        The Detective graph ARN
        """
        return pulumi.get(self, "graph_arn")


class AwaitableGetOrganizationAdminResult(GetOrganizationAdminResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOrganizationAdminResult(
            graph_arn=self.graph_arn)


def get_organization_admin(account_id: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOrganizationAdminResult:
    """
    Resource schema for AWS::Detective::OrganizationAdmin


    :param str account_id: The account ID of the account that should be registered as your Organization's delegated administrator for Detective
    """
    __args__ = dict()
    __args__['accountId'] = account_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:detective:getOrganizationAdmin', __args__, opts=opts, typ=GetOrganizationAdminResult).value

    return AwaitableGetOrganizationAdminResult(
        graph_arn=pulumi.get(__ret__, 'graph_arn'))


@_utilities.lift_output_func(get_organization_admin)
def get_organization_admin_output(account_id: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOrganizationAdminResult]:
    """
    Resource schema for AWS::Detective::OrganizationAdmin


    :param str account_id: The account ID of the account that should be registered as your Organization's delegated administrator for Detective
    """
    ...
