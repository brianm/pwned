.PHONY: clean full-clean bloom

pwned: .build/bin/rice
	go build
	$(PWD)/.rice/bin/rice append --exec pwned

.build:
	mkdir .build

.build/bin/rice: .build	
	GOPATH=$(PWD)/.build go get github.com/GeertJohan/go.rice/rice

clean:
	rm -f pwned

full-clean: clean
	rm -rf .build

pwned-passwords-2.0.txt:
	@echo "please download pwned-passwords-2.0.txt from https://haveibeenpwned.com/Passwords"
	@echo "and place it in this directory."

bloom: pwned-passwords-2.0.txt .build/bin/bloom
	rm -rf data
	mkdir data
	cut -f 1 -d : ./pwned-passwords-2.0.txt | .build/bin/bloom cr -p 0.000001 -n 501636842 data/pwned.bloom	

.build/bin/bloom:
	GOPATH=$(PWD)/.build go get github.com/DCSO/bloom/bloom	
