# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetTopicResult',
    'AwaitableGetTopicResult',
    'get_topic',
    'get_topic_output',
]

@pulumi.output_type
class GetTopicResult:
    def __init__(__self__, content_based_deduplication=None, display_name=None, id=None, kms_master_key_id=None, subscription=None, tags=None):
        if content_based_deduplication and not isinstance(content_based_deduplication, bool):
            raise TypeError("Expected argument 'content_based_deduplication' to be a bool")
        pulumi.set(__self__, "content_based_deduplication", content_based_deduplication)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if kms_master_key_id and not isinstance(kms_master_key_id, str):
            raise TypeError("Expected argument 'kms_master_key_id' to be a str")
        pulumi.set(__self__, "kms_master_key_id", kms_master_key_id)
        if subscription and not isinstance(subscription, list):
            raise TypeError("Expected argument 'subscription' to be a list")
        pulumi.set(__self__, "subscription", subscription)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="contentBasedDeduplication")
    def content_based_deduplication(self) -> Optional[bool]:
        return pulumi.get(self, "content_based_deduplication")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="kmsMasterKeyId")
    def kms_master_key_id(self) -> Optional[str]:
        return pulumi.get(self, "kms_master_key_id")

    @property
    @pulumi.getter
    def subscription(self) -> Optional[Sequence['outputs.TopicSubscription']]:
        return pulumi.get(self, "subscription")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.TopicTag']]:
        return pulumi.get(self, "tags")


class AwaitableGetTopicResult(GetTopicResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTopicResult(
            content_based_deduplication=self.content_based_deduplication,
            display_name=self.display_name,
            id=self.id,
            kms_master_key_id=self.kms_master_key_id,
            subscription=self.subscription,
            tags=self.tags)


def get_topic(id: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTopicResult:
    """
    Resource Type definition for AWS::SNS::Topic
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:sns:getTopic', __args__, opts=opts, typ=GetTopicResult).value

    return AwaitableGetTopicResult(
        content_based_deduplication=__ret__.content_based_deduplication,
        display_name=__ret__.display_name,
        id=__ret__.id,
        kms_master_key_id=__ret__.kms_master_key_id,
        subscription=__ret__.subscription,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_topic)
def get_topic_output(id: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTopicResult]:
    """
    Resource Type definition for AWS::SNS::Topic
    """
    ...