---
title: "Cli Flags"
slug: "Cli Flags"
description: ""
summary: ""
date: 2023-08-15T21:01:37+01:00
lastmod: 2023-08-15T21:01:37+01:00
draft: false
menu:
  docs:
    parent: ""
    identifier: "cli-flags-85030aefbc53456496023ea81b6941f9"
weight: 7006
toc: true
seo:
  title: "" # custom title (optional)
  description: "" # custom description (recommended)
  canonical: "" # custom canonical URL (optional)
  noindex: false # false (default) or true
---
The CLI flags provide a method for executing the permguard binaries via the command line.
There are multiple executables that can be run, generally they can be divided between servers and provisioners.

## Servers

Irrespective of the chosen distribution, the binary supports the following flags:

---
**--debug**: *Enable debug mode (default `false`).*

---
**--log.level**: *Set log level (default `INFO`, options: `DEBUG`, `INFO`, `WARN`, `ERROR`, `DPANIC`, `PANIC`, `FATAL`).*

<details>
  <summary>Options</summary>

| LEVEL     | MEANING                                                                                                          |
|-----------|------------------------------------------------------------------------------------------------------------------|
| DEBUG     | Debug logs are typically voluminous, and are usually disabled in production.                                     |
| INFO      | Info is the default logging priority.                                                                            |
| WARN      | Warn logs are more important than Info, but don't need individual human review.                                  |
| ERROR     | Error logs are high-priority. If an application is running smoothly, it shouldn't generate any error-level logs. |
| DPANIC    | DPanic logs are particularly important errors. In development the logger panics after writing the message.       |
| PANIC     | Panic logs a message, then panics.                                                                               |
| FATAL     | Fatal logs a message, then calls os.Exit(1).                                                                     |

</details>

---

**--storage.central.engine**: *Data storage engine to be used for central data (default `postgres`).*

---

**--storage.proximity.engine**: *Data storage engine to be used for proximity data (default `badger`).*

---

**Storage Engines**: Storage engine flags are used to configure the storage engine to be used for the central and proximity data.

<details>
  <summary>PostgreSQL</summary>

**--stroage.engine.postgres.host**: *postgres host (default `localhost`).*

---

**--stroage.engine.postgres.port**: *postgres port (default `5432`).*

---

**--stroage.engine.postgres.sslmode**: *postgres port (default `disable`).*

<details>
  <summary>Options</summary>

| LEVEL       | MEANING                                                                                                                                                          |
|-------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| DISABLE     | The user doesn't care about security and doesn't want to pay the overhead of encryption.                                                                         |
| REQUIRE     | The user wants their data to be encrypted and accepts the overhead. They trust that the network will ensure they always connect to the intended server.          |
| VERIFY-CA   | The user wants their data encrypted and accepts the overhead. They want to ensure they connect to a server that they trust.                                      |
| VERIFY-FULL | The user wants their data encrypted and accepts the overhead. They want to ensure they connect to a server they trust and verify that it's the one they specify. |

</details>

---

**--stroage.engine.postgres.auth.username**: *postgres username (default `admin`).*

---

**--stroage.engine.postgres.auth.password**: *postgres password (default `admin`).*

---

**--stroage.engine.postgres.database**: *postgres database (default `permguard`).*

---

</details>

<details>
  <summary>Badger</summary>

</details>

---

**--server.appdata**: *Directory to be used as application data (default `./`).*

---

### server-aap

{{< callout >}} Account Administration Point. {{< /callout >}}

**--server.aap.grpc.port int**: *port to be used for exposing the aap grpc services. (default `9091`).*

---

**--server.aap.http.port int**: *port to be used for exposing the aap http services. (default `8081`).*

---

### server-pap

{{< callout >}} Policy Administration Point. {{< /callout >}}

**--server.pap.grpc.port int**: *port to be used for exposing the pap grpc services. (default `9092`).*

---

**--server.pap.http.port int**: *port to be used for exposing the pap http services. (default `8082`).*

---

### server-prp

{{< callout >}} Policy Retrieval Point. {{< /callout >}}

### server-pdp

{{< callout >}} Policy Decision Point. {{< /callout >}}

**--server.pdp.grpc.port int**: *port to be used for exposing the pdp grpc services. (default `9096`).*

---

**--server.pdp.http.port int**: *port to be used for exposing the pdp http services. (default `8086`).*

---

## Provisioners

Irrespective of the chosen distribution, the binary supports the following flags:

---
**--debug**: *Enable debug mode (default `false`).*

---
**--log.level**: *Set log level (default `INFO`, options: `DEBUG`, `INFO`, `WARN`, `ERROR`, `DPANIC`, `PANIC`, `FATAL`).*

<details>
  <summary>Options</summary>

| LEVEL     | MEANING                                                                                                          |
|-----------|------------------------------------------------------------------------------------------------------------------|
| DEBUG     | Debug logs are typically voluminous, and are usually disabled in production.                                     |
| INFO      | Info is the default logging priority.                                                                            |
| WARN      | Warn logs are more important than Info, but don't need individual human review.                                  |
| ERROR     | Error logs are high-priority. If an application is running smoothly, it shouldn't generate any error-level logs. |
| DPANIC    | DPanic logs are particularly important errors. In development the logger panics after writing the message.       |
| PANIC     | Panic logs a message, then panics.                                                                               |
| FATAL     | Fatal logs a message, then calls os.Exit(1).                                                                     |

</details>

---

### provisioner-db-postgress

{{< callout >}} Provisioner of the postgress storage. {{< /callout >}}

**--stroage.engine.postgres.host**: *postgres host (default `localhost`).*

---

**--stroage.engine.postgres.port**: *postgres port (default `5432`).*

---

**--stroage.engine.postgres.sslmode**: *postgres port (default `disable`).*

<details>
  <summary>Options</summary>

| LEVEL       | MEANING                                                                                                                                                          |
|-------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| DISABLE     | The user doesn't care about security and doesn't want to pay the overhead of encryption.                                                                         |
| REQUIRE     | The user wants their data to be encrypted and accepts the overhead. They trust that the network will ensure they always connect to the intended server.          |
| VERIFY-CA   | The user wants their data encrypted and accepts the overhead. They want to ensure they connect to a server that they trust.                                      |
| VERIFY-FULL | The user wants their data encrypted and accepts the overhead. They want to ensure they connect to a server they trust and verify that it's the one they specify. |

</details>

---

**--stroage.engine.postgres.auth.username**: *postgres username (default `admin`).*

---

**--stroage.engine.postgres.auth.password**: *postgres password (default `admin`).*

---

**--stroage.engine.postgres.database**: *postgres database (default `permguard`).*

---
