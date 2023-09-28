### OneOf generator

This is one of generators that can randomly pick one element from list of provided values or pick one from a table column.



Examples
---

1) Reference  column from another table.

    we should also specify connection to use, this way we can reference columns in another database too.

    ```
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
                        "column":"name",
                        "conname":"pgone"
                    }
                },
                "sequence": false,
                "cardinality": 0
            }
    ```

2) List of values option


    ```
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
                        
                    },
                    "items": [ "accounts", "billing", "quality", "manuf", "delivery", "sales"]
                },
                "sequence": false,
                "cardinality": 0
            }
    ```