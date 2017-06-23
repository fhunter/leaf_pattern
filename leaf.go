package main

import (
    "fmt"
    "math/rand"
    "math/cmplx"
    "math"
    "time"
//    "tree"
    "graphviz_dump"
    "openscad_dump"
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
        weights:= calc_weights(tree,len(veinNodes))
        graphviz_dump.Dumpall(growpoints, veinNodes, tree, influence, weights)
        openscad_dump.Dumpall(growpoints, veinNodes, tree, influence, weights)
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
    weights:= calc_weights(tree,len(veinNodes))
    graphviz_dump.Dumpall(growpoints, veinNodes, tree, make([]int,0), weights)
    openscad_dump.Dumpall(growpoints, veinNodes, tree, make([]int,0), weights)


}
