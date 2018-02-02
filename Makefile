all: leaf.gif

leaf.gif: leaf.py Makefile datatypes.py
	./leaf.py leaf.gif

clean:
	-@rm *.dot
	-@rm *.scad
	-@rm *.png
	-@rm *.gif
	-@rm *.pyc

