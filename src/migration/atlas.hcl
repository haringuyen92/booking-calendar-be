data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./",
  ]
}

locals {
  url = "mysql://root:password@localhost:3306/db"
}

env "local" {
  src = data.external_schema.gorm.url
  dev = "docker://mysql/8/dev"
  url = local.url
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}