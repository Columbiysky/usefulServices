data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./databases/models",
    "--dialect", "postgres", // | postgres | sqlite
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "postgresql://postgres:0o9i*U&Y@cbsky.ru:5432/postgres?sslmode=disable&&TimeZone=Europe/London"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}