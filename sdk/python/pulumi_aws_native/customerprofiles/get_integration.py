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
    'GetIntegrationResult',
    'AwaitableGetIntegrationResult',
    'get_integration',
    'get_integration_output',
]

@pulumi.output_type
class GetIntegrationResult:
    def __init__(__self__, created_at=None, last_updated_at=None, object_type_name=None, tags=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if last_updated_at and not isinstance(last_updated_at, str):
            raise TypeError("Expected argument 'last_updated_at' to be a str")
        pulumi.set(__self__, "last_updated_at", last_updated_at)
        if object_type_name and not isinstance(object_type_name, str):
            raise TypeError("Expected argument 'object_type_name' to be a str")
        pulumi.set(__self__, "object_type_name", object_type_name)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[str]:
        """
        The time of this integration got created
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="lastUpdatedAt")
    def last_updated_at(self) -> Optional[str]:
        """
        The time of this integration got last updated at
        """
        return pulumi.get(self, "last_updated_at")

    @property
    @pulumi.getter(name="objectTypeName")
    def object_type_name(self) -> Optional[str]:
        """
        The name of the ObjectType defined for the 3rd party data in Profile Service
        """
        return pulumi.get(self, "object_type_name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.IntegrationTag']]:
        """
        The tags (keys and values) associated with the integration
        """
        return pulumi.get(self, "tags")


class AwaitableGetIntegrationResult(GetIntegrationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIntegrationResult(
            created_at=self.created_at,
            last_updated_at=self.last_updated_at,
            object_type_name=self.object_type_name,
            tags=self.tags)


def get_integration(domain_name: Optional[str] = None,
                    uri: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIntegrationResult:
    """
    The resource schema for creating an Amazon Connect Customer Profiles Integration.


    :param str domain_name: The unique name of the domain.
    :param str uri: The URI of the S3 bucket or any other type of data source.
    """
    __args__ = dict()
    __args__['domainName'] = domain_name
    __args__['uri'] = uri
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:customerprofiles:getIntegration', __args__, opts=opts, typ=GetIntegrationResult).value

    return AwaitableGetIntegrationResult(
        created_at=__ret__.created_at,
        last_updated_at=__ret__.last_updated_at,
        object_type_name=__ret__.object_type_name,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_integration)
def get_integration_output(domain_name: Optional[pulumi.Input[str]] = None,
                           uri: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIntegrationResult]:
    """
    The resource schema for creating an Amazon Connect Customer Profiles Integration.


    :param str domain_name: The unique name of the domain.
    :param str uri: The URI of the S3 bucket or any other type of data source.
    """
    ...
