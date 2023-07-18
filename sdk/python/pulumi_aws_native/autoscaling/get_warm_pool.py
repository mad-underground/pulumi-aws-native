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
    'GetWarmPoolResult',
    'AwaitableGetWarmPoolResult',
    'get_warm_pool',
    'get_warm_pool_output',
]

@pulumi.output_type
class GetWarmPoolResult:
    def __init__(__self__, instance_reuse_policy=None, max_group_prepared_capacity=None, min_size=None, pool_state=None):
        if instance_reuse_policy and not isinstance(instance_reuse_policy, dict):
            raise TypeError("Expected argument 'instance_reuse_policy' to be a dict")
        pulumi.set(__self__, "instance_reuse_policy", instance_reuse_policy)
        if max_group_prepared_capacity and not isinstance(max_group_prepared_capacity, int):
            raise TypeError("Expected argument 'max_group_prepared_capacity' to be a int")
        pulumi.set(__self__, "max_group_prepared_capacity", max_group_prepared_capacity)
        if min_size and not isinstance(min_size, int):
            raise TypeError("Expected argument 'min_size' to be a int")
        pulumi.set(__self__, "min_size", min_size)
        if pool_state and not isinstance(pool_state, str):
            raise TypeError("Expected argument 'pool_state' to be a str")
        pulumi.set(__self__, "pool_state", pool_state)

    @property
    @pulumi.getter(name="instanceReusePolicy")
    def instance_reuse_policy(self) -> Optional['outputs.WarmPoolInstanceReusePolicy']:
        return pulumi.get(self, "instance_reuse_policy")

    @property
    @pulumi.getter(name="maxGroupPreparedCapacity")
    def max_group_prepared_capacity(self) -> Optional[int]:
        return pulumi.get(self, "max_group_prepared_capacity")

    @property
    @pulumi.getter(name="minSize")
    def min_size(self) -> Optional[int]:
        return pulumi.get(self, "min_size")

    @property
    @pulumi.getter(name="poolState")
    def pool_state(self) -> Optional[str]:
        return pulumi.get(self, "pool_state")


class AwaitableGetWarmPoolResult(GetWarmPoolResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWarmPoolResult(
            instance_reuse_policy=self.instance_reuse_policy,
            max_group_prepared_capacity=self.max_group_prepared_capacity,
            min_size=self.min_size,
            pool_state=self.pool_state)


def get_warm_pool(auto_scaling_group_name: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWarmPoolResult:
    """
    Resource schema for AWS::AutoScaling::WarmPool.
    """
    __args__ = dict()
    __args__['autoScalingGroupName'] = auto_scaling_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:autoscaling:getWarmPool', __args__, opts=opts, typ=GetWarmPoolResult).value

    return AwaitableGetWarmPoolResult(
        instance_reuse_policy=pulumi.get(__ret__, 'instance_reuse_policy'),
        max_group_prepared_capacity=pulumi.get(__ret__, 'max_group_prepared_capacity'),
        min_size=pulumi.get(__ret__, 'min_size'),
        pool_state=pulumi.get(__ret__, 'pool_state'))


@_utilities.lift_output_func(get_warm_pool)
def get_warm_pool_output(auto_scaling_group_name: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetWarmPoolResult]:
    """
    Resource schema for AWS::AutoScaling::WarmPool.
    """
    ...
