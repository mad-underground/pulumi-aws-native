# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetDBParameterGroupResult',
    'AwaitableGetDBParameterGroupResult',
    'get_db_parameter_group',
    'get_db_parameter_group_output',
]

@pulumi.output_type
class GetDBParameterGroupResult:
    def __init__(__self__, d_b_parameter_group_name=None, tags=None):
        if d_b_parameter_group_name and not isinstance(d_b_parameter_group_name, str):
            raise TypeError("Expected argument 'd_b_parameter_group_name' to be a str")
        pulumi.set(__self__, "d_b_parameter_group_name", d_b_parameter_group_name)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="dBParameterGroupName")
    def d_b_parameter_group_name(self) -> Optional[str]:
        """
        Specifies the name of the DB parameter group
        """
        return pulumi.get(self, "d_b_parameter_group_name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.DBParameterGroupTag']]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")


class AwaitableGetDBParameterGroupResult(GetDBParameterGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDBParameterGroupResult(
            d_b_parameter_group_name=self.d_b_parameter_group_name,
            tags=self.tags)


def get_db_parameter_group(d_b_parameter_group_name: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDBParameterGroupResult:
    """
    The AWS::RDS::DBParameterGroup resource creates a custom parameter group for an RDS database family


    :param str d_b_parameter_group_name: Specifies the name of the DB parameter group
    """
    __args__ = dict()
    __args__['dBParameterGroupName'] = d_b_parameter_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:rds:getDBParameterGroup', __args__, opts=opts, typ=GetDBParameterGroupResult).value

    return AwaitableGetDBParameterGroupResult(
        d_b_parameter_group_name=__ret__.d_b_parameter_group_name,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_db_parameter_group)
def get_db_parameter_group_output(d_b_parameter_group_name: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDBParameterGroupResult]:
    """
    The AWS::RDS::DBParameterGroup resource creates a custom parameter group for an RDS database family


    :param str d_b_parameter_group_name: Specifies the name of the DB parameter group
    """
    ...
