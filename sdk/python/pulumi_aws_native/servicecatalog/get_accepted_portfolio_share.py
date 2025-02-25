# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetAcceptedPortfolioShareResult',
    'AwaitableGetAcceptedPortfolioShareResult',
    'get_accepted_portfolio_share',
    'get_accepted_portfolio_share_output',
]

@pulumi.output_type
class GetAcceptedPortfolioShareResult:
    def __init__(__self__, id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")


class AwaitableGetAcceptedPortfolioShareResult(GetAcceptedPortfolioShareResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAcceptedPortfolioShareResult(
            id=self.id)


def get_accepted_portfolio_share(id: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAcceptedPortfolioShareResult:
    """
    Resource Type definition for AWS::ServiceCatalog::AcceptedPortfolioShare
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:servicecatalog:getAcceptedPortfolioShare', __args__, opts=opts, typ=GetAcceptedPortfolioShareResult).value

    return AwaitableGetAcceptedPortfolioShareResult(
        id=pulumi.get(__ret__, 'id'))


@_utilities.lift_output_func(get_accepted_portfolio_share)
def get_accepted_portfolio_share_output(id: Optional[pulumi.Input[str]] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAcceptedPortfolioShareResult]:
    """
    Resource Type definition for AWS::ServiceCatalog::AcceptedPortfolioShare
    """
    ...
