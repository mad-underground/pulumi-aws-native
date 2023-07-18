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
from ._enums import *

__all__ = [
    'GetDataRepositoryAssociationResult',
    'AwaitableGetDataRepositoryAssociationResult',
    'get_data_repository_association',
    'get_data_repository_association_output',
]

@pulumi.output_type
class GetDataRepositoryAssociationResult:
    def __init__(__self__, association_id=None, imported_file_chunk_size=None, resource_arn=None, s3=None, tags=None):
        if association_id and not isinstance(association_id, str):
            raise TypeError("Expected argument 'association_id' to be a str")
        pulumi.set(__self__, "association_id", association_id)
        if imported_file_chunk_size and not isinstance(imported_file_chunk_size, int):
            raise TypeError("Expected argument 'imported_file_chunk_size' to be a int")
        pulumi.set(__self__, "imported_file_chunk_size", imported_file_chunk_size)
        if resource_arn and not isinstance(resource_arn, str):
            raise TypeError("Expected argument 'resource_arn' to be a str")
        pulumi.set(__self__, "resource_arn", resource_arn)
        if s3 and not isinstance(s3, dict):
            raise TypeError("Expected argument 's3' to be a dict")
        pulumi.set(__self__, "s3", s3)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="associationId")
    def association_id(self) -> Optional[str]:
        """
        The system-generated, unique ID of the data repository association.
        """
        return pulumi.get(self, "association_id")

    @property
    @pulumi.getter(name="importedFileChunkSize")
    def imported_file_chunk_size(self) -> Optional[int]:
        """
        For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. The maximum number of disks that a single file can be striped across is limited by the total number of disks that make up the file system.
        """
        return pulumi.get(self, "imported_file_chunk_size")

    @property
    @pulumi.getter(name="resourceARN")
    def resource_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) for a given resource. ARNs uniquely identify Amazon Web Services resources. We require an ARN when you need to specify a resource unambiguously across all of Amazon Web Services. For more information, see Amazon Resource Names (ARNs) in the Amazon Web Services General Reference.
        """
        return pulumi.get(self, "resource_arn")

    @property
    @pulumi.getter
    def s3(self) -> Optional['outputs.DataRepositoryAssociationS3']:
        """
        The configuration for an Amazon S3 data repository linked to an Amazon FSx Lustre file system with a data repository association. The configuration defines which file events (new, changed, or deleted files or directories) are automatically imported from the linked data repository to the file system or automatically exported from the file system to the data repository.
        """
        return pulumi.get(self, "s3")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.DataRepositoryAssociationTag']]:
        """
        A list of Tag values, with a maximum of 50 elements.
        """
        return pulumi.get(self, "tags")


class AwaitableGetDataRepositoryAssociationResult(GetDataRepositoryAssociationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDataRepositoryAssociationResult(
            association_id=self.association_id,
            imported_file_chunk_size=self.imported_file_chunk_size,
            resource_arn=self.resource_arn,
            s3=self.s3,
            tags=self.tags)


def get_data_repository_association(association_id: Optional[str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDataRepositoryAssociationResult:
    """
    Resource Type definition for AWS::FSx::DataRepositoryAssociation


    :param str association_id: The system-generated, unique ID of the data repository association.
    """
    __args__ = dict()
    __args__['associationId'] = association_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:fsx:getDataRepositoryAssociation', __args__, opts=opts, typ=GetDataRepositoryAssociationResult).value

    return AwaitableGetDataRepositoryAssociationResult(
        association_id=pulumi.get(__ret__, 'association_id'),
        imported_file_chunk_size=pulumi.get(__ret__, 'imported_file_chunk_size'),
        resource_arn=pulumi.get(__ret__, 'resource_arn'),
        s3=pulumi.get(__ret__, 's3'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_data_repository_association)
def get_data_repository_association_output(association_id: Optional[pulumi.Input[str]] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDataRepositoryAssociationResult]:
    """
    Resource Type definition for AWS::FSx::DataRepositoryAssociation


    :param str association_id: The system-generated, unique ID of the data repository association.
    """
    ...
