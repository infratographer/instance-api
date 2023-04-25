-- +goose Up
-- +goose StatementBegin

CREATE TABLE instance_providers (
  id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  name STRING NOT NULL,
  tenant_id UUID NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMPTZ,
  INDEX idx_instances_tenant_id (tenant_id)
);

CREATE TABLE instances (
  id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  name STRING NOT NULL,
  tenant_id UUID NOT NULL,
  location_id UUID NOT NULL,
  instance_provider_id UUID NOT NULL REFERENCES instance_providers(id),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMPTZ,
  INDEX idx_instances_tenant_id (tenant_id),
  INDEX idx_instances_location_id (location_id),
  INDEX idx_instances_instance_provider_id (instance_provider_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE instances;
DROP TABLE instance_providers;

-- +goose StatementEnd
