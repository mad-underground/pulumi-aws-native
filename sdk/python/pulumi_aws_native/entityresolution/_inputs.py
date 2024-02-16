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
    'IdMappingWorkflowIdMappingTechniquesArgs',
    'IdMappingWorkflowInputSourceArgs',
    'IdMappingWorkflowIntermediateSourceConfigurationArgs',
    'IdMappingWorkflowOutputSourceArgs',
    'IdMappingWorkflowProviderPropertiesArgs',
    'IdMappingWorkflowTagArgs',
    'MatchingWorkflowInputSourceArgs',
    'MatchingWorkflowIntermediateSourceConfigurationArgs',
    'MatchingWorkflowOutputAttributeArgs',
    'MatchingWorkflowOutputSourceArgs',
    'MatchingWorkflowProviderPropertiesArgs',
    'MatchingWorkflowResolutionTechniquesArgs',
    'MatchingWorkflowRuleBasedPropertiesArgs',
    'MatchingWorkflowRuleArgs',
    'MatchingWorkflowTagArgs',
    'SchemaMappingSchemaInputAttributeArgs',
    'SchemaMappingTagArgs',
]

@pulumi.input_type
class IdMappingWorkflowIdMappingTechniquesArgs:
    def __init__(__self__, *,
                 id_mapping_type: Optional[pulumi.Input['IdMappingWorkflowIdMappingTechniquesIdMappingType']] = None,
                 provider_properties: Optional[pulumi.Input['IdMappingWorkflowProviderPropertiesArgs']] = None):
        if id_mapping_type is not None:
            pulumi.set(__self__, "id_mapping_type", id_mapping_type)
        if provider_properties is not None:
            pulumi.set(__self__, "provider_properties", provider_properties)

    @property
    @pulumi.getter(name="idMappingType")
    def id_mapping_type(self) -> Optional[pulumi.Input['IdMappingWorkflowIdMappingTechniquesIdMappingType']]:
        return pulumi.get(self, "id_mapping_type")

    @id_mapping_type.setter
    def id_mapping_type(self, value: Optional[pulumi.Input['IdMappingWorkflowIdMappingTechniquesIdMappingType']]):
        pulumi.set(self, "id_mapping_type", value)

    @property
    @pulumi.getter(name="providerProperties")
    def provider_properties(self) -> Optional[pulumi.Input['IdMappingWorkflowProviderPropertiesArgs']]:
        return pulumi.get(self, "provider_properties")

    @provider_properties.setter
    def provider_properties(self, value: Optional[pulumi.Input['IdMappingWorkflowProviderPropertiesArgs']]):
        pulumi.set(self, "provider_properties", value)


@pulumi.input_type
class IdMappingWorkflowInputSourceArgs:
    def __init__(__self__, *,
                 input_source_arn: pulumi.Input[str],
                 schema_arn: pulumi.Input[str]):
        """
        :param pulumi.Input[str] input_source_arn: An Glue table ARN for the input source table
        """
        pulumi.set(__self__, "input_source_arn", input_source_arn)
        pulumi.set(__self__, "schema_arn", schema_arn)

    @property
    @pulumi.getter(name="inputSourceArn")
    def input_source_arn(self) -> pulumi.Input[str]:
        """
        An Glue table ARN for the input source table
        """
        return pulumi.get(self, "input_source_arn")

    @input_source_arn.setter
    def input_source_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "input_source_arn", value)

    @property
    @pulumi.getter(name="schemaArn")
    def schema_arn(self) -> pulumi.Input[str]:
        return pulumi.get(self, "schema_arn")

    @schema_arn.setter
    def schema_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "schema_arn", value)


