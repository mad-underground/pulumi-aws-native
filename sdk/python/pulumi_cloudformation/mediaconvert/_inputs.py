# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'JobTemplateAccelerationSettingsArgs',
    'JobTemplateHopDestinationArgs',
    'JobTemplatePropertiesArgs',
    'PresetPropertiesArgs',
    'QueuePropertiesArgs',
]

@pulumi.input_type
class JobTemplateAccelerationSettingsArgs:
    def __init__(__self__, *,
                 mode: pulumi.Input[str]):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconvert-jobtemplate-accelerationsettings.html
        :param pulumi.Input[str] mode: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconvert-jobtemplate-accelerationsettings.html#cfn-mediaconvert-jobtemplate-accelerationsettings-mode
        """
        pulumi.set(__self__, "mode", mode)

    @property
    @pulumi.getter(name="Mode")
    def mode(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconvert-jobtemplate-accelerationsettings.html#cfn-mediaconvert-jobtemplate-accelerationsettings-mode
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "mode", value)


@pulumi.input_type
class JobTemplateHopDestinationArgs:
    def __init__(__self__, *,
                 priority: Optional[pulumi.Input[int]] = None,
                 queue: Optional[pulumi.Input[str]] = None,
                 wait_minutes: Optional[pulumi.Input[int]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconvert-jobtemplate-hopdestination.html
        :param pulumi.Input[int] priority: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconvert-jobtemplate-hopdestination.html#cfn-mediaconvert-jobtemplate-hopdestination-priority
        :param pulumi.Input[str] queue: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconvert-jobtemplate-hopdestination.html#cfn-mediaconvert-jobtemplate-hopdestination-queue
        :param pulumi.Input[int] wait_minutes: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconvert-jobtemplate-hopdestination.html#cfn-mediaconvert-jobtemplate-hopdestination-waitminutes
        """
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if queue is not None:
            pulumi.set(__self__, "queue", queue)
        if wait_minutes is not None:
            pulumi.set(__self__, "wait_minutes", wait_minutes)

    @property
    @pulumi.getter(name="Priority")
    def priority(self) -> Optional[pulumi.Input[int]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconvert-jobtemplate-hopdestination.html#cfn-mediaconvert-jobtemplate-hopdestination-priority
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="Queue")
    def queue(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconvert-jobtemplate-hopdestination.html#cfn-mediaconvert-jobtemplate-hopdestination-queue
        """
        return pulumi.get(self, "queue")

    @queue.setter
    def queue(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "queue", value)

    @property
    @pulumi.getter(name="WaitMinutes")
    def wait_minutes(self) -> Optional[pulumi.Input[int]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconvert-jobtemplate-hopdestination.html#cfn-mediaconvert-jobtemplate-hopdestination-waitminutes
        """
        return pulumi.get(self, "wait_minutes")

    @wait_minutes.setter
    def wait_minutes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "wait_minutes", value)


