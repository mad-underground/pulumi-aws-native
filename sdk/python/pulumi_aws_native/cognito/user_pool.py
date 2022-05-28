# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['UserPoolArgs', 'UserPool']

@pulumi.input_type
class UserPoolArgs:
    def __init__(__self__, *,
                 account_recovery_setting: Optional[pulumi.Input['UserPoolAccountRecoverySettingArgs']] = None,
                 admin_create_user_config: Optional[pulumi.Input['UserPoolAdminCreateUserConfigArgs']] = None,
                 alias_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 auto_verified_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 device_configuration: Optional[pulumi.Input['UserPoolDeviceConfigurationArgs']] = None,
                 email_configuration: Optional[pulumi.Input['UserPoolEmailConfigurationArgs']] = None,
                 email_verification_message: Optional[pulumi.Input[str]] = None,
                 email_verification_subject: Optional[pulumi.Input[str]] = None,
                 enabled_mfas: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 lambda_config: Optional[pulumi.Input['UserPoolLambdaConfigArgs']] = None,
                 mfa_configuration: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input['UserPoolPoliciesArgs']] = None,
                 schema: Optional[pulumi.Input[Sequence[pulumi.Input['UserPoolSchemaAttributeArgs']]]] = None,
                 sms_authentication_message: Optional[pulumi.Input[str]] = None,
                 sms_configuration: Optional[pulumi.Input['UserPoolSmsConfigurationArgs']] = None,
                 sms_verification_message: Optional[pulumi.Input[str]] = None,
                 user_attribute_update_settings: Optional[pulumi.Input['UserPoolUserAttributeUpdateSettingsArgs']] = None,
                 user_pool_add_ons: Optional[pulumi.Input['UserPoolAddOnsArgs']] = None,
                 user_pool_name: Optional[pulumi.Input[str]] = None,
                 user_pool_tags: Optional[Any] = None,
                 username_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 username_configuration: Optional[pulumi.Input['UserPoolUsernameConfigurationArgs']] = None,
                 verification_message_template: Optional[pulumi.Input['UserPoolVerificationMessageTemplateArgs']] = None):
        """
        The set of arguments for constructing a UserPool resource.
        """
        if account_recovery_setting is not None:
            pulumi.set(__self__, "account_recovery_setting", account_recovery_setting)
        if admin_create_user_config is not None:
            pulumi.set(__self__, "admin_create_user_config", admin_create_user_config)
        if alias_attributes is not None:
            pulumi.set(__self__, "alias_attributes", alias_attributes)
        if auto_verified_attributes is not None:
            pulumi.set(__self__, "auto_verified_attributes", auto_verified_attributes)
        if device_configuration is not None:
            pulumi.set(__self__, "device_configuration", device_configuration)
        if email_configuration is not None:
            pulumi.set(__self__, "email_configuration", email_configuration)
        if email_verification_message is not None:
            pulumi.set(__self__, "email_verification_message", email_verification_message)
        if email_verification_subject is not None:
            pulumi.set(__self__, "email_verification_subject", email_verification_subject)
        if enabled_mfas is not None:
            pulumi.set(__self__, "enabled_mfas", enabled_mfas)
        if lambda_config is not None:
            pulumi.set(__self__, "lambda_config", lambda_config)
        if mfa_configuration is not None:
            pulumi.set(__self__, "mfa_configuration", mfa_configuration)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)
        if schema is not None:
            pulumi.set(__self__, "schema", schema)
        if sms_authentication_message is not None:
            pulumi.set(__self__, "sms_authentication_message", sms_authentication_message)
        if sms_configuration is not None:
            pulumi.set(__self__, "sms_configuration", sms_configuration)
        if sms_verification_message is not None:
            pulumi.set(__self__, "sms_verification_message", sms_verification_message)
        if user_attribute_update_settings is not None:
            pulumi.set(__self__, "user_attribute_update_settings", user_attribute_update_settings)
        if user_pool_add_ons is not None:
            pulumi.set(__self__, "user_pool_add_ons", user_pool_add_ons)
        if user_pool_name is not None:
            pulumi.set(__self__, "user_pool_name", user_pool_name)
        if user_pool_tags is not None:
            pulumi.set(__self__, "user_pool_tags", user_pool_tags)
        if username_attributes is not None:
            pulumi.set(__self__, "username_attributes", username_attributes)
        if username_configuration is not None:
            pulumi.set(__self__, "username_configuration", username_configuration)
        if verification_message_template is not None:
            pulumi.set(__self__, "verification_message_template", verification_message_template)

    @property
    @pulumi.getter(name="accountRecoverySetting")
    def account_recovery_setting(self) -> Optional[pulumi.Input['UserPoolAccountRecoverySettingArgs']]:
        return pulumi.get(self, "account_recovery_setting")

    @account_recovery_setting.setter
    def account_recovery_setting(self, value: Optional[pulumi.Input['UserPoolAccountRecoverySettingArgs']]):
        pulumi.set(self, "account_recovery_setting", value)

    @property
    @pulumi.getter(name="adminCreateUserConfig")
    def admin_create_user_config(self) -> Optional[pulumi.Input['UserPoolAdminCreateUserConfigArgs']]:
        return pulumi.get(self, "admin_create_user_config")

    @admin_create_user_config.setter
    def admin_create_user_config(self, value: Optional[pulumi.Input['UserPoolAdminCreateUserConfigArgs']]):
        pulumi.set(self, "admin_create_user_config", value)

    @property
    @pulumi.getter(name="aliasAttributes")
    def alias_attributes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "alias_attributes")

    @alias_attributes.setter
    def alias_attributes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "alias_attributes", value)

    @property
    @pulumi.getter(name="autoVerifiedAttributes")
    def auto_verified_attributes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "auto_verified_attributes")

    @auto_verified_attributes.setter
    def auto_verified_attributes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "auto_verified_attributes", value)

    @property
    @pulumi.getter(name="deviceConfiguration")
    def device_configuration(self) -> Optional[pulumi.Input['UserPoolDeviceConfigurationArgs']]:
        return pulumi.get(self, "device_configuration")

    @device_configuration.setter
    def device_configuration(self, value: Optional[pulumi.Input['UserPoolDeviceConfigurationArgs']]):
        pulumi.set(self, "device_configuration", value)

    @property
    @pulumi.getter(name="emailConfiguration")
    def email_configuration(self) -> Optional[pulumi.Input['UserPoolEmailConfigurationArgs']]:
        return pulumi.get(self, "email_configuration")

    @email_configuration.setter
    def email_configuration(self, value: Optional[pulumi.Input['UserPoolEmailConfigurationArgs']]):
        pulumi.set(self, "email_configuration", value)

    @property
    @pulumi.getter(name="emailVerificationMessage")
    def email_verification_message(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "email_verification_message")

    @email_verification_message.setter
    def email_verification_message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email_verification_message", value)

    @property
    @pulumi.getter(name="emailVerificationSubject")
    def email_verification_subject(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "email_verification_subject")

    @email_verification_subject.setter
    def email_verification_subject(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email_verification_subject", value)

    @property
    @pulumi.getter(name="enabledMfas")
    def enabled_mfas(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "enabled_mfas")

    @enabled_mfas.setter
    def enabled_mfas(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "enabled_mfas", value)

    @property
    @pulumi.getter(name="lambdaConfig")
    def lambda_config(self) -> Optional[pulumi.Input['UserPoolLambdaConfigArgs']]:
        return pulumi.get(self, "lambda_config")

    @lambda_config.setter
    def lambda_config(self, value: Optional[pulumi.Input['UserPoolLambdaConfigArgs']]):
        pulumi.set(self, "lambda_config", value)

    @property
    @pulumi.getter(name="mfaConfiguration")
    def mfa_configuration(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mfa_configuration")

    @mfa_configuration.setter
    def mfa_configuration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mfa_configuration", value)

    @property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input['UserPoolPoliciesArgs']]:
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input['UserPoolPoliciesArgs']]):
        pulumi.set(self, "policies", value)

    @property
    @pulumi.getter
    def schema(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UserPoolSchemaAttributeArgs']]]]:
        return pulumi.get(self, "schema")

    @schema.setter
    def schema(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UserPoolSchemaAttributeArgs']]]]):
        pulumi.set(self, "schema", value)

    @property
    @pulumi.getter(name="smsAuthenticationMessage")
    def sms_authentication_message(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "sms_authentication_message")

    @sms_authentication_message.setter
    def sms_authentication_message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sms_authentication_message", value)

    @property
    @pulumi.getter(name="smsConfiguration")
    def sms_configuration(self) -> Optional[pulumi.Input['UserPoolSmsConfigurationArgs']]:
        return pulumi.get(self, "sms_configuration")

    @sms_configuration.setter
    def sms_configuration(self, value: Optional[pulumi.Input['UserPoolSmsConfigurationArgs']]):
        pulumi.set(self, "sms_configuration", value)

    @property
    @pulumi.getter(name="smsVerificationMessage")
    def sms_verification_message(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "sms_verification_message")

    @sms_verification_message.setter
    def sms_verification_message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sms_verification_message", value)

    @property
    @pulumi.getter(name="userAttributeUpdateSettings")
    def user_attribute_update_settings(self) -> Optional[pulumi.Input['UserPoolUserAttributeUpdateSettingsArgs']]:
        return pulumi.get(self, "user_attribute_update_settings")

    @user_attribute_update_settings.setter
    def user_attribute_update_settings(self, value: Optional[pulumi.Input['UserPoolUserAttributeUpdateSettingsArgs']]):
        pulumi.set(self, "user_attribute_update_settings", value)

    @property
    @pulumi.getter(name="userPoolAddOns")
    def user_pool_add_ons(self) -> Optional[pulumi.Input['UserPoolAddOnsArgs']]:
        return pulumi.get(self, "user_pool_add_ons")

    @user_pool_add_ons.setter
    def user_pool_add_ons(self, value: Optional[pulumi.Input['UserPoolAddOnsArgs']]):
        pulumi.set(self, "user_pool_add_ons", value)

    @property
    @pulumi.getter(name="userPoolName")
    def user_pool_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "user_pool_name")

    @user_pool_name.setter
    def user_pool_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_pool_name", value)

    @property
    @pulumi.getter(name="userPoolTags")
    def user_pool_tags(self) -> Optional[Any]:
        return pulumi.get(self, "user_pool_tags")

    @user_pool_tags.setter
    def user_pool_tags(self, value: Optional[Any]):
        pulumi.set(self, "user_pool_tags", value)

    @property
    @pulumi.getter(name="usernameAttributes")
    def username_attributes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "username_attributes")

    @username_attributes.setter
    def username_attributes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "username_attributes", value)

    @property
    @pulumi.getter(name="usernameConfiguration")
    def username_configuration(self) -> Optional[pulumi.Input['UserPoolUsernameConfigurationArgs']]:
        return pulumi.get(self, "username_configuration")

    @username_configuration.setter
    def username_configuration(self, value: Optional[pulumi.Input['UserPoolUsernameConfigurationArgs']]):
        pulumi.set(self, "username_configuration", value)

    @property
    @pulumi.getter(name="verificationMessageTemplate")
    def verification_message_template(self) -> Optional[pulumi.Input['UserPoolVerificationMessageTemplateArgs']]:
        return pulumi.get(self, "verification_message_template")

    @verification_message_template.setter
    def verification_message_template(self, value: Optional[pulumi.Input['UserPoolVerificationMessageTemplateArgs']]):
        pulumi.set(self, "verification_message_template", value)


