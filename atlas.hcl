data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "./cmd/loader",
  ]
}

env "local" {
  src = data.external_schema.gorm.url
  dev = "docker://postgres/latest/dev?search_path=public"
  url = "postgres://postgres:postgres@127.0.0.1:5432/pgsql?search_path=public&sslmode=disable"
  migration {
    dir = "file://migrations"
  }
}
