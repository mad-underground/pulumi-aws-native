# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'AssociationInstanceAssociationOutputLocationArgs',
    'AssociationParameterValuesArgs',
    'AssociationS3OutputLocationArgs',
    'AssociationTargetArgs',
]

@pulumi.input_type
class AssociationInstanceAssociationOutputLocationArgs:
    def __init__(__self__, *,
                 s3_location: Optional[pulumi.Input['AssociationS3OutputLocationArgs']] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-instanceassociationoutputlocation.html
        :param pulumi.Input['AssociationS3OutputLocationArgs'] s3_location: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-instanceassociationoutputlocation.html#cfn-ssm-association-instanceassociationoutputlocation-s3location
        """
        if s3_location is not None:
            pulumi.set(__self__, "s3_location", s3_location)

    @property
    @pulumi.getter(name="S3Location")
    def s3_location(self) -> Optional[pulumi.Input['AssociationS3OutputLocationArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-instanceassociationoutputlocation.html#cfn-ssm-association-instanceassociationoutputlocation-s3location
        """
        return pulumi.get(self, "s3_location")

    @s3_location.setter
    def s3_location(self, value: Optional[pulumi.Input['AssociationS3OutputLocationArgs']]):
        pulumi.set(self, "s3_location", value)


@pulumi.input_type
class AssociationParameterValuesArgs:
    def __init__(__self__, *,
                 parameter_values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-parametervalues.html
        :param pulumi.Input[Sequence[pulumi.Input[str]]] parameter_values: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-parametervalues.html#cfn-ssm-association-parametervalues-parametervalues
        """
        if parameter_values is not None:
            pulumi.set(__self__, "parameter_values", parameter_values)

    @property
    @pulumi.getter(name="ParameterValues")
    def parameter_values(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-parametervalues.html#cfn-ssm-association-parametervalues-parametervalues
        """
        return pulumi.get(self, "parameter_values")

    @parameter_values.setter
    def parameter_values(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "parameter_values", value)


@pulumi.input_type
class AssociationS3OutputLocationArgs:
    def __init__(__self__, *,
                 output_s3_bucket_name: Optional[pulumi.Input[str]] = None,
                 output_s3_key_prefix: Optional[pulumi.Input[str]] = None,
                 output_s3_region: Optional[pulumi.Input[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-s3outputlocation.html
        :param pulumi.Input[str] output_s3_bucket_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-s3outputlocation.html#cfn-ssm-association-s3outputlocation-outputs3bucketname
        :param pulumi.Input[str] output_s3_key_prefix: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-s3outputlocation.html#cfn-ssm-association-s3outputlocation-outputs3keyprefix
        :param pulumi.Input[str] output_s3_region: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-s3outputlocation.html#cfn-ssm-association-s3outputlocation-outputs3region
        """
        if output_s3_bucket_name is not None:
            pulumi.set(__self__, "output_s3_bucket_name", output_s3_bucket_name)
        if output_s3_key_prefix is not None:
            pulumi.set(__self__, "output_s3_key_prefix", output_s3_key_prefix)
        if output_s3_region is not None:
            pulumi.set(__self__, "output_s3_region", output_s3_region)

    @property
    @pulumi.getter(name="OutputS3BucketName")
    def output_s3_bucket_name(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-s3outputlocation.html#cfn-ssm-association-s3outputlocation-outputs3bucketname
        """
        return pulumi.get(self, "output_s3_bucket_name")

    @output_s3_bucket_name.setter
    def output_s3_bucket_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "output_s3_bucket_name", value)

    @property
    @pulumi.getter(name="OutputS3KeyPrefix")
    def output_s3_key_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-s3outputlocation.html#cfn-ssm-association-s3outputlocation-outputs3keyprefix
        """
        return pulumi.get(self, "output_s3_key_prefix")

    @output_s3_key_prefix.setter
    def output_s3_key_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "output_s3_key_prefix", value)

    @property
    @pulumi.getter(name="OutputS3Region")
    def output_s3_region(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-s3outputlocation.html#cfn-ssm-association-s3outputlocation-outputs3region
        """
        return pulumi.get(self, "output_s3_region")

    @output_s3_region.setter
    def output_s3_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "output_s3_region", value)


@pulumi.input_type
class AssociationTargetArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 values: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-target.html
        :param pulumi.Input[str] key: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-target.html#cfn-ssm-association-target-key
        :param pulumi.Input[Sequence[pulumi.Input[str]]] values: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-target.html#cfn-ssm-association-target-values
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter(name="Key")
    def key(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-target.html#cfn-ssm-association-target-key
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="Values")
    def values(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-target.html#cfn-ssm-association-target-values
        """
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "values", value)


