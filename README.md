# JSONtoXML
The program that converts XML to JSON and vice versa and compares its content

## Description

This project was developed as a part of my studying at School 21. It consists of 3 parts:

### ex00 
Task: Parse the file in JSON format and print its content in XML format to stdout. 

The program is runnable like this: 

`~$ ./readDB --f original_database.xml`

### ex01 
Task: Compare the contents of 2 files (XML or JSON), the differences may be:
1) New cake is added or old one removed
2) Cooking time is different for the same cake
3) New ingredient is added or removed for the same cake. *Important:* the order of ingredients doesn't matter. Only the names are.
4) The count of units for the same ingredient has changed
5) The unit itself for measuring the ingredient has changed
6) Ingredient unit is missing or added

The program is runnable like this: 

`~$ ./compareDB --old original_database.xml --new stolen_database.json`

### ex02 
Compare the contents of 2 .txt files. The files can be huge, so both of them may not fit into RAM on the same time.

The program is runnable like this:

`~$ ./compareFS --old snapshots/snapshot1.txt --new snapshots/snapshot2.txt`

## Important notes

- Makefile is provided to start the program conveniently;
- Functions that perform program logic are covered by unit tests.


