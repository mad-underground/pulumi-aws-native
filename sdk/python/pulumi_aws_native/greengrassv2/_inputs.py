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
    'ComponentVersionComponentDependencyRequirementArgs',
    'ComponentVersionComponentPlatformArgs',
    'ComponentVersionLambdaContainerParamsArgs',
    'ComponentVersionLambdaDeviceMountArgs',
    'ComponentVersionLambdaEventSourceArgs',
    'ComponentVersionLambdaExecutionParametersArgs',
    'ComponentVersionLambdaFunctionRecipeSourceArgs',
    'ComponentVersionLambdaLinuxProcessParamsArgs',
    'ComponentVersionLambdaVolumeMountArgs',
    'DeploymentComponentConfigurationUpdateArgs',
    'DeploymentComponentDeploymentSpecificationArgs',
    'DeploymentComponentRunWithArgs',
    'DeploymentComponentUpdatePolicyArgs',
    'DeploymentConfigurationValidationPolicyArgs',
    'DeploymentIoTJobAbortConfigArgs',
    'DeploymentIoTJobAbortCriteriaArgs',
    'DeploymentIoTJobConfigurationArgs',
    'DeploymentIoTJobExecutionsRolloutConfigArgs',
    'DeploymentIoTJobExponentialRolloutRateArgs',
    'DeploymentIoTJobRateIncreaseCriteriaArgs',
    'DeploymentIoTJobTimeoutConfigArgs',
    'DeploymentPoliciesArgs',
    'DeploymentSystemResourceLimitsArgs',
]

