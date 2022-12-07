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
    'ReplicationSetRegionConfigurationArgs',
    'ReplicationSetReplicationRegionArgs',
    'ResponsePlanActionArgs',
    'ResponsePlanChatChannelArgs',
    'ResponsePlanDynamicSsmParameterValueArgs',
    'ResponsePlanDynamicSsmParameterArgs',
    'ResponsePlanIncidentTemplateArgs',
    'ResponsePlanIntegrationArgs',
    'ResponsePlanNotificationTargetItemArgs',
    'ResponsePlanPagerDutyConfigurationArgs',
    'ResponsePlanPagerDutyIncidentConfigurationArgs',
    'ResponsePlanSsmAutomationArgs',
    'ResponsePlanSsmParameterArgs',
    'ResponsePlanTagArgs',
]

@pulumi.input_type
class ReplicationSetRegionConfigurationArgs:
    def __init__(__self__, *,
                 sse_kms_key_id: pulumi.Input[str]):
        """
        The ReplicationSet regional configuration.
        """
        pulumi.set(__self__, "sse_kms_key_id", sse_kms_key_id)

    @property
    @pulumi.getter(name="sseKmsKeyId")
    def sse_kms_key_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "sse_kms_key_id")

    @sse_kms_key_id.setter
    def sse_kms_key_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "sse_kms_key_id", value)


@pulumi.input_type
class ReplicationSetReplicationRegionArgs:
    def __init__(__self__, *,
                 region_configuration: Optional[pulumi.Input['ReplicationSetRegionConfigurationArgs']] = None,
                 region_name: Optional[pulumi.Input[str]] = None):
        """
        The ReplicationSet regional configuration.
        """
        if region_configuration is not None:
            pulumi.set(__self__, "region_configuration", region_configuration)
        if region_name is not None:
            pulumi.set(__self__, "region_name", region_name)

    @property
    @pulumi.getter(name="regionConfiguration")
    def region_configuration(self) -> Optional[pulumi.Input['ReplicationSetRegionConfigurationArgs']]:
        return pulumi.get(self, "region_configuration")

    @region_configuration.setter
    def region_configuration(self, value: Optional[pulumi.Input['ReplicationSetRegionConfigurationArgs']]):
        pulumi.set(self, "region_configuration", value)

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region_name")

    @region_name.setter
    def region_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region_name", value)


@pulumi.input_type
class ResponsePlanActionArgs:
    def __init__(__self__, *,
                 ssm_automation: Optional[pulumi.Input['ResponsePlanSsmAutomationArgs']] = None):
        """
        The automation configuration to launch.
        """
        if ssm_automation is not None:
            pulumi.set(__self__, "ssm_automation", ssm_automation)

    @property
    @pulumi.getter(name="ssmAutomation")
    def ssm_automation(self) -> Optional[pulumi.Input['ResponsePlanSsmAutomationArgs']]:
        return pulumi.get(self, "ssm_automation")

    @ssm_automation.setter
    def ssm_automation(self, value: Optional[pulumi.Input['ResponsePlanSsmAutomationArgs']]):
        pulumi.set(self, "ssm_automation", value)


