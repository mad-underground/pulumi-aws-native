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
    'AppMonitorConfiguration',
    'AppMonitorCustomEvents',
    'AppMonitorMetricDefinition',
    'AppMonitorMetricDestination',
]

@pulumi.output_type
class AppMonitorConfiguration(dict):
    """
    AppMonitor configuration
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "allowCookies":
            suggest = "allow_cookies"
        elif key == "enableXRay":
            suggest = "enable_x_ray"
        elif key == "excludedPages":
            suggest = "excluded_pages"
        elif key == "favoritePages":
            suggest = "favorite_pages"
        elif key == "guestRoleArn":
            suggest = "guest_role_arn"
        elif key == "identityPoolId":
            suggest = "identity_pool_id"
        elif key == "includedPages":
            suggest = "included_pages"
        elif key == "metricDestinations":
            suggest = "metric_destinations"
        elif key == "sessionSampleRate":
            suggest = "session_sample_rate"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AppMonitorConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AppMonitorConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AppMonitorConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 allow_cookies: Optional[bool] = None,
                 enable_x_ray: Optional[bool] = None,
                 excluded_pages: Optional[Sequence[str]] = None,
                 favorite_pages: Optional[Sequence[str]] = None,
                 guest_role_arn: Optional[str] = None,
                 identity_pool_id: Optional[str] = None,
                 included_pages: Optional[Sequence[str]] = None,
                 metric_destinations: Optional[Sequence['outputs.AppMonitorMetricDestination']] = None,
                 session_sample_rate: Optional[float] = None,
                 telemetries: Optional[Sequence['AppMonitorTelemetry']] = None):
        """
        AppMonitor configuration
        :param bool allow_cookies: If you set this to true, the RUM web client sets two cookies, a session cookie and a user cookie. The cookies allow the RUM web client to collect data relating to the number of users an application has and the behavior of the application across a sequence of events. Cookies are stored in the top-level domain of the current page.
        :param bool enable_x_ray: If you set this to true, RUM enables xray tracing for the user sessions that RUM samples. RUM adds an xray trace header to allowed HTTP requests. It also records an xray segment for allowed HTTP requests. You can see traces and segments from these user sessions in the xray console and the CW ServiceLens console.
        :param Sequence[str] excluded_pages: A list of URLs in your website or application to exclude from RUM data collection. You can't include both ExcludedPages and IncludedPages in the same operation.
        :param Sequence[str] favorite_pages: A list of pages in the RUM console that are to be displayed with a favorite icon.
        :param str guest_role_arn: The ARN of the guest IAM role that is attached to the identity pool that is used to authorize the sending of data to RUM.
        :param str identity_pool_id: The ID of the identity pool that is used to authorize the sending of data to RUM.
        :param Sequence[str] included_pages: If this app monitor is to collect data from only certain pages in your application, this structure lists those pages. You can't include both ExcludedPages and IncludedPages in the same operation.
        :param Sequence['AppMonitorMetricDestination'] metric_destinations: An array of structures which define the destinations and the metrics that you want to send.
        :param float session_sample_rate: Specifies the percentage of user sessions to use for RUM data collection. Choosing a higher percentage gives you more data but also incurs more costs. The number you specify is the percentage of user sessions that will be used. If you omit this parameter, the default of 10 is used.
        :param Sequence['AppMonitorTelemetry'] telemetries: An array that lists the types of telemetry data that this app monitor is to collect.
        """
        if allow_cookies is not None:
            pulumi.set(__self__, "allow_cookies", allow_cookies)
        if enable_x_ray is not None:
            pulumi.set(__self__, "enable_x_ray", enable_x_ray)
        if excluded_pages is not None:
            pulumi.set(__self__, "excluded_pages", excluded_pages)
        if favorite_pages is not None:
            pulumi.set(__self__, "favorite_pages", favorite_pages)
        if guest_role_arn is not None:
            pulumi.set(__self__, "guest_role_arn", guest_role_arn)
        if identity_pool_id is not None:
            pulumi.set(__self__, "identity_pool_id", identity_pool_id)
        if included_pages is not None:
            pulumi.set(__self__, "included_pages", included_pages)
        if metric_destinations is not None:
            pulumi.set(__self__, "metric_destinations", metric_destinations)
        if session_sample_rate is not None:
            pulumi.set(__self__, "session_sample_rate", session_sample_rate)
        if telemetries is not None:
            pulumi.set(__self__, "telemetries", telemetries)

    @property
    @pulumi.getter(name="allowCookies")
    def allow_cookies(self) -> Optional[bool]:
        """
        If you set this to true, the RUM web client sets two cookies, a session cookie and a user cookie. The cookies allow the RUM web client to collect data relating to the number of users an application has and the behavior of the application across a sequence of events. Cookies are stored in the top-level domain of the current page.
        """
        return pulumi.get(self, "allow_cookies")

    @property
    @pulumi.getter(name="enableXRay")
    def enable_x_ray(self) -> Optional[bool]:
        """
        If you set this to true, RUM enables xray tracing for the user sessions that RUM samples. RUM adds an xray trace header to allowed HTTP requests. It also records an xray segment for allowed HTTP requests. You can see traces and segments from these user sessions in the xray console and the CW ServiceLens console.
        """
        return pulumi.get(self, "enable_x_ray")

    @property
    @pulumi.getter(name="excludedPages")
    def excluded_pages(self) -> Optional[Sequence[str]]:
        """
        A list of URLs in your website or application to exclude from RUM data collection. You can't include both ExcludedPages and IncludedPages in the same operation.
        """
        return pulumi.get(self, "excluded_pages")

    @property
    @pulumi.getter(name="favoritePages")
    def favorite_pages(self) -> Optional[Sequence[str]]:
        """
        A list of pages in the RUM console that are to be displayed with a favorite icon.
        """
        return pulumi.get(self, "favorite_pages")

    @property
    @pulumi.getter(name="guestRoleArn")
    def guest_role_arn(self) -> Optional[str]:
        """
        The ARN of the guest IAM role that is attached to the identity pool that is used to authorize the sending of data to RUM.
        """
        return pulumi.get(self, "guest_role_arn")

    @property
    @pulumi.getter(name="identityPoolId")
    def identity_pool_id(self) -> Optional[str]:
        """
        The ID of the identity pool that is used to authorize the sending of data to RUM.
        """
        return pulumi.get(self, "identity_pool_id")

    @property
    @pulumi.getter(name="includedPages")
    def included_pages(self) -> Optional[Sequence[str]]:
        """
        If this app monitor is to collect data from only certain pages in your application, this structure lists those pages. You can't include both ExcludedPages and IncludedPages in the same operation.
        """
        return pulumi.get(self, "included_pages")

    @property
    @pulumi.getter(name="metricDestinations")
    def metric_destinations(self) -> Optional[Sequence['outputs.AppMonitorMetricDestination']]:
        """
        An array of structures which define the destinations and the metrics that you want to send.
        """
        return pulumi.get(self, "metric_destinations")

    @property
    @pulumi.getter(name="sessionSampleRate")
    def session_sample_rate(self) -> Optional[float]:
        """
        Specifies the percentage of user sessions to use for RUM data collection. Choosing a higher percentage gives you more data but also incurs more costs. The number you specify is the percentage of user sessions that will be used. If you omit this parameter, the default of 10 is used.
        """
        return pulumi.get(self, "session_sample_rate")

    @property
    @pulumi.getter
    def telemetries(self) -> Optional[Sequence['AppMonitorTelemetry']]:
        """
        An array that lists the types of telemetry data that this app monitor is to collect.
        """
        return pulumi.get(self, "telemetries")


@pulumi.output_type
class AppMonitorCustomEvents(dict):
    """
    AppMonitor custom events configuration
    """
    def __init__(__self__, *,
                 status: Optional['AppMonitorCustomEventsStatus'] = None):
        """
        AppMonitor custom events configuration
        :param 'AppMonitorCustomEventsStatus' status: Indicates whether AppMonitor accepts custom events.
        """
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def status(self) -> Optional['AppMonitorCustomEventsStatus']:
        """
        Indicates whether AppMonitor accepts custom events.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class AppMonitorMetricDefinition(dict):
    """
    A single metric definition
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "dimensionKeys":
            suggest = "dimension_keys"
        elif key == "eventPattern":
            suggest = "event_pattern"
        elif key == "unitLabel":
            suggest = "unit_label"
        elif key == "valueKey":
            suggest = "value_key"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AppMonitorMetricDefinition. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AppMonitorMetricDefinition.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AppMonitorMetricDefinition.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 name: str,
                 dimension_keys: Optional[Mapping[str, str]] = None,
                 event_pattern: Optional[str] = None,
                 namespace: Optional[str] = None,
                 unit_label: Optional[str] = None,
                 value_key: Optional[str] = None):
        """
        A single metric definition
        :param str name: The name for the metric that is defined in this structure. For extended metrics, valid values are the following:
               
               PerformanceNavigationDuration
               
               PerformanceResourceDuration
               
               NavigationSatisfiedTransaction
               
               NavigationToleratedTransaction
               
               NavigationFrustratedTransaction
               
               WebVitalsCumulativeLayoutShift
               
               WebVitalsFirstInputDelay
               
               WebVitalsLargestContentfulPaint
               
               JsErrorCount
               
               HttpErrorCount
               
               SessionCount
        :param Mapping[str, str] dimension_keys: Use this field only if you are sending the metric to CloudWatch.
               
               This field is a map of field paths to dimension names. It defines the dimensions to associate with this metric in CloudWatch. For extended metrics, valid values for the entries in this field are the following:
               
               "metadata.pageId": "PageId"
               
               "metadata.browserName": "BrowserName"
               
               "metadata.deviceType": "DeviceType"
               
               "metadata.osName": "OSName"
               
               "metadata.countryCode": "CountryCode"
               
               "event_details.fileType": "FileType"
               
               All dimensions listed in this field must also be included in EventPattern.
        :param str event_pattern: The pattern that defines the metric, specified as a JSON object. RUM checks events that happen in a user's session against the pattern, and events that match the pattern are sent to the metric destination.
               
               When you define extended metrics, the metric definition is not valid if EventPattern is omitted.
               
               Example event patterns:
               
               '{ "event_type": ["com.amazon.rum.js_error_event"], "metadata": { "browserName": [ "Chrome", "Safari" ], } }'
               
               '{ "event_type": ["com.amazon.rum.performance_navigation_event"], "metadata": { "browserName": [ "Chrome", "Firefox" ] }, "event_details": { "duration": [{ "numeric": [ "<", 2000 ] }] } }'
               
               '{ "event_type": ["com.amazon.rum.performance_navigation_event"], "metadata": { "browserName": [ "Chrome", "Safari" ], "countryCode": [ "US" ] }, "event_details": { "duration": [{ "numeric": [ ">=", 2000, "<", 8000 ] }] } }'
               
               If the metrics destination' is CloudWatch and the event also matches a value in DimensionKeys, then the metric is published with the specified dimensions.
        :param str namespace: The namespace used by CloudWatch Metrics for the metric that is defined in this structure
        :param str unit_label: The CloudWatch metric unit to use for this metric. If you omit this field, the metric is recorded with no unit.
        :param str value_key: The field within the event object that the metric value is sourced from.
               
               If you omit this field, a hardcoded value of 1 is pushed as the metric value. This is useful if you just want to count the number of events that the filter catches.
               
               If this metric is sent to Evidently, this field will be passed to Evidently raw and Evidently will handle data extraction from the event.
        """
        pulumi.set(__self__, "name", name)
        if dimension_keys is not None:
            pulumi.set(__self__, "dimension_keys", dimension_keys)
        if event_pattern is not None:
            pulumi.set(__self__, "event_pattern", event_pattern)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if unit_label is not None:
            pulumi.set(__self__, "unit_label", unit_label)
        if value_key is not None:
            pulumi.set(__self__, "value_key", value_key)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name for the metric that is defined in this structure. For extended metrics, valid values are the following:

        PerformanceNavigationDuration

        PerformanceResourceDuration

        NavigationSatisfiedTransaction

        NavigationToleratedTransaction

        NavigationFrustratedTransaction

        WebVitalsCumulativeLayoutShift

        WebVitalsFirstInputDelay

        WebVitalsLargestContentfulPaint

        JsErrorCount

        HttpErrorCount

        SessionCount
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="dimensionKeys")
    def dimension_keys(self) -> Optional[Mapping[str, str]]:
        """
        Use this field only if you are sending the metric to CloudWatch.

        This field is a map of field paths to dimension names. It defines the dimensions to associate with this metric in CloudWatch. For extended metrics, valid values for the entries in this field are the following:

        "metadata.pageId": "PageId"

        "metadata.browserName": "BrowserName"

        "metadata.deviceType": "DeviceType"

        "metadata.osName": "OSName"

        "metadata.countryCode": "CountryCode"

        "event_details.fileType": "FileType"

        All dimensions listed in this field must also be included in EventPattern.
        """
        return pulumi.get(self, "dimension_keys")

    @property
    @pulumi.getter(name="eventPattern")
    def event_pattern(self) -> Optional[str]:
        """
        The pattern that defines the metric, specified as a JSON object. RUM checks events that happen in a user's session against the pattern, and events that match the pattern are sent to the metric destination.

        When you define extended metrics, the metric definition is not valid if EventPattern is omitted.

        Example event patterns:

        '{ "event_type": ["com.amazon.rum.js_error_event"], "metadata": { "browserName": [ "Chrome", "Safari" ], } }'

        '{ "event_type": ["com.amazon.rum.performance_navigation_event"], "metadata": { "browserName": [ "Chrome", "Firefox" ] }, "event_details": { "duration": [{ "numeric": [ "<", 2000 ] }] } }'

        '{ "event_type": ["com.amazon.rum.performance_navigation_event"], "metadata": { "browserName": [ "Chrome", "Safari" ], "countryCode": [ "US" ] }, "event_details": { "duration": [{ "numeric": [ ">=", 2000, "<", 8000 ] }] } }'

        If the metrics destination' is CloudWatch and the event also matches a value in DimensionKeys, then the metric is published with the specified dimensions.
        """
        return pulumi.get(self, "event_pattern")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[str]:
        """
        The namespace used by CloudWatch Metrics for the metric that is defined in this structure
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="unitLabel")
    def unit_label(self) -> Optional[str]:
        """
        The CloudWatch metric unit to use for this metric. If you omit this field, the metric is recorded with no unit.
        """
        return pulumi.get(self, "unit_label")

    @property
    @pulumi.getter(name="valueKey")
    def value_key(self) -> Optional[str]:
        """
        The field within the event object that the metric value is sourced from.

        If you omit this field, a hardcoded value of 1 is pushed as the metric value. This is useful if you just want to count the number of events that the filter catches.

        If this metric is sent to Evidently, this field will be passed to Evidently raw and Evidently will handle data extraction from the event.
        """
        return pulumi.get(self, "value_key")


