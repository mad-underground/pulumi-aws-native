// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./project";
export * from "./reportGroup";
export * from "./sourceCredential";

// Import resources to register:
import { Project } from "./project";
import { ReportGroup } from "./reportGroup";
import { SourceCredential } from "./sourceCredential";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "cloudformation:CodeBuild:Project":
                return new Project(name, <any>undefined, { urn })
            case "cloudformation:CodeBuild:ReportGroup":
                return new ReportGroup(name, <any>undefined, { urn })
            case "cloudformation:CodeBuild:SourceCredential":
                return new SourceCredential(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("cloudformation", "CodeBuild", _module)
