# _Trello_ Open Microservice

> This is an trello service

[![Open Microservice Specification Version](https://img.shields.io/badge/Open%20Microservice-1.0-477bf3.svg)](https://openmicroservices.org) [![Open Microservices Spectrum Chat](https://withspectrum.github.io/badge/badge.svg)](https://spectrum.chat/open-microservices) [![Open Microservices Code of Conduct](https://img.shields.io/badge/Contributor%20Covenant-v1.4%20adopted-ff69b4.svg)](https://github.com/oms-services/.github/blob/master/CODE_OF_CONDUCT.md) [![Open Microservices Commitzen](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg)](http://commitizen.github.io/cz-cli/) [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

## Introduction

This project is an example implementation of the [Open Microservice Specification](https://openmicroservices.org), a standard originally created at [Storyscript](https://storyscript.io) for building highly-portable "microservices" that expose the events, actions, and APIs inside containerized software.

## Getting Started

The `oms` command-line interface allows you to interact with Open Microservices. If you're interested in creating an Open Microservice the CLI also helps validate, test, and debug your `oms.yml` implementation!

See the [oms-cli](https://github.com/microservices/oms) project to learn more!

### Installation

```
npm install -g @microservices/oms
```

## Usage

### Open Microservices CLI Usage

Once you have the [oms-cli](https://github.com/microservices/oms) installed, you can run any of the following commands from within this project's root directory:

#### Actions

##### createBoard

> Create new board in trello.

##### Action Arguments

| Argument Name | Type     | Required | Default | Description                         |
| :------------ | :------- | :------- | :------ | :---------------------------------- |
| boardName     | `string` | `true`   | None    | The name for new board.             |
| API_KEY       | `string` | `true`   | None    | The API key of trello account.      |
| ACCESS_TOKEN  | `string` | `true`   | None    | The access token of trello account. |

```shell
oms run createBoard \
    -a boardName='*****' \
    -e API_KEY=$API_KEY \
    -e ACCESS_TOKEN=$ACCESS_TOKEN
```

##### getBoard

> Get board details from trello.

##### Action Arguments

| Argument Name | Type     | Required | Default | Description                         |
| :------------ | :------- | :------- | :------ | :---------------------------------- |
| boardId       | `string` | `true`   | None    | The board ID of trello account.     |
| API_KEY       | `string` | `true`   | None    | The API key of trello account.      |
| ACCESS_TOKEN  | `string` | `true`   | None    | The access token of trello account. |

```shell
oms run getBoard \
    -a boardId='*****' \
    -e API_KEY=$API_KEY \
    -e ACCESS_TOKEN=$ACCESS_TOKEN
```

##### createList

> Create new list in trello.

##### Action Arguments

| Argument Name | Type     | Required | Default | Description                           |
| :------------ | :------- | :------- | :------ | :------------------------------------ |
| boardId       | `string` | `true`   | None    | The board ID to where to create list. |
| listName      | `string` | `true`   | None    | The name for new list.                |
| API_KEY       | `string` | `true`   | None    | The API key of trello account.        |
| ACCESS_TOKEN  | `string` | `true`   | None    | The access token of trello account.   |

```shell
oms run createList \
    -a boardId='*****' \
    -a listName='*****' \
    -e API_KEY=$API_KEY \
    -e ACCESS_TOKEN=$ACCESS_TOKEN
```

##### getLists

> Get all list details from trello board.

##### Action Arguments

| Argument Name | Type     | Required | Default | Description                         |
| :------------ | :------- | :------- | :------ | :---------------------------------- |
| boardId       | `string` | `true`   | None    | The board ID of trello account.     |
| API_KEY       | `string` | `true`   | None    | The API key of trello account.      |
| ACCESS_TOKEN  | `string` | `true`   | None    | The access token of trello account. |

```shell
oms run getLists \
    -a boardId='*****' \
    -e API_KEY=$API_KEY \
    -e ACCESS_TOKEN=$ACCESS_TOKEN
```

##### addCard

> Create new card on list.

##### Action Arguments

| Argument Name | Type     | Required | Default | Description                           |
| :------------ | :------- | :------- | :------ | :------------------------------------ |
| name          | `string` | `true`   | None    | The name of card to be added.         |
| description   | `string` | `false`  | None    | The description for the card.         |
| listId        | `string` | `true`   | None    | The ID of list where card has to add. |
| API_KEY       | `string` | `true`   | None    | The API key of trello account.        |
| ACCESS_TOKEN  | `string` | `true`   | None    | The access token of trello account.   |

```shell
oms run addCard \
    -a name='*****' \
    -a description='*****' \
    -a listId='*****' \
    -e API_KEY=$API_KEY \
    -e ACCESS_TOKEN=$ACCESS_TOKEN
```

##### getCards

> Get all card details from all lists.

##### Action Arguments

| Argument Name | Type     | Required | Default | Description                         |
| :------------ | :------- | :------- | :------ | :---------------------------------- |
| boardId       | `string` | `true`   | None    | The board ID of trello account.     |
| API_KEY       | `string` | `true`   | None    | The API key of trello account.      |
| ACCESS_TOKEN  | `string` | `true`   | None    | The access token of trello account. |

```shell
oms run getCards \
    -a boardId='*****' \
    -e API_KEY=$API_KEY \
    -e ACCESS_TOKEN=$ACCESS_TOKEN
```

##### moveCard

> Move card from one list to another.

##### Action Arguments

| Argument Name | Type     | Required | Default | Description                            |
| :------------ | :------- | :------- | :------ | :------------------------------------- |
| cardId        | `string` | `true`   | None    | The ID of card to move.                |
| listId        | `string` | `true`   | None    | The ID of list where to move the card. |
| API_KEY       | `string` | `true`   | None    | The API key of trello account.         |
| ACCESS_TOKEN  | `string` | `true`   | None    | The access token of trello account.    |

```shell
oms run moveCard \
    -a cardId='*****' \
    -a listId='*****' \
    -e API_KEY=$API_KEY \
    -e ACCESS_TOKEN=$ACCESS_TOKEN
```

##### copyCard

> Copy card from one list to another.

##### Action Arguments

| Argument Name | Type     | Required | Default | Description                            |
| :------------ | :------- | :------- | :------ | :------------------------------------- |
| cardId        | `string` | `true`   | None    | The ID of card to copy.                |
| listId        | `string` | `true`   | None    | The ID of list where to copy the card. |
| API_KEY       | `string` | `true`   | None    | The API key of trello account.         |
| ACCESS_TOKEN  | `string` | `true`   | None    | The access token of trello account.    |

```shell
oms run copyCard \
    -a cardId='*****' \
    -a listId='*****' \
    -e API_KEY=$API_KEY \
    -e ACCESS_TOKEN=$ACCESS_TOKEN
```

##### card

> Triggered anytime when any new card is added in list.

##### Action Arguments

| Argument Name | Type      | Required | Default | Description                                                                                         |
| :------------ | :-------- | :------- | :------ | :-------------------------------------------------------------------------------------------------- |
| boardId       | `string`  | `true`   | None    | The board ID to subscribe.                                                                          |
| listId        | `string`  | `false`  | None    | The list ID to subscribe the list.                                                                  |
| existing      | `boolean` | `true`   | None    | Set true to get all existing cards or false to get only new card added to board after subscription. |
| API_KEY       | `string`  | `true`   | None    | The API key of trello account.                                                                      |
| ACCESS_TOKEN  | `string`  | `true`   | None    | The access token of trello account.                                                                 |

```shell
oms subscribe card \
    -a boardId='*****' \
    -a listId='*****' \
    -a existing='*****' \
    -e API_KEY=$API_KEY \
    -e ACCESS_TOKEN=$ACCESS_TOKEN
```

##### deleteBoard

> Delete board from trello.

##### Action Arguments

| Argument Name | Type     | Required | Default | Description                         |
| :------------ | :------- | :------- | :------ | :---------------------------------- |
| boardId       | `string` | `true`   | None    | The board ID to delete.             |
| API_KEY       | `string` | `true`   | None    | The API key of trello account.      |
| ACCESS_TOKEN  | `string` | `true`   | None    | The access token of trello account. |

```shell
oms run deleteBoard \
    -a boardId='*****' \
    -e API_KEY=$API_KEY \
    -e ACCESS_TOKEN=$ACCESS_TOKEN
```

##### getAllBoards

> Get board details from trello.

##### Action Arguments

| Argument Name | Type     | Required | Default | Description                         |
| :------------ | :------- | :------- | :------ | :---------------------------------- |
| username      | `string` | `true`   | None    | The username of trello account.     |
| API_KEY       | `string` | `true`   | None    | The API key of trello account.      |
| ACCESS_TOKEN  | `string` | `true`   | None    | The access token of trello account. |

```shell
oms run getAllBoards \
    -a username='*****' \
    -e API_KEY=$API_KEY \
    -e ACCESS_TOKEN=$ACCESS_TOKEN
```

## Contributing

All suggestions in how to improve the specification and this guide are very welcome. Feel free share your thoughts in the Issue tracker, or even better, fork the repository to implement your own ideas and submit a pull request.

[![Edit trello on CodeSandbox](https://codesandbox.io/static/img/play-codesandbox.svg)](https://codesandbox.io/s/github/oms-services/trello)

This project is guided by [Contributor Covenant](https://github.com/oms-services/.github/blob/master/CODE_OF_CONDUCT.md). Please read out full [Contribution Guidelines](https://github.com/oms-services/.github/blob/master/CONTRIBUTING.md).

## Additional Resources

- [Install the CLI](https://github.com/microservices/oms) - The OMS CLI helps developers create, test, validate, and build microservices.
- [Example OMS Services](https://github.com/oms-services) - Examples of OMS-compliant services written in a variety of languages.
- [Example Language Implementations](https://github.com/microservices) - Find tooling & language implementations in Node, Python, Scala, Java, Clojure.
- [Storyscript Hub](https://hub.storyscript.io) - A public registry of OMS services.
- [Community Chat](https://spectrum.chat/open-microservices) - Have ideas? Questions? Join us on Spectrum.
