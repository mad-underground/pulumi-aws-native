# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = [
    'AccessLogSubscriptionTagArgs',
    'ListenerDefaultActionArgs',
    'ListenerFixedResponseArgs',
    'ListenerForwardArgs',
    'ListenerTagArgs',
    'ListenerWeightedTargetGroupArgs',
    'RuleActionArgs',
    'RuleFixedResponseArgs',
    'RuleForwardArgs',
    'RuleHeaderMatchTypeArgs',
    'RuleHeaderMatchArgs',
    'RuleHttpMatchArgs',
    'RuleMatchArgs',
    'RulePathMatchTypeArgs',
    'RulePathMatchArgs',
    'RuleTagArgs',
    'RuleWeightedTargetGroupArgs',
    'ServiceDnsEntryArgs',
    'ServiceNetworkServiceAssociationDnsEntryArgs',
    'ServiceNetworkServiceAssociationTagArgs',
    'ServiceNetworkTagArgs',
    'ServiceNetworkVpcAssociationTagArgs',
    'ServiceTagArgs',
    'TargetGroupConfigArgs',
    'TargetGroupHealthCheckConfigArgs',
    'TargetGroupMatcherArgs',
    'TargetGroupTagArgs',
    'TargetGroupTargetArgs',
]

@pulumi.input_type
class AccessLogSubscriptionTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ListenerDefaultActionArgs:
    def __init__(__self__, *,
                 fixed_response: Optional[pulumi.Input['ListenerFixedResponseArgs']] = None,
                 forward: Optional[pulumi.Input['ListenerForwardArgs']] = None):
        if fixed_response is not None:
            pulumi.set(__self__, "fixed_response", fixed_response)
        if forward is not None:
            pulumi.set(__self__, "forward", forward)

    @property
    @pulumi.getter(name="fixedResponse")
    def fixed_response(self) -> Optional[pulumi.Input['ListenerFixedResponseArgs']]:
        return pulumi.get(self, "fixed_response")

    @fixed_response.setter
    def fixed_response(self, value: Optional[pulumi.Input['ListenerFixedResponseArgs']]):
        pulumi.set(self, "fixed_response", value)

    @property
    @pulumi.getter
    def forward(self) -> Optional[pulumi.Input['ListenerForwardArgs']]:
        return pulumi.get(self, "forward")

    @forward.setter
    def forward(self, value: Optional[pulumi.Input['ListenerForwardArgs']]):
        pulumi.set(self, "forward", value)


@pulumi.input_type
class ListenerFixedResponseArgs:
    def __init__(__self__, *,
                 status_code: pulumi.Input[int]):
        pulumi.set(__self__, "status_code", status_code)

    @property
    @pulumi.getter(name="statusCode")
    def status_code(self) -> pulumi.Input[int]:
        return pulumi.get(self, "status_code")

    @status_code.setter
    def status_code(self, value: pulumi.Input[int]):
        pulumi.set(self, "status_code", value)


@pulumi.input_type
class ListenerForwardArgs:
    def __init__(__self__, *,
                 target_groups: pulumi.Input[Sequence[pulumi.Input['ListenerWeightedTargetGroupArgs']]]):
        pulumi.set(__self__, "target_groups", target_groups)

    @property
    @pulumi.getter(name="targetGroups")
    def target_groups(self) -> pulumi.Input[Sequence[pulumi.Input['ListenerWeightedTargetGroupArgs']]]:
        return pulumi.get(self, "target_groups")

    @target_groups.setter
    def target_groups(self, value: pulumi.Input[Sequence[pulumi.Input['ListenerWeightedTargetGroupArgs']]]):
        pulumi.set(self, "target_groups", value)


