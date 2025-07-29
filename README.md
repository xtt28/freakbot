# Welcome to Freakbot's source repository.



People in the ATCS 2028 Discord server are known for saying some rather
questionable things. This bot keeps track of who sends the most questionable
messages using the OpenAI moderation API and a SQLite database. The data is
presented in a leaderboard accessible via a slash command.

## Internals

In each guild, Freakbot creates a leaderboard and gives each user a counter of
"freaky" messages sent.

When a user sends a message, we use the free-of-charge OpenAI moderation API to
see if the message is flagged for inappropriate content. If it is indeed flagged,
we increment the user's freaky message counter in that guild. The data is stored
in a SQLite database. We also cache the data in a hash table that persists for as
long as the bot process is running.

## Commands

### `/about`

Shows information about the bot. The information is determined at compile time
and stored as constants in the internal/manifest/manifest.go file.

### `/freakerboard`

Shows the top 10 freakiest people in your Discord server in descending order.
Freakiness is determined by the amount of messages that a user has sent that
were flagged by the OpenAI moderation API.

## Running this bot

At the moment, we don't have a shared instance of Freakbot available for
everyone to use. However, whoever wants to use Freakbot is free to download the
bot and host it on their own server. One instance of Freakbot can be used to
manage multiple Discord guilds.

The preferred method of running this bot is Docker. We publish a Docker image by
means of GitHub Packages for everyone to use. The following tutorial assumes
that you will use Docker.

### Prepare your secrets and .env file

You must have an OpenAI API key and a Discord bot token to run this bot. After
getting these, please enter their values in the `sample.env` file in the root of
this repository, and rename the `sample.env` file to `.env`.

### Pull the Docker image

Pull the latest, bleeding-edge Docker image:

    docker pull ghcr.io/xtt28/freakbot:main

Or pull the Docker image of any release:

    docker pull ghcr.io/xtt28/freakbot:v1.2.3

### Run the Docker container

Run the container as follows:

```shell
docker run \
    --env-file .env \ # import .env file values
    -v freakbot-data:/data \ # persist data
    -d \ # run in background
    ghcr.io/xtt28/freakbot:main # replace main with your version
```

## License

This project is licensed under:

    SPDX-License-Identifier: AGPL-3.0-or-later

being in concordance with the contents of the LICENSE file in the root of this
repository.