package openscad_dump

import (
    "fmt"
    "os"
    )

var frame int = 0

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func Dumpall(growpoints []complex128, veinNodes []complex128, tree map[int] []int, influence []int, weights []float64) {
    dumpscad_str(veinNodes, tree, weights, "")
    str:= fmt.Sprintf("%06d", frame)
    frame++
    dumpscad_str(veinNodes, tree, weights, str)
}

func dumpscad_str(veinNodes []complex128, tree map[int] []int, weights []float64, postfix string) {
    f,err := os.Create("./data"+postfix+".scad")
    check(err)
    defer f.Close()
    //Print nodes
    for i, t:= range veinNodes {
        fmt.Fprint(f,"node(p1=[",real(t),",",imag(t),"],width=",1,",ht=",(1+weights[i]),");")
        fmt.Fprintln(f,"")
    }
    fmt.Fprintln(f," ")
    fmt.Fprintln(f," ")
    //Print links
    /*
    for i, t:= range tree {
        for _, k:= range t {
            fmt.Fprint(f,"branch(p1=[",real(veinNodes[i]),",",imag(veinNodes[i]),"],p2=[",real(veinNodes[k]),",",imag(veinNodes[k]),"],width1=",weights[i],",width2=",weights[k],");")
            fmt.Fprintln(f,"")
        }
    }
    */
    fmt.Fprintln(f," ")
    fmt.Fprintln(f," ")
}

