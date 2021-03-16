# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'TableBillingMode',
    'TableClusteringKeyColumn',
    'TableColumn',
    'TableProvisionedThroughput',
]

@pulumi.output_type
class TableBillingMode(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-billingmode.html
    """
    def __init__(__self__, *,
                 mode: str,
                 provisioned_throughput: Optional['outputs.TableProvisionedThroughput'] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-billingmode.html
        :param str mode: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-billingmode.html#cfn-cassandra-table-billingmode-mode
        :param 'TableProvisionedThroughputArgs' provisioned_throughput: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-billingmode.html#cfn-cassandra-table-billingmode-provisionedthroughput
        """
        pulumi.set(__self__, "mode", mode)
        if provisioned_throughput is not None:
            pulumi.set(__self__, "provisioned_throughput", provisioned_throughput)

    @property
    @pulumi.getter
    def mode(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-billingmode.html#cfn-cassandra-table-billingmode-mode
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter(name="provisionedThroughput")
    def provisioned_throughput(self) -> Optional['outputs.TableProvisionedThroughput']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-billingmode.html#cfn-cassandra-table-billingmode-provisionedthroughput
        """
        return pulumi.get(self, "provisioned_throughput")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TableClusteringKeyColumn(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-clusteringkeycolumn.html
    """
    def __init__(__self__, *,
                 column: 'outputs.TableColumn',
                 order_by: Optional[str] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-clusteringkeycolumn.html
        :param 'TableColumnArgs' column: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-clusteringkeycolumn.html#cfn-cassandra-table-clusteringkeycolumn-column
        :param str order_by: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-clusteringkeycolumn.html#cfn-cassandra-table-clusteringkeycolumn-orderby
        """
        pulumi.set(__self__, "column", column)
        if order_by is not None:
            pulumi.set(__self__, "order_by", order_by)

    @property
    @pulumi.getter
    def column(self) -> 'outputs.TableColumn':
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-clusteringkeycolumn.html#cfn-cassandra-table-clusteringkeycolumn-column
        """
        return pulumi.get(self, "column")

    @property
    @pulumi.getter(name="orderBy")
    def order_by(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-clusteringkeycolumn.html#cfn-cassandra-table-clusteringkeycolumn-orderby
        """
        return pulumi.get(self, "order_by")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TableColumn(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-column.html
    """
    def __init__(__self__, *,
                 column_name: str,
                 column_type: str):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-column.html
        :param str column_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-column.html#cfn-cassandra-table-column-columnname
        :param str column_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-column.html#cfn-cassandra-table-column-columntype
        """
        pulumi.set(__self__, "column_name", column_name)
        pulumi.set(__self__, "column_type", column_type)

    @property
    @pulumi.getter(name="columnName")
    def column_name(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-column.html#cfn-cassandra-table-column-columnname
        """
        return pulumi.get(self, "column_name")

    @property
    @pulumi.getter(name="columnType")
    def column_type(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-column.html#cfn-cassandra-table-column-columntype
        """
        return pulumi.get(self, "column_type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TableProvisionedThroughput(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-provisionedthroughput.html
    """
    def __init__(__self__, *,
                 read_capacity_units: int,
                 write_capacity_units: int):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-provisionedthroughput.html
        :param int read_capacity_units: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-provisionedthroughput.html#cfn-cassandra-table-provisionedthroughput-readcapacityunits
        :param int write_capacity_units: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-provisionedthroughput.html#cfn-cassandra-table-provisionedthroughput-writecapacityunits
        """
        pulumi.set(__self__, "read_capacity_units", read_capacity_units)
        pulumi.set(__self__, "write_capacity_units", write_capacity_units)

    @property
    @pulumi.getter(name="readCapacityUnits")
    def read_capacity_units(self) -> int:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-provisionedthroughput.html#cfn-cassandra-table-provisionedthroughput-readcapacityunits
        """
        return pulumi.get(self, "read_capacity_units")

    @property
    @pulumi.getter(name="writeCapacityUnits")
    def write_capacity_units(self) -> int:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-provisionedthroughput.html#cfn-cassandra-table-provisionedthroughput-writecapacityunits
        """
        return pulumi.get(self, "write_capacity_units")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

