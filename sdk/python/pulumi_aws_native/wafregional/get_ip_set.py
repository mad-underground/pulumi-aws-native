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
    'GetIPSetResult',
    'AwaitableGetIPSetResult',
    'get_ip_set',
    'get_ip_set_output',
]

@pulumi.output_type
class GetIPSetResult:
    def __init__(__self__, i_p_set_descriptors=None, id=None):
        if i_p_set_descriptors and not isinstance(i_p_set_descriptors, list):
            raise TypeError("Expected argument 'i_p_set_descriptors' to be a list")
        pulumi.set(__self__, "i_p_set_descriptors", i_p_set_descriptors)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="iPSetDescriptors")
    def i_p_set_descriptors(self) -> Optional[Sequence['outputs.IPSetDescriptor']]:
        return pulumi.get(self, "i_p_set_descriptors")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")


class AwaitableGetIPSetResult(GetIPSetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIPSetResult(
            i_p_set_descriptors=self.i_p_set_descriptors,
            id=self.id)


def get_ip_set(id: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIPSetResult:
    """
    Resource Type definition for AWS::WAFRegional::IPSet
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:wafregional:getIPSet', __args__, opts=opts, typ=GetIPSetResult).value

    return AwaitableGetIPSetResult(
        i_p_set_descriptors=pulumi.get(__ret__, 'i_p_set_descriptors'),
        id=pulumi.get(__ret__, 'id'))


@_utilities.lift_output_func(get_ip_set)
def get_ip_set_output(id: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIPSetResult]:
    """
    Resource Type definition for AWS::WAFRegional::IPSet
    """
    ...
