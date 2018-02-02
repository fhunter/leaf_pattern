all: leaf.gif

leaf.gif: main.py Makefile Leaf.py Tree.py Point.py
	./main.py leaf.gif

clean:
	-@rm *.dot
	-@rm *.scad
	-@rm *.png
	-@rm *.gif
	-@rm *.pyc

