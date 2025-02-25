# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'PrivateDnsNamespacePrivateDnsPropertiesMutableArgs',
    'PrivateDnsNamespacePropertiesArgs',
    'PrivateDnsNamespaceSoaArgs',
    'PublicDnsNamespacePropertiesArgs',
    'PublicDnsNamespacePublicDnsPropertiesMutableArgs',
    'PublicDnsNamespaceSoaArgs',
    'ServiceDnsConfigArgs',
    'ServiceDnsRecordArgs',
    'ServiceHealthCheckConfigArgs',
    'ServiceHealthCheckCustomConfigArgs',
]

@pulumi.input_type
class PrivateDnsNamespacePrivateDnsPropertiesMutableArgs:
    def __init__(__self__, *,
                 soa: Optional[pulumi.Input['PrivateDnsNamespaceSoaArgs']] = None):
        if soa is not None:
            pulumi.set(__self__, "soa", soa)

    @property
    @pulumi.getter
    def soa(self) -> Optional[pulumi.Input['PrivateDnsNamespaceSoaArgs']]:
        return pulumi.get(self, "soa")

    @soa.setter
    def soa(self, value: Optional[pulumi.Input['PrivateDnsNamespaceSoaArgs']]):
        pulumi.set(self, "soa", value)


@pulumi.input_type
class PrivateDnsNamespacePropertiesArgs:
    def __init__(__self__, *,
                 dns_properties: Optional[pulumi.Input['PrivateDnsNamespacePrivateDnsPropertiesMutableArgs']] = None):
        if dns_properties is not None:
            pulumi.set(__self__, "dns_properties", dns_properties)

    @property
    @pulumi.getter(name="dnsProperties")
    def dns_properties(self) -> Optional[pulumi.Input['PrivateDnsNamespacePrivateDnsPropertiesMutableArgs']]:
        return pulumi.get(self, "dns_properties")

    @dns_properties.setter
    def dns_properties(self, value: Optional[pulumi.Input['PrivateDnsNamespacePrivateDnsPropertiesMutableArgs']]):
        pulumi.set(self, "dns_properties", value)


@pulumi.input_type
class PrivateDnsNamespaceSoaArgs:
    def __init__(__self__, *,
                 ttl: Optional[pulumi.Input[float]] = None):
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "ttl", value)


@pulumi.input_type
class PublicDnsNamespacePropertiesArgs:
    def __init__(__self__, *,
                 dns_properties: Optional[pulumi.Input['PublicDnsNamespacePublicDnsPropertiesMutableArgs']] = None):
        if dns_properties is not None:
            pulumi.set(__self__, "dns_properties", dns_properties)

    @property
    @pulumi.getter(name="dnsProperties")
    def dns_properties(self) -> Optional[pulumi.Input['PublicDnsNamespacePublicDnsPropertiesMutableArgs']]:
        return pulumi.get(self, "dns_properties")

    @dns_properties.setter
    def dns_properties(self, value: Optional[pulumi.Input['PublicDnsNamespacePublicDnsPropertiesMutableArgs']]):
        pulumi.set(self, "dns_properties", value)


@pulumi.input_type
class PublicDnsNamespacePublicDnsPropertiesMutableArgs:
    def __init__(__self__, *,
                 soa: Optional[pulumi.Input['PublicDnsNamespaceSoaArgs']] = None):
        if soa is not None:
            pulumi.set(__self__, "soa", soa)

    @property
    @pulumi.getter
    def soa(self) -> Optional[pulumi.Input['PublicDnsNamespaceSoaArgs']]:
        return pulumi.get(self, "soa")

    @soa.setter
    def soa(self, value: Optional[pulumi.Input['PublicDnsNamespaceSoaArgs']]):
        pulumi.set(self, "soa", value)


@pulumi.input_type
class PublicDnsNamespaceSoaArgs:
    def __init__(__self__, *,
                 ttl: Optional[pulumi.Input[float]] = None):
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "ttl", value)


@pulumi.input_type
class ServiceDnsConfigArgs:
    def __init__(__self__, *,
                 dns_records: pulumi.Input[Sequence[pulumi.Input['ServiceDnsRecordArgs']]],
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 routing_policy: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "dns_records", dns_records)
        if namespace_id is not None:
            pulumi.set(__self__, "namespace_id", namespace_id)
        if routing_policy is not None:
            pulumi.set(__self__, "routing_policy", routing_policy)

    @property
    @pulumi.getter(name="dnsRecords")
    def dns_records(self) -> pulumi.Input[Sequence[pulumi.Input['ServiceDnsRecordArgs']]]:
        return pulumi.get(self, "dns_records")

    @dns_records.setter
    def dns_records(self, value: pulumi.Input[Sequence[pulumi.Input['ServiceDnsRecordArgs']]]):
        pulumi.set(self, "dns_records", value)

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "namespace_id")

    @namespace_id.setter
    def namespace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_id", value)

    @property
    @pulumi.getter(name="routingPolicy")
    def routing_policy(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "routing_policy")

    @routing_policy.setter
    def routing_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "routing_policy", value)


@pulumi.input_type
class ServiceDnsRecordArgs:
    def __init__(__self__, *,
                 ttl: pulumi.Input[float],
                 type: pulumi.Input[str]):
        pulumi.set(__self__, "ttl", ttl)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Input[float]:
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: pulumi.Input[float]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class ServiceHealthCheckConfigArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 failure_threshold: Optional[pulumi.Input[float]] = None,
                 resource_path: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "type", type)
        if failure_threshold is not None:
            pulumi.set(__self__, "failure_threshold", failure_threshold)
        if resource_path is not None:
            pulumi.set(__self__, "resource_path", resource_path)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="failureThreshold")
    def failure_threshold(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "failure_threshold")

    @failure_threshold.setter
    def failure_threshold(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "failure_threshold", value)

    @property
    @pulumi.getter(name="resourcePath")
    def resource_path(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "resource_path")

    @resource_path.setter
    def resource_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_path", value)


@pulumi.input_type
class ServiceHealthCheckCustomConfigArgs:
    def __init__(__self__, *,
                 failure_threshold: Optional[pulumi.Input[float]] = None):
        if failure_threshold is not None:
            pulumi.set(__self__, "failure_threshold", failure_threshold)

    @property
    @pulumi.getter(name="failureThreshold")
    def failure_threshold(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "failure_threshold")

    @failure_threshold.setter
    def failure_threshold(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "failure_threshold", value)