@pulumi.input_type
class JobTemplatePropertiesArgs:
    def __init__(__self__, *,
                 settings_json: pulumi.Input[Union[Any, str]],
                 acceleration_settings: Optional[pulumi.Input['JobTemplateAccelerationSettingsArgs']] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 hop_destinations: Optional[pulumi.Input[Sequence[pulumi.Input['JobTemplateHopDestinationArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 queue: Optional[pulumi.Input[str]] = None,
                 status_update_interval: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Union[Any, str]]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html
        :param pulumi.Input[Union[Any, str]] settings_json: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-settingsjson
        :param pulumi.Input['JobTemplateAccelerationSettingsArgs'] acceleration_settings: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-accelerationsettings
        :param pulumi.Input[str] category: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-category
        :param pulumi.Input[str] description: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-description
        :param pulumi.Input[Sequence[pulumi.Input['JobTemplateHopDestinationArgs']]] hop_destinations: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-hopdestinations
        :param pulumi.Input[str] name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-name
        :param pulumi.Input[int] priority: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-priority
        :param pulumi.Input[str] queue: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-queue
        :param pulumi.Input[str] status_update_interval: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-statusupdateinterval
        :param pulumi.Input[Union[Any, str]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-tags
        """
        pulumi.set(__self__, "settings_json", settings_json)
        if acceleration_settings is not None:
            pulumi.set(__self__, "acceleration_settings", acceleration_settings)
        if category is not None:
            pulumi.set(__self__, "category", category)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if hop_destinations is not None:
            pulumi.set(__self__, "hop_destinations", hop_destinations)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if queue is not None:
            pulumi.set(__self__, "queue", queue)
        if status_update_interval is not None:
            pulumi.set(__self__, "status_update_interval", status_update_interval)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="SettingsJson")
    def settings_json(self) -> pulumi.Input[Union[Any, str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-settingsjson
        """
        return pulumi.get(self, "settings_json")

    @settings_json.setter
    def settings_json(self, value: pulumi.Input[Union[Any, str]]):
        pulumi.set(self, "settings_json", value)

    @property
    @pulumi.getter(name="AccelerationSettings")
    def acceleration_settings(self) -> Optional[pulumi.Input['JobTemplateAccelerationSettingsArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-accelerationsettings
        """
        return pulumi.get(self, "acceleration_settings")

    @acceleration_settings.setter
    def acceleration_settings(self, value: Optional[pulumi.Input['JobTemplateAccelerationSettingsArgs']]):
        pulumi.set(self, "acceleration_settings", value)

    @property
    @pulumi.getter(name="Category")
    def category(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-category
        """
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "category", value)

    @property
    @pulumi.getter(name="Description")
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="HopDestinations")
    def hop_destinations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['JobTemplateHopDestinationArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-hopdestinations
        """
        return pulumi.get(self, "hop_destinations")

    @hop_destinations.setter
    def hop_destinations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['JobTemplateHopDestinationArgs']]]]):
        pulumi.set(self, "hop_destinations", value)

    @property
    @pulumi.getter(name="Name")
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="Priority")
    def priority(self) -> Optional[pulumi.Input[int]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-priority
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="Queue")
    def queue(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-queue
        """
        return pulumi.get(self, "queue")

    @queue.setter
    def queue(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "queue", value)

    @property
    @pulumi.getter(name="StatusUpdateInterval")
    def status_update_interval(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-statusupdateinterval
        """
        return pulumi.get(self, "status_update_interval")

    @status_update_interval.setter
    def status_update_interval(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status_update_interval", value)

    @property
    @pulumi.getter(name="Tags")
    def tags(self) -> Optional[pulumi.Input[Union[Any, str]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Union[Any, str]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class PresetPropertiesArgs:
    def __init__(__self__, *,
                 settings_json: pulumi.Input[Union[Any, str]],
                 category: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Union[Any, str]]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-preset.html
        :param pulumi.Input[Union[Any, str]] settings_json: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-preset.html#cfn-mediaconvert-preset-settingsjson
        :param pulumi.Input[str] category: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-preset.html#cfn-mediaconvert-preset-category
        :param pulumi.Input[str] description: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-preset.html#cfn-mediaconvert-preset-description
        :param pulumi.Input[str] name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-preset.html#cfn-mediaconvert-preset-name
        :param pulumi.Input[Union[Any, str]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-preset.html#cfn-mediaconvert-preset-tags
        """
        pulumi.set(__self__, "settings_json", settings_json)
        if category is not None:
            pulumi.set(__self__, "category", category)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="SettingsJson")
    def settings_json(self) -> pulumi.Input[Union[Any, str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-preset.html#cfn-mediaconvert-preset-settingsjson
        """
        return pulumi.get(self, "settings_json")

    @settings_json.setter
    def settings_json(self, value: pulumi.Input[Union[Any, str]]):
        pulumi.set(self, "settings_json", value)

    @property
    @pulumi.getter(name="Category")
    def category(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-preset.html#cfn-mediaconvert-preset-category
        """
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "category", value)

    @property
    @pulumi.getter(name="Description")
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-preset.html#cfn-mediaconvert-preset-description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="Name")
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-preset.html#cfn-mediaconvert-preset-name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="Tags")
    def tags(self) -> Optional[pulumi.Input[Union[Any, str]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-preset.html#cfn-mediaconvert-preset-tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Union[Any, str]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class QueuePropertiesArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pricing_plan: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Union[Any, str]]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-queue.html
        :param pulumi.Input[str] description: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-queue.html#cfn-mediaconvert-queue-description
        :param pulumi.Input[str] name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-queue.html#cfn-mediaconvert-queue-name
        :param pulumi.Input[str] pricing_plan: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-queue.html#cfn-mediaconvert-queue-pricingplan
        :param pulumi.Input[str] status: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-queue.html#cfn-mediaconvert-queue-status
        :param pulumi.Input[Union[Any, str]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-queue.html#cfn-mediaconvert-queue-tags
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if pricing_plan is not None:
            pulumi.set(__self__, "pricing_plan", pricing_plan)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="Description")
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-queue.html#cfn-mediaconvert-queue-description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="Name")
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-queue.html#cfn-mediaconvert-queue-name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="PricingPlan")
    def pricing_plan(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-queue.html#cfn-mediaconvert-queue-pricingplan
        """
        return pulumi.get(self, "pricing_plan")

    @pricing_plan.setter
    def pricing_plan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pricing_plan", value)

    @property
    @pulumi.getter(name="Status")
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-queue.html#cfn-mediaconvert-queue-status
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="Tags")
    def tags(self) -> Optional[pulumi.Input[Union[Any, str]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-queue.html#cfn-mediaconvert-queue-tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Union[Any, str]]]):
        pulumi.set(self, "tags", value)


