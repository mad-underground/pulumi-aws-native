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
    'GetGraphQLApiResult',
    'AwaitableGetGraphQLApiResult',
    'get_graph_ql_api',
    'get_graph_ql_api_output',
]

@pulumi.output_type
class GetGraphQLApiResult:
    def __init__(__self__, additional_authentication_providers=None, api_id=None, arn=None, authentication_type=None, graph_ql_url=None, id=None, lambda_authorizer_config=None, log_config=None, name=None, open_id_connect_config=None, tags=None, user_pool_config=None, xray_enabled=None):
        if additional_authentication_providers and not isinstance(additional_authentication_providers, dict):
            raise TypeError("Expected argument 'additional_authentication_providers' to be a dict")
        pulumi.set(__self__, "additional_authentication_providers", additional_authentication_providers)
        if api_id and not isinstance(api_id, str):
            raise TypeError("Expected argument 'api_id' to be a str")
        pulumi.set(__self__, "api_id", api_id)
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if authentication_type and not isinstance(authentication_type, str):
            raise TypeError("Expected argument 'authentication_type' to be a str")
        pulumi.set(__self__, "authentication_type", authentication_type)
        if graph_ql_url and not isinstance(graph_ql_url, str):
            raise TypeError("Expected argument 'graph_ql_url' to be a str")
        pulumi.set(__self__, "graph_ql_url", graph_ql_url)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lambda_authorizer_config and not isinstance(lambda_authorizer_config, dict):
            raise TypeError("Expected argument 'lambda_authorizer_config' to be a dict")
        pulumi.set(__self__, "lambda_authorizer_config", lambda_authorizer_config)
        if log_config and not isinstance(log_config, dict):
            raise TypeError("Expected argument 'log_config' to be a dict")
        pulumi.set(__self__, "log_config", log_config)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if open_id_connect_config and not isinstance(open_id_connect_config, dict):
            raise TypeError("Expected argument 'open_id_connect_config' to be a dict")
        pulumi.set(__self__, "open_id_connect_config", open_id_connect_config)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if user_pool_config and not isinstance(user_pool_config, dict):
            raise TypeError("Expected argument 'user_pool_config' to be a dict")
        pulumi.set(__self__, "user_pool_config", user_pool_config)
        if xray_enabled and not isinstance(xray_enabled, bool):
            raise TypeError("Expected argument 'xray_enabled' to be a bool")
        pulumi.set(__self__, "xray_enabled", xray_enabled)

    @property
    @pulumi.getter(name="additionalAuthenticationProviders")
    def additional_authentication_providers(self) -> Optional['outputs.GraphQLApiAdditionalAuthenticationProviders']:
        return pulumi.get(self, "additional_authentication_providers")

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> Optional[str]:
        return pulumi.get(self, "api_id")

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="authenticationType")
    def authentication_type(self) -> Optional[str]:
        return pulumi.get(self, "authentication_type")

    @property
    @pulumi.getter(name="graphQLUrl")
    def graph_ql_url(self) -> Optional[str]:
        return pulumi.get(self, "graph_ql_url")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lambdaAuthorizerConfig")
    def lambda_authorizer_config(self) -> Optional['outputs.GraphQLApiLambdaAuthorizerConfig']:
        return pulumi.get(self, "lambda_authorizer_config")

    @property
    @pulumi.getter(name="logConfig")
    def log_config(self) -> Optional['outputs.GraphQLApiLogConfig']:
        return pulumi.get(self, "log_config")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="openIDConnectConfig")
    def open_id_connect_config(self) -> Optional['outputs.GraphQLApiOpenIDConnectConfig']:
        return pulumi.get(self, "open_id_connect_config")

    @property
    @pulumi.getter
    def tags(self) -> Optional['outputs.GraphQLApiTags']:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="userPoolConfig")
    def user_pool_config(self) -> Optional['outputs.GraphQLApiUserPoolConfig']:
        return pulumi.get(self, "user_pool_config")

    @property
    @pulumi.getter(name="xrayEnabled")
    def xray_enabled(self) -> Optional[bool]:
        return pulumi.get(self, "xray_enabled")


class AwaitableGetGraphQLApiResult(GetGraphQLApiResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGraphQLApiResult(
            additional_authentication_providers=self.additional_authentication_providers,
            api_id=self.api_id,
            arn=self.arn,
            authentication_type=self.authentication_type,
            graph_ql_url=self.graph_ql_url,
            id=self.id,
            lambda_authorizer_config=self.lambda_authorizer_config,
            log_config=self.log_config,
            name=self.name,
            open_id_connect_config=self.open_id_connect_config,
            tags=self.tags,
            user_pool_config=self.user_pool_config,
            xray_enabled=self.xray_enabled)


def get_graph_ql_api(id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGraphQLApiResult:
    """
    Resource Type definition for AWS::AppSync::GraphQLApi
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:appsync:getGraphQLApi', __args__, opts=opts, typ=GetGraphQLApiResult).value

    return AwaitableGetGraphQLApiResult(
        additional_authentication_providers=__ret__.additional_authentication_providers,
        api_id=__ret__.api_id,
        arn=__ret__.arn,
        authentication_type=__ret__.authentication_type,
        graph_ql_url=__ret__.graph_ql_url,
        id=__ret__.id,
        lambda_authorizer_config=__ret__.lambda_authorizer_config,
        log_config=__ret__.log_config,
        name=__ret__.name,
        open_id_connect_config=__ret__.open_id_connect_config,
        tags=__ret__.tags,
        user_pool_config=__ret__.user_pool_config,
        xray_enabled=__ret__.xray_enabled)


@_utilities.lift_output_func(get_graph_ql_api)
def get_graph_ql_api_output(id: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGraphQLApiResult]:
    """
    Resource Type definition for AWS::AppSync::GraphQLApi
    """
    ...
