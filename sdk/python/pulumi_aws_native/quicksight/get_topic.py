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
    'GetTopicResult',
    'AwaitableGetTopicResult',
    'get_topic',
    'get_topic_output',
]

@pulumi.output_type
class GetTopicResult:
    def __init__(__self__, arn=None, data_sets=None, description=None, name=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if data_sets and not isinstance(data_sets, list):
            raise TypeError("Expected argument 'data_sets' to be a list")
        pulumi.set(__self__, "data_sets", data_sets)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="dataSets")
    def data_sets(self) -> Optional[Sequence['outputs.TopicDatasetMetadata']]:
        return pulumi.get(self, "data_sets")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")


class AwaitableGetTopicResult(GetTopicResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTopicResult(
            arn=self.arn,
            data_sets=self.data_sets,
            description=self.description,
            name=self.name)


def get_topic(aws_account_id: Optional[str] = None,
              topic_id: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTopicResult:
    """
    Definition of the AWS::QuickSight::Topic Resource Type.
    """
    __args__ = dict()
    __args__['awsAccountId'] = aws_account_id
    __args__['topicId'] = topic_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:quicksight:getTopic', __args__, opts=opts, typ=GetTopicResult).value

    return AwaitableGetTopicResult(
        arn=__ret__.arn,
        data_sets=__ret__.data_sets,
        description=__ret__.description,
        name=__ret__.name)


@_utilities.lift_output_func(get_topic)
def get_topic_output(aws_account_id: Optional[pulumi.Input[str]] = None,
                     topic_id: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTopicResult]:
    """
    Definition of the AWS::QuickSight::Topic Resource Type.
    """
    ...