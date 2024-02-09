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
from ._enums import *

__all__ = [
    'GetDomainConfigurationResult',
    'AwaitableGetDomainConfigurationResult',
    'get_domain_configuration',
    'get_domain_configuration_output',
]

@pulumi.output_type
class GetDomainConfigurationResult:
    def __init__(__self__, arn=None, authorizer_config=None, domain_configuration_status=None, domain_type=None, server_certificate_config=None, server_certificates=None, tags=None, tls_config=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if authorizer_config and not isinstance(authorizer_config, dict):
            raise TypeError("Expected argument 'authorizer_config' to be a dict")
        pulumi.set(__self__, "authorizer_config", authorizer_config)
        if domain_configuration_status and not isinstance(domain_configuration_status, str):
            raise TypeError("Expected argument 'domain_configuration_status' to be a str")
        pulumi.set(__self__, "domain_configuration_status", domain_configuration_status)
        if domain_type and not isinstance(domain_type, str):
            raise TypeError("Expected argument 'domain_type' to be a str")
        pulumi.set(__self__, "domain_type", domain_type)
        if server_certificate_config and not isinstance(server_certificate_config, dict):
            raise TypeError("Expected argument 'server_certificate_config' to be a dict")
        pulumi.set(__self__, "server_certificate_config", server_certificate_config)
        if server_certificates and not isinstance(server_certificates, list):
            raise TypeError("Expected argument 'server_certificates' to be a list")
        pulumi.set(__self__, "server_certificates", server_certificates)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if tls_config and not isinstance(tls_config, dict):
            raise TypeError("Expected argument 'tls_config' to be a dict")
        pulumi.set(__self__, "tls_config", tls_config)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="authorizerConfig")
    def authorizer_config(self) -> Optional['outputs.DomainConfigurationAuthorizerConfig']:
        return pulumi.get(self, "authorizer_config")

    @property
    @pulumi.getter(name="domainConfigurationStatus")
    def domain_configuration_status(self) -> Optional['DomainConfigurationStatus']:
        return pulumi.get(self, "domain_configuration_status")

    @property
    @pulumi.getter(name="domainType")
    def domain_type(self) -> Optional['DomainConfigurationDomainType']:
        return pulumi.get(self, "domain_type")

    @property
    @pulumi.getter(name="serverCertificateConfig")
    def server_certificate_config(self) -> Optional['outputs.DomainConfigurationServerCertificateConfig']:
        return pulumi.get(self, "server_certificate_config")

    @property
    @pulumi.getter(name="serverCertificates")
    def server_certificates(self) -> Optional[Sequence['outputs.DomainConfigurationServerCertificateSummary']]:
        return pulumi.get(self, "server_certificates")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.DomainConfigurationTag']]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tlsConfig")
    def tls_config(self) -> Optional['outputs.DomainConfigurationTlsConfig']:
        return pulumi.get(self, "tls_config")


class AwaitableGetDomainConfigurationResult(GetDomainConfigurationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDomainConfigurationResult(
            arn=self.arn,
            authorizer_config=self.authorizer_config,
            domain_configuration_status=self.domain_configuration_status,
            domain_type=self.domain_type,
            server_certificate_config=self.server_certificate_config,
            server_certificates=self.server_certificates,
            tags=self.tags,
            tls_config=self.tls_config)


def get_domain_configuration(domain_configuration_name: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDomainConfigurationResult:
    """
    Create and manage a Domain Configuration
    """
    __args__ = dict()
    __args__['domainConfigurationName'] = domain_configuration_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:iot:getDomainConfiguration', __args__, opts=opts, typ=GetDomainConfigurationResult).value

    return AwaitableGetDomainConfigurationResult(
        arn=pulumi.get(__ret__, 'arn'),
        authorizer_config=pulumi.get(__ret__, 'authorizer_config'),
        domain_configuration_status=pulumi.get(__ret__, 'domain_configuration_status'),
        domain_type=pulumi.get(__ret__, 'domain_type'),
        server_certificate_config=pulumi.get(__ret__, 'server_certificate_config'),
        server_certificates=pulumi.get(__ret__, 'server_certificates'),
        tags=pulumi.get(__ret__, 'tags'),
        tls_config=pulumi.get(__ret__, 'tls_config'))


@_utilities.lift_output_func(get_domain_configuration)
def get_domain_configuration_output(domain_configuration_name: Optional[pulumi.Input[str]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDomainConfigurationResult]:
    """
    Create and manage a Domain Configuration
    """
    ...
