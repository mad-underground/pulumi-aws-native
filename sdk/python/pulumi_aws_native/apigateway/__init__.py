# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .api_key import *
from .client_certificate import *
from .documentation_version import *
from .domain_name import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi
    from .. import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "aws-native:ApiGateway:ApiKey":
                return ApiKey(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws-native:ApiGateway:ClientCertificate":
                return ClientCertificate(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws-native:ApiGateway:DocumentationVersion":
                return DocumentationVersion(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws-native:ApiGateway:DomainName":
                return DomainName(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("aws-native", "ApiGateway", _module_instance)

_register_module()
