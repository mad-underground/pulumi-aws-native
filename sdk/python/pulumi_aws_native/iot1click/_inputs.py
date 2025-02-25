# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ProjectPlacementTemplateArgs',
]

@pulumi.input_type
class ProjectPlacementTemplateArgs:
    def __init__(__self__, *,
                 default_attributes: Optional[Any] = None,
                 device_templates: Optional[Any] = None):
        if default_attributes is not None:
            pulumi.set(__self__, "default_attributes", default_attributes)
        if device_templates is not None:
            pulumi.set(__self__, "device_templates", device_templates)

    @property
    @pulumi.getter(name="defaultAttributes")
    def default_attributes(self) -> Optional[Any]:
        return pulumi.get(self, "default_attributes")

    @default_attributes.setter
    def default_attributes(self, value: Optional[Any]):
        pulumi.set(self, "default_attributes", value)

    @property
    @pulumi.getter(name="deviceTemplates")
    def device_templates(self) -> Optional[Any]:
        return pulumi.get(self, "device_templates")

    @device_templates.setter
    def device_templates(self, value: Optional[Any]):
        pulumi.set(self, "device_templates", value)


