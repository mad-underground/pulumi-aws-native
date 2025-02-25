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

__all__ = ['PolicyArgs', 'Policy']

@pulumi.input_type
class PolicyArgs:
    def __init__(__self__, *,
                 exclude_resource_tags: pulumi.Input[bool],
                 remediation_enabled: pulumi.Input[bool],
                 security_service_policy_data: pulumi.Input['PolicySecurityServicePolicyDataArgs'],
                 delete_all_policy_resources: Optional[pulumi.Input[bool]] = None,
                 exclude_map: Optional[pulumi.Input['PolicyIeMapArgs']] = None,
                 include_map: Optional[pulumi.Input['PolicyIeMapArgs']] = None,
                 policy_description: Optional[pulumi.Input[str]] = None,
                 policy_name: Optional[pulumi.Input[str]] = None,
                 resource_set_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_tags: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyResourceTagArgs']]]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 resource_type_list: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resources_clean_up: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a Policy resource.
        """
        pulumi.set(__self__, "exclude_resource_tags", exclude_resource_tags)
        pulumi.set(__self__, "remediation_enabled", remediation_enabled)
        pulumi.set(__self__, "security_service_policy_data", security_service_policy_data)
        if delete_all_policy_resources is not None:
            pulumi.set(__self__, "delete_all_policy_resources", delete_all_policy_resources)
        if exclude_map is not None:
            pulumi.set(__self__, "exclude_map", exclude_map)
        if include_map is not None:
            pulumi.set(__self__, "include_map", include_map)
        if policy_description is not None:
            pulumi.set(__self__, "policy_description", policy_description)
        if policy_name is not None:
            pulumi.set(__self__, "policy_name", policy_name)
        if resource_set_ids is not None:
            pulumi.set(__self__, "resource_set_ids", resource_set_ids)
        if resource_tags is not None:
            pulumi.set(__self__, "resource_tags", resource_tags)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)
        if resource_type_list is not None:
            pulumi.set(__self__, "resource_type_list", resource_type_list)
        if resources_clean_up is not None:
            pulumi.set(__self__, "resources_clean_up", resources_clean_up)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="excludeResourceTags")
    def exclude_resource_tags(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "exclude_resource_tags")

    @exclude_resource_tags.setter
    def exclude_resource_tags(self, value: pulumi.Input[bool]):
        pulumi.set(self, "exclude_resource_tags", value)

    @property
    @pulumi.getter(name="remediationEnabled")
    def remediation_enabled(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "remediation_enabled")

    @remediation_enabled.setter
    def remediation_enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "remediation_enabled", value)

    @property
    @pulumi.getter(name="securityServicePolicyData")
    def security_service_policy_data(self) -> pulumi.Input['PolicySecurityServicePolicyDataArgs']:
        return pulumi.get(self, "security_service_policy_data")

    @security_service_policy_data.setter
    def security_service_policy_data(self, value: pulumi.Input['PolicySecurityServicePolicyDataArgs']):
        pulumi.set(self, "security_service_policy_data", value)

    @property
    @pulumi.getter(name="deleteAllPolicyResources")
    def delete_all_policy_resources(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "delete_all_policy_resources")

    @delete_all_policy_resources.setter
    def delete_all_policy_resources(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_all_policy_resources", value)

    @property
    @pulumi.getter(name="excludeMap")
    def exclude_map(self) -> Optional[pulumi.Input['PolicyIeMapArgs']]:
        return pulumi.get(self, "exclude_map")

    @exclude_map.setter
    def exclude_map(self, value: Optional[pulumi.Input['PolicyIeMapArgs']]):
        pulumi.set(self, "exclude_map", value)

    @property
    @pulumi.getter(name="includeMap")
    def include_map(self) -> Optional[pulumi.Input['PolicyIeMapArgs']]:
        return pulumi.get(self, "include_map")

    @include_map.setter
    def include_map(self, value: Optional[pulumi.Input['PolicyIeMapArgs']]):
        pulumi.set(self, "include_map", value)

    @property
    @pulumi.getter(name="policyDescription")
    def policy_description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "policy_description")

    @policy_description.setter
    def policy_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_description", value)

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "policy_name")

    @policy_name.setter
    def policy_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_name", value)

    @property
    @pulumi.getter(name="resourceSetIds")
    def resource_set_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "resource_set_ids")

    @resource_set_ids.setter
    def resource_set_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "resource_set_ids", value)

    @property
    @pulumi.getter(name="resourceTags")
    def resource_tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PolicyResourceTagArgs']]]]:
        return pulumi.get(self, "resource_tags")

    @resource_tags.setter
    def resource_tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyResourceTagArgs']]]]):
        pulumi.set(self, "resource_tags", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_type", value)

    @property
    @pulumi.getter(name="resourceTypeList")
    def resource_type_list(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "resource_type_list")

    @resource_type_list.setter
    def resource_type_list(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "resource_type_list", value)

    @property
    @pulumi.getter(name="resourcesCleanUp")
    def resources_clean_up(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "resources_clean_up")

    @resources_clean_up.setter
    def resources_clean_up(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "resources_clean_up", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


class Policy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delete_all_policy_resources: Optional[pulumi.Input[bool]] = None,
                 exclude_map: Optional[pulumi.Input[pulumi.InputType['PolicyIeMapArgs']]] = None,
                 exclude_resource_tags: Optional[pulumi.Input[bool]] = None,
                 include_map: Optional[pulumi.Input[pulumi.InputType['PolicyIeMapArgs']]] = None,
                 policy_description: Optional[pulumi.Input[str]] = None,
                 policy_name: Optional[pulumi.Input[str]] = None,
                 remediation_enabled: Optional[pulumi.Input[bool]] = None,
                 resource_set_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyResourceTagArgs']]]]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 resource_type_list: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resources_clean_up: Optional[pulumi.Input[bool]] = None,
                 security_service_policy_data: Optional[pulumi.Input[pulumi.InputType['PolicySecurityServicePolicyDataArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None):
        """
        Creates an AWS Firewall Manager policy.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates an AWS Firewall Manager policy.

        :param str resource_name: The name of the resource.
        :param PolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delete_all_policy_resources: Optional[pulumi.Input[bool]] = None,
                 exclude_map: Optional[pulumi.Input[pulumi.InputType['PolicyIeMapArgs']]] = None,
                 exclude_resource_tags: Optional[pulumi.Input[bool]] = None,
                 include_map: Optional[pulumi.Input[pulumi.InputType['PolicyIeMapArgs']]] = None,
                 policy_description: Optional[pulumi.Input[str]] = None,
                 policy_name: Optional[pulumi.Input[str]] = None,
                 remediation_enabled: Optional[pulumi.Input[bool]] = None,
                 resource_set_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyResourceTagArgs']]]]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 resource_type_list: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resources_clean_up: Optional[pulumi.Input[bool]] = None,
                 security_service_policy_data: Optional[pulumi.Input[pulumi.InputType['PolicySecurityServicePolicyDataArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PolicyArgs.__new__(PolicyArgs)

            __props__.__dict__["delete_all_policy_resources"] = delete_all_policy_resources
            __props__.__dict__["exclude_map"] = exclude_map
            if exclude_resource_tags is None and not opts.urn:
                raise TypeError("Missing required property 'exclude_resource_tags'")
            __props__.__dict__["exclude_resource_tags"] = exclude_resource_tags
            __props__.__dict__["include_map"] = include_map
            __props__.__dict__["policy_description"] = policy_description
            __props__.__dict__["policy_name"] = policy_name
            if remediation_enabled is None and not opts.urn:
                raise TypeError("Missing required property 'remediation_enabled'")
            __props__.__dict__["remediation_enabled"] = remediation_enabled
            __props__.__dict__["resource_set_ids"] = resource_set_ids
            __props__.__dict__["resource_tags"] = resource_tags
            __props__.__dict__["resource_type"] = resource_type
            __props__.__dict__["resource_type_list"] = resource_type_list
            __props__.__dict__["resources_clean_up"] = resources_clean_up
            if security_service_policy_data is None and not opts.urn:
                raise TypeError("Missing required property 'security_service_policy_data'")
            __props__.__dict__["security_service_policy_data"] = security_service_policy_data
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["aws_id"] = None
        super(Policy, __self__).__init__(
            'aws-native:fms:Policy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Policy':
        """
        Get an existing Policy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PolicyArgs.__new__(PolicyArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["aws_id"] = None
        __props__.__dict__["delete_all_policy_resources"] = None
        __props__.__dict__["exclude_map"] = None
        __props__.__dict__["exclude_resource_tags"] = None
        __props__.__dict__["include_map"] = None
        __props__.__dict__["policy_description"] = None
        __props__.__dict__["policy_name"] = None
        __props__.__dict__["remediation_enabled"] = None
        __props__.__dict__["resource_set_ids"] = None
        __props__.__dict__["resource_tags"] = None
        __props__.__dict__["resource_type"] = None
        __props__.__dict__["resource_type_list"] = None
        __props__.__dict__["resources_clean_up"] = None
        __props__.__dict__["security_service_policy_data"] = None
        __props__.__dict__["tags"] = None
        return Policy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="awsId")
    def aws_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "aws_id")

    @property
    @pulumi.getter(name="deleteAllPolicyResources")
    def delete_all_policy_resources(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "delete_all_policy_resources")

    @property
    @pulumi.getter(name="excludeMap")
    def exclude_map(self) -> pulumi.Output[Optional['outputs.PolicyIeMap']]:
        return pulumi.get(self, "exclude_map")

    @property
    @pulumi.getter(name="excludeResourceTags")
    def exclude_resource_tags(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "exclude_resource_tags")

    @property
    @pulumi.getter(name="includeMap")
    def include_map(self) -> pulumi.Output[Optional['outputs.PolicyIeMap']]:
        return pulumi.get(self, "include_map")

    @property
    @pulumi.getter(name="policyDescription")
    def policy_description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "policy_description")

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "policy_name")

    @property
    @pulumi.getter(name="remediationEnabled")
    def remediation_enabled(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "remediation_enabled")

    @property
    @pulumi.getter(name="resourceSetIds")
    def resource_set_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "resource_set_ids")

    @property
    @pulumi.getter(name="resourceTags")
    def resource_tags(self) -> pulumi.Output[Optional[Sequence['outputs.PolicyResourceTag']]]:
        return pulumi.get(self, "resource_tags")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter(name="resourceTypeList")
    def resource_type_list(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "resource_type_list")

    @property
    @pulumi.getter(name="resourcesCleanUp")
    def resources_clean_up(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "resources_clean_up")

    @property
    @pulumi.getter(name="securityServicePolicyData")
    def security_service_policy_data(self) -> pulumi.Output['outputs.PolicySecurityServicePolicyData']:
        return pulumi.get(self, "security_service_policy_data")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        return pulumi.get(self, "tags")

