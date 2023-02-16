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
    'DataCatalogTag',
    'WorkGroupConfiguration',
    'WorkGroupConfigurationUpdates',
    'WorkGroupEncryptionConfiguration',
    'WorkGroupEngineVersion',
    'WorkGroupResultConfiguration',
    'WorkGroupResultConfigurationUpdates',
    'WorkGroupTag',
]

@pulumi.output_type
class DataCatalogTag(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


@pulumi.output_type
class WorkGroupConfiguration(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "bytesScannedCutoffPerQuery":
            suggest = "bytes_scanned_cutoff_per_query"
        elif key == "enforceWorkGroupConfiguration":
            suggest = "enforce_work_group_configuration"
        elif key == "engineVersion":
            suggest = "engine_version"
        elif key == "publishCloudWatchMetricsEnabled":
            suggest = "publish_cloud_watch_metrics_enabled"
        elif key == "requesterPaysEnabled":
            suggest = "requester_pays_enabled"
        elif key == "resultConfiguration":
            suggest = "result_configuration"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in WorkGroupConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        WorkGroupConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        WorkGroupConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 bytes_scanned_cutoff_per_query: Optional[int] = None,
                 enforce_work_group_configuration: Optional[bool] = None,
                 engine_version: Optional['outputs.WorkGroupEngineVersion'] = None,
                 publish_cloud_watch_metrics_enabled: Optional[bool] = None,
                 requester_pays_enabled: Optional[bool] = None,
                 result_configuration: Optional['outputs.WorkGroupResultConfiguration'] = None):
        if bytes_scanned_cutoff_per_query is not None:
            pulumi.set(__self__, "bytes_scanned_cutoff_per_query", bytes_scanned_cutoff_per_query)
        if enforce_work_group_configuration is not None:
            pulumi.set(__self__, "enforce_work_group_configuration", enforce_work_group_configuration)
        if engine_version is not None:
            pulumi.set(__self__, "engine_version", engine_version)
        if publish_cloud_watch_metrics_enabled is not None:
            pulumi.set(__self__, "publish_cloud_watch_metrics_enabled", publish_cloud_watch_metrics_enabled)
        if requester_pays_enabled is not None:
            pulumi.set(__self__, "requester_pays_enabled", requester_pays_enabled)
        if result_configuration is not None:
            pulumi.set(__self__, "result_configuration", result_configuration)

    @property
    @pulumi.getter(name="bytesScannedCutoffPerQuery")
    def bytes_scanned_cutoff_per_query(self) -> Optional[int]:
        return pulumi.get(self, "bytes_scanned_cutoff_per_query")

    @property
    @pulumi.getter(name="enforceWorkGroupConfiguration")
    def enforce_work_group_configuration(self) -> Optional[bool]:
        return pulumi.get(self, "enforce_work_group_configuration")

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> Optional['outputs.WorkGroupEngineVersion']:
        return pulumi.get(self, "engine_version")

    @property
    @pulumi.getter(name="publishCloudWatchMetricsEnabled")
    def publish_cloud_watch_metrics_enabled(self) -> Optional[bool]:
        return pulumi.get(self, "publish_cloud_watch_metrics_enabled")

    @property
    @pulumi.getter(name="requesterPaysEnabled")
    def requester_pays_enabled(self) -> Optional[bool]:
        return pulumi.get(self, "requester_pays_enabled")

    @property
    @pulumi.getter(name="resultConfiguration")
    def result_configuration(self) -> Optional['outputs.WorkGroupResultConfiguration']:
        return pulumi.get(self, "result_configuration")


@pulumi.output_type
class WorkGroupConfigurationUpdates(dict):
    """
    The configuration information that will be updated for this workgroup, which includes the location in Amazon S3 where query results are stored, the encryption option, if any, used for query results, whether the Amazon CloudWatch Metrics are enabled for the workgroup, whether the workgroup settings override the client-side settings, and the data usage limit for the amount of bytes scanned per query, if it is specified. 
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "bytesScannedCutoffPerQuery":
            suggest = "bytes_scanned_cutoff_per_query"
        elif key == "enforceWorkGroupConfiguration":
            suggest = "enforce_work_group_configuration"
        elif key == "engineVersion":
            suggest = "engine_version"
        elif key == "publishCloudWatchMetricsEnabled":
            suggest = "publish_cloud_watch_metrics_enabled"
        elif key == "removeBytesScannedCutoffPerQuery":
            suggest = "remove_bytes_scanned_cutoff_per_query"
        elif key == "requesterPaysEnabled":
            suggest = "requester_pays_enabled"
        elif key == "resultConfigurationUpdates":
            suggest = "result_configuration_updates"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in WorkGroupConfigurationUpdates. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        WorkGroupConfigurationUpdates.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        WorkGroupConfigurationUpdates.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 bytes_scanned_cutoff_per_query: Optional[int] = None,
                 enforce_work_group_configuration: Optional[bool] = None,
                 engine_version: Optional['outputs.WorkGroupEngineVersion'] = None,
                 publish_cloud_watch_metrics_enabled: Optional[bool] = None,
                 remove_bytes_scanned_cutoff_per_query: Optional[bool] = None,
                 requester_pays_enabled: Optional[bool] = None,
                 result_configuration_updates: Optional['outputs.WorkGroupResultConfigurationUpdates'] = None):
        """
        The configuration information that will be updated for this workgroup, which includes the location in Amazon S3 where query results are stored, the encryption option, if any, used for query results, whether the Amazon CloudWatch Metrics are enabled for the workgroup, whether the workgroup settings override the client-side settings, and the data usage limit for the amount of bytes scanned per query, if it is specified. 
        """
        if bytes_scanned_cutoff_per_query is not None:
            pulumi.set(__self__, "bytes_scanned_cutoff_per_query", bytes_scanned_cutoff_per_query)
        if enforce_work_group_configuration is not None:
            pulumi.set(__self__, "enforce_work_group_configuration", enforce_work_group_configuration)
        if engine_version is not None:
            pulumi.set(__self__, "engine_version", engine_version)
        if publish_cloud_watch_metrics_enabled is not None:
            pulumi.set(__self__, "publish_cloud_watch_metrics_enabled", publish_cloud_watch_metrics_enabled)
        if remove_bytes_scanned_cutoff_per_query is not None:
            pulumi.set(__self__, "remove_bytes_scanned_cutoff_per_query", remove_bytes_scanned_cutoff_per_query)
        if requester_pays_enabled is not None:
            pulumi.set(__self__, "requester_pays_enabled", requester_pays_enabled)
        if result_configuration_updates is not None:
            pulumi.set(__self__, "result_configuration_updates", result_configuration_updates)

    @property
    @pulumi.getter(name="bytesScannedCutoffPerQuery")
    def bytes_scanned_cutoff_per_query(self) -> Optional[int]:
        return pulumi.get(self, "bytes_scanned_cutoff_per_query")

    @property
    @pulumi.getter(name="enforceWorkGroupConfiguration")
    def enforce_work_group_configuration(self) -> Optional[bool]:
        return pulumi.get(self, "enforce_work_group_configuration")

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> Optional['outputs.WorkGroupEngineVersion']:
        return pulumi.get(self, "engine_version")

    @property
    @pulumi.getter(name="publishCloudWatchMetricsEnabled")
    def publish_cloud_watch_metrics_enabled(self) -> Optional[bool]:
        return pulumi.get(self, "publish_cloud_watch_metrics_enabled")

    @property
    @pulumi.getter(name="removeBytesScannedCutoffPerQuery")
    def remove_bytes_scanned_cutoff_per_query(self) -> Optional[bool]:
        return pulumi.get(self, "remove_bytes_scanned_cutoff_per_query")

    @property
    @pulumi.getter(name="requesterPaysEnabled")
    def requester_pays_enabled(self) -> Optional[bool]:
        return pulumi.get(self, "requester_pays_enabled")

    @property
    @pulumi.getter(name="resultConfigurationUpdates")
    def result_configuration_updates(self) -> Optional['outputs.WorkGroupResultConfigurationUpdates']:
        return pulumi.get(self, "result_configuration_updates")


@pulumi.output_type
class WorkGroupEncryptionConfiguration(dict):
    """
    If query results are encrypted in Amazon S3, indicates the encryption option used (for example, SSE-KMS or CSE-KMS) and key information.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "encryptionOption":
            suggest = "encryption_option"
        elif key == "kmsKey":
            suggest = "kms_key"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in WorkGroupEncryptionConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        WorkGroupEncryptionConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        WorkGroupEncryptionConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 encryption_option: 'WorkGroupEncryptionOption',
                 kms_key: Optional[str] = None):
        """
        If query results are encrypted in Amazon S3, indicates the encryption option used (for example, SSE-KMS or CSE-KMS) and key information.
        """
        pulumi.set(__self__, "encryption_option", encryption_option)
        if kms_key is not None:
            pulumi.set(__self__, "kms_key", kms_key)

    @property
    @pulumi.getter(name="encryptionOption")
    def encryption_option(self) -> 'WorkGroupEncryptionOption':
        return pulumi.get(self, "encryption_option")

    @property
    @pulumi.getter(name="kmsKey")
    def kms_key(self) -> Optional[str]:
        return pulumi.get(self, "kms_key")


