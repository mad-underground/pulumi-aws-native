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
    'GetEventSubscriptionResult',
    'AwaitableGetEventSubscriptionResult',
    'get_event_subscription',
    'get_event_subscription_output',
]

@pulumi.output_type
class GetEventSubscriptionResult:
    def __init__(__self__, enabled=None, event_categories=None, source_ids=None, source_type=None, tags=None):
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if event_categories and not isinstance(event_categories, list):
            raise TypeError("Expected argument 'event_categories' to be a list")
        pulumi.set(__self__, "event_categories", event_categories)
        if source_ids and not isinstance(source_ids, list):
            raise TypeError("Expected argument 'source_ids' to be a list")
        pulumi.set(__self__, "source_ids", source_ids)
        if source_type and not isinstance(source_type, str):
            raise TypeError("Expected argument 'source_type' to be a str")
        pulumi.set(__self__, "source_type", source_type)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        A Boolean value; set to true to activate the subscription, set to false to create the subscription but not active it.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="eventCategories")
    def event_categories(self) -> Optional[Sequence[str]]:
        """
        A list of event categories for a SourceType that you want to subscribe to. You can see a list of the categories for a given SourceType in the Events topic in the Amazon RDS User Guide or by using the DescribeEventCategories action.
        """
        return pulumi.get(self, "event_categories")

    @property
    @pulumi.getter(name="sourceIds")
    def source_ids(self) -> Optional[Sequence[str]]:
        """
        The list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it cannot end with a hyphen or contain two consecutive hyphens.
        """
        return pulumi.get(self, "source_ids")

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> Optional['EventSubscriptionSourceType']:
        """
        The type of source that will be generating the events. For example, if you want to be notified of events generated by a DB instance, you would set this parameter to db-instance. if this value is not specified, all events are returned.
        """
        return pulumi.get(self, "source_type")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.EventSubscriptionTag']]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")


class AwaitableGetEventSubscriptionResult(GetEventSubscriptionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEventSubscriptionResult(
            enabled=self.enabled,
            event_categories=self.event_categories,
            source_ids=self.source_ids,
            source_type=self.source_type,
            tags=self.tags)


def get_event_subscription(subscription_name: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEventSubscriptionResult:
    """
    The AWS::RDS::EventSubscription resource allows you to receive notifications for Amazon Relational Database Service events through the Amazon Simple Notification Service (Amazon SNS). For more information, see Using Amazon RDS Event Notification in the Amazon RDS User Guide.


    :param str subscription_name: The name of the subscription.
    """
    __args__ = dict()
    __args__['subscriptionName'] = subscription_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:rds:getEventSubscription', __args__, opts=opts, typ=GetEventSubscriptionResult).value

    return AwaitableGetEventSubscriptionResult(
        enabled=__ret__.enabled,
        event_categories=__ret__.event_categories,
        source_ids=__ret__.source_ids,
        source_type=__ret__.source_type,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_event_subscription)
def get_event_subscription_output(subscription_name: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEventSubscriptionResult]:
    """
    The AWS::RDS::EventSubscription resource allows you to receive notifications for Amazon Relational Database Service events through the Amazon Simple Notification Service (Amazon SNS). For more information, see Using Amazon RDS Event Notification in the Amazon RDS User Guide.


    :param str subscription_name: The name of the subscription.
    """
    ...
