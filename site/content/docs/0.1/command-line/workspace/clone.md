---
title: "Clone"
description: ""
summary: ""
date: 2023-08-17T11:47:15+01:00
lastmod: 2023-08-17T11:47:15+01:00
draft: false
menu:
  docs:
    parent: ""
    identifier: "clone-64624414-985f-47bf-ad58-34bad2ecd6ba"
weight: 5305
toc: true
seo:
  title: "" # custom title (optional)
  description: "" # custom description (recommended)
  canonical: "" # custom canonical URL (optional)
  noindex: false # false (default) or true
---
Using the `clone` command, it is possible to clone a remote repository locally.

```text
  ____                                               _
 |  _ \ ___ _ __ _ __ ___   __ _ _   _  __ _ _ __ __| |
 | |_) / _ \ '__| '_ ` _ \ / _` | | | |/ _` | '__/ _` |
 |  __/  __/ |  | | | | | | (_| | |_| | (_| | | | (_| |
 |_|   \___|_|  |_| |_| |_|\__, |\__,_|\__,_|_|  \__,_|
                           |___/

The official Permguard Command Line Interface - Copyright © 2022 Nitro Agility S.r.l.

This command clones a remote repository to the local permguard workspace.

Examples:
  # clone a remote repository to the local permguard workspace
  permguard clone 268786704340/magicfarmacia

  Find more information at: https://www.permguard.com/docs/0.1/using-the-cli/how-to-use/

Usage:
  permguard clone [flags]

Flags:
      --aap int   aap port (default 9091)
  -h, --help      help for clone
      --pap int   pap port (default 9092)

Global Flags:
  -o, --output string    output format (default "terminal")
  -v, --verbose          true for verbose output
  -w, --workdir string   workdir (default ".")
```

{{< callout context="caution" icon="alert-triangle" >}}
The output from your current version of Permguard may differ from the example provided on this page.
{{< /callout >}}

## Clone a repository

The `permguard clone` command allows you to clone a remote repository locally.

```bash
permguard clone permguard@localhost/273165098782/magicfarmacia
```

output:

```bash
Initialized empty permguard repository in 'magicfarmacia'.
Remote origin has been added.
Repo magicfarmacia has been added.
```

<details>
  <summary>
    JSON Output
  </summary>

```bash
permguard clone permguard@localhost/273165098782/magicfarmacia --output json
```

output:

```bash
{
  "repos": [
    {
      "repo": "magicfarmacia"
    }
  ]
}
```

</details>