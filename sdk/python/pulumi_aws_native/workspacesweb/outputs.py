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
    'BrowserSettingsEncryptionContextMap',
    'BrowserSettingsTag',
    'IdentityProviderDetails',
    'IpAccessSettingsEncryptionContextMap',
    'IpAccessSettingsIpRule',
    'IpAccessSettingsTag',
    'NetworkSettingsTag',
    'PortalEncryptionContextMap',
    'PortalTag',
    'TrustStoreTag',
    'UserAccessLoggingSettingsTag',
    'UserSettingsCookieSpecification',
    'UserSettingsCookieSynchronizationConfiguration',
    'UserSettingsEncryptionContextMap',
    'UserSettingsTag',
]

@pulumi.output_type
class BrowserSettingsEncryptionContextMap(dict):
    def __init__(__self__):
        pass


@pulumi.output_type
class BrowserSettingsTag(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


@pulumi.output_type
class IdentityProviderDetails(dict):
    def __init__(__self__):
        pass


@pulumi.output_type
class IpAccessSettingsEncryptionContextMap(dict):
    def __init__(__self__):
        pass


@pulumi.output_type
class IpAccessSettingsIpRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "ipRange":
            suggest = "ip_range"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in IpAccessSettingsIpRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        IpAccessSettingsIpRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        IpAccessSettingsIpRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 ip_range: str,
                 description: Optional[str] = None):
        """
        :param str ip_range: A single IP address or an IP address range in CIDR notation
        """
        pulumi.set(__self__, "ip_range", ip_range)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="ipRange")
    def ip_range(self) -> str:
        """
        A single IP address or an IP address range in CIDR notation
        """
        return pulumi.get(self, "ip_range")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")


@pulumi.output_type
class IpAccessSettingsTag(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


@pulumi.output_type
class NetworkSettingsTag(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


@pulumi.output_type
class PortalEncryptionContextMap(dict):
    def __init__(__self__):
        pass


@pulumi.output_type
class PortalTag(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


@pulumi.output_type
class TrustStoreTag(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


@pulumi.output_type
class UserAccessLoggingSettingsTag(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


@pulumi.output_type
class UserSettingsCookieSpecification(dict):
    def __init__(__self__, *,
                 domain: str,
                 name: Optional[str] = None,
                 path: Optional[str] = None):
        pulumi.set(__self__, "domain", domain)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if path is not None:
            pulumi.set(__self__, "path", path)

    @property
    @pulumi.getter
    def domain(self) -> str:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def path(self) -> Optional[str]:
        return pulumi.get(self, "path")


@pulumi.output_type
class UserSettingsCookieSynchronizationConfiguration(dict):
    def __init__(__self__, *,
                 allowlist: Sequence['outputs.UserSettingsCookieSpecification'],
                 blocklist: Optional[Sequence['outputs.UserSettingsCookieSpecification']] = None):
        pulumi.set(__self__, "allowlist", allowlist)
        if blocklist is not None:
            pulumi.set(__self__, "blocklist", blocklist)

    @property
    @pulumi.getter
    def allowlist(self) -> Sequence['outputs.UserSettingsCookieSpecification']:
        return pulumi.get(self, "allowlist")

    @property
    @pulumi.getter
    def blocklist(self) -> Optional[Sequence['outputs.UserSettingsCookieSpecification']]:
        return pulumi.get(self, "blocklist")


@pulumi.output_type
class UserSettingsEncryptionContextMap(dict):
    def __init__(__self__):
        pass


@pulumi.output_type
class UserSettingsTag(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")

