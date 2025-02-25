# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = [
    'AnalysisTemplateAnalysisParameterArgs',
    'AnalysisTemplateAnalysisSourceArgs',
    'CollaborationDataEncryptionMetadataArgs',
    'CollaborationMemberSpecificationArgs',
    'CollaborationPaymentConfigurationArgs',
    'CollaborationQueryComputePaymentConfigArgs',
    'ConfiguredTableAggregateColumnArgs',
    'ConfiguredTableAggregationConstraintArgs',
    'ConfiguredTableAnalysisRuleAggregationArgs',
    'ConfiguredTableAnalysisRuleCustomArgs',
    'ConfiguredTableAnalysisRuleListArgs',
    'ConfiguredTableAnalysisRulePolicyV10PropertiesArgs',
    'ConfiguredTableAnalysisRulePolicyV11PropertiesArgs',
    'ConfiguredTableAnalysisRulePolicyV12PropertiesArgs',
    'ConfiguredTableAnalysisRulePolicyArgs',
    'ConfiguredTableAnalysisRuleArgs',
    'ConfiguredTableGlueTableReferenceArgs',
    'ConfiguredTableTableReferenceArgs',
    'MembershipPaymentConfigurationArgs',
    'MembershipProtectedQueryOutputConfigurationArgs',
    'MembershipProtectedQueryResultConfigurationArgs',
    'MembershipProtectedQueryS3OutputConfigurationArgs',
    'MembershipQueryComputePaymentConfigArgs',
]

@pulumi.input_type
class AnalysisTemplateAnalysisParameterArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 type: pulumi.Input['AnalysisTemplateAnalysisParameterType'],
                 default_value: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)
        if default_value is not None:
            pulumi.set(__self__, "default_value", default_value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input['AnalysisTemplateAnalysisParameterType']:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input['AnalysisTemplateAnalysisParameterType']):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="defaultValue")
    def default_value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "default_value")

    @default_value.setter
    def default_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_value", value)


@pulumi.input_type
class AnalysisTemplateAnalysisSourceArgs:
    def __init__(__self__, *,
                 text: pulumi.Input[str]):
        pulumi.set(__self__, "text", text)

    @property
    @pulumi.getter
    def text(self) -> pulumi.Input[str]:
        return pulumi.get(self, "text")

    @text.setter
    def text(self, value: pulumi.Input[str]):
        pulumi.set(self, "text", value)


@pulumi.input_type
class CollaborationDataEncryptionMetadataArgs:
    def __init__(__self__, *,
                 allow_cleartext: pulumi.Input[bool],
                 allow_duplicates: pulumi.Input[bool],
                 allow_joins_on_columns_with_different_names: pulumi.Input[bool],
                 preserve_nulls: pulumi.Input[bool]):
        pulumi.set(__self__, "allow_cleartext", allow_cleartext)
        pulumi.set(__self__, "allow_duplicates", allow_duplicates)
        pulumi.set(__self__, "allow_joins_on_columns_with_different_names", allow_joins_on_columns_with_different_names)
        pulumi.set(__self__, "preserve_nulls", preserve_nulls)

    @property
    @pulumi.getter(name="allowCleartext")
    def allow_cleartext(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "allow_cleartext")

    @allow_cleartext.setter
    def allow_cleartext(self, value: pulumi.Input[bool]):
        pulumi.set(self, "allow_cleartext", value)

    @property
    @pulumi.getter(name="allowDuplicates")
    def allow_duplicates(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "allow_duplicates")

    @allow_duplicates.setter
    def allow_duplicates(self, value: pulumi.Input[bool]):
        pulumi.set(self, "allow_duplicates", value)

    @property
    @pulumi.getter(name="allowJoinsOnColumnsWithDifferentNames")
    def allow_joins_on_columns_with_different_names(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "allow_joins_on_columns_with_different_names")

    @allow_joins_on_columns_with_different_names.setter
    def allow_joins_on_columns_with_different_names(self, value: pulumi.Input[bool]):
        pulumi.set(self, "allow_joins_on_columns_with_different_names", value)

    @property
    @pulumi.getter(name="preserveNulls")
    def preserve_nulls(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "preserve_nulls")

    @preserve_nulls.setter
    def preserve_nulls(self, value: pulumi.Input[bool]):
        pulumi.set(self, "preserve_nulls", value)


