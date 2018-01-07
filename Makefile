all: leaf.gif

leaf.gif: leaf.py Makefile
	./leaf.py leaf.gif

leaf: leaf.py
	./build.sh

clean:
	-@rm *.dot
	-@rm *.scad
	-@rm *.png
	-@rm *.gif
	-@rm *.pyc

