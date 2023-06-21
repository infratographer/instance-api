-- Create "instance_providers" table
CREATE TABLE "instance_providers" ("id" character varying NOT NULL, "name" text NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- Create "instances" table
CREATE TABLE "instances" ("id" character varying NOT NULL, "name" text NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "tenant_id" character varying NOT NULL, "instance_provider_id" character varying NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "instances_instance_providers_instance_provider" FOREIGN KEY ("instance_provider_id") REFERENCES "instance_providers" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create "instance_metadata" table
CREATE TABLE "instance_metadata" ("id" character varying NOT NULL, "namespace" text NOT NULL, "data" jsonb NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "instance_id" character varying NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "instance_metadata_instances_instance" FOREIGN KEY ("instance_id") REFERENCES "instances" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
