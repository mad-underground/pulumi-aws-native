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
from ._enums import *
from ._inputs import *

__all__ = ['OriginEndpointArgs', 'OriginEndpoint']

@pulumi.input_type
class OriginEndpointArgs:
    def __init__(__self__, *,
                 channel_group_name: pulumi.Input[str],
                 channel_name: pulumi.Input[str],
                 container_type: Optional[pulumi.Input['OriginEndpointContainerType']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 hls_manifests: Optional[pulumi.Input[Sequence[pulumi.Input['OriginEndpointHlsManifestConfigurationArgs']]]] = None,
                 low_latency_hls_manifests: Optional[pulumi.Input[Sequence[pulumi.Input['OriginEndpointLowLatencyHlsManifestConfigurationArgs']]]] = None,
                 origin_endpoint_name: Optional[pulumi.Input[str]] = None,
                 segment: Optional[pulumi.Input['OriginEndpointSegmentArgs']] = None,
                 startover_window_seconds: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a OriginEndpoint resource.
        :param pulumi.Input[str] description: <p>Enter any descriptive text that helps you to identify the origin endpoint.</p>
        :param pulumi.Input[Sequence[pulumi.Input['OriginEndpointHlsManifestConfigurationArgs']]] hls_manifests: <p>An HTTP live streaming (HLS) manifest configuration.</p>
        :param pulumi.Input[Sequence[pulumi.Input['OriginEndpointLowLatencyHlsManifestConfigurationArgs']]] low_latency_hls_manifests: <p>A low-latency HLS manifest configuration.</p>
        :param pulumi.Input[int] startover_window_seconds: <p>The size of the window (in seconds) to create a window of the live stream that's available for on-demand viewing. Viewers can start-over or catch-up on content that falls within the window. The maximum startover window is 1,209,600 seconds (14 days).</p>
        """
        pulumi.set(__self__, "channel_group_name", channel_group_name)
        pulumi.set(__self__, "channel_name", channel_name)
        if container_type is not None:
            pulumi.set(__self__, "container_type", container_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if hls_manifests is not None:
            pulumi.set(__self__, "hls_manifests", hls_manifests)
        if low_latency_hls_manifests is not None:
            pulumi.set(__self__, "low_latency_hls_manifests", low_latency_hls_manifests)
        if origin_endpoint_name is not None:
            pulumi.set(__self__, "origin_endpoint_name", origin_endpoint_name)
        if segment is not None:
            pulumi.set(__self__, "segment", segment)
        if startover_window_seconds is not None:
            pulumi.set(__self__, "startover_window_seconds", startover_window_seconds)
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
    def channel_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "channel_name")

    @channel_name.setter
    def channel_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "channel_name", value)

    @property
    @pulumi.getter(name="containerType")
    def container_type(self) -> Optional[pulumi.Input['OriginEndpointContainerType']]:
        return pulumi.get(self, "container_type")

    @container_type.setter
    def container_type(self, value: Optional[pulumi.Input['OriginEndpointContainerType']]):
        pulumi.set(self, "container_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        <p>Enter any descriptive text that helps you to identify the origin endpoint.</p>
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="hlsManifests")
    def hls_manifests(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['OriginEndpointHlsManifestConfigurationArgs']]]]:
        """
        <p>An HTTP live streaming (HLS) manifest configuration.</p>
        """
        return pulumi.get(self, "hls_manifests")

    @hls_manifests.setter
    def hls_manifests(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['OriginEndpointHlsManifestConfigurationArgs']]]]):
        pulumi.set(self, "hls_manifests", value)

    @property
    @pulumi.getter(name="lowLatencyHlsManifests")
    def low_latency_hls_manifests(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['OriginEndpointLowLatencyHlsManifestConfigurationArgs']]]]:
        """
        <p>A low-latency HLS manifest configuration.</p>
        """
        return pulumi.get(self, "low_latency_hls_manifests")

    @low_latency_hls_manifests.setter
    def low_latency_hls_manifests(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['OriginEndpointLowLatencyHlsManifestConfigurationArgs']]]]):
        pulumi.set(self, "low_latency_hls_manifests", value)

    @property
    @pulumi.getter(name="originEndpointName")
    def origin_endpoint_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "origin_endpoint_name")

    @origin_endpoint_name.setter
    def origin_endpoint_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "origin_endpoint_name", value)

    @property
    @pulumi.getter
    def segment(self) -> Optional[pulumi.Input['OriginEndpointSegmentArgs']]:
        return pulumi.get(self, "segment")

    @segment.setter
    def segment(self, value: Optional[pulumi.Input['OriginEndpointSegmentArgs']]):
        pulumi.set(self, "segment", value)

    @property
    @pulumi.getter(name="startoverWindowSeconds")
    def startover_window_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        <p>The size of the window (in seconds) to create a window of the live stream that's available for on-demand viewing. Viewers can start-over or catch-up on content that falls within the window. The maximum startover window is 1,209,600 seconds (14 days).</p>
        """
        return pulumi.get(self, "startover_window_seconds")

    @startover_window_seconds.setter
    def startover_window_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "startover_window_seconds", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


class OriginEndpoint(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 channel_group_name: Optional[pulumi.Input[str]] = None,
                 channel_name: Optional[pulumi.Input[str]] = None,
                 container_type: Optional[pulumi.Input['OriginEndpointContainerType']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 hls_manifests: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OriginEndpointHlsManifestConfigurationArgs']]]]] = None,
                 low_latency_hls_manifests: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OriginEndpointLowLatencyHlsManifestConfigurationArgs']]]]] = None,
                 origin_endpoint_name: Optional[pulumi.Input[str]] = None,
                 segment: Optional[pulumi.Input[pulumi.InputType['OriginEndpointSegmentArgs']]] = None,
                 startover_window_seconds: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None):
        """
        <p>Represents an origin endpoint that is associated with a channel, offering a dynamically repackaged version of its content through various streaming media protocols. The content can be efficiently disseminated to end-users via a Content Delivery Network (CDN), like Amazon CloudFront.</p>

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: <p>Enter any descriptive text that helps you to identify the origin endpoint.</p>
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OriginEndpointHlsManifestConfigurationArgs']]]] hls_manifests: <p>An HTTP live streaming (HLS) manifest configuration.</p>
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OriginEndpointLowLatencyHlsManifestConfigurationArgs']]]] low_latency_hls_manifests: <p>A low-latency HLS manifest configuration.</p>
        :param pulumi.Input[int] startover_window_seconds: <p>The size of the window (in seconds) to create a window of the live stream that's available for on-demand viewing. Viewers can start-over or catch-up on content that falls within the window. The maximum startover window is 1,209,600 seconds (14 days).</p>
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OriginEndpointArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        <p>Represents an origin endpoint that is associated with a channel, offering a dynamically repackaged version of its content through various streaming media protocols. The content can be efficiently disseminated to end-users via a Content Delivery Network (CDN), like Amazon CloudFront.</p>

        :param str resource_name: The name of the resource.
        :param OriginEndpointArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OriginEndpointArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 channel_group_name: Optional[pulumi.Input[str]] = None,
                 channel_name: Optional[pulumi.Input[str]] = None,
                 container_type: Optional[pulumi.Input['OriginEndpointContainerType']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 hls_manifests: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OriginEndpointHlsManifestConfigurationArgs']]]]] = None,
                 low_latency_hls_manifests: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OriginEndpointLowLatencyHlsManifestConfigurationArgs']]]]] = None,
                 origin_endpoint_name: Optional[pulumi.Input[str]] = None,
                 segment: Optional[pulumi.Input[pulumi.InputType['OriginEndpointSegmentArgs']]] = None,
                 startover_window_seconds: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OriginEndpointArgs.__new__(OriginEndpointArgs)

            if channel_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'channel_group_name'")
            __props__.__dict__["channel_group_name"] = channel_group_name
            if channel_name is None and not opts.urn:
                raise TypeError("Missing required property 'channel_name'")
            __props__.__dict__["channel_name"] = channel_name
            __props__.__dict__["container_type"] = container_type
            __props__.__dict__["description"] = description
            __props__.__dict__["hls_manifests"] = hls_manifests
            __props__.__dict__["low_latency_hls_manifests"] = low_latency_hls_manifests
            __props__.__dict__["origin_endpoint_name"] = origin_endpoint_name
            __props__.__dict__["segment"] = segment
            __props__.__dict__["startover_window_seconds"] = startover_window_seconds
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["modified_at"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["channel_group_name", "channel_name", "origin_endpoint_name"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(OriginEndpoint, __self__).__init__(
            'aws-native:mediapackagev2:OriginEndpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'OriginEndpoint':
        """
        Get an existing OriginEndpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = OriginEndpointArgs.__new__(OriginEndpointArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["channel_group_name"] = None
        __props__.__dict__["channel_name"] = None
        __props__.__dict__["container_type"] = None
        __props__.__dict__["created_at"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["hls_manifests"] = None
        __props__.__dict__["low_latency_hls_manifests"] = None
        __props__.__dict__["modified_at"] = None
        __props__.__dict__["origin_endpoint_name"] = None
        __props__.__dict__["segment"] = None
        __props__.__dict__["startover_window_seconds"] = None
        __props__.__dict__["tags"] = None
        return OriginEndpoint(resource_name, opts=opts, __props__=__props__)

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
    @pulumi.getter(name="containerType")
    def container_type(self) -> pulumi.Output[Optional['OriginEndpointContainerType']]:
        return pulumi.get(self, "container_type")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        <p>The date and time the origin endpoint was created.</p>
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        <p>Enter any descriptive text that helps you to identify the origin endpoint.</p>
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="hlsManifests")
    def hls_manifests(self) -> pulumi.Output[Optional[Sequence['outputs.OriginEndpointHlsManifestConfiguration']]]:
        """
        <p>An HTTP live streaming (HLS) manifest configuration.</p>
        """
        return pulumi.get(self, "hls_manifests")

    @property
    @pulumi.getter(name="lowLatencyHlsManifests")
    def low_latency_hls_manifests(self) -> pulumi.Output[Optional[Sequence['outputs.OriginEndpointLowLatencyHlsManifestConfiguration']]]:
        """
        <p>A low-latency HLS manifest configuration.</p>
        """
        return pulumi.get(self, "low_latency_hls_manifests")

    @property
    @pulumi.getter(name="modifiedAt")
    def modified_at(self) -> pulumi.Output[str]:
        """
        <p>The date and time the origin endpoint was modified.</p>
        """
        return pulumi.get(self, "modified_at")

    @property
    @pulumi.getter(name="originEndpointName")
    def origin_endpoint_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "origin_endpoint_name")

    @property
    @pulumi.getter
    def segment(self) -> pulumi.Output[Optional['outputs.OriginEndpointSegment']]:
        return pulumi.get(self, "segment")

    @property
    @pulumi.getter(name="startoverWindowSeconds")
    def startover_window_seconds(self) -> pulumi.Output[Optional[int]]:
        """
        <p>The size of the window (in seconds) to create a window of the live stream that's available for on-demand viewing. Viewers can start-over or catch-up on content that falls within the window. The maximum startover window is 1,209,600 seconds (14 days).</p>
        """
        return pulumi.get(self, "startover_window_seconds")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        return pulumi.get(self, "tags")

