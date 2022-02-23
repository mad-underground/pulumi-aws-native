# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetClusterResult',
    'AwaitableGetClusterResult',
    'get_cluster',
    'get_cluster_output',
]

@pulumi.output_type
class GetClusterResult:
    def __init__(__self__, broker_node_group_info=None, client_authentication=None, configuration_info=None, encryption_info=None, enhanced_monitoring=None, id=None, kafka_version=None, logging_info=None, number_of_broker_nodes=None, open_monitoring=None):
        if broker_node_group_info and not isinstance(broker_node_group_info, dict):
            raise TypeError("Expected argument 'broker_node_group_info' to be a dict")
        pulumi.set(__self__, "broker_node_group_info", broker_node_group_info)
        if client_authentication and not isinstance(client_authentication, dict):
            raise TypeError("Expected argument 'client_authentication' to be a dict")
        pulumi.set(__self__, "client_authentication", client_authentication)
        if configuration_info and not isinstance(configuration_info, dict):
            raise TypeError("Expected argument 'configuration_info' to be a dict")
        pulumi.set(__self__, "configuration_info", configuration_info)
        if encryption_info and not isinstance(encryption_info, dict):
            raise TypeError("Expected argument 'encryption_info' to be a dict")
        pulumi.set(__self__, "encryption_info", encryption_info)
        if enhanced_monitoring and not isinstance(enhanced_monitoring, str):
            raise TypeError("Expected argument 'enhanced_monitoring' to be a str")
        pulumi.set(__self__, "enhanced_monitoring", enhanced_monitoring)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
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
    @pulumi.getter(name="encryptionInfo")
    def encryption_info(self) -> Optional['outputs.ClusterEncryptionInfo']:
        return pulumi.get(self, "encryption_info")

    @property
    @pulumi.getter(name="enhancedMonitoring")
    def enhanced_monitoring(self) -> Optional[str]:
        return pulumi.get(self, "enhanced_monitoring")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

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


class AwaitableGetClusterResult(GetClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClusterResult(
            broker_node_group_info=self.broker_node_group_info,
            client_authentication=self.client_authentication,
            configuration_info=self.configuration_info,
            encryption_info=self.encryption_info,
            enhanced_monitoring=self.enhanced_monitoring,
            id=self.id,
            kafka_version=self.kafka_version,
            logging_info=self.logging_info,
            number_of_broker_nodes=self.number_of_broker_nodes,
            open_monitoring=self.open_monitoring)


def get_cluster(id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClusterResult:
    """
    Resource Type definition for AWS::MSK::Cluster
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:msk:getCluster', __args__, opts=opts, typ=GetClusterResult).value

    return AwaitableGetClusterResult(
        broker_node_group_info=__ret__.broker_node_group_info,
        client_authentication=__ret__.client_authentication,
        configuration_info=__ret__.configuration_info,
        encryption_info=__ret__.encryption_info,
        enhanced_monitoring=__ret__.enhanced_monitoring,
        id=__ret__.id,
        kafka_version=__ret__.kafka_version,
        logging_info=__ret__.logging_info,
        number_of_broker_nodes=__ret__.number_of_broker_nodes,
        open_monitoring=__ret__.open_monitoring)


@_utilities.lift_output_func(get_cluster)
def get_cluster_output(id: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetClusterResult]:
    """
    Resource Type definition for AWS::MSK::Cluster
    """
    ...