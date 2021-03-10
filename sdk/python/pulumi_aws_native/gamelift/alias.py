# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Alias']


class Alias(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 routing_strategy: Optional[pulumi.Input[pulumi.InputType['AliasRoutingStrategyArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html#cfn-gamelift-alias-description
        :param pulumi.Input[str] name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html#cfn-gamelift-alias-name
        :param pulumi.Input[pulumi.InputType['AliasRoutingStrategyArgs']] routing_strategy: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html#cfn-gamelift-alias-routingstrategy
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['description'] = description
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if routing_strategy is None and not opts.urn:
                raise TypeError("Missing required property 'routing_strategy'")
            __props__['routing_strategy'] = routing_strategy
            __props__['alias_id'] = None
        super(Alias, __self__).__init__(
            'aws-native:GameLift:Alias',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Alias':
        """
        Get an existing Alias resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Alias(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="AliasId")
    def alias_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "alias_id")

    @property
    @pulumi.getter(name="Description")
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html#cfn-gamelift-alias-description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="Name")
    def name(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html#cfn-gamelift-alias-name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="RoutingStrategy")
    def routing_strategy(self) -> pulumi.Output['outputs.AliasRoutingStrategy']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html#cfn-gamelift-alias-routingstrategy
        """
        return pulumi.get(self, "routing_strategy")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

