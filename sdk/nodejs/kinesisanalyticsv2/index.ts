// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ApplicationArgs } from "./application";
export type Application = import("./application").Application;
export const Application: typeof import("./application").Application = null as any;
utilities.lazyLoad(exports, ["Application"], () => require("./application"));

export { ApplicationCloudWatchLoggingOptionArgs } from "./applicationCloudWatchLoggingOption";
export type ApplicationCloudWatchLoggingOption = import("./applicationCloudWatchLoggingOption").ApplicationCloudWatchLoggingOption;
export const ApplicationCloudWatchLoggingOption: typeof import("./applicationCloudWatchLoggingOption").ApplicationCloudWatchLoggingOption = null as any;
utilities.lazyLoad(exports, ["ApplicationCloudWatchLoggingOption"], () => require("./applicationCloudWatchLoggingOption"));

export { ApplicationOutputResourceArgs } from "./applicationOutputResource";
export type ApplicationOutputResource = import("./applicationOutputResource").ApplicationOutputResource;
export const ApplicationOutputResource: typeof import("./applicationOutputResource").ApplicationOutputResource = null as any;
utilities.lazyLoad(exports, ["ApplicationOutputResource"], () => require("./applicationOutputResource"));

export { ApplicationReferenceDataSourceArgs } from "./applicationReferenceDataSource";
export type ApplicationReferenceDataSource = import("./applicationReferenceDataSource").ApplicationReferenceDataSource;
export const ApplicationReferenceDataSource: typeof import("./applicationReferenceDataSource").ApplicationReferenceDataSource = null as any;
utilities.lazyLoad(exports, ["ApplicationReferenceDataSource"], () => require("./applicationReferenceDataSource"));

export { GetApplicationArgs, GetApplicationResult, GetApplicationOutputArgs } from "./getApplication";
export const getApplication: typeof import("./getApplication").getApplication = null as any;
export const getApplicationOutput: typeof import("./getApplication").getApplicationOutput = null as any;
utilities.lazyLoad(exports, ["getApplication","getApplicationOutput"], () => require("./getApplication"));

export { GetApplicationCloudWatchLoggingOptionArgs, GetApplicationCloudWatchLoggingOptionResult, GetApplicationCloudWatchLoggingOptionOutputArgs } from "./getApplicationCloudWatchLoggingOption";
export const getApplicationCloudWatchLoggingOption: typeof import("./getApplicationCloudWatchLoggingOption").getApplicationCloudWatchLoggingOption = null as any;
export const getApplicationCloudWatchLoggingOptionOutput: typeof import("./getApplicationCloudWatchLoggingOption").getApplicationCloudWatchLoggingOptionOutput = null as any;
utilities.lazyLoad(exports, ["getApplicationCloudWatchLoggingOption","getApplicationCloudWatchLoggingOptionOutput"], () => require("./getApplicationCloudWatchLoggingOption"));

export { GetApplicationOutputResourceArgs, GetApplicationOutputResourceResult, GetApplicationOutputResourceOutputArgs } from "./getApplicationOutputResource";
export const getApplicationOutputResource: typeof import("./getApplicationOutputResource").getApplicationOutputResource = null as any;
export const getApplicationOutputResourceOutput: typeof import("./getApplicationOutputResource").getApplicationOutputResourceOutput = null as any;
utilities.lazyLoad(exports, ["getApplicationOutputResource","getApplicationOutputResourceOutput"], () => require("./getApplicationOutputResource"));

export { GetApplicationReferenceDataSourceArgs, GetApplicationReferenceDataSourceResult, GetApplicationReferenceDataSourceOutputArgs } from "./getApplicationReferenceDataSource";
export const getApplicationReferenceDataSource: typeof import("./getApplicationReferenceDataSource").getApplicationReferenceDataSource = null as any;
export const getApplicationReferenceDataSourceOutput: typeof import("./getApplicationReferenceDataSource").getApplicationReferenceDataSourceOutput = null as any;
utilities.lazyLoad(exports, ["getApplicationReferenceDataSource","getApplicationReferenceDataSourceOutput"], () => require("./getApplicationReferenceDataSource"));


// Export enums:
export * from "../types/enums/kinesisanalyticsv2";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:kinesisanalyticsv2:Application":
                return new Application(name, <any>undefined, { urn })
            case "aws-native:kinesisanalyticsv2:ApplicationCloudWatchLoggingOption":
                return new ApplicationCloudWatchLoggingOption(name, <any>undefined, { urn })
            case "aws-native:kinesisanalyticsv2:ApplicationOutputResource":
                return new ApplicationOutputResource(name, <any>undefined, { urn })
            case "aws-native:kinesisanalyticsv2:ApplicationReferenceDataSource":
                return new ApplicationReferenceDataSource(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "kinesisanalyticsv2", _module)