@pulumi.input_type
class ComponentVersionComponentDependencyRequirementArgs:
    def __init__(__self__, *,
                 dependency_type: Optional[pulumi.Input['ComponentVersionComponentDependencyRequirementDependencyType']] = None,
                 version_requirement: Optional[pulumi.Input[str]] = None):
        if dependency_type is not None:
            pulumi.set(__self__, "dependency_type", dependency_type)
        if version_requirement is not None:
            pulumi.set(__self__, "version_requirement", version_requirement)

    @property
    @pulumi.getter(name="dependencyType")
    def dependency_type(self) -> Optional[pulumi.Input['ComponentVersionComponentDependencyRequirementDependencyType']]:
        return pulumi.get(self, "dependency_type")

    @dependency_type.setter
    def dependency_type(self, value: Optional[pulumi.Input['ComponentVersionComponentDependencyRequirementDependencyType']]):
        pulumi.set(self, "dependency_type", value)

    @property
    @pulumi.getter(name="versionRequirement")
    def version_requirement(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "version_requirement")

    @version_requirement.setter
    def version_requirement(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version_requirement", value)


@pulumi.input_type
class ComponentVersionComponentPlatformArgs:
    def __init__(__self__, *,
                 attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        if attributes is not None:
            pulumi.set(__self__, "attributes", attributes)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def attributes(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "attributes")

    @attributes.setter
    def attributes(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "attributes", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class ComponentVersionLambdaContainerParamsArgs:
    def __init__(__self__, *,
                 devices: Optional[pulumi.Input[Sequence[pulumi.Input['ComponentVersionLambdaDeviceMountArgs']]]] = None,
                 memory_size_in_kb: Optional[pulumi.Input[int]] = None,
                 mount_ro_sysfs: Optional[pulumi.Input[bool]] = None,
                 volumes: Optional[pulumi.Input[Sequence[pulumi.Input['ComponentVersionLambdaVolumeMountArgs']]]] = None):
        if devices is not None:
            pulumi.set(__self__, "devices", devices)
        if memory_size_in_kb is not None:
            pulumi.set(__self__, "memory_size_in_kb", memory_size_in_kb)
        if mount_ro_sysfs is not None:
            pulumi.set(__self__, "mount_ro_sysfs", mount_ro_sysfs)
        if volumes is not None:
            pulumi.set(__self__, "volumes", volumes)

    @property
    @pulumi.getter
    def devices(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ComponentVersionLambdaDeviceMountArgs']]]]:
        return pulumi.get(self, "devices")

    @devices.setter
    def devices(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ComponentVersionLambdaDeviceMountArgs']]]]):
        pulumi.set(self, "devices", value)

    @property
    @pulumi.getter(name="memorySizeInKb")
    def memory_size_in_kb(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "memory_size_in_kb")

    @memory_size_in_kb.setter
    def memory_size_in_kb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "memory_size_in_kb", value)

    @property
    @pulumi.getter(name="mountRoSysfs")
    def mount_ro_sysfs(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "mount_ro_sysfs")

    @mount_ro_sysfs.setter
    def mount_ro_sysfs(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "mount_ro_sysfs", value)

    @property
    @pulumi.getter
    def volumes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ComponentVersionLambdaVolumeMountArgs']]]]:
        return pulumi.get(self, "volumes")

    @volumes.setter
    def volumes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ComponentVersionLambdaVolumeMountArgs']]]]):
        pulumi.set(self, "volumes", value)


@pulumi.input_type
class ComponentVersionLambdaDeviceMountArgs:
    def __init__(__self__, *,
                 add_group_owner: Optional[pulumi.Input[bool]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 permission: Optional[pulumi.Input['ComponentVersionLambdaFilesystemPermission']] = None):
        if add_group_owner is not None:
            pulumi.set(__self__, "add_group_owner", add_group_owner)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if permission is not None:
            pulumi.set(__self__, "permission", permission)

    @property
    @pulumi.getter(name="addGroupOwner")
    def add_group_owner(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "add_group_owner")

    @add_group_owner.setter
    def add_group_owner(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "add_group_owner", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter
    def permission(self) -> Optional[pulumi.Input['ComponentVersionLambdaFilesystemPermission']]:
        return pulumi.get(self, "permission")

    @permission.setter
    def permission(self, value: Optional[pulumi.Input['ComponentVersionLambdaFilesystemPermission']]):
        pulumi.set(self, "permission", value)


@pulumi.input_type
class ComponentVersionLambdaEventSourceArgs:
    def __init__(__self__, *,
                 topic: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input['ComponentVersionLambdaEventSourceType']] = None):
        if topic is not None:
            pulumi.set(__self__, "topic", topic)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def topic(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "topic")

    @topic.setter
    def topic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "topic", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input['ComponentVersionLambdaEventSourceType']]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input['ComponentVersionLambdaEventSourceType']]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class ComponentVersionLambdaExecutionParametersArgs:
    def __init__(__self__, *,
                 environment_variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 event_sources: Optional[pulumi.Input[Sequence[pulumi.Input['ComponentVersionLambdaEventSourceArgs']]]] = None,
                 exec_args: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 input_payload_encoding_type: Optional[pulumi.Input['ComponentVersionLambdaExecutionParametersInputPayloadEncodingType']] = None,
                 linux_process_params: Optional[pulumi.Input['ComponentVersionLambdaLinuxProcessParamsArgs']] = None,
                 max_idle_time_in_seconds: Optional[pulumi.Input[int]] = None,
                 max_instances_count: Optional[pulumi.Input[int]] = None,
                 max_queue_size: Optional[pulumi.Input[int]] = None,
                 pinned: Optional[pulumi.Input[bool]] = None,
                 status_timeout_in_seconds: Optional[pulumi.Input[int]] = None,
                 timeout_in_seconds: Optional[pulumi.Input[int]] = None):
        if environment_variables is not None:
            pulumi.set(__self__, "environment_variables", environment_variables)
        if event_sources is not None:
            pulumi.set(__self__, "event_sources", event_sources)
        if exec_args is not None:
            pulumi.set(__self__, "exec_args", exec_args)
        if input_payload_encoding_type is not None:
            pulumi.set(__self__, "input_payload_encoding_type", input_payload_encoding_type)
        if linux_process_params is not None:
            pulumi.set(__self__, "linux_process_params", linux_process_params)
        if max_idle_time_in_seconds is not None:
            pulumi.set(__self__, "max_idle_time_in_seconds", max_idle_time_in_seconds)
        if max_instances_count is not None:
            pulumi.set(__self__, "max_instances_count", max_instances_count)
        if max_queue_size is not None:
            pulumi.set(__self__, "max_queue_size", max_queue_size)
        if pinned is not None:
            pulumi.set(__self__, "pinned", pinned)
        if status_timeout_in_seconds is not None:
            pulumi.set(__self__, "status_timeout_in_seconds", status_timeout_in_seconds)
        if timeout_in_seconds is not None:
            pulumi.set(__self__, "timeout_in_seconds", timeout_in_seconds)

    @property
    @pulumi.getter(name="environmentVariables")
    def environment_variables(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "environment_variables")

    @environment_variables.setter
    def environment_variables(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "environment_variables", value)

    @property
    @pulumi.getter(name="eventSources")
    def event_sources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ComponentVersionLambdaEventSourceArgs']]]]:
        return pulumi.get(self, "event_sources")

    @event_sources.setter
    def event_sources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ComponentVersionLambdaEventSourceArgs']]]]):
        pulumi.set(self, "event_sources", value)

    @property
    @pulumi.getter(name="execArgs")
    def exec_args(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "exec_args")

    @exec_args.setter
    def exec_args(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "exec_args", value)

    @property
    @pulumi.getter(name="inputPayloadEncodingType")
    def input_payload_encoding_type(self) -> Optional[pulumi.Input['ComponentVersionLambdaExecutionParametersInputPayloadEncodingType']]:
        return pulumi.get(self, "input_payload_encoding_type")

    @input_payload_encoding_type.setter
    def input_payload_encoding_type(self, value: Optional[pulumi.Input['ComponentVersionLambdaExecutionParametersInputPayloadEncodingType']]):
        pulumi.set(self, "input_payload_encoding_type", value)

    @property
    @pulumi.getter(name="linuxProcessParams")
    def linux_process_params(self) -> Optional[pulumi.Input['ComponentVersionLambdaLinuxProcessParamsArgs']]:
        return pulumi.get(self, "linux_process_params")

    @linux_process_params.setter
    def linux_process_params(self, value: Optional[pulumi.Input['ComponentVersionLambdaLinuxProcessParamsArgs']]):
        pulumi.set(self, "linux_process_params", value)

    @property
    @pulumi.getter(name="maxIdleTimeInSeconds")
    def max_idle_time_in_seconds(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "max_idle_time_in_seconds")

    @max_idle_time_in_seconds.setter
    def max_idle_time_in_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_idle_time_in_seconds", value)

    @property
    @pulumi.getter(name="maxInstancesCount")
    def max_instances_count(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "max_instances_count")

    @max_instances_count.setter
    def max_instances_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_instances_count", value)

    @property
    @pulumi.getter(name="maxQueueSize")
    def max_queue_size(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "max_queue_size")

    @max_queue_size.setter
    def max_queue_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_queue_size", value)

    @property
    @pulumi.getter
    def pinned(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "pinned")

    @pinned.setter
    def pinned(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "pinned", value)

    @property
    @pulumi.getter(name="statusTimeoutInSeconds")
    def status_timeout_in_seconds(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "status_timeout_in_seconds")

    @status_timeout_in_seconds.setter
    def status_timeout_in_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "status_timeout_in_seconds", value)

    @property
    @pulumi.getter(name="timeoutInSeconds")
    def timeout_in_seconds(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "timeout_in_seconds")

    @timeout_in_seconds.setter
    def timeout_in_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout_in_seconds", value)


@pulumi.input_type
class ComponentVersionLambdaFunctionRecipeSourceArgs:
    def __init__(__self__, *,
                 component_dependencies: Optional[pulumi.Input[Mapping[str, pulumi.Input['ComponentVersionComponentDependencyRequirementArgs']]]] = None,
                 component_lambda_parameters: Optional[pulumi.Input['ComponentVersionLambdaExecutionParametersArgs']] = None,
                 component_name: Optional[pulumi.Input[str]] = None,
                 component_platforms: Optional[pulumi.Input[Sequence[pulumi.Input['ComponentVersionComponentPlatformArgs']]]] = None,
                 component_version: Optional[pulumi.Input[str]] = None,
                 lambda_arn: Optional[pulumi.Input[str]] = None):
        if component_dependencies is not None:
            pulumi.set(__self__, "component_dependencies", component_dependencies)
        if component_lambda_parameters is not None:
            pulumi.set(__self__, "component_lambda_parameters", component_lambda_parameters)
        if component_name is not None:
            pulumi.set(__self__, "component_name", component_name)
        if component_platforms is not None:
            pulumi.set(__self__, "component_platforms", component_platforms)
        if component_version is not None:
            pulumi.set(__self__, "component_version", component_version)
        if lambda_arn is not None:
            pulumi.set(__self__, "lambda_arn", lambda_arn)

    @property
    @pulumi.getter(name="componentDependencies")
    def component_dependencies(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input['ComponentVersionComponentDependencyRequirementArgs']]]]:
        return pulumi.get(self, "component_dependencies")

    @component_dependencies.setter
    def component_dependencies(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input['ComponentVersionComponentDependencyRequirementArgs']]]]):
        pulumi.set(self, "component_dependencies", value)

    @property
    @pulumi.getter(name="componentLambdaParameters")
    def component_lambda_parameters(self) -> Optional[pulumi.Input['ComponentVersionLambdaExecutionParametersArgs']]:
        return pulumi.get(self, "component_lambda_parameters")

    @component_lambda_parameters.setter
    def component_lambda_parameters(self, value: Optional[pulumi.Input['ComponentVersionLambdaExecutionParametersArgs']]):
        pulumi.set(self, "component_lambda_parameters", value)

    @property
    @pulumi.getter(name="componentName")
    def component_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "component_name")

    @component_name.setter
    def component_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "component_name", value)

    @property
    @pulumi.getter(name="componentPlatforms")
    def component_platforms(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ComponentVersionComponentPlatformArgs']]]]:
        return pulumi.get(self, "component_platforms")

    @component_platforms.setter
    def component_platforms(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ComponentVersionComponentPlatformArgs']]]]):
        pulumi.set(self, "component_platforms", value)

    @property
    @pulumi.getter(name="componentVersion")
    def component_version(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "component_version")

    @component_version.setter
    def component_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "component_version", value)

    @property
    @pulumi.getter(name="lambdaArn")
    def lambda_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "lambda_arn")

    @lambda_arn.setter
    def lambda_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lambda_arn", value)


