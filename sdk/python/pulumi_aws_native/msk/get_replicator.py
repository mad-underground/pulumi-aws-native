# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'GetReplicatorResult',
    'AwaitableGetReplicatorResult',
    'get_replicator',
    'get_replicator_output',
]

@pulumi.output_type
class GetReplicatorResult:
    def __init__(__self__, current_version=None, replication_info_list=None, replicator_arn=None, tags=None):
        if current_version and not isinstance(current_version, str):
            raise TypeError("Expected argument 'current_version' to be a str")
        pulumi.set(__self__, "current_version", current_version)
        if replication_info_list and not isinstance(replication_info_list, list):
            raise TypeError("Expected argument 'replication_info_list' to be a list")
        pulumi.set(__self__, "replication_info_list", replication_info_list)
        if replicator_arn and not isinstance(replicator_arn, str):
            raise TypeError("Expected argument 'replicator_arn' to be a str")
        pulumi.set(__self__, "replicator_arn", replicator_arn)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="currentVersion")
    def current_version(self) -> Optional[str]:
        """
        The current version of the MSK replicator.
        """
        return pulumi.get(self, "current_version")

    @property
    @pulumi.getter(name="replicationInfoList")
    def replication_info_list(self) -> Optional[Sequence['outputs.ReplicatorReplicationInfo']]:
        """
        A list of replication configurations, where each configuration targets a given source cluster to target cluster replication flow.
        """
        return pulumi.get(self, "replication_info_list")

    @property
    @pulumi.getter(name="replicatorArn")
    def replicator_arn(self) -> Optional[str]:
        """
        Amazon Resource Name for the created replicator.
        """
        return pulumi.get(self, "replicator_arn")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.ReplicatorTag']]:
        """
        A collection of tags associated with a resource
        """
        return pulumi.get(self, "tags")


class AwaitableGetReplicatorResult(GetReplicatorResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetReplicatorResult(
            current_version=self.current_version,
            replication_info_list=self.replication_info_list,
            replicator_arn=self.replicator_arn,
            tags=self.tags)


def get_replicator(replicator_arn: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetReplicatorResult:
    """
    Resource Type definition for AWS::MSK::Replicator


    :param str replicator_arn: Amazon Resource Name for the created replicator.
    """
    __args__ = dict()
    __args__['replicatorArn'] = replicator_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:msk:getReplicator', __args__, opts=opts, typ=GetReplicatorResult).value

    return AwaitableGetReplicatorResult(
        current_version=pulumi.get(__ret__, 'current_version'),
        replication_info_list=pulumi.get(__ret__, 'replication_info_list'),
        replicator_arn=pulumi.get(__ret__, 'replicator_arn'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_replicator)
def get_replicator_output(replicator_arn: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetReplicatorResult]:
    """
    Resource Type definition for AWS::MSK::Replicator


    :param str replicator_arn: Amazon Resource Name for the created replicator.
    """
    ...
