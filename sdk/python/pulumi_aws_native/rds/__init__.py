# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .db_proxy import *
from .db_proxy_target_group import *
from .global_cluster import *
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
            if typ == "aws-native:RDS:DBProxy":
                return DBProxy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws-native:RDS:DBProxyTargetGroup":
                return DBProxyTargetGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws-native:RDS:GlobalCluster":
                return GlobalCluster(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("aws-native", "RDS", _module_instance)

_register_module()
