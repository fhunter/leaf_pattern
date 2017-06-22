package graphviz_dump

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
    dumpall_str(growpoints, veinNodes, tree, influence, weights, "") //This is working dump
    str:= fmt.Sprintf("%06d", frame)
    frame++
    dumpall_str(growpoints, veinNodes, tree, influence, weights, str) //This is frame dump
}

func dumpall_str(growpoints []complex128, veinNodes []complex128, tree map[int] []int, influence []int, weights []float64, postfix string) {
    f,err := os.Create("./leaf"+postfix+".dot")
    check(err)
    defer f.Close()
    fmt.Fprintln(f,"graph T {")
    // Вывод точек роста
    fmt.Fprintln(f,"node [color=\"green\"]")
    for i, t := range growpoints {
        //Draw position - xxx [ label = i, pos = "0,0!" ]
        fmt.Fprint(f,"grownode",i," [label=\"",i,"\", pos=\"",real(t)*32,",",imag(t)*32,"!\" ] ")
        fmt.Fprintln(f,"")
    }

    fmt.Fprintln(f,"node [shape=none,color=\"red\", fillcolor=\"transparent\", bgcolor=\"transparent\"]")
    //Print veinNodes
    for i, t := range veinNodes {
//        fmt.Fprint(f,"veinNode",i," [label=\"",i,"\", pos=\"",real(t)*32,",",imag(t)*32,"!\" ] ")
        fmt.Fprint(f,"veinNode",i," [label=\"\", pos=\"",real(t)*32,",",imag(t)*32,"!\" ] ")
        fmt.Fprintln(f,"")
    }

    //Print links
    fmt.Fprintln(f,"edge [tailclip=false,headclip=false,color=\"red\"]")
    for i, t:= range tree {
        for _, k:= range t {
            fmt.Fprint(f,"veinNode",i,"--veinNode",k,"[penwidth=",int(weights[i]*12+0.5),"]")
            fmt.Fprintln(f,"")
        }
    }
    fmt.Fprintln(f,"")

    // Print influence
    fmt.Fprintln(f,"edge [tailclip=true,headclip=false,color=\"blue\",style=\"dotted\"]")
    for i, t:= range influence {
        if t<len(veinNodes) {
            fmt.Fprint(f,"grownode",i,"--veinNode",t)
            fmt.Fprintln(f,"")
        }
    }
    fmt.Fprintln(f,"}")
    fmt.Fprintln(f," ")
}
