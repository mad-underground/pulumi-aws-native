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
from ._inputs import *

__all__ = ['DataLakeSettingsArgs', 'DataLakeSettings']

@pulumi.input_type
class DataLakeSettingsArgs:
    def __init__(__self__, *,
                 admins: Optional[pulumi.Input['DataLakeSettingsAdminsArgs']] = None,
                 allow_external_data_filtering: Optional[pulumi.Input[bool]] = None,
                 allow_full_table_external_data_access: Optional[pulumi.Input[bool]] = None,
                 authorized_session_tag_value_list: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 create_database_default_permissions: Optional[pulumi.Input['DataLakeSettingsCreateDatabaseDefaultPermissionsArgs']] = None,
                 create_table_default_permissions: Optional[pulumi.Input['DataLakeSettingsCreateTableDefaultPermissionsArgs']] = None,
                 external_data_filtering_allow_list: Optional[pulumi.Input['DataLakeSettingsExternalDataFilteringAllowListArgs']] = None,
                 mutation_type: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[Any] = None,
                 trusted_resource_owners: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a DataLakeSettings resource.
        :param Any parameters: Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::LakeFormation::DataLakeSettings` for more information about the expected schema for this property.
        """
        if admins is not None:
            pulumi.set(__self__, "admins", admins)
        if allow_external_data_filtering is not None:
            pulumi.set(__self__, "allow_external_data_filtering", allow_external_data_filtering)
        if allow_full_table_external_data_access is not None:
            pulumi.set(__self__, "allow_full_table_external_data_access", allow_full_table_external_data_access)
        if authorized_session_tag_value_list is not None:
            pulumi.set(__self__, "authorized_session_tag_value_list", authorized_session_tag_value_list)
        if create_database_default_permissions is not None:
            pulumi.set(__self__, "create_database_default_permissions", create_database_default_permissions)
        if create_table_default_permissions is not None:
            pulumi.set(__self__, "create_table_default_permissions", create_table_default_permissions)
        if external_data_filtering_allow_list is not None:
            pulumi.set(__self__, "external_data_filtering_allow_list", external_data_filtering_allow_list)
        if mutation_type is not None:
            pulumi.set(__self__, "mutation_type", mutation_type)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if trusted_resource_owners is not None:
            pulumi.set(__self__, "trusted_resource_owners", trusted_resource_owners)

    @property
    @pulumi.getter
    def admins(self) -> Optional[pulumi.Input['DataLakeSettingsAdminsArgs']]:
        return pulumi.get(self, "admins")

    @admins.setter
    def admins(self, value: Optional[pulumi.Input['DataLakeSettingsAdminsArgs']]):
        pulumi.set(self, "admins", value)

    @property
    @pulumi.getter(name="allowExternalDataFiltering")
    def allow_external_data_filtering(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "allow_external_data_filtering")

    @allow_external_data_filtering.setter
    def allow_external_data_filtering(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_external_data_filtering", value)

    @property
    @pulumi.getter(name="allowFullTableExternalDataAccess")
    def allow_full_table_external_data_access(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "allow_full_table_external_data_access")

    @allow_full_table_external_data_access.setter
    def allow_full_table_external_data_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_full_table_external_data_access", value)

    @property
    @pulumi.getter(name="authorizedSessionTagValueList")
    def authorized_session_tag_value_list(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "authorized_session_tag_value_list")

    @authorized_session_tag_value_list.setter
    def authorized_session_tag_value_list(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "authorized_session_tag_value_list", value)

    @property
    @pulumi.getter(name="createDatabaseDefaultPermissions")
    def create_database_default_permissions(self) -> Optional[pulumi.Input['DataLakeSettingsCreateDatabaseDefaultPermissionsArgs']]:
        return pulumi.get(self, "create_database_default_permissions")

    @create_database_default_permissions.setter
    def create_database_default_permissions(self, value: Optional[pulumi.Input['DataLakeSettingsCreateDatabaseDefaultPermissionsArgs']]):
        pulumi.set(self, "create_database_default_permissions", value)

    @property
    @pulumi.getter(name="createTableDefaultPermissions")
    def create_table_default_permissions(self) -> Optional[pulumi.Input['DataLakeSettingsCreateTableDefaultPermissionsArgs']]:
        return pulumi.get(self, "create_table_default_permissions")

    @create_table_default_permissions.setter
    def create_table_default_permissions(self, value: Optional[pulumi.Input['DataLakeSettingsCreateTableDefaultPermissionsArgs']]):
        pulumi.set(self, "create_table_default_permissions", value)

    @property
    @pulumi.getter(name="externalDataFilteringAllowList")
    def external_data_filtering_allow_list(self) -> Optional[pulumi.Input['DataLakeSettingsExternalDataFilteringAllowListArgs']]:
        return pulumi.get(self, "external_data_filtering_allow_list")

    @external_data_filtering_allow_list.setter
    def external_data_filtering_allow_list(self, value: Optional[pulumi.Input['DataLakeSettingsExternalDataFilteringAllowListArgs']]):
        pulumi.set(self, "external_data_filtering_allow_list", value)

    @property
    @pulumi.getter(name="mutationType")
    def mutation_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mutation_type")

    @mutation_type.setter
    def mutation_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mutation_type", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[Any]:
        """
        Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::LakeFormation::DataLakeSettings` for more information about the expected schema for this property.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[Any]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter(name="trustedResourceOwners")
    def trusted_resource_owners(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "trusted_resource_owners")

    @trusted_resource_owners.setter
    def trusted_resource_owners(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "trusted_resource_owners", value)


warnings.warn("""DataLakeSettings is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class DataLakeSettings(pulumi.CustomResource):
    warnings.warn("""DataLakeSettings is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admins: Optional[pulumi.Input[pulumi.InputType['DataLakeSettingsAdminsArgs']]] = None,
                 allow_external_data_filtering: Optional[pulumi.Input[bool]] = None,
                 allow_full_table_external_data_access: Optional[pulumi.Input[bool]] = None,
                 authorized_session_tag_value_list: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 create_database_default_permissions: Optional[pulumi.Input[pulumi.InputType['DataLakeSettingsCreateDatabaseDefaultPermissionsArgs']]] = None,
                 create_table_default_permissions: Optional[pulumi.Input[pulumi.InputType['DataLakeSettingsCreateTableDefaultPermissionsArgs']]] = None,
                 external_data_filtering_allow_list: Optional[pulumi.Input[pulumi.InputType['DataLakeSettingsExternalDataFilteringAllowListArgs']]] = None,
                 mutation_type: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[Any] = None,
                 trusted_resource_owners: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::LakeFormation::DataLakeSettings

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param Any parameters: Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::LakeFormation::DataLakeSettings` for more information about the expected schema for this property.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[DataLakeSettingsArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::LakeFormation::DataLakeSettings

        :param str resource_name: The name of the resource.
        :param DataLakeSettingsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DataLakeSettingsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admins: Optional[pulumi.Input[pulumi.InputType['DataLakeSettingsAdminsArgs']]] = None,
                 allow_external_data_filtering: Optional[pulumi.Input[bool]] = None,
                 allow_full_table_external_data_access: Optional[pulumi.Input[bool]] = None,
                 authorized_session_tag_value_list: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 create_database_default_permissions: Optional[pulumi.Input[pulumi.InputType['DataLakeSettingsCreateDatabaseDefaultPermissionsArgs']]] = None,
                 create_table_default_permissions: Optional[pulumi.Input[pulumi.InputType['DataLakeSettingsCreateTableDefaultPermissionsArgs']]] = None,
                 external_data_filtering_allow_list: Optional[pulumi.Input[pulumi.InputType['DataLakeSettingsExternalDataFilteringAllowListArgs']]] = None,
                 mutation_type: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[Any] = None,
                 trusted_resource_owners: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        pulumi.log.warn("""DataLakeSettings is deprecated: DataLakeSettings is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DataLakeSettingsArgs.__new__(DataLakeSettingsArgs)

            __props__.__dict__["admins"] = admins
            __props__.__dict__["allow_external_data_filtering"] = allow_external_data_filtering
            __props__.__dict__["allow_full_table_external_data_access"] = allow_full_table_external_data_access
            __props__.__dict__["authorized_session_tag_value_list"] = authorized_session_tag_value_list
            __props__.__dict__["create_database_default_permissions"] = create_database_default_permissions
            __props__.__dict__["create_table_default_permissions"] = create_table_default_permissions
            __props__.__dict__["external_data_filtering_allow_list"] = external_data_filtering_allow_list
            __props__.__dict__["mutation_type"] = mutation_type
            __props__.__dict__["parameters"] = parameters
            __props__.__dict__["trusted_resource_owners"] = trusted_resource_owners
            __props__.__dict__["aws_id"] = None
        super(DataLakeSettings, __self__).__init__(
            'aws-native:lakeformation:DataLakeSettings',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DataLakeSettings':
        """
        Get an existing DataLakeSettings resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DataLakeSettingsArgs.__new__(DataLakeSettingsArgs)

        __props__.__dict__["admins"] = None
        __props__.__dict__["allow_external_data_filtering"] = None
        __props__.__dict__["allow_full_table_external_data_access"] = None
        __props__.__dict__["authorized_session_tag_value_list"] = None
        __props__.__dict__["aws_id"] = None
        __props__.__dict__["create_database_default_permissions"] = None
        __props__.__dict__["create_table_default_permissions"] = None
        __props__.__dict__["external_data_filtering_allow_list"] = None
        __props__.__dict__["mutation_type"] = None
        __props__.__dict__["parameters"] = None
        __props__.__dict__["trusted_resource_owners"] = None
        return DataLakeSettings(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def admins(self) -> pulumi.Output[Optional['outputs.DataLakeSettingsAdmins']]:
        return pulumi.get(self, "admins")

    @property
    @pulumi.getter(name="allowExternalDataFiltering")
    def allow_external_data_filtering(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "allow_external_data_filtering")

    @property
    @pulumi.getter(name="allowFullTableExternalDataAccess")
    def allow_full_table_external_data_access(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "allow_full_table_external_data_access")

    @property
    @pulumi.getter(name="authorizedSessionTagValueList")
    def authorized_session_tag_value_list(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "authorized_session_tag_value_list")

    @property
    @pulumi.getter(name="awsId")
    def aws_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "aws_id")

    @property
    @pulumi.getter(name="createDatabaseDefaultPermissions")
    def create_database_default_permissions(self) -> pulumi.Output[Optional['outputs.DataLakeSettingsCreateDatabaseDefaultPermissions']]:
        return pulumi.get(self, "create_database_default_permissions")

    @property
    @pulumi.getter(name="createTableDefaultPermissions")
    def create_table_default_permissions(self) -> pulumi.Output[Optional['outputs.DataLakeSettingsCreateTableDefaultPermissions']]:
        return pulumi.get(self, "create_table_default_permissions")

    @property
    @pulumi.getter(name="externalDataFilteringAllowList")
    def external_data_filtering_allow_list(self) -> pulumi.Output[Optional['outputs.DataLakeSettingsExternalDataFilteringAllowList']]:
        return pulumi.get(self, "external_data_filtering_allow_list")

    @property
    @pulumi.getter(name="mutationType")
    def mutation_type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "mutation_type")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[Any]]:
        """
        Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::LakeFormation::DataLakeSettings` for more information about the expected schema for this property.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="trustedResourceOwners")
    def trusted_resource_owners(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "trusted_resource_owners")

