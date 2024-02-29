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
from .. import outputs as _root_outputs

__all__ = [
    'GetRoleResult',
    'AwaitableGetRoleResult',
    'get_role',
    'get_role_output',
]

@pulumi.output_type
class GetRoleResult:
    def __init__(__self__, arn=None, assume_role_policy_document=None, description=None, managed_policy_arns=None, max_session_duration=None, permissions_boundary=None, policies=None, role_id=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if assume_role_policy_document and not isinstance(assume_role_policy_document, dict):
            raise TypeError("Expected argument 'assume_role_policy_document' to be a dict")
        pulumi.set(__self__, "assume_role_policy_document", assume_role_policy_document)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if managed_policy_arns and not isinstance(managed_policy_arns, list):
            raise TypeError("Expected argument 'managed_policy_arns' to be a list")
        pulumi.set(__self__, "managed_policy_arns", managed_policy_arns)
        if max_session_duration and not isinstance(max_session_duration, int):
            raise TypeError("Expected argument 'max_session_duration' to be a int")
        pulumi.set(__self__, "max_session_duration", max_session_duration)
        if permissions_boundary and not isinstance(permissions_boundary, str):
            raise TypeError("Expected argument 'permissions_boundary' to be a str")
        pulumi.set(__self__, "permissions_boundary", permissions_boundary)
        if policies and not isinstance(policies, list):
            raise TypeError("Expected argument 'policies' to be a list")
        pulumi.set(__self__, "policies", policies)
        if role_id and not isinstance(role_id, str):
            raise TypeError("Expected argument 'role_id' to be a str")
        pulumi.set(__self__, "role_id", role_id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="assumeRolePolicyDocument")
    def assume_role_policy_document(self) -> Optional[Any]:
        """
        The trust policy that is associated with this role. Trust policies define which entities can assume the role. You can associate only one trust policy with a role. For an example of a policy that can be used to assume a role, see [Template Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html#aws-resource-iam-role--examples). For more information about the elements that you can use in an IAM policy, see [Policy Elements Reference](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements.html) in the *User Guide*.

        Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::IAM::Role` for more information about the expected schema for this property.
        """
        return pulumi.get(self, "assume_role_policy_document")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        A description of the role that you provide.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="managedPolicyArns")
    def managed_policy_arns(self) -> Optional[Sequence[str]]:
        """
        A list of Amazon Resource Names (ARNs) of the IAM managed policies that you want to attach to the role.
         For more information about ARNs, see [Amazon Resource Names (ARNs) and Service Namespaces](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the *General Reference*.
        """
        return pulumi.get(self, "managed_policy_arns")

    @property
    @pulumi.getter(name="maxSessionDuration")
    def max_session_duration(self) -> Optional[int]:
        """
        The maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default value of one hour is applied. This setting can have a value from 1 hour to 12 hours.
         Anyone who assumes the role from the CLI or API can use the ``DurationSeconds`` API parameter or the ``duration-seconds`` CLI parameter to request a longer session. The ``MaxSessionDuration`` setting determines the maximum duration that can be requested using the ``DurationSeconds`` parameter. If users don't specify a value for the ``DurationSeconds`` parameter, their security credentials are valid for one hour by default. This applies when you use the ``AssumeRole*`` API operations or the ``assume-role*`` CLI operations but does not apply when you use those operations to create a console URL. For more information, see [Using IAM roles](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html) in the *IAM User Guide*.
        """
        return pulumi.get(self, "max_session_duration")

    @property
    @pulumi.getter(name="permissionsBoundary")
    def permissions_boundary(self) -> Optional[str]:
        """
        The ARN of the policy used to set the permissions boundary for the role.
         For more information about permissions boundaries, see [Permissions boundaries for IAM identities](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html) in the *IAM User Guide*.
        """
        return pulumi.get(self, "permissions_boundary")

    @property
    @pulumi.getter
    def policies(self) -> Optional[Sequence['outputs.RolePolicy']]:
        """
        Adds or updates an inline policy document that is embedded in the specified IAM role.
         When you embed an inline policy in a role, the inline policy is used as part of the role's access (permissions) policy. The role's trust policy is created at the same time as the role. You can update a role's trust policy later. For more information about IAM roles, go to [Using Roles to Delegate Permissions and Federate Identities](https://docs.aws.amazon.com/IAM/latest/UserGuide/roles-toplevel.html).
         A role can also have an attached managed policy. For information about policies, see [Managed Policies and Inline Policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html) in the *User Guide*.
         For information about limits on the number of inline policies that you can embed with a role, see [Limitations on Entities](https://docs.aws.amazon.com/IAM/latest/UserGuide/LimitationsOnEntities.html) in the *User Guide*.
          If an external policy (such as ``AWS::IAM::Policy`` or
        """
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> Optional[str]:
        return pulumi.get(self, "role_id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        A list of tags that are attached to the role. For more information about tagging, see [Tagging IAM resources](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the *IAM User Guide*.
        """
        return pulumi.get(self, "tags")


class AwaitableGetRoleResult(GetRoleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRoleResult(
            arn=self.arn,
            assume_role_policy_document=self.assume_role_policy_document,
            description=self.description,
            managed_policy_arns=self.managed_policy_arns,
            max_session_duration=self.max_session_duration,
            permissions_boundary=self.permissions_boundary,
            policies=self.policies,
            role_id=self.role_id,
            tags=self.tags)


def get_role(role_name: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRoleResult:
    """
    Creates a new role for your AWS-account.
      For more information about roles, see [IAM roles](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html) in the *IAM User Guide*. For information about quotas for role names and the number of roles you can create, see [IAM and quotas](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html) in the *IAM User Guide*.


    :param str role_name: A name for the IAM role, up to 64 characters in length. For valid values, see the ``RoleName`` parameter for the [CreateRole](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html) action in the *User Guide*.
            This parameter allows (per its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-. The role name must be unique within the account. Role names are not distinguished by case. For example, you cannot create roles named both "Role1" and "role1".
            If you don't specify a name, CFN generates a unique physical ID and uses that ID for the role name.
            If you specify a name, you must specify the ``CAPABILITY_NAMED_IAM`` value to acknowledge your template's capabilities. For more information, see [Acknowledging Resources in Templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/Use
    """
    __args__ = dict()
    __args__['roleName'] = role_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:iam:getRole', __args__, opts=opts, typ=GetRoleResult).value

    return AwaitableGetRoleResult(
        arn=pulumi.get(__ret__, 'arn'),
        assume_role_policy_document=pulumi.get(__ret__, 'assume_role_policy_document'),
        description=pulumi.get(__ret__, 'description'),
        managed_policy_arns=pulumi.get(__ret__, 'managed_policy_arns'),
        max_session_duration=pulumi.get(__ret__, 'max_session_duration'),
        permissions_boundary=pulumi.get(__ret__, 'permissions_boundary'),
        policies=pulumi.get(__ret__, 'policies'),
        role_id=pulumi.get(__ret__, 'role_id'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_role)
def get_role_output(role_name: Optional[pulumi.Input[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRoleResult]:
    """
    Creates a new role for your AWS-account.
      For more information about roles, see [IAM roles](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html) in the *IAM User Guide*. For information about quotas for role names and the number of roles you can create, see [IAM and quotas](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html) in the *IAM User Guide*.


    :param str role_name: A name for the IAM role, up to 64 characters in length. For valid values, see the ``RoleName`` parameter for the [CreateRole](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html) action in the *User Guide*.
            This parameter allows (per its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-. The role name must be unique within the account. Role names are not distinguished by case. For example, you cannot create roles named both "Role1" and "role1".
            If you don't specify a name, CFN generates a unique physical ID and uses that ID for the role name.
            If you specify a name, you must specify the ``CAPABILITY_NAMED_IAM`` value to acknowledge your template's capabilities. For more information, see [Acknowledging Resources in Templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/Use
    """
    ...
