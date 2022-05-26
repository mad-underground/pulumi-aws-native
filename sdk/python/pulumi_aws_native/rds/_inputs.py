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
    'DBClusterParameterGroupTagArgs',
    'DBClusterRoleArgs',
    'DBClusterScalingConfigurationArgs',
    'DBClusterTagArgs',
    'DBInstanceProcessorFeatureArgs',
    'DBInstanceRoleArgs',
    'DBInstanceTagArgs',
    'DBParameterGroupTagArgs',
    'DBProxyAuthFormatArgs',
    'DBProxyEndpointTagFormatArgs',
    'DBProxyTagFormatArgs',
    'DBProxyTargetGroupConnectionPoolConfigurationInfoFormatArgs',
    'DBSecurityGroupIngressArgs',
    'DBSecurityGroupTagArgs',
    'DBSubnetGroupTagArgs',
    'EventSubscriptionTagArgs',
    'OptionGroupOptionConfigurationArgs',
    'OptionGroupOptionSettingArgs',
    'OptionGroupTagArgs',
]

@pulumi.input_type
class DBClusterParameterGroupTagArgs:
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
class DBClusterRoleArgs:
    def __init__(__self__, *,
                 role_arn: pulumi.Input[str],
                 feature_name: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "role_arn", role_arn)
        if feature_name is not None:
            pulumi.set(__self__, "feature_name", feature_name)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Input[str]:
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter(name="featureName")
    def feature_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "feature_name")

    @feature_name.setter
    def feature_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "feature_name", value)


@pulumi.input_type
class DBClusterScalingConfigurationArgs:
    def __init__(__self__, *,
                 auto_pause: Optional[pulumi.Input[bool]] = None,
                 max_capacity: Optional[pulumi.Input[int]] = None,
                 min_capacity: Optional[pulumi.Input[int]] = None,
                 seconds_until_auto_pause: Optional[pulumi.Input[int]] = None):
        if auto_pause is not None:
            pulumi.set(__self__, "auto_pause", auto_pause)
        if max_capacity is not None:
            pulumi.set(__self__, "max_capacity", max_capacity)
        if min_capacity is not None:
            pulumi.set(__self__, "min_capacity", min_capacity)
        if seconds_until_auto_pause is not None:
            pulumi.set(__self__, "seconds_until_auto_pause", seconds_until_auto_pause)

    @property
    @pulumi.getter(name="autoPause")
    def auto_pause(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "auto_pause")

    @auto_pause.setter
    def auto_pause(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_pause", value)

    @property
    @pulumi.getter(name="maxCapacity")
    def max_capacity(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "max_capacity")

    @max_capacity.setter
    def max_capacity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_capacity", value)

    @property
    @pulumi.getter(name="minCapacity")
    def min_capacity(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "min_capacity")

    @min_capacity.setter
    def min_capacity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_capacity", value)

    @property
    @pulumi.getter(name="secondsUntilAutoPause")
    def seconds_until_auto_pause(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "seconds_until_auto_pause")

    @seconds_until_auto_pause.setter
    def seconds_until_auto_pause(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "seconds_until_auto_pause", value)


@pulumi.input_type
class DBClusterTagArgs:
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
class DBInstanceProcessorFeatureArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        if name is not None:
            pulumi.set(__self__, "name", name)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class DBInstanceRoleArgs:
    def __init__(__self__, *,
                 feature_name: pulumi.Input[str],
                 role_arn: pulumi.Input[str]):
        pulumi.set(__self__, "feature_name", feature_name)
        pulumi.set(__self__, "role_arn", role_arn)

    @property
    @pulumi.getter(name="featureName")
    def feature_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "feature_name")

    @feature_name.setter
    def feature_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "feature_name", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Input[str]:
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_arn", value)