@pulumi.output_type
class AppMonitorMetricDestination(dict):
    """
    An structure which defines the destination and the metrics that you want to send.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "destinationArn":
            suggest = "destination_arn"
        elif key == "iamRoleArn":
            suggest = "iam_role_arn"
        elif key == "metricDefinitions":
            suggest = "metric_definitions"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AppMonitorMetricDestination. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AppMonitorMetricDestination.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AppMonitorMetricDestination.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 destination: 'AppMonitorMetricDestinationDestination',
                 destination_arn: Optional[str] = None,
                 iam_role_arn: Optional[str] = None,
                 metric_definitions: Optional[Sequence['outputs.AppMonitorMetricDefinition']] = None):
        """
        An structure which defines the destination and the metrics that you want to send.
        :param 'AppMonitorMetricDestinationDestination' destination: Defines the destination to send the metrics to. Valid values are CloudWatch and Evidently. If you specify Evidently, you must also specify the ARN of the Evidently experiment that is to be the destination and an IAM role that has permission to write to the experiment.
        :param str destination_arn: Use this parameter only if Destination is Evidently. This parameter specifies the ARN of the Evidently experiment that will receive the extended metrics.
        :param str iam_role_arn: This parameter is required if Destination is Evidently. If Destination is CloudWatch, do not use this parameter.
               
               This parameter specifies the ARN of an IAM role that RUM will assume to write to the Evidently experiment that you are sending metrics to. This role must have permission to write to that experiment.
        :param Sequence['AppMonitorMetricDefinition'] metric_definitions: An array of structures which define the metrics that you want to send.
        """
        pulumi.set(__self__, "destination", destination)
        if destination_arn is not None:
            pulumi.set(__self__, "destination_arn", destination_arn)
        if iam_role_arn is not None:
            pulumi.set(__self__, "iam_role_arn", iam_role_arn)
        if metric_definitions is not None:
            pulumi.set(__self__, "metric_definitions", metric_definitions)

    @property
    @pulumi.getter
    def destination(self) -> 'AppMonitorMetricDestinationDestination':
        """
        Defines the destination to send the metrics to. Valid values are CloudWatch and Evidently. If you specify Evidently, you must also specify the ARN of the Evidently experiment that is to be the destination and an IAM role that has permission to write to the experiment.
        """
        return pulumi.get(self, "destination")

    @property
    @pulumi.getter(name="destinationArn")
    def destination_arn(self) -> Optional[str]:
        """
        Use this parameter only if Destination is Evidently. This parameter specifies the ARN of the Evidently experiment that will receive the extended metrics.
        """
        return pulumi.get(self, "destination_arn")

    @property
    @pulumi.getter(name="iamRoleArn")
    def iam_role_arn(self) -> Optional[str]:
        """
        This parameter is required if Destination is Evidently. If Destination is CloudWatch, do not use this parameter.

        This parameter specifies the ARN of an IAM role that RUM will assume to write to the Evidently experiment that you are sending metrics to. This role must have permission to write to that experiment.
        """
        return pulumi.get(self, "iam_role_arn")

    @property
    @pulumi.getter(name="metricDefinitions")
    def metric_definitions(self) -> Optional[Sequence['outputs.AppMonitorMetricDefinition']]:
        """
        An array of structures which define the metrics that you want to send.
        """
        return pulumi.get(self, "metric_definitions")


