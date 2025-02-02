// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetServerlessContainer
    {
        /// <summary>
        /// Get information about a Yandex Cloud Serverless Container. 
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Yandex = Pulumi.Yandex;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var my_container = Yandex.GetServerlessContainer.Invoke(new()
        ///     {
        ///         ContainerId = "are1samplecontainer11",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// This data source is used to define Yandex Cloud Container that can be used by other resources.
        /// </summary>
        public static Task<GetServerlessContainerResult> InvokeAsync(GetServerlessContainerArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServerlessContainerResult>("yandex:index/getServerlessContainer:getServerlessContainer", args ?? new GetServerlessContainerArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Yandex Cloud Serverless Container. 
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Yandex = Pulumi.Yandex;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var my_container = Yandex.GetServerlessContainer.Invoke(new()
        ///     {
        ///         ContainerId = "are1samplecontainer11",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// This data source is used to define Yandex Cloud Container that can be used by other resources.
        /// </summary>
        public static Output<GetServerlessContainerResult> Invoke(GetServerlessContainerInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServerlessContainerResult>("yandex:index/getServerlessContainer:getServerlessContainer", args ?? new GetServerlessContainerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServerlessContainerArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Network access. If specified the revision will be attached to specified network
        /// * `connectivity.0.network_id` - Network the revision will have access to
        /// </summary>
        [Input("connectivity")]
        public Inputs.GetServerlessContainerConnectivityArgs? Connectivity { get; set; }

        /// <summary>
        /// Yandex Cloud Serverless Container id used to define container
        /// </summary>
        [Input("containerId")]
        public string? ContainerId { get; set; }

        /// <summary>
        /// Folder ID for the Yandex Cloud Serverless Container
        /// </summary>
        [Input("folderId")]
        public string? FolderId { get; set; }

        /// <summary>
        /// Yandex Cloud Serverless Container name used to define container
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        [Input("secrets")]
        private List<Inputs.GetServerlessContainerSecretArgs>? _secrets;

        /// <summary>
        /// Secrets for Yandex Cloud Serverless Container
        /// * `image.0.url` - URL of image that deployed as Yandex Cloud Serverless Container
        /// * `image.0.work_dir` - Working directory of Yandex Cloud Serverless Container
        /// * `image.0.digest` - Digest of image that deployed as Yandex Cloud Serverless Container
        /// * `image.0.command` - List of commands of the Yandex Cloud Serverless Container
        /// * `image.0.args` - List of arguments of the Yandex Cloud Serverless Container
        /// * `image.0.environment` -  A set of key/value environment variable pairs of Yandex Cloud Serverless Container
        /// </summary>
        public List<Inputs.GetServerlessContainerSecretArgs> Secrets
        {
            get => _secrets ?? (_secrets = new List<Inputs.GetServerlessContainerSecretArgs>());
            set => _secrets = value;
        }

        public GetServerlessContainerArgs()
        {
        }
        public static new GetServerlessContainerArgs Empty => new GetServerlessContainerArgs();
    }

    public sealed class GetServerlessContainerInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Network access. If specified the revision will be attached to specified network
        /// * `connectivity.0.network_id` - Network the revision will have access to
        /// </summary>
        [Input("connectivity")]
        public Input<Inputs.GetServerlessContainerConnectivityInputArgs>? Connectivity { get; set; }

        /// <summary>
        /// Yandex Cloud Serverless Container id used to define container
        /// </summary>
        [Input("containerId")]
        public Input<string>? ContainerId { get; set; }

        /// <summary>
        /// Folder ID for the Yandex Cloud Serverless Container
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        /// <summary>
        /// Yandex Cloud Serverless Container name used to define container
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("secrets")]
        private InputList<Inputs.GetServerlessContainerSecretInputArgs>? _secrets;

        /// <summary>
        /// Secrets for Yandex Cloud Serverless Container
        /// * `image.0.url` - URL of image that deployed as Yandex Cloud Serverless Container
        /// * `image.0.work_dir` - Working directory of Yandex Cloud Serverless Container
        /// * `image.0.digest` - Digest of image that deployed as Yandex Cloud Serverless Container
        /// * `image.0.command` - List of commands of the Yandex Cloud Serverless Container
        /// * `image.0.args` - List of arguments of the Yandex Cloud Serverless Container
        /// * `image.0.environment` -  A set of key/value environment variable pairs of Yandex Cloud Serverless Container
        /// </summary>
        public InputList<Inputs.GetServerlessContainerSecretInputArgs> Secrets
        {
            get => _secrets ?? (_secrets = new InputList<Inputs.GetServerlessContainerSecretInputArgs>());
            set => _secrets = value;
        }

        public GetServerlessContainerInvokeArgs()
        {
        }
        public static new GetServerlessContainerInvokeArgs Empty => new GetServerlessContainerInvokeArgs();
    }


    [OutputType]
    public sealed class GetServerlessContainerResult
    {
        /// <summary>
        /// Concurrency of Yandex Cloud Serverless Container
        /// </summary>
        public readonly int Concurrency;
        /// <summary>
        /// Network access. If specified the revision will be attached to specified network
        /// * `connectivity.0.network_id` - Network the revision will have access to
        /// </summary>
        public readonly Outputs.GetServerlessContainerConnectivityResult? Connectivity;
        public readonly string? ContainerId;
        /// <summary>
        /// Core fraction (**0...100**) of the Yandex Cloud Serverless Container
        /// </summary>
        public readonly int CoreFraction;
        public readonly int Cores;
        /// <summary>
        /// Creation timestamp of the Yandex Cloud Serverless Container
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// Description of the Yandex Cloud Serverless Container
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Execution timeout (duration format) of Yandex Cloud Serverless Container
        /// </summary>
        public readonly string ExecutionTimeout;
        public readonly string? FolderId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetServerlessContainerImageResult> Images;
        /// <summary>
        /// A set of key/value label pairs assigned to the Yandex Cloud Serverless Container
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Memory in megabytes of Yandex Cloud Serverless Container
        /// </summary>
        public readonly int Memory;
        public readonly string? Name;
        /// <summary>
        /// Last revision ID of the Yandex Cloud Serverless Container
        /// </summary>
        public readonly string RevisionId;
        /// <summary>
        /// Secrets for Yandex Cloud Serverless Container
        /// * `image.0.url` - URL of image that deployed as Yandex Cloud Serverless Container
        /// * `image.0.work_dir` - Working directory of Yandex Cloud Serverless Container
        /// * `image.0.digest` - Digest of image that deployed as Yandex Cloud Serverless Container
        /// * `image.0.command` - List of commands of the Yandex Cloud Serverless Container
        /// * `image.0.args` - List of arguments of the Yandex Cloud Serverless Container
        /// * `image.0.environment` -  A set of key/value environment variable pairs of Yandex Cloud Serverless Container
        /// </summary>
        public readonly ImmutableArray<Outputs.GetServerlessContainerSecretResult> Secrets;
        /// <summary>
        /// Service account ID of Yandex Cloud Serverless Container
        /// </summary>
        public readonly string ServiceAccountId;
        /// <summary>
        /// Invoke URL of the Yandex Cloud Serverless Container
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private GetServerlessContainerResult(
            int concurrency,

            Outputs.GetServerlessContainerConnectivityResult? connectivity,

            string? containerId,

            int coreFraction,

            int cores,

            string createdAt,

            string description,

            string executionTimeout,

            string? folderId,

            string id,

            ImmutableArray<Outputs.GetServerlessContainerImageResult> images,

            ImmutableDictionary<string, string> labels,

            int memory,

            string? name,

            string revisionId,

            ImmutableArray<Outputs.GetServerlessContainerSecretResult> secrets,

            string serviceAccountId,

            string url)
        {
            Concurrency = concurrency;
            Connectivity = connectivity;
            ContainerId = containerId;
            CoreFraction = coreFraction;
            Cores = cores;
            CreatedAt = createdAt;
            Description = description;
            ExecutionTimeout = executionTimeout;
            FolderId = folderId;
            Id = id;
            Images = images;
            Labels = labels;
            Memory = memory;
            Name = name;
            RevisionId = revisionId;
            Secrets = secrets;
            ServiceAccountId = serviceAccountId;
            Url = url;
        }
    }
}
