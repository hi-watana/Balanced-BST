avl:
	gcc -g avl_tree.c

splay:
	gcc -g splay_tree.c

run:
	./a.out

debug:
	lldb a.out

clean:
	rm -rfv a.out a.out.dSYM/a.out a.out.dSYM/
