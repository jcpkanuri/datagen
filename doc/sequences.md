### Sequences

- Sequences as name suggests is about sequences. 
- Column element "sequence" = true, enables sequence.
- Sequences configuration has below elements

    ```
    SequenceConfig {
        Start int
        Step  int
        Max   int
        Cycle bool
    }
    ```

Sequences can be declared at column level exclusively.
So a table can have multiple sequences.


example

sample\seq_test.json
```
{
    "tableName": "seq_test",
    "columns": [
        {
            "name": "id",
            "type": "int(11)",
            "nullsAllowed": "NO",
            "key": "PRI",
            "default": "",
            "extra": "",
            "generator": {
                "foreignKey": {}
            },
            "sequence": true,
            "cardinality": 0,
            "seqConfig": {
                "Start": 100,
                "Step": 3,
                "Max": 120,
                "Cycle": false
            }
        },
        {
            "name": "seq1",
            "type": "int(11)",
            "nullsAllowed": "YES",
            "key": "",
            "default": "",
            "extra": "",
            "generator": {
                "foreignKey": {}
            },
            "sequence": true,
            "cardinality": 0,
            "seqConfig": {
                "Start": 10,
                "Step": 1,
                "Max": 20,
                "Cycle": true
            }
        },
        {
            "name": "seq2",
            "type": "int(11)",
            "nullsAllowed": "YES",
            "key": "",
            "default": "",
            "extra": "",
            "generator": {
                "foreignKey": {}
            },
            "sequence": true,
            "cardinality": 0,
            "seqConfig": {
                "Start": 20,
                "Step": 7,
                "Max": 999,
                "Cycle": true
            }
        }
    ],
    "rows": null
}
```


You may run below cmd to generate 100 records using this config.

./datagen.exe gendata --conname mysinglestore --tblname seq_test --rcount 100 --configfile ./sample/seq_test.json --inline

