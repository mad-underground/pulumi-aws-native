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
    'ApplicationAutoStartConfiguration',
    'ApplicationAutoStopConfiguration',
    'ApplicationCloudWatchLoggingConfiguration',
    'ApplicationConfigurationObject',
    'ApplicationImageConfigurationInput',
    'ApplicationInitialCapacityConfig',
    'ApplicationInitialCapacityConfigKeyValuePair',
    'ApplicationLogTypeMapKeyValuePair',
    'ApplicationManagedPersistenceMonitoringConfiguration',
    'ApplicationMaximumAllowedResources',
    'ApplicationMonitoringConfiguration',
    'ApplicationNetworkConfiguration',
    'ApplicationS3MonitoringConfiguration',
    'ApplicationWorkerConfiguration',
    'ApplicationWorkerTypeSpecificationInput',
]

@pulumi.output_type
class ApplicationAutoStartConfiguration(dict):
    """
    Configuration for Auto Start of Application
    """
    def __init__(__self__, *,
                 enabled: Optional[bool] = None):
        """
        Configuration for Auto Start of Application
        :param bool enabled: If set to true, the Application will automatically start. Defaults to true.
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        If set to true, the Application will automatically start. Defaults to true.
        """
        return pulumi.get(self, "enabled")


@pulumi.output_type
class ApplicationAutoStopConfiguration(dict):
    """
    Configuration for Auto Stop of Application
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "idleTimeoutMinutes":
            suggest = "idle_timeout_minutes"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApplicationAutoStopConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApplicationAutoStopConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApplicationAutoStopConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 enabled: Optional[bool] = None,
                 idle_timeout_minutes: Optional[int] = None):
        """
        Configuration for Auto Stop of Application
        :param bool enabled: If set to true, the Application will automatically stop after being idle. Defaults to true.
        :param int idle_timeout_minutes: The amount of time [in minutes] to wait before auto stopping the Application when idle. Defaults to 15 minutes.
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if idle_timeout_minutes is not None:
            pulumi.set(__self__, "idle_timeout_minutes", idle_timeout_minutes)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        If set to true, the Application will automatically stop after being idle. Defaults to true.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="idleTimeoutMinutes")
    def idle_timeout_minutes(self) -> Optional[int]:
        """
        The amount of time [in minutes] to wait before auto stopping the Application when idle. Defaults to 15 minutes.
        """
        return pulumi.get(self, "idle_timeout_minutes")


