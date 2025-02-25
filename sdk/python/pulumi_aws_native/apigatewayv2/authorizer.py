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

__all__ = ['AuthorizerArgs', 'Authorizer']

@pulumi.input_type
class AuthorizerArgs:
    def __init__(__self__, *,
                 api_id: pulumi.Input[str],
                 authorizer_type: pulumi.Input[str],
                 authorizer_credentials_arn: Optional[pulumi.Input[str]] = None,
                 authorizer_payload_format_version: Optional[pulumi.Input[str]] = None,
                 authorizer_result_ttl_in_seconds: Optional[pulumi.Input[int]] = None,
                 authorizer_uri: Optional[pulumi.Input[str]] = None,
                 enable_simple_responses: Optional[pulumi.Input[bool]] = None,
                 identity_source: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 identity_validation_expression: Optional[pulumi.Input[str]] = None,
                 jwt_configuration: Optional[pulumi.Input['AuthorizerJwtConfigurationArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Authorizer resource.
        """
        pulumi.set(__self__, "api_id", api_id)
        pulumi.set(__self__, "authorizer_type", authorizer_type)
        if authorizer_credentials_arn is not None:
            pulumi.set(__self__, "authorizer_credentials_arn", authorizer_credentials_arn)
        if authorizer_payload_format_version is not None:
            pulumi.set(__self__, "authorizer_payload_format_version", authorizer_payload_format_version)
        if authorizer_result_ttl_in_seconds is not None:
            pulumi.set(__self__, "authorizer_result_ttl_in_seconds", authorizer_result_ttl_in_seconds)
        if authorizer_uri is not None:
            pulumi.set(__self__, "authorizer_uri", authorizer_uri)
        if enable_simple_responses is not None:
            pulumi.set(__self__, "enable_simple_responses", enable_simple_responses)
        if identity_source is not None:
            pulumi.set(__self__, "identity_source", identity_source)
        if identity_validation_expression is not None:
            pulumi.set(__self__, "identity_validation_expression", identity_validation_expression)
        if jwt_configuration is not None:
            pulumi.set(__self__, "jwt_configuration", jwt_configuration)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "api_id")

    @api_id.setter
    def api_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_id", value)

    @property
    @pulumi.getter(name="authorizerType")
    def authorizer_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "authorizer_type")

    @authorizer_type.setter
    def authorizer_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "authorizer_type", value)

    @property
    @pulumi.getter(name="authorizerCredentialsArn")
    def authorizer_credentials_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "authorizer_credentials_arn")

    @authorizer_credentials_arn.setter
    def authorizer_credentials_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorizer_credentials_arn", value)

    @property
    @pulumi.getter(name="authorizerPayloadFormatVersion")
    def authorizer_payload_format_version(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "authorizer_payload_format_version")

    @authorizer_payload_format_version.setter
    def authorizer_payload_format_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorizer_payload_format_version", value)

    @property
    @pulumi.getter(name="authorizerResultTtlInSeconds")
    def authorizer_result_ttl_in_seconds(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "authorizer_result_ttl_in_seconds")

    @authorizer_result_ttl_in_seconds.setter
    def authorizer_result_ttl_in_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "authorizer_result_ttl_in_seconds", value)

    @property
    @pulumi.getter(name="authorizerUri")
    def authorizer_uri(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "authorizer_uri")

    @authorizer_uri.setter
    def authorizer_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorizer_uri", value)

    @property
    @pulumi.getter(name="enableSimpleResponses")
    def enable_simple_responses(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enable_simple_responses")

    @enable_simple_responses.setter
    def enable_simple_responses(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_simple_responses", value)

    @property
    @pulumi.getter(name="identitySource")
    def identity_source(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "identity_source")

    @identity_source.setter
    def identity_source(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "identity_source", value)

    @property
    @pulumi.getter(name="identityValidationExpression")
    def identity_validation_expression(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "identity_validation_expression")

    @identity_validation_expression.setter
    def identity_validation_expression(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "identity_validation_expression", value)

    @property
    @pulumi.getter(name="jwtConfiguration")
    def jwt_configuration(self) -> Optional[pulumi.Input['AuthorizerJwtConfigurationArgs']]:
        return pulumi.get(self, "jwt_configuration")

    @jwt_configuration.setter
    def jwt_configuration(self, value: Optional[pulumi.Input['AuthorizerJwtConfigurationArgs']]):
        pulumi.set(self, "jwt_configuration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class Authorizer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 authorizer_credentials_arn: Optional[pulumi.Input[str]] = None,
                 authorizer_payload_format_version: Optional[pulumi.Input[str]] = None,
                 authorizer_result_ttl_in_seconds: Optional[pulumi.Input[int]] = None,
                 authorizer_type: Optional[pulumi.Input[str]] = None,
                 authorizer_uri: Optional[pulumi.Input[str]] = None,
                 enable_simple_responses: Optional[pulumi.Input[bool]] = None,
                 identity_source: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 identity_validation_expression: Optional[pulumi.Input[str]] = None,
                 jwt_configuration: Optional[pulumi.Input[pulumi.InputType['AuthorizerJwtConfigurationArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::ApiGatewayV2::Authorizer

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AuthorizerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::ApiGatewayV2::Authorizer

        :param str resource_name: The name of the resource.
        :param AuthorizerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AuthorizerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 authorizer_credentials_arn: Optional[pulumi.Input[str]] = None,
                 authorizer_payload_format_version: Optional[pulumi.Input[str]] = None,
                 authorizer_result_ttl_in_seconds: Optional[pulumi.Input[int]] = None,
                 authorizer_type: Optional[pulumi.Input[str]] = None,
                 authorizer_uri: Optional[pulumi.Input[str]] = None,
                 enable_simple_responses: Optional[pulumi.Input[bool]] = None,
                 identity_source: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 identity_validation_expression: Optional[pulumi.Input[str]] = None,
                 jwt_configuration: Optional[pulumi.Input[pulumi.InputType['AuthorizerJwtConfigurationArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AuthorizerArgs.__new__(AuthorizerArgs)

            if api_id is None and not opts.urn:
                raise TypeError("Missing required property 'api_id'")
            __props__.__dict__["api_id"] = api_id
            __props__.__dict__["authorizer_credentials_arn"] = authorizer_credentials_arn
            __props__.__dict__["authorizer_payload_format_version"] = authorizer_payload_format_version
            __props__.__dict__["authorizer_result_ttl_in_seconds"] = authorizer_result_ttl_in_seconds
            if authorizer_type is None and not opts.urn:
                raise TypeError("Missing required property 'authorizer_type'")
            __props__.__dict__["authorizer_type"] = authorizer_type
            __props__.__dict__["authorizer_uri"] = authorizer_uri
            __props__.__dict__["enable_simple_responses"] = enable_simple_responses
            __props__.__dict__["identity_source"] = identity_source
            __props__.__dict__["identity_validation_expression"] = identity_validation_expression
            __props__.__dict__["jwt_configuration"] = jwt_configuration
            __props__.__dict__["name"] = name
            __props__.__dict__["authorizer_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["api_id"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Authorizer, __self__).__init__(
            'aws-native:apigatewayv2:Authorizer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Authorizer':
        """
        Get an existing Authorizer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AuthorizerArgs.__new__(AuthorizerArgs)

        __props__.__dict__["api_id"] = None
        __props__.__dict__["authorizer_credentials_arn"] = None
        __props__.__dict__["authorizer_id"] = None
        __props__.__dict__["authorizer_payload_format_version"] = None
        __props__.__dict__["authorizer_result_ttl_in_seconds"] = None
        __props__.__dict__["authorizer_type"] = None
        __props__.__dict__["authorizer_uri"] = None
        __props__.__dict__["enable_simple_responses"] = None
        __props__.__dict__["identity_source"] = None
        __props__.__dict__["identity_validation_expression"] = None
        __props__.__dict__["jwt_configuration"] = None
        __props__.__dict__["name"] = None
        return Authorizer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "api_id")

    @property
    @pulumi.getter(name="authorizerCredentialsArn")
    def authorizer_credentials_arn(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "authorizer_credentials_arn")

    @property
    @pulumi.getter(name="authorizerId")
    def authorizer_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "authorizer_id")

    @property
    @pulumi.getter(name="authorizerPayloadFormatVersion")
    def authorizer_payload_format_version(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "authorizer_payload_format_version")

    @property
    @pulumi.getter(name="authorizerResultTtlInSeconds")
    def authorizer_result_ttl_in_seconds(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "authorizer_result_ttl_in_seconds")

    @property
    @pulumi.getter(name="authorizerType")
    def authorizer_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "authorizer_type")

    @property
    @pulumi.getter(name="authorizerUri")
    def authorizer_uri(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "authorizer_uri")

    @property
    @pulumi.getter(name="enableSimpleResponses")
    def enable_simple_responses(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "enable_simple_responses")

    @property
    @pulumi.getter(name="identitySource")
    def identity_source(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "identity_source")

    @property
    @pulumi.getter(name="identityValidationExpression")
    def identity_validation_expression(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "identity_validation_expression")

    @property
    @pulumi.getter(name="jwtConfiguration")
    def jwt_configuration(self) -> pulumi.Output[Optional['outputs.AuthorizerJwtConfiguration']]:
        return pulumi.get(self, "jwt_configuration")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

