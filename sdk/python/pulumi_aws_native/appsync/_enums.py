# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'SourceApiAssociationConfigMergeType',
    'SourceApiAssociationStatus',
]


class SourceApiAssociationConfigMergeType(str, Enum):
    """
    Configuration of the merged behavior for the association. For example when it could be auto or has to be manual.
    """
    AUTO_MERGE = "AUTO_MERGE"
    MANUAL_MERGE = "MANUAL_MERGE"


class SourceApiAssociationStatus(str, Enum):
    """
    Current status of SourceApiAssociation.
    """
    MERGE_SCHEDULED = "MERGE_SCHEDULED"
    MERGE_FAILED = "MERGE_FAILED"
    MERGE_SUCCESS = "MERGE_SUCCESS"
    MERGE_IN_PROGRESS = "MERGE_IN_PROGRESS"
    AUTO_MERGE_SCHEDULE_FAILED = "AUTO_MERGE_SCHEDULE_FAILED"
    DELETION_SCHEDULED = "DELETION_SCHEDULED"
    DELETION_IN_PROGRESS = "DELETION_IN_PROGRESS"
    DELETION_FAILED = "DELETION_FAILED"
