# datagen
database data generator

    - A command line utility to generate data for a given table in singlestore database in accordance with constraints and limits configured via a json/yaml file.


```
$ ./datagen.exe -h


*********************************************************************
**************          DATAGENERATOR           *********************
*********************************************************************



A cmd line utility to generate mock data for database table

Usage:
  datagen [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  genconf     Generate default config for a given table in database
  gendata     Command to generate table as configured
  help        Help about any command

Flags:
      --conname string   Database Connection Name in datagen.json
      --debug            enable debug logs
  -h, --help             help for datagen
      --tblname string   Table name to use for the action (generate config | generate data)

Use "datagen [command] --help" for more information about a command.
```

### How to build

Pre-requisites:

Go 1.21 installed.

Steps to use as docker image
---

 1)  Build Docker Image

    -   git clone git@github.com:jkanuri/datagenerator.git
    -   cd datagen
    -   docker build -t datagen:1 .

2) Prepare application config

    - mkdir config output
    - cp datagen_template.json ./config/datagen.json
    - edit ./datagen.json file and update parameters as needed.
    

3) Mount config file as volume and run docker container

    docker run --rm --name datagen -v ./config:/etc/datagen -v ./output:/app/output <subcmd> 
    
    Examples

    * generate configuration 

    ```
     docker run --rm --name datagen -v ./config:/etc/datagen -v ./output:/app/output genconf --conname mysinglestore --tblname employee --out employee.json
    ```
    Above command reads database table structure and outputs table information as json. 
    
    This file is needed by gendata command to generate data as per table structure
    

    *   generate data

    ```
     docker run --rm --name datagen -v ./config:/etc/datagen -v ./output:/app/output gendata --conname mysinglestore --tblname employee --configfile employee.json --inline --rcount 100
    ```

     Above generates 100 records and inserts them into database.


Steps to use as executable
---

 1)  Build 

    -   git clone git@github.com:jkanuri/datagenerator.git
    -   cd datagen
    -   go mod tidy
    -   go build

2) Prepare application config

    - mkdir output $HOME/.datagen
    - cp datagen_template.json $HOME/.datagen/datagen.json
    - edit ./config/datagen.json file and update parameters as needed.
    

3) Generate configuration

    ```
    ./datagen.exe genconf -h


    *********************************************************************
    **************          DATAGENERATOR           *********************
    *********************************************************************



    Generate default config for a given table in database either in JSON or YAML format

    Usage:
    datagen genconf [flags]

    Flags:
    -h, --help         help for genconf
        --out string   output file name

    Global Flags:
        --conname string   Database Connection Name in datagen.json
        --debug            enable debug logs
        --tblname string   Table name to use for the action (generate config | generate data)

    ```
4) Generate Data

    ```
    ./datagen.exe genconf -h


    *********************************************************************
    **************          DATAGENERATOR           *********************
    *********************************************************************



    Generate default config for a given table in database either in JSON or YAML format

    Usage:
    datagen genconf [flags]

    Flags:
    -h, --help         help for genconf
        --out string   output file name

    Global Flags:
        --conname string   Database Connection Name in datagen.json
        --debug            enable debug logs
        --tblname string   Table name to use for the action (generate config | generate data)

    ```


Examples
---
* generate configuration 

```
./datagen.exe genconf --conname mysinglestore --tblname employee --out employee.json
```
Above command reads database table structure and outputs table information as json. 

This file is needed by gendata command to generate data as per table structure


*   generate data

```
    docker run --rm --name datagen -v ./config:/etc/datagen -v ./output:/app/output gendata --conname mysinglestore --tblname employee --configfile employee.json --inline --rcount 100
```

Above generates 100 records and inserts them into database.


Source Databases Supported 
---

| Name | Support |
|---|---|
|   Mysql   | [ x ] |
| Postgresql | [ x ] |
| Oracle | [ ] |

Note: Target Database will be singlestore.( Tool designed to support data generation into single store)

Documentation
---

| Feature | Link |
|---|---|
|Generators| [link](doc/basic_generators.md) |
|Oneof | [link](doc/oneof.md) |
|Sequences| [link](doc/sequences.md) |
|Geography| [link](doc/geo.md) |
|Configuration| [link](doc/config.md) |
