# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .access_key import *
from .group import *
from .instance_profile import *
from .managed_policy import *
from .policy import *
from .role import *
from .service_linked_role import *
from .user import *
from .user_to_group_addition import *
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
            if typ == "cloudformation:IAM:AccessKey":
                return AccessKey(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "cloudformation:IAM:Group":
                return Group(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "cloudformation:IAM:InstanceProfile":
                return InstanceProfile(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "cloudformation:IAM:ManagedPolicy":
                return ManagedPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "cloudformation:IAM:Policy":
                return Policy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "cloudformation:IAM:Role":
                return Role(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "cloudformation:IAM:ServiceLinkedRole":
                return ServiceLinkedRole(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "cloudformation:IAM:User":
                return User(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "cloudformation:IAM:UserToGroupAddition":
                return UserToGroupAddition(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("cloudformation", "IAM", _module_instance)

_register_module()