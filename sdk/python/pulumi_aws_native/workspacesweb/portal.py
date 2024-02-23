# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs
from ._enums import *

__all__ = ['PortalArgs', 'Portal']

@pulumi.input_type
class PortalArgs:
    def __init__(__self__, *,
                 additional_encryption_context: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 authentication_type: Optional[pulumi.Input['PortalAuthenticationType']] = None,
                 browser_settings_arn: Optional[pulumi.Input[str]] = None,
                 customer_managed_key: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 ip_access_settings_arn: Optional[pulumi.Input[str]] = None,
                 network_settings_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None,
                 trust_store_arn: Optional[pulumi.Input[str]] = None,
                 user_access_logging_settings_arn: Optional[pulumi.Input[str]] = None,
                 user_settings_arn: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Portal resource.
        """
        if additional_encryption_context is not None:
            pulumi.set(__self__, "additional_encryption_context", additional_encryption_context)
        if authentication_type is not None:
            pulumi.set(__self__, "authentication_type", authentication_type)
        if browser_settings_arn is not None:
            pulumi.set(__self__, "browser_settings_arn", browser_settings_arn)
        if customer_managed_key is not None:
            pulumi.set(__self__, "customer_managed_key", customer_managed_key)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if ip_access_settings_arn is not None:
            pulumi.set(__self__, "ip_access_settings_arn", ip_access_settings_arn)
        if network_settings_arn is not None:
            pulumi.set(__self__, "network_settings_arn", network_settings_arn)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if trust_store_arn is not None:
            pulumi.set(__self__, "trust_store_arn", trust_store_arn)
        if user_access_logging_settings_arn is not None:
            pulumi.set(__self__, "user_access_logging_settings_arn", user_access_logging_settings_arn)
        if user_settings_arn is not None:
            pulumi.set(__self__, "user_settings_arn", user_settings_arn)

    @property
    @pulumi.getter(name="additionalEncryptionContext")
    def additional_encryption_context(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "additional_encryption_context")

    @additional_encryption_context.setter
    def additional_encryption_context(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "additional_encryption_context", value)

    @property
    @pulumi.getter(name="authenticationType")
    def authentication_type(self) -> Optional[pulumi.Input['PortalAuthenticationType']]:
        return pulumi.get(self, "authentication_type")

    @authentication_type.setter
    def authentication_type(self, value: Optional[pulumi.Input['PortalAuthenticationType']]):
        pulumi.set(self, "authentication_type", value)

    @property
    @pulumi.getter(name="browserSettingsArn")
    def browser_settings_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "browser_settings_arn")

    @browser_settings_arn.setter
    def browser_settings_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "browser_settings_arn", value)

    @property
    @pulumi.getter(name="customerManagedKey")
    def customer_managed_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "customer_managed_key")

    @customer_managed_key.setter
    def customer_managed_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "customer_managed_key", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="ipAccessSettingsArn")
    def ip_access_settings_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ip_access_settings_arn")

    @ip_access_settings_arn.setter
    def ip_access_settings_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_access_settings_arn", value)

    @property
    @pulumi.getter(name="networkSettingsArn")
    def network_settings_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "network_settings_arn")

    @network_settings_arn.setter
    def network_settings_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_settings_arn", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="trustStoreArn")
    def trust_store_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "trust_store_arn")

    @trust_store_arn.setter
    def trust_store_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trust_store_arn", value)

    @property
    @pulumi.getter(name="userAccessLoggingSettingsArn")
    def user_access_logging_settings_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "user_access_logging_settings_arn")

    @user_access_logging_settings_arn.setter
    def user_access_logging_settings_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_access_logging_settings_arn", value)

    @property
    @pulumi.getter(name="userSettingsArn")
    def user_settings_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "user_settings_arn")

    @user_settings_arn.setter
    def user_settings_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_settings_arn", value)


class Portal(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_encryption_context: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 authentication_type: Optional[pulumi.Input['PortalAuthenticationType']] = None,
                 browser_settings_arn: Optional[pulumi.Input[str]] = None,
                 customer_managed_key: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 ip_access_settings_arn: Optional[pulumi.Input[str]] = None,
                 network_settings_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 trust_store_arn: Optional[pulumi.Input[str]] = None,
                 user_access_logging_settings_arn: Optional[pulumi.Input[str]] = None,
                 user_settings_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Definition of AWS::WorkSpacesWeb::Portal Resource Type

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[PortalArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of AWS::WorkSpacesWeb::Portal Resource Type

        :param str resource_name: The name of the resource.
        :param PortalArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PortalArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_encryption_context: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 authentication_type: Optional[pulumi.Input['PortalAuthenticationType']] = None,
                 browser_settings_arn: Optional[pulumi.Input[str]] = None,
                 customer_managed_key: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 ip_access_settings_arn: Optional[pulumi.Input[str]] = None,
                 network_settings_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 trust_store_arn: Optional[pulumi.Input[str]] = None,
                 user_access_logging_settings_arn: Optional[pulumi.Input[str]] = None,
                 user_settings_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PortalArgs.__new__(PortalArgs)

            __props__.__dict__["additional_encryption_context"] = additional_encryption_context
            __props__.__dict__["authentication_type"] = authentication_type
            __props__.__dict__["browser_settings_arn"] = browser_settings_arn
            __props__.__dict__["customer_managed_key"] = customer_managed_key
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["ip_access_settings_arn"] = ip_access_settings_arn
            __props__.__dict__["network_settings_arn"] = network_settings_arn
            __props__.__dict__["tags"] = tags
            __props__.__dict__["trust_store_arn"] = trust_store_arn
            __props__.__dict__["user_access_logging_settings_arn"] = user_access_logging_settings_arn
            __props__.__dict__["user_settings_arn"] = user_settings_arn
            __props__.__dict__["browser_type"] = None
            __props__.__dict__["creation_date"] = None
            __props__.__dict__["portal_arn"] = None
            __props__.__dict__["portal_endpoint"] = None
            __props__.__dict__["portal_status"] = None
            __props__.__dict__["renderer_type"] = None
            __props__.__dict__["service_provider_saml_metadata"] = None
            __props__.__dict__["status_reason"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["additional_encryption_context.*", "customer_managed_key"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Portal, __self__).__init__(
            'aws-native:workspacesweb:Portal',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Portal':
        """
        Get an existing Portal resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PortalArgs.__new__(PortalArgs)

        __props__.__dict__["additional_encryption_context"] = None
        __props__.__dict__["authentication_type"] = None
        __props__.__dict__["browser_settings_arn"] = None
        __props__.__dict__["browser_type"] = None
        __props__.__dict__["creation_date"] = None
        __props__.__dict__["customer_managed_key"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["ip_access_settings_arn"] = None
        __props__.__dict__["network_settings_arn"] = None
        __props__.__dict__["portal_arn"] = None
        __props__.__dict__["portal_endpoint"] = None
        __props__.__dict__["portal_status"] = None
        __props__.__dict__["renderer_type"] = None
        __props__.__dict__["service_provider_saml_metadata"] = None
        __props__.__dict__["status_reason"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["trust_store_arn"] = None
        __props__.__dict__["user_access_logging_settings_arn"] = None
        __props__.__dict__["user_settings_arn"] = None
        return Portal(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalEncryptionContext")
    def additional_encryption_context(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "additional_encryption_context")

    @property
    @pulumi.getter(name="authenticationType")
    def authentication_type(self) -> pulumi.Output[Optional['PortalAuthenticationType']]:
        return pulumi.get(self, "authentication_type")

    @property
    @pulumi.getter(name="browserSettingsArn")
    def browser_settings_arn(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "browser_settings_arn")

    @property
    @pulumi.getter(name="browserType")
    def browser_type(self) -> pulumi.Output['PortalBrowserType']:
        return pulumi.get(self, "browser_type")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> pulumi.Output[str]:
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter(name="customerManagedKey")
    def customer_managed_key(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "customer_managed_key")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="ipAccessSettingsArn")
    def ip_access_settings_arn(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "ip_access_settings_arn")

    @property
    @pulumi.getter(name="networkSettingsArn")
    def network_settings_arn(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "network_settings_arn")

    @property
    @pulumi.getter(name="portalArn")
    def portal_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "portal_arn")

    @property
    @pulumi.getter(name="portalEndpoint")
    def portal_endpoint(self) -> pulumi.Output[str]:
        return pulumi.get(self, "portal_endpoint")

    @property
    @pulumi.getter(name="portalStatus")
    def portal_status(self) -> pulumi.Output['PortalStatus']:
        return pulumi.get(self, "portal_status")

    @property
    @pulumi.getter(name="rendererType")
    def renderer_type(self) -> pulumi.Output['PortalRendererType']:
        return pulumi.get(self, "renderer_type")

    @property
    @pulumi.getter(name="serviceProviderSamlMetadata")
    def service_provider_saml_metadata(self) -> pulumi.Output[str]:
        return pulumi.get(self, "service_provider_saml_metadata")

    @property
    @pulumi.getter(name="statusReason")
    def status_reason(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status_reason")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="trustStoreArn")
    def trust_store_arn(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "trust_store_arn")

    @property
    @pulumi.getter(name="userAccessLoggingSettingsArn")
    def user_access_logging_settings_arn(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "user_access_logging_settings_arn")

    @property
    @pulumi.getter(name="userSettingsArn")
    def user_settings_arn(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "user_settings_arn")

