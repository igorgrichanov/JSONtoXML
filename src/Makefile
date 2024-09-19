all: ex_00 ex_01 ex_02

TASK00=readDB
TASK01=compareDB
TASK02=compareFS
MAIN_GO=main.go

ex_00: ex_00_1 ex_00_2

ex_01: ex_01_1 ex_01_2 ex_01_3

ex_00_1: clean
	go build -o $(TASK00) $(MAIN_GO)
	./$(TASK00) -f original_database.xml

ex_00_2: clean
	go build -o $(TASK00) $(MAIN_GO)
	./$(TASK00) -f stolen_database.json

ex_01_1: clean
	go build -o $(TASK01) $(MAIN_GO)
	./$(TASK01) --old original_database.xml --new original_database_2.xml

ex_01_2: clean
	go build -o $(TASK01) $(MAIN_GO)
	./$(TASK01) --old original_database.xml --new stolen_database.json

ex_01_3: clean
	go build -o $(TASK01) $(MAIN_GO)
	./$(TASK01) --old stolen_database.json --new stolen_database.json

ex_01_4: clean
	go build -o $(TASK01) $(MAIN_GO)
	./$(TASK01) --old stolen_database.json --new original_database_2.xml

ex_02: clean
	go build -o $(TASK02) $(MAIN_GO)
	./$(TASK02) --old snapshots/snapshot1.txt --new snapshots/snapshot2.txt

format:
	gofmt -w .
	golines -w -m 80 .

install_godoc:
	go install -v golang.org/x/tools/cmd/godoc@latest

run_godoc:
	godoc -http=:6060

save_doc:
	wget -p -k http://localhost:6060/pkg/

zip_doc:
	zip -r doc.zip localhost:6060

clean:
	rm -f $(TASK00) $(TASK01) $(TASK02)