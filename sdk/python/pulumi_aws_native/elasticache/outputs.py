# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'UserAuthentication',
    'UserGroupReplicationGroupIdList',
    'UserGroupUserGroupPendingChanges',
    'UserGroupUserIdList',
    'UserPasswordList',
    'UserUserGroupIdList',
]

@pulumi.output_type
class UserAuthentication(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-user-authentication.html
    """
    def __init__(__self__, *,
                 password_count: Optional[int] = None,
                 type: Optional[str] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-user-authentication.html
        :param int password_count: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-user-authentication.html#cfn-elasticache-user-authentication-passwordcount
        :param str type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-user-authentication.html#cfn-elasticache-user-authentication-type
        """
        if password_count is not None:
            pulumi.set(__self__, "password_count", password_count)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="passwordCount")
    def password_count(self) -> Optional[int]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-user-authentication.html#cfn-elasticache-user-authentication-passwordcount
        """
        return pulumi.get(self, "password_count")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-user-authentication.html#cfn-elasticache-user-authentication-type
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class UserGroupReplicationGroupIdList(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-usergroup-replicationgroupidlist.html
    """
    def __init__(__self__, *,
                 replication_group_id_list: Optional[Sequence[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-usergroup-replicationgroupidlist.html
        :param Sequence[str] replication_group_id_list: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-usergroup-replicationgroupidlist.html#cfn-elasticache-usergroup-replicationgroupidlist-replicationgroupidlist
        """
        if replication_group_id_list is not None:
            pulumi.set(__self__, "replication_group_id_list", replication_group_id_list)

    @property
    @pulumi.getter(name="replicationGroupIdList")
    def replication_group_id_list(self) -> Optional[Sequence[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-usergroup-replicationgroupidlist.html#cfn-elasticache-usergroup-replicationgroupidlist-replicationgroupidlist
        """
        return pulumi.get(self, "replication_group_id_list")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class UserGroupUserGroupPendingChanges(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-usergroup-usergrouppendingchanges.html
    """
    def __init__(__self__, *,
                 user_ids_to_add: Optional[Sequence[str]] = None,
                 user_ids_to_remove: Optional[Sequence[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-usergroup-usergrouppendingchanges.html
        :param Sequence[str] user_ids_to_add: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-usergroup-usergrouppendingchanges.html#cfn-elasticache-usergroup-usergrouppendingchanges-useridstoadd
        :param Sequence[str] user_ids_to_remove: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-usergroup-usergrouppendingchanges.html#cfn-elasticache-usergroup-usergrouppendingchanges-useridstoremove
        """
        if user_ids_to_add is not None:
            pulumi.set(__self__, "user_ids_to_add", user_ids_to_add)
        if user_ids_to_remove is not None:
            pulumi.set(__self__, "user_ids_to_remove", user_ids_to_remove)

    @property
    @pulumi.getter(name="userIdsToAdd")
    def user_ids_to_add(self) -> Optional[Sequence[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-usergroup-usergrouppendingchanges.html#cfn-elasticache-usergroup-usergrouppendingchanges-useridstoadd
        """
        return pulumi.get(self, "user_ids_to_add")

    @property
    @pulumi.getter(name="userIdsToRemove")
    def user_ids_to_remove(self) -> Optional[Sequence[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-usergroup-usergrouppendingchanges.html#cfn-elasticache-usergroup-usergrouppendingchanges-useridstoremove
        """
        return pulumi.get(self, "user_ids_to_remove")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class UserGroupUserIdList(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-usergroup-useridlist.html
    """
    def __init__(__self__, *,
                 user_id_list: Optional[Sequence[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-usergroup-useridlist.html
        :param Sequence[str] user_id_list: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-usergroup-useridlist.html#cfn-elasticache-usergroup-useridlist-useridlist
        """
        if user_id_list is not None:
            pulumi.set(__self__, "user_id_list", user_id_list)

    @property
    @pulumi.getter(name="userIdList")
    def user_id_list(self) -> Optional[Sequence[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-usergroup-useridlist.html#cfn-elasticache-usergroup-useridlist-useridlist
        """
        return pulumi.get(self, "user_id_list")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class UserPasswordList(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-user-passwordlist.html
    """
    def __init__(__self__, *,
                 password_list: Optional[Sequence[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-user-passwordlist.html
        :param Sequence[str] password_list: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-user-passwordlist.html#cfn-elasticache-user-passwordlist-passwordlist
        """
        if password_list is not None:
            pulumi.set(__self__, "password_list", password_list)

    @property
    @pulumi.getter(name="passwordList")
    def password_list(self) -> Optional[Sequence[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-user-passwordlist.html#cfn-elasticache-user-passwordlist-passwordlist
        """
        return pulumi.get(self, "password_list")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class UserUserGroupIdList(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-user-usergroupidlist.html
    """
    def __init__(__self__, *,
                 user_group_id_list: Optional[Sequence[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-user-usergroupidlist.html
        :param Sequence[str] user_group_id_list: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-user-usergroupidlist.html#cfn-elasticache-user-usergroupidlist-usergroupidlist
        """
        if user_group_id_list is not None:
            pulumi.set(__self__, "user_group_id_list", user_group_id_list)

    @property
    @pulumi.getter(name="userGroupIdList")
    def user_group_id_list(self) -> Optional[Sequence[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-user-usergroupidlist.html#cfn-elasticache-user-usergroupidlist-usergroupidlist
        """
        return pulumi.get(self, "user_group_id_list")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

