.PHONY: run

HELP_FUN = \
		%help; while(<>){push@{$$help{$$2//'options'}},[$$1,$$3] \
		if/^([\w-_]+)\s*\s*:.*\#\#(?:@(\w+))?\s(.*)$$/}; \
		print"$$_:\n", map"  $$_->[0]".(" "x(20-length($$_->[0])))."$$_->[1]\n",\
		@{$$help{$$_}},"\n" for keys %help; \

run: #@Ex just run it
	go test ./...

