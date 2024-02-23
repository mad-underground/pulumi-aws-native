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
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs
from ._enums import *
from ._inputs import *

__all__ = ['WebAclArgs', 'WebAcl']

@pulumi.input_type
class WebAclArgs:
    def __init__(__self__, *,
                 default_action: pulumi.Input['WebAclDefaultActionArgs'],
                 scope: pulumi.Input['WebAclScope'],
                 visibility_config: pulumi.Input['WebAclVisibilityConfigArgs'],
                 association_config: Optional[pulumi.Input['WebAclAssociationConfigArgs']] = None,
                 captcha_config: Optional[pulumi.Input['WebAclCaptchaConfigArgs']] = None,
                 challenge_config: Optional[pulumi.Input['WebAclChallengeConfigArgs']] = None,
                 custom_response_bodies: Optional[pulumi.Input[Mapping[str, pulumi.Input['WebAclCustomResponseBodyArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['WebAclRuleArgs']]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None,
                 token_domains: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a WebAcl resource.
        :param pulumi.Input[Sequence[pulumi.Input['WebAclRuleArgs']]] rules: Collection of Rules.
        """
        pulumi.set(__self__, "default_action", default_action)
        pulumi.set(__self__, "scope", scope)
        pulumi.set(__self__, "visibility_config", visibility_config)
        if association_config is not None:
            pulumi.set(__self__, "association_config", association_config)
        if captcha_config is not None:
            pulumi.set(__self__, "captcha_config", captcha_config)
        if challenge_config is not None:
            pulumi.set(__self__, "challenge_config", challenge_config)
        if custom_response_bodies is not None:
            pulumi.set(__self__, "custom_response_bodies", custom_response_bodies)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if token_domains is not None:
            pulumi.set(__self__, "token_domains", token_domains)

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> pulumi.Input['WebAclDefaultActionArgs']:
        return pulumi.get(self, "default_action")

    @default_action.setter
    def default_action(self, value: pulumi.Input['WebAclDefaultActionArgs']):
        pulumi.set(self, "default_action", value)

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Input['WebAclScope']:
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: pulumi.Input['WebAclScope']):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter(name="visibilityConfig")
    def visibility_config(self) -> pulumi.Input['WebAclVisibilityConfigArgs']:
        return pulumi.get(self, "visibility_config")

    @visibility_config.setter
    def visibility_config(self, value: pulumi.Input['WebAclVisibilityConfigArgs']):
        pulumi.set(self, "visibility_config", value)

    @property
    @pulumi.getter(name="associationConfig")
    def association_config(self) -> Optional[pulumi.Input['WebAclAssociationConfigArgs']]:
        return pulumi.get(self, "association_config")

    @association_config.setter
    def association_config(self, value: Optional[pulumi.Input['WebAclAssociationConfigArgs']]):
        pulumi.set(self, "association_config", value)

    @property
    @pulumi.getter(name="captchaConfig")
    def captcha_config(self) -> Optional[pulumi.Input['WebAclCaptchaConfigArgs']]:
        return pulumi.get(self, "captcha_config")

    @captcha_config.setter
    def captcha_config(self, value: Optional[pulumi.Input['WebAclCaptchaConfigArgs']]):
        pulumi.set(self, "captcha_config", value)

    @property
    @pulumi.getter(name="challengeConfig")
    def challenge_config(self) -> Optional[pulumi.Input['WebAclChallengeConfigArgs']]:
        return pulumi.get(self, "challenge_config")

    @challenge_config.setter
    def challenge_config(self, value: Optional[pulumi.Input['WebAclChallengeConfigArgs']]):
        pulumi.set(self, "challenge_config", value)

    @property
    @pulumi.getter(name="customResponseBodies")
    def custom_response_bodies(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input['WebAclCustomResponseBodyArgs']]]]:
        return pulumi.get(self, "custom_response_bodies")

    @custom_response_bodies.setter
    def custom_response_bodies(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input['WebAclCustomResponseBodyArgs']]]]):
        pulumi.set(self, "custom_response_bodies", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WebAclRuleArgs']]]]:
        """
        Collection of Rules.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WebAclRuleArgs']]]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tokenDomains")
    def token_domains(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "token_domains")

    @token_domains.setter
    def token_domains(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "token_domains", value)


class WebAcl(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 association_config: Optional[pulumi.Input[pulumi.InputType['WebAclAssociationConfigArgs']]] = None,
                 captcha_config: Optional[pulumi.Input[pulumi.InputType['WebAclCaptchaConfigArgs']]] = None,
                 challenge_config: Optional[pulumi.Input[pulumi.InputType['WebAclChallengeConfigArgs']]] = None,
                 custom_response_bodies: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['WebAclCustomResponseBodyArgs']]]]] = None,
                 default_action: Optional[pulumi.Input[pulumi.InputType['WebAclDefaultActionArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WebAclRuleArgs']]]]] = None,
                 scope: Optional[pulumi.Input['WebAclScope']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 token_domains: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 visibility_config: Optional[pulumi.Input[pulumi.InputType['WebAclVisibilityConfigArgs']]] = None,
                 __props__=None):
        """
        Contains the Rules that identify the requests that you want to allow, block, or count. In a WebACL, you also specify a default action (ALLOW or BLOCK), and the action for each Rule that you add to a WebACL, for example, block requests from specified IP addresses or block requests from specified referrers. You also associate the WebACL with a CloudFront distribution to identify the requests that you want AWS WAF to filter. If you add more than one Rule to a WebACL, a request needs to match only one of the specifications to be allowed, blocked, or counted.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WebAclRuleArgs']]]] rules: Collection of Rules.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WebAclArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Contains the Rules that identify the requests that you want to allow, block, or count. In a WebACL, you also specify a default action (ALLOW or BLOCK), and the action for each Rule that you add to a WebACL, for example, block requests from specified IP addresses or block requests from specified referrers. You also associate the WebACL with a CloudFront distribution to identify the requests that you want AWS WAF to filter. If you add more than one Rule to a WebACL, a request needs to match only one of the specifications to be allowed, blocked, or counted.

        :param str resource_name: The name of the resource.
        :param WebAclArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebAclArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 association_config: Optional[pulumi.Input[pulumi.InputType['WebAclAssociationConfigArgs']]] = None,
                 captcha_config: Optional[pulumi.Input[pulumi.InputType['WebAclCaptchaConfigArgs']]] = None,
                 challenge_config: Optional[pulumi.Input[pulumi.InputType['WebAclChallengeConfigArgs']]] = None,
                 custom_response_bodies: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['WebAclCustomResponseBodyArgs']]]]] = None,
                 default_action: Optional[pulumi.Input[pulumi.InputType['WebAclDefaultActionArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WebAclRuleArgs']]]]] = None,
                 scope: Optional[pulumi.Input['WebAclScope']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 token_domains: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 visibility_config: Optional[pulumi.Input[pulumi.InputType['WebAclVisibilityConfigArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WebAclArgs.__new__(WebAclArgs)

            __props__.__dict__["association_config"] = association_config
            __props__.__dict__["captcha_config"] = captcha_config
            __props__.__dict__["challenge_config"] = challenge_config
            __props__.__dict__["custom_response_bodies"] = custom_response_bodies
            if default_action is None and not opts.urn:
                raise TypeError("Missing required property 'default_action'")
            __props__.__dict__["default_action"] = default_action
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["rules"] = rules
            if scope is None and not opts.urn:
                raise TypeError("Missing required property 'scope'")
            __props__.__dict__["scope"] = scope
            __props__.__dict__["tags"] = tags
            __props__.__dict__["token_domains"] = token_domains
            if visibility_config is None and not opts.urn:
                raise TypeError("Missing required property 'visibility_config'")
            __props__.__dict__["visibility_config"] = visibility_config
            __props__.__dict__["arn"] = None
            __props__.__dict__["capacity"] = None
            __props__.__dict__["label_namespace"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["name", "scope"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(WebAcl, __self__).__init__(
            'aws-native:wafv2:WebAcl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'WebAcl':
        """
        Get an existing WebAcl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = WebAclArgs.__new__(WebAclArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["association_config"] = None
        __props__.__dict__["capacity"] = None
        __props__.__dict__["captcha_config"] = None
        __props__.__dict__["challenge_config"] = None
        __props__.__dict__["custom_response_bodies"] = None
        __props__.__dict__["default_action"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["label_namespace"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["rules"] = None
        __props__.__dict__["scope"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["token_domains"] = None
        __props__.__dict__["visibility_config"] = None
        return WebAcl(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="associationConfig")
    def association_config(self) -> pulumi.Output[Optional['outputs.WebAclAssociationConfig']]:
        return pulumi.get(self, "association_config")

    @property
    @pulumi.getter
    def capacity(self) -> pulumi.Output[int]:
        return pulumi.get(self, "capacity")

    @property
    @pulumi.getter(name="captchaConfig")
    def captcha_config(self) -> pulumi.Output[Optional['outputs.WebAclCaptchaConfig']]:
        return pulumi.get(self, "captcha_config")

    @property
    @pulumi.getter(name="challengeConfig")
    def challenge_config(self) -> pulumi.Output[Optional['outputs.WebAclChallengeConfig']]:
        return pulumi.get(self, "challenge_config")

    @property
    @pulumi.getter(name="customResponseBodies")
    def custom_response_bodies(self) -> pulumi.Output[Optional[Mapping[str, 'outputs.WebAclCustomResponseBody']]]:
        return pulumi.get(self, "custom_response_bodies")

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> pulumi.Output['outputs.WebAclDefaultAction']:
        return pulumi.get(self, "default_action")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="labelNamespace")
    def label_namespace(self) -> pulumi.Output[str]:
        return pulumi.get(self, "label_namespace")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Optional[Sequence['outputs.WebAclRule']]]:
        """
        Collection of Rules.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output['WebAclScope']:
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tokenDomains")
    def token_domains(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "token_domains")

    @property
    @pulumi.getter(name="visibilityConfig")
    def visibility_config(self) -> pulumi.Output['outputs.WebAclVisibilityConfig']:
        return pulumi.get(self, "visibility_config")

