### Generators


In table config json, we may configure below generators

| Name | Description |
|--|--|
| string|generates random string of given length |
|firstname| generates fake first name |
|lastname | generates fake last name |
|lexify| generates string by replacing '?' in exp provided |
|numerify | generates number by replacing '#' char with random integer |
| bothify | does both lexify and numerify e.g. '??? - ###' gives 'ACV - 712' |
| oneof | as name suggests picks a random element from provided list or foreign column reference |

Example configuration

```
{
    "tableName": "employee",
    "columns": [
        {
            "name": "id",
            "type": "bigint(11)",
            "nullsAllowed": "NO",
            "key": "UNI",
            "default": "",
            "extra": "auto_increment",
            "generator": {
                "foreignKey": {}
            },
            "sequence": false
        },
        {
            "name": "ename",
            "type": "varchar(100)",
            "nullsAllowed": "YES",
            "key": "",
            "default": "",
            "extra": "",
            "generator": {
                "name": "firstname",
                "foreignKey": {}
            },
            "sequence": false
        },
        {
            "name": "dept",
            "type": "varchar(10)",
            "nullsAllowed": "YES",
            "key": "",
            "default": "",
            "extra": "",
            "generator": {
                "name":"oneof",
                "foreignKey": {
                    "table":"department",
                    "column":"name"
                }
            },
            "sequence": false
        }
    ],
    "rows": null
}

```