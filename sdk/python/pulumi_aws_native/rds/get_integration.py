# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from .. import outputs as _root_outputs

__all__ = [
    'GetIntegrationResult',
    'AwaitableGetIntegrationResult',
    'get_integration',
    'get_integration_output',
]

@pulumi.output_type
class GetIntegrationResult:
    def __init__(__self__, create_time=None, integration_arn=None, tags=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if integration_arn and not isinstance(integration_arn, str):
            raise TypeError("Expected argument 'integration_arn' to be a str")
        pulumi.set(__self__, "integration_arn", integration_arn)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[str]:
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="integrationArn")
    def integration_arn(self) -> Optional[str]:
        """
        The ARN of the integration.
        """
        return pulumi.get(self, "integration_arn")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")


class AwaitableGetIntegrationResult(GetIntegrationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIntegrationResult(
            create_time=self.create_time,
            integration_arn=self.integration_arn,
            tags=self.tags)


def get_integration(integration_arn: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIntegrationResult:
    """
    Creates a zero-ETL integration with Amazon Redshift.


    :param str integration_arn: The ARN of the integration.
    """
    __args__ = dict()
    __args__['integrationArn'] = integration_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:rds:getIntegration', __args__, opts=opts, typ=GetIntegrationResult).value

    return AwaitableGetIntegrationResult(
        create_time=pulumi.get(__ret__, 'create_time'),
        integration_arn=pulumi.get(__ret__, 'integration_arn'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_integration)
def get_integration_output(integration_arn: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIntegrationResult]:
    """
    Creates a zero-ETL integration with Amazon Redshift.


    :param str integration_arn: The ARN of the integration.
    """
    ...
