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
from ._inputs import *

__all__ = ['EventTypeArgs', 'EventType']

@pulumi.input_type
class EventTypeArgs:
    def __init__(__self__, *,
                 entity_types: pulumi.Input[Sequence[pulumi.Input['EventTypeEntityTypeArgs']]],
                 event_variables: pulumi.Input[Sequence[pulumi.Input['EventTypeEventVariableArgs']]],
                 labels: pulumi.Input[Sequence[pulumi.Input['EventTypeLabelArgs']]],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['EventTypeTagArgs']]]] = None):
        """
        The set of arguments for constructing a EventType resource.
        :param pulumi.Input[str] description: The description of the event type.
        :param pulumi.Input[str] name: The name for the event type
        :param pulumi.Input[Sequence[pulumi.Input['EventTypeTagArgs']]] tags: Tags associated with this event type.
        """
        pulumi.set(__self__, "entity_types", entity_types)
        pulumi.set(__self__, "event_variables", event_variables)
        pulumi.set(__self__, "labels", labels)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="entityTypes")
    def entity_types(self) -> pulumi.Input[Sequence[pulumi.Input['EventTypeEntityTypeArgs']]]:
        return pulumi.get(self, "entity_types")

    @entity_types.setter
    def entity_types(self, value: pulumi.Input[Sequence[pulumi.Input['EventTypeEntityTypeArgs']]]):
        pulumi.set(self, "entity_types", value)

    @property
    @pulumi.getter(name="eventVariables")
    def event_variables(self) -> pulumi.Input[Sequence[pulumi.Input['EventTypeEventVariableArgs']]]:
        return pulumi.get(self, "event_variables")

    @event_variables.setter
    def event_variables(self, value: pulumi.Input[Sequence[pulumi.Input['EventTypeEventVariableArgs']]]):
        pulumi.set(self, "event_variables", value)

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Input[Sequence[pulumi.Input['EventTypeLabelArgs']]]:
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: pulumi.Input[Sequence[pulumi.Input['EventTypeLabelArgs']]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the event type.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name for the event type
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EventTypeTagArgs']]]]:
        """
        Tags associated with this event type.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EventTypeTagArgs']]]]):
        pulumi.set(self, "tags", value)


class EventType(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 entity_types: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EventTypeEntityTypeArgs']]]]] = None,
                 event_variables: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EventTypeEventVariableArgs']]]]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EventTypeLabelArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EventTypeTagArgs']]]]] = None,
                 __props__=None):
        """
        A resource schema for an EventType in Amazon Fraud Detector.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the event type.
        :param pulumi.Input[str] name: The name for the event type
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EventTypeTagArgs']]]] tags: Tags associated with this event type.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EventTypeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A resource schema for an EventType in Amazon Fraud Detector.

        :param str resource_name: The name of the resource.
        :param EventTypeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EventTypeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 entity_types: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EventTypeEntityTypeArgs']]]]] = None,
                 event_variables: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EventTypeEventVariableArgs']]]]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EventTypeLabelArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EventTypeTagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EventTypeArgs.__new__(EventTypeArgs)

            __props__.__dict__["description"] = description
            if entity_types is None and not opts.urn:
                raise TypeError("Missing required property 'entity_types'")
            __props__.__dict__["entity_types"] = entity_types
            if event_variables is None and not opts.urn:
                raise TypeError("Missing required property 'event_variables'")
            __props__.__dict__["event_variables"] = event_variables
            if labels is None and not opts.urn:
                raise TypeError("Missing required property 'labels'")
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["created_time"] = None
            __props__.__dict__["last_updated_time"] = None
        super(EventType, __self__).__init__(
            'aws-native:frauddetector:EventType',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'EventType':
        """
        Get an existing EventType resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = EventTypeArgs.__new__(EventTypeArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["created_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["entity_types"] = None
        __props__.__dict__["event_variables"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["last_updated_time"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["tags"] = None
        return EventType(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the event type.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> pulumi.Output[str]:
        """
        The time when the event type was created.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the event type.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="entityTypes")
    def entity_types(self) -> pulumi.Output[Sequence['outputs.EventTypeEntityType']]:
        return pulumi.get(self, "entity_types")

    @property
    @pulumi.getter(name="eventVariables")
    def event_variables(self) -> pulumi.Output[Sequence['outputs.EventTypeEventVariable']]:
        return pulumi.get(self, "event_variables")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Sequence['outputs.EventTypeLabel']]:
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="lastUpdatedTime")
    def last_updated_time(self) -> pulumi.Output[str]:
        """
        The time when the event type was last updated.
        """
        return pulumi.get(self, "last_updated_time")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name for the event type
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.EventTypeTag']]]:
        """
        Tags associated with this event type.
        """
        return pulumi.get(self, "tags")

