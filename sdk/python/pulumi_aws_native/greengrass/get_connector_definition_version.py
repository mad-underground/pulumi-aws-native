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
    'GetConnectorDefinitionVersionResult',
    'AwaitableGetConnectorDefinitionVersionResult',
    'get_connector_definition_version',
    'get_connector_definition_version_output',
]

@pulumi.output_type
class GetConnectorDefinitionVersionResult:
    def __init__(__self__, id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")


class AwaitableGetConnectorDefinitionVersionResult(GetConnectorDefinitionVersionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConnectorDefinitionVersionResult(
            id=self.id)


def get_connector_definition_version(id: Optional[str] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConnectorDefinitionVersionResult:
    """
    Resource Type definition for AWS::Greengrass::ConnectorDefinitionVersion
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:greengrass:getConnectorDefinitionVersion', __args__, opts=opts, typ=GetConnectorDefinitionVersionResult).value

    return AwaitableGetConnectorDefinitionVersionResult(
        id=pulumi.get(__ret__, 'id'))


@_utilities.lift_output_func(get_connector_definition_version)
def get_connector_definition_version_output(id: Optional[pulumi.Input[str]] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetConnectorDefinitionVersionResult]:
    """
    Resource Type definition for AWS::Greengrass::ConnectorDefinitionVersion
    """
    ...
