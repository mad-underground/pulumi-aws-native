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

__all__ = ['NotificationChannelArgs', 'NotificationChannel']

@pulumi.input_type
class NotificationChannelArgs:
    def __init__(__self__, *,
                 config: pulumi.Input['NotificationChannelConfigArgs']):
        """
        The set of arguments for constructing a NotificationChannel resource.
        """
        pulumi.set(__self__, "config", config)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Input['NotificationChannelConfigArgs']:
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: pulumi.Input['NotificationChannelConfigArgs']):
        pulumi.set(self, "config", value)


class NotificationChannel(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[pulumi.InputType['NotificationChannelConfigArgs']]] = None,
                 __props__=None):
        """
        This resource schema represents the NotificationChannel resource in the Amazon DevOps Guru.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NotificationChannelArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource schema represents the NotificationChannel resource in the Amazon DevOps Guru.

        :param str resource_name: The name of the resource.
        :param NotificationChannelArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NotificationChannelArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[pulumi.InputType['NotificationChannelConfigArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NotificationChannelArgs.__new__(NotificationChannelArgs)

            if config is None and not opts.urn:
                raise TypeError("Missing required property 'config'")
            __props__.__dict__["config"] = config
            __props__.__dict__["aws_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["config"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(NotificationChannel, __self__).__init__(
            'aws-native:devopsguru:NotificationChannel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'NotificationChannel':
        """
        Get an existing NotificationChannel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = NotificationChannelArgs.__new__(NotificationChannelArgs)

        __props__.__dict__["aws_id"] = None
        __props__.__dict__["config"] = None
        return NotificationChannel(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="awsId")
    def aws_id(self) -> pulumi.Output[str]:
        """
        The ID of a notification channel.
        """
        return pulumi.get(self, "aws_id")

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output['outputs.NotificationChannelConfig']:
        return pulumi.get(self, "config")

