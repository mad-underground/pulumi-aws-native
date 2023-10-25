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
    'GetModelResult',
    'AwaitableGetModelResult',
    'get_model',
    'get_model_output',
]

@pulumi.output_type
class GetModelResult:
    def __init__(__self__, description=None, schema=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if schema and not isinstance(schema, dict):
            raise TypeError("Expected argument 'schema' to be a dict")
        pulumi.set(__self__, "schema", schema)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the model.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def schema(self) -> Optional[Any]:
        """
        The schema for the model. For ``application/json`` models, this should be JSON schema draft 4 model. Do not include "\\*/" characters in the description of any properties because such "\\*/" characters may be interpreted as the closing marker for comments in some languages, such as Java or JavaScript, causing the installation of your API's SDK generated by API Gateway to fail.
        """
        return pulumi.get(self, "schema")


class AwaitableGetModelResult(GetModelResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetModelResult(
            description=self.description,
            schema=self.schema)


def get_model(name: Optional[str] = None,
              rest_api_id: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetModelResult:
    """
    The ``AWS::ApiGateway::Model`` resource defines the structure of a request or response payload for an API method.


    :param str name: A name for the model. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the model name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
             If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
    :param str rest_api_id: The string identifier of the associated RestApi.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['restApiId'] = rest_api_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:apigateway:getModel', __args__, opts=opts, typ=GetModelResult).value

    return AwaitableGetModelResult(
        description=pulumi.get(__ret__, 'description'),
        schema=pulumi.get(__ret__, 'schema'))


@_utilities.lift_output_func(get_model)
def get_model_output(name: Optional[pulumi.Input[str]] = None,
                     rest_api_id: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetModelResult]:
    """
    The ``AWS::ApiGateway::Model`` resource defines the structure of a request or response payload for an API method.


    :param str name: A name for the model. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the model name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
             If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
    :param str rest_api_id: The string identifier of the associated RestApi.
    """
    ...
