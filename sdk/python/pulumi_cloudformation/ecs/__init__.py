# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .capacity_provider import *
from .cluster import *
from .primary_task_set import *
from .service import *
from .task_definition import *
from .task_set import *
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
            if typ == "cloudformation:ECS:CapacityProvider":
                return CapacityProvider(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "cloudformation:ECS:Cluster":
                return Cluster(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "cloudformation:ECS:PrimaryTaskSet":
                return PrimaryTaskSet(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "cloudformation:ECS:Service":
                return Service(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "cloudformation:ECS:TaskDefinition":
                return TaskDefinition(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "cloudformation:ECS:TaskSet":
                return TaskSet(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("cloudformation", "ECS", _module_instance)

_register_module()