@pulumi.input_type
class IdMappingWorkflowIntermediateSourceConfigurationArgs:
    def __init__(__self__, *,
                 intermediate_s3_path: pulumi.Input[str]):
        """
        :param pulumi.Input[str] intermediate_s3_path: The s3 path that would be used to stage the intermediate data being generated during workflow execution.
        """
        pulumi.set(__self__, "intermediate_s3_path", intermediate_s3_path)

    @property
    @pulumi.getter(name="intermediateS3Path")
    def intermediate_s3_path(self) -> pulumi.Input[str]:
        """
        The s3 path that would be used to stage the intermediate data being generated during workflow execution.
        """
        return pulumi.get(self, "intermediate_s3_path")

    @intermediate_s3_path.setter
    def intermediate_s3_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "intermediate_s3_path", value)


@pulumi.input_type
class IdMappingWorkflowOutputSourceArgs:
    def __init__(__self__, *,
                 output_s3_path: pulumi.Input[str],
                 kms_arn: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] output_s3_path: The S3 path to which Entity Resolution will write the output table
        """
        pulumi.set(__self__, "output_s3_path", output_s3_path)
        if kms_arn is not None:
            pulumi.set(__self__, "kms_arn", kms_arn)

    @property
    @pulumi.getter(name="outputS3Path")
    def output_s3_path(self) -> pulumi.Input[str]:
        """
        The S3 path to which Entity Resolution will write the output table
        """
        return pulumi.get(self, "output_s3_path")

    @output_s3_path.setter
    def output_s3_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "output_s3_path", value)

    @property
    @pulumi.getter(name="kmsArn")
    def kms_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "kms_arn")

    @kms_arn.setter
    def kms_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_arn", value)


@pulumi.input_type
class IdMappingWorkflowProviderPropertiesArgs:
    def __init__(__self__, *,
                 provider_service_arn: pulumi.Input[str],
                 intermediate_source_configuration: Optional[pulumi.Input['IdMappingWorkflowIntermediateSourceConfigurationArgs']] = None,
                 provider_configuration: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[str] provider_service_arn: Arn of the Provider Service being used.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] provider_configuration: Additional Provider configuration that would be required for the provider service. The Configuration must be in JSON string format
        """
        pulumi.set(__self__, "provider_service_arn", provider_service_arn)
        if intermediate_source_configuration is not None:
            pulumi.set(__self__, "intermediate_source_configuration", intermediate_source_configuration)
        if provider_configuration is not None:
            pulumi.set(__self__, "provider_configuration", provider_configuration)

    @property
    @pulumi.getter(name="providerServiceArn")
    def provider_service_arn(self) -> pulumi.Input[str]:
        """
        Arn of the Provider Service being used.
        """
        return pulumi.get(self, "provider_service_arn")

    @provider_service_arn.setter
    def provider_service_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "provider_service_arn", value)

    @property
    @pulumi.getter(name="intermediateSourceConfiguration")
    def intermediate_source_configuration(self) -> Optional[pulumi.Input['IdMappingWorkflowIntermediateSourceConfigurationArgs']]:
        return pulumi.get(self, "intermediate_source_configuration")

    @intermediate_source_configuration.setter
    def intermediate_source_configuration(self, value: Optional[pulumi.Input['IdMappingWorkflowIntermediateSourceConfigurationArgs']]):
        pulumi.set(self, "intermediate_source_configuration", value)

    @property
    @pulumi.getter(name="providerConfiguration")
    def provider_configuration(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Additional Provider configuration that would be required for the provider service. The Configuration must be in JSON string format
        """
        return pulumi.get(self, "provider_configuration")

    @provider_configuration.setter
    def provider_configuration(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "provider_configuration", value)


@pulumi.input_type
class IdMappingWorkflowTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        A key-value pair to associate with a resource
        :param pulumi.Input[str] key: The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        :param pulumi.Input[str] value: The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class MatchingWorkflowInputSourceArgs:
    def __init__(__self__, *,
                 input_source_arn: pulumi.Input[str],
                 schema_arn: pulumi.Input[str],
                 apply_normalization: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[str] input_source_arn: An Glue table ARN for the input source table
        """
        pulumi.set(__self__, "input_source_arn", input_source_arn)
        pulumi.set(__self__, "schema_arn", schema_arn)
        if apply_normalization is not None:
            pulumi.set(__self__, "apply_normalization", apply_normalization)

    @property
    @pulumi.getter(name="inputSourceArn")
    def input_source_arn(self) -> pulumi.Input[str]:
        """
        An Glue table ARN for the input source table
        """
        return pulumi.get(self, "input_source_arn")

    @input_source_arn.setter
    def input_source_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "input_source_arn", value)

    @property
    @pulumi.getter(name="schemaArn")
    def schema_arn(self) -> pulumi.Input[str]:
        return pulumi.get(self, "schema_arn")

    @schema_arn.setter
    def schema_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "schema_arn", value)

    @property
    @pulumi.getter(name="applyNormalization")
    def apply_normalization(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "apply_normalization")

    @apply_normalization.setter
    def apply_normalization(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "apply_normalization", value)


