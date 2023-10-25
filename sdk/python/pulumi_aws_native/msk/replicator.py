# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['ReplicatorArgs', 'Replicator']

@pulumi.input_type
class ReplicatorArgs:
    def __init__(__self__, *,
                 kafka_clusters: pulumi.Input[Sequence[pulumi.Input['ReplicatorKafkaClusterArgs']]],
                 replication_info_list: pulumi.Input[Sequence[pulumi.Input['ReplicatorReplicationInfoArgs']]],
                 service_execution_role_arn: pulumi.Input[str],
                 current_version: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 replicator_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['ReplicatorTagArgs']]]] = None):
        """
        The set of arguments for constructing a Replicator resource.
        :param pulumi.Input[Sequence[pulumi.Input['ReplicatorKafkaClusterArgs']]] kafka_clusters: Specifies a list of Kafka clusters which are targets of the replicator.
        :param pulumi.Input[Sequence[pulumi.Input['ReplicatorReplicationInfoArgs']]] replication_info_list: A list of replication configurations, where each configuration targets a given source cluster to target cluster replication flow.
        :param pulumi.Input[str] service_execution_role_arn: The Amazon Resource Name (ARN) of the IAM role used by the replicator to access external resources.
        :param pulumi.Input[str] current_version: The current version of the MSK replicator.
        :param pulumi.Input[str] description: A summary description of the replicator.
        :param pulumi.Input[str] replicator_name: The name of the replicator.
        :param pulumi.Input[Sequence[pulumi.Input['ReplicatorTagArgs']]] tags: A collection of tags associated with a resource
        """
        ReplicatorArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            kafka_clusters=kafka_clusters,
            replication_info_list=replication_info_list,
            service_execution_role_arn=service_execution_role_arn,
            current_version=current_version,
            description=description,
            replicator_name=replicator_name,
            tags=tags,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             kafka_clusters: pulumi.Input[Sequence[pulumi.Input['ReplicatorKafkaClusterArgs']]],
             replication_info_list: pulumi.Input[Sequence[pulumi.Input['ReplicatorReplicationInfoArgs']]],
             service_execution_role_arn: pulumi.Input[str],
             current_version: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             replicator_name: Optional[pulumi.Input[str]] = None,
             tags: Optional[pulumi.Input[Sequence[pulumi.Input['ReplicatorTagArgs']]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("kafka_clusters", kafka_clusters)
        _setter("replication_info_list", replication_info_list)
        _setter("service_execution_role_arn", service_execution_role_arn)
        if current_version is not None:
            _setter("current_version", current_version)
        if description is not None:
            _setter("description", description)
        if replicator_name is not None:
            _setter("replicator_name", replicator_name)
        if tags is not None:
            _setter("tags", tags)

    @property
    @pulumi.getter(name="kafkaClusters")
    def kafka_clusters(self) -> pulumi.Input[Sequence[pulumi.Input['ReplicatorKafkaClusterArgs']]]:
        """
        Specifies a list of Kafka clusters which are targets of the replicator.
        """
        return pulumi.get(self, "kafka_clusters")

    @kafka_clusters.setter
    def kafka_clusters(self, value: pulumi.Input[Sequence[pulumi.Input['ReplicatorKafkaClusterArgs']]]):
        pulumi.set(self, "kafka_clusters", value)

    @property
    @pulumi.getter(name="replicationInfoList")
    def replication_info_list(self) -> pulumi.Input[Sequence[pulumi.Input['ReplicatorReplicationInfoArgs']]]:
        """
        A list of replication configurations, where each configuration targets a given source cluster to target cluster replication flow.
        """
        return pulumi.get(self, "replication_info_list")

    @replication_info_list.setter
    def replication_info_list(self, value: pulumi.Input[Sequence[pulumi.Input['ReplicatorReplicationInfoArgs']]]):
        pulumi.set(self, "replication_info_list", value)

    @property
    @pulumi.getter(name="serviceExecutionRoleArn")
    def service_execution_role_arn(self) -> pulumi.Input[str]:
        """
        The Amazon Resource Name (ARN) of the IAM role used by the replicator to access external resources.
        """
        return pulumi.get(self, "service_execution_role_arn")

    @service_execution_role_arn.setter
    def service_execution_role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_execution_role_arn", value)

    @property
    @pulumi.getter(name="currentVersion")
    def current_version(self) -> Optional[pulumi.Input[str]]:
        """
        The current version of the MSK replicator.
        """
        return pulumi.get(self, "current_version")

    @current_version.setter
    def current_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "current_version", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A summary description of the replicator.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="replicatorName")
    def replicator_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the replicator.
        """
        return pulumi.get(self, "replicator_name")

    @replicator_name.setter
    def replicator_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "replicator_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ReplicatorTagArgs']]]]:
        """
        A collection of tags associated with a resource
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ReplicatorTagArgs']]]]):
        pulumi.set(self, "tags", value)


class Replicator(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 current_version: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 kafka_clusters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReplicatorKafkaClusterArgs']]]]] = None,
                 replication_info_list: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReplicatorReplicationInfoArgs']]]]] = None,
                 replicator_name: Optional[pulumi.Input[str]] = None,
                 service_execution_role_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReplicatorTagArgs']]]]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::MSK::Replicator

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] current_version: The current version of the MSK replicator.
        :param pulumi.Input[str] description: A summary description of the replicator.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReplicatorKafkaClusterArgs']]]] kafka_clusters: Specifies a list of Kafka clusters which are targets of the replicator.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReplicatorReplicationInfoArgs']]]] replication_info_list: A list of replication configurations, where each configuration targets a given source cluster to target cluster replication flow.
        :param pulumi.Input[str] replicator_name: The name of the replicator.
        :param pulumi.Input[str] service_execution_role_arn: The Amazon Resource Name (ARN) of the IAM role used by the replicator to access external resources.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReplicatorTagArgs']]]] tags: A collection of tags associated with a resource
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReplicatorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::MSK::Replicator

        :param str resource_name: The name of the resource.
        :param ReplicatorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReplicatorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ReplicatorArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 current_version: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 kafka_clusters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReplicatorKafkaClusterArgs']]]]] = None,
                 replication_info_list: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReplicatorReplicationInfoArgs']]]]] = None,
                 replicator_name: Optional[pulumi.Input[str]] = None,
                 service_execution_role_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReplicatorTagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ReplicatorArgs.__new__(ReplicatorArgs)

            __props__.__dict__["current_version"] = current_version
            __props__.__dict__["description"] = description
            if kafka_clusters is None and not opts.urn:
                raise TypeError("Missing required property 'kafka_clusters'")
            __props__.__dict__["kafka_clusters"] = kafka_clusters
            if replication_info_list is None and not opts.urn:
                raise TypeError("Missing required property 'replication_info_list'")
            __props__.__dict__["replication_info_list"] = replication_info_list
            __props__.__dict__["replicator_name"] = replicator_name
            if service_execution_role_arn is None and not opts.urn:
                raise TypeError("Missing required property 'service_execution_role_arn'")
            __props__.__dict__["service_execution_role_arn"] = service_execution_role_arn
            __props__.__dict__["tags"] = tags
            __props__.__dict__["replicator_arn"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["description", "kafka_clusters[*]", "replicator_name", "service_execution_role_arn"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Replicator, __self__).__init__(
            'aws-native:msk:Replicator',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Replicator':
        """
        Get an existing Replicator resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ReplicatorArgs.__new__(ReplicatorArgs)

        __props__.__dict__["current_version"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["kafka_clusters"] = None
        __props__.__dict__["replication_info_list"] = None
        __props__.__dict__["replicator_arn"] = None
        __props__.__dict__["replicator_name"] = None
        __props__.__dict__["service_execution_role_arn"] = None
        __props__.__dict__["tags"] = None
        return Replicator(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="currentVersion")
    def current_version(self) -> pulumi.Output[Optional[str]]:
        """
        The current version of the MSK replicator.
        """
        return pulumi.get(self, "current_version")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A summary description of the replicator.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="kafkaClusters")
    def kafka_clusters(self) -> pulumi.Output[Sequence['outputs.ReplicatorKafkaCluster']]:
        """
        Specifies a list of Kafka clusters which are targets of the replicator.
        """
        return pulumi.get(self, "kafka_clusters")

    @property
    @pulumi.getter(name="replicationInfoList")
    def replication_info_list(self) -> pulumi.Output[Sequence['outputs.ReplicatorReplicationInfo']]:
        """
        A list of replication configurations, where each configuration targets a given source cluster to target cluster replication flow.
        """
        return pulumi.get(self, "replication_info_list")

    @property
    @pulumi.getter(name="replicatorArn")
    def replicator_arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name for the created replicator.
        """
        return pulumi.get(self, "replicator_arn")

    @property
    @pulumi.getter(name="replicatorName")
    def replicator_name(self) -> pulumi.Output[str]:
        """
        The name of the replicator.
        """
        return pulumi.get(self, "replicator_name")

    @property
    @pulumi.getter(name="serviceExecutionRoleArn")
    def service_execution_role_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the IAM role used by the replicator to access external resources.
        """
        return pulumi.get(self, "service_execution_role_arn")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.ReplicatorTag']]]:
        """
        A collection of tags associated with a resource
        """
        return pulumi.get(self, "tags")

