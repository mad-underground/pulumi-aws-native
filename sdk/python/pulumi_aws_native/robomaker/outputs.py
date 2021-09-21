# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = [
    'FleetTags',
    'RobotTags',
    'SimulationApplicationRenderingEngine',
    'SimulationApplicationRobotSoftwareSuite',
    'SimulationApplicationSimulationSoftwareSuite',
    'SimulationApplicationSourceConfig',
    'SimulationApplicationTags',
]

@pulumi.output_type
class FleetTags(dict):
    """
    A key-value pair to associate with a resource.
    """
    def __init__(__self__):
        """
        A key-value pair to associate with a resource.
        """
        pass


@pulumi.output_type
class RobotTags(dict):
    """
    A key-value pair to associate with a resource.
    """
    def __init__(__self__):
        """
        A key-value pair to associate with a resource.
        """
        pass


@pulumi.output_type
class SimulationApplicationRenderingEngine(dict):
    """
    Information about a rendering engine.
    """
    def __init__(__self__, *,
                 name: 'SimulationApplicationRenderingEngineName',
                 version: str):
        """
        Information about a rendering engine.
        :param 'SimulationApplicationRenderingEngineName' name: The name of the rendering engine.
        :param str version: The version of the rendering engine.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def name(self) -> 'SimulationApplicationRenderingEngineName':
        """
        The name of the rendering engine.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        The version of the rendering engine.
        """
        return pulumi.get(self, "version")


@pulumi.output_type
class SimulationApplicationRobotSoftwareSuite(dict):
    """
    Information about a robot software suite (ROS distribution).
    """
    def __init__(__self__, *,
                 name: 'SimulationApplicationRobotSoftwareSuiteName',
                 version: 'SimulationApplicationRobotSoftwareSuiteVersion'):
        """
        Information about a robot software suite (ROS distribution).
        :param 'SimulationApplicationRobotSoftwareSuiteName' name: The name of the robot software suite (ROS distribution).
        :param 'SimulationApplicationRobotSoftwareSuiteVersion' version: The version of the robot software suite (ROS distribution).
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def name(self) -> 'SimulationApplicationRobotSoftwareSuiteName':
        """
        The name of the robot software suite (ROS distribution).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def version(self) -> 'SimulationApplicationRobotSoftwareSuiteVersion':
        """
        The version of the robot software suite (ROS distribution).
        """
        return pulumi.get(self, "version")


@pulumi.output_type
class SimulationApplicationSimulationSoftwareSuite(dict):
    """
    Information about a simulation software suite.
    """
    def __init__(__self__, *,
                 name: 'SimulationApplicationSimulationSoftwareSuiteName',
                 version: 'SimulationApplicationSimulationSoftwareSuiteVersion'):
        """
        Information about a simulation software suite.
        :param 'SimulationApplicationSimulationSoftwareSuiteName' name: The name of the simulation software suite.
        :param 'SimulationApplicationSimulationSoftwareSuiteVersion' version: The version of the simulation software suite.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def name(self) -> 'SimulationApplicationSimulationSoftwareSuiteName':
        """
        The name of the simulation software suite.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def version(self) -> 'SimulationApplicationSimulationSoftwareSuiteVersion':
        """
        The version of the simulation software suite.
        """
        return pulumi.get(self, "version")


@pulumi.output_type
class SimulationApplicationSourceConfig(dict):
    """
    Information about a source configuration.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "s3Bucket":
            suggest = "s3_bucket"
        elif key == "s3Key":
            suggest = "s3_key"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SimulationApplicationSourceConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SimulationApplicationSourceConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SimulationApplicationSourceConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 architecture: 'SimulationApplicationSourceConfigArchitecture',
                 s3_bucket: str,
                 s3_key: str):
        """
        Information about a source configuration.
        :param 'SimulationApplicationSourceConfigArchitecture' architecture: The target processor architecture for the application.
        :param str s3_bucket: The Amazon S3 bucket name.
        :param str s3_key: The s3 object key.
        """
        pulumi.set(__self__, "architecture", architecture)
        pulumi.set(__self__, "s3_bucket", s3_bucket)
        pulumi.set(__self__, "s3_key", s3_key)

    @property
    @pulumi.getter
    def architecture(self) -> 'SimulationApplicationSourceConfigArchitecture':
        """
        The target processor architecture for the application.
        """
        return pulumi.get(self, "architecture")

    @property
    @pulumi.getter(name="s3Bucket")
    def s3_bucket(self) -> str:
        """
        The Amazon S3 bucket name.
        """
        return pulumi.get(self, "s3_bucket")

    @property
    @pulumi.getter(name="s3Key")
    def s3_key(self) -> str:
        """
        The s3 object key.
        """
        return pulumi.get(self, "s3_key")


@pulumi.output_type
class SimulationApplicationTags(dict):
    """
    A key-value pair to associate with a resource.
    """
    def __init__(__self__):
        """
        A key-value pair to associate with a resource.
        """
        pass

