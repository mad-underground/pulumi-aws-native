# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'StreamStreamEncryptionArgs',
]

@pulumi.input_type
class StreamStreamEncryptionArgs:
    def __init__(__self__, *,
                 encryption_type: pulumi.Input[str],
                 key_id: pulumi.Input[str]):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-stream-streamencryption.html
        :param pulumi.Input[str] encryption_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-stream-streamencryption.html#cfn-kinesis-stream-streamencryption-encryptiontype
        :param pulumi.Input[str] key_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-stream-streamencryption.html#cfn-kinesis-stream-streamencryption-keyid
        """
        pulumi.set(__self__, "encryption_type", encryption_type)
        pulumi.set(__self__, "key_id", key_id)

    @property
    @pulumi.getter(name="EncryptionType")
    def encryption_type(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-stream-streamencryption.html#cfn-kinesis-stream-streamencryption-encryptiontype
        """
        return pulumi.get(self, "encryption_type")

    @encryption_type.setter
    def encryption_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "encryption_type", value)

    @property
    @pulumi.getter(name="KeyId")
    def key_id(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-stream-streamencryption.html#cfn-kinesis-stream-streamencryption-keyid
        """
        return pulumi.get(self, "key_id")

    @key_id.setter
    def key_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_id", value)


