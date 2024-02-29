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

__all__ = [
    'GetDistributionResult',
    'AwaitableGetDistributionResult',
    'get_distribution',
    'get_distribution_output',
]

@pulumi.output_type
class GetDistributionResult:
    def __init__(__self__, distribution_config=None, domain_name=None, id=None, tags=None):
        if distribution_config and not isinstance(distribution_config, dict):
            raise TypeError("Expected argument 'distribution_config' to be a dict")
        pulumi.set(__self__, "distribution_config", distribution_config)
        if domain_name and not isinstance(domain_name, str):
            raise TypeError("Expected argument 'domain_name' to be a str")
        pulumi.set(__self__, "domain_name", domain_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="distributionConfig")
    def distribution_config(self) -> Optional['outputs.DistributionConfig']:
        """
        The distribution's configuration.
        """
        return pulumi.get(self, "distribution_config")

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> Optional[str]:
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        A complex type that contains zero or more ``Tag`` elements.
        """
        return pulumi.get(self, "tags")


class AwaitableGetDistributionResult(GetDistributionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDistributionResult(
            distribution_config=self.distribution_config,
            domain_name=self.domain_name,
            id=self.id,
            tags=self.tags)


def get_distribution(id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDistributionResult:
    """
    A distribution tells CloudFront where you want content to be delivered from, and the details about how to track and manage content delivery.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:cloudfront:getDistribution', __args__, opts=opts, typ=GetDistributionResult).value

    return AwaitableGetDistributionResult(
        distribution_config=pulumi.get(__ret__, 'distribution_config'),
        domain_name=pulumi.get(__ret__, 'domain_name'),
        id=pulumi.get(__ret__, 'id'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_distribution)
def get_distribution_output(id: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDistributionResult]:
    """
    A distribution tells CloudFront where you want content to be delivered from, and the details about how to track and manage content delivery.
    """
    ...
