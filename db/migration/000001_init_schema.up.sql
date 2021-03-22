CREATE TABLE "contas" (
  "id" bigserial PRIMARY KEY,
  "proprietario" varchar NOT NULL,
  "saldo" bigint NOT NULL,
  "moeda" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entradas" (
  "id" bigserial PRIMARY KEY,
  "id_conta" bigint NOT NULL,
  "quantidade" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transferencias" (
  "id" bigserial PRIMARY KEY,
  "emissor_id_conta" bigint NOT NULL,
  "receptor_id_conta" bigint NOT NULL,
  "quantidade" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "entradas" ADD FOREIGN KEY ("id_conta") REFERENCES "contas" ("id");

ALTER TABLE "transferencias" ADD FOREIGN KEY ("emissor_id_conta") REFERENCES "contas" ("id");

ALTER TABLE "transferencias" ADD FOREIGN KEY ("receptor_id_conta") REFERENCES "contas" ("id");

CREATE INDEX ON "contas" ("proprietario");

CREATE INDEX ON "entradas" ("id_conta");

CREATE INDEX ON "transferencias" ("emissor_id_conta");

CREATE INDEX ON "transferencias" ("receptor_id_conta");

CREATE INDEX ON "transferencias" ("emissor_id_conta", "receptor_id_conta");