@pulumi.output_type
class WorkGroupEngineVersion(dict):
    """
    The Athena engine version for running queries.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "effectiveEngineVersion":
            suggest = "effective_engine_version"
        elif key == "selectedEngineVersion":
            suggest = "selected_engine_version"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in WorkGroupEngineVersion. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        WorkGroupEngineVersion.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        WorkGroupEngineVersion.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 effective_engine_version: Optional[str] = None,
                 selected_engine_version: Optional[str] = None):
        """
        The Athena engine version for running queries.
        """
        if effective_engine_version is not None:
            pulumi.set(__self__, "effective_engine_version", effective_engine_version)
        if selected_engine_version is not None:
            pulumi.set(__self__, "selected_engine_version", selected_engine_version)

    @property
    @pulumi.getter(name="effectiveEngineVersion")
    def effective_engine_version(self) -> Optional[str]:
        return pulumi.get(self, "effective_engine_version")

    @property
    @pulumi.getter(name="selectedEngineVersion")
    def selected_engine_version(self) -> Optional[str]:
        return pulumi.get(self, "selected_engine_version")


@pulumi.output_type
class WorkGroupResultConfiguration(dict):
    """
    The location in Amazon S3 where query results are stored and the encryption option, if any, used for query results. These are known as "client-side settings". If workgroup settings override client-side settings, then the query uses the workgroup settings.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "encryptionConfiguration":
            suggest = "encryption_configuration"
        elif key == "outputLocation":
            suggest = "output_location"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in WorkGroupResultConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        WorkGroupResultConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        WorkGroupResultConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 encryption_configuration: Optional['outputs.WorkGroupEncryptionConfiguration'] = None,
                 output_location: Optional[str] = None):
        """
        The location in Amazon S3 where query results are stored and the encryption option, if any, used for query results. These are known as "client-side settings". If workgroup settings override client-side settings, then the query uses the workgroup settings.
        """
        if encryption_configuration is not None:
            pulumi.set(__self__, "encryption_configuration", encryption_configuration)
        if output_location is not None:
            pulumi.set(__self__, "output_location", output_location)

    @property
    @pulumi.getter(name="encryptionConfiguration")
    def encryption_configuration(self) -> Optional['outputs.WorkGroupEncryptionConfiguration']:
        return pulumi.get(self, "encryption_configuration")

    @property
    @pulumi.getter(name="outputLocation")
    def output_location(self) -> Optional[str]:
        return pulumi.get(self, "output_location")


