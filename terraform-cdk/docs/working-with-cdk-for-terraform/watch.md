# `cdktf watch` command

> ⚠️ Warning: The `watch` command is a new experimental command. **ONLY use in development environments**

The `watch` command watches a directory for changes and automatically synthesizes and deploys changes as they happen. It allows for rapid iterations when developing infrastructure, especially when working with serverless services. It currently supports only one stack at a time.

## Using `watch`

Watch reads your root `.gitignore` file to determine which files trigger a `synth` and it uses the `outdir` of your `cdktf.json` (or the default `cdktf.out`) to do so for a deploy. It will generate a checksum of the subdirectory containing the Terraform code for your stack to skip deploys if the synthesized Terraform config did not change.

### Checking your root `.gitignore`

Your root `.gitignore` needs to contain all files that are generated by a synth and a deploy. If that is not the case the watch will trigger itself and continuously try to deploy. Although Git supports nesting `.gitignore` files, the watch command currently only reads the root `.gitignore` file.

### Checking your environment

Watch should only be used for development environments. It is best to make sure that the terminal you're running watch in has no access keys that allow the cdktf-cli to deploy to your production environment.

### Running watch

An exemplary invocation of watch could be:

```
cdktf watch --stack dev --auto-approve
```

Please note that watch currently automatically deploys all changes and does not ask for confirmation!

### Troubleshooting

To troubleshoot watch set the `CDKTF_LOG_LEVEL` environment variable to `all`. By supplying `CDKTF_DISABLE_LOGGING=false` the debug output will be directed to a `cdktf.log` file in your projects root directory. The log will contain information about detected file system changes and the actions they triggered.