@pulumi.output_type
class ApplicationCloudWatchLoggingConfiguration(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "encryptionKeyArn":
            suggest = "encryption_key_arn"
        elif key == "logGroupName":
            suggest = "log_group_name"
        elif key == "logStreamNamePrefix":
            suggest = "log_stream_name_prefix"
        elif key == "logTypeMap":
            suggest = "log_type_map"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApplicationCloudWatchLoggingConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApplicationCloudWatchLoggingConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApplicationCloudWatchLoggingConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 enabled: Optional[bool] = None,
                 encryption_key_arn: Optional[str] = None,
                 log_group_name: Optional[str] = None,
                 log_stream_name_prefix: Optional[str] = None,
                 log_type_map: Optional[Sequence['outputs.ApplicationLogTypeMapKeyValuePair']] = None):
        """
        :param bool enabled: If set to false, CloudWatch logging will be turned off. Defaults to false.
        :param str encryption_key_arn: KMS key ARN to encrypt the logs stored in given CloudWatch log-group.
        :param str log_group_name: Log-group name to produce log-streams on CloudWatch. If undefined, logs will be produced in a default log-group /aws/emr-serverless
        :param str log_stream_name_prefix: Log-stream name prefix by which log-stream names will start in the CloudWatch Log-group.
        :param Sequence['ApplicationLogTypeMapKeyValuePair'] log_type_map: The specific log-streams which need to be uploaded to CloudWatch.
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if encryption_key_arn is not None:
            pulumi.set(__self__, "encryption_key_arn", encryption_key_arn)
        if log_group_name is not None:
            pulumi.set(__self__, "log_group_name", log_group_name)
        if log_stream_name_prefix is not None:
            pulumi.set(__self__, "log_stream_name_prefix", log_stream_name_prefix)
        if log_type_map is not None:
            pulumi.set(__self__, "log_type_map", log_type_map)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        If set to false, CloudWatch logging will be turned off. Defaults to false.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="encryptionKeyArn")
    def encryption_key_arn(self) -> Optional[str]:
        """
        KMS key ARN to encrypt the logs stored in given CloudWatch log-group.
        """
        return pulumi.get(self, "encryption_key_arn")

    @property
    @pulumi.getter(name="logGroupName")
    def log_group_name(self) -> Optional[str]:
        """
        Log-group name to produce log-streams on CloudWatch. If undefined, logs will be produced in a default log-group /aws/emr-serverless
        """
        return pulumi.get(self, "log_group_name")

    @property
    @pulumi.getter(name="logStreamNamePrefix")
    def log_stream_name_prefix(self) -> Optional[str]:
        """
        Log-stream name prefix by which log-stream names will start in the CloudWatch Log-group.
        """
        return pulumi.get(self, "log_stream_name_prefix")

    @property
    @pulumi.getter(name="logTypeMap")
    def log_type_map(self) -> Optional[Sequence['outputs.ApplicationLogTypeMapKeyValuePair']]:
        """
        The specific log-streams which need to be uploaded to CloudWatch.
        """
        return pulumi.get(self, "log_type_map")


@pulumi.output_type
class ApplicationConfigurationObject(dict):
    """
    Configuration for a JobRun.
    """
    def __init__(__self__, *,
                 classification: str,
                 configurations: Optional[Sequence['outputs.ApplicationConfigurationObject']] = None,
                 properties: Optional[Mapping[str, str]] = None):
        """
        Configuration for a JobRun.
        :param str classification: String with a maximum length of 1024.
        """
        pulumi.set(__self__, "classification", classification)
        if configurations is not None:
            pulumi.set(__self__, "configurations", configurations)
        if properties is not None:
            pulumi.set(__self__, "properties", properties)

    @property
    @pulumi.getter
    def classification(self) -> str:
        """
        String with a maximum length of 1024.
        """
        return pulumi.get(self, "classification")

    @property
    @pulumi.getter
    def configurations(self) -> Optional[Sequence['outputs.ApplicationConfigurationObject']]:
        return pulumi.get(self, "configurations")

    @property
    @pulumi.getter
    def properties(self) -> Optional[Mapping[str, str]]:
        return pulumi.get(self, "properties")


@pulumi.output_type
class ApplicationImageConfigurationInput(dict):
    """
    The image configuration.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "imageUri":
            suggest = "image_uri"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApplicationImageConfigurationInput. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApplicationImageConfigurationInput.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApplicationImageConfigurationInput.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 image_uri: Optional[str] = None):
        """
        The image configuration.
        :param str image_uri: The URI of an image in the Amazon ECR registry. This field is required when you create a new application. If you leave this field blank in an update, Amazon EMR will remove the image configuration.
        """
        if image_uri is not None:
            pulumi.set(__self__, "image_uri", image_uri)

    @property
    @pulumi.getter(name="imageUri")
    def image_uri(self) -> Optional[str]:
        """
        The URI of an image in the Amazon ECR registry. This field is required when you create a new application. If you leave this field blank in an update, Amazon EMR will remove the image configuration.
        """
        return pulumi.get(self, "image_uri")


@pulumi.output_type
class ApplicationInitialCapacityConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "workerConfiguration":
            suggest = "worker_configuration"
        elif key == "workerCount":
            suggest = "worker_count"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApplicationInitialCapacityConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApplicationInitialCapacityConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApplicationInitialCapacityConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 worker_configuration: 'outputs.ApplicationWorkerConfiguration',
                 worker_count: int):
        """
        :param int worker_count: Initial count of workers to be initialized when an Application is started. This count will be continued to be maintained until the Application is stopped
        """
        pulumi.set(__self__, "worker_configuration", worker_configuration)
        pulumi.set(__self__, "worker_count", worker_count)

    @property
    @pulumi.getter(name="workerConfiguration")
    def worker_configuration(self) -> 'outputs.ApplicationWorkerConfiguration':
        return pulumi.get(self, "worker_configuration")

    @property
    @pulumi.getter(name="workerCount")
    def worker_count(self) -> int:
        """
        Initial count of workers to be initialized when an Application is started. This count will be continued to be maintained until the Application is stopped
        """
        return pulumi.get(self, "worker_count")


