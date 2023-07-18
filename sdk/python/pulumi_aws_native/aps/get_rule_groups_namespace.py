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
    'GetRuleGroupsNamespaceResult',
    'AwaitableGetRuleGroupsNamespaceResult',
    'get_rule_groups_namespace',
    'get_rule_groups_namespace_output',
]

@pulumi.output_type
class GetRuleGroupsNamespaceResult:
    def __init__(__self__, arn=None, data=None, tags=None, workspace=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if data and not isinstance(data, str):
            raise TypeError("Expected argument 'data' to be a str")
        pulumi.set(__self__, "data", data)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if workspace and not isinstance(workspace, str):
            raise TypeError("Expected argument 'workspace' to be a str")
        pulumi.set(__self__, "workspace", workspace)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        The RuleGroupsNamespace ARN.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def data(self) -> Optional[str]:
        """
        The RuleGroupsNamespace data.
        """
        return pulumi.get(self, "data")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.RuleGroupsNamespaceTag']]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def workspace(self) -> Optional[str]:
        """
        Required to identify a specific APS Workspace associated with this RuleGroupsNamespace.
        """
        return pulumi.get(self, "workspace")


class AwaitableGetRuleGroupsNamespaceResult(GetRuleGroupsNamespaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRuleGroupsNamespaceResult(
            arn=self.arn,
            data=self.data,
            tags=self.tags,
            workspace=self.workspace)


def get_rule_groups_namespace(arn: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRuleGroupsNamespaceResult:
    """
    RuleGroupsNamespace schema for cloudformation.


    :param str arn: The RuleGroupsNamespace ARN.
    """
    __args__ = dict()
    __args__['arn'] = arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:aps:getRuleGroupsNamespace', __args__, opts=opts, typ=GetRuleGroupsNamespaceResult).value

    return AwaitableGetRuleGroupsNamespaceResult(
        arn=pulumi.get(__ret__, 'arn'),
        data=pulumi.get(__ret__, 'data'),
        tags=pulumi.get(__ret__, 'tags'),
        workspace=pulumi.get(__ret__, 'workspace'))


@_utilities.lift_output_func(get_rule_groups_namespace)
def get_rule_groups_namespace_output(arn: Optional[pulumi.Input[str]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRuleGroupsNamespaceResult]:
    """
    RuleGroupsNamespace schema for cloudformation.


    :param str arn: The RuleGroupsNamespace ARN.
    """
    ...
