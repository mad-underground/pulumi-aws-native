# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetDeploymentResult',
    'AwaitableGetDeploymentResult',
    'get_deployment',
    'get_deployment_output',
]

@pulumi.output_type
class GetDeploymentResult:
    def __init__(__self__, deployment_id=None, description=None):
        if deployment_id and not isinstance(deployment_id, str):
            raise TypeError("Expected argument 'deployment_id' to be a str")
        pulumi.set(__self__, "deployment_id", deployment_id)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="deploymentId")
    def deployment_id(self) -> Optional[str]:
        """
        Primary Id for this resource
        """
        return pulumi.get(self, "deployment_id")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description for the Deployment resource to create.
        """
        return pulumi.get(self, "description")


class AwaitableGetDeploymentResult(GetDeploymentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDeploymentResult(
            deployment_id=self.deployment_id,
            description=self.description)


def get_deployment(deployment_id: Optional[str] = None,
                   rest_api_id: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDeploymentResult:
    """
    The ``AWS::ApiGateway::Deployment`` resource deploys an API Gateway ``RestApi`` resource to a stage so that clients can call the API over the internet. The stage acts as an environment.


    :param str deployment_id: Primary Id for this resource
    :param str rest_api_id: The string identifier of the associated RestApi.
    """
    __args__ = dict()
    __args__['deploymentId'] = deployment_id
    __args__['restApiId'] = rest_api_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:apigateway:getDeployment', __args__, opts=opts, typ=GetDeploymentResult).value

    return AwaitableGetDeploymentResult(
        deployment_id=pulumi.get(__ret__, 'deployment_id'),
        description=pulumi.get(__ret__, 'description'))


@_utilities.lift_output_func(get_deployment)
def get_deployment_output(deployment_id: Optional[pulumi.Input[str]] = None,
                          rest_api_id: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDeploymentResult]:
    """
    The ``AWS::ApiGateway::Deployment`` resource deploys an API Gateway ``RestApi`` resource to a stage so that clients can call the API over the internet. The stage acts as an environment.


    :param str deployment_id: Primary Id for this resource
    :param str rest_api_id: The string identifier of the associated RestApi.
    """
    ...
