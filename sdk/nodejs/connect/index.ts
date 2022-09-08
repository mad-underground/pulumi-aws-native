// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ContactFlowArgs } from "./contactFlow";
export type ContactFlow = import("./contactFlow").ContactFlow;
export const ContactFlow: typeof import("./contactFlow").ContactFlow = null as any;

export { ContactFlowModuleArgs } from "./contactFlowModule";
export type ContactFlowModule = import("./contactFlowModule").ContactFlowModule;
export const ContactFlowModule: typeof import("./contactFlowModule").ContactFlowModule = null as any;

export { GetContactFlowArgs, GetContactFlowResult, GetContactFlowOutputArgs } from "./getContactFlow";
export const getContactFlow: typeof import("./getContactFlow").getContactFlow = null as any;
export const getContactFlowOutput: typeof import("./getContactFlow").getContactFlowOutput = null as any;

export { GetContactFlowModuleArgs, GetContactFlowModuleResult, GetContactFlowModuleOutputArgs } from "./getContactFlowModule";
export const getContactFlowModule: typeof import("./getContactFlowModule").getContactFlowModule = null as any;
export const getContactFlowModuleOutput: typeof import("./getContactFlowModule").getContactFlowModuleOutput = null as any;

export { GetHoursOfOperationArgs, GetHoursOfOperationResult, GetHoursOfOperationOutputArgs } from "./getHoursOfOperation";
export const getHoursOfOperation: typeof import("./getHoursOfOperation").getHoursOfOperation = null as any;
export const getHoursOfOperationOutput: typeof import("./getHoursOfOperation").getHoursOfOperationOutput = null as any;

export { GetInstanceArgs, GetInstanceResult, GetInstanceOutputArgs } from "./getInstance";
export const getInstance: typeof import("./getInstance").getInstance = null as any;
export const getInstanceOutput: typeof import("./getInstance").getInstanceOutput = null as any;

export { GetInstanceStorageConfigArgs, GetInstanceStorageConfigResult, GetInstanceStorageConfigOutputArgs } from "./getInstanceStorageConfig";
export const getInstanceStorageConfig: typeof import("./getInstanceStorageConfig").getInstanceStorageConfig = null as any;
export const getInstanceStorageConfigOutput: typeof import("./getInstanceStorageConfig").getInstanceStorageConfigOutput = null as any;

export { GetPhoneNumberArgs, GetPhoneNumberResult, GetPhoneNumberOutputArgs } from "./getPhoneNumber";
export const getPhoneNumber: typeof import("./getPhoneNumber").getPhoneNumber = null as any;
export const getPhoneNumberOutput: typeof import("./getPhoneNumber").getPhoneNumberOutput = null as any;

export { GetQuickConnectArgs, GetQuickConnectResult, GetQuickConnectOutputArgs } from "./getQuickConnect";
export const getQuickConnect: typeof import("./getQuickConnect").getQuickConnect = null as any;
export const getQuickConnectOutput: typeof import("./getQuickConnect").getQuickConnectOutput = null as any;

export { GetTaskTemplateArgs, GetTaskTemplateResult, GetTaskTemplateOutputArgs } from "./getTaskTemplate";
export const getTaskTemplate: typeof import("./getTaskTemplate").getTaskTemplate = null as any;
export const getTaskTemplateOutput: typeof import("./getTaskTemplate").getTaskTemplateOutput = null as any;

export { GetUserArgs, GetUserResult, GetUserOutputArgs } from "./getUser";
export const getUser: typeof import("./getUser").getUser = null as any;
export const getUserOutput: typeof import("./getUser").getUserOutput = null as any;

export { GetUserHierarchyGroupArgs, GetUserHierarchyGroupResult, GetUserHierarchyGroupOutputArgs } from "./getUserHierarchyGroup";
export const getUserHierarchyGroup: typeof import("./getUserHierarchyGroup").getUserHierarchyGroup = null as any;
export const getUserHierarchyGroupOutput: typeof import("./getUserHierarchyGroup").getUserHierarchyGroupOutput = null as any;

export { HoursOfOperationArgs } from "./hoursOfOperation";
export type HoursOfOperation = import("./hoursOfOperation").HoursOfOperation;
export const HoursOfOperation: typeof import("./hoursOfOperation").HoursOfOperation = null as any;

export { InstanceArgs } from "./instance";
export type Instance = import("./instance").Instance;
export const Instance: typeof import("./instance").Instance = null as any;