@pulumi.output_type
class WorkGroupResultConfigurationUpdates(dict):
    """
    The result configuration information about the queries in this workgroup that will be updated. Includes the updated results location and an updated option for encrypting query results. 
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "encryptionConfiguration":
            suggest = "encryption_configuration"
        elif key == "outputLocation":
            suggest = "output_location"
        elif key == "removeEncryptionConfiguration":
            suggest = "remove_encryption_configuration"
        elif key == "removeOutputLocation":
            suggest = "remove_output_location"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in WorkGroupResultConfigurationUpdates. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        WorkGroupResultConfigurationUpdates.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        WorkGroupResultConfigurationUpdates.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 encryption_configuration: Optional['outputs.WorkGroupEncryptionConfiguration'] = None,
                 output_location: Optional[str] = None,
                 remove_encryption_configuration: Optional[bool] = None,
                 remove_output_location: Optional[bool] = None):
        """
        The result configuration information about the queries in this workgroup that will be updated. Includes the updated results location and an updated option for encrypting query results. 
        """
        if encryption_configuration is not None:
            pulumi.set(__self__, "encryption_configuration", encryption_configuration)
        if output_location is not None:
            pulumi.set(__self__, "output_location", output_location)
        if remove_encryption_configuration is not None:
            pulumi.set(__self__, "remove_encryption_configuration", remove_encryption_configuration)
        if remove_output_location is not None:
            pulumi.set(__self__, "remove_output_location", remove_output_location)

    @property
    @pulumi.getter(name="encryptionConfiguration")
    def encryption_configuration(self) -> Optional['outputs.WorkGroupEncryptionConfiguration']:
        return pulumi.get(self, "encryption_configuration")

    @property
    @pulumi.getter(name="outputLocation")
    def output_location(self) -> Optional[str]:
        return pulumi.get(self, "output_location")

    @property
    @pulumi.getter(name="removeEncryptionConfiguration")
    def remove_encryption_configuration(self) -> Optional[bool]:
        return pulumi.get(self, "remove_encryption_configuration")

    @property
    @pulumi.getter(name="removeOutputLocation")
    def remove_output_location(self) -> Optional[bool]:
        return pulumi.get(self, "remove_output_location")


@pulumi.output_type
class WorkGroupTag(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


