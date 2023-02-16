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

__all__ = ['IntegrationArgs', 'Integration']

@pulumi.input_type
class IntegrationArgs:
    def __init__(__self__, *,
                 api_id: pulumi.Input[str],
                 integration_type: pulumi.Input[str],
                 connection_id: Optional[pulumi.Input[str]] = None,
                 connection_type: Optional[pulumi.Input[str]] = None,
                 content_handling_strategy: Optional[pulumi.Input[str]] = None,
                 credentials_arn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 integration_method: Optional[pulumi.Input[str]] = None,
                 integration_subtype: Optional[pulumi.Input[str]] = None,
                 integration_uri: Optional[pulumi.Input[str]] = None,
                 passthrough_behavior: Optional[pulumi.Input[str]] = None,
                 payload_format_version: Optional[pulumi.Input[str]] = None,
                 request_parameters: Optional[Any] = None,
                 request_templates: Optional[Any] = None,
                 response_parameters: Optional[Any] = None,
                 template_selection_expression: Optional[pulumi.Input[str]] = None,
                 timeout_in_millis: Optional[pulumi.Input[int]] = None,
                 tls_config: Optional[pulumi.Input['IntegrationTlsConfigArgs']] = None):
        """
        The set of arguments for constructing a Integration resource.
        """
        pulumi.set(__self__, "api_id", api_id)
        pulumi.set(__self__, "integration_type", integration_type)
        if connection_id is not None:
            pulumi.set(__self__, "connection_id", connection_id)
        if connection_type is not None:
            pulumi.set(__self__, "connection_type", connection_type)
        if content_handling_strategy is not None:
            pulumi.set(__self__, "content_handling_strategy", content_handling_strategy)
        if credentials_arn is not None:
            pulumi.set(__self__, "credentials_arn", credentials_arn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if integration_method is not None:
            pulumi.set(__self__, "integration_method", integration_method)
        if integration_subtype is not None:
            pulumi.set(__self__, "integration_subtype", integration_subtype)
        if integration_uri is not None:
            pulumi.set(__self__, "integration_uri", integration_uri)
        if passthrough_behavior is not None:
            pulumi.set(__self__, "passthrough_behavior", passthrough_behavior)
        if payload_format_version is not None:
            pulumi.set(__self__, "payload_format_version", payload_format_version)
        if request_parameters is not None:
            pulumi.set(__self__, "request_parameters", request_parameters)
        if request_templates is not None:
            pulumi.set(__self__, "request_templates", request_templates)
        if response_parameters is not None:
            pulumi.set(__self__, "response_parameters", response_parameters)
        if template_selection_expression is not None:
            pulumi.set(__self__, "template_selection_expression", template_selection_expression)
        if timeout_in_millis is not None:
            pulumi.set(__self__, "timeout_in_millis", timeout_in_millis)
        if tls_config is not None:
            pulumi.set(__self__, "tls_config", tls_config)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "api_id")

    @api_id.setter
    def api_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_id", value)

    @property
    @pulumi.getter(name="integrationType")
    def integration_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "integration_type")

    @integration_type.setter
    def integration_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "integration_type", value)

    @property
    @pulumi.getter(name="connectionId")
    def connection_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "connection_id")

    @connection_id.setter
    def connection_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_id", value)

    @property
    @pulumi.getter(name="connectionType")
    def connection_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "connection_type")

    @connection_type.setter
    def connection_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_type", value)

    @property
    @pulumi.getter(name="contentHandlingStrategy")
    def content_handling_strategy(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "content_handling_strategy")

    @content_handling_strategy.setter
    def content_handling_strategy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_handling_strategy", value)

    @property
    @pulumi.getter(name="credentialsArn")
    def credentials_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "credentials_arn")

    @credentials_arn.setter
    def credentials_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "credentials_arn", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="integrationMethod")
    def integration_method(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "integration_method")

    @integration_method.setter
    def integration_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "integration_method", value)

    @property
    @pulumi.getter(name="integrationSubtype")
    def integration_subtype(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "integration_subtype")

    @integration_subtype.setter
    def integration_subtype(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "integration_subtype", value)

    @property
    @pulumi.getter(name="integrationUri")
    def integration_uri(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "integration_uri")

    @integration_uri.setter
    def integration_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "integration_uri", value)

    @property
    @pulumi.getter(name="passthroughBehavior")
    def passthrough_behavior(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "passthrough_behavior")

    @passthrough_behavior.setter
    def passthrough_behavior(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "passthrough_behavior", value)

    @property
    @pulumi.getter(name="payloadFormatVersion")
    def payload_format_version(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "payload_format_version")

    @payload_format_version.setter
    def payload_format_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "payload_format_version", value)

    @property
    @pulumi.getter(name="requestParameters")
    def request_parameters(self) -> Optional[Any]:
        return pulumi.get(self, "request_parameters")

    @request_parameters.setter
    def request_parameters(self, value: Optional[Any]):
        pulumi.set(self, "request_parameters", value)

    @property
    @pulumi.getter(name="requestTemplates")
    def request_templates(self) -> Optional[Any]:
        return pulumi.get(self, "request_templates")

    @request_templates.setter
    def request_templates(self, value: Optional[Any]):
        pulumi.set(self, "request_templates", value)

    @property
    @pulumi.getter(name="responseParameters")
    def response_parameters(self) -> Optional[Any]:
        return pulumi.get(self, "response_parameters")

    @response_parameters.setter
    def response_parameters(self, value: Optional[Any]):
        pulumi.set(self, "response_parameters", value)

    @property
    @pulumi.getter(name="templateSelectionExpression")
    def template_selection_expression(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "template_selection_expression")

    @template_selection_expression.setter
    def template_selection_expression(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template_selection_expression", value)

    @property
    @pulumi.getter(name="timeoutInMillis")
    def timeout_in_millis(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "timeout_in_millis")

    @timeout_in_millis.setter
    def timeout_in_millis(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout_in_millis", value)

    @property
    @pulumi.getter(name="tlsConfig")
    def tls_config(self) -> Optional[pulumi.Input['IntegrationTlsConfigArgs']]:
        return pulumi.get(self, "tls_config")

    @tls_config.setter
    def tls_config(self, value: Optional[pulumi.Input['IntegrationTlsConfigArgs']]):
        pulumi.set(self, "tls_config", value)


warnings.warn("""Integration is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class Integration(pulumi.CustomResource):
    warnings.warn("""Integration is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 connection_id: Optional[pulumi.Input[str]] = None,
                 connection_type: Optional[pulumi.Input[str]] = None,
                 content_handling_strategy: Optional[pulumi.Input[str]] = None,
                 credentials_arn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 integration_method: Optional[pulumi.Input[str]] = None,
                 integration_subtype: Optional[pulumi.Input[str]] = None,
                 integration_type: Optional[pulumi.Input[str]] = None,
                 integration_uri: Optional[pulumi.Input[str]] = None,
                 passthrough_behavior: Optional[pulumi.Input[str]] = None,
                 payload_format_version: Optional[pulumi.Input[str]] = None,
                 request_parameters: Optional[Any] = None,
                 request_templates: Optional[Any] = None,
                 response_parameters: Optional[Any] = None,
                 template_selection_expression: Optional[pulumi.Input[str]] = None,
                 timeout_in_millis: Optional[pulumi.Input[int]] = None,
                 tls_config: Optional[pulumi.Input[pulumi.InputType['IntegrationTlsConfigArgs']]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::ApiGatewayV2::Integration

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IntegrationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::ApiGatewayV2::Integration

        :param str resource_name: The name of the resource.
        :param IntegrationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IntegrationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 connection_id: Optional[pulumi.Input[str]] = None,
                 connection_type: Optional[pulumi.Input[str]] = None,
                 content_handling_strategy: Optional[pulumi.Input[str]] = None,
                 credentials_arn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 integration_method: Optional[pulumi.Input[str]] = None,
                 integration_subtype: Optional[pulumi.Input[str]] = None,
                 integration_type: Optional[pulumi.Input[str]] = None,
                 integration_uri: Optional[pulumi.Input[str]] = None,
                 passthrough_behavior: Optional[pulumi.Input[str]] = None,
                 payload_format_version: Optional[pulumi.Input[str]] = None,
                 request_parameters: Optional[Any] = None,
                 request_templates: Optional[Any] = None,
                 response_parameters: Optional[Any] = None,
                 template_selection_expression: Optional[pulumi.Input[str]] = None,
                 timeout_in_millis: Optional[pulumi.Input[int]] = None,
                 tls_config: Optional[pulumi.Input[pulumi.InputType['IntegrationTlsConfigArgs']]] = None,
                 __props__=None):
        pulumi.log.warn("""Integration is deprecated: Integration is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IntegrationArgs.__new__(IntegrationArgs)

            if api_id is None and not opts.urn:
                raise TypeError("Missing required property 'api_id'")
            __props__.__dict__["api_id"] = api_id
            __props__.__dict__["connection_id"] = connection_id
            __props__.__dict__["connection_type"] = connection_type
            __props__.__dict__["content_handling_strategy"] = content_handling_strategy
            __props__.__dict__["credentials_arn"] = credentials_arn
            __props__.__dict__["description"] = description
            __props__.__dict__["integration_method"] = integration_method
            __props__.__dict__["integration_subtype"] = integration_subtype
            if integration_type is None and not opts.urn:
                raise TypeError("Missing required property 'integration_type'")
            __props__.__dict__["integration_type"] = integration_type
            __props__.__dict__["integration_uri"] = integration_uri
            __props__.__dict__["passthrough_behavior"] = passthrough_behavior
            __props__.__dict__["payload_format_version"] = payload_format_version
            __props__.__dict__["request_parameters"] = request_parameters
            __props__.__dict__["request_templates"] = request_templates
            __props__.__dict__["response_parameters"] = response_parameters
            __props__.__dict__["template_selection_expression"] = template_selection_expression
            __props__.__dict__["timeout_in_millis"] = timeout_in_millis
            __props__.__dict__["tls_config"] = tls_config
        super(Integration, __self__).__init__(
            'aws-native:apigatewayv2:Integration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Integration':
        """
        Get an existing Integration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = IntegrationArgs.__new__(IntegrationArgs)

        __props__.__dict__["api_id"] = None
        __props__.__dict__["connection_id"] = None
        __props__.__dict__["connection_type"] = None
        __props__.__dict__["content_handling_strategy"] = None
        __props__.__dict__["credentials_arn"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["integration_method"] = None
        __props__.__dict__["integration_subtype"] = None
        __props__.__dict__["integration_type"] = None
        __props__.__dict__["integration_uri"] = None
        __props__.__dict__["passthrough_behavior"] = None
        __props__.__dict__["payload_format_version"] = None
        __props__.__dict__["request_parameters"] = None
        __props__.__dict__["request_templates"] = None
        __props__.__dict__["response_parameters"] = None
        __props__.__dict__["template_selection_expression"] = None
        __props__.__dict__["timeout_in_millis"] = None
        __props__.__dict__["tls_config"] = None
        return Integration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "api_id")

    @property
    @pulumi.getter(name="connectionId")
    def connection_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "connection_id")

    @property
    @pulumi.getter(name="connectionType")
    def connection_type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "connection_type")

    @property
    @pulumi.getter(name="contentHandlingStrategy")
    def content_handling_strategy(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "content_handling_strategy")

    @property
    @pulumi.getter(name="credentialsArn")
    def credentials_arn(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "credentials_arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="integrationMethod")
    def integration_method(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "integration_method")

    @property
    @pulumi.getter(name="integrationSubtype")
    def integration_subtype(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "integration_subtype")

    @property
    @pulumi.getter(name="integrationType")
    def integration_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "integration_type")

    @property
    @pulumi.getter(name="integrationUri")
    def integration_uri(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "integration_uri")

    @property
    @pulumi.getter(name="passthroughBehavior")
    def passthrough_behavior(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "passthrough_behavior")

    @property
    @pulumi.getter(name="payloadFormatVersion")
    def payload_format_version(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "payload_format_version")

    @property
    @pulumi.getter(name="requestParameters")
    def request_parameters(self) -> pulumi.Output[Optional[Any]]:
        return pulumi.get(self, "request_parameters")

    @property
    @pulumi.getter(name="requestTemplates")
    def request_templates(self) -> pulumi.Output[Optional[Any]]:
        return pulumi.get(self, "request_templates")

    @property
    @pulumi.getter(name="responseParameters")
    def response_parameters(self) -> pulumi.Output[Optional[Any]]:
        return pulumi.get(self, "response_parameters")

    @property
    @pulumi.getter(name="templateSelectionExpression")
    def template_selection_expression(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "template_selection_expression")

    @property
    @pulumi.getter(name="timeoutInMillis")
    def timeout_in_millis(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "timeout_in_millis")

    @property
    @pulumi.getter(name="tlsConfig")
    def tls_config(self) -> pulumi.Output[Optional['outputs.IntegrationTlsConfig']]:
        return pulumi.get(self, "tls_config")

