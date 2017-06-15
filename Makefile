all: leaf.pdf

leaf.pdf: leaf.dot
	dot -Kfdp -n -Tpdf -o leaf.pdf leaf.dot

leaf.dot: leaf
	./leaf > leaf.dot

leaf: leaf.go
	./build.sh
