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
    'GetApplicationReferenceDataSourceResult',
    'AwaitableGetApplicationReferenceDataSourceResult',
    'get_application_reference_data_source',
    'get_application_reference_data_source_output',
]

@pulumi.output_type
class GetApplicationReferenceDataSourceResult:
    def __init__(__self__, id=None, reference_data_source=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if reference_data_source and not isinstance(reference_data_source, dict):
            raise TypeError("Expected argument 'reference_data_source' to be a dict")
        pulumi.set(__self__, "reference_data_source", reference_data_source)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="referenceDataSource")
    def reference_data_source(self) -> Optional['outputs.ApplicationReferenceDataSourceReferenceDataSource']:
        return pulumi.get(self, "reference_data_source")


class AwaitableGetApplicationReferenceDataSourceResult(GetApplicationReferenceDataSourceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApplicationReferenceDataSourceResult(
            id=self.id,
            reference_data_source=self.reference_data_source)


def get_application_reference_data_source(id: Optional[str] = None,
                                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApplicationReferenceDataSourceResult:
    """
    Resource Type definition for AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:kinesisanalyticsv2:getApplicationReferenceDataSource', __args__, opts=opts, typ=GetApplicationReferenceDataSourceResult).value

    return AwaitableGetApplicationReferenceDataSourceResult(
        id=__ret__.id,
        reference_data_source=__ret__.reference_data_source)


@_utilities.lift_output_func(get_application_reference_data_source)
def get_application_reference_data_source_output(id: Optional[pulumi.Input[str]] = None,
                                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetApplicationReferenceDataSourceResult]:
    """
    Resource Type definition for AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource
    """
    ...