// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AppBlockArgs } from "./appBlock";
export type AppBlock = import("./appBlock").AppBlock;
export const AppBlock: typeof import("./appBlock").AppBlock = null as any;

export { ApplicationArgs } from "./application";
export type Application = import("./application").Application;
export const Application: typeof import("./application").Application = null as any;

export { ApplicationEntitlementAssociationArgs } from "./applicationEntitlementAssociation";
export type ApplicationEntitlementAssociation = import("./applicationEntitlementAssociation").ApplicationEntitlementAssociation;
export const ApplicationEntitlementAssociation: typeof import("./applicationEntitlementAssociation").ApplicationEntitlementAssociation = null as any;

export { ApplicationFleetAssociationArgs } from "./applicationFleetAssociation";
export type ApplicationFleetAssociation = import("./applicationFleetAssociation").ApplicationFleetAssociation;
export const ApplicationFleetAssociation: typeof import("./applicationFleetAssociation").ApplicationFleetAssociation = null as any;

export { DirectoryConfigArgs } from "./directoryConfig";
export type DirectoryConfig = import("./directoryConfig").DirectoryConfig;
export const DirectoryConfig: typeof import("./directoryConfig").DirectoryConfig = null as any;

export { EntitlementArgs } from "./entitlement";
export type Entitlement = import("./entitlement").Entitlement;
export const Entitlement: typeof import("./entitlement").Entitlement = null as any;

export { FleetArgs } from "./fleet";
export type Fleet = import("./fleet").Fleet;
export const Fleet: typeof import("./fleet").Fleet = null as any;

export { GetAppBlockArgs, GetAppBlockResult, GetAppBlockOutputArgs } from "./getAppBlock";
export const getAppBlock: typeof import("./getAppBlock").getAppBlock = null as any;
export const getAppBlockOutput: typeof import("./getAppBlock").getAppBlockOutput = null as any;

export { GetApplicationArgs, GetApplicationResult, GetApplicationOutputArgs } from "./getApplication";
export const getApplication: typeof import("./getApplication").getApplication = null as any;
export const getApplicationOutput: typeof import("./getApplication").getApplicationOutput = null as any;

export { GetDirectoryConfigArgs, GetDirectoryConfigResult, GetDirectoryConfigOutputArgs } from "./getDirectoryConfig";
export const getDirectoryConfig: typeof import("./getDirectoryConfig").getDirectoryConfig = null as any;
export const getDirectoryConfigOutput: typeof import("./getDirectoryConfig").getDirectoryConfigOutput = null as any;

export { GetEntitlementArgs, GetEntitlementResult, GetEntitlementOutputArgs } from "./getEntitlement";
export const getEntitlement: typeof import("./getEntitlement").getEntitlement = null as any;
export const getEntitlementOutput: typeof import("./getEntitlement").getEntitlementOutput = null as any;

export { GetFleetArgs, GetFleetResult, GetFleetOutputArgs } from "./getFleet";
export const getFleet: typeof import("./getFleet").getFleet = null as any;
export const getFleetOutput: typeof import("./getFleet").getFleetOutput = null as any;

export { GetImageBuilderArgs, GetImageBuilderResult, GetImageBuilderOutputArgs } from "./getImageBuilder";
export const getImageBuilder: typeof import("./getImageBuilder").getImageBuilder = null as any;
export const getImageBuilderOutput: typeof import("./getImageBuilder").getImageBuilderOutput = null as any;

export { GetStackArgs, GetStackResult, GetStackOutputArgs } from "./getStack";
export const getStack: typeof import("./getStack").getStack = null as any;
export const getStackOutput: typeof import("./getStack").getStackOutput = null as any;

export { GetStackFleetAssociationArgs, GetStackFleetAssociationResult, GetStackFleetAssociationOutputArgs } from "./getStackFleetAssociation";
export const getStackFleetAssociation: typeof import("./getStackFleetAssociation").getStackFleetAssociation = null as any;
export const getStackFleetAssociationOutput: typeof import("./getStackFleetAssociation").getStackFleetAssociationOutput = null as any;

export { GetStackUserAssociationArgs, GetStackUserAssociationResult, GetStackUserAssociationOutputArgs } from "./getStackUserAssociation";
export const getStackUserAssociation: typeof import("./getStackUserAssociation").getStackUserAssociation = null as any;
export const getStackUserAssociationOutput: typeof import("./getStackUserAssociation").getStackUserAssociationOutput = null as any;

export { GetUserArgs, GetUserResult, GetUserOutputArgs } from "./getUser";
export const getUser: typeof import("./getUser").getUser = null as any;
export const getUserOutput: typeof import("./getUser").getUserOutput = null as any;

