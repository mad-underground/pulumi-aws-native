# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .fleet import *
from .robot import *
from .robot_application import *
from .robot_application_version import *
from .simulation_application import *
from .simulation_application_version import *
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
            if typ == "cloudformation:RoboMaker:Fleet":
                return Fleet(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "cloudformation:RoboMaker:Robot":
                return Robot(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "cloudformation:RoboMaker:RobotApplication":
                return RobotApplication(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "cloudformation:RoboMaker:RobotApplicationVersion":
                return RobotApplicationVersion(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "cloudformation:RoboMaker:SimulationApplication":
                return SimulationApplication(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "cloudformation:RoboMaker:SimulationApplicationVersion":
                return SimulationApplicationVersion(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("cloudformation", "RoboMaker", _module_instance)

_register_module()
