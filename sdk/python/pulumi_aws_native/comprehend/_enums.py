# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'FlywheelDocumentClassificationConfigMode',
    'FlywheelModelType',
    'FlywheelTaskConfigLanguageCode',
]


class FlywheelDocumentClassificationConfigMode(str, Enum):
    MULTI_CLASS = "MULTI_CLASS"
    MULTI_LABEL = "MULTI_LABEL"


class FlywheelModelType(str, Enum):
    DOCUMENT_CLASSIFIER = "DOCUMENT_CLASSIFIER"
    ENTITY_RECOGNIZER = "ENTITY_RECOGNIZER"


class FlywheelTaskConfigLanguageCode(str, Enum):
    EN = "en"
    ES = "es"
    FR = "fr"
    IT = "it"
    DE = "de"
    PT = "pt"