// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ProviderType.Submod
{
    [ProviderTypeResourceType("providerType:submod:provider")]
    public partial class Provider : global::Pulumi.CustomResource
    {
        [Output("a")]
        public Output<bool?> A { get; private set; } = null!;


        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs? args = null, CustomResourceOptions? options = null)
            : base("providerType:submod:provider", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Provider(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("providerType:submod:provider", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Provider resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Provider Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Provider(name, id, options);
        }
    }

    public sealed class ProviderArgs : global::Pulumi.ResourceArgs
    {
        [Input("a")]
        public bool? A { get; set; }

        public ProviderArgs()
        {
        }
        public static new ProviderArgs Empty => new ProviderArgs();
    }
}