@pulumi.input_type
class ComponentVersionLambdaLinuxProcessParamsArgs:
    def __init__(__self__, *,
                 container_params: Optional[pulumi.Input['ComponentVersionLambdaContainerParamsArgs']] = None,
                 isolation_mode: Optional[pulumi.Input['ComponentVersionLambdaLinuxProcessParamsIsolationMode']] = None):
        if container_params is not None:
            pulumi.set(__self__, "container_params", container_params)
        if isolation_mode is not None:
            pulumi.set(__self__, "isolation_mode", isolation_mode)

    @property
    @pulumi.getter(name="containerParams")
    def container_params(self) -> Optional[pulumi.Input['ComponentVersionLambdaContainerParamsArgs']]:
        return pulumi.get(self, "container_params")

    @container_params.setter
    def container_params(self, value: Optional[pulumi.Input['ComponentVersionLambdaContainerParamsArgs']]):
        pulumi.set(self, "container_params", value)

    @property
    @pulumi.getter(name="isolationMode")
    def isolation_mode(self) -> Optional[pulumi.Input['ComponentVersionLambdaLinuxProcessParamsIsolationMode']]:
        return pulumi.get(self, "isolation_mode")

    @isolation_mode.setter
    def isolation_mode(self, value: Optional[pulumi.Input['ComponentVersionLambdaLinuxProcessParamsIsolationMode']]):
        pulumi.set(self, "isolation_mode", value)


