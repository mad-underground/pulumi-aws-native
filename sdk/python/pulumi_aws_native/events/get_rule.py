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

__all__ = [
    'GetRuleResult',
    'AwaitableGetRuleResult',
    'get_rule',
    'get_rule_output',
]

@pulumi.output_type
class GetRuleResult:
    def __init__(__self__, arn=None, description=None, event_pattern=None, id=None, role_arn=None, schedule_expression=None, state=None, targets=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if event_pattern and not isinstance(event_pattern, dict):
            raise TypeError("Expected argument 'event_pattern' to be a dict")
        pulumi.set(__self__, "event_pattern", event_pattern)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if role_arn and not isinstance(role_arn, str):
            raise TypeError("Expected argument 'role_arn' to be a str")
        pulumi.set(__self__, "role_arn", role_arn)
        if schedule_expression and not isinstance(schedule_expression, str):
            raise TypeError("Expected argument 'schedule_expression' to be a str")
        pulumi.set(__self__, "schedule_expression", schedule_expression)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if targets and not isinstance(targets, list):
            raise TypeError("Expected argument 'targets' to be a list")
        pulumi.set(__self__, "targets", targets)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="eventPattern")
    def event_pattern(self) -> Optional[Any]:
        return pulumi.get(self, "event_pattern")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[str]:
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="scheduleExpression")
    def schedule_expression(self) -> Optional[str]:
        return pulumi.get(self, "schedule_expression")

    @property
    @pulumi.getter
    def state(self) -> Optional[str]:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def targets(self) -> Optional[Sequence['outputs.RuleTarget']]:
        return pulumi.get(self, "targets")


class AwaitableGetRuleResult(GetRuleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRuleResult(
            arn=self.arn,
            description=self.description,
            event_pattern=self.event_pattern,
            id=self.id,
            role_arn=self.role_arn,
            schedule_expression=self.schedule_expression,
            state=self.state,
            targets=self.targets)


def get_rule(id: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRuleResult:
    """
    Resource Type definition for AWS::Events::Rule
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:events:getRule', __args__, opts=opts, typ=GetRuleResult).value

    return AwaitableGetRuleResult(
        arn=pulumi.get(__ret__, 'arn'),
        description=pulumi.get(__ret__, 'description'),
        event_pattern=pulumi.get(__ret__, 'event_pattern'),
        id=pulumi.get(__ret__, 'id'),
        role_arn=pulumi.get(__ret__, 'role_arn'),
        schedule_expression=pulumi.get(__ret__, 'schedule_expression'),
        state=pulumi.get(__ret__, 'state'),
        targets=pulumi.get(__ret__, 'targets'))


@_utilities.lift_output_func(get_rule)
def get_rule_output(id: Optional[pulumi.Input[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRuleResult]:
    """
    Resource Type definition for AWS::Events::Rule
    """
    ...
