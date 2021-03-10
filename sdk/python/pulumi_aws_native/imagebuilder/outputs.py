# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'DistributionConfigurationDistribution',
    'ImageImageTestsConfiguration',
    'ImagePipelineImageTestsConfiguration',
    'ImagePipelineSchedule',
    'ImageRecipeComponentConfiguration',
    'ImageRecipeEbsInstanceBlockDeviceSpecification',
    'ImageRecipeInstanceBlockDeviceMapping',
]

@pulumi.output_type
class DistributionConfigurationDistribution(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-distribution.html
    """
    def __init__(__self__, *,
                 region: str,
                 ami_distribution_configuration: Optional[str] = None,
                 license_configuration_arns: Optional[Sequence[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-distribution.html
        :param str region: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-distribution.html#cfn-imagebuilder-distributionconfiguration-distribution-region
        :param Union[Any, str] ami_distribution_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-distribution.html#cfn-imagebuilder-distributionconfiguration-distribution-amidistributionconfiguration
        :param Sequence[str] license_configuration_arns: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-distribution.html#cfn-imagebuilder-distributionconfiguration-distribution-licenseconfigurationarns
        """
        pulumi.set(__self__, "region", region)
        if ami_distribution_configuration is not None:
            pulumi.set(__self__, "ami_distribution_configuration", ami_distribution_configuration)
        if license_configuration_arns is not None:
            pulumi.set(__self__, "license_configuration_arns", license_configuration_arns)

    @property
    @pulumi.getter(name="Region")
    def region(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-distribution.html#cfn-imagebuilder-distributionconfiguration-distribution-region
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="AmiDistributionConfiguration")
    def ami_distribution_configuration(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-distribution.html#cfn-imagebuilder-distributionconfiguration-distribution-amidistributionconfiguration
        """
        return pulumi.get(self, "ami_distribution_configuration")

    @property
    @pulumi.getter(name="LicenseConfigurationArns")
    def license_configuration_arns(self) -> Optional[Sequence[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-distribution.html#cfn-imagebuilder-distributionconfiguration-distribution-licenseconfigurationarns
        """
        return pulumi.get(self, "license_configuration_arns")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ImageImageTestsConfiguration(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-imagetestsconfiguration.html
    """
    def __init__(__self__, *,
                 image_tests_enabled: Optional[bool] = None,
                 timeout_minutes: Optional[int] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-imagetestsconfiguration.html
        :param bool image_tests_enabled: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-imagetestsconfiguration.html#cfn-imagebuilder-image-imagetestsconfiguration-imagetestsenabled
        :param int timeout_minutes: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-imagetestsconfiguration.html#cfn-imagebuilder-image-imagetestsconfiguration-timeoutminutes
        """
        if image_tests_enabled is not None:
            pulumi.set(__self__, "image_tests_enabled", image_tests_enabled)
        if timeout_minutes is not None:
            pulumi.set(__self__, "timeout_minutes", timeout_minutes)

    @property
    @pulumi.getter(name="ImageTestsEnabled")
    def image_tests_enabled(self) -> Optional[bool]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-imagetestsconfiguration.html#cfn-imagebuilder-image-imagetestsconfiguration-imagetestsenabled
        """
        return pulumi.get(self, "image_tests_enabled")

    @property
    @pulumi.getter(name="TimeoutMinutes")
    def timeout_minutes(self) -> Optional[int]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-imagetestsconfiguration.html#cfn-imagebuilder-image-imagetestsconfiguration-timeoutminutes
        """
        return pulumi.get(self, "timeout_minutes")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ImagePipelineImageTestsConfiguration(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-imagetestsconfiguration.html
    """
    def __init__(__self__, *,
                 image_tests_enabled: Optional[bool] = None,
                 timeout_minutes: Optional[int] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-imagetestsconfiguration.html
        :param bool image_tests_enabled: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-imagetestsconfiguration.html#cfn-imagebuilder-imagepipeline-imagetestsconfiguration-imagetestsenabled
        :param int timeout_minutes: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-imagetestsconfiguration.html#cfn-imagebuilder-imagepipeline-imagetestsconfiguration-timeoutminutes
        """
        if image_tests_enabled is not None:
            pulumi.set(__self__, "image_tests_enabled", image_tests_enabled)
        if timeout_minutes is not None:
            pulumi.set(__self__, "timeout_minutes", timeout_minutes)

    @property
    @pulumi.getter(name="ImageTestsEnabled")
    def image_tests_enabled(self) -> Optional[bool]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-imagetestsconfiguration.html#cfn-imagebuilder-imagepipeline-imagetestsconfiguration-imagetestsenabled
        """
        return pulumi.get(self, "image_tests_enabled")

    @property
    @pulumi.getter(name="TimeoutMinutes")
    def timeout_minutes(self) -> Optional[int]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-imagetestsconfiguration.html#cfn-imagebuilder-imagepipeline-imagetestsconfiguration-timeoutminutes
        """
        return pulumi.get(self, "timeout_minutes")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ImagePipelineSchedule(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-schedule.html
    """
    def __init__(__self__, *,
                 pipeline_execution_start_condition: Optional[str] = None,
                 schedule_expression: Optional[str] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-schedule.html
        :param str pipeline_execution_start_condition: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-schedule.html#cfn-imagebuilder-imagepipeline-schedule-pipelineexecutionstartcondition
        :param str schedule_expression: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-schedule.html#cfn-imagebuilder-imagepipeline-schedule-scheduleexpression
        """
        if pipeline_execution_start_condition is not None:
            pulumi.set(__self__, "pipeline_execution_start_condition", pipeline_execution_start_condition)
        if schedule_expression is not None:
            pulumi.set(__self__, "schedule_expression", schedule_expression)

    @property
    @pulumi.getter(name="PipelineExecutionStartCondition")
    def pipeline_execution_start_condition(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-schedule.html#cfn-imagebuilder-imagepipeline-schedule-pipelineexecutionstartcondition
        """
        return pulumi.get(self, "pipeline_execution_start_condition")

    @property
    @pulumi.getter(name="ScheduleExpression")
    def schedule_expression(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-schedule.html#cfn-imagebuilder-imagepipeline-schedule-scheduleexpression
        """
        return pulumi.get(self, "schedule_expression")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ImageRecipeComponentConfiguration(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-componentconfiguration.html
    """
    def __init__(__self__, *,
                 component_arn: Optional[str] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-componentconfiguration.html
        :param str component_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-componentconfiguration.html#cfn-imagebuilder-imagerecipe-componentconfiguration-componentarn
        """
        if component_arn is not None:
            pulumi.set(__self__, "component_arn", component_arn)

    @property
    @pulumi.getter(name="ComponentArn")
    def component_arn(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-componentconfiguration.html#cfn-imagebuilder-imagerecipe-componentconfiguration-componentarn
        """
        return pulumi.get(self, "component_arn")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ImageRecipeEbsInstanceBlockDeviceSpecification(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html
    """
    def __init__(__self__, *,
                 delete_on_termination: Optional[bool] = None,
                 encrypted: Optional[bool] = None,
                 iops: Optional[int] = None,
                 kms_key_id: Optional[str] = None,
                 snapshot_id: Optional[str] = None,
                 volume_size: Optional[int] = None,
                 volume_type: Optional[str] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html
        :param bool delete_on_termination: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification-deleteontermination
        :param bool encrypted: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification-encrypted
        :param int iops: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification-iops
        :param str kms_key_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification-kmskeyid
        :param str snapshot_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification-snapshotid
        :param int volume_size: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification-volumesize
        :param str volume_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification-volumetype
        """
        if delete_on_termination is not None:
            pulumi.set(__self__, "delete_on_termination", delete_on_termination)
        if encrypted is not None:
            pulumi.set(__self__, "encrypted", encrypted)
        if iops is not None:
            pulumi.set(__self__, "iops", iops)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)
        if snapshot_id is not None:
            pulumi.set(__self__, "snapshot_id", snapshot_id)
        if volume_size is not None:
            pulumi.set(__self__, "volume_size", volume_size)
        if volume_type is not None:
            pulumi.set(__self__, "volume_type", volume_type)

    @property
    @pulumi.getter(name="DeleteOnTermination")
    def delete_on_termination(self) -> Optional[bool]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification-deleteontermination
        """
        return pulumi.get(self, "delete_on_termination")

    @property
    @pulumi.getter(name="Encrypted")
    def encrypted(self) -> Optional[bool]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification-encrypted
        """
        return pulumi.get(self, "encrypted")

    @property
    @pulumi.getter(name="Iops")
    def iops(self) -> Optional[int]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification-iops
        """
        return pulumi.get(self, "iops")

    @property
    @pulumi.getter(name="KmsKeyId")
    def kms_key_id(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification-kmskeyid
        """
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter(name="SnapshotId")
    def snapshot_id(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification-snapshotid
        """
        return pulumi.get(self, "snapshot_id")

    @property
    @pulumi.getter(name="VolumeSize")
    def volume_size(self) -> Optional[int]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification-volumesize
        """
        return pulumi.get(self, "volume_size")

    @property
    @pulumi.getter(name="VolumeType")
    def volume_type(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification-volumetype
        """
        return pulumi.get(self, "volume_type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ImageRecipeInstanceBlockDeviceMapping(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-instanceblockdevicemapping.html
    """
    def __init__(__self__, *,
                 device_name: Optional[str] = None,
                 ebs: Optional['outputs.ImageRecipeEbsInstanceBlockDeviceSpecification'] = None,
                 no_device: Optional[str] = None,
                 virtual_name: Optional[str] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-instanceblockdevicemapping.html
        :param str device_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-instanceblockdevicemapping.html#cfn-imagebuilder-imagerecipe-instanceblockdevicemapping-devicename
        :param 'ImageRecipeEbsInstanceBlockDeviceSpecificationArgs' ebs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-instanceblockdevicemapping.html#cfn-imagebuilder-imagerecipe-instanceblockdevicemapping-ebs
        :param str no_device: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-instanceblockdevicemapping.html#cfn-imagebuilder-imagerecipe-instanceblockdevicemapping-nodevice
        :param str virtual_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-instanceblockdevicemapping.html#cfn-imagebuilder-imagerecipe-instanceblockdevicemapping-virtualname
        """
        if device_name is not None:
            pulumi.set(__self__, "device_name", device_name)
        if ebs is not None:
            pulumi.set(__self__, "ebs", ebs)
        if no_device is not None:
            pulumi.set(__self__, "no_device", no_device)
        if virtual_name is not None:
            pulumi.set(__self__, "virtual_name", virtual_name)

    @property
    @pulumi.getter(name="DeviceName")
    def device_name(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-instanceblockdevicemapping.html#cfn-imagebuilder-imagerecipe-instanceblockdevicemapping-devicename
        """
        return pulumi.get(self, "device_name")

    @property
    @pulumi.getter(name="Ebs")
    def ebs(self) -> Optional['outputs.ImageRecipeEbsInstanceBlockDeviceSpecification']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-instanceblockdevicemapping.html#cfn-imagebuilder-imagerecipe-instanceblockdevicemapping-ebs
        """
        return pulumi.get(self, "ebs")

    @property
    @pulumi.getter(name="NoDevice")
    def no_device(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-instanceblockdevicemapping.html#cfn-imagebuilder-imagerecipe-instanceblockdevicemapping-nodevice
        """
        return pulumi.get(self, "no_device")

    @property
    @pulumi.getter(name="VirtualName")
    def virtual_name(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-instanceblockdevicemapping.html#cfn-imagebuilder-imagerecipe-instanceblockdevicemapping-virtualname
        """
        return pulumi.get(self, "virtual_name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


