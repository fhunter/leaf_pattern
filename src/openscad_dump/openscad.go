package openscad_dump

import (
    "fmt"
    "os"
    "math/cmplx"
    "math"
    )

var frame int = 0

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func makeCoord(point complex128) string {
	str:= fmt.Sprintf("[%.4f,%.4f]",real(point),imag(point))
	return str
}


func Dumpall(growpoints []complex128, veinNodes []complex128, tree map[int] []int, influence []int, weights []float64) {
    dumpscad_str(growpoints, veinNodes, tree, weights, "")
    str:= fmt.Sprintf("%06d", frame)
    frame++
    dumpscad_str(growpoints, veinNodes, tree, weights, str)
}

func dumpscad_str(growpoints []complex128, veinNodes []complex128, tree map[int] []int, weights []float64, postfix string) {
    f,err := os.Create("./data"+postfix+".scad")
    check(err)
    defer f.Close()
    for _, t:= range growpoints {
        fmt.Fprint(f,"growpoint(p1=",makeCoord(t),");")
        fmt.Fprintln(f,"")
    }
    //Print nodes
    for i, t:= range veinNodes {
        fmt.Fprint(f,"node(p1=",makeCoord(t),",width=",(weights[i]),",ht=1);")
        fmt.Fprintln(f,"")
    }
    fmt.Fprintln(f," ")
    fmt.Fprintln(f," ")
    //Print links
    for i, t:= range tree {
        for _, k:= range t {
            var temp []complex128;
            angle:= cmplx.Phase(veinNodes[k]-veinNodes[i])
            angle+= math.Pi/2
            vector1:= cmplx.Rect(weights[i]/2,angle)
            vector2:= cmplx.Rect(weights[k]/2,angle)
            p1_1:=veinNodes[i]+vector1
            p1_2:=veinNodes[i]-vector1
            p2_1:=veinNodes[k]+vector2
            p2_2:=veinNodes[k]-vector2
            temp = append(temp,p1_1,p1_2,p2_2,p2_1)
            fmt.Fprint(f,"branch(p1=[")
            flag:=false
            for _,point:= range temp {
                if(flag) {
                    fmt.Fprint(f,",")
                }else{
                    flag=true
                }
                fmt.Fprint(f,makeCoord(point));
            }
            fmt.Fprint(f,"],ht=1);")
            fmt.Fprintln(f," ")
        }
    }
    fmt.Fprintln(f," ")
    fmt.Fprintln(f," ")
}