@pulumi.input_type
class CollaborationMemberSpecificationArgs:
    def __init__(__self__, *,
                 account_id: pulumi.Input[str],
                 display_name: pulumi.Input[str],
                 member_abilities: pulumi.Input[Sequence[pulumi.Input['CollaborationMemberAbility']]],
                 payment_configuration: Optional[pulumi.Input['CollaborationPaymentConfigurationArgs']] = None):
        pulumi.set(__self__, "account_id", account_id)
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "member_abilities", member_abilities)
        if payment_configuration is not None:
            pulumi.set(__self__, "payment_configuration", payment_configuration)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="memberAbilities")
    def member_abilities(self) -> pulumi.Input[Sequence[pulumi.Input['CollaborationMemberAbility']]]:
        return pulumi.get(self, "member_abilities")

    @member_abilities.setter
    def member_abilities(self, value: pulumi.Input[Sequence[pulumi.Input['CollaborationMemberAbility']]]):
        pulumi.set(self, "member_abilities", value)

    @property
    @pulumi.getter(name="paymentConfiguration")
    def payment_configuration(self) -> Optional[pulumi.Input['CollaborationPaymentConfigurationArgs']]:
        return pulumi.get(self, "payment_configuration")

    @payment_configuration.setter
    def payment_configuration(self, value: Optional[pulumi.Input['CollaborationPaymentConfigurationArgs']]):
        pulumi.set(self, "payment_configuration", value)


@pulumi.input_type
class CollaborationPaymentConfigurationArgs:
    def __init__(__self__, *,
                 query_compute: pulumi.Input['CollaborationQueryComputePaymentConfigArgs']):
        pulumi.set(__self__, "query_compute", query_compute)

    @property
    @pulumi.getter(name="queryCompute")
    def query_compute(self) -> pulumi.Input['CollaborationQueryComputePaymentConfigArgs']:
        return pulumi.get(self, "query_compute")

    @query_compute.setter
    def query_compute(self, value: pulumi.Input['CollaborationQueryComputePaymentConfigArgs']):
        pulumi.set(self, "query_compute", value)


@pulumi.input_type
class CollaborationQueryComputePaymentConfigArgs:
    def __init__(__self__, *,
                 is_responsible: pulumi.Input[bool]):
        pulumi.set(__self__, "is_responsible", is_responsible)

    @property
    @pulumi.getter(name="isResponsible")
    def is_responsible(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "is_responsible")

    @is_responsible.setter
    def is_responsible(self, value: pulumi.Input[bool]):
        pulumi.set(self, "is_responsible", value)


@pulumi.input_type
class ConfiguredTableAggregateColumnArgs:
    def __init__(__self__, *,
                 column_names: pulumi.Input[Sequence[pulumi.Input[str]]],
                 function: pulumi.Input['ConfiguredTableAggregateFunctionName']):
        pulumi.set(__self__, "column_names", column_names)
        pulumi.set(__self__, "function", function)

    @property
    @pulumi.getter(name="columnNames")
    def column_names(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "column_names")

    @column_names.setter
    def column_names(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "column_names", value)

    @property
    @pulumi.getter
    def function(self) -> pulumi.Input['ConfiguredTableAggregateFunctionName']:
        return pulumi.get(self, "function")

    @function.setter
    def function(self, value: pulumi.Input['ConfiguredTableAggregateFunctionName']):
        pulumi.set(self, "function", value)


@pulumi.input_type
class ConfiguredTableAggregationConstraintArgs:
    def __init__(__self__, *,
                 column_name: pulumi.Input[str],
                 minimum: pulumi.Input[float],
                 type: pulumi.Input['ConfiguredTableAggregationType']):
        pulumi.set(__self__, "column_name", column_name)
        pulumi.set(__self__, "minimum", minimum)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="columnName")
    def column_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "column_name")

    @column_name.setter
    def column_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "column_name", value)

    @property
    @pulumi.getter
    def minimum(self) -> pulumi.Input[float]:
        return pulumi.get(self, "minimum")

    @minimum.setter
    def minimum(self, value: pulumi.Input[float]):
        pulumi.set(self, "minimum", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input['ConfiguredTableAggregationType']:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input['ConfiguredTableAggregationType']):
        pulumi.set(self, "type", value)


