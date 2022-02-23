# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetVirtualClusterResult',
    'AwaitableGetVirtualClusterResult',
    'get_virtual_cluster',
    'get_virtual_cluster_output',
]

@pulumi.output_type
class GetVirtualClusterResult:
    def __init__(__self__, arn=None, id=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        Id of the virtual cluster.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.VirtualClusterTag']]:
        """
        An array of key-value pairs to apply to this virtual cluster.
        """
        return pulumi.get(self, "tags")


class AwaitableGetVirtualClusterResult(GetVirtualClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVirtualClusterResult(
            arn=self.arn,
            id=self.id,
            tags=self.tags)


def get_virtual_cluster(id: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVirtualClusterResult:
    """
    Resource Schema of AWS::EMRContainers::VirtualCluster Type


    :param str id: Id of the virtual cluster.
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:emrcontainers:getVirtualCluster', __args__, opts=opts, typ=GetVirtualClusterResult).value

    return AwaitableGetVirtualClusterResult(
        arn=__ret__.arn,
        id=__ret__.id,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_virtual_cluster)
def get_virtual_cluster_output(id: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVirtualClusterResult]:
    """
    Resource Schema of AWS::EMRContainers::VirtualCluster Type


    :param str id: Id of the virtual cluster.
    """
    ...