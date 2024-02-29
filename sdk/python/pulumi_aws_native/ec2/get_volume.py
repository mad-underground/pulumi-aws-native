# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from .. import outputs as _root_outputs

__all__ = [
    'GetVolumeResult',
    'AwaitableGetVolumeResult',
    'get_volume',
    'get_volume_output',
]

@pulumi.output_type
class GetVolumeResult:
    def __init__(__self__, auto_enable_io=None, availability_zone=None, encrypted=None, iops=None, kms_key_id=None, multi_attach_enabled=None, outpost_arn=None, size=None, snapshot_id=None, tags=None, throughput=None, volume_id=None, volume_type=None):
        if auto_enable_io and not isinstance(auto_enable_io, bool):
            raise TypeError("Expected argument 'auto_enable_io' to be a bool")
        pulumi.set(__self__, "auto_enable_io", auto_enable_io)
        if availability_zone and not isinstance(availability_zone, str):
            raise TypeError("Expected argument 'availability_zone' to be a str")
        pulumi.set(__self__, "availability_zone", availability_zone)
        if encrypted and not isinstance(encrypted, bool):
            raise TypeError("Expected argument 'encrypted' to be a bool")
        pulumi.set(__self__, "encrypted", encrypted)
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
        if volume_id and not isinstance(volume_id, str):
            raise TypeError("Expected argument 'volume_id' to be a str")
        pulumi.set(__self__, "volume_id", volume_id)
        if volume_type and not isinstance(volume_type, str):
            raise TypeError("Expected argument 'volume_type' to be a str")
        pulumi.set(__self__, "volume_type", volume_type)

    @property
    @pulumi.getter(name="autoEnableIo")
    def auto_enable_io(self) -> Optional[bool]:
        """
        Indicates whether the volume is auto-enabled for I/O operations. By default, Amazon EBS disables I/O to the volume from attached EC2 instances when it determines that a volume's data is potentially inconsistent. If the consistency of the volume is not a concern, and you prefer that the volume be made available immediately if it's impaired, you can configure the volume to automatically enable I/O.
        """
        return pulumi.get(self, "auto_enable_io")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[str]:
        """
        The ID of the Availability Zone in which to create the volume. For example, ``us-east-1a``.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter
    def encrypted(self) -> Optional[bool]:
        """
        Indicates whether the volume should be encrypted. The effect of setting the encryption state to ``true`` depends on the volume origin (new or from a snapshot), starting encryption state, ownership, and whether encryption by default is enabled. For more information, see [Encryption by default](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#encryption-by-default) in the *Amazon Elastic Compute Cloud User Guide*.
         Encrypted Amazon EBS volumes must be attached to instances that support Amazon EBS encryption. For more information, see [Supported instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#EBSEncryption_supported_instances).
        """
        return pulumi.get(self, "encrypted")

    @property
    @pulumi.getter
    def iops(self) -> Optional[int]:
        """
        The number of I/O operations per second (IOPS). For ``gp3``, ``io1``, and ``io2`` volumes, this represents the number of IOPS that are provisioned for the volume. For ``gp2`` volumes, this represents the baseline performance of the volume and the rate at which the volume accumulates I/O credits for bursting.
         The following are the supported values for each volume type:
          +   ``gp3``: 3,000 - 16,000 IOPS
          +   ``io1``: 100 - 64,000 IOPS
          +   ``io2``: 100 - 256,000 IOPS
          
         For ``io2`` volumes, you can achieve up to 256,000 IOPS on [instances built on the Nitro System](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#ec2-nitro-instances). On other instances, you can achieve performance up to 32,000 IOPS.
         This parameter is required for ``io1`` and ``io2`` volumes. The default for ``gp3`` volumes is 3,000 IOPS. This parameter is not supported for ``gp2``, ``st1``, ``sc1``, or ``standard`` volumes.
        """
        return pulumi.get(self, "iops")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[str]:
        """
        The identifier of the kms-key-long to use for Amazon EBS encryption. If ``KmsKeyId`` is specified, the encrypted state must be ``true``.
         If you omit this property and your account is enabled for encryption by default, or *Encrypted* is set to ``true``, then the volume is encrypted using the default key specified for your account. If your account does not have a default key, then the volume is encrypted using the aws-managed-key.
         Alternatively, if you want to specify a different key, you can specify one of the following:
          +  Key ID. For example, 1234abcd-12ab-34cd-56ef-1234567890ab.
          +  Key alias. Specify the alias for the key, prefixed with ``alias/``. For example, for a key with the alias ``my_cmk``, use ``alias/my_cmk``. Or to specify the aws-managed-key, use ``alias/aws/ebs``.
          +  Key ARN. For example, arn:aws:kms:us-east-1:012345678910:key/1234abcd-12ab-34cd-56ef-1234567890ab.
          +  Alias ARN. For example, arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.
        """
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter(name="multiAttachEnabled")
    def multi_attach_enabled(self) -> Optional[bool]:
        """
        Indicates whether Amazon EBS Multi-Attach is enabled.
         CFNlong does not currently support updating a single-attach volume to be multi-attach enabled, updating a multi-attach enabled volume to be single-attach, or updating the size or number of I/O operations per second (IOPS) of a multi-attach enabled volume.
        """
        return pulumi.get(self, "multi_attach_enabled")

    @property
    @pulumi.getter(name="outpostArn")
    def outpost_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the Outpost.
        """
        return pulumi.get(self, "outpost_arn")

    @property
    @pulumi.getter
    def size(self) -> Optional[int]:
        """
        The size of the volume, in GiBs. You must specify either a snapshot ID or a volume size. If you specify a snapshot, the default is the snapshot size. You can specify a volume size that is equal to or larger than the snapshot size.
         The following are the supported volumes sizes for each volume type:
          +   ``gp2`` and ``gp3``: 1 - 16,384 GiB
          +   ``io1``: 4 - 16,384 GiB
          +   ``io2``: 4 - 65,536 GiB
          +   ``st1`` and ``sc1``: 125 - 16,384 GiB
          +   ``standard``: 1 - 1024 GiB
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> Optional[str]:
        """
        The snapshot from which to create the volume. You must specify either a snapshot ID or a volume size.
        """
        return pulumi.get(self, "snapshot_id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        The tags to apply to the volume during creation.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def throughput(self) -> Optional[int]:
        """
        The throughput to provision for a volume, with a maximum of 1,000 MiB/s.
         This parameter is valid only for ``gp3`` volumes. The default value is 125.
         Valid Range: Minimum value of 125. Maximum value of 1000.
        """
        return pulumi.get(self, "throughput")

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> Optional[str]:
        return pulumi.get(self, "volume_id")

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> Optional[str]:
        """
        The volume type. This parameter can be one of the following values:
          +  General Purpose SSD: ``gp2`` | ``gp3`` 
          +  Provisioned IOPS SSD: ``io1`` | ``io2`` 
          +  Throughput Optimized HDD: ``st1`` 
          +  Cold HDD: ``sc1`` 
          +  Magnetic: ``standard`` 
          
         For more information, see [Amazon EBS volume types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html) in the *Amazon Elastic Compute Cloud User Guide*.
         Default: ``gp2``
        """
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
            iops=self.iops,
            kms_key_id=self.kms_key_id,
            multi_attach_enabled=self.multi_attach_enabled,
            outpost_arn=self.outpost_arn,
            size=self.size,
            snapshot_id=self.snapshot_id,
            tags=self.tags,
            throughput=self.throughput,
            volume_id=self.volume_id,
            volume_type=self.volume_type)


def get_volume(volume_id: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVolumeResult:
    """
    Specifies an Amazon Elastic Block Store (Amazon EBS) volume.
     When you use CFNlong to update an Amazon EBS volume that modifies ``Iops``, ``Size``, or ``VolumeType``, there is a cooldown period before another operation can occur. This can cause your stack to report being in ``UPDATE_IN_PROGRESS`` or ``UPDATE_ROLLBACK_IN_PROGRESS`` for long periods of time.
     Amazon EBS does not support sizing down an Amazon EBS volume. CFNlong does not attempt to modify an Amazon EBS volume to a smaller size on rollback.
     Some common scenarios when you might encounter a cooldown period for Amazon EBS include:
      +  You successfully update an Amazon EBS volume and the update succeeds. When you attempt another update within the cooldown window, that update will be subject to a cooldown period.
      +  You successfully update an Amazon EBS volume and the update succeeds but another change in your ``update-stack`` call fails. The rollback will be subject to a cooldown period.

     For more information on the coo
    """
    __args__ = dict()
    __args__['volumeId'] = volume_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:ec2:getVolume', __args__, opts=opts, typ=GetVolumeResult).value

    return AwaitableGetVolumeResult(
        auto_enable_io=pulumi.get(__ret__, 'auto_enable_io'),
        availability_zone=pulumi.get(__ret__, 'availability_zone'),
        encrypted=pulumi.get(__ret__, 'encrypted'),
        iops=pulumi.get(__ret__, 'iops'),
        kms_key_id=pulumi.get(__ret__, 'kms_key_id'),
        multi_attach_enabled=pulumi.get(__ret__, 'multi_attach_enabled'),
        outpost_arn=pulumi.get(__ret__, 'outpost_arn'),
        size=pulumi.get(__ret__, 'size'),
        snapshot_id=pulumi.get(__ret__, 'snapshot_id'),
        tags=pulumi.get(__ret__, 'tags'),
        throughput=pulumi.get(__ret__, 'throughput'),
        volume_id=pulumi.get(__ret__, 'volume_id'),
        volume_type=pulumi.get(__ret__, 'volume_type'))


@_utilities.lift_output_func(get_volume)
def get_volume_output(volume_id: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVolumeResult]:
    """
    Specifies an Amazon Elastic Block Store (Amazon EBS) volume.
     When you use CFNlong to update an Amazon EBS volume that modifies ``Iops``, ``Size``, or ``VolumeType``, there is a cooldown period before another operation can occur. This can cause your stack to report being in ``UPDATE_IN_PROGRESS`` or ``UPDATE_ROLLBACK_IN_PROGRESS`` for long periods of time.
     Amazon EBS does not support sizing down an Amazon EBS volume. CFNlong does not attempt to modify an Amazon EBS volume to a smaller size on rollback.
     Some common scenarios when you might encounter a cooldown period for Amazon EBS include:
      +  You successfully update an Amazon EBS volume and the update succeeds. When you attempt another update within the cooldown window, that update will be subject to a cooldown period.
      +  You successfully update an Amazon EBS volume and the update succeeds but another change in your ``update-stack`` call fails. The rollback will be subject to a cooldown period.

     For more information on the coo
    """
    ...
