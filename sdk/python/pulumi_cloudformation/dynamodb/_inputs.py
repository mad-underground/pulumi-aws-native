# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from .. import _inputs as _root_inputs

__all__ = [
    'TableAttributeDefinitionArgs',
    'TableGlobalSecondaryIndexArgs',
    'TableKeySchemaArgs',
    'TableLocalSecondaryIndexArgs',
    'TablePointInTimeRecoverySpecificationArgs',
    'TableProjectionArgs',
    'TablePropertiesArgs',
    'TableProvisionedThroughputArgs',
    'TableSSESpecificationArgs',
    'TableStreamSpecificationArgs',
    'TableTimeToLiveSpecificationArgs',
]

@pulumi.input_type
class TableAttributeDefinitionArgs:
    def __init__(__self__, *,
                 attribute_name: pulumi.Input[str],
                 attribute_type: pulumi.Input[str]):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-attributedef.html
        :param pulumi.Input[str] attribute_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-attributedef.html#cfn-dynamodb-attributedef-attributename
        :param pulumi.Input[str] attribute_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-attributedef.html#cfn-dynamodb-attributedef-attributename-attributetype
        """
        pulumi.set(__self__, "attribute_name", attribute_name)
        pulumi.set(__self__, "attribute_type", attribute_type)

    @property
    @pulumi.getter(name="AttributeName")
    def attribute_name(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-attributedef.html#cfn-dynamodb-attributedef-attributename
        """
        return pulumi.get(self, "attribute_name")

    @attribute_name.setter
    def attribute_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "attribute_name", value)

    @property
    @pulumi.getter(name="AttributeType")
    def attribute_type(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-attributedef.html#cfn-dynamodb-attributedef-attributename-attributetype
        """
        return pulumi.get(self, "attribute_type")

    @attribute_type.setter
    def attribute_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "attribute_type", value)


@pulumi.input_type
class TableGlobalSecondaryIndexArgs:
    def __init__(__self__, *,
                 index_name: pulumi.Input[str],
                 key_schema: pulumi.Input[Sequence[pulumi.Input['TableKeySchemaArgs']]],
                 projection: pulumi.Input['TableProjectionArgs'],
                 provisioned_throughput: Optional[pulumi.Input['TableProvisionedThroughputArgs']] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html
        :param pulumi.Input[str] index_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html#cfn-dynamodb-gsi-indexname
        :param pulumi.Input[Sequence[pulumi.Input['TableKeySchemaArgs']]] key_schema: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html#cfn-dynamodb-gsi-keyschema
        :param pulumi.Input['TableProjectionArgs'] projection: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html#cfn-dynamodb-gsi-projection
        :param pulumi.Input['TableProvisionedThroughputArgs'] provisioned_throughput: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html#cfn-dynamodb-gsi-provisionedthroughput
        """
        pulumi.set(__self__, "index_name", index_name)
        pulumi.set(__self__, "key_schema", key_schema)
        pulumi.set(__self__, "projection", projection)
        if provisioned_throughput is not None:
            pulumi.set(__self__, "provisioned_throughput", provisioned_throughput)

    @property
    @pulumi.getter(name="IndexName")
    def index_name(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html#cfn-dynamodb-gsi-indexname
        """
        return pulumi.get(self, "index_name")

    @index_name.setter
    def index_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "index_name", value)

    @property
    @pulumi.getter(name="KeySchema")
    def key_schema(self) -> pulumi.Input[Sequence[pulumi.Input['TableKeySchemaArgs']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html#cfn-dynamodb-gsi-keyschema
        """
        return pulumi.get(self, "key_schema")

    @key_schema.setter
    def key_schema(self, value: pulumi.Input[Sequence[pulumi.Input['TableKeySchemaArgs']]]):
        pulumi.set(self, "key_schema", value)

    @property
    @pulumi.getter(name="Projection")
    def projection(self) -> pulumi.Input['TableProjectionArgs']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html#cfn-dynamodb-gsi-projection
        """
        return pulumi.get(self, "projection")

    @projection.setter
    def projection(self, value: pulumi.Input['TableProjectionArgs']):
        pulumi.set(self, "projection", value)

    @property
    @pulumi.getter(name="ProvisionedThroughput")
    def provisioned_throughput(self) -> Optional[pulumi.Input['TableProvisionedThroughputArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html#cfn-dynamodb-gsi-provisionedthroughput
        """
        return pulumi.get(self, "provisioned_throughput")

    @provisioned_throughput.setter
    def provisioned_throughput(self, value: Optional[pulumi.Input['TableProvisionedThroughputArgs']]):
        pulumi.set(self, "provisioned_throughput", value)


@pulumi.input_type
class TableKeySchemaArgs:
    def __init__(__self__, *,
                 attribute_name: pulumi.Input[str],
                 key_type: pulumi.Input[str]):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-keyschema.html
        :param pulumi.Input[str] attribute_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-keyschema.html#aws-properties-dynamodb-keyschema-attributename
        :param pulumi.Input[str] key_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-keyschema.html#aws-properties-dynamodb-keyschema-keytype
        """
        pulumi.set(__self__, "attribute_name", attribute_name)
        pulumi.set(__self__, "key_type", key_type)

    @property
    @pulumi.getter(name="AttributeName")
    def attribute_name(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-keyschema.html#aws-properties-dynamodb-keyschema-attributename
        """
        return pulumi.get(self, "attribute_name")

    @attribute_name.setter
    def attribute_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "attribute_name", value)

    @property
    @pulumi.getter(name="KeyType")
    def key_type(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-keyschema.html#aws-properties-dynamodb-keyschema-keytype
        """
        return pulumi.get(self, "key_type")

    @key_type.setter
    def key_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_type", value)


@pulumi.input_type
class TableLocalSecondaryIndexArgs:
    def __init__(__self__, *,
                 index_name: pulumi.Input[str],
                 key_schema: pulumi.Input[Sequence[pulumi.Input['TableKeySchemaArgs']]],
                 projection: pulumi.Input['TableProjectionArgs']):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-lsi.html
        :param pulumi.Input[str] index_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-lsi.html#cfn-dynamodb-lsi-indexname
        :param pulumi.Input[Sequence[pulumi.Input['TableKeySchemaArgs']]] key_schema: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-lsi.html#cfn-dynamodb-lsi-keyschema
        :param pulumi.Input['TableProjectionArgs'] projection: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-lsi.html#cfn-dynamodb-lsi-projection
        """
        pulumi.set(__self__, "index_name", index_name)
        pulumi.set(__self__, "key_schema", key_schema)
        pulumi.set(__self__, "projection", projection)

    @property
    @pulumi.getter(name="IndexName")
    def index_name(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-lsi.html#cfn-dynamodb-lsi-indexname
        """
        return pulumi.get(self, "index_name")

    @index_name.setter
    def index_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "index_name", value)

    @property
    @pulumi.getter(name="KeySchema")
    def key_schema(self) -> pulumi.Input[Sequence[pulumi.Input['TableKeySchemaArgs']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-lsi.html#cfn-dynamodb-lsi-keyschema
        """
        return pulumi.get(self, "key_schema")

    @key_schema.setter
    def key_schema(self, value: pulumi.Input[Sequence[pulumi.Input['TableKeySchemaArgs']]]):
        pulumi.set(self, "key_schema", value)

    @property
    @pulumi.getter(name="Projection")
    def projection(self) -> pulumi.Input['TableProjectionArgs']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-lsi.html#cfn-dynamodb-lsi-projection
        """
        return pulumi.get(self, "projection")

    @projection.setter
    def projection(self, value: pulumi.Input['TableProjectionArgs']):
        pulumi.set(self, "projection", value)


@pulumi.input_type
class TablePointInTimeRecoverySpecificationArgs:
    def __init__(__self__, *,
                 point_in_time_recovery_enabled: Optional[pulumi.Input[bool]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-pointintimerecoveryspecification.html
        :param pulumi.Input[bool] point_in_time_recovery_enabled: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-pointintimerecoveryspecification.html#cfn-dynamodb-table-pointintimerecoveryspecification-pointintimerecoveryenabled
        """
        if point_in_time_recovery_enabled is not None:
            pulumi.set(__self__, "point_in_time_recovery_enabled", point_in_time_recovery_enabled)

    @property
    @pulumi.getter(name="PointInTimeRecoveryEnabled")
    def point_in_time_recovery_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-pointintimerecoveryspecification.html#cfn-dynamodb-table-pointintimerecoveryspecification-pointintimerecoveryenabled
        """
        return pulumi.get(self, "point_in_time_recovery_enabled")

    @point_in_time_recovery_enabled.setter
    def point_in_time_recovery_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "point_in_time_recovery_enabled", value)


@pulumi.input_type
class TableProjectionArgs:
    def __init__(__self__, *,
                 non_key_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 projection_type: Optional[pulumi.Input[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-projectionobject.html
        :param pulumi.Input[Sequence[pulumi.Input[str]]] non_key_attributes: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-projectionobject.html#cfn-dynamodb-projectionobj-nonkeyatt
        :param pulumi.Input[str] projection_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-projectionobject.html#cfn-dynamodb-projectionobj-projtype
        """
        if non_key_attributes is not None:
            pulumi.set(__self__, "non_key_attributes", non_key_attributes)
        if projection_type is not None:
            pulumi.set(__self__, "projection_type", projection_type)

    @property
    @pulumi.getter(name="NonKeyAttributes")
    def non_key_attributes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-projectionobject.html#cfn-dynamodb-projectionobj-nonkeyatt
        """
        return pulumi.get(self, "non_key_attributes")

    @non_key_attributes.setter
    def non_key_attributes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "non_key_attributes", value)

    @property
    @pulumi.getter(name="ProjectionType")
    def projection_type(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-projectionobject.html#cfn-dynamodb-projectionobj-projtype
        """
        return pulumi.get(self, "projection_type")

    @projection_type.setter
    def projection_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "projection_type", value)


@pulumi.input_type
class TablePropertiesArgs:
    def __init__(__self__, *,
                 key_schema: pulumi.Input[Sequence[pulumi.Input['TableKeySchemaArgs']]],
                 attribute_definitions: Optional[pulumi.Input[Sequence[pulumi.Input['TableAttributeDefinitionArgs']]]] = None,
                 billing_mode: Optional[pulumi.Input[str]] = None,
                 global_secondary_indexes: Optional[pulumi.Input[Sequence[pulumi.Input['TableGlobalSecondaryIndexArgs']]]] = None,
                 local_secondary_indexes: Optional[pulumi.Input[Sequence[pulumi.Input['TableLocalSecondaryIndexArgs']]]] = None,
                 point_in_time_recovery_specification: Optional[pulumi.Input['TablePointInTimeRecoverySpecificationArgs']] = None,
                 provisioned_throughput: Optional[pulumi.Input['TableProvisionedThroughputArgs']] = None,
                 sse_specification: Optional[pulumi.Input['TableSSESpecificationArgs']] = None,
                 stream_specification: Optional[pulumi.Input['TableStreamSpecificationArgs']] = None,
                 table_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None,
                 time_to_live_specification: Optional[pulumi.Input['TableTimeToLiveSpecificationArgs']] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html
        :param pulumi.Input[Sequence[pulumi.Input['TableKeySchemaArgs']]] key_schema: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-keyschema
        :param pulumi.Input[Sequence[pulumi.Input['TableAttributeDefinitionArgs']]] attribute_definitions: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-attributedef
        :param pulumi.Input[str] billing_mode: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-billingmode
        :param pulumi.Input[Sequence[pulumi.Input['TableGlobalSecondaryIndexArgs']]] global_secondary_indexes: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-gsi
        :param pulumi.Input[Sequence[pulumi.Input['TableLocalSecondaryIndexArgs']]] local_secondary_indexes: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-lsi
        :param pulumi.Input['TablePointInTimeRecoverySpecificationArgs'] point_in_time_recovery_specification: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-pointintimerecoveryspecification
        :param pulumi.Input['TableProvisionedThroughputArgs'] provisioned_throughput: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-provisionedthroughput
        :param pulumi.Input['TableSSESpecificationArgs'] sse_specification: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-ssespecification
        :param pulumi.Input['TableStreamSpecificationArgs'] stream_specification: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-streamspecification
        :param pulumi.Input[str] table_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-tablename
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-tags
        :param pulumi.Input['TableTimeToLiveSpecificationArgs'] time_to_live_specification: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-timetolivespecification
        """
        pulumi.set(__self__, "key_schema", key_schema)
        if attribute_definitions is not None:
            pulumi.set(__self__, "attribute_definitions", attribute_definitions)
        if billing_mode is not None:
            pulumi.set(__self__, "billing_mode", billing_mode)
        if global_secondary_indexes is not None:
            pulumi.set(__self__, "global_secondary_indexes", global_secondary_indexes)
        if local_secondary_indexes is not None:
            pulumi.set(__self__, "local_secondary_indexes", local_secondary_indexes)
        if point_in_time_recovery_specification is not None:
            pulumi.set(__self__, "point_in_time_recovery_specification", point_in_time_recovery_specification)
        if provisioned_throughput is not None:
            pulumi.set(__self__, "provisioned_throughput", provisioned_throughput)
        if sse_specification is not None:
            pulumi.set(__self__, "sse_specification", sse_specification)
        if stream_specification is not None:
            pulumi.set(__self__, "stream_specification", stream_specification)
        if table_name is not None:
            pulumi.set(__self__, "table_name", table_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if time_to_live_specification is not None:
            pulumi.set(__self__, "time_to_live_specification", time_to_live_specification)

    @property
    @pulumi.getter(name="KeySchema")
    def key_schema(self) -> pulumi.Input[Sequence[pulumi.Input['TableKeySchemaArgs']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-keyschema
        """
        return pulumi.get(self, "key_schema")

    @key_schema.setter
    def key_schema(self, value: pulumi.Input[Sequence[pulumi.Input['TableKeySchemaArgs']]]):
        pulumi.set(self, "key_schema", value)

    @property
    @pulumi.getter(name="AttributeDefinitions")
    def attribute_definitions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TableAttributeDefinitionArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-attributedef
        """
        return pulumi.get(self, "attribute_definitions")

    @attribute_definitions.setter
    def attribute_definitions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TableAttributeDefinitionArgs']]]]):
        pulumi.set(self, "attribute_definitions", value)

    @property
    @pulumi.getter(name="BillingMode")
    def billing_mode(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-billingmode
        """
        return pulumi.get(self, "billing_mode")

    @billing_mode.setter
    def billing_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "billing_mode", value)

    @property
    @pulumi.getter(name="GlobalSecondaryIndexes")
    def global_secondary_indexes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TableGlobalSecondaryIndexArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-gsi
        """
        return pulumi.get(self, "global_secondary_indexes")

    @global_secondary_indexes.setter
    def global_secondary_indexes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TableGlobalSecondaryIndexArgs']]]]):
        pulumi.set(self, "global_secondary_indexes", value)

    @property
    @pulumi.getter(name="LocalSecondaryIndexes")
    def local_secondary_indexes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TableLocalSecondaryIndexArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-lsi
        """
        return pulumi.get(self, "local_secondary_indexes")

    @local_secondary_indexes.setter
    def local_secondary_indexes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TableLocalSecondaryIndexArgs']]]]):
        pulumi.set(self, "local_secondary_indexes", value)

    @property
    @pulumi.getter(name="PointInTimeRecoverySpecification")
    def point_in_time_recovery_specification(self) -> Optional[pulumi.Input['TablePointInTimeRecoverySpecificationArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-pointintimerecoveryspecification
        """
        return pulumi.get(self, "point_in_time_recovery_specification")

    @point_in_time_recovery_specification.setter
    def point_in_time_recovery_specification(self, value: Optional[pulumi.Input['TablePointInTimeRecoverySpecificationArgs']]):
        pulumi.set(self, "point_in_time_recovery_specification", value)

    @property
    @pulumi.getter(name="ProvisionedThroughput")
    def provisioned_throughput(self) -> Optional[pulumi.Input['TableProvisionedThroughputArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-provisionedthroughput
        """
        return pulumi.get(self, "provisioned_throughput")

    @provisioned_throughput.setter
    def provisioned_throughput(self, value: Optional[pulumi.Input['TableProvisionedThroughputArgs']]):
        pulumi.set(self, "provisioned_throughput", value)

    @property
    @pulumi.getter(name="SSESpecification")
    def sse_specification(self) -> Optional[pulumi.Input['TableSSESpecificationArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-ssespecification
        """
        return pulumi.get(self, "sse_specification")

    @sse_specification.setter
    def sse_specification(self, value: Optional[pulumi.Input['TableSSESpecificationArgs']]):
        pulumi.set(self, "sse_specification", value)

    @property
    @pulumi.getter(name="StreamSpecification")
    def stream_specification(self) -> Optional[pulumi.Input['TableStreamSpecificationArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-streamspecification
        """
        return pulumi.get(self, "stream_specification")

    @stream_specification.setter
    def stream_specification(self, value: Optional[pulumi.Input['TableStreamSpecificationArgs']]):
        pulumi.set(self, "stream_specification", value)

    @property
    @pulumi.getter(name="TableName")
    def table_name(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-tablename
        """
        return pulumi.get(self, "table_name")

    @table_name.setter
    def table_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "table_name", value)

    @property
    @pulumi.getter(name="Tags")
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="TimeToLiveSpecification")
    def time_to_live_specification(self) -> Optional[pulumi.Input['TableTimeToLiveSpecificationArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-timetolivespecification
        """
        return pulumi.get(self, "time_to_live_specification")

    @time_to_live_specification.setter
    def time_to_live_specification(self, value: Optional[pulumi.Input['TableTimeToLiveSpecificationArgs']]):
        pulumi.set(self, "time_to_live_specification", value)


@pulumi.input_type
class TableProvisionedThroughputArgs:
    def __init__(__self__, *,
                 read_capacity_units: pulumi.Input[int],
                 write_capacity_units: pulumi.Input[int]):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-provisionedthroughput.html
        :param pulumi.Input[int] read_capacity_units: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-provisionedthroughput.html#cfn-dynamodb-provisionedthroughput-readcapacityunits
        :param pulumi.Input[int] write_capacity_units: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-provisionedthroughput.html#cfn-dynamodb-provisionedthroughput-writecapacityunits
        """
        pulumi.set(__self__, "read_capacity_units", read_capacity_units)
        pulumi.set(__self__, "write_capacity_units", write_capacity_units)

    @property
    @pulumi.getter(name="ReadCapacityUnits")
    def read_capacity_units(self) -> pulumi.Input[int]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-provisionedthroughput.html#cfn-dynamodb-provisionedthroughput-readcapacityunits
        """
        return pulumi.get(self, "read_capacity_units")

    @read_capacity_units.setter
    def read_capacity_units(self, value: pulumi.Input[int]):
        pulumi.set(self, "read_capacity_units", value)

    @property
    @pulumi.getter(name="WriteCapacityUnits")
    def write_capacity_units(self) -> pulumi.Input[int]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-provisionedthroughput.html#cfn-dynamodb-provisionedthroughput-writecapacityunits
        """
        return pulumi.get(self, "write_capacity_units")

    @write_capacity_units.setter
    def write_capacity_units(self, value: pulumi.Input[int]):
        pulumi.set(self, "write_capacity_units", value)


@pulumi.input_type
class TableSSESpecificationArgs:
    def __init__(__self__, *,
                 sse_enabled: pulumi.Input[bool],
                 kms_master_key_id: Optional[pulumi.Input[str]] = None,
                 sse_type: Optional[pulumi.Input[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-ssespecification.html
        :param pulumi.Input[bool] sse_enabled: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-ssespecification.html#cfn-dynamodb-table-ssespecification-sseenabled
        :param pulumi.Input[str] kms_master_key_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-ssespecification.html#cfn-dynamodb-table-ssespecification-kmsmasterkeyid
        :param pulumi.Input[str] sse_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-ssespecification.html#cfn-dynamodb-table-ssespecification-ssetype
        """
        pulumi.set(__self__, "sse_enabled", sse_enabled)
        if kms_master_key_id is not None:
            pulumi.set(__self__, "kms_master_key_id", kms_master_key_id)
        if sse_type is not None:
            pulumi.set(__self__, "sse_type", sse_type)

    @property
    @pulumi.getter(name="SSEEnabled")
    def sse_enabled(self) -> pulumi.Input[bool]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-ssespecification.html#cfn-dynamodb-table-ssespecification-sseenabled
        """
        return pulumi.get(self, "sse_enabled")

    @sse_enabled.setter
    def sse_enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "sse_enabled", value)

    @property
    @pulumi.getter(name="KMSMasterKeyId")
    def kms_master_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-ssespecification.html#cfn-dynamodb-table-ssespecification-kmsmasterkeyid
        """
        return pulumi.get(self, "kms_master_key_id")

    @kms_master_key_id.setter
    def kms_master_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_master_key_id", value)

    @property
    @pulumi.getter(name="SSEType")
    def sse_type(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-ssespecification.html#cfn-dynamodb-table-ssespecification-ssetype
        """
        return pulumi.get(self, "sse_type")

    @sse_type.setter
    def sse_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sse_type", value)


@pulumi.input_type
class TableStreamSpecificationArgs:
    def __init__(__self__, *,
                 stream_view_type: pulumi.Input[str]):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-streamspecification.html
        :param pulumi.Input[str] stream_view_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-streamspecification.html#cfn-dynamodb-streamspecification-streamviewtype
        """
        pulumi.set(__self__, "stream_view_type", stream_view_type)

    @property
    @pulumi.getter(name="StreamViewType")
    def stream_view_type(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-streamspecification.html#cfn-dynamodb-streamspecification-streamviewtype
        """
        return pulumi.get(self, "stream_view_type")

    @stream_view_type.setter
    def stream_view_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "stream_view_type", value)


@pulumi.input_type
class TableTimeToLiveSpecificationArgs:
    def __init__(__self__, *,
                 attribute_name: pulumi.Input[str],
                 enabled: pulumi.Input[bool]):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-timetolivespecification.html
        :param pulumi.Input[str] attribute_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-timetolivespecification.html#cfn-dynamodb-timetolivespecification-attributename
        :param pulumi.Input[bool] enabled: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-timetolivespecification.html#cfn-dynamodb-timetolivespecification-enabled
        """
        pulumi.set(__self__, "attribute_name", attribute_name)
        pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter(name="AttributeName")
    def attribute_name(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-timetolivespecification.html#cfn-dynamodb-timetolivespecification-attributename
        """
        return pulumi.get(self, "attribute_name")

    @attribute_name.setter
    def attribute_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "attribute_name", value)

    @property
    @pulumi.getter(name="Enabled")
    def enabled(self) -> pulumi.Input[bool]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-timetolivespecification.html#cfn-dynamodb-timetolivespecification-enabled
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)


