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
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs
from ._inputs import *

__all__ = ['StreamArgs', 'Stream']

@pulumi.input_type
class StreamArgs:
    def __init__(__self__, *,
                 inclusive_start_time: pulumi.Input[str],
                 kinesis_configuration: pulumi.Input['StreamKinesisConfigurationArgs'],
                 ledger_name: pulumi.Input[str],
                 role_arn: pulumi.Input[str],
                 exclusive_end_time: Optional[pulumi.Input[str]] = None,
                 stream_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a Stream resource.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: An array of key-value pairs to apply to this resource.
        """
        pulumi.set(__self__, "inclusive_start_time", inclusive_start_time)
        pulumi.set(__self__, "kinesis_configuration", kinesis_configuration)
        pulumi.set(__self__, "ledger_name", ledger_name)
        pulumi.set(__self__, "role_arn", role_arn)
        if exclusive_end_time is not None:
            pulumi.set(__self__, "exclusive_end_time", exclusive_end_time)
        if stream_name is not None:
            pulumi.set(__self__, "stream_name", stream_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="inclusiveStartTime")
    def inclusive_start_time(self) -> pulumi.Input[str]:
        return pulumi.get(self, "inclusive_start_time")

    @inclusive_start_time.setter
    def inclusive_start_time(self, value: pulumi.Input[str]):
        pulumi.set(self, "inclusive_start_time", value)

    @property
    @pulumi.getter(name="kinesisConfiguration")
    def kinesis_configuration(self) -> pulumi.Input['StreamKinesisConfigurationArgs']:
        return pulumi.get(self, "kinesis_configuration")

    @kinesis_configuration.setter
    def kinesis_configuration(self, value: pulumi.Input['StreamKinesisConfigurationArgs']):
        pulumi.set(self, "kinesis_configuration", value)

    @property
    @pulumi.getter(name="ledgerName")
    def ledger_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "ledger_name")

    @ledger_name.setter
    def ledger_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "ledger_name", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Input[str]:
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter(name="exclusiveEndTime")
    def exclusive_end_time(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "exclusive_end_time")

    @exclusive_end_time.setter
    def exclusive_end_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "exclusive_end_time", value)

    @property
    @pulumi.getter(name="streamName")
    def stream_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "stream_name")

    @stream_name.setter
    def stream_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stream_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


class Stream(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 exclusive_end_time: Optional[pulumi.Input[str]] = None,
                 inclusive_start_time: Optional[pulumi.Input[str]] = None,
                 kinesis_configuration: Optional[pulumi.Input[pulumi.InputType['StreamKinesisConfigurationArgs']]] = None,
                 ledger_name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 stream_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None):
        """
        Resource schema for AWS::QLDB::Stream.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]] tags: An array of key-value pairs to apply to this resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StreamArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource schema for AWS::QLDB::Stream.

        :param str resource_name: The name of the resource.
        :param StreamArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StreamArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 exclusive_end_time: Optional[pulumi.Input[str]] = None,
                 inclusive_start_time: Optional[pulumi.Input[str]] = None,
                 kinesis_configuration: Optional[pulumi.Input[pulumi.InputType['StreamKinesisConfigurationArgs']]] = None,
                 ledger_name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 stream_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StreamArgs.__new__(StreamArgs)

            __props__.__dict__["exclusive_end_time"] = exclusive_end_time
            if inclusive_start_time is None and not opts.urn:
                raise TypeError("Missing required property 'inclusive_start_time'")
            __props__.__dict__["inclusive_start_time"] = inclusive_start_time
            if kinesis_configuration is None and not opts.urn:
                raise TypeError("Missing required property 'kinesis_configuration'")
            __props__.__dict__["kinesis_configuration"] = kinesis_configuration
            if ledger_name is None and not opts.urn:
                raise TypeError("Missing required property 'ledger_name'")
            __props__.__dict__["ledger_name"] = ledger_name
            if role_arn is None and not opts.urn:
                raise TypeError("Missing required property 'role_arn'")
            __props__.__dict__["role_arn"] = role_arn
            __props__.__dict__["stream_name"] = stream_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["aws_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["exclusive_end_time", "inclusive_start_time", "kinesis_configuration", "ledger_name", "role_arn", "stream_name"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Stream, __self__).__init__(
            'aws-native:qldb:Stream',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Stream':
        """
        Get an existing Stream resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = StreamArgs.__new__(StreamArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["aws_id"] = None
        __props__.__dict__["exclusive_end_time"] = None
        __props__.__dict__["inclusive_start_time"] = None
        __props__.__dict__["kinesis_configuration"] = None
        __props__.__dict__["ledger_name"] = None
        __props__.__dict__["role_arn"] = None
        __props__.__dict__["stream_name"] = None
        __props__.__dict__["tags"] = None
        return Stream(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="awsId")
    def aws_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "aws_id")

    @property
    @pulumi.getter(name="exclusiveEndTime")
    def exclusive_end_time(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "exclusive_end_time")

    @property
    @pulumi.getter(name="inclusiveStartTime")
    def inclusive_start_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "inclusive_start_time")

    @property
    @pulumi.getter(name="kinesisConfiguration")
    def kinesis_configuration(self) -> pulumi.Output['outputs.StreamKinesisConfiguration']:
        return pulumi.get(self, "kinesis_configuration")

    @property
    @pulumi.getter(name="ledgerName")
    def ledger_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ledger_name")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="streamName")
    def stream_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "stream_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

