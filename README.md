![License](https://img.shields.io/github/license/bradleyGamiMarques/PersonaGrimoire?style=plastic)
![Issues](https://img.shields.io/github/issues/bradleyGamiMarques/PersonaGrimoire?style=plastic)
![Go](https://img.shields.io/github/go-mod/go-version/bradleyGamiMarques/PersonaGrimoire?style=plastic)
![Code Size](https://img.shields.io/github/languages/code-size/bradleyGamiMarques/PersonaGrimoire?style=plastic)
# Persona Grimoire

## Project Description
* What does this application do?
  * This application provides a RESTful API for Create, Read, Update, and Delete operations on Personas from the Shin Megami Tensei Persona spin-off series of video games.

* What is a Persona?
  * According to the [Megami Tensei Wiki](https://megamitensei.fandom.com/wiki/Persona_(concept)) A Persona is a manifestation of a Persona user's personality in the Persona series.

  * During gameplay these Personas are what the protagonists of the series use to fight during the turn-based combat sections.

* Why use the technologies in this project?
  ### Go
   * Go was selected as it is a popular language for building fast server-side applications. I am currently using Go as the language of choice for my part-time work at [Streemtech](https://github.com/streemtech).

      Writing the project in Go gives me additional practice for the work I am doing at Streemtech.
      
      
  ### Open API
   * Open API was selected because it provides a standard interface for documentation and code-generation.
  
  ### Echo
   * High performance web framework for Go. Minimal configuration required to get a server ready.

  ### PostgreSQL
  * Currently using this at Streemtech and its robust feature set supports the data types I wish to use.

  ### GORM
  * Another case of using this at Streemtech. It allows me to programatically map my objects to tables thereby reducing the mental load for creating and managing my databases.

* Why build this project?
  * I already have a project called [PersonaCompendium](https://github.com/bradleyGamiMarques/PersonaCompendium).
  * I decided to revisit this idea in Go because I want to have a project to show to potential employers who are looking for Go developers.
  * The original PersonaCompendium app was not written for fun.
  
    I was tasked with writing a MERN stack application in addition to my duties at Walmart as part of a PIP. This was a decidedly low point in my career.
    
    Revisiting this concept for fun will help motivate me to complete it. Discipline will allow me to follow through.

## Installation
### Docker - Coming  Soon! 
  Will be available at bradleygamimarques/personagrimoire
### Windows
 1. Clone down the project to your local machine.
 2. Install Make via [chocolatey](https://chocolatey.org/install)


 ![Chocolatey command](https://www.technewstoday.com/wp-content/uploads/2021/11/install-choco-make.webp)
 
 3. From the command prompt/powershell in the project folder run `make all`.
 
    This command will generate the API types, build the project, and start the server on PORT 5000.
    
    If you want to take it step by step here are the other commands that are run when `make all` is used.

    `make gen` - Generates the API Types.

    `make build` - Builds the executable.

    `make run` - Runs the executable. Starts the server on PORT 5000.
    
4. In another terminal run docker-compose -f docker-compose.yaml up.

   This will pull the Postgres image down and start the container.
   
## Usage
1. Use a HTTP client, such as Postman, to make requests to the API.
2. Routes are found in openapi.yaml
3. Make a request `http://localhost:5000/grimoire/v1/p5/arcana/name/Fool`
4. Example response:


![example](https://user-images.githubusercontent.com/18665288/192122219-60caa118-8380-472f-a9cf-f9c854667223.png)


## Contributing to the project
Notice something? Say something by [filing an issue](https://github.com/bradleyGamiMarques/PersonaGrimoire/issues/new).

When creating a branch that references a particular issue please follow the format {username}-issue#{issue number}.

Where possible, prefer rebase over squash even when squash is possible.
