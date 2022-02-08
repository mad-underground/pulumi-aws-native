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
    'GetWorkspaceResult',
    'AwaitableGetWorkspaceResult',
    'get_workspace',
    'get_workspace_output',
]

@pulumi.output_type
class GetWorkspaceResult:
    def __init__(__self__, bundle_id=None, directory_id=None, id=None, root_volume_encryption_enabled=None, tags=None, user_volume_encryption_enabled=None, volume_encryption_key=None, workspace_properties=None):
        if bundle_id and not isinstance(bundle_id, str):
            raise TypeError("Expected argument 'bundle_id' to be a str")
        pulumi.set(__self__, "bundle_id", bundle_id)
        if directory_id and not isinstance(directory_id, str):
            raise TypeError("Expected argument 'directory_id' to be a str")
        pulumi.set(__self__, "directory_id", directory_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if root_volume_encryption_enabled and not isinstance(root_volume_encryption_enabled, bool):
            raise TypeError("Expected argument 'root_volume_encryption_enabled' to be a bool")
        pulumi.set(__self__, "root_volume_encryption_enabled", root_volume_encryption_enabled)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if user_volume_encryption_enabled and not isinstance(user_volume_encryption_enabled, bool):
            raise TypeError("Expected argument 'user_volume_encryption_enabled' to be a bool")
        pulumi.set(__self__, "user_volume_encryption_enabled", user_volume_encryption_enabled)
        if volume_encryption_key and not isinstance(volume_encryption_key, str):
            raise TypeError("Expected argument 'volume_encryption_key' to be a str")
        pulumi.set(__self__, "volume_encryption_key", volume_encryption_key)
        if workspace_properties and not isinstance(workspace_properties, dict):
            raise TypeError("Expected argument 'workspace_properties' to be a dict")
        pulumi.set(__self__, "workspace_properties", workspace_properties)

    @property
    @pulumi.getter(name="bundleId")
    def bundle_id(self) -> Optional[str]:
        return pulumi.get(self, "bundle_id")

    @property
    @pulumi.getter(name="directoryId")
    def directory_id(self) -> Optional[str]:
        return pulumi.get(self, "directory_id")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="rootVolumeEncryptionEnabled")
    def root_volume_encryption_enabled(self) -> Optional[bool]:
        return pulumi.get(self, "root_volume_encryption_enabled")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.WorkspaceTag']]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="userVolumeEncryptionEnabled")
    def user_volume_encryption_enabled(self) -> Optional[bool]:
        return pulumi.get(self, "user_volume_encryption_enabled")

    @property
    @pulumi.getter(name="volumeEncryptionKey")
    def volume_encryption_key(self) -> Optional[str]:
        return pulumi.get(self, "volume_encryption_key")

    @property
    @pulumi.getter(name="workspaceProperties")
    def workspace_properties(self) -> Optional['outputs.WorkspaceProperties']:
        return pulumi.get(self, "workspace_properties")


class AwaitableGetWorkspaceResult(GetWorkspaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWorkspaceResult(
            bundle_id=self.bundle_id,
            directory_id=self.directory_id,
            id=self.id,
            root_volume_encryption_enabled=self.root_volume_encryption_enabled,
            tags=self.tags,
            user_volume_encryption_enabled=self.user_volume_encryption_enabled,
            volume_encryption_key=self.volume_encryption_key,
            workspace_properties=self.workspace_properties)


def get_workspace(id: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWorkspaceResult:
    """
    Resource Type definition for AWS::WorkSpaces::Workspace
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:workspaces:getWorkspace', __args__, opts=opts, typ=GetWorkspaceResult).value

    return AwaitableGetWorkspaceResult(
        bundle_id=__ret__.bundle_id,
        directory_id=__ret__.directory_id,
        id=__ret__.id,
        root_volume_encryption_enabled=__ret__.root_volume_encryption_enabled,
        tags=__ret__.tags,
        user_volume_encryption_enabled=__ret__.user_volume_encryption_enabled,
        volume_encryption_key=__ret__.volume_encryption_key,
        workspace_properties=__ret__.workspace_properties)


@_utilities.lift_output_func(get_workspace)
def get_workspace_output(id: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetWorkspaceResult]:
    """
    Resource Type definition for AWS::WorkSpaces::Workspace
    """
    ...
