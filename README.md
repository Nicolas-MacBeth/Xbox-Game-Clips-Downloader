# Xbox-Game-Clips-Downloader
A quickly-built utility for easily bulk-downloading Xbox game DVR clips hosted on the Xbox network.

## How to run
Either download the built executable for your platform through GitHub releases, or run the program yourself. You'll need `go` installed. Clone this repo and build the executable using `go build ./`. Then, run it!

You'll need a valid authentication token (API key) from [xapi.us](https://xapi.us/). Create an account if you don't currently have one, connect your Microsoft account and copy your token. Input it into the program when you are asked to do so.

On MacOS and Linux, you'll probably need to run `chmod +x ./path-to-binary` to allow running the executable.

## Motivation
Both the official Xbox app and third-party Xbox clips downloaders are missing the option to bulk-download all a user's clips and screenshots hosted on the Xbox network. This utility does just that.

## Future
Reverse engineer the **private** Xbox REST API to have more functionality like deleting gameclips.
