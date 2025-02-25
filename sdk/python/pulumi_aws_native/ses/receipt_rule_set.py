# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ReceiptRuleSetArgs', 'ReceiptRuleSet']

@pulumi.input_type
class ReceiptRuleSetArgs:
    def __init__(__self__, *,
                 rule_set_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ReceiptRuleSet resource.
        """
        if rule_set_name is not None:
            pulumi.set(__self__, "rule_set_name", rule_set_name)

    @property
    @pulumi.getter(name="ruleSetName")
    def rule_set_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "rule_set_name")

    @rule_set_name.setter
    def rule_set_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rule_set_name", value)


warnings.warn("""ReceiptRuleSet is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class ReceiptRuleSet(pulumi.CustomResource):
    warnings.warn("""ReceiptRuleSet is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 rule_set_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::SES::ReceiptRuleSet

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ReceiptRuleSetArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::SES::ReceiptRuleSet

        :param str resource_name: The name of the resource.
        :param ReceiptRuleSetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReceiptRuleSetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 rule_set_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""ReceiptRuleSet is deprecated: ReceiptRuleSet is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ReceiptRuleSetArgs.__new__(ReceiptRuleSetArgs)

            __props__.__dict__["rule_set_name"] = rule_set_name
            __props__.__dict__["aws_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["rule_set_name"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(ReceiptRuleSet, __self__).__init__(
            'aws-native:ses:ReceiptRuleSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ReceiptRuleSet':
        """
        Get an existing ReceiptRuleSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ReceiptRuleSetArgs.__new__(ReceiptRuleSetArgs)

        __props__.__dict__["aws_id"] = None
        __props__.__dict__["rule_set_name"] = None
        return ReceiptRuleSet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="awsId")
    def aws_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "aws_id")

    @property
    @pulumi.getter(name="ruleSetName")
    def rule_set_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "rule_set_name")

