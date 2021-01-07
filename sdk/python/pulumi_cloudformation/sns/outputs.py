# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from .. import outputs as _root_outputs

__all__ = [
    'SubscriptionAttributes',
    'SubscriptionProperties',
    'TopicAttributes',
    'TopicPolicyAttributes',
    'TopicPolicyProperties',
    'TopicProperties',
    'TopicSubscription',
]

@pulumi.output_type
class SubscriptionAttributes(dict):
    def __init__(__self__):
        pass

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SubscriptionProperties(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html
    """
    def __init__(__self__, *,
                 protocol: str,
                 topic_arn: str,
                 delivery_policy: Optional[str] = None,
                 endpoint: Optional[str] = None,
                 filter_policy: Optional[str] = None,
                 raw_message_delivery: Optional[bool] = None,
                 redrive_policy: Optional[str] = None,
                 region: Optional[str] = None,
                 subscription_role_arn: Optional[str] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html
        :param str protocol: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-protocol
        :param str topic_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#topicarn
        :param Union[Any, str] delivery_policy: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-subscription-deliverypolicy
        :param str endpoint: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-endpoint
        :param Union[Any, str] filter_policy: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-subscription-filterpolicy
        :param bool raw_message_delivery: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-subscription-rawmessagedelivery
        :param Union[Any, str] redrive_policy: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-subscription-redrivepolicy
        :param str region: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-subscription-region
        :param str subscription_role_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-subscription-subscriptionrolearn
        """
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "topic_arn", topic_arn)
        if delivery_policy is not None:
            pulumi.set(__self__, "delivery_policy", delivery_policy)
        if endpoint is not None:
            pulumi.set(__self__, "endpoint", endpoint)
        if filter_policy is not None:
            pulumi.set(__self__, "filter_policy", filter_policy)
        if raw_message_delivery is not None:
            pulumi.set(__self__, "raw_message_delivery", raw_message_delivery)
        if redrive_policy is not None:
            pulumi.set(__self__, "redrive_policy", redrive_policy)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if subscription_role_arn is not None:
            pulumi.set(__self__, "subscription_role_arn", subscription_role_arn)

    @property
    @pulumi.getter(name="Protocol")
    def protocol(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-protocol
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="TopicArn")
    def topic_arn(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#topicarn
        """
        return pulumi.get(self, "topic_arn")

    @property
    @pulumi.getter(name="DeliveryPolicy")
    def delivery_policy(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-subscription-deliverypolicy
        """
        return pulumi.get(self, "delivery_policy")

    @property
    @pulumi.getter(name="Endpoint")
    def endpoint(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-endpoint
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="FilterPolicy")
    def filter_policy(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-subscription-filterpolicy
        """
        return pulumi.get(self, "filter_policy")

    @property
    @pulumi.getter(name="RawMessageDelivery")
    def raw_message_delivery(self) -> Optional[bool]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-subscription-rawmessagedelivery
        """
        return pulumi.get(self, "raw_message_delivery")

    @property
    @pulumi.getter(name="RedrivePolicy")
    def redrive_policy(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-subscription-redrivepolicy
        """
        return pulumi.get(self, "redrive_policy")

    @property
    @pulumi.getter(name="Region")
    def region(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-subscription-region
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="SubscriptionRoleArn")
    def subscription_role_arn(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-subscription-subscriptionrolearn
        """
        return pulumi.get(self, "subscription_role_arn")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TopicAttributes(dict):
    def __init__(__self__, *,
                 topic_name: str):
        pulumi.set(__self__, "topic_name", topic_name)

    @property
    @pulumi.getter(name="TopicName")
    def topic_name(self) -> str:
        return pulumi.get(self, "topic_name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TopicPolicyAttributes(dict):
    def __init__(__self__):
        pass

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TopicPolicyProperties(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html
    """
    def __init__(__self__, *,
                 policy_document: str,
                 topics: Sequence[str]):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html
        :param Union[Any, str] policy_document: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html#cfn-sns-topicpolicy-policydocument
        :param Sequence[str] topics: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html#cfn-sns-topicpolicy-topics
        """
        pulumi.set(__self__, "policy_document", policy_document)
        pulumi.set(__self__, "topics", topics)

    @property
    @pulumi.getter(name="PolicyDocument")
    def policy_document(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html#cfn-sns-topicpolicy-policydocument
        """
        return pulumi.get(self, "policy_document")

    @property
    @pulumi.getter(name="Topics")
    def topics(self) -> Sequence[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html#cfn-sns-topicpolicy-topics
        """
        return pulumi.get(self, "topics")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TopicProperties(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html
    """
    def __init__(__self__, *,
                 content_based_deduplication: Optional[bool] = None,
                 display_name: Optional[str] = None,
                 fifo_topic: Optional[bool] = None,
                 kms_master_key_id: Optional[str] = None,
                 subscription: Optional[Sequence['outputs.TopicSubscription']] = None,
                 tags: Optional[Sequence['_root_outputs.Tag']] = None,
                 topic_name: Optional[str] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html
        :param bool content_based_deduplication: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html#cfn-sns-topic-contentbaseddeduplication
        :param str display_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html#cfn-sns-topic-displayname
        :param bool fifo_topic: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html#cfn-sns-topic-fifotopic
        :param str kms_master_key_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html#cfn-sns-topic-kmsmasterkeyid
        :param Sequence['TopicSubscriptionArgs'] subscription: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html#cfn-sns-topic-subscription
        :param Sequence['_root_inputs.TagArgs'] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html#cfn-sns-topic-tags
        :param str topic_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html#cfn-sns-topic-topicname
        """
        if content_based_deduplication is not None:
            pulumi.set(__self__, "content_based_deduplication", content_based_deduplication)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if fifo_topic is not None:
            pulumi.set(__self__, "fifo_topic", fifo_topic)
        if kms_master_key_id is not None:
            pulumi.set(__self__, "kms_master_key_id", kms_master_key_id)
        if subscription is not None:
            pulumi.set(__self__, "subscription", subscription)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if topic_name is not None:
            pulumi.set(__self__, "topic_name", topic_name)

    @property
    @pulumi.getter(name="ContentBasedDeduplication")
    def content_based_deduplication(self) -> Optional[bool]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html#cfn-sns-topic-contentbaseddeduplication
        """
        return pulumi.get(self, "content_based_deduplication")

    @property
    @pulumi.getter(name="DisplayName")
    def display_name(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html#cfn-sns-topic-displayname
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="FifoTopic")
    def fifo_topic(self) -> Optional[bool]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html#cfn-sns-topic-fifotopic
        """
        return pulumi.get(self, "fifo_topic")

    @property
    @pulumi.getter(name="KmsMasterKeyId")
    def kms_master_key_id(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html#cfn-sns-topic-kmsmasterkeyid
        """
        return pulumi.get(self, "kms_master_key_id")

    @property
    @pulumi.getter(name="Subscription")
    def subscription(self) -> Optional[Sequence['outputs.TopicSubscription']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html#cfn-sns-topic-subscription
        """
        return pulumi.get(self, "subscription")

    @property
    @pulumi.getter(name="Tags")
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html#cfn-sns-topic-tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="TopicName")
    def topic_name(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html#cfn-sns-topic-topicname
        """
        return pulumi.get(self, "topic_name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TopicSubscription(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-subscription.html
    """
    def __init__(__self__, *,
                 endpoint: str,
                 protocol: str):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-subscription.html
        :param str endpoint: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-subscription.html#cfn-sns-topic-subscription-endpoint
        :param str protocol: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-subscription.html#cfn-sns-topic-subscription-protocol
        """
        pulumi.set(__self__, "endpoint", endpoint)
        pulumi.set(__self__, "protocol", protocol)

    @property
    @pulumi.getter(name="Endpoint")
    def endpoint(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-subscription.html#cfn-sns-topic-subscription-endpoint
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="Protocol")
    def protocol(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-subscription.html#cfn-sns-topic-subscription-protocol
        """
        return pulumi.get(self, "protocol")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

