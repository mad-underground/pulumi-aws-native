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

__all__ = [
    'GroupInsightsConfiguration',
    'GroupTag',
    'SamplingRule',
    'SamplingRuleRecord',
    'SamplingRuleTag',
    'SamplingRuleUpdate',
]

@pulumi.output_type
class GroupInsightsConfiguration(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "insightsEnabled":
            suggest = "insights_enabled"
        elif key == "notificationsEnabled":
            suggest = "notifications_enabled"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in GroupInsightsConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        GroupInsightsConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        GroupInsightsConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 insights_enabled: Optional[bool] = None,
                 notifications_enabled: Optional[bool] = None):
        """
        :param bool insights_enabled: Set the InsightsEnabled value to true to enable insights or false to disable insights.
        :param bool notifications_enabled: Set the NotificationsEnabled value to true to enable insights notifications. Notifications can only be enabled on a group with InsightsEnabled set to true.
        """
        if insights_enabled is not None:
            pulumi.set(__self__, "insights_enabled", insights_enabled)
        if notifications_enabled is not None:
            pulumi.set(__self__, "notifications_enabled", notifications_enabled)

    @property
    @pulumi.getter(name="insightsEnabled")
    def insights_enabled(self) -> Optional[bool]:
        """
        Set the InsightsEnabled value to true to enable insights or false to disable insights.
        """
        return pulumi.get(self, "insights_enabled")

    @property
    @pulumi.getter(name="notificationsEnabled")
    def notifications_enabled(self) -> Optional[bool]:
        """
        Set the NotificationsEnabled value to true to enable insights notifications. Notifications can only be enabled on a group with InsightsEnabled set to true.
        """
        return pulumi.get(self, "notifications_enabled")


@pulumi.output_type
class GroupTag(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        """
        :param str key: The key name of the tag.
        :param str value: The value for the tag.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The key name of the tag.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value for the tag.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class SamplingRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "fixedRate":
            suggest = "fixed_rate"
        elif key == "httpMethod":
            suggest = "http_method"
        elif key == "reservoirSize":
            suggest = "reservoir_size"
        elif key == "resourceArn":
            suggest = "resource_arn"
        elif key == "serviceName":
            suggest = "service_name"
        elif key == "serviceType":
            suggest = "service_type"
        elif key == "urlPath":
            suggest = "url_path"
        elif key == "ruleArn":
            suggest = "rule_arn"
        elif key == "ruleName":
            suggest = "rule_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SamplingRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SamplingRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SamplingRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 fixed_rate: float,
                 host: str,
                 http_method: str,
                 priority: int,
                 reservoir_size: int,
                 resource_arn: str,
                 service_name: str,
                 service_type: str,
                 url_path: str,
                 attributes: Optional[Mapping[str, str]] = None,
                 rule_arn: Optional[str] = None,
                 rule_name: Optional[str] = None,
                 version: Optional[int] = None):
        """
        :param float fixed_rate: The percentage of matching requests to instrument, after the reservoir is exhausted.
        :param str host: Matches the hostname from a request URL.
        :param str http_method: Matches the HTTP method from a request URL.
        :param int priority: The priority of the sampling rule.
        :param int reservoir_size: A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.
        :param str resource_arn: Matches the ARN of the AWS resource on which the service runs.
        :param str service_name: Matches the name that the service uses to identify itself in segments.
        :param str service_type: Matches the origin that the service uses to identify its type in segments.
        :param str url_path: Matches the path from a request URL.
        :param Mapping[str, str] attributes: Matches attributes derived from the request.
        :param int version: The version of the sampling rule format (1)
        """
        pulumi.set(__self__, "fixed_rate", fixed_rate)
        pulumi.set(__self__, "host", host)
        pulumi.set(__self__, "http_method", http_method)
        pulumi.set(__self__, "priority", priority)
        pulumi.set(__self__, "reservoir_size", reservoir_size)
        pulumi.set(__self__, "resource_arn", resource_arn)
        pulumi.set(__self__, "service_name", service_name)
        pulumi.set(__self__, "service_type", service_type)
        pulumi.set(__self__, "url_path", url_path)
        if attributes is not None:
            pulumi.set(__self__, "attributes", attributes)
        if rule_arn is not None:
            pulumi.set(__self__, "rule_arn", rule_arn)
        if rule_name is not None:
            pulumi.set(__self__, "rule_name", rule_name)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="fixedRate")
    def fixed_rate(self) -> float:
        """
        The percentage of matching requests to instrument, after the reservoir is exhausted.
        """
        return pulumi.get(self, "fixed_rate")

    @property
    @pulumi.getter
    def host(self) -> str:
        """
        Matches the hostname from a request URL.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter(name="httpMethod")
    def http_method(self) -> str:
        """
        Matches the HTTP method from a request URL.
        """
        return pulumi.get(self, "http_method")

    @property
    @pulumi.getter
    def priority(self) -> int:
        """
        The priority of the sampling rule.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="reservoirSize")
    def reservoir_size(self) -> int:
        """
        A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.
        """
        return pulumi.get(self, "reservoir_size")

    @property
    @pulumi.getter(name="resourceArn")
    def resource_arn(self) -> str:
        """
        Matches the ARN of the AWS resource on which the service runs.
        """
        return pulumi.get(self, "resource_arn")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        """
        Matches the name that the service uses to identify itself in segments.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="serviceType")
    def service_type(self) -> str:
        """
        Matches the origin that the service uses to identify its type in segments.
        """
        return pulumi.get(self, "service_type")

    @property
    @pulumi.getter(name="urlPath")
    def url_path(self) -> str:
        """
        Matches the path from a request URL.
        """
        return pulumi.get(self, "url_path")

    @property
    @pulumi.getter
    def attributes(self) -> Optional[Mapping[str, str]]:
        """
        Matches attributes derived from the request.
        """
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter(name="ruleArn")
    def rule_arn(self) -> Optional[str]:
        return pulumi.get(self, "rule_arn")

    @property
    @pulumi.getter(name="ruleName")
    def rule_name(self) -> Optional[str]:
        return pulumi.get(self, "rule_name")

    @property
    @pulumi.getter
    def version(self) -> Optional[int]:
        """
        The version of the sampling rule format (1)
        """
        return pulumi.get(self, "version")


