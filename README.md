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
 * MySQL Backend
 * [Stattleship API Key](https://api.stattleship.com/)
 
### Settings

By default demspirals tries to look in its current directory and then in the config/ directory for a setting file named **settings**. Settings can be in YAML,TOML or JSON format. A sample json formatted file is provided in **config/settings.json**
