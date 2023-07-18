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
    'GetEnvironmentAccountConnectionResult',
    'AwaitableGetEnvironmentAccountConnectionResult',
    'get_environment_account_connection',
    'get_environment_account_connection_output',
]

@pulumi.output_type
class GetEnvironmentAccountConnectionResult:
    def __init__(__self__, arn=None, codebuild_role_arn=None, component_role_arn=None, environment_account_id=None, environment_name=None, id=None, management_account_id=None, role_arn=None, status=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if codebuild_role_arn and not isinstance(codebuild_role_arn, str):
            raise TypeError("Expected argument 'codebuild_role_arn' to be a str")
        pulumi.set(__self__, "codebuild_role_arn", codebuild_role_arn)
        if component_role_arn and not isinstance(component_role_arn, str):
            raise TypeError("Expected argument 'component_role_arn' to be a str")
        pulumi.set(__self__, "component_role_arn", component_role_arn)
        if environment_account_id and not isinstance(environment_account_id, str):
            raise TypeError("Expected argument 'environment_account_id' to be a str")
        pulumi.set(__self__, "environment_account_id", environment_account_id)
        if environment_name and not isinstance(environment_name, str):
            raise TypeError("Expected argument 'environment_name' to be a str")
        pulumi.set(__self__, "environment_name", environment_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if management_account_id and not isinstance(management_account_id, str):
            raise TypeError("Expected argument 'management_account_id' to be a str")
        pulumi.set(__self__, "management_account_id", management_account_id)
        if role_arn and not isinstance(role_arn, str):
            raise TypeError("Expected argument 'role_arn' to be a str")
        pulumi.set(__self__, "role_arn", role_arn)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the environment account connection.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="codebuildRoleArn")
    def codebuild_role_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of an IAM service role in the environment account. AWS Proton uses this role to provision infrastructure resources using CodeBuild-based provisioning in the associated environment account.
        """
        return pulumi.get(self, "codebuild_role_arn")

    @property
    @pulumi.getter(name="componentRoleArn")
    def component_role_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the IAM service role that AWS Proton uses when provisioning directly defined components in the associated environment account. It determines the scope of infrastructure that a component can provision in the account.
        """
        return pulumi.get(self, "component_role_arn")

    @property
    @pulumi.getter(name="environmentAccountId")
    def environment_account_id(self) -> Optional[str]:
        """
        The environment account that's connected to the environment account connection.
        """
        return pulumi.get(self, "environment_account_id")

    @property
    @pulumi.getter(name="environmentName")
    def environment_name(self) -> Optional[str]:
        """
        The name of the AWS Proton environment that's created in the associated management account.
        """
        return pulumi.get(self, "environment_name")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The ID of the environment account connection.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="managementAccountId")
    def management_account_id(self) -> Optional[str]:
        """
        The ID of the management account that accepts or rejects the environment account connection. You create an manage the AWS Proton environment in this account. If the management account accepts the environment account connection, AWS Proton can use the associated IAM role to provision environment infrastructure resources in the associated environment account.
        """
        return pulumi.get(self, "management_account_id")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the IAM service role that's created in the environment account. AWS Proton uses this role to provision infrastructure resources in the associated environment account.
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter
    def status(self) -> Optional['EnvironmentAccountConnectionStatus']:
        """
        The status of the environment account connection.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.EnvironmentAccountConnectionTag']]:
        """
        <p>An optional list of metadata items that you can associate with the Proton environment account connection. A tag is a key-value pair.</p>
                 <p>For more information, see <a href="https://docs.aws.amazon.com/proton/latest/userguide/resources.html">Proton resources and tagging</a> in the
                <i>Proton User Guide</i>.</p>
        """
        return pulumi.get(self, "tags")


class AwaitableGetEnvironmentAccountConnectionResult(GetEnvironmentAccountConnectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEnvironmentAccountConnectionResult(
            arn=self.arn,
            codebuild_role_arn=self.codebuild_role_arn,
            component_role_arn=self.component_role_arn,
            environment_account_id=self.environment_account_id,
            environment_name=self.environment_name,
            id=self.id,
            management_account_id=self.management_account_id,
            role_arn=self.role_arn,
            status=self.status,
            tags=self.tags)


def get_environment_account_connection(arn: Optional[str] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEnvironmentAccountConnectionResult:
    """
    Resource Schema describing various properties for AWS Proton Environment Account Connections resources.


    :param str arn: The Amazon Resource Name (ARN) of the environment account connection.
    """
    __args__ = dict()
    __args__['arn'] = arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:proton:getEnvironmentAccountConnection', __args__, opts=opts, typ=GetEnvironmentAccountConnectionResult).value

    return AwaitableGetEnvironmentAccountConnectionResult(
        arn=pulumi.get(__ret__, 'arn'),
        codebuild_role_arn=pulumi.get(__ret__, 'codebuild_role_arn'),
        component_role_arn=pulumi.get(__ret__, 'component_role_arn'),
        environment_account_id=pulumi.get(__ret__, 'environment_account_id'),
        environment_name=pulumi.get(__ret__, 'environment_name'),
        id=pulumi.get(__ret__, 'id'),
        management_account_id=pulumi.get(__ret__, 'management_account_id'),
        role_arn=pulumi.get(__ret__, 'role_arn'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_environment_account_connection)
def get_environment_account_connection_output(arn: Optional[pulumi.Input[str]] = None,
                                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEnvironmentAccountConnectionResult]:
    """
    Resource Schema describing various properties for AWS Proton Environment Account Connections resources.


    :param str arn: The Amazon Resource Name (ARN) of the environment account connection.
    """
    ...
