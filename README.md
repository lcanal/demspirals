# Demspirals
Simple player stat exploration tool for Fantasy Football use.

## Setup
### Build

You'll need the Go build environment from Go's website:

* https://golang.org/dl/

And NodeJS 7+:

* https://nodejs.org/en/
 
#### Dependencies 

The dependencies call both be managed in Go and NodeJS via their respective official* dependency management tools. As of this writing build is still manual, build tool will come later (makefile?). 
* Go's isn't official as of me writing this yet, but it's planned to be.

Install Go's dep manager:

    $ go get -u github.com/golang/dep/cmd/dep
    
And install necessary depenencies for both Go and NodeJS. Make sure dep is in your path:

    $ dep ensure
    $ cd client
    $ npm install

Build the binary and JS bundles:

    $ go build
    $ cd client
    $ npm run build

Note: If you're building on windows, you will get errors about gcc not being in your path. For answers, look [here](https://github.com/mattn/go-sqlite3/issues/212#issuecomment-273531789)
but essentially:

    1. Download and install "tdm64-gcc-5.1.0-2.exe" from http://tdm-gcc.tdragon.net/download.
    2. Go to Program Files and click on "MinGW Command Prompt". This will start a DOS prompt with the correct environment for using MinGW with GCC.
    3. Within this DOS prompt window, navigate to your GOPATH. For example, I went to C:\go-apps.
    4. Enter the following commands: go get -u github.com/mattn/go-sqlite3. Then enter go install github.com/mattn/go-sqlite3
    5. You are done! go-sqlite3 is now installed and ready for use.
#### Build Options

You can enable cross compilation via the GOOS environment varibable:
	
	$ GOOS=windows go build

Produces a *opseb.exe* binary file.
Full options:

| GOOS Value         | Notes                                                                                            |
|:----------------------|:-------------------------------------------------------------------------------------------------------|
| `windows`             | Name of the application, will be used to construct the node service context name.    
| `darwin`         | Instance of the application (dev/tst/stg). This will be used to construct the node service context name. |
| `linux`           | Absolute path on build server where the built artifact files are stored |
    

### Run-time Dependencies
In order to run this app you'll need:
 * A MySQL Backend
 * A [Stattleship API Key](https://api.stattleship.com/)
 
 MySQL will just need a database location. It takes care of creating its own schemas and table definitions.
### Settings

By default demspirals tries to look in its current directory and then in the config/ directory for a setting file named **settings**. Settings can be in YAML,TOML or JSON format. A sample json formatted file is provided in **config/settings.json**
