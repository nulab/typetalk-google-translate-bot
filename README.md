# Typetalk Google Translate Bot

This is one of Typetalk webhook example.

When Typetalk users post in a language other than English, this Bot try to translate in English.
Then, this Bot reply a message translated in English.

## How to use

### 1. Get API key from Google

Refer to Google Documents for getting API key.
https://support.google.com/cloud/answer/6158862

### 2. Clone this repository

Clone this repository on an externally accessible server.

```
$ git clone typetalk-google-translate
```

### 3. Run application

Set Google API key you created.

```
$ go get all
$ export $GOOGLE_API_KEY=????????????????
$ go run app.go
```

### 4. Create Typetalk bot

Create Typetalk bot in a topic you want to translate messages.
See more detail Typetalk document.
https://developer.nulab-inc.com/docs/typetalk/#webhook

Webhook URL should be you are publishing host name, and port is 12345.

e.g. https://example.com:12345/

We strongly recommend you should use reverse proxy to not use the port for secure if you want to use in real.

### 5. Post a message not in English

Try to post a message as "こんにちは".
You can see the bot reply a message as "Hello".
