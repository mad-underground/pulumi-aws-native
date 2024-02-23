# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AnnotationStoreAnnotationType',
    'AnnotationStoreEncryptionType',
    'AnnotationStoreSchemaValueType',
    'AnnotationStoreStoreFormat',
    'AnnotationStoreStoreStatus',
    'ReferenceStoreEncryptionType',
    'SequenceStoreEncryptionType',
    'VariantStoreEncryptionType',
    'VariantStoreStoreStatus',
    'WorkflowAccelerators',
    'WorkflowEngine',
    'WorkflowStatus',
    'WorkflowType',
]


class AnnotationStoreAnnotationType(str, Enum):
    GENERIC = "GENERIC"
    CHR_POS = "CHR_POS"
    CHR_POS_REF_ALT = "CHR_POS_REF_ALT"
    CHR_START_END_ONE_BASE = "CHR_START_END_ONE_BASE"
    CHR_START_END_REF_ALT_ONE_BASE = "CHR_START_END_REF_ALT_ONE_BASE"
    CHR_START_END_ZERO_BASE = "CHR_START_END_ZERO_BASE"
    CHR_START_END_REF_ALT_ZERO_BASE = "CHR_START_END_REF_ALT_ZERO_BASE"


class AnnotationStoreEncryptionType(str, Enum):
    KMS = "KMS"


class AnnotationStoreSchemaValueType(str, Enum):
    LONG = "LONG"
    INT = "INT"
    STRING = "STRING"
    FLOAT = "FLOAT"
    DOUBLE = "DOUBLE"
    BOOLEAN = "BOOLEAN"


class AnnotationStoreStoreFormat(str, Enum):
    GFF = "GFF"
    TSV = "TSV"
    VCF = "VCF"


class AnnotationStoreStoreStatus(str, Enum):
    CREATING = "CREATING"
    UPDATING = "UPDATING"
    DELETING = "DELETING"
    ACTIVE = "ACTIVE"
    FAILED = "FAILED"


class ReferenceStoreEncryptionType(str, Enum):
    KMS = "KMS"


class SequenceStoreEncryptionType(str, Enum):
    KMS = "KMS"


class VariantStoreEncryptionType(str, Enum):
    KMS = "KMS"


class VariantStoreStoreStatus(str, Enum):
    CREATING = "CREATING"
    UPDATING = "UPDATING"
    DELETING = "DELETING"
    ACTIVE = "ACTIVE"
    FAILED = "FAILED"


class WorkflowAccelerators(str, Enum):
    GPU = "GPU"


class WorkflowEngine(str, Enum):
    WDL = "WDL"
    NEXTFLOW = "NEXTFLOW"
    CWL = "CWL"


class WorkflowStatus(str, Enum):
    CREATING = "CREATING"
    ACTIVE = "ACTIVE"
    UPDATING = "UPDATING"
    DELETED = "DELETED"
    FAILED = "FAILED"


class WorkflowType(str, Enum):
    PRIVATE = "PRIVATE"
