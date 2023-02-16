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
from ._inputs import *

__all__ = ['TaskArgs', 'Task']

@pulumi.input_type
class TaskArgs:
    def __init__(__self__, *,
                 destination_location_arn: pulumi.Input[str],
                 source_location_arn: pulumi.Input[str],
                 cloud_watch_log_group_arn: Optional[pulumi.Input[str]] = None,
                 excludes: Optional[pulumi.Input[Sequence[pulumi.Input['TaskFilterRuleArgs']]]] = None,
                 includes: Optional[pulumi.Input[Sequence[pulumi.Input['TaskFilterRuleArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 options: Optional[pulumi.Input['TaskOptionsArgs']] = None,
                 schedule: Optional[pulumi.Input['TaskScheduleArgs']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['TaskTagArgs']]]] = None):
        """
        The set of arguments for constructing a Task resource.
        :param pulumi.Input[str] destination_location_arn: The ARN of an AWS storage resource's location.
        :param pulumi.Input[str] source_location_arn: The ARN of the source location for the task.
        :param pulumi.Input[str] cloud_watch_log_group_arn: The ARN of the Amazon CloudWatch log group that is used to monitor and log events in the task.
        :param pulumi.Input[str] name: The name of a task. This value is a text reference that is used to identify the task in the console.
        :param pulumi.Input[Sequence[pulumi.Input['TaskTagArgs']]] tags: An array of key-value pairs to apply to this resource.
        """
        pulumi.set(__self__, "destination_location_arn", destination_location_arn)
        pulumi.set(__self__, "source_location_arn", source_location_arn)
        if cloud_watch_log_group_arn is not None:
            pulumi.set(__self__, "cloud_watch_log_group_arn", cloud_watch_log_group_arn)
        if excludes is not None:
            pulumi.set(__self__, "excludes", excludes)
        if includes is not None:
            pulumi.set(__self__, "includes", includes)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if options is not None:
            pulumi.set(__self__, "options", options)
        if schedule is not None:
            pulumi.set(__self__, "schedule", schedule)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="destinationLocationArn")
    def destination_location_arn(self) -> pulumi.Input[str]:
        """
        The ARN of an AWS storage resource's location.
        """
        return pulumi.get(self, "destination_location_arn")

    @destination_location_arn.setter
    def destination_location_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination_location_arn", value)

    @property
    @pulumi.getter(name="sourceLocationArn")
    def source_location_arn(self) -> pulumi.Input[str]:
        """
        The ARN of the source location for the task.
        """
        return pulumi.get(self, "source_location_arn")

    @source_location_arn.setter
    def source_location_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "source_location_arn", value)

    @property
    @pulumi.getter(name="cloudWatchLogGroupArn")
    def cloud_watch_log_group_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the Amazon CloudWatch log group that is used to monitor and log events in the task.
        """
        return pulumi.get(self, "cloud_watch_log_group_arn")

    @cloud_watch_log_group_arn.setter
    def cloud_watch_log_group_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cloud_watch_log_group_arn", value)

    @property
    @pulumi.getter
    def excludes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TaskFilterRuleArgs']]]]:
        return pulumi.get(self, "excludes")

    @excludes.setter
    def excludes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TaskFilterRuleArgs']]]]):
        pulumi.set(self, "excludes", value)

    @property
    @pulumi.getter
    def includes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TaskFilterRuleArgs']]]]:
        return pulumi.get(self, "includes")

    @includes.setter
    def includes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TaskFilterRuleArgs']]]]):
        pulumi.set(self, "includes", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of a task. This value is a text reference that is used to identify the task in the console.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def options(self) -> Optional[pulumi.Input['TaskOptionsArgs']]:
        return pulumi.get(self, "options")

    @options.setter
    def options(self, value: Optional[pulumi.Input['TaskOptionsArgs']]):
        pulumi.set(self, "options", value)

    @property
    @pulumi.getter
    def schedule(self) -> Optional[pulumi.Input['TaskScheduleArgs']]:
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: Optional[pulumi.Input['TaskScheduleArgs']]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TaskTagArgs']]]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TaskTagArgs']]]]):
        pulumi.set(self, "tags", value)


class Task(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cloud_watch_log_group_arn: Optional[pulumi.Input[str]] = None,
                 destination_location_arn: Optional[pulumi.Input[str]] = None,
                 excludes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskFilterRuleArgs']]]]] = None,
                 includes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskFilterRuleArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 options: Optional[pulumi.Input[pulumi.InputType['TaskOptionsArgs']]] = None,
                 schedule: Optional[pulumi.Input[pulumi.InputType['TaskScheduleArgs']]] = None,
                 source_location_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskTagArgs']]]]] = None,
                 __props__=None):
        """
        Resource schema for AWS::DataSync::Task.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cloud_watch_log_group_arn: The ARN of the Amazon CloudWatch log group that is used to monitor and log events in the task.
        :param pulumi.Input[str] destination_location_arn: The ARN of an AWS storage resource's location.
        :param pulumi.Input[str] name: The name of a task. This value is a text reference that is used to identify the task in the console.
        :param pulumi.Input[str] source_location_arn: The ARN of the source location for the task.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskTagArgs']]]] tags: An array of key-value pairs to apply to this resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TaskArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource schema for AWS::DataSync::Task.

        :param str resource_name: The name of the resource.
        :param TaskArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TaskArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cloud_watch_log_group_arn: Optional[pulumi.Input[str]] = None,
                 destination_location_arn: Optional[pulumi.Input[str]] = None,
                 excludes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskFilterRuleArgs']]]]] = None,
                 includes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskFilterRuleArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 options: Optional[pulumi.Input[pulumi.InputType['TaskOptionsArgs']]] = None,
                 schedule: Optional[pulumi.Input[pulumi.InputType['TaskScheduleArgs']]] = None,
                 source_location_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskTagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TaskArgs.__new__(TaskArgs)

            __props__.__dict__["cloud_watch_log_group_arn"] = cloud_watch_log_group_arn
            if destination_location_arn is None and not opts.urn:
                raise TypeError("Missing required property 'destination_location_arn'")
            __props__.__dict__["destination_location_arn"] = destination_location_arn
            __props__.__dict__["excludes"] = excludes
            __props__.__dict__["includes"] = includes
            __props__.__dict__["name"] = name
            __props__.__dict__["options"] = options
            __props__.__dict__["schedule"] = schedule
            if source_location_arn is None and not opts.urn:
                raise TypeError("Missing required property 'source_location_arn'")
            __props__.__dict__["source_location_arn"] = source_location_arn
            __props__.__dict__["tags"] = tags
            __props__.__dict__["destination_network_interface_arns"] = None
            __props__.__dict__["source_network_interface_arns"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["task_arn"] = None
        super(Task, __self__).__init__(
            'aws-native:datasync:Task',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Task':
        """
        Get an existing Task resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = TaskArgs.__new__(TaskArgs)

        __props__.__dict__["cloud_watch_log_group_arn"] = None
        __props__.__dict__["destination_location_arn"] = None
        __props__.__dict__["destination_network_interface_arns"] = None
        __props__.__dict__["excludes"] = None
        __props__.__dict__["includes"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["options"] = None
        __props__.__dict__["schedule"] = None
        __props__.__dict__["source_location_arn"] = None
        __props__.__dict__["source_network_interface_arns"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["task_arn"] = None
        return Task(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cloudWatchLogGroupArn")
    def cloud_watch_log_group_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The ARN of the Amazon CloudWatch log group that is used to monitor and log events in the task.
        """
        return pulumi.get(self, "cloud_watch_log_group_arn")

    @property
    @pulumi.getter(name="destinationLocationArn")
    def destination_location_arn(self) -> pulumi.Output[str]:
        """
        The ARN of an AWS storage resource's location.
        """
        return pulumi.get(self, "destination_location_arn")

    @property
    @pulumi.getter(name="destinationNetworkInterfaceArns")
    def destination_network_interface_arns(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "destination_network_interface_arns")

    @property
    @pulumi.getter
    def excludes(self) -> pulumi.Output[Optional[Sequence['outputs.TaskFilterRule']]]:
        return pulumi.get(self, "excludes")

    @property
    @pulumi.getter
    def includes(self) -> pulumi.Output[Optional[Sequence['outputs.TaskFilterRule']]]:
        return pulumi.get(self, "includes")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of a task. This value is a text reference that is used to identify the task in the console.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def options(self) -> pulumi.Output[Optional['outputs.TaskOptions']]:
        return pulumi.get(self, "options")

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Output[Optional['outputs.TaskSchedule']]:
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter(name="sourceLocationArn")
    def source_location_arn(self) -> pulumi.Output[str]:
        """
        The ARN of the source location for the task.
        """
        return pulumi.get(self, "source_location_arn")

    @property
    @pulumi.getter(name="sourceNetworkInterfaceArns")
    def source_network_interface_arns(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "source_network_interface_arns")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output['TaskStatus']:
        """
        The status of the task that was described.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.TaskTag']]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="taskArn")
    def task_arn(self) -> pulumi.Output[str]:
        """
        The ARN of the task.
        """
        return pulumi.get(self, "task_arn")

