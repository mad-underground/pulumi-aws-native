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
    'BrowserSettingsEncryptionContextMapArgs',
    'BrowserSettingsTagArgs',
    'IdentityProviderDetailsArgs',
    'IpAccessSettingsEncryptionContextMapArgs',
    'IpAccessSettingsIpRuleArgs',
    'IpAccessSettingsTagArgs',
    'NetworkSettingsTagArgs',
    'PortalEncryptionContextMapArgs',
    'PortalTagArgs',
    'TrustStoreTagArgs',
    'UserAccessLoggingSettingsTagArgs',
    'UserSettingsCookieSpecificationArgs',
    'UserSettingsCookieSynchronizationConfigurationArgs',
    'UserSettingsEncryptionContextMapArgs',
    'UserSettingsTagArgs',
]

@pulumi.input_type
class BrowserSettingsEncryptionContextMapArgs:
    def __init__(__self__):
        pass


@pulumi.input_type
class BrowserSettingsTagArgs:
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
class IdentityProviderDetailsArgs:
    def __init__(__self__):
        pass


@pulumi.input_type
class IpAccessSettingsEncryptionContextMapArgs:
    def __init__(__self__):
        pass


@pulumi.input_type
class IpAccessSettingsIpRuleArgs:
    def __init__(__self__, *,
                 ip_range: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] ip_range: A single IP address or an IP address range in CIDR notation
        """
        pulumi.set(__self__, "ip_range", ip_range)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="ipRange")
    def ip_range(self) -> pulumi.Input[str]:
        """
        A single IP address or an IP address range in CIDR notation
        """
        return pulumi.get(self, "ip_range")

    @ip_range.setter
    def ip_range(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip_range", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class IpAccessSettingsTagArgs:
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
class NetworkSettingsTagArgs:
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
class PortalEncryptionContextMapArgs:
    def __init__(__self__):
        pass


@pulumi.input_type
class PortalTagArgs:
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
class TrustStoreTagArgs:
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
class UserAccessLoggingSettingsTagArgs:
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
class UserSettingsCookieSpecificationArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "domain", domain)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if path is not None:
            pulumi.set(__self__, "path", path)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)


@pulumi.input_type
class UserSettingsCookieSynchronizationConfigurationArgs:
    def __init__(__self__, *,
                 allowlist: pulumi.Input[Sequence[pulumi.Input['UserSettingsCookieSpecificationArgs']]],
                 blocklist: Optional[pulumi.Input[Sequence[pulumi.Input['UserSettingsCookieSpecificationArgs']]]] = None):
        pulumi.set(__self__, "allowlist", allowlist)
        if blocklist is not None:
            pulumi.set(__self__, "blocklist", blocklist)

    @property
    @pulumi.getter
    def allowlist(self) -> pulumi.Input[Sequence[pulumi.Input['UserSettingsCookieSpecificationArgs']]]:
        return pulumi.get(self, "allowlist")

    @allowlist.setter
    def allowlist(self, value: pulumi.Input[Sequence[pulumi.Input['UserSettingsCookieSpecificationArgs']]]):
        pulumi.set(self, "allowlist", value)

    @property
    @pulumi.getter
    def blocklist(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UserSettingsCookieSpecificationArgs']]]]:
        return pulumi.get(self, "blocklist")

    @blocklist.setter
    def blocklist(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UserSettingsCookieSpecificationArgs']]]]):
        pulumi.set(self, "blocklist", value)


@pulumi.input_type
class UserSettingsEncryptionContextMapArgs:
    def __init__(__self__):
        pass


@pulumi.input_type
class UserSettingsTagArgs:
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

