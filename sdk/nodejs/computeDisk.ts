// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class ComputeDisk extends pulumi.CustomResource {
    /**
     * Get an existing ComputeDisk resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ComputeDiskState, opts?: pulumi.CustomResourceOptions): ComputeDisk {
        return new ComputeDisk(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/computeDisk:ComputeDisk';

    /**
     * Returns true if the given object is an instance of ComputeDisk.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ComputeDisk {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ComputeDisk.__pulumiType;
    }

    public readonly allowRecreate!: pulumi.Output<boolean | undefined>;
    public readonly blockSize!: pulumi.Output<number | undefined>;
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly diskPlacementPolicy!: pulumi.Output<outputs.ComputeDiskDiskPlacementPolicy>;
    public readonly folderId!: pulumi.Output<string>;
    public readonly imageId!: pulumi.Output<string | undefined>;
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly productIds!: pulumi.Output<string[]>;
    public readonly size!: pulumi.Output<number | undefined>;
    public readonly snapshotId!: pulumi.Output<string | undefined>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    public readonly type!: pulumi.Output<string | undefined>;
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a ComputeDisk resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ComputeDiskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ComputeDiskArgs | ComputeDiskState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ComputeDiskState | undefined;
            resourceInputs["allowRecreate"] = state ? state.allowRecreate : undefined;
            resourceInputs["blockSize"] = state ? state.blockSize : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["diskPlacementPolicy"] = state ? state.diskPlacementPolicy : undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["imageId"] = state ? state.imageId : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["productIds"] = state ? state.productIds : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["snapshotId"] = state ? state.snapshotId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as ComputeDiskArgs | undefined;
            resourceInputs["allowRecreate"] = args ? args.allowRecreate : undefined;
            resourceInputs["blockSize"] = args ? args.blockSize : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["diskPlacementPolicy"] = args ? args.diskPlacementPolicy : undefined;
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["imageId"] = args ? args.imageId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["snapshotId"] = args ? args.snapshotId : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["productIds"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ComputeDisk.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ComputeDisk resources.
 */
export interface ComputeDiskState {
    allowRecreate?: pulumi.Input<boolean>;
    blockSize?: pulumi.Input<number>;
    createdAt?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    diskPlacementPolicy?: pulumi.Input<inputs.ComputeDiskDiskPlacementPolicy>;
    folderId?: pulumi.Input<string>;
    imageId?: pulumi.Input<string>;
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    name?: pulumi.Input<string>;
    productIds?: pulumi.Input<pulumi.Input<string>[]>;
    size?: pulumi.Input<number>;
    snapshotId?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ComputeDisk resource.
 */
export interface ComputeDiskArgs {
    allowRecreate?: pulumi.Input<boolean>;
    blockSize?: pulumi.Input<number>;
    description?: pulumi.Input<string>;
    diskPlacementPolicy?: pulumi.Input<inputs.ComputeDiskDiskPlacementPolicy>;
    folderId?: pulumi.Input<string>;
    imageId?: pulumi.Input<string>;
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    name?: pulumi.Input<string>;
    size?: pulumi.Input<number>;
    snapshotId?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    zone?: pulumi.Input<string>;
}