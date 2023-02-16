# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ApplicationMaxAgeRuleArgs',
    'ApplicationMaxCountRuleArgs',
    'ApplicationResourceLifecycleConfigArgs',
    'ApplicationVersionLifecycleConfigArgs',
    'ApplicationVersionSourceBundleArgs',
    'ConfigurationTemplateConfigurationOptionSettingArgs',
    'ConfigurationTemplateSourceConfigurationArgs',
    'EnvironmentOptionSettingArgs',
    'EnvironmentTagArgs',
    'EnvironmentTierArgs',
]

@pulumi.input_type
class ApplicationMaxAgeRuleArgs:
    def __init__(__self__, *,
                 delete_source_from_s3: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 max_age_in_days: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[bool] delete_source_from_s3: Set to true to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.
        :param pulumi.Input[bool] enabled: Specify true to apply the rule, or false to disable it.
        :param pulumi.Input[int] max_age_in_days: Specify the number of days to retain an application versions.
        """
        if delete_source_from_s3 is not None:
            pulumi.set(__self__, "delete_source_from_s3", delete_source_from_s3)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if max_age_in_days is not None:
            pulumi.set(__self__, "max_age_in_days", max_age_in_days)

    @property
    @pulumi.getter(name="deleteSourceFromS3")
    def delete_source_from_s3(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to true to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.
        """
        return pulumi.get(self, "delete_source_from_s3")

    @delete_source_from_s3.setter
    def delete_source_from_s3(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_source_from_s3", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specify true to apply the rule, or false to disable it.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="maxAgeInDays")
    def max_age_in_days(self) -> Optional[pulumi.Input[int]]:
        """
        Specify the number of days to retain an application versions.
        """
        return pulumi.get(self, "max_age_in_days")

    @max_age_in_days.setter
    def max_age_in_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_age_in_days", value)


@pulumi.input_type
class ApplicationMaxCountRuleArgs:
    def __init__(__self__, *,
                 delete_source_from_s3: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 max_count: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[bool] delete_source_from_s3: Set to true to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.
        :param pulumi.Input[bool] enabled: Specify true to apply the rule, or false to disable it.
        :param pulumi.Input[int] max_count: Specify the maximum number of application versions to retain.
        """
        if delete_source_from_s3 is not None:
            pulumi.set(__self__, "delete_source_from_s3", delete_source_from_s3)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if max_count is not None:
            pulumi.set(__self__, "max_count", max_count)

    @property
    @pulumi.getter(name="deleteSourceFromS3")
    def delete_source_from_s3(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to true to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.
        """
        return pulumi.get(self, "delete_source_from_s3")

    @delete_source_from_s3.setter
    def delete_source_from_s3(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_source_from_s3", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specify true to apply the rule, or false to disable it.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="maxCount")
    def max_count(self) -> Optional[pulumi.Input[int]]:
        """
        Specify the maximum number of application versions to retain.
        """
        return pulumi.get(self, "max_count")

    @max_count.setter
    def max_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_count", value)


@pulumi.input_type
class ApplicationResourceLifecycleConfigArgs:
    def __init__(__self__, *,
                 service_role: Optional[pulumi.Input[str]] = None,
                 version_lifecycle_config: Optional[pulumi.Input['ApplicationVersionLifecycleConfigArgs']] = None):
        """
        :param pulumi.Input[str] service_role: The ARN of an IAM service role that Elastic Beanstalk has permission to assume. The ServiceRole property is required the first time that you provide a ResourceLifecycleConfig for the application. After you provide it once, Elastic Beanstalk persists the Service Role with the application, and you don't need to specify it again. You can, however, specify it in subsequent updates to change the Service Role to another value.
        :param pulumi.Input['ApplicationVersionLifecycleConfigArgs'] version_lifecycle_config: Defines lifecycle settings for application versions.
        """
        if service_role is not None:
            pulumi.set(__self__, "service_role", service_role)
        if version_lifecycle_config is not None:
            pulumi.set(__self__, "version_lifecycle_config", version_lifecycle_config)

    @property
    @pulumi.getter(name="serviceRole")
    def service_role(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of an IAM service role that Elastic Beanstalk has permission to assume. The ServiceRole property is required the first time that you provide a ResourceLifecycleConfig for the application. After you provide it once, Elastic Beanstalk persists the Service Role with the application, and you don't need to specify it again. You can, however, specify it in subsequent updates to change the Service Role to another value.
        """
        return pulumi.get(self, "service_role")

    @service_role.setter
    def service_role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_role", value)

    @property
    @pulumi.getter(name="versionLifecycleConfig")
    def version_lifecycle_config(self) -> Optional[pulumi.Input['ApplicationVersionLifecycleConfigArgs']]:
        """
        Defines lifecycle settings for application versions.
        """
        return pulumi.get(self, "version_lifecycle_config")

    @version_lifecycle_config.setter
    def version_lifecycle_config(self, value: Optional[pulumi.Input['ApplicationVersionLifecycleConfigArgs']]):
        pulumi.set(self, "version_lifecycle_config", value)


@pulumi.input_type
class ApplicationVersionLifecycleConfigArgs:
    def __init__(__self__, *,
                 max_age_rule: Optional[pulumi.Input['ApplicationMaxAgeRuleArgs']] = None,
                 max_count_rule: Optional[pulumi.Input['ApplicationMaxCountRuleArgs']] = None):
        """
        :param pulumi.Input['ApplicationMaxAgeRuleArgs'] max_age_rule: Specify a max age rule to restrict the length of time that application versions are retained for an application.
        :param pulumi.Input['ApplicationMaxCountRuleArgs'] max_count_rule: Specify a max count rule to restrict the number of application versions that are retained for an application.
        """
        if max_age_rule is not None:
            pulumi.set(__self__, "max_age_rule", max_age_rule)
        if max_count_rule is not None:
            pulumi.set(__self__, "max_count_rule", max_count_rule)

    @property
    @pulumi.getter(name="maxAgeRule")
    def max_age_rule(self) -> Optional[pulumi.Input['ApplicationMaxAgeRuleArgs']]:
        """
        Specify a max age rule to restrict the length of time that application versions are retained for an application.
        """
        return pulumi.get(self, "max_age_rule")

    @max_age_rule.setter
    def max_age_rule(self, value: Optional[pulumi.Input['ApplicationMaxAgeRuleArgs']]):
        pulumi.set(self, "max_age_rule", value)

    @property
    @pulumi.getter(name="maxCountRule")
    def max_count_rule(self) -> Optional[pulumi.Input['ApplicationMaxCountRuleArgs']]:
        """
        Specify a max count rule to restrict the number of application versions that are retained for an application.
        """
        return pulumi.get(self, "max_count_rule")

    @max_count_rule.setter
    def max_count_rule(self, value: Optional[pulumi.Input['ApplicationMaxCountRuleArgs']]):
        pulumi.set(self, "max_count_rule", value)


@pulumi.input_type
class ApplicationVersionSourceBundleArgs:
    def __init__(__self__, *,
                 s3_bucket: pulumi.Input[str],
                 s3_key: pulumi.Input[str]):
        """
        :param pulumi.Input[str] s3_bucket: The Amazon S3 bucket where the data is located.
        :param pulumi.Input[str] s3_key: The Amazon S3 key where the data is located.
        """
        pulumi.set(__self__, "s3_bucket", s3_bucket)
        pulumi.set(__self__, "s3_key", s3_key)

    @property
    @pulumi.getter(name="s3Bucket")
    def s3_bucket(self) -> pulumi.Input[str]:
        """
        The Amazon S3 bucket where the data is located.
        """
        return pulumi.get(self, "s3_bucket")

    @s3_bucket.setter
    def s3_bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "s3_bucket", value)

    @property
    @pulumi.getter(name="s3Key")
    def s3_key(self) -> pulumi.Input[str]:
        """
        The Amazon S3 key where the data is located.
        """
        return pulumi.get(self, "s3_key")

    @s3_key.setter
    def s3_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "s3_key", value)


@pulumi.input_type
class ConfigurationTemplateConfigurationOptionSettingArgs:
    def __init__(__self__, *,
                 namespace: pulumi.Input[str],
                 option_name: pulumi.Input[str],
                 resource_name: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] namespace: A unique namespace that identifies the option's associated AWS resource.
        :param pulumi.Input[str] option_name: The name of the configuration option.
        :param pulumi.Input[str] resource_name: A unique resource name for the option setting. Use it for a time–based scaling configuration option. 
        :param pulumi.Input[str] value: The current value for the configuration option.
        """
        pulumi.set(__self__, "namespace", namespace)
        pulumi.set(__self__, "option_name", option_name)
        if resource_name is not None:
            pulumi.set(__self__, "resource_name", resource_name)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Input[str]:
        """
        A unique namespace that identifies the option's associated AWS resource.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="optionName")
    def option_name(self) -> pulumi.Input[str]:
        """
        The name of the configuration option.
        """
        return pulumi.get(self, "option_name")

    @option_name.setter
    def option_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "option_name", value)

    @property
    @pulumi.getter(name="resourceName")
    def resource_name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique resource name for the option setting. Use it for a time–based scaling configuration option. 
        """
        return pulumi.get(self, "resource_name")

    @resource_name.setter
    def resource_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_name", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The current value for the configuration option.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ConfigurationTemplateSourceConfigurationArgs:
    def __init__(__self__, *,
                 application_name: pulumi.Input[str],
                 template_name: pulumi.Input[str]):
        """
        :param pulumi.Input[str] application_name: The name of the application associated with the configuration.
        :param pulumi.Input[str] template_name: The name of the configuration template.
        """
        pulumi.set(__self__, "application_name", application_name)
        pulumi.set(__self__, "template_name", template_name)

    @property
    @pulumi.getter(name="applicationName")
    def application_name(self) -> pulumi.Input[str]:
        """
        The name of the application associated with the configuration.
        """
        return pulumi.get(self, "application_name")

    @application_name.setter
    def application_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "application_name", value)

    @property
    @pulumi.getter(name="templateName")
    def template_name(self) -> pulumi.Input[str]:
        """
        The name of the configuration template.
        """
        return pulumi.get(self, "template_name")

    @template_name.setter
    def template_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "template_name", value)


@pulumi.input_type
class EnvironmentOptionSettingArgs:
    def __init__(__self__, *,
                 namespace: pulumi.Input[str],
                 option_name: pulumi.Input[str],
                 resource_name: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] namespace: A unique namespace that identifies the option's associated AWS resource.
        :param pulumi.Input[str] option_name: The name of the configuration option.
        :param pulumi.Input[str] resource_name: A unique resource name for the option setting. Use it for a time–based scaling configuration option.
        :param pulumi.Input[str] value: The current value for the configuration option.
        """
        pulumi.set(__self__, "namespace", namespace)
        pulumi.set(__self__, "option_name", option_name)
        if resource_name is not None:
            pulumi.set(__self__, "resource_name", resource_name)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Input[str]:
        """
        A unique namespace that identifies the option's associated AWS resource.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="optionName")
    def option_name(self) -> pulumi.Input[str]:
        """
        The name of the configuration option.
        """
        return pulumi.get(self, "option_name")

    @option_name.setter
    def option_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "option_name", value)

    @property
    @pulumi.getter(name="resourceName")
    def resource_name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique resource name for the option setting. Use it for a time–based scaling configuration option.
        """
        return pulumi.get(self, "resource_name")

    @resource_name.setter
    def resource_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_name", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The current value for the configuration option.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class EnvironmentTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        :param pulumi.Input[str] key: The key name of the tag.
        :param pulumi.Input[str] value: The value for the tag.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The key name of the tag.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value for the tag.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class EnvironmentTierArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: The name of this environment tier.
        :param pulumi.Input[str] type: The type of this environment tier.
        :param pulumi.Input[str] version: The version of this environment tier. When you don't set a value to it, Elastic Beanstalk uses the latest compatible worker tier version.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of this environment tier.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of this environment tier.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of this environment tier. When you don't set a value to it, Elastic Beanstalk uses the latest compatible worker tier version.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


