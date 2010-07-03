COMPILER=6g
LINKER=6l

goget: goget.6
	$(LINKER) -o $@ $<

goget.6: goget.go
	$(COMPILER) $<

clean:
	$(RM) -f *.6