@pulumi.input_type
class ComponentVersionLambdaVolumeMountArgs:
    def __init__(__self__, *,
                 add_group_owner: Optional[pulumi.Input[bool]] = None,
                 destination_path: Optional[pulumi.Input[str]] = None,
                 permission: Optional[pulumi.Input['ComponentVersionLambdaFilesystemPermission']] = None,
                 source_path: Optional[pulumi.Input[str]] = None):
        if add_group_owner is not None:
            pulumi.set(__self__, "add_group_owner", add_group_owner)
        if destination_path is not None:
            pulumi.set(__self__, "destination_path", destination_path)
        if permission is not None:
            pulumi.set(__self__, "permission", permission)
        if source_path is not None:
            pulumi.set(__self__, "source_path", source_path)

    @property
    @pulumi.getter(name="addGroupOwner")
    def add_group_owner(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "add_group_owner")

    @add_group_owner.setter
    def add_group_owner(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "add_group_owner", value)

    @property
    @pulumi.getter(name="destinationPath")
    def destination_path(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "destination_path")

    @destination_path.setter
    def destination_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_path", value)

    @property
    @pulumi.getter
    def permission(self) -> Optional[pulumi.Input['ComponentVersionLambdaFilesystemPermission']]:
        return pulumi.get(self, "permission")

    @permission.setter
    def permission(self, value: Optional[pulumi.Input['ComponentVersionLambdaFilesystemPermission']]):
        pulumi.set(self, "permission", value)

    @property
    @pulumi.getter(name="sourcePath")
    def source_path(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "source_path")

    @source_path.setter
    def source_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_path", value)


