# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Table']


class Table(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 billing_mode: Optional[pulumi.Input[pulumi.InputType['TableBillingModeArgs']]] = None,
                 clustering_key_columns: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TableClusteringKeyColumnArgs']]]]] = None,
                 keyspace_name: Optional[pulumi.Input[str]] = None,
                 partition_key_columns: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TableColumnArgs']]]]] = None,
                 regular_columns: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TableColumnArgs']]]]] = None,
                 table_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['TableBillingModeArgs']] billing_mode: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-billingmode
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TableClusteringKeyColumnArgs']]]] clustering_key_columns: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-clusteringkeycolumns
        :param pulumi.Input[str] keyspace_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-keyspacename
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TableColumnArgs']]]] partition_key_columns: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-partitionkeycolumns
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TableColumnArgs']]]] regular_columns: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-regularcolumns
        :param pulumi.Input[str] table_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-tablename
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['billing_mode'] = billing_mode
            __props__['clustering_key_columns'] = clustering_key_columns
            if keyspace_name is None and not opts.urn:
                raise TypeError("Missing required property 'keyspace_name'")
            __props__['keyspace_name'] = keyspace_name
            if partition_key_columns is None and not opts.urn:
                raise TypeError("Missing required property 'partition_key_columns'")
            __props__['partition_key_columns'] = partition_key_columns
            __props__['regular_columns'] = regular_columns
            __props__['table_name'] = table_name
        super(Table, __self__).__init__(
            'aws-native:Cassandra:Table',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Table':
        """
        Get an existing Table resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Table(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="billingMode")
    def billing_mode(self) -> pulumi.Output[Optional['outputs.TableBillingMode']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-billingmode
        """
        return pulumi.get(self, "billing_mode")

    @property
    @pulumi.getter(name="clusteringKeyColumns")
    def clustering_key_columns(self) -> pulumi.Output[Optional[Sequence['outputs.TableClusteringKeyColumn']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-clusteringkeycolumns
        """
        return pulumi.get(self, "clustering_key_columns")

    @property
    @pulumi.getter(name="keyspaceName")
    def keyspace_name(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-keyspacename
        """
        return pulumi.get(self, "keyspace_name")

    @property
    @pulumi.getter(name="partitionKeyColumns")
    def partition_key_columns(self) -> pulumi.Output[Sequence['outputs.TableColumn']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-partitionkeycolumns
        """
        return pulumi.get(self, "partition_key_columns")

    @property
    @pulumi.getter(name="regularColumns")
    def regular_columns(self) -> pulumi.Output[Optional[Sequence['outputs.TableColumn']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-regularcolumns
        """
        return pulumi.get(self, "regular_columns")

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-tablename
        """
        return pulumi.get(self, "table_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
