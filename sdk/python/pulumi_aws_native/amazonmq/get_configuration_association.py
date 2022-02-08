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
    'GetConfigurationAssociationResult',
    'AwaitableGetConfigurationAssociationResult',
    'get_configuration_association',
    'get_configuration_association_output',
]

@pulumi.output_type
class GetConfigurationAssociationResult:
    def __init__(__self__, configuration=None, id=None):
        if configuration and not isinstance(configuration, dict):
            raise TypeError("Expected argument 'configuration' to be a dict")
        pulumi.set(__self__, "configuration", configuration)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def configuration(self) -> Optional['outputs.ConfigurationAssociationConfigurationId']:
        return pulumi.get(self, "configuration")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")


class AwaitableGetConfigurationAssociationResult(GetConfigurationAssociationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConfigurationAssociationResult(
            configuration=self.configuration,
            id=self.id)


def get_configuration_association(id: Optional[str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConfigurationAssociationResult:
    """
    Resource Type definition for AWS::AmazonMQ::ConfigurationAssociation
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:amazonmq:getConfigurationAssociation', __args__, opts=opts, typ=GetConfigurationAssociationResult).value

    return AwaitableGetConfigurationAssociationResult(
        configuration=__ret__.configuration,
        id=__ret__.id)


@_utilities.lift_output_func(get_configuration_association)
def get_configuration_association_output(id: Optional[pulumi.Input[str]] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetConfigurationAssociationResult]:
    """
    Resource Type definition for AWS::AmazonMQ::ConfigurationAssociation
    """
    ...
