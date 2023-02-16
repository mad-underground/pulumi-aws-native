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
from ._inputs import *

__all__ = ['GraphQLApiArgs', 'GraphQLApi']

@pulumi.input_type
class GraphQLApiArgs:
    def __init__(__self__, *,
                 authentication_type: pulumi.Input[str],
                 additional_authentication_providers: Optional[pulumi.Input[Sequence[pulumi.Input['GraphQLApiAdditionalAuthenticationProviderArgs']]]] = None,
                 lambda_authorizer_config: Optional[pulumi.Input['GraphQLApiLambdaAuthorizerConfigArgs']] = None,
                 log_config: Optional[pulumi.Input['GraphQLApiLogConfigArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 open_id_connect_config: Optional[pulumi.Input['GraphQLApiOpenIDConnectConfigArgs']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['GraphQLApiTagArgs']]]] = None,
                 user_pool_config: Optional[pulumi.Input['GraphQLApiUserPoolConfigArgs']] = None,
                 xray_enabled: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a GraphQLApi resource.
        """
        pulumi.set(__self__, "authentication_type", authentication_type)
        if additional_authentication_providers is not None:
            pulumi.set(__self__, "additional_authentication_providers", additional_authentication_providers)
        if lambda_authorizer_config is not None:
            pulumi.set(__self__, "lambda_authorizer_config", lambda_authorizer_config)
        if log_config is not None:
            pulumi.set(__self__, "log_config", log_config)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if open_id_connect_config is not None:
            pulumi.set(__self__, "open_id_connect_config", open_id_connect_config)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if user_pool_config is not None:
            pulumi.set(__self__, "user_pool_config", user_pool_config)
        if xray_enabled is not None:
            pulumi.set(__self__, "xray_enabled", xray_enabled)

    @property
    @pulumi.getter(name="authenticationType")
    def authentication_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "authentication_type")

    @authentication_type.setter
    def authentication_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "authentication_type", value)

    @property
    @pulumi.getter(name="additionalAuthenticationProviders")
    def additional_authentication_providers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GraphQLApiAdditionalAuthenticationProviderArgs']]]]:
        return pulumi.get(self, "additional_authentication_providers")

    @additional_authentication_providers.setter
    def additional_authentication_providers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GraphQLApiAdditionalAuthenticationProviderArgs']]]]):
        pulumi.set(self, "additional_authentication_providers", value)

    @property
    @pulumi.getter(name="lambdaAuthorizerConfig")
    def lambda_authorizer_config(self) -> Optional[pulumi.Input['GraphQLApiLambdaAuthorizerConfigArgs']]:
        return pulumi.get(self, "lambda_authorizer_config")

    @lambda_authorizer_config.setter
    def lambda_authorizer_config(self, value: Optional[pulumi.Input['GraphQLApiLambdaAuthorizerConfigArgs']]):
        pulumi.set(self, "lambda_authorizer_config", value)

    @property
    @pulumi.getter(name="logConfig")
    def log_config(self) -> Optional[pulumi.Input['GraphQLApiLogConfigArgs']]:
        return pulumi.get(self, "log_config")

    @log_config.setter
    def log_config(self, value: Optional[pulumi.Input['GraphQLApiLogConfigArgs']]):
        pulumi.set(self, "log_config", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="openIDConnectConfig")
    def open_id_connect_config(self) -> Optional[pulumi.Input['GraphQLApiOpenIDConnectConfigArgs']]:
        return pulumi.get(self, "open_id_connect_config")

    @open_id_connect_config.setter
    def open_id_connect_config(self, value: Optional[pulumi.Input['GraphQLApiOpenIDConnectConfigArgs']]):
        pulumi.set(self, "open_id_connect_config", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GraphQLApiTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GraphQLApiTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="userPoolConfig")
    def user_pool_config(self) -> Optional[pulumi.Input['GraphQLApiUserPoolConfigArgs']]:
        return pulumi.get(self, "user_pool_config")

    @user_pool_config.setter
    def user_pool_config(self, value: Optional[pulumi.Input['GraphQLApiUserPoolConfigArgs']]):
        pulumi.set(self, "user_pool_config", value)

    @property
    @pulumi.getter(name="xrayEnabled")
    def xray_enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "xray_enabled")

    @xray_enabled.setter
    def xray_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "xray_enabled", value)


warnings.warn("""GraphQLApi is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class GraphQLApi(pulumi.CustomResource):
    warnings.warn("""GraphQLApi is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_authentication_providers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GraphQLApiAdditionalAuthenticationProviderArgs']]]]] = None,
                 authentication_type: Optional[pulumi.Input[str]] = None,
                 lambda_authorizer_config: Optional[pulumi.Input[pulumi.InputType['GraphQLApiLambdaAuthorizerConfigArgs']]] = None,
                 log_config: Optional[pulumi.Input[pulumi.InputType['GraphQLApiLogConfigArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 open_id_connect_config: Optional[pulumi.Input[pulumi.InputType['GraphQLApiOpenIDConnectConfigArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GraphQLApiTagArgs']]]]] = None,
                 user_pool_config: Optional[pulumi.Input[pulumi.InputType['GraphQLApiUserPoolConfigArgs']]] = None,
                 xray_enabled: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::AppSync::GraphQLApi

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GraphQLApiArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::AppSync::GraphQLApi

        :param str resource_name: The name of the resource.
        :param GraphQLApiArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GraphQLApiArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_authentication_providers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GraphQLApiAdditionalAuthenticationProviderArgs']]]]] = None,
                 authentication_type: Optional[pulumi.Input[str]] = None,
                 lambda_authorizer_config: Optional[pulumi.Input[pulumi.InputType['GraphQLApiLambdaAuthorizerConfigArgs']]] = None,
                 log_config: Optional[pulumi.Input[pulumi.InputType['GraphQLApiLogConfigArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 open_id_connect_config: Optional[pulumi.Input[pulumi.InputType['GraphQLApiOpenIDConnectConfigArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GraphQLApiTagArgs']]]]] = None,
                 user_pool_config: Optional[pulumi.Input[pulumi.InputType['GraphQLApiUserPoolConfigArgs']]] = None,
                 xray_enabled: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        pulumi.log.warn("""GraphQLApi is deprecated: GraphQLApi is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GraphQLApiArgs.__new__(GraphQLApiArgs)

            __props__.__dict__["additional_authentication_providers"] = additional_authentication_providers
            if authentication_type is None and not opts.urn:
                raise TypeError("Missing required property 'authentication_type'")
            __props__.__dict__["authentication_type"] = authentication_type
            __props__.__dict__["lambda_authorizer_config"] = lambda_authorizer_config
            __props__.__dict__["log_config"] = log_config
            __props__.__dict__["name"] = name
            __props__.__dict__["open_id_connect_config"] = open_id_connect_config
            __props__.__dict__["tags"] = tags
            __props__.__dict__["user_pool_config"] = user_pool_config
            __props__.__dict__["xray_enabled"] = xray_enabled
            __props__.__dict__["api_id"] = None
            __props__.__dict__["arn"] = None
            __props__.__dict__["graph_ql_url"] = None
        super(GraphQLApi, __self__).__init__(
            'aws-native:appsync:GraphQLApi',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'GraphQLApi':
        """
        Get an existing GraphQLApi resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = GraphQLApiArgs.__new__(GraphQLApiArgs)

        __props__.__dict__["additional_authentication_providers"] = None
        __props__.__dict__["api_id"] = None
        __props__.__dict__["arn"] = None
        __props__.__dict__["authentication_type"] = None
        __props__.__dict__["graph_ql_url"] = None
        __props__.__dict__["lambda_authorizer_config"] = None
        __props__.__dict__["log_config"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["open_id_connect_config"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["user_pool_config"] = None
        __props__.__dict__["xray_enabled"] = None
        return GraphQLApi(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalAuthenticationProviders")
    def additional_authentication_providers(self) -> pulumi.Output[Optional[Sequence['outputs.GraphQLApiAdditionalAuthenticationProvider']]]:
        return pulumi.get(self, "additional_authentication_providers")

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "api_id")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="authenticationType")
    def authentication_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "authentication_type")

    @property
    @pulumi.getter(name="graphQLUrl")
    def graph_ql_url(self) -> pulumi.Output[str]:
        return pulumi.get(self, "graph_ql_url")

    @property
    @pulumi.getter(name="lambdaAuthorizerConfig")
    def lambda_authorizer_config(self) -> pulumi.Output[Optional['outputs.GraphQLApiLambdaAuthorizerConfig']]:
        return pulumi.get(self, "lambda_authorizer_config")

    @property
    @pulumi.getter(name="logConfig")
    def log_config(self) -> pulumi.Output[Optional['outputs.GraphQLApiLogConfig']]:
        return pulumi.get(self, "log_config")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="openIDConnectConfig")
    def open_id_connect_config(self) -> pulumi.Output[Optional['outputs.GraphQLApiOpenIDConnectConfig']]:
        return pulumi.get(self, "open_id_connect_config")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.GraphQLApiTag']]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="userPoolConfig")
    def user_pool_config(self) -> pulumi.Output[Optional['outputs.GraphQLApiUserPoolConfig']]:
        return pulumi.get(self, "user_pool_config")

    @property
    @pulumi.getter(name="xrayEnabled")
    def xray_enabled(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "xray_enabled")

