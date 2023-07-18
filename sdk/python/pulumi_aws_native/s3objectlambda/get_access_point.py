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

__all__ = [
    'GetAccessPointResult',
    'AwaitableGetAccessPointResult',
    'get_access_point',
    'get_access_point_output',
]

@pulumi.output_type
class GetAccessPointResult:
    def __init__(__self__, alias=None, arn=None, creation_date=None, object_lambda_configuration=None, policy_status=None, public_access_block_configuration=None):
        if alias and not isinstance(alias, dict):
            raise TypeError("Expected argument 'alias' to be a dict")
        pulumi.set(__self__, "alias", alias)
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if creation_date and not isinstance(creation_date, str):
            raise TypeError("Expected argument 'creation_date' to be a str")
        pulumi.set(__self__, "creation_date", creation_date)
        if object_lambda_configuration and not isinstance(object_lambda_configuration, dict):
            raise TypeError("Expected argument 'object_lambda_configuration' to be a dict")
        pulumi.set(__self__, "object_lambda_configuration", object_lambda_configuration)
        if policy_status and not isinstance(policy_status, dict):
            raise TypeError("Expected argument 'policy_status' to be a dict")
        pulumi.set(__self__, "policy_status", policy_status)
        if public_access_block_configuration and not isinstance(public_access_block_configuration, dict):
            raise TypeError("Expected argument 'public_access_block_configuration' to be a dict")
        pulumi.set(__self__, "public_access_block_configuration", public_access_block_configuration)

    @property
    @pulumi.getter
    def alias(self) -> Optional['outputs.AccessPointAlias']:
        return pulumi.get(self, "alias")

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> Optional[str]:
        """
        The date and time when the Object lambda Access Point was created.
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter(name="objectLambdaConfiguration")
    def object_lambda_configuration(self) -> Optional['outputs.AccessPointObjectLambdaConfiguration']:
        """
        The Object lambda Access Point Configuration that configures transformations to be applied on the objects on specified S3 Actions
        """
        return pulumi.get(self, "object_lambda_configuration")

    @property
    @pulumi.getter(name="policyStatus")
    def policy_status(self) -> Optional['outputs.AccessPointPolicyStatus']:
        return pulumi.get(self, "policy_status")

    @property
    @pulumi.getter(name="publicAccessBlockConfiguration")
    def public_access_block_configuration(self) -> Optional['outputs.AccessPointPublicAccessBlockConfiguration']:
        """
        The PublicAccessBlock configuration that you want to apply to this Access Point. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide.
        """
        return pulumi.get(self, "public_access_block_configuration")


class AwaitableGetAccessPointResult(GetAccessPointResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccessPointResult(
            alias=self.alias,
            arn=self.arn,
            creation_date=self.creation_date,
            object_lambda_configuration=self.object_lambda_configuration,
            policy_status=self.policy_status,
            public_access_block_configuration=self.public_access_block_configuration)


def get_access_point(name: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccessPointResult:
    """
    The AWS::S3ObjectLambda::AccessPoint resource is an Amazon S3ObjectLambda resource type that you can use to add computation to S3 actions


    :param str name: The name you want to assign to this Object lambda Access Point.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:s3objectlambda:getAccessPoint', __args__, opts=opts, typ=GetAccessPointResult).value

    return AwaitableGetAccessPointResult(
        alias=pulumi.get(__ret__, 'alias'),
        arn=pulumi.get(__ret__, 'arn'),
        creation_date=pulumi.get(__ret__, 'creation_date'),
        object_lambda_configuration=pulumi.get(__ret__, 'object_lambda_configuration'),
        policy_status=pulumi.get(__ret__, 'policy_status'),
        public_access_block_configuration=pulumi.get(__ret__, 'public_access_block_configuration'))


@_utilities.lift_output_func(get_access_point)
def get_access_point_output(name: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAccessPointResult]:
    """
    The AWS::S3ObjectLambda::AccessPoint resource is an Amazon S3ObjectLambda resource type that you can use to add computation to S3 actions


    :param str name: The name you want to assign to this Object lambda Access Point.
    """
    ...
