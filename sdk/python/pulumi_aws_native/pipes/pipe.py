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

__all__ = ['PipeArgs', 'Pipe']

@pulumi.input_type
class PipeArgs:
    def __init__(__self__, *,
                 role_arn: pulumi.Input[str],
                 source: pulumi.Input[str],
                 target: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 desired_state: Optional[pulumi.Input['PipeRequestedPipeState']] = None,
                 enrichment: Optional[pulumi.Input[str]] = None,
                 enrichment_parameters: Optional[pulumi.Input['PipeEnrichmentParametersArgs']] = None,
                 log_configuration: Optional[pulumi.Input['PipeLogConfigurationArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 source_parameters: Optional[pulumi.Input['PipeSourceParametersArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target_parameters: Optional[pulumi.Input['PipeTargetParametersArgs']] = None):
        """
        The set of arguments for constructing a Pipe resource.
        """
        pulumi.set(__self__, "role_arn", role_arn)
        pulumi.set(__self__, "source", source)
        pulumi.set(__self__, "target", target)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if desired_state is not None:
            pulumi.set(__self__, "desired_state", desired_state)
        if enrichment is not None:
            pulumi.set(__self__, "enrichment", enrichment)
        if enrichment_parameters is not None:
            pulumi.set(__self__, "enrichment_parameters", enrichment_parameters)
        if log_configuration is not None:
            pulumi.set(__self__, "log_configuration", log_configuration)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if source_parameters is not None:
            pulumi.set(__self__, "source_parameters", source_parameters)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if target_parameters is not None:
            pulumi.set(__self__, "target_parameters", target_parameters)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Input[str]:
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter
    def source(self) -> pulumi.Input[str]:
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: pulumi.Input[str]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter
    def target(self) -> pulumi.Input[str]:
        return pulumi.get(self, "target")

    @target.setter
    def target(self, value: pulumi.Input[str]):
        pulumi.set(self, "target", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="desiredState")
    def desired_state(self) -> Optional[pulumi.Input['PipeRequestedPipeState']]:
        return pulumi.get(self, "desired_state")

    @desired_state.setter
    def desired_state(self, value: Optional[pulumi.Input['PipeRequestedPipeState']]):
        pulumi.set(self, "desired_state", value)

    @property
    @pulumi.getter
    def enrichment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "enrichment")

    @enrichment.setter
    def enrichment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enrichment", value)

    @property
    @pulumi.getter(name="enrichmentParameters")
    def enrichment_parameters(self) -> Optional[pulumi.Input['PipeEnrichmentParametersArgs']]:
        return pulumi.get(self, "enrichment_parameters")

    @enrichment_parameters.setter
    def enrichment_parameters(self, value: Optional[pulumi.Input['PipeEnrichmentParametersArgs']]):
        pulumi.set(self, "enrichment_parameters", value)

    @property
    @pulumi.getter(name="logConfiguration")
    def log_configuration(self) -> Optional[pulumi.Input['PipeLogConfigurationArgs']]:
        return pulumi.get(self, "log_configuration")

    @log_configuration.setter
    def log_configuration(self, value: Optional[pulumi.Input['PipeLogConfigurationArgs']]):
        pulumi.set(self, "log_configuration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="sourceParameters")
    def source_parameters(self) -> Optional[pulumi.Input['PipeSourceParametersArgs']]:
        return pulumi.get(self, "source_parameters")

    @source_parameters.setter
    def source_parameters(self, value: Optional[pulumi.Input['PipeSourceParametersArgs']]):
        pulumi.set(self, "source_parameters", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="targetParameters")
    def target_parameters(self) -> Optional[pulumi.Input['PipeTargetParametersArgs']]:
        return pulumi.get(self, "target_parameters")

    @target_parameters.setter
    def target_parameters(self, value: Optional[pulumi.Input['PipeTargetParametersArgs']]):
        pulumi.set(self, "target_parameters", value)


class Pipe(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 desired_state: Optional[pulumi.Input['PipeRequestedPipeState']] = None,
                 enrichment: Optional[pulumi.Input[str]] = None,
                 enrichment_parameters: Optional[pulumi.Input[pulumi.InputType['PipeEnrichmentParametersArgs']]] = None,
                 log_configuration: Optional[pulumi.Input[pulumi.InputType['PipeLogConfigurationArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 source_parameters: Optional[pulumi.Input[pulumi.InputType['PipeSourceParametersArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 target_parameters: Optional[pulumi.Input[pulumi.InputType['PipeTargetParametersArgs']]] = None,
                 __props__=None):
        """
        Definition of AWS::Pipes::Pipe Resource Type

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PipeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of AWS::Pipes::Pipe Resource Type

        :param str resource_name: The name of the resource.
        :param PipeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PipeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 desired_state: Optional[pulumi.Input['PipeRequestedPipeState']] = None,
                 enrichment: Optional[pulumi.Input[str]] = None,
                 enrichment_parameters: Optional[pulumi.Input[pulumi.InputType['PipeEnrichmentParametersArgs']]] = None,
                 log_configuration: Optional[pulumi.Input[pulumi.InputType['PipeLogConfigurationArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 source_parameters: Optional[pulumi.Input[pulumi.InputType['PipeSourceParametersArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 target_parameters: Optional[pulumi.Input[pulumi.InputType['PipeTargetParametersArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PipeArgs.__new__(PipeArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["desired_state"] = desired_state
            __props__.__dict__["enrichment"] = enrichment
            __props__.__dict__["enrichment_parameters"] = enrichment_parameters
            __props__.__dict__["log_configuration"] = log_configuration
            __props__.__dict__["name"] = name
            if role_arn is None and not opts.urn:
                raise TypeError("Missing required property 'role_arn'")
            __props__.__dict__["role_arn"] = role_arn
            if source is None and not opts.urn:
                raise TypeError("Missing required property 'source'")
            __props__.__dict__["source"] = source
            __props__.__dict__["source_parameters"] = source_parameters
            __props__.__dict__["tags"] = tags
            if target is None and not opts.urn:
                raise TypeError("Missing required property 'target'")
            __props__.__dict__["target"] = target
            __props__.__dict__["target_parameters"] = target_parameters
            __props__.__dict__["arn"] = None
            __props__.__dict__["creation_time"] = None
            __props__.__dict__["current_state"] = None
            __props__.__dict__["last_modified_time"] = None
            __props__.__dict__["state_reason"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["name", "source", "source_parameters.active_mq_broker_parameters.queue_name", "source_parameters.dynamo_db_stream_parameters.starting_position", "source_parameters.kinesis_stream_parameters.starting_position", "source_parameters.kinesis_stream_parameters.starting_position_timestamp", "source_parameters.managed_streaming_kafka_parameters.consumer_group_id", "source_parameters.managed_streaming_kafka_parameters.starting_position", "source_parameters.managed_streaming_kafka_parameters.topic_name", "source_parameters.rabbit_mq_broker_parameters.queue_name", "source_parameters.rabbit_mq_broker_parameters.virtual_host"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Pipe, __self__).__init__(
            'aws-native:pipes:Pipe',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Pipe':
        """
        Get an existing Pipe resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PipeArgs.__new__(PipeArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["creation_time"] = None
        __props__.__dict__["current_state"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["desired_state"] = None
        __props__.__dict__["enrichment"] = None
        __props__.__dict__["enrichment_parameters"] = None
        __props__.__dict__["last_modified_time"] = None
        __props__.__dict__["log_configuration"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["role_arn"] = None
        __props__.__dict__["source"] = None
        __props__.__dict__["source_parameters"] = None
        __props__.__dict__["state_reason"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["target"] = None
        __props__.__dict__["target_parameters"] = None
        return Pipe(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="currentState")
    def current_state(self) -> pulumi.Output['PipeState']:
        return pulumi.get(self, "current_state")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="desiredState")
    def desired_state(self) -> pulumi.Output[Optional['PipeRequestedPipeState']]:
        return pulumi.get(self, "desired_state")

    @property
    @pulumi.getter
    def enrichment(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "enrichment")

    @property
    @pulumi.getter(name="enrichmentParameters")
    def enrichment_parameters(self) -> pulumi.Output[Optional['outputs.PipeEnrichmentParameters']]:
        return pulumi.get(self, "enrichment_parameters")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter(name="logConfiguration")
    def log_configuration(self) -> pulumi.Output[Optional['outputs.PipeLogConfiguration']]:
        return pulumi.get(self, "log_configuration")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output[str]:
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="sourceParameters")
    def source_parameters(self) -> pulumi.Output[Optional['outputs.PipeSourceParameters']]:
        return pulumi.get(self, "source_parameters")

    @property
    @pulumi.getter(name="stateReason")
    def state_reason(self) -> pulumi.Output[str]:
        return pulumi.get(self, "state_reason")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def target(self) -> pulumi.Output[str]:
        return pulumi.get(self, "target")

    @property
    @pulumi.getter(name="targetParameters")
    def target_parameters(self) -> pulumi.Output[Optional['outputs.PipeTargetParameters']]:
        return pulumi.get(self, "target_parameters")