export { ImageBuilderArgs } from "./imageBuilder";
export type ImageBuilder = import("./imageBuilder").ImageBuilder;
export const ImageBuilder: typeof import("./imageBuilder").ImageBuilder = null as any;

export { StackArgs } from "./stack";
export type Stack = import("./stack").Stack;
export const Stack: typeof import("./stack").Stack = null as any;

export { StackFleetAssociationArgs } from "./stackFleetAssociation";
export type StackFleetAssociation = import("./stackFleetAssociation").StackFleetAssociation;
export const StackFleetAssociation: typeof import("./stackFleetAssociation").StackFleetAssociation = null as any;

export { StackUserAssociationArgs } from "./stackUserAssociation";
export type StackUserAssociation = import("./stackUserAssociation").StackUserAssociation;
export const StackUserAssociation: typeof import("./stackUserAssociation").StackUserAssociation = null as any;

export { UserArgs } from "./user";
export type User = import("./user").User;
export const User: typeof import("./user").User = null as any;

utilities.lazyLoad(exports, ["AppBlock"], () => require("./appBlock"));
utilities.lazyLoad(exports, ["Application"], () => require("./application"));
utilities.lazyLoad(exports, ["ApplicationEntitlementAssociation"], () => require("./applicationEntitlementAssociation"));
utilities.lazyLoad(exports, ["ApplicationFleetAssociation"], () => require("./applicationFleetAssociation"));
utilities.lazyLoad(exports, ["DirectoryConfig"], () => require("./directoryConfig"));
utilities.lazyLoad(exports, ["Entitlement"], () => require("./entitlement"));
utilities.lazyLoad(exports, ["Fleet"], () => require("./fleet"));
utilities.lazyLoad(exports, ["getAppBlock","getAppBlockOutput"], () => require("./getAppBlock"));
utilities.lazyLoad(exports, ["getApplication","getApplicationOutput"], () => require("./getApplication"));
utilities.lazyLoad(exports, ["getDirectoryConfig","getDirectoryConfigOutput"], () => require("./getDirectoryConfig"));
utilities.lazyLoad(exports, ["getEntitlement","getEntitlementOutput"], () => require("./getEntitlement"));
utilities.lazyLoad(exports, ["getFleet","getFleetOutput"], () => require("./getFleet"));
utilities.lazyLoad(exports, ["getImageBuilder","getImageBuilderOutput"], () => require("./getImageBuilder"));
utilities.lazyLoad(exports, ["getStack","getStackOutput"], () => require("./getStack"));
utilities.lazyLoad(exports, ["getStackFleetAssociation","getStackFleetAssociationOutput"], () => require("./getStackFleetAssociation"));
utilities.lazyLoad(exports, ["getStackUserAssociation","getStackUserAssociationOutput"], () => require("./getStackUserAssociation"));
utilities.lazyLoad(exports, ["getUser","getUserOutput"], () => require("./getUser"));
utilities.lazyLoad(exports, ["ImageBuilder"], () => require("./imageBuilder"));
utilities.lazyLoad(exports, ["Stack"], () => require("./stack"));
utilities.lazyLoad(exports, ["StackFleetAssociation"], () => require("./stackFleetAssociation"));
utilities.lazyLoad(exports, ["StackUserAssociation"], () => require("./stackUserAssociation"));
utilities.lazyLoad(exports, ["User"], () => require("./user"));

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:appstream:AppBlock":
                return new AppBlock(name, <any>undefined, { urn })
            case "aws-native:appstream:Application":
                return new Application(name, <any>undefined, { urn })
            case "aws-native:appstream:ApplicationEntitlementAssociation":
                return new ApplicationEntitlementAssociation(name, <any>undefined, { urn })
            case "aws-native:appstream:ApplicationFleetAssociation":
                return new ApplicationFleetAssociation(name, <any>undefined, { urn })
            case "aws-native:appstream:DirectoryConfig":
                return new DirectoryConfig(name, <any>undefined, { urn })
            case "aws-native:appstream:Entitlement":
                return new Entitlement(name, <any>undefined, { urn })
            case "aws-native:appstream:Fleet":
                return new Fleet(name, <any>undefined, { urn })
            case "aws-native:appstream:ImageBuilder":
                return new ImageBuilder(name, <any>undefined, { urn })
            case "aws-native:appstream:Stack":
                return new Stack(name, <any>undefined, { urn })
            case "aws-native:appstream:StackFleetAssociation":
                return new StackFleetAssociation(name, <any>undefined, { urn })
            case "aws-native:appstream:StackUserAssociation":
                return new StackUserAssociation(name, <any>undefined, { urn })
            case "aws-native:appstream:User":
                return new User(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "appstream", _module)
