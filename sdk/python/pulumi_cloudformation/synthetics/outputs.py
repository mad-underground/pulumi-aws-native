# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from .. import outputs as _root_outputs

__all__ = [
    'CanaryAttributes',
    'CanaryCode',
    'CanaryProperties',
    'CanaryRunConfig',
    'CanarySchedule',
    'CanaryVPCConfig',
]

@pulumi.output_type
class CanaryAttributes(dict):
    def __init__(__self__, *,
                 id: str,
                 state: str):
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="Id")
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="State")
    def state(self) -> str:
        return pulumi.get(self, "state")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CanaryCode(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-code.html
    """
    def __init__(__self__, *,
                 handler: Optional[str] = None,
                 s3_bucket: Optional[str] = None,
                 s3_key: Optional[str] = None,
                 s3_object_version: Optional[str] = None,
                 script: Optional[str] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-code.html
        :param str handler: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-code.html#cfn-synthetics-canary-code-handler
        :param str s3_bucket: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-code.html#cfn-synthetics-canary-code-s3bucket
        :param str s3_key: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-code.html#cfn-synthetics-canary-code-s3key
        :param str s3_object_version: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-code.html#cfn-synthetics-canary-code-s3objectversion
        :param str script: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-code.html#cfn-synthetics-canary-code-script
        """
        if handler is not None:
            pulumi.set(__self__, "handler", handler)
        if s3_bucket is not None:
            pulumi.set(__self__, "s3_bucket", s3_bucket)
        if s3_key is not None:
            pulumi.set(__self__, "s3_key", s3_key)
        if s3_object_version is not None:
            pulumi.set(__self__, "s3_object_version", s3_object_version)
        if script is not None:
            pulumi.set(__self__, "script", script)

    @property
    @pulumi.getter(name="Handler")
    def handler(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-code.html#cfn-synthetics-canary-code-handler
        """
        return pulumi.get(self, "handler")

    @property
    @pulumi.getter(name="S3Bucket")
    def s3_bucket(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-code.html#cfn-synthetics-canary-code-s3bucket
        """
        return pulumi.get(self, "s3_bucket")

    @property
    @pulumi.getter(name="S3Key")
    def s3_key(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-code.html#cfn-synthetics-canary-code-s3key
        """
        return pulumi.get(self, "s3_key")

    @property
    @pulumi.getter(name="S3ObjectVersion")
    def s3_object_version(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-code.html#cfn-synthetics-canary-code-s3objectversion
        """
        return pulumi.get(self, "s3_object_version")

    @property
    @pulumi.getter(name="Script")
    def script(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-code.html#cfn-synthetics-canary-code-script
        """
        return pulumi.get(self, "script")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CanaryProperties(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html
    """
    def __init__(__self__, *,
                 artifact_s3_location: str,
                 code: 'outputs.CanaryCode',
                 execution_role_arn: str,
                 name: str,
                 runtime_version: str,
                 schedule: 'outputs.CanarySchedule',
                 start_canary_after_creation: bool,
                 failure_retention_period: Optional[int] = None,
                 run_config: Optional['outputs.CanaryRunConfig'] = None,
                 success_retention_period: Optional[int] = None,
                 tags: Optional[Sequence['_root_outputs.Tag']] = None,
                 vpc_config: Optional['outputs.CanaryVPCConfig'] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html
        :param str artifact_s3_location: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-artifacts3location
        :param 'CanaryCodeArgs' code: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-code
        :param str execution_role_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-executionrolearn
        :param str name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-name
        :param str runtime_version: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-runtimeversion
        :param 'CanaryScheduleArgs' schedule: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-schedule
        :param bool start_canary_after_creation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-startcanaryaftercreation
        :param int failure_retention_period: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-failureretentionperiod
        :param 'CanaryRunConfigArgs' run_config: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-runconfig
        :param int success_retention_period: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-successretentionperiod
        :param Sequence['_root_inputs.TagArgs'] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-tags
        :param 'CanaryVPCConfigArgs' vpc_config: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-vpcconfig
        """
        pulumi.set(__self__, "artifact_s3_location", artifact_s3_location)
        pulumi.set(__self__, "code", code)
        pulumi.set(__self__, "execution_role_arn", execution_role_arn)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "runtime_version", runtime_version)
        pulumi.set(__self__, "schedule", schedule)
        pulumi.set(__self__, "start_canary_after_creation", start_canary_after_creation)
        if failure_retention_period is not None:
            pulumi.set(__self__, "failure_retention_period", failure_retention_period)
        if run_config is not None:
            pulumi.set(__self__, "run_config", run_config)
        if success_retention_period is not None:
            pulumi.set(__self__, "success_retention_period", success_retention_period)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vpc_config is not None:
            pulumi.set(__self__, "vpc_config", vpc_config)

    @property
    @pulumi.getter(name="ArtifactS3Location")
    def artifact_s3_location(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-artifacts3location
        """
        return pulumi.get(self, "artifact_s3_location")

    @property
    @pulumi.getter(name="Code")
    def code(self) -> 'outputs.CanaryCode':
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-code
        """
        return pulumi.get(self, "code")

    @property
    @pulumi.getter(name="ExecutionRoleArn")
    def execution_role_arn(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-executionrolearn
        """
        return pulumi.get(self, "execution_role_arn")

    @property
    @pulumi.getter(name="Name")
    def name(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="RuntimeVersion")
    def runtime_version(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-runtimeversion
        """
        return pulumi.get(self, "runtime_version")

    @property
    @pulumi.getter(name="Schedule")
    def schedule(self) -> 'outputs.CanarySchedule':
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-schedule
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter(name="StartCanaryAfterCreation")
    def start_canary_after_creation(self) -> bool:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-startcanaryaftercreation
        """
        return pulumi.get(self, "start_canary_after_creation")

    @property
    @pulumi.getter(name="FailureRetentionPeriod")
    def failure_retention_period(self) -> Optional[int]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-failureretentionperiod
        """
        return pulumi.get(self, "failure_retention_period")

    @property
    @pulumi.getter(name="RunConfig")
    def run_config(self) -> Optional['outputs.CanaryRunConfig']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-runconfig
        """
        return pulumi.get(self, "run_config")

    @property
    @pulumi.getter(name="SuccessRetentionPeriod")
    def success_retention_period(self) -> Optional[int]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-successretentionperiod
        """
        return pulumi.get(self, "success_retention_period")

    @property
    @pulumi.getter(name="Tags")
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="VPCConfig")
    def vpc_config(self) -> Optional['outputs.CanaryVPCConfig']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-vpcconfig
        """
        return pulumi.get(self, "vpc_config")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CanaryRunConfig(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-runconfig.html
    """
    def __init__(__self__, *,
                 timeout_in_seconds: int,
                 active_tracing: Optional[bool] = None,
                 environment_variables: Optional[Mapping[str, str]] = None,
                 memory_in_mb: Optional[int] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-runconfig.html
        :param int timeout_in_seconds: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-runconfig.html#cfn-synthetics-canary-runconfig-timeoutinseconds
        :param bool active_tracing: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-runconfig.html#cfn-synthetics-canary-runconfig-activetracing
        :param Mapping[str, str] environment_variables: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-runconfig.html#cfn-synthetics-canary-runconfig-environmentvariables
        :param int memory_in_mb: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-runconfig.html#cfn-synthetics-canary-runconfig-memoryinmb
        """
        pulumi.set(__self__, "timeout_in_seconds", timeout_in_seconds)
        if active_tracing is not None:
            pulumi.set(__self__, "active_tracing", active_tracing)
        if environment_variables is not None:
            pulumi.set(__self__, "environment_variables", environment_variables)
        if memory_in_mb is not None:
            pulumi.set(__self__, "memory_in_mb", memory_in_mb)

    @property
    @pulumi.getter(name="TimeoutInSeconds")
    def timeout_in_seconds(self) -> int:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-runconfig.html#cfn-synthetics-canary-runconfig-timeoutinseconds
        """
        return pulumi.get(self, "timeout_in_seconds")

    @property
    @pulumi.getter(name="ActiveTracing")
    def active_tracing(self) -> Optional[bool]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-runconfig.html#cfn-synthetics-canary-runconfig-activetracing
        """
        return pulumi.get(self, "active_tracing")

    @property
    @pulumi.getter(name="EnvironmentVariables")
    def environment_variables(self) -> Optional[Mapping[str, str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-runconfig.html#cfn-synthetics-canary-runconfig-environmentvariables
        """
        return pulumi.get(self, "environment_variables")

    @property
    @pulumi.getter(name="MemoryInMB")
    def memory_in_mb(self) -> Optional[int]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-runconfig.html#cfn-synthetics-canary-runconfig-memoryinmb
        """
        return pulumi.get(self, "memory_in_mb")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CanarySchedule(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-schedule.html
    """
    def __init__(__self__, *,
                 expression: str,
                 duration_in_seconds: Optional[str] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-schedule.html
        :param str expression: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-schedule.html#cfn-synthetics-canary-schedule-expression
        :param str duration_in_seconds: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-schedule.html#cfn-synthetics-canary-schedule-durationinseconds
        """
        pulumi.set(__self__, "expression", expression)
        if duration_in_seconds is not None:
            pulumi.set(__self__, "duration_in_seconds", duration_in_seconds)

    @property
    @pulumi.getter(name="Expression")
    def expression(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-schedule.html#cfn-synthetics-canary-schedule-expression
        """
        return pulumi.get(self, "expression")

    @property
    @pulumi.getter(name="DurationInSeconds")
    def duration_in_seconds(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-schedule.html#cfn-synthetics-canary-schedule-durationinseconds
        """
        return pulumi.get(self, "duration_in_seconds")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CanaryVPCConfig(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-vpcconfig.html
    """
    def __init__(__self__, *,
                 security_group_ids: Sequence[str],
                 subnet_ids: Sequence[str],
                 vpc_id: Optional[str] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-vpcconfig.html
        :param Sequence[str] security_group_ids: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-vpcconfig.html#cfn-synthetics-canary-vpcconfig-securitygroupids
        :param Sequence[str] subnet_ids: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-vpcconfig.html#cfn-synthetics-canary-vpcconfig-subnetids
        :param str vpc_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-vpcconfig.html#cfn-synthetics-canary-vpcconfig-vpcid
        """
        pulumi.set(__self__, "security_group_ids", security_group_ids)
        pulumi.set(__self__, "subnet_ids", subnet_ids)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="SecurityGroupIds")
    def security_group_ids(self) -> Sequence[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-vpcconfig.html#cfn-synthetics-canary-vpcconfig-securitygroupids
        """
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter(name="SubnetIds")
    def subnet_ids(self) -> Sequence[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-vpcconfig.html#cfn-synthetics-canary-vpcconfig-subnetids
        """
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter(name="VpcId")
    def vpc_id(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-vpcconfig.html#cfn-synthetics-canary-vpcconfig-vpcid
        """
        return pulumi.get(self, "vpc_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