@pulumi.input_type
class DeploymentComponentConfigurationUpdateArgs:
    def __init__(__self__, *,
                 merge: Optional[pulumi.Input[str]] = None,
                 reset: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        if merge is not None:
            pulumi.set(__self__, "merge", merge)
        if reset is not None:
            pulumi.set(__self__, "reset", reset)

    @property
    @pulumi.getter
    def merge(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "merge")

    @merge.setter
    def merge(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "merge", value)

    @property
    @pulumi.getter
    def reset(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "reset")

    @reset.setter
    def reset(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "reset", value)


@pulumi.input_type
class DeploymentComponentDeploymentSpecificationArgs:
    def __init__(__self__, *,
                 component_version: Optional[pulumi.Input[str]] = None,
                 configuration_update: Optional[pulumi.Input['DeploymentComponentConfigurationUpdateArgs']] = None,
                 run_with: Optional[pulumi.Input['DeploymentComponentRunWithArgs']] = None):
        if component_version is not None:
            pulumi.set(__self__, "component_version", component_version)
        if configuration_update is not None:
            pulumi.set(__self__, "configuration_update", configuration_update)
        if run_with is not None:
            pulumi.set(__self__, "run_with", run_with)

    @property
    @pulumi.getter(name="componentVersion")
    def component_version(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "component_version")

    @component_version.setter
    def component_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "component_version", value)

    @property
    @pulumi.getter(name="configurationUpdate")
    def configuration_update(self) -> Optional[pulumi.Input['DeploymentComponentConfigurationUpdateArgs']]:
        return pulumi.get(self, "configuration_update")

    @configuration_update.setter
    def configuration_update(self, value: Optional[pulumi.Input['DeploymentComponentConfigurationUpdateArgs']]):
        pulumi.set(self, "configuration_update", value)

    @property
    @pulumi.getter(name="runWith")
    def run_with(self) -> Optional[pulumi.Input['DeploymentComponentRunWithArgs']]:
        return pulumi.get(self, "run_with")

    @run_with.setter
    def run_with(self, value: Optional[pulumi.Input['DeploymentComponentRunWithArgs']]):
        pulumi.set(self, "run_with", value)


@pulumi.input_type
class DeploymentComponentRunWithArgs:
    def __init__(__self__, *,
                 posix_user: Optional[pulumi.Input[str]] = None,
                 system_resource_limits: Optional[pulumi.Input['DeploymentSystemResourceLimitsArgs']] = None,
                 windows_user: Optional[pulumi.Input[str]] = None):
        if posix_user is not None:
            pulumi.set(__self__, "posix_user", posix_user)
        if system_resource_limits is not None:
            pulumi.set(__self__, "system_resource_limits", system_resource_limits)
        if windows_user is not None:
            pulumi.set(__self__, "windows_user", windows_user)

    @property
    @pulumi.getter(name="posixUser")
    def posix_user(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "posix_user")

    @posix_user.setter
    def posix_user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "posix_user", value)

    @property
    @pulumi.getter(name="systemResourceLimits")
    def system_resource_limits(self) -> Optional[pulumi.Input['DeploymentSystemResourceLimitsArgs']]:
        return pulumi.get(self, "system_resource_limits")

    @system_resource_limits.setter
    def system_resource_limits(self, value: Optional[pulumi.Input['DeploymentSystemResourceLimitsArgs']]):
        pulumi.set(self, "system_resource_limits", value)

    @property
    @pulumi.getter(name="windowsUser")
    def windows_user(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "windows_user")

    @windows_user.setter
    def windows_user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "windows_user", value)


