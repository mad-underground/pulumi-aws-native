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
from ._enums import *

__all__ = [
    'DocumentClassifierAugmentedManifestsListItem',
    'DocumentClassifierDocumentReaderConfig',
    'DocumentClassifierDocuments',
    'DocumentClassifierInputDataConfig',
    'DocumentClassifierOutputDataConfig',
    'DocumentClassifierVpcConfig',
    'FlywheelDataSecurityConfig',
    'FlywheelDocumentClassificationConfig',
    'FlywheelEntityRecognitionConfig',
    'FlywheelEntityTypesListItem',
    'FlywheelTaskConfig',
    'FlywheelVpcConfig',
]

@pulumi.output_type
class DocumentClassifierAugmentedManifestsListItem(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "attributeNames":
            suggest = "attribute_names"
        elif key == "s3Uri":
            suggest = "s3_uri"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DocumentClassifierAugmentedManifestsListItem. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DocumentClassifierAugmentedManifestsListItem.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DocumentClassifierAugmentedManifestsListItem.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 attribute_names: Sequence[str],
                 s3_uri: str,
                 split: Optional['DocumentClassifierAugmentedManifestsListItemSplit'] = None):
        pulumi.set(__self__, "attribute_names", attribute_names)
        pulumi.set(__self__, "s3_uri", s3_uri)
        if split is not None:
            pulumi.set(__self__, "split", split)

    @property
    @pulumi.getter(name="attributeNames")
    def attribute_names(self) -> Sequence[str]:
        return pulumi.get(self, "attribute_names")

    @property
    @pulumi.getter(name="s3Uri")
    def s3_uri(self) -> str:
        return pulumi.get(self, "s3_uri")

    @property
    @pulumi.getter
    def split(self) -> Optional['DocumentClassifierAugmentedManifestsListItemSplit']:
        return pulumi.get(self, "split")


@pulumi.output_type
class DocumentClassifierDocumentReaderConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "documentReadAction":
            suggest = "document_read_action"
        elif key == "documentReadMode":
            suggest = "document_read_mode"
        elif key == "featureTypes":
            suggest = "feature_types"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DocumentClassifierDocumentReaderConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DocumentClassifierDocumentReaderConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DocumentClassifierDocumentReaderConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 document_read_action: 'DocumentClassifierDocumentReaderConfigDocumentReadAction',
                 document_read_mode: Optional['DocumentClassifierDocumentReaderConfigDocumentReadMode'] = None,
                 feature_types: Optional[Sequence['DocumentClassifierDocumentReaderConfigFeatureTypesItem']] = None):
        pulumi.set(__self__, "document_read_action", document_read_action)
        if document_read_mode is not None:
            pulumi.set(__self__, "document_read_mode", document_read_mode)
        if feature_types is not None:
            pulumi.set(__self__, "feature_types", feature_types)

    @property
    @pulumi.getter(name="documentReadAction")
    def document_read_action(self) -> 'DocumentClassifierDocumentReaderConfigDocumentReadAction':
        return pulumi.get(self, "document_read_action")

    @property
    @pulumi.getter(name="documentReadMode")
    def document_read_mode(self) -> Optional['DocumentClassifierDocumentReaderConfigDocumentReadMode']:
        return pulumi.get(self, "document_read_mode")

    @property
    @pulumi.getter(name="featureTypes")
    def feature_types(self) -> Optional[Sequence['DocumentClassifierDocumentReaderConfigFeatureTypesItem']]:
        return pulumi.get(self, "feature_types")


@pulumi.output_type
class DocumentClassifierDocuments(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "s3Uri":
            suggest = "s3_uri"
        elif key == "testS3Uri":
            suggest = "test_s3_uri"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DocumentClassifierDocuments. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DocumentClassifierDocuments.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DocumentClassifierDocuments.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 s3_uri: str,
                 test_s3_uri: Optional[str] = None):
        pulumi.set(__self__, "s3_uri", s3_uri)
        if test_s3_uri is not None:
            pulumi.set(__self__, "test_s3_uri", test_s3_uri)

    @property
    @pulumi.getter(name="s3Uri")
    def s3_uri(self) -> str:
        return pulumi.get(self, "s3_uri")

    @property
    @pulumi.getter(name="testS3Uri")
    def test_s3_uri(self) -> Optional[str]:
        return pulumi.get(self, "test_s3_uri")