@pulumi.input_type
class ResponsePlanChatChannelArgs:
    def __init__(__self__, *,
                 chatbot_sns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The chat channel configuration.
        """
        if chatbot_sns is not None:
            pulumi.set(__self__, "chatbot_sns", chatbot_sns)

    @property
    @pulumi.getter(name="chatbotSns")
    def chatbot_sns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "chatbot_sns")

    @chatbot_sns.setter
    def chatbot_sns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "chatbot_sns", value)


@pulumi.input_type
class ResponsePlanDynamicSsmParameterValueArgs:
    def __init__(__self__, *,
                 variable: Optional[pulumi.Input['ResponsePlanVariableType']] = None):
        """
        Value of the dynamic parameter to set when starting the SSM automation document.
        """
        if variable is not None:
            pulumi.set(__self__, "variable", variable)

    @property
    @pulumi.getter
    def variable(self) -> Optional[pulumi.Input['ResponsePlanVariableType']]:
        return pulumi.get(self, "variable")

    @variable.setter
    def variable(self, value: Optional[pulumi.Input['ResponsePlanVariableType']]):
        pulumi.set(self, "variable", value)


@pulumi.input_type
class ResponsePlanDynamicSsmParameterArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input['ResponsePlanDynamicSsmParameterValueArgs']):
        """
        A parameter with a dynamic value to set when starting the SSM automation document.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input['ResponsePlanDynamicSsmParameterValueArgs']:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input['ResponsePlanDynamicSsmParameterValueArgs']):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ResponsePlanIncidentTemplateArgs:
    def __init__(__self__, *,
                 impact: pulumi.Input[int],
                 title: pulumi.Input[str],
                 dedupe_string: Optional[pulumi.Input[str]] = None,
                 incident_tags: Optional[pulumi.Input[Sequence[pulumi.Input['ResponsePlanTagArgs']]]] = None,
                 notification_targets: Optional[pulumi.Input[Sequence[pulumi.Input['ResponsePlanNotificationTargetItemArgs']]]] = None,
                 summary: Optional[pulumi.Input[str]] = None):
        """
        The incident template configuration.
        :param pulumi.Input[int] impact: The impact value.
        :param pulumi.Input[str] title: The title string.
        :param pulumi.Input[str] dedupe_string: The deduplication string.
        :param pulumi.Input[Sequence[pulumi.Input['ResponsePlanTagArgs']]] incident_tags: Tags that get applied to incidents created by the StartIncident API action.
        :param pulumi.Input[Sequence[pulumi.Input['ResponsePlanNotificationTargetItemArgs']]] notification_targets: The list of notification targets.
        :param pulumi.Input[str] summary: The summary string.
        """
        pulumi.set(__self__, "impact", impact)
        pulumi.set(__self__, "title", title)
        if dedupe_string is not None:
            pulumi.set(__self__, "dedupe_string", dedupe_string)
        if incident_tags is not None:
            pulumi.set(__self__, "incident_tags", incident_tags)
        if notification_targets is not None:
            pulumi.set(__self__, "notification_targets", notification_targets)
        if summary is not None:
            pulumi.set(__self__, "summary", summary)

    @property
    @pulumi.getter
    def impact(self) -> pulumi.Input[int]:
        """
        The impact value.
        """
        return pulumi.get(self, "impact")

    @impact.setter
    def impact(self, value: pulumi.Input[int]):
        pulumi.set(self, "impact", value)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        """
        The title string.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter(name="dedupeString")
    def dedupe_string(self) -> Optional[pulumi.Input[str]]:
        """
        The deduplication string.
        """
        return pulumi.get(self, "dedupe_string")

    @dedupe_string.setter
    def dedupe_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dedupe_string", value)

    @property
    @pulumi.getter(name="incidentTags")
    def incident_tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ResponsePlanTagArgs']]]]:
        """
        Tags that get applied to incidents created by the StartIncident API action.
        """
        return pulumi.get(self, "incident_tags")

    @incident_tags.setter
    def incident_tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ResponsePlanTagArgs']]]]):
        pulumi.set(self, "incident_tags", value)

    @property
    @pulumi.getter(name="notificationTargets")
    def notification_targets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ResponsePlanNotificationTargetItemArgs']]]]:
        """
        The list of notification targets.
        """
        return pulumi.get(self, "notification_targets")

    @notification_targets.setter
    def notification_targets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ResponsePlanNotificationTargetItemArgs']]]]):
        pulumi.set(self, "notification_targets", value)

    @property
    @pulumi.getter
    def summary(self) -> Optional[pulumi.Input[str]]:
        """
        The summary string.
        """
        return pulumi.get(self, "summary")

    @summary.setter
    def summary(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "summary", value)


@pulumi.input_type
class ResponsePlanIntegrationArgs:
    def __init__(__self__, *,
                 pager_duty_configuration: Optional[pulumi.Input['ResponsePlanPagerDutyConfigurationArgs']] = None):
        if pager_duty_configuration is not None:
            pulumi.set(__self__, "pager_duty_configuration", pager_duty_configuration)

    @property
    @pulumi.getter(name="pagerDutyConfiguration")
    def pager_duty_configuration(self) -> Optional[pulumi.Input['ResponsePlanPagerDutyConfigurationArgs']]:
        return pulumi.get(self, "pager_duty_configuration")

    @pager_duty_configuration.setter
    def pager_duty_configuration(self, value: Optional[pulumi.Input['ResponsePlanPagerDutyConfigurationArgs']]):
        pulumi.set(self, "pager_duty_configuration", value)


@pulumi.input_type
class ResponsePlanNotificationTargetItemArgs:
    def __init__(__self__, *,
                 sns_topic_arn: Optional[pulumi.Input[str]] = None):
        """
        A notification target.
        """
        if sns_topic_arn is not None:
            pulumi.set(__self__, "sns_topic_arn", sns_topic_arn)

    @property
    @pulumi.getter(name="snsTopicArn")
    def sns_topic_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "sns_topic_arn")

    @sns_topic_arn.setter
    def sns_topic_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sns_topic_arn", value)


@pulumi.input_type
class ResponsePlanPagerDutyConfigurationArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 pager_duty_incident_configuration: pulumi.Input['ResponsePlanPagerDutyIncidentConfigurationArgs'],
                 secret_id: pulumi.Input[str]):
        """
        The pagerDuty configuration to use when starting the incident.
        :param pulumi.Input[str] name: The name of the pagerDuty configuration.
        :param pulumi.Input[str] secret_id: The AWS secrets manager secretId storing the pagerDuty token.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "pager_duty_incident_configuration", pager_duty_incident_configuration)
        pulumi.set(__self__, "secret_id", secret_id)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the pagerDuty configuration.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="pagerDutyIncidentConfiguration")
    def pager_duty_incident_configuration(self) -> pulumi.Input['ResponsePlanPagerDutyIncidentConfigurationArgs']:
        return pulumi.get(self, "pager_duty_incident_configuration")

    @pager_duty_incident_configuration.setter
    def pager_duty_incident_configuration(self, value: pulumi.Input['ResponsePlanPagerDutyIncidentConfigurationArgs']):
        pulumi.set(self, "pager_duty_incident_configuration", value)

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> pulumi.Input[str]:
        """
        The AWS secrets manager secretId storing the pagerDuty token.
        """
        return pulumi.get(self, "secret_id")

    @secret_id.setter
    def secret_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "secret_id", value)


