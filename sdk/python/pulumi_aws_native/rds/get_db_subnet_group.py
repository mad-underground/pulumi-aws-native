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
    'GetDBSubnetGroupResult',
    'AwaitableGetDBSubnetGroupResult',
    'get_db_subnet_group',
    'get_db_subnet_group_output',
]

@pulumi.output_type
class GetDBSubnetGroupResult:
    def __init__(__self__, d_b_subnet_group_description=None, tags=None):
        if d_b_subnet_group_description and not isinstance(d_b_subnet_group_description, str):
            raise TypeError("Expected argument 'd_b_subnet_group_description' to be a str")
        pulumi.set(__self__, "d_b_subnet_group_description", d_b_subnet_group_description)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="dBSubnetGroupDescription")
    def d_b_subnet_group_description(self) -> Optional[str]:
        return pulumi.get(self, "d_b_subnet_group_description")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.DBSubnetGroupTag']]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")


class AwaitableGetDBSubnetGroupResult(GetDBSubnetGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDBSubnetGroupResult(
            d_b_subnet_group_description=self.d_b_subnet_group_description,
            tags=self.tags)


def get_db_subnet_group(d_b_subnet_group_name: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDBSubnetGroupResult:
    """
    The AWS::RDS::DBSubnetGroup resource creates a database subnet group. Subnet groups must contain at least two subnets in two different Availability Zones in the same region.
    """
    __args__ = dict()
    __args__['dBSubnetGroupName'] = d_b_subnet_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:rds:getDBSubnetGroup', __args__, opts=opts, typ=GetDBSubnetGroupResult).value

    return AwaitableGetDBSubnetGroupResult(
        d_b_subnet_group_description=pulumi.get(__ret__, 'd_b_subnet_group_description'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_db_subnet_group)
def get_db_subnet_group_output(d_b_subnet_group_name: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDBSubnetGroupResult]:
    """
    The AWS::RDS::DBSubnetGroup resource creates a database subnet group. Subnet groups must contain at least two subnets in two different Availability Zones in the same region.
    """
    ...
