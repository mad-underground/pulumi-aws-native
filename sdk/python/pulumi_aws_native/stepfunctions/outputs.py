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
    'StateMachineAttributes',
    'StateMachineCloudWatchLogsLogGroup',
    'StateMachineDefinitionSubstitutions',
    'StateMachineLogDestination',
    'StateMachineLoggingConfiguration',
    'StateMachineProperties',
    'StateMachineS3Location',
    'StateMachineTagsEntry',
    'StateMachineTracingConfiguration',
]

@pulumi.output_type
class StateMachineAttributes(dict):
    def __init__(__self__, *,
                 arn: str,
                 name: str):
        pulumi.set(__self__, "arn", arn)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="Arn")
    def arn(self) -> str:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="Name")
    def name(self) -> str:
        return pulumi.get(self, "name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class StateMachineCloudWatchLogsLogGroup(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-cloudwatchlogsloggroup.html
    """
    def __init__(__self__, *,
                 log_group_arn: Optional[str] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-cloudwatchlogsloggroup.html
        :param str log_group_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-cloudwatchlogsloggroup.html#cfn-stepfunctions-statemachine-cloudwatchlogsloggroup-loggrouparn
        """
        if log_group_arn is not None:
            pulumi.set(__self__, "log_group_arn", log_group_arn)

    @property
    @pulumi.getter(name="LogGroupArn")
    def log_group_arn(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-cloudwatchlogsloggroup.html#cfn-stepfunctions-statemachine-cloudwatchlogsloggroup-loggrouparn
        """
        return pulumi.get(self, "log_group_arn")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class StateMachineDefinitionSubstitutions(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-definitionsubstitutions.html
    """
    def __init__(__self__):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-definitionsubstitutions.html
        """
        pass

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class StateMachineLogDestination(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-logdestination.html
    """
    def __init__(__self__, *,
                 cloud_watch_logs_log_group: Optional['outputs.StateMachineCloudWatchLogsLogGroup'] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-logdestination.html
        :param 'StateMachineCloudWatchLogsLogGroupArgs' cloud_watch_logs_log_group: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-logdestination.html#cfn-stepfunctions-statemachine-logdestination-cloudwatchlogsloggroup
        """
        if cloud_watch_logs_log_group is not None:
            pulumi.set(__self__, "cloud_watch_logs_log_group", cloud_watch_logs_log_group)

    @property
    @pulumi.getter(name="CloudWatchLogsLogGroup")
    def cloud_watch_logs_log_group(self) -> Optional['outputs.StateMachineCloudWatchLogsLogGroup']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-logdestination.html#cfn-stepfunctions-statemachine-logdestination-cloudwatchlogsloggroup
        """
        return pulumi.get(self, "cloud_watch_logs_log_group")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class StateMachineLoggingConfiguration(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-loggingconfiguration.html
    """
    def __init__(__self__, *,
                 destinations: Optional[Sequence['outputs.StateMachineLogDestination']] = None,
                 include_execution_data: Optional[bool] = None,
                 level: Optional[str] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-loggingconfiguration.html
        :param Sequence['StateMachineLogDestinationArgs'] destinations: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-loggingconfiguration.html#cfn-stepfunctions-statemachine-loggingconfiguration-destinations
        :param bool include_execution_data: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-loggingconfiguration.html#cfn-stepfunctions-statemachine-loggingconfiguration-includeexecutiondata
        :param str level: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-loggingconfiguration.html#cfn-stepfunctions-statemachine-loggingconfiguration-level
        """
        if destinations is not None:
            pulumi.set(__self__, "destinations", destinations)
        if include_execution_data is not None:
            pulumi.set(__self__, "include_execution_data", include_execution_data)
        if level is not None:
            pulumi.set(__self__, "level", level)

    @property
    @pulumi.getter(name="Destinations")
    def destinations(self) -> Optional[Sequence['outputs.StateMachineLogDestination']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-loggingconfiguration.html#cfn-stepfunctions-statemachine-loggingconfiguration-destinations
        """
        return pulumi.get(self, "destinations")

    @property
    @pulumi.getter(name="IncludeExecutionData")
    def include_execution_data(self) -> Optional[bool]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-loggingconfiguration.html#cfn-stepfunctions-statemachine-loggingconfiguration-includeexecutiondata
        """
        return pulumi.get(self, "include_execution_data")

    @property
    @pulumi.getter(name="Level")
    def level(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-loggingconfiguration.html#cfn-stepfunctions-statemachine-loggingconfiguration-level
        """
        return pulumi.get(self, "level")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class StateMachineProperties(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html
    """
    def __init__(__self__, *,
                 role_arn: str,
                 definition_s3_location: Optional['outputs.StateMachineS3Location'] = None,
                 definition_string: Optional[str] = None,
                 definition_substitutions: Optional['outputs.StateMachineDefinitionSubstitutions'] = None,
                 logging_configuration: Optional['outputs.StateMachineLoggingConfiguration'] = None,
                 state_machine_name: Optional[str] = None,
                 state_machine_type: Optional[str] = None,
                 tags: Optional[Sequence['outputs.StateMachineTagsEntry']] = None,
                 tracing_configuration: Optional['outputs.StateMachineTracingConfiguration'] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html
        :param str role_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-rolearn
        :param 'StateMachineS3LocationArgs' definition_s3_location: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definitions3location
        :param str definition_string: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definitionstring
        :param 'StateMachineDefinitionSubstitutionsArgs' definition_substitutions: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definitionsubstitutions
        :param 'StateMachineLoggingConfigurationArgs' logging_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-loggingconfiguration
        :param str state_machine_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-statemachinename
        :param str state_machine_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-statemachinetype
        :param Sequence['StateMachineTagsEntryArgs'] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-tags
        :param 'StateMachineTracingConfigurationArgs' tracing_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-tracingconfiguration
        """
        pulumi.set(__self__, "role_arn", role_arn)
        if definition_s3_location is not None:
            pulumi.set(__self__, "definition_s3_location", definition_s3_location)
        if definition_string is not None:
            pulumi.set(__self__, "definition_string", definition_string)
        if definition_substitutions is not None:
            pulumi.set(__self__, "definition_substitutions", definition_substitutions)
        if logging_configuration is not None:
            pulumi.set(__self__, "logging_configuration", logging_configuration)
        if state_machine_name is not None:
            pulumi.set(__self__, "state_machine_name", state_machine_name)
        if state_machine_type is not None:
            pulumi.set(__self__, "state_machine_type", state_machine_type)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tracing_configuration is not None:
            pulumi.set(__self__, "tracing_configuration", tracing_configuration)

    @property
    @pulumi.getter(name="RoleArn")
    def role_arn(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-rolearn
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="DefinitionS3Location")
    def definition_s3_location(self) -> Optional['outputs.StateMachineS3Location']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definitions3location
        """
        return pulumi.get(self, "definition_s3_location")

    @property
    @pulumi.getter(name="DefinitionString")
    def definition_string(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definitionstring
        """
        return pulumi.get(self, "definition_string")

    @property
    @pulumi.getter(name="DefinitionSubstitutions")
    def definition_substitutions(self) -> Optional['outputs.StateMachineDefinitionSubstitutions']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definitionsubstitutions
        """
        return pulumi.get(self, "definition_substitutions")

    @property
    @pulumi.getter(name="LoggingConfiguration")
    def logging_configuration(self) -> Optional['outputs.StateMachineLoggingConfiguration']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-loggingconfiguration
        """
        return pulumi.get(self, "logging_configuration")

    @property
    @pulumi.getter(name="StateMachineName")
    def state_machine_name(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-statemachinename
        """
        return pulumi.get(self, "state_machine_name")

    @property
    @pulumi.getter(name="StateMachineType")
    def state_machine_type(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-statemachinetype
        """
        return pulumi.get(self, "state_machine_type")

    @property
    @pulumi.getter(name="Tags")
    def tags(self) -> Optional[Sequence['outputs.StateMachineTagsEntry']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="TracingConfiguration")
    def tracing_configuration(self) -> Optional['outputs.StateMachineTracingConfiguration']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-tracingconfiguration
        """
        return pulumi.get(self, "tracing_configuration")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class StateMachineS3Location(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-s3location.html
    """
    def __init__(__self__, *,
                 bucket: str,
                 key: str,
                 version: Optional[str] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-s3location.html
        :param str bucket: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-s3location.html#cfn-stepfunctions-statemachine-s3location-bucket
        :param str key: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-s3location.html#cfn-stepfunctions-statemachine-s3location-key
        :param str version: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-s3location.html#cfn-stepfunctions-statemachine-s3location-version
        """
        pulumi.set(__self__, "bucket", bucket)
        pulumi.set(__self__, "key", key)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="Bucket")
    def bucket(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-s3location.html#cfn-stepfunctions-statemachine-s3location-bucket
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter(name="Key")
    def key(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-s3location.html#cfn-stepfunctions-statemachine-s3location-key
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="Version")
    def version(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-s3location.html#cfn-stepfunctions-statemachine-s3location-version
        """
        return pulumi.get(self, "version")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class StateMachineTagsEntry(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-tagsentry.html
    """
    def __init__(__self__, *,
                 key: str,
                 value: str):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-tagsentry.html
        :param str key: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-tagsentry.html#cfn-stepfunctions-statemachine-tagsentry-key
        :param str value: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-tagsentry.html#cfn-stepfunctions-statemachine-tagsentry-value
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="Key")
    def key(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-tagsentry.html#cfn-stepfunctions-statemachine-tagsentry-key
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="Value")
    def value(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-tagsentry.html#cfn-stepfunctions-statemachine-tagsentry-value
        """
        return pulumi.get(self, "value")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class StateMachineTracingConfiguration(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-tracingconfiguration.html
    """
    def __init__(__self__, *,
                 enabled: Optional[bool] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-tracingconfiguration.html
        :param bool enabled: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-tracingconfiguration.html#cfn-stepfunctions-statemachine-tracingconfiguration-enabled
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter(name="Enabled")
    def enabled(self) -> Optional[bool]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-tracingconfiguration.html#cfn-stepfunctions-statemachine-tracingconfiguration-enabled
        """
        return pulumi.get(self, "enabled")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


