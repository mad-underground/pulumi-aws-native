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

__all__ = [
    'GetResiliencyPolicyResult',
    'AwaitableGetResiliencyPolicyResult',
    'get_resiliency_policy',
    'get_resiliency_policy_output',
]

@pulumi.output_type
class GetResiliencyPolicyResult:
    def __init__(__self__, data_location_constraint=None, policy=None, policy_arn=None, policy_description=None, policy_name=None, tags=None, tier=None):
        if data_location_constraint and not isinstance(data_location_constraint, str):
            raise TypeError("Expected argument 'data_location_constraint' to be a str")
        pulumi.set(__self__, "data_location_constraint", data_location_constraint)
        if policy and not isinstance(policy, dict):
            raise TypeError("Expected argument 'policy' to be a dict")
        pulumi.set(__self__, "policy", policy)
        if policy_arn and not isinstance(policy_arn, str):
            raise TypeError("Expected argument 'policy_arn' to be a str")
        pulumi.set(__self__, "policy_arn", policy_arn)
        if policy_description and not isinstance(policy_description, str):
            raise TypeError("Expected argument 'policy_description' to be a str")
        pulumi.set(__self__, "policy_description", policy_description)
        if policy_name and not isinstance(policy_name, str):
            raise TypeError("Expected argument 'policy_name' to be a str")
        pulumi.set(__self__, "policy_name", policy_name)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if tier and not isinstance(tier, str):
            raise TypeError("Expected argument 'tier' to be a str")
        pulumi.set(__self__, "tier", tier)

    @property
    @pulumi.getter(name="dataLocationConstraint")
    def data_location_constraint(self) -> Optional['ResiliencyPolicyDataLocationConstraint']:
        """
        Data Location Constraint of the Policy.
        """
        return pulumi.get(self, "data_location_constraint")

    @property
    @pulumi.getter
    def policy(self) -> Optional[Mapping[str, 'outputs.ResiliencyPolicyFailurePolicy']]:
        return pulumi.get(self, "policy")

    @property
    @pulumi.getter(name="policyArn")
    def policy_arn(self) -> Optional[str]:
        """
        Amazon Resource Name (ARN) of the Resiliency Policy.
        """
        return pulumi.get(self, "policy_arn")

    @property
    @pulumi.getter(name="policyDescription")
    def policy_description(self) -> Optional[str]:
        """
        Description of Resiliency Policy.
        """
        return pulumi.get(self, "policy_description")

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> Optional[str]:
        """
        Name of Resiliency Policy.
        """
        return pulumi.get(self, "policy_name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def tier(self) -> Optional['ResiliencyPolicyTier']:
        """
        Resiliency Policy Tier.
        """
        return pulumi.get(self, "tier")


class AwaitableGetResiliencyPolicyResult(GetResiliencyPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResiliencyPolicyResult(
            data_location_constraint=self.data_location_constraint,
            policy=self.policy,
            policy_arn=self.policy_arn,
            policy_description=self.policy_description,
            policy_name=self.policy_name,
            tags=self.tags,
            tier=self.tier)


def get_resiliency_policy(policy_arn: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetResiliencyPolicyResult:
    """
    Resource Type Definition for Resiliency Policy.


    :param str policy_arn: Amazon Resource Name (ARN) of the Resiliency Policy.
    """
    __args__ = dict()
    __args__['policyArn'] = policy_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:resiliencehub:getResiliencyPolicy', __args__, opts=opts, typ=GetResiliencyPolicyResult).value

    return AwaitableGetResiliencyPolicyResult(
        data_location_constraint=pulumi.get(__ret__, 'data_location_constraint'),
        policy=pulumi.get(__ret__, 'policy'),
        policy_arn=pulumi.get(__ret__, 'policy_arn'),
        policy_description=pulumi.get(__ret__, 'policy_description'),
        policy_name=pulumi.get(__ret__, 'policy_name'),
        tags=pulumi.get(__ret__, 'tags'),
        tier=pulumi.get(__ret__, 'tier'))


@_utilities.lift_output_func(get_resiliency_policy)
def get_resiliency_policy_output(policy_arn: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetResiliencyPolicyResult]:
    """
    Resource Type Definition for Resiliency Policy.


    :param str policy_arn: Amazon Resource Name (ARN) of the Resiliency Policy.
    """
    ...