@pulumi.output_type
class SamplingRuleRecord(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createdAt":
            suggest = "created_at"
        elif key == "modifiedAt":
            suggest = "modified_at"
        elif key == "samplingRule":
            suggest = "sampling_rule"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SamplingRuleRecord. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SamplingRuleRecord.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SamplingRuleRecord.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 created_at: Optional[str] = None,
                 modified_at: Optional[str] = None,
                 sampling_rule: Optional['outputs.SamplingRule'] = None):
        """
        :param str created_at: When the rule was created, in Unix time seconds.
        :param str modified_at: When the rule was modified, in Unix time seconds.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if modified_at is not None:
            pulumi.set(__self__, "modified_at", modified_at)
        if sampling_rule is not None:
            pulumi.set(__self__, "sampling_rule", sampling_rule)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[str]:
        """
        When the rule was created, in Unix time seconds.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="modifiedAt")
    def modified_at(self) -> Optional[str]:
        """
        When the rule was modified, in Unix time seconds.
        """
        return pulumi.get(self, "modified_at")

    @property
    @pulumi.getter(name="samplingRule")
    def sampling_rule(self) -> Optional['outputs.SamplingRule']:
        return pulumi.get(self, "sampling_rule")


@pulumi.output_type
class SamplingRuleTag(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        """
        :param str key: The key name of the tag.
        :param str value: The value for the tag.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The key name of the tag.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value for the tag.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class SamplingRuleUpdate(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "fixedRate":
            suggest = "fixed_rate"
        elif key == "httpMethod":
            suggest = "http_method"
        elif key == "reservoirSize":
            suggest = "reservoir_size"
        elif key == "resourceArn":
            suggest = "resource_arn"
        elif key == "ruleArn":
            suggest = "rule_arn"
        elif key == "ruleName":
            suggest = "rule_name"
        elif key == "serviceName":
            suggest = "service_name"
        elif key == "serviceType":
            suggest = "service_type"
        elif key == "urlPath":
            suggest = "url_path"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SamplingRuleUpdate. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SamplingRuleUpdate.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SamplingRuleUpdate.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 attributes: Optional[Mapping[str, str]] = None,
                 fixed_rate: Optional[float] = None,
                 host: Optional[str] = None,
                 http_method: Optional[str] = None,
                 priority: Optional[int] = None,
                 reservoir_size: Optional[int] = None,
                 resource_arn: Optional[str] = None,
                 rule_arn: Optional[str] = None,
                 rule_name: Optional[str] = None,
                 service_name: Optional[str] = None,
                 service_type: Optional[str] = None,
                 url_path: Optional[str] = None):
        """
        :param Mapping[str, str] attributes: Matches attributes derived from the request.
        :param float fixed_rate: The percentage of matching requests to instrument, after the reservoir is exhausted.
        :param str host: Matches the hostname from a request URL.
        :param str http_method: Matches the HTTP method from a request URL.
        :param int priority: The priority of the sampling rule.
        :param int reservoir_size: A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.
        :param str resource_arn: Matches the ARN of the AWS resource on which the service runs.
        :param str service_name: Matches the name that the service uses to identify itself in segments.
        :param str service_type: Matches the origin that the service uses to identify its type in segments.
        :param str url_path: Matches the path from a request URL.
        """
        if attributes is not None:
            pulumi.set(__self__, "attributes", attributes)
        if fixed_rate is not None:
            pulumi.set(__self__, "fixed_rate", fixed_rate)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if http_method is not None:
            pulumi.set(__self__, "http_method", http_method)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if reservoir_size is not None:
            pulumi.set(__self__, "reservoir_size", reservoir_size)
        if resource_arn is not None:
            pulumi.set(__self__, "resource_arn", resource_arn)
        if rule_arn is not None:
            pulumi.set(__self__, "rule_arn", rule_arn)
        if rule_name is not None:
            pulumi.set(__self__, "rule_name", rule_name)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if service_type is not None:
            pulumi.set(__self__, "service_type", service_type)
        if url_path is not None:
            pulumi.set(__self__, "url_path", url_path)

    @property
    @pulumi.getter
    def attributes(self) -> Optional[Mapping[str, str]]:
        """
        Matches attributes derived from the request.
        """
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter(name="fixedRate")
    def fixed_rate(self) -> Optional[float]:
        """
        The percentage of matching requests to instrument, after the reservoir is exhausted.
        """
        return pulumi.get(self, "fixed_rate")

    @property
    @pulumi.getter
    def host(self) -> Optional[str]:
        """
        Matches the hostname from a request URL.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter(name="httpMethod")
    def http_method(self) -> Optional[str]:
        """
        Matches the HTTP method from a request URL.
        """
        return pulumi.get(self, "http_method")

    @property
    @pulumi.getter
    def priority(self) -> Optional[int]:
        """
        The priority of the sampling rule.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="reservoirSize")
    def reservoir_size(self) -> Optional[int]:
        """
        A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.
        """
        return pulumi.get(self, "reservoir_size")

    @property
    @pulumi.getter(name="resourceArn")
    def resource_arn(self) -> Optional[str]:
        """
        Matches the ARN of the AWS resource on which the service runs.
        """
        return pulumi.get(self, "resource_arn")

    @property
    @pulumi.getter(name="ruleArn")
    def rule_arn(self) -> Optional[str]:
        return pulumi.get(self, "rule_arn")

    @property
    @pulumi.getter(name="ruleName")
    def rule_name(self) -> Optional[str]:
        return pulumi.get(self, "rule_name")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[str]:
        """
        Matches the name that the service uses to identify itself in segments.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="serviceType")
    def service_type(self) -> Optional[str]:
        """
        Matches the origin that the service uses to identify its type in segments.
        """
        return pulumi.get(self, "service_type")

    @property
    @pulumi.getter(name="urlPath")
    def url_path(self) -> Optional[str]:
        """
        Matches the path from a request URL.
        """
        return pulumi.get(self, "url_path")


