---
layout: "docs"
page_title: "Configuring Local Values"
sidebar_current: "docs-config-locals"
description: |-
  Local values assign a name to an expression that can then be used multiple times
  within a module.
---

# Variable Configuration

Local values assign a name to an expression, that can then be used multiple
times within a module.

Comparing modules to functions in a traditional programming language,
if [variables](./variables.html) are analogous to function arguments and
[outputs](./outputs.html) are analogous to function return values then
_local values_ are comparable to a function's local variables.

This page assumes you're familiar with the
[configuration syntax](/docs/configuration/syntax.html)
already.

## Example

Local values are defined in `locals` blocks:

```hcl

locals {
  greeting      = "hello"
  instance_ids  = "${aws_instance.foo.*.id}"
  defaulted_foo = "${var.foo != "" ? var.foo : "default_foo"}"
}
```

## Description

The `locals` block defines one or more local variables within a module.
Each `locals` block can have as many locals as needed, and there can be any
number of `locals` blocks within a module.

The names given for the items in the `locals` block must be unique throughout
a module. The given value can be any expression that is valid within
the current module.

The expression of a local value can refer to other locals, but as usual
reference cycles are not allowed. That is, a local cannot refer to itself
or to a variable that refers (directly or indirectly) back to it.
