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
    'AssistantAssociationAssociationDataArgs',
    'AssistantAssociationTagArgs',
    'AssistantServerSideEncryptionConfigurationArgs',
    'AssistantTagArgs',
    'KnowledgeBaseAppIntegrationsConfigurationArgs',
    'KnowledgeBaseRenderingConfigurationArgs',
    'KnowledgeBaseServerSideEncryptionConfigurationArgs',
    'KnowledgeBaseSourceConfigurationArgs',
    'KnowledgeBaseTagArgs',
]

@pulumi.input_type
class AssistantAssociationAssociationDataArgs:
    def __init__(__self__, *,
                 knowledge_base_id: pulumi.Input[str]):
        pulumi.set(__self__, "knowledge_base_id", knowledge_base_id)

    @property
    @pulumi.getter(name="knowledgeBaseId")
    def knowledge_base_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "knowledge_base_id")

    @knowledge_base_id.setter
    def knowledge_base_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "knowledge_base_id", value)


@pulumi.input_type
class AssistantAssociationTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class AssistantServerSideEncryptionConfigurationArgs:
    def __init__(__self__, *,
                 kms_key_id: Optional[pulumi.Input[str]] = None):
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_id", value)


@pulumi.input_type
class AssistantTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class KnowledgeBaseAppIntegrationsConfigurationArgs:
    def __init__(__self__, *,
                 app_integration_arn: pulumi.Input[str],
                 object_fields: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(__self__, "app_integration_arn", app_integration_arn)
        pulumi.set(__self__, "object_fields", object_fields)

    @property
    @pulumi.getter(name="appIntegrationArn")
    def app_integration_arn(self) -> pulumi.Input[str]:
        return pulumi.get(self, "app_integration_arn")

    @app_integration_arn.setter
    def app_integration_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "app_integration_arn", value)

    @property
    @pulumi.getter(name="objectFields")
    def object_fields(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "object_fields")

    @object_fields.setter
    def object_fields(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "object_fields", value)


@pulumi.input_type
class KnowledgeBaseRenderingConfigurationArgs:
    def __init__(__self__, *,
                 template_uri: Optional[pulumi.Input[str]] = None):
        if template_uri is not None:
            pulumi.set(__self__, "template_uri", template_uri)

    @property
    @pulumi.getter(name="templateUri")
    def template_uri(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "template_uri")

    @template_uri.setter
    def template_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template_uri", value)


@pulumi.input_type
class KnowledgeBaseServerSideEncryptionConfigurationArgs:
    def __init__(__self__, *,
                 kms_key_id: Optional[pulumi.Input[str]] = None):
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_id", value)


@pulumi.input_type
class KnowledgeBaseSourceConfigurationArgs:
    def __init__(__self__, *,
                 app_integrations: Optional[pulumi.Input['KnowledgeBaseAppIntegrationsConfigurationArgs']] = None):
        if app_integrations is not None:
            pulumi.set(__self__, "app_integrations", app_integrations)

    @property
    @pulumi.getter(name="appIntegrations")
    def app_integrations(self) -> Optional[pulumi.Input['KnowledgeBaseAppIntegrationsConfigurationArgs']]:
        return pulumi.get(self, "app_integrations")

    @app_integrations.setter
    def app_integrations(self, value: Optional[pulumi.Input['KnowledgeBaseAppIntegrationsConfigurationArgs']]):
        pulumi.set(self, "app_integrations", value)


@pulumi.input_type
class KnowledgeBaseTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