export { InstanceStorageConfigArgs } from "./instanceStorageConfig";
export type InstanceStorageConfig = import("./instanceStorageConfig").InstanceStorageConfig;
export const InstanceStorageConfig: typeof import("./instanceStorageConfig").InstanceStorageConfig = null as any;

export { PhoneNumberArgs } from "./phoneNumber";
export type PhoneNumber = import("./phoneNumber").PhoneNumber;
export const PhoneNumber: typeof import("./phoneNumber").PhoneNumber = null as any;

export { QuickConnectArgs } from "./quickConnect";
export type QuickConnect = import("./quickConnect").QuickConnect;
export const QuickConnect: typeof import("./quickConnect").QuickConnect = null as any;

export { TaskTemplateArgs } from "./taskTemplate";
export type TaskTemplate = import("./taskTemplate").TaskTemplate;
export const TaskTemplate: typeof import("./taskTemplate").TaskTemplate = null as any;

export { UserArgs } from "./user";
export type User = import("./user").User;
export const User: typeof import("./user").User = null as any;

export { UserHierarchyGroupArgs } from "./userHierarchyGroup";
export type UserHierarchyGroup = import("./userHierarchyGroup").UserHierarchyGroup;
export const UserHierarchyGroup: typeof import("./userHierarchyGroup").UserHierarchyGroup = null as any;

utilities.lazyLoad(exports, ["ContactFlow"], () => require("./contactFlow"));
utilities.lazyLoad(exports, ["ContactFlowModule"], () => require("./contactFlowModule"));
utilities.lazyLoad(exports, ["getContactFlow","getContactFlowOutput"], () => require("./getContactFlow"));
utilities.lazyLoad(exports, ["getContactFlowModule","getContactFlowModuleOutput"], () => require("./getContactFlowModule"));
utilities.lazyLoad(exports, ["getHoursOfOperation","getHoursOfOperationOutput"], () => require("./getHoursOfOperation"));
utilities.lazyLoad(exports, ["getInstance","getInstanceOutput"], () => require("./getInstance"));
utilities.lazyLoad(exports, ["getInstanceStorageConfig","getInstanceStorageConfigOutput"], () => require("./getInstanceStorageConfig"));
utilities.lazyLoad(exports, ["getPhoneNumber","getPhoneNumberOutput"], () => require("./getPhoneNumber"));
utilities.lazyLoad(exports, ["getQuickConnect","getQuickConnectOutput"], () => require("./getQuickConnect"));
utilities.lazyLoad(exports, ["getTaskTemplate","getTaskTemplateOutput"], () => require("./getTaskTemplate"));
utilities.lazyLoad(exports, ["getUser","getUserOutput"], () => require("./getUser"));
utilities.lazyLoad(exports, ["getUserHierarchyGroup","getUserHierarchyGroupOutput"], () => require("./getUserHierarchyGroup"));
utilities.lazyLoad(exports, ["HoursOfOperation"], () => require("./hoursOfOperation"));
utilities.lazyLoad(exports, ["Instance"], () => require("./instance"));
utilities.lazyLoad(exports, ["InstanceStorageConfig"], () => require("./instanceStorageConfig"));
utilities.lazyLoad(exports, ["PhoneNumber"], () => require("./phoneNumber"));
utilities.lazyLoad(exports, ["QuickConnect"], () => require("./quickConnect"));
utilities.lazyLoad(exports, ["TaskTemplate"], () => require("./taskTemplate"));
utilities.lazyLoad(exports, ["User"], () => require("./user"));
utilities.lazyLoad(exports, ["UserHierarchyGroup"], () => require("./userHierarchyGroup"));

// Export enums:
export * from "../types/enums/connect";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:connect:ContactFlow":
                return new ContactFlow(name, <any>undefined, { urn })
            case "aws-native:connect:ContactFlowModule":
                return new ContactFlowModule(name, <any>undefined, { urn })
            case "aws-native:connect:HoursOfOperation":
                return new HoursOfOperation(name, <any>undefined, { urn })
            case "aws-native:connect:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "aws-native:connect:InstanceStorageConfig":
                return new InstanceStorageConfig(name, <any>undefined, { urn })
            case "aws-native:connect:PhoneNumber":
                return new PhoneNumber(name, <any>undefined, { urn })
            case "aws-native:connect:QuickConnect":
                return new QuickConnect(name, <any>undefined, { urn })
            case "aws-native:connect:TaskTemplate":
                return new TaskTemplate(name, <any>undefined, { urn })
            case "aws-native:connect:User":
                return new User(name, <any>undefined, { urn })
            case "aws-native:connect:UserHierarchyGroup":
                return new UserHierarchyGroup(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "connect", _module)