@pulumi.input_type
class ResponsePlanPagerDutyIncidentConfigurationArgs:
    def __init__(__self__, *,
                 service_id: pulumi.Input[str]):
        """
        The pagerDuty incident configuration.
        :param pulumi.Input[str] service_id: The pagerDuty serviceId.
        """
        pulumi.set(__self__, "service_id", service_id)

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> pulumi.Input[str]:
        """
        The pagerDuty serviceId.
        """
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_id", value)


@pulumi.input_type
class ResponsePlanSsmAutomationArgs:
    def __init__(__self__, *,
                 document_name: pulumi.Input[str],
                 role_arn: pulumi.Input[str],
                 document_version: Optional[pulumi.Input[str]] = None,
                 dynamic_parameters: Optional[pulumi.Input[Sequence[pulumi.Input['ResponsePlanDynamicSsmParameterArgs']]]] = None,
                 parameters: Optional[pulumi.Input[Sequence[pulumi.Input['ResponsePlanSsmParameterArgs']]]] = None,
                 target_account: Optional[pulumi.Input['ResponsePlanSsmAutomationTargetAccount']] = None):
        """
        The configuration to use when starting the SSM automation document.
        :param pulumi.Input[str] document_name: The document name to use when starting the SSM automation document.
        :param pulumi.Input[str] role_arn: The role ARN to use when starting the SSM automation document.
        :param pulumi.Input[str] document_version: The version of the document to use when starting the SSM automation document.
        :param pulumi.Input[Sequence[pulumi.Input['ResponsePlanDynamicSsmParameterArgs']]] dynamic_parameters: The parameters with dynamic values to set when starting the SSM automation document.
        :param pulumi.Input[Sequence[pulumi.Input['ResponsePlanSsmParameterArgs']]] parameters: The parameters to set when starting the SSM automation document.
        :param pulumi.Input['ResponsePlanSsmAutomationTargetAccount'] target_account: The account type to use when starting the SSM automation document.
        """
        pulumi.set(__self__, "document_name", document_name)
        pulumi.set(__self__, "role_arn", role_arn)
        if document_version is not None:
            pulumi.set(__self__, "document_version", document_version)
        if dynamic_parameters is not None:
            pulumi.set(__self__, "dynamic_parameters", dynamic_parameters)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if target_account is not None:
            pulumi.set(__self__, "target_account", target_account)

    @property
    @pulumi.getter(name="documentName")
    def document_name(self) -> pulumi.Input[str]:
        """
        The document name to use when starting the SSM automation document.
        """
        return pulumi.get(self, "document_name")

    @document_name.setter
    def document_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "document_name", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Input[str]:
        """
        The role ARN to use when starting the SSM automation document.
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter(name="documentVersion")
    def document_version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of the document to use when starting the SSM automation document.
        """
        return pulumi.get(self, "document_version")

    @document_version.setter
    def document_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "document_version", value)

    @property
    @pulumi.getter(name="dynamicParameters")
    def dynamic_parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ResponsePlanDynamicSsmParameterArgs']]]]:
        """
        The parameters with dynamic values to set when starting the SSM automation document.
        """
        return pulumi.get(self, "dynamic_parameters")

    @dynamic_parameters.setter
    def dynamic_parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ResponsePlanDynamicSsmParameterArgs']]]]):
        pulumi.set(self, "dynamic_parameters", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ResponsePlanSsmParameterArgs']]]]:
        """
        The parameters to set when starting the SSM automation document.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ResponsePlanSsmParameterArgs']]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter(name="targetAccount")
    def target_account(self) -> Optional[pulumi.Input['ResponsePlanSsmAutomationTargetAccount']]:
        """
        The account type to use when starting the SSM automation document.
        """
        return pulumi.get(self, "target_account")

    @target_account.setter
    def target_account(self, value: Optional[pulumi.Input['ResponsePlanSsmAutomationTargetAccount']]):
        pulumi.set(self, "target_account", value)


@pulumi.input_type
class ResponsePlanSsmParameterArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 values: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        A parameter to set when starting the SSM automation document.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def values(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "values", value)


@pulumi.input_type
class ResponsePlanTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        A key-value pair to tag a resource.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


