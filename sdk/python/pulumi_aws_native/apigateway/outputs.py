# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'ApiKeyStageKey',
    'DomainNameEndpointConfiguration',
    'DomainNameMutualTlsAuthentication',
]

@pulumi.output_type
class ApiKeyStageKey(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-apikey-stagekey.html
    """
    def __init__(__self__, *,
                 rest_api_id: Optional[str] = None,
                 stage_name: Optional[str] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-apikey-stagekey.html
        :param str rest_api_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-apikey-stagekey.html#cfn-apigateway-apikey-stagekey-restapiid
        :param str stage_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-apikey-stagekey.html#cfn-apigateway-apikey-stagekey-stagename
        """
        if rest_api_id is not None:
            pulumi.set(__self__, "rest_api_id", rest_api_id)
        if stage_name is not None:
            pulumi.set(__self__, "stage_name", stage_name)

    @property
    @pulumi.getter(name="RestApiId")
    def rest_api_id(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-apikey-stagekey.html#cfn-apigateway-apikey-stagekey-restapiid
        """
        return pulumi.get(self, "rest_api_id")

    @property
    @pulumi.getter(name="StageName")
    def stage_name(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-apikey-stagekey.html#cfn-apigateway-apikey-stagekey-stagename
        """
        return pulumi.get(self, "stage_name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DomainNameEndpointConfiguration(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-domainname-endpointconfiguration.html
    """
    def __init__(__self__, *,
                 types: Optional[Sequence[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-domainname-endpointconfiguration.html
        :param Sequence[str] types: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-domainname-endpointconfiguration.html#cfn-apigateway-domainname-endpointconfiguration-types
        """
        if types is not None:
            pulumi.set(__self__, "types", types)

    @property
    @pulumi.getter(name="Types")
    def types(self) -> Optional[Sequence[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-domainname-endpointconfiguration.html#cfn-apigateway-domainname-endpointconfiguration-types
        """
        return pulumi.get(self, "types")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DomainNameMutualTlsAuthentication(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-domainname-mutualtlsauthentication.html
    """
    def __init__(__self__, *,
                 truststore_uri: Optional[str] = None,
                 truststore_version: Optional[str] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-domainname-mutualtlsauthentication.html
        :param str truststore_uri: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-domainname-mutualtlsauthentication.html#cfn-apigateway-domainname-mutualtlsauthentication-truststoreuri
        :param str truststore_version: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-domainname-mutualtlsauthentication.html#cfn-apigateway-domainname-mutualtlsauthentication-truststoreversion
        """
        if truststore_uri is not None:
            pulumi.set(__self__, "truststore_uri", truststore_uri)
        if truststore_version is not None:
            pulumi.set(__self__, "truststore_version", truststore_version)

    @property
    @pulumi.getter(name="TruststoreUri")
    def truststore_uri(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-domainname-mutualtlsauthentication.html#cfn-apigateway-domainname-mutualtlsauthentication-truststoreuri
        """
        return pulumi.get(self, "truststore_uri")

    @property
    @pulumi.getter(name="TruststoreVersion")
    def truststore_version(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-domainname-mutualtlsauthentication.html#cfn-apigateway-domainname-mutualtlsauthentication-truststoreversion
        """
        return pulumi.get(self, "truststore_version")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