@pulumi.output_type
class DocumentClassifierInputDataConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "augmentedManifests":
            suggest = "augmented_manifests"
        elif key == "dataFormat":
            suggest = "data_format"
        elif key == "documentReaderConfig":
            suggest = "document_reader_config"
        elif key == "documentType":
            suggest = "document_type"
        elif key == "labelDelimiter":
            suggest = "label_delimiter"
        elif key == "s3Uri":
            suggest = "s3_uri"
        elif key == "testS3Uri":
            suggest = "test_s3_uri"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DocumentClassifierInputDataConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DocumentClassifierInputDataConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DocumentClassifierInputDataConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 augmented_manifests: Optional[Sequence['outputs.DocumentClassifierAugmentedManifestsListItem']] = None,
                 data_format: Optional['DocumentClassifierInputDataConfigDataFormat'] = None,
                 document_reader_config: Optional['outputs.DocumentClassifierDocumentReaderConfig'] = None,
                 document_type: Optional['DocumentClassifierInputDataConfigDocumentType'] = None,
                 documents: Optional['outputs.DocumentClassifierDocuments'] = None,
                 label_delimiter: Optional[str] = None,
                 s3_uri: Optional[str] = None,
                 test_s3_uri: Optional[str] = None):
        if augmented_manifests is not None:
            pulumi.set(__self__, "augmented_manifests", augmented_manifests)
        if data_format is not None:
            pulumi.set(__self__, "data_format", data_format)
        if document_reader_config is not None:
            pulumi.set(__self__, "document_reader_config", document_reader_config)
        if document_type is not None:
            pulumi.set(__self__, "document_type", document_type)
        if documents is not None:
            pulumi.set(__self__, "documents", documents)
        if label_delimiter is not None:
            pulumi.set(__self__, "label_delimiter", label_delimiter)
        if s3_uri is not None:
            pulumi.set(__self__, "s3_uri", s3_uri)
        if test_s3_uri is not None:
            pulumi.set(__self__, "test_s3_uri", test_s3_uri)

    @property
    @pulumi.getter(name="augmentedManifests")
    def augmented_manifests(self) -> Optional[Sequence['outputs.DocumentClassifierAugmentedManifestsListItem']]:
        return pulumi.get(self, "augmented_manifests")

    @property
    @pulumi.getter(name="dataFormat")
    def data_format(self) -> Optional['DocumentClassifierInputDataConfigDataFormat']:
        return pulumi.get(self, "data_format")

    @property
    @pulumi.getter(name="documentReaderConfig")
    def document_reader_config(self) -> Optional['outputs.DocumentClassifierDocumentReaderConfig']:
        return pulumi.get(self, "document_reader_config")

    @property
    @pulumi.getter(name="documentType")
    def document_type(self) -> Optional['DocumentClassifierInputDataConfigDocumentType']:
        return pulumi.get(self, "document_type")

    @property
    @pulumi.getter
    def documents(self) -> Optional['outputs.DocumentClassifierDocuments']:
        return pulumi.get(self, "documents")

    @property
    @pulumi.getter(name="labelDelimiter")
    def label_delimiter(self) -> Optional[str]:
        return pulumi.get(self, "label_delimiter")

    @property
    @pulumi.getter(name="s3Uri")
    def s3_uri(self) -> Optional[str]:
        return pulumi.get(self, "s3_uri")

    @property
    @pulumi.getter(name="testS3Uri")
    def test_s3_uri(self) -> Optional[str]:
        return pulumi.get(self, "test_s3_uri")


@pulumi.output_type
class DocumentClassifierOutputDataConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "kmsKeyId":
            suggest = "kms_key_id"
        elif key == "s3Uri":
            suggest = "s3_uri"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DocumentClassifierOutputDataConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DocumentClassifierOutputDataConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DocumentClassifierOutputDataConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 kms_key_id: Optional[str] = None,
                 s3_uri: Optional[str] = None):
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)
        if s3_uri is not None:
            pulumi.set(__self__, "s3_uri", s3_uri)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[str]:
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter(name="s3Uri")
    def s3_uri(self) -> Optional[str]:
        return pulumi.get(self, "s3_uri")


@pulumi.output_type
class DocumentClassifierVpcConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "securityGroupIds":
            suggest = "security_group_ids"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DocumentClassifierVpcConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DocumentClassifierVpcConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DocumentClassifierVpcConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 security_group_ids: Sequence[str],
                 subnets: Sequence[str]):
        pulumi.set(__self__, "security_group_ids", security_group_ids)
        pulumi.set(__self__, "subnets", subnets)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Sequence[str]:
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter
    def subnets(self) -> Sequence[str]:
        return pulumi.get(self, "subnets")


