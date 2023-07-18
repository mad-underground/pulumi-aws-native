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
    'GetRoutingControlResult',
    'AwaitableGetRoutingControlResult',
    'get_routing_control',
    'get_routing_control_output',
]

@pulumi.output_type
class GetRoutingControlResult:
    def __init__(__self__, name=None, routing_control_arn=None, status=None):
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if routing_control_arn and not isinstance(routing_control_arn, str):
            raise TypeError("Expected argument 'routing_control_arn' to be a str")
        pulumi.set(__self__, "routing_control_arn", routing_control_arn)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the routing control. You can use any non-white space character in the name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="routingControlArn")
    def routing_control_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the routing control.
        """
        return pulumi.get(self, "routing_control_arn")

    @property
    @pulumi.getter
    def status(self) -> Optional['RoutingControlStatus']:
        """
        The deployment status of the routing control. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.
        """
        return pulumi.get(self, "status")


class AwaitableGetRoutingControlResult(GetRoutingControlResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRoutingControlResult(
            name=self.name,
            routing_control_arn=self.routing_control_arn,
            status=self.status)


def get_routing_control(routing_control_arn: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRoutingControlResult:
    """
    AWS Route53 Recovery Control Routing Control resource schema .


    :param str routing_control_arn: The Amazon Resource Name (ARN) of the routing control.
    """
    __args__ = dict()
    __args__['routingControlArn'] = routing_control_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:route53recoverycontrol:getRoutingControl', __args__, opts=opts, typ=GetRoutingControlResult).value

    return AwaitableGetRoutingControlResult(
        name=pulumi.get(__ret__, 'name'),
        routing_control_arn=pulumi.get(__ret__, 'routing_control_arn'),
        status=pulumi.get(__ret__, 'status'))


@_utilities.lift_output_func(get_routing_control)
def get_routing_control_output(routing_control_arn: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRoutingControlResult]:
    """
    AWS Route53 Recovery Control Routing Control resource schema .


    :param str routing_control_arn: The Amazon Resource Name (ARN) of the routing control.
    """
    ...
