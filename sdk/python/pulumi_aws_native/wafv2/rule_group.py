# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs
from ._inputs import *

__all__ = ['RuleGroup']


class RuleGroup(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 capacity: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleGroupRuleArgs']]]]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 visibility_config: Optional[pulumi.Input[pulumi.InputType['RuleGroupVisibilityConfigArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-rulegroup.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] capacity: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-rulegroup.html#cfn-wafv2-rulegroup-capacity
        :param pulumi.Input[str] description: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-rulegroup.html#cfn-wafv2-rulegroup-description
        :param pulumi.Input[str] name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-rulegroup.html#cfn-wafv2-rulegroup-name
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleGroupRuleArgs']]]] rules: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-rulegroup.html#cfn-wafv2-rulegroup-rules
        :param pulumi.Input[str] scope: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-rulegroup.html#cfn-wafv2-rulegroup-scope
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-rulegroup.html#cfn-wafv2-rulegroup-tags
        :param pulumi.Input[pulumi.InputType['RuleGroupVisibilityConfigArgs']] visibility_config: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-rulegroup.html#cfn-wafv2-rulegroup-visibilityconfig
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

            if capacity is None and not opts.urn:
                raise TypeError("Missing required property 'capacity'")
            __props__['capacity'] = capacity
            __props__['description'] = description
            __props__['name'] = name
            __props__['rules'] = rules
            if scope is None and not opts.urn:
                raise TypeError("Missing required property 'scope'")
            __props__['scope'] = scope
            __props__['tags'] = tags
            if visibility_config is None and not opts.urn:
                raise TypeError("Missing required property 'visibility_config'")
            __props__['visibility_config'] = visibility_config
            __props__['arn'] = None
            __props__['id'] = None
        super(RuleGroup, __self__).__init__(
            'aws-native:WAFv2:RuleGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'RuleGroup':
        """
        Get an existing RuleGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return RuleGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def capacity(self) -> pulumi.Output[int]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-rulegroup.html#cfn-wafv2-rulegroup-capacity
        """
        return pulumi.get(self, "capacity")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-rulegroup.html#cfn-wafv2-rulegroup-description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-rulegroup.html#cfn-wafv2-rulegroup-name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Optional[Sequence['outputs.RuleGroupRule']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-rulegroup.html#cfn-wafv2-rulegroup-rules
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-rulegroup.html#cfn-wafv2-rulegroup-scope
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-rulegroup.html#cfn-wafv2-rulegroup-tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="visibilityConfig")
    def visibility_config(self) -> pulumi.Output['outputs.RuleGroupVisibilityConfig']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-rulegroup.html#cfn-wafv2-rulegroup-visibilityconfig
        """
        return pulumi.get(self, "visibility_config")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

