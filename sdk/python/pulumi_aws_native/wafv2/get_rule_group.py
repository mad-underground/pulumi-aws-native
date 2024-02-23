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
from .. import outputs as _root_outputs
from ._enums import *

__all__ = [
    'GetRuleGroupResult',
    'AwaitableGetRuleGroupResult',
    'get_rule_group',
    'get_rule_group_output',
]

@pulumi.output_type
class GetRuleGroupResult:
    def __init__(__self__, arn=None, available_labels=None, capacity=None, consumed_labels=None, custom_response_bodies=None, description=None, id=None, label_namespace=None, rules=None, tags=None, visibility_config=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if available_labels and not isinstance(available_labels, list):
            raise TypeError("Expected argument 'available_labels' to be a list")
        pulumi.set(__self__, "available_labels", available_labels)
        if capacity and not isinstance(capacity, int):
            raise TypeError("Expected argument 'capacity' to be a int")
        pulumi.set(__self__, "capacity", capacity)
        if consumed_labels and not isinstance(consumed_labels, list):
            raise TypeError("Expected argument 'consumed_labels' to be a list")
        pulumi.set(__self__, "consumed_labels", consumed_labels)
        if custom_response_bodies and not isinstance(custom_response_bodies, dict):
            raise TypeError("Expected argument 'custom_response_bodies' to be a dict")
        pulumi.set(__self__, "custom_response_bodies", custom_response_bodies)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if label_namespace and not isinstance(label_namespace, str):
            raise TypeError("Expected argument 'label_namespace' to be a str")
        pulumi.set(__self__, "label_namespace", label_namespace)
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        pulumi.set(__self__, "rules", rules)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if visibility_config and not isinstance(visibility_config, dict):
            raise TypeError("Expected argument 'visibility_config' to be a dict")
        pulumi.set(__self__, "visibility_config", visibility_config)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="availableLabels")
    def available_labels(self) -> Optional[Sequence['outputs.RuleGroupLabelSummary']]:
        """
        Collection of Available Labels.
        """
        return pulumi.get(self, "available_labels")

    @property
    @pulumi.getter
    def capacity(self) -> Optional[int]:
        return pulumi.get(self, "capacity")

    @property
    @pulumi.getter(name="consumedLabels")
    def consumed_labels(self) -> Optional[Sequence['outputs.RuleGroupLabelSummary']]:
        """
        Collection of Consumed Labels.
        """
        return pulumi.get(self, "consumed_labels")

    @property
    @pulumi.getter(name="customResponseBodies")
    def custom_response_bodies(self) -> Optional[Mapping[str, 'outputs.RuleGroupCustomResponseBody']]:
        return pulumi.get(self, "custom_response_bodies")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="labelNamespace")
    def label_namespace(self) -> Optional[str]:
        return pulumi.get(self, "label_namespace")

    @property
    @pulumi.getter
    def rules(self) -> Optional[Sequence['outputs.RuleGroupRule']]:
        """
        Collection of Rules.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="visibilityConfig")
    def visibility_config(self) -> Optional['outputs.RuleGroupVisibilityConfig']:
        return pulumi.get(self, "visibility_config")


class AwaitableGetRuleGroupResult(GetRuleGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRuleGroupResult(
            arn=self.arn,
            available_labels=self.available_labels,
            capacity=self.capacity,
            consumed_labels=self.consumed_labels,
            custom_response_bodies=self.custom_response_bodies,
            description=self.description,
            id=self.id,
            label_namespace=self.label_namespace,
            rules=self.rules,
            tags=self.tags,
            visibility_config=self.visibility_config)


def get_rule_group(id: Optional[str] = None,
                   name: Optional[str] = None,
                   scope: Optional['RuleGroupScope'] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRuleGroupResult:
    """
    Contains the Rules that identify the requests that you want to allow, block, or count. In a RuleGroup, you also specify a default action (ALLOW or BLOCK), and the action for each Rule that you add to a RuleGroup, for example, block requests from specified IP addresses or block requests from specified referrers. You also associate the RuleGroup with a CloudFront distribution to identify the requests that you want AWS WAF to filter. If you add more than one Rule to a RuleGroup, a request needs to match only one of the specifications to be allowed, blocked, or counted.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    __args__['scope'] = scope
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:wafv2:getRuleGroup', __args__, opts=opts, typ=GetRuleGroupResult).value

    return AwaitableGetRuleGroupResult(
        arn=pulumi.get(__ret__, 'arn'),
        available_labels=pulumi.get(__ret__, 'available_labels'),
        capacity=pulumi.get(__ret__, 'capacity'),
        consumed_labels=pulumi.get(__ret__, 'consumed_labels'),
        custom_response_bodies=pulumi.get(__ret__, 'custom_response_bodies'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        label_namespace=pulumi.get(__ret__, 'label_namespace'),
        rules=pulumi.get(__ret__, 'rules'),
        tags=pulumi.get(__ret__, 'tags'),
        visibility_config=pulumi.get(__ret__, 'visibility_config'))


@_utilities.lift_output_func(get_rule_group)
def get_rule_group_output(id: Optional[pulumi.Input[str]] = None,
                          name: Optional[pulumi.Input[str]] = None,
                          scope: Optional[pulumi.Input['RuleGroupScope']] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRuleGroupResult]:
    """
    Contains the Rules that identify the requests that you want to allow, block, or count. In a RuleGroup, you also specify a default action (ALLOW or BLOCK), and the action for each Rule that you add to a RuleGroup, for example, block requests from specified IP addresses or block requests from specified referrers. You also associate the RuleGroup with a CloudFront distribution to identify the requests that you want AWS WAF to filter. If you add more than one Rule to a RuleGroup, a request needs to match only one of the specifications to be allowed, blocked, or counted.
    """
    ...