@pulumi.input_type
class ConfiguredTableAnalysisRuleAggregationArgs:
    def __init__(__self__, *,
                 aggregate_columns: pulumi.Input[Sequence[pulumi.Input['ConfiguredTableAggregateColumnArgs']]],
                 dimension_columns: pulumi.Input[Sequence[pulumi.Input[str]]],
                 join_columns: pulumi.Input[Sequence[pulumi.Input[str]]],
                 output_constraints: pulumi.Input[Sequence[pulumi.Input['ConfiguredTableAggregationConstraintArgs']]],
                 scalar_functions: pulumi.Input[Sequence[pulumi.Input['ConfiguredTableScalarFunctions']]],
                 allowed_join_operators: Optional[pulumi.Input[Sequence[pulumi.Input['ConfiguredTableJoinOperator']]]] = None,
                 join_required: Optional[pulumi.Input['ConfiguredTableJoinRequiredOption']] = None):
        pulumi.set(__self__, "aggregate_columns", aggregate_columns)
        pulumi.set(__self__, "dimension_columns", dimension_columns)
        pulumi.set(__self__, "join_columns", join_columns)
        pulumi.set(__self__, "output_constraints", output_constraints)
        pulumi.set(__self__, "scalar_functions", scalar_functions)
        if allowed_join_operators is not None:
            pulumi.set(__self__, "allowed_join_operators", allowed_join_operators)
        if join_required is not None:
            pulumi.set(__self__, "join_required", join_required)

    @property
    @pulumi.getter(name="aggregateColumns")
    def aggregate_columns(self) -> pulumi.Input[Sequence[pulumi.Input['ConfiguredTableAggregateColumnArgs']]]:
        return pulumi.get(self, "aggregate_columns")

    @aggregate_columns.setter
    def aggregate_columns(self, value: pulumi.Input[Sequence[pulumi.Input['ConfiguredTableAggregateColumnArgs']]]):
        pulumi.set(self, "aggregate_columns", value)

    @property
    @pulumi.getter(name="dimensionColumns")
    def dimension_columns(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "dimension_columns")

    @dimension_columns.setter
    def dimension_columns(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "dimension_columns", value)

    @property
    @pulumi.getter(name="joinColumns")
    def join_columns(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "join_columns")

    @join_columns.setter
    def join_columns(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "join_columns", value)

    @property
    @pulumi.getter(name="outputConstraints")
    def output_constraints(self) -> pulumi.Input[Sequence[pulumi.Input['ConfiguredTableAggregationConstraintArgs']]]:
        return pulumi.get(self, "output_constraints")

    @output_constraints.setter
    def output_constraints(self, value: pulumi.Input[Sequence[pulumi.Input['ConfiguredTableAggregationConstraintArgs']]]):
        pulumi.set(self, "output_constraints", value)

    @property
    @pulumi.getter(name="scalarFunctions")
    def scalar_functions(self) -> pulumi.Input[Sequence[pulumi.Input['ConfiguredTableScalarFunctions']]]:
        return pulumi.get(self, "scalar_functions")

    @scalar_functions.setter
    def scalar_functions(self, value: pulumi.Input[Sequence[pulumi.Input['ConfiguredTableScalarFunctions']]]):
        pulumi.set(self, "scalar_functions", value)

    @property
    @pulumi.getter(name="allowedJoinOperators")
    def allowed_join_operators(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ConfiguredTableJoinOperator']]]]:
        return pulumi.get(self, "allowed_join_operators")

    @allowed_join_operators.setter
    def allowed_join_operators(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ConfiguredTableJoinOperator']]]]):
        pulumi.set(self, "allowed_join_operators", value)

    @property
    @pulumi.getter(name="joinRequired")
    def join_required(self) -> Optional[pulumi.Input['ConfiguredTableJoinRequiredOption']]:
        return pulumi.get(self, "join_required")

    @join_required.setter
    def join_required(self, value: Optional[pulumi.Input['ConfiguredTableJoinRequiredOption']]):
        pulumi.set(self, "join_required", value)


@pulumi.input_type
class ConfiguredTableAnalysisRuleCustomArgs:
    def __init__(__self__, *,
                 allowed_analyses: pulumi.Input[Sequence[pulumi.Input[str]]],
                 allowed_analysis_providers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        pulumi.set(__self__, "allowed_analyses", allowed_analyses)
        if allowed_analysis_providers is not None:
            pulumi.set(__self__, "allowed_analysis_providers", allowed_analysis_providers)

    @property
    @pulumi.getter(name="allowedAnalyses")
    def allowed_analyses(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "allowed_analyses")

    @allowed_analyses.setter
    def allowed_analyses(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "allowed_analyses", value)

    @property
    @pulumi.getter(name="allowedAnalysisProviders")
    def allowed_analysis_providers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "allowed_analysis_providers")

    @allowed_analysis_providers.setter
    def allowed_analysis_providers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_analysis_providers", value)


@pulumi.input_type
class ConfiguredTableAnalysisRuleListArgs:
    def __init__(__self__, *,
                 join_columns: pulumi.Input[Sequence[pulumi.Input[str]]],
                 list_columns: pulumi.Input[Sequence[pulumi.Input[str]]],
                 allowed_join_operators: Optional[pulumi.Input[Sequence[pulumi.Input['ConfiguredTableJoinOperator']]]] = None):
        pulumi.set(__self__, "join_columns", join_columns)
        pulumi.set(__self__, "list_columns", list_columns)
        if allowed_join_operators is not None:
            pulumi.set(__self__, "allowed_join_operators", allowed_join_operators)

    @property
    @pulumi.getter(name="joinColumns")
    def join_columns(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "join_columns")

    @join_columns.setter
    def join_columns(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "join_columns", value)

    @property
    @pulumi.getter(name="listColumns")
    def list_columns(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "list_columns")

    @list_columns.setter
    def list_columns(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "list_columns", value)

    @property
    @pulumi.getter(name="allowedJoinOperators")
    def allowed_join_operators(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ConfiguredTableJoinOperator']]]]:
        return pulumi.get(self, "allowed_join_operators")

    @allowed_join_operators.setter
    def allowed_join_operators(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ConfiguredTableJoinOperator']]]]):
        pulumi.set(self, "allowed_join_operators", value)


@pulumi.input_type
class ConfiguredTableAnalysisRulePolicyV10PropertiesArgs:
    def __init__(__self__, *,
                 list: pulumi.Input['ConfiguredTableAnalysisRuleListArgs']):
        pulumi.set(__self__, "list", list)

    @property
    @pulumi.getter
    def list(self) -> pulumi.Input['ConfiguredTableAnalysisRuleListArgs']:
        return pulumi.get(self, "list")

    @list.setter
    def list(self, value: pulumi.Input['ConfiguredTableAnalysisRuleListArgs']):
        pulumi.set(self, "list", value)


@pulumi.input_type
class ConfiguredTableAnalysisRulePolicyV11PropertiesArgs:
    def __init__(__self__, *,
                 aggregation: pulumi.Input['ConfiguredTableAnalysisRuleAggregationArgs']):
        pulumi.set(__self__, "aggregation", aggregation)

    @property
    @pulumi.getter
    def aggregation(self) -> pulumi.Input['ConfiguredTableAnalysisRuleAggregationArgs']:
        return pulumi.get(self, "aggregation")

    @aggregation.setter
    def aggregation(self, value: pulumi.Input['ConfiguredTableAnalysisRuleAggregationArgs']):
        pulumi.set(self, "aggregation", value)


@pulumi.input_type
class ConfiguredTableAnalysisRulePolicyV12PropertiesArgs:
    def __init__(__self__, *,
                 custom: pulumi.Input['ConfiguredTableAnalysisRuleCustomArgs']):
        pulumi.set(__self__, "custom", custom)

    @property
    @pulumi.getter
    def custom(self) -> pulumi.Input['ConfiguredTableAnalysisRuleCustomArgs']:
        return pulumi.get(self, "custom")

    @custom.setter
    def custom(self, value: pulumi.Input['ConfiguredTableAnalysisRuleCustomArgs']):
        pulumi.set(self, "custom", value)


@pulumi.input_type
class ConfiguredTableAnalysisRulePolicyArgs:
    def __init__(__self__, *,
                 v1: pulumi.Input[Union['ConfiguredTableAnalysisRulePolicyV10PropertiesArgs', 'ConfiguredTableAnalysisRulePolicyV11PropertiesArgs', 'ConfiguredTableAnalysisRulePolicyV12PropertiesArgs']]):
        pulumi.set(__self__, "v1", v1)

    @property
    @pulumi.getter
    def v1(self) -> pulumi.Input[Union['ConfiguredTableAnalysisRulePolicyV10PropertiesArgs', 'ConfiguredTableAnalysisRulePolicyV11PropertiesArgs', 'ConfiguredTableAnalysisRulePolicyV12PropertiesArgs']]:
        return pulumi.get(self, "v1")

    @v1.setter
    def v1(self, value: pulumi.Input[Union['ConfiguredTableAnalysisRulePolicyV10PropertiesArgs', 'ConfiguredTableAnalysisRulePolicyV11PropertiesArgs', 'ConfiguredTableAnalysisRulePolicyV12PropertiesArgs']]):
        pulumi.set(self, "v1", value)


@pulumi.input_type
class ConfiguredTableAnalysisRuleArgs:
    def __init__(__self__, *,
                 policy: pulumi.Input['ConfiguredTableAnalysisRulePolicyArgs'],
                 type: pulumi.Input['ConfiguredTableAnalysisRuleType']):
        pulumi.set(__self__, "policy", policy)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Input['ConfiguredTableAnalysisRulePolicyArgs']:
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: pulumi.Input['ConfiguredTableAnalysisRulePolicyArgs']):
        pulumi.set(self, "policy", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input['ConfiguredTableAnalysisRuleType']:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input['ConfiguredTableAnalysisRuleType']):
        pulumi.set(self, "type", value)


@pulumi.input_type
class ConfiguredTableGlueTableReferenceArgs:
    def __init__(__self__, *,
                 database_name: pulumi.Input[str],
                 table_name: pulumi.Input[str]):
        pulumi.set(__self__, "database_name", database_name)
        pulumi.set(__self__, "table_name", table_name)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "table_name")

    @table_name.setter
    def table_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "table_name", value)


@pulumi.input_type
class ConfiguredTableTableReferenceArgs:
    def __init__(__self__, *,
                 glue: pulumi.Input['ConfiguredTableGlueTableReferenceArgs']):
        pulumi.set(__self__, "glue", glue)

    @property
    @pulumi.getter
    def glue(self) -> pulumi.Input['ConfiguredTableGlueTableReferenceArgs']:
        return pulumi.get(self, "glue")

    @glue.setter
    def glue(self, value: pulumi.Input['ConfiguredTableGlueTableReferenceArgs']):
        pulumi.set(self, "glue", value)


@pulumi.input_type
class MembershipPaymentConfigurationArgs:
    def __init__(__self__, *,
                 query_compute: pulumi.Input['MembershipQueryComputePaymentConfigArgs']):
        pulumi.set(__self__, "query_compute", query_compute)

    @property
    @pulumi.getter(name="queryCompute")
    def query_compute(self) -> pulumi.Input['MembershipQueryComputePaymentConfigArgs']:
        return pulumi.get(self, "query_compute")

    @query_compute.setter
    def query_compute(self, value: pulumi.Input['MembershipQueryComputePaymentConfigArgs']):
        pulumi.set(self, "query_compute", value)


@pulumi.input_type
class MembershipProtectedQueryOutputConfigurationArgs:
    def __init__(__self__, *,
                 s3: pulumi.Input['MembershipProtectedQueryS3OutputConfigurationArgs']):
        pulumi.set(__self__, "s3", s3)

    @property
    @pulumi.getter
    def s3(self) -> pulumi.Input['MembershipProtectedQueryS3OutputConfigurationArgs']:
        return pulumi.get(self, "s3")

    @s3.setter
    def s3(self, value: pulumi.Input['MembershipProtectedQueryS3OutputConfigurationArgs']):
        pulumi.set(self, "s3", value)


@pulumi.input_type
class MembershipProtectedQueryResultConfigurationArgs:
    def __init__(__self__, *,
                 output_configuration: pulumi.Input['MembershipProtectedQueryOutputConfigurationArgs'],
                 role_arn: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "output_configuration", output_configuration)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)

    @property
    @pulumi.getter(name="outputConfiguration")
    def output_configuration(self) -> pulumi.Input['MembershipProtectedQueryOutputConfigurationArgs']:
        return pulumi.get(self, "output_configuration")

    @output_configuration.setter
    def output_configuration(self, value: pulumi.Input['MembershipProtectedQueryOutputConfigurationArgs']):
        pulumi.set(self, "output_configuration", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_arn", value)


@pulumi.input_type
class MembershipProtectedQueryS3OutputConfigurationArgs:
    def __init__(__self__, *,
                 bucket: pulumi.Input[str],
                 result_format: pulumi.Input['MembershipResultFormat'],
                 key_prefix: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "bucket", bucket)
        pulumi.set(__self__, "result_format", result_format)
        if key_prefix is not None:
            pulumi.set(__self__, "key_prefix", key_prefix)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Input[str]:
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter(name="resultFormat")
    def result_format(self) -> pulumi.Input['MembershipResultFormat']:
        return pulumi.get(self, "result_format")

    @result_format.setter
    def result_format(self, value: pulumi.Input['MembershipResultFormat']):
        pulumi.set(self, "result_format", value)

    @property
    @pulumi.getter(name="keyPrefix")
    def key_prefix(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "key_prefix")

    @key_prefix.setter
    def key_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_prefix", value)


@pulumi.input_type
class MembershipQueryComputePaymentConfigArgs:
    def __init__(__self__, *,
                 is_responsible: pulumi.Input[bool]):
        pulumi.set(__self__, "is_responsible", is_responsible)

    @property
    @pulumi.getter(name="isResponsible")
    def is_responsible(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "is_responsible")

    @is_responsible.setter
    def is_responsible(self, value: pulumi.Input[bool]):
        pulumi.set(self, "is_responsible", value)


