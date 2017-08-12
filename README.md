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

    $ npm install

To start a development instance with live reloading on the frontend side:

    $ npm start

To get ready to package a build for distribution, you can build the binary and JS bundles:

    $ npm run build

Package the build:
    
    $ npm run package

You should now have a fully usable application in the 'package' directory.
#### Build Options

You can enable cross compilation via the GOOS environment varibable:
	
	$ GOOS=windows go build

Produces a *opseb.exe* binary file.
Full options:

| GOOS Value         | Notes                                                                                            |
|:----------------------|:-------------------------------------------------------------------------------------------------------|
| `windows`             | Name of the application, will be used to construct the node service context name.    
| `darwin`              | Instance of the application (dev/tst/stg). This will be used to construct the node service context name. |
| `linux`               | Absolute path on build server where the built artifact files are stored |
    

### Run-time Dependencies
In order to run this app you'll need:
 * A MySQL Backend
 * A [MySportsFeed API Key](https://www.mysportsfeeds.com/data-feeds/)
 
 Database portion will just need a database location. It takes care of creating its own schemas and table definitions.
### Settings

By default demspirals tries to look in its current directory and then in the config/ directory for a setting file named **settings**. Settings can be in YAML,TOML or JSON format. A sample json formatted file is provided in **config/settings.json**


## Running
### Load Options

Loads are *currently* done in a "load once, load manually" fashion. Only available as startup flags. Add in a -h to see what the most up to date flag options are:

	Usage of ./backend:
	  -doloads
        Run initial loads for loading teams, players, stats.
	  -droptables
        Drop tables. Must be set along with doloads to run.