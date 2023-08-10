# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'DataSourceConditionOperator',
    'DataSourceConfluenceAttachmentFieldName',
    'DataSourceConfluenceBlogFieldName',
    'DataSourceConfluencePageFieldName',
    'DataSourceConfluenceSpaceFieldName',
    'DataSourceConfluenceVersion',
    'DataSourceDatabaseEngineType',
    'DataSourceQueryIdentifiersEnclosingOption',
    'DataSourceSalesforceChatterFeedIncludeFilterType',
    'DataSourceSalesforceKnowledgeArticleState',
    'DataSourceSalesforceStandardObjectName',
    'DataSourceServiceNowAuthenticationType',
    'DataSourceServiceNowBuildVersionType',
    'DataSourceSharePointConfigurationSharePointVersion',
    'DataSourceType',
    'DataSourceWebCrawlerSeedUrlConfigurationWebCrawlerMode',
    'FaqFileFormat',
    'IndexDocumentAttributeValueType',
    'IndexEdition',
    'IndexKeyLocation',
    'IndexOrder',
    'IndexUserContextPolicy',
]


class DataSourceConditionOperator(str, Enum):
    GREATER_THAN = "GreaterThan"
    GREATER_THAN_OR_EQUALS = "GreaterThanOrEquals"
    LESS_THAN = "LessThan"
    LESS_THAN_OR_EQUALS = "LessThanOrEquals"
    EQUALS = "Equals"
    NOT_EQUALS = "NotEquals"
    CONTAINS = "Contains"
    NOT_CONTAINS = "NotContains"
    EXISTS = "Exists"
    NOT_EXISTS = "NotExists"
    BEGINS_WITH = "BeginsWith"


class DataSourceConfluenceAttachmentFieldName(str, Enum):
    AUTHOR = "AUTHOR"
    CONTENT_TYPE = "CONTENT_TYPE"
    CREATED_DATE = "CREATED_DATE"
    DISPLAY_URL = "DISPLAY_URL"
    FILE_SIZE = "FILE_SIZE"
    ITEM_TYPE = "ITEM_TYPE"
    PARENT_ID = "PARENT_ID"
    SPACE_KEY = "SPACE_KEY"
    SPACE_NAME = "SPACE_NAME"
    URL = "URL"
    VERSION = "VERSION"


class DataSourceConfluenceBlogFieldName(str, Enum):
    AUTHOR = "AUTHOR"
    DISPLAY_URL = "DISPLAY_URL"
    ITEM_TYPE = "ITEM_TYPE"
    LABELS = "LABELS"
    PUBLISH_DATE = "PUBLISH_DATE"
    SPACE_KEY = "SPACE_KEY"
    SPACE_NAME = "SPACE_NAME"
    URL = "URL"
    VERSION = "VERSION"


class DataSourceConfluencePageFieldName(str, Enum):
    AUTHOR = "AUTHOR"
    CONTENT_STATUS = "CONTENT_STATUS"
    CREATED_DATE = "CREATED_DATE"
    DISPLAY_URL = "DISPLAY_URL"
    ITEM_TYPE = "ITEM_TYPE"
    LABELS = "LABELS"
    MODIFIED_DATE = "MODIFIED_DATE"
    PARENT_ID = "PARENT_ID"
    SPACE_KEY = "SPACE_KEY"
    SPACE_NAME = "SPACE_NAME"
    URL = "URL"
    VERSION = "VERSION"


class DataSourceConfluenceSpaceFieldName(str, Enum):
    DISPLAY_URL = "DISPLAY_URL"
    ITEM_TYPE = "ITEM_TYPE"
    SPACE_KEY = "SPACE_KEY"
    URL = "URL"


class DataSourceConfluenceVersion(str, Enum):
    CLOUD = "CLOUD"
    SERVER = "SERVER"


class DataSourceDatabaseEngineType(str, Enum):
    RDS_AURORA_MYSQL = "RDS_AURORA_MYSQL"
    RDS_AURORA_POSTGRESQL = "RDS_AURORA_POSTGRESQL"
    RDS_MYSQL = "RDS_MYSQL"
    RDS_POSTGRESQL = "RDS_POSTGRESQL"


