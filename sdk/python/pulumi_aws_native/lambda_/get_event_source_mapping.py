# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'GetEventSourceMappingResult',
    'AwaitableGetEventSourceMappingResult',
    'get_event_source_mapping',
    'get_event_source_mapping_output',
]

@pulumi.output_type
class GetEventSourceMappingResult:
    def __init__(__self__, batch_size=None, bisect_batch_on_function_error=None, destination_config=None, enabled=None, filter_criteria=None, function_name=None, function_response_types=None, id=None, maximum_batching_window_in_seconds=None, maximum_record_age_in_seconds=None, maximum_retry_attempts=None, parallelization_factor=None, queues=None, source_access_configurations=None, starting_position_timestamp=None, topics=None, tumbling_window_in_seconds=None):
        if batch_size and not isinstance(batch_size, int):
            raise TypeError("Expected argument 'batch_size' to be a int")
        pulumi.set(__self__, "batch_size", batch_size)
        if bisect_batch_on_function_error and not isinstance(bisect_batch_on_function_error, bool):
            raise TypeError("Expected argument 'bisect_batch_on_function_error' to be a bool")
        pulumi.set(__self__, "bisect_batch_on_function_error", bisect_batch_on_function_error)
        if destination_config and not isinstance(destination_config, dict):
            raise TypeError("Expected argument 'destination_config' to be a dict")
        pulumi.set(__self__, "destination_config", destination_config)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if filter_criteria and not isinstance(filter_criteria, dict):
            raise TypeError("Expected argument 'filter_criteria' to be a dict")
        pulumi.set(__self__, "filter_criteria", filter_criteria)
        if function_name and not isinstance(function_name, str):
            raise TypeError("Expected argument 'function_name' to be a str")
        pulumi.set(__self__, "function_name", function_name)
        if function_response_types and not isinstance(function_response_types, list):
            raise TypeError("Expected argument 'function_response_types' to be a list")
        pulumi.set(__self__, "function_response_types", function_response_types)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if maximum_batching_window_in_seconds and not isinstance(maximum_batching_window_in_seconds, int):
            raise TypeError("Expected argument 'maximum_batching_window_in_seconds' to be a int")
        pulumi.set(__self__, "maximum_batching_window_in_seconds", maximum_batching_window_in_seconds)
        if maximum_record_age_in_seconds and not isinstance(maximum_record_age_in_seconds, int):
            raise TypeError("Expected argument 'maximum_record_age_in_seconds' to be a int")
        pulumi.set(__self__, "maximum_record_age_in_seconds", maximum_record_age_in_seconds)
        if maximum_retry_attempts and not isinstance(maximum_retry_attempts, int):
            raise TypeError("Expected argument 'maximum_retry_attempts' to be a int")
        pulumi.set(__self__, "maximum_retry_attempts", maximum_retry_attempts)
        if parallelization_factor and not isinstance(parallelization_factor, int):
            raise TypeError("Expected argument 'parallelization_factor' to be a int")
        pulumi.set(__self__, "parallelization_factor", parallelization_factor)
        if queues and not isinstance(queues, list):
            raise TypeError("Expected argument 'queues' to be a list")
        pulumi.set(__self__, "queues", queues)
        if source_access_configurations and not isinstance(source_access_configurations, list):
            raise TypeError("Expected argument 'source_access_configurations' to be a list")
        pulumi.set(__self__, "source_access_configurations", source_access_configurations)
        if starting_position_timestamp and not isinstance(starting_position_timestamp, float):
            raise TypeError("Expected argument 'starting_position_timestamp' to be a float")
        pulumi.set(__self__, "starting_position_timestamp", starting_position_timestamp)
        if topics and not isinstance(topics, list):
            raise TypeError("Expected argument 'topics' to be a list")
        pulumi.set(__self__, "topics", topics)
        if tumbling_window_in_seconds and not isinstance(tumbling_window_in_seconds, int):
            raise TypeError("Expected argument 'tumbling_window_in_seconds' to be a int")
        pulumi.set(__self__, "tumbling_window_in_seconds", tumbling_window_in_seconds)

    @property
    @pulumi.getter(name="batchSize")
    def batch_size(self) -> Optional[int]:
        """
        The maximum number of items to retrieve in a single batch.
        """
        return pulumi.get(self, "batch_size")

    @property
    @pulumi.getter(name="bisectBatchOnFunctionError")
    def bisect_batch_on_function_error(self) -> Optional[bool]:
        """
        (Streams) If the function returns an error, split the batch in two and retry.
        """
        return pulumi.get(self, "bisect_batch_on_function_error")

    @property
    @pulumi.getter(name="destinationConfig")
    def destination_config(self) -> Optional['outputs.EventSourceMappingDestinationConfig']:
        """
        (Streams) An Amazon SQS queue or Amazon SNS topic destination for discarded records.
        """
        return pulumi.get(self, "destination_config")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        Disables the event source mapping to pause polling and invocation.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="filterCriteria")
    def filter_criteria(self) -> Optional['outputs.FilterCriteriaProperties']:
        """
        The filter criteria to control event filtering.
        """
        return pulumi.get(self, "filter_criteria")

    @property
    @pulumi.getter(name="functionName")
    def function_name(self) -> Optional[str]:
        """
        The name of the Lambda function.
        """
        return pulumi.get(self, "function_name")

    @property
    @pulumi.getter(name="functionResponseTypes")
    def function_response_types(self) -> Optional[Sequence['EventSourceMappingFunctionResponseTypesItem']]:
        """
        (Streams) A list of response types supported by the function.
        """
        return pulumi.get(self, "function_response_types")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        Event Source Mapping Identifier UUID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="maximumBatchingWindowInSeconds")
    def maximum_batching_window_in_seconds(self) -> Optional[int]:
        """
        (Streams) The maximum amount of time to gather records before invoking the function, in seconds.
        """
        return pulumi.get(self, "maximum_batching_window_in_seconds")

    @property
    @pulumi.getter(name="maximumRecordAgeInSeconds")
    def maximum_record_age_in_seconds(self) -> Optional[int]:
        """
        (Streams) The maximum age of a record that Lambda sends to a function for processing.
        """
        return pulumi.get(self, "maximum_record_age_in_seconds")

    @property
    @pulumi.getter(name="maximumRetryAttempts")
    def maximum_retry_attempts(self) -> Optional[int]:
        """
        (Streams) The maximum number of times to retry when the function returns an error.
        """
        return pulumi.get(self, "maximum_retry_attempts")

    @property
    @pulumi.getter(name="parallelizationFactor")
    def parallelization_factor(self) -> Optional[int]:
        """
        (Streams) The number of batches to process from each shard concurrently.
        """
        return pulumi.get(self, "parallelization_factor")

    @property
    @pulumi.getter
    def queues(self) -> Optional[Sequence[str]]:
        """
        (ActiveMQ) A list of ActiveMQ queues.
        """
        return pulumi.get(self, "queues")

    @property
    @pulumi.getter(name="sourceAccessConfigurations")
    def source_access_configurations(self) -> Optional[Sequence['outputs.EventSourceMappingSourceAccessConfiguration']]:
        """
        A list of SourceAccessConfiguration.
        """
        return pulumi.get(self, "source_access_configurations")

    @property
    @pulumi.getter(name="startingPositionTimestamp")
    def starting_position_timestamp(self) -> Optional[float]:
        """
        With StartingPosition set to AT_TIMESTAMP, the time from which to start reading, in Unix time seconds.
        """
        return pulumi.get(self, "starting_position_timestamp")

    @property
    @pulumi.getter
    def topics(self) -> Optional[Sequence[str]]:
        """
        (Kafka) A list of Kafka topics.
        """
        return pulumi.get(self, "topics")

    @property
    @pulumi.getter(name="tumblingWindowInSeconds")
    def tumbling_window_in_seconds(self) -> Optional[int]:
        """
        (Streams) Tumbling window (non-overlapping time window) duration to perform aggregations.
        """
        return pulumi.get(self, "tumbling_window_in_seconds")


class AwaitableGetEventSourceMappingResult(GetEventSourceMappingResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEventSourceMappingResult(
            batch_size=self.batch_size,
            bisect_batch_on_function_error=self.bisect_batch_on_function_error,
            destination_config=self.destination_config,
            enabled=self.enabled,
            filter_criteria=self.filter_criteria,
            function_name=self.function_name,
            function_response_types=self.function_response_types,
            id=self.id,
            maximum_batching_window_in_seconds=self.maximum_batching_window_in_seconds,
            maximum_record_age_in_seconds=self.maximum_record_age_in_seconds,
            maximum_retry_attempts=self.maximum_retry_attempts,
            parallelization_factor=self.parallelization_factor,
            queues=self.queues,
            source_access_configurations=self.source_access_configurations,
            starting_position_timestamp=self.starting_position_timestamp,
            topics=self.topics,
            tumbling_window_in_seconds=self.tumbling_window_in_seconds)


def get_event_source_mapping(id: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEventSourceMappingResult:
    """
    Resource Type definition for AWS::Lambda::EventSourceMapping


    :param str id: Event Source Mapping Identifier UUID.
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:lambda:getEventSourceMapping', __args__, opts=opts, typ=GetEventSourceMappingResult).value

    return AwaitableGetEventSourceMappingResult(
        batch_size=__ret__.batch_size,
        bisect_batch_on_function_error=__ret__.bisect_batch_on_function_error,
        destination_config=__ret__.destination_config,
        enabled=__ret__.enabled,
        filter_criteria=__ret__.filter_criteria,
        function_name=__ret__.function_name,
        function_response_types=__ret__.function_response_types,
        id=__ret__.id,
        maximum_batching_window_in_seconds=__ret__.maximum_batching_window_in_seconds,
        maximum_record_age_in_seconds=__ret__.maximum_record_age_in_seconds,
        maximum_retry_attempts=__ret__.maximum_retry_attempts,
        parallelization_factor=__ret__.parallelization_factor,
        queues=__ret__.queues,
        source_access_configurations=__ret__.source_access_configurations,
        starting_position_timestamp=__ret__.starting_position_timestamp,
        topics=__ret__.topics,
        tumbling_window_in_seconds=__ret__.tumbling_window_in_seconds)


@_utilities.lift_output_func(get_event_source_mapping)
def get_event_source_mapping_output(id: Optional[pulumi.Input[str]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEventSourceMappingResult]:
    """
    Resource Type definition for AWS::Lambda::EventSourceMapping


    :param str id: Event Source Mapping Identifier UUID.
    """
    ...
