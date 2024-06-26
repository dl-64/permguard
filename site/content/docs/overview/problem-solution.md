---
title: "Problem and Solution"
description: ""
summary: ""
date: 2023-08-27T10:21:48+01:00
lastmod: 2023-08-27T10:21:48+01:00
draft: false
menu:
  docs:
    parent: ""
    identifier: "problem-solution-fd262b5997389b467f3cf77720aafe70"
weight: 1002
toc: true
seo:
  title: "" # custom title (optional)
  description: "" # custom description (recommended)
  canonical: "" # custom canonical URL (optional)
  noindex: false # false (default) or true
---
Companies and software solutions commonly adopt an `Authentication Provider` as a `central point to manage users`. By centrally managing users, different applications can avoid the need for custom authentication implementations and delegate authentication to the Authentication Provider.

Similarly, software often necessitates a centralized solution for authorization. `PermGuard` emerges as an `Authorization Provider` by implementing a `central point to manage authorizations`.

{{< callout context="note" icon="info-circle" >}}
Authorization is a complex aspect that should not be implemented separately in each application, similar to authentication. Building an authentication layer compliant with the latest security standards isn't a simple task, and the same applies to the authorization layer.
{{< /callout >}}

## Level of Authorizations

There are serval level of authroization that can be implemented, essentially they can be summaries as followting:

- Level 1: Application and API Access
- Level 2: Functional Access
- Level 3: Data Access

PermGuard has been specifically designed to cover and facilitate the implementation of all these levels of authorization.

### Level 1: Application and API Access
This level of authorization covers the access to the application and the API. It is the most common level of authorization and it is the one that is typically implemented by an Authentication Provider.

### Level 2: Functional Access
This level of authorization covers the access to the different functionalities of the application. It is the level of authorization that is typically implemented by the application itself.

### Level 3: Data Access
This level of authorization covers the access to the different data of the application. It is the level of authorization that is typically implemented by the application itself.

## The Problem

In a scenario where is missing an authorization layer, users are typically annotated with custom role metadata, and applications implement custom business logic based on the roles associated with users.

However, this approach presents several drawbacks:

- **Tight Coupled Authorization Logic**: The authorization logic is tightly coupled with the application code, leading to challenges in management and maintenance. If an administrator decides to create a new role, it requires code changes in the application code base.
- **Duplicated Code**: Logic for evaluating permissions and enforcing them is replicated across various sections of the application's code base. This duplication introduces maintenance complexities and the risk of inconsistencies in permission enforcement.
- **Limited Flexibility**: The application's authorization logic is limited to the capabilities offered by the Authentication Provider, which may pose challenges when implementing complex authorization requirements. And of course an Authentication Provider is not designed to manage permissions.
- **Delegation of Permissions**: Delegation is a multifaceted task that involves implementing custom logic to grant permissions to another identity. This is necessary in situations where an identity must pass on permissions to another identity, or where application components like workers or services need to act on behalf of another identity. While companies often use a functional account with extensive permissions for such operations, this poses a security risk. Efforts have been made to create identity tokens for delegating permissions, but they may not be applicable in all scenarios. For example, consider a worker operating on behalf of an identity after receiving a request via Apache Kafka messaging. In such cases, traditional authentication tokens like JWT may not be available or reliable.
- **Security Risks**: A missing authorization central point open the door to security risks, as it is difficult to track the permissions of each identity on different applications.

## The Solution

`PermGuard` implements an authorization layer. With this approach the previous drawbacks are mitigated:

- **Tight Coupled Authorization Logic**: This problem is fixed as the authorization logic is loosely coupled from the application code. Administrators can create new roles and permissions without requiring changes to the application's code base. The application retrieves the latest policies from the authorization layer without any code modifications and evaluates them using PermGuard.
- **Duplicated Code**: This problem is fixed as the authorization evaluation logic is implemented within PermGuard, eliminating duplicated code as the policies are authored externally to the application code.
- **Limited Flexibility**: This problem is fixed as using the PermGuard Policy Language or JSON-based API, it is possible to define complex authorization policies. This allows administrators to utilize a configuration language that is highly expressive, enabling the creation of custom and complex permissions.
- **Delegation of Permissions**: This problem is addressed through the implementation of Trusted Delegation, a feature that enables the creation of a security sandbox within which the application can handle and fulfill requests on behalf of another identity, eliminating the necessity for a functional account with elevated permissions. Additionally, this feature resolves the challenge of processing requests triggered by messages from messaging systems such as Apache Kafka.
- **Security Risks**: This problem is fixed as having a central point for managing the authorization layer enables the tracking of permissions for each identity across different applications.

{{< callout context="note" icon="info-circle" >}}
`PermGuard` further enhances the concept of the authorization layer by allowing the creation of multiple types of identities, including Users and Roles. Moreover, it implements features to delegate permissions to another identity.
{{< /callout >}}
