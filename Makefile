all: leaf.pdf

leaf.pdf: leaf.dot Makefile
	dot -Kneato -n2 -Tpdf -o leaf.pdf leaf.dot

#leaf.dot: leaf
#	./leaf 

leaf: leaf.go src/openscad_dump/openscad.go src/graphviz_dump/graphviz.go src/tree/tree.go
	./build.sh

clean:
	-@rm *.dot
	-@rm *.scad