@pulumi.input_type
class DBInstanceTagArgs:
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
class DBParameterGroupTagArgs:
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
class DBProxyAuthFormatArgs:
    def __init__(__self__, *,
                 auth_scheme: Optional[pulumi.Input['DBProxyAuthFormatAuthScheme']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 i_am_auth: Optional[pulumi.Input['DBProxyAuthFormatIAMAuth']] = None,
                 secret_arn: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input['DBProxyAuthFormatAuthScheme'] auth_scheme: The type of authentication that the proxy uses for connections from the proxy to the underlying database. 
        :param pulumi.Input[str] description: A user-specified description about the authentication used by a proxy to log in as a specific database user. 
        :param pulumi.Input['DBProxyAuthFormatIAMAuth'] i_am_auth: Whether to require or disallow AWS Identity and Access Management (IAM) authentication for connections to the proxy. 
        :param pulumi.Input[str] secret_arn: The Amazon Resource Name (ARN) representing the secret that the proxy uses to authenticate to the RDS DB instance or Aurora DB cluster. These secrets are stored within Amazon Secrets Manager. 
        :param pulumi.Input[str] user_name: The name of the database user to which the proxy connects.
        """
        if auth_scheme is not None:
            pulumi.set(__self__, "auth_scheme", auth_scheme)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if i_am_auth is not None:
            pulumi.set(__self__, "i_am_auth", i_am_auth)
        if secret_arn is not None:
            pulumi.set(__self__, "secret_arn", secret_arn)
        if user_name is not None:
            pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="authScheme")
    def auth_scheme(self) -> Optional[pulumi.Input['DBProxyAuthFormatAuthScheme']]:
        """
        The type of authentication that the proxy uses for connections from the proxy to the underlying database. 
        """
        return pulumi.get(self, "auth_scheme")

    @auth_scheme.setter
    def auth_scheme(self, value: Optional[pulumi.Input['DBProxyAuthFormatAuthScheme']]):
        pulumi.set(self, "auth_scheme", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A user-specified description about the authentication used by a proxy to log in as a specific database user. 
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="iAMAuth")
    def i_am_auth(self) -> Optional[pulumi.Input['DBProxyAuthFormatIAMAuth']]:
        """
        Whether to require or disallow AWS Identity and Access Management (IAM) authentication for connections to the proxy. 
        """
        return pulumi.get(self, "i_am_auth")

    @i_am_auth.setter
    def i_am_auth(self, value: Optional[pulumi.Input['DBProxyAuthFormatIAMAuth']]):
        pulumi.set(self, "i_am_auth", value)

    @property
    @pulumi.getter(name="secretArn")
    def secret_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) representing the secret that the proxy uses to authenticate to the RDS DB instance or Aurora DB cluster. These secrets are stored within Amazon Secrets Manager. 
        """
        return pulumi.get(self, "secret_arn")

    @secret_arn.setter
    def secret_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_arn", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the database user to which the proxy connects.
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_name", value)


@pulumi.input_type
class DBProxyEndpointTagFormatArgs:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        if key is not None:
            pulumi.set(__self__, "key", key)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class DBProxyTagFormatArgs:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        if key is not None:
            pulumi.set(__self__, "key", key)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class DBProxyTargetGroupConnectionPoolConfigurationInfoFormatArgs:
    def __init__(__self__, *,
                 connection_borrow_timeout: Optional[pulumi.Input[int]] = None,
                 init_query: Optional[pulumi.Input[str]] = None,
                 max_connections_percent: Optional[pulumi.Input[int]] = None,
                 max_idle_connections_percent: Optional[pulumi.Input[int]] = None,
                 session_pinning_filters: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[int] connection_borrow_timeout: The number of seconds for a proxy to wait for a connection to become available in the connection pool.
        :param pulumi.Input[str] init_query: One or more SQL statements for the proxy to run when opening each new database connection.
        :param pulumi.Input[int] max_connections_percent: The maximum size of the connection pool for each target in a target group.
        :param pulumi.Input[int] max_idle_connections_percent: Controls how actively the proxy closes idle database connections in the connection pool.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] session_pinning_filters: Each item in the list represents a class of SQL operations that normally cause all later statements in a session using a proxy to be pinned to the same underlying database connection.
        """
        if connection_borrow_timeout is not None:
            pulumi.set(__self__, "connection_borrow_timeout", connection_borrow_timeout)
        if init_query is not None:
            pulumi.set(__self__, "init_query", init_query)
        if max_connections_percent is not None:
            pulumi.set(__self__, "max_connections_percent", max_connections_percent)
        if max_idle_connections_percent is not None:
            pulumi.set(__self__, "max_idle_connections_percent", max_idle_connections_percent)
        if session_pinning_filters is not None:
            pulumi.set(__self__, "session_pinning_filters", session_pinning_filters)

    @property
    @pulumi.getter(name="connectionBorrowTimeout")
    def connection_borrow_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        The number of seconds for a proxy to wait for a connection to become available in the connection pool.
        """
        return pulumi.get(self, "connection_borrow_timeout")

    @connection_borrow_timeout.setter
    def connection_borrow_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "connection_borrow_timeout", value)

    @property
    @pulumi.getter(name="initQuery")
    def init_query(self) -> Optional[pulumi.Input[str]]:
        """
        One or more SQL statements for the proxy to run when opening each new database connection.
        """
        return pulumi.get(self, "init_query")

    @init_query.setter
    def init_query(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "init_query", value)

    @property
    @pulumi.getter(name="maxConnectionsPercent")
    def max_connections_percent(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum size of the connection pool for each target in a target group.
        """
        return pulumi.get(self, "max_connections_percent")

    @max_connections_percent.setter
    def max_connections_percent(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_connections_percent", value)

    @property
    @pulumi.getter(name="maxIdleConnectionsPercent")
    def max_idle_connections_percent(self) -> Optional[pulumi.Input[int]]:
        """
        Controls how actively the proxy closes idle database connections in the connection pool.
        """
        return pulumi.get(self, "max_idle_connections_percent")

    @max_idle_connections_percent.setter
    def max_idle_connections_percent(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_idle_connections_percent", value)

    @property
    @pulumi.getter(name="sessionPinningFilters")
    def session_pinning_filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Each item in the list represents a class of SQL operations that normally cause all later statements in a session using a proxy to be pinned to the same underlying database connection.
        """
        return pulumi.get(self, "session_pinning_filters")

    @session_pinning_filters.setter
    def session_pinning_filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "session_pinning_filters", value)


@pulumi.input_type
class DBSecurityGroupIngressArgs:
    def __init__(__self__, *,
                 c_idrip: Optional[pulumi.Input[str]] = None,
                 e_c2_security_group_id: Optional[pulumi.Input[str]] = None,
                 e_c2_security_group_name: Optional[pulumi.Input[str]] = None,
                 e_c2_security_group_owner_id: Optional[pulumi.Input[str]] = None):
        if c_idrip is not None:
            pulumi.set(__self__, "c_idrip", c_idrip)
        if e_c2_security_group_id is not None:
            pulumi.set(__self__, "e_c2_security_group_id", e_c2_security_group_id)
        if e_c2_security_group_name is not None:
            pulumi.set(__self__, "e_c2_security_group_name", e_c2_security_group_name)
        if e_c2_security_group_owner_id is not None:
            pulumi.set(__self__, "e_c2_security_group_owner_id", e_c2_security_group_owner_id)

    @property
    @pulumi.getter(name="cIDRIP")
    def c_idrip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "c_idrip")

    @c_idrip.setter
    def c_idrip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "c_idrip", value)

    @property
    @pulumi.getter(name="eC2SecurityGroupId")
    def e_c2_security_group_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "e_c2_security_group_id")

    @e_c2_security_group_id.setter
    def e_c2_security_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "e_c2_security_group_id", value)

    @property
    @pulumi.getter(name="eC2SecurityGroupName")
    def e_c2_security_group_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "e_c2_security_group_name")

    @e_c2_security_group_name.setter
    def e_c2_security_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "e_c2_security_group_name", value)

    @property
    @pulumi.getter(name="eC2SecurityGroupOwnerId")
    def e_c2_security_group_owner_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "e_c2_security_group_owner_id")

    @e_c2_security_group_owner_id.setter
    def e_c2_security_group_owner_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "e_c2_security_group_owner_id", value)


