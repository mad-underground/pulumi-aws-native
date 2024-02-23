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
    'DataIntegrationFileConfigurationArgs',
    'DataIntegrationScheduleConfigArgs',
    'EventIntegrationEventFilterArgs',
]

@pulumi.input_type
class DataIntegrationFileConfigurationArgs:
    def __init__(__self__, *,
                 folders: pulumi.Input[Sequence[pulumi.Input[str]]],
                 filters: Optional[pulumi.Input[Mapping[str, pulumi.Input[Sequence[pulumi.Input[str]]]]]] = None):
        """
        The configuration for what files should be pulled from the source.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] folders: Identifiers for the source folders to pull all files from recursively.
        :param pulumi.Input[Mapping[str, pulumi.Input[Sequence[pulumi.Input[str]]]]] filters: Restrictions for what files should be pulled from the source.
        """
        pulumi.set(__self__, "folders", folders)
        if filters is not None:
            pulumi.set(__self__, "filters", filters)

    @property
    @pulumi.getter
    def folders(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Identifiers for the source folders to pull all files from recursively.
        """
        return pulumi.get(self, "folders")

    @folders.setter
    def folders(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "folders", value)

    @property
    @pulumi.getter
    def filters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[Sequence[pulumi.Input[str]]]]]]:
        """
        Restrictions for what files should be pulled from the source.
        """
        return pulumi.get(self, "filters")

    @filters.setter
    def filters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[Sequence[pulumi.Input[str]]]]]]):
        pulumi.set(self, "filters", value)


@pulumi.input_type
class DataIntegrationScheduleConfigArgs:
    def __init__(__self__, *,
                 schedule_expression: pulumi.Input[str],
                 first_execution_from: Optional[pulumi.Input[str]] = None,
                 object: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] schedule_expression: How often the data should be pulled from data source.
        :param pulumi.Input[str] first_execution_from: The start date for objects to import in the first flow run. Epoch or ISO timestamp format is supported.
        :param pulumi.Input[str] object: The name of the object to pull from the data source.
        """
        pulumi.set(__self__, "schedule_expression", schedule_expression)
        if first_execution_from is not None:
            pulumi.set(__self__, "first_execution_from", first_execution_from)
        if object is not None:
            pulumi.set(__self__, "object", object)

    @property
    @pulumi.getter(name="scheduleExpression")
    def schedule_expression(self) -> pulumi.Input[str]:
        """
        How often the data should be pulled from data source.
        """
        return pulumi.get(self, "schedule_expression")

    @schedule_expression.setter
    def schedule_expression(self, value: pulumi.Input[str]):
        pulumi.set(self, "schedule_expression", value)

    @property
    @pulumi.getter(name="firstExecutionFrom")
    def first_execution_from(self) -> Optional[pulumi.Input[str]]:
        """
        The start date for objects to import in the first flow run. Epoch or ISO timestamp format is supported.
        """
        return pulumi.get(self, "first_execution_from")

    @first_execution_from.setter
    def first_execution_from(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "first_execution_from", value)

    @property
    @pulumi.getter
    def object(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the object to pull from the data source.
        """
        return pulumi.get(self, "object")

    @object.setter
    def object(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "object", value)


@pulumi.input_type
class EventIntegrationEventFilterArgs:
    def __init__(__self__, *,
                 source: pulumi.Input[str]):
        """
        :param pulumi.Input[str] source: The source of the events.
        """
        pulumi.set(__self__, "source", source)

    @property
    @pulumi.getter
    def source(self) -> pulumi.Input[str]:
        """
        The source of the events.
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: pulumi.Input[str]):
        pulumi.set(self, "source", value)