@pulumi.input_type
class DeploymentComponentUpdatePolicyArgs:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input['DeploymentComponentUpdatePolicyAction']] = None,
                 timeout_in_seconds: Optional[pulumi.Input[int]] = None):
        if action is not None:
            pulumi.set(__self__, "action", action)
        if timeout_in_seconds is not None:
            pulumi.set(__self__, "timeout_in_seconds", timeout_in_seconds)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input['DeploymentComponentUpdatePolicyAction']]:
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input['DeploymentComponentUpdatePolicyAction']]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="timeoutInSeconds")
    def timeout_in_seconds(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "timeout_in_seconds")

    @timeout_in_seconds.setter
    def timeout_in_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout_in_seconds", value)


@pulumi.input_type
class DeploymentConfigurationValidationPolicyArgs:
    def __init__(__self__, *,
                 timeout_in_seconds: Optional[pulumi.Input[int]] = None):
        if timeout_in_seconds is not None:
            pulumi.set(__self__, "timeout_in_seconds", timeout_in_seconds)

    @property
    @pulumi.getter(name="timeoutInSeconds")
    def timeout_in_seconds(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "timeout_in_seconds")

    @timeout_in_seconds.setter
    def timeout_in_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout_in_seconds", value)


@pulumi.input_type
class DeploymentIoTJobAbortConfigArgs:
    def __init__(__self__, *,
                 criteria_list: pulumi.Input[Sequence[pulumi.Input['DeploymentIoTJobAbortCriteriaArgs']]]):
        pulumi.set(__self__, "criteria_list", criteria_list)

    @property
    @pulumi.getter(name="criteriaList")
    def criteria_list(self) -> pulumi.Input[Sequence[pulumi.Input['DeploymentIoTJobAbortCriteriaArgs']]]:
        return pulumi.get(self, "criteria_list")

    @criteria_list.setter
    def criteria_list(self, value: pulumi.Input[Sequence[pulumi.Input['DeploymentIoTJobAbortCriteriaArgs']]]):
        pulumi.set(self, "criteria_list", value)


@pulumi.input_type
class DeploymentIoTJobAbortCriteriaArgs:
    def __init__(__self__, *,
                 action: pulumi.Input['DeploymentIoTJobAbortCriteriaAction'],
                 failure_type: pulumi.Input['DeploymentIoTJobAbortCriteriaFailureType'],
                 min_number_of_executed_things: pulumi.Input[int],
                 threshold_percentage: pulumi.Input[float]):
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "failure_type", failure_type)
        pulumi.set(__self__, "min_number_of_executed_things", min_number_of_executed_things)
        pulumi.set(__self__, "threshold_percentage", threshold_percentage)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input['DeploymentIoTJobAbortCriteriaAction']:
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input['DeploymentIoTJobAbortCriteriaAction']):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="failureType")
    def failure_type(self) -> pulumi.Input['DeploymentIoTJobAbortCriteriaFailureType']:
        return pulumi.get(self, "failure_type")

    @failure_type.setter
    def failure_type(self, value: pulumi.Input['DeploymentIoTJobAbortCriteriaFailureType']):
        pulumi.set(self, "failure_type", value)

    @property
    @pulumi.getter(name="minNumberOfExecutedThings")
    def min_number_of_executed_things(self) -> pulumi.Input[int]:
        return pulumi.get(self, "min_number_of_executed_things")

    @min_number_of_executed_things.setter
    def min_number_of_executed_things(self, value: pulumi.Input[int]):
        pulumi.set(self, "min_number_of_executed_things", value)

    @property
    @pulumi.getter(name="thresholdPercentage")
    def threshold_percentage(self) -> pulumi.Input[float]:
        return pulumi.get(self, "threshold_percentage")

    @threshold_percentage.setter
    def threshold_percentage(self, value: pulumi.Input[float]):
        pulumi.set(self, "threshold_percentage", value)


