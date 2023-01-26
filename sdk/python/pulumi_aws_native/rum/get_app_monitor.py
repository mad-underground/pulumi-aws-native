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
    'GetAppMonitorResult',
    'AwaitableGetAppMonitorResult',
    'get_app_monitor',
    'get_app_monitor_output',
]

@pulumi.output_type
class GetAppMonitorResult:
    def __init__(__self__, app_monitor_configuration=None, custom_events=None, cw_log_enabled=None, domain=None, tags=None):
        if app_monitor_configuration and not isinstance(app_monitor_configuration, dict):
            raise TypeError("Expected argument 'app_monitor_configuration' to be a dict")
        pulumi.set(__self__, "app_monitor_configuration", app_monitor_configuration)
        if custom_events and not isinstance(custom_events, dict):
            raise TypeError("Expected argument 'custom_events' to be a dict")
        pulumi.set(__self__, "custom_events", custom_events)
        if cw_log_enabled and not isinstance(cw_log_enabled, bool):
            raise TypeError("Expected argument 'cw_log_enabled' to be a bool")
        pulumi.set(__self__, "cw_log_enabled", cw_log_enabled)
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        pulumi.set(__self__, "domain", domain)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="appMonitorConfiguration")
    def app_monitor_configuration(self) -> Optional['outputs.AppMonitorConfiguration']:
        return pulumi.get(self, "app_monitor_configuration")

    @property
    @pulumi.getter(name="customEvents")
    def custom_events(self) -> Optional['outputs.AppMonitorCustomEvents']:
        return pulumi.get(self, "custom_events")

    @property
    @pulumi.getter(name="cwLogEnabled")
    def cw_log_enabled(self) -> Optional[bool]:
        """
        Data collected by RUM is kept by RUM for 30 days and then deleted. This parameter specifies whether RUM sends a copy of this telemetry data to CWLlong in your account. This enables you to keep the telemetry data for more than 30 days, but it does incur CWLlong charges. If you omit this parameter, the default is false
        """
        return pulumi.get(self, "cw_log_enabled")

    @property
    @pulumi.getter
    def domain(self) -> Optional[str]:
        """
        The top-level internet domain name for which your application has administrative authority.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.AppMonitorTag']]:
        return pulumi.get(self, "tags")


class AwaitableGetAppMonitorResult(GetAppMonitorResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppMonitorResult(
            app_monitor_configuration=self.app_monitor_configuration,
            custom_events=self.custom_events,
            cw_log_enabled=self.cw_log_enabled,
            domain=self.domain,
            tags=self.tags)


def get_app_monitor(name: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppMonitorResult:
    """
    Resource Type definition for AWS::RUM::AppMonitor


    :param str name: A name for the app monitor
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:rum:getAppMonitor', __args__, opts=opts, typ=GetAppMonitorResult).value

    return AwaitableGetAppMonitorResult(
        app_monitor_configuration=__ret__.app_monitor_configuration,
        custom_events=__ret__.custom_events,
        cw_log_enabled=__ret__.cw_log_enabled,
        domain=__ret__.domain,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_app_monitor)
def get_app_monitor_output(name: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAppMonitorResult]:
    """
    Resource Type definition for AWS::RUM::AppMonitor


    :param str name: A name for the app monitor
    """
    ...
