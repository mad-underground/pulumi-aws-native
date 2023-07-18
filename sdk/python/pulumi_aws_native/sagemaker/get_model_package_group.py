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
    'GetModelPackageGroupResult',
    'AwaitableGetModelPackageGroupResult',
    'get_model_package_group',
    'get_model_package_group_output',
]

@pulumi.output_type
class GetModelPackageGroupResult:
    def __init__(__self__, creation_time=None, model_package_group_arn=None, model_package_group_policy=None, model_package_group_status=None, tags=None):
        if creation_time and not isinstance(creation_time, str):
            raise TypeError("Expected argument 'creation_time' to be a str")
        pulumi.set(__self__, "creation_time", creation_time)
        if model_package_group_arn and not isinstance(model_package_group_arn, str):
            raise TypeError("Expected argument 'model_package_group_arn' to be a str")
        pulumi.set(__self__, "model_package_group_arn", model_package_group_arn)
        if model_package_group_policy and not isinstance(model_package_group_policy, dict):
            raise TypeError("Expected argument 'model_package_group_policy' to be a dict")
        pulumi.set(__self__, "model_package_group_policy", model_package_group_policy)
        if model_package_group_status and not isinstance(model_package_group_status, str):
            raise TypeError("Expected argument 'model_package_group_status' to be a str")
        pulumi.set(__self__, "model_package_group_status", model_package_group_status)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[str]:
        """
        The time at which the model package group was created.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="modelPackageGroupArn")
    def model_package_group_arn(self) -> Optional[str]:
        return pulumi.get(self, "model_package_group_arn")

    @property
    @pulumi.getter(name="modelPackageGroupPolicy")
    def model_package_group_policy(self) -> Optional[Any]:
        return pulumi.get(self, "model_package_group_policy")

    @property
    @pulumi.getter(name="modelPackageGroupStatus")
    def model_package_group_status(self) -> Optional['ModelPackageGroupStatus']:
        """
        The status of a modelpackage group job.
        """
        return pulumi.get(self, "model_package_group_status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.ModelPackageGroupTag']]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")


class AwaitableGetModelPackageGroupResult(GetModelPackageGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetModelPackageGroupResult(
            creation_time=self.creation_time,
            model_package_group_arn=self.model_package_group_arn,
            model_package_group_policy=self.model_package_group_policy,
            model_package_group_status=self.model_package_group_status,
            tags=self.tags)


def get_model_package_group(model_package_group_arn: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetModelPackageGroupResult:
    """
    Resource Type definition for AWS::SageMaker::ModelPackageGroup
    """
    __args__ = dict()
    __args__['modelPackageGroupArn'] = model_package_group_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:sagemaker:getModelPackageGroup', __args__, opts=opts, typ=GetModelPackageGroupResult).value

    return AwaitableGetModelPackageGroupResult(
        creation_time=pulumi.get(__ret__, 'creation_time'),
        model_package_group_arn=pulumi.get(__ret__, 'model_package_group_arn'),
        model_package_group_policy=pulumi.get(__ret__, 'model_package_group_policy'),
        model_package_group_status=pulumi.get(__ret__, 'model_package_group_status'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_model_package_group)
def get_model_package_group_output(model_package_group_arn: Optional[pulumi.Input[str]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetModelPackageGroupResult]:
    """
    Resource Type definition for AWS::SageMaker::ModelPackageGroup
    """
    ...
