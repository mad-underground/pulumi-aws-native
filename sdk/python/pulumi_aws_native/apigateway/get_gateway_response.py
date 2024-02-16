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
    'GetGatewayResponseResult',
    'AwaitableGetGatewayResponseResult',
    'get_gateway_response',
    'get_gateway_response_output',
]

@pulumi.output_type
class GetGatewayResponseResult:
    def __init__(__self__, id=None, response_parameters=None, response_templates=None, status_code=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if response_parameters and not isinstance(response_parameters, dict):
            raise TypeError("Expected argument 'response_parameters' to be a dict")
        pulumi.set(__self__, "response_parameters", response_parameters)
        if response_templates and not isinstance(response_templates, dict):
            raise TypeError("Expected argument 'response_templates' to be a dict")
        pulumi.set(__self__, "response_templates", response_templates)
        if status_code and not isinstance(status_code, str):
            raise TypeError("Expected argument 'status_code' to be a str")
        pulumi.set(__self__, "status_code", status_code)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="responseParameters")
    def response_parameters(self) -> Optional[Mapping[str, str]]:
        """
        Response parameters (paths, query strings and headers) of the GatewayResponse as a string-to-string map of key-value pairs.
        """
        return pulumi.get(self, "response_parameters")

    @property
    @pulumi.getter(name="responseTemplates")
    def response_templates(self) -> Optional[Mapping[str, str]]:
        """
        Response templates of the GatewayResponse as a string-to-string map of key-value pairs.
        """
        return pulumi.get(self, "response_templates")

    @property
    @pulumi.getter(name="statusCode")
    def status_code(self) -> Optional[str]:
        """
        The HTTP status code for this GatewayResponse.
        """
        return pulumi.get(self, "status_code")


class AwaitableGetGatewayResponseResult(GetGatewayResponseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGatewayResponseResult(
            id=self.id,
            response_parameters=self.response_parameters,
            response_templates=self.response_templates,
            status_code=self.status_code)


def get_gateway_response(id: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGatewayResponseResult:
    """
    The ``AWS::ApiGateway::GatewayResponse`` resource creates a gateway response for your API. For more information, see [API Gateway Responses](https://docs.aws.amazon.com/apigateway/latest/developerguide/customize-gateway-responses.html#api-gateway-gatewayResponse-definition) in the *API Gateway Developer Guide*.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:apigateway:getGatewayResponse', __args__, opts=opts, typ=GetGatewayResponseResult).value

    return AwaitableGetGatewayResponseResult(
        id=pulumi.get(__ret__, 'id'),
        response_parameters=pulumi.get(__ret__, 'response_parameters'),
        response_templates=pulumi.get(__ret__, 'response_templates'),
        status_code=pulumi.get(__ret__, 'status_code'))


@_utilities.lift_output_func(get_gateway_response)
def get_gateway_response_output(id: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGatewayResponseResult]:
    """
    The ``AWS::ApiGateway::GatewayResponse`` resource creates a gateway response for your API. For more information, see [API Gateway Responses](https://docs.aws.amazon.com/apigateway/latest/developerguide/customize-gateway-responses.html#api-gateway-gatewayResponse-definition) in the *API Gateway Developer Guide*.
    """
    ...
