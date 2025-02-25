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
    'DetectorEntityType',
    'DetectorEventType',
    'DetectorEventVariable',
    'DetectorLabel',
    'DetectorModel',
    'DetectorOutcome',
    'DetectorRule',
    'DetectorTag',
    'EventTypeEntityType',
    'EventTypeEventVariable',
    'EventTypeLabel',
    'EventTypeTag',
]

@pulumi.output_type
class DetectorEntityType(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createdTime":
            suggest = "created_time"
        elif key == "lastUpdatedTime":
            suggest = "last_updated_time"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DetectorEntityType. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DetectorEntityType.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DetectorEntityType.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 arn: Optional[str] = None,
                 created_time: Optional[str] = None,
                 description: Optional[str] = None,
                 inline: Optional[bool] = None,
                 last_updated_time: Optional[str] = None,
                 name: Optional[str] = None,
                 tags: Optional[Sequence['outputs.DetectorTag']] = None):
        """
        :param str created_time: The time when the entity type was created.
        :param str description: The description.
        :param str last_updated_time: The time when the entity type was last updated.
        :param Sequence['DetectorTag'] tags: Tags associated with this entity type.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if created_time is not None:
            pulumi.set(__self__, "created_time", created_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if inline is not None:
            pulumi.set(__self__, "inline", inline)
        if last_updated_time is not None:
            pulumi.set(__self__, "last_updated_time", last_updated_time)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> Optional[str]:
        """
        The time when the entity type was created.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def inline(self) -> Optional[bool]:
        return pulumi.get(self, "inline")

    @property
    @pulumi.getter(name="lastUpdatedTime")
    def last_updated_time(self) -> Optional[str]:
        """
        The time when the entity type was last updated.
        """
        return pulumi.get(self, "last_updated_time")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.DetectorTag']]:
        """
        Tags associated with this entity type.
        """
        return pulumi.get(self, "tags")


@pulumi.output_type
class DetectorEventType(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createdTime":
            suggest = "created_time"
        elif key == "entityTypes":
            suggest = "entity_types"
        elif key == "eventVariables":
            suggest = "event_variables"
        elif key == "lastUpdatedTime":
            suggest = "last_updated_time"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DetectorEventType. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DetectorEventType.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DetectorEventType.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 arn: Optional[str] = None,
                 created_time: Optional[str] = None,
                 description: Optional[str] = None,
                 entity_types: Optional[Sequence['outputs.DetectorEntityType']] = None,
                 event_variables: Optional[Sequence['outputs.DetectorEventVariable']] = None,
                 inline: Optional[bool] = None,
                 labels: Optional[Sequence['outputs.DetectorLabel']] = None,
                 last_updated_time: Optional[str] = None,
                 name: Optional[str] = None,
                 tags: Optional[Sequence['outputs.DetectorTag']] = None):
        """
        :param str arn: The ARN of the event type.
        :param str created_time: The time when the event type was created.
        :param str description: The description of the event type.
        :param str last_updated_time: The time when the event type was last updated.
        :param str name: The name for the event type
        :param Sequence['DetectorTag'] tags: Tags associated with this event type.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if created_time is not None:
            pulumi.set(__self__, "created_time", created_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if entity_types is not None:
            pulumi.set(__self__, "entity_types", entity_types)
        if event_variables is not None:
            pulumi.set(__self__, "event_variables", event_variables)
        if inline is not None:
            pulumi.set(__self__, "inline", inline)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if last_updated_time is not None:
            pulumi.set(__self__, "last_updated_time", last_updated_time)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        The ARN of the event type.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> Optional[str]:
        """
        The time when the event type was created.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the event type.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="entityTypes")
    def entity_types(self) -> Optional[Sequence['outputs.DetectorEntityType']]:
        return pulumi.get(self, "entity_types")

    @property
    @pulumi.getter(name="eventVariables")
    def event_variables(self) -> Optional[Sequence['outputs.DetectorEventVariable']]:
        return pulumi.get(self, "event_variables")

    @property
    @pulumi.getter
    def inline(self) -> Optional[bool]:
        return pulumi.get(self, "inline")

    @property
    @pulumi.getter
    def labels(self) -> Optional[Sequence['outputs.DetectorLabel']]:
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="lastUpdatedTime")
    def last_updated_time(self) -> Optional[str]:
        """
        The time when the event type was last updated.
        """
        return pulumi.get(self, "last_updated_time")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name for the event type
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.DetectorTag']]:
        """
        Tags associated with this event type.
        """
        return pulumi.get(self, "tags")


