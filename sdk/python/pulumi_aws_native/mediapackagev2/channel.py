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

__all__ = ['ChannelArgs', 'Channel']

@pulumi.input_type
class ChannelArgs:
    def __init__(__self__, *,
                 channel_group_name: pulumi.Input[str],
                 channel_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a Channel resource.
        :param pulumi.Input[str] description: <p>Enter any descriptive text that helps you to identify the channel.</p>
        """
        pulumi.set(__self__, "channel_group_name", channel_group_name)
        if channel_name is not None:
            pulumi.set(__self__, "channel_name", channel_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="channelGroupName")
    def channel_group_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "channel_group_name")

    @channel_group_name.setter
    def channel_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "channel_group_name", value)

    @property
    @pulumi.getter(name="channelName")
    def channel_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "channel_name")

    @channel_name.setter
    def channel_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "channel_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        <p>Enter any descriptive text that helps you to identify the channel.</p>
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


class Channel(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 channel_group_name: Optional[pulumi.Input[str]] = None,
                 channel_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None):
        """
        <p>Represents an entry point into AWS Elemental MediaPackage for an ABR video content stream sent from an upstream encoder such as AWS Elemental MediaLive. The channel continuously analyzes the content that it receives and prepares it to be distributed to consumers via one or more origin endpoints.</p>

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: <p>Enter any descriptive text that helps you to identify the channel.</p>
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ChannelArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        <p>Represents an entry point into AWS Elemental MediaPackage for an ABR video content stream sent from an upstream encoder such as AWS Elemental MediaLive. The channel continuously analyzes the content that it receives and prepares it to be distributed to consumers via one or more origin endpoints.</p>

        :param str resource_name: The name of the resource.
        :param ChannelArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ChannelArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 channel_group_name: Optional[pulumi.Input[str]] = None,
                 channel_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ChannelArgs.__new__(ChannelArgs)

            if channel_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'channel_group_name'")
            __props__.__dict__["channel_group_name"] = channel_group_name
            __props__.__dict__["channel_name"] = channel_name
            __props__.__dict__["description"] = description
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["ingest_endpoints"] = None
            __props__.__dict__["modified_at"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["channel_group_name", "channel_name"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Channel, __self__).__init__(
            'aws-native:mediapackagev2:Channel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Channel':
        """
        Get an existing Channel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ChannelArgs.__new__(ChannelArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["channel_group_name"] = None
        __props__.__dict__["channel_name"] = None
        __props__.__dict__["created_at"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["ingest_endpoints"] = None
        __props__.__dict__["modified_at"] = None
        __props__.__dict__["tags"] = None
        return Channel(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        <p>The Amazon Resource Name (ARN) associated with the resource.</p>
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="channelGroupName")
    def channel_group_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "channel_group_name")

    @property
    @pulumi.getter(name="channelName")
    def channel_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "channel_name")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        <p>The date and time the channel was created.</p>
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        <p>Enter any descriptive text that helps you to identify the channel.</p>
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="ingestEndpoints")
    def ingest_endpoints(self) -> pulumi.Output[Sequence['outputs.ChannelIngestEndpoint']]:
        """
        <p>The list of ingest endpoints.</p>
        """
        return pulumi.get(self, "ingest_endpoints")

    @property
    @pulumi.getter(name="modifiedAt")
    def modified_at(self) -> pulumi.Output[str]:
        """
        <p>The date and time the channel was modified.</p>
        """
        return pulumi.get(self, "modified_at")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        return pulumi.get(self, "tags")

