# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetTransitGatewayRouteResult',
    'AwaitableGetTransitGatewayRouteResult',
    'get_transit_gateway_route',
    'get_transit_gateway_route_output',
]

@pulumi.output_type
class GetTransitGatewayRouteResult:
    def __init__(__self__, id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")


class AwaitableGetTransitGatewayRouteResult(GetTransitGatewayRouteResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTransitGatewayRouteResult(
            id=self.id)


def get_transit_gateway_route(id: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTransitGatewayRouteResult:
    """
    Resource Type definition for AWS::EC2::TransitGatewayRoute
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:ec2:getTransitGatewayRoute', __args__, opts=opts, typ=GetTransitGatewayRouteResult).value

    return AwaitableGetTransitGatewayRouteResult(
        id=__ret__.id)


@_utilities.lift_output_func(get_transit_gateway_route)
def get_transit_gateway_route_output(id: Optional[pulumi.Input[str]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTransitGatewayRouteResult]:
    """
    Resource Type definition for AWS::EC2::TransitGatewayRoute
    """
    ...
