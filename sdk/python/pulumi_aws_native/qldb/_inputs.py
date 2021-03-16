# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'StreamKinesisConfigurationArgs',
]

@pulumi.input_type
class StreamKinesisConfigurationArgs:
    def __init__(__self__, *,
                 aggregation_enabled: Optional[pulumi.Input[bool]] = None,
                 stream_arn: Optional[pulumi.Input[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qldb-stream-kinesisconfiguration.html
        :param pulumi.Input[bool] aggregation_enabled: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qldb-stream-kinesisconfiguration.html#cfn-qldb-stream-kinesisconfiguration-aggregationenabled
        :param pulumi.Input[str] stream_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qldb-stream-kinesisconfiguration.html#cfn-qldb-stream-kinesisconfiguration-streamarn
        """
        if aggregation_enabled is not None:
            pulumi.set(__self__, "aggregation_enabled", aggregation_enabled)
        if stream_arn is not None:
            pulumi.set(__self__, "stream_arn", stream_arn)

    @property
    @pulumi.getter(name="aggregationEnabled")
    def aggregation_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qldb-stream-kinesisconfiguration.html#cfn-qldb-stream-kinesisconfiguration-aggregationenabled
        """
        return pulumi.get(self, "aggregation_enabled")

    @aggregation_enabled.setter
    def aggregation_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "aggregation_enabled", value)

    @property
    @pulumi.getter(name="streamArn")
    def stream_arn(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qldb-stream-kinesisconfiguration.html#cfn-qldb-stream-kinesisconfiguration-streamarn
        """
        return pulumi.get(self, "stream_arn")

    @stream_arn.setter
    def stream_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stream_arn", value)

