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

__all__ = [
    'GetClusterResult',
    'AwaitableGetClusterResult',
    'get_cluster',
    'get_cluster_output',
]

@pulumi.output_type
class GetClusterResult:
    def __init__(__self__, arn=None, broker_node_group_info=None, client_authentication=None, configuration_info=None, current_version=None, encryption_info=None, enhanced_monitoring=None, kafka_version=None, logging_info=None, number_of_broker_nodes=None, open_monitoring=None, storage_mode=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if broker_node_group_info and not isinstance(broker_node_group_info, dict):
            raise TypeError("Expected argument 'broker_node_group_info' to be a dict")
        pulumi.set(__self__, "broker_node_group_info", broker_node_group_info)
        if client_authentication and not isinstance(client_authentication, dict):
            raise TypeError("Expected argument 'client_authentication' to be a dict")
        pulumi.set(__self__, "client_authentication", client_authentication)
        if configuration_info and not isinstance(configuration_info, dict):
            raise TypeError("Expected argument 'configuration_info' to be a dict")
        pulumi.set(__self__, "configuration_info", configuration_info)
        if current_version and not isinstance(current_version, str):
            raise TypeError("Expected argument 'current_version' to be a str")
        pulumi.set(__self__, "current_version", current_version)
        if encryption_info and not isinstance(encryption_info, dict):
            raise TypeError("Expected argument 'encryption_info' to be a dict")
        pulumi.set(__self__, "encryption_info", encryption_info)
        if enhanced_monitoring and not isinstance(enhanced_monitoring, str):
            raise TypeError("Expected argument 'enhanced_monitoring' to be a str")
        pulumi.set(__self__, "enhanced_monitoring", enhanced_monitoring)
        if kafka_version and not isinstance(kafka_version, str):
            raise TypeError("Expected argument 'kafka_version' to be a str")
        pulumi.set(__self__, "kafka_version", kafka_version)
        if logging_info and not isinstance(logging_info, dict):
            raise TypeError("Expected argument 'logging_info' to be a dict")
        pulumi.set(__self__, "logging_info", logging_info)
        if number_of_broker_nodes and not isinstance(number_of_broker_nodes, int):
            raise TypeError("Expected argument 'number_of_broker_nodes' to be a int")
        pulumi.set(__self__, "number_of_broker_nodes", number_of_broker_nodes)
        if open_monitoring and not isinstance(open_monitoring, dict):
            raise TypeError("Expected argument 'open_monitoring' to be a dict")
        pulumi.set(__self__, "open_monitoring", open_monitoring)
        if storage_mode and not isinstance(storage_mode, str):
            raise TypeError("Expected argument 'storage_mode' to be a str")
        pulumi.set(__self__, "storage_mode", storage_mode)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="brokerNodeGroupInfo")
    def broker_node_group_info(self) -> Optional['outputs.ClusterBrokerNodeGroupInfo']:
        return pulumi.get(self, "broker_node_group_info")

    @property
    @pulumi.getter(name="clientAuthentication")
    def client_authentication(self) -> Optional['outputs.ClusterClientAuthentication']:
        return pulumi.get(self, "client_authentication")

    @property
    @pulumi.getter(name="configurationInfo")
    def configuration_info(self) -> Optional['outputs.ClusterConfigurationInfo']:
        return pulumi.get(self, "configuration_info")

    @property
    @pulumi.getter(name="currentVersion")
    def current_version(self) -> Optional[str]:
        """
        The current version of the MSK cluster
        """
        return pulumi.get(self, "current_version")

    @property
    @pulumi.getter(name="encryptionInfo")
    def encryption_info(self) -> Optional['outputs.ClusterEncryptionInfo']:
        return pulumi.get(self, "encryption_info")

    @property
    @pulumi.getter(name="enhancedMonitoring")
    def enhanced_monitoring(self) -> Optional['ClusterEnhancedMonitoring']:
        return pulumi.get(self, "enhanced_monitoring")

    @property
    @pulumi.getter(name="kafkaVersion")
    def kafka_version(self) -> Optional[str]:
        return pulumi.get(self, "kafka_version")

    @property
    @pulumi.getter(name="loggingInfo")
    def logging_info(self) -> Optional['outputs.ClusterLoggingInfo']:
        return pulumi.get(self, "logging_info")

    @property
    @pulumi.getter(name="numberOfBrokerNodes")
    def number_of_broker_nodes(self) -> Optional[int]:
        return pulumi.get(self, "number_of_broker_nodes")

    @property
    @pulumi.getter(name="openMonitoring")
    def open_monitoring(self) -> Optional['outputs.ClusterOpenMonitoring']:
        return pulumi.get(self, "open_monitoring")

    @property
    @pulumi.getter(name="storageMode")
    def storage_mode(self) -> Optional['ClusterStorageMode']:
        return pulumi.get(self, "storage_mode")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        A key-value pair to associate with a resource.
        """
        return pulumi.get(self, "tags")


class AwaitableGetClusterResult(GetClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClusterResult(
            arn=self.arn,
            broker_node_group_info=self.broker_node_group_info,
            client_authentication=self.client_authentication,
            configuration_info=self.configuration_info,
            current_version=self.current_version,
            encryption_info=self.encryption_info,
            enhanced_monitoring=self.enhanced_monitoring,
            kafka_version=self.kafka_version,
            logging_info=self.logging_info,
            number_of_broker_nodes=self.number_of_broker_nodes,
            open_monitoring=self.open_monitoring,
            storage_mode=self.storage_mode,
            tags=self.tags)


def get_cluster(arn: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClusterResult:
    """
    Resource Type definition for AWS::MSK::Cluster
    """
    __args__ = dict()
    __args__['arn'] = arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:msk:getCluster', __args__, opts=opts, typ=GetClusterResult).value

    return AwaitableGetClusterResult(
        arn=pulumi.get(__ret__, 'arn'),
        broker_node_group_info=pulumi.get(__ret__, 'broker_node_group_info'),
        client_authentication=pulumi.get(__ret__, 'client_authentication'),
        configuration_info=pulumi.get(__ret__, 'configuration_info'),
        current_version=pulumi.get(__ret__, 'current_version'),
        encryption_info=pulumi.get(__ret__, 'encryption_info'),
        enhanced_monitoring=pulumi.get(__ret__, 'enhanced_monitoring'),
        kafka_version=pulumi.get(__ret__, 'kafka_version'),
        logging_info=pulumi.get(__ret__, 'logging_info'),
        number_of_broker_nodes=pulumi.get(__ret__, 'number_of_broker_nodes'),
        open_monitoring=pulumi.get(__ret__, 'open_monitoring'),
        storage_mode=pulumi.get(__ret__, 'storage_mode'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_cluster)
def get_cluster_output(arn: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetClusterResult]:
    """
    Resource Type definition for AWS::MSK::Cluster
    """
    ...
