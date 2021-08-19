import { Node } from "constructs";
import fs = require("fs");
import path = require("path");
import os = require("os");
import { App } from "../lib";
import { TerraformStack } from "./terraform-stack";
import { FUTURE_FLAGS } from "./features";

/**
 * Testing utilities for cdktf applications.
 */
export class Testing {
  /**
   * Returns an app for testing with the following properties:
   * - Output directory is a temp dir.
   */
  public static app(): App {
    const outdir = fs.mkdtempSync(path.join(os.tmpdir(), "cdktf.outdir."));
    const app = new App({ outdir, stackTraces: false });
    return this.stubVersion(this.enableFutureFlags(app));
  }

  public static stubVersion(app: App): App {
    Node.of(app).setContext("cdktfVersion", "stubbed");
    return app;
  }

  public static enableFutureFlags(app: App): App {
    const node = Node.of(app);
    Object.entries(FUTURE_FLAGS).forEach(([key, value]) =>
      node.setContext(key, value)
    );
    return app;
  }

  /**
   * Returns the Terraform synthesized JSON.
   */
  public static synth(stack: TerraformStack) {
    const errors = Node.of(stack).validate();
    if (errors.length > 0) {
      throw new Error(`${errors.length} Error found in stack:

${errors.map((err) => `${err.source}: ${err.message}`).join("\n")}`);
    }
    const tfConfig = stack.toTerraform();

    return JSON.stringify(tfConfig, null, 2);
  }

  /* istanbul ignore next */
  private constructor() {
    return;
  }
}
