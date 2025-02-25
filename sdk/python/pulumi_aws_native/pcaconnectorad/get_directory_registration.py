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
    'GetDirectoryRegistrationResult',
    'AwaitableGetDirectoryRegistrationResult',
    'get_directory_registration',
    'get_directory_registration_output',
]

@pulumi.output_type
class GetDirectoryRegistrationResult:
    def __init__(__self__, directory_registration_arn=None):
        if directory_registration_arn and not isinstance(directory_registration_arn, str):
            raise TypeError("Expected argument 'directory_registration_arn' to be a str")
        pulumi.set(__self__, "directory_registration_arn", directory_registration_arn)

    @property
    @pulumi.getter(name="directoryRegistrationArn")
    def directory_registration_arn(self) -> Optional[str]:
        return pulumi.get(self, "directory_registration_arn")


class AwaitableGetDirectoryRegistrationResult(GetDirectoryRegistrationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDirectoryRegistrationResult(
            directory_registration_arn=self.directory_registration_arn)


def get_directory_registration(directory_registration_arn: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDirectoryRegistrationResult:
    """
    Definition of AWS::PCAConnectorAD::DirectoryRegistration Resource Type
    """
    __args__ = dict()
    __args__['directoryRegistrationArn'] = directory_registration_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:pcaconnectorad:getDirectoryRegistration', __args__, opts=opts, typ=GetDirectoryRegistrationResult).value

    return AwaitableGetDirectoryRegistrationResult(
        directory_registration_arn=pulumi.get(__ret__, 'directory_registration_arn'))


@_utilities.lift_output_func(get_directory_registration)
def get_directory_registration_output(directory_registration_arn: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDirectoryRegistrationResult]:
    """
    Definition of AWS::PCAConnectorAD::DirectoryRegistration Resource Type
    """
    ...
