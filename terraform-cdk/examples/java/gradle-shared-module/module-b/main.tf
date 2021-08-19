
provider "aws" {
  region                     = "us-west-1"
  alias                      = "external"
  skip_requesting_account_id = true
  profile                    = "external"
}

resource "aws_secretsmanager_secret" "local-secret" {
  name     = "internal-secret"
  // Here we're referencing a provider generated by the
  // base application module that was imported via gradle
  provider = aws.default
}


resource "aws_secretsmanager_secret" "external-secret" {
  name     = "external-secret"
  provider = aws.external
}
