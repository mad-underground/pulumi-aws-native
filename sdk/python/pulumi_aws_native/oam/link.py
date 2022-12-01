# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = ['LinkArgs', 'Link']

@pulumi.input_type
class LinkArgs:
    def __init__(__self__, *,
                 label_template: pulumi.Input[str],
                 resource_types: pulumi.Input[Sequence[pulumi.Input['LinkResourceType']]],
                 sink_identifier: pulumi.Input[str],
                 tags: Optional[Any] = None):
        """
        The set of arguments for constructing a Link resource.
        :param Any tags: Tags to apply to the link
        """
        pulumi.set(__self__, "label_template", label_template)
        pulumi.set(__self__, "resource_types", resource_types)
        pulumi.set(__self__, "sink_identifier", sink_identifier)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="labelTemplate")
    def label_template(self) -> pulumi.Input[str]:
        return pulumi.get(self, "label_template")

    @label_template.setter
    def label_template(self, value: pulumi.Input[str]):
        pulumi.set(self, "label_template", value)

    @property
    @pulumi.getter(name="resourceTypes")
    def resource_types(self) -> pulumi.Input[Sequence[pulumi.Input['LinkResourceType']]]:
        return pulumi.get(self, "resource_types")

    @resource_types.setter
    def resource_types(self, value: pulumi.Input[Sequence[pulumi.Input['LinkResourceType']]]):
        pulumi.set(self, "resource_types", value)

    @property
    @pulumi.getter(name="sinkIdentifier")
    def sink_identifier(self) -> pulumi.Input[str]:
        return pulumi.get(self, "sink_identifier")

    @sink_identifier.setter
    def sink_identifier(self, value: pulumi.Input[str]):
        pulumi.set(self, "sink_identifier", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[Any]:
        """
        Tags to apply to the link
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[Any]):
        pulumi.set(self, "tags", value)


class Link(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 label_template: Optional[pulumi.Input[str]] = None,
                 resource_types: Optional[pulumi.Input[Sequence[pulumi.Input['LinkResourceType']]]] = None,
                 sink_identifier: Optional[pulumi.Input[str]] = None,
                 tags: Optional[Any] = None,
                 __props__=None):
        """
        Definition of AWS::Oam::Link Resource Type

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param Any tags: Tags to apply to the link
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LinkArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of AWS::Oam::Link Resource Type

        :param str resource_name: The name of the resource.
        :param LinkArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LinkArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 label_template: Optional[pulumi.Input[str]] = None,
                 resource_types: Optional[pulumi.Input[Sequence[pulumi.Input['LinkResourceType']]]] = None,
                 sink_identifier: Optional[pulumi.Input[str]] = None,
                 tags: Optional[Any] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LinkArgs.__new__(LinkArgs)

            if label_template is None and not opts.urn:
                raise TypeError("Missing required property 'label_template'")
            __props__.__dict__["label_template"] = label_template
            if resource_types is None and not opts.urn:
                raise TypeError("Missing required property 'resource_types'")
            __props__.__dict__["resource_types"] = resource_types
            if sink_identifier is None and not opts.urn:
                raise TypeError("Missing required property 'sink_identifier'")
            __props__.__dict__["sink_identifier"] = sink_identifier
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["label"] = None
        super(Link, __self__).__init__(
            'aws-native:oam:Link',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Link':
        """
        Get an existing Link resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = LinkArgs.__new__(LinkArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["label"] = None
        __props__.__dict__["label_template"] = None
        __props__.__dict__["resource_types"] = None
        __props__.__dict__["sink_identifier"] = None
        __props__.__dict__["tags"] = None
        return Link(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def label(self) -> pulumi.Output[str]:
        return pulumi.get(self, "label")

    @property
    @pulumi.getter(name="labelTemplate")
    def label_template(self) -> pulumi.Output[str]:
        return pulumi.get(self, "label_template")

    @property
    @pulumi.getter(name="resourceTypes")
    def resource_types(self) -> pulumi.Output[Sequence['LinkResourceType']]:
        return pulumi.get(self, "resource_types")

    @property
    @pulumi.getter(name="sinkIdentifier")
    def sink_identifier(self) -> pulumi.Output[str]:
        return pulumi.get(self, "sink_identifier")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Any]]:
        """
        Tags to apply to the link
        """
        return pulumi.get(self, "tags")