@pulumi.input_type
class DBSecurityGroupTagArgs:
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
class DBSubnetGroupTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        A key-value pair to associate with a resource.
        :param pulumi.Input[str] key: The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
        :param pulumi.Input[str] value: The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. 
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class EventSubscriptionTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        A key-value pair to associate with a resource.
        :param pulumi.Input[str] key: The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        :param pulumi.Input[str] value: The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class OptionGroupOptionConfigurationArgs:
    def __init__(__self__, *,
                 option_name: pulumi.Input[str],
                 d_b_security_group_memberships: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 option_settings: Optional[pulumi.Input[Sequence[pulumi.Input['OptionGroupOptionSettingArgs']]]] = None,
                 option_version: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 vpc_security_group_memberships: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        pulumi.set(__self__, "option_name", option_name)
        if d_b_security_group_memberships is not None:
            pulumi.set(__self__, "d_b_security_group_memberships", d_b_security_group_memberships)
        if option_settings is not None:
            pulumi.set(__self__, "option_settings", option_settings)
        if option_version is not None:
            pulumi.set(__self__, "option_version", option_version)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if vpc_security_group_memberships is not None:
            pulumi.set(__self__, "vpc_security_group_memberships", vpc_security_group_memberships)

    @property
    @pulumi.getter(name="optionName")
    def option_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "option_name")

    @option_name.setter
    def option_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "option_name", value)

    @property
    @pulumi.getter(name="dBSecurityGroupMemberships")
    def d_b_security_group_memberships(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "d_b_security_group_memberships")

    @d_b_security_group_memberships.setter
    def d_b_security_group_memberships(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "d_b_security_group_memberships", value)

    @property
    @pulumi.getter(name="optionSettings")
    def option_settings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['OptionGroupOptionSettingArgs']]]]:
        return pulumi.get(self, "option_settings")

    @option_settings.setter
    def option_settings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['OptionGroupOptionSettingArgs']]]]):
        pulumi.set(self, "option_settings", value)

    @property
    @pulumi.getter(name="optionVersion")
    def option_version(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "option_version")

    @option_version.setter
    def option_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "option_version", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="vpcSecurityGroupMemberships")
    def vpc_security_group_memberships(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "vpc_security_group_memberships")

    @vpc_security_group_memberships.setter
    def vpc_security_group_memberships(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "vpc_security_group_memberships", value)


@pulumi.input_type
class OptionGroupOptionSettingArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        if name is not None:
            pulumi.set(__self__, "name", name)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class OptionGroupTagArgs:
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


