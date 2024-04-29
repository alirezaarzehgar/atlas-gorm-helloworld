# How to integrate atlas with GORM ?

Install atlas:
```bash
curl -sSf https://atlasgo.sh | sh
```

First of all you should run a database instance:
```bash
	docker run \
	--name devdb \
	-e POSTGRES_PASSWORD=admin \
	-e POSTGRES_USER=admin \
	-e POSTGRES_DB=db \
	-p 5432:5432 -d postgres:alpine3.18
```

Import ariga.io/atlas-provider-gorm/gormschema on project:
```go
// tools.go
//go:build tools
package main

import _ "ariga.io/atlas-provider-gorm/gormschema"
```

Create HCL file:
```hcl
data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./models",
    "--dialect", "postgres", // | mysql | sqlite | sqlserver
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "docker://postgres/alpine3.18"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
```

Create model:
```go
// models/model.go
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Age      uint
}
```

Then you should migrate your database:
```bash
atlas migrate diff --env gorm
```

Feel free to change model and save diff. Then you can migrate database using following command:
```bash
atlas schema apply \
    --url "postgres://admin:admin@localhost:5432/db?sslmode=disable" \
    --to "file://migrations" \
    --dev-url "docker://postgres/alpine3.18"
```
