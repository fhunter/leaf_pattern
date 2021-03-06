package main

import (
    "fmt"
    "math/rand"
    "math/cmplx"
    "math"
    "os"
    "time"
)

func distance(p1 complex128, p2 complex128) float64 {
    return cmplx.Abs(p2-p1)
}

func Norm(p complex128) complex128 {
    return cmplx.Rect(1, cmplx.Phase(p))
}

func findClosest(p complex128, vec []complex128) int {
    if len(vec)==0 {
        return math.MaxInt32
    }
    minDistance:= distance(p,vec[0])
    minIndex:= 0
    for i, j:= range vec {
        if minDistance > distance(p, j) {
            minIndex = i
            minDistance = distance(p, j)
        }
    }
    return minIndex
}

func makeInitialGrowPoints(growpoints []complex128, maxx float64, maxy float64, pointnum int) []complex128 {
    for i:=0; i< pointnum; i++ {
        //Make points here
        growpoints[i] = complex(float64(rand.Intn(int(maxx))), float64(rand.Intn(int(maxy))))
    }
    return growpoints
}

func addGrowPoints(growpoints []complex128, maxx float64, maxy float64, density float64) []complex128 {
    numberOfPointsToAdd:=int(maxx*maxy*density)
    for i:=0;i<numberOfPointsToAdd;i++ {
        growpoints = append(growpoints,complex(float64(rand.Intn(int(maxx))), float64(rand.Intn(int(maxy)))))
    }
    return growpoints
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func weight(tree map[int] []int, node int) float64 {
    var ret float64 = 0.5
    if node > len(tree) {
        return ret
    }
    if len(tree[node]) == 0 {
        //This is a leaf
        ret = 1
    } else {
        ret = 0
        //This is a junction - 
        for _, k:= range tree[node] {
            tmp:= weight(tree, k)
            ret += tmp*tmp
        }
        ret= math.Sqrt(ret)
    }
    return ret
}

func calc_weights(tree map[int] []int,length int) []float64 {
    weights:= make([]float64, length + 1, length+1)
    for i, _:= range tree {
        weights[i]= weight(tree, i)
    }
    return weights
}

var frame int = 0

func dumpall(growpoints []complex128, veinNodes []complex128, tree map[int] []int, influence []int) {
//    dumpall_str(growpoints, veinNodes, tree, influence, "") //This is working dump
    dumpscad_str(veinNodes, tree, "")
    str:= fmt.Sprintf("%06d", frame)
    frame++
//    dumpall_str(growpoints, veinNodes, tree, influence, str) //This is frame dump
    dumpscad_str(veinNodes, tree, str)
}

func dumpall_str(growpoints []complex128, veinNodes []complex128, tree map[int] []int, influence []int, postfix string) {
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
    weights:= calc_weights(tree,len(veinNodes))
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

func dumpscad_str(veinNodes []complex128, tree map[int] []int, postfix string) {
    f,err := os.Create("./data"+postfix+".scad")
    check(err)
    defer f.Close()
    weights:= calc_weights(tree,len(veinNodes))
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


const pointnum int = 10000
const maxveinpoints int = 40000
const deathdistance float64 = 2
const growthSpeed float64 = 1
const addGrowthDensity float64 = 0.000

func main() {
    //Заполнение точками. Квадрат 100x100
    rand.Seed(time.Now().Unix())
    growpoints := make([]complex128, pointnum, pointnum)
    growpoints = makeInitialGrowPoints(growpoints, 100.0, 100.0, pointnum)

    veinNodes := make([]complex128, 0, maxveinpoints)

    tree := make(map[int] []int)

    veinNodes = append(veinNodes,50)

    tree[0] = make([]int,0, maxveinpoints)

    for (len(growpoints) > 0) {
        //Make lists of influence
        influence := make([]int, len(growpoints), len(growpoints)) //each growpoint is an influence for specific vein point
        for i, _:= range influence {
            //Go over all influence points and gather distances. fill the closest
            influence[i] = findClosest(growpoints[i],veinNodes)
        }
        dumpall(growpoints, veinNodes, tree, influence)
        //Calculate growth vectors
        {
            newNodes := make([]complex128, 0, maxveinpoints)
            for i, _:= range veinNodes {
                p:= 0+0i //Initial vector
                needAdd := false
                for j, k:= range influence {
                    //Summ vectors
                    if k == i {
                        pnt := growpoints[j] - veinNodes[i]
                        pnt = Norm(pnt)
                        p += pnt
                        needAdd = true
                    }
                }
                if needAdd {
                    p = Norm(p)
                    p = p * complex128(growthSpeed)
                    p = veinNodes[i] + p
                    newNodes = append(newNodes,p)
                    tree[i] = append(tree[i], len(veinNodes)-1 + len(newNodes))
                }
            }
            for _, t:= range newNodes {
                veinNodes = append(veinNodes, t)
            }

        }
        growpoints = addGrowPoints(growpoints,100.0,100.0,addGrowthDensity)
        newGrowPoints:= make([]complex128,0,len(growpoints))
        //Delete any growth points that are too close
        for _, p:= range growpoints {
            flag := true
            for _, t:= range veinNodes {
                if distance(p, t) < deathdistance {
                    flag = false
                    break
                }
            }
            if flag {
                newGrowPoints = append(newGrowPoints, p)
            }
        }
        fmt.Println("# growpoints",len(growpoints))
        growpoints = newGrowPoints
    }
    dumpall(growpoints, veinNodes, tree, make([]int,0))


}
