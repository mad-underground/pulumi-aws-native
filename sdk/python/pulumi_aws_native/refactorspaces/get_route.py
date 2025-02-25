# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from .. import outputs as _root_outputs

__all__ = [
    'GetRouteResult',
    'AwaitableGetRouteResult',
    'get_route',
    'get_route_output',
]

@pulumi.output_type
class GetRouteResult:
    def __init__(__self__, arn=None, path_resource_to_id=None, route_identifier=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if path_resource_to_id and not isinstance(path_resource_to_id, str):
            raise TypeError("Expected argument 'path_resource_to_id' to be a str")
        pulumi.set(__self__, "path_resource_to_id", path_resource_to_id)
        if route_identifier and not isinstance(route_identifier, str):
            raise TypeError("Expected argument 'route_identifier' to be a str")
        pulumi.set(__self__, "route_identifier", route_identifier)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="pathResourceToId")
    def path_resource_to_id(self) -> Optional[str]:
        return pulumi.get(self, "path_resource_to_id")

    @property
    @pulumi.getter(name="routeIdentifier")
    def route_identifier(self) -> Optional[str]:
        return pulumi.get(self, "route_identifier")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.
        """
        return pulumi.get(self, "tags")


class AwaitableGetRouteResult(GetRouteResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouteResult(
            arn=self.arn,
            path_resource_to_id=self.path_resource_to_id,
            route_identifier=self.route_identifier,
            tags=self.tags)


def get_route(application_identifier: Optional[str] = None,
              environment_identifier: Optional[str] = None,
              route_identifier: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouteResult:
    """
    Definition of AWS::RefactorSpaces::Route Resource Type
    """
    __args__ = dict()
    __args__['applicationIdentifier'] = application_identifier
    __args__['environmentIdentifier'] = environment_identifier
    __args__['routeIdentifier'] = route_identifier
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:refactorspaces:getRoute', __args__, opts=opts, typ=GetRouteResult).value

    return AwaitableGetRouteResult(
        arn=pulumi.get(__ret__, 'arn'),
        path_resource_to_id=pulumi.get(__ret__, 'path_resource_to_id'),
        route_identifier=pulumi.get(__ret__, 'route_identifier'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_route)
def get_route_output(application_identifier: Optional[pulumi.Input[str]] = None,
                     environment_identifier: Optional[pulumi.Input[str]] = None,
                     route_identifier: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRouteResult]:
    """
    Definition of AWS::RefactorSpaces::Route Resource Type
    """
    ...
