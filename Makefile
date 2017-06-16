all: leaf.pdf

leaf.pdf: leaf.dot Makefile
	dot -Kneato -n2 -Tpdf -o leaf.pdf leaf.dot

#leaf.dot: leaf
#	./leaf 

leaf: leaf.go
	./build.sh
