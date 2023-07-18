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

__all__ = [
    'GetAppResult',
    'AwaitableGetAppResult',
    'get_app',
    'get_app_output',
]

@pulumi.output_type
class GetAppResult:
    def __init__(__self__, app_source=None, attributes=None, data_sources=None, description=None, domains=None, enable_ssl=None, environment=None, id=None, name=None, ssl_configuration=None, type=None):
        if app_source and not isinstance(app_source, dict):
            raise TypeError("Expected argument 'app_source' to be a dict")
        pulumi.set(__self__, "app_source", app_source)
        if attributes and not isinstance(attributes, dict):
            raise TypeError("Expected argument 'attributes' to be a dict")
        pulumi.set(__self__, "attributes", attributes)
        if data_sources and not isinstance(data_sources, list):
            raise TypeError("Expected argument 'data_sources' to be a list")
        pulumi.set(__self__, "data_sources", data_sources)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if domains and not isinstance(domains, list):
            raise TypeError("Expected argument 'domains' to be a list")
        pulumi.set(__self__, "domains", domains)
        if enable_ssl and not isinstance(enable_ssl, bool):
            raise TypeError("Expected argument 'enable_ssl' to be a bool")
        pulumi.set(__self__, "enable_ssl", enable_ssl)
        if environment and not isinstance(environment, list):
            raise TypeError("Expected argument 'environment' to be a list")
        pulumi.set(__self__, "environment", environment)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if ssl_configuration and not isinstance(ssl_configuration, dict):
            raise TypeError("Expected argument 'ssl_configuration' to be a dict")
        pulumi.set(__self__, "ssl_configuration", ssl_configuration)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="appSource")
    def app_source(self) -> Optional['outputs.AppSource']:
        return pulumi.get(self, "app_source")

    @property
    @pulumi.getter
    def attributes(self) -> Optional[Any]:
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter(name="dataSources")
    def data_sources(self) -> Optional[Sequence['outputs.AppDataSource']]:
        return pulumi.get(self, "data_sources")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def domains(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "domains")

    @property
    @pulumi.getter(name="enableSsl")
    def enable_ssl(self) -> Optional[bool]:
        return pulumi.get(self, "enable_ssl")

    @property
    @pulumi.getter
    def environment(self) -> Optional[Sequence['outputs.AppEnvironmentVariable']]:
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="sslConfiguration")
    def ssl_configuration(self) -> Optional['outputs.AppSslConfiguration']:
        return pulumi.get(self, "ssl_configuration")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")


class AwaitableGetAppResult(GetAppResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppResult(
            app_source=self.app_source,
            attributes=self.attributes,
            data_sources=self.data_sources,
            description=self.description,
            domains=self.domains,
            enable_ssl=self.enable_ssl,
            environment=self.environment,
            id=self.id,
            name=self.name,
            ssl_configuration=self.ssl_configuration,
            type=self.type)


def get_app(id: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppResult:
    """
    Resource Type definition for AWS::OpsWorks::App
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:opsworks:getApp', __args__, opts=opts, typ=GetAppResult).value

    return AwaitableGetAppResult(
        app_source=pulumi.get(__ret__, 'app_source'),
        attributes=pulumi.get(__ret__, 'attributes'),
        data_sources=pulumi.get(__ret__, 'data_sources'),
        description=pulumi.get(__ret__, 'description'),
        domains=pulumi.get(__ret__, 'domains'),
        enable_ssl=pulumi.get(__ret__, 'enable_ssl'),
        environment=pulumi.get(__ret__, 'environment'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        ssl_configuration=pulumi.get(__ret__, 'ssl_configuration'),
        type=pulumi.get(__ret__, 'type'))


@_utilities.lift_output_func(get_app)
def get_app_output(id: Optional[pulumi.Input[str]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAppResult]:
    """
    Resource Type definition for AWS::OpsWorks::App
    """
    ...