class DataSourceQueryIdentifiersEnclosingOption(str, Enum):
    DOUBLE_QUOTES = "DOUBLE_QUOTES"
    NONE = "NONE"


class DataSourceSalesforceChatterFeedIncludeFilterType(str, Enum):
    ACTIVE_USER = "ACTIVE_USER"
    STANDARD_USER = "STANDARD_USER"


class DataSourceSalesforceKnowledgeArticleState(str, Enum):
    DRAFT = "DRAFT"
    PUBLISHED = "PUBLISHED"
    ARCHIVED = "ARCHIVED"


class DataSourceSalesforceStandardObjectName(str, Enum):
    ACCOUNT = "ACCOUNT"
    CAMPAIGN = "CAMPAIGN"
    CASE = "CASE"
    CONTACT = "CONTACT"
    CONTRACT = "CONTRACT"
    DOCUMENT = "DOCUMENT"
    GROUP = "GROUP"
    IDEA = "IDEA"
    LEAD = "LEAD"
    OPPORTUNITY = "OPPORTUNITY"
    PARTNER = "PARTNER"
    PRICEBOOK = "PRICEBOOK"
    PRODUCT = "PRODUCT"
    PROFILE = "PROFILE"
    SOLUTION = "SOLUTION"
    TASK = "TASK"
    USER = "USER"


class DataSourceServiceNowAuthenticationType(str, Enum):
    HTTP_BASIC = "HTTP_BASIC"
    OAUTH2 = "OAUTH2"


class DataSourceServiceNowBuildVersionType(str, Enum):
    LONDON = "LONDON"
    OTHERS = "OTHERS"


class DataSourceSharePointConfigurationSharePointVersion(str, Enum):
    SHAREPOINT_ONLINE = "SHAREPOINT_ONLINE"
    SHAREPOINT2013 = "SHAREPOINT_2013"
    SHAREPOINT2016 = "SHAREPOINT_2016"


class DataSourceType(str, Enum):
    """
    Data source type
    """
    S3 = "S3"
    SHAREPOINT = "SHAREPOINT"
    SALESFORCE = "SALESFORCE"
    ONEDRIVE = "ONEDRIVE"
    SERVICENOW = "SERVICENOW"
    DATABASE = "DATABASE"
    CUSTOM = "CUSTOM"
    CONFLUENCE = "CONFLUENCE"
    GOOGLEDRIVE = "GOOGLEDRIVE"
    WEBCRAWLER = "WEBCRAWLER"
    WORKDOCS = "WORKDOCS"
    TEMPLATE = "TEMPLATE"


class DataSourceWebCrawlerSeedUrlConfigurationWebCrawlerMode(str, Enum):
    HOST_ONLY = "HOST_ONLY"
    SUBDOMAINS = "SUBDOMAINS"
    EVERYTHING = "EVERYTHING"


class FaqFileFormat(str, Enum):
    """
    Format of the input file
    """
    CSV = "CSV"
    CSV_WITH_HEADER = "CSV_WITH_HEADER"
    JSON = "JSON"


class IndexDocumentAttributeValueType(str, Enum):
    STRING_VALUE = "STRING_VALUE"
    STRING_LIST_VALUE = "STRING_LIST_VALUE"
    LONG_VALUE = "LONG_VALUE"
    DATE_VALUE = "DATE_VALUE"


class IndexEdition(str, Enum):
    """
    Edition of index
    """
    DEVELOPER_EDITION = "DEVELOPER_EDITION"
    ENTERPRISE_EDITION = "ENTERPRISE_EDITION"


class IndexKeyLocation(str, Enum):
    URL = "URL"
    SECRET_MANAGER = "SECRET_MANAGER"


class IndexOrder(str, Enum):
    ASCENDING = "ASCENDING"
    DESCENDING = "DESCENDING"


class IndexUserContextPolicy(str, Enum):
    ATTRIBUTE_FILTER = "ATTRIBUTE_FILTER"
    USER_TOKEN = "USER_TOKEN"
