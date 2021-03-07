// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./user";
export * from "./userGroup";

// Import resources to register:
import { User } from "./user";
import { UserGroup } from "./userGroup";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:ElastiCache:User":
                return new User(name, <any>undefined, { urn })
            case "aws-native:ElastiCache:UserGroup":
                return new UserGroup(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "ElastiCache", _module)
