---
title: "Init"
description: ""
summary: ""
date: 2023-08-17T11:47:15+01:00
lastmod: 2023-08-17T11:47:15+01:00
draft: false
menu:
  docs:
    parent: ""
    identifier: "init-46f6439f-06f3-41fe-afab-279f3de9b67f"
weight: 5301
toc: true
seo:
  title: "" # custom title (optional)
  description: "" # custom description (recommended)
  canonical: "" # custom canonical URL (optional)
  noindex: false # false (default) or true
---
Using the `Init` commands, it is possible to initialize a new repo in the working directory.

```text
  ____                                               _
 |  _ \ ___ _ __ _ __ ___   __ _ _   _  __ _ _ __ __| |
 | |_) / _ \ '__| '_ ` _ \ / _` | | | |/ _` | '__/ _` |
 |  __/  __/ |  | | | | | | (_| | |_| | (_| | | | (_| |
 |_|   \___|_|  |_| |_| |_|\__, |\__,_|\__,_|_|  \__,_|
                           |___/

The official Permguard Command Line Interface - Copyright © 2022 Nitro Agility S.r.l.

This command initializes a new repo in the working directory.

Examples:
  # initialize a new repo in the working directory
  permguard init

	Find more information at: https://www.permguard.com/docs/using-the-cli/how-to-use/

Usage:
  permguard init [flags]

Flags:
  -h, --help   help for init

Global Flags:
  -o, --output string    output format (default "terminal")
  -v, --verbose          true for verbose output
  -w, --workdir string   workdir (default ".")
```

{{< callout context="caution" icon="alert-triangle" >}}
The output from your current version of Permguard may differ from the example provided on this page.
{{< /callout >}}

## Initialize a workspace

The `permguard init` command initializes a new repo in the working directory.

```bash
permguard init
```

output:

```bash
Initialized empty permguard repository in '.permguard'.
```

<details>
  <summary>
    JSON Output
  </summary>

```bash
permguard checkout origin/273165098782/v1.0 --output json
```

output:

```bash
{
  "workspace": {
    "cwd": ".permguard"
  }
}
```

</details>