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
from ._enums import *

__all__ = [
    'GetTableResult',
    'AwaitableGetTableResult',
    'get_table',
    'get_table_output',
]

@pulumi.output_type
class GetTableResult:
    def __init__(__self__, arn=None, attribute_definitions=None, billing_mode=None, contributor_insights_specification=None, deletion_protection_enabled=None, global_secondary_indexes=None, key_schema=None, kinesis_stream_specification=None, local_secondary_indexes=None, point_in_time_recovery_specification=None, provisioned_throughput=None, sse_specification=None, stream_arn=None, stream_specification=None, table_class=None, tags=None, time_to_live_specification=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if attribute_definitions and not isinstance(attribute_definitions, list):
            raise TypeError("Expected argument 'attribute_definitions' to be a list")
        pulumi.set(__self__, "attribute_definitions", attribute_definitions)
        if billing_mode and not isinstance(billing_mode, str):
            raise TypeError("Expected argument 'billing_mode' to be a str")
        pulumi.set(__self__, "billing_mode", billing_mode)
        if contributor_insights_specification and not isinstance(contributor_insights_specification, dict):
            raise TypeError("Expected argument 'contributor_insights_specification' to be a dict")
        pulumi.set(__self__, "contributor_insights_specification", contributor_insights_specification)
        if deletion_protection_enabled and not isinstance(deletion_protection_enabled, bool):
            raise TypeError("Expected argument 'deletion_protection_enabled' to be a bool")
        pulumi.set(__self__, "deletion_protection_enabled", deletion_protection_enabled)
        if global_secondary_indexes and not isinstance(global_secondary_indexes, list):
            raise TypeError("Expected argument 'global_secondary_indexes' to be a list")
        pulumi.set(__self__, "global_secondary_indexes", global_secondary_indexes)
        if key_schema and not isinstance(key_schema, dict):
            raise TypeError("Expected argument 'key_schema' to be a dict")
        pulumi.set(__self__, "key_schema", key_schema)
        if kinesis_stream_specification and not isinstance(kinesis_stream_specification, dict):
            raise TypeError("Expected argument 'kinesis_stream_specification' to be a dict")
        pulumi.set(__self__, "kinesis_stream_specification", kinesis_stream_specification)
        if local_secondary_indexes and not isinstance(local_secondary_indexes, list):
            raise TypeError("Expected argument 'local_secondary_indexes' to be a list")
        pulumi.set(__self__, "local_secondary_indexes", local_secondary_indexes)
        if point_in_time_recovery_specification and not isinstance(point_in_time_recovery_specification, dict):
            raise TypeError("Expected argument 'point_in_time_recovery_specification' to be a dict")
        pulumi.set(__self__, "point_in_time_recovery_specification", point_in_time_recovery_specification)
        if provisioned_throughput and not isinstance(provisioned_throughput, dict):
            raise TypeError("Expected argument 'provisioned_throughput' to be a dict")
        pulumi.set(__self__, "provisioned_throughput", provisioned_throughput)
        if sse_specification and not isinstance(sse_specification, dict):
            raise TypeError("Expected argument 'sse_specification' to be a dict")
        pulumi.set(__self__, "sse_specification", sse_specification)
        if stream_arn and not isinstance(stream_arn, str):
            raise TypeError("Expected argument 'stream_arn' to be a str")
        pulumi.set(__self__, "stream_arn", stream_arn)
        if stream_specification and not isinstance(stream_specification, dict):
            raise TypeError("Expected argument 'stream_specification' to be a dict")
        pulumi.set(__self__, "stream_specification", stream_specification)
        if table_class and not isinstance(table_class, str):
            raise TypeError("Expected argument 'table_class' to be a str")
        pulumi.set(__self__, "table_class", table_class)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if time_to_live_specification and not isinstance(time_to_live_specification, dict):
            raise TypeError("Expected argument 'time_to_live_specification' to be a dict")
        pulumi.set(__self__, "time_to_live_specification", time_to_live_specification)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="attributeDefinitions")
    def attribute_definitions(self) -> Optional[Sequence['outputs.TableAttributeDefinition']]:
        """
        A list of attributes that describe the key schema for the table and indexes.
         This property is required to create a DDB table.
         Update requires: [Some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt). Replacement if you edit an existing AttributeDefinition.
        """
        return pulumi.get(self, "attribute_definitions")

    @property
    @pulumi.getter(name="billingMode")
    def billing_mode(self) -> Optional[str]:
        """
        Specify how you are charged for read and write throughput and how you manage capacity.
         Valid values include:
          +   ``PROVISIONED`` - We recommend using ``PROVISIONED`` for predictable workloads. ``PROVISIONED`` sets the billing mode to [Provisioned Mode](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.ProvisionedThroughput.Manual).
          +   ``PAY_PER_REQUEST`` - We recommend using ``PAY_PER_REQUEST`` for unpredictable workloads. ``PAY_PER_REQUEST`` sets the billing mode to [On-Demand Mode](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.OnDemand).
          
         If not specified, the default is ``PROVISIONED``.
        """
        return pulumi.get(self, "billing_mode")

    @property
    @pulumi.getter(name="contributorInsightsSpecification")
    def contributor_insights_specification(self) -> Optional['outputs.TableContributorInsightsSpecification']:
        """
        The settings used to enable or disable CloudWatch Contributor Insights for the specified table.
        """
        return pulumi.get(self, "contributor_insights_specification")

    @property
    @pulumi.getter(name="deletionProtectionEnabled")
    def deletion_protection_enabled(self) -> Optional[bool]:
        """
        Determines if a table is protected from deletion. When enabled, the table cannot be deleted by any user or process. This setting is disabled by default. For more information, see [Using deletion protection](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.Basics.html#WorkingWithTables.Basics.DeletionProtection) in the *Developer Guide*.
        """
        return pulumi.get(self, "deletion_protection_enabled")

    @property
    @pulumi.getter(name="globalSecondaryIndexes")
    def global_secondary_indexes(self) -> Optional[Sequence['outputs.TableGlobalSecondaryIndex']]:
        """
        Global secondary indexes to be created on the table. You can create up to 20 global secondary indexes.
          If you update a table to include a new global secondary index, CFNlong initiates the index creation and then proceeds with the stack update. CFNlong doesn't wait for the index to complete creation because the backfilling phase can take a long time, depending on the size of the table. You can't use the index or update the table until the index's status is ``ACTIVE``. You can track its status by using the DynamoDB [DescribeTable](https://docs.aws.amazon.com/cli/latest/reference/dynamodb/describe-table.html) command.
         If you add or delete an index during an update, we recommend that you don't update any other resources. If your stack fails to update and is rolled back while adding a new index, you must manually delete the index. 
         Updates are not supported. The following are exceptions:
          +  If you update either the contributor insights specification or the provisioned throughput value
        """
        return pulumi.get(self, "global_secondary_indexes")

    @property
    @pulumi.getter(name="keySchema")
    def key_schema(self) -> Optional[Any]:
        """
        Specifies the attributes that make up the primary key for the table. The attributes in the ``KeySchema`` property must also be defined in the ``AttributeDefinitions`` property.
        """
        return pulumi.get(self, "key_schema")

    @property
    @pulumi.getter(name="kinesisStreamSpecification")
    def kinesis_stream_specification(self) -> Optional['outputs.TableKinesisStreamSpecification']:
        """
        The Kinesis Data Streams configuration for the specified table.
        """
        return pulumi.get(self, "kinesis_stream_specification")

    @property
    @pulumi.getter(name="localSecondaryIndexes")
    def local_secondary_indexes(self) -> Optional[Sequence['outputs.TableLocalSecondaryIndex']]:
        """
        Local secondary indexes to be created on the table. You can create up to 5 local secondary indexes. Each index is scoped to a given hash key value. The size of each hash key can be up to 10 gigabytes.
        """
        return pulumi.get(self, "local_secondary_indexes")

    @property
    @pulumi.getter(name="pointInTimeRecoverySpecification")
    def point_in_time_recovery_specification(self) -> Optional['outputs.TablePointInTimeRecoverySpecification']:
        """
        The settings used to enable point in time recovery.
        """
        return pulumi.get(self, "point_in_time_recovery_specification")

    @property
    @pulumi.getter(name="provisionedThroughput")
    def provisioned_throughput(self) -> Optional['outputs.TableProvisionedThroughput']:
        """
        Throughput for the specified table, which consists of values for ``ReadCapacityUnits`` and ``WriteCapacityUnits``. For more information about the contents of a provisioned throughput structure, see [Amazon DynamoDB Table ProvisionedThroughput](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ProvisionedThroughput.html). 
         If you set ``BillingMode`` as ``PROVISIONED``, you must specify this property. If you set ``BillingMode`` as ``PAY_PER_REQUEST``, you cannot specify this property.
        """
        return pulumi.get(self, "provisioned_throughput")

    @property
    @pulumi.getter(name="sseSpecification")
    def sse_specification(self) -> Optional['outputs.TableSseSpecification']:
        """
        Specifies the settings to enable server-side encryption.
        """
        return pulumi.get(self, "sse_specification")

    @property
    @pulumi.getter(name="streamArn")
    def stream_arn(self) -> Optional[str]:
        return pulumi.get(self, "stream_arn")

    @property
    @pulumi.getter(name="streamSpecification")
    def stream_specification(self) -> Optional['outputs.TableStreamSpecification']:
        """
        The settings for the DDB table stream, which capture changes to items stored in the table.
        """
        return pulumi.get(self, "stream_specification")

    @property
    @pulumi.getter(name="tableClass")
    def table_class(self) -> Optional[str]:
        """
        The table class of the new table. Valid values are ``STANDARD`` and ``STANDARD_INFREQUENT_ACCESS``.
        """
        return pulumi.get(self, "table_class")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        An array of key-value pairs to apply to this resource.
         For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="timeToLiveSpecification")
    def time_to_live_specification(self) -> Optional['outputs.TableTimeToLiveSpecification']:
        """
        Specifies the Time to Live (TTL) settings for the table.
          For detailed information about the limits in DynamoDB, see [Limits in Amazon DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the Amazon DynamoDB Developer Guide.
        """
        return pulumi.get(self, "time_to_live_specification")