@pulumi.output_type
class ApplicationInitialCapacityConfigKeyValuePair(dict):
    def __init__(__self__, *,
                 key: str,
                 value: 'outputs.ApplicationInitialCapacityConfig'):
        """
        :param str key: Worker type for an analytics framework.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        Worker type for an analytics framework.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> 'outputs.ApplicationInitialCapacityConfig':
        return pulumi.get(self, "value")


@pulumi.output_type
class ApplicationLogTypeMapKeyValuePair(dict):
    def __init__(__self__, *,
                 key: str,
                 value: Sequence[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> Sequence[str]:
        return pulumi.get(self, "value")


@pulumi.output_type
class ApplicationManagedPersistenceMonitoringConfiguration(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "encryptionKeyArn":
            suggest = "encryption_key_arn"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApplicationManagedPersistenceMonitoringConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApplicationManagedPersistenceMonitoringConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApplicationManagedPersistenceMonitoringConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 enabled: Optional[bool] = None,
                 encryption_key_arn: Optional[str] = None):
        """
        :param bool enabled: If set to false, managed logging will be turned off. Defaults to true.
        :param str encryption_key_arn: KMS key ARN to encrypt the logs stored in managed persistence
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if encryption_key_arn is not None:
            pulumi.set(__self__, "encryption_key_arn", encryption_key_arn)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        If set to false, managed logging will be turned off. Defaults to true.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="encryptionKeyArn")
    def encryption_key_arn(self) -> Optional[str]:
        """
        KMS key ARN to encrypt the logs stored in managed persistence
        """
        return pulumi.get(self, "encryption_key_arn")


@pulumi.output_type
class ApplicationMaximumAllowedResources(dict):
    def __init__(__self__, *,
                 cpu: str,
                 memory: str,
                 disk: Optional[str] = None):
        """
        :param str cpu: Per worker CPU resource. vCPU is the only supported unit and specifying vCPU is optional.
        :param str memory: Per worker memory resource. GB is the only supported unit and specifying GB is optional.
        :param str disk: Per worker Disk resource. GB is the only supported unit and specifying GB is optional
        """
        pulumi.set(__self__, "cpu", cpu)
        pulumi.set(__self__, "memory", memory)
        if disk is not None:
            pulumi.set(__self__, "disk", disk)

    @property
    @pulumi.getter
    def cpu(self) -> str:
        """
        Per worker CPU resource. vCPU is the only supported unit and specifying vCPU is optional.
        """
        return pulumi.get(self, "cpu")

    @property
    @pulumi.getter
    def memory(self) -> str:
        """
        Per worker memory resource. GB is the only supported unit and specifying GB is optional.
        """
        return pulumi.get(self, "memory")

    @property
    @pulumi.getter
    def disk(self) -> Optional[str]:
        """
        Per worker Disk resource. GB is the only supported unit and specifying GB is optional
        """
        return pulumi.get(self, "disk")


