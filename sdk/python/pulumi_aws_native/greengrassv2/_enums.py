# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'ComponentVersionComponentDependencyRequirementDependencyType',
    'ComponentVersionLambdaEventSourceType',
    'ComponentVersionLambdaExecutionParametersInputPayloadEncodingType',
    'ComponentVersionLambdaFilesystemPermission',
    'ComponentVersionLambdaLinuxProcessParamsIsolationMode',
    'DeploymentComponentUpdatePolicyAction',
    'DeploymentIoTJobAbortCriteriaAction',
    'DeploymentIoTJobAbortCriteriaFailureType',
    'DeploymentPoliciesFailureHandlingPolicy',
]


class ComponentVersionComponentDependencyRequirementDependencyType(str, Enum):
    SOFT = "SOFT"
    HARD = "HARD"


class ComponentVersionLambdaEventSourceType(str, Enum):
    PUB_SUB = "PUB_SUB"
    IOT_CORE = "IOT_CORE"


class ComponentVersionLambdaExecutionParametersInputPayloadEncodingType(str, Enum):
    JSON = "json"
    BINARY = "binary"


class ComponentVersionLambdaFilesystemPermission(str, Enum):
    RO = "ro"
    RW = "rw"


class ComponentVersionLambdaLinuxProcessParamsIsolationMode(str, Enum):
    GREENGRASS_CONTAINER = "GreengrassContainer"
    NO_CONTAINER = "NoContainer"


class DeploymentComponentUpdatePolicyAction(str, Enum):
    NOTIFY_COMPONENTS = "NOTIFY_COMPONENTS"
    SKIP_NOTIFY_COMPONENTS = "SKIP_NOTIFY_COMPONENTS"


class DeploymentIoTJobAbortCriteriaAction(str, Enum):
    CANCEL = "CANCEL"


class DeploymentIoTJobAbortCriteriaFailureType(str, Enum):
    FAILED = "FAILED"
    REJECTED = "REJECTED"
    TIMED_OUT = "TIMED_OUT"
    ALL = "ALL"


class DeploymentPoliciesFailureHandlingPolicy(str, Enum):
    ROLLBACK = "ROLLBACK"
    DO_NOTHING = "DO_NOTHING"
