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
    'TopicLoggingConfig',
    'TopicSubscription',
    'TopicTag',
]

@pulumi.output_type
class TopicLoggingConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "failureFeedbackRoleArn":
            suggest = "failure_feedback_role_arn"
        elif key == "successFeedbackRoleArn":
            suggest = "success_feedback_role_arn"
        elif key == "successFeedbackSampleRate":
            suggest = "success_feedback_sample_rate"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TopicLoggingConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TopicLoggingConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TopicLoggingConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 protocol: 'TopicLoggingConfigProtocol',
                 failure_feedback_role_arn: Optional[str] = None,
                 success_feedback_role_arn: Optional[str] = None,
                 success_feedback_sample_rate: Optional[str] = None):
        """
        :param 'TopicLoggingConfigProtocol' protocol: Indicates one of the supported protocols for the SNS topic
        :param str failure_feedback_role_arn: The IAM role ARN to be used when logging failed message deliveries in Amazon CloudWatch
        :param str success_feedback_role_arn: The IAM role ARN to be used when logging successful message deliveries in Amazon CloudWatch
        :param str success_feedback_sample_rate: The percentage of successful message deliveries to be logged in Amazon CloudWatch. Valid percentage values range from 0 to 100
        """
        pulumi.set(__self__, "protocol", protocol)
        if failure_feedback_role_arn is not None:
            pulumi.set(__self__, "failure_feedback_role_arn", failure_feedback_role_arn)
        if success_feedback_role_arn is not None:
            pulumi.set(__self__, "success_feedback_role_arn", success_feedback_role_arn)
        if success_feedback_sample_rate is not None:
            pulumi.set(__self__, "success_feedback_sample_rate", success_feedback_sample_rate)

    @property
    @pulumi.getter
    def protocol(self) -> 'TopicLoggingConfigProtocol':
        """
        Indicates one of the supported protocols for the SNS topic
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="failureFeedbackRoleArn")
    def failure_feedback_role_arn(self) -> Optional[str]:
        """
        The IAM role ARN to be used when logging failed message deliveries in Amazon CloudWatch
        """
        return pulumi.get(self, "failure_feedback_role_arn")

    @property
    @pulumi.getter(name="successFeedbackRoleArn")
    def success_feedback_role_arn(self) -> Optional[str]:
        """
        The IAM role ARN to be used when logging successful message deliveries in Amazon CloudWatch
        """
        return pulumi.get(self, "success_feedback_role_arn")

    @property
    @pulumi.getter(name="successFeedbackSampleRate")
    def success_feedback_sample_rate(self) -> Optional[str]:
        """
        The percentage of successful message deliveries to be logged in Amazon CloudWatch. Valid percentage values range from 0 to 100
        """
        return pulumi.get(self, "success_feedback_sample_rate")


@pulumi.output_type
class TopicSubscription(dict):
    def __init__(__self__, *,
                 endpoint: str,
                 protocol: str):
        pulumi.set(__self__, "endpoint", endpoint)
        pulumi.set(__self__, "protocol", protocol)

    @property
    @pulumi.getter
    def endpoint(self) -> str:
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        return pulumi.get(self, "protocol")


@pulumi.output_type
class TopicTag(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        """
        :param str key: The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, `_`, `.`, `/`, `=`, `+`, and `-`.
        :param str value: The value for the tag. You can specify a value that is 0 to 256 characters in length.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, `_`, `.`, `/`, `=`, `+`, and `-`.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value for the tag. You can specify a value that is 0 to 256 characters in length.
        """
        return pulumi.get(self, "value")