@pulumi.input_type
class DeploymentIoTJobConfigurationArgs:
    def __init__(__self__, *,
                 abort_config: Optional[pulumi.Input['DeploymentIoTJobAbortConfigArgs']] = None,
                 job_executions_rollout_config: Optional[pulumi.Input['DeploymentIoTJobExecutionsRolloutConfigArgs']] = None,
                 timeout_config: Optional[pulumi.Input['DeploymentIoTJobTimeoutConfigArgs']] = None):
        if abort_config is not None:
            pulumi.set(__self__, "abort_config", abort_config)
        if job_executions_rollout_config is not None:
            pulumi.set(__self__, "job_executions_rollout_config", job_executions_rollout_config)
        if timeout_config is not None:
            pulumi.set(__self__, "timeout_config", timeout_config)

    @property
    @pulumi.getter(name="abortConfig")
    def abort_config(self) -> Optional[pulumi.Input['DeploymentIoTJobAbortConfigArgs']]:
        return pulumi.get(self, "abort_config")

    @abort_config.setter
    def abort_config(self, value: Optional[pulumi.Input['DeploymentIoTJobAbortConfigArgs']]):
        pulumi.set(self, "abort_config", value)

    @property
    @pulumi.getter(name="jobExecutionsRolloutConfig")
    def job_executions_rollout_config(self) -> Optional[pulumi.Input['DeploymentIoTJobExecutionsRolloutConfigArgs']]:
        return pulumi.get(self, "job_executions_rollout_config")

    @job_executions_rollout_config.setter
    def job_executions_rollout_config(self, value: Optional[pulumi.Input['DeploymentIoTJobExecutionsRolloutConfigArgs']]):
        pulumi.set(self, "job_executions_rollout_config", value)

    @property
    @pulumi.getter(name="timeoutConfig")
    def timeout_config(self) -> Optional[pulumi.Input['DeploymentIoTJobTimeoutConfigArgs']]:
        return pulumi.get(self, "timeout_config")

    @timeout_config.setter
    def timeout_config(self, value: Optional[pulumi.Input['DeploymentIoTJobTimeoutConfigArgs']]):
        pulumi.set(self, "timeout_config", value)


@pulumi.input_type
class DeploymentIoTJobExecutionsRolloutConfigArgs:
    def __init__(__self__, *,
                 exponential_rate: Optional[pulumi.Input['DeploymentIoTJobExponentialRolloutRateArgs']] = None,
                 maximum_per_minute: Optional[pulumi.Input[int]] = None):
        if exponential_rate is not None:
            pulumi.set(__self__, "exponential_rate", exponential_rate)
        if maximum_per_minute is not None:
            pulumi.set(__self__, "maximum_per_minute", maximum_per_minute)

    @property
    @pulumi.getter(name="exponentialRate")
    def exponential_rate(self) -> Optional[pulumi.Input['DeploymentIoTJobExponentialRolloutRateArgs']]:
        return pulumi.get(self, "exponential_rate")

    @exponential_rate.setter
    def exponential_rate(self, value: Optional[pulumi.Input['DeploymentIoTJobExponentialRolloutRateArgs']]):
        pulumi.set(self, "exponential_rate", value)

    @property
    @pulumi.getter(name="maximumPerMinute")
    def maximum_per_minute(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "maximum_per_minute")

    @maximum_per_minute.setter
    def maximum_per_minute(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "maximum_per_minute", value)


@pulumi.input_type
class DeploymentIoTJobExponentialRolloutRateArgs:
    def __init__(__self__, *,
                 base_rate_per_minute: pulumi.Input[int],
                 increment_factor: pulumi.Input[float],
                 rate_increase_criteria: pulumi.Input['DeploymentIoTJobRateIncreaseCriteriaArgs']):
        pulumi.set(__self__, "base_rate_per_minute", base_rate_per_minute)
        pulumi.set(__self__, "increment_factor", increment_factor)
        pulumi.set(__self__, "rate_increase_criteria", rate_increase_criteria)

    @property
    @pulumi.getter(name="baseRatePerMinute")
    def base_rate_per_minute(self) -> pulumi.Input[int]:
        return pulumi.get(self, "base_rate_per_minute")

    @base_rate_per_minute.setter
    def base_rate_per_minute(self, value: pulumi.Input[int]):
        pulumi.set(self, "base_rate_per_minute", value)

    @property
    @pulumi.getter(name="incrementFactor")
    def increment_factor(self) -> pulumi.Input[float]:
        return pulumi.get(self, "increment_factor")

    @increment_factor.setter
    def increment_factor(self, value: pulumi.Input[float]):
        pulumi.set(self, "increment_factor", value)

    @property
    @pulumi.getter(name="rateIncreaseCriteria")
    def rate_increase_criteria(self) -> pulumi.Input['DeploymentIoTJobRateIncreaseCriteriaArgs']:
        return pulumi.get(self, "rate_increase_criteria")

    @rate_increase_criteria.setter
    def rate_increase_criteria(self, value: pulumi.Input['DeploymentIoTJobRateIncreaseCriteriaArgs']):
        pulumi.set(self, "rate_increase_criteria", value)


