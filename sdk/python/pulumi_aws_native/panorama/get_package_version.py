# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = [
    'GetPackageVersionResult',
    'AwaitableGetPackageVersionResult',
    'get_package_version',
    'get_package_version_output',
]

@pulumi.output_type
class GetPackageVersionResult:
    def __init__(__self__, is_latest_patch=None, mark_latest=None, package_arn=None, package_name=None, registered_time=None, status=None, status_description=None, updated_latest_patch_version=None):
        if is_latest_patch and not isinstance(is_latest_patch, bool):
            raise TypeError("Expected argument 'is_latest_patch' to be a bool")
        pulumi.set(__self__, "is_latest_patch", is_latest_patch)
        if mark_latest and not isinstance(mark_latest, bool):
            raise TypeError("Expected argument 'mark_latest' to be a bool")
        pulumi.set(__self__, "mark_latest", mark_latest)
        if package_arn and not isinstance(package_arn, str):
            raise TypeError("Expected argument 'package_arn' to be a str")
        pulumi.set(__self__, "package_arn", package_arn)
        if package_name and not isinstance(package_name, str):
            raise TypeError("Expected argument 'package_name' to be a str")
        pulumi.set(__self__, "package_name", package_name)
        if registered_time and not isinstance(registered_time, int):
            raise TypeError("Expected argument 'registered_time' to be a int")
        pulumi.set(__self__, "registered_time", registered_time)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if status_description and not isinstance(status_description, str):
            raise TypeError("Expected argument 'status_description' to be a str")
        pulumi.set(__self__, "status_description", status_description)
        if updated_latest_patch_version and not isinstance(updated_latest_patch_version, str):
            raise TypeError("Expected argument 'updated_latest_patch_version' to be a str")
        pulumi.set(__self__, "updated_latest_patch_version", updated_latest_patch_version)

    @property
    @pulumi.getter(name="isLatestPatch")
    def is_latest_patch(self) -> Optional[bool]:
        return pulumi.get(self, "is_latest_patch")

    @property
    @pulumi.getter(name="markLatest")
    def mark_latest(self) -> Optional[bool]:
        return pulumi.get(self, "mark_latest")

    @property
    @pulumi.getter(name="packageArn")
    def package_arn(self) -> Optional[str]:
        return pulumi.get(self, "package_arn")

    @property
    @pulumi.getter(name="packageName")
    def package_name(self) -> Optional[str]:
        return pulumi.get(self, "package_name")

    @property
    @pulumi.getter(name="registeredTime")
    def registered_time(self) -> Optional[int]:
        return pulumi.get(self, "registered_time")

    @property
    @pulumi.getter
    def status(self) -> Optional['PackageVersionStatus']:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusDescription")
    def status_description(self) -> Optional[str]:
        return pulumi.get(self, "status_description")

    @property
    @pulumi.getter(name="updatedLatestPatchVersion")
    def updated_latest_patch_version(self) -> Optional[str]:
        return pulumi.get(self, "updated_latest_patch_version")


class AwaitableGetPackageVersionResult(GetPackageVersionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPackageVersionResult(
            is_latest_patch=self.is_latest_patch,
            mark_latest=self.mark_latest,
            package_arn=self.package_arn,
            package_name=self.package_name,
            registered_time=self.registered_time,
            status=self.status,
            status_description=self.status_description,
            updated_latest_patch_version=self.updated_latest_patch_version)


def get_package_version(package_id: Optional[str] = None,
                        package_version: Optional[str] = None,
                        patch_version: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPackageVersionResult:
    """
    Schema for PackageVersion Resource Type
    """
    __args__ = dict()
    __args__['packageId'] = package_id
    __args__['packageVersion'] = package_version
    __args__['patchVersion'] = patch_version
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:panorama:getPackageVersion', __args__, opts=opts, typ=GetPackageVersionResult).value

    return AwaitableGetPackageVersionResult(
        is_latest_patch=__ret__.is_latest_patch,
        mark_latest=__ret__.mark_latest,
        package_arn=__ret__.package_arn,
        package_name=__ret__.package_name,
        registered_time=__ret__.registered_time,
        status=__ret__.status,
        status_description=__ret__.status_description,
        updated_latest_patch_version=__ret__.updated_latest_patch_version)


@_utilities.lift_output_func(get_package_version)
def get_package_version_output(package_id: Optional[pulumi.Input[str]] = None,
                               package_version: Optional[pulumi.Input[str]] = None,
                               patch_version: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPackageVersionResult]:
    """
    Schema for PackageVersion Resource Type
    """
    ...