warnings.warn("""UserPool is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class UserPool(pulumi.CustomResource):
    warnings.warn("""UserPool is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_recovery_setting: Optional[pulumi.Input[pulumi.InputType['UserPoolAccountRecoverySettingArgs']]] = None,
                 admin_create_user_config: Optional[pulumi.Input[pulumi.InputType['UserPoolAdminCreateUserConfigArgs']]] = None,
                 alias_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 auto_verified_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 device_configuration: Optional[pulumi.Input[pulumi.InputType['UserPoolDeviceConfigurationArgs']]] = None,
                 email_configuration: Optional[pulumi.Input[pulumi.InputType['UserPoolEmailConfigurationArgs']]] = None,
                 email_verification_message: Optional[pulumi.Input[str]] = None,
                 email_verification_subject: Optional[pulumi.Input[str]] = None,
                 enabled_mfas: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 lambda_config: Optional[pulumi.Input[pulumi.InputType['UserPoolLambdaConfigArgs']]] = None,
                 mfa_configuration: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[pulumi.InputType['UserPoolPoliciesArgs']]] = None,
                 schema: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UserPoolSchemaAttributeArgs']]]]] = None,
                 sms_authentication_message: Optional[pulumi.Input[str]] = None,
                 sms_configuration: Optional[pulumi.Input[pulumi.InputType['UserPoolSmsConfigurationArgs']]] = None,
                 sms_verification_message: Optional[pulumi.Input[str]] = None,
                 user_attribute_update_settings: Optional[pulumi.Input[pulumi.InputType['UserPoolUserAttributeUpdateSettingsArgs']]] = None,
                 user_pool_add_ons: Optional[pulumi.Input[pulumi.InputType['UserPoolAddOnsArgs']]] = None,
                 user_pool_name: Optional[pulumi.Input[str]] = None,
                 user_pool_tags: Optional[Any] = None,
                 username_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 username_configuration: Optional[pulumi.Input[pulumi.InputType['UserPoolUsernameConfigurationArgs']]] = None,
                 verification_message_template: Optional[pulumi.Input[pulumi.InputType['UserPoolVerificationMessageTemplateArgs']]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::Cognito::UserPool

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[UserPoolArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::Cognito::UserPool

        :param str resource_name: The name of the resource.
        :param UserPoolArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserPoolArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_recovery_setting: Optional[pulumi.Input[pulumi.InputType['UserPoolAccountRecoverySettingArgs']]] = None,
                 admin_create_user_config: Optional[pulumi.Input[pulumi.InputType['UserPoolAdminCreateUserConfigArgs']]] = None,
                 alias_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 auto_verified_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 device_configuration: Optional[pulumi.Input[pulumi.InputType['UserPoolDeviceConfigurationArgs']]] = None,
                 email_configuration: Optional[pulumi.Input[pulumi.InputType['UserPoolEmailConfigurationArgs']]] = None,
                 email_verification_message: Optional[pulumi.Input[str]] = None,
                 email_verification_subject: Optional[pulumi.Input[str]] = None,
                 enabled_mfas: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 lambda_config: Optional[pulumi.Input[pulumi.InputType['UserPoolLambdaConfigArgs']]] = None,
                 mfa_configuration: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[pulumi.InputType['UserPoolPoliciesArgs']]] = None,
                 schema: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UserPoolSchemaAttributeArgs']]]]] = None,
                 sms_authentication_message: Optional[pulumi.Input[str]] = None,
                 sms_configuration: Optional[pulumi.Input[pulumi.InputType['UserPoolSmsConfigurationArgs']]] = None,
                 sms_verification_message: Optional[pulumi.Input[str]] = None,
                 user_attribute_update_settings: Optional[pulumi.Input[pulumi.InputType['UserPoolUserAttributeUpdateSettingsArgs']]] = None,
                 user_pool_add_ons: Optional[pulumi.Input[pulumi.InputType['UserPoolAddOnsArgs']]] = None,
                 user_pool_name: Optional[pulumi.Input[str]] = None,
                 user_pool_tags: Optional[Any] = None,
                 username_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 username_configuration: Optional[pulumi.Input[pulumi.InputType['UserPoolUsernameConfigurationArgs']]] = None,
                 verification_message_template: Optional[pulumi.Input[pulumi.InputType['UserPoolVerificationMessageTemplateArgs']]] = None,
                 __props__=None):
        pulumi.log.warn("""UserPool is deprecated: UserPool is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserPoolArgs.__new__(UserPoolArgs)

            __props__.__dict__["account_recovery_setting"] = account_recovery_setting
            __props__.__dict__["admin_create_user_config"] = admin_create_user_config
            __props__.__dict__["alias_attributes"] = alias_attributes
            __props__.__dict__["auto_verified_attributes"] = auto_verified_attributes
            __props__.__dict__["device_configuration"] = device_configuration
            __props__.__dict__["email_configuration"] = email_configuration
            __props__.__dict__["email_verification_message"] = email_verification_message
            __props__.__dict__["email_verification_subject"] = email_verification_subject
            __props__.__dict__["enabled_mfas"] = enabled_mfas
            __props__.__dict__["lambda_config"] = lambda_config
            __props__.__dict__["mfa_configuration"] = mfa_configuration
            __props__.__dict__["policies"] = policies
            __props__.__dict__["schema"] = schema
            __props__.__dict__["sms_authentication_message"] = sms_authentication_message
            __props__.__dict__["sms_configuration"] = sms_configuration
            __props__.__dict__["sms_verification_message"] = sms_verification_message
            __props__.__dict__["user_attribute_update_settings"] = user_attribute_update_settings
            __props__.__dict__["user_pool_add_ons"] = user_pool_add_ons
            __props__.__dict__["user_pool_name"] = user_pool_name
            __props__.__dict__["user_pool_tags"] = user_pool_tags
            __props__.__dict__["username_attributes"] = username_attributes
            __props__.__dict__["username_configuration"] = username_configuration
            __props__.__dict__["verification_message_template"] = verification_message_template
            __props__.__dict__["arn"] = None
            __props__.__dict__["provider_name"] = None
            __props__.__dict__["provider_url"] = None
        super(UserPool, __self__).__init__(
            'aws-native:cognito:UserPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'UserPool':
        """
        Get an existing UserPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = UserPoolArgs.__new__(UserPoolArgs)

        __props__.__dict__["account_recovery_setting"] = None
        __props__.__dict__["admin_create_user_config"] = None
        __props__.__dict__["alias_attributes"] = None
        __props__.__dict__["arn"] = None
        __props__.__dict__["auto_verified_attributes"] = None
        __props__.__dict__["device_configuration"] = None
        __props__.__dict__["email_configuration"] = None
        __props__.__dict__["email_verification_message"] = None
        __props__.__dict__["email_verification_subject"] = None
        __props__.__dict__["enabled_mfas"] = None
        __props__.__dict__["lambda_config"] = None
        __props__.__dict__["mfa_configuration"] = None
        __props__.__dict__["policies"] = None
        __props__.__dict__["provider_name"] = None
        __props__.__dict__["provider_url"] = None
        __props__.__dict__["schema"] = None
        __props__.__dict__["sms_authentication_message"] = None
        __props__.__dict__["sms_configuration"] = None
        __props__.__dict__["sms_verification_message"] = None
        __props__.__dict__["user_attribute_update_settings"] = None
        __props__.__dict__["user_pool_add_ons"] = None
        __props__.__dict__["user_pool_name"] = None
        __props__.__dict__["user_pool_tags"] = None
        __props__.__dict__["username_attributes"] = None
        __props__.__dict__["username_configuration"] = None
        __props__.__dict__["verification_message_template"] = None
        return UserPool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountRecoverySetting")
    def account_recovery_setting(self) -> pulumi.Output[Optional['outputs.UserPoolAccountRecoverySetting']]:
        return pulumi.get(self, "account_recovery_setting")

    @property
    @pulumi.getter(name="adminCreateUserConfig")
    def admin_create_user_config(self) -> pulumi.Output[Optional['outputs.UserPoolAdminCreateUserConfig']]:
        return pulumi.get(self, "admin_create_user_config")

    @property
    @pulumi.getter(name="aliasAttributes")
    def alias_attributes(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "alias_attributes")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="autoVerifiedAttributes")
    def auto_verified_attributes(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "auto_verified_attributes")

    @property
    @pulumi.getter(name="deviceConfiguration")
    def device_configuration(self) -> pulumi.Output[Optional['outputs.UserPoolDeviceConfiguration']]:
        return pulumi.get(self, "device_configuration")

    @property
    @pulumi.getter(name="emailConfiguration")
    def email_configuration(self) -> pulumi.Output[Optional['outputs.UserPoolEmailConfiguration']]:
        return pulumi.get(self, "email_configuration")

    @property
    @pulumi.getter(name="emailVerificationMessage")
    def email_verification_message(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "email_verification_message")

    @property
    @pulumi.getter(name="emailVerificationSubject")
    def email_verification_subject(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "email_verification_subject")

    @property
    @pulumi.getter(name="enabledMfas")
    def enabled_mfas(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "enabled_mfas")

    @property
    @pulumi.getter(name="lambdaConfig")
    def lambda_config(self) -> pulumi.Output[Optional['outputs.UserPoolLambdaConfig']]:
        return pulumi.get(self, "lambda_config")

    @property
    @pulumi.getter(name="mfaConfiguration")
    def mfa_configuration(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "mfa_configuration")

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Output[Optional['outputs.UserPoolPolicies']]:
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter(name="providerName")
    def provider_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "provider_name")

    @property
    @pulumi.getter(name="providerURL")
    def provider_url(self) -> pulumi.Output[str]:
        return pulumi.get(self, "provider_url")

    @property
    @pulumi.getter
    def schema(self) -> pulumi.Output[Optional[Sequence['outputs.UserPoolSchemaAttribute']]]:
        return pulumi.get(self, "schema")

    @property
    @pulumi.getter(name="smsAuthenticationMessage")
    def sms_authentication_message(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "sms_authentication_message")

    @property
    @pulumi.getter(name="smsConfiguration")
    def sms_configuration(self) -> pulumi.Output[Optional['outputs.UserPoolSmsConfiguration']]:
        return pulumi.get(self, "sms_configuration")

    @property
    @pulumi.getter(name="smsVerificationMessage")
    def sms_verification_message(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "sms_verification_message")

    @property
    @pulumi.getter(name="userAttributeUpdateSettings")
    def user_attribute_update_settings(self) -> pulumi.Output[Optional['outputs.UserPoolUserAttributeUpdateSettings']]:
        return pulumi.get(self, "user_attribute_update_settings")

    @property
    @pulumi.getter(name="userPoolAddOns")
    def user_pool_add_ons(self) -> pulumi.Output[Optional['outputs.UserPoolAddOns']]:
        return pulumi.get(self, "user_pool_add_ons")

    @property
    @pulumi.getter(name="userPoolName")
    def user_pool_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "user_pool_name")

    @property
    @pulumi.getter(name="userPoolTags")
    def user_pool_tags(self) -> pulumi.Output[Optional[Any]]:
        return pulumi.get(self, "user_pool_tags")

    @property
    @pulumi.getter(name="usernameAttributes")
    def username_attributes(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "username_attributes")

    @property
    @pulumi.getter(name="usernameConfiguration")
    def username_configuration(self) -> pulumi.Output[Optional['outputs.UserPoolUsernameConfiguration']]:
        return pulumi.get(self, "username_configuration")

    @property
    @pulumi.getter(name="verificationMessageTemplate")
    def verification_message_template(self) -> pulumi.Output[Optional['outputs.UserPoolVerificationMessageTemplate']]:
        return pulumi.get(self, "verification_message_template")