class AwaitableGetTableResult(GetTableResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTableResult(
            arn=self.arn,
            attribute_definitions=self.attribute_definitions,
            billing_mode=self.billing_mode,
            contributor_insights_specification=self.contributor_insights_specification,
            deletion_protection_enabled=self.deletion_protection_enabled,
            global_secondary_indexes=self.global_secondary_indexes,
            key_schema=self.key_schema,
            kinesis_stream_specification=self.kinesis_stream_specification,
            local_secondary_indexes=self.local_secondary_indexes,
            point_in_time_recovery_specification=self.point_in_time_recovery_specification,
            provisioned_throughput=self.provisioned_throughput,
            sse_specification=self.sse_specification,
            stream_arn=self.stream_arn,
            stream_specification=self.stream_specification,
            table_class=self.table_class,
            tags=self.tags,
            time_to_live_specification=self.time_to_live_specification)


def get_table(table_name: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTableResult:
    """
    The ``AWS::DynamoDB::Table`` resource creates a DDB table. For more information, see [CreateTable](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_CreateTable.html) in the *API Reference*.
     You should be aware of the following behaviors when working with DDB tables:
      +  CFNlong typically creates DDB tables in parallel. However, if your template includes multiple DDB tables with indexes, you must declare dependencies so that the tables are created sequentially. DDBlong limits the number of tables with secondary indexes that are in the creating state. If you create multiple tables with indexes at the same time, DDB returns an error and the stack operation fails. For an example, see [DynamoDB Table with a DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#aws-resource-dynamodb-table--examples--DynamoDB_Table_with_a_DependsOn_Attribute).

       Our guidance is to use the latest schema documented here for y


    :param str table_name: A name for the table. If you don't specify a name, CFNlong generates a unique physical ID and uses that ID for the table name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
             If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
    """
    __args__ = dict()
    __args__['tableName'] = table_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:dynamodb:getTable', __args__, opts=opts, typ=GetTableResult).value

    return AwaitableGetTableResult(
        arn=pulumi.get(__ret__, 'arn'),
        attribute_definitions=pulumi.get(__ret__, 'attribute_definitions'),
        billing_mode=pulumi.get(__ret__, 'billing_mode'),
        contributor_insights_specification=pulumi.get(__ret__, 'contributor_insights_specification'),
        deletion_protection_enabled=pulumi.get(__ret__, 'deletion_protection_enabled'),
        global_secondary_indexes=pulumi.get(__ret__, 'global_secondary_indexes'),
        key_schema=pulumi.get(__ret__, 'key_schema'),
        kinesis_stream_specification=pulumi.get(__ret__, 'kinesis_stream_specification'),
        local_secondary_indexes=pulumi.get(__ret__, 'local_secondary_indexes'),
        point_in_time_recovery_specification=pulumi.get(__ret__, 'point_in_time_recovery_specification'),
        provisioned_throughput=pulumi.get(__ret__, 'provisioned_throughput'),
        sse_specification=pulumi.get(__ret__, 'sse_specification'),
        stream_arn=pulumi.get(__ret__, 'stream_arn'),
        stream_specification=pulumi.get(__ret__, 'stream_specification'),
        table_class=pulumi.get(__ret__, 'table_class'),
        tags=pulumi.get(__ret__, 'tags'),
        time_to_live_specification=pulumi.get(__ret__, 'time_to_live_specification'))


@_utilities.lift_output_func(get_table)
def get_table_output(table_name: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTableResult]:
    """
    The ``AWS::DynamoDB::Table`` resource creates a DDB table. For more information, see [CreateTable](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_CreateTable.html) in the *API Reference*.
     You should be aware of the following behaviors when working with DDB tables:
      +  CFNlong typically creates DDB tables in parallel. However, if your template includes multiple DDB tables with indexes, you must declare dependencies so that the tables are created sequentially. DDBlong limits the number of tables with secondary indexes that are in the creating state. If you create multiple tables with indexes at the same time, DDB returns an error and the stack operation fails. For an example, see [DynamoDB Table with a DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#aws-resource-dynamodb-table--examples--DynamoDB_Table_with_a_DependsOn_Attribute).

       Our guidance is to use the latest schema documented here for y


    :param str table_name: A name for the table. If you don't specify a name, CFNlong generates a unique physical ID and uses that ID for the table name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
             If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
    """
    ...