@pulumi.input_type
class ListenerTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ListenerWeightedTargetGroupArgs:
    def __init__(__self__, *,
                 target_group_identifier: pulumi.Input[str],
                 weight: Optional[pulumi.Input[int]] = None):
        pulumi.set(__self__, "target_group_identifier", target_group_identifier)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter(name="targetGroupIdentifier")
    def target_group_identifier(self) -> pulumi.Input[str]:
        return pulumi.get(self, "target_group_identifier")

    @target_group_identifier.setter
    def target_group_identifier(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_group_identifier", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "weight", value)


@pulumi.input_type
class RuleActionArgs:
    def __init__(__self__, *,
                 fixed_response: Optional[pulumi.Input['RuleFixedResponseArgs']] = None,
                 forward: Optional[pulumi.Input['RuleForwardArgs']] = None):
        if fixed_response is not None:
            pulumi.set(__self__, "fixed_response", fixed_response)
        if forward is not None:
            pulumi.set(__self__, "forward", forward)

    @property
    @pulumi.getter(name="fixedResponse")
    def fixed_response(self) -> Optional[pulumi.Input['RuleFixedResponseArgs']]:
        return pulumi.get(self, "fixed_response")

    @fixed_response.setter
    def fixed_response(self, value: Optional[pulumi.Input['RuleFixedResponseArgs']]):
        pulumi.set(self, "fixed_response", value)

    @property
    @pulumi.getter
    def forward(self) -> Optional[pulumi.Input['RuleForwardArgs']]:
        return pulumi.get(self, "forward")

    @forward.setter
    def forward(self, value: Optional[pulumi.Input['RuleForwardArgs']]):
        pulumi.set(self, "forward", value)


@pulumi.input_type
class RuleFixedResponseArgs:
    def __init__(__self__, *,
                 status_code: pulumi.Input[int]):
        pulumi.set(__self__, "status_code", status_code)

    @property
    @pulumi.getter(name="statusCode")
    def status_code(self) -> pulumi.Input[int]:
        return pulumi.get(self, "status_code")

    @status_code.setter
    def status_code(self, value: pulumi.Input[int]):
        pulumi.set(self, "status_code", value)


@pulumi.input_type
class RuleForwardArgs:
    def __init__(__self__, *,
                 target_groups: pulumi.Input[Sequence[pulumi.Input['RuleWeightedTargetGroupArgs']]]):
        pulumi.set(__self__, "target_groups", target_groups)

    @property
    @pulumi.getter(name="targetGroups")
    def target_groups(self) -> pulumi.Input[Sequence[pulumi.Input['RuleWeightedTargetGroupArgs']]]:
        return pulumi.get(self, "target_groups")

    @target_groups.setter
    def target_groups(self, value: pulumi.Input[Sequence[pulumi.Input['RuleWeightedTargetGroupArgs']]]):
        pulumi.set(self, "target_groups", value)


@pulumi.input_type
class RuleHeaderMatchTypeArgs:
    def __init__(__self__, *,
                 contains: Optional[pulumi.Input[str]] = None,
                 exact: Optional[pulumi.Input[str]] = None,
                 prefix: Optional[pulumi.Input[str]] = None):
        if contains is not None:
            pulumi.set(__self__, "contains", contains)
        if exact is not None:
            pulumi.set(__self__, "exact", exact)
        if prefix is not None:
            pulumi.set(__self__, "prefix", prefix)

    @property
    @pulumi.getter
    def contains(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "contains")

    @contains.setter
    def contains(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "contains", value)

    @property
    @pulumi.getter
    def exact(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "exact")

    @exact.setter
    def exact(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "exact", value)

    @property
    @pulumi.getter
    def prefix(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "prefix")

    @prefix.setter
    def prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix", value)


@pulumi.input_type
class RuleHeaderMatchArgs:
    def __init__(__self__, *,
                 match: pulumi.Input['RuleHeaderMatchTypeArgs'],
                 name: pulumi.Input[str],
                 case_sensitive: Optional[pulumi.Input[bool]] = None):
        pulumi.set(__self__, "match", match)
        pulumi.set(__self__, "name", name)
        if case_sensitive is not None:
            pulumi.set(__self__, "case_sensitive", case_sensitive)

    @property
    @pulumi.getter
    def match(self) -> pulumi.Input['RuleHeaderMatchTypeArgs']:
        return pulumi.get(self, "match")

    @match.setter
    def match(self, value: pulumi.Input['RuleHeaderMatchTypeArgs']):
        pulumi.set(self, "match", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="caseSensitive")
    def case_sensitive(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "case_sensitive")

    @case_sensitive.setter
    def case_sensitive(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "case_sensitive", value)


@pulumi.input_type
class RuleHttpMatchArgs:
    def __init__(__self__, *,
                 header_matches: Optional[pulumi.Input[Sequence[pulumi.Input['RuleHeaderMatchArgs']]]] = None,
                 method: Optional[pulumi.Input['RuleHttpMatchMethod']] = None,
                 path_match: Optional[pulumi.Input['RulePathMatchArgs']] = None):
        if header_matches is not None:
            pulumi.set(__self__, "header_matches", header_matches)
        if method is not None:
            pulumi.set(__self__, "method", method)
        if path_match is not None:
            pulumi.set(__self__, "path_match", path_match)

    @property
    @pulumi.getter(name="headerMatches")
    def header_matches(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RuleHeaderMatchArgs']]]]:
        return pulumi.get(self, "header_matches")

    @header_matches.setter
    def header_matches(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RuleHeaderMatchArgs']]]]):
        pulumi.set(self, "header_matches", value)

    @property
    @pulumi.getter
    def method(self) -> Optional[pulumi.Input['RuleHttpMatchMethod']]:
        return pulumi.get(self, "method")

    @method.setter
    def method(self, value: Optional[pulumi.Input['RuleHttpMatchMethod']]):
        pulumi.set(self, "method", value)

    @property
    @pulumi.getter(name="pathMatch")
    def path_match(self) -> Optional[pulumi.Input['RulePathMatchArgs']]:
        return pulumi.get(self, "path_match")

    @path_match.setter
    def path_match(self, value: Optional[pulumi.Input['RulePathMatchArgs']]):
        pulumi.set(self, "path_match", value)


@pulumi.input_type
class RuleMatchArgs:
    def __init__(__self__, *,
                 http_match: pulumi.Input['RuleHttpMatchArgs']):
        pulumi.set(__self__, "http_match", http_match)

    @property
    @pulumi.getter(name="httpMatch")
    def http_match(self) -> pulumi.Input['RuleHttpMatchArgs']:
        return pulumi.get(self, "http_match")

    @http_match.setter
    def http_match(self, value: pulumi.Input['RuleHttpMatchArgs']):
        pulumi.set(self, "http_match", value)


@pulumi.input_type
class RulePathMatchTypeArgs:
    def __init__(__self__, *,
                 exact: Optional[pulumi.Input[str]] = None,
                 prefix: Optional[pulumi.Input[str]] = None):
        if exact is not None:
            pulumi.set(__self__, "exact", exact)
        if prefix is not None:
            pulumi.set(__self__, "prefix", prefix)

    @property
    @pulumi.getter
    def exact(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "exact")

    @exact.setter
    def exact(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "exact", value)

    @property
    @pulumi.getter
    def prefix(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "prefix")

    @prefix.setter
    def prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix", value)


@pulumi.input_type
class RulePathMatchArgs:
    def __init__(__self__, *,
                 match: pulumi.Input['RulePathMatchTypeArgs'],
                 case_sensitive: Optional[pulumi.Input[bool]] = None):
        pulumi.set(__self__, "match", match)
        if case_sensitive is not None:
            pulumi.set(__self__, "case_sensitive", case_sensitive)

    @property
    @pulumi.getter
    def match(self) -> pulumi.Input['RulePathMatchTypeArgs']:
        return pulumi.get(self, "match")

    @match.setter
    def match(self, value: pulumi.Input['RulePathMatchTypeArgs']):
        pulumi.set(self, "match", value)

    @property
    @pulumi.getter(name="caseSensitive")
    def case_sensitive(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "case_sensitive")

    @case_sensitive.setter
    def case_sensitive(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "case_sensitive", value)


@pulumi.input_type
class RuleTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class RuleWeightedTargetGroupArgs:
    def __init__(__self__, *,
                 target_group_identifier: pulumi.Input[str],
                 weight: Optional[pulumi.Input[int]] = None):
        pulumi.set(__self__, "target_group_identifier", target_group_identifier)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter(name="targetGroupIdentifier")
    def target_group_identifier(self) -> pulumi.Input[str]:
        return pulumi.get(self, "target_group_identifier")

    @target_group_identifier.setter
    def target_group_identifier(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_group_identifier", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "weight", value)


@pulumi.input_type
class ServiceDnsEntryArgs:
    def __init__(__self__, *,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 hosted_zone_id: Optional[pulumi.Input[str]] = None):
        if domain_name is not None:
            pulumi.set(__self__, "domain_name", domain_name)
        if hosted_zone_id is not None:
            pulumi.set(__self__, "hosted_zone_id", hosted_zone_id)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_name", value)

    @property
    @pulumi.getter(name="hostedZoneId")
    def hosted_zone_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "hosted_zone_id")

    @hosted_zone_id.setter
    def hosted_zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hosted_zone_id", value)


@pulumi.input_type
class ServiceNetworkServiceAssociationDnsEntryArgs:
    def __init__(__self__, *,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 hosted_zone_id: Optional[pulumi.Input[str]] = None):
        if domain_name is not None:
            pulumi.set(__self__, "domain_name", domain_name)
        if hosted_zone_id is not None:
            pulumi.set(__self__, "hosted_zone_id", hosted_zone_id)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_name", value)

    @property
    @pulumi.getter(name="hostedZoneId")
    def hosted_zone_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "hosted_zone_id")

    @hosted_zone_id.setter
    def hosted_zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hosted_zone_id", value)


@pulumi.input_type
class ServiceNetworkServiceAssociationTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ServiceNetworkTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ServiceNetworkVpcAssociationTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ServiceTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class TargetGroupConfigArgs:
    def __init__(__self__, *,
                 port: pulumi.Input[int],
                 protocol: pulumi.Input['TargetGroupConfigProtocol'],
                 vpc_identifier: pulumi.Input[str],
                 health_check: Optional[pulumi.Input['TargetGroupHealthCheckConfigArgs']] = None,
                 ip_address_type: Optional[pulumi.Input['TargetGroupConfigIpAddressType']] = None,
                 protocol_version: Optional[pulumi.Input['TargetGroupConfigProtocolVersion']] = None):
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "vpc_identifier", vpc_identifier)
        if health_check is not None:
            pulumi.set(__self__, "health_check", health_check)
        if ip_address_type is not None:
            pulumi.set(__self__, "ip_address_type", ip_address_type)
        if protocol_version is not None:
            pulumi.set(__self__, "protocol_version", protocol_version)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[int]:
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[int]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input['TargetGroupConfigProtocol']:
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input['TargetGroupConfigProtocol']):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="vpcIdentifier")
    def vpc_identifier(self) -> pulumi.Input[str]:
        return pulumi.get(self, "vpc_identifier")

    @vpc_identifier.setter
    def vpc_identifier(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_identifier", value)

    @property
    @pulumi.getter(name="healthCheck")
    def health_check(self) -> Optional[pulumi.Input['TargetGroupHealthCheckConfigArgs']]:
        return pulumi.get(self, "health_check")

    @health_check.setter
    def health_check(self, value: Optional[pulumi.Input['TargetGroupHealthCheckConfigArgs']]):
        pulumi.set(self, "health_check", value)

    @property
    @pulumi.getter(name="ipAddressType")
    def ip_address_type(self) -> Optional[pulumi.Input['TargetGroupConfigIpAddressType']]:
        return pulumi.get(self, "ip_address_type")

    @ip_address_type.setter
    def ip_address_type(self, value: Optional[pulumi.Input['TargetGroupConfigIpAddressType']]):
        pulumi.set(self, "ip_address_type", value)

    @property
    @pulumi.getter(name="protocolVersion")
    def protocol_version(self) -> Optional[pulumi.Input['TargetGroupConfigProtocolVersion']]:
        return pulumi.get(self, "protocol_version")

    @protocol_version.setter
    def protocol_version(self, value: Optional[pulumi.Input['TargetGroupConfigProtocolVersion']]):
        pulumi.set(self, "protocol_version", value)


@pulumi.input_type
class TargetGroupHealthCheckConfigArgs:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 health_check_interval_seconds: Optional[pulumi.Input[int]] = None,
                 health_check_timeout_seconds: Optional[pulumi.Input[int]] = None,
                 healthy_threshold_count: Optional[pulumi.Input[int]] = None,
                 matcher: Optional[pulumi.Input['TargetGroupMatcherArgs']] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input['TargetGroupHealthCheckConfigProtocol']] = None,
                 protocol_version: Optional[pulumi.Input['TargetGroupHealthCheckConfigProtocolVersion']] = None,
                 unhealthy_threshold_count: Optional[pulumi.Input[int]] = None):
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if health_check_interval_seconds is not None:
            pulumi.set(__self__, "health_check_interval_seconds", health_check_interval_seconds)
        if health_check_timeout_seconds is not None:
            pulumi.set(__self__, "health_check_timeout_seconds", health_check_timeout_seconds)
        if healthy_threshold_count is not None:
            pulumi.set(__self__, "healthy_threshold_count", healthy_threshold_count)
        if matcher is not None:
            pulumi.set(__self__, "matcher", matcher)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if protocol_version is not None:
            pulumi.set(__self__, "protocol_version", protocol_version)
        if unhealthy_threshold_count is not None:
            pulumi.set(__self__, "unhealthy_threshold_count", unhealthy_threshold_count)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="healthCheckIntervalSeconds")
    def health_check_interval_seconds(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "health_check_interval_seconds")

    @health_check_interval_seconds.setter
    def health_check_interval_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "health_check_interval_seconds", value)

    @property
    @pulumi.getter(name="healthCheckTimeoutSeconds")
    def health_check_timeout_seconds(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "health_check_timeout_seconds")

    @health_check_timeout_seconds.setter
    def health_check_timeout_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "health_check_timeout_seconds", value)

    @property
    @pulumi.getter(name="healthyThresholdCount")
    def healthy_threshold_count(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "healthy_threshold_count")

    @healthy_threshold_count.setter
    def healthy_threshold_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "healthy_threshold_count", value)

    @property
    @pulumi.getter
    def matcher(self) -> Optional[pulumi.Input['TargetGroupMatcherArgs']]:
        return pulumi.get(self, "matcher")

    @matcher.setter
    def matcher(self, value: Optional[pulumi.Input['TargetGroupMatcherArgs']]):
        pulumi.set(self, "matcher", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input['TargetGroupHealthCheckConfigProtocol']]:
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input['TargetGroupHealthCheckConfigProtocol']]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="protocolVersion")
    def protocol_version(self) -> Optional[pulumi.Input['TargetGroupHealthCheckConfigProtocolVersion']]:
        return pulumi.get(self, "protocol_version")

    @protocol_version.setter
    def protocol_version(self, value: Optional[pulumi.Input['TargetGroupHealthCheckConfigProtocolVersion']]):
        pulumi.set(self, "protocol_version", value)

    @property
    @pulumi.getter(name="unhealthyThresholdCount")
    def unhealthy_threshold_count(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "unhealthy_threshold_count")

    @unhealthy_threshold_count.setter
    def unhealthy_threshold_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "unhealthy_threshold_count", value)


@pulumi.input_type
class TargetGroupMatcherArgs:
    def __init__(__self__, *,
                 http_code: pulumi.Input[str]):
        pulumi.set(__self__, "http_code", http_code)

    @property
    @pulumi.getter(name="httpCode")
    def http_code(self) -> pulumi.Input[str]:
        return pulumi.get(self, "http_code")

    @http_code.setter
    def http_code(self, value: pulumi.Input[str]):
        pulumi.set(self, "http_code", value)


@pulumi.input_type
class TargetGroupTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class TargetGroupTargetArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str],
                 port: Optional[pulumi.Input[int]] = None):
        pulumi.set(__self__, "id", id)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