@pulumi.output_type
class DetectorEventVariable(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createdTime":
            suggest = "created_time"
        elif key == "dataSource":
            suggest = "data_source"
        elif key == "dataType":
            suggest = "data_type"
        elif key == "defaultValue":
            suggest = "default_value"
        elif key == "lastUpdatedTime":
            suggest = "last_updated_time"
        elif key == "variableType":
            suggest = "variable_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DetectorEventVariable. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DetectorEventVariable.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DetectorEventVariable.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 arn: Optional[str] = None,
                 created_time: Optional[str] = None,
                 data_source: Optional['DetectorEventVariableDataSource'] = None,
                 data_type: Optional['DetectorEventVariableDataType'] = None,
                 default_value: Optional[str] = None,
                 description: Optional[str] = None,
                 inline: Optional[bool] = None,
                 last_updated_time: Optional[str] = None,
                 name: Optional[str] = None,
                 tags: Optional[Sequence['outputs.DetectorTag']] = None,
                 variable_type: Optional['DetectorEventVariableVariableType'] = None):
        """
        :param str created_time: The time when the event variable was created.
        :param str description: The description.
        :param str last_updated_time: The time when the event variable was last updated.
        :param Sequence['DetectorTag'] tags: Tags associated with this event variable.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if created_time is not None:
            pulumi.set(__self__, "created_time", created_time)
        if data_source is not None:
            pulumi.set(__self__, "data_source", data_source)
        if data_type is not None:
            pulumi.set(__self__, "data_type", data_type)
        if default_value is not None:
            pulumi.set(__self__, "default_value", default_value)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if inline is not None:
            pulumi.set(__self__, "inline", inline)
        if last_updated_time is not None:
            pulumi.set(__self__, "last_updated_time", last_updated_time)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if variable_type is not None:
            pulumi.set(__self__, "variable_type", variable_type)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> Optional[str]:
        """
        The time when the event variable was created.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter(name="dataSource")
    def data_source(self) -> Optional['DetectorEventVariableDataSource']:
        return pulumi.get(self, "data_source")

    @property
    @pulumi.getter(name="dataType")
    def data_type(self) -> Optional['DetectorEventVariableDataType']:
        return pulumi.get(self, "data_type")

    @property
    @pulumi.getter(name="defaultValue")
    def default_value(self) -> Optional[str]:
        return pulumi.get(self, "default_value")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def inline(self) -> Optional[bool]:
        return pulumi.get(self, "inline")

    @property
    @pulumi.getter(name="lastUpdatedTime")
    def last_updated_time(self) -> Optional[str]:
        """
        The time when the event variable was last updated.
        """
        return pulumi.get(self, "last_updated_time")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.DetectorTag']]:
        """
        Tags associated with this event variable.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="variableType")
    def variable_type(self) -> Optional['DetectorEventVariableVariableType']:
        return pulumi.get(self, "variable_type")


@pulumi.output_type
class DetectorLabel(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createdTime":
            suggest = "created_time"
        elif key == "lastUpdatedTime":
            suggest = "last_updated_time"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DetectorLabel. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DetectorLabel.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DetectorLabel.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 arn: Optional[str] = None,
                 created_time: Optional[str] = None,
                 description: Optional[str] = None,
                 inline: Optional[bool] = None,
                 last_updated_time: Optional[str] = None,
                 name: Optional[str] = None,
                 tags: Optional[Sequence['outputs.DetectorTag']] = None):
        """
        :param str created_time: The time when the label was created.
        :param str description: The description.
        :param str last_updated_time: The time when the label was last updated.
        :param Sequence['DetectorTag'] tags: Tags associated with this label.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if created_time is not None:
            pulumi.set(__self__, "created_time", created_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if inline is not None:
            pulumi.set(__self__, "inline", inline)
        if last_updated_time is not None:
            pulumi.set(__self__, "last_updated_time", last_updated_time)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> Optional[str]:
        """
        The time when the label was created.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def inline(self) -> Optional[bool]:
        return pulumi.get(self, "inline")

    @property
    @pulumi.getter(name="lastUpdatedTime")
    def last_updated_time(self) -> Optional[str]:
        """
        The time when the label was last updated.
        """
        return pulumi.get(self, "last_updated_time")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.DetectorTag']]:
        """
        Tags associated with this label.
        """
        return pulumi.get(self, "tags")


@pulumi.output_type
class DetectorModel(dict):
    """
    A model to associate with a detector.
    """
    def __init__(__self__, *,
                 arn: Optional[str] = None):
        """
        A model to associate with a detector.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")


@pulumi.output_type
class DetectorOutcome(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createdTime":
            suggest = "created_time"
        elif key == "lastUpdatedTime":
            suggest = "last_updated_time"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DetectorOutcome. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DetectorOutcome.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DetectorOutcome.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 arn: Optional[str] = None,
                 created_time: Optional[str] = None,
                 description: Optional[str] = None,
                 inline: Optional[bool] = None,
                 last_updated_time: Optional[str] = None,
                 name: Optional[str] = None,
                 tags: Optional[Sequence['outputs.DetectorTag']] = None):
        """
        :param str created_time: The time when the outcome was created.
        :param str description: The description.
        :param str last_updated_time: The time when the outcome was last updated.
        :param Sequence['DetectorTag'] tags: Tags associated with this outcome.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if created_time is not None:
            pulumi.set(__self__, "created_time", created_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if inline is not None:
            pulumi.set(__self__, "inline", inline)
        if last_updated_time is not None:
            pulumi.set(__self__, "last_updated_time", last_updated_time)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> Optional[str]:
        """
        The time when the outcome was created.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def inline(self) -> Optional[bool]:
        return pulumi.get(self, "inline")

    @property
    @pulumi.getter(name="lastUpdatedTime")
    def last_updated_time(self) -> Optional[str]:
        """
        The time when the outcome was last updated.
        """
        return pulumi.get(self, "last_updated_time")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.DetectorTag']]:
        """
        Tags associated with this outcome.
        """
        return pulumi.get(self, "tags")


@pulumi.output_type
class DetectorRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createdTime":
            suggest = "created_time"
        elif key == "detectorId":
            suggest = "detector_id"
        elif key == "lastUpdatedTime":
            suggest = "last_updated_time"
        elif key == "ruleId":
            suggest = "rule_id"
        elif key == "ruleVersion":
            suggest = "rule_version"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DetectorRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DetectorRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DetectorRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 arn: Optional[str] = None,
                 created_time: Optional[str] = None,
                 description: Optional[str] = None,
                 detector_id: Optional[str] = None,
                 expression: Optional[str] = None,
                 language: Optional['DetectorRuleLanguage'] = None,
                 last_updated_time: Optional[str] = None,
                 outcomes: Optional[Sequence['outputs.DetectorOutcome']] = None,
                 rule_id: Optional[str] = None,
                 rule_version: Optional[str] = None,
                 tags: Optional[Sequence['outputs.DetectorTag']] = None):
        """
        :param str created_time: The time when the event type was created.
        :param str description: The description.
        :param str last_updated_time: The time when the event type was last updated.
        :param Sequence['DetectorTag'] tags: Tags associated with this event type.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if created_time is not None:
            pulumi.set(__self__, "created_time", created_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if detector_id is not None:
            pulumi.set(__self__, "detector_id", detector_id)
        if expression is not None:
            pulumi.set(__self__, "expression", expression)
        if language is not None:
            pulumi.set(__self__, "language", language)
        if last_updated_time is not None:
            pulumi.set(__self__, "last_updated_time", last_updated_time)
        if outcomes is not None:
            pulumi.set(__self__, "outcomes", outcomes)
        if rule_id is not None:
            pulumi.set(__self__, "rule_id", rule_id)
        if rule_version is not None:
            pulumi.set(__self__, "rule_version", rule_version)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> Optional[str]:
        """
        The time when the event type was created.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="detectorId")
    def detector_id(self) -> Optional[str]:
        return pulumi.get(self, "detector_id")

    @property
    @pulumi.getter
    def expression(self) -> Optional[str]:
        return pulumi.get(self, "expression")

    @property
    @pulumi.getter
    def language(self) -> Optional['DetectorRuleLanguage']:
        return pulumi.get(self, "language")

    @property
    @pulumi.getter(name="lastUpdatedTime")
    def last_updated_time(self) -> Optional[str]:
        """
        The time when the event type was last updated.
        """
        return pulumi.get(self, "last_updated_time")

    @property
    @pulumi.getter
    def outcomes(self) -> Optional[Sequence['outputs.DetectorOutcome']]:
        return pulumi.get(self, "outcomes")

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> Optional[str]:
        return pulumi.get(self, "rule_id")

    @property
    @pulumi.getter(name="ruleVersion")
    def rule_version(self) -> Optional[str]:
        return pulumi.get(self, "rule_version")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.DetectorTag']]:
        """
        Tags associated with this event type.
        """
        return pulumi.get(self, "tags")


@pulumi.output_type
class DetectorTag(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


@pulumi.output_type
class EventTypeEntityType(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createdTime":
            suggest = "created_time"
        elif key == "lastUpdatedTime":
            suggest = "last_updated_time"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in EventTypeEntityType. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        EventTypeEntityType.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        EventTypeEntityType.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 arn: Optional[str] = None,
                 created_time: Optional[str] = None,
                 description: Optional[str] = None,
                 inline: Optional[bool] = None,
                 last_updated_time: Optional[str] = None,
                 name: Optional[str] = None,
                 tags: Optional[Sequence['outputs.EventTypeTag']] = None):
        """
        :param str created_time: The time when the event type was created.
        :param str description: The description.
        :param str last_updated_time: The time when the event type was last updated.
        :param Sequence['EventTypeTag'] tags: Tags associated with this event type.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if created_time is not None:
            pulumi.set(__self__, "created_time", created_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if inline is not None:
            pulumi.set(__self__, "inline", inline)
        if last_updated_time is not None:
            pulumi.set(__self__, "last_updated_time", last_updated_time)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> Optional[str]:
        """
        The time when the event type was created.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def inline(self) -> Optional[bool]:
        return pulumi.get(self, "inline")

    @property
    @pulumi.getter(name="lastUpdatedTime")
    def last_updated_time(self) -> Optional[str]:
        """
        The time when the event type was last updated.
        """
        return pulumi.get(self, "last_updated_time")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.EventTypeTag']]:
        """
        Tags associated with this event type.
        """
        return pulumi.get(self, "tags")


@pulumi.output_type
class EventTypeEventVariable(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createdTime":
            suggest = "created_time"
        elif key == "dataSource":
            suggest = "data_source"
        elif key == "dataType":
            suggest = "data_type"
        elif key == "defaultValue":
            suggest = "default_value"
        elif key == "lastUpdatedTime":
            suggest = "last_updated_time"
        elif key == "variableType":
            suggest = "variable_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in EventTypeEventVariable. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        EventTypeEventVariable.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        EventTypeEventVariable.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 arn: Optional[str] = None,
                 created_time: Optional[str] = None,
                 data_source: Optional['EventTypeEventVariableDataSource'] = None,
                 data_type: Optional['EventTypeEventVariableDataType'] = None,
                 default_value: Optional[str] = None,
                 description: Optional[str] = None,
                 inline: Optional[bool] = None,
                 last_updated_time: Optional[str] = None,
                 name: Optional[str] = None,
                 tags: Optional[Sequence['outputs.EventTypeTag']] = None,
                 variable_type: Optional['EventTypeEventVariableVariableType'] = None):
        """
        :param str created_time: The time when the event type was created.
        :param str description: The description.
        :param str last_updated_time: The time when the event type was last updated.
        :param Sequence['EventTypeTag'] tags: Tags associated with this event type.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if created_time is not None:
            pulumi.set(__self__, "created_time", created_time)
        if data_source is not None:
            pulumi.set(__self__, "data_source", data_source)
        if data_type is not None:
            pulumi.set(__self__, "data_type", data_type)
        if default_value is not None:
            pulumi.set(__self__, "default_value", default_value)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if inline is not None:
            pulumi.set(__self__, "inline", inline)
        if last_updated_time is not None:
            pulumi.set(__self__, "last_updated_time", last_updated_time)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if variable_type is not None:
            pulumi.set(__self__, "variable_type", variable_type)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> Optional[str]:
        """
        The time when the event type was created.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter(name="dataSource")
    def data_source(self) -> Optional['EventTypeEventVariableDataSource']:
        return pulumi.get(self, "data_source")

    @property
    @pulumi.getter(name="dataType")
    def data_type(self) -> Optional['EventTypeEventVariableDataType']:
        return pulumi.get(self, "data_type")

    @property
    @pulumi.getter(name="defaultValue")
    def default_value(self) -> Optional[str]:
        return pulumi.get(self, "default_value")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def inline(self) -> Optional[bool]:
        return pulumi.get(self, "inline")

    @property
    @pulumi.getter(name="lastUpdatedTime")
    def last_updated_time(self) -> Optional[str]:
        """
        The time when the event type was last updated.
        """
        return pulumi.get(self, "last_updated_time")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.EventTypeTag']]:
        """
        Tags associated with this event type.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="variableType")
    def variable_type(self) -> Optional['EventTypeEventVariableVariableType']:
        return pulumi.get(self, "variable_type")


@pulumi.output_type
class EventTypeLabel(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createdTime":
            suggest = "created_time"
        elif key == "lastUpdatedTime":
            suggest = "last_updated_time"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in EventTypeLabel. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        EventTypeLabel.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        EventTypeLabel.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 arn: Optional[str] = None,
                 created_time: Optional[str] = None,
                 description: Optional[str] = None,
                 inline: Optional[bool] = None,
                 last_updated_time: Optional[str] = None,
                 name: Optional[str] = None,
                 tags: Optional[Sequence['outputs.EventTypeTag']] = None):
        """
        :param str created_time: The time when the event type was created.
        :param str description: The description.
        :param str last_updated_time: The time when the event type was last updated.
        :param Sequence['EventTypeTag'] tags: Tags associated with this event type.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if created_time is not None:
            pulumi.set(__self__, "created_time", created_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if inline is not None:
            pulumi.set(__self__, "inline", inline)
        if last_updated_time is not None:
            pulumi.set(__self__, "last_updated_time", last_updated_time)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> Optional[str]:
        """
        The time when the event type was created.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def inline(self) -> Optional[bool]:
        return pulumi.get(self, "inline")

    @property
    @pulumi.getter(name="lastUpdatedTime")
    def last_updated_time(self) -> Optional[str]:
        """
        The time when the event type was last updated.
        """
        return pulumi.get(self, "last_updated_time")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.EventTypeTag']]:
        """
        Tags associated with this event type.
        """
        return pulumi.get(self, "tags")


@pulumi.output_type
class EventTypeTag(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