@pulumi.input_type
class MatchingWorkflowIntermediateSourceConfigurationArgs:
    def __init__(__self__, *,
                 intermediate_s3_path: pulumi.Input[str]):
        """
        :param pulumi.Input[str] intermediate_s3_path: The s3 path that would be used to stage the intermediate data being generated during workflow execution.
        """
        pulumi.set(__self__, "intermediate_s3_path", intermediate_s3_path)

    @property
    @pulumi.getter(name="intermediateS3Path")
    def intermediate_s3_path(self) -> pulumi.Input[str]:
        """
        The s3 path that would be used to stage the intermediate data being generated during workflow execution.
        """
        return pulumi.get(self, "intermediate_s3_path")

    @intermediate_s3_path.setter
    def intermediate_s3_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "intermediate_s3_path", value)


@pulumi.input_type
class MatchingWorkflowOutputAttributeArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 hashed: Optional[pulumi.Input[bool]] = None):
        pulumi.set(__self__, "name", name)
        if hashed is not None:
            pulumi.set(__self__, "hashed", hashed)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def hashed(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "hashed")

    @hashed.setter
    def hashed(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "hashed", value)


@pulumi.input_type
class MatchingWorkflowOutputSourceArgs:
    def __init__(__self__, *,
                 output: pulumi.Input[Sequence[pulumi.Input['MatchingWorkflowOutputAttributeArgs']]],
                 output_s3_path: pulumi.Input[str],
                 apply_normalization: Optional[pulumi.Input[bool]] = None,
                 kms_arn: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] output_s3_path: The S3 path to which Entity Resolution will write the output table
        """
        pulumi.set(__self__, "output", output)
        pulumi.set(__self__, "output_s3_path", output_s3_path)
        if apply_normalization is not None:
            pulumi.set(__self__, "apply_normalization", apply_normalization)
        if kms_arn is not None:
            pulumi.set(__self__, "kms_arn", kms_arn)

    @property
    @pulumi.getter
    def output(self) -> pulumi.Input[Sequence[pulumi.Input['MatchingWorkflowOutputAttributeArgs']]]:
        return pulumi.get(self, "output")

    @output.setter
    def output(self, value: pulumi.Input[Sequence[pulumi.Input['MatchingWorkflowOutputAttributeArgs']]]):
        pulumi.set(self, "output", value)

    @property
    @pulumi.getter(name="outputS3Path")
    def output_s3_path(self) -> pulumi.Input[str]:
        """
        The S3 path to which Entity Resolution will write the output table
        """
        return pulumi.get(self, "output_s3_path")

    @output_s3_path.setter
    def output_s3_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "output_s3_path", value)

    @property
    @pulumi.getter(name="applyNormalization")
    def apply_normalization(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "apply_normalization")

    @apply_normalization.setter
    def apply_normalization(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "apply_normalization", value)

    @property
    @pulumi.getter(name="kmsArn")
    def kms_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "kms_arn")

    @kms_arn.setter
    def kms_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_arn", value)


@pulumi.input_type
class MatchingWorkflowProviderPropertiesArgs:
    def __init__(__self__, *,
                 provider_service_arn: pulumi.Input[str],
                 intermediate_source_configuration: Optional[pulumi.Input['MatchingWorkflowIntermediateSourceConfigurationArgs']] = None,
                 provider_configuration: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[str] provider_service_arn: Arn of the Provider service being used.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] provider_configuration: Additional Provider configuration that would be required for the provider service. The Configuration must be in JSON string format
        """
        pulumi.set(__self__, "provider_service_arn", provider_service_arn)
        if intermediate_source_configuration is not None:
            pulumi.set(__self__, "intermediate_source_configuration", intermediate_source_configuration)
        if provider_configuration is not None:
            pulumi.set(__self__, "provider_configuration", provider_configuration)

    @property
    @pulumi.getter(name="providerServiceArn")
    def provider_service_arn(self) -> pulumi.Input[str]:
        """
        Arn of the Provider service being used.
        """
        return pulumi.get(self, "provider_service_arn")

    @provider_service_arn.setter
    def provider_service_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "provider_service_arn", value)

    @property
    @pulumi.getter(name="intermediateSourceConfiguration")
    def intermediate_source_configuration(self) -> Optional[pulumi.Input['MatchingWorkflowIntermediateSourceConfigurationArgs']]:
        return pulumi.get(self, "intermediate_source_configuration")

    @intermediate_source_configuration.setter
    def intermediate_source_configuration(self, value: Optional[pulumi.Input['MatchingWorkflowIntermediateSourceConfigurationArgs']]):
        pulumi.set(self, "intermediate_source_configuration", value)

    @property
    @pulumi.getter(name="providerConfiguration")
    def provider_configuration(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Additional Provider configuration that would be required for the provider service. The Configuration must be in JSON string format
        """
        return pulumi.get(self, "provider_configuration")

    @provider_configuration.setter
    def provider_configuration(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "provider_configuration", value)


@pulumi.input_type
class MatchingWorkflowResolutionTechniquesArgs:
    def __init__(__self__, *,
                 provider_properties: Optional[pulumi.Input['MatchingWorkflowProviderPropertiesArgs']] = None,
                 resolution_type: Optional[pulumi.Input['MatchingWorkflowResolutionTechniquesResolutionType']] = None,
                 rule_based_properties: Optional[pulumi.Input['MatchingWorkflowRuleBasedPropertiesArgs']] = None):
        if provider_properties is not None:
            pulumi.set(__self__, "provider_properties", provider_properties)
        if resolution_type is not None:
            pulumi.set(__self__, "resolution_type", resolution_type)
        if rule_based_properties is not None:
            pulumi.set(__self__, "rule_based_properties", rule_based_properties)

    @property
    @pulumi.getter(name="providerProperties")
    def provider_properties(self) -> Optional[pulumi.Input['MatchingWorkflowProviderPropertiesArgs']]:
        return pulumi.get(self, "provider_properties")

    @provider_properties.setter
    def provider_properties(self, value: Optional[pulumi.Input['MatchingWorkflowProviderPropertiesArgs']]):
        pulumi.set(self, "provider_properties", value)

    @property
    @pulumi.getter(name="resolutionType")
    def resolution_type(self) -> Optional[pulumi.Input['MatchingWorkflowResolutionTechniquesResolutionType']]:
        return pulumi.get(self, "resolution_type")

    @resolution_type.setter
    def resolution_type(self, value: Optional[pulumi.Input['MatchingWorkflowResolutionTechniquesResolutionType']]):
        pulumi.set(self, "resolution_type", value)

    @property
    @pulumi.getter(name="ruleBasedProperties")
    def rule_based_properties(self) -> Optional[pulumi.Input['MatchingWorkflowRuleBasedPropertiesArgs']]:
        return pulumi.get(self, "rule_based_properties")

    @rule_based_properties.setter
    def rule_based_properties(self, value: Optional[pulumi.Input['MatchingWorkflowRuleBasedPropertiesArgs']]):
        pulumi.set(self, "rule_based_properties", value)


@pulumi.input_type
class MatchingWorkflowRuleBasedPropertiesArgs:
    def __init__(__self__, *,
                 attribute_matching_model: pulumi.Input['MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel'],
                 rules: pulumi.Input[Sequence[pulumi.Input['MatchingWorkflowRuleArgs']]]):
        pulumi.set(__self__, "attribute_matching_model", attribute_matching_model)
        pulumi.set(__self__, "rules", rules)

    @property
    @pulumi.getter(name="attributeMatchingModel")
    def attribute_matching_model(self) -> pulumi.Input['MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel']:
        return pulumi.get(self, "attribute_matching_model")

    @attribute_matching_model.setter
    def attribute_matching_model(self, value: pulumi.Input['MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel']):
        pulumi.set(self, "attribute_matching_model", value)

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Input[Sequence[pulumi.Input['MatchingWorkflowRuleArgs']]]:
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: pulumi.Input[Sequence[pulumi.Input['MatchingWorkflowRuleArgs']]]):
        pulumi.set(self, "rules", value)


@pulumi.input_type
class MatchingWorkflowRuleArgs:
    def __init__(__self__, *,
                 matching_keys: pulumi.Input[Sequence[pulumi.Input[str]]],
                 rule_name: pulumi.Input[str]):
        pulumi.set(__self__, "matching_keys", matching_keys)
        pulumi.set(__self__, "rule_name", rule_name)

    @property
    @pulumi.getter(name="matchingKeys")
    def matching_keys(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "matching_keys")

    @matching_keys.setter
    def matching_keys(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "matching_keys", value)

    @property
    @pulumi.getter(name="ruleName")
    def rule_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "rule_name")

    @rule_name.setter
    def rule_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_name", value)


@pulumi.input_type
class MatchingWorkflowTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        A key-value pair to associate with a resource
        :param pulumi.Input[str] key: The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        :param pulumi.Input[str] value: The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class SchemaMappingSchemaInputAttributeArgs:
    def __init__(__self__, *,
                 field_name: pulumi.Input[str],
                 type: pulumi.Input['SchemaMappingSchemaAttributeType'],
                 group_name: Optional[pulumi.Input[str]] = None,
                 match_key: Optional[pulumi.Input[str]] = None,
                 sub_type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] sub_type: The subtype of the Attribute. Would be required only when type is PROVIDER_ID
        """
        pulumi.set(__self__, "field_name", field_name)
        pulumi.set(__self__, "type", type)
        if group_name is not None:
            pulumi.set(__self__, "group_name", group_name)
        if match_key is not None:
            pulumi.set(__self__, "match_key", match_key)
        if sub_type is not None:
            pulumi.set(__self__, "sub_type", sub_type)

    @property
    @pulumi.getter(name="fieldName")
    def field_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "field_name")

    @field_name.setter
    def field_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "field_name", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input['SchemaMappingSchemaAttributeType']:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input['SchemaMappingSchemaAttributeType']):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "group_name")

    @group_name.setter
    def group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_name", value)

    @property
    @pulumi.getter(name="matchKey")
    def match_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "match_key")

    @match_key.setter
    def match_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "match_key", value)

    @property
    @pulumi.getter(name="subType")
    def sub_type(self) -> Optional[pulumi.Input[str]]:
        """
        The subtype of the Attribute. Would be required only when type is PROVIDER_ID
        """
        return pulumi.get(self, "sub_type")

    @sub_type.setter
    def sub_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sub_type", value)


@pulumi.input_type
class SchemaMappingTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        A key-value pair to associate with a resource
        :param pulumi.Input[str] key: The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        :param pulumi.Input[str] value: The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


