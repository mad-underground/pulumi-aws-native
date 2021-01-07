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
    'DeviceAttributes',
    'DeviceProperties',
    'PlacementAttributes',
    'PlacementProperties',
    'ProjectAttributes',
    'ProjectPlacementTemplate',
    'ProjectProperties',
]

@pulumi.output_type
class DeviceAttributes(dict):
    def __init__(__self__, *,
                 arn: str,
                 device_id: str,
                 enabled: bool):
        pulumi.set(__self__, "arn", arn)
        pulumi.set(__self__, "device_id", device_id)
        pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter(name="Arn")
    def arn(self) -> str:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="DeviceId")
    def device_id(self) -> str:
        return pulumi.get(self, "device_id")

    @property
    @pulumi.getter(name="Enabled")
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DeviceProperties(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-device.html
    """
    def __init__(__self__, *,
                 device_id: str,
                 enabled: bool):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-device.html
        :param str device_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-device.html#cfn-iot1click-device-deviceid
        :param bool enabled: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-device.html#cfn-iot1click-device-enabled
        """
        pulumi.set(__self__, "device_id", device_id)
        pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter(name="DeviceId")
    def device_id(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-device.html#cfn-iot1click-device-deviceid
        """
        return pulumi.get(self, "device_id")

    @property
    @pulumi.getter(name="Enabled")
    def enabled(self) -> bool:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-device.html#cfn-iot1click-device-enabled
        """
        return pulumi.get(self, "enabled")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PlacementAttributes(dict):
    def __init__(__self__, *,
                 placement_name: str,
                 project_name: str):
        pulumi.set(__self__, "placement_name", placement_name)
        pulumi.set(__self__, "project_name", project_name)

    @property
    @pulumi.getter(name="PlacementName")
    def placement_name(self) -> str:
        return pulumi.get(self, "placement_name")

    @property
    @pulumi.getter(name="ProjectName")
    def project_name(self) -> str:
        return pulumi.get(self, "project_name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PlacementProperties(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-placement.html
    """
    def __init__(__self__, *,
                 project_name: str,
                 associated_devices: Optional[str] = None,
                 attributes: Optional[str] = None,
                 placement_name: Optional[str] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-placement.html
        :param str project_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-placement.html#cfn-iot1click-placement-projectname
        :param Union[Any, str] associated_devices: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-placement.html#cfn-iot1click-placement-associateddevices
        :param Union[Any, str] attributes: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-placement.html#cfn-iot1click-placement-attributes
        :param str placement_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-placement.html#cfn-iot1click-placement-placementname
        """
        pulumi.set(__self__, "project_name", project_name)
        if associated_devices is not None:
            pulumi.set(__self__, "associated_devices", associated_devices)
        if attributes is not None:
            pulumi.set(__self__, "attributes", attributes)
        if placement_name is not None:
            pulumi.set(__self__, "placement_name", placement_name)

    @property
    @pulumi.getter(name="ProjectName")
    def project_name(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-placement.html#cfn-iot1click-placement-projectname
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter(name="AssociatedDevices")
    def associated_devices(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-placement.html#cfn-iot1click-placement-associateddevices
        """
        return pulumi.get(self, "associated_devices")

    @property
    @pulumi.getter(name="Attributes")
    def attributes(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-placement.html#cfn-iot1click-placement-attributes
        """
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter(name="PlacementName")
    def placement_name(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-placement.html#cfn-iot1click-placement-placementname
        """
        return pulumi.get(self, "placement_name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ProjectAttributes(dict):
    def __init__(__self__, *,
                 arn: str,
                 project_name: str):
        pulumi.set(__self__, "arn", arn)
        pulumi.set(__self__, "project_name", project_name)

    @property
    @pulumi.getter(name="Arn")
    def arn(self) -> str:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="ProjectName")
    def project_name(self) -> str:
        return pulumi.get(self, "project_name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ProjectPlacementTemplate(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot1click-project-placementtemplate.html
    """
    def __init__(__self__, *,
                 default_attributes: Optional[str] = None,
                 device_templates: Optional[str] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot1click-project-placementtemplate.html
        :param Union[Any, str] default_attributes: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot1click-project-placementtemplate.html#cfn-iot1click-project-placementtemplate-defaultattributes
        :param Union[Any, str] device_templates: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot1click-project-placementtemplate.html#cfn-iot1click-project-placementtemplate-devicetemplates
        """
        if default_attributes is not None:
            pulumi.set(__self__, "default_attributes", default_attributes)
        if device_templates is not None:
            pulumi.set(__self__, "device_templates", device_templates)

    @property
    @pulumi.getter(name="DefaultAttributes")
    def default_attributes(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot1click-project-placementtemplate.html#cfn-iot1click-project-placementtemplate-defaultattributes
        """
        return pulumi.get(self, "default_attributes")

    @property
    @pulumi.getter(name="DeviceTemplates")
    def device_templates(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot1click-project-placementtemplate.html#cfn-iot1click-project-placementtemplate-devicetemplates
        """
        return pulumi.get(self, "device_templates")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ProjectProperties(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-project.html
    """
    def __init__(__self__, *,
                 placement_template: 'outputs.ProjectPlacementTemplate',
                 description: Optional[str] = None,
                 project_name: Optional[str] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-project.html
        :param 'ProjectPlacementTemplateArgs' placement_template: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-project.html#cfn-iot1click-project-placementtemplate
        :param str description: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-project.html#cfn-iot1click-project-description
        :param str project_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-project.html#cfn-iot1click-project-projectname
        """
        pulumi.set(__self__, "placement_template", placement_template)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)

    @property
    @pulumi.getter(name="PlacementTemplate")
    def placement_template(self) -> 'outputs.ProjectPlacementTemplate':
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-project.html#cfn-iot1click-project-placementtemplate
        """
        return pulumi.get(self, "placement_template")

    @property
    @pulumi.getter(name="Description")
    def description(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-project.html#cfn-iot1click-project-description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="ProjectName")
    def project_name(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-project.html#cfn-iot1click-project-projectname
        """
        return pulumi.get(self, "project_name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

