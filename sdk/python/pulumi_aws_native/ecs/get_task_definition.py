# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from .. import outputs as _root_outputs

__all__ = [
    'GetTaskDefinitionResult',
    'AwaitableGetTaskDefinitionResult',
    'get_task_definition',
    'get_task_definition_output',
]

@pulumi.output_type
class GetTaskDefinitionResult:
    def __init__(__self__, tags=None, task_definition_arn=None):
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if task_definition_arn and not isinstance(task_definition_arn, str):
            raise TypeError("Expected argument 'task_definition_arn' to be a str")
        pulumi.set(__self__, "task_definition_arn", task_definition_arn)

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        The metadata that you apply to the task definition to help you categorize and organize them. Each tag consists of a key and an optional value. You define both of them.
         The following basic restrictions apply to tags:
          +  Maximum number of tags per resource - 50
          +  For each resource, each tag key must be unique, and each tag key can have only one value.
          +  Maximum key length - 128 Unicode characters in UTF-8
          +  Maximum value length - 256 Unicode characters in UTF-8
          +  If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.
          +  Tag keys and values are case-sensitive.
          +  Do not use ``aws:``, ``AWS:``, or any upper or lowercase combination of such as a prefix for either keys or values as it is reserved for AWS use. You cannot edit or delete tag keys or values
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="taskDefinitionArn")
    def task_definition_arn(self) -> Optional[str]:
        return pulumi.get(self, "task_definition_arn")


class AwaitableGetTaskDefinitionResult(GetTaskDefinitionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTaskDefinitionResult(
            tags=self.tags,
            task_definition_arn=self.task_definition_arn)


def get_task_definition(task_definition_arn: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTaskDefinitionResult:
    """
    Registers a new task definition from the supplied ``family`` and ``containerDefinitions``. Optionally, you can add data volumes to your containers with the ``volumes`` parameter. For more information about task definition parameters and defaults, see [Amazon ECS Task Definitions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html) in the *Amazon Elastic Container Service Developer Guide*.
     You can specify a role for your task with the ``taskRoleArn`` parameter. When you specify a role for a task, its containers can then use the latest versions of the CLI or SDKs to make API requests to the AWS services that are specified in the policy that's associated with the role. For more information, see [IAM Roles for Tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html) in the *Amazon Elastic Container Service Developer Guide*.
     You can specify a Docker networking mode for the containers in your task definition with the ``networkMod
    """
    __args__ = dict()
    __args__['taskDefinitionArn'] = task_definition_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:ecs:getTaskDefinition', __args__, opts=opts, typ=GetTaskDefinitionResult).value

    return AwaitableGetTaskDefinitionResult(
        tags=pulumi.get(__ret__, 'tags'),
        task_definition_arn=pulumi.get(__ret__, 'task_definition_arn'))


@_utilities.lift_output_func(get_task_definition)
def get_task_definition_output(task_definition_arn: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTaskDefinitionResult]:
    """
    Registers a new task definition from the supplied ``family`` and ``containerDefinitions``. Optionally, you can add data volumes to your containers with the ``volumes`` parameter. For more information about task definition parameters and defaults, see [Amazon ECS Task Definitions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html) in the *Amazon Elastic Container Service Developer Guide*.
     You can specify a role for your task with the ``taskRoleArn`` parameter. When you specify a role for a task, its containers can then use the latest versions of the CLI or SDKs to make API requests to the AWS services that are specified in the policy that's associated with the role. For more information, see [IAM Roles for Tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html) in the *Amazon Elastic Container Service Developer Guide*.
     You can specify a Docker networking mode for the containers in your task definition with the ``networkMod
    """
    ...
