package main

import (
    "fmt"
    "math/rand"
    "math"
)

type Point2D struct {
    x float64
    y float64
}

func distance(p1 Point2D, p2 Point2D) float64 {
    distance := math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
    return distance
}

func Norm(p Point2D) Point2D {
    dst:= distance(Point2D{0,0},p)
    if dst == 0 { 
        return p
    }
    p.x/=dst
    p.y/=dst
    return p
}

func findClosest(p Point2D, vec []Point2D) int {
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


const pointnum int = 40
const maxveinpoints int = 40000000
const deathdistance float64 = 10
const growthSpeed float64 = 0.5

func main() {
    fmt.Println("digraph T {")
    defer fmt.Println("}")
    //Заполнение точками. Квадрат 100x100
    growpoints := make([]Point2D, pointnum, pointnum)
    for i:=0; i< pointnum; i++ {
        //Make points here
        growpoints[i].x = rand.Float64()*100.0
        growpoints[i].y = rand.Float64()*100.0
    }

    // Вывод точек роста
    for i, t := range growpoints {
        //Draw position - xxx [ label = i, pos = "0,0!" ]
        fmt.Print("grownode",i," [label=\"",i,"\", color=\"green\", pos=\"",t.x,",",t.y,"!\" ] ")
        fmt.Println("")
    }

    veinNodes := make([]Point2D, 0, maxveinpoints)

    tree := make(map[int] []int)

    veinNodes = append(veinNodes,Point2D{50,0})

    tree[0] = make([]int,0, maxveinpoints)
    //tree[0] = append(tree[0],0)

    for (len(growpoints) > 0) {
        //Make lists of influence
        influence := make([]int, len(growpoints), len(growpoints)) //each growpoint is an influence for specific vein point
        for i, _:= range influence {
            //Go over all influence points and gather distances. fill the closest
            influence[i] = findClosest(growpoints[i],veinNodes)
        }
        fmt.Println("#",influence)
        //Calculate growth vectors
        {
            newNodes := make([]Point2D, 0, maxveinpoints)
            for i, _:= range veinNodes {
                p:= Point2D{0,0} //Initial vector
                needAdd := false
                for j, k:= range influence {
                    //Summ vectors
                    if k == i {
                        pnt := Point2D{growpoints[j].x - veinNodes[i].x, growpoints[j].y - veinNodes[i].y}
                        pnt = Norm(pnt)
                        p.x += pnt.x
                        p.y += pnt.y
                        needAdd = true
                    }
                }
                if needAdd {
                    p = Norm(p)
                    p.x = p.x * growthSpeed
                    p.y = p.y * growthSpeed
                    fmt.Println("#",p)
                    p.x = veinNodes[i].x + p.x
                    p.y = veinNodes[i].y + p.y
                    newNodes = append(newNodes,p)
                    tree[i] = append(tree[i], len(veinNodes)-1 + len(newNodes))
                }
            }
            for _, t:= range newNodes {
                veinNodes = append(veinNodes, t)
            }

        }
        newGrowPoints:= make([]Point2D,0,len(growpoints))
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
        fmt.Println("# growpoints",len(growpoints))
    }

    //Print veinNodes

    for i, t := range veinNodes {
        fmt.Print("veinNode",i," [label=\"",i,"\", color=\"red\", pos=\"",t.x,",",t.y,"!\" ] ")
        fmt.Println("")
    }

    //Print links

    for i, t:= range tree {
        for _, k:= range t {
            fmt.Print("veinNode",i,"->veinNode",k,"[ color=\"red\" ]")
            fmt.Println("")
        }
    }

}
