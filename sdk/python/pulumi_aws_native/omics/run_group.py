# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RunGroupArgs', 'RunGroup']

@pulumi.input_type
class RunGroupArgs:
    def __init__(__self__, *,
                 max_cpus: Optional[pulumi.Input[float]] = None,
                 max_duration: Optional[pulumi.Input[float]] = None,
                 max_gpus: Optional[pulumi.Input[float]] = None,
                 max_runs: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a RunGroup resource.
        """
        if max_cpus is not None:
            pulumi.set(__self__, "max_cpus", max_cpus)
        if max_duration is not None:
            pulumi.set(__self__, "max_duration", max_duration)
        if max_gpus is not None:
            pulumi.set(__self__, "max_gpus", max_gpus)
        if max_runs is not None:
            pulumi.set(__self__, "max_runs", max_runs)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="maxCpus")
    def max_cpus(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "max_cpus")

    @max_cpus.setter
    def max_cpus(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "max_cpus", value)

    @property
    @pulumi.getter(name="maxDuration")
    def max_duration(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "max_duration")

    @max_duration.setter
    def max_duration(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "max_duration", value)

    @property
    @pulumi.getter(name="maxGpus")
    def max_gpus(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "max_gpus")

    @max_gpus.setter
    def max_gpus(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "max_gpus", value)

    @property
    @pulumi.getter(name="maxRuns")
    def max_runs(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "max_runs")

    @max_runs.setter
    def max_runs(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "max_runs", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class RunGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 max_cpus: Optional[pulumi.Input[float]] = None,
                 max_duration: Optional[pulumi.Input[float]] = None,
                 max_gpus: Optional[pulumi.Input[float]] = None,
                 max_runs: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Definition of AWS::Omics::RunGroup Resource Type

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[RunGroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of AWS::Omics::RunGroup Resource Type

        :param str resource_name: The name of the resource.
        :param RunGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RunGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 max_cpus: Optional[pulumi.Input[float]] = None,
                 max_duration: Optional[pulumi.Input[float]] = None,
                 max_gpus: Optional[pulumi.Input[float]] = None,
                 max_runs: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RunGroupArgs.__new__(RunGroupArgs)

            __props__.__dict__["max_cpus"] = max_cpus
            __props__.__dict__["max_duration"] = max_duration
            __props__.__dict__["max_gpus"] = max_gpus
            __props__.__dict__["max_runs"] = max_runs
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["creation_time"] = None
        super(RunGroup, __self__).__init__(
            'aws-native:omics:RunGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'RunGroup':
        """
        Get an existing RunGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = RunGroupArgs.__new__(RunGroupArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["creation_time"] = None
        __props__.__dict__["max_cpus"] = None
        __props__.__dict__["max_duration"] = None
        __props__.__dict__["max_gpus"] = None
        __props__.__dict__["max_runs"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["tags"] = None
        return RunGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="maxCpus")
    def max_cpus(self) -> pulumi.Output[Optional[float]]:
        return pulumi.get(self, "max_cpus")

    @property
    @pulumi.getter(name="maxDuration")
    def max_duration(self) -> pulumi.Output[Optional[float]]:
        return pulumi.get(self, "max_duration")

    @property
    @pulumi.getter(name="maxGpus")
    def max_gpus(self) -> pulumi.Output[Optional[float]]:
        return pulumi.get(self, "max_gpus")

    @property
    @pulumi.getter(name="maxRuns")
    def max_runs(self) -> pulumi.Output[Optional[float]]:
        return pulumi.get(self, "max_runs")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "tags")

