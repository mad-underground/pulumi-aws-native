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
    'GetApplicationResult',
    'AwaitableGetApplicationResult',
    'get_application',
    'get_application_output',
]

@pulumi.output_type
class GetApplicationResult:
    def __init__(__self__, application_arn=None, auto_configuration_enabled=None, c_we_monitor_enabled=None, component_monitoring_settings=None, custom_components=None, grouping_type=None, log_pattern_sets=None, ops_center_enabled=None, ops_item_sns_topic_arn=None, tags=None):
        if application_arn and not isinstance(application_arn, str):
            raise TypeError("Expected argument 'application_arn' to be a str")
        pulumi.set(__self__, "application_arn", application_arn)
        if auto_configuration_enabled and not isinstance(auto_configuration_enabled, bool):
            raise TypeError("Expected argument 'auto_configuration_enabled' to be a bool")
        pulumi.set(__self__, "auto_configuration_enabled", auto_configuration_enabled)
        if c_we_monitor_enabled and not isinstance(c_we_monitor_enabled, bool):
            raise TypeError("Expected argument 'c_we_monitor_enabled' to be a bool")
        pulumi.set(__self__, "c_we_monitor_enabled", c_we_monitor_enabled)
        if component_monitoring_settings and not isinstance(component_monitoring_settings, list):
            raise TypeError("Expected argument 'component_monitoring_settings' to be a list")
        pulumi.set(__self__, "component_monitoring_settings", component_monitoring_settings)
        if custom_components and not isinstance(custom_components, list):
            raise TypeError("Expected argument 'custom_components' to be a list")
        pulumi.set(__self__, "custom_components", custom_components)
        if grouping_type and not isinstance(grouping_type, str):
            raise TypeError("Expected argument 'grouping_type' to be a str")
        pulumi.set(__self__, "grouping_type", grouping_type)
        if log_pattern_sets and not isinstance(log_pattern_sets, list):
            raise TypeError("Expected argument 'log_pattern_sets' to be a list")
        pulumi.set(__self__, "log_pattern_sets", log_pattern_sets)
        if ops_center_enabled and not isinstance(ops_center_enabled, bool):
            raise TypeError("Expected argument 'ops_center_enabled' to be a bool")
        pulumi.set(__self__, "ops_center_enabled", ops_center_enabled)
        if ops_item_sns_topic_arn and not isinstance(ops_item_sns_topic_arn, str):
            raise TypeError("Expected argument 'ops_item_sns_topic_arn' to be a str")
        pulumi.set(__self__, "ops_item_sns_topic_arn", ops_item_sns_topic_arn)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="applicationARN")
    def application_arn(self) -> Optional[str]:
        """
        The ARN of the ApplicationInsights application.
        """
        return pulumi.get(self, "application_arn")

    @property
    @pulumi.getter(name="autoConfigurationEnabled")
    def auto_configuration_enabled(self) -> Optional[bool]:
        """
        If set to true, application will be configured with recommended monitoring configuration.
        """
        return pulumi.get(self, "auto_configuration_enabled")

    @property
    @pulumi.getter(name="cWEMonitorEnabled")
    def c_we_monitor_enabled(self) -> Optional[bool]:
        """
        Indicates whether Application Insights can listen to CloudWatch events for the application resources.
        """
        return pulumi.get(self, "c_we_monitor_enabled")

    @property
    @pulumi.getter(name="componentMonitoringSettings")
    def component_monitoring_settings(self) -> Optional[Sequence['outputs.ApplicationComponentMonitoringSetting']]:
        """
        The monitoring settings of the components.
        """
        return pulumi.get(self, "component_monitoring_settings")

    @property
    @pulumi.getter(name="customComponents")
    def custom_components(self) -> Optional[Sequence['outputs.ApplicationCustomComponent']]:
        """
        The custom grouped components.
        """
        return pulumi.get(self, "custom_components")

    @property
    @pulumi.getter(name="groupingType")
    def grouping_type(self) -> Optional['ApplicationGroupingType']:
        """
        The grouping type of the application
        """
        return pulumi.get(self, "grouping_type")

    @property
    @pulumi.getter(name="logPatternSets")
    def log_pattern_sets(self) -> Optional[Sequence['outputs.ApplicationLogPatternSet']]:
        """
        The log pattern sets.
        """
        return pulumi.get(self, "log_pattern_sets")

    @property
    @pulumi.getter(name="opsCenterEnabled")
    def ops_center_enabled(self) -> Optional[bool]:
        """
        When set to true, creates opsItems for any problems detected on an application.
        """
        return pulumi.get(self, "ops_center_enabled")

    @property
    @pulumi.getter(name="opsItemSNSTopicArn")
    def ops_item_sns_topic_arn(self) -> Optional[str]:
        """
        The SNS topic provided to Application Insights that is associated to the created opsItem.
        """
        return pulumi.get(self, "ops_item_sns_topic_arn")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.ApplicationTag']]:
        """
        The tags of Application Insights application.
        """
        return pulumi.get(self, "tags")


class AwaitableGetApplicationResult(GetApplicationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApplicationResult(
            application_arn=self.application_arn,
            auto_configuration_enabled=self.auto_configuration_enabled,
            c_we_monitor_enabled=self.c_we_monitor_enabled,
            component_monitoring_settings=self.component_monitoring_settings,
            custom_components=self.custom_components,
            grouping_type=self.grouping_type,
            log_pattern_sets=self.log_pattern_sets,
            ops_center_enabled=self.ops_center_enabled,
            ops_item_sns_topic_arn=self.ops_item_sns_topic_arn,
            tags=self.tags)


def get_application(application_arn: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApplicationResult:
    """
    Resource schema for AWS::ApplicationInsights::Application


    :param str application_arn: The ARN of the ApplicationInsights application.
    """
    __args__ = dict()
    __args__['applicationARN'] = application_arn
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:applicationinsights:getApplication', __args__, opts=opts, typ=GetApplicationResult).value

    return AwaitableGetApplicationResult(
        application_arn=__ret__.application_arn,
        auto_configuration_enabled=__ret__.auto_configuration_enabled,
        c_we_monitor_enabled=__ret__.c_we_monitor_enabled,
        component_monitoring_settings=__ret__.component_monitoring_settings,
        custom_components=__ret__.custom_components,
        grouping_type=__ret__.grouping_type,
        log_pattern_sets=__ret__.log_pattern_sets,
        ops_center_enabled=__ret__.ops_center_enabled,
        ops_item_sns_topic_arn=__ret__.ops_item_sns_topic_arn,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_application)
def get_application_output(application_arn: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetApplicationResult]:
    """
    Resource schema for AWS::ApplicationInsights::Application


    :param str application_arn: The ARN of the ApplicationInsights application.
    """
    ...
