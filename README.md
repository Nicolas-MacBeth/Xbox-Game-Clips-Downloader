# Xbox-Game-Clips-Downloader
A quickly-built utility for easily downloading Xbox game DVR clips hosted on the Xbox network.

## How to run
Either download the built executable for your platform, or run the program yourself. You'll need `go` installed. Clone this repo and build the executable using `go build ./`. Then, run it!

You'll need a valid authentication token (API key) from [xapi.us](https://xapi.us/). Create an account if you don't currently have one, connect your Microsoft account and copy your token. Input it into the program when you are asked to do so.

On MacOS and Linux, you'll probable need to run `chmod +x ./Xbox-Game-Clips-Downloader` to allow running the executable.

## Future
Reverse engineer the private Xbox REST API to have more functionality like deleting gameclips.
