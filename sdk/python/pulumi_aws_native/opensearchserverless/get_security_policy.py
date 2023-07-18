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

__all__ = [
    'GetSecurityPolicyResult',
    'AwaitableGetSecurityPolicyResult',
    'get_security_policy',
    'get_security_policy_output',
]

@pulumi.output_type
class GetSecurityPolicyResult:
    def __init__(__self__, description=None, policy=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if policy and not isinstance(policy, str):
            raise TypeError("Expected argument 'policy' to be a str")
        pulumi.set(__self__, "policy", policy)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the policy
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def policy(self) -> Optional[str]:
        """
        The JSON policy document that is the content for the policy
        """
        return pulumi.get(self, "policy")


class AwaitableGetSecurityPolicyResult(GetSecurityPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecurityPolicyResult(
            description=self.description,
            policy=self.policy)


def get_security_policy(name: Optional[str] = None,
                        type: Optional['SecurityPolicyType'] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSecurityPolicyResult:
    """
    Amazon OpenSearchServerless security policy resource


    :param str name: The name of the policy
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['type'] = type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:opensearchserverless:getSecurityPolicy', __args__, opts=opts, typ=GetSecurityPolicyResult).value

    return AwaitableGetSecurityPolicyResult(
        description=pulumi.get(__ret__, 'description'),
        policy=pulumi.get(__ret__, 'policy'))


@_utilities.lift_output_func(get_security_policy)
def get_security_policy_output(name: Optional[pulumi.Input[str]] = None,
                               type: Optional[pulumi.Input['SecurityPolicyType']] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSecurityPolicyResult]:
    """
    Amazon OpenSearchServerless security policy resource


    :param str name: The name of the policy
    """
    ...
