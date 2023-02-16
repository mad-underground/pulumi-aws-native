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
    'GetTableResult',
    'AwaitableGetTableResult',
    'get_table',
    'get_table_output',
]

@pulumi.output_type
class GetTableResult:
    def __init__(__self__, arn=None, magnetic_store_write_properties=None, name=None, retention_properties=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if magnetic_store_write_properties and not isinstance(magnetic_store_write_properties, dict):
            raise TypeError("Expected argument 'magnetic_store_write_properties' to be a dict")
        pulumi.set(__self__, "magnetic_store_write_properties", magnetic_store_write_properties)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if retention_properties and not isinstance(retention_properties, dict):
            raise TypeError("Expected argument 'retention_properties' to be a dict")
        pulumi.set(__self__, "retention_properties", retention_properties)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="magneticStoreWriteProperties")
    def magnetic_store_write_properties(self) -> Optional['outputs.MagneticStoreWritePropertiesProperties']:
        """
        The properties that determine whether magnetic store writes are enabled.
        """
        return pulumi.get(self, "magnetic_store_write_properties")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The table name exposed as a read-only attribute.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="retentionProperties")
    def retention_properties(self) -> Optional['outputs.RetentionPropertiesProperties']:
        """
        The retention duration of the memory store and the magnetic store.
        """
        return pulumi.get(self, "retention_properties")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.TableTag']]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")


class AwaitableGetTableResult(GetTableResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTableResult(
            arn=self.arn,
            magnetic_store_write_properties=self.magnetic_store_write_properties,
            name=self.name,
            retention_properties=self.retention_properties,
            tags=self.tags)


def get_table(database_name: Optional[str] = None,
              table_name: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTableResult:
    """
    The AWS::Timestream::Table resource creates a Timestream Table.


    :param str database_name: The name for the database which the table to be created belongs to.
    :param str table_name: The name for the table. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the table name.
    """
    __args__ = dict()
    __args__['databaseName'] = database_name
    __args__['tableName'] = table_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:timestream:getTable', __args__, opts=opts, typ=GetTableResult).value

    return AwaitableGetTableResult(
        arn=__ret__.arn,
        magnetic_store_write_properties=__ret__.magnetic_store_write_properties,
        name=__ret__.name,
        retention_properties=__ret__.retention_properties,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_table)
def get_table_output(database_name: Optional[pulumi.Input[str]] = None,
                     table_name: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTableResult]:
    """
    The AWS::Timestream::Table resource creates a Timestream Table.


    :param str database_name: The name for the database which the table to be created belongs to.
    :param str table_name: The name for the table. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the table name.
    """
    ...
