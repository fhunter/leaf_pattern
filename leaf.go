package main

import (
    "fmt"
    "math/rand"
    "math/cmplx"
    "os"
    "bufio"
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
        return 0
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

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func dumpall(growpoints []complex128, veinNodes []complex128, tree map[int] []int, influence []int) {
    f,err := os.Create("./leaf.dot")
    check(err)
    defer f.Close()
    fmt.Fprintln(f,"graph T {")
    // Вывод точек роста
    for i, t := range growpoints {
        //Draw position - xxx [ label = i, pos = "0,0!" ]
//        fmt.Fprint(f,"grownode",i," [scale=0.01, label=\"",i,"\", color=\"green\", pos=\"",real(t),",",imag(t),"!\" ] ")
        fmt.Fprint(f,"grownode",i," [label=\"",i,"\", color=\"green\", pos=\"",real(t)*32,",",imag(t)*32,"!\" ] ")
        fmt.Fprintln(f,"")
    }
    //Print veinNodes

    for i, t := range veinNodes {
//        fmt.Fprint(f,"veinNode",i," [scale=0.01, label=\"",i,"\", color=\"red\", pos=\"",real(t),",",imag(t),"!\" ] ")
        fmt.Fprint(f,"veinNode",i," [label=\"",i,"\", color=\"red\", pos=\"",real(t)*32,",",imag(t)*32,"!\" ] ")
        fmt.Fprintln(f,"")
    }

    //Print links

    for i, t:= range tree {
        for _, k:= range t {
            fmt.Fprint(f,"veinNode",i,"--veinNode",k,"[ color=\"red\" ]")
            fmt.Fprintln(f,"")
        }
    }
    fmt.Fprintln(f,"")

    // Print influence
    for i, t:= range influence {
	if t<len(veinNodes) {
		fmt.Fprint(f,"grownode",i,"--veinNode",t,"[style=\"dotted\",color=\"blue\"]")
	        fmt.Fprintln(f,"")
	}
    }
    fmt.Fprintln(f,"}")
}


const pointnum int = 10000
const maxveinpoints int = 40000
const deathdistance float64 = 3
const growthSpeed float64 = 0.9

func main() {
    //Заполнение точками. Квадрат 100x100
    rand.Seed(time.Now().Unix())
    growpoints := make([]complex128, pointnum, pointnum)
    for i:=0; i< pointnum; i++ {
        //Make points here
        growpoints[i] = complex(float64(rand.Intn(100)), float64(rand.Intn(100)))
    }

    veinNodes := make([]complex128, 0, maxveinpoints)

    tree := make(map[int] []int)

    veinNodes = append(veinNodes,50)

    tree[0] = make([]int,0, maxveinpoints)
    //tree[0] = append(tree[0],0)

    for (len(growpoints) > 0) {
        //Make lists of influence
        influence := make([]int, len(growpoints), len(growpoints)) //each growpoint is an influence for specific vein point
        for i, _:= range influence {
            //Go over all influence points and gather distances. fill the closest
            influence[i] = findClosest(growpoints[i],veinNodes)
        }
        fmt.Println("# influence",influence)
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
                    fmt.Println("# point",p)
                    newNodes = append(newNodes,p)
                    tree[i] = append(tree[i], len(veinNodes)-1 + len(newNodes))
                }
            }
            for _, t:= range newNodes {
                veinNodes = append(veinNodes, t)
            }

        }
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
        fmt.Println("# growpoints",growpoints)
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n') 
        growpoints = newGrowPoints
    }


}