@pulumi.input_type
class DeploymentIoTJobRateIncreaseCriteriaArgs:
    def __init__(__self__):
        pass


@pulumi.input_type
class DeploymentIoTJobTimeoutConfigArgs:
    def __init__(__self__, *,
                 in_progress_timeout_in_minutes: Optional[pulumi.Input[int]] = None):
        if in_progress_timeout_in_minutes is not None:
            pulumi.set(__self__, "in_progress_timeout_in_minutes", in_progress_timeout_in_minutes)

    @property
    @pulumi.getter(name="inProgressTimeoutInMinutes")
    def in_progress_timeout_in_minutes(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "in_progress_timeout_in_minutes")

    @in_progress_timeout_in_minutes.setter
    def in_progress_timeout_in_minutes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "in_progress_timeout_in_minutes", value)


@pulumi.input_type
class DeploymentPoliciesArgs:
    def __init__(__self__, *,
                 component_update_policy: Optional[pulumi.Input['DeploymentComponentUpdatePolicyArgs']] = None,
                 configuration_validation_policy: Optional[pulumi.Input['DeploymentConfigurationValidationPolicyArgs']] = None,
                 failure_handling_policy: Optional[pulumi.Input['DeploymentPoliciesFailureHandlingPolicy']] = None):
        if component_update_policy is not None:
            pulumi.set(__self__, "component_update_policy", component_update_policy)
        if configuration_validation_policy is not None:
            pulumi.set(__self__, "configuration_validation_policy", configuration_validation_policy)
        if failure_handling_policy is not None:
            pulumi.set(__self__, "failure_handling_policy", failure_handling_policy)

    @property
    @pulumi.getter(name="componentUpdatePolicy")
    def component_update_policy(self) -> Optional[pulumi.Input['DeploymentComponentUpdatePolicyArgs']]:
        return pulumi.get(self, "component_update_policy")

    @component_update_policy.setter
    def component_update_policy(self, value: Optional[pulumi.Input['DeploymentComponentUpdatePolicyArgs']]):
        pulumi.set(self, "component_update_policy", value)

    @property
    @pulumi.getter(name="configurationValidationPolicy")
    def configuration_validation_policy(self) -> Optional[pulumi.Input['DeploymentConfigurationValidationPolicyArgs']]:
        return pulumi.get(self, "configuration_validation_policy")

    @configuration_validation_policy.setter
    def configuration_validation_policy(self, value: Optional[pulumi.Input['DeploymentConfigurationValidationPolicyArgs']]):
        pulumi.set(self, "configuration_validation_policy", value)

    @property
    @pulumi.getter(name="failureHandlingPolicy")
    def failure_handling_policy(self) -> Optional[pulumi.Input['DeploymentPoliciesFailureHandlingPolicy']]:
        return pulumi.get(self, "failure_handling_policy")

    @failure_handling_policy.setter
    def failure_handling_policy(self, value: Optional[pulumi.Input['DeploymentPoliciesFailureHandlingPolicy']]):
        pulumi.set(self, "failure_handling_policy", value)


@pulumi.input_type
class DeploymentSystemResourceLimitsArgs:
    def __init__(__self__, *,
                 cpus: Optional[pulumi.Input[float]] = None,
                 memory: Optional[pulumi.Input[int]] = None):
        if cpus is not None:
            pulumi.set(__self__, "cpus", cpus)
        if memory is not None:
            pulumi.set(__self__, "memory", memory)

    @property
    @pulumi.getter
    def cpus(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "cpus")

    @cpus.setter
    def cpus(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "cpus", value)

    @property
    @pulumi.getter
    def memory(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "memory")

    @memory.setter
    def memory(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "memory", value)


