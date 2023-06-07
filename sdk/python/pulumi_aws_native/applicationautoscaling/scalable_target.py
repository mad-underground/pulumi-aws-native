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
from ._inputs import *

__all__ = ['ScalableTargetArgs', 'ScalableTarget']

@pulumi.input_type
class ScalableTargetArgs:
    def __init__(__self__, *,
                 max_capacity: pulumi.Input[int],
                 min_capacity: pulumi.Input[int],
                 resource_id: pulumi.Input[str],
                 scalable_dimension: pulumi.Input[str],
                 service_namespace: pulumi.Input[str],
                 role_arn: Optional[pulumi.Input[str]] = None,
                 scheduled_actions: Optional[pulumi.Input[Sequence[pulumi.Input['ScalableTargetScheduledActionArgs']]]] = None,
                 suspended_state: Optional[pulumi.Input['ScalableTargetSuspendedStateArgs']] = None):
        """
        The set of arguments for constructing a ScalableTarget resource.
        :param pulumi.Input[int] max_capacity: The maximum value that you plan to scale in to. When a scaling policy is in effect, Application Auto Scaling can scale in (contract) as needed to the minimum capacity limit in response to changing demand
        :param pulumi.Input[int] min_capacity: The minimum value that you plan to scale in to. When a scaling policy is in effect, Application Auto Scaling can scale in (contract) as needed to the minimum capacity limit in response to changing demand
        :param pulumi.Input[str] resource_id: The identifier of the resource associated with the scalable target
        :param pulumi.Input[str] scalable_dimension: The scalable dimension associated with the scalable target. This string consists of the service namespace, resource type, and scaling property
        :param pulumi.Input[str] service_namespace: The namespace of the AWS service that provides the resource, or a custom-resource
        :param pulumi.Input[str] role_arn: Specify the Amazon Resource Name (ARN) of an Identity and Access Management (IAM) role that allows Application Auto Scaling to modify the scalable target on your behalf. 
        :param pulumi.Input[Sequence[pulumi.Input['ScalableTargetScheduledActionArgs']]] scheduled_actions: The scheduled actions for the scalable target. Duplicates aren't allowed.
        :param pulumi.Input['ScalableTargetSuspendedStateArgs'] suspended_state: An embedded object that contains attributes and attribute values that are used to suspend and resume automatic scaling. Setting the value of an attribute to true suspends the specified scaling activities. Setting it to false (default) resumes the specified scaling activities.
        """
        pulumi.set(__self__, "max_capacity", max_capacity)
        pulumi.set(__self__, "min_capacity", min_capacity)
        pulumi.set(__self__, "resource_id", resource_id)
        pulumi.set(__self__, "scalable_dimension", scalable_dimension)
        pulumi.set(__self__, "service_namespace", service_namespace)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)
        if scheduled_actions is not None:
            pulumi.set(__self__, "scheduled_actions", scheduled_actions)
        if suspended_state is not None:
            pulumi.set(__self__, "suspended_state", suspended_state)

    @property
    @pulumi.getter(name="maxCapacity")
    def max_capacity(self) -> pulumi.Input[int]:
        """
        The maximum value that you plan to scale in to. When a scaling policy is in effect, Application Auto Scaling can scale in (contract) as needed to the minimum capacity limit in response to changing demand
        """
        return pulumi.get(self, "max_capacity")

    @max_capacity.setter
    def max_capacity(self, value: pulumi.Input[int]):
        pulumi.set(self, "max_capacity", value)

    @property
    @pulumi.getter(name="minCapacity")
    def min_capacity(self) -> pulumi.Input[int]:
        """
        The minimum value that you plan to scale in to. When a scaling policy is in effect, Application Auto Scaling can scale in (contract) as needed to the minimum capacity limit in response to changing demand
        """
        return pulumi.get(self, "min_capacity")

    @min_capacity.setter
    def min_capacity(self, value: pulumi.Input[int]):
        pulumi.set(self, "min_capacity", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Input[str]:
        """
        The identifier of the resource associated with the scalable target
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="scalableDimension")
    def scalable_dimension(self) -> pulumi.Input[str]:
        """
        The scalable dimension associated with the scalable target. This string consists of the service namespace, resource type, and scaling property
        """
        return pulumi.get(self, "scalable_dimension")

    @scalable_dimension.setter
    def scalable_dimension(self, value: pulumi.Input[str]):
        pulumi.set(self, "scalable_dimension", value)

    @property
    @pulumi.getter(name="serviceNamespace")
    def service_namespace(self) -> pulumi.Input[str]:
        """
        The namespace of the AWS service that provides the resource, or a custom-resource
        """
        return pulumi.get(self, "service_namespace")

    @service_namespace.setter
    def service_namespace(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_namespace", value)

    @property
    @pulumi.getter(name="roleARN")
    def role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        Specify the Amazon Resource Name (ARN) of an Identity and Access Management (IAM) role that allows Application Auto Scaling to modify the scalable target on your behalf. 
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter(name="scheduledActions")
    def scheduled_actions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ScalableTargetScheduledActionArgs']]]]:
        """
        The scheduled actions for the scalable target. Duplicates aren't allowed.
        """
        return pulumi.get(self, "scheduled_actions")

    @scheduled_actions.setter
    def scheduled_actions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ScalableTargetScheduledActionArgs']]]]):
        pulumi.set(self, "scheduled_actions", value)

    @property
    @pulumi.getter(name="suspendedState")
    def suspended_state(self) -> Optional[pulumi.Input['ScalableTargetSuspendedStateArgs']]:
        """
        An embedded object that contains attributes and attribute values that are used to suspend and resume automatic scaling. Setting the value of an attribute to true suspends the specified scaling activities. Setting it to false (default) resumes the specified scaling activities.
        """
        return pulumi.get(self, "suspended_state")

    @suspended_state.setter
    def suspended_state(self, value: Optional[pulumi.Input['ScalableTargetSuspendedStateArgs']]):
        pulumi.set(self, "suspended_state", value)


class ScalableTarget(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 max_capacity: Optional[pulumi.Input[int]] = None,
                 min_capacity: Optional[pulumi.Input[int]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 scalable_dimension: Optional[pulumi.Input[str]] = None,
                 scheduled_actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ScalableTargetScheduledActionArgs']]]]] = None,
                 service_namespace: Optional[pulumi.Input[str]] = None,
                 suspended_state: Optional[pulumi.Input[pulumi.InputType['ScalableTargetSuspendedStateArgs']]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::ApplicationAutoScaling::ScalableTarget

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] max_capacity: The maximum value that you plan to scale in to. When a scaling policy is in effect, Application Auto Scaling can scale in (contract) as needed to the minimum capacity limit in response to changing demand
        :param pulumi.Input[int] min_capacity: The minimum value that you plan to scale in to. When a scaling policy is in effect, Application Auto Scaling can scale in (contract) as needed to the minimum capacity limit in response to changing demand
        :param pulumi.Input[str] resource_id: The identifier of the resource associated with the scalable target
        :param pulumi.Input[str] role_arn: Specify the Amazon Resource Name (ARN) of an Identity and Access Management (IAM) role that allows Application Auto Scaling to modify the scalable target on your behalf. 
        :param pulumi.Input[str] scalable_dimension: The scalable dimension associated with the scalable target. This string consists of the service namespace, resource type, and scaling property
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ScalableTargetScheduledActionArgs']]]] scheduled_actions: The scheduled actions for the scalable target. Duplicates aren't allowed.
        :param pulumi.Input[str] service_namespace: The namespace of the AWS service that provides the resource, or a custom-resource
        :param pulumi.Input[pulumi.InputType['ScalableTargetSuspendedStateArgs']] suspended_state: An embedded object that contains attributes and attribute values that are used to suspend and resume automatic scaling. Setting the value of an attribute to true suspends the specified scaling activities. Setting it to false (default) resumes the specified scaling activities.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ScalableTargetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::ApplicationAutoScaling::ScalableTarget

        :param str resource_name: The name of the resource.
        :param ScalableTargetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ScalableTargetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 max_capacity: Optional[pulumi.Input[int]] = None,
                 min_capacity: Optional[pulumi.Input[int]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 scalable_dimension: Optional[pulumi.Input[str]] = None,
                 scheduled_actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ScalableTargetScheduledActionArgs']]]]] = None,
                 service_namespace: Optional[pulumi.Input[str]] = None,
                 suspended_state: Optional[pulumi.Input[pulumi.InputType['ScalableTargetSuspendedStateArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ScalableTargetArgs.__new__(ScalableTargetArgs)

            if max_capacity is None and not opts.urn:
                raise TypeError("Missing required property 'max_capacity'")
            __props__.__dict__["max_capacity"] = max_capacity
            if min_capacity is None and not opts.urn:
                raise TypeError("Missing required property 'min_capacity'")
            __props__.__dict__["min_capacity"] = min_capacity
            if resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'resource_id'")
            __props__.__dict__["resource_id"] = resource_id
            __props__.__dict__["role_arn"] = role_arn
            if scalable_dimension is None and not opts.urn:
                raise TypeError("Missing required property 'scalable_dimension'")
            __props__.__dict__["scalable_dimension"] = scalable_dimension
            __props__.__dict__["scheduled_actions"] = scheduled_actions
            if service_namespace is None and not opts.urn:
                raise TypeError("Missing required property 'service_namespace'")
            __props__.__dict__["service_namespace"] = service_namespace
            __props__.__dict__["suspended_state"] = suspended_state
        super(ScalableTarget, __self__).__init__(
            'aws-native:applicationautoscaling:ScalableTarget',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ScalableTarget':
        """
        Get an existing ScalableTarget resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ScalableTargetArgs.__new__(ScalableTargetArgs)

        __props__.__dict__["max_capacity"] = None
        __props__.__dict__["min_capacity"] = None
        __props__.__dict__["resource_id"] = None
        __props__.__dict__["role_arn"] = None
        __props__.__dict__["scalable_dimension"] = None
        __props__.__dict__["scheduled_actions"] = None
        __props__.__dict__["service_namespace"] = None
        __props__.__dict__["suspended_state"] = None
        return ScalableTarget(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="maxCapacity")
    def max_capacity(self) -> pulumi.Output[int]:
        """
        The maximum value that you plan to scale in to. When a scaling policy is in effect, Application Auto Scaling can scale in (contract) as needed to the minimum capacity limit in response to changing demand
        """
        return pulumi.get(self, "max_capacity")

    @property
    @pulumi.getter(name="minCapacity")
    def min_capacity(self) -> pulumi.Output[int]:
        """
        The minimum value that you plan to scale in to. When a scaling policy is in effect, Application Auto Scaling can scale in (contract) as needed to the minimum capacity limit in response to changing demand
        """
        return pulumi.get(self, "min_capacity")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[str]:
        """
        The identifier of the resource associated with the scalable target
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="roleARN")
    def role_arn(self) -> pulumi.Output[Optional[str]]:
        """
        Specify the Amazon Resource Name (ARN) of an Identity and Access Management (IAM) role that allows Application Auto Scaling to modify the scalable target on your behalf. 
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="scalableDimension")
    def scalable_dimension(self) -> pulumi.Output[str]:
        """
        The scalable dimension associated with the scalable target. This string consists of the service namespace, resource type, and scaling property
        """
        return pulumi.get(self, "scalable_dimension")

    @property
    @pulumi.getter(name="scheduledActions")
    def scheduled_actions(self) -> pulumi.Output[Optional[Sequence['outputs.ScalableTargetScheduledAction']]]:
        """
        The scheduled actions for the scalable target. Duplicates aren't allowed.
        """
        return pulumi.get(self, "scheduled_actions")

    @property
    @pulumi.getter(name="serviceNamespace")
    def service_namespace(self) -> pulumi.Output[str]:
        """
        The namespace of the AWS service that provides the resource, or a custom-resource
        """
        return pulumi.get(self, "service_namespace")

    @property
    @pulumi.getter(name="suspendedState")
    def suspended_state(self) -> pulumi.Output[Optional['outputs.ScalableTargetSuspendedState']]:
        """
        An embedded object that contains attributes and attribute values that are used to suspend and resume automatic scaling. Setting the value of an attribute to true suspends the specified scaling activities. Setting it to false (default) resumes the specified scaling activities.
        """
        return pulumi.get(self, "suspended_state")

