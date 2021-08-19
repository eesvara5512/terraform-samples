/*

Starts a Docker container with an Nginx server.

Steps:

- Start Docker on your machine
- Run `cdktf get` to generate provider bindings
- Run `go mod tidy` to install dependencies (e.g. jsii-runtime-go)
- Run `cdktf deploy` to run the Nginx container
- Visit http://localhost:8000/ to view the Nginx default page
- Optional: `docker ps` will show the running container
- Destroy with `cdktf destroy`

*/

package main

import (
	"github.com/hashicorp/terraform-cdk/examples/go/docker/generated/kreuzwerker/docker"

	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewExampleCdktfGoDockerStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	docker.NewDockerProvider(stack, jsii.String("provider"), &docker.DockerProviderConfig{})

	dockerImage := docker.NewImage(stack, jsii.String("nginxImage"), &docker.ImageConfig{
		Name:        jsii.String("nginx:latest"),
		KeepLocally: jsii.Bool(false),
	})

	docker.NewContainer(stack, jsii.String("nginxContainer"), &docker.ContainerConfig{
		Image: dockerImage.Latest(),
		Name:  jsii.String("tutorial"),
		Ports: &[]*docker.ContainerPorts{{
			Internal: jsii.Number(80), External: jsii.Number(8000),
		}},
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewExampleCdktfGoDockerStack(app, "ExampleCdktfGoDockerStack")

	app.Synth()
}