@pulumi.output_type
class FlywheelDataSecurityConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "dataLakeKmsKeyId":
            suggest = "data_lake_kms_key_id"
        elif key == "modelKmsKeyId":
            suggest = "model_kms_key_id"
        elif key == "volumeKmsKeyId":
            suggest = "volume_kms_key_id"
        elif key == "vpcConfig":
            suggest = "vpc_config"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FlywheelDataSecurityConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FlywheelDataSecurityConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FlywheelDataSecurityConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 data_lake_kms_key_id: Optional[str] = None,
                 model_kms_key_id: Optional[str] = None,
                 volume_kms_key_id: Optional[str] = None,
                 vpc_config: Optional['outputs.FlywheelVpcConfig'] = None):
        if data_lake_kms_key_id is not None:
            pulumi.set(__self__, "data_lake_kms_key_id", data_lake_kms_key_id)
        if model_kms_key_id is not None:
            pulumi.set(__self__, "model_kms_key_id", model_kms_key_id)
        if volume_kms_key_id is not None:
            pulumi.set(__self__, "volume_kms_key_id", volume_kms_key_id)
        if vpc_config is not None:
            pulumi.set(__self__, "vpc_config", vpc_config)

    @property
    @pulumi.getter(name="dataLakeKmsKeyId")
    def data_lake_kms_key_id(self) -> Optional[str]:
        return pulumi.get(self, "data_lake_kms_key_id")

    @property
    @pulumi.getter(name="modelKmsKeyId")
    def model_kms_key_id(self) -> Optional[str]:
        return pulumi.get(self, "model_kms_key_id")

    @property
    @pulumi.getter(name="volumeKmsKeyId")
    def volume_kms_key_id(self) -> Optional[str]:
        return pulumi.get(self, "volume_kms_key_id")

    @property
    @pulumi.getter(name="vpcConfig")
    def vpc_config(self) -> Optional['outputs.FlywheelVpcConfig']:
        return pulumi.get(self, "vpc_config")


@pulumi.output_type
class FlywheelDocumentClassificationConfig(dict):
    def __init__(__self__, *,
                 mode: 'FlywheelDocumentClassificationConfigMode',
                 labels: Optional[Sequence[str]] = None):
        pulumi.set(__self__, "mode", mode)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)

    @property
    @pulumi.getter
    def mode(self) -> 'FlywheelDocumentClassificationConfigMode':
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def labels(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "labels")


@pulumi.output_type
class FlywheelEntityRecognitionConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "entityTypes":
            suggest = "entity_types"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FlywheelEntityRecognitionConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FlywheelEntityRecognitionConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FlywheelEntityRecognitionConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 entity_types: Optional[Sequence['outputs.FlywheelEntityTypesListItem']] = None):
        if entity_types is not None:
            pulumi.set(__self__, "entity_types", entity_types)

    @property
    @pulumi.getter(name="entityTypes")
    def entity_types(self) -> Optional[Sequence['outputs.FlywheelEntityTypesListItem']]:
        return pulumi.get(self, "entity_types")


@pulumi.output_type
class FlywheelEntityTypesListItem(dict):
    def __init__(__self__, *,
                 type: str):
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")


@pulumi.output_type
class FlywheelTaskConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "languageCode":
            suggest = "language_code"
        elif key == "documentClassificationConfig":
            suggest = "document_classification_config"
        elif key == "entityRecognitionConfig":
            suggest = "entity_recognition_config"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FlywheelTaskConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FlywheelTaskConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FlywheelTaskConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 language_code: 'FlywheelTaskConfigLanguageCode',
                 document_classification_config: Optional['outputs.FlywheelDocumentClassificationConfig'] = None,
                 entity_recognition_config: Optional['outputs.FlywheelEntityRecognitionConfig'] = None):
        pulumi.set(__self__, "language_code", language_code)
        if document_classification_config is not None:
            pulumi.set(__self__, "document_classification_config", document_classification_config)
        if entity_recognition_config is not None:
            pulumi.set(__self__, "entity_recognition_config", entity_recognition_config)

    @property
    @pulumi.getter(name="languageCode")
    def language_code(self) -> 'FlywheelTaskConfigLanguageCode':
        return pulumi.get(self, "language_code")

    @property
    @pulumi.getter(name="documentClassificationConfig")
    def document_classification_config(self) -> Optional['outputs.FlywheelDocumentClassificationConfig']:
        return pulumi.get(self, "document_classification_config")

    @property
    @pulumi.getter(name="entityRecognitionConfig")
    def entity_recognition_config(self) -> Optional['outputs.FlywheelEntityRecognitionConfig']:
        return pulumi.get(self, "entity_recognition_config")


@pulumi.output_type
class FlywheelVpcConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "securityGroupIds":
            suggest = "security_group_ids"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FlywheelVpcConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FlywheelVpcConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FlywheelVpcConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 security_group_ids: Sequence[str],
                 subnets: Sequence[str]):
        pulumi.set(__self__, "security_group_ids", security_group_ids)
        pulumi.set(__self__, "subnets", subnets)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Sequence[str]:
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter
    def subnets(self) -> Sequence[str]:
        return pulumi.get(self, "subnets")