@pulumi.output_type
class ApplicationMonitoringConfiguration(dict):
    """
    Monitoring configuration for batch and interactive JobRun.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "cloudWatchLoggingConfiguration":
            suggest = "cloud_watch_logging_configuration"
        elif key == "managedPersistenceMonitoringConfiguration":
            suggest = "managed_persistence_monitoring_configuration"
        elif key == "s3MonitoringConfiguration":
            suggest = "s3_monitoring_configuration"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApplicationMonitoringConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApplicationMonitoringConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApplicationMonitoringConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 cloud_watch_logging_configuration: Optional['outputs.ApplicationCloudWatchLoggingConfiguration'] = None,
                 managed_persistence_monitoring_configuration: Optional['outputs.ApplicationManagedPersistenceMonitoringConfiguration'] = None,
                 s3_monitoring_configuration: Optional['outputs.ApplicationS3MonitoringConfiguration'] = None):
        """
        Monitoring configuration for batch and interactive JobRun.
        :param 'ApplicationCloudWatchLoggingConfiguration' cloud_watch_logging_configuration: CloudWatch logging configurations for a JobRun.
        :param 'ApplicationManagedPersistenceMonitoringConfiguration' managed_persistence_monitoring_configuration: Managed log persistence configurations for a JobRun.
        :param 'ApplicationS3MonitoringConfiguration' s3_monitoring_configuration: S3 monitoring configurations for a JobRun.
        """
        if cloud_watch_logging_configuration is not None:
            pulumi.set(__self__, "cloud_watch_logging_configuration", cloud_watch_logging_configuration)
        if managed_persistence_monitoring_configuration is not None:
            pulumi.set(__self__, "managed_persistence_monitoring_configuration", managed_persistence_monitoring_configuration)
        if s3_monitoring_configuration is not None:
            pulumi.set(__self__, "s3_monitoring_configuration", s3_monitoring_configuration)

    @property
    @pulumi.getter(name="cloudWatchLoggingConfiguration")
    def cloud_watch_logging_configuration(self) -> Optional['outputs.ApplicationCloudWatchLoggingConfiguration']:
        """
        CloudWatch logging configurations for a JobRun.
        """
        return pulumi.get(self, "cloud_watch_logging_configuration")

    @property
    @pulumi.getter(name="managedPersistenceMonitoringConfiguration")
    def managed_persistence_monitoring_configuration(self) -> Optional['outputs.ApplicationManagedPersistenceMonitoringConfiguration']:
        """
        Managed log persistence configurations for a JobRun.
        """
        return pulumi.get(self, "managed_persistence_monitoring_configuration")

    @property
    @pulumi.getter(name="s3MonitoringConfiguration")
    def s3_monitoring_configuration(self) -> Optional['outputs.ApplicationS3MonitoringConfiguration']:
        """
        S3 monitoring configurations for a JobRun.
        """
        return pulumi.get(self, "s3_monitoring_configuration")


@pulumi.output_type
class ApplicationNetworkConfiguration(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "securityGroupIds":
            suggest = "security_group_ids"
        elif key == "subnetIds":
            suggest = "subnet_ids"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApplicationNetworkConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApplicationNetworkConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApplicationNetworkConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 security_group_ids: Optional[Sequence[str]] = None,
                 subnet_ids: Optional[Sequence[str]] = None):
        """
        :param Sequence[str] security_group_ids: The ID of the security groups in the VPC to which you want to connect your job or application.
        :param Sequence[str] subnet_ids: The ID of the subnets in the VPC to which you want to connect your job or application.
        """
        if security_group_ids is not None:
            pulumi.set(__self__, "security_group_ids", security_group_ids)
        if subnet_ids is not None:
            pulumi.set(__self__, "subnet_ids", subnet_ids)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[Sequence[str]]:
        """
        The ID of the security groups in the VPC to which you want to connect your job or application.
        """
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> Optional[Sequence[str]]:
        """
        The ID of the subnets in the VPC to which you want to connect your job or application.
        """
        return pulumi.get(self, "subnet_ids")


@pulumi.output_type
class ApplicationS3MonitoringConfiguration(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "encryptionKeyArn":
            suggest = "encryption_key_arn"
        elif key == "logUri":
            suggest = "log_uri"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApplicationS3MonitoringConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApplicationS3MonitoringConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApplicationS3MonitoringConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 encryption_key_arn: Optional[str] = None,
                 log_uri: Optional[str] = None):
        """
        :param str encryption_key_arn: KMS key ARN to encrypt the logs stored in given s3
        """
        if encryption_key_arn is not None:
            pulumi.set(__self__, "encryption_key_arn", encryption_key_arn)
        if log_uri is not None:
            pulumi.set(__self__, "log_uri", log_uri)

    @property
    @pulumi.getter(name="encryptionKeyArn")
    def encryption_key_arn(self) -> Optional[str]:
        """
        KMS key ARN to encrypt the logs stored in given s3
        """
        return pulumi.get(self, "encryption_key_arn")

    @property
    @pulumi.getter(name="logUri")
    def log_uri(self) -> Optional[str]:
        return pulumi.get(self, "log_uri")


@pulumi.output_type
class ApplicationWorkerConfiguration(dict):
    def __init__(__self__, *,
                 cpu: str,
                 memory: str,
                 disk: Optional[str] = None):
        """
        :param str cpu: Per worker CPU resource. vCPU is the only supported unit and specifying vCPU is optional.
        :param str memory: Per worker memory resource. GB is the only supported unit and specifying GB is optional.
        :param str disk: Per worker Disk resource. GB is the only supported unit and specifying GB is optional
        """
        pulumi.set(__self__, "cpu", cpu)
        pulumi.set(__self__, "memory", memory)
        if disk is not None:
            pulumi.set(__self__, "disk", disk)

    @property
    @pulumi.getter
    def cpu(self) -> str:
        """
        Per worker CPU resource. vCPU is the only supported unit and specifying vCPU is optional.
        """
        return pulumi.get(self, "cpu")

    @property
    @pulumi.getter
    def memory(self) -> str:
        """
        Per worker memory resource. GB is the only supported unit and specifying GB is optional.
        """
        return pulumi.get(self, "memory")

    @property
    @pulumi.getter
    def disk(self) -> Optional[str]:
        """
        Per worker Disk resource. GB is the only supported unit and specifying GB is optional
        """
        return pulumi.get(self, "disk")


@pulumi.output_type
class ApplicationWorkerTypeSpecificationInput(dict):
    """
    The specifications for a worker type.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "imageConfiguration":
            suggest = "image_configuration"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApplicationWorkerTypeSpecificationInput. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApplicationWorkerTypeSpecificationInput.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApplicationWorkerTypeSpecificationInput.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 image_configuration: Optional['outputs.ApplicationImageConfigurationInput'] = None):
        """
        The specifications for a worker type.
        """
        if image_configuration is not None:
            pulumi.set(__self__, "image_configuration", image_configuration)

    @property
    @pulumi.getter(name="imageConfiguration")
    def image_configuration(self) -> Optional['outputs.ApplicationImageConfigurationInput']:
        return pulumi.get(self, "image_configuration")


