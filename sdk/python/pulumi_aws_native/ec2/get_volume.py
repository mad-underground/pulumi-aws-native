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
    'GetVolumeResult',
    'AwaitableGetVolumeResult',
    'get_volume',
    'get_volume_output',
]

@pulumi.output_type
class GetVolumeResult:
    def __init__(__self__, auto_enable_io=None, availability_zone=None, encrypted=None, id=None, iops=None, kms_key_id=None, multi_attach_enabled=None, outpost_arn=None, size=None, snapshot_id=None, tags=None, throughput=None, volume_type=None):
        if auto_enable_io and not isinstance(auto_enable_io, bool):
            raise TypeError("Expected argument 'auto_enable_io' to be a bool")
        pulumi.set(__self__, "auto_enable_io", auto_enable_io)
        if availability_zone and not isinstance(availability_zone, str):
            raise TypeError("Expected argument 'availability_zone' to be a str")
        pulumi.set(__self__, "availability_zone", availability_zone)
        if encrypted and not isinstance(encrypted, bool):
            raise TypeError("Expected argument 'encrypted' to be a bool")
        pulumi.set(__self__, "encrypted", encrypted)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if iops and not isinstance(iops, int):
            raise TypeError("Expected argument 'iops' to be a int")
        pulumi.set(__self__, "iops", iops)
        if kms_key_id and not isinstance(kms_key_id, str):
            raise TypeError("Expected argument 'kms_key_id' to be a str")
        pulumi.set(__self__, "kms_key_id", kms_key_id)
        if multi_attach_enabled and not isinstance(multi_attach_enabled, bool):
            raise TypeError("Expected argument 'multi_attach_enabled' to be a bool")
        pulumi.set(__self__, "multi_attach_enabled", multi_attach_enabled)
        if outpost_arn and not isinstance(outpost_arn, str):
            raise TypeError("Expected argument 'outpost_arn' to be a str")
        pulumi.set(__self__, "outpost_arn", outpost_arn)
        if size and not isinstance(size, int):
            raise TypeError("Expected argument 'size' to be a int")
        pulumi.set(__self__, "size", size)
        if snapshot_id and not isinstance(snapshot_id, str):
            raise TypeError("Expected argument 'snapshot_id' to be a str")
        pulumi.set(__self__, "snapshot_id", snapshot_id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if throughput and not isinstance(throughput, int):
            raise TypeError("Expected argument 'throughput' to be a int")
        pulumi.set(__self__, "throughput", throughput)
        if volume_type and not isinstance(volume_type, str):
            raise TypeError("Expected argument 'volume_type' to be a str")
        pulumi.set(__self__, "volume_type", volume_type)

    @property
    @pulumi.getter(name="autoEnableIO")
    def auto_enable_io(self) -> Optional[bool]:
        return pulumi.get(self, "auto_enable_io")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[str]:
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter
    def encrypted(self) -> Optional[bool]:
        return pulumi.get(self, "encrypted")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def iops(self) -> Optional[int]:
        return pulumi.get(self, "iops")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[str]:
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter(name="multiAttachEnabled")
    def multi_attach_enabled(self) -> Optional[bool]:
        return pulumi.get(self, "multi_attach_enabled")

    @property
    @pulumi.getter(name="outpostArn")
    def outpost_arn(self) -> Optional[str]:
        return pulumi.get(self, "outpost_arn")

    @property
    @pulumi.getter
    def size(self) -> Optional[int]:
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> Optional[str]:
        return pulumi.get(self, "snapshot_id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.VolumeTag']]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def throughput(self) -> Optional[int]:
        return pulumi.get(self, "throughput")

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> Optional[str]:
        return pulumi.get(self, "volume_type")


class AwaitableGetVolumeResult(GetVolumeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVolumeResult(
            auto_enable_io=self.auto_enable_io,
            availability_zone=self.availability_zone,
            encrypted=self.encrypted,
            id=self.id,
            iops=self.iops,
            kms_key_id=self.kms_key_id,
            multi_attach_enabled=self.multi_attach_enabled,
            outpost_arn=self.outpost_arn,
            size=self.size,
            snapshot_id=self.snapshot_id,
            tags=self.tags,
            throughput=self.throughput,
            volume_type=self.volume_type)


def get_volume(id: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVolumeResult:
    """
    Resource Type definition for AWS::EC2::Volume
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:ec2:getVolume', __args__, opts=opts, typ=GetVolumeResult).value

    return AwaitableGetVolumeResult(
        auto_enable_io=__ret__.auto_enable_io,
        availability_zone=__ret__.availability_zone,
        encrypted=__ret__.encrypted,
        id=__ret__.id,
        iops=__ret__.iops,
        kms_key_id=__ret__.kms_key_id,
        multi_attach_enabled=__ret__.multi_attach_enabled,
        outpost_arn=__ret__.outpost_arn,
        size=__ret__.size,
        snapshot_id=__ret__.snapshot_id,
        tags=__ret__.tags,
        throughput=__ret__.throughput,
        volume_type=__ret__.volume_type)


@_utilities.lift_output_func(get_volume)
def get_volume_output(id: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVolumeResult]:
    """
    Resource Type definition for AWS::EC2::Volume
    """
    ...