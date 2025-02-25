# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = [
    'ApplicationDefinition0PropertiesArgs',
    'ApplicationDefinition1PropertiesArgs',
    'EnvironmentHighAvailabilityConfigArgs',
    'EnvironmentStorageConfigurationArgs',
]

@pulumi.input_type
class ApplicationDefinition0PropertiesArgs:
    def __init__(__self__, *,
                 s3_location: pulumi.Input[str]):
        pulumi.set(__self__, "s3_location", s3_location)

    @property
    @pulumi.getter(name="s3Location")
    def s3_location(self) -> pulumi.Input[str]:
        return pulumi.get(self, "s3_location")

    @s3_location.setter
    def s3_location(self, value: pulumi.Input[str]):
        pulumi.set(self, "s3_location", value)


@pulumi.input_type
class ApplicationDefinition1PropertiesArgs:
    def __init__(__self__, *,
                 content: pulumi.Input[str]):
        pulumi.set(__self__, "content", content)

    @property
    @pulumi.getter
    def content(self) -> pulumi.Input[str]:
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: pulumi.Input[str]):
        pulumi.set(self, "content", value)


@pulumi.input_type
class EnvironmentHighAvailabilityConfigArgs:
    def __init__(__self__, *,
                 desired_capacity: pulumi.Input[int]):
        """
        Defines the details of a high availability configuration.
        """
        pulumi.set(__self__, "desired_capacity", desired_capacity)

    @property
    @pulumi.getter(name="desiredCapacity")
    def desired_capacity(self) -> pulumi.Input[int]:
        return pulumi.get(self, "desired_capacity")

    @desired_capacity.setter
    def desired_capacity(self, value: pulumi.Input[int]):
        pulumi.set(self, "desired_capacity", value)


@pulumi.input_type
class EnvironmentStorageConfigurationArgs:
    def __init__(__self__):
        """
        Defines the storage configuration for an environment.
        """
        pass